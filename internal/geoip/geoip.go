// SPDX-License-Identifier: GPL-3.0-or-later

// Package geoip maintains xray-format geoip.dat and geosite.dat datasets on
// disk and expands a category code into the CIDR list (geoip) or domain-suffix
// list (geosite) that mieru's egress rules understand. mieru has no geo engine
// of its own, so the panel does the translation.
//
// Both formats are the same ones 3x-ui/xray use, so any published .dat works
// (Loyalsoldier or runetfreedom for geoip, runetfreedom or v2fly for geosite),
// downloaded online or mounted into the datasets directory. geoip files are
// stored as "<name>.dat" and geosite files as "<name>.site.dat" so the two
// kinds coexist in one directory.
package geoip

import (
	"context"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"sync"
	"syscall"
	"time"
)

// File suffixes distinguish the two dataset kinds in one directory. geosite
// files carry ".site.dat" so a geoip scan skips them and vice versa.
const (
	geoipSuffix = ".dat"
	siteSuffix  = ".site.dat"
)

// Manager owns the datasets directory (geoip *.dat and geosite *.site.dat).
type Manager struct {
	dir    string
	client *http.Client

	mu           sync.Mutex
	cache        map[string][]string // geoip: category code -> CIDRs, merged
	cacheKey     string              // fingerprint of loaded geoip files
	siteCache    map[string][]string // geosite: category code -> domain suffixes
	siteCacheKey string              // fingerprint of loaded geosite files
}

// New returns a Manager storing datasets under dir (created if missing).
func New(dir string) (*Manager, error) {
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return nil, fmt.Errorf("create geoip dir: %w", err)
	}
	return &Manager{dir: dir, client: ssrfSafeClient()}, nil
}

// ssrfSafeClient is an HTTP client for fetching operator-supplied dataset URLs
// that refuses to connect to non-public IP addresses. The check runs in the
// dialer's Control hook, i.e. on the actual resolved IP just before the socket
// is opened, so it also covers HTTP redirects and DNS-rebinding — a URL that
// passes a naive up-front check but resolves/redirects to an internal address
// is still blocked at dial time.
func ssrfSafeClient() *http.Client {
	dialer := &net.Dialer{Timeout: 30 * time.Second, Control: blockPrivateDial}
	return &http.Client{
		Timeout:   90 * time.Second,
		Transport: &http.Transport{DialContext: dialer.DialContext},
	}
}

// blockPrivateDial rejects a dial to any address that is not a routable public
// IP, closing off SSRF to loopback, private ranges, and link-local (which
// includes the cloud metadata endpoint 169.254.169.254).
func blockPrivateDial(_, address string, _ syscall.RawConn) error {
	host, _, err := net.SplitHostPort(address)
	if err != nil {
		return err
	}
	ip := net.ParseIP(host)
	if ip == nil {
		return fmt.Errorf("cannot resolve dial address %q", address)
	}
	if !isPublicIP(ip) {
		return fmt.Errorf("refusing to connect to non-public address %s", ip)
	}
	return nil
}

// isPublicIP reports whether ip is a globally routable unicast address.
func isPublicIP(ip net.IP) bool {
	// IsPrivate covers RFC1918 and IPv6 unique-local (fc00::/7);
	// IsLinkLocalUnicast covers 169.254.0.0/16 (cloud metadata) and fe80::/10.
	if ip.IsLoopback() || ip.IsPrivate() || ip.IsUnspecified() ||
		ip.IsLinkLocalUnicast() || ip.IsLinkLocalMulticast() || ip.IsMulticast() {
		return false
	}
	return true
}

// Dataset is one .dat file on disk.
type Dataset struct {
	Name  string `json:"name"`
	Bytes int64  `json:"bytes"`
}

// Category is a geoip code present across the loaded datasets.
type Category struct {
	Code  string `json:"code"`
	CIDRs int    `json:"cidrs"`
}

// SiteCategory is a geosite code present across the loaded datasets.
type SiteCategory struct {
	Code    string `json:"code"`
	Domains int    `json:"domains"`
}

// isSiteFile reports whether a directory entry is a geosite (.site.dat) file.
func isSiteFile(name string) bool { return strings.HasSuffix(name, siteSuffix) }

// isGeoipFile reports whether a directory entry is a geoip (.dat, not .site.dat)
// file.
func isGeoipFile(name string) bool {
	return strings.HasSuffix(name, geoipSuffix) && !isSiteFile(name)
}

