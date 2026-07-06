// SPDX-License-Identifier: GPL-3.0-or-later

package geoip

import (
	"context"
	"net"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"google.golang.org/protobuf/encoding/protowire"
)

// buildDat encodes a minimal geoip.dat with the given category -> CIDRs.
func buildDat(t *testing.T, cats map[string][]string) []byte {
	t.Helper()
	var out []byte
	for code, cidrs := range cats {
		var geoip []byte
		geoip = protowire.AppendTag(geoip, 1, protowire.BytesType)
		geoip = protowire.AppendBytes(geoip, []byte(code))
		for _, c := range cidrs {
			_, ipnet, err := net.ParseCIDR(c)
			if err != nil {
				t.Fatal(err)
			}
			ip := ipnet.IP.To4()
			if ip == nil {
				ip = ipnet.IP.To16()
			}
			ones, _ := ipnet.Mask.Size()
			var cidr []byte
			cidr = protowire.AppendTag(cidr, 1, protowire.BytesType)
			cidr = protowire.AppendBytes(cidr, ip)
			cidr = protowire.AppendTag(cidr, 2, protowire.VarintType)
			cidr = protowire.AppendVarint(cidr, uint64(ones))
			geoip = protowire.AppendTag(geoip, 2, protowire.BytesType)
			geoip = protowire.AppendBytes(geoip, cidr)
		}
		out = protowire.AppendTag(out, 1, protowire.BytesType)
		out = protowire.AppendBytes(out, geoip)
	}
	return out
}

// buildSiteDat encodes a minimal geosite.dat with the given category ->
// []domain{value, type}. type: 2=Domain(suffix), 3=Full, 0=Plain, 1=Regex.
func buildSiteDat(t *testing.T, cats map[string][]struct {
	Value string
	Type  uint64
}) []byte {
	t.Helper()
	var out []byte
	for code, domains := range cats {
		var site []byte
		site = protowire.AppendTag(site, 1, protowire.BytesType)
		site = protowire.AppendBytes(site, []byte(code))
		for _, d := range domains {
			var dom []byte
			dom = protowire.AppendTag(dom, 1, protowire.VarintType)
			dom = protowire.AppendVarint(dom, d.Type)
			dom = protowire.AppendTag(dom, 2, protowire.BytesType)
			dom = protowire.AppendBytes(dom, []byte(d.Value))
			site = protowire.AppendTag(site, 2, protowire.BytesType)
			site = protowire.AppendBytes(site, dom)
		}
		out = protowire.AppendTag(out, 1, protowire.BytesType)
		out = protowire.AppendBytes(out, site)
	}
	return out
}

func TestParseGeoSite(t *testing.T) {
	dir := t.TempDir()
	type dom = struct {
		Value string
		Type  uint64
	}
	dat := buildSiteDat(t, map[string][]dom{
		"telegram": {{"t.me", 2}, {"telegram.org", 2}, {"exact.telegram.org", 3}},
		"ads":      {{"keyword", 0}, {"re.*gex", 1}}, // both unsupported -> dropped
	})
	if err := os.WriteFile(filepath.Join(dir, "sites.site.dat"), dat, 0o644); err != nil {
		t.Fatal(err)
	}
	m, err := New(dir)
	if err != nil {
		t.Fatal(err)
	}

	tg, err := m.Domains("telegram")
	if err != nil || len(tg) != 3 {
		t.Fatalf("telegram domains: %v %v", tg, err)
	}
	// Plain/Regex-only category yields no usable domains.
	if _, err := m.Domains("ads"); err == nil {
		t.Fatal("expected ads (plain/regex only) to have no domains")
	}

	// Geosite files must not appear as geoip datasets, and vice versa.
	if ds, _ := m.Datasets(); len(ds) != 0 {
		t.Fatalf("geosite file leaked into geoip datasets: %+v", ds)
	}
	sds, err := m.SiteDatasets()
	if err != nil || len(sds) != 1 || sds[0].Name != "sites" {
		t.Fatalf("site datasets: %v %+v", err, sds)
	}
	sc, err := m.SiteCategories()
	if err != nil || len(sc) != 1 || sc[0].Code != "telegram" || sc[0].Domains != 3 {
		t.Fatalf("site categories: %v %+v", err, sc)
	}
}

