package quagga

import (
	"context"
	"github.com/browningluke/opnsense-go/pkg/api"
)

var BGPCommunityListOpts = api.ReqOpts{
	AddEndpoint:         "/quagga/bgp/addCommunitylist",
	GetEndpoint:         "/quagga/bgp/getCommunitylist",
	UpdateEndpoint:      "/quagga/bgp/setCommunitylist",
	DeleteEndpoint:      "/quagga/bgp/delCommunitylist",
	ReconfigureEndpoint: quaggaReconfigureEndpoint,
	Monad:               "communitylist",
}

// Data structs

type BGPCommunityList struct {
	Enabled        string          `json:"enabled"`
	Description    string          `json:"description"`
	Number         string          `json:"number"`
	SequenceNumber string          `json:"seqnumber"`
	Action         api.SelectedMap `json:"action"`
	Community      string          `json:"community"`
}

// CRUD operations

func (c *Controller) AddBGPCommunityList(ctx context.Context, resource *BGPCommunityList) (string, error) {
	return api.Add(c.Client(), ctx, BGPCommunityListOpts, resource)
}

func (c *Controller) GetBGPCommunityList(ctx context.Context, id string) (*BGPCommunityList, error) {
	return api.Get(c.Client(), ctx, BGPCommunityListOpts, &BGPCommunityList{}, id)
}

func (c *Controller) UpdateBGPCommunityList(ctx context.Context, id string, resource *BGPCommunityList) error {
	return api.Update(c.Client(), ctx, BGPCommunityListOpts, resource, id)
}

func (c *Controller) DeleteBGPCommunityList(ctx context.Context, id string) error {
	return api.Delete(c.Client(), ctx, BGPCommunityListOpts, id)
}
