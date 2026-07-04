// SPDX-License-Identifier: GPL-3.0-or-later

// Package metrics exposes panel and mita metrics via OpenTelemetry.
//
// By default it pushes over OTLP (works with OpenObserve and any OTLP
// backend) using the standard OTEL_EXPORTER_OTLP_* environment variables, so
// no code config is needed. A Prometheus /metrics endpoint can be enabled too.
package metrics

import (
	"context"
	"net/http"
	"os"
	"strconv"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetrichttp"
	otelprom "go.opentelemetry.io/otel/exporters/prometheus"
	"go.opentelemetry.io/otel/metric"
	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	semconv "go.opentelemetry.io/otel/semconv/v1.26.0"
)

// Snapshot is the current panel/mita state sampled on each collection.
type Snapshot struct {
	MitaUp         int64
	Users          int64
	Sessions       int64
	DownloadBytes  int64
	UploadBytes    int64
	PeersTotal     int64
	PeersConnected int64
	MitaRestarts   int64
}

// Collector samples the current state. It is called on each metrics export.
type Collector func(ctx context.Context) Snapshot

// Metrics holds the running exporter pipeline.
type Metrics struct {
	provider   *sdkmetric.MeterProvider
	prom       http.Handler
	reqCounter metric.Int64Counter
}

// Enabled reports whether any exporter is configured via the environment.
func Enabled() bool {
	return os.Getenv("OTEL_EXPORTER_OTLP_ENDPOINT") != "" ||
		os.Getenv("OTEL_EXPORTER_OTLP_METRICS_ENDPOINT") != "" ||
		os.Getenv("PANEL_METRICS_PROMETHEUS") == "true"
}

// Setup builds the meter provider with an OTLP exporter (when an OTLP endpoint
// is set) and/or a Prometheus exporter (when PANEL_METRICS_PROMETHEUS=true),
// and registers the observable gauges backed by collector.
func Setup(ctx context.Context, version string, collector Collector) (*Metrics, error) {
	res, err := resource.New(ctx,
		resource.WithAttributes(
			semconv.ServiceName("mieru-web-ui"),
			semconv.ServiceVersion(version),
		),
	)
	if err != nil {
		return nil, err
	}

	var readers []sdkmetric.Reader
	if os.Getenv("OTEL_EXPORTER_OTLP_ENDPOINT") != "" || os.Getenv("OTEL_EXPORTER_OTLP_METRICS_ENDPOINT") != "" {
		// Reads OTEL_EXPORTER_OTLP_* from the environment (endpoint, headers,
		// protocol) — the OpenObserve OTLP HTTP endpoint + Authorization header
		// is enough.
		exp, err := otlpmetrichttp.New(ctx)
		if err != nil {
			return nil, err
		}
		readers = append(readers, sdkmetric.NewPeriodicReader(exp))
	}

	var promHandler http.Handler
	if os.Getenv("PANEL_METRICS_PROMETHEUS") == "true" {
		reg := prometheus.NewRegistry()
		promExp, err := otelprom.New(otelprom.WithRegisterer(reg))
		if err != nil {
			return nil, err
		}
		readers = append(readers, promExp)
		promHandler = promhttp.HandlerFor(reg, promhttp.HandlerOpts{})
	}

	opts := []sdkmetric.Option{sdkmetric.WithResource(res)}
	for _, r := range readers {
		opts = append(opts, sdkmetric.WithReader(r))
	}
	provider := sdkmetric.NewMeterProvider(opts...)
	otel.SetMeterProvider(provider)

	m := &Metrics{provider: provider, prom: promHandler}
	if err := registerGauges(provider.Meter("mieru-web-ui"), collector, m); err != nil {
		return nil, err
	}
	return m, nil
}

func registerGauges(meter metric.Meter, collector Collector, m *Metrics) error {
	g := func(name, desc string) metric.Int64ObservableGauge {
		gauge, _ := meter.Int64ObservableGauge(name, metric.WithDescription(desc))
		return gauge
	}
	up := g("mieru.mita.up", "1 if the mita proxy is running")
	users := g("mieru.users", "number of mieru users")
	sessions := g("mieru.sessions.active", "active proxy sessions")
	down := g("mieru.traffic.download_bytes", "total downloaded bytes")
	upBytes := g("mieru.traffic.upload_bytes", "total uploaded bytes")
	peersTotal := g("mieru.peers.total", "configured upstream chain peers")
	peersUp := g("mieru.peers.connected", "connected upstream chain peers")
	restarts := g("mieru.mita.restarts", "mita restart count")

	_, err := meter.RegisterCallback(func(ctx context.Context, o metric.Observer) error {
		s := collector(ctx)
		o.ObserveInt64(up, s.MitaUp)
		o.ObserveInt64(users, s.Users)
		o.ObserveInt64(sessions, s.Sessions)
		o.ObserveInt64(down, s.DownloadBytes)
		o.ObserveInt64(upBytes, s.UploadBytes)
		o.ObserveInt64(peersTotal, s.PeersTotal)
		o.ObserveInt64(peersUp, s.PeersConnected)
		o.ObserveInt64(restarts, s.MitaRestarts)
		return nil
	}, up, users, sessions, down, upBytes, peersTotal, peersUp, restarts)
	if err != nil {
		return err
	}

	m.reqCounter, err = meter.Int64Counter("mieru.http.requests", metric.WithDescription("panel HTTP requests"))
	return err
}

// PromHandler returns the Prometheus /metrics handler, or nil if disabled.
func (m *Metrics) PromHandler() http.Handler { return m.prom }

// Middleware counts HTTP requests by status code.
func (m *Metrics) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sw := &statusWriter{ResponseWriter: w, status: http.StatusOK}
		next.ServeHTTP(sw, r)
		m.reqCounter.Add(r.Context(), 1, metric.WithAttributes(
			attribute.String("code", strconv.Itoa(sw.status)),
		))
	})
}

// Shutdown flushes and stops the exporters.
func (m *Metrics) Shutdown(ctx context.Context) error {
	return m.provider.Shutdown(ctx)
}

type statusWriter struct {
	http.ResponseWriter
	status int
}

func (w *statusWriter) WriteHeader(code int) {
	w.status = code
	w.ResponseWriter.WriteHeader(code)
}
