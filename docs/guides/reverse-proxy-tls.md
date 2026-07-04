# Reverse proxy and TLS

Always reach the panel over HTTPS. Two options.

## Option A: nginx in front (recommended)

Keep the panel bound to loopback (the default compose publishes
`127.0.0.1:8686`) and terminate TLS in nginx.

```nginx
server {
    listen 443 ssl;
    server_name vpn.example.com;

    ssl_certificate     /etc/letsencrypt/live/vpn.example.com/fullchain.pem;
    ssl_certificate_key /etc/letsencrypt/live/vpn.example.com/privkey.pem;

    location / {
        proxy_pass http://127.0.0.1:8686;
        proxy_set_header Host $host;
        proxy_set_header X-Forwarded-Proto $scheme;
        proxy_set_header X-Forwarded-Host $host;
    }
}
```

Then set **Panel URL** in Settings to `https://vpn.example.com`. The panel uses
it for share links and rejects requests with any other Host.

If you run the panel under a [secret base path](hardening.md), proxy only that
path (e.g. `location /a7Fq2xK/ { ... }`) so nothing answers at `/`.

## Option B: built-in TLS (no proxy)

Mount a certificate and key and point the panel at them:

```env
PANEL_TLS_CERT=/data/tls/cert.pem
PANEL_TLS_KEY=/data/tls/key.pem
```

Put the files under the `/data` volume. The panel then serves HTTPS directly;
publish its port (e.g. `443:8686`) instead of binding to loopback.

## Why it matters

Over plain HTTP the admin password, session cookie, and client credentials
travel in the clear. The dashboard shows a warning when it detects plain HTTP on
a non-loopback address.
