// SPDX-License-Identifier: GPL-3.0-or-later

// Package auth implements panel admin authentication: bcrypt password
// verification, opaque session tokens (SHA-256 hashed at rest), the HTTP
// middleware guarding the API, and login rate limiting.
package auth

import (
	"crypto/rand"
	"crypto/sha256"
	"crypto/subtle"
	"encoding/hex"
	"errors"
	"log"
	"net"
	"net/http"
	"sync"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/fjcrux/mieru-web-ui/internal/store"
)

const (
	CookieName = "mieru_web_ui_session"

	bcryptCost = 12

	loginWindow      = time.Minute
	loginMaxAttempts = 5
)

// DefaultSessionTTL is used when the session_ttl_hours setting is unset.
const DefaultSessionTTL = 7 * 24 * time.Hour

type Auth struct {
	store      store.Store
	sessionTTL time.Duration
	secure     bool // set Secure flag on cookies (TLS enabled)

	mu       sync.Mutex
	attempts map[string][]time.Time
}

func New(st store.Store, sessionTTL time.Duration, secureCookies bool) *Auth {
	if sessionTTL <= 0 {
		sessionTTL = DefaultSessionTTL
	}
	return &Auth{
		store:      st,
		sessionTTL: sessionTTL,
		secure:     secureCookies,
		attempts:   make(map[string][]time.Time),
	}
}

// HashPassword bcrypt-hashes a panel admin password.
func HashPassword(password string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(password), bcryptCost)
	return string(b), err
}

// GenerateToken returns a random 256-bit hex token.
func GenerateToken() (string, error) {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return hex.EncodeToString(b), nil
}

func hashToken(token string) string {
	h := sha256.Sum256([]byte(token))
	return hex.EncodeToString(h[:])
}

var ErrRateLimited = errors.New("too many login attempts")
var ErrBadCredentials = errors.New("invalid username or password")

// Login verifies credentials and returns a new session token.
func (a *Auth) Login(remoteIP, username, password string) (string, error) {
	if !a.allowAttempt(remoteIP) {
		return "", ErrRateLimited
	}
	adminUser, adminHash, err := a.store.Admin()
	if err != nil {
		return "", err
	}
	userOK := subtle.ConstantTimeCompare([]byte(adminUser), []byte(username)) == 1
	passErr := bcrypt.CompareHashAndPassword([]byte(adminHash), []byte(password))
	if !userOK || passErr != nil {
		return "", ErrBadCredentials
	}
	token, err := GenerateToken()
	if err != nil {
		return "", err
	}
	if err := a.store.CreateSession(hashToken(token), time.Now().Add(a.sessionTTL)); err != nil {
		return "", err
	}
	a.clearAttempts(remoteIP)
	return token, nil
}

// VerifyPassword checks the admin password without creating a session
// (used for re-authentication before sensitive changes).
func (a *Auth) VerifyPassword(password string) error {
	_, adminHash, err := a.store.Admin()
	if err != nil {
		return err
	}
	if bcrypt.CompareHashAndPassword([]byte(adminHash), []byte(password)) != nil {
		return ErrBadCredentials
	}
	return nil
}

// Logout invalidates the session token.
func (a *Auth) Logout(token string) error {
	return a.store.DeleteSession(hashToken(token))
}

// Valid reports whether the session token is active.
func (a *Auth) Valid(token string) bool {
	ok, err := a.store.SessionValid(hashToken(token))
	if err != nil {
		log.Printf("session lookup failed: %v", err)
		return false
	}
	return ok
}

// SetCookie writes the session cookie on a successful login.
func (a *Auth) SetCookie(w http.ResponseWriter, token string) {
	http.SetCookie(w, &http.Cookie{
		Name:     CookieName,
		Value:    token,
		Path:     "/",
		MaxAge:   int(a.sessionTTL.Seconds()),
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
		Secure:   a.secure,
	})
}

// ClearCookie expires the session cookie (logout).
func (a *Auth) ClearCookie(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name:     CookieName,
		Value:    "",
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
		Secure:   a.secure,
	})
}

// Middleware rejects requests without a valid session. Mutating methods
// must also carry X-Requested-With: XMLHttpRequest as a CSRF belt over
// the SameSite=Lax suspenders.
func (a *Auth) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c, err := r.Cookie(CookieName)
		if err != nil || !a.Valid(c.Value) {
			http.Error(w, `{"error":"unauthorized"}`, http.StatusUnauthorized)
			return
		}
		switch r.Method {
		case http.MethodPost, http.MethodPut, http.MethodPatch, http.MethodDelete:
			if r.Header.Get("X-Requested-With") != "XMLHttpRequest" {
				http.Error(w, `{"error":"missing X-Requested-With header"}`, http.StatusForbidden)
				return
			}
		}
		next.ServeHTTP(w, r)
	})
}

// allowAttempt implements a sliding-window limit of loginMaxAttempts
// per loginWindow per IP.
func (a *Auth) allowAttempt(ip string) bool {
	a.mu.Lock()
	defer a.mu.Unlock()
	now := time.Now()
	recent := a.attempts[ip][:0]
	for _, t := range a.attempts[ip] {
		if now.Sub(t) < loginWindow {
			recent = append(recent, t)
		}
	}
	if len(recent) >= loginMaxAttempts {
		a.attempts[ip] = recent
		return false
	}
	a.attempts[ip] = append(recent, now)
	return true
}

func (a *Auth) clearAttempts(ip string) {
	a.mu.Lock()
	defer a.mu.Unlock()
	delete(a.attempts, ip)
}

// RemoteIP extracts the client IP from a request (no proxy headers -
// the panel is expected to be accessed directly or the operator accepts
// per-proxy-IP rate limiting).
func RemoteIP(r *http.Request) string {
	host, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return r.RemoteAddr
	}
	return host
}
