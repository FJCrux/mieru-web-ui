// SPDX-License-Identifier: GPL-3.0-or-later

// Package clashsub renders Clash/Mihomo subscription YAML for a mieru user.
// mihomo supports `type: mieru` outbounds since v1.19.0.
package clashsub

import (
	"fmt"

	pb "github.com/enfein/mieru/v3/pkg/appctl/appctlpb"
	yaml "go.yaml.in/yaml/v2"
)

// Params describes one user's client setup, mirroring sharelink.Params.
type Params struct {
	Username string
	Password string
	// Host is the public address clients connect to (IP or domain).
	Host string
	// PortBindings are copied from the server config; only TCP bindings
	// are used, others are skipped.
	PortBindings []*pb.PortBinding
	Multiplexing string // MULTIPLEXING_OFF/LOW/MIDDLE/HIGH; "" omits the key
}

type proxy struct {
	Name         string `yaml:"name"`
	Type         string `yaml:"type"`
	Server       string `yaml:"server"`
	Port         int32  `yaml:"port,omitempty"`
	PortRange    string `yaml:"port-range,omitempty"`
	Transport    string `yaml:"transport"`
	UDP          bool   `yaml:"udp"`
	Username     string `yaml:"username"`
	Password     string `yaml:"password"`
	Multiplexing string `yaml:"multiplexing,omitempty"`
}

type group struct {
	Name     string   `yaml:"name"`
	Type     string   `yaml:"type"`
	Proxies  []string `yaml:"proxies"`
	URL      string   `yaml:"url,omitempty"`
	Interval int      `yaml:"interval,omitempty"`
}

type profile struct {
	MixedPort   int      `yaml:"mixed-port"`
	AllowLan    bool     `yaml:"allow-lan"`
	Mode        string   `yaml:"mode"`
	LogLevel    string   `yaml:"log-level"`
	Proxies     []proxy  `yaml:"proxies"`
	ProxyGroups []group  `yaml:"proxy-groups"`
	Rules       []string `yaml:"rules"`
}

type proxiesOnly struct {
	Proxies []proxy `yaml:"proxies"`
}

// Private-range rules keep LAN traffic direct without requiring any geo
// database on the client.
var baseRules = []string{
	"IP-CIDR,127.0.0.0/8,DIRECT,no-resolve",
	"IP-CIDR,10.0.0.0/8,DIRECT,no-resolve",
	"IP-CIDR,172.16.0.0/12,DIRECT,no-resolve",
	"IP-CIDR,192.168.0.0/16,DIRECT,no-resolve",
	"IP-CIDR,100.64.0.0/10,DIRECT,no-resolve",
	"IP-CIDR6,::1/128,DIRECT,no-resolve",
	"IP-CIDR6,fc00::/7,DIRECT,no-resolve",
	"MATCH,PROXY",
}

// BuildProfile renders a complete minimal profile (proxies + groups + rules)
// that a Clash client can import and use as-is.
func BuildProfile(p Params) ([]byte, error) {
	proxies, err := buildProxies(p)
	if err != nil {
		return nil, err
	}

	names := make([]string, len(proxies))
	for i, px := range proxies {
		names[i] = px.Name
	}
	groups := []group{{Name: "PROXY", Type: "select", Proxies: names}}
	if len(proxies) > 1 {
		auto := group{
			Name:     "AUTO",
			Type:     "url-test",
			Proxies:  names,
			URL:      "https://cp.cloudflare.com/generate_204",
			Interval: 300,
		}
		groups = []group{
			{Name: "PROXY", Type: "select", Proxies: append([]string{"AUTO"}, names...)},
			auto,
		}
	}

	return yaml.Marshal(profile{
		MixedPort:   7890,
		AllowLan:    false,
		Mode:        "rule",
		LogLevel:    "info",
		Proxies:     proxies,
		ProxyGroups: groups,
		Rules:       baseRules,
	})
}

// BuildProxies renders a proxies-only document for use with proxy-provider.
func BuildProxies(p Params) ([]byte, error) {
	proxies, err := buildProxies(p)
	if err != nil {
		return nil, err
	}
	return yaml.Marshal(proxiesOnly{Proxies: proxies})
}

// buildProxies validates params and produces one proxy entry per TCP port
// binding (a Clash entry holds exactly one port or port-range).
func buildProxies(p Params) ([]proxy, error) {
	if p.Username == "" || p.Password == "" {
		return nil, fmt.Errorf("username and password are required")
	}
	if p.Host == "" {
		return nil, fmt.Errorf("public host is not configured")
	}
	if p.Multiplexing != "" {
		if _, ok := pb.MultiplexingLevel_value[p.Multiplexing]; !ok {
			return nil, fmt.Errorf("unknown multiplexing level %q", p.Multiplexing)
		}
	}

	var tcp []*pb.PortBinding
	for _, b := range p.PortBindings {
		if b.GetProtocol() == pb.TransportProtocol_TCP {
			tcp = append(tcp, b)
		}
	}
	if len(tcp) == 0 {
		return nil, fmt.Errorf("server has no TCP port bindings")
	}

	proxies := make([]proxy, 0, len(tcp))
	for i, b := range tcp {
		name := p.Username
		if len(tcp) > 1 {
			name = fmt.Sprintf("%s-%d", p.Username, i+1)
		}
		px := proxy{
			Name:         name,
			Type:         "mieru",
			Server:       p.Host,
			Transport:    "TCP",
			UDP:          true,
			Username:     p.Username,
			Password:     p.Password,
			Multiplexing: p.Multiplexing,
		}
		if r := b.GetPortRange(); r != "" {
			px.PortRange = r
		} else {
			px.Port = b.GetPort()
		}
		proxies = append(proxies, px)
	}
	return proxies, nil
}
