package quagga

import (
	"context"
	"github.com/browningluke/opnsense-go/pkg/api"
)

var BGPASPathOpts = api.ReqOpts{
	AddEndpoint:         "/quagga/bgp/addAspath",
	GetEndpoint:         "/quagga/bgp/getAspath",
	UpdateEndpoint:      "/quagga/bgp/setAspath",
	DeleteEndpoint:      "/quagga/bgp/delAspath",
	ReconfigureEndpoint: quaggaReconfigureEndpoint,
	Monad:               "aspath",
}

// Data structs

type BGPASPath struct {
	Enabled     string          `json:"enabled"`
	Description string          `json:"description"`
	Number      string          `json:"number"`
	Action      api.SelectedMap `json:"action"`
	AS          string          `json:"as"`
}

// CRUD operations

func (c *Controller) AddBGPASPath(ctx context.Context, resource *BGPASPath) (string, error) {
	return api.Add(c.Client(), ctx, BGPASPathOpts, resource)
}

func (c *Controller) GetBGPASPath(ctx context.Context, id string) (*BGPASPath, error) {
	return api.Get(c.Client(), ctx, BGPASPathOpts, &BGPASPath{}, id)
}

func (c *Controller) UpdateBGPASPath(ctx context.Context, id string, resource *BGPASPath) error {
	return api.Update(c.Client(), ctx, BGPASPathOpts, resource, id)
}

func (c *Controller) DeleteBGPASPath(ctx context.Context, id string) error {
	return api.Delete(c.Client(), ctx, BGPASPathOpts, id)
}
