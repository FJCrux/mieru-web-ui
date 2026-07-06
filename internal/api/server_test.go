// SPDX-License-Identifier: GPL-3.0-or-later

// HTTP-level tests: drive the real router (auth middleware, JSON decoding,
// status codes, response shapes) with the real store and auth, and in-memory
// fakes for mita, geoip, and chain peers.
package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"strings"
	"testing"
	"time"

	pb "github.com/enfein/mieru/v3/pkg/appctl/appctlpb"
	"google.golang.org/protobuf/proto"

	"github.com/fjcrux/mieru-web-ui/internal/auth"
	"github.com/fjcrux/mieru-web-ui/internal/backup"
	"github.com/fjcrux/mieru-web-ui/internal/chain"
	"github.com/fjcrux/mieru-web-ui/internal/geoip"
	"github.com/fjcrux/mieru-web-ui/internal/store"
)

// --- fakes ---

type fakeMita struct {
	cfg    *pb.ServerConfig
	status string
}

func newFakeMita() *fakeMita {
	return &fakeMita{cfg: &pb.ServerConfig{}, status: "RUNNING"}
}

func (f *fakeMita) Status(context.Context) (string, error)  { return f.status, nil }
func (f *fakeMita) Version(context.Context) (string, error) { return "v0.0.0-test", nil }
func (f *fakeMita) Start(context.Context) error             { f.status = "RUNNING"; return nil }
func (f *fakeMita) Stop(context.Context) error              { f.status = "IDLE"; return nil }
func (f *fakeMita) Reload(context.Context) error            { return nil }
func (f *fakeMita) GetConfig(context.Context) (*pb.ServerConfig, error) {
	return proto.Clone(f.cfg).(*pb.ServerConfig), nil
}
func (f *fakeMita) SetConfig(_ context.Context, cfg *pb.ServerConfig) (*pb.ServerConfig, error) {
	f.cfg = proto.Clone(cfg).(*pb.ServerConfig)
	return f.cfg, nil
}
func (f *fakeMita) UpdateConfig(ctx context.Context, mutate func(*pb.ServerConfig) error) (*pb.ServerConfig, error) {
	cfg := proto.Clone(f.cfg).(*pb.ServerConfig)
	if err := mutate(cfg); err != nil {
		return nil, err
	}
	return f.SetConfig(ctx, cfg)
}
func (f *fakeMita) Metrics(context.Context) (string, error) { return "{}", nil }
func (f *fakeMita) Users(context.Context) (*pb.UserWithMetricsList, error) {
	out := &pb.UserWithMetricsList{}
	for _, u := range f.cfg.GetUsers() {
		out.Items = append(out.Items, &pb.UserWithMetrics{User: proto.Clone(u).(*pb.User)})
	}
	return out, nil
}
func (f *fakeMita) Sessions(context.Context) (*pb.SessionInfoList, error) {
	return &pb.SessionInfoList{}, nil
}

type fakeGeo struct{}

func (fakeGeo) Datasets() ([]geoip.Dataset, error)    { return nil, nil }
func (fakeGeo) Categories() ([]geoip.Category, error) { return nil, nil }
func (fakeGeo) CIDRs(code string) ([]string, error) {
	if code == "ru" {
		return []string{"5.8.0.0/19"}, nil
	}
	return nil, nil
}
func (fakeGeo) AddDataset(context.Context, string, string) error { return nil }
func (fakeGeo) DeleteDataset(string) error                       { return nil }
func (fakeGeo) SiteDatasets() ([]geoip.Dataset, error)           { return nil, nil }
func (fakeGeo) SiteCategories() ([]geoip.SiteCategory, error)    { return nil, nil }
func (fakeGeo) Domains(code string) ([]string, error) {
	if code == "telegram" {
		return []string{"t.me", "telegram.org"}, nil
	}
	return nil, fmt.Errorf("geosite category %q not found", code)
}
func (fakeGeo) AddSiteDataset(context.Context, string, string) error { return nil }
func (fakeGeo) DeleteSiteDataset(string) error                       { return nil }

type fakePeers struct {
	names map[string]bool
}

func (f *fakePeers) List() []chain.Status       { return nil }
func (f *fakePeers) Add(name, key string) error { return nil }
func (f *fakePeers) Remove(name string) error   { return nil }
func (f *fakePeers) Names() map[string]bool     { return f.names }
func (f *fakePeers) EgressProxies() []*pb.EgressProxy {
	var out []*pb.EgressProxy
	for name := range f.names {
		out = append(out, &pb.EgressProxy{
			Name:     proto.String(name),
			Protocol: pb.ProxyProtocol_SOCKS5_PROXY_PROTOCOL.Enum(),
			Host:     proto.String("127.0.0.1"),
			Port:     proto.Int32(42000),
		})
	}
	return out
}

