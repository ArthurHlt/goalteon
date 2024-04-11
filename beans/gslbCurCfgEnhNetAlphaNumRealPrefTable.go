package beans

import (
	"fmt"
	"reflect"
)

// GslbCurCfgEnhNetAlphaNumRealPrefTable The table of network remote real  preference in the cur_config.
// Note:This mib is not supported for VX instance of Virtualization.
type GslbCurCfgEnhNetAlphaNumRealPrefTable struct {
	// The new network remote real preference  table index.
	GslbCurCfgEnhNetAlphaNumRealIndx int32
	// The new network remote real preference server table index.
	GslbCurCfgEnhNetAlphaNumRealServerIndx string
	Params                                 *GslbCurCfgEnhNetAlphaNumRealPrefTableParams
}

func NewGslbCurCfgEnhNetAlphaNumRealPrefTableList() *GslbCurCfgEnhNetAlphaNumRealPrefTable {
	return &GslbCurCfgEnhNetAlphaNumRealPrefTable{}
}

func NewGslbCurCfgEnhNetAlphaNumRealPrefTable(
	gslbCurCfgEnhNetAlphaNumRealIndx int32,
	gslbCurCfgEnhNetAlphaNumRealServerIndx string,
	params *GslbCurCfgEnhNetAlphaNumRealPrefTableParams,
) *GslbCurCfgEnhNetAlphaNumRealPrefTable {
	return &GslbCurCfgEnhNetAlphaNumRealPrefTable{
		GslbCurCfgEnhNetAlphaNumRealIndx:       gslbCurCfgEnhNetAlphaNumRealIndx,
		GslbCurCfgEnhNetAlphaNumRealServerIndx: gslbCurCfgEnhNetAlphaNumRealServerIndx,
		Params:                                 params,
	}
}

func (c *GslbCurCfgEnhNetAlphaNumRealPrefTable) Name() string {
	return "GslbCurCfgEnhNetAlphaNumRealPrefTable"
}

func (c *GslbCurCfgEnhNetAlphaNumRealPrefTable) GetParams() BeanType {
	return c.Params
}

func (c *GslbCurCfgEnhNetAlphaNumRealPrefTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *GslbCurCfgEnhNetAlphaNumRealPrefTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.GslbCurCfgEnhNetAlphaNumRealIndx).IsZero() &&
		reflect.ValueOf(c.GslbCurCfgEnhNetAlphaNumRealServerIndx).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.GslbCurCfgEnhNetAlphaNumRealIndx) + "/" + fmt.Sprint(c.GslbCurCfgEnhNetAlphaNumRealServerIndx)
}

type GslbCurCfgEnhNetAlphaNumRealPrefTableParams struct {
	// The new network remote real preference  table index.
	Indx int32 `json:"Indx,omitempty"`
	// The new network remote real preference server table index.
	ServerIndx string `json:"ServerIndx,omitempty"`
	// Add remote real server preference to network
	ServerPref int32 `json:"ServerPref,omitempty"`
}

func (p GslbCurCfgEnhNetAlphaNumRealPrefTableParams) iMABean() {}
