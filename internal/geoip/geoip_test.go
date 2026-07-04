// SPDX-License-Identifier: GPL-3.0-or-later

package geoip

import (
	"net"
	"os"
	"path/filepath"
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
