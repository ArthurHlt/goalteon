package beans

import (
	"fmt"
	"reflect"
)

// GslbInfoRemRealServerTable The Global SLB virtual server information table.
// Note:This mib is not supported for VX instance of Virtualization.
type GslbInfoRemRealServerTable struct {
	// The remote real server number that identifies the remote real server.
	GslbInfoRemRealServerIdx int32
	Params                   *GslbInfoRemRealServerTableParams
}

func NewGslbInfoRemRealServerTableList() *GslbInfoRemRealServerTable {
	return &GslbInfoRemRealServerTable{}
}

func NewGslbInfoRemRealServerTable(
	gslbInfoRemRealServerIdx int32,
	params *GslbInfoRemRealServerTableParams,
) *GslbInfoRemRealServerTable {
	return &GslbInfoRemRealServerTable{
		GslbInfoRemRealServerIdx: gslbInfoRemRealServerIdx,
		Params:                   params,
	}
}

func (c *GslbInfoRemRealServerTable) Name() string {
	return "GslbInfoRemRealServerTable"
}

func (c *GslbInfoRemRealServerTable) GetParams() BeanType {
	return c.Params
}

func (c *GslbInfoRemRealServerTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *GslbInfoRemRealServerTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.GslbInfoRemRealServerIdx).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.GslbInfoRemRealServerIdx)
}

type GslbInfoRemRealServerTableState int32

const (
	GslbInfoRemRealServerTableState_Running     GslbInfoRemRealServerTableState = 2
	GslbInfoRemRealServerTableState_Failed      GslbInfoRemRealServerTableState = 3
	GslbInfoRemRealServerTableState_Disabled    GslbInfoRemRealServerTableState = 4
	GslbInfoRemRealServerTableState_Unsupported GslbInfoRemRealServerTableState = 2147483647
)

type GslbInfoRemRealServerTableParams struct {
	// The remote real server number that identifies the remote real server.
	Idx int32 `json:"Idx,omitempty"`
	// IP address of the remote real server.
	IpAddr string `json:"IpAddr,omitempty"`
	// The name of the remote real server.
	Name string `json:"Name,omitempty"`
	// The state of the remote real server.
	State GslbInfoRemRealServerTableState `json:"State,omitempty"`
}

func (p GslbInfoRemRealServerTableParams) iMABean() {}
