package firewall

import (
	"context"
	"github.com/browningluke/opnsense-go/pkg/api"
)

var NATOpts = api.ReqOpts{
	AddEndpoint:         "/firewall/source_nat/addRule",
	GetEndpoint:         "/firewall/source_nat/getRule",
	UpdateEndpoint:      "/firewall/source_nat/setRule",
	DeleteEndpoint:      "/firewall/source_nat/delRule",
	ReconfigureEndpoint: "/firewall/source_nat/apply",
	Monad:               "rule",
}

// Data structs

type NAT struct {
	Enabled           string          `json:"enabled"`
	DisableNAT        string          `json:"nonat"`
	Sequence          string          `json:"sequence"`
	Interface         api.SelectedMap `json:"interface"`
	IPProtocol        api.SelectedMap `json:"ipprotocol"`
	Protocol          api.SelectedMap `json:"protocol"`
	SourceNet         string          `json:"source_net"`
	SourcePort        string          `json:"source_port"`
	SourceInvert      string          `json:"source_not"`
	DestinationNet    string          `json:"destination_net"`
	DestinationPort   string          `json:"destination_port"`
	DestinationInvert string          `json:"destination_not"`
	Target            string          `json:"target"`
	TargetPort        string          `json:"target_port"`
	Log               string          `json:"log"`
	Description       string          `json:"description"`
}

// CRUD operations

func (c *Controller) AddNAT(ctx context.Context, resource *NAT) (string, error) {
	return api.Add(c.Client(), ctx, NATOpts, resource)
}

func (c *Controller) GetNAT(ctx context.Context, id string) (*NAT, error) {
	return api.Get(c.Client(), ctx, NATOpts, &NAT{}, id)
}

func (c *Controller) UpdateNAT(ctx context.Context, id string, resource *NAT) error {
	return api.Update(c.Client(), ctx, NATOpts, resource, id)
}

func (c *Controller) DeleteNAT(ctx context.Context, id string) error {
	return api.Delete(c.Client(), ctx, NATOpts, id)
}
