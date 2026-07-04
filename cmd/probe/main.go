// SPDX-License-Identifier: GPL-3.0-or-later

// Command probe checks that a mita daemon behaves the way the panel expects.
// It is a development tool, not part of the shipped image.
//
// It verifies that:
//   - SetConfig replaces the whole config rather than merging users
//   - a user's hashed password survives a GetConfig -> SetConfig round-trip
//   - port binding changes take effect after Reload
//
// and prints the shapes of GetMetrics, GetUsers, and GetSessionInfoList.
//
// Usage: MITA_RPC_TARGET=tcp://localhost:9090 go run ./cmd/probe
package main

import (
	"context"
	"fmt"
	"net"
	"os"
	"time"

	pb "github.com/enfein/mieru/v3/pkg/appctl/appctlpb"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"

	"github.com/fjcrux/mieru-web-ui/internal/mitaclient"
)

var pj = protojson.MarshalOptions{Multiline: true, Indent: "  ", EmitUnpopulated: false}

func dump(label string, m proto.Message) {
	b, err := pj.Marshal(m)
	if err != nil {
		fmt.Printf("%s: marshal error: %v\n", label, err)
		return
	}
	fmt.Printf("--- %s ---\n%s\n", label, b)
}

func main() {
	target := os.Getenv("MITA_RPC_TARGET")
	if target == "" {
		target = "unix://" + mitaclient.DefaultUDSPath
	}
	fmt.Printf("probe: connecting to %s\n", target)

	c, err := mitaclient.Dial(target)
	if err != nil {
		fmt.Printf("dial failed: %v\n", err)
		os.Exit(1)
	}
	defer c.Close()
	ctx := context.Background()

	// PROBE_MODE=restart: just Stop+Start the proxy and report status.
	if os.Getenv("PROBE_MODE") == "restart" {
		must("Stop", c.Stop(ctx))
		time.Sleep(1 * time.Second)
		must("Start", c.Start(ctx))
		time.Sleep(1 * time.Second)
		status, err := c.Status(ctx)
		must("GetStatus", err)
		fmt.Printf("restart done, status=%s\n", status)
		return
	}

	version, err := c.Version(ctx)
	must("GetVersion", err)
	status, err := c.Status(ctx)
	must("GetStatus", err)
	fmt.Printf("mita version=%s status=%s\n\n", version, status)

	// Seed: two users + one port binding, then start the proxy.
	seed := &pb.ServerConfig{
		PortBindings: []*pb.PortBinding{
			{Port: proto.Int32(2012), Protocol: pb.TransportProtocol_TCP.Enum()},
		},
		Users: []*pb.User{
			{Name: proto.String("alice"), Password: proto.String("alice-pass-1")},
			{Name: proto.String("bob"), Password: proto.String("bob-pass-1")},
		},
		LoggingLevel: pb.LoggingLevel_INFO.Enum(),
		Mtu:          proto.Int32(1400),
	}
	_, err = c.SetConfig(ctx, seed)
	must("SetConfig(seed)", err)
	must("Start", c.Start(ctx))
	time.Sleep(500 * time.Millisecond)

	cfg, err := c.GetConfig(ctx)
	must("GetConfig(after seed)", err)
	dump("config after seed (expect alice+bob, hashed passwords)", cfg)

	// check: a user's hashed password should survive a config round-trip
	cfg3, err := c.GetConfig(ctx)
	must("GetConfig(roundtrip)", err)
	aliceHashBefore := findUser(cfg3, "alice").GetHashedPassword()
	_, err = c.SetConfig(ctx, cfg3) // write back verbatim
	must("SetConfig(roundtrip write)", err)
	cfg3b, err := c.GetConfig(ctx)
	must("GetConfig(roundtrip readback)", err)
	aliceHashAfter := findUser(cfg3b, "alice").GetHashedPassword()
	fmt.Printf("alice hashedPassword preserved after round-trip: %v (before=%q after=%q)\n\n",
		aliceHashBefore == aliceHashAfter && aliceHashAfter != "", aliceHashBefore, aliceHashAfter)

	// check: SetConfig should replace the whole config (alice should be gone)
	cfg1, err := c.GetConfig(ctx)
	must("GetConfig(baseline)", err)
	var keep []*pb.User
	for _, u := range cfg1.GetUsers() {
		if u.GetName() != "alice" {
			keep = append(keep, u)
		}
	}
	cfg1.Users = keep
	_, err = c.SetConfig(ctx, cfg1)
	must("SetConfig(drop alice)", err)
	cfg1b, err := c.GetConfig(ctx)
	must("GetConfig(replace readback)", err)
	if findUser(cfg1b, "alice") == nil {
		fmt.Printf("SetConfig is REPLACE (alice removed)\n\n")
	} else {
		fmt.Printf("SetConfig is MERGE (alice still present!) - panel delete flow must change\n\n")
	}

	// check: port binding changes should take effect
	_, err = c.UpdateConfig(ctx, func(cfg *pb.ServerConfig) error {
		cfg.PortBindings = []*pb.PortBinding{
			{Port: proto.Int32(2013), Protocol: pb.TransportProtocol_TCP.Enum()},
		}
		return nil
	})
	must("UpdateConfig(port change)", err)
	time.Sleep(1 * time.Second)
	cfg2, err := c.GetConfig(ctx)
	must("GetConfig(port readback)", err)
	dump("config after port change + Reload", cfg2)
	probeHost := os.Getenv("PROBE_PORT_HOST")
	if probeHost != "" {
		for _, p := range []string{"2012", "2013"} {
			conn, err := net.DialTimeout("tcp", net.JoinHostPort(probeHost, p), 2*time.Second)
			if err != nil {
				fmt.Printf("tcp connect %s: CLOSED (%v)\n", p, err)
			} else {
				conn.Close()
				fmt.Printf("tcp connect %s: OPEN\n", p)
			}
		}
		fmt.Println("expected after Reload if ports apply live: 2012 CLOSED, 2013 OPEN")
	} else {
		fmt.Println("set PROBE_PORT_HOST to test real port connectability")
	}
	fmt.Println()

	// --- Schema dumps for the dashboard ---
	metricsJSON, err := c.Metrics(ctx)
	must("GetMetrics", err)
	fmt.Printf("--- GetMetrics JSON ---\n%s\n\n", metricsJSON)

	users, err := c.Users(ctx)
	must("GetUsers", err)
	dump("GetUsers", users)

	sessions, err := c.Sessions(ctx)
	must("GetSessionInfoList", err)
	dump("GetSessionInfoList", sessions)

	fmt.Println("probe: done")
}

func findUser(cfg *pb.ServerConfig, name string) *pb.User {
	for _, u := range cfg.GetUsers() {
		if u.GetName() == name {
			return u
		}
	}
	return nil
}

func must(op string, err error) {
	if err != nil {
		fmt.Printf("%s failed: %v\n", op, err)
		os.Exit(1)
	}
}