// --- harness ---

type testEnv struct {
	srv    *Server
	h      http.Handler
	mita   *fakeMita
	store  store.Store
	cookie *http.Cookie
}

func newTestEnv(t *testing.T) *testEnv {
	t.Helper()
	st, err := store.OpenSQLite(filepath.Join(t.TempDir(), "panel.db"))
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { st.Close() })
	if err := auth.Bootstrap(st, "admin", "secret-pass"); err != nil {
		t.Fatal(err)
	}
	if err := st.SetSetting("desired_ports", "2012"); err != nil {
		t.Fatal(err)
	}

	mita := newFakeMita()
	dir := t.TempDir()
	srv := &Server{
		Mita:      mita,
		Store:     st,
		Auth:      auth.New(st, time.Hour, false),
		Geo:       fakeGeo{},
		Peers:     &fakePeers{names: map[string]bool{}},
		SharePath: "/s",
		Version:   "test",
		Backup: &backup.Paths{
			DBPath:         filepath.Join(t.TempDir(), "panel.db"),
			GeoDir:         filepath.Join(dir, "geoip"),
			PeersDir:       filepath.Join(dir, "peers"),
			MitaConfig:     filepath.Join(dir, "server.conf.pb"),
			RestoreStaging: filepath.Join(dir, ".restore"),
		},
	}
	// Point the backup DB path at the live store's file.
	srv.Backup.DBPath = filepath.Join(dir, "unused-snapshot-base")

	env := &testEnv{srv: srv, h: srv.Routes(), mita: mita, store: st}
	env.cookie = env.login(t, "admin", "secret-pass", http.StatusOK)
	return env
}

// login performs POST /api/login and returns the session cookie (nil unless
// wantStatus is 200).
func (e *testEnv) login(t *testing.T, user, pass string, wantStatus int) *http.Cookie {
	t.Helper()
	rec := httptest.NewRecorder()
	body, _ := json.Marshal(map[string]string{"username": user, "password": pass})
	req := httptest.NewRequest("POST", "/api/login", bytes.NewReader(body))
	e.h.ServeHTTP(rec, req)
	if rec.Code != wantStatus {
		t.Fatalf("login: want %d, got %d (%s)", wantStatus, rec.Code, rec.Body.String())
	}
	for _, c := range rec.Result().Cookies() {
		if c.Name == auth.CookieName {
			return c
		}
	}
	if wantStatus == http.StatusOK {
		t.Fatal("login: no session cookie in response")
	}
	return nil
}

// do sends an authenticated request through the router.
func (e *testEnv) do(t *testing.T, method, path string, body any) *httptest.ResponseRecorder {
	t.Helper()
	var rd *bytes.Reader
	if body != nil {
		b, _ := json.Marshal(body)
		rd = bytes.NewReader(b)
	} else {
		rd = bytes.NewReader(nil)
	}
	req := httptest.NewRequest(method, path, rd)
	req.AddCookie(e.cookie)
	req.Header.Set("X-Requested-With", "XMLHttpRequest")
	rec := httptest.NewRecorder()
	e.h.ServeHTTP(rec, req)
	return rec
}

// --- tests ---

func TestLoginAndAuthGate(t *testing.T) {
	env := newTestEnv(t)

	// Bad credentials are rejected.
	env.login(t, "admin", "wrong-pass", http.StatusUnauthorized)

	// No cookie -> 401 on any API route.
	rec := httptest.NewRecorder()
	env.h.ServeHTTP(rec, httptest.NewRequest("GET", "/api/users", nil))
	if rec.Code != http.StatusUnauthorized {
		t.Fatalf("unauthenticated: want 401, got %d", rec.Code)
	}

	// Mutation without the CSRF header -> 403 even with a valid session.
	req := httptest.NewRequest("POST", "/api/users", strings.NewReader("{}"))
	req.AddCookie(env.cookie)
	rec = httptest.NewRecorder()
	env.h.ServeHTTP(rec, req)
	if rec.Code != http.StatusForbidden {
		t.Fatalf("missing CSRF header: want 403, got %d", rec.Code)
	}

	// Authenticated /api/me works.
	rec = env.do(t, "GET", "/api/me", nil)
	if rec.Code != http.StatusOK || !strings.Contains(rec.Body.String(), "admin") {
		t.Fatalf("/api/me: got %d %s", rec.Code, rec.Body.String())
	}
}

