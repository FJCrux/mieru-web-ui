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
	"golang.org/x/net/idna"
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
	Geo []string `json:"geo"`
	// Sites are geosite category codes (e.g. "telegram", "youtube") expanded
	// to domain suffixes from the loaded datasets when pushed to mita.
	Sites   []string `json:"sites"`
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
	// Chain peers are injected into mita's egress at build time (see
	// buildEgress); they are not part of the user-editable proxy list. Echoing
	// them back here would make the next PUT collide with the peer names.
	peerNames := map[string]bool{}
	if s.Peers != nil {
		peerNames = s.Peers.Names()
	}
	for _, p := range eg.GetProxies() {
		if peerNames[p.GetName()] {
			continue
		}
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
			Sites:   []string{},
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
	peerNames := map[string]bool{}
	var peerProxies []*pb.EgressProxy
	if s.Peers != nil {
		peerProxies = s.Peers.EgressProxies()
		peerNames = s.Peers.Names()
		for name := range peerNames {
			names[name] = true
		}
	}
	proxies := make([]*pb.EgressProxy, 0, len(in.Proxies)+len(peerProxies))
	for _, p := range in.Proxies {
		if p.Name == "" {
			return nil, fmt.Errorf("every outbound proxy needs a name")
		}
		if peerNames[p.Name] {
			return nil, fmt.Errorf("proxy name %q is already used by a chain peer (peers are added to egress automatically - reference %q in a rule instead of re-adding it)", p.Name, p.Name)
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
		for j, d := range domains {
			nd, err := normalizeDomain(d)
			if err != nil {
				return nil, fmt.Errorf("rule %d has an invalid domain %q: %v", i+1, d, err)
			}
			domains[j] = nd
		}
		cidrs := cleanList(rule.Cidrs)
		geo := cleanList(rule.Geo)
		sites := cleanList(rule.Sites)
		if len(domains) == 0 && len(cidrs) == 0 && len(geo) == 0 && len(sites) == 0 {
			return nil, fmt.Errorf("rule %d matches nothing (add a domain, CIDR, geoip, or geosite category, or use \"*\")", i+1)
		}
		for _, c := range cidrs {
			if c == "*" {
				continue
			}
			if _, _, err := net.ParseCIDR(c); err != nil {
				return nil, fmt.Errorf("rule %d has an invalid CIDR %q", i+1, c)
			}
		}
		// Expand each geoip category into CIDRs from the datasets. mieru
		// matches on CIDRs; the panel does the geo lookup.
		for _, code := range geo {
			expanded, err := s.Geo.CIDRs(code)
			if err != nil {
				return nil, fmt.Errorf("rule %d: %v", i+1, err)
			}
			cidrs = append(cidrs, expanded...)
		}
		// Expand each geosite category into domain suffixes. Entries come from
		// a large published list, so a single malformed domain is skipped
		// rather than failing the whole rule.
		for _, code := range sites {
			expanded, err := s.Geo.Domains(code)
			if err != nil {
				return nil, fmt.Errorf("rule %d: %v", i+1, err)
			}
			for _, d := range expanded {
				if nd, err := normalizeDomain(d); err == nil {
					domains = append(domains, nd)
				}
			}
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

// normalizeDomain converts a UI domain pattern into what mita actually
// matches. mita has no wildcards other than a bare "*": an entry matches the
// domain exactly or as a dot-separated suffix ("ru" matches "2ip.ru"). So
// "*.example.com" is rewritten to "example.com", and IDN names are converted
// to punycode because clients send ASCII FQDNs ("рф" would never match).
func normalizeDomain(d string) (string, error) {
	if d == "*" {
		return d, nil
	}
	d = strings.ToLower(strings.TrimSuffix(d, "."))
	d = strings.TrimPrefix(d, "*.")
	d = strings.TrimPrefix(d, ".")
	if d == "" {
		return "", fmt.Errorf("empty domain")
	}
	ascii, err := idna.Lookup.ToASCII(d)
	if err != nil {
		return "", err
	}
	return ascii, nil
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
