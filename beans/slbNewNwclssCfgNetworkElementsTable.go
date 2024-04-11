package beans

import (
	"fmt"
	"reflect"
)

// SlbNewNwclssCfgNetworkElementsTable New network elements table.
// Note:This mib is not supported for VX instance of Virtualization.
type SlbNewNwclssCfgNetworkElementsTable struct {
	// ID of network class which owns network element.
	SlbNewNwclssCfgNetworkElementsNcId string
	// Network class ID.
	SlbNewNwclssCfgNetworkElementsId string
	Params                           *SlbNewNwclssCfgNetworkElementsTableParams
}

func NewSlbNewNwclssCfgNetworkElementsTableList() *SlbNewNwclssCfgNetworkElementsTable {
	return &SlbNewNwclssCfgNetworkElementsTable{}
}

func NewSlbNewNwclssCfgNetworkElementsTable(
	slbNewNwclssCfgNetworkElementsNcId string,
	slbNewNwclssCfgNetworkElementsId string,
	params *SlbNewNwclssCfgNetworkElementsTableParams,
) *SlbNewNwclssCfgNetworkElementsTable {
	return &SlbNewNwclssCfgNetworkElementsTable{
		SlbNewNwclssCfgNetworkElementsNcId: slbNewNwclssCfgNetworkElementsNcId,
		SlbNewNwclssCfgNetworkElementsId:   slbNewNwclssCfgNetworkElementsId,
		Params:                             params,
	}
}

func (c *SlbNewNwclssCfgNetworkElementsTable) Name() string {
	return "SlbNewNwclssCfgNetworkElementsTable"
}

func (c *SlbNewNwclssCfgNetworkElementsTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbNewNwclssCfgNetworkElementsTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbNewNwclssCfgNetworkElementsTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbNewNwclssCfgNetworkElementsNcId).IsZero() &&
		reflect.ValueOf(c.SlbNewNwclssCfgNetworkElementsId).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbNewNwclssCfgNetworkElementsNcId) + "/" + fmt.Sprint(c.SlbNewNwclssCfgNetworkElementsId)
}

type SlbNewNwclssCfgNetworkElementsTableNetType int32

const (
	SlbNewNwclssCfgNetworkElementsTableNetType_Subnet      SlbNewNwclssCfgNetworkElementsTableNetType = 1
	SlbNewNwclssCfgNetworkElementsTableNetType_Range       SlbNewNwclssCfgNetworkElementsTableNetType = 2
	SlbNewNwclssCfgNetworkElementsTableNetType_Geo         SlbNewNwclssCfgNetworkElementsTableNetType = 3
	SlbNewNwclssCfgNetworkElementsTableNetType_Isp         SlbNewNwclssCfgNetworkElementsTableNetType = 4
	SlbNewNwclssCfgNetworkElementsTableNetType_Unsupported SlbNewNwclssCfgNetworkElementsTableNetType = 2147483647
)

type SlbNewNwclssCfgNetworkElementsTableMatchType int32

const (
	SlbNewNwclssCfgNetworkElementsTableMatchType_Include     SlbNewNwclssCfgNetworkElementsTableMatchType = 1
	SlbNewNwclssCfgNetworkElementsTableMatchType_Exclude     SlbNewNwclssCfgNetworkElementsTableMatchType = 2
	SlbNewNwclssCfgNetworkElementsTableMatchType_Unsupported SlbNewNwclssCfgNetworkElementsTableMatchType = 2147483647
)

type SlbNewNwclssCfgNetworkElementsTableDel int32

const (
	SlbNewNwclssCfgNetworkElementsTableDel_Other       SlbNewNwclssCfgNetworkElementsTableDel = 1
	SlbNewNwclssCfgNetworkElementsTableDel_Delete      SlbNewNwclssCfgNetworkElementsTableDel = 2
	SlbNewNwclssCfgNetworkElementsTableDel_Unsupported SlbNewNwclssCfgNetworkElementsTableDel = 2147483647
)

type SlbNewNwclssCfgNetworkElementsTableTypeAddrOrRegn int32

const (
	SlbNewNwclssCfgNetworkElementsTableTypeAddrOrRegn_Address     SlbNewNwclssCfgNetworkElementsTableTypeAddrOrRegn = 1
	SlbNewNwclssCfgNetworkElementsTableTypeAddrOrRegn_Region      SlbNewNwclssCfgNetworkElementsTableTypeAddrOrRegn = 2
	SlbNewNwclssCfgNetworkElementsTableTypeAddrOrRegn_Unsupported SlbNewNwclssCfgNetworkElementsTableTypeAddrOrRegn = 2147483647
)

type SlbNewNwclssCfgNetworkElementsTableParams struct {
	// ID of network class which owns network element.
	NcId string `json:"NcId,omitempty"`
	// Network class ID.
	Id string `json:"Id,omitempty"`
	// The type of network.
	NetType SlbNewNwclssCfgNetworkElementsTableNetType `json:"NetType,omitempty"`
	// IP address of network element.
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
	MatchType SlbNewNwclssCfgNetworkElementsTableMatchType `json:"MatchType,omitempty"`
	// By setting the value to delete(2), the entire row is deleted.
	Del SlbNewNwclssCfgNetworkElementsTableDel `json:"Del,omitempty"`
	// Continent name of the region for network elements.
	RegCont string `json:"RegCont,omitempty"`
	// Country/Geolocation name of the region for network elements.
	RegCountry string `json:"RegCountry,omitempty"`
	// State name of the region for network elements.
	RegState string `json:"RegState,omitempty"`
	// The network class type should be either 'address' or 'region' from network class.
	TypeAddrOrRegn SlbNewNwclssCfgNetworkElementsTableTypeAddrOrRegn `json:"TypeAddrOrRegn,omitempty"`
	// ISP name of the isp region type for network elements.
	IspName string `json:"IspName,omitempty"`
}

func (p SlbNewNwclssCfgNetworkElementsTableParams) iMABean() {}
