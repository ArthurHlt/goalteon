package beans

import (
	"fmt"
	"reflect"
)

// SlbNewCfgEnhRealServerTable The table of Real Server configuration in the new_config.
// Note:This mib is not supported on VX instance of Virtualization.
type SlbNewCfgEnhRealServerTable struct {
	// The real server number
	SlbNewCfgEnhRealServerIndex string
	Params                      *SlbNewCfgEnhRealServerTableParams
}

func NewSlbNewCfgEnhRealServerTableList() *SlbNewCfgEnhRealServerTable {
	return &SlbNewCfgEnhRealServerTable{}
}

func NewSlbNewCfgEnhRealServerTable(
	slbNewCfgEnhRealServerIndex string,
	params *SlbNewCfgEnhRealServerTableParams,
) *SlbNewCfgEnhRealServerTable {
	return &SlbNewCfgEnhRealServerTable{
		SlbNewCfgEnhRealServerIndex: slbNewCfgEnhRealServerIndex,
		Params:                      params,
	}
}

func (c *SlbNewCfgEnhRealServerTable) Name() string {
	return "SlbNewCfgEnhRealServerTable"
}

func (c *SlbNewCfgEnhRealServerTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbNewCfgEnhRealServerTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbNewCfgEnhRealServerTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbNewCfgEnhRealServerIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbNewCfgEnhRealServerIndex)
}

type SlbNewCfgEnhRealServerTableState int32

const (
	SlbNewCfgEnhRealServerTableState_Enabled                    SlbNewCfgEnhRealServerTableState = 2
	SlbNewCfgEnhRealServerTableState_Disabled                   SlbNewCfgEnhRealServerTableState = 3
	SlbNewCfgEnhRealServerTableState_DisabledWithFastage        SlbNewCfgEnhRealServerTableState = 4
	SlbNewCfgEnhRealServerTableState_ShutdownConnection         SlbNewCfgEnhRealServerTableState = 5
	SlbNewCfgEnhRealServerTableState_ShutdownPersistentSessions SlbNewCfgEnhRealServerTableState = 6
	SlbNewCfgEnhRealServerTableState_Unsupported                SlbNewCfgEnhRealServerTableState = 2147483647
)

type SlbNewCfgEnhRealServerTableDelete int32

const (
	SlbNewCfgEnhRealServerTableDelete_Other       SlbNewCfgEnhRealServerTableDelete = 1
	SlbNewCfgEnhRealServerTableDelete_Delete      SlbNewCfgEnhRealServerTableDelete = 2
	SlbNewCfgEnhRealServerTableDelete_Unsupported SlbNewCfgEnhRealServerTableDelete = 2147483647
)

type SlbNewCfgEnhRealServerTableType int32

const (
	SlbNewCfgEnhRealServerTableType_LocalServer  SlbNewCfgEnhRealServerTableType = 1
	SlbNewCfgEnhRealServerTableType_RemoteServer SlbNewCfgEnhRealServerTableType = 2
	SlbNewCfgEnhRealServerTableType_Unsupported  SlbNewCfgEnhRealServerTableType = 2147483647
)

type SlbNewCfgEnhRealServerTableCookie int32

const (
	SlbNewCfgEnhRealServerTableCookie_Enabled     SlbNewCfgEnhRealServerTableCookie = 1
	SlbNewCfgEnhRealServerTableCookie_Disabled    SlbNewCfgEnhRealServerTableCookie = 2
	SlbNewCfgEnhRealServerTableCookie_Unsupported SlbNewCfgEnhRealServerTableCookie = 2147483647
)

type SlbNewCfgEnhRealServerTableExcludeStr int32

const (
	SlbNewCfgEnhRealServerTableExcludeStr_Enabled     SlbNewCfgEnhRealServerTableExcludeStr = 1
	SlbNewCfgEnhRealServerTableExcludeStr_Disabled    SlbNewCfgEnhRealServerTableExcludeStr = 2
	SlbNewCfgEnhRealServerTableExcludeStr_Unsupported SlbNewCfgEnhRealServerTableExcludeStr = 2147483647
)

type SlbNewCfgEnhRealServerTableSubmac int32

