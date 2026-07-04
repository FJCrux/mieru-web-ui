// SPDX-License-Identifier: GPL-3.0-or-later

package mitaclient

import (
	"context"

	pb "github.com/enfein/mieru/v3/pkg/appctl/appctlpb"
	"google.golang.org/protobuf/types/known/emptypb"
)

// Metrics returns the daemon metrics dump as a raw JSON string.
func (c *Client) Metrics(ctx context.Context) (string, error) {
	ctx, cancel := c.withTimeout(ctx)
	defer cancel()
	m, err := c.rpc.GetMetrics(ctx, &emptypb.Empty{})
	if err != nil {
		return "", err
	}
	return m.GetJson(), nil
}

// Users returns per-user config plus runtime traffic metrics.
func (c *Client) Users(ctx context.Context) (*pb.UserWithMetricsList, error) {
	ctx, cancel := c.withTimeout(ctx)
	defer cancel()
	return c.rpc.GetUsers(ctx, &emptypb.Empty{})
}

// Sessions returns the list of active proxy sessions.
func (c *Client) Sessions(ctx context.Context) (*pb.SessionInfoList, error) {
	ctx, cancel := c.withTimeout(ctx)
	defer cancel()
	return c.rpc.GetSessionInfoList(ctx, &emptypb.Empty{})
}
