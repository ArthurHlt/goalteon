package beans

import (
	"fmt"
	"reflect"
)

// GslbNewEnhDnsResVipTable Note:This mib is not supported for VX instance of Virtualization.
type GslbNewEnhDnsResVipTable struct {
	// DNS Responder VIP Index 1.
	GslbNewEnhDnsResVipIndex1 string
	// DNS Responder VIP Index 2.
	GslbNewEnhDnsResVipIndex2 string
	Params                    *GslbNewEnhDnsResVipTableParams
}

func NewGslbNewEnhDnsResVipTableList() *GslbNewEnhDnsResVipTable {
	return &GslbNewEnhDnsResVipTable{}
}

func NewGslbNewEnhDnsResVipTable(
	gslbNewEnhDnsResVipIndex1 string,
	gslbNewEnhDnsResVipIndex2 string,
	params *GslbNewEnhDnsResVipTableParams,
) *GslbNewEnhDnsResVipTable {
	return &GslbNewEnhDnsResVipTable{
		GslbNewEnhDnsResVipIndex1: gslbNewEnhDnsResVipIndex1,
		GslbNewEnhDnsResVipIndex2: gslbNewEnhDnsResVipIndex2,
		Params:                    params,
	}
}

func (c *GslbNewEnhDnsResVipTable) Name() string {
	return "GslbNewEnhDnsResVipTable"
}

func (c *GslbNewEnhDnsResVipTable) GetParams() BeanType {
	return c.Params
}

func (c *GslbNewEnhDnsResVipTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *GslbNewEnhDnsResVipTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.GslbNewEnhDnsResVipIndex1).IsZero() &&
		reflect.ValueOf(c.GslbNewEnhDnsResVipIndex2).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.GslbNewEnhDnsResVipIndex1) + "/" + fmt.Sprint(c.GslbNewEnhDnsResVipIndex2)
}

type GslbNewEnhDnsResVipTableIPVer int32

const (
	GslbNewEnhDnsResVipTableIPVer_Ipv4        GslbNewEnhDnsResVipTableIPVer = 1
	GslbNewEnhDnsResVipTableIPVer_Ipv6        GslbNewEnhDnsResVipTableIPVer = 2
	GslbNewEnhDnsResVipTableIPVer_Unsupported GslbNewEnhDnsResVipTableIPVer = 2147483647
)

type GslbNewEnhDnsResVipTableStatus int32

const (
	GslbNewEnhDnsResVipTableStatus_Enable      GslbNewEnhDnsResVipTableStatus = 1
	GslbNewEnhDnsResVipTableStatus_Disable     GslbNewEnhDnsResVipTableStatus = 2
	GslbNewEnhDnsResVipTableStatus_Unsupported GslbNewEnhDnsResVipTableStatus = 2147483647
)

type GslbNewEnhDnsResVipTableDelete int32

const (
	GslbNewEnhDnsResVipTableDelete_Other       GslbNewEnhDnsResVipTableDelete = 1
	GslbNewEnhDnsResVipTableDelete_Delete      GslbNewEnhDnsResVipTableDelete = 2
	GslbNewEnhDnsResVipTableDelete_Unsupported GslbNewEnhDnsResVipTableDelete = 2147483647
)

type GslbNewEnhDnsResVipTableRtsrcmac int32

const (
	GslbNewEnhDnsResVipTableRtsrcmac_Enable      GslbNewEnhDnsResVipTableRtsrcmac = 1
	GslbNewEnhDnsResVipTableRtsrcmac_Disable     GslbNewEnhDnsResVipTableRtsrcmac = 2
	GslbNewEnhDnsResVipTableRtsrcmac_Unsupported GslbNewEnhDnsResVipTableRtsrcmac = 2147483647
)

type GslbNewEnhDnsResVipTableParams struct {
	// DNS Responder VIP Index 1.
	Index1 string `json:"Index1,omitempty"`
	// DNS Responder VIP Index 2.
	Index2 string `json:"Index2,omitempty"`
	// DNS Responder VIP Name.
	Name string `json:"Name,omitempty"`
	// The type of IP address of DNS Responder VIP.
	IPVer GslbNewEnhDnsResVipTableIPVer `json:"IPVer,omitempty"`
	// IPV4 address of the DNS Responder VIP.
	V4 string `json:"V4,omitempty"`
	// IPV6 address of the DNS Responder VIP
	// Address should be 4-byte hexadecimal colon notation. Valid IPv6 address should be in
	// any of the following forms xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx or xxxx::xxxx:xxxx:xxxx:xxxx or ::xxxx
	V6 string `json:"V6,omitempty"`
	// Zone status - Enabled/Disabled.
	Status GslbNewEnhDnsResVipTableStatus `json:"Status,omitempty"`
	// By setting the value to delete(2), the entire row is deleted.
	Delete GslbNewEnhDnsResVipTableDelete `json:"Delete,omitempty"`
	// Enabling/Disabling this will enable/disable rtsrcmac in both virtual servers that comprise this DNS Responder.
	Rtsrcmac GslbNewEnhDnsResVipTableRtsrcmac `json:"Rtsrcmac,omitempty"`
}

func (p GslbNewEnhDnsResVipTableParams) iMABean() {}