const (
	SlbNewCfgEnhRealServerTableSubmac_Enabled     SlbNewCfgEnhRealServerTableSubmac = 1
	SlbNewCfgEnhRealServerTableSubmac_Disabled    SlbNewCfgEnhRealServerTableSubmac = 2
	SlbNewCfgEnhRealServerTableSubmac_Unsupported SlbNewCfgEnhRealServerTableSubmac = 2147483647
)

type SlbNewCfgEnhRealServerTableIpVer int32

const (
	SlbNewCfgEnhRealServerTableIpVer_Ipv4        SlbNewCfgEnhRealServerTableIpVer = 1
	SlbNewCfgEnhRealServerTableIpVer_Ipv6        SlbNewCfgEnhRealServerTableIpVer = 2
	SlbNewCfgEnhRealServerTableIpVer_Unsupported SlbNewCfgEnhRealServerTableIpVer = 2147483647
)

type SlbNewCfgEnhRealServerTableLLBType int32

const (
	SlbNewCfgEnhRealServerTableLLBType_Local       SlbNewCfgEnhRealServerTableLLBType = 0
	SlbNewCfgEnhRealServerTableLLBType_Remote      SlbNewCfgEnhRealServerTableLLBType = 1
	SlbNewCfgEnhRealServerTableLLBType_Wanlink     SlbNewCfgEnhRealServerTableLLBType = 2
	SlbNewCfgEnhRealServerTableLLBType_Unsupported SlbNewCfgEnhRealServerTableLLBType = 2147483647
)

type SlbNewCfgEnhRealServerTableSecType int32

const (
	SlbNewCfgEnhRealServerTableSecType_None        SlbNewCfgEnhRealServerTableSecType = 1
	SlbNewCfgEnhRealServerTableSecType_Virtual     SlbNewCfgEnhRealServerTableSecType = 2
	SlbNewCfgEnhRealServerTableSecType_Layer       SlbNewCfgEnhRealServerTableSecType = 3
	SlbNewCfgEnhRealServerTableSecType_Passive     SlbNewCfgEnhRealServerTableSecType = 4
	SlbNewCfgEnhRealServerTableSecType_L3sw        SlbNewCfgEnhRealServerTableSecType = 5
	SlbNewCfgEnhRealServerTableSecType_Unsupported SlbNewCfgEnhRealServerTableSecType = 2147483647
)

type SlbNewCfgEnhRealServerTableSecDeviceFlag int32

const (
	SlbNewCfgEnhRealServerTableSecDeviceFlag_None        SlbNewCfgEnhRealServerTableSecDeviceFlag = 1
	SlbNewCfgEnhRealServerTableSecDeviceFlag_Security    SlbNewCfgEnhRealServerTableSecDeviceFlag = 2
	SlbNewCfgEnhRealServerTableSecDeviceFlag_Unsupported SlbNewCfgEnhRealServerTableSecDeviceFlag = 2147483647
)

