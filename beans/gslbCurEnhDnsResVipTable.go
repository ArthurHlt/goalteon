package beans

import (
	"fmt"
	"reflect"
)

// GslbCurEnhDnsResVipTable Note:This mib is not supported for VX instance of Virtualization.
type GslbCurEnhDnsResVipTable struct {
	// DNS Responder VIP Index 1.
	GslbCurEnhDnsResVipIndex1 string
	// DNS Responder VIP Index 2.
	GslbCurEnhDnsResVipIndex2 string
	Params                    *GslbCurEnhDnsResVipTableParams
}

func NewGslbCurEnhDnsResVipTableList() *GslbCurEnhDnsResVipTable {
	return &GslbCurEnhDnsResVipTable{}
}

func NewGslbCurEnhDnsResVipTable(
	gslbCurEnhDnsResVipIndex1 string,
	gslbCurEnhDnsResVipIndex2 string,
	params *GslbCurEnhDnsResVipTableParams,
) *GslbCurEnhDnsResVipTable {
	return &GslbCurEnhDnsResVipTable{
		GslbCurEnhDnsResVipIndex1: gslbCurEnhDnsResVipIndex1,
		GslbCurEnhDnsResVipIndex2: gslbCurEnhDnsResVipIndex2,
		Params:                    params,
	}
}

func (c *GslbCurEnhDnsResVipTable) Name() string {
	return "GslbCurEnhDnsResVipTable"
}

func (c *GslbCurEnhDnsResVipTable) GetParams() BeanType {
	return c.Params
}

func (c *GslbCurEnhDnsResVipTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *GslbCurEnhDnsResVipTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.GslbCurEnhDnsResVipIndex1).IsZero() &&
		reflect.ValueOf(c.GslbCurEnhDnsResVipIndex2).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.GslbCurEnhDnsResVipIndex1) + "/" + fmt.Sprint(c.GslbCurEnhDnsResVipIndex2)
}

type GslbCurEnhDnsResVipTableIPVer int32

const (
	GslbCurEnhDnsResVipTableIPVer_Ipv4        GslbCurEnhDnsResVipTableIPVer = 1
	GslbCurEnhDnsResVipTableIPVer_Ipv6        GslbCurEnhDnsResVipTableIPVer = 2
	GslbCurEnhDnsResVipTableIPVer_Unsupported GslbCurEnhDnsResVipTableIPVer = 2147483647
)

type GslbCurEnhDnsResVipTableStatus int32

const (
	GslbCurEnhDnsResVipTableStatus_Enable      GslbCurEnhDnsResVipTableStatus = 1
	GslbCurEnhDnsResVipTableStatus_Disable     GslbCurEnhDnsResVipTableStatus = 2
	GslbCurEnhDnsResVipTableStatus_Unsupported GslbCurEnhDnsResVipTableStatus = 2147483647
)

type GslbCurEnhDnsResVipTableRtsrcmac int32

const (
	GslbCurEnhDnsResVipTableRtsrcmac_Enable      GslbCurEnhDnsResVipTableRtsrcmac = 1
	GslbCurEnhDnsResVipTableRtsrcmac_Disable     GslbCurEnhDnsResVipTableRtsrcmac = 2
	GslbCurEnhDnsResVipTableRtsrcmac_Unsupported GslbCurEnhDnsResVipTableRtsrcmac = 2147483647
)

type GslbCurEnhDnsResVipTableParams struct {
	// DNS Responder VIP Index 1.
	Index1 string `json:"Index1,omitempty"`
	// DNS Responder VIP Index 2.
	Index2 string `json:"Index2,omitempty"`
	// DNS Responder VIP Name.
	Name string `json:"Name,omitempty"`
	// The type of IP address of DNS Responder VIP.
	IPVer GslbCurEnhDnsResVipTableIPVer `json:"IPVer,omitempty"`
	// IPV4 address of the DNS Responder VIP.
	V4 string `json:"V4,omitempty"`
	// IPV6 address of the DNS Responder VIP
	// Address should be 4-byte hexadecimal colon notation. Valid IPv6 address should be in
	// any of the following forms xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx or xxxx::xxxx:xxxx:xxxx:xxxx or ::xxxx
	V6 string `json:"V6,omitempty"`
	// Zone status - Enabled/Disabled.
	Status GslbCurEnhDnsResVipTableStatus `json:"Status,omitempty"`
	// Enabling/Disabling this will enable/disable rtsrcmac in both virtual servers that comprise this DNS Responder.
	Rtsrcmac GslbCurEnhDnsResVipTableRtsrcmac `json:"Rtsrcmac,omitempty"`
}

func (p GslbCurEnhDnsResVipTableParams) iMABean() {}
