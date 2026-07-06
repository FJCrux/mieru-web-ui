// Package store is the panel's persistence layer.
//
// It holds only what mita cannot: the panel admin account, sessions, panel
// settings, and the plaintext passwords of mieru users the panel created
// (mita stores password hashes only, but client configs and share links
// need plaintext).
//
// Store is an interface so the backend can be swapped. SQLite is the only
// implementation today (see sqlite.go); a Postgres or MySQL backend just
// needs a type satisfying this interface plus an Open* constructor. The rest
// of the app depends on the interface, never on a concrete driver.
package store

import (
	"errors"
	"time"
)

// ErrNotFound is returned when a requested row does not exist.
var ErrNotFound = errors.New("not found")

// Store is the panel persistence contract.
type Store interface {
	// Admin returns the panel admin credentials, or ErrNotFound if unset.
	Admin() (username, passwordHash string, err error)
	// SetAdmin creates or replaces the panel admin account.
	SetAdmin(username, passwordHash string) error

	// CreateSession stores a session token hash valid until expiresAt.
	CreateSession(tokenHash string, expiresAt time.Time) error
	// SessionValid reports whether the token hash exists and is unexpired.
	SessionValid(tokenHash string) (bool, error)
	// DeleteSession removes a single session (logout).
	DeleteSession(tokenHash string) error
	// DeleteAllSessions removes every session (e.g. after a password change).
	DeleteAllSessions() error

	// SetUserSecret stores/updates the plaintext password of a mieru user.
	SetUserSecret(username, password string) error
	// UserSecret returns the stored plaintext password, or ErrNotFound.
	UserSecret(username string) (string, error)
	// UserSecretNames lists usernames that have a stored secret.
	UserSecretNames() (map[string]bool, error)
	// DeleteUserSecret removes a stored password.
	DeleteUserSecret(username string) error

	// Setting returns a setting value, or "" if unset.
	Setting(key string) (string, error)
	// SetSetting stores a key/value pair.
	SetSetting(key, value string) error

	// CreatePeer stores an upstream chain peer (a supervised mieru client).
	CreatePeer(p PeerRecord) error
	// Peers lists stored chain peers.
	Peers() ([]PeerRecord, error)
	// DeletePeer removes a chain peer.
	DeletePeer(name string) error

	// CreateShareToken stores a one-off share link token (hashed) that
	// resolves to a user's client config until expiresAt.
	CreateShareToken(tokenHash string, t ShareTokenData, expiresAt time.Time) error
	// ShareToken returns the data behind a token hash, or ErrNotFound if it
	// is missing or expired.
	ShareToken(tokenHash string) (ShareTokenData, error)
	// DeleteShareTokensForUser drops every share token for a user (called
	// when the user is deleted or their password changes).
	DeleteShareTokensForUser(username string) error

	// UpsertSubToken creates or rotates the single permanent subscription
	// token for t.Username (any previous token is invalidated atomically).
	UpsertSubToken(token string, t SubTokenData) error
	// SubToken resolves a subscription token, or ErrNotFound.
	SubToken(token string) (SubTokenData, error)
	// SubTokenForUser returns the user's subscription token, or ErrNotFound.
	SubTokenForUser(username string) (SubTokenData, error)
	// SubTokenUsernames lists usernames that have a subscription token.
	SubTokenUsernames() (map[string]bool, error)
	// DeleteSubTokenForUser revokes the user's subscription token.
	DeleteSubTokenForUser(username string) error

	// Backup writes a consistent snapshot of the store to dstPath (which must
	// not already exist). Used by the panel's backup export.
	Backup(dstPath string) error

	// Close releases the backend.
	Close() error
}

// ShareTokenData is what a share link resolves to.
type ShareTokenData struct {
	Username     string
	Host         string
	Multiplexing string
}

// SubTokenData is what a permanent subscription token resolves to. The token
// is stored in plaintext so the panel can re-display the subscription URL
// (the same database already holds the users' plaintext passwords).
type SubTokenData struct {
	Token        string
	Username     string
	Host         string // "" = use the public_host setting at fetch time
	Multiplexing string
	CreatedAt    time.Time
}

// PeerRecord is a stored upstream chain peer.
type PeerRecord struct {
	Name       string
	ConfigJSON string // mieru client config (with assigned local ports)
	Socks5Port int32  // local SOCKS5 port the peer's client listens on
}
