package quagga

import (
	"context"
	"github.com/browningluke/opnsense-go/pkg/api"
)

var BGPPrefixListOpts = api.ReqOpts{
	AddEndpoint:         "/quagga/bgp/addPrefixlist",
	GetEndpoint:         "/quagga/bgp/getPrefixlist",
	UpdateEndpoint:      "/quagga/bgp/setPrefixlist",
	DeleteEndpoint:      "/quagga/bgp/delPrefixlist",
	ReconfigureEndpoint: quaggaReconfigureEndpoint,
	Monad:               "prefixlist",
}

// Data structs

type BGPPrefixList struct {
	Enabled        string          `json:"enabled"`
	Description    string          `json:"description"`
	Name           string          `json:"name"`
	IPVersion      api.SelectedMap `json:"version"`
	SequenceNumber string          `json:"seqnumber"`
	Action         api.SelectedMap `json:"action"`
	Network        string          `json:"network"`
}

// CRUD operations

func (c *Controller) AddBGPPrefixList(ctx context.Context, resource *BGPPrefixList) (string, error) {
	return api.Add(c.Client(), ctx, BGPPrefixListOpts, resource)
}

func (c *Controller) GetBGPPrefixList(ctx context.Context, id string) (*BGPPrefixList, error) {
	return api.Get(c.Client(), ctx, BGPPrefixListOpts, &BGPPrefixList{}, id)
}

func (c *Controller) UpdateBGPPrefixList(ctx context.Context, id string, resource *BGPPrefixList) error {
	return api.Update(c.Client(), ctx, BGPPrefixListOpts, resource, id)
}

func (c *Controller) DeleteBGPPrefixList(ctx context.Context, id string) error {
	return api.Delete(c.Client(), ctx, BGPPrefixListOpts, id)
}
