#!/usr/bin/env bash
# Bump the pinned mieru/mita version everywhere it lives, keeping go.mod (the
# Go module / gRPC stubs) in sync with the Docker base image (the daemon).
#
# Usage: scripts/update-mieru.sh v3.35.0
set -euo pipefail

VERSION="${1:-}"
if [[ -z "$VERSION" ]]; then
  echo "usage: $0 vX.Y.Z   (e.g. $0 v3.35.0)" >&2
  exit 1
fi
if [[ "$VERSION" != v*.*.* ]]; then
  echo "error: version must look like vX.Y.Z (with the leading v)" >&2
  exit 1
fi

cd "$(dirname "$0")/.."

echo "==> Updating Go module to $VERSION"
go get "github.com/enfein/mieru/v3@$VERSION"
go mod tidy

echo "==> Syncing Dockerfile / .env.example version pins"
# perl -i is portable across GNU and BSD (unlike sed -i).
perl -pi -e "s/^(ARG MITA_VERSION=)v[0-9]+\.[0-9]+\.[0-9]+/\${1}$VERSION/" \
  docker/Dockerfile docker/dev/Dockerfile
perl -pi -e "s/^(MITA_VERSION=)v[0-9]+\.[0-9]+\.[0-9]+/\${1}$VERSION/" .env.example
# Sync the compose fallback default too.
perl -pi -e "s/(MITA_VERSION:-)v[0-9]+\.[0-9]+\.[0-9]+/\${1}$VERSION/" docker-compose.yml

echo "==> Verifying consistency"
make check-version

cat <<EOF

Version bumped to $VERSION.

IMPORTANT: re-verify runtime behaviour against the new mita. The appctl gRPC
API has no formal stability guarantee, so a version bump can silently change
what the panel assumes (config replace semantics, password round-trips,
port reload):

  make dev-mita
  MITA_RPC_TARGET=tcp://localhost:9090 PROBE_PORT_HOST=127.0.0.1 go run ./cmd/probe

Then run the full checks:

  make test && make docker
EOF
