package beans

import (
	"fmt"
	"reflect"
)

// Layer7NewCfgURLFilteringTable The table for configuring URL Filtering.
// Note:This mib is not supported for VX instance of virtualization.
type Layer7NewCfgURLFilteringTable struct {
	// The URL Filtering policy ID(key id).
	Layer7NewCfgURLFilteringID string
	Params                     *Layer7NewCfgURLFilteringTableParams
}

func NewLayer7NewCfgURLFilteringTableList() *Layer7NewCfgURLFilteringTable {
	return &Layer7NewCfgURLFilteringTable{}
}

func NewLayer7NewCfgURLFilteringTable(
	layer7NewCfgURLFilteringID string,
	params *Layer7NewCfgURLFilteringTableParams,
) *Layer7NewCfgURLFilteringTable {
	return &Layer7NewCfgURLFilteringTable{
		Layer7NewCfgURLFilteringID: layer7NewCfgURLFilteringID,
		Params:                     params,
	}
}

func (c *Layer7NewCfgURLFilteringTable) Name() string {
	return "Layer7NewCfgURLFilteringTable"
}

func (c *Layer7NewCfgURLFilteringTable) GetParams() BeanType {
	return c.Params
}

func (c *Layer7NewCfgURLFilteringTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *Layer7NewCfgURLFilteringTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.Layer7NewCfgURLFilteringID).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.Layer7NewCfgURLFilteringID)
}

type Layer7NewCfgURLFilteringTableDelete int32

const (
	Layer7NewCfgURLFilteringTableDelete_Other       Layer7NewCfgURLFilteringTableDelete = 1
	Layer7NewCfgURLFilteringTableDelete_Delete      Layer7NewCfgURLFilteringTableDelete = 2
	Layer7NewCfgURLFilteringTableDelete_Unsupported Layer7NewCfgURLFilteringTableDelete = 2147483647
)

type Layer7NewCfgURLFilteringTableParams struct {
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
	// By setting the value to delete(2), the entire row is deleted.
	Delete Layer7NewCfgURLFilteringTableDelete `json:"Delete,omitempty"`
	// The URL Filtering policy fallback categories srting.
	FallbackCatgs string `json:"FallbackCatgs,omitempty"`
}

func (p Layer7NewCfgURLFilteringTableParams) iMABean() {}