// Datasets lists the geoip .dat files present.
func (m *Manager) Datasets() ([]Dataset, error) {
	return m.datasets(isGeoipFile, geoipSuffix)
}

// SiteDatasets lists the geosite .site.dat files present.
func (m *Manager) SiteDatasets() ([]Dataset, error) {
	return m.datasets(isSiteFile, siteSuffix)
}

func (m *Manager) datasets(match func(string) bool, suffix string) ([]Dataset, error) {
	entries, err := os.ReadDir(m.dir)
	if err != nil {
		return nil, err
	}
	var out []Dataset
	for _, e := range entries {
		if e.IsDir() || !match(e.Name()) {
			continue
		}
		info, err := e.Info()
		if err != nil {
			continue
		}
		out = append(out, Dataset{Name: strings.TrimSuffix(e.Name(), suffix), Bytes: info.Size()})
	}
	sort.Slice(out, func(i, j int) bool { return out[i].Name < out[j].Name })
	return out, nil
}

// Categories lists every category code across all datasets with its CIDR count.
func (m *Manager) Categories() ([]Category, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if err := m.reload(); err != nil {
		return nil, err
	}
	out := make([]Category, 0, len(m.cache))
	for code, cidrs := range m.cache {
		out = append(out, Category{Code: code, CIDRs: len(cidrs)})
	}
	sort.Slice(out, func(i, j int) bool { return out[i].Code < out[j].Code })
	return out, nil
}

// CIDRs returns the CIDR list for a geoip category code (merged across
// datasets).
func (m *Manager) CIDRs(code string) ([]string, error) {
	code = strings.ToLower(strings.TrimSpace(code))
	m.mu.Lock()
	defer m.mu.Unlock()
	if err := m.reload(); err != nil {
		return nil, err
	}
	cidrs, ok := m.cache[code]
	if !ok {
		return nil, fmt.Errorf("geoip category %q not found - add a dataset that provides it", code)
	}
	return cidrs, nil
}

// SiteCategories lists every geosite category code with its domain count.
func (m *Manager) SiteCategories() ([]SiteCategory, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if err := m.reloadSites(); err != nil {
		return nil, err
	}
	out := make([]SiteCategory, 0, len(m.siteCache))
	for code, domains := range m.siteCache {
		out = append(out, SiteCategory{Code: code, Domains: len(domains)})
	}
	sort.Slice(out, func(i, j int) bool { return out[i].Code < out[j].Code })
	return out, nil
}

// Domains returns the domain-suffix list for a geosite category code (merged
// across datasets).
func (m *Manager) Domains(code string) ([]string, error) {
	code = strings.ToLower(strings.TrimSpace(code))
	m.mu.Lock()
	defer m.mu.Unlock()
	if err := m.reloadSites(); err != nil {
		return nil, err
	}
	domains, ok := m.siteCache[code]
	if !ok {
		return nil, fmt.Errorf("geosite category %q not found - add a dataset that provides it", code)
	}
	return domains, nil
}

// AddDataset downloads a geoip.dat from url and stores it as name.dat after
// verifying it parses.
func (m *Manager) AddDataset(ctx context.Context, name, rawURL string) error {
	return m.addDataset(ctx, name, rawURL, geoipSuffix, func(b []byte) error {
		if _, err := parseGeoIPDat(b); err != nil {
			return fmt.Errorf("not a valid geoip.dat: %w", err)
		}
		return nil
	})
}

// AddSiteDataset downloads a geosite.dat from url and stores it as
// name.site.dat after verifying it parses.
func (m *Manager) AddSiteDataset(ctx context.Context, name, rawURL string) error {
	return m.addDataset(ctx, name, rawURL, siteSuffix, func(b []byte) error {
		if _, err := parseGeoSiteDat(b); err != nil {
			return fmt.Errorf("not a valid geosite.dat: %w", err)
		}
		return nil
	})
}

