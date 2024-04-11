package beans

import (
	"fmt"
	"reflect"
)

// GslbEnhNetworkVirtsTable The table of network virtual servers in the new config.
// Note:This mib is not supported for VX instance of Virtualization.
type GslbEnhNetworkVirtsTable struct {
	// The network table index.
	GslbEnhNetworkIndex int32
	// The network virtual server table AlphaNumeric index.
	GslbEnhNetworkVirtIndx string
	Params                 *GslbEnhNetworkVirtsTableParams
}

func NewGslbEnhNetworkVirtsTableList() *GslbEnhNetworkVirtsTable {
	return &GslbEnhNetworkVirtsTable{}
}

func NewGslbEnhNetworkVirtsTable(
	gslbEnhNetworkIndex int32,
	gslbEnhNetworkVirtIndx string,
	params *GslbEnhNetworkVirtsTableParams,
) *GslbEnhNetworkVirtsTable {
	return &GslbEnhNetworkVirtsTable{
		GslbEnhNetworkIndex:    gslbEnhNetworkIndex,
		GslbEnhNetworkVirtIndx: gslbEnhNetworkVirtIndx,
		Params:                 params,
	}
}

func (c *GslbEnhNetworkVirtsTable) Name() string {
	return "GslbEnhNetworkVirtsTable"
}

func (c *GslbEnhNetworkVirtsTable) GetParams() BeanType {
	return c.Params
}

func (c *GslbEnhNetworkVirtsTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *GslbEnhNetworkVirtsTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.GslbEnhNetworkIndex).IsZero() &&
		reflect.ValueOf(c.GslbEnhNetworkVirtIndx).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.GslbEnhNetworkIndex) + "/" + fmt.Sprint(c.GslbEnhNetworkVirtIndx)
}

type GslbEnhNetworkVirtsTableParams struct {
	// The network table index.
	Index int32 `json:"Index,omitempty"`
	// The network virtual server table AlphaNumeric index.
	Indx string `json:"Indx,omitempty"`
}

func (p GslbEnhNetworkVirtsTableParams) iMABean() {}
