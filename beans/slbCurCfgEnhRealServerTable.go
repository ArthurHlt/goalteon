package beans

import (
	"fmt"
	"reflect"
)

// SlbCurCfgEnhRealServerTable The table of Real Server configuration in the current_config.
// Note:This mib is not supported on VX instance of Virtualization.
type SlbCurCfgEnhRealServerTable struct {
	// The real server number
	SlbCurCfgEnhRealServerIndex string
	Params                      *SlbCurCfgEnhRealServerTableParams
}

func NewSlbCurCfgEnhRealServerTableList() *SlbCurCfgEnhRealServerTable {
	return &SlbCurCfgEnhRealServerTable{}
}

func NewSlbCurCfgEnhRealServerTable(
	slbCurCfgEnhRealServerIndex string,
	params *SlbCurCfgEnhRealServerTableParams,
) *SlbCurCfgEnhRealServerTable {
	return &SlbCurCfgEnhRealServerTable{
		SlbCurCfgEnhRealServerIndex: slbCurCfgEnhRealServerIndex,
		Params:                      params,
	}
}

func (c *SlbCurCfgEnhRealServerTable) Name() string {
	return "SlbCurCfgEnhRealServerTable"
}

func (c *SlbCurCfgEnhRealServerTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbCurCfgEnhRealServerTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbCurCfgEnhRealServerTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbCurCfgEnhRealServerIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbCurCfgEnhRealServerIndex)
}

type SlbCurCfgEnhRealServerTableState int32

const (
	SlbCurCfgEnhRealServerTableState_Enabled                    SlbCurCfgEnhRealServerTableState = 2
	SlbCurCfgEnhRealServerTableState_Disabled                   SlbCurCfgEnhRealServerTableState = 3
	SlbCurCfgEnhRealServerTableState_DisabledWithFastage        SlbCurCfgEnhRealServerTableState = 4
	SlbCurCfgEnhRealServerTableState_ShutdownConnection         SlbCurCfgEnhRealServerTableState = 5
	SlbCurCfgEnhRealServerTableState_ShutdownPersistentSessions SlbCurCfgEnhRealServerTableState = 6
	SlbCurCfgEnhRealServerTableState_Unsupported                SlbCurCfgEnhRealServerTableState = 2147483647
)

type SlbCurCfgEnhRealServerTableType int32

const (
	SlbCurCfgEnhRealServerTableType_LocalServer  SlbCurCfgEnhRealServerTableType = 1
	SlbCurCfgEnhRealServerTableType_RemoteServer SlbCurCfgEnhRealServerTableType = 2
	SlbCurCfgEnhRealServerTableType_Unsupported  SlbCurCfgEnhRealServerTableType = 2147483647
)

type SlbCurCfgEnhRealServerTableCookie int32

const (
	SlbCurCfgEnhRealServerTableCookie_Enabled     SlbCurCfgEnhRealServerTableCookie = 1
	SlbCurCfgEnhRealServerTableCookie_Disabled    SlbCurCfgEnhRealServerTableCookie = 2
	SlbCurCfgEnhRealServerTableCookie_Unsupported SlbCurCfgEnhRealServerTableCookie = 2147483647
)

type SlbCurCfgEnhRealServerTableExcludeStr int32

const (
	SlbCurCfgEnhRealServerTableExcludeStr_Enabled     SlbCurCfgEnhRealServerTableExcludeStr = 1
	SlbCurCfgEnhRealServerTableExcludeStr_Disabled    SlbCurCfgEnhRealServerTableExcludeStr = 2
	SlbCurCfgEnhRealServerTableExcludeStr_Unsupported SlbCurCfgEnhRealServerTableExcludeStr = 2147483647
)

type SlbCurCfgEnhRealServerTableSubmac int32

