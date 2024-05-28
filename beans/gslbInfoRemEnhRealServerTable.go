package beans

import (
	"fmt"
	"reflect"
)

// GslbInfoRemEnhRealServerTable The Global SLB virtual server information table.
// Note:This mib is not supported for VX instance of Virtualization.
type GslbInfoRemEnhRealServerTable struct {
	// The remote real server id that identifies the remote real server.
	GslbInfoRemEnhRealServerIdx string
	Params                      *GslbInfoRemEnhRealServerTableParams
}

func NewGslbInfoRemEnhRealServerTableList() *GslbInfoRemEnhRealServerTable {
	return &GslbInfoRemEnhRealServerTable{}
}

func NewGslbInfoRemEnhRealServerTable(
	gslbInfoRemEnhRealServerIdx string,
	params *GslbInfoRemEnhRealServerTableParams,
) *GslbInfoRemEnhRealServerTable {
	return &GslbInfoRemEnhRealServerTable{
		GslbInfoRemEnhRealServerIdx: gslbInfoRemEnhRealServerIdx,
		Params:                      params,
	}
}

func (c *GslbInfoRemEnhRealServerTable) Name() string {
	return "GslbInfoRemEnhRealServerTable"
}

func (c *GslbInfoRemEnhRealServerTable) GetParams() BeanType {
	return c.Params
}

func (c *GslbInfoRemEnhRealServerTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *GslbInfoRemEnhRealServerTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.GslbInfoRemEnhRealServerIdx).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.GslbInfoRemEnhRealServerIdx)
}

type GslbInfoRemEnhRealServerTableState int32

const (
	GslbInfoRemEnhRealServerTableState_Running     GslbInfoRemEnhRealServerTableState = 2
	GslbInfoRemEnhRealServerTableState_Failed      GslbInfoRemEnhRealServerTableState = 3
	GslbInfoRemEnhRealServerTableState_Disabled    GslbInfoRemEnhRealServerTableState = 4
	GslbInfoRemEnhRealServerTableState_Unsupported GslbInfoRemEnhRealServerTableState = 2147483647
)

type GslbInfoRemEnhRealServerTableParams struct {
	// The remote real server id that identifies the remote real server.
	Idx string `json:"Idx,omitempty"`
	// IP address of the remote real server.
	IpAddr string `json:"IpAddr,omitempty"`
	// The name of the remote real server.
	Name string `json:"Name,omitempty"`
	// The state of the remote real server.
	State GslbInfoRemEnhRealServerTableState `json:"State,omitempty"`
}

func (p GslbInfoRemEnhRealServerTableParams) iMABean() {}
