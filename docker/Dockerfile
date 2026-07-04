# syntax=docker/dockerfile:1

# mieru/mita version. Keep in sync with go.mod (the Go module provides the
# gRPC stubs; this image provides the daemon - they must match). The canonical
# source is go.mod; `make docker` passes its version here automatically and
# `make check-version` guards against drift. Bump both with scripts/update-mieru.sh.
ARG MITA_VERSION=v3.34.0

# 1. Build the Vue SPA.
FROM node:26-alpine AS web
WORKDIR /web
COPY web/package.json web/package-lock.json ./
RUN npm ci
COPY web/ ./
RUN npm run build

# 2. Build the Go binary with the SPA embedded.
FROM golang:1.25-alpine AS go
WORKDIR /src
RUN apk add --no-cache git
COPY go.mod go.sum ./
RUN go mod download
COPY . .
# Bring in the freshly built SPA (vite.config outDir is ../internal/webfs/dist,
# which from the /web workdir resolves to /internal/webfs/dist).
COPY --from=web /internal/webfs/dist ./internal/webfs/dist
ARG VERSION=dev
RUN CGO_ENABLED=0 go build -trimpath -ldflags="-s -w -X main.version=${VERSION}" -o /mieru-web-ui ./cmd/mieru-web-ui
# The mieru client, used to chain this panel's egress through another panel.
RUN CGO_ENABLED=0 go build -trimpath -ldflags="-s -w" -o /mieru github.com/enfein/mieru/v3/cmd/mieru

# 3. Final image: mita daemon + panel in one container.
FROM ghcr.io/enfein/mita:${MITA_VERSION}
COPY --from=go /mieru-web-ui /usr/local/bin/mieru-web-ui
COPY --from=go /mieru /usr/local/bin/mieru

# The panel supervises `mita run`; run as the built-in mita user and skip
# the socket chown that requires the mita user/group to own /var/run/mita.
ENV MITA_INSECURE_UDS=true \
    PANEL_PORT=8686 \
    PANEL_DATA_DIR=/data \
    MITA_BINARY=/usr/local/bin/mita \
    MIERU_BINARY=/usr/local/bin/mieru \
    MITA_UDS_PATH=/var/run/mita/mita.sock

EXPOSE 8686
VOLUME ["/data", "/etc/mita"]

# Works whether the panel serves HTTP or native TLS.
HEALTHCHECK --interval=30s --timeout=5s --start-period=10s --retries=3 \
    CMD wget -qO- "http://127.0.0.1:${PANEL_PORT}/healthz" >/dev/null 2>&1 \
     || wget -qO- --no-check-certificate "https://127.0.0.1:${PANEL_PORT}/healthz" >/dev/null 2>&1 \
     || exit 1

ENTRYPOINT ["/usr/local/bin/mieru-web-ui"]
