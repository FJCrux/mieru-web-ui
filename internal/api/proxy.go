// SPDX-License-Identifier: GPL-3.0-or-later

package api

import (
	"context"
	"log"
	"time"

	pb "github.com/enfein/mieru/v3/pkg/appctl/appctlpb"
)

// mita couples users and ports: SetConfig rejects a config with users but no
// ports ("port binding is not set"), and `mita run` crashloops with ports but
// no users. So the panel keeps the desired ports in its own store and, on every
// config write, sets mita's bindings to match this invariant:
//
//	users present  -> ports = desired_ports
//	no users       -> no ports
//
// wantBindings returns the bindings a config should have. Because it depends on
// the users already in cfg, callers set it inside the same UpdateConfig that
// changes users, keeping the write atomic and always valid.
func (s *Server) wantBindings(cfg *pb.ServerConfig) []*pb.PortBinding {
	if len(cfg.GetUsers()) == 0 {
		return nil
	}
	spec, _ := s.Store.Setting("desired_ports")
	if spec == "" {
		return nil
	}
	b, err := parseSpecToBindings(spec)
	if err != nil {
		log.Printf("wantBindings: bad desired_ports %q: %v", spec, err)
		return nil
	}
	return b
}

// applyPortInvariant sets cfg.PortBindings for the invariant. Call it inside an
// UpdateConfig mutation after changing users.
func (s *Server) applyPortInvariant(cfg *pb.ServerConfig) {
	cfg.PortBindings = s.wantBindings(cfg)
	if len(cfg.PortBindings) > 0 && cfg.GetMtu() == 0 {
		cfg.Mtu = proto32(1400)
	}
}

// reconcileProxy fixes the ports invariant if needed (e.g. healing an old
// ports-without-users config) and brings the run state in line: proxying runs
// when there are users and ports, and is stopped otherwise.
//
// forceRestart=true does a Stop+Start even if already running, required after
// port changes since Reload opens new ports but never closes removed ones.
func (s *Server) reconcileProxy(ctx context.Context, forceRestart bool) {
	cfg, err := s.Mita.GetConfig(ctx)
	if err != nil {
		log.Printf("reconcileProxy: GetConfig: %v", err)
		return
	}
	target := s.wantBindings(cfg)
	if !samePortBindings(cfg.GetPortBindings(), target) {
		if _, err := s.Mita.UpdateConfig(ctx, func(c *pb.ServerConfig) error {
			s.applyPortInvariant(c)
			return nil
		}); err != nil {
			log.Printf("reconcileProxy: set ports: %v", err)
			return
		}
	}

	needRun := len(cfg.GetUsers()) > 0 && len(target) > 0
	status, err := s.Mita.Status(ctx)
	if err != nil {
		log.Printf("reconcileProxy: Status: %v", err)
		return
	}
	running := status == "RUNNING"

	switch {
	case !needRun:
		if running {
			if err := s.Mita.Stop(ctx); err != nil {
				log.Printf("reconcileProxy: Stop: %v", err)
			}
		}
	case forceRestart:
		if running {
			_ = s.Mita.Stop(ctx)
		}
		if err := s.Mita.Start(ctx); err != nil {
			log.Printf("reconcileProxy: Start: %v", err)
		}
	case !running:
		if err := s.Mita.Start(ctx); err != nil {
			log.Printf("reconcileProxy: Start: %v", err)
		}
	}
}

// StartupReconcile brings mita into a safe state on boot, retrying so it can
// heal a crashloop from an old ports-without-users config: mita's RPC is briefly
// up before it crashes, and clearing the ports once makes it stable.
func (s *Server) StartupReconcile(ctx context.Context) {
	for i := 0; i < 120; i++ {
		if _, err := s.Mita.Status(ctx); err == nil {
			s.reconcileProxy(ctx, false)
			if cfg, err := s.Mita.GetConfig(ctx); err == nil {
				// Safe once ports and users agree (both present or both absent).
				if (len(cfg.GetUsers()) > 0) == (len(cfg.GetPortBindings()) > 0) {
					return
				}
			}
		}
		select {
		case <-ctx.Done():
			return
		case <-time.After(400 * time.Millisecond):
		}
	}
}

func samePortBindings(a, b []*pb.PortBinding) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i].GetPort() != b[i].GetPort() ||
			a[i].GetPortRange() != b[i].GetPortRange() ||
			a[i].GetProtocol() != b[i].GetProtocol() {
			return false
		}
	}
	return true
}

func proto32(v int32) *int32 { return &v }
