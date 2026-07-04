// SPDX-License-Identifier: GPL-3.0-or-later

// Package webfs serves the embedded SPA build with an index.html fallback for
// client-side routes. It injects the base path so the panel can live under a
// secret prefix.
package webfs

import (
	"bytes"
	"embed"
	"io/fs"
	"net/http"
	"strings"
)

//go:embed all:dist
var dist embed.FS

// Handler serves the SPA. basePath is the panel's URL prefix ("" for root,
// e.g. "/panel"); it is injected into index.html as <base> and window.__BASE__
// so the SPA's router and API calls resolve under it. Unknown non-asset paths
// fall back to index.html so vue-router deep links work.
func Handler(basePath string) http.Handler {
	sub, err := fs.Sub(dist, "dist")
	if err != nil {
		panic(err)
	}
	index := buildIndex(sub, basePath)
	fileServer := http.FileServer(http.FS(sub))

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path := strings.TrimPrefix(r.URL.Path, "/")
		if path != "" && path != "index.html" {
			if f, err := sub.Open(path); err == nil {
				f.Close()
				fileServer.ServeHTTP(w, r)
				return
			}
		}
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Write(index)
	})
}

// buildIndex reads index.html and injects the base tag and window.__BASE__.
func buildIndex(sub fs.FS, basePath string) []byte {
	raw, err := fs.ReadFile(sub, "index.html")
	if err != nil {
		panic(err)
	}
	inject := []byte("<base href=\"" + basePath + "/\">\n" +
		"<script>window.__BASE__=" + jsString(basePath) + "</script>\n")
	if i := bytes.Index(raw, []byte("<head>")); i >= 0 {
		out := make([]byte, 0, len(raw)+len(inject))
		out = append(out, raw[:i+len("<head>")]...)
		out = append(out, '\n')
		out = append(out, inject...)
		out = append(out, raw[i+len("<head>"):]...)
		return out
	}
	return append(inject, raw...)
}

func jsString(s string) string {
	return "\"" + strings.ReplaceAll(s, "\"", "\\\"") + "\""
}
