package beans

import (
	"fmt"
	"reflect"
)

// GslbNewCfgEnhNetworkRealPrefTable The table of network remote real preference in the new_config.
// Note:This mib is not supported for VX instance of Virtualization.
type GslbNewCfgEnhNetworkRealPrefTable struct {
	// The new network preference table index.
	GslbNewCfgEnhNetworkRealIndx int32
	// The new network preference server table index.
	GslbNewCfgEnhNetworkRealServerIndx int32
	Params                             *GslbNewCfgEnhNetworkRealPrefTableParams
}

func NewGslbNewCfgEnhNetworkRealPrefTableList() *GslbNewCfgEnhNetworkRealPrefTable {
	return &GslbNewCfgEnhNetworkRealPrefTable{}
}

func NewGslbNewCfgEnhNetworkRealPrefTable(
	gslbNewCfgEnhNetworkRealIndx int32,
	gslbNewCfgEnhNetworkRealServerIndx int32,
	params *GslbNewCfgEnhNetworkRealPrefTableParams,
) *GslbNewCfgEnhNetworkRealPrefTable {
	return &GslbNewCfgEnhNetworkRealPrefTable{
		GslbNewCfgEnhNetworkRealIndx:       gslbNewCfgEnhNetworkRealIndx,
		GslbNewCfgEnhNetworkRealServerIndx: gslbNewCfgEnhNetworkRealServerIndx,
		Params:                             params,
	}
}

func (c *GslbNewCfgEnhNetworkRealPrefTable) Name() string {
	return "GslbNewCfgEnhNetworkRealPrefTable"
}

func (c *GslbNewCfgEnhNetworkRealPrefTable) GetParams() BeanType {
	return c.Params
}

func (c *GslbNewCfgEnhNetworkRealPrefTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *GslbNewCfgEnhNetworkRealPrefTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.GslbNewCfgEnhNetworkRealIndx).IsZero() &&
		reflect.ValueOf(c.GslbNewCfgEnhNetworkRealServerIndx).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.GslbNewCfgEnhNetworkRealIndx) + "/" + fmt.Sprint(c.GslbNewCfgEnhNetworkRealServerIndx)
}

type GslbNewCfgEnhNetworkRealPrefTableParams struct {
	// The new network preference table index.
	Indx int32 `json:"Indx,omitempty"`
	// The new network preference server table index.
	ServerIndx int32 `json:"ServerIndx,omitempty"`
	// Add remote real server preference to network.
	// Set Zero to remove the remote real server from network
	ServerPref int32 `json:"ServerPref,omitempty"`
}

func (p GslbNewCfgEnhNetworkRealPrefTableParams) iMABean() {}
