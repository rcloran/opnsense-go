package wireguard

import (
	"context"
	"github.com/browningluke/opnsense-go/pkg/api"
)

var ServerOpts = api.ReqOpts{
	AddEndpoint:         "/wireguard/server/addServer",
	GetEndpoint:         "/wireguard/server/getServer",
	UpdateEndpoint:      "/wireguard/server/setServer",
	DeleteEndpoint:      "/wireguard/server/delServer",
	ReconfigureEndpoint: wireguardReconfigureEndpoint,
	Monad:               "server",
}

// Data structs

type Server struct {
	Enabled       string              `json:"enabled"`
	Name          string              `json:"name"`
	Instance      string              `json:"instance,omitempty"`
	PublicKey     string              `json:"pubkey"`
	PrivateKey    string              `json:"privkey"`
	Port          string              `json:"port"`
	MTU           string              `json:"mtu"`
	DNS           api.SelectedMapList `json:"dns"`
	TunnelAddress api.SelectedMapList `json:"tunneladdress"`
	DisableRoutes string              `json:"disableroutes"`
	Gateway       string              `json:"gateway"`
	Peers         api.SelectedMapList `json:"peers"`
}

// CRUD operations

func (c *Controller) AddServer(ctx context.Context, resource *Server) (string, error) {
	return api.Add(c.Client(), ctx, ServerOpts, resource)
}

func (c *Controller) GetServer(ctx context.Context, id string) (*Server, error) {
	return api.Get(c.Client(), ctx, ServerOpts, &Server{}, id)
}

func (c *Controller) UpdateServer(ctx context.Context, id string, resource *Server) error {
	return api.Update(c.Client(), ctx, ServerOpts, resource, id)
}

func (c *Controller) DeleteServer(ctx context.Context, id string) error {
	return api.Delete(c.Client(), ctx, ServerOpts, id)
}
