package firewall

import (
	"context"
	"github.com/browningluke/opnsense-go/pkg/api"
)

var CategoryOpts = api.ReqOpts{
	AddEndpoint:         "/firewall/category/addItem",
	GetEndpoint:         "/firewall/category/getItem",
	UpdateEndpoint:      "/firewall/category/setItem",
	DeleteEndpoint:      "/firewall/category/delItem",
	ReconfigureEndpoint: "",
	Monad:               "category",
}

// Data structs

type Category struct {
	Automatic string `json:"auto"`
	Name      string `json:"name"`
	Color     string `json:"color"`
}

// CRUD operations

func (c *Controller) AddCategory(ctx context.Context, resource *Category) (string, error) {
	return api.Add(c.Client(), ctx, CategoryOpts, resource)
}

func (c *Controller) GetCategory(ctx context.Context, id string) (*Category, error) {
	return api.Get(c.Client(), ctx, CategoryOpts, &Category{}, id)
}

func (c *Controller) UpdateCategory(ctx context.Context, id string, resource *Category) error {
	return api.Update(c.Client(), ctx, CategoryOpts, resource, id)
}

func (c *Controller) DeleteCategory(ctx context.Context, id string) error {
	return api.Delete(c.Client(), ctx, CategoryOpts, id)
}
