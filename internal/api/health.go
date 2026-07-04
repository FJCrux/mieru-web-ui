// SPDX-License-Identifier: GPL-3.0-or-later

package api

import (
	"net"
	"net/http"
	"net/url"
	"strings"
)

// HandleHealth is a public liveness probe: 200 while the panel serves. It also
// reports the mita proxy state for convenience. Used by the Docker HEALTHCHECK.
func (s *Server) HandleHealth(w http.ResponseWriter, r *http.Request) {
	mitaStatus := "unknown"
	if st, err := s.Mita.Status(r.Context()); err == nil {
		mitaStatus = st
	} else {
		mitaStatus = "unreachable"
	}
	writeJSON(w, http.StatusOK, map[string]string{"status": "ok", "mita": mitaStatus})
}

// HostGuard rejects requests whose Host header does not match the configured
// panel URL. This blocks Host-header injection (share links are built from the
// panel URL) and stray access by IP. Loopback is always allowed so the health
// check and local tooling work. When panel_url is unset, all hosts are allowed.
func (s *Server) HostGuard(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Infra endpoints are scraped/probed by host or IP, not the panel URL.
		if r.URL.Path == "/healthz" || r.URL.Path == "/metrics" {
			next.ServeHTTP(w, r)
			return
		}
		if allowed := s.allowedHost(); allowed != "" {
			if !hostMatches(r.Host, allowed) && !isLoopbackHost(r.Host) {
				http.Error(w, "unexpected host", http.StatusMisdirectedRequest)
				return
			}
		}
		next.ServeHTTP(w, r)
	})
}

// allowedHost extracts the host[:port] from the panel_url setting.
func (s *Server) allowedHost() string {
	raw, err := s.Store.Setting("panel_url")
	if err != nil || raw == "" {
		return ""
	}
	u, err := url.Parse(raw)
	if err != nil {
		return ""
	}
	return u.Host
}

// hostMatches compares request Host to the allowed host, ignoring the port on
// the allowed side (proxies may terminate on 443 and forward on another port).
func hostMatches(reqHost, allowed string) bool {
	rh := hostname(reqHost)
	ah := hostname(allowed)
	return rh != "" && strings.EqualFold(rh, ah)
}

func hostname(h string) string {
	if host, _, err := net.SplitHostPort(h); err == nil {
		return host
	}
	return h
}

func isLoopbackHost(h string) bool {
	host := hostname(h)
	if host == "localhost" {
		return true
	}
	ip := net.ParseIP(host)
	return ip != nil && ip.IsLoopback()
}