func TestUserLifecycle(t *testing.T) {
	env := newTestEnv(t)

	// Invalid input is rejected.
	if rec := env.do(t, "POST", "/api/users", map[string]any{"name": "bad name!", "password": "longenough"}); rec.Code != http.StatusBadRequest {
		t.Fatalf("bad name: want 400, got %d", rec.Code)
	}
	if rec := env.do(t, "POST", "/api/users", map[string]any{"name": "alice", "password": "short"}); rec.Code != http.StatusBadRequest {
		t.Fatalf("short password: want 400, got %d", rec.Code)
	}

	// Create: user lands in mita's config together with the port bindings
	// (the users<->ports invariant must hold in the same write).
	if rec := env.do(t, "POST", "/api/users", map[string]any{"name": "alice", "password": "longenough"}); rec.Code != http.StatusCreated {
		t.Fatalf("create: want 201, got %d (%s)", rec.Code, rec.Body.String())
	}
	if got := len(env.mita.cfg.GetUsers()); got != 1 {
		t.Fatalf("mita config: want 1 user, got %d", got)
	}
	if got := len(env.mita.cfg.GetPortBindings()); got != 1 {
		t.Fatalf("mita config: want port bindings with a user present, got %d", got)
	}

	// Duplicate name is rejected.
	if rec := env.do(t, "POST", "/api/users", map[string]any{"name": "alice", "password": "longenough"}); rec.Code != http.StatusBadRequest {
		t.Fatalf("duplicate: want 400, got %d", rec.Code)
	}

	// List reports the stored secret.
	rec := env.do(t, "GET", "/api/users", nil)
	if rec.Code != http.StatusOK {
		t.Fatalf("list: want 200, got %d", rec.Code)
	}
	var users []struct {
		Name      string `json:"name"`
		HasSecret bool   `json:"hasSecret"`
	}
	if err := json.Unmarshal(rec.Body.Bytes(), &users); err != nil {
		t.Fatal(err)
	}
	if len(users) != 1 || users[0].Name != "alice" || !users[0].HasSecret {
		t.Fatalf("list: unexpected %+v", users)
	}

	// Delete: last user removed -> ports must be cleared (crashloop guard).
	if rec := env.do(t, "DELETE", "/api/users/alice", nil); rec.Code != http.StatusOK {
		t.Fatalf("delete: want 200, got %d", rec.Code)
	}
	if got := len(env.mita.cfg.GetPortBindings()); got != 0 {
		t.Fatalf("mita config: ports must be cleared with no users, got %d bindings", got)
	}
}

func TestBackupEndpoint(t *testing.T) {
	env := newTestEnv(t)

	// Plain backup: a gzip stream with an attachment disposition.
	rec := env.do(t, "POST", "/api/backup", map[string]string{})
	if rec.Code != http.StatusOK {
		t.Fatalf("backup: want 200, got %d (%s)", rec.Code, rec.Body.String())
	}
	if cd := rec.Header().Get("Content-Disposition"); !strings.Contains(cd, "attachment") {
		t.Fatalf("backup: missing attachment disposition, got %q", cd)
	}
	if b := rec.Body.Bytes(); len(b) < 2 || b[0] != 0x1f || b[1] != 0x8b {
		t.Fatal("backup: body is not gzip")
	}

	// Encrypted backup: carries the magic header, not gzip.
	rec = env.do(t, "POST", "/api/backup", map[string]string{"passphrase": "correct horse"})
	if rec.Code != http.StatusOK {
		t.Fatalf("encrypted backup: want 200, got %d", rec.Code)
	}
	if !bytes.HasPrefix(rec.Body.Bytes(), []byte("IMUGIBK1")) {
		t.Fatal("encrypted backup: missing magic header")
	}
}

// Regression test: mita matches domains as exact-or-suffix and has no "*."
// wildcards, so the panel must rewrite UI patterns into that form or the
// rule silently never matches (geo-routing "not working").
func TestEgressDomainNormalization(t *testing.T) {
	env := newTestEnv(t)

	rec := env.do(t, "PUT", "/api/config/egress", egressConfig{
		Proxies: []egressProxy{},
		Rules: []egressRule{
			{Domains: []string{"*.ru", "РФ", "Example.COM.", "*"}, Action: "DIRECT"},
		},
	})
	if rec.Code != http.StatusOK {
		t.Fatalf("put egress: want 200, got %d (%s)", rec.Code, rec.Body.String())
	}
	rules := env.mita.cfg.GetEgress().GetRules()
	if len(rules) != 1 {
		t.Fatalf("want 1 rule, got %d", len(rules))
	}
	want := []string{"ru", "xn--p1ai", "example.com", "*"}
	got := rules[0].GetDomainNames()
	if len(got) != len(want) {
		t.Fatalf("domains: want %v, got %v", want, got)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("domains[%d]: want %q, got %q", i, want[i], got[i])
		}
	}

	// Garbage that can't be a domain is rejected, not silently stored.
	rec = env.do(t, "PUT", "/api/config/egress", egressConfig{
		Rules: []egressRule{{Domains: []string{"exa mple.com"}, Action: "DIRECT"}},
	})
	if rec.Code != http.StatusBadRequest {
		t.Fatalf("invalid domain: want 400, got %d (%s)", rec.Code, rec.Body.String())
	}
}

