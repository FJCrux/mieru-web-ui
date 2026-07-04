#!/bin/sh
set -e

/usr/local/bin/mita run &
MITA_PID=$!

# Wait for the management socket, then bridge it to TCP :9090.
for i in $(seq 1 50); do
  [ -S /var/run/mita/mita.sock ] && break
  sleep 0.2
done
if [ ! -S /var/run/mita/mita.sock ]; then
  echo "mita.sock did not appear" >&2
  exit 1
fi

socat TCP-LISTEN:9090,fork,reuseaddr UNIX-CONNECT:/var/run/mita/mita.sock &

wait $MITA_PID
