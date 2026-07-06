// SPDX-License-Identifier: GPL-3.0-or-later

// Command mieru-web-ui is the web admin panel for the mieru proxy server.
// It supervises the mita daemon as a child process (unless --no-supervise)
// and serves the management SPA + REST API.
package main

import (
	"context"
	"encoding/json"
	"errors"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/fjcrux/mieru-web-ui/internal/api"
	"github.com/fjcrux/mieru-web-ui/internal/auth"
	"github.com/fjcrux/mieru-web-ui/internal/backup"
	"github.com/fjcrux/mieru-web-ui/internal/chain"
	"github.com/fjcrux/mieru-web-ui/internal/geoip"
	"github.com/fjcrux/mieru-web-ui/internal/metrics"
	"github.com/fjcrux/mieru-web-ui/internal/mitaclient"
	"github.com/fjcrux/mieru-web-ui/internal/store"
	"github.com/fjcrux/mieru-web-ui/internal/supervisor"
	"github.com/fjcrux/mieru-web-ui/internal/webfs"
)

// version is the panel version, overridable via -ldflags "-X main.version=...".
var version = "dev"

func envOr(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}

// boolEnv reads a boolean environment flag (true/1/yes/on, case-insensitive).
func boolEnv(key string) bool {
	switch strings.ToLower(strings.TrimSpace(os.Getenv(key))) {
	case "1", "true", "yes", "on":
		return true
	}
	return false
}

// seedSetting stores value under key only if the key is currently unset and
// value is non-empty (first-run seeding from the environment).
func seedSetting(st store.Store, key, value string) {
	if value == "" {
		return
	}
	if cur, _ := st.Setting(key); cur == "" {
		_ = st.SetSetting(key, value)
	}
}

// settingOr returns the stored setting, or def when unset.
func settingOr(st store.Store, key, def string) string {
	if v, err := st.Setting(key); err == nil && v != "" {
		return v
	}
	return def
}

// normalizePath returns a clean URL prefix with a leading slash and no
// trailing slash, or "" for root.
func normalizePath(p string) string {
	p = strings.TrimSpace(p)
	if p == "" || p == "/" {
		return ""
	}
	if !strings.HasPrefix(p, "/") {
		p = "/" + p
	}
	return strings.TrimRight(p, "/")
}

// noindex tells crawlers to stay away from every page the panel serves
// (admin SPA, share pages, health endpoint). robots.txt below is the polite
// version; the header also covers crawlers that only honor per-response tags.
func noindex(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-Robots-Tag", "noindex, nofollow, noarchive")
		next.ServeHTTP(w, r)
	})
}

// securityHeaders sets browser hardening headers on every response. The CSP
// fits the embedded SPA: everything is same-origin except the inline base
// script webfs injects (allowed by hash) and naive-ui's runtime styles.
// Handlers that need a different policy (the share page) override it.
func securityHeaders(hsts bool, baseScriptHash string) func(http.Handler) http.Handler {
	csp := "default-src 'self'; script-src 'self' '" + baseScriptHash + "'; " +
		"style-src 'self' 'unsafe-inline'; img-src 'self' data: blob:; font-src 'self' data:; " +
		"connect-src 'self'; object-src 'none'; frame-ancestors 'none'; base-uri 'self'; form-action 'self'"
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			h := w.Header()
			h.Set("Content-Security-Policy", csp)
			h.Set("X-Content-Type-Options", "nosniff")
			h.Set("X-Frame-Options", "DENY")
			h.Set("Referrer-Policy", "no-referrer")
			if hsts {
				h.Set("Strict-Transport-Security", "max-age=63072000")
			}
			next.ServeHTTP(w, r)
		})
	}
}

// metricsCollector samples panel/mita state for the OTel exporter.
func metricsCollector(mita *mitaclient.Client, peers *chain.Manager, sup *supervisor.Supervisor) metrics.Collector {
	return func(ctx context.Context) metrics.Snapshot {
		s := metrics.Snapshot{}
		if st, err := mita.Status(ctx); err == nil && st == "RUNNING" {
			s.MitaUp = 1
		}
		if u, err := mita.Users(ctx); err == nil {
			s.Users = int64(len(u.GetItems()))
		}
		if sess, err := mita.Sessions(ctx); err == nil {
			s.Sessions = int64(len(sess.GetItems()))
		}
		if raw, err := mita.Metrics(ctx); err == nil {
			var m map[string]map[string]int64
			if json.Unmarshal([]byte(raw), &m) == nil {
				s.DownloadBytes = m["traffic"]["DownloadBytes"]
				s.UploadBytes = m["traffic"]["UploadBytes"]
			}
		}
		for _, p := range peers.List() {
			s.PeersTotal++
			if p.Running {
				s.PeersConnected++
			}
		}
		if sup != nil {
			r, _ := sup.Stats()
			s.MitaRestarts = int64(r)
		}
		return s
	}
}