const (
	SlbCurCfgEnhRealServerTableSubmac_Enabled     SlbCurCfgEnhRealServerTableSubmac = 1
	SlbCurCfgEnhRealServerTableSubmac_Disabled    SlbCurCfgEnhRealServerTableSubmac = 2
	SlbCurCfgEnhRealServerTableSubmac_Unsupported SlbCurCfgEnhRealServerTableSubmac = 2147483647
)

type SlbCurCfgEnhRealServerTableIpVer int32

const (
	SlbCurCfgEnhRealServerTableIpVer_Ipv4        SlbCurCfgEnhRealServerTableIpVer = 1
	SlbCurCfgEnhRealServerTableIpVer_Ipv6        SlbCurCfgEnhRealServerTableIpVer = 2
	SlbCurCfgEnhRealServerTableIpVer_Unsupported SlbCurCfgEnhRealServerTableIpVer = 2147483647
)

type SlbCurCfgEnhRealServerTableLLBType int32

const (
	SlbCurCfgEnhRealServerTableLLBType_Local       SlbCurCfgEnhRealServerTableLLBType = 0
	SlbCurCfgEnhRealServerTableLLBType_Remote      SlbCurCfgEnhRealServerTableLLBType = 1
	SlbCurCfgEnhRealServerTableLLBType_Wanlink     SlbCurCfgEnhRealServerTableLLBType = 2
	SlbCurCfgEnhRealServerTableLLBType_Unsupported SlbCurCfgEnhRealServerTableLLBType = 2147483647
)

type SlbCurCfgEnhRealServerTableSecType int32

const (
	SlbCurCfgEnhRealServerTableSecType_None        SlbCurCfgEnhRealServerTableSecType = 1
	SlbCurCfgEnhRealServerTableSecType_Virtual     SlbCurCfgEnhRealServerTableSecType = 2
	SlbCurCfgEnhRealServerTableSecType_Layer       SlbCurCfgEnhRealServerTableSecType = 3
	SlbCurCfgEnhRealServerTableSecType_Passive     SlbCurCfgEnhRealServerTableSecType = 4
	SlbCurCfgEnhRealServerTableSecType_L3sw        SlbCurCfgEnhRealServerTableSecType = 5
	SlbCurCfgEnhRealServerTableSecType_Unsupported SlbCurCfgEnhRealServerTableSecType = 2147483647
)

type SlbCurCfgEnhRealServerTableIpVerDuplicate int32

const (
	SlbCurCfgEnhRealServerTableIpVerDuplicate_Ipv4        SlbCurCfgEnhRealServerTableIpVerDuplicate = 1
	SlbCurCfgEnhRealServerTableIpVerDuplicate_Ipv6        SlbCurCfgEnhRealServerTableIpVerDuplicate = 2
	SlbCurCfgEnhRealServerTableIpVerDuplicate_Unsupported SlbCurCfgEnhRealServerTableIpVerDuplicate = 2147483647
)

type SlbCurCfgEnhRealServerTableSecDeviceFlag int32

const (
	SlbCurCfgEnhRealServerTableSecDeviceFlag_None        SlbCurCfgEnhRealServerTableSecDeviceFlag = 1
	SlbCurCfgEnhRealServerTableSecDeviceFlag_Security    SlbCurCfgEnhRealServerTableSecDeviceFlag = 2
	SlbCurCfgEnhRealServerTableSecDeviceFlag_Unsupported SlbCurCfgEnhRealServerTableSecDeviceFlag = 2147483647
)

