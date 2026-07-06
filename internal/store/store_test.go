// SPDX-License-Identifier: GPL-3.0-or-later

package store

import (
	"errors"
	"path/filepath"
	"testing"
	"time"
)

func testStore(t *testing.T) Store {
	t.Helper()
	st, err := OpenSQLite(filepath.Join(t.TempDir(), "panel.db"))
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { st.Close() })
	return st
}

func TestAdminRoundTrip(t *testing.T) {
	st := testStore(t)
	if _, _, err := st.Admin(); !errors.Is(err, ErrNotFound) {
		t.Fatalf("want ErrNotFound, got %v", err)
	}
	if err := st.SetAdmin("admin", "hash1"); err != nil {
		t.Fatal(err)
	}
	u, h, err := st.Admin()
	if err != nil || u != "admin" || h != "hash1" {
		t.Fatalf("got %q %q %v", u, h, err)
	}
	if err := st.SetAdmin("admin", "hash2"); err != nil {
		t.Fatal(err)
	}
	_, h, _ = st.Admin()
	if h != "hash2" {
		t.Fatalf("hash not updated: %q", h)
	}
}

func TestSessions(t *testing.T) {
	st := testStore(t)
	if err := st.CreateSession("tok1", time.Now().Add(time.Hour)); err != nil {
		t.Fatal(err)
	}
	if ok, _ := st.SessionValid("tok1"); !ok {
		t.Fatal("tok1 should be valid")
	}
	if ok, _ := st.SessionValid("unknown"); ok {
		t.Fatal("unknown token should be invalid")
	}
	// Expired sessions are dropped.
	if err := st.CreateSession("tok2", time.Now().Add(-time.Minute)); err != nil {
		t.Fatal(err)
	}
	if ok, _ := st.SessionValid("tok2"); ok {
		t.Fatal("expired token should be invalid")
	}
	if err := st.DeleteAllSessions(); err != nil {
		t.Fatal(err)
	}
	if ok, _ := st.SessionValid("tok1"); ok {
		t.Fatal("tok1 should be gone after DeleteAllSessions")
	}
}

func TestUserSecrets(t *testing.T) {
	st := testStore(t)
	if _, err := st.UserSecret("alice"); !errors.Is(err, ErrNotFound) {
		t.Fatalf("want ErrNotFound, got %v", err)
	}
	if err := st.SetUserSecret("alice", "pw1"); err != nil {
		t.Fatal(err)
	}
	if err := st.SetUserSecret("alice", "pw2"); err != nil {
		t.Fatal(err)
	}
	pw, err := st.UserSecret("alice")
	if err != nil || pw != "pw2" {
		t.Fatalf("got %q %v", pw, err)
	}
	names, err := st.UserSecretNames()
	if err != nil || !names["alice"] {
		t.Fatalf("got %v %v", names, err)
	}
	if err := st.DeleteUserSecret("alice"); err != nil {
		t.Fatal(err)
	}
	if _, err := st.UserSecret("alice"); !errors.Is(err, ErrNotFound) {
		t.Fatal("secret should be deleted")
	}
}

func TestShareTokens(t *testing.T) {
	st := testStore(t)
	data := ShareTokenData{Username: "alice", Host: "1.2.3.4", Multiplexing: "MULTIPLEXING_MIDDLE"}

	if err := st.CreateShareToken("tok1", data, time.Now().Add(time.Hour)); err != nil {
		t.Fatal(err)
	}
	got, err := st.ShareToken("tok1")
	if err != nil || got != data {
		t.Fatalf("got %+v %v", got, err)
	}
	if _, err := st.ShareToken("nope"); !errors.Is(err, ErrNotFound) {
		t.Fatalf("want ErrNotFound, got %v", err)
	}

	// Expired tokens are dropped.
	if err := st.CreateShareToken("tok2", data, time.Now().Add(-time.Minute)); err != nil {
		t.Fatal(err)
	}
	if _, err := st.ShareToken("tok2"); !errors.Is(err, ErrNotFound) {
		t.Fatal("expired token should be gone")
	}

	// Invalidation by user drops every token for that user.
	if err := st.CreateShareToken("tok3", data, time.Now().Add(time.Hour)); err != nil {
		t.Fatal(err)
	}
	if err := st.DeleteShareTokensForUser("alice"); err != nil {
		t.Fatal(err)
	}
	if _, err := st.ShareToken("tok1"); !errors.Is(err, ErrNotFound) {
		t.Fatal("tok1 should be gone after user invalidation")
	}
}

func TestSubTokens(t *testing.T) {
	st := testStore(t)
	data := SubTokenData{Token: "tok1", Username: "alice", Host: "", Multiplexing: ""}

	if _, err := st.SubTokenForUser("alice"); !errors.Is(err, ErrNotFound) {
		t.Fatalf("want ErrNotFound, got %v", err)
	}
	if err := st.UpsertSubToken("tok1", data); err != nil {
		t.Fatal(err)
	}
	got, err := st.SubToken("tok1")
	if err != nil || got.Username != "alice" || got.Token != "tok1" {
		t.Fatalf("got %+v %v", got, err)
	}
	if got.CreatedAt.IsZero() {
		t.Fatal("CreatedAt not set")
	}
	if _, err := st.SubToken("nope"); !errors.Is(err, ErrNotFound) {
		t.Fatalf("want ErrNotFound, got %v", err)
	}

	// Rotation: a second upsert for the same user invalidates the old token.
	if err := st.UpsertSubToken("tok2", SubTokenData{Username: "alice", Host: "1.2.3.4"}); err != nil {
		t.Fatal(err)
	}
	if _, err := st.SubToken("tok1"); !errors.Is(err, ErrNotFound) {
		t.Fatal("old token should be gone after rotation")
	}
	got, err = st.SubTokenForUser("alice")
	if err != nil || got.Token != "tok2" || got.Host != "1.2.3.4" {
		t.Fatalf("got %+v %v", got, err)
	}

	names, err := st.SubTokenUsernames()
	if err != nil || !names["alice"] || len(names) != 1 {
		t.Fatalf("SubTokenUsernames: got %v %v", names, err)
	}

	if err := st.DeleteSubTokenForUser("alice"); err != nil {
		t.Fatal(err)
	}
	if _, err := st.SubToken("tok2"); !errors.Is(err, ErrNotFound) {
		t.Fatal("token should be gone after revocation")
	}
	names, err = st.SubTokenUsernames()
	if err != nil || len(names) != 0 {
		t.Fatalf("SubTokenUsernames after delete: got %v %v", names, err)
	}
}

func TestSettings(t *testing.T) {
	st := testStore(t)
	v, err := st.Setting("public_host")
	if err != nil || v != "" {
		t.Fatalf("unset setting: got %q %v", v, err)
	}
	if err := st.SetSetting("public_host", "1.2.3.4"); err != nil {
		t.Fatal(err)
	}
	if err := st.SetSetting("public_host", "5.6.7.8"); err != nil {
		t.Fatal(err)
	}
	v, _ = st.Setting("public_host")
	if v != "5.6.7.8" {
		t.Fatalf("got %q", v)
	}
}
