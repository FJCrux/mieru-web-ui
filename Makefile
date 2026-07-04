.PHONY: build build-web build-go test check-version docker \
        up dev down logs dev-mita clean

# Single source of truth: the mieru module version pinned in go.mod.
# `make docker` passes it as the mita base-image tag so they cannot diverge.
MITA_VERSION := $(shell awk '/enfein\/mieru\/v3/ {print $$2}' go.mod)

COMPOSE     := docker compose
COMPOSE_DEV := docker compose -f docker-compose.yml -f docker-compose.dev.yml

## --- run (Docker Compose) ---

# Production: pull the published image and start (as in the README).
up:
	$(COMPOSE) up -d

# Development: build the image from source and start.
dev:
	$(COMPOSE_DEV) up -d --build

down:
	$(COMPOSE) down

logs:
	$(COMPOSE) logs -f

## --- build ---

build: build-web build-go

build-web:
	cd web && npm ci && npm run build

build-go:
	CGO_ENABLED=0 go build -trimpath -ldflags="-s -w" -o mieru-web-ui ./cmd/mieru-web-ui

test:
	go test ./...

# Build the production image directly (used by CI is the compose/publish path).
docker: check-version
	docker build -f docker/Dockerfile --build-arg MITA_VERSION=$(MITA_VERSION) -t mieru-web-ui .

# Fail if the pinned mieru version drifts between go.mod, the Dockerfiles, and
# .env.example. Run in CI; fix mismatches with scripts/update-mieru.sh.
check-version:
	@go_ver=$$(awk '/enfein\/mieru\/v3/ {print $$2}' go.mod); \
	 df_ver=$$(awk -F= '/^ARG MITA_VERSION=/ {print $$2}' docker/Dockerfile); \
	 dev_ver=$$(awk -F= '/^ARG MITA_VERSION=/ {print $$2}' docker/dev/Dockerfile); \
	 env_ver=$$(awk -F= '/^MITA_VERSION=/ {print $$2}' .env.example); \
	 echo "go.mod=$$go_ver Dockerfile=$$df_ver dev=$$dev_ver .env.example=$$env_ver"; \
	 if [ "$$go_ver" != "$$df_ver" ] || [ "$$go_ver" != "$$dev_ver" ] || [ "$$go_ver" != "$$env_ver" ]; then \
	   echo "ERROR: mieru version mismatch - run scripts/update-mieru.sh $$go_ver"; exit 1; \
	 fi

## --- host development (run the Go panel on the host) ---

# Runs a real mita in Docker with its management socket bridged to
# tcp://localhost:9090 (Docker Desktop on macOS can't share unix sockets with
# the host). Then run the panel binary on the host against it:
#   MITA_RPC_TARGET=tcp://localhost:9090 go run ./cmd/mieru-web-ui --no-supervise
# See docs/DEVELOPMENT.md.
dev-mita:
	docker build -f docker/dev/Dockerfile --build-arg MITA_VERSION=$(MITA_VERSION) -t mita-dev docker
	docker rm -f mita-dev 2>/dev/null || true
	docker run -d --name mita-dev \
		-p 127.0.0.1:9090:9090 -p 127.0.0.1:2012-2014:2012-2014 \
		mita-dev

clean:
	rm -f mieru-web-ui
	rm -rf web/node_modules internal/webfs/dist
