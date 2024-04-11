package beans

import (
	"fmt"
	"reflect"
)

// GslbNewCfgLocalDnsTable Gslb LLB Proxmity Local Dns table,
// Note:This mib is not supported for VX instance of Virtualization.
type GslbNewCfgLocalDnsTable struct {
	// The index for local dns.
	GslbNewCfgLocalDnsIndex int32
	Params                  *GslbNewCfgLocalDnsTableParams
}

func NewGslbNewCfgLocalDnsTableList() *GslbNewCfgLocalDnsTable {
	return &GslbNewCfgLocalDnsTable{}
}

func NewGslbNewCfgLocalDnsTable(
	gslbNewCfgLocalDnsIndex int32,
	params *GslbNewCfgLocalDnsTableParams,
) *GslbNewCfgLocalDnsTable {
	return &GslbNewCfgLocalDnsTable{
		GslbNewCfgLocalDnsIndex: gslbNewCfgLocalDnsIndex,
		Params:                  params,
	}
}

func (c *GslbNewCfgLocalDnsTable) Name() string {
	return "GslbNewCfgLocalDnsTable"
}

func (c *GslbNewCfgLocalDnsTable) GetParams() BeanType {
	return c.Params
}

func (c *GslbNewCfgLocalDnsTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *GslbNewCfgLocalDnsTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.GslbNewCfgLocalDnsIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.GslbNewCfgLocalDnsIndex)
}

type GslbNewCfgLocalDnsTableDelete int32

const (
	GslbNewCfgLocalDnsTableDelete_Other       GslbNewCfgLocalDnsTableDelete = 1
	GslbNewCfgLocalDnsTableDelete_Delete      GslbNewCfgLocalDnsTableDelete = 2
	GslbNewCfgLocalDnsTableDelete_Unsupported GslbNewCfgLocalDnsTableDelete = 2147483647
)

type GslbNewCfgLocalDnsTableVersion int32

const (
	GslbNewCfgLocalDnsTableVersion_Unknown     GslbNewCfgLocalDnsTableVersion = 0
	GslbNewCfgLocalDnsTableVersion_Ipv4        GslbNewCfgLocalDnsTableVersion = 4
	GslbNewCfgLocalDnsTableVersion_Ipv6        GslbNewCfgLocalDnsTableVersion = 6
	GslbNewCfgLocalDnsTableVersion_Unsupported GslbNewCfgLocalDnsTableVersion = 2147483647
)

type GslbNewCfgLocalDnsTableParams struct {
	// The index for local dns.
	Index int32 `json:"Index,omitempty"`
	// local DNS server IPv4 address .
	Addr string `json:"Addr,omitempty"`
	// local DNS server IPv6 address.
	// Address should be 16-byte hexadecimal colon notation.
	// Valid IPv6 address should be in any of the following forms
	// xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx or
	// xxxx::xxxx:xxxx:xxxx:xxxx or ::xxxx
	// Only new entry can be created,cannot modify existing ip.
	Dnsv6Addr string `json:"Dnsv6Addr,omitempty"`
	// By setting the value to delete(2), the entire row is
	// deleted.
	Delete GslbNewCfgLocalDnsTableDelete `json:"Delete,omitempty"`
	// The IP version of the local dns server.
	Version GslbNewCfgLocalDnsTableVersion `json:"Version,omitempty"`
}

func (p GslbNewCfgLocalDnsTableParams) iMABean() {}
