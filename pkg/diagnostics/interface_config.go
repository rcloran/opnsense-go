package diagnostics

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/browningluke/opnsense-go/pkg/api"
)

var InterfaceConfigOpts = api.ReqOpts{
	GetEndpoint: "/diagnostics/interface/getInterfaceConfig",
}

// Data structs

type Ipv4Config struct {
	IpAddr     string `json:"ipaddr"`
	SubnetBits int64  `json:"subnetbits"`
	Tunnel     bool   `json:"tunnel"`
}

type Ipv6Config struct {
	IpAddr     string `json:"ipaddr"`
	SubnetBits int64  `json:"subnetbits"`
	Tunnel     bool   `json:"tunnel"`
	Autoconf   bool   `json:"autoconf"`
	Deprecated bool   `json:"deprecated"`
	LinkLocal  bool   `json:"link-local"`
	Tentative  bool   `json:"tentative"`
}

type InterfaceConfig struct {
	Device     string `json:"device"`
	Media      string `json:"media"`
	MediaRaw   string `json:"media_raw"`
	MacAddr    string `json:"macaddr"`
	IsPhysical bool   `json:"is_physical"`
	Mtu        string `json:"mtu"`
	Status     string `json:"status"`

	Flags          []string `json:"flags"`
	Capabilities   []string `json:"capabilities"`
	Options        []string `json:"options"`
	SupportedMedia []string `json:"supported_media"`
	Groups         []string `json:"groups"`

	Ipv4 []Ipv4Config `json:"ipv4"`
	Ipv6 []Ipv6Config `json:"ipv6"`
}

// CRUD operations
func (c *Controller) GetInterfaceConfig(ctx context.Context) (map[string]InterfaceConfig, error) {
	var respData json.RawMessage
	var client = c.Client()

	err := client.DoRequest(ctx, "GET", InterfaceConfigOpts.GetEndpoint, nil, &respData)

	if err != nil {
		switch err.(type) {
		case *json.UnmarshalTypeError:
			return nil, errors.New("Hello") // errs.NewNotFoundError()
		}
		return nil, err
	}

	var res map[string]InterfaceConfig
	if err := json.Unmarshal(respData, &res); err != nil {
		return nil, err
	}

	return res, nil

	// return map[string]InterfaceConfig{}, nil
}