func main() {
	noSupervise := flag.Bool("no-supervise", false, "do not run mita as a child process (connect to an external daemon)")
	flag.Parse()

	panelPort := envOr("PANEL_PORT", "8686")
	panelBind := envOr("PANEL_BIND", "")
	dataDir := envOr("PANEL_DATA_DIR", "/data")
	udsPath := envOr("MITA_UDS_PATH", mitaclient.DefaultUDSPath)
	rpcTarget := envOr("MITA_RPC_TARGET", "unix://"+udsPath)
	mitaBinary := envOr("MITA_BINARY", "/usr/local/bin/mita")
	tlsCert := os.Getenv("PANEL_TLS_CERT")
	tlsKey := os.Getenv("PANEL_TLS_KEY")
	// Trust X-Forwarded-* from a reverse proxy / load balancer in front of the
	// panel (e.g. a TLS-terminating balancer). Enable only when such a proxy is
	// present, or clients could forge these headers.
	trustProxy := boolEnv("PANEL_TRUST_PROXY")

	// Resolve the paths that make up the panel's state, shared by the store,
	// geoip/chain init, and the backup/restore endpoints.
	statePaths := backup.Paths{
		DBPath:         filepath.Join(dataDir, "panel.db"),
		GeoDir:         envOr("GEOIP_DIR", filepath.Join(dataDir, "geoip")),
		PeersDir:       filepath.Join(dataDir, "peers"),
		MitaConfig:     envOr("MITA_CONFIG_FILE", "/etc/mita/server.conf.pb"),
		RestoreStaging: filepath.Join(dataDir, ".restore"),
	}
	dbPath := statePaths.DBPath
	geoDir := statePaths.GeoDir
	mitaConfigFile := statePaths.MitaConfig

	// Apply a restore staged by the restore endpoint before opening anything,
	// so swapping the live panel.db is safe.
	if err := backup.ApplyPending(statePaths); err != nil {
		log.Printf("apply pending restore: %v", err)
	}

	st, err := store.OpenSQLite(dbPath)
	if err != nil {
		log.Fatalf("open store: %v", err)
	}
	defer st.Close()

	if err := auth.Bootstrap(st, os.Getenv("PANEL_ADMIN_USER"), os.Getenv("PANEL_ADMIN_PASSWORD")); err != nil {
		log.Fatalf("bootstrap admin: %v", err)
	}

	// Seed URL/path settings from the environment on first run; they are
	// editable in Settings afterwards. Path changes take effect on restart.
	seedSetting(st, "public_host", strings.TrimSpace(os.Getenv("PUBLIC_HOST")))
	seedSetting(st, "panel_url", strings.TrimRight(os.Getenv("PANEL_URL"), "/"))
	seedSetting(st, "base_path", os.Getenv("PANEL_BASE_PATH"))
	seedSetting(st, "share_path", os.Getenv("SHARE_PATH"))
	seedSetting(st, "sub_path", os.Getenv("SUB_PATH"))
	seedSetting(st, "sub_port", os.Getenv("SUB_PORT"))

	// Desired ports live in the panel store and are applied to mita only when a
	// user exists (mita crashloops with ports but no users). PROXY_PORTS, when
	// set, is the source of truth and overwrites the stored value each start;
	// otherwise seed a default on first run.
	proxyPorts := os.Getenv("PROXY_PORTS")
	if proxyPorts != "" {
		_ = st.SetSetting("desired_ports", proxyPorts)
	} else {
		seedSetting(st, "desired_ports", envOr("DEFAULT_PORTS", "2012"))
	}

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	// Before starting mita, fix its on-disk config so it can't crashloop from a
	// ports-without-users state (heals deployments created by older versions).
	if !*noSupervise {
		spec := settingOr(st, "desired_ports", "")
		if healed, err := mitaclient.HealConfigFile(mitaConfigFile, spec); err != nil {
			log.Printf("heal mita config: %v", err)
		} else if healed {
			log.Print("healed mita config (removed ports without users)")
		}
	}

	var sup *supervisor.Supervisor
	if !*noSupervise {
		sup = supervisor.New("mita", mitaBinary, []string{
			"MITA_UDS_PATH=" + udsPath,
			"MITA_INSECURE_UDS=true",
		})
		go func() {
			if err := sup.Run(ctx); err != nil && !errors.Is(err, context.Canceled) {
				log.Printf("supervisor stopped: %v", err)
			}
		}()
	}

	mita, err := mitaclient.Dial(rpcTarget)
	if err != nil {
		log.Fatalf("dial mita: %v", err)
	}
	defer mita.Close()

	geo, err := geoip.New(geoDir)
	if err != nil {
		log.Fatalf("init geoip: %v", err)
	}

	mieruBinary := envOr("MIERU_BINARY", "/usr/local/bin/mieru")
	peers, err := chain.New(mieruBinary, statePaths.PeersDir, st)
	if err != nil {
		log.Fatalf("init chain: %v", err)
	}
	if err := peers.Start(ctx); err != nil {
		log.Printf("start chain peers: %v", err)
	}

	basePath := normalizePath(settingOr(st, "base_path", "")) // "" = root
	sharePath := normalizePath(settingOr(st, "share_path", "/s"))
	if sharePath == "" {
		sharePath = "/s"
	}
	subPath := normalizePath(settingOr(st, "sub_path", "/sub"))
	if subPath == "" {
		subPath = "/sub"
	}
	// The settings API rejects colliding paths, but an env seed on first run
	// could still produce one, which would panic mux registration.
	if subPath == sharePath || subPath == basePath {
		fallback := "/sub"
		if fallback == sharePath || fallback == basePath {
			fallback = "/clashsub"
		}
		log.Printf("sub path %q collides with another path, using %s", subPath, fallback)
		subPath = fallback
	}

	// A dedicated subscription port gets its own listener serving only
	// subscription routes. Subscriptions embed user passwords, so the
	// dedicated listener requires TLS: its own SUB_TLS_CERT/KEY, else the
	// panel's certs, else refuse to start (SUB_INSECURE_HTTP=1 overrides for
	// setups where a TLS-terminating reverse proxy fronts the port).
	subPort := strings.TrimSpace(settingOr(st, "sub_port", ""))
	if subPort != "" {
		if n, err := strconv.Atoi(subPort); err != nil || n < 1 || n > 65535 || subPort == panelPort {
			log.Printf("invalid sub port %q, serving subscriptions on the panel port", subPort)
			subPort = ""
		}
	}
	subTLSCert := envOr("SUB_TLS_CERT", tlsCert)
	subTLSKey := envOr("SUB_TLS_KEY", tlsKey)
	if subPort != "" && subTLSCert == "" {
		if boolEnv("SUB_INSECURE_HTTP") {
			log.Print("WARNING: subscription port serves plain HTTP (SUB_INSECURE_HTTP=1); subscriptions contain user passwords - make sure a TLS-terminating proxy fronts this port")
		} else {
			log.Fatal("SUB_PORT is set but no TLS is configured: subscriptions contain user passwords and must not travel over plain HTTP. Set SUB_TLS_CERT/SUB_TLS_KEY (or PANEL_TLS_CERT/PANEL_TLS_KEY), or set SUB_INSECURE_HTTP=1 if a TLS-terminating reverse proxy fronts this port.")
		}
	}

	// Show the effective paths on startup, in case a secret base path was set
	// and forgotten.
	adminPath := basePath + "/"
	subAt := subPath
	if subPort != "" {
		subAt = ":" + subPort + subPath
	}
	log.Printf("paths: admin=%s  share=%s/<token>  sub=%s/<token>", adminPath, sharePath, subAt)
	if pu := settingOr(st, "panel_url", ""); pu != "" {
		log.Printf("panel URL: %s%s", pu, basePath)
	}

	// Secure cookies whenever the browser reaches the panel over HTTPS: either
	// the panel terminates TLS itself, or a trusted reverse proxy does.
	a := auth.New(st, auth.DefaultSessionTTL, tlsCert != "" || trustProxy)
	srv := &api.Server{Mita: mita, Store: st, Auth: a, Geo: geo, Peers: peers, SharePath: sharePath, PortsManaged: proxyPorts != "", PanelPort: panelPort, TLSEnabled: tlsCert != "", TrustProxy: trustProxy, Version: version, Updates: api.NewUpdateChecker(!boolEnv("PANEL_DISABLE_UPDATE_CHECK")),
		Backup: &statePaths}
	srv.ActiveBasePath = basePath
	srv.ActiveSharePath = sharePath
	srv.SubPath = subPath
	srv.SubPort = subPort
	srv.ActiveSubPath = subPath
	srv.ActiveSubPort = subPort
	// Restart = graceful shutdown; a Docker restart policy brings the panel
	// back with the new config. Cancelling ctx runs the shutdown path below.
	srv.Restart = stop

	// Bring mita into a safe state (and heal any ports-without-users crashloop).
	go srv.StartupReconcile(ctx)
	// Assign only when present so Sup stays a nil interface (not a typed-nil)
	// under --no-supervise, keeping the s.Sup == nil checks correct.
	if sup != nil {
		srv.Sup = sup
	}

	// The admin app (API + SPA) is served under basePath; the SPA learns its
	// base via webfs injection.
	adminMux := http.NewServeMux()
	adminMux.Handle("/api/", srv.Routes())
	adminMux.Handle("/", webfs.Handler(basePath))

	root := http.NewServeMux()
	root.HandleFunc("GET /healthz", srv.HandleHealth)
	root.HandleFunc("GET /robots.txt", func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		_, _ = w.Write([]byte("User-agent: *\nDisallow: /\n"))
	})
	// Public share pages live under their own prefix, not under basePath, so a
	// shared link never reveals the admin path.
	root.HandleFunc("GET "+sharePath+"/{token}", srv.HandlePublicSharePage)
	// Subscriptions live on the panel port unless a dedicated port is set, in
	// which case they get their own listener below and stay off this mux.
	if subPort == "" {
		root.HandleFunc("GET "+subPath+"/{token}", srv.HandlePublicSubscription)
	}
	if basePath == "" {
		root.Handle("/", adminMux)
	} else {
		root.Handle(basePath+"/", http.StripPrefix(basePath, adminMux))
		root.HandleFunc(basePath, func(w http.ResponseWriter, r *http.Request) {
			http.Redirect(w, r, basePath+"/", http.StatusFound)
		})
		log.Printf("panel served under base path %s", basePath)
	}

	// OpenTelemetry metrics (OTLP push and/or a Prometheus endpoint).
	var reqMiddleware func(http.Handler) http.Handler
	if metrics.Enabled() {
		mtr, err := metrics.Setup(ctx, version, metricsCollector(mita, peers, sup))
		if err != nil {
			log.Fatalf("init metrics: %v", err)
		}
		defer func() {
			shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()
			_ = mtr.Shutdown(shutdownCtx)
		}()
		if h := mtr.PromHandler(); h != nil {
			root.Handle("GET /metrics", h)
		}
		reqMiddleware = mtr.Middleware
		log.Print("metrics enabled")
	}

	// Reject requests with an unexpected Host (when a panel URL is configured),
	// then count them if metrics are on. HSTS only makes sense when browsers
	// reach the panel over HTTPS (own TLS or a trusted terminating proxy).
	harden := securityHeaders(tlsCert != "" || trustProxy, webfs.BaseScriptHash(basePath))
	var handler http.Handler = harden(noindex(srv.HostGuard(root)))
	if reqMiddleware != nil {
		handler = reqMiddleware(handler)
	}

	httpServer := &http.Server{
		Addr:              panelBind + ":" + panelPort,
		Handler:           handler,
		ReadHeaderTimeout: 10 * time.Second,
	}

	// Dedicated subscription listener: same middleware stack, but the mux
	// carries only the public subscription route.
	if subPort != "" {
		subMux := http.NewServeMux()
		subMux.HandleFunc("GET "+subPath+"/{token}", srv.HandlePublicSubscription)
		var subHandler http.Handler = harden(noindex(srv.HostGuard(subMux)))
		if reqMiddleware != nil {
			subHandler = reqMiddleware(subHandler)
		}
		subServer := &http.Server{
			Addr:              panelBind + ":" + subPort,
			Handler:           subHandler,
			ReadHeaderTimeout: 10 * time.Second,
		}
		go func() {
			<-ctx.Done()
			shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()
			_ = subServer.Shutdown(shutdownCtx)
		}()
		go func() {
			subScheme := "http"
			if subTLSCert != "" {
				subScheme = "https"
			}
			log.Printf("subscriptions listening on %s://%s", subScheme, subServer.Addr)
			var err error
			if subTLSCert != "" {
				err = subServer.ListenAndServeTLS(subTLSCert, subTLSKey)
			} else {
				err = subServer.ListenAndServe()
			}
			if err != nil && !errors.Is(err, http.ErrServerClosed) {
				log.Fatalf("subscription server: %v", err)
			}
		}()
	}

	go func() {
		<-ctx.Done()
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		_ = httpServer.Shutdown(shutdownCtx)
	}()

	scheme := "http"
	if tlsCert != "" {
		scheme = "https"
	}
	log.Printf("mieru-web-ui listening on %s://%s", scheme, httpServer.Addr)
	if tlsCert != "" {
		err = httpServer.ListenAndServeTLS(tlsCert, tlsKey)
	} else {
		err = httpServer.ListenAndServe()
	}
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatalf("http server: %v", err)
	}
	log.Print("shutdown complete")
}
