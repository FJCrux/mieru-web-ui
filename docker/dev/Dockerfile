# Development-only image: real mita + socat bridge exposing the management
# unix socket as TCP :9090, because Docker Desktop on macOS cannot share
# unix sockets with the host. Production image lives in docker/Dockerfile.
#
# MITA_VERSION must match go.mod; see scripts/update-mieru.sh + make check-version.
ARG MITA_VERSION=v3.34.0
FROM ghcr.io/enfein/mita:${MITA_VERSION}

USER root
RUN apk add --no-cache socat

COPY dev/entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh

ENV MITA_INSECURE_UDS=true
ENTRYPOINT ["/entrypoint.sh"]
