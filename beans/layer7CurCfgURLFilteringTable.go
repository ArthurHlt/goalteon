package beans

import (
	"fmt"
	"reflect"
)

// Layer7CurCfgURLFilteringTable The table for configuring URL Filtering.
// Note:This mib is not supported for VX instance of virtualization.
type Layer7CurCfgURLFilteringTable struct {
	// The URL Filtering policy ID(key id).
	Layer7CurCfgURLFilteringID string
	Params                     *Layer7CurCfgURLFilteringTableParams
}

func NewLayer7CurCfgURLFilteringTableList() *Layer7CurCfgURLFilteringTable {
	return &Layer7CurCfgURLFilteringTable{}
}

func NewLayer7CurCfgURLFilteringTable(
	layer7CurCfgURLFilteringID string,
	params *Layer7CurCfgURLFilteringTableParams,
) *Layer7CurCfgURLFilteringTable {
	return &Layer7CurCfgURLFilteringTable{
		Layer7CurCfgURLFilteringID: layer7CurCfgURLFilteringID,
		Params:                     params,
	}
}

func (c *Layer7CurCfgURLFilteringTable) Name() string {
	return "Layer7CurCfgURLFilteringTable"
}

func (c *Layer7CurCfgURLFilteringTable) GetParams() BeanType {
	return c.Params
}

func (c *Layer7CurCfgURLFilteringTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *Layer7CurCfgURLFilteringTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.Layer7CurCfgURLFilteringID).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.Layer7CurCfgURLFilteringID)
}

type Layer7CurCfgURLFilteringTableParams struct {
	// The URL Filtering policy ID(key id).
	ID string `json:"ID,omitempty"`
	// The URL Filtering policy name.
	Name string `json:"Name,omitempty"`
	// The URL Filtering policy security categories srting.
	SecCatgs string `json:"SecCatgs,omitempty"`
	// The URL Filtering policy compliance categories srting.
	CompCatgs string `json:"CompCatgs,omitempty"`
	// The URL Filtering policy productivity categories srting.
	ProdCatgs string `json:"ProdCatgs,omitempty"`
	// The URL Filtering policy fallback categories srting.
	FallbackCatgs string `json:"FallbackCatgs,omitempty"`
}

func (p Layer7CurCfgURLFilteringTableParams) iMABean() {}
