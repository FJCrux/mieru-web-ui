// SPDX-License-Identifier: GPL-3.0-or-later

package api

import (
	"errors"
	"net/http"

	"github.com/fjcrux/mieru-web-ui/internal/auth"
)

type loginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (s *Server) handleLogin(w http.ResponseWriter, r *http.Request) {
	var req loginRequest
	if !readJSON(w, r, &req) {
		return
	}
	token, err := s.Auth.Login(auth.RemoteIP(r), req.Username, req.Password)
	switch {
	case errors.Is(err, auth.ErrRateLimited):
		writeErr(w, http.StatusTooManyRequests, "too many login attempts, try again later")
		return
	case errors.Is(err, auth.ErrBadCredentials):
		writeErr(w, http.StatusUnauthorized, "invalid username or password")
		return
	case err != nil:
		writeErr(w, http.StatusInternalServerError, "login failed")
		return
	}
	s.Auth.SetCookie(w, token)
	writeJSON(w, http.StatusOK, map[string]bool{"ok": true})
}

func (s *Server) handleLogout(w http.ResponseWriter, r *http.Request) {
	if c, err := r.Cookie(auth.CookieName); err == nil {
		_ = s.Auth.Logout(c.Value)
	}
	s.Auth.ClearCookie(w)
	writeJSON(w, http.StatusOK, map[string]bool{"ok": true})
}

func (s *Server) handleMe(w http.ResponseWriter, r *http.Request) {
	username, _, err := s.Store.Admin()
	if err != nil {
		writeErr(w, http.StatusInternalServerError, "failed to load admin")
		return
	}
	writeJSON(w, http.StatusOK, map[string]string{"username": username})
}

type changePasswordRequest struct {
	CurrentPassword string `json:"currentPassword"`
	NewUsername     string `json:"newUsername"`
	NewPassword     string `json:"newPassword"`
}

func (s *Server) handleChangePassword(w http.ResponseWriter, r *http.Request) {
	var req changePasswordRequest
	if !readJSON(w, r, &req) {
		return
	}
	if len(req.NewPassword) < 8 {
		writeErr(w, http.StatusBadRequest, "new password must be at least 8 characters")
		return
	}
	// Re-verify the current password before changing credentials.
	if err := s.Auth.VerifyPassword(req.CurrentPassword); err != nil {
		writeErr(w, http.StatusUnauthorized, "current password is incorrect")
		return
	}
	username := req.NewUsername
	if username == "" {
		username = currentUsername(s)
	}
	hash, err := auth.HashPassword(req.NewPassword)
	if err != nil {
		writeErr(w, http.StatusInternalServerError, "failed to hash password")
		return
	}
	if err := s.Store.SetAdmin(username, hash); err != nil {
		writeErr(w, http.StatusInternalServerError, "failed to save admin")
		return
	}
	// Invalidate every session, including this one; the user logs in again.
	_ = s.Store.DeleteAllSessions()
	s.Auth.ClearCookie(w)
	writeJSON(w, http.StatusOK, map[string]bool{"ok": true})
}

func currentUsername(s *Server) string {
	u, _, _ := s.Store.Admin()
	return u
}