func (m *Manager) addDataset(ctx context.Context, name, rawURL, suffix string, verify func([]byte) error) error {
	name = sanitizeName(name)
	if name == "" {
		return fmt.Errorf("invalid dataset name")
	}
	if rawURL == "" {
		return fmt.Errorf("url is required")
	}
	u, err := url.Parse(rawURL)
	if err != nil {
		return fmt.Errorf("invalid url: %w", err)
	}
	if u.Scheme != "http" && u.Scheme != "https" {
		return fmt.Errorf("url must be http or https")
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u.String(), nil)
	if err != nil {
		return err
	}
	resp, err := m.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("download failed: status %d", resp.StatusCode)
	}
	data, err := io.ReadAll(io.LimitReader(resp.Body, 128<<20))
	if err != nil {
		return err
	}
	if err := verify(data); err != nil {
		return err
	}

	m.mu.Lock()
	defer m.mu.Unlock()
	tmp := filepath.Join(m.dir, name+suffix+".tmp")
	if err := os.WriteFile(tmp, data, 0o644); err != nil {
		return err
	}
	if err := os.Rename(tmp, filepath.Join(m.dir, name+suffix)); err != nil {
		return err
	}
	m.cacheKey = "" // force reload of both kinds
	m.siteCacheKey = ""
	return nil
}

// DeleteDataset removes a geoip dataset file.
func (m *Manager) DeleteDataset(name string) error {
	return m.deleteDataset(name, geoipSuffix)
}

// DeleteSiteDataset removes a geosite dataset file.
func (m *Manager) DeleteSiteDataset(name string) error {
	return m.deleteDataset(name, siteSuffix)
}

func (m *Manager) deleteDataset(name, suffix string) error {
	name = sanitizeName(name)
	if name == "" {
		return fmt.Errorf("invalid dataset name")
	}
	m.mu.Lock()
	defer m.mu.Unlock()
	m.cacheKey = ""
	m.siteCacheKey = ""
	err := os.Remove(filepath.Join(m.dir, name+suffix))
	if os.IsNotExist(err) {
		return nil
	}
	return err
}

// reload re-parses the geoip datasets if the set of files or their mtimes
// changed. Caller must hold m.mu.
func (m *Manager) reload() error {
	key, files, err := m.scan(isGeoipFile)
	if err != nil {
		return err
	}
	if key == m.cacheKey && m.cache != nil {
		return nil
	}
	merged := map[string][]string{}
	for _, f := range files {
		data, err := os.ReadFile(f)
		if err != nil {
			return err
		}
		cats, err := parseGeoIPDat(data)
		if err != nil {
			return fmt.Errorf("parse %s: %w", filepath.Base(f), err)
		}
		for code, cidrs := range cats {
			merged[code] = append(merged[code], cidrs...)
		}
	}
	m.cache = merged
	m.cacheKey = key
	return nil
}

// reloadSites re-parses the geosite datasets if they changed. Caller must hold
// m.mu.
func (m *Manager) reloadSites() error {
	key, files, err := m.scan(isSiteFile)
	if err != nil {
		return err
	}
	if key == m.siteCacheKey && m.siteCache != nil {
		return nil
	}
	merged := map[string][]string{}
	for _, f := range files {
		data, err := os.ReadFile(f)
		if err != nil {
			return err
		}
		cats, err := parseGeoSiteDat(data)
		if err != nil {
			return fmt.Errorf("parse %s: %w", filepath.Base(f), err)
		}
		for code, domains := range cats {
			merged[code] = append(merged[code], domains...)
		}
	}
	m.siteCache = merged
	m.siteCacheKey = key
	return nil
}

// scan returns a fingerprint of the matching dataset files and their paths.
// Caller must hold m.mu.
func (m *Manager) scan(match func(string) bool) (string, []string, error) {
	entries, err := os.ReadDir(m.dir)
	if err != nil {
		return "", nil, err
	}
	var key strings.Builder
	var files []string
	for _, e := range entries {
		if e.IsDir() || !match(e.Name()) {
			continue
		}
		info, err := e.Info()
		if err != nil {
			continue
		}
		fmt.Fprintf(&key, "%s:%d:%d;", e.Name(), info.Size(), info.ModTime().UnixNano())
		files = append(files, filepath.Join(m.dir, e.Name()))
	}
	return key.String(), files, nil
}

func sanitizeName(name string) string {
	name = strings.ToLower(strings.TrimSpace(name))
	name = strings.TrimSuffix(name, ".dat")
	for _, r := range name {
		if !(r >= 'a' && r <= 'z') && !(r >= '0' && r <= '9') && r != '-' && r != '_' {
			return ""
		}
	}
	if name == "" || len(name) > 40 {
		return ""
	}
	return name
}
