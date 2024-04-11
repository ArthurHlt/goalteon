package beans

import (
	"fmt"
	"reflect"
)

// GslbNewDnsResVipTable Note:This mib is not supported for VX instance of Virtualization.
type GslbNewDnsResVipTable struct {
	// DNS Responder VIP Index 1.
	GslbNewDnsResVipIndex1 uint32
	// DNS Responder VIP Index 2.
	GslbNewDnsResVipIndex2 uint32
	Params                 *GslbNewDnsResVipTableParams
}

func NewGslbNewDnsResVipTableList() *GslbNewDnsResVipTable {
	return &GslbNewDnsResVipTable{}
}

func NewGslbNewDnsResVipTable(
	gslbNewDnsResVipIndex1 uint32,
	gslbNewDnsResVipIndex2 uint32,
	params *GslbNewDnsResVipTableParams,
) *GslbNewDnsResVipTable {
	return &GslbNewDnsResVipTable{
		GslbNewDnsResVipIndex1: gslbNewDnsResVipIndex1,
		GslbNewDnsResVipIndex2: gslbNewDnsResVipIndex2,
		Params:                 params,
	}
}

func (c *GslbNewDnsResVipTable) Name() string {
	return "GslbNewDnsResVipTable"
}

func (c *GslbNewDnsResVipTable) GetParams() BeanType {
	return c.Params
}

func (c *GslbNewDnsResVipTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *GslbNewDnsResVipTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.GslbNewDnsResVipIndex1).IsZero() &&
		reflect.ValueOf(c.GslbNewDnsResVipIndex2).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.GslbNewDnsResVipIndex1) + "/" + fmt.Sprint(c.GslbNewDnsResVipIndex2)
}

type GslbNewDnsResVipTableIPVer int32

const (
	GslbNewDnsResVipTableIPVer_Ipv4        GslbNewDnsResVipTableIPVer = 1
	GslbNewDnsResVipTableIPVer_Ipv6        GslbNewDnsResVipTableIPVer = 2
	GslbNewDnsResVipTableIPVer_Unsupported GslbNewDnsResVipTableIPVer = 2147483647
)

type GslbNewDnsResVipTableStatus int32

const (
	GslbNewDnsResVipTableStatus_Enable      GslbNewDnsResVipTableStatus = 1
	GslbNewDnsResVipTableStatus_Disable     GslbNewDnsResVipTableStatus = 2
	GslbNewDnsResVipTableStatus_Unsupported GslbNewDnsResVipTableStatus = 2147483647
)

type GslbNewDnsResVipTableDelete int32

const (
	GslbNewDnsResVipTableDelete_Other       GslbNewDnsResVipTableDelete = 1
	GslbNewDnsResVipTableDelete_Delete      GslbNewDnsResVipTableDelete = 2
	GslbNewDnsResVipTableDelete_Unsupported GslbNewDnsResVipTableDelete = 2147483647
)

type GslbNewDnsResVipTableParams struct {
	// DNS Responder VIP Index 1.
	Index1 uint32 `json:"Index1,omitempty"`
	// DNS Responder VIP Index 2.
	Index2 uint32 `json:"Index2,omitempty"`
	// DNS Responder VIP Name.
	Name string `json:"Name,omitempty"`
	// The type of IP address of DNS Responder VIP.
	IPVer GslbNewDnsResVipTableIPVer `json:"IPVer,omitempty"`
	// IPV4 address of the DNS Responder VIP.
	V4 string `json:"V4,omitempty"`
	// IPV6 address of the DNS Responder VIP
	// Address should be 4-byte hexadecimal colon notation. Valid IPv6 address should be in
	// any of the following forms xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx or xxxx::xxxx:xxxx:xxxx:xxxx or ::xxxx
	V6 string `json:"V6,omitempty"`
	// Zone status - Enabled/Disabled.
	Status GslbNewDnsResVipTableStatus `json:"Status,omitempty"`
	// By setting the value to delete(2), the entire row is deleted.
	Delete GslbNewDnsResVipTableDelete `json:"Delete,omitempty"`
}

func (p GslbNewDnsResVipTableParams) iMABean() {}
