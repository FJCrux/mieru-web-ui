# Metrics

The panel exports OpenTelemetry metrics: `mieru_mita_up`, `mieru_users`,
`mieru_sessions_active`, `mieru_traffic_download_bytes` /
`mieru_traffic_upload_bytes`, `mieru_peers_total` / `mieru_peers_connected`,
`mieru_mita_restarts`, and `mieru_http_requests_total`.

Metrics are off until you configure an exporter.

## OTLP (OpenObserve and other OTLP backends)

Set the standard OTEL variables; OpenObserve ingests them directly, no collector
needed:

```env
OTEL_EXPORTER_OTLP_ENDPOINT=https://openobserve.example.com/api/default
OTEL_EXPORTER_OTLP_HEADERS=Authorization=Basic <base64(user:token)>
OTEL_EXPORTER_OTLP_PROTOCOL=http/protobuf
```

The exporter posts to `<endpoint>/v1/metrics`, which is where OpenObserve expects
OTLP HTTP metrics for the given organization.

## Prometheus

Scrape the panel directly:

```env
PANEL_METRICS_PROMETHEUS=true
```

Then read `GET /metrics`. It stays at the root path (not behind
`PANEL_BASE_PATH`), so scrapers reach it regardless of the panel prefix.
