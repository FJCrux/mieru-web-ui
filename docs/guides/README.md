# Guides

Task-oriented how-tos. Start with the [main README](../../README.md) for
install and configuration.

- [How it works](how-it-works.md) - the moving parts: container, ports, paths,
  request flow.
- [Split deployment](split-deployment.md) - run the panel and the proxy on
  separate hosts.

- [Reverse proxy and TLS](reverse-proxy-tls.md) - serve the panel over HTTPS,
  with nginx or the panel's built-in TLS.
- [Hardening](hardening.md) - hide the panel behind a secret path, restrict the
  Host, lock down access.
- [Outbound routing and GeoIP](geoip-routing.md) - keep local traffic direct and
  proxy the rest, or proxy only blocked ranges.
- [Chaining two panels](chaining-panels.md) - route one server's egress through
  another.
- [Sharing client configs](sharing-configs.md) - hand out connection links and
  QR codes.
- [Metrics](metrics.md) - export to OpenObserve or Prometheus.
