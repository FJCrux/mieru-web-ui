# Contributing

Thanks for your interest in mieru-web-ui.

## Development setup

See [docs/DEVELOPMENT.md](docs/DEVELOPMENT.md) for the full loop. In short:

```bash
make dev-mita     # real mita in Docker, socket bridged to tcp://localhost:9090
# backend (host):
MITA_RPC_TARGET=tcp://localhost:9090 go run ./cmd/mieru-web-ui --no-supervise
# frontend (host):
cd web && npm install && npm run dev
```

## Before opening a PR

```bash
make check-version   # mieru version pins in sync
go vet ./... && go test ./...
cd web && npm run build   # SPA type-checks and builds
```

CI runs the same checks plus a full Docker build.

## Bumping the mieru version

Don't edit the version pins by hand, they live in four places that must
agree. Use the helper:

```bash
scripts/update-mieru.sh v3.35.0
```

Then run the probe against the new mita to confirm it still works with the
panel:

```bash
make dev-mita
MITA_RPC_TARGET=tcp://localhost:9090 PROBE_PORT_HOST=127.0.0.1 go run ./cmd/probe
```

## Code layout

`internal/mitaclient` is the only package that imports mieru's appctl types.
Keep it that way; everything else talks to it through panel-defined types.

## License

By contributing you agree that your contributions are licensed under GPL-3.0.
