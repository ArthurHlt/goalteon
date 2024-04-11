package beans

import (
	"fmt"
	"reflect"
)

// SlbCurCfgSapAslrTable The SAP ASLR connection management  table.
// Note:This mib is not supported on VX instance of Virtualization.
type SlbCurCfgSapAslrTable struct {
	// The Sap Aslr number
	SlbCurCfgSapAslrIndex int32
	Params                *SlbCurCfgSapAslrTableParams
}

func NewSlbCurCfgSapAslrTableList() *SlbCurCfgSapAslrTable {
	return &SlbCurCfgSapAslrTable{}
}

func NewSlbCurCfgSapAslrTable(
	slbCurCfgSapAslrIndex int32,
	params *SlbCurCfgSapAslrTableParams,
) *SlbCurCfgSapAslrTable {
	return &SlbCurCfgSapAslrTable{
		SlbCurCfgSapAslrIndex: slbCurCfgSapAslrIndex,
		Params:                params,
	}
}

func (c *SlbCurCfgSapAslrTable) Name() string {
	return "SlbCurCfgSapAslrTable"
}

func (c *SlbCurCfgSapAslrTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbCurCfgSapAslrTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbCurCfgSapAslrTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbCurCfgSapAslrIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbCurCfgSapAslrIndex)
}

type SlbCurCfgSapAslrTableIpVer int32

const (
	SlbCurCfgSapAslrTableIpVer_Ipv4        SlbCurCfgSapAslrTableIpVer = 1
	SlbCurCfgSapAslrTableIpVer_Ipv6        SlbCurCfgSapAslrTableIpVer = 2
	SlbCurCfgSapAslrTableIpVer_Unsupported SlbCurCfgSapAslrTableIpVer = 2147483647
)

type SlbCurCfgSapAslrTableVipIpVer int32

const (
	SlbCurCfgSapAslrTableVipIpVer_Ipv4        SlbCurCfgSapAslrTableVipIpVer = 1
	SlbCurCfgSapAslrTableVipIpVer_Ipv6        SlbCurCfgSapAslrTableVipIpVer = 2
	SlbCurCfgSapAslrTableVipIpVer_Unsupported SlbCurCfgSapAslrTableVipIpVer = 2147483647
)

type SlbCurCfgSapAslrTableState int32

const (
	SlbCurCfgSapAslrTableState_Enabled     SlbCurCfgSapAslrTableState = 1
	SlbCurCfgSapAslrTableState_Disabled    SlbCurCfgSapAslrTableState = 2
	SlbCurCfgSapAslrTableState_Unsupported SlbCurCfgSapAslrTableState = 2147483647
)

type SlbCurCfgSapAslrTableAutoConfig int32

const (
	SlbCurCfgSapAslrTableAutoConfig_Basic       SlbCurCfgSapAslrTableAutoConfig = 1
	SlbCurCfgSapAslrTableAutoConfig_Full        SlbCurCfgSapAslrTableAutoConfig = 2
	SlbCurCfgSapAslrTableAutoConfig_Unsupported SlbCurCfgSapAslrTableAutoConfig = 2147483647
)

type SlbCurCfgSapAslrTableDelete int32

const (
	SlbCurCfgSapAslrTableDelete_Other       SlbCurCfgSapAslrTableDelete = 1
	SlbCurCfgSapAslrTableDelete_Delete      SlbCurCfgSapAslrTableDelete = 2
	SlbCurCfgSapAslrTableDelete_Unsupported SlbCurCfgSapAslrTableDelete = 2147483647
)

type SlbCurCfgSapAslrTableParams struct {
	// The Sap Aslr number
	Index int32 `json:"Index,omitempty"`
	// IP address of the Sap Aslr identified by the instance of the
	// slbCurCfgSapAslrIndex.
	IpAddr string `json:"IpAddr,omitempty"`
	// IPV6 address of the Sap Aslr manager identified by the instance of the slbCurCfgSapAslrIndex
	// Address should be 4-byte hexadecimal colon notation. Valid IPv6 address should be in
	// any of the following forms xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx or xxxx::xxxx:xxxx:xxxx:xxxx or ::xxxx
	Ipv6Addr string `json:"Ipv6Addr,omitempty"`
	// The type of IP address.
	IpVer SlbCurCfgSapAslrTableIpVer `json:"IpVer,omitempty"`
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
	VipIpVer SlbCurCfgSapAslrTableVipIpVer `json:"VipIpVer,omitempty"`
	// Set the Sap Aslr server inspection interval.
	Interval uint64 `json:"Interval,omitempty"`
	// Enable or disable the Sap Aslr server sampling.
	State SlbCurCfgSapAslrTableState `json:"State,omitempty"`
	// The name of the Sap Aslr connection.
	Name string `json:"Name,omitempty"`
	// The Sap Aslr connection last activity time stamp.
	LastAct string `json:"LastAct,omitempty"`
	// Defines the configuration update process.
	// When set the value to basic (1), the configuration is limited,
	// 	 comparing to full (2)
	AutoConfig SlbCurCfgSapAslrTableAutoConfig `json:"AutoConfig,omitempty"`
	// Set the Sap Aslr session information.
	// 	When read, zero is returned.
	SessionInfo int32 `json:"SessionInfo,omitempty"`
	// When set to the value of 2 (delete), the entire row is deleted.
	// When read, other(1) is returned. Setting the value to anything
	// other than 2(delete) has no effect on the state of the row.
	Delete SlbCurCfgSapAslrTableDelete `json:"Delete,omitempty"`
	// The Sap Aslr connection signature.
	Signature string `json:"Signature,omitempty"`
	// Server Certificate name associated with this virtual service.
	ServCert string `json:"ServCert,omitempty"`
}

func (p SlbCurCfgSapAslrTableParams) iMABean() {}
