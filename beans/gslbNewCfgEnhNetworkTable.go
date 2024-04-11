package beans

import (
	"fmt"
	"reflect"
)

// GslbNewCfgEnhNetworkTable The table of network preference in the new_config.
// Note:This mib is not supported for VX instance of Virtualization.
type GslbNewCfgEnhNetworkTable struct {
	// The new network preference table index.
	GslbNewCfgEnhNetworkIndx int32
	Params                   *GslbNewCfgEnhNetworkTableParams
}

func NewGslbNewCfgEnhNetworkTableList() *GslbNewCfgEnhNetworkTable {
	return &GslbNewCfgEnhNetworkTable{}
}

func NewGslbNewCfgEnhNetworkTable(
	gslbNewCfgEnhNetworkIndx int32,
	params *GslbNewCfgEnhNetworkTableParams,
) *GslbNewCfgEnhNetworkTable {
	return &GslbNewCfgEnhNetworkTable{
		GslbNewCfgEnhNetworkIndx: gslbNewCfgEnhNetworkIndx,
		Params:                   params,
	}
}

func (c *GslbNewCfgEnhNetworkTable) Name() string {
	return "GslbNewCfgEnhNetworkTable"
}

func (c *GslbNewCfgEnhNetworkTable) GetParams() BeanType {
	return c.Params
}

func (c *GslbNewCfgEnhNetworkTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *GslbNewCfgEnhNetworkTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.GslbNewCfgEnhNetworkIndx).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.GslbNewCfgEnhNetworkIndx)
}

type GslbNewCfgEnhNetworkTableState int32

const (
	GslbNewCfgEnhNetworkTableState_Enabled     GslbNewCfgEnhNetworkTableState = 1
	GslbNewCfgEnhNetworkTableState_Disabled    GslbNewCfgEnhNetworkTableState = 2
	GslbNewCfgEnhNetworkTableState_Unsupported GslbNewCfgEnhNetworkTableState = 2147483647
)

type GslbNewCfgEnhNetworkTableDelete int32

const (
	GslbNewCfgEnhNetworkTableDelete_Other       GslbNewCfgEnhNetworkTableDelete = 1
	GslbNewCfgEnhNetworkTableDelete_Delete      GslbNewCfgEnhNetworkTableDelete = 2
	GslbNewCfgEnhNetworkTableDelete_Unsupported GslbNewCfgEnhNetworkTableDelete = 2147483647
)

type GslbNewCfgEnhNetworkTableVer int32

const (
	GslbNewCfgEnhNetworkTableVer_Ipv4        GslbNewCfgEnhNetworkTableVer = 1
	GslbNewCfgEnhNetworkTableVer_Ipv6        GslbNewCfgEnhNetworkTableVer = 2
	GslbNewCfgEnhNetworkTableVer_Unsupported GslbNewCfgEnhNetworkTableVer = 2147483647
)

type GslbNewCfgEnhNetworkTableClientAddrSrc int32

const (
	GslbNewCfgEnhNetworkTableClientAddrSrc_Ldns        GslbNewCfgEnhNetworkTableClientAddrSrc = 1
	GslbNewCfgEnhNetworkTableClientAddrSrc_Ecs         GslbNewCfgEnhNetworkTableClientAddrSrc = 2
	GslbNewCfgEnhNetworkTableClientAddrSrc_Unsupported GslbNewCfgEnhNetworkTableClientAddrSrc = 2147483647
)

type GslbNewCfgEnhNetworkTableServType int32

const (
	GslbNewCfgEnhNetworkTableServType_Group       GslbNewCfgEnhNetworkTableServType = 0
	GslbNewCfgEnhNetworkTableServType_Server      GslbNewCfgEnhNetworkTableServType = 1
	GslbNewCfgEnhNetworkTableServType_Unsupported GslbNewCfgEnhNetworkTableServType = 2147483647
)

type GslbNewCfgEnhNetworkTableSrcAddrType int32

