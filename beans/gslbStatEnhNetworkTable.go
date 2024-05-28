package beans

import (
	"fmt"
	"reflect"
)

// GslbStatEnhNetworkTable The Global SLB network statistics table.
// Note:This mib is not supported for VX instance of Virtualization.
type GslbStatEnhNetworkTable struct {
	// The network number that identifies the network.
	GslbStatEnhNetworkIdx int32
	Params                *GslbStatEnhNetworkTableParams
}

func NewGslbStatEnhNetworkTableList() *GslbStatEnhNetworkTable {
	return &GslbStatEnhNetworkTable{}
}

func NewGslbStatEnhNetworkTable(
	gslbStatEnhNetworkIdx int32,
	params *GslbStatEnhNetworkTableParams,
) *GslbStatEnhNetworkTable {
	return &GslbStatEnhNetworkTable{
		GslbStatEnhNetworkIdx: gslbStatEnhNetworkIdx,
		Params:                params,
	}
}

func (c *GslbStatEnhNetworkTable) Name() string {
	return "GslbStatEnhNetworkTable"
}

func (c *GslbStatEnhNetworkTable) GetParams() BeanType {
	return c.Params
}

func (c *GslbStatEnhNetworkTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *GslbStatEnhNetworkTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.GslbStatEnhNetworkIdx).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.GslbStatEnhNetworkIdx)
}

type GslbStatEnhNetworkTableParams struct {
	// The network number that identifies the network.
	Idx int32 `json:"Idx,omitempty"`
	// The number of times network is selected.
	Hit uint32 `json:"Hit,omitempty"`
	// The network source IP address.
	IpAddr string `json:"IpAddr,omitempty"`
	// The network source IPv6 IP address.
	Ipv6Addr string `json:"Ipv6Addr,omitempty"`
	// Source address type (1)address (2)nwclass.
	SrcType int32 `json:"SrcType,omitempty"`
	// The network class and its region.
	ClassId string `json:"ClassId,omitempty"`
}

func (p GslbStatEnhNetworkTableParams) iMABean() {}
