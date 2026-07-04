// SPDX-License-Identifier: GPL-3.0-or-later

// Package sharelink builds mieru client configs and share links
// (mieru:// and mierus://) for a given server user, reusing the
// official encoders from the mieru module.
package sharelink

import (
	"fmt"
	"net"

	"github.com/enfein/mieru/v3/pkg/appctl"
	pb "github.com/enfein/mieru/v3/pkg/appctl/appctlpb"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// Params describes one user's client setup.
type Params struct {
	Username string
	Password string
	// Host is the public address clients connect to (IP or domain).
	Host string
	// PortBindings are copied from the server config.
	PortBindings []*pb.PortBinding
	MTU          int32
	Multiplexing string // MULTIPLEXING_OFF/LOW/MIDDLE/HIGH; default MIDDLE
}

// Links is what the panel's share endpoint returns.
type Links struct {
	// ClientConfigJSON is a complete `mieru apply config` file.
	ClientConfigJSON string `json:"clientConfigJson"`
	// MieruURL is the mieru:// link (base64 protobuf, full fidelity).
	MieruURL string `json:"mieruUrl"`
	// MierusURLs are human-readable mierus:// links, one per server.
	MierusURLs []string `json:"mierusUrls"`
}

// Build assembles the client config and both link formats.
func Build(p Params) (*Links, error) {
	if p.Username == "" || p.Password == "" {
		return nil, fmt.Errorf("username and password are required")
	}
	if p.Host == "" {
		return nil, fmt.Errorf("public host is not configured")
	}
	if len(p.PortBindings) == 0 {
		return nil, fmt.Errorf("server has no port bindings")
	}
	if p.MTU == 0 {
		p.MTU = 1400
	}
	mux, ok := pb.MultiplexingLevel_value[p.Multiplexing]
	if !ok || p.Multiplexing == "" {
		mux = int32(pb.MultiplexingLevel_MULTIPLEXING_MIDDLE)
	}

	server := &pb.ServerEndpoint{PortBindings: p.PortBindings}
	// The proto has separate fields for IP address and domain name.
	if ipLike(p.Host) {
		server.IpAddress = proto.String(p.Host)
	} else {
		server.DomainName = proto.String(p.Host)
	}

	profile := &pb.ClientProfile{
		ProfileName: proto.String("default"),
		User: &pb.User{
			Name:     proto.String(p.Username),
			Password: proto.String(p.Password),
		},
		Servers: []*pb.ServerEndpoint{server},
		Mtu:     proto.Int32(p.MTU),
		Multiplexing: &pb.MultiplexingConfig{
			Level: pb.MultiplexingLevel(mux).Enum(),
		},
	}
	cfg := &pb.ClientConfig{
		Profiles:      []*pb.ClientProfile{profile},
		ActiveProfile: proto.String("default"),
		RpcPort:       proto.Int32(8964),
		Socks5Port:    proto.Int32(1080),
		LoggingLevel:  pb.LoggingLevel_INFO.Enum(),
		HttpProxyPort: proto.Int32(8080),
	}

	cfgJSON, err := protojson.MarshalOptions{Multiline: true, Indent: "    "}.Marshal(cfg)
	if err != nil {
		return nil, fmt.Errorf("marshal client config: %w", err)
	}
	mieruURL, err := appctl.ClientConfigToURL(cfg)
	if err != nil {
		return nil, fmt.Errorf("build mieru:// link: %w", err)
	}
	mierusURLs, err := appctl.ClientProfileToMultiURLs(profile)
	if err != nil {
		return nil, fmt.Errorf("build mierus:// links: %w", err)
	}

	return &Links{
		ClientConfigJSON: string(cfgJSON),
		MieruURL:         mieruURL,
		MierusURLs:       mierusURLs,
	}, nil
}

// ipLike reports whether s parses as an IP address.
func ipLike(s string) bool {
	return net.ParseIP(s) != nil
}
