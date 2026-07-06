// SPDX-License-Identifier: GPL-3.0-or-later

// SQLite implementation of Store, using the pure-Go modernc.org/sqlite
// driver so the binary stays static (no cgo). To add another backend
// (Postgres, MySQL, ...), write a type that satisfies the Store interface
// in store.go and a corresponding Open* constructor.

package store

import (
	"database/sql"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"time"

	_ "modernc.org/sqlite"
)

type sqliteStore struct {
	db *sql.DB
}

// OpenSQLite creates/opens the SQLite-backed panel store at path and applies
// migrations.
func OpenSQLite(path string) (Store, error) {
	if err := os.MkdirAll(filepath.Dir(path), 0o700); err != nil {
		return nil, fmt.Errorf("create data dir: %w", err)
	}
	db, err := sql.Open("sqlite", path+"?_pragma=journal_mode(WAL)&_pragma=busy_timeout(5000)&_pragma=foreign_keys(1)")
	if err != nil {
		return nil, err
	}
	// SQLite handles one writer at a time; a single connection avoids
	// SQLITE_BUSY under concurrent API calls.
	db.SetMaxOpenConns(1)
	if err := migrate(db); err != nil {
		db.Close()
		return nil, err
	}
	if err := os.Chmod(path, 0o600); err != nil {
		db.Close()
		return nil, fmt.Errorf("chmod db: %w", err)
	}
	return &sqliteStore{db: db}, nil
}

func (s *sqliteStore) Close() error { return s.db.Close() }

// Backup writes a consistent snapshot to dstPath. VACUUM INTO produces a clean
// copy that folds in the WAL, so the snapshot is transactionally consistent
// even while the panel keeps running.
func (s *sqliteStore) Backup(dstPath string) error {
	if _, err := os.Stat(dstPath); err == nil {
		return fmt.Errorf("backup destination already exists: %s", dstPath)
	}
	_, err := s.db.Exec("VACUUM INTO ?", dstPath)
	return err
}

func migrate(db *sql.DB) error {
	_, err := db.Exec(`
CREATE TABLE IF NOT EXISTS admin (
    id INTEGER PRIMARY KEY CHECK (id = 1),
    username TEXT NOT NULL,
    password_hash TEXT NOT NULL
);
CREATE TABLE IF NOT EXISTS sessions (
    token_hash TEXT PRIMARY KEY,
    created_at INTEGER NOT NULL,
    expires_at INTEGER NOT NULL
);
CREATE TABLE IF NOT EXISTS user_secrets (
    mieru_username TEXT PRIMARY KEY,
    password TEXT NOT NULL,
    created_at INTEGER NOT NULL
);
CREATE TABLE IF NOT EXISTS settings (
    key TEXT PRIMARY KEY,
    value TEXT NOT NULL
);
CREATE TABLE IF NOT EXISTS share_tokens (
    token_hash TEXT PRIMARY KEY,
    mieru_username TEXT NOT NULL,
    host TEXT NOT NULL,
    multiplexing TEXT NOT NULL,
    expires_at INTEGER NOT NULL
);
CREATE TABLE IF NOT EXISTS sub_tokens (
    token TEXT PRIMARY KEY,
    mieru_username TEXT NOT NULL UNIQUE,
    host TEXT NOT NULL,
    multiplexing TEXT NOT NULL,
    created_at INTEGER NOT NULL
);
CREATE TABLE IF NOT EXISTS peers (
    name TEXT PRIMARY KEY,
    config_json TEXT NOT NULL,
    socks5_port INTEGER NOT NULL,
    created_at INTEGER NOT NULL
);`)
	return err
}

// --- admin ---

// Admin returns the panel admin credentials.
func (s *sqliteStore) Admin() (username, passwordHash string, err error) {
	err = s.db.QueryRow(`SELECT username, password_hash FROM admin WHERE id = 1`).Scan(&username, &passwordHash)
	if errors.Is(err, sql.ErrNoRows) {
		return "", "", ErrNotFound
	}
	return username, passwordHash, err
}

