// SPDX-License-Identifier: GPL-3.0-or-later

package chain

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/enfein/mieru/v3/pkg/appctl"
	pb "github.com/enfein/mieru/v3/pkg/appctl/appctlpb"
	"google.golang.org/protobuf/proto"

	"github.com/fjcrux/mieru-web-ui/internal/store"
)

func testKey(t *testing.T) string {
	t.Helper()
	cfg := &pb.ClientConfig{
		Profiles: []*pb.ClientProfile{{
			ProfileName: proto.String("default"),
			User:        &pb.User{Name: proto.String("u"), Password: proto.String("p")},
			Servers: []*pb.ServerEndpoint{{
				IpAddress:    proto.String("1.2.3.4"),
				PortBindings: []*pb.PortBinding{{Port: proto.Int32(2012), Protocol: pb.TransportProtocol_TCP.Enum()}},
			}},
		}},
		ActiveProfile: proto.String("default"),
	}
	key, err := appctl.ClientConfigToURL(cfg)
	if err != nil {
		t.Fatal(err)
	}
	return key
}

func newTestManager(t *testing.T) (*Manager, store.Store) {
	t.Helper()
	st, err := store.OpenSQLite(filepath.Join(t.TempDir(), "panel.db"))
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { st.Close() })
	// /bin/true as a stand-in binary; without Start(ctx) no process is run.
	m, err := New("/bin/true", t.TempDir(), st)
	if err != nil {
		t.Fatal(err)
	}
	return m, st
}

func TestAddPersistsAndAssignsPorts(t *testing.T) {
	m, st := newTestManager(t)
	key := testKey(t)

	if err := m.Add("first", key); err != nil {
		t.Fatal(err)
	}
	if err := m.Add("second", key); err != nil {
		t.Fatal(err)
	}
	// Duplicate name rejected.
	if err := m.Add("first", key); err == nil {
		t.Fatal("expected duplicate name error")
	}
	// Invalid key rejected.
	if err := m.Add("bad", "not-a-key"); err == nil {
		t.Fatal("expected invalid key error")
	}

	peers, err := st.Peers()
	if err != nil {
		t.Fatal(err)
	}
	if len(peers) != 2 {
		t.Fatalf("got %d peers", len(peers))
	}
	byName := map[string]int32{}
	for _, p := range peers {
		byName[p.Name] = p.Socks5Port
		if _, err := os.Stat(filepath.Join(m.dir, p.Name+".json")); err != nil {
			t.Fatalf("config for %s missing: %v", p.Name, err)
		}
	}
	if byName["first"] != socks5PortBase || byName["second"] != socks5PortBase+1 {
		t.Fatalf("unexpected ports: %v", byName)
	}
}

func TestRemove(t *testing.T) {
	m, st := newTestManager(t)
	if err := m.Add("p", testKey(t)); err != nil {
		t.Fatal(err)
	}
	if err := m.Remove("p"); err != nil {
		t.Fatal(err)
	}
	peers, _ := st.Peers()
	if len(peers) != 0 {
		t.Fatalf("peer not removed: %+v", peers)
	}
	if _, err := os.Stat(filepath.Join(m.dir, "p.json")); !os.IsNotExist(err) {
		t.Fatal("config file not removed")
	}
}
