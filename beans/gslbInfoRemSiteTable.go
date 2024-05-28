package beans

import (
	"fmt"
	"reflect"
)

// GslbInfoRemSiteTable The Global SLB remote site information table.
// Note:This mib is not supported for VX instance of Virtualization.
type GslbInfoRemSiteTable struct {
	// The remote site number that identifies the remote site.
	GslbInfoRemSiteIdx int32
	Params             *GslbInfoRemSiteTableParams
}

func NewGslbInfoRemSiteTableList() *GslbInfoRemSiteTable {
	return &GslbInfoRemSiteTable{}
}

func NewGslbInfoRemSiteTable(
	gslbInfoRemSiteIdx int32,
	params *GslbInfoRemSiteTableParams,
) *GslbInfoRemSiteTable {
	return &GslbInfoRemSiteTable{
		GslbInfoRemSiteIdx: gslbInfoRemSiteIdx,
		Params:             params,
	}
}

func (c *GslbInfoRemSiteTable) Name() string {
	return "GslbInfoRemSiteTable"
}

func (c *GslbInfoRemSiteTable) GetParams() BeanType {
	return c.Params
}

func (c *GslbInfoRemSiteTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *GslbInfoRemSiteTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.GslbInfoRemSiteIdx).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.GslbInfoRemSiteIdx)
}

type GslbInfoRemSiteTableState int32

const (
	GslbInfoRemSiteTableState_Running     GslbInfoRemSiteTableState = 2
	GslbInfoRemSiteTableState_Failed      GslbInfoRemSiteTableState = 3
	GslbInfoRemSiteTableState_Disabled    GslbInfoRemSiteTableState = 4
	GslbInfoRemSiteTableState_Unsupported GslbInfoRemSiteTableState = 2147483647
)

type GslbInfoRemSiteTableParams struct {
	// The remote site number that identifies the remote site.
	Idx int32 `json:"Idx,omitempty"`
	// The primary IP address of the remote site.
	PrimaryIp string `json:"PrimaryIp,omitempty"`
	// The secondary IP address of the remote site.
	SecondaryIp string `json:"SecondaryIp,omitempty"`
	// The name of the remote site.
	Name string `json:"Name,omitempty"`
	// The state of the remote site.
	State GslbInfoRemSiteTableState `json:"State,omitempty"`
}

func (p GslbInfoRemSiteTableParams) iMABean() {}
