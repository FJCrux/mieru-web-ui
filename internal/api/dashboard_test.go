// SPDX-License-Identifier: GPL-3.0-or-later

package api

import (
	"net/http/httptest"
	"testing"
)

func TestIsLoopback(t *testing.T) {
	cases := map[string]bool{
		"localhost:8686":    true,
		"127.0.0.1:8686":    true,
		"127.0.0.1":         true,
		"[::1]:8686":        true,
		"203.0.113.10:8686": false,
		"vpn.example.com":   false,
		"10.0.0.5:8686":     false,
	}
	for host, want := range cases {
		r := httptest.NewRequest("GET", "/api/dashboard", nil)
		r.Host = host
		if got := isLoopback(r); got != want {
			t.Errorf("isLoopback(%q) = %v, want %v", host, got, want)
		}
	}
}