type SlbNewCfgEnhRealServerTableParams struct {
	// The real server number
	Index string `json:"Index,omitempty"`
	// IP address of the real server identified by the instance of
	// slbRealServerIndex.
	IpAddr string `json:"IpAddr,omitempty"`
	// The server weight.
	Weight uint64 `json:"Weight,omitempty"`
	// The maximum number of connections that are allowed.
	MaxConns uint32 `json:"MaxConns,omitempty"`
	// The maximum number of minutes an inactive connection remains open.
	TimeOut uint32 `json:"TimeOut,omitempty"`
	// The interval between keep-alive (ping) attempts in number of
	// seconds. Zero means disabling ping attempt.
	PingInterval uint64 `json:"PingInterval,omitempty"`
	// The number of failed attempts to declare this server DOWN.
	FailRetry uint64 `json:"FailRetry,omitempty"`
	// The number of successful attempts to declare a server UP.
	SuccRetry uint64 `json:"SuccRetry,omitempty"`
	// Enable or disable the server and remove the existing sessions using disabled-with-fastage option.
	State SlbNewCfgEnhRealServerTableState `json:"State,omitempty"`
	// By setting the value to delete(2), the entire row is deleted.
	Delete SlbNewCfgEnhRealServerTableDelete `json:"Delete,omitempty"`
	// The server type.  It participates in global server load balancing
	// when it is configured as remote-server.
	Type SlbNewCfgEnhRealServerTableType `json:"Type,omitempty"`
	// The name of the real server.
	Name string `json:"Name,omitempty"`
	// The URL Path (slbCurCfgUrlLbPathIndex) to be added to the
	// real server. A zero is returned when read.
	AddUrl int32 `json:"AddUrl,omitempty"`
	// The URL Path (slbCurCfgUrlLbPathIndex) to be removed from
	// the real server. A zero is returned when read.
	RemUrl int32 `json:"RemUrl,omitempty"`
	// Enable or disable real server to handle client requests
	// that don't contain a cookie if cookie loadbalance is enabled.
	Cookie SlbNewCfgEnhRealServerTableCookie `json:"Cookie,omitempty"`
	// Enable or disable exclusionary matching string on real server.
	ExcludeStr SlbNewCfgEnhRealServerTableExcludeStr `json:"ExcludeStr,omitempty"`
	// The real server config to enable/disable MAC SA substitution for
	// L4 traffic. If disabled (the default) we will NOT substitute the
	// MAC SA of client-to-server frames.  If enabled, we will substitute
	// the MAC SA.
	Submac SlbNewCfgEnhRealServerTableSubmac `json:"Submac,omitempty"`
	// The port to be connected to IDS server.
	Idsport int32 `json:"Idsport,omitempty"`
	// The type of IP address.
	IpVer SlbNewCfgEnhRealServerTableIpVer `json:"IpVer,omitempty"`
	// IPV6 address of the real server identified by the instance of the slbRealServerIndex.
	// Address should be 4-byte hexadecimal colon notation. Valid IPv6 address should be in
	// any of the following forms xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx or xxxx::xxxx:xxxx:xxxx:xxxx or ::xxxx
	Ipv6Addr string `json:"Ipv6Addr,omitempty"`
	// The next available free slot index number, to add the real port
	// to the server. Value 0 will be returned if no free slot available.
	NxtRportIdx uint64 `json:"NxtRportIdx,omitempty"`
	// The next available free slot Buddy index number, to add the Buddy Server
	// to the Real server. Value 0 will be returned if no free slot available.
	NxtBuddyIdx uint32 `json:"NxtBuddyIdx,omitempty"`
	// The server type.
	LLBType SlbNewCfgEnhRealServerTableLLBType `json:"LLBType,omitempty"`
	// The alphanumeric index of the new copy to be created.
	Copy string `json:"Copy,omitempty"`
	// List of ingress ports attached to the real server (security device), used for SSL inspection WebUI wizard.
	// 	Using the following format: xx, xx, xx
	PortsIngress string `json:"PortsIngress,omitempty"`
	// List of egress ports attached to the real server (security device), used for SSL inspection WebUI wizard.
	// 	Using the following format: xx, xx, xx
	PortsEgress string `json:"PortsEgress,omitempty"`
	// The ingress port to be added to the specified security device.
	// 	  A '0' value is returned when read.
	AddPortsIngress int32 `json:"AddPortsIngress,omitempty"`
	// The ingress port to be removed to the specified security device.
	// 	 A zero is returned when read.
	RemPortsIngress int32 `json:"RemPortsIngress,omitempty"`
	// The egress port to be added to the specified security device.
	// 	 A zero is returned when read.
	AddPortsEgress int32 `json:"AddPortsEgress,omitempty"`
	// The egress port to be removed to the specified security device.
	// 	 A zero is returned when read.
	RemPortsEgress int32 `json:"RemPortsEgress,omitempty"`
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
	SecType SlbNewCfgEnhRealServerTableSecType `json:"SecType,omitempty"`
	// The ingress interface specified on security device.
	// 	 Used for SSL wizard
	IngressIf int32 `json:"IngressIf,omitempty"`
	// The Real security device flag.
	SecDeviceFlag SlbNewCfgEnhRealServerTableSecDeviceFlag `json:"SecDeviceFlag,omitempty"`
	// The ingress port to be connected to IDS server.
	Ingport int32 `json:"Ingport,omitempty"`
}

func (p SlbNewCfgEnhRealServerTableParams) iMABean() {}
