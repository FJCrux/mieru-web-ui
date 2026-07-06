// SPDX-License-Identifier: GPL-3.0-or-later

package api

import (
	"encoding/base64"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// subEnv creates a user with a stored secret and a public host, and returns a
// mux carrying the public subscription route (which lives on the root mux in
// main.go, so the router alone can't serve it).
func subEnv(t *testing.T) (*testEnv, *http.ServeMux) {
	t.Helper()
	env := newTestEnv(t)
	env.srv.SubPath = "/sub"
	if rec := env.do(t, "POST", "/api/users", map[string]any{"name": "alice", "password": "longenough"}); rec.Code != http.StatusCreated {
		t.Fatalf("create user: want 201, got %d (%s)", rec.Code, rec.Body.String())
	}
	if err := env.store.SetSetting("public_host", "203.0.113.10"); err != nil {
		t.Fatal(err)
	}
	pub := http.NewServeMux()
	pub.HandleFunc("GET /sub/{token}", env.srv.HandlePublicSubscription)
	return env, pub
}

func mintSub(t *testing.T, env *testEnv) string {
	t.Helper()
	rec := env.do(t, "POST", "/api/users/alice/subscription", map[string]any{})
	if rec.Code != http.StatusCreated {
		t.Fatalf("mint: want 201, got %d (%s)", rec.Code, rec.Body.String())
	}
	var resp struct {
		URL       string `json:"url"`
		CreatedAt int64  `json:"createdAt"`
	}
	if err := json.Unmarshal(rec.Body.Bytes(), &resp); err != nil {
		t.Fatal(err)
	}
	if resp.URL == "" || resp.CreatedAt == 0 {
		t.Fatalf("mint: unexpected response %+v", resp)
	}
	return resp.URL[strings.LastIndex(resp.URL, "/")+1:]
}

func fetchSub(pub *http.ServeMux, path string) *httptest.ResponseRecorder {
	rec := httptest.NewRecorder()
	pub.ServeHTTP(rec, httptest.NewRequest("GET", path, nil))
	return rec
}

func TestSubscriptionLifecycle(t *testing.T) {
	env, pub := subEnv(t)

	// No token yet.
	rec := env.do(t, "GET", "/api/users/alice/subscription", nil)
	if rec.Code != http.StatusOK || !strings.Contains(rec.Body.String(), `"exists":false`) {
		t.Fatalf("status: want exists:false, got %d (%s)", rec.Code, rec.Body.String())
	}

	token := mintSub(t, env)

	// Status now re-displays the URL (plaintext token storage).
	rec = env.do(t, "GET", "/api/users/alice/subscription", nil)
	if !strings.Contains(rec.Body.String(), token) {
		t.Fatalf("status should include the URL: %s", rec.Body.String())
	}

	// Default format: full Clash profile with headers.
	rec = fetchSub(pub, "/sub/"+token)
	if rec.Code != http.StatusOK {
		t.Fatalf("fetch: want 200, got %d (%s)", rec.Code, rec.Body.String())
	}
	if ct := rec.Header().Get("Content-Type"); ct != "text/yaml; charset=utf-8" {
		t.Fatalf("content-type: %q", ct)
	}
	if rec.Header().Get("Profile-Update-Interval") != "24" {
		t.Fatal("missing profile-update-interval")
	}
	if ui := rec.Header().Get("Subscription-Userinfo"); !strings.Contains(ui, "upload=") || !strings.Contains(ui, "download=") {
		t.Fatalf("subscription-userinfo: %q", ui)
	}
	body := rec.Body.String()
	for _, want := range []string{"type: mieru", "server: 203.0.113.10", "proxy-groups", "MATCH,PROXY"} {
		if !strings.Contains(body, want) {
			t.Fatalf("profile missing %q:\n%s", want, body)
		}
	}

	// Rotation invalidates the old token.
	newToken := mintSub(t, env)
	if newToken == token {
		t.Fatal("rotation should mint a different token")
	}
	if rec := fetchSub(pub, "/sub/"+token); rec.Code != http.StatusNotFound {
		t.Fatalf("old token: want 404, got %d", rec.Code)
	}
	if rec := fetchSub(pub, "/sub/"+newToken); rec.Code != http.StatusOK {
		t.Fatalf("new token: want 200, got %d", rec.Code)
	}

	// Revocation.
	if rec := env.do(t, "DELETE", "/api/users/alice/subscription", nil); rec.Code != http.StatusOK {
		t.Fatalf("revoke: want 200, got %d", rec.Code)
	}
	if rec := fetchSub(pub, "/sub/"+newToken); rec.Code != http.StatusNotFound {
		t.Fatalf("revoked token: want 404, got %d", rec.Code)
	}
}

func TestSubscriptionFormats(t *testing.T) {
	env, pub := subEnv(t)
	token := mintSub(t, env)

	// proxies-only flavor has no groups or rules.
	rec := fetchSub(pub, "/sub/"+token+"?flavor=proxies")
	if rec.Code != http.StatusOK {
		t.Fatalf("flavor=proxies: want 200, got %d", rec.Code)
	}
	if body := rec.Body.String(); !strings.Contains(body, "type: mieru") || strings.Contains(body, "proxy-groups") {
		t.Fatalf("flavor=proxies body:\n%s", body)
	}

	// mierus: base64-encoded line list of mierus:// links.
	rec = fetchSub(pub, "/sub/"+token+"?format=mierus")
	if rec.Code != http.StatusOK {
		t.Fatalf("format=mierus: want 200, got %d", rec.Code)
	}
	if ct := rec.Header().Get("Content-Type"); ct != "text/plain; charset=utf-8" {
		t.Fatalf("mierus content-type: %q", ct)
	}
	decoded, err := base64.StdEncoding.DecodeString(rec.Body.String())
	if err != nil {
		t.Fatalf("mierus body is not base64: %v", err)
	}
	for _, line := range strings.Split(string(decoded), "\n") {
		if !strings.HasPrefix(line, "mierus://") {
			t.Fatalf("unexpected line %q", line)
		}
	}

	// mieru: base64-encoded mieru:// link.
	rec = fetchSub(pub, "/sub/"+token+"?format=mieru")
	decoded, err = base64.StdEncoding.DecodeString(rec.Body.String())
	if err != nil || !strings.HasPrefix(string(decoded), "mieru://") {
		t.Fatalf("format=mieru: %v %q", err, decoded)
	}

	// Unknown format/flavor and garbage tokens are indistinguishable 404s.
	for _, path := range []string{
		"/sub/" + token + "?format=singbox",
		"/sub/" + token + "?flavor=bogus",
		"/sub/" + token + "?format=mierus&flavor=proxies",
		"/sub/deadbeef",
	} {
		if rec := fetchSub(pub, path); rec.Code != http.StatusNotFound {
			t.Fatalf("%s: want 404, got %d", path, rec.Code)
		}
	}
}

func TestSubscriptionDeletedUser(t *testing.T) {
	env, pub := subEnv(t)
	token := mintSub(t, env)

	garbage := fetchSub(pub, "/sub/deadbeef")
	if rec := env.do(t, "DELETE", "/api/users/alice", nil); rec.Code != http.StatusOK {
		t.Fatalf("delete user: want 200, got %d", rec.Code)
	}
	rec := fetchSub(pub, "/sub/"+token)
	if rec.Code != http.StatusNotFound || rec.Body.String() != garbage.Body.String() {
		t.Fatalf("deleted user must look like a garbage token: %d %q vs %q",
			rec.Code, rec.Body.String(), garbage.Body.String())
	}
}

func TestSubscriptionURLPort(t *testing.T) {
	env, _ := subEnv(t)
	env.srv.SubPort = "8688"
	if err := env.store.SetSetting("panel_url", "https://vpn.example.com"); err != nil {
		t.Fatal(err)
	}
	rec := env.do(t, "POST", "/api/users/alice/subscription", map[string]any{})
	if rec.Code != http.StatusCreated {
		t.Fatalf("mint: want 201, got %d", rec.Code)
	}
	if !strings.Contains(rec.Body.String(), "https://vpn.example.com:8688/sub/") {
		t.Fatalf("URL should carry the sub port: %s", rec.Body.String())
	}
}

func TestSettingsSubValidation(t *testing.T) {
	env := newTestEnv(t)
	env.srv.PanelPort = "8686"

	base := map[string]any{"publicHost": "", "panelUrl": "", "basePath": "", "sharePath": "/s"}
	put := func(overrides map[string]any) *httptest.ResponseRecorder {
		body := map[string]any{"subPath": "", "subPort": "", "restartPending": false}
		for k, v := range base {
			body[k] = v
		}
		for k, v := range overrides {
			body[k] = v
		}
		return env.do(t, "PUT", "/api/settings", body)
	}

	if rec := put(map[string]any{"subPath": "/s"}); rec.Code != http.StatusBadRequest {
		t.Fatalf("subPath == sharePath: want 400, got %d", rec.Code)
	}
	for _, port := range []string{"0", "70000", "abc", "8686"} {
		if rec := put(map[string]any{"subPort": port}); rec.Code != http.StatusBadRequest {
			t.Fatalf("subPort %q: want 400, got %d", port, rec.Code)
		}
	}
	if rec := put(map[string]any{"subPath": "/sub", "subPort": "8688"}); rec.Code != http.StatusOK {
		t.Fatalf("valid settings: want 200, got %d (%s)", rec.Code, rec.Body.String())
	}

	// A stored sub path differing from the active one flags a restart.
	env.srv.ActiveSharePath = "/s"
	env.srv.ActiveSubPath = "/sub"
	env.srv.ActiveSubPort = "8688"
	rec := env.do(t, "GET", "/api/settings", nil)
	if strings.Contains(rec.Body.String(), `"restartPending":true`) {
		t.Fatalf("no restart should be pending: %s", rec.Body.String())
	}
	if rec := put(map[string]any{"subPath": "/newsub", "subPort": "8688"}); rec.Code != http.StatusOK {
		t.Fatal("save failed")
	}
	rec = env.do(t, "GET", "/api/settings", nil)
	if !strings.Contains(rec.Body.String(), `"restartPending":true`) {
		t.Fatalf("restart should be pending after sub path change: %s", rec.Body.String())
	}
}
