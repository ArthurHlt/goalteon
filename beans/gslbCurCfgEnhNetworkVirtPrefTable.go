package beans

import (
	"fmt"
	"reflect"
)

// GslbCurCfgEnhNetworkVirtPrefTable The table of network virt  preference in the cur_config.
// Note:This mib is not supported for VX instance of Virtualization.
type GslbCurCfgEnhNetworkVirtPrefTable struct {
	// The current network preference table index.
	GslbCurCfgEnhNetworkVirtIndx int32
	// The current network preference table server index.
	GslbCurCfgEnhNetworkVirtServerIndx int32
	Params                             *GslbCurCfgEnhNetworkVirtPrefTableParams
}

func NewGslbCurCfgEnhNetworkVirtPrefTableList() *GslbCurCfgEnhNetworkVirtPrefTable {
	return &GslbCurCfgEnhNetworkVirtPrefTable{}
}

func NewGslbCurCfgEnhNetworkVirtPrefTable(
	gslbCurCfgEnhNetworkVirtIndx int32,
	gslbCurCfgEnhNetworkVirtServerIndx int32,
	params *GslbCurCfgEnhNetworkVirtPrefTableParams,
) *GslbCurCfgEnhNetworkVirtPrefTable {
	return &GslbCurCfgEnhNetworkVirtPrefTable{
		GslbCurCfgEnhNetworkVirtIndx:       gslbCurCfgEnhNetworkVirtIndx,
		GslbCurCfgEnhNetworkVirtServerIndx: gslbCurCfgEnhNetworkVirtServerIndx,
		Params:                             params,
	}
}

func (c *GslbCurCfgEnhNetworkVirtPrefTable) Name() string {
	return "GslbCurCfgEnhNetworkVirtPrefTable"
}

func (c *GslbCurCfgEnhNetworkVirtPrefTable) GetParams() BeanType {
	return c.Params
}

func (c *GslbCurCfgEnhNetworkVirtPrefTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *GslbCurCfgEnhNetworkVirtPrefTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.GslbCurCfgEnhNetworkVirtIndx).IsZero() &&
		reflect.ValueOf(c.GslbCurCfgEnhNetworkVirtServerIndx).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.GslbCurCfgEnhNetworkVirtIndx) + "/" + fmt.Sprint(c.GslbCurCfgEnhNetworkVirtServerIndx)
}

type GslbCurCfgEnhNetworkVirtPrefTableParams struct {
	// The current network preference table index.
	Indx int32 `json:"Indx,omitempty"`
	// The current network preference table server index.
	ServerIndx int32 `json:"ServerIndx,omitempty"`
	// current virtual server preference to network
	ServerPref int32 `json:"ServerPref,omitempty"`
}

func (p GslbCurCfgEnhNetworkVirtPrefTableParams) iMABean() {}
