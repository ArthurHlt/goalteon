package beans

import (
	"fmt"
	"reflect"
)

// GslbNewDnsSecZoneTable Note:This mib is not supported for VX instance of Virtualization.
type GslbNewDnsSecZoneTable struct {
	// Zone ID.
	GslbNewDnsSecZoneID string
	Params              *GslbNewDnsSecZoneTableParams
}

func NewGslbNewDnsSecZoneTableList() *GslbNewDnsSecZoneTable {
	return &GslbNewDnsSecZoneTable{}
}

func NewGslbNewDnsSecZoneTable(
	gslbNewDnsSecZoneID string,
	params *GslbNewDnsSecZoneTableParams,
) *GslbNewDnsSecZoneTable {
	return &GslbNewDnsSecZoneTable{
		GslbNewDnsSecZoneID: gslbNewDnsSecZoneID,
		Params:              params,
	}
}

func (c *GslbNewDnsSecZoneTable) Name() string {
	return "GslbNewDnsSecZoneTable"
}

func (c *GslbNewDnsSecZoneTable) GetParams() BeanType {
	return c.Params
}

func (c *GslbNewDnsSecZoneTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *GslbNewDnsSecZoneTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.GslbNewDnsSecZoneID).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.GslbNewDnsSecZoneID)
}

type GslbNewDnsSecZoneTableStatus int32

const (
	GslbNewDnsSecZoneTableStatus_Enable      GslbNewDnsSecZoneTableStatus = 1
	GslbNewDnsSecZoneTableStatus_Disable     GslbNewDnsSecZoneTableStatus = 2
	GslbNewDnsSecZoneTableStatus_Unsupported GslbNewDnsSecZoneTableStatus = 2147483647
)

type GslbNewDnsSecZoneTableParentIPVer int32

const (
	GslbNewDnsSecZoneTableParentIPVer_Ipv4        GslbNewDnsSecZoneTableParentIPVer = 1
	GslbNewDnsSecZoneTableParentIPVer_Ipv6        GslbNewDnsSecZoneTableParentIPVer = 2
	GslbNewDnsSecZoneTableParentIPVer_Unsupported GslbNewDnsSecZoneTableParentIPVer = 2147483647
)

type GslbNewDnsSecZoneTableDelete int32

const (
	GslbNewDnsSecZoneTableDelete_Other       GslbNewDnsSecZoneTableDelete = 1
	GslbNewDnsSecZoneTableDelete_Delete      GslbNewDnsSecZoneTableDelete = 2
	GslbNewDnsSecZoneTableDelete_Unsupported GslbNewDnsSecZoneTableDelete = 2147483647
)

type GslbNewDnsSecZoneTableParams struct {
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
	// ZSK 3 of the zone.
	ZSK3 string `json:"ZSK3,omitempty"`
	// Zone status - Enabled/Disabled.
	Status GslbNewDnsSecZoneTableStatus `json:"Status,omitempty"`
	// The type of IP address.
	ParentIPVer GslbNewDnsSecZoneTableParentIPVer `json:"ParentIPVer,omitempty"`
	// IPV4 address of the Parent IP.
	ParentIPv4 string `json:"ParentIPv4,omitempty"`
	// IPV6 address of the Parent IP
	// Address should be 4-byte hexadecimal colon notation. Valid IPv6 address should be in
	// any of the following forms xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx or xxxx::xxxx:xxxx:xxxx:xxxx or ::xxxx
	ParentIPv6 string `json:"ParentIPv6,omitempty"`
	// By setting the value to delete(2), the entire row is deleted.
	Delete GslbNewDnsSecZoneTableDelete `json:"Delete,omitempty"`
	// Parent Hostname/IPv4 Address/IPv6 Address.
	ParentHost string `json:"ParentHost,omitempty"`
}

func (p GslbNewDnsSecZoneTableParams) iMABean() {}
