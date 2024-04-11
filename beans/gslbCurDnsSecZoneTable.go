package beans

import (
	"fmt"
	"reflect"
)

// GslbCurDnsSecZoneTable Note:This mib is not supported for VX instance of Virtualization.
type GslbCurDnsSecZoneTable struct {
	// Zone ID.
	GslbCurDnsSecZoneID string
	Params              *GslbCurDnsSecZoneTableParams
}

func NewGslbCurDnsSecZoneTableList() *GslbCurDnsSecZoneTable {
	return &GslbCurDnsSecZoneTable{}
}

func NewGslbCurDnsSecZoneTable(
	gslbCurDnsSecZoneID string,
	params *GslbCurDnsSecZoneTableParams,
) *GslbCurDnsSecZoneTable {
	return &GslbCurDnsSecZoneTable{
		GslbCurDnsSecZoneID: gslbCurDnsSecZoneID,
		Params:              params,
	}
}

func (c *GslbCurDnsSecZoneTable) Name() string {
	return "GslbCurDnsSecZoneTable"
}

func (c *GslbCurDnsSecZoneTable) GetParams() BeanType {
	return c.Params
}

func (c *GslbCurDnsSecZoneTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *GslbCurDnsSecZoneTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.GslbCurDnsSecZoneID).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.GslbCurDnsSecZoneID)
}

type GslbCurDnsSecZoneTableStatus int32

const (
	GslbCurDnsSecZoneTableStatus_Enable      GslbCurDnsSecZoneTableStatus = 1
	GslbCurDnsSecZoneTableStatus_Disable     GslbCurDnsSecZoneTableStatus = 2
	GslbCurDnsSecZoneTableStatus_Unsupported GslbCurDnsSecZoneTableStatus = 2147483647
)

type GslbCurDnsSecZoneTableParentIPVer int32

const (
	GslbCurDnsSecZoneTableParentIPVer_Ipv4        GslbCurDnsSecZoneTableParentIPVer = 1
	GslbCurDnsSecZoneTableParentIPVer_Ipv6        GslbCurDnsSecZoneTableParentIPVer = 2
	GslbCurDnsSecZoneTableParentIPVer_Unsupported GslbCurDnsSecZoneTableParentIPVer = 2147483647
)

type GslbCurDnsSecZoneTableParams struct {
	// Zone ID.
	ID string `json:"ID,omitempty"`
	// Zone name (virt's dname).
	Name string `json:"Name,omitempty"`
	// KSK 1 of the zone.
	KSK1 string `json:"KSK1,omitempty"`
	// KSK 2 of the zone.
	KSK2 string `json:"KSK2,omitempty"`
	// KSK 3 of the zone.
	KSK3 string `json:"KSK3,omitempty"`
	// ZSK 1 of the zone.
	ZSK1 string `json:"ZSK1,omitempty"`
	// ZSK 2 of the zone.
	ZSK2 string `json:"ZSK2,omitempty"`
	// ZSK 1 of the zone.
	ZSK3 string `json:"ZSK3,omitempty"`
	// Zone status - Enabled/Disabled.
	Status GslbCurDnsSecZoneTableStatus `json:"Status,omitempty"`
	// The type of IP address.
	ParentIPVer GslbCurDnsSecZoneTableParentIPVer `json:"ParentIPVer,omitempty"`
	// Parent IP.
	ParentIPv4 string `json:"ParentIPv4,omitempty"`
	// IPV6 address of the Parent IP
	// Address should be 4-byte hexadecimal colon notation. Valid IPv6 address should be in
	// any of the following forms xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx or xxxx::xxxx:xxxx:xxxx:xxxx or ::xxxx
	ParentIPv6 string `json:"ParentIPv6,omitempty"`
	// Parent Hostname/IPv4 Address/IPv6 Address.
	ParentHost string `json:"ParentHost,omitempty"`
}

func (p GslbCurDnsSecZoneTableParams) iMABean() {}
