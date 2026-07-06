// SPDX-License-Identifier: GPL-3.0-or-later

package clashsub

import (
	"strings"
	"testing"

	pb "github.com/enfein/mieru/v3/pkg/appctl/appctlpb"
	yaml "go.yaml.in/yaml/v2"
	"google.golang.org/protobuf/proto"
)

func tcpPort(p int32) *pb.PortBinding {
	return &pb.PortBinding{Port: proto.Int32(p), Protocol: pb.TransportProtocol_TCP.Enum()}
}

func tcpRange(r string) *pb.PortBinding {
	return &pb.PortBinding{PortRange: proto.String(r), Protocol: pb.TransportProtocol_TCP.Enum()}
}

func params(bindings ...*pb.PortBinding) Params {
	return Params{
		Username:     "alice",
		Password:     "secret",
		Host:         "203.0.113.10",
		PortBindings: bindings,
	}
}

// unmarshal proves the output is valid YAML and returns it for inspection.
func unmarshal(t *testing.T, b []byte) map[string]any {
	t.Helper()
	var m map[string]any
	if err := yaml.Unmarshal(b, &m); err != nil {
		t.Fatalf("output is not valid YAML: %v\n%s", err, b)
	}
	return m
}

func TestBuildProfileSinglePort(t *testing.T) {
	out, err := BuildProfile(params(tcpPort(2012)))
	if err != nil {
		t.Fatal(err)
	}
	s := string(out)
	for _, want := range []string{
		"name: alice", "type: mieru", "server: 203.0.113.10", "port: 2012",
		"transport: TCP", "udp: true", "username: alice", "password: secret",
		"MATCH,PROXY", "mixed-port: 7890",
	} {
		if !strings.Contains(s, want) {
			t.Errorf("profile missing %q:\n%s", want, s)
		}
	}
	if strings.Contains(s, "multiplexing") {
		t.Errorf("multiplexing key should be omitted when unset:\n%s", s)
	}
	if strings.Contains(s, "port-range") {
		t.Errorf("port-range should be absent for a single port:\n%s", s)
	}
	if strings.Contains(s, "AUTO") {
		t.Errorf("AUTO group should be absent with one proxy:\n%s", s)
	}
	m := unmarshal(t, out)
	if _, ok := m["proxy-groups"]; !ok {
		t.Error("proxy-groups missing")
	}
}

func TestBuildProfilePortRange(t *testing.T) {
	out, err := BuildProfile(params(tcpRange("2090-2099")))
	if err != nil {
		t.Fatal(err)
	}
	s := string(out)
	if !strings.Contains(s, "port-range: 2090-2099") {
		t.Errorf("missing port-range:\n%s", s)
	}
	m := unmarshal(t, out)
	px := m["proxies"].([]any)[0].(map[any]any)
	if _, ok := px["port"]; ok {
		t.Errorf("port should be absent when port-range is set:\n%s", s)
	}
}

func TestBuildProfileMultipleBindings(t *testing.T) {
	out, err := BuildProfile(params(tcpPort(2012), tcpRange("2090-2099")))
	if err != nil {
		t.Fatal(err)
	}
	s := string(out)
	for _, want := range []string{"name: alice-1", "name: alice-2", "name: AUTO", "url-test"} {
		if !strings.Contains(s, want) {
			t.Errorf("profile missing %q:\n%s", want, s)
		}
	}
	unmarshal(t, out)
}

func TestBuildProfileMultiplexing(t *testing.T) {
	p := params(tcpPort(2012))
	p.Multiplexing = "MULTIPLEXING_LOW"
	out, err := BuildProfile(p)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(out), "multiplexing: MULTIPLEXING_LOW") {
		t.Errorf("missing multiplexing:\n%s", out)
	}

	p.Multiplexing = "MULTIPLEXING_BOGUS"
	if _, err := BuildProfile(p); err == nil {
		t.Error("expected error for unknown multiplexing level")
	}
}

func TestBuildProxiesOnly(t *testing.T) {
	out, err := BuildProxies(params(tcpPort(2012)))
	if err != nil {
		t.Fatal(err)
	}
	s := string(out)
	if !strings.Contains(s, "type: mieru") {
		t.Errorf("missing proxy entry:\n%s", s)
	}
	for _, absent := range []string{"proxy-groups", "rules", "mixed-port"} {
		if strings.Contains(s, absent) {
			t.Errorf("proxies-only output should not contain %q:\n%s", absent, s)
		}
	}
	unmarshal(t, out)
}

func TestNonTCPFiltered(t *testing.T) {
	udp := &pb.PortBinding{Port: proto.Int32(2012), Protocol: pb.TransportProtocol_UDP.Enum()}
	out, err := BuildProfile(params(udp, tcpPort(2013)))
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(out), "port: 2013") || strings.Contains(string(out), "port: 2012") {
		t.Errorf("expected only the TCP binding:\n%s", out)
	}

	if _, err := BuildProfile(params(udp)); err == nil {
		t.Error("expected error when no TCP bindings remain")
	}
}

func TestValidation(t *testing.T) {
	cases := []Params{
		{Password: "x", Host: "h", PortBindings: []*pb.PortBinding{tcpPort(1)}},
		{Username: "u", Host: "h", PortBindings: []*pb.PortBinding{tcpPort(1)}},
		{Username: "u", Password: "x", PortBindings: []*pb.PortBinding{tcpPort(1)}},
		{Username: "u", Password: "x", Host: "h"},
	}
	for i, p := range cases {
		if _, err := BuildProfile(p); err == nil {
			t.Errorf("case %d: expected validation error", i)
		}
	}
}
