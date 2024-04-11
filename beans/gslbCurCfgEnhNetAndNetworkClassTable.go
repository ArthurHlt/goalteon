package beans

import (
	"fmt"
	"reflect"
)

// GslbCurCfgEnhNetAndNetworkClassTable The network and network class table in gslb.
// Note:This mib is not supported for VX instance of Virtualization.
type GslbCurCfgEnhNetAndNetworkClassTable struct {
	// The current network table index.
	GslbCurCfgEnhNetAndNetworkClassNetIndx int32
	// The current network class index.
	GslbCurCfgEnhNetAndNetworkClassNwclsIndx string
	Params                                   *GslbCurCfgEnhNetAndNetworkClassTableParams
}

func NewGslbCurCfgEnhNetAndNetworkClassTableList() *GslbCurCfgEnhNetAndNetworkClassTable {
	return &GslbCurCfgEnhNetAndNetworkClassTable{}
}

func NewGslbCurCfgEnhNetAndNetworkClassTable(
	gslbCurCfgEnhNetAndNetworkClassNetIndx int32,
	gslbCurCfgEnhNetAndNetworkClassNwclsIndx string,
	params *GslbCurCfgEnhNetAndNetworkClassTableParams,
) *GslbCurCfgEnhNetAndNetworkClassTable {
	return &GslbCurCfgEnhNetAndNetworkClassTable{
		GslbCurCfgEnhNetAndNetworkClassNetIndx:   gslbCurCfgEnhNetAndNetworkClassNetIndx,
		GslbCurCfgEnhNetAndNetworkClassNwclsIndx: gslbCurCfgEnhNetAndNetworkClassNwclsIndx,
		Params:                                   params,
	}
}

func (c *GslbCurCfgEnhNetAndNetworkClassTable) Name() string {
	return "GslbCurCfgEnhNetAndNetworkClassTable"
}

func (c *GslbCurCfgEnhNetAndNetworkClassTable) GetParams() BeanType {
	return c.Params
}

func (c *GslbCurCfgEnhNetAndNetworkClassTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *GslbCurCfgEnhNetAndNetworkClassTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.GslbCurCfgEnhNetAndNetworkClassNetIndx).IsZero() &&
		reflect.ValueOf(c.GslbCurCfgEnhNetAndNetworkClassNwclsIndx).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.GslbCurCfgEnhNetAndNetworkClassNetIndx) + "/" + fmt.Sprint(c.GslbCurCfgEnhNetAndNetworkClassNwclsIndx)
}

type GslbCurCfgEnhNetAndNetworkClassTableParams struct {
	// The current network table index.
	NetIndx int32 `json:"NetIndx,omitempty"`
	// The current network class index.
	NwclsIndx string `json:"NwclsIndx,omitempty"`
}

func (p GslbCurCfgEnhNetAndNetworkClassTableParams) iMABean() {}
