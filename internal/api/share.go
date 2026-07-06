// SPDX-License-Identifier: GPL-3.0-or-later

package api

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"net/http"
	"time"

	pb "github.com/enfein/mieru/v3/pkg/appctl/appctlpb"
	"github.com/fjcrux/mieru-web-ui/internal/sharelink"
	"github.com/fjcrux/mieru-web-ui/internal/store"
)

// maxShareTTL caps how long a generated link can live.
const maxShareTTL = 7 * 24 * time.Hour

// resolveUserParams fetches everything a link or subscription builder needs:
// the user's plaintext password, the public host (falling back to the
// public_host setting when host is empty), and the live server config. The
// int is an HTTP status to use on error (0 on success).
func (s *Server) resolveUserParams(ctx context.Context, name, host string) (password, resolvedHost string, cfg *pb.ServerConfig, status int, err error) {
	password, err = s.Store.UserSecret(name)
	if errors.Is(err, store.ErrNotFound) {
		return "", "", nil, http.StatusConflict, errors.New("no stored password for this user (created outside the panel) - reset the password to enable sharing")
	} else if err != nil {
		return "", "", nil, http.StatusInternalServerError, err
	}

	if host == "" {
		host, err = s.Store.Setting("public_host")
		if err != nil || host == "" {
			return "", "", nil, http.StatusConflict, errors.New("public host is not configured - set it in Settings")
		}
	}

	cfg, err = s.Mita.GetConfig(ctx)
	if err != nil {
		return "", "", nil, http.StatusBadGateway, err
	}
	return password, host, cfg, 0, nil
}

// buildLinks resolves a user's client config and share links. host falls back
// to the public_host setting when empty. The int is an HTTP status to use on
// error (0 on success).
func (s *Server) buildLinks(ctx context.Context, name, host, multiplexing string) (*sharelink.Links, int, error) {
	password, host, cfg, status, err := s.resolveUserParams(ctx, name, host)
	if err != nil {
		return nil, status, err
	}

	links, err := sharelink.Build(sharelink.Params{
		Username:     name,
		Password:     password,
		Host:         host,
		PortBindings: cfg.GetPortBindings(),
		MTU:          cfg.GetMtu(),
		Multiplexing: multiplexing,
	})
	if err != nil {
		return nil, http.StatusBadRequest, err
	}
	return links, 0, nil
}

// handleShare returns a user's links directly (admin-only, for the panel UI).
func (s *Server) handleShare(w http.ResponseWriter, r *http.Request) {
	links, status, err := s.buildLinks(r.Context(), r.PathValue("name"),
		r.URL.Query().Get("host"), r.URL.Query().Get("multiplexing"))
	if err != nil {
		writeErr(w, status, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, links)
}

type createShareTokenRequest struct {
	TTLMinutes   int    `json:"ttlMinutes"`
	Host         string `json:"host"`
	Multiplexing string `json:"multiplexing"`
}

type createShareTokenResponse struct {
	URL       string `json:"url"`
	ExpiresAt int64  `json:"expiresAt"`
}

// handleCreateShareToken mints an expiring public link for a user (admin-only).
func (s *Server) handleCreateShareToken(w http.ResponseWriter, r *http.Request) {
	name := r.PathValue("name")
	var req createShareTokenRequest
	if !readJSON(w, r, &req) {
		return
	}

	ttl := time.Duration(req.TTLMinutes) * time.Minute
	if ttl <= 0 {
		ttl = time.Hour
	}
	if ttl > maxShareTTL {
		ttl = maxShareTTL
	}

	host := req.Host
	if host == "" {
		host, _ = s.Store.Setting("public_host")
	}
	// Validate the link can actually be built before handing out a token.
	if _, status, err := s.buildLinks(r.Context(), name, host, req.Multiplexing); err != nil {
		writeErr(w, status, err.Error())
		return
	}

	token, err := randomToken()
	if err != nil {
		writeErr(w, http.StatusInternalServerError, "failed to generate token")
		return
	}
	expiresAt := time.Now().Add(ttl)
	if err := s.Store.CreateShareToken(hashShareToken(token), store.ShareTokenData{
		Username:     name,
		Host:         host,
		Multiplexing: req.Multiplexing,
	}, expiresAt); err != nil {
		writeErr(w, http.StatusInternalServerError, err.Error())
		return
	}

	writeJSON(w, http.StatusCreated, createShareTokenResponse{
		URL:       s.panelBaseURL(r) + s.SharePath + "/" + token,
		ExpiresAt: expiresAt.Unix(),
	})
}

// HandlePublicSharePage renders a self-contained share page for whoever holds
// the token. No auth (the token is the credential) and no admin SPA, so it can
// live under its own public prefix without revealing the panel base path.
// Invalid or expired tokens return 404.
func (s *Server) HandlePublicSharePage(w http.ResponseWriter, r *http.Request) {
	data, err := s.Store.ShareToken(hashShareToken(r.PathValue("token")))
	if err != nil {
		http.Error(w, "this link is invalid or has expired", http.StatusNotFound)
		return
	}
	links, status, err := s.buildLinks(r.Context(), data.Username, data.Host, data.Multiplexing)
	if err != nil {
		http.Error(w, "this link is no longer available", statusOr(status, http.StatusNotFound))
		return
	}
	renderSharePage(w, data.Username, links)
}

func statusOr(s, def int) int {
	if s == http.StatusConflict {
		return http.StatusNotFound
	}
	return def
}

func randomToken() (string, error) {
	b := make([]byte, 24)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return hex.EncodeToString(b), nil
}

func hashShareToken(token string) string {
	h := sha256.Sum256([]byte(token))
	return hex.EncodeToString(h[:])
}

// panelBaseURL returns the configured panel URL when set, else falls back to
// the request's own origin.
func (s *Server) panelBaseURL(r *http.Request) string {
	if u, err := s.Store.Setting("panel_url"); err == nil && u != "" {
		return u
	}
	return requestBaseURL(r)
}

// requestBaseURL reconstructs the panel's own origin, honouring a reverse
// proxy's X-Forwarded-Proto/Host when present.
func requestBaseURL(r *http.Request) string {
	scheme := "http"
	if r.TLS != nil {
		scheme = "https"
	}
	if xf := r.Header.Get("X-Forwarded-Proto"); xf != "" {
		scheme = xf
	}
	host := r.Host
	if xh := r.Header.Get("X-Forwarded-Host"); xh != "" {
		host = xh
	}
	return scheme + "://" + host
}
