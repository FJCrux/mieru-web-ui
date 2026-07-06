// SPDX-License-Identifier: GPL-3.0-or-later

package api

import (
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type settingsResponse struct {
	// PublicHost is the mieru server address embedded in client configs and
	// share links (the address clients connect to).
	PublicHost string `json:"publicHost"`
	// PanelURL is this panel's external URL (e.g. https://vpn.example.com),
	// used as the base for share links and to validate the Host header.
	PanelURL string `json:"panelUrl"`
	// BasePath is the secret admin prefix; SharePath is the public share
	// prefix; SubPath is the public subscription prefix. All take effect
	// after a restart.
	BasePath  string `json:"basePath"`
	SharePath string `json:"sharePath"`
	SubPath   string `json:"subPath"`
	// SubPort is the dedicated subscription listener port; empty serves
	// subscriptions on the panel port. Takes effect after a restart.
	SubPort string `json:"subPort"`
	// RestartPending is true when a saved path differs from the running one.
	RestartPending bool `json:"restartPending"`
}

// restartPending reports whether stored path settings differ from the ones the
// process is currently running with (so a restart would change behaviour).
func (s *Server) restartPending() bool {
	storedBase, _ := s.Store.Setting("base_path")
	base, _ := cleanURLPath(storedBase)
	storedShare, _ := s.Store.Setting("share_path")
	share, _ := cleanURLPath(storedShare)
	if share == "" {
		share = "/s"
	}
	storedSub, _ := s.Store.Setting("sub_path")
	sub, _ := cleanURLPath(storedSub)
	if sub == "" {
		sub = "/sub"
	}
	storedSubPort, _ := s.Store.Setting("sub_port")
	return base != s.ActiveBasePath || share != s.ActiveSharePath ||
		sub != s.ActiveSubPath || storedSubPort != s.ActiveSubPort
}

func (s *Server) handleGetSettings(w http.ResponseWriter, r *http.Request) {
	host, err := s.Store.Setting("public_host")
	if err != nil {
		writeErr(w, http.StatusInternalServerError, err.Error())
		return
	}
	panelURL, _ := s.Store.Setting("panel_url")
	basePath, _ := s.Store.Setting("base_path")
	sharePath, _ := s.Store.Setting("share_path")
	subPath, _ := s.Store.Setting("sub_path")
	subPort, _ := s.Store.Setting("sub_port")
	writeJSON(w, http.StatusOK, settingsResponse{
		PublicHost: host, PanelURL: panelURL, BasePath: basePath, SharePath: sharePath,
		SubPath: subPath, SubPort: subPort,
		RestartPending: s.restartPending(),
	})
}

func (s *Server) handlePutSettings(w http.ResponseWriter, r *http.Request) {
	var req settingsResponse
	if !readJSON(w, r, &req) {
		return
	}
	req.PanelURL = strings.TrimRight(strings.TrimSpace(req.PanelURL), "/")
	if req.PanelURL != "" {
		u, err := url.Parse(req.PanelURL)
		if err != nil || (u.Scheme != "http" && u.Scheme != "https") || u.Host == "" {
			writeErr(w, http.StatusBadRequest, "panel URL must look like https://host[:port]")
			return
		}
	}
	base, ok := cleanURLPath(req.BasePath)
	if !ok {
		writeErr(w, http.StatusBadRequest, "base path must look like /something")
		return
	}
	share, ok := cleanURLPath(req.SharePath)
	if !ok {
		writeErr(w, http.StatusBadRequest, "share path must look like /something")
		return
	}
	if share != "" && share == base {
		writeErr(w, http.StatusBadRequest, "share path must differ from the base path")
		return
	}
	sub, ok := cleanURLPath(req.SubPath)
	if !ok {
		writeErr(w, http.StatusBadRequest, "subscription path must look like /something")
		return
	}
	if sub != "" && (sub == base || sub == share) {
		writeErr(w, http.StatusBadRequest, "subscription path must differ from the base and share paths")
		return
	}
	subPort := strings.TrimSpace(req.SubPort)
	if subPort != "" {
		n, err := strconv.Atoi(subPort)
		if err != nil || n < 1 || n > 65535 {
			writeErr(w, http.StatusBadRequest, "subscription port must be 1-65535 or empty")
			return
		}
		if subPort == s.PanelPort {
			writeErr(w, http.StatusBadRequest, "subscription port must differ from the panel port")
			return
		}
	}

	for k, v := range map[string]string{
		"public_host": req.PublicHost, "panel_url": req.PanelURL,
		"base_path": base, "share_path": share,
		"sub_path": sub, "sub_port": subPort,
	} {
		if err := s.Store.SetSetting(k, v); err != nil {
			writeErr(w, http.StatusInternalServerError, err.Error())
			return
		}
	}
	writeJSON(w, http.StatusOK, map[string]bool{"ok": true})
}

// cleanURLPath normalises a URL prefix to "/x" (no trailing slash) or "".
func cleanURLPath(p string) (string, bool) {
	p = strings.TrimSpace(p)
	if p == "" || p == "/" {
		return "", true
	}
	if !strings.HasPrefix(p, "/") {
		p = "/" + p
	}
	p = strings.TrimRight(p, "/")
	// One segment, url-safe characters only.
	if strings.Count(p, "/") != 1 {
		return "", false
	}
	for _, r := range p[1:] {
		if !(r >= 'a' && r <= 'z') && !(r >= 'A' && r <= 'Z') && !(r >= '0' && r <= '9') && r != '-' && r != '_' {
			return "", false
		}
	}
	return p, true
}
