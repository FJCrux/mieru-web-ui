// SPDX-License-Identifier: GPL-3.0-or-later

// Package api implements the panel's REST surface consumed by the SPA.
package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/fjcrux/mieru-web-ui/internal/auth"
	"github.com/fjcrux/mieru-web-ui/internal/backup"
	"github.com/fjcrux/mieru-web-ui/internal/store"
)

// Server bundles the ports handlers depend on (see ports.go).
type Server struct {
	Mita  MitaController
	Store store.Store
	Auth  *auth.Auth
	// Sup is nil when running with --no-supervise (external mita).
	Sup ProcSupervisor
	// Geo expands GeoIP categories into CIDRs for egress rules.
	Geo GeoResolver
	// Peers runs upstream chain peers.
	Peers PeerManager
	// SharePath is the public URL prefix for share links (e.g. "/s"),
	// decoupled from the admin base path so links don't reveal it.
	SharePath string
	// SubPath is the public URL prefix for subscription links (e.g. "/sub").
	SubPath string
	// SubPort is the dedicated subscription listener port ("" = the panel
	// port); subscription URLs are built with it.
	SubPort string
	// ActiveBasePath / ActiveSharePath / ActiveSubPath / ActiveSubPort are
	// the path settings the running process started with; comparing them to
	// the stored settings tells us a restart is pending.
	ActiveBasePath  string
	ActiveSharePath string
	ActiveSubPath   string
	ActiveSubPort   string
	// PortsManaged is true when port bindings are enforced from PROXY_PORTS;
	// the Network UI then shows them read-only.
	PortsManaged bool
	// PanelPort is the configured panel port, used for a hardening hint.
	PanelPort string
	// Restart triggers a graceful shutdown of the panel process. Under a
	// Docker restart policy the container comes back with the new config
	// (this is how path changes are applied). nil disables the endpoint.
	Restart func()
	// TLSEnabled is true when the panel serves HTTPS; when false the
	// dashboard shows an insecure-connection warning on non-loopback access.
	TLSEnabled bool
	// TrustProxy makes the panel believe X-Forwarded-Proto from the reverse
	// proxy / load balancer in front of it (so a TLS-terminating balancer
	// silences the plain-HTTP warning). Only enable when such a proxy is
	// actually present and reachable clients can't forge the header.
	TrustProxy bool
	// Version is the running panel version, shown in the UI.
	Version string
	// Updates checks GitHub for a newer release (nil disables the check).
	Updates *updateChecker
	// Backup locates panel state for the backup/restore endpoints (nil = off).
	Backup *backup.Paths
}

// Routes returns the /api handler tree with auth applied.
func (s *Server) Routes() http.Handler {
	authed := http.NewServeMux()
	authed.HandleFunc("POST /api/logout", s.handleLogout)
	authed.HandleFunc("GET /api/me", s.handleMe)
	authed.HandleFunc("GET /api/version", s.handleVersion)

	authed.HandleFunc("GET /api/dashboard", s.handleDashboard)
	authed.HandleFunc("GET /api/sessions", s.handleSessions)

	authed.HandleFunc("GET /api/users", s.handleListUsers)
	authed.HandleFunc("POST /api/users", s.handleCreateUser)
	authed.HandleFunc("PUT /api/users/{name}", s.handleUpdateUser)
	authed.HandleFunc("DELETE /api/users/{name}", s.handleDeleteUser)
	authed.HandleFunc("GET /api/users/{name}/share", s.handleShare)
	authed.HandleFunc("POST /api/users/{name}/share-token", s.handleCreateShareToken)
	authed.HandleFunc("GET /api/users/{name}/subscription", s.handleGetSubscription)
	authed.HandleFunc("POST /api/users/{name}/subscription", s.handleCreateSubscription)
	authed.HandleFunc("DELETE /api/users/{name}/subscription", s.handleDeleteSubscription)

	authed.HandleFunc("GET /api/config/network", s.handleGetNetwork)
	authed.HandleFunc("PUT /api/config/network", s.handlePutNetwork)

	authed.HandleFunc("GET /api/config/advanced", s.handleGetAdvanced)
	authed.HandleFunc("PUT /api/config/advanced", s.handlePutAdvanced)

	authed.HandleFunc("GET /api/config/egress", s.handleGetEgress)
	authed.HandleFunc("PUT /api/config/egress", s.handlePutEgress)

	authed.HandleFunc("GET /api/geoip", s.handleGeoipList)
	authed.HandleFunc("POST /api/geoip/datasets", s.handleGeoipAddDataset)
	authed.HandleFunc("DELETE /api/geoip/datasets/{name}", s.handleGeoipDeleteDataset)

	authed.HandleFunc("GET /api/chain", s.handleChainList)
	authed.HandleFunc("POST /api/chain", s.handleChainAdd)
	authed.HandleFunc("DELETE /api/chain/{name}", s.handleChainRemove)
	authed.HandleFunc("POST /api/chain/key", s.handleChainKey)

	authed.HandleFunc("GET /api/mita/logs", s.handleMitaLogs)
	authed.HandleFunc("POST /api/mita/restart", s.handleMitaRestart)
	authed.HandleFunc("POST /api/panel/restart", s.handlePanelRestart)

	authed.HandleFunc("GET /api/settings", s.handleGetSettings)
	authed.HandleFunc("PUT /api/settings", s.handlePutSettings)
	authed.HandleFunc("PUT /api/settings/password", s.handleChangePassword)

	authed.HandleFunc("POST /api/backup", s.handleBackup)
	authed.HandleFunc("POST /api/restore", s.handleRestore)

	mux := http.NewServeMux()
	mux.HandleFunc("POST /api/login", s.handleLogin)
	mux.Handle("/api/", s.Auth.Middleware(authed))
	return mux
}

// --- helpers ---

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(v); err != nil {
		log.Printf("write response: %v", err)
	}
}

func writeErr(w http.ResponseWriter, status int, msg string) {
	writeJSON(w, status, map[string]string{"error": msg})
}

func readJSON(w http.ResponseWriter, r *http.Request, v any) bool {
	dec := json.NewDecoder(http.MaxBytesReader(w, r.Body, 1<<20))
	dec.DisallowUnknownFields()
	if err := dec.Decode(v); err != nil {
		writeErr(w, http.StatusBadRequest, "invalid request body: "+err.Error())
		return false
	}
	return true
}
