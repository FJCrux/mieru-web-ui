// SPDX-License-Identifier: GPL-3.0-or-later

package auth

import (
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"testing"
	"time"

	"github.com/fjcrux/mieru-web-ui/internal/store"
)

func testAuth(t *testing.T) (*Auth, store.Store) {
	t.Helper()
	st, err := store.OpenSQLite(filepath.Join(t.TempDir(), "panel.db"))
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { st.Close() })
	if err := Bootstrap(st, "admin", "secret-pass"); err != nil {
		t.Fatal(err)
	}
	return New(st, time.Hour, false), st
}

func TestLoginLogout(t *testing.T) {
	a, _ := testAuth(t)

	if _, err := a.Login("1.2.3.4", "admin", "wrong"); err != ErrBadCredentials {
		t.Fatalf("want ErrBadCredentials, got %v", err)
	}
	if _, err := a.Login("1.2.3.4", "nobody", "secret-pass"); err != ErrBadCredentials {
		t.Fatalf("want ErrBadCredentials, got %v", err)
	}

	token, err := a.Login("1.2.3.4", "admin", "secret-pass")
	if err != nil {
		t.Fatal(err)
	}
	if !a.Valid(token) {
		t.Fatal("fresh token should be valid")
	}
	if err := a.Logout(token); err != nil {
		t.Fatal(err)
	}
	if a.Valid(token) {
		t.Fatal("token should be invalid after logout")
	}
}

func TestRateLimit(t *testing.T) {
	a, _ := testAuth(t)
	for i := 0; i < 5; i++ {
		if _, err := a.Login("9.9.9.9", "admin", "wrong"); err != ErrBadCredentials {
			t.Fatalf("attempt %d: want ErrBadCredentials, got %v", i, err)
		}
	}
	if _, err := a.Login("9.9.9.9", "admin", "secret-pass"); err != ErrRateLimited {
		t.Fatalf("want ErrRateLimited, got %v", err)
	}
	// Other IPs are unaffected.
	if _, err := a.Login("8.8.8.8", "admin", "secret-pass"); err != nil {
		t.Fatalf("other IP should not be limited: %v", err)
	}
}

func TestMiddleware(t *testing.T) {
	a, _ := testAuth(t)
	token, err := a.Login("1.2.3.4", "admin", "secret-pass")
	if err != nil {
		t.Fatal(err)
	}

	ok := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	})
	h := a.Middleware(ok)

	// No cookie -> 401.
	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, httptest.NewRequest("GET", "/api/users", nil))
	if rec.Code != http.StatusUnauthorized {
		t.Fatalf("no cookie: want 401, got %d", rec.Code)
	}

	// Valid cookie GET -> pass.
	req := httptest.NewRequest("GET", "/api/users", nil)
	req.AddCookie(&http.Cookie{Name: CookieName, Value: token})
	rec = httptest.NewRecorder()
	h.ServeHTTP(rec, req)
	if rec.Code != http.StatusNoContent {
		t.Fatalf("valid GET: want 204, got %d", rec.Code)
	}

	// Mutating request without CSRF header -> 403.
	req = httptest.NewRequest("POST", "/api/users", nil)
	req.AddCookie(&http.Cookie{Name: CookieName, Value: token})
	rec = httptest.NewRecorder()
	h.ServeHTTP(rec, req)
	if rec.Code != http.StatusForbidden {
		t.Fatalf("POST without header: want 403, got %d", rec.Code)
	}

	// Mutating request with CSRF header -> pass.
	req = httptest.NewRequest("POST", "/api/users", nil)
	req.AddCookie(&http.Cookie{Name: CookieName, Value: token})
	req.Header.Set("X-Requested-With", "XMLHttpRequest")
	rec = httptest.NewRecorder()
	h.ServeHTTP(rec, req)
	if rec.Code != http.StatusNoContent {
		t.Fatalf("POST with header: want 204, got %d", rec.Code)
	}
}

func TestBootstrapIdempotent(t *testing.T) {
	st, err := store.OpenSQLite(filepath.Join(t.TempDir(), "panel.db"))
	if err != nil {
		t.Fatal(err)
	}
	defer st.Close()
	if err := Bootstrap(st, "admin", "first"); err != nil {
		t.Fatal(err)
	}
	// Second bootstrap must not overwrite the existing account.
	if err := Bootstrap(st, "other", "second"); err != nil {
		t.Fatal(err)
	}
	user, _, err := st.Admin()
	if err != nil {
		t.Fatal(err)
	}
	if user != "admin" {
		t.Fatalf("admin overwritten: %s", user)
	}
}
