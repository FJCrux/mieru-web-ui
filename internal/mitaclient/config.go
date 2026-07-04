// SPDX-License-Identifier: GPL-3.0-or-later

package mitaclient

import (
	"context"
	"sync"

	pb "github.com/enfein/mieru/v3/pkg/appctl/appctlpb"
	"google.golang.org/protobuf/types/known/emptypb"
)

// configMu serializes all read-modify-write cycles against the daemon.
// SetConfig replaces the whole stored config, so concurrent updates
// would silently drop each other's changes.
var configMu sync.Mutex

// GetConfig fetches the full server config. User passwords come back
// hashed only; the plaintext password field is always empty.
func (c *Client) GetConfig(ctx context.Context) (*pb.ServerConfig, error) {
	ctx, cancel := c.withTimeout(ctx)
	defer cancel()
	return c.rpc.GetConfig(ctx, &emptypb.Empty{})
}

// SetConfig replaces the stored server config with cfg.
func (c *Client) SetConfig(ctx context.Context, cfg *pb.ServerConfig) (*pb.ServerConfig, error) {
	ctx, cancel := c.withTimeout(ctx)
	defer cancel()
	return c.rpc.SetConfig(ctx, cfg)
}

// UpdateConfig runs the universal write flow:
// lock → GetConfig → mutate → SetConfig → Reload.
// The mutate callback edits the config in place; returning an error aborts
// the update without writing anything.
func (c *Client) UpdateConfig(ctx context.Context, mutate func(*pb.ServerConfig) error) (*pb.ServerConfig, error) {
	configMu.Lock()
	defer configMu.Unlock()

	cfg, err := c.GetConfig(ctx)
	if err != nil {
		return nil, err
	}
	if err := mutate(cfg); err != nil {
		return nil, err
	}
	updated, err := c.SetConfig(ctx, cfg)
	if err != nil {
		return nil, err
	}
	if err := c.Reload(ctx); err != nil {
		return updated, err
	}
	return updated, nil
}
