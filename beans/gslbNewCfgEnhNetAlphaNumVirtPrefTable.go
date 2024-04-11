package beans

import (
	"fmt"
	"reflect"
)

// GslbNewCfgEnhNetAlphaNumVirtPrefTable The table of network virt preference in the new_config.
// Note:This mib is not supported for VX instance of Virtualization.
type GslbNewCfgEnhNetAlphaNumVirtPrefTable struct {
	// The new network preference table index.
	GslbNewCfgEnhNetAlphaNumVirtIndx int32
	// The new network preference server table index.
	GslbNewCfgEnhNetAlphaNumVirtServerIndx string
	Params                                 *GslbNewCfgEnhNetAlphaNumVirtPrefTableParams
}

func NewGslbNewCfgEnhNetAlphaNumVirtPrefTableList() *GslbNewCfgEnhNetAlphaNumVirtPrefTable {
	return &GslbNewCfgEnhNetAlphaNumVirtPrefTable{}
}

func NewGslbNewCfgEnhNetAlphaNumVirtPrefTable(
	gslbNewCfgEnhNetAlphaNumVirtIndx int32,
	gslbNewCfgEnhNetAlphaNumVirtServerIndx string,
	params *GslbNewCfgEnhNetAlphaNumVirtPrefTableParams,
) *GslbNewCfgEnhNetAlphaNumVirtPrefTable {
	return &GslbNewCfgEnhNetAlphaNumVirtPrefTable{
		GslbNewCfgEnhNetAlphaNumVirtIndx:       gslbNewCfgEnhNetAlphaNumVirtIndx,
		GslbNewCfgEnhNetAlphaNumVirtServerIndx: gslbNewCfgEnhNetAlphaNumVirtServerIndx,
		Params:                                 params,
	}
}

func (c *GslbNewCfgEnhNetAlphaNumVirtPrefTable) Name() string {
	return "GslbNewCfgEnhNetAlphaNumVirtPrefTable"
}

func (c *GslbNewCfgEnhNetAlphaNumVirtPrefTable) GetParams() BeanType {
	return c.Params
}

func (c *GslbNewCfgEnhNetAlphaNumVirtPrefTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *GslbNewCfgEnhNetAlphaNumVirtPrefTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.GslbNewCfgEnhNetAlphaNumVirtIndx).IsZero() &&
		reflect.ValueOf(c.GslbNewCfgEnhNetAlphaNumVirtServerIndx).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.GslbNewCfgEnhNetAlphaNumVirtIndx) + "/" + fmt.Sprint(c.GslbNewCfgEnhNetAlphaNumVirtServerIndx)
}

type GslbNewCfgEnhNetAlphaNumVirtPrefTableParams struct {
	// The new network preference table index.
	Indx int32 `json:"Indx,omitempty"`
	// The new network preference server table index.
	ServerIndx string `json:"ServerIndx,omitempty"`
	// Add virtual server preference to network.
	// Set Zero to remove virtual server from network.
	// 	When read, the value '0' is always returned.
	ServerPref int32 `json:"ServerPref,omitempty"`
	// rem virtual server preference to network.
	// Set Zero to remove virtual server from network.
	// When read, the value '0' is always returned
	ServerDelete int32 `json:"ServerDelete,omitempty"`
}

func (p GslbNewCfgEnhNetAlphaNumVirtPrefTableParams) iMABean() {}
