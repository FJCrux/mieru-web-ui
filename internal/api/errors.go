// SPDX-License-Identifier: GPL-3.0-or-later

package api

import (
	"net/http"
	"strings"
)

// mitaMessage turns a raw gRPC/mita error into something a user can read:
// it drops the "rpc error: code = ... desc = " wrapper and maps the common
// "not set up yet" states to plain language.
func mitaMessage(err error) string {
	s := err.Error()
	if i := strings.LastIndex(s, "desc = "); i >= 0 {
		s = s[i+len("desc = "):]
	}
	low := strings.ToLower(s)
	switch {
	case strings.Contains(low, "multiplex") && strings.Contains(low, "unavailable"):
		return "the proxy is not running yet — add a user and a port first"
	case strings.Contains(low, "config is empty"):
		return "no server configuration yet — add a port and a user"
	case strings.Contains(low, "connection refused"), strings.Contains(low, "transport:"):
		return "cannot reach the mita daemon"
	}
	return s
}

// writeMitaErr reports a mita failure with a cleaned-up message.
func writeMitaErr(w http.ResponseWriter, err error) {
	writeErr(w, http.StatusBadGateway, mitaMessage(err))
}

// proxyNotReady reports the "nothing configured / proxy stopped" state, so read
// endpoints can return an empty result instead of a scary error on a fresh
// install.
func proxyNotReady(err error) bool {
	low := strings.ToLower(err.Error())
	return strings.Contains(low, "multiplex") ||
		strings.Contains(low, "config is empty") ||
		strings.Contains(low, "not running")
}