// SetAdmin creates or replaces the panel admin account.
func (s *sqliteStore) SetAdmin(username, passwordHash string) error {
	_, err := s.db.Exec(`
INSERT INTO admin (id, username, password_hash) VALUES (1, ?, ?)
ON CONFLICT (id) DO UPDATE SET username = excluded.username, password_hash = excluded.password_hash`,
		username, passwordHash)
	return err
}

// --- sessions ---

// CreateSession stores a session token hash valid until expiresAt.
func (s *sqliteStore) CreateSession(tokenHash string, expiresAt time.Time) error {
	_, err := s.db.Exec(`INSERT INTO sessions (token_hash, created_at, expires_at) VALUES (?, ?, ?)`,
		tokenHash, time.Now().Unix(), expiresAt.Unix())
	return err
}

// SessionValid reports whether the token hash exists and has not expired.
// Expired rows are lazily deleted.
func (s *sqliteStore) SessionValid(tokenHash string) (bool, error) {
	if _, err := s.db.Exec(`DELETE FROM sessions WHERE expires_at < ?`, time.Now().Unix()); err != nil {
		return false, err
	}
	var one int
	err := s.db.QueryRow(`SELECT 1 FROM sessions WHERE token_hash = ?`, tokenHash).Scan(&one)
	if errors.Is(err, sql.ErrNoRows) {
		return false, nil
	}
	return err == nil, err
}

// DeleteSession removes a session (logout).
func (s *sqliteStore) DeleteSession(tokenHash string) error {
	_, err := s.db.Exec(`DELETE FROM sessions WHERE token_hash = ?`, tokenHash)
	return err
}

// DeleteAllSessions removes every session (e.g. after a password change).
func (s *sqliteStore) DeleteAllSessions() error {
	_, err := s.db.Exec(`DELETE FROM sessions`)
	return err
}

// --- user secrets ---

// SetUserSecret stores the plaintext password for a mieru user the panel
// created or whose password it reset.
func (s *sqliteStore) SetUserSecret(username, password string) error {
	_, err := s.db.Exec(`
INSERT INTO user_secrets (mieru_username, password, created_at) VALUES (?, ?, ?)
ON CONFLICT (mieru_username) DO UPDATE SET password = excluded.password, created_at = excluded.created_at`,
		username, password, time.Now().Unix())
	return err
}

// UserSecret returns the stored plaintext password for username.
func (s *sqliteStore) UserSecret(username string) (string, error) {
	var pw string
	err := s.db.QueryRow(`SELECT password FROM user_secrets WHERE mieru_username = ?`, username).Scan(&pw)
	if errors.Is(err, sql.ErrNoRows) {
		return "", ErrNotFound
	}
	return pw, err
}

