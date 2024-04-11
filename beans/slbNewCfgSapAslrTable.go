package beans

import (
	"fmt"
	"reflect"
)

// SlbNewCfgSapAslrTable The SAP ASLR connection management  table.
// Note:This mib is not supported on VX instance of Virtualization.
type SlbNewCfgSapAslrTable struct {
	// The Sap Aslr number
	SlbNewCfgSapAslrIndex int32
	Params                *SlbNewCfgSapAslrTableParams
}

func NewSlbNewCfgSapAslrTableList() *SlbNewCfgSapAslrTable {
	return &SlbNewCfgSapAslrTable{}
}

func NewSlbNewCfgSapAslrTable(
	slbNewCfgSapAslrIndex int32,
	params *SlbNewCfgSapAslrTableParams,
) *SlbNewCfgSapAslrTable {
	return &SlbNewCfgSapAslrTable{
		SlbNewCfgSapAslrIndex: slbNewCfgSapAslrIndex,
		Params:                params,
	}
}

func (c *SlbNewCfgSapAslrTable) Name() string {
	return "SlbNewCfgSapAslrTable"
}

func (c *SlbNewCfgSapAslrTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbNewCfgSapAslrTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbNewCfgSapAslrTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbNewCfgSapAslrIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbNewCfgSapAslrIndex)
}

type SlbNewCfgSapAslrTableIpVer int32

const (
	SlbNewCfgSapAslrTableIpVer_Ipv4        SlbNewCfgSapAslrTableIpVer = 1
	SlbNewCfgSapAslrTableIpVer_Ipv6        SlbNewCfgSapAslrTableIpVer = 2
	SlbNewCfgSapAslrTableIpVer_Unsupported SlbNewCfgSapAslrTableIpVer = 2147483647
)

type SlbNewCfgSapAslrTableVipIpVer int32

const (
	SlbNewCfgSapAslrTableVipIpVer_Ipv4        SlbNewCfgSapAslrTableVipIpVer = 1
	SlbNewCfgSapAslrTableVipIpVer_Ipv6        SlbNewCfgSapAslrTableVipIpVer = 2
	SlbNewCfgSapAslrTableVipIpVer_Unsupported SlbNewCfgSapAslrTableVipIpVer = 2147483647
)

type SlbNewCfgSapAslrTableState int32

const (
	SlbNewCfgSapAslrTableState_Enabled     SlbNewCfgSapAslrTableState = 1
	SlbNewCfgSapAslrTableState_Disabled    SlbNewCfgSapAslrTableState = 2
	SlbNewCfgSapAslrTableState_Unsupported SlbNewCfgSapAslrTableState = 2147483647
)

type SlbNewCfgSapAslrTableAutoConfig int32

const (
	SlbNewCfgSapAslrTableAutoConfig_Basic       SlbNewCfgSapAslrTableAutoConfig = 1
	SlbNewCfgSapAslrTableAutoConfig_Full        SlbNewCfgSapAslrTableAutoConfig = 2
	SlbNewCfgSapAslrTableAutoConfig_Unsupported SlbNewCfgSapAslrTableAutoConfig = 2147483647
)

type SlbNewCfgSapAslrTableDelete int32

const (
	SlbNewCfgSapAslrTableDelete_Other       SlbNewCfgSapAslrTableDelete = 1
	SlbNewCfgSapAslrTableDelete_Delete      SlbNewCfgSapAslrTableDelete = 2
	SlbNewCfgSapAslrTableDelete_Unsupported SlbNewCfgSapAslrTableDelete = 2147483647
)

type SlbNewCfgSapAslrTableParams struct {
	// The Sap Aslr number
	Index int32 `json:"Index,omitempty"`
	// IP address of the Sap Aslr identified by the instance of the
	// slbNewCfgSapAslrIndex.
	IpAddr string `json:"IpAddr,omitempty"`
	// IPV6 address of the Sap Aslr manager identified by the instance of the slbNewCfgSapAslrIndex
	// Address should be 4-byte hexadecimal colon notation. Valid IPv6 address should be in
	// any of the following forms xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx or xxxx::xxxx:xxxx:xxxx:xxxx or ::xxxx
	Ipv6Addr string `json:"Ipv6Addr,omitempty"`
	// The type of IP address.
	IpVer SlbNewCfgSapAslrTableIpVer `json:"IpVer,omitempty"`
	// The Sap Aslr service port number.
	PortNum uint64 `json:"PortNum,omitempty"`
	// IP address of the VIP identified by the instance of the
	// slbNewCfgSapAslrIndex.
	VipAddr string `json:"VipAddr,omitempty"`
	// IPV6 address of the VIP identified by the instance of the slbNewCfgSapAslrIndex
	// Address should be 4-byte hexadecimal colon notation. Valid IPv6 address should be in
	// any of the following forms xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx or xxxx::xxxx:xxxx:xxxx:xxxx or ::xxxx
	Vipv6Addr string `json:"Vipv6Addr,omitempty"`
	// The type of VIP IP address.
	VipIpVer SlbNewCfgSapAslrTableVipIpVer `json:"VipIpVer,omitempty"`
	// Set the Sap Aslr server inspection interval.
	Interval uint64 `json:"Interval,omitempty"`
	// Enable or disable the Sap Aslr server sampling.
	State SlbNewCfgSapAslrTableState `json:"State,omitempty"`
	// The name of the Sap Aslr connection.
	Name string `json:"Name,omitempty"`
	// The Sap Aslr last activity time stamp.
	LastAct string `json:"LastAct,omitempty"`
	// Defines the configuration update process.
	// When set the value to basic (1), the configuration is limited,
	// 	 comparing to full (2)
	AutoConfig SlbNewCfgSapAslrTableAutoConfig `json:"AutoConfig,omitempty"`
	// Set the Sap Aslr session information.
	// 	When read, zero is returned.
	SessionInfo int32 `json:"SessionInfo,omitempty"`
	// When set to the value of 2 (delete), the entire row is deleted.
	// When read, other(1) is returned. Setting the value to anything
	// other than 2(delete) has no effect on the state of the row.
	Delete SlbNewCfgSapAslrTableDelete `json:"Delete,omitempty"`
	// The Sap Aslr connection signature.
	Signature string `json:"Signature,omitempty"`
	// Server Certificate name associated with this virtual service.
	ServCert string `json:"ServCert,omitempty"`
}

func (p SlbNewCfgSapAslrTableParams) iMABean() {}
