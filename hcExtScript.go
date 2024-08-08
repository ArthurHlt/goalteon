package goalteon

import (
	"fmt"
	"github.com/ArthurHlt/goalteon/beans"
	"net/url"
)

type ImpExpHcExtScriptParams struct {
	ID      string
	Name    string
	Disable bool
}

func (i *ImpExpHcExtScriptParams) Values() url.Values {
	v := url.Values{}
	v.Set("id", i.ID)
	if i.Name == "" {
		i.Name = i.ID
	}
	v.Set("name", i.Name)
	return v
}

func (c *Client) ExportHcExtScript(params *ImpExpHcExtScriptParams) ([]byte, error) {
	if params == nil {
		return nil, fmt.Errorf("params are nil")
	}
	if params.ID == "" {
		return nil, fmt.Errorf("id is required")
	}
	res, err := c.Get(&beans.GetHcExtScriptBean{}, params.Values())
	if err != nil {
		return nil, err
	}
	data := Translate[beans.BytesParams](res)
	return data, nil
}

func (c *Client) ImportHcExtScript(data []byte, params *ImpExpHcExtScriptParams) (*StatusResponse, error) {
	if params == nil {
		return nil, fmt.Errorf("params are nil")
	}
	if params.ID == "" {
		return nil, fmt.Errorf("id is required")
	}
	return c.Post(&beans.HcExtScriptImportBean{Content: data}, params.Values())
}

func (c *Client) ImportHcExtScriptAsBulk(data []byte, params *ImpExpHcExtScriptParams) (*BulkItem, error) {
	if params == nil {
		return nil, fmt.Errorf("params are nil")
	}
	if params.ID == "" {
		return nil, fmt.Errorf("id is required")
	}
	return NewBulkItem(BulkMethodPost, &beans.HcExtScriptImportBean{Content: data}, params.Values()), nil
}
