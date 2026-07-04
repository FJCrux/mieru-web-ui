// SPDX-License-Identifier: GPL-3.0-or-later

// Package mitaclient wraps the mita ServerManagementService gRPC API.
// All appctl types are kept inside this package so a mieru version bump
// only touches code here.
package mitaclient

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/enfein/mieru/v3/pkg/appctl/appctlgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"
)

const (
	// DefaultUDSPath is where mita serves its management API.
	DefaultUDSPath = "/var/run/mita/mita.sock"

	// maxRecvMsgSize mirrors appctl.MaxRecvMsgSize (16 MiB).
	maxRecvMsgSize = 16 * 1024 * 1024

	rpcTimeout = 10 * time.Second
)

// Client is a thin typed wrapper over the mita management gRPC service.
type Client struct {
	conn *grpc.ClientConn
	rpc  appctlgrpc.ServerManagementServiceClient
}

// Dial connects to the mita management API.
// target accepts "unix:///path/to/mita.sock" or "tcp://host:port"
// (the latter is used in development through a socat bridge).
// A bare filesystem path is treated as a unix socket path.
func Dial(target string) (*Client, error) {
	var addr string
	switch {
	case strings.HasPrefix(target, "unix://"):
		addr = target
	case strings.HasPrefix(target, "tcp://"):
		addr = strings.TrimPrefix(target, "tcp://")
	case strings.HasPrefix(target, "/"):
		addr = "unix://" + target
	default:
		return nil, fmt.Errorf("unsupported mita RPC target %q", target)
	}
	conn, err := grpc.NewClient(addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(maxRecvMsgSize)))
	if err != nil {
		return nil, fmt.Errorf("grpc.NewClient(%s) failed: %w", addr, err)
	}
	return &Client{conn: conn, rpc: appctlgrpc.NewServerManagementServiceClient(conn)}, nil
}

func (c *Client) Close() error {
	return c.conn.Close()
}

func (c *Client) withTimeout(ctx context.Context) (context.Context, context.CancelFunc) {
	return context.WithTimeout(ctx, rpcTimeout)
}

// Status returns the daemon proxy state (e.g. "RUNNING", "IDLE").
func (c *Client) Status(ctx context.Context) (string, error) {
	ctx, cancel := c.withTimeout(ctx)
	defer cancel()
	st, err := c.rpc.GetStatus(ctx, &emptypb.Empty{})
	if err != nil {
		return "", err
	}
	return st.GetStatus().String(), nil
}

// Version returns the mita daemon version as "major.minor.patch".
func (c *Client) Version(ctx context.Context) (string, error) {
	ctx, cancel := c.withTimeout(ctx)
	defer cancel()
	v, err := c.rpc.GetVersion(ctx, &emptypb.Empty{})
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%d.%d.%d", v.GetMajor(), v.GetMinor(), v.GetPatch()), nil
}

// Start begins proxying.
func (c *Client) Start(ctx context.Context) error {
	ctx, cancel := c.withTimeout(ctx)
	defer cancel()
	_, err := c.rpc.Start(ctx, &emptypb.Empty{})
	return err
}

// Stop halts proxying (daemon keeps running).
func (c *Client) Stop(ctx context.Context) error {
	ctx, cancel := c.withTimeout(ctx)
	defer cancel()
	_, err := c.rpc.Stop(ctx, &emptypb.Empty{})
	return err
}

// Reload re-applies users and logging config without dropping connections.
func (c *Client) Reload(ctx context.Context) error {
	ctx, cancel := c.withTimeout(ctx)
	defer cancel()
	_, err := c.rpc.Reload(ctx, &emptypb.Empty{})
	return err
}
