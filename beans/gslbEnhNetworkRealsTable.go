package beans

import (
	"fmt"
	"reflect"
)

// GslbEnhNetworkRealsTable The table of network Reals in the new config.
// Note:This mib is not supported for VX instance of Virtualization.
type GslbEnhNetworkRealsTable struct {
	// The network table index.
	GslbEnhNetworkIndx int32
	// The network real table AlphaNumeric index.
	GslbEnhNetworkRealIndx string
	Params                 *GslbEnhNetworkRealsTableParams
}

func NewGslbEnhNetworkRealsTableList() *GslbEnhNetworkRealsTable {
	return &GslbEnhNetworkRealsTable{}
}

func NewGslbEnhNetworkRealsTable(
	gslbEnhNetworkIndx int32,
	gslbEnhNetworkRealIndx string,
	params *GslbEnhNetworkRealsTableParams,
) *GslbEnhNetworkRealsTable {
	return &GslbEnhNetworkRealsTable{
		GslbEnhNetworkIndx:     gslbEnhNetworkIndx,
		GslbEnhNetworkRealIndx: gslbEnhNetworkRealIndx,
		Params:                 params,
	}
}

func (c *GslbEnhNetworkRealsTable) Name() string {
	return "GslbEnhNetworkRealsTable"
}

func (c *GslbEnhNetworkRealsTable) GetParams() BeanType {
	return c.Params
}

func (c *GslbEnhNetworkRealsTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *GslbEnhNetworkRealsTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.GslbEnhNetworkIndx).IsZero() &&
		reflect.ValueOf(c.GslbEnhNetworkRealIndx).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.GslbEnhNetworkIndx) + "/" + fmt.Sprint(c.GslbEnhNetworkRealIndx)
}

type GslbEnhNetworkRealsTableParams struct {
	// The network table index.
	Indx int32 `json:"Indx,omitempty"`
	// The network real table AlphaNumeric index.
	RealIndx string `json:"RealIndx,omitempty"`
}

func (p GslbEnhNetworkRealsTableParams) iMABean() {}