type SlbCurCfgEnhRealServerTableParams struct {
	// The real server number
	Index string `json:"Index,omitempty"`
	// IP address of the real server identified by the instance of the
	// slbRealServerIndex.
	IpAddr string `json:"IpAddr,omitempty"`
	// The server weight.
	Weight uint64 `json:"Weight,omitempty"`
	// The maximum number of connections that are allowed.
	MaxConns uint32 `json:"MaxConns,omitempty"`
	// The maximum number of minutes an inactive connection remains open.
	TimeOut uint32 `json:"TimeOut,omitempty"`
	// The interval between keep-alive (ping) attempts in number of seconds.
	// Zero means disabling ping attempt.
	PingInterval uint64 `json:"PingInterval,omitempty"`
	// The number of failed attempts to declare this server DOWN.
	FailRetry uint64 `json:"FailRetry,omitempty"`
	// The number of successful attempts to declare a server UP.
	SuccRetry uint64 `json:"SuccRetry,omitempty"`
	// Enable or disable the server and remove the existing sessions using disabled-with-fastage option.
	State SlbCurCfgEnhRealServerTableState `json:"State,omitempty"`
	// The server type.  It participates in global server
	// load balancing when it is configured as remote-server.
	Type SlbCurCfgEnhRealServerTableType `json:"Type,omitempty"`
	// The name of the real server.
	Name string `json:"Name,omitempty"`
	// The real server that will handle client requests that doesn't
	// contain an URL cookie if Cookie loadbalance is enabled.
	Cookie SlbCurCfgEnhRealServerTableCookie `json:"Cookie,omitempty"`
	// The real server will handle requests that don't match the
	// loadbalance string if it is enabled.
	ExcludeStr SlbCurCfgEnhRealServerTableExcludeStr `json:"ExcludeStr,omitempty"`
	// The real server config to enable/disable MAC SA substitution for
	// L4 traffic. If disabled (the default) we will NOT substitute the
	// MAC SA of client-to-server frames.  If enabled, we will substitute
	// the MAC SA.
	Submac SlbCurCfgEnhRealServerTableSubmac `json:"Submac,omitempty"`
	// The port to be connected to the IDS server.
	Idsport int32 `json:"Idsport,omitempty"`
	// The type of IP address.
	IpVer SlbCurCfgEnhRealServerTableIpVer `json:"IpVer,omitempty"`
	// IPV6 address of the real server identified by the instance of the slbRealServerIndex.
	// Address should be 4-byte hexadecimal colon notation. Valid IPv6 address should be in
	// any of the following forms xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx or xxxx::xxxx:xxxx:xxxx:xxxx or ::xxxx
	Ipv6Addr string `json:"Ipv6Addr,omitempty"`
	// The server type.
	LLBType SlbCurCfgEnhRealServerTableLLBType `json:"LLBType,omitempty"`
	// List of ingress ports attached to the real server (security device), used for SSL inspection WebUI wizard.
	// 	Using the following format: xx, xx, xx
	PortsIngress string `json:"PortsIngress,omitempty"`
	// List of egress ports attached to the real server (security device), used for SSL inspection WebUI wizard.
	// 	Using the following format: xx, xx, xx
	PortsEgress string `json:"PortsEgress,omitempty"`
	// The ingress Vlan specified on security device.
	// 	 Used for SSL wizard
	VlanIngress int32 `json:"VlanIngress,omitempty"`
	// The egress Vlan specified on security device.
	// 	 Used for SSL wizard
	VlanEgress int32 `json:"VlanEgress,omitempty"`
	// The egress interface specified on security device.
	// 	 Used for SSL wizard
	EgressIf int32 `json:"EgressIf,omitempty"`
	// The security device  type.
	SecType SlbCurCfgEnhRealServerTableSecType `json:"SecType,omitempty"`
	// The ingress interface specified on security device.
	// 	 Used for SSL wizard
	IngressIf int32 `json:"IngressIf,omitempty"`
	// The type of IP address.
	IpVerDuplicate SlbCurCfgEnhRealServerTableIpVerDuplicate `json:"IpVerDuplicate,omitempty"`
	// The Real security device flag
	SecDeviceFlag SlbCurCfgEnhRealServerTableSecDeviceFlag `json:"SecDeviceFlag,omitempty"`
	// The port to be connected to the IDS server.
	Ingport int32 `json:"Ingport,omitempty"`
}

func (p SlbCurCfgEnhRealServerTableParams) iMABean() {}
