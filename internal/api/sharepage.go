// SPDX-License-Identifier: GPL-3.0-or-later

package api

import (
	"encoding/base64"
	"html/template"
	"net/http"

	qrcode "github.com/skip2/go-qrcode"

	"github.com/fjcrux/mieru-web-ui/internal/sharelink"
)

var sharePageTmpl = template.Must(template.New("share").Parse(`<!doctype html>
<html lang="en"><head>
<meta charset="utf-8">
<meta name="viewport" content="width=device-width, initial-scale=1">
<meta name="robots" content="noindex,nofollow">
<title>mieru connection</title>
<style>
:root{color-scheme:dark}
body{margin:0;background:#101014;color:#e8e8ea;font-family:-apple-system,BlinkMacSystemFont,'Segoe UI',Roboto,sans-serif;
 display:flex;min-height:100vh;align-items:center;justify-content:center;padding:16px}
.card{background:#18181c;border:1px solid #2a2a30;border-radius:12px;padding:24px;max-width:420px;width:100%}
h1{font-size:18px;margin:0 0 16px}
img{display:block;width:260px;height:260px;margin:0 auto 16px;background:#fff;border-radius:8px;padding:8px;box-sizing:border-box}
textarea{width:100%;box-sizing:border-box;background:#0d0d10;color:#c8c8cc;border:1px solid #2a2a30;border-radius:6px;
 padding:8px;font-family:ui-monospace,monospace;font-size:12px;resize:vertical}
button{margin-top:8px;background:#63e2b7;color:#08130f;border:0;border-radius:6px;padding:8px 14px;font-weight:600;cursor:pointer}
.hint{margin-top:14px;font-size:12px;color:#8a8a90}
</style></head><body>
<div class="card">
<h1>Connect: {{.User}}</h1>
<img src="data:image/png;base64,{{.QR}}" alt="QR code">
<textarea id="l" rows="3" readonly>{{.Link}}</textarea>
<button onclick="navigator.clipboard.writeText(document.getElementById('l').value)">Copy link</button>
<p class="hint">Import this into your mieru client. Keep it private.</p>
</div></body></html>`))

func renderSharePage(w http.ResponseWriter, user string, links *sharelink.Links) {
	link := ""
	if len(links.MierusURLs) > 0 {
		link = links.MierusURLs[0]
	} else {
		link = links.MieruURL
	}
	png, err := qrcode.Encode(link, qrcode.Medium, 512)
	if err != nil {
		http.Error(w, "failed to render", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("X-Robots-Tag", "noindex")
	_ = sharePageTmpl.Execute(w, struct {
		User string
		QR   string
		Link string
	}{
		User: user,
		QR:   base64.StdEncoding.EncodeToString(png),
		Link: link,
	})
}
