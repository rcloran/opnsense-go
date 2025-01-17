package unbound

import (
	"context"
	"github.com/browningluke/opnsense-go/pkg/api"
)

var HostOverrideOpts = api.ReqOpts{
	AddEndpoint:         "/unbound/settings/addHostOverride",
	GetEndpoint:         "/unbound/settings/getHostOverride",
	UpdateEndpoint:      "/unbound/settings/setHostOverride",
	DeleteEndpoint:      "/unbound/settings/delHostOverride",
	ReconfigureEndpoint: unboundReconfigureEndpoint,
	Monad:               "host",
}

// Data structs

type HostOverride struct {
	Enabled     string          `json:"enabled"`
	Hostname    string          `json:"hostname"`
	Domain      string          `json:"domain"`
	Type        api.SelectedMap `json:"rr"`
	Server      string          `json:"server"`
	MXPriority  string          `json:"mxprio"`
	MXDomain    string          `json:"mx"`
	Description string          `json:"description"`
}

// CRUD operations

func (c *Controller) AddHostOverride(ctx context.Context, resource *HostOverride) (string, error) {
	return api.Add(c.Client(), ctx, HostOverrideOpts, resource)
}

func (c *Controller) GetHostOverride(ctx context.Context, id string) (*HostOverride, error) {
	return api.Get(c.Client(), ctx, HostOverrideOpts, &HostOverride{}, id)
}

func (c *Controller) UpdateHostOverride(ctx context.Context, id string, resource *HostOverride) error {
	return api.Update(c.Client(), ctx, HostOverrideOpts, resource, id)
}

func (c *Controller) DeleteHostOverride(ctx context.Context, id string) error {
	return api.Delete(c.Client(), ctx, HostOverrideOpts, id)
}
