package beans

import (
	"fmt"
	"reflect"
)

// GslbCurCfgLocalDnsTable The table of proximity localdns in the current_config.
// Note:This mib is not supported for VX instance of Virtualization.
type GslbCurCfgLocalDnsTable struct {
	// The index for localdns.
	GslbCurCfgLocalDnsIndex int32
	Params                  *GslbCurCfgLocalDnsTableParams
}

func NewGslbCurCfgLocalDnsTableList() *GslbCurCfgLocalDnsTable {
	return &GslbCurCfgLocalDnsTable{}
}

func NewGslbCurCfgLocalDnsTable(
	gslbCurCfgLocalDnsIndex int32,
	params *GslbCurCfgLocalDnsTableParams,
) *GslbCurCfgLocalDnsTable {
	return &GslbCurCfgLocalDnsTable{
		GslbCurCfgLocalDnsIndex: gslbCurCfgLocalDnsIndex,
		Params:                  params,
	}
}

func (c *GslbCurCfgLocalDnsTable) Name() string {
	return "GslbCurCfgLocalDnsTable"
}

func (c *GslbCurCfgLocalDnsTable) GetParams() BeanType {
	return c.Params
}

func (c *GslbCurCfgLocalDnsTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *GslbCurCfgLocalDnsTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.GslbCurCfgLocalDnsIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.GslbCurCfgLocalDnsIndex)
}

type GslbCurCfgLocalDnsTableVersion int32

const (
	GslbCurCfgLocalDnsTableVersion_Unknown     GslbCurCfgLocalDnsTableVersion = 0
	GslbCurCfgLocalDnsTableVersion_Ipv4        GslbCurCfgLocalDnsTableVersion = 4
	GslbCurCfgLocalDnsTableVersion_Ipv6        GslbCurCfgLocalDnsTableVersion = 6
	GslbCurCfgLocalDnsTableVersion_Unsupported GslbCurCfgLocalDnsTableVersion = 2147483647
)

type GslbCurCfgLocalDnsTableParams struct {
	// The index for localdns.
	DnsIndex int32 `json:"DnsIndex,omitempty"`
	// local DNS server IPv4 address.
	Addr string `json:"Addr,omitempty"`
	// local DNS server IPv6 address.
	// Address should be 16-byte hexadecimal colon notation.
	// Valid IPv6 address should be in any of the following forms
	// xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx or
	// xxxx::xxxx:xxxx:xxxx:xxxx or ::xxxx
	Dnsv6Addr string `json:"Dnsv6Addr,omitempty"`
	// The IP version of the localdns.
	Version GslbCurCfgLocalDnsTableVersion `json:"Version,omitempty"`
}

func (p GslbCurCfgLocalDnsTableParams) iMABean() {}
