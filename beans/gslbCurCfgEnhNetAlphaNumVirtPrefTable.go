package beans

import (
	"fmt"
	"reflect"
)

// GslbCurCfgEnhNetAlphaNumVirtPrefTable The table of network virt  preference in the cur_config.
// Note:This mib is not supported for VX instance of Virtualization.
type GslbCurCfgEnhNetAlphaNumVirtPrefTable struct {
	// The current network preference table index.
	GslbCurCfgEnhNetAlphaNumVirtIndx int32
	// The current network preference table server index.
	GslbCurCfgEnhNetAlphaNumVirtServerIndx string
	Params                                 *GslbCurCfgEnhNetAlphaNumVirtPrefTableParams
}

func NewGslbCurCfgEnhNetAlphaNumVirtPrefTableList() *GslbCurCfgEnhNetAlphaNumVirtPrefTable {
	return &GslbCurCfgEnhNetAlphaNumVirtPrefTable{}
}

func NewGslbCurCfgEnhNetAlphaNumVirtPrefTable(
	gslbCurCfgEnhNetAlphaNumVirtIndx int32,
	gslbCurCfgEnhNetAlphaNumVirtServerIndx string,
	params *GslbCurCfgEnhNetAlphaNumVirtPrefTableParams,
) *GslbCurCfgEnhNetAlphaNumVirtPrefTable {
	return &GslbCurCfgEnhNetAlphaNumVirtPrefTable{
		GslbCurCfgEnhNetAlphaNumVirtIndx:       gslbCurCfgEnhNetAlphaNumVirtIndx,
		GslbCurCfgEnhNetAlphaNumVirtServerIndx: gslbCurCfgEnhNetAlphaNumVirtServerIndx,
		Params:                                 params,
	}
}

func (c *GslbCurCfgEnhNetAlphaNumVirtPrefTable) Name() string {
	return "GslbCurCfgEnhNetAlphaNumVirtPrefTable"
}

func (c *GslbCurCfgEnhNetAlphaNumVirtPrefTable) GetParams() BeanType {
	return c.Params
}

func (c *GslbCurCfgEnhNetAlphaNumVirtPrefTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *GslbCurCfgEnhNetAlphaNumVirtPrefTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.GslbCurCfgEnhNetAlphaNumVirtIndx).IsZero() &&
		reflect.ValueOf(c.GslbCurCfgEnhNetAlphaNumVirtServerIndx).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.GslbCurCfgEnhNetAlphaNumVirtIndx) + "/" + fmt.Sprint(c.GslbCurCfgEnhNetAlphaNumVirtServerIndx)
}

type GslbCurCfgEnhNetAlphaNumVirtPrefTableParams struct {
	// The current network preference table index.
	Indx int32 `json:"Indx,omitempty"`
	// The current network preference table server index.
	ServerIndx string `json:"ServerIndx,omitempty"`
	// current virtual server preference to network
	ServerPref int32 `json:"ServerPref,omitempty"`
}

func (p GslbCurCfgEnhNetAlphaNumVirtPrefTableParams) iMABean() {}
