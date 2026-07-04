# How it works

A quick mental model of the moving parts.

## One container, three processes

The image runs the panel as PID 1. The panel supervises `mita` (the mieru
server) as a child process and, when you chain to another panel, one `mieru`
client per upstream. If `mita` crashes, the panel restarts it with backoff. The
panel talks to `mita` over its local gRPC socket to read and write config.

State lives in the `/data` volume (a SQLite DB with the admin account,
sessions, per-user secrets, and settings) and `/etc/mita` (mita's own config).

## Ports

mieru needs at least one proxy port. One is enough; a range enables port-hopping
(spreading traffic across ports, harder to block).

You manage ports in one place: `PROXY_PORTS`. The compose file uses it both to
publish the ports and to tell the panel what to enforce, so the published ports
and mita always agree. On every start the panel applies `PROXY_PORTS` to mita,
and the Network page shows them read-only. If you don't set `PROXY_PORTS`, the
panel seeds a single default port on first run and you manage ports from the
Network page instead.

## Addresses and paths

- **PANEL_BIND_IP** (compose): which host IP the panel port is published on.
  Default `127.0.0.1` (loopback only). Change it to expose the panel, but only
  behind TLS.
- **Public host** (Settings): the address clients connect to, embedded in client
  configs and share links.
- **Panel URL** (Settings): the panel's own external address. It's the base for
  share links and the allowed `Host` - requests with a different Host are
  rejected.
- **Base path** (Settings; restart the panel container to apply): serves the whole admin panel under
  a secret prefix; everything else returns 404.
- **Share path** (Settings; restart the panel container to apply): the public prefix for share links,
  kept separate from the base path so a shared link never reveals it.

## Request flow

```
client  ->  mita (proxy ports)  ->  internet   (or -> egress proxy / chained panel)
browser ->  panel  ->  mita gRPC  (manage users, ports, routing)
recipient -> panel /s/<token>     (server-rendered share page, no login)
```

## First run

1. The panel creates an admin account (random password in the logs unless you
   set `PANEL_ADMIN_PASSWORD`).
2. It seeds a default port binding (or applies `PROXY_PORTS`).
3. You set the Public host, add a user, and share the link. The proxy starts
   once there's at least one user and one port.

The Dashboard shows hardening hints (no TLS, no Panel URL, default path/port) so
you know what's left to lock down.
