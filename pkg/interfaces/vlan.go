package interfaces

import (
	"context"
	"github.com/browningluke/opnsense-go/pkg/api"
)

var VlanOpts = api.ReqOpts{
	AddEndpoint:         "/interfaces/vlan_settings/addItem",
	GetEndpoint:         "/interfaces/vlan_settings/getItem",
	UpdateEndpoint:      "/interfaces/vlan_settings/setItem",
	DeleteEndpoint:      "/interfaces/vlan_settings/delItem",
	ReconfigureEndpoint: "/interfaces/vlan_settings/reconfigure",
	Monad:               "vlan",
}

// Data structs

type Vlan struct {
	Description string          `json:"descr"`
	Tag         string          `json:"tag"`
	Priority    api.SelectedMap `json:"pcp"`
	Parent      api.SelectedMap `json:"if"`
	Device      string          `json:"vlanif"`
}

// CRUD operations

func (c *Controller) AddVlan(ctx context.Context, resource *Vlan) (string, error) {
	return api.Add(c.Client(), ctx, VlanOpts, resource)
}

func (c *Controller) GetVlan(ctx context.Context, id string) (*Vlan, error) {
	return api.Get(c.Client(), ctx, VlanOpts, &Vlan{}, id)
}

func (c *Controller) UpdateVlan(ctx context.Context, id string, resource *Vlan) error {
	return api.Update(c.Client(), ctx, VlanOpts, resource, id)
}

func (c *Controller) DeleteVlan(ctx context.Context, id string) error {
	return api.Delete(c.Client(), ctx, VlanOpts, id)
}