// UserSecretNames lists usernames that have a stored secret.
func (s *sqliteStore) UserSecretNames() (map[string]bool, error) {
	rows, err := s.db.Query(`SELECT mieru_username FROM user_secrets`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	names := make(map[string]bool)
	for rows.Next() {
		var n string
		if err := rows.Scan(&n); err != nil {
			return nil, err
		}
		names[n] = true
	}
	return names, rows.Err()
}

// DeleteUserSecret removes the stored password for username.
func (s *sqliteStore) DeleteUserSecret(username string) error {
	_, err := s.db.Exec(`DELETE FROM user_secrets WHERE mieru_username = ?`, username)
	return err
}

// --- settings ---

// Setting returns the value for key, or "" if unset.
func (s *sqliteStore) Setting(key string) (string, error) {
	var v string
	err := s.db.QueryRow(`SELECT value FROM settings WHERE key = ?`, key).Scan(&v)
	if errors.Is(err, sql.ErrNoRows) {
		return "", nil
	}
	return v, err
}

// SetSetting stores a key/value pair.
func (s *sqliteStore) SetSetting(key, value string) error {
	_, err := s.db.Exec(`
INSERT INTO settings (key, value) VALUES (?, ?)
ON CONFLICT (key) DO UPDATE SET value = excluded.value`, key, value)
	return err
}

// --- chain peers ---

func (s *sqliteStore) CreatePeer(p PeerRecord) error {
	_, err := s.db.Exec(`
INSERT INTO peers (name, config_json, socks5_port, created_at) VALUES (?, ?, ?, ?)`,
		p.Name, p.ConfigJSON, p.Socks5Port, time.Now().Unix())
	return err
}

func (s *sqliteStore) Peers() ([]PeerRecord, error) {
	rows, err := s.db.Query(`SELECT name, config_json, socks5_port FROM peers ORDER BY name`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var out []PeerRecord
	for rows.Next() {
		var p PeerRecord
		if err := rows.Scan(&p.Name, &p.ConfigJSON, &p.Socks5Port); err != nil {
			return nil, err
		}
		out = append(out, p)
	}
	return out, rows.Err()
}

func (s *sqliteStore) DeletePeer(name string) error {
	_, err := s.db.Exec(`DELETE FROM peers WHERE name = ?`, name)
	return err
}

// --- share tokens ---

func (s *sqliteStore) CreateShareToken(tokenHash string, t ShareTokenData, expiresAt time.Time) error {
	_, err := s.db.Exec(`
INSERT INTO share_tokens (token_hash, mieru_username, host, multiplexing, expires_at)
VALUES (?, ?, ?, ?, ?)`,
		tokenHash, t.Username, t.Host, t.Multiplexing, expiresAt.Unix())
	return err
}

func (s *sqliteStore) ShareToken(tokenHash string) (ShareTokenData, error) {
	if _, err := s.db.Exec(`DELETE FROM share_tokens WHERE expires_at < ?`, time.Now().Unix()); err != nil {
		return ShareTokenData{}, err
	}
	var t ShareTokenData
	err := s.db.QueryRow(`SELECT mieru_username, host, multiplexing FROM share_tokens WHERE token_hash = ?`, tokenHash).
		Scan(&t.Username, &t.Host, &t.Multiplexing)
	if errors.Is(err, sql.ErrNoRows) {
		return ShareTokenData{}, ErrNotFound
	}
	return t, err
}

func (s *sqliteStore) DeleteShareTokensForUser(username string) error {
	_, err := s.db.Exec(`DELETE FROM share_tokens WHERE mieru_username = ?`, username)
	return err
}

// --- subscription tokens ---

func (s *sqliteStore) UpsertSubToken(token string, t SubTokenData) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	if _, err := tx.Exec(`DELETE FROM sub_tokens WHERE mieru_username = ?`, t.Username); err != nil {
		return err
	}
	if _, err := tx.Exec(`
INSERT INTO sub_tokens (token, mieru_username, host, multiplexing, created_at)
VALUES (?, ?, ?, ?, ?)`,
		token, t.Username, t.Host, t.Multiplexing, time.Now().Unix()); err != nil {
		return err
	}
	return tx.Commit()
}

func (s *sqliteStore) SubToken(token string) (SubTokenData, error) {
	return s.subTokenBy(`token = ?`, token)
}

func (s *sqliteStore) SubTokenForUser(username string) (SubTokenData, error) {
	return s.subTokenBy(`mieru_username = ?`, username)
}

func (s *sqliteStore) subTokenBy(where string, arg any) (SubTokenData, error) {
	var t SubTokenData
	var createdAt int64
	err := s.db.QueryRow(`SELECT token, mieru_username, host, multiplexing, created_at FROM sub_tokens WHERE `+where, arg).
		Scan(&t.Token, &t.Username, &t.Host, &t.Multiplexing, &createdAt)
	if errors.Is(err, sql.ErrNoRows) {
		return SubTokenData{}, ErrNotFound
	}
	t.CreatedAt = time.Unix(createdAt, 0)
	return t, err
}

// SubTokenUsernames lists usernames that have a subscription token.
func (s *sqliteStore) SubTokenUsernames() (map[string]bool, error) {
	rows, err := s.db.Query(`SELECT mieru_username FROM sub_tokens`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	names := make(map[string]bool)
	for rows.Next() {
		var n string
		if err := rows.Scan(&n); err != nil {
			return nil, err
		}
		names[n] = true
	}
	return names, rows.Err()
}

func (s *sqliteStore) DeleteSubTokenForUser(username string) error {
	_, err := s.db.Exec(`DELETE FROM sub_tokens WHERE mieru_username = ?`, username)
	return err
}
