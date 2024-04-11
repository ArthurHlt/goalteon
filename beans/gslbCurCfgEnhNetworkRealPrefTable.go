package beans

import (
	"fmt"
	"reflect"
)

// GslbCurCfgEnhNetworkRealPrefTable The table of network remote real  preference in the cur_config.
// Note:This mib is not supported for VX instance of Virtualization.
type GslbCurCfgEnhNetworkRealPrefTable struct {
	// The new network remote real preference  table index.
	GslbCurCfgEnhNetworkRealIndx int32
	// The new network remote real preference server table index.
	GslbCurCfgEnhNetworkRealServerIndx int32
	Params                             *GslbCurCfgEnhNetworkRealPrefTableParams
}

func NewGslbCurCfgEnhNetworkRealPrefTableList() *GslbCurCfgEnhNetworkRealPrefTable {
	return &GslbCurCfgEnhNetworkRealPrefTable{}
}

func NewGslbCurCfgEnhNetworkRealPrefTable(
	gslbCurCfgEnhNetworkRealIndx int32,
	gslbCurCfgEnhNetworkRealServerIndx int32,
	params *GslbCurCfgEnhNetworkRealPrefTableParams,
) *GslbCurCfgEnhNetworkRealPrefTable {
	return &GslbCurCfgEnhNetworkRealPrefTable{
		GslbCurCfgEnhNetworkRealIndx:       gslbCurCfgEnhNetworkRealIndx,
		GslbCurCfgEnhNetworkRealServerIndx: gslbCurCfgEnhNetworkRealServerIndx,
		Params:                             params,
	}
}

func (c *GslbCurCfgEnhNetworkRealPrefTable) Name() string {
	return "GslbCurCfgEnhNetworkRealPrefTable"
}

func (c *GslbCurCfgEnhNetworkRealPrefTable) GetParams() BeanType {
	return c.Params
}

func (c *GslbCurCfgEnhNetworkRealPrefTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *GslbCurCfgEnhNetworkRealPrefTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.GslbCurCfgEnhNetworkRealIndx).IsZero() &&
		reflect.ValueOf(c.GslbCurCfgEnhNetworkRealServerIndx).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.GslbCurCfgEnhNetworkRealIndx) + "/" + fmt.Sprint(c.GslbCurCfgEnhNetworkRealServerIndx)
}

type GslbCurCfgEnhNetworkRealPrefTableParams struct {
	// The new network remote real preference  table index.
	Indx int32 `json:"Indx,omitempty"`
	// The new network remote real preference server table index.
	ServerIndx int32 `json:"ServerIndx,omitempty"`
	// Add remote real server preference to network
	ServerPref int32 `json:"ServerPref,omitempty"`
}

func (p GslbCurCfgEnhNetworkRealPrefTableParams) iMABean() {}