func TestGeoipAndGeositeCoexist(t *testing.T) {
	dir := t.TempDir()
	os.WriteFile(filepath.Join(dir, "ip.dat"), buildDat(t, map[string][]string{"ru": {"2.56.0.0/16"}}), 0o644)
	os.WriteFile(filepath.Join(dir, "site.site.dat"), buildSiteDat(t, map[string][]struct {
		Value string
		Type  uint64
	}{"youtube": {{"youtube.com", 2}}}), 0o644)
	m, _ := New(dir)

	if ds, _ := m.Datasets(); len(ds) != 1 || ds[0].Name != "ip" {
		t.Fatalf("geoip datasets: %+v", ds)
	}
	if sds, _ := m.SiteDatasets(); len(sds) != 1 || sds[0].Name != "site" {
		t.Fatalf("site datasets: %+v", sds)
	}
	if _, err := m.CIDRs("ru"); err != nil {
		t.Fatalf("geoip lookup failed: %v", err)
	}
	if _, err := m.Domains("youtube"); err != nil {
		t.Fatalf("geosite lookup failed: %v", err)
	}
}

func TestParseAndExpand(t *testing.T) {
	dir := t.TempDir()
	dat := buildDat(t, map[string][]string{
		"CN": {"1.0.1.0/24", "1.0.2.0/23"},
		"RU": {"2.56.0.0/16"},
	})
	if err := os.WriteFile(filepath.Join(dir, "countries.dat"), dat, 0o644); err != nil {
		t.Fatal(err)
	}
	m, err := New(dir)
	if err != nil {
		t.Fatal(err)
	}

	cn, err := m.CIDRs("cn") // case-insensitive
	if err != nil {
		t.Fatal(err)
	}
	if len(cn) != 2 || cn[0] != "1.0.1.0/24" {
		t.Fatalf("got %v", cn)
	}
	if _, err := m.CIDRs("us"); err == nil {
		t.Fatal("expected error for missing category")
	}

	cats, err := m.Categories()
	if err != nil {
		t.Fatal(err)
	}
	if len(cats) != 2 {
		t.Fatalf("got %+v", cats)
	}
}

func TestMergeAcrossDatasets(t *testing.T) {
	dir := t.TempDir()
	os.WriteFile(filepath.Join(dir, "a.dat"), buildDat(t, map[string][]string{"cn": {"1.0.1.0/24"}}), 0o644)
	os.WriteFile(filepath.Join(dir, "b.dat"), buildDat(t, map[string][]string{"ru-blocked": {"2.56.0.0/16"}}), 0o644)
	m, _ := New(dir)

	ds, err := m.Datasets()
	if err != nil || len(ds) != 2 {
		t.Fatalf("datasets: %v %+v", err, ds)
	}
	blocked, err := m.CIDRs("ru-blocked")
	if err != nil || len(blocked) != 1 {
		t.Fatalf("got %v %v", blocked, err)
	}
}

func TestIsPublicIP(t *testing.T) {
	cases := map[string]bool{
		"8.8.8.8":         true,
		"1.1.1.1":         true,
		"2606:4700::1111": true,
		"127.0.0.1":       false, // loopback
		"10.0.0.5":        false, // RFC1918
		"192.168.1.1":     false,
		"172.16.0.1":      false,
		"169.254.169.254": false, // cloud metadata (link-local)
		"0.0.0.0":         false, // unspecified
		"::1":             false, // IPv6 loopback
		"fd00::1":         false, // IPv6 unique-local
		"fe80::1":         false, // IPv6 link-local
	}
	for s, want := range cases {
		if got := isPublicIP(net.ParseIP(s)); got != want {
			t.Errorf("isPublicIP(%s) = %v, want %v", s, got, want)
		}
	}
}

func TestAddDatasetRejects(t *testing.T) {
	m, err := New(t.TempDir())
	if err != nil {
		t.Fatal(err)
	}
	ctx := context.Background()

	// A server on loopback must be refused at dial time (SSRF guard).
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Write([]byte("should never be reached"))
	}))
	defer srv.Close()
	if err := m.AddDataset(ctx, "loop", srv.URL); err == nil {
		t.Error("expected loopback URL to be blocked, got nil")
	} else if !strings.Contains(err.Error(), "non-public") {
		t.Errorf("expected non-public block, got: %v", err)
	}

	// Non-http schemes are rejected before any dial.
	if err := m.AddDataset(ctx, "file", "file:///etc/passwd"); err == nil {
		t.Error("expected file:// scheme to be rejected")
	}
}
