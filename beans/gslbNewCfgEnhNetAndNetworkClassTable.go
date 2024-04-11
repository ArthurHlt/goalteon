package beans

import (
	"fmt"
	"reflect"
)

// GslbNewCfgEnhNetAndNetworkClassTable The network and network class table for gslb in the new_config.
// Note:This mib is not supported for VX instance of Virtualization.
type GslbNewCfgEnhNetAndNetworkClassTable struct {
	// The new gslb network table index.
	GslbNewCfgEnhNetAndNetworkClassNetIndx int32
	// The new network class table index.
	GslbNewCfgEnhNetAndNetworkClassNwclsIndx string
	Params                                   *GslbNewCfgEnhNetAndNetworkClassTableParams
}

func NewGslbNewCfgEnhNetAndNetworkClassTableList() *GslbNewCfgEnhNetAndNetworkClassTable {
	return &GslbNewCfgEnhNetAndNetworkClassTable{}
}

func NewGslbNewCfgEnhNetAndNetworkClassTable(
	gslbNewCfgEnhNetAndNetworkClassNetIndx int32,
	gslbNewCfgEnhNetAndNetworkClassNwclsIndx string,
	params *GslbNewCfgEnhNetAndNetworkClassTableParams,
) *GslbNewCfgEnhNetAndNetworkClassTable {
	return &GslbNewCfgEnhNetAndNetworkClassTable{
		GslbNewCfgEnhNetAndNetworkClassNetIndx:   gslbNewCfgEnhNetAndNetworkClassNetIndx,
		GslbNewCfgEnhNetAndNetworkClassNwclsIndx: gslbNewCfgEnhNetAndNetworkClassNwclsIndx,
		Params:                                   params,
	}
}

func (c *GslbNewCfgEnhNetAndNetworkClassTable) Name() string {
	return "GslbNewCfgEnhNetAndNetworkClassTable"
}

func (c *GslbNewCfgEnhNetAndNetworkClassTable) GetParams() BeanType {
	return c.Params
}

func (c *GslbNewCfgEnhNetAndNetworkClassTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *GslbNewCfgEnhNetAndNetworkClassTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.GslbNewCfgEnhNetAndNetworkClassNetIndx).IsZero() &&
		reflect.ValueOf(c.GslbNewCfgEnhNetAndNetworkClassNwclsIndx).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.GslbNewCfgEnhNetAndNetworkClassNetIndx) + "/" + fmt.Sprint(c.GslbNewCfgEnhNetAndNetworkClassNwclsIndx)
}

type GslbNewCfgEnhNetAndNetworkClassTableDelete int32

const (
	GslbNewCfgEnhNetAndNetworkClassTableDelete_Other       GslbNewCfgEnhNetAndNetworkClassTableDelete = 1
	GslbNewCfgEnhNetAndNetworkClassTableDelete_Delete      GslbNewCfgEnhNetAndNetworkClassTableDelete = 2
	GslbNewCfgEnhNetAndNetworkClassTableDelete_Unsupported GslbNewCfgEnhNetAndNetworkClassTableDelete = 2147483647
)

type GslbNewCfgEnhNetAndNetworkClassTableParams struct {
	// The new gslb network table index.
	NetIndx int32 `json:"NetIndx,omitempty"`
	// The new network class table index.
	NwclsIndx string `json:"NwclsIndx,omitempty"`
	// By setting the value to delete(2), the network class will be removed from gslb network table.
	Delete GslbNewCfgEnhNetAndNetworkClassTableDelete `json:"Delete,omitempty"`
}

func (p GslbNewCfgEnhNetAndNetworkClassTableParams) iMABean() {}
