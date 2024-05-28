package beans

import (
	"fmt"
	"reflect"
)

// GslbStatRemEnhRealServerTable The GSLB remote real server statistics table.
// Note:This mib is not supported for VX instance of Virtualization.
type GslbStatRemEnhRealServerTable struct {
	// The server number that identifies the remote real server.
	GslbStatRemEnhRealServerIndex string
	Params                        *GslbStatRemEnhRealServerTableParams
}

func NewGslbStatRemEnhRealServerTableList() *GslbStatRemEnhRealServerTable {
	return &GslbStatRemEnhRealServerTable{}
}

func NewGslbStatRemEnhRealServerTable(
	gslbStatRemEnhRealServerIndex string,
	params *GslbStatRemEnhRealServerTableParams,
) *GslbStatRemEnhRealServerTable {
	return &GslbStatRemEnhRealServerTable{
		GslbStatRemEnhRealServerIndex: gslbStatRemEnhRealServerIndex,
		Params:                        params,
	}
}

func (c *GslbStatRemEnhRealServerTable) Name() string {
	return "GslbStatRemEnhRealServerTable"
}

func (c *GslbStatRemEnhRealServerTable) GetParams() BeanType {
	return c.Params
}

func (c *GslbStatRemEnhRealServerTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *GslbStatRemEnhRealServerTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.GslbStatRemEnhRealServerIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.GslbStatRemEnhRealServerIndex)
}

type GslbStatRemEnhRealServerTableIpVer int32

const (
	GslbStatRemEnhRealServerTableIpVer_Ipv4        GslbStatRemEnhRealServerTableIpVer = 1
	GslbStatRemEnhRealServerTableIpVer_Ipv6        GslbStatRemEnhRealServerTableIpVer = 2
	GslbStatRemEnhRealServerTableIpVer_Unsupported GslbStatRemEnhRealServerTableIpVer = 2147483647
)

type GslbStatRemEnhRealServerTableParams struct {
	// The server number that identifies the remote real server.
	Index string `json:"Index,omitempty"`
	// The number of DNS handoffs by the remote real server.
	DnsHandoffs uint32 `json:"DnsHandoffs,omitempty"`
	// The number of HTTP redirections by the remote real server.
	HttpRedirs uint32 `json:"HttpRedirs,omitempty"`
	// The total number of hits exceeded threshold for remote real server.
	ThresholdExceeded uint32 `json:"ThresholdExceeded,omitempty"`
	// The IP address of the remote real server.
	IpAddress string `json:"IpAddress,omitempty"`
	// The IPv6 address of the remote real server.
	Ipv6Address string `json:"Ipv6Address,omitempty"`
	// The type of IP address of the remote real server.
	IpVer GslbStatRemEnhRealServerTableIpVer `json:"IpVer,omitempty"`
}

func (p GslbStatRemEnhRealServerTableParams) iMABean() {}
