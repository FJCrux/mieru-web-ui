// SPDX-License-Identifier: GPL-3.0-or-later

package mitaclient

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	pb "github.com/enfein/mieru/v3/pkg/appctl/appctlpb"
	"google.golang.org/protobuf/proto"
)

// HealConfigFile fixes mita's on-disk config so it can start without
// crashlooping. mita couples users and ports (it rejects users without ports
// and crashes on `run` with ports but no users), so this rewrites the port
// bindings to match: ports (from spec) when there are users, none otherwise.
//
// It edits the protobuf config file directly (mita stores it as a plain
// proto.Marshal of ServerConfig), so it must run before mita is started. User
// entries and their hashed passwords are left untouched. Reports whether it
// changed anything.
func HealConfigFile(path, spec string) (bool, error) {
	b, err := os.ReadFile(path)
	if os.IsNotExist(err) {
		return false, nil // no config yet, nothing to heal
	}
	if err != nil {
		return false, err
	}
	cfg := &pb.ServerConfig{}
	if err := proto.Unmarshal(b, cfg); err != nil {
		return false, fmt.Errorf("parse %s: %w", path, err)
	}

	var want []*pb.PortBinding
	if len(cfg.GetUsers()) > 0 {
		// Users need ports. Only rewrite from a valid spec; otherwise leave the
		// existing (valid) ports so we never break a working config.
		if spec == "" {
			return false, nil
		}
		parsed, err := parsePortSpec(spec)
		if err != nil {
			return false, nil
		}
		want = parsed
	}
	// No users -> want stays nil (clear ports; this is the crashloop fix).
	if samePorts(cfg.GetPortBindings(), want) {
		return false, nil
	}
	cfg.PortBindings = want
	if len(want) > 0 && cfg.GetMtu() == 0 {
		cfg.Mtu = proto.Int32(1400)
	}
	nb, err := proto.Marshal(cfg)
	if err != nil {
		return false, err
	}
	return true, os.WriteFile(path, nb, 0o660)
}

func parsePortSpec(spec string) ([]*pb.PortBinding, error) {
	var out []*pb.PortBinding
	for _, tok := range strings.Split(spec, ",") {
		tok = strings.TrimSpace(tok)
		if tok == "" {
			continue
		}
		b := &pb.PortBinding{Protocol: pb.TransportProtocol_TCP.Enum()}
		if strings.Contains(tok, "-") {
			b.PortRange = proto.String(tok)
		} else {
			p, err := strconv.Atoi(tok)
			if err != nil {
				return nil, fmt.Errorf("invalid port %q", tok)
			}
			b.Port = proto.Int32(int32(p))
		}
		out = append(out, b)
	}
	if len(out) == 0 {
		return nil, fmt.Errorf("no ports")
	}
	return out, nil
}

func samePorts(a, b []*pb.PortBinding) bool {
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
