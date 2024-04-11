package beans

import (
	"fmt"
	"reflect"
)

// GslbDnsResVipEmptyIndexesTable Note:This mib is not supported for VX instance of Virtualization.
type GslbDnsResVipEmptyIndexesTable struct {
	// DNS Responder VIP Index.
	GslbDnsResVipEmptyIndexesIndex uint32
	Params                         *GslbDnsResVipEmptyIndexesTableParams
}

func NewGslbDnsResVipEmptyIndexesTableList() *GslbDnsResVipEmptyIndexesTable {
	return &GslbDnsResVipEmptyIndexesTable{}
}

func NewGslbDnsResVipEmptyIndexesTable(
	gslbDnsResVipEmptyIndexesIndex uint32,
	params *GslbDnsResVipEmptyIndexesTableParams,
) *GslbDnsResVipEmptyIndexesTable {
	return &GslbDnsResVipEmptyIndexesTable{
		GslbDnsResVipEmptyIndexesIndex: gslbDnsResVipEmptyIndexesIndex,
		Params:                         params,
	}
}

func (c *GslbDnsResVipEmptyIndexesTable) Name() string {
	return "GslbDnsResVipEmptyIndexesTable"
}

func (c *GslbDnsResVipEmptyIndexesTable) GetParams() BeanType {
	return c.Params
}

func (c *GslbDnsResVipEmptyIndexesTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *GslbDnsResVipEmptyIndexesTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.GslbDnsResVipEmptyIndexesIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.GslbDnsResVipEmptyIndexesIndex)
}

type GslbDnsResVipEmptyIndexesTableParams struct {
	// DNS Responder VIP Index.
	Index uint32 `json:"Index,omitempty"`
}

func (p GslbDnsResVipEmptyIndexesTableParams) iMABean() {}
