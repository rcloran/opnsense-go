package unbound

import (
	"context"
	"github.com/browningluke/opnsense-go/pkg/api"
)

var ForwardOpts = api.ReqOpts{
	AddEndpoint:         "/unbound/settings/addDot",
	GetEndpoint:         "/unbound/settings/getDot",
	UpdateEndpoint:      "/unbound/settings/setDot",
	DeleteEndpoint:      "/unbound/settings/delDot",
	ReconfigureEndpoint: unboundReconfigureEndpoint,
	Monad:               "dot",
}

// Data structs

type Forward struct {
	Enabled  string          `json:"enabled"`
	Domain   string          `json:"domain"`
	Type     api.SelectedMap `json:"type"`
	Server   string          `json:"server"`
	Port     string          `json:"port"`
	VerifyCN string          `json:"verify"`
}

// CRUD operations

func (c *Controller) AddForward(ctx context.Context, resource *Forward) (string, error) {
	return api.Add(c.Client(), ctx, ForwardOpts, resource)
}

func (c *Controller) GetForward(ctx context.Context, id string) (*Forward, error) {
	return api.Get(c.Client(), ctx, ForwardOpts, &Forward{}, id)
}

func (c *Controller) UpdateForward(ctx context.Context, id string, resource *Forward) error {
	return api.Update(c.Client(), ctx, ForwardOpts, resource, id)
}

func (c *Controller) DeleteForward(ctx context.Context, id string) error {
	return api.Delete(c.Client(), ctx, ForwardOpts, id)
}
