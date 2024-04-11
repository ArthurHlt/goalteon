package beans

import (
	"fmt"
	"reflect"
)

// GslbNewCfgEnhNetworkVirtPrefTable The table of network virt preference in the new_config.
// Note:This mib is not supported for VX instance of Virtualization.
type GslbNewCfgEnhNetworkVirtPrefTable struct {
	// The new network preference table index.
	GslbNewCfgEnhNetworkVirtIndx int32
	// The new network preference server table index.
	GslbNewCfgEnhNetworkVirtServerIndx int32
	Params                             *GslbNewCfgEnhNetworkVirtPrefTableParams
}

func NewGslbNewCfgEnhNetworkVirtPrefTableList() *GslbNewCfgEnhNetworkVirtPrefTable {
	return &GslbNewCfgEnhNetworkVirtPrefTable{}
}

func NewGslbNewCfgEnhNetworkVirtPrefTable(
	gslbNewCfgEnhNetworkVirtIndx int32,
	gslbNewCfgEnhNetworkVirtServerIndx int32,
	params *GslbNewCfgEnhNetworkVirtPrefTableParams,
) *GslbNewCfgEnhNetworkVirtPrefTable {
	return &GslbNewCfgEnhNetworkVirtPrefTable{
		GslbNewCfgEnhNetworkVirtIndx:       gslbNewCfgEnhNetworkVirtIndx,
		GslbNewCfgEnhNetworkVirtServerIndx: gslbNewCfgEnhNetworkVirtServerIndx,
		Params:                             params,
	}
}

func (c *GslbNewCfgEnhNetworkVirtPrefTable) Name() string {
	return "GslbNewCfgEnhNetworkVirtPrefTable"
}

func (c *GslbNewCfgEnhNetworkVirtPrefTable) GetParams() BeanType {
	return c.Params
}

func (c *GslbNewCfgEnhNetworkVirtPrefTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *GslbNewCfgEnhNetworkVirtPrefTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.GslbNewCfgEnhNetworkVirtIndx).IsZero() &&
		reflect.ValueOf(c.GslbNewCfgEnhNetworkVirtServerIndx).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.GslbNewCfgEnhNetworkVirtIndx) + "/" + fmt.Sprint(c.GslbNewCfgEnhNetworkVirtServerIndx)
}

type GslbNewCfgEnhNetworkVirtPrefTableParams struct {
	// The new network preference table index.
	Indx int32 `json:"Indx,omitempty"`
	// The new network preference server table index.
	ServerIndx int32 `json:"ServerIndx,omitempty"`
	// Add virtual server preference to network.
	// Set Zero to remove virtual server from network
	ServerPref int32 `json:"ServerPref,omitempty"`
}

func (p GslbNewCfgEnhNetworkVirtPrefTableParams) iMABean() {}
