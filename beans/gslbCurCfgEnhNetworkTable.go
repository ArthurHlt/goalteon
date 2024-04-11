package beans

import (
	"fmt"
	"reflect"
)

// GslbCurCfgEnhNetworkTable The table of network preference in the current_config.
// Note:This mib is not supported for VX instance of Virtualization.
type GslbCurCfgEnhNetworkTable struct {
	// The current network preference table index.
	GslbCurCfgEnhNetworkIndx int32
	Params                   *GslbCurCfgEnhNetworkTableParams
}

func NewGslbCurCfgEnhNetworkTableList() *GslbCurCfgEnhNetworkTable {
	return &GslbCurCfgEnhNetworkTable{}
}

func NewGslbCurCfgEnhNetworkTable(
	gslbCurCfgEnhNetworkIndx int32,
	params *GslbCurCfgEnhNetworkTableParams,
) *GslbCurCfgEnhNetworkTable {
	return &GslbCurCfgEnhNetworkTable{
		GslbCurCfgEnhNetworkIndx: gslbCurCfgEnhNetworkIndx,
		Params:                   params,
	}
}

func (c *GslbCurCfgEnhNetworkTable) Name() string {
	return "GslbCurCfgEnhNetworkTable"
}

func (c *GslbCurCfgEnhNetworkTable) GetParams() BeanType {
	return c.Params
}

func (c *GslbCurCfgEnhNetworkTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *GslbCurCfgEnhNetworkTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.GslbCurCfgEnhNetworkIndx).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.GslbCurCfgEnhNetworkIndx)
}

type GslbCurCfgEnhNetworkTableState int32

const (
	GslbCurCfgEnhNetworkTableState_Enabled     GslbCurCfgEnhNetworkTableState = 1
	GslbCurCfgEnhNetworkTableState_Disabled    GslbCurCfgEnhNetworkTableState = 2
	GslbCurCfgEnhNetworkTableState_Unsupported GslbCurCfgEnhNetworkTableState = 2147483647
)

type GslbCurCfgEnhNetworkTableVer int32

const (
	GslbCurCfgEnhNetworkTableVer_Ipv4        GslbCurCfgEnhNetworkTableVer = 1
	GslbCurCfgEnhNetworkTableVer_Ipv6        GslbCurCfgEnhNetworkTableVer = 2
	GslbCurCfgEnhNetworkTableVer_Unsupported GslbCurCfgEnhNetworkTableVer = 2147483647
)

type GslbCurCfgEnhNetworkTableClientAddrSrc int32

const (
	GslbCurCfgEnhNetworkTableClientAddrSrc_Ldns        GslbCurCfgEnhNetworkTableClientAddrSrc = 1
	GslbCurCfgEnhNetworkTableClientAddrSrc_Ecs         GslbCurCfgEnhNetworkTableClientAddrSrc = 2
	GslbCurCfgEnhNetworkTableClientAddrSrc_Unsupported GslbCurCfgEnhNetworkTableClientAddrSrc = 2147483647
)

type GslbCurCfgEnhNetworkTableServType int32

const (
	GslbCurCfgEnhNetworkTableServType_Group       GslbCurCfgEnhNetworkTableServType = 0
	GslbCurCfgEnhNetworkTableServType_Server      GslbCurCfgEnhNetworkTableServType = 1
	GslbCurCfgEnhNetworkTableServType_Unsupported GslbCurCfgEnhNetworkTableServType = 2147483647
)

type GslbCurCfgEnhNetworkTableSrcAddrType int32

const (
	GslbCurCfgEnhNetworkTableSrcAddrType_Address     GslbCurCfgEnhNetworkTableSrcAddrType = 1
	GslbCurCfgEnhNetworkTableSrcAddrType_Network     GslbCurCfgEnhNetworkTableSrcAddrType = 2
	GslbCurCfgEnhNetworkTableSrcAddrType_Unsupported GslbCurCfgEnhNetworkTableSrcAddrType = 2147483647
)

type GslbCurCfgEnhNetworkTableParams struct {
	// The current network preference table index.
	Indx int32 `json:"Indx,omitempty"`
	// Enable/Disable Global SLB for the network in the
	// current_configuration block.
	State GslbCurCfgEnhNetworkTableState `json:"State,omitempty"`
	// The Source IP address of the network preference table.
	SourceIp string `json:"SourceIp,omitempty"`
	// The Net mask of the network preference table.
	NetMask string `json:"NetMask,omitempty"`
	// The Source IPV6 address of the network table
	// Address should be 4-byte hexadecimal colon notation. Valid IPv6 address should be in
	// any of the following forms xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx or xxxx::xxxx:xxxx:xxxx:xxxx or ::xxxx
	SourceIpV6 string `json:"SourceIpV6,omitempty"`
	// The secondary IP address version of the remote site.
	Ver GslbCurCfgEnhNetworkTableVer `json:"Ver,omitempty"`
	// The prefix length associated with the source IPV6 address of the network table.
	Sprefix uint32 `json:"Sprefix,omitempty"`
	// This is for gslb network sip value as any or network class ID.
	ClassId string `json:"ClassId,omitempty"`
	// Client address source either LDNS or ECS.
	ClientAddrSrc GslbCurCfgEnhNetworkTableClientAddrSrc `json:"ClientAddrSrc,omitempty"`
	// Global SLB local server type for Smart nat in the new_configuration block.
	ServType GslbCurCfgEnhNetworkTableServType `json:"ServType,omitempty"`
	// The Local Server IP address.IPV4/V6
	ServIp string `json:"ServIp,omitempty"`
	// Global SLB network type fron any one as address or network.
	SrcAddrType GslbCurCfgEnhNetworkTableSrcAddrType `json:"SrcAddrType,omitempty"`
	// The Local Server WAN Group (optional), applicable only for servtype server.
	WanGrp string `json:"WanGrp,omitempty"`
	// The client network descriptive name.
	Name string `json:"Name,omitempty"`
}

func (p GslbCurCfgEnhNetworkTableParams) iMABean() {}
