package goalteaon

import (
	"fmt"
	"github.com/ArthurHlt/goalteon/beans"
	"net/url"
)

type ImpExpAppShapeParams struct {
	ID      string
	Name    string
	Disable bool
}

func (i *ImpExpAppShapeParams) Values() url.Values {
	v := url.Values{}
	v.Set("id", i.ID)
	if i.Name == "" {
		i.Name = i.ID
	}
	v.Set("name", i.Name)
	state := "ena"
	if i.Disable {
		state = "dis"
	}
	v.Set("state", state)
	return v
}

func (c *Client) ExportAppShape(params *ImpExpAppShapeParams) ([]byte, error) {
	if params == nil {
		return nil, fmt.Errorf("params are nil")
	}
	if params.ID == "" {
		return nil, fmt.Errorf("id is required")
	}
	res, err := c.Get(&beans.GetAppShapeBean{}, params.Values())
	if err != nil {
		return nil, err
	}
	data := Translate[beans.BytesParams](res)
	return data, nil
}

func (c *Client) ImportAppShape(data []byte, params *ImpExpAppShapeParams) (*StatusResponse, error) {
	if params == nil {
		return nil, fmt.Errorf("params are nil")
	}
	if params.ID == "" {
		return nil, fmt.Errorf("id is required")
	}
	return c.Post(&beans.AppShapeImportBean{Content: data}, params.Values())
}