const (
	GslbNewCfgEnhNetworkTableSrcAddrType_Address     GslbNewCfgEnhNetworkTableSrcAddrType = 1
	GslbNewCfgEnhNetworkTableSrcAddrType_Network     GslbNewCfgEnhNetworkTableSrcAddrType = 2
	GslbNewCfgEnhNetworkTableSrcAddrType_Unsupported GslbNewCfgEnhNetworkTableSrcAddrType = 2147483647
)

type GslbNewCfgEnhNetworkTableParams struct {
	// The new network preference table index.
	Indx int32 `json:"Indx,omitempty"`
	// Enable/Disable Global SLB for the network in
	// the new_configuration block.
	State GslbNewCfgEnhNetworkTableState `json:"State,omitempty"`
	// The Source IP address of the network table.
	SourceIp string `json:"SourceIp,omitempty"`
	// The Net mask of the network preference table.
	NetMask string `json:"NetMask,omitempty"`
	// By setting the value to delete(2), the entire row is deleted.
	Delete GslbNewCfgEnhNetworkTableDelete `json:"Delete,omitempty"`
	// The virtual server to be added to the network. When read, 0 is returned.
	AddVirtServer int32 `json:"AddVirtServer,omitempty"`
	// The virtual server to be removed from the network. When read, 0 is returned.
	RemoveVirtServer int32 `json:"RemoveVirtServer,omitempty"`
	// The remote real server to be added to the network. When read, 0 is returned.
	AddRemRealServer int32 `json:"AddRemRealServer,omitempty"`
	// The remote real server to be removed from the network. When read, 0 is returned.
	RemoveRemRealServer int32 `json:"RemoveRemRealServer,omitempty"`
	// The Source IPV6 address of the network table
	// Address should be 4-byte hexadecimal colon notation. Valid IPv6 address should be in
	// any of the following forms xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx or xxxx::xxxx:xxxx:xxxx:xxxx or ::xxxx
	SourceIpV6 string `json:"SourceIpV6,omitempty"`
	// The secondary IP address version of the remote site.
	Ver GslbNewCfgEnhNetworkTableVer `json:"Ver,omitempty"`
	// The prefix length associated with the source IPV6 address of the network table.
	Sprefix uint32 `json:"Sprefix,omitempty"`
	// The remote real server in alpha numeric to be added to the network.
	AddRealServerAlphaNum string `json:"AddRealServerAlphaNum,omitempty"`
	// The remote real server to be removed from the network.
	RemRealServerAlphaNum string `json:"RemRealServerAlphaNum,omitempty"`
	// The remote virtual server in alpha numeric to be added to the network.
	AddEnhVirtServer string `json:"AddEnhVirtServer,omitempty"`
	// The remote virtual server to be removed from the network.
	RemoveEnhVirtServer string `json:"RemoveEnhVirtServer,omitempty"`
	// The gslb network sip value as any or network class ID to be added to the network.
	ClassId string `json:"ClassId,omitempty"`
	// Client address source configuration either LDNS or ECS to be added to the network.
	ClientAddrSrc GslbNewCfgEnhNetworkTableClientAddrSrc `json:"ClientAddrSrc,omitempty"`
	// Global SLB local server type for Smart nat in the new_configuration block.
	ServType GslbNewCfgEnhNetworkTableServType `json:"ServType,omitempty"`
	// The Local Server IP address.IPV4/V6
	ServIp string `json:"ServIp,omitempty"`
	// The gslb network type either address or network to be added to the network.
	SrcAddrType GslbNewCfgEnhNetworkTableSrcAddrType `json:"SrcAddrType,omitempty"`
	// The Local Server WAN Group (optional), applicable only for servtype server.
	WanGrp string `json:"WanGrp,omitempty"`
	// The client network descriptive name.
	Name string `json:"Name,omitempty"`
	// The network class alpha numeric index to be added to the network.
	AddNetworkClass string `json:"AddNetworkClass,omitempty"`
	// The network class index to be removed from the network.
	RemoveNetworkClass string `json:"RemoveNetworkClass,omitempty"`
}

func (p GslbNewCfgEnhNetworkTableParams) iMABean() {}
