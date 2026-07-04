# Running the panel and the proxy on separate hosts

By default the panel and `mita` live in one container. You can also run the
panel on one host and `mita` on another: the panel talks to `mita` over gRPC,
and that target is configurable.

## How

Run the panel with `--no-supervise` (it won't start its own `mita`) and point it
at the remote daemon:

```yaml
services:
  mieru-web-ui:
    image: ghcr.io/fjcrux/mieru-web-ui:latest
    command: ["--no-supervise"]
    environment:
      MITA_RPC_TARGET: "tcp://10.8.0.2:9090"   # the mita host, over a private net
    ports:
      - "127.0.0.1:8686:8686"
    volumes:
      - ./panel-data:/data
```

On the `mita` host, run `mita` (its official image or a systemd install) and
expose its management socket as TCP for the panel to reach. mita only listens on
a unix socket, so bridge it, e.g. with socat:

```bash
socat TCP-LISTEN:9090,fork,reuseaddr UNIX-CONNECT:/var/run/mita/mita.sock
```

## Security: this gRPC is unauthenticated

mita's management API has no auth or TLS. **Never expose the bridged port to the
internet.** Put the panel and mita on a private network you control:

- a WireGuard/Tailscale tunnel between the hosts (target the tunnel IP), or
- an SSH tunnel (`ssh -L 9090:127.0.0.1:9090 mita-host`), or
- an mTLS proxy in front of the socket.

## What works, and what doesn't

Over the remote gRPC the panel manages **users, ports, routing/GeoIP, metrics,
dashboard, and sharing** normally.

**Chaining panels does not work split.** A chained upstream runs a `mieru`
client next to `mita` and exposes a local SOCKS5 port that `mita`'s egress dials
at `127.0.0.1`. If the panel is on a different host than `mita`, that loopback
address isn't the panel's client. Use chaining only in the single-host setup, or
run the client on the `mita` host yourself and add it as a plain SOCKS5 outbound.