// A geosite category on a rule expands to domain suffixes, and a geoip
// category to CIDRs, in the same rule.
func TestEgressGeoExpansion(t *testing.T) {
	env := newTestEnv(t)

	rec := env.do(t, "PUT", "/api/config/egress", egressConfig{
		Rules: []egressRule{
			{Sites: []string{"telegram"}, Geo: []string{"ru"}, Action: "DIRECT"},
		},
	})
	if rec.Code != http.StatusOK {
		t.Fatalf("put egress: want 200, got %d (%s)", rec.Code, rec.Body.String())
	}
	rules := env.mita.cfg.GetEgress().GetRules()
	if len(rules) != 1 {
		t.Fatalf("want 1 rule, got %d", len(rules))
	}
	if got := rules[0].GetDomainNames(); len(got) != 2 || got[0] != "t.me" {
		t.Fatalf("geosite domains: got %v", got)
	}
	if got := rules[0].GetIpRanges(); len(got) != 1 || got[0] != "5.8.0.0/19" {
		t.Fatalf("geoip cidrs: got %v", got)
	}

	// An unknown geosite category is rejected.
	rec = env.do(t, "PUT", "/api/config/egress", egressConfig{
		Rules: []egressRule{{Sites: []string{"nope"}, Action: "DIRECT"}},
	})
	if rec.Code != http.StatusBadRequest {
		t.Fatalf("unknown geosite: want 400, got %d", rec.Code)
	}
}

// Regression test: a chain peer's auto-injected egress proxy must not be
// echoed back by GET (the next PUT would then collide with the peer name).
func TestEgressPeerNameCollision(t *testing.T) {
	env := newTestEnv(t)
	env.srv.Peers = &fakePeers{names: map[string]bool{"sg": true}}

	// Simulate mita's config already carrying the injected peer proxy.
	env.mita.cfg.Egress = &pb.Egress{Proxies: (&fakePeers{names: map[string]bool{"sg": true}}).EgressProxies()}

	// GET falls back to mita's config (nothing stored yet) and must filter
	// out the peer-owned proxy.
	rec := env.do(t, "GET", "/api/config/egress", nil)
	if rec.Code != http.StatusOK {
		t.Fatalf("get egress: want 200, got %d", rec.Code)
	}
	var got struct {
		Proxies []egressProxy `json:"proxies"`
	}
	if err := json.Unmarshal(rec.Body.Bytes(), &got); err != nil {
		t.Fatal(err)
	}
	if len(got.Proxies) != 0 {
		t.Fatalf("get egress: peer proxy leaked into editable list: %+v", got.Proxies)
	}

	// Manually adding a proxy with a peer's name is a clear error.
	rec = env.do(t, "PUT", "/api/config/egress", egressConfig{
		Proxies: []egressProxy{{Name: "sg", Host: "127.0.0.1", Port: 42000}},
		Rules:   []egressRule{},
	})
	if rec.Code != http.StatusBadRequest || !strings.Contains(rec.Body.String(), "chain peer") {
		t.Fatalf("put egress: want 400 mentioning chain peer, got %d (%s)", rec.Code, rec.Body.String())
	}

	// Referencing the peer from a rule (the intended way) works.
	rec = env.do(t, "PUT", "/api/config/egress", egressConfig{
		Proxies: []egressProxy{},
		Rules:   []egressRule{{Geo: []string{"ru"}, Action: "PROXY", Proxies: []string{"sg"}}},
	})
	if rec.Code != http.StatusOK {
		t.Fatalf("put egress with peer rule: want 200, got %d (%s)", rec.Code, rec.Body.String())
	}
	// The pushed config carries the peer proxy exactly once.
	count := 0
	for _, p := range env.mita.cfg.GetEgress().GetProxies() {
		if p.GetName() == "sg" {
			count++
		}
	}
	if count != 1 {
		t.Fatalf("pushed egress: want peer proxy once, got %d", count)
	}
}
