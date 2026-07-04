# Development

## Prerequisites

- Go 1.25+
- Node.js 20+
- Docker (for a real `mita` to talk to)

`mita` exposes its management API over a unix socket. For local development
`docker/dev/Dockerfile` runs `mita` with a `socat` bridge that republishes the
socket as TCP `:9090`, so the panel (running on the host) can reach it with
`MITA_RPC_TARGET=tcp://localhost:9090`. On Linux you can instead volume-mount
the socket directory and use the default `unix://` target.

## Run the stack

```bash
# 1. Start mita with the management socket bridged to tcp://localhost:9090
make dev-mita

# 2. Backend (host), pointed at the bridge, not supervising its own mita
PANEL_DATA_DIR=./.devdata \
PANEL_ADMIN_USER=admin PANEL_ADMIN_PASSWORD=devpass \
MITA_RPC_TARGET=tcp://localhost:9090 \
go run ./cmd/mieru-web-ui --no-supervise

# 3. Frontend dev server (proxies /api to :8686)
cd web && npm install && npm run dev
# open http://localhost:5173
```

`--no-supervise` tells the panel not to spawn its own `mita`; it talks to the
dockerized one instead. In production (the single-container image) the flag is
omitted and the panel supervises `mita run` itself.

## Build the embedded binary

```bash
make build        # builds the SPA, embeds it, produces ./mieru-web-ui
make test         # Go unit tests
make docker       # production image
```

## The probe tool

`cmd/probe` runs a set of read/write calls against a mita daemon to check it
still behaves as the panel expects. Run it against a dockerized mita after a
mieru version bump.

```bash
MITA_RPC_TARGET=tcp://localhost:9090 PROBE_PORT_HOST=127.0.0.1 go run ./cmd/probe
```

## Layout

- `cmd/mieru-web-ui` - entrypoint, wiring, PID-1 signal handling
- `internal/mitaclient` - gRPC wrapper; the only package that imports mieru
  appctl types, isolating version churn
- `internal/supervisor` - mita child lifecycle
- `internal/store` - SQLite persistence
- `internal/auth` - panel authentication
- `internal/api` - REST handlers
- `internal/sharelink` - client config + share link generation
- `internal/webfs` - embedded SPA
- `web/` - Vue 3 + Vite SPA
