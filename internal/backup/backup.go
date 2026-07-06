// SPDX-License-Identifier: GPL-3.0-or-later

// Package backup builds, encrypts and restores portable panel backups.
// The HTTP handlers live in internal/api; this package owns the archive
// format, the crypto and the staged-restore mechanics.
package backup

import (
	"archive/tar"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/fjcrux/mieru-web-ui/internal/store"
)

// Paths tells the backup/restore code where the panel's state lives.
type Paths struct {
	DBPath     string // panel.db
	GeoDir     string // geoip datasets dir
	PeersDir   string // chain peers dir
	MitaConfig string // mita server.conf.pb
	// RestoreStaging is where an uploaded restore is unpacked; the panel
	// applies it on the next start (see ApplyPending). It must live on the
	// same filesystem as the targets so the swap is an atomic rename.
	RestoreStaging string
}

// Archive layout (inside the .tar.gz):
//
//	manifest.json          format/version + timestamp
//	panel.db               consistent VACUUM INTO snapshot
//	mita/server.conf.pb    mita daemon config
//	geoip/<name>.dat       geoip / geosite (.site.dat) datasets
//	peers/...              chain peer state
const Format = "imugi-panel-backup/1"

type manifest struct {
	Format       string `json:"format"`
	CreatedAt    string `json:"createdAt"`
	PanelVersion string `json:"panelVersion"`
}

// Build writes the tar.gz backup to w.
func Build(w io.Writer, st store.Store, p Paths, panelVersion string) error {
	gz := gzip.NewWriter(w)
	tw := tar.NewWriter(gz)

	m, _ := json.Marshal(manifest{
		Format:       Format,
		CreatedAt:    time.Now().UTC().Format(time.RFC3339),
		PanelVersion: panelVersion,
	})
	if err := writeTarBytes(tw, "manifest.json", m); err != nil {
		return err
	}

	// panel.db — consistent snapshot via VACUUM INTO.
	snap := p.DBPath + ".backup-" + fmt.Sprint(time.Now().UnixNano())
	if err := st.Backup(snap); err != nil {
		return fmt.Errorf("snapshot db: %w", err)
	}
	defer os.Remove(snap)
	if err := writeTarFile(tw, "panel.db", snap); err != nil {
		return err
	}

	// mita config (single file, may be absent on a fresh install).
	if p.MitaConfig != "" {
		if _, err := os.Stat(p.MitaConfig); err == nil {
			if err := writeTarFile(tw, "mita/server.conf.pb", p.MitaConfig); err != nil {
				return err
			}
		}
	}

	// geoip and peers directories (optional).
	if err := writeTarDir(tw, "geoip", p.GeoDir); err != nil {
		return err
	}
	if err := writeTarDir(tw, "peers", p.PeersDir); err != nil {
		return err
	}

	if err := tw.Close(); err != nil {
		return err
	}
	return gz.Close()
}

// --- tar helpers ---

func writeTarBytes(tw *tar.Writer, name string, data []byte) error {
	if err := tw.WriteHeader(&tar.Header{Name: name, Mode: 0o600, Size: int64(len(data)), ModTime: time.Now()}); err != nil {
		return err
	}
	_, err := tw.Write(data)
	return err
}

func writeTarFile(tw *tar.Writer, name, path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()
	info, err := f.Stat()
	if err != nil {
		return err
	}
	if err := tw.WriteHeader(&tar.Header{Name: name, Mode: 0o600, Size: info.Size(), ModTime: info.ModTime()}); err != nil {
		return err
	}
	_, err = io.Copy(tw, f)
	return err
}

func writeTarDir(tw *tar.Writer, prefix, dir string) error {
	if dir == "" {
		return nil
	}
	entries, err := os.ReadDir(dir)
	if os.IsNotExist(err) {
		return nil
	}
	if err != nil {
		return err
	}
	for _, e := range entries {
		if e.IsDir() {
			if err := writeTarDir(tw, prefix+"/"+e.Name(), filepath.Join(dir, e.Name())); err != nil {
				return err
			}
			continue
		}
		if err := writeTarFile(tw, prefix+"/"+e.Name(), filepath.Join(dir, e.Name())); err != nil {
			return err
		}
	}
	return nil
}

// Extract unpacks a tar.gz into dst, rejecting path traversal.
func Extract(r io.Reader, dst string) error {
	gz, err := gzip.NewReader(r)
	if err != nil {
		return fmt.Errorf("not a gzip archive: %w", err)
	}
	defer gz.Close()
	tr := tar.NewReader(gz)
	for {
		hdr, err := tr.Next()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		// Reject absolute paths and traversal (".." components).
		clean := filepath.Clean(hdr.Name)
		if strings.HasPrefix(clean, "..") || strings.HasPrefix(clean, "/") || strings.Contains(clean, ".."+string(os.PathSeparator)) {
			return fmt.Errorf("unsafe path in archive: %q", hdr.Name)
		}
		target := filepath.Join(dst, clean)
		if hdr.FileInfo().IsDir() {
			if err := os.MkdirAll(target, 0o700); err != nil {
				return err
			}
			continue
		}
		if err := os.MkdirAll(filepath.Dir(target), 0o700); err != nil {
			return err
		}
		out, err := os.OpenFile(target, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0o600)
		if err != nil {
			return err
		}
		if _, err := io.Copy(out, tr); err != nil {
			out.Close()
			return err
		}
		out.Close()
	}
}
