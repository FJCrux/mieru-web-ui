# mieru-web-ui

[![CI](https://github.com/fjcrux/mieru-web-ui/actions/workflows/ci.yml/badge.svg)](https://github.com/fjcrux/mieru-web-ui/actions/workflows/ci.yml)
[![License: GPL v3](https://img.shields.io/badge/License-GPLv3-blue.svg)](LICENSE)

A web admin panel for the [mieru](https://github.com/enfein/mieru) proxy server
(`mita`), in the spirit of 3x-ui for Xray. Manage users, ports, routing, and
share links from the browser.

Single Go binary with the Vue SPA embedded, shipped as one Docker image that
bundles and supervises `mita`.

## Features

- **Users & quotas** - CRUD mieru users, traffic quotas, private-IP toggle.
- **Network** - port bindings (single ports and ranges, TCP/UDP), MTU, logging.
- **Dashboard** - mita status, live sessions, traffic, logs; restart from the UI.
- **Sharing** - per-user `mieru://` / `mierus://` links, QR codes, and expiring
  tokenized links on a separate path. See [sharing](docs/guides/sharing-configs.md).
- **Outbound routing & GeoIP** - SOCKS5 outbounds and rules by domain, CIDR, or
  GeoIP category (xray `geoip.dat`). See [routing](docs/guides/geoip-routing.md).
- **Panel chaining** - route egress through another panel with a pasted key.
  See [chaining](docs/guides/chaining-panels.md).
- **Metrics** - OpenTelemetry (OTLP / Prometheus). See [metrics](docs/guides/metrics.md).
- **Hardening** - TLS, secret base path, Host validation, rate-limited login.
  See [hardening](docs/guides/hardening.md).

## Quick start (Docker)

```bash
git clone https://github.com/fjcrux/mieru-web-ui
cd mieru-web-ui
cp .env.example .env        # edit as needed
docker compose up -d        # pulls ghcr.io/fjcrux/mieru-web-ui:latest
docker compose logs | grep -A4 "Panel admin"   # generated password on first run
```

To build from source instead:

```bash
docker compose -f docker-compose.yml -f docker-compose.dev.yml up -d --build
```

The compose file publishes the panel on `127.0.0.1:8686` and the proxy ports in
bridge mode. Reach the panel via an SSH tunnel (or a
[reverse proxy](docs/guides/reverse-proxy-tls.md)):

```bash
ssh -L 8686:127.0.0.1:8686 user@your-server   # then open http://localhost:8686
```

First login is `admin` with the printed password. Change it under **Settings**.

### First-time setup

1. **Settings**: set the Public host (address clients connect to) and, behind a
   proxy, the Panel URL.
2. **Network**: add a port binding (open it in your firewall / `PROXY_PORTS`).
3. **Users**: add a user; the proxy starts once there's a user and a port.
4. **Users → Share**: hand out the link or QR.

## Configuration

All options live in `.env` - see [`.env.example`](.env.example) for the full,
commented list (ports, admin, TLS, panel URL, base path, share path, GeoIP,
metrics). `docker compose` reads `.env` automatically.

## How it works

See [docs/guides/how-it-works.md](docs/guides/how-it-works.md) for the mental
model: one container supervising mita (and any chained clients), ports driven by
a single `PROXY_PORTS`, and the admin/share paths.

## Guides

Task-oriented how-tos are in [docs/guides](docs/guides/): how it works,
TLS/reverse-proxy, hardening, routing & GeoIP, chaining panels, sharing, metrics.

## Updating mieru/mita

```bash
scripts/update-mieru.sh v3.35.0   # bumps go.mod, Dockerfiles, .env, compose
make test && make docker          # verify and rebuild
```

## Releases

Pushing a `vX.Y.Z` tag builds and publishes the image to
`ghcr.io/fjcrux/mieru-web-ui` (`X.Y.Z`, `X.Y`, `latest`) and creates a GitHub
Release. Nothing is published from branch pushes.

## Security

See [SECURITY.md](SECURITY.md). Key points: serve over TLS; the panel stores
mieru user passwords in plaintext (mode 0600) to build client configs, so
protect the `/data` volume; treat share links as secrets.

## Development

See [docs/DEVELOPMENT.md](docs/DEVELOPMENT.md).

## License

GPL-3.0 - mieru-web-ui imports the GPL-3.0 `mieru` module. See [LICENSE](LICENSE).
