package firewall

import (
	"context"
	"github.com/browningluke/opnsense-go/pkg/api"
)

var AliasOpts = api.ReqOpts{
	AddEndpoint:         "/firewall/alias/addItem",
	GetEndpoint:         "/firewall/alias/getItem",
	UpdateEndpoint:      "/firewall/alias/setItem",
	DeleteEndpoint:      "/firewall/alias/delItem",
	ReconfigureEndpoint: "/firewall/alias/reconfigure",
	Monad:               "alias",
}

// Data structs

type Alias struct {
	Enabled     string                `json:"enabled"`
	Name        string                `json:"name"`
	Type        api.SelectedMap       `json:"type"`
	IPProtocol  api.SelectedMap       `json:"proto"`
	Interface   api.SelectedMap       `json:"interface"`
	Content     api.SelectedMapListNL `json:"content"`
	Categories  api.SelectedMapList   `json:"categories"`
	UpdateFreq  string                `json:"updatefreq"`
	Statistics  string                `json:"counters"`
	Description string                `json:"description"`
}

// CRUD operations

func (c *Controller) AddAlias(ctx context.Context, resource *Alias) (string, error) {
	return api.Add(c.Client(), ctx, AliasOpts, resource)
}

func (c *Controller) GetAlias(ctx context.Context, id string) (*Alias, error) {
	return api.Get(c.Client(), ctx, AliasOpts, &Alias{}, id)
}

func (c *Controller) UpdateAlias(ctx context.Context, id string, resource *Alias) error {
	return api.Update(c.Client(), ctx, AliasOpts, resource, id)
}

func (c *Controller) DeleteAlias(ctx context.Context, id string) error {
	return api.Delete(c.Client(), ctx, AliasOpts, id)
}
