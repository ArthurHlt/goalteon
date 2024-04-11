package beans

import (
	"fmt"
	"reflect"
)

// GslbCurDnsResVipTable Note:This mib is not supported for VX instance of Virtualization.
type GslbCurDnsResVipTable struct {
	// DNS Responder VIP Index 1.
	GslbCurDnsResVipIndex1 uint32
	// DNS Responder VIP Index 2.
	GslbCurDnsResVipIndex2 uint32
	Params                 *GslbCurDnsResVipTableParams
}

func NewGslbCurDnsResVipTableList() *GslbCurDnsResVipTable {
	return &GslbCurDnsResVipTable{}
}

func NewGslbCurDnsResVipTable(
	gslbCurDnsResVipIndex1 uint32,
	gslbCurDnsResVipIndex2 uint32,
	params *GslbCurDnsResVipTableParams,
) *GslbCurDnsResVipTable {
	return &GslbCurDnsResVipTable{
		GslbCurDnsResVipIndex1: gslbCurDnsResVipIndex1,
		GslbCurDnsResVipIndex2: gslbCurDnsResVipIndex2,
		Params:                 params,
	}
}

func (c *GslbCurDnsResVipTable) Name() string {
	return "GslbCurDnsResVipTable"
}

func (c *GslbCurDnsResVipTable) GetParams() BeanType {
	return c.Params
}

func (c *GslbCurDnsResVipTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *GslbCurDnsResVipTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.GslbCurDnsResVipIndex1).IsZero() &&
		reflect.ValueOf(c.GslbCurDnsResVipIndex2).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.GslbCurDnsResVipIndex1) + "/" + fmt.Sprint(c.GslbCurDnsResVipIndex2)
}

type GslbCurDnsResVipTableIPVer int32

const (
	GslbCurDnsResVipTableIPVer_Ipv4        GslbCurDnsResVipTableIPVer = 1
	GslbCurDnsResVipTableIPVer_Ipv6        GslbCurDnsResVipTableIPVer = 2
	GslbCurDnsResVipTableIPVer_Unsupported GslbCurDnsResVipTableIPVer = 2147483647
)

type GslbCurDnsResVipTableStatus int32

const (
	GslbCurDnsResVipTableStatus_Enable      GslbCurDnsResVipTableStatus = 1
	GslbCurDnsResVipTableStatus_Disable     GslbCurDnsResVipTableStatus = 2
	GslbCurDnsResVipTableStatus_Unsupported GslbCurDnsResVipTableStatus = 2147483647
)

type GslbCurDnsResVipTableParams struct {
	// DNS Responder VIP Index 1.
	Index1 uint32 `json:"Index1,omitempty"`
	// DNS Responder VIP Index 2.
	Index2 uint32 `json:"Index2,omitempty"`
	// DNS Responder VIP Name.
	Name string `json:"Name,omitempty"`
	// The type of IP address of DNS Responder VIP.
	IPVer GslbCurDnsResVipTableIPVer `json:"IPVer,omitempty"`
	// IPV4 address of the DNS Responder VIP.
	V4 string `json:"V4,omitempty"`
	// IPV6 address of the DNS Responder VIP
	// Address should be 4-byte hexadecimal colon notation. Valid IPv6 address should be in
	// any of the following forms xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx or xxxx::xxxx:xxxx:xxxx:xxxx or ::xxxx
	V6 string `json:"V6,omitempty"`
	// Zone status - Enabled/Disabled.
	Status GslbCurDnsResVipTableStatus `json:"Status,omitempty"`
}

func (p GslbCurDnsResVipTableParams) iMABean() {}
