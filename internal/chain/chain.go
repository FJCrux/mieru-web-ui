// SPDX-License-Identifier: GPL-3.0-or-later

// Package chain links this panel to an upstream mieru server ("peer") so its
// egress can be routed through it. mieru's egress only dials SOCKS5, and a
// mita server is not a SOCKS5 endpoint, so for each peer the panel runs a
// supervised mieru client that connects to the peer and exposes a local
// SOCKS5 port. That port is then registered as an egress proxy.
//
// A peer is added from a "peer key" (a mieru:// client config the upstream
// panel hands out); the panel decodes it, assigns local ports, and starts
// the client. No manual client setup required.
package chain

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sync"

	"github.com/enfein/mieru/v3/pkg/appctl"
	pb "github.com/enfein/mieru/v3/pkg/appctl/appctlpb"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"

	"github.com/fjcrux/mieru-web-ui/internal/store"
	"github.com/fjcrux/mieru-web-ui/internal/supervisor"
)

const (
	socks5PortBase = 42000
	rpcPortBase    = 43000
	maxPeers       = 100
)

var nameRe = regexp.MustCompile(`^[a-zA-Z0-9._-]{1,40}$`)

// Manager runs and tracks upstream chain peers.
type Manager struct {
	binary string // mieru client binary
	dir    string // per-peer config directory
	store  store.Store

	mu    sync.Mutex
	ctx   context.Context
	procs map[string]*proc
}

type proc struct {
	sup        *supervisor.Supervisor
	cancel     context.CancelFunc
	socks5Port int32
}

// Status is what the API reports for a peer.
type Status struct {
	Name       string `json:"name"`
	Socks5Port int32  `json:"socks5Port"`
	Running    bool   `json:"running"`
	Restarts   int    `json:"restarts"`
}

// New creates a Manager. clientBinary is the mieru client; dir holds the
// generated per-peer configs.
func New(clientBinary, dir string, st store.Store) (*Manager, error) {
	if err := os.MkdirAll(dir, 0o700); err != nil {
		return nil, fmt.Errorf("create peers dir: %w", err)
	}
	return &Manager{
		binary: clientBinary,
		dir:    dir,
		store:  st,
		procs:  map[string]*proc{},
	}, nil
}

// Start launches the clients for all stored peers and keeps them supervised
// for the lifetime of ctx.
func (m *Manager) Start(ctx context.Context) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.ctx = ctx
	peers, err := m.store.Peers()
	if err != nil {
		return err
	}
	for _, p := range peers {
		m.launch(p.Name, p.Socks5Port)
	}
	return nil
}

// Add decodes a peer key, assigns local ports, starts the client, and
// persists the peer. It does not touch mita's egress; the caller re-applies
// egress afterwards (via EgressProxies).
func (m *Manager) Add(name, key string) error {
	if !nameRe.MatchString(name) {
		return fmt.Errorf("invalid peer name (letters, digits, . _ -, max 40)")
	}
	cfg, err := appctl.URLToClientConfig(key)
	if err != nil {
		return fmt.Errorf("invalid peer key: %w", err)
	}
	if len(cfg.GetProfiles()) == 0 {
		return fmt.Errorf("peer key has no profile")
	}

	m.mu.Lock()
	defer m.mu.Unlock()
	if _, exists := m.procs[name]; exists {
		return fmt.Errorf("peer %q already exists", name)
	}
	peers, err := m.store.Peers()
	if err != nil {
		return err
	}
	if len(peers) >= maxPeers {
		return fmt.Errorf("too many peers")
	}
	for _, p := range peers {
		if p.Name == name {
			return fmt.Errorf("peer %q already exists", name)
		}
	}
	socks5Port, rpcPort := m.freePorts(peers)

	// Force local-only listeners on the assigned ports; disable the HTTP
	// proxy so peers can't collide on it.
	cfg.Socks5Port = proto.Int32(socks5Port)
	cfg.RpcPort = proto.Int32(rpcPort)
	cfg.Socks5ListenLAN = proto.Bool(false)
	cfg.HttpProxyPort = nil

	confJSON, err := protojson.Marshal(cfg)
	if err != nil {
		return fmt.Errorf("marshal client config: %w", err)
	}
	if err := os.WriteFile(m.configPath(name), confJSON, 0o600); err != nil {
		return err
	}
	if err := m.store.CreatePeer(store.PeerRecord{
		Name: name, ConfigJSON: string(confJSON), Socks5Port: socks5Port,
	}); err != nil {
		os.Remove(m.configPath(name))
		return err
	}
	m.launch(name, socks5Port)
	return nil
}

// Remove stops a peer's client and deletes it.
func (m *Manager) Remove(name string) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	if p, ok := m.procs[name]; ok {
		p.cancel()
		delete(m.procs, name)
	}
	os.Remove(m.configPath(name))
	return m.store.DeletePeer(name)
}

// List reports every peer's status.
func (m *Manager) List() []Status {
	m.mu.Lock()
	defer m.mu.Unlock()
	out := []Status{}
	for name, p := range m.procs {
		restarts, startedAt := p.sup.Stats()
		out = append(out, Status{
			Name:       name,
			Socks5Port: p.socks5Port,
			Running:    !startedAt.IsZero(),
			Restarts:   restarts,
		})
	}
	return out
}

// EgressProxies returns an egress proxy per peer, pointing at its local
// SOCKS5 port. The caller merges these into mita's egress config.
func (m *Manager) EgressProxies() []*pb.EgressProxy {
	m.mu.Lock()
	defer m.mu.Unlock()
	out := make([]*pb.EgressProxy, 0, len(m.procs))
	for name, p := range m.procs {
		out = append(out, &pb.EgressProxy{
			Name:     proto.String(name),
			Protocol: pb.ProxyProtocol_SOCKS5_PROXY_PROTOCOL.Enum(),
			Host:     proto.String("127.0.0.1"),
			Port:     proto.Int32(p.socks5Port),
		})
	}
	return out
}

// Names returns the peer names (used to reserve egress-proxy names).
func (m *Manager) Names() map[string]bool {
	m.mu.Lock()
	defer m.mu.Unlock()
	out := map[string]bool{}
	for name := range m.procs {
		out[name] = true
	}
	return out
}

// launch starts a supervised client for a peer. Caller holds m.mu.
func (m *Manager) launch(name string, socks5Port int32) {
	if m.ctx == nil {
		return // Start not called yet
	}
	// Only the JSON config path: the client loads MIERU_CONFIG_FILE
	// (protobuf) with priority, so setting it to a non-existent path would
	// make `mieru run` fail before it reads our JSON.
	sup := supervisor.New("chain:"+name, m.binary, []string{
		"MIERU_CONFIG_JSON_FILE=" + m.configPath(name),
	})
	ctx, cancel := context.WithCancel(m.ctx)
	m.procs[name] = &proc{sup: sup, cancel: cancel, socks5Port: socks5Port}
	go sup.Run(ctx)
}

func (m *Manager) configPath(name string) string {
	return filepath.Join(m.dir, name+".json")
}

// freePorts returns the lowest unused socks5/rpc port pair. Caller holds m.mu.
func (m *Manager) freePorts(peers []store.PeerRecord) (int32, int32) {
	used := map[int32]bool{}
	for _, p := range peers {
		used[p.Socks5Port] = true
	}
	for i := int32(0); i < maxPeers; i++ {
		if !used[socks5PortBase+i] {
			return socks5PortBase + i, rpcPortBase + i
		}
	}
	return socks5PortBase, rpcPortBase
}
