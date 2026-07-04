// SPDX-License-Identifier: GPL-3.0-or-later

// Package geoip maintains xray-format geoip.dat datasets on disk and expands
// a category code (e.g. "cn", "ru-blocked") into the CIDR list that mieru's
// egress rules understand. mieru has no geoip engine of its own, so the panel
// does the translation.
//
// geoip.dat is the same format 3x-ui/xray use, so any published .dat works
// (Loyalsoldier for countries, runetfreedom for RU-blocked ranges, etc.),
// downloaded online or mounted into the datasets directory as "<name>.dat".
package geoip

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"sync"
	"time"
)

// Manager owns the datasets directory (*.dat files).
type Manager struct {
	dir    string
	client *http.Client

	mu       sync.Mutex
	cache    map[string][]string // category code -> CIDRs, merged across files
	cacheKey string              // fingerprint of the loaded files
}

// New returns a Manager storing datasets under dir (created if missing).
func New(dir string) (*Manager, error) {
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return nil, fmt.Errorf("create geoip dir: %w", err)
	}
	return &Manager{dir: dir, client: &http.Client{Timeout: 90 * time.Second}}, nil
}

// Dataset is one .dat file on disk.
type Dataset struct {
	Name  string `json:"name"`
	Bytes int64  `json:"bytes"`
}

// Category is a code present across the loaded datasets.
type Category struct {
	Code  string `json:"code"`
	CIDRs int    `json:"cidrs"`
}

// Datasets lists the .dat files present.
func (m *Manager) Datasets() ([]Dataset, error) {
	entries, err := os.ReadDir(m.dir)
	if err != nil {
		return nil, err
	}
	var out []Dataset
	for _, e := range entries {
		if e.IsDir() || !strings.HasSuffix(e.Name(), ".dat") {
			continue
		}
		info, err := e.Info()
		if err != nil {
			continue
		}
		out = append(out, Dataset{Name: strings.TrimSuffix(e.Name(), ".dat"), Bytes: info.Size()})
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

// CIDRs returns the CIDR list for a category code (merged across datasets).
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

// AddDataset downloads a geoip.dat from url and stores it as name.dat after
// verifying it parses.
func (m *Manager) AddDataset(ctx context.Context, name, url string) error {
	name = sanitizeName(name)
	if name == "" {
		return fmt.Errorf("invalid dataset name")
	}
	if url == "" {
		return fmt.Errorf("url is required")
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
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
	if _, err := parseGeoIPDat(data); err != nil {
		return fmt.Errorf("not a valid geoip.dat: %w", err)
	}

	m.mu.Lock()
	defer m.mu.Unlock()
	tmp := filepath.Join(m.dir, name+".dat.tmp")
	if err := os.WriteFile(tmp, data, 0o644); err != nil {
		return err
	}
	if err := os.Rename(tmp, filepath.Join(m.dir, name+".dat")); err != nil {
		return err
	}
	m.cacheKey = "" // force reload
	return nil
}

// DeleteDataset removes a dataset file.
func (m *Manager) DeleteDataset(name string) error {
	name = sanitizeName(name)
	if name == "" {
		return fmt.Errorf("invalid dataset name")
	}
	m.mu.Lock()
	defer m.mu.Unlock()
	m.cacheKey = ""
	err := os.Remove(filepath.Join(m.dir, name+".dat"))
	if os.IsNotExist(err) {
		return nil
	}
	return err
}

// reload re-parses the datasets if the set of files or their mtimes changed.
// Caller must hold m.mu.
func (m *Manager) reload() error {
	entries, err := os.ReadDir(m.dir)
	if err != nil {
		return err
	}
	var key strings.Builder
	var files []string
	for _, e := range entries {
		if e.IsDir() || !strings.HasSuffix(e.Name(), ".dat") {
			continue
		}
		info, err := e.Info()
		if err != nil {
			continue
		}
		fmt.Fprintf(&key, "%s:%d:%d;", e.Name(), info.Size(), info.ModTime().UnixNano())
		files = append(files, filepath.Join(m.dir, e.Name()))
	}
	if key.String() == m.cacheKey && m.cache != nil {
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
	m.cacheKey = key.String()
	return nil
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
