package beans

import (
	"fmt"
	"reflect"
)

// GslbEnhDnsResVipEmptyIndexesTable The table will display first six empty DNS Responders VIP indexes.
// Note:This mib is not supported for VX instance of Virtualization.
type GslbEnhDnsResVipEmptyIndexesTable struct {
	// DNS Responder VIP Index.
	GslbEnhDnsResVipEmptyIndexesIndex string
	Params                            *GslbEnhDnsResVipEmptyIndexesTableParams
}

func NewGslbEnhDnsResVipEmptyIndexesTableList() *GslbEnhDnsResVipEmptyIndexesTable {
	return &GslbEnhDnsResVipEmptyIndexesTable{}
}

func NewGslbEnhDnsResVipEmptyIndexesTable(
	gslbEnhDnsResVipEmptyIndexesIndex string,
	params *GslbEnhDnsResVipEmptyIndexesTableParams,
) *GslbEnhDnsResVipEmptyIndexesTable {
	return &GslbEnhDnsResVipEmptyIndexesTable{
		GslbEnhDnsResVipEmptyIndexesIndex: gslbEnhDnsResVipEmptyIndexesIndex,
		Params:                            params,
	}
}

func (c *GslbEnhDnsResVipEmptyIndexesTable) Name() string {
	return "GslbEnhDnsResVipEmptyIndexesTable"
}

func (c *GslbEnhDnsResVipEmptyIndexesTable) GetParams() BeanType {
	return c.Params
}

func (c *GslbEnhDnsResVipEmptyIndexesTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *GslbEnhDnsResVipEmptyIndexesTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.GslbEnhDnsResVipEmptyIndexesIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.GslbEnhDnsResVipEmptyIndexesIndex)
}

type GslbEnhDnsResVipEmptyIndexesTableParams struct {
	// DNS Responder VIP Index.
	Index string `json:"Index,omitempty"`
}

func (p GslbEnhDnsResVipEmptyIndexesTableParams) iMABean() {}
