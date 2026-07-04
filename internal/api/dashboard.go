// SPDX-License-Identifier: GPL-3.0-or-later

package api

import (
	"encoding/json"
	"net"
	"net/http"
	"time"
)

// isLoopback reports whether the request came in over the loopback
// interface (accessed via localhost / an SSH tunnel), where serving plain
// HTTP is acceptable.
func isLoopback(r *http.Request) bool {
	host := r.Host
	if h, _, err := net.SplitHostPort(r.Host); err == nil {
		host = h
	}
	if host == "localhost" {
		return true
	}
	ip := net.ParseIP(host)
	return ip != nil && ip.IsLoopback()
}

type dashboardResponse struct {
	MitaStatus     string          `json:"mitaStatus"`
	MitaVersion    string          `json:"mitaVersion"`
	Restarts       int             `json:"restarts"`
	MitaUptime     int64           `json:"mitaUptimeSeconds"`
	SessionCount   int             `json:"sessionCount"`
	UserCount      int             `json:"userCount"`
	Metrics        json.RawMessage `json:"metrics"`
	InsecureAccess bool            `json:"insecureAccess"`
	// Warnings are hardening hints shown at the top of the dashboard.
	Warnings []string `json:"warnings"`
}

// hardeningWarnings collects non-fatal hints about the panel's exposure.
func (s *Server) hardeningWarnings(r *http.Request) []string {
	var w []string
	if !s.TLSEnabled && !isLoopback(r) {
		w = append(w, "Served over plain HTTP on a non-loopback address. Use TLS or a reverse proxy, or tunnel to 127.0.0.1.")
	}
	if p, _ := s.Store.Setting("panel_url"); p == "" {
		w = append(w, "Panel URL is not set, so the Host header isn't validated. Set it in Settings.")
	}
	if bp, _ := s.Store.Setting("base_path"); bp == "" {
		w = append(w, "The panel is at the default root path. Set a secret base path in Settings to hide it from scanners.")
	}
	if s.PanelPort == "8686" {
		w = append(w, "The panel is on the default port 8686. Consider changing it or putting it behind a reverse proxy.")
	}
	if s.restartPending() {
		w = append(w, "Path settings changed but not applied yet. Restart the panel (Settings) to apply them.")
	}
	return w
}

func (s *Server) handleDashboard(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	resp := dashboardResponse{
		Metrics:        json.RawMessage("{}"),
		InsecureAccess: !s.TLSEnabled && !isLoopback(r),
		Warnings:       s.hardeningWarnings(r),
	}

	status, err := s.Mita.Status(ctx)
	if err != nil {
		// mita being down is a state the dashboard must render, not a 5xx.
		resp.MitaStatus = "UNREACHABLE"
		writeJSON(w, http.StatusOK, resp)
		return
	}
	resp.MitaStatus = status

	if v, err := s.Mita.Version(ctx); err == nil {
		resp.MitaVersion = v
	}
	if m, err := s.Mita.Metrics(ctx); err == nil && json.Valid([]byte(m)) {
		resp.Metrics = json.RawMessage(m)
	}
	if sessions, err := s.Mita.Sessions(ctx); err == nil {
		resp.SessionCount = len(sessions.GetItems())
	}
	if users, err := s.Mita.Users(ctx); err == nil {
		resp.UserCount = len(users.GetItems())
	}
	if s.Sup != nil {
		restarts, startedAt := s.Sup.Stats()
		resp.Restarts = restarts
		if !startedAt.IsZero() {
			resp.MitaUptime = int64(time.Since(startedAt).Seconds())
		}
	}
	writeJSON(w, http.StatusOK, resp)
}

type sessionInfo struct {
	ID         string `json:"id"`
	Protocol   string `json:"protocol"`
	LocalAddr  string `json:"localAddr"`
	RemoteAddr string `json:"remoteAddr"`
	State      string `json:"state"`
}

func (s *Server) handleSessions(w http.ResponseWriter, r *http.Request) {
	sessions, err := s.Mita.Sessions(r.Context())
	if err != nil {
		if proxyNotReady(err) {
			writeJSON(w, http.StatusOK, []sessionInfo{})
			return
		}
		writeMitaErr(w, err)
		return
	}
	out := []sessionInfo{}
	for _, it := range sessions.GetItems() {
		out = append(out, sessionInfo{
			ID:         it.GetId(),
			Protocol:   it.GetProtocol(),
			LocalAddr:  it.GetLocalAddr(),
			RemoteAddr: it.GetRemoteAddr(),
			State:      it.GetState(),
		})
	}
	writeJSON(w, http.StatusOK, out)
}
