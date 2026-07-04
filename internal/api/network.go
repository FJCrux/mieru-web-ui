// SPDX-License-Identifier: GPL-3.0-or-later

package api

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	pb "github.com/enfein/mieru/v3/pkg/appctl/appctlpb"
	"google.golang.org/protobuf/proto"
)

type portBinding struct {
	// Port is set for a single port, PortRange ("2012-2022") for a range.
	Port      int32  `json:"port,omitempty"`
	PortRange string `json:"portRange,omitempty"`
	Protocol  string `json:"protocol"` // TCP or UDP
}

type networkConfig struct {
	PortBindings []portBinding `json:"portBindings"`
	MTU          int32         `json:"mtu"`
	LoggingLevel string        `json:"loggingLevel"`
	// PortsManaged is true when ports come from PROXY_PORTS (read-only in UI).
	PortsManaged bool `json:"portsManaged"`
}

func (s *Server) handleGetNetwork(w http.ResponseWriter, r *http.Request) {
	out := networkConfig{
		PortBindings: []portBinding{},
		MTU:          1400,
		LoggingLevel: "INFO",
		PortsManaged: s.PortsManaged,
	}
	// Ports come from the panel's desired_ports (shown even with no users, and
	// safe when mita is momentarily unreachable). MTU/logging come from mita.
	if spec, _ := s.Store.Setting("desired_ports"); spec != "" {
		if bindings, err := parseSpecToBindings(spec); err == nil {
			for _, b := range bindings {
				out.PortBindings = append(out.PortBindings, portBinding{
					Port:      b.GetPort(),
					PortRange: b.GetPortRange(),
					Protocol:  b.GetProtocol().String(),
				})
			}
		}
	}
	if cfg, err := s.Mita.GetConfig(r.Context()); err == nil {
		out.MTU = cfg.GetMtu()
		out.LoggingLevel = cfg.GetLoggingLevel().String()
	}
	writeJSON(w, http.StatusOK, out)
}

// handlePutNetwork applies port/MTU/logging changes. Per the M1 probe,
// Reload only opens new ports and never closes removed ones, so this
// path uses Stop+Start (brief proxy interruption) to fully apply.
func (s *Server) handlePutNetwork(w http.ResponseWriter, r *http.Request) {
	var req networkConfig
	if !readJSON(w, r, &req) {
		return
	}
	// Ports are stored as desired_ports (applied to mita only when a user
	// exists). When PROXY_PORTS manages them, the UI can't change them.
	if !s.PortsManaged {
		spec, err := bindingsToSpec(req.PortBindings)
		if err != nil {
			writeErr(w, http.StatusBadRequest, err.Error())
			return
		}
		if err := s.Store.SetSetting("desired_ports", spec); err != nil {
			writeErr(w, http.StatusInternalServerError, err.Error())
			return
		}
	}
	if req.MTU != 0 && (req.MTU < 1280 || req.MTU > 1500) {
		writeErr(w, http.StatusBadRequest, "mtu must be between 1280 and 1500")
		return
	}
	level, ok := pb.LoggingLevel_value[req.LoggingLevel]
	if req.LoggingLevel != "" && !ok {
		writeErr(w, http.StatusBadRequest, "invalid logging level")
		return
	}

	ctx := r.Context()
	if req.MTU != 0 || req.LoggingLevel != "" {
		if _, err := s.Mita.UpdateConfig(ctx, func(cfg *pb.ServerConfig) error {
			if req.MTU != 0 {
				cfg.Mtu = proto.Int32(req.MTU)
			}
			if req.LoggingLevel != "" {
				cfg.LoggingLevel = pb.LoggingLevel(level).Enum()
			}
			return nil
		}); err != nil {
			writeMitaErr(w, err)
			return
		}
	}
	// Apply the ports (and restart) based on whether users exist.
	s.reconcileProxy(ctx, true)
	writeJSON(w, http.StatusOK, map[string]bool{"ok": true})
}

// bindingsToSpec turns the UI port bindings into a desired_ports spec string,
// validating each entry.
func bindingsToSpec(in []portBinding) (string, error) {
	if len(in) == 0 {
		return "", nil
	}
	var toks []string
	for _, b := range in {
		if b.Protocol != "" && b.Protocol != "TCP" {
			return "", fmt.Errorf("only TCP port bindings are supported here")
		}
		if b.PortRange != "" {
			if _, err := parseSpecToBindings(b.PortRange); err != nil {
				return "", err
			}
			toks = append(toks, b.PortRange)
		} else {
			if err := validPort(b.Port); err != nil {
				return "", err
			}
			toks = append(toks, strconv.Itoa(int(b.Port)))
		}
	}
	spec := strings.Join(toks, ",")
	if _, err := parseSpecToBindings(spec); err != nil {
		return "", err
	}
	return spec, nil
}

// parseSpecToBindings parses "2012", "2012-2022", or "2012,2020,2027" into TCP
// port bindings.
func parseSpecToBindings(spec string) ([]*pb.PortBinding, error) {
	var out []*pb.PortBinding
	for _, tok := range strings.Split(spec, ",") {
		tok = strings.TrimSpace(tok)
		if tok == "" {
			continue
		}
		b := &pb.PortBinding{Protocol: pb.TransportProtocol_TCP.Enum()}
		if strings.Contains(tok, "-") {
			lo, hi, err := parsePortRange(tok)
			if err != nil {
				return nil, err
			}
			_ = lo
			_ = hi
			b.PortRange = proto.String(tok)
		} else {
			p, err := strconv.Atoi(tok)
			if err != nil || validPort(int32(p)) != nil {
				return nil, fmt.Errorf("invalid port %q", tok)
			}
			b.Port = proto.Int32(int32(p))
		}
		out = append(out, b)
	}
	if len(out) == 0 {
		return nil, fmt.Errorf("no ports in %q", spec)
	}
	return out, nil
}

func validPort(p int32) error {
	if p < 1025 || p > 65535 {
		return fmt.Errorf("port %d out of range (1025-65535)", p)
	}
	return nil
}

func parsePortRange(s string) (int32, int32, error) {
	parts := strings.SplitN(s, "-", 2)
	if len(parts) != 2 {
		return 0, 0, fmt.Errorf("port range must look like 2012-2022")
	}
	lo, err1 := strconv.Atoi(strings.TrimSpace(parts[0]))
	hi, err2 := strconv.Atoi(strings.TrimSpace(parts[1]))
	if err1 != nil || err2 != nil {
		return 0, 0, fmt.Errorf("port range must look like 2012-2022")
	}
	if err := validPort(int32(lo)); err != nil {
		return 0, 0, err
	}
	if err := validPort(int32(hi)); err != nil {
		return 0, 0, err
	}
	if lo >= hi {
		return 0, 0, fmt.Errorf("port range start must be below end")
	}
	return int32(lo), int32(hi), nil
}
