# Hardening

Defense in depth for the panel. None of these replace TLS and a strong admin
password - they reduce exposure.

## Hide the panel behind a secret path

By default the panel answers at `/`. Set a secret prefix and everything else
returns 404, so scanners and bots hitting `/` find nothing:

```env
PANEL_BASE_PATH=/a7Fq2xK9
```

The panel is then served only under `https://host/a7Fq2xK9/`. Use an
unguessable value - a fixed word like `/panel` adds nothing. You can also set
this in Settings; either way, restart the panel container (`docker compose
restart`) to apply it.

Share links are **not** placed under this path (see below), so sending a link
never reveals it.

## Keep share links on a separate path

Public share pages live under their own prefix (default `/s`), independent of
the admin path:

```env
SHARE_PATH=/s
```

A share link looks like `https://host/s/<token>` and reveals nothing about where
the admin panel lives. The token is unguessable and the link expires.

## Restrict the Host

Set **Panel URL** (Settings, or `PANEL_URL`) to your real address. The panel
then rejects any request whose `Host` header doesn't match (loopback always
allowed, so health checks work). This blocks Host-header injection and stray
access by raw IP.

## Bind to loopback

The default compose publishes the panel on `127.0.0.1` only. Reach it through an
SSH tunnel (`ssh -L 8686:127.0.0.1:8686 user@server`) or a reverse proxy. Only
expose it publicly behind TLS.

## Rotate the bootstrap password

On first run the panel prints a generated admin password (unless you set
`PANEL_ADMIN_PASSWORD`). Change it under Settings; changing credentials
invalidates every session.
