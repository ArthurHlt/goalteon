package beans

import (
	"fmt"
	"reflect"
)

// GslbNewCfgEnhNetAlphaNumRealPrefTable The table of network remote real preference in the new_config.
// Note:This mib is not supported for VX instance of Virtualization.
type GslbNewCfgEnhNetAlphaNumRealPrefTable struct {
	// The new network preference table index.
	GslbNewCfgEnhNetAlphaNumRealIndx int32
	// The new network preference server table index.
	GslbNewCfgEnhNetAlphaNumRealServerIndx string
	Params                                 *GslbNewCfgEnhNetAlphaNumRealPrefTableParams
}

func NewGslbNewCfgEnhNetAlphaNumRealPrefTableList() *GslbNewCfgEnhNetAlphaNumRealPrefTable {
	return &GslbNewCfgEnhNetAlphaNumRealPrefTable{}
}

func NewGslbNewCfgEnhNetAlphaNumRealPrefTable(
	gslbNewCfgEnhNetAlphaNumRealIndx int32,
	gslbNewCfgEnhNetAlphaNumRealServerIndx string,
	params *GslbNewCfgEnhNetAlphaNumRealPrefTableParams,
) *GslbNewCfgEnhNetAlphaNumRealPrefTable {
	return &GslbNewCfgEnhNetAlphaNumRealPrefTable{
		GslbNewCfgEnhNetAlphaNumRealIndx:       gslbNewCfgEnhNetAlphaNumRealIndx,
		GslbNewCfgEnhNetAlphaNumRealServerIndx: gslbNewCfgEnhNetAlphaNumRealServerIndx,
		Params:                                 params,
	}
}

func (c *GslbNewCfgEnhNetAlphaNumRealPrefTable) Name() string {
	return "GslbNewCfgEnhNetAlphaNumRealPrefTable"
}

func (c *GslbNewCfgEnhNetAlphaNumRealPrefTable) GetParams() BeanType {
	return c.Params
}

func (c *GslbNewCfgEnhNetAlphaNumRealPrefTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *GslbNewCfgEnhNetAlphaNumRealPrefTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.GslbNewCfgEnhNetAlphaNumRealIndx).IsZero() &&
		reflect.ValueOf(c.GslbNewCfgEnhNetAlphaNumRealServerIndx).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.GslbNewCfgEnhNetAlphaNumRealIndx) + "/" + fmt.Sprint(c.GslbNewCfgEnhNetAlphaNumRealServerIndx)
}

type GslbNewCfgEnhNetAlphaNumRealPrefTableParams struct {
	// The new network preference table index.
	Indx int32 `json:"Indx,omitempty"`
	// The new network preference server table index.
	ServerIndx string `json:"ServerIndx,omitempty"`
	// Add remote real server preference to network.
	// Set Zero to remove the remote real server from network
	ServerPref int32 `json:"ServerPref,omitempty"`
	// Del remote real server preference to network.
	// Set Zero to remove the remote real server from network.
	// When read, the value '0' is always returned
	ServerDelete int32 `json:"ServerDelete,omitempty"`
}

func (p GslbNewCfgEnhNetAlphaNumRealPrefTableParams) iMABean() {}
