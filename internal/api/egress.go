// SPDX-License-Identifier: GPL-3.0-or-later

package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"strings"

	pb "github.com/enfein/mieru/v3/pkg/appctl/appctlpb"
	"google.golang.org/protobuf/proto"
)

// egressSettingKey stores the panel's semantic egress config (rules keep
// their country codes instead of the expanded CIDRs) so the UI stays clean
// and re-expands after a geoip update.
const egressSettingKey = "egress_config"

// egressProxy is one outbound SOCKS5 proxy (a hop in a chain).
type egressProxy struct {
	Name     string `json:"name"`
	Host     string `json:"host"`
	Port     int32  `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// egressRule routes matching traffic. Rules are evaluated in order; the
// first match wins, and the default when nothing matches is DIRECT.
type egressRule struct {
	Domains []string `json:"domains"`
	Cidrs   []string `json:"cidrs"`
	// Geo are geoip category codes (e.g. "ru", "ru-blocked") expanded to
	// CIDRs from the loaded datasets when pushed to mita. Kept separate so
	// the UI shows "ru" instead of thousands of ranges.
	Geo     []string `json:"geo"`
	Action  string   `json:"action"` // PROXY, DIRECT, REJECT
	Proxies []string `json:"proxies"`
}

type egressConfig struct {
	Proxies []egressProxy `json:"proxies"`
	Rules   []egressRule  `json:"rules"`
}

func (s *Server) handleGetEgress(w http.ResponseWriter, r *http.Request) {
	// The panel is the source of truth (it keeps country codes). Return the
	// stored semantic config; fall back to reading mita on first use.
	if raw, err := s.Store.Setting(egressSettingKey); err == nil && raw != "" {
		var out egressConfig
		if json.Unmarshal([]byte(raw), &out) == nil {
			writeJSON(w, http.StatusOK, out)
			return
		}
	}

	cfg, err := s.Mita.GetConfig(r.Context())
	if err != nil {
		writeMitaErr(w, err)
		return
	}
	out := egressConfig{Proxies: []egressProxy{}, Rules: []egressRule{}}
	eg := cfg.GetEgress()
	for _, p := range eg.GetProxies() {
		out.Proxies = append(out.Proxies, egressProxy{
			Name:     p.GetName(),
			Host:     p.GetHost(),
			Port:     p.GetPort(),
			Username: p.GetSocks5Authentication().GetUser(),
			Password: p.GetSocks5Authentication().GetPassword(),
		})
	}
	for _, rule := range eg.GetRules() {
		out.Rules = append(out.Rules, egressRule{
			Domains: orEmpty(rule.GetDomainNames()),
			Cidrs:   orEmpty(rule.GetIpRanges()),
			Geo:     []string{},
			Action:  rule.GetAction().String(),
			Proxies: orEmpty(rule.GetProxyNames()),
		})
	}
	writeJSON(w, http.StatusOK, out)
}

func (s *Server) handlePutEgress(w http.ResponseWriter, r *http.Request) {
	var req egressConfig
	if !readJSON(w, r, &req) {
		return
	}
	// Validate (and expand) before persisting.
	if _, err := s.buildEgress(req); err != nil {
		writeErr(w, http.StatusBadRequest, err.Error())
		return
	}
	if raw, err := json.Marshal(req); err == nil {
		_ = s.Store.SetSetting(egressSettingKey, string(raw))
	}
	if err := s.applyEgress(r.Context()); err != nil {
		writeMitaErr(w, err)
		return
	}
	writeJSON(w, http.StatusOK, map[string]bool{"ok": true})
}

// applyEgress rebuilds mita's egress from the stored semantic config plus the
// current chain peers, and pushes it. Called after egress or peer changes.
func (s *Server) applyEgress(ctx context.Context) error {
	var cfg egressConfig
	if raw, err := s.Store.Setting(egressSettingKey); err == nil && raw != "" {
		_ = json.Unmarshal([]byte(raw), &cfg)
	}
	egress, err := s.buildEgress(cfg)
	if err != nil {
		return err
	}
	_, err = s.Mita.UpdateConfig(ctx, func(c *pb.ServerConfig) error {
		c.Egress = egress
		return nil
	})
	return err
}

func (s *Server) buildEgress(in egressConfig) (*pb.Egress, error) {
	names := map[string]bool{}
	// Chain peers appear as egress proxies too, and rules may reference them.
	var peerProxies []*pb.EgressProxy
	if s.Peers != nil {
		peerProxies = s.Peers.EgressProxies()
		for name := range s.Peers.Names() {
			names[name] = true
		}
	}
	proxies := make([]*pb.EgressProxy, 0, len(in.Proxies)+len(peerProxies))
	for _, p := range in.Proxies {
		if p.Name == "" {
			return nil, fmt.Errorf("every outbound proxy needs a name")
		}
		if names[p.Name] {
			return nil, fmt.Errorf("duplicate proxy name %q", p.Name)
		}
		names[p.Name] = true
		if p.Host == "" {
			return nil, fmt.Errorf("proxy %q needs a host", p.Name)
		}
		if p.Port < 1 || p.Port > 65535 {
			return nil, fmt.Errorf("proxy %q has an invalid port", p.Name)
		}
		pb2 := &pb.EgressProxy{
			Name:     proto.String(p.Name),
			Protocol: pb.ProxyProtocol_SOCKS5_PROXY_PROTOCOL.Enum(),
			Host:     proto.String(p.Host),
			Port:     proto.Int32(p.Port),
		}
		if p.Username != "" || p.Password != "" {
			pb2.Socks5Authentication = &pb.Auth{
				User:     proto.String(p.Username),
				Password: proto.String(p.Password),
			}
		}
		proxies = append(proxies, pb2)
	}

	rules := make([]*pb.EgressRule, 0, len(in.Rules))
	for i, rule := range in.Rules {
		action, ok := pb.EgressAction_value[rule.Action]
		if !ok {
			return nil, fmt.Errorf("rule %d has an invalid action %q", i+1, rule.Action)
		}
		domains := cleanList(rule.Domains)
		cidrs := cleanList(rule.Cidrs)
		geo := cleanList(rule.Geo)
		if len(domains) == 0 && len(cidrs) == 0 && len(geo) == 0 {
			return nil, fmt.Errorf("rule %d matches nothing (add a domain, CIDR, or geo category, or use \"*\")", i+1)
		}
		for _, c := range cidrs {
			if c == "*" {
				continue
			}
			if _, _, err := net.ParseCIDR(c); err != nil {
				return nil, fmt.Errorf("rule %d has an invalid CIDR %q", i+1, c)
			}
		}
		// Expand each geo category into CIDRs from the datasets. mieru
		// matches on CIDRs; the panel does the geo lookup.
		for _, code := range geo {
			expanded, err := s.Geo.CIDRs(code)
			if err != nil {
				return nil, fmt.Errorf("rule %d: %v", i+1, err)
			}
			cidrs = append(cidrs, expanded...)
		}
		if pb.EgressAction(action) == pb.EgressAction_PROXY {
			if len(rule.Proxies) == 0 {
				return nil, fmt.Errorf("rule %d uses PROXY but names no proxy", i+1)
			}
			for _, name := range rule.Proxies {
				if !names[name] {
					return nil, fmt.Errorf("rule %d references unknown proxy %q", i+1, name)
				}
			}
		}
		rules = append(rules, &pb.EgressRule{
			DomainNames: domains,
			IpRanges:    cidrs,
			Action:      pb.EgressAction(action).Enum(),
			ProxyNames:  rule.Proxies,
		})
	}

	// Chain peers are always available as egress proxies.
	proxies = append(proxies, peerProxies...)

	if len(proxies) == 0 && len(rules) == 0 {
		return nil, nil // clearing egress
	}
	return &pb.Egress{Proxies: proxies, Rules: rules}, nil
}

func orEmpty(in []string) []string {
	if in == nil {
		return []string{}
	}
	return in
}

func cleanList(in []string) []string {
	out := make([]string, 0, len(in))
	for _, v := range in {
		v = strings.TrimSpace(v)
		if v != "" {
			out = append(out, v)
		}
	}
	return out
}
