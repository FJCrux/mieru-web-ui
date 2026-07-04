// SPDX-License-Identifier: GPL-3.0-or-later

package api

import (
	"net/http"
	"strconv"
	"time"
)

func (s *Server) handleMitaLogs(w http.ResponseWriter, r *http.Request) {
	if s.Sup == nil {
		writeErr(w, http.StatusNotImplemented, "panel is not supervising mita (external daemon)")
		return
	}
	n := 200
	if q := r.URL.Query().Get("lines"); q != "" {
		if v, err := strconv.Atoi(q); err == nil && v > 0 {
			n = v
		}
	}
	writeJSON(w, http.StatusOK, map[string]any{"lines": s.Sup.Logs(n)})
}

func (s *Server) handleMitaRestart(w http.ResponseWriter, r *http.Request) {
	if s.Sup == nil {
		writeErr(w, http.StatusNotImplemented, "panel is not supervising mita (external daemon)")
		return
	}
	s.Sup.Restart()
	writeJSON(w, http.StatusOK, map[string]bool{"ok": true})
}

// handlePanelRestart gracefully stops the panel process. Under a Docker restart
// policy the container comes back with any new config applied (e.g. path
// changes). It replies first, then exits shortly after.
func (s *Server) handlePanelRestart(w http.ResponseWriter, r *http.Request) {
	if s.Restart == nil {
		writeErr(w, http.StatusNotImplemented, "restart is not available")
		return
	}
	writeJSON(w, http.StatusOK, map[string]bool{"ok": true})
	go func() {
		time.Sleep(300 * time.Millisecond)
		s.Restart()
	}()
}
