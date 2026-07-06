// SPDX-License-Identifier: GPL-3.0-or-later

package api

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"strings"

	"github.com/fjcrux/mieru-web-ui/internal/clashsub"
	"github.com/fjcrux/mieru-web-ui/internal/sharelink"
	"github.com/fjcrux/mieru-web-ui/internal/store"
)

type subscriptionStatus struct {
	Exists    bool   `json:"exists"`
	URL       string `json:"url,omitempty"`
	CreatedAt int64  `json:"createdAt,omitempty"`
}

type createSubscriptionRequest struct {
	Host         string `json:"host"`         // "" = follow public_host
	Multiplexing string `json:"multiplexing"` // "" = client default
}

type createSubscriptionResponse struct {
	URL       string `json:"url"`
	CreatedAt int64  `json:"createdAt"`
}

// handleGetSubscription reports whether the user has a subscription token and
// re-displays its URL (tokens are stored in plaintext for exactly this).
func (s *Server) handleGetSubscription(w http.ResponseWriter, r *http.Request) {
	data, err := s.Store.SubTokenForUser(r.PathValue("name"))
	if errors.Is(err, store.ErrNotFound) {
		writeJSON(w, http.StatusOK, subscriptionStatus{Exists: false})
		return
	} else if err != nil {
		writeErr(w, http.StatusInternalServerError, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, subscriptionStatus{
		Exists:    true,
		URL:       s.subscriptionURL(r, data.Token),
		CreatedAt: data.CreatedAt.Unix(),
	})
}

// handleCreateSubscription mints the user's permanent subscription token,
// rotating (invalidating) any previous one.
func (s *Server) handleCreateSubscription(w http.ResponseWriter, r *http.Request) {
	name := r.PathValue("name")
	var req createSubscriptionRequest
	if !readJSON(w, r, &req) {
		return
	}

	// Validate the subscription can actually be built before minting.
	if _, status, err := s.buildLinks(r.Context(), name, req.Host, req.Multiplexing); err != nil {
		writeErr(w, status, err.Error())
		return
	}

	token, err := randomToken()
	if err != nil {
		writeErr(w, http.StatusInternalServerError, "failed to generate token")
		return
	}
	if err := s.Store.UpsertSubToken(token, store.SubTokenData{
		Username:     name,
		Host:         req.Host,
		Multiplexing: req.Multiplexing,
	}); err != nil {
		writeErr(w, http.StatusInternalServerError, err.Error())
		return
	}
	data, err := s.Store.SubTokenForUser(name)
	if err != nil {
		writeErr(w, http.StatusInternalServerError, err.Error())
		return
	}
	writeJSON(w, http.StatusCreated, createSubscriptionResponse{
		URL:       s.subscriptionURL(r, token),
		CreatedAt: data.CreatedAt.Unix(),
	})
}

func (s *Server) handleDeleteSubscription(w http.ResponseWriter, r *http.Request) {
	if err := s.Store.DeleteSubTokenForUser(r.PathValue("name")); err != nil {
		writeErr(w, http.StatusInternalServerError, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, map[string]bool{"ok": true})
}

// HandlePublicSubscription serves a client subscription for the permanent
// token. No auth (the token is the credential). Invalid tokens, unknown
// formats, and unresolvable users all return the same 404 so the endpoint
// leaks nothing (no enumeration, no format oracle).
//
// Formats: default/?format=clash is a full Clash/Mihomo profile
// (?flavor=proxies strips it to a proxies-only document for proxy-provider);
// ?format=mierus and ?format=mieru return the native links as a
// base64-encoded line list (the v2rayN-style URI subscription convention).
// A future ?format=singbox can branch here without changing routing.
func (s *Server) HandlePublicSubscription(w http.ResponseWriter, r *http.Request) {
	notFound := func() { http.Error(w, "this link is invalid", http.StatusNotFound) }

	data, err := s.Store.SubToken(r.PathValue("token"))
	if err != nil {
		notFound()
		return
	}

	format := r.URL.Query().Get("format")
	flavor := r.URL.Query().Get("flavor")
	if format == "" {
		format = "clash"
	}
	if flavor != "" && (format != "clash" || flavor != "proxies") {
		notFound()
		return
	}

	password, host, cfg, _, err := s.resolveUserParams(r.Context(), data.Username, data.Host)
	if err != nil {
		notFound()
		return
	}

	var body []byte
	var contentType, filename string
	switch format {
	case "clash":
		p := clashsub.Params{
			Username:     data.Username,
			Password:     password,
			Host:         host,
			PortBindings: cfg.GetPortBindings(),
			Multiplexing: data.Multiplexing,
		}
		if flavor == "proxies" {
			body, err = clashsub.BuildProxies(p)
		} else {
			body, err = clashsub.BuildProfile(p)
		}
		contentType = "text/yaml; charset=utf-8"
		filename = "mieru-" + data.Username + ".yaml"
	case "mieru", "mierus":
		var links *sharelink.Links
		links, err = sharelink.Build(sharelink.Params{
			Username:     data.Username,
			Password:     password,
			Host:         host,
			PortBindings: cfg.GetPortBindings(),
			MTU:          cfg.GetMtu(),
			Multiplexing: data.Multiplexing,
		})
		if err == nil {
			raw := links.MieruURL
			if format == "mierus" {
				raw = strings.Join(links.MierusURLs, "\n")
			}
			body = []byte(base64.StdEncoding.EncodeToString([]byte(raw)))
		}
		contentType = "text/plain; charset=utf-8"
		filename = "mieru-" + data.Username + ".txt"
	default:
		notFound()
		return
	}
	if err != nil {
		notFound()
		return
	}

	w.Header().Set("Content-Type", contentType)
	w.Header().Set("Cache-Control", "no-store")
	w.Header().Set("Profile-Update-Interval", "24")
	w.Header().Set("Profile-Title", "mieru-"+data.Username)
	w.Header().Set("Content-Disposition", `attachment; filename="`+filename+`"`)
	if ui := s.subscriptionUserInfo(r.Context(), data.Username); ui != "" {
		w.Header().Set("Subscription-Userinfo", ui)
	}
	_, _ = w.Write(body)
}

// subscriptionUserInfo builds the subscription-userinfo header value from
// mita's per-user traffic counters and quotas. upload/download are cumulative
// bytes (mita persists them across restarts via its metrics dump); total is
// the tightest quota when quotas exist. No expire: mieru quotas are rolling
// windows, so a fixed expiry timestamp does not exist. Returns "" (header
// omitted) when mita is unreachable or the user is unknown.
func (s *Server) subscriptionUserInfo(ctx context.Context, name string) string {
	users, err := s.Mita.Users(ctx)
	if err != nil {
		return ""
	}
	for _, item := range users.GetItems() {
		if item.GetUser().GetName() != name {
			continue
		}
		var upload, download int64
		for _, m := range item.GetMetrics() {
			switch m.GetName() {
			case "UploadBytes":
				upload = m.GetValue()
			case "DownloadBytes":
				download = m.GetValue()
			}
		}
		info := fmt.Sprintf("upload=%d; download=%d", upload, download)
		var total int64
		for _, q := range item.GetUser().GetQuotas() {
			cap := int64(q.GetMegabytes()) * 1024 * 1024
			if total == 0 || cap < total {
				total = cap
			}
		}
		if total > 0 {
			info += fmt.Sprintf("; total=%d", total)
		}
		return info
	}
	return ""
}

// subscriptionURL builds the public subscription URL for a token. With a
// dedicated subscription port configured, the panel port in the base URL is
// swapped for it.
func (s *Server) subscriptionURL(r *http.Request, token string) string {
	base := s.panelBaseURL(r)
	if s.SubPort != "" {
		base = replaceURLPort(base, s.SubPort)
	}
	return base + s.SubPath + "/" + token
}

// replaceURLPort swaps the port of a scheme://host[:port] base URL.
func replaceURLPort(base, port string) string {
	u, err := url.Parse(base)
	if err != nil || u.Host == "" {
		return base
	}
	u.Host = net.JoinHostPort(u.Hostname(), port)
	return u.String()
}
