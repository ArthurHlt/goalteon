package beans

import (
	"fmt"
	"reflect"
)

// SlbCurNwclssCfgNetworkElementsTable Current network elements table.
// Note:This mib is not supported for VX instance of Virtualization.
type SlbCurNwclssCfgNetworkElementsTable struct {
	// Id of network class which owns network element.
	SlbCurNwclssCfgNetworkElementsNcId string
	// Network element id.
	SlbCurNwclssCfgNetworkElementsId string
	Params                           *SlbCurNwclssCfgNetworkElementsTableParams
}

func NewSlbCurNwclssCfgNetworkElementsTableList() *SlbCurNwclssCfgNetworkElementsTable {
	return &SlbCurNwclssCfgNetworkElementsTable{}
}

func NewSlbCurNwclssCfgNetworkElementsTable(
	slbCurNwclssCfgNetworkElementsNcId string,
	slbCurNwclssCfgNetworkElementsId string,
	params *SlbCurNwclssCfgNetworkElementsTableParams,
) *SlbCurNwclssCfgNetworkElementsTable {
	return &SlbCurNwclssCfgNetworkElementsTable{
		SlbCurNwclssCfgNetworkElementsNcId: slbCurNwclssCfgNetworkElementsNcId,
		SlbCurNwclssCfgNetworkElementsId:   slbCurNwclssCfgNetworkElementsId,
		Params:                             params,
	}
}

func (c *SlbCurNwclssCfgNetworkElementsTable) Name() string {
	return "SlbCurNwclssCfgNetworkElementsTable"
}

func (c *SlbCurNwclssCfgNetworkElementsTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbCurNwclssCfgNetworkElementsTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbCurNwclssCfgNetworkElementsTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbCurNwclssCfgNetworkElementsNcId).IsZero() &&
		reflect.ValueOf(c.SlbCurNwclssCfgNetworkElementsId).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbCurNwclssCfgNetworkElementsNcId) + "/" + fmt.Sprint(c.SlbCurNwclssCfgNetworkElementsId)
}

type SlbCurNwclssCfgNetworkElementsTableNetType int32

const (
	SlbCurNwclssCfgNetworkElementsTableNetType_Subnet      SlbCurNwclssCfgNetworkElementsTableNetType = 1
	SlbCurNwclssCfgNetworkElementsTableNetType_Range       SlbCurNwclssCfgNetworkElementsTableNetType = 2
	SlbCurNwclssCfgNetworkElementsTableNetType_Geo         SlbCurNwclssCfgNetworkElementsTableNetType = 3
	SlbCurNwclssCfgNetworkElementsTableNetType_Isp         SlbCurNwclssCfgNetworkElementsTableNetType = 4
	SlbCurNwclssCfgNetworkElementsTableNetType_Unsupported SlbCurNwclssCfgNetworkElementsTableNetType = 2147483647
)

type SlbCurNwclssCfgNetworkElementsTableMatchType int32

const (
	SlbCurNwclssCfgNetworkElementsTableMatchType_Include     SlbCurNwclssCfgNetworkElementsTableMatchType = 1
	SlbCurNwclssCfgNetworkElementsTableMatchType_Exclude     SlbCurNwclssCfgNetworkElementsTableMatchType = 2
	SlbCurNwclssCfgNetworkElementsTableMatchType_Unsupported SlbCurNwclssCfgNetworkElementsTableMatchType = 2147483647
)

type SlbCurNwclssCfgNetworkElementsTableTypeAddrOrRegn int32

const (
	SlbCurNwclssCfgNetworkElementsTableTypeAddrOrRegn_Address     SlbCurNwclssCfgNetworkElementsTableTypeAddrOrRegn = 1
	SlbCurNwclssCfgNetworkElementsTableTypeAddrOrRegn_Region      SlbCurNwclssCfgNetworkElementsTableTypeAddrOrRegn = 2
	SlbCurNwclssCfgNetworkElementsTableTypeAddrOrRegn_Unsupported SlbCurNwclssCfgNetworkElementsTableTypeAddrOrRegn = 2147483647
)

type SlbCurNwclssCfgNetworkElementsTableParams struct {
	// Id of network class which owns network element.
	NcId string `json:"NcId,omitempty"`
	// Network element id.
	Id string `json:"Id,omitempty"`
	// The type of network.
	NetType SlbCurNwclssCfgNetworkElementsTableNetType `json:"NetType,omitempty"`
	// IPv4 address of network element.
	Ip string `json:"Ip,omitempty"`
	// Subnet mask of network element.
	Mask string `json:"Mask,omitempty"`
	// From IP address of network element.
	FromIp string `json:"FromIp,omitempty"`
	// To IP address of network element.
	ToIp string `json:"ToIp,omitempty"`
	// The IPv6 address of network element.
	// 	 Address should be 4-byte hexadecimal colon notation.
	// Valid IPv6 address should be in any of the following forms
	// xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx or
	// xxxx::xxxx:xxxx:xxxx:xxxx or ::xxxx
	Ipv6Addr string `json:"Ipv6Addr,omitempty"`
	// The prefix length associated with the network IP address .
	PrefixLen uint32 `json:"PrefixLen,omitempty"`
	// The from IPv6 address of network element.
	// 	 Address should be 4-byte hexadecimal colon notation.
	// Valid IPv6 address should be in any of the following forms
	// xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx or
	// xxxx::xxxx:xxxx:xxxx:xxxx or ::xxxx
	FromIpv6Addr string `json:"FromIpv6Addr,omitempty"`
	// The to IPv6 address of network element.
	// 	 Address should be 4-byte hexadecimal colon notation.
	// Valid IPv6 address should be in any of the following forms
	// xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx or
	// xxxx::xxxx:xxxx:xxxx:xxxx or ::xxxx
	ToIpv6Addr string `json:"ToIpv6Addr,omitempty"`
	// The match type of network element.
	MatchType SlbCurNwclssCfgNetworkElementsTableMatchType `json:"MatchType,omitempty"`
	// Continent name of the region to be added for network element.
	RegCont string `json:"RegCont,omitempty"`
	// Country/Geolocation name of the region to be added for network element.
	RegCountry string `json:"RegCountry,omitempty"`
	// State name of the region to be added for network element.
	RegState string `json:"RegState,omitempty"`
	// The network class type should be either 'address' or 'region' from network class.
	TypeAddrOrRegn SlbCurNwclssCfgNetworkElementsTableTypeAddrOrRegn `json:"TypeAddrOrRegn,omitempty"`
	// ISP name of the isp region type for network elements.
	IspName string `json:"IspName,omitempty"`
}

func (p SlbCurNwclssCfgNetworkElementsTableParams) iMABean() {}
