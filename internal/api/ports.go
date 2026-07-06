// SPDX-License-Identifier: GPL-3.0-or-later

package api

import (
	"context"
	"time"

	pb "github.com/enfein/mieru/v3/pkg/appctl/appctlpb"

	"github.com/fjcrux/mieru-web-ui/internal/chain"
	"github.com/fjcrux/mieru-web-ui/internal/geoip"
)

// The api layer depends on these ports (interfaces), not on concrete
// infrastructure, so the dependency direction points inward and handlers can
// be tested with fakes. The concrete implementations live in their own
// packages (mitaclient, geoip, supervisor) and satisfy these by shape.

// MitaController manages the mita daemon over its gRPC API.
type MitaController interface {
	Status(ctx context.Context) (string, error)
	Version(ctx context.Context) (string, error)
	Start(ctx context.Context) error
	Stop(ctx context.Context) error
	Reload(ctx context.Context) error
	GetConfig(ctx context.Context) (*pb.ServerConfig, error)
	SetConfig(ctx context.Context, cfg *pb.ServerConfig) (*pb.ServerConfig, error)
	UpdateConfig(ctx context.Context, mutate func(*pb.ServerConfig) error) (*pb.ServerConfig, error)
	Metrics(ctx context.Context) (string, error)
	Users(ctx context.Context) (*pb.UserWithMetricsList, error)
	Sessions(ctx context.Context) (*pb.SessionInfoList, error)
}

// GeoResolver manages GeoIP/GeoSite datasets and expands categories to CIDRs
// (geoip) or domain suffixes (geosite).
type GeoResolver interface {
	Datasets() ([]geoip.Dataset, error)
	Categories() ([]geoip.Category, error)
	CIDRs(code string) ([]string, error)
	AddDataset(ctx context.Context, name, url string) error
	DeleteDataset(name string) error

	SiteDatasets() ([]geoip.Dataset, error)
	SiteCategories() ([]geoip.SiteCategory, error)
	Domains(code string) ([]string, error)
	AddSiteDataset(ctx context.Context, name, url string) error
	DeleteSiteDataset(name string) error
}

// ProcSupervisor is the subset of the process supervisor the panel exposes.
type ProcSupervisor interface {
	Logs(n int) []string
	Restart()
	Stats() (restarts int, startedAt time.Time)
}

// PeerManager runs upstream chain peers (supervised mieru clients).
type PeerManager interface {
	List() []chain.Status
	Add(name, key string) error
	Remove(name string) error
	EgressProxies() []*pb.EgressProxy
	Names() map[string]bool
}
