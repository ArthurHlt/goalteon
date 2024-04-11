package beans

import (
	"fmt"
	"reflect"
)

// SlbCurCfgRealServerTable The table of Real Server configuration in the current_config.
// Note:This mib is not supported on VX instance of Virtualization.
type SlbCurCfgRealServerTable struct {
	// The real server number
	SlbCurCfgRealServerIndex int32
	Params                   *SlbCurCfgRealServerTableParams
}

func NewSlbCurCfgRealServerTableList() *SlbCurCfgRealServerTable {
	return &SlbCurCfgRealServerTable{}
}

func NewSlbCurCfgRealServerTable(
	slbCurCfgRealServerIndex int32,
	params *SlbCurCfgRealServerTableParams,
) *SlbCurCfgRealServerTable {
	return &SlbCurCfgRealServerTable{
		SlbCurCfgRealServerIndex: slbCurCfgRealServerIndex,
		Params:                   params,
	}
}

func (c *SlbCurCfgRealServerTable) Name() string {
	return "SlbCurCfgRealServerTable"
}

func (c *SlbCurCfgRealServerTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbCurCfgRealServerTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbCurCfgRealServerTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbCurCfgRealServerIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbCurCfgRealServerIndex)
}

type SlbCurCfgRealServerTableState int32

const (
	SlbCurCfgRealServerTableState_Enabled             SlbCurCfgRealServerTableState = 2
	SlbCurCfgRealServerTableState_Disabled            SlbCurCfgRealServerTableState = 3
	SlbCurCfgRealServerTableState_DisabledWithFastage SlbCurCfgRealServerTableState = 4
	SlbCurCfgRealServerTableState_Unsupported         SlbCurCfgRealServerTableState = 2147483647
)

type SlbCurCfgRealServerTableType int32

const (
	SlbCurCfgRealServerTableType_LocalServer  SlbCurCfgRealServerTableType = 1
	SlbCurCfgRealServerTableType_RemoteServer SlbCurCfgRealServerTableType = 2
	SlbCurCfgRealServerTableType_Unsupported  SlbCurCfgRealServerTableType = 2147483647
)

type SlbCurCfgRealServerTableCookie int32

const (
	SlbCurCfgRealServerTableCookie_Enabled     SlbCurCfgRealServerTableCookie = 1
	SlbCurCfgRealServerTableCookie_Disabled    SlbCurCfgRealServerTableCookie = 2
	SlbCurCfgRealServerTableCookie_Unsupported SlbCurCfgRealServerTableCookie = 2147483647
)

type SlbCurCfgRealServerTableExcludeStr int32

const (
	SlbCurCfgRealServerTableExcludeStr_Enabled     SlbCurCfgRealServerTableExcludeStr = 1
	SlbCurCfgRealServerTableExcludeStr_Disabled    SlbCurCfgRealServerTableExcludeStr = 2
	SlbCurCfgRealServerTableExcludeStr_Unsupported SlbCurCfgRealServerTableExcludeStr = 2147483647
)

type SlbCurCfgRealServerTableSubmac int32

const (
	SlbCurCfgRealServerTableSubmac_Enabled     SlbCurCfgRealServerTableSubmac = 1
	SlbCurCfgRealServerTableSubmac_Disabled    SlbCurCfgRealServerTableSubmac = 2
	SlbCurCfgRealServerTableSubmac_Unsupported SlbCurCfgRealServerTableSubmac = 2147483647
)

type SlbCurCfgRealServerTableProxy int32

const (
	SlbCurCfgRealServerTableProxy_Enabled     SlbCurCfgRealServerTableProxy = 1
	SlbCurCfgRealServerTableProxy_Disabled    SlbCurCfgRealServerTableProxy = 2
	SlbCurCfgRealServerTableProxy_Unsupported SlbCurCfgRealServerTableProxy = 2147483647
)

type SlbCurCfgRealServerTableLdapwr int32

const (
	SlbCurCfgRealServerTableLdapwr_Enabled     SlbCurCfgRealServerTableLdapwr = 1
	SlbCurCfgRealServerTableLdapwr_Disabled    SlbCurCfgRealServerTableLdapwr = 2
	SlbCurCfgRealServerTableLdapwr_Unsupported SlbCurCfgRealServerTableLdapwr = 2147483647
)

type SlbCurCfgRealServerTableFastHealthCheck int32

const (
	SlbCurCfgRealServerTableFastHealthCheck_Enabled     SlbCurCfgRealServerTableFastHealthCheck = 1
	SlbCurCfgRealServerTableFastHealthCheck_Disabled    SlbCurCfgRealServerTableFastHealthCheck = 2
	SlbCurCfgRealServerTableFastHealthCheck_Unsupported SlbCurCfgRealServerTableFastHealthCheck = 2147483647
)

type SlbCurCfgRealServerTableSubdmac int32

const (
	SlbCurCfgRealServerTableSubdmac_Enabled     SlbCurCfgRealServerTableSubdmac = 1
	SlbCurCfgRealServerTableSubdmac_Disabled    SlbCurCfgRealServerTableSubdmac = 2
	SlbCurCfgRealServerTableSubdmac_Unsupported SlbCurCfgRealServerTableSubdmac = 2147483647
)

type SlbCurCfgRealServerTableOverflow int32

const (
	SlbCurCfgRealServerTableOverflow_Enabled     SlbCurCfgRealServerTableOverflow = 1
	SlbCurCfgRealServerTableOverflow_Disabled    SlbCurCfgRealServerTableOverflow = 2
	SlbCurCfgRealServerTableOverflow_Unsupported SlbCurCfgRealServerTableOverflow = 2147483647
)

type SlbCurCfgRealServerTableBkpPreempt int32

const (
	SlbCurCfgRealServerTableBkpPreempt_Enabled     SlbCurCfgRealServerTableBkpPreempt = 1
	SlbCurCfgRealServerTableBkpPreempt_Disabled    SlbCurCfgRealServerTableBkpPreempt = 2
	SlbCurCfgRealServerTableBkpPreempt_Unsupported SlbCurCfgRealServerTableBkpPreempt = 2147483647
)

type SlbCurCfgRealServerTableIpVer int32

const (
	SlbCurCfgRealServerTableIpVer_Ipv4        SlbCurCfgRealServerTableIpVer = 1
	SlbCurCfgRealServerTableIpVer_Ipv6        SlbCurCfgRealServerTableIpVer = 2
	SlbCurCfgRealServerTableIpVer_Unsupported SlbCurCfgRealServerTableIpVer = 2147483647
)

type SlbCurCfgRealServerTableMode int32

const (
	SlbCurCfgRealServerTableMode_Physical    SlbCurCfgRealServerTableMode = 1
	SlbCurCfgRealServerTableMode_Logical     SlbCurCfgRealServerTableMode = 2
	SlbCurCfgRealServerTableMode_Unsupported SlbCurCfgRealServerTableMode = 2147483647
)

type SlbCurCfgRealServerTableProxyIpMode int32

const (
	SlbCurCfgRealServerTableProxyIpMode_Enable      SlbCurCfgRealServerTableProxyIpMode = 0
	SlbCurCfgRealServerTableProxyIpMode_Address     SlbCurCfgRealServerTableProxyIpMode = 2
	SlbCurCfgRealServerTableProxyIpMode_Nwclss      SlbCurCfgRealServerTableProxyIpMode = 3
	SlbCurCfgRealServerTableProxyIpMode_Disable     SlbCurCfgRealServerTableProxyIpMode = 4
	SlbCurCfgRealServerTableProxyIpMode_Unsupported SlbCurCfgRealServerTableProxyIpMode = 2147483647
)

type SlbCurCfgRealServerTableProxyIpPersistency int32

const (
	SlbCurCfgRealServerTableProxyIpPersistency_Disable     SlbCurCfgRealServerTableProxyIpPersistency = 0
	SlbCurCfgRealServerTableProxyIpPersistency_Client      SlbCurCfgRealServerTableProxyIpPersistency = 1
	SlbCurCfgRealServerTableProxyIpPersistency_Host        SlbCurCfgRealServerTableProxyIpPersistency = 2
	SlbCurCfgRealServerTableProxyIpPersistency_Unsupported SlbCurCfgRealServerTableProxyIpPersistency = 2147483647
)

type SlbCurCfgRealServerTableProxyIpNWclassPersistency int32

const (
	SlbCurCfgRealServerTableProxyIpNWclassPersistency_Disable     SlbCurCfgRealServerTableProxyIpNWclassPersistency = 0
	SlbCurCfgRealServerTableProxyIpNWclassPersistency_Client      SlbCurCfgRealServerTableProxyIpNWclassPersistency = 1
	SlbCurCfgRealServerTableProxyIpNWclassPersistency_Unsupported SlbCurCfgRealServerTableProxyIpNWclassPersistency = 2147483647
)

type SlbCurCfgRealServerTableParams struct {
	// The real server number
	Index int32 `json:"Index,omitempty"`
	// IP address of the real server identified by the instance of the
	// slbRealServerIndex.
	IpAddr string `json:"IpAddr,omitempty"`
	// The server weight.
	Weight uint64 `json:"Weight,omitempty"`
	// The maximum number of connections that are allowed.
	MaxConns uint32 `json:"MaxConns,omitempty"`
	// The maximum number of minutes an inactive connection remains open.
	TimeOut uint32 `json:"TimeOut,omitempty"`
	// The backup server number for this server.
	BackUp int32 `json:"BackUp,omitempty"`
	// The interval between keep-alive (ping) attempts in number of seconds.
	// Zero means disabling ping attempt.
	PingInterval uint64 `json:"PingInterval,omitempty"`
	// The number of failed attempts to declare this server DOWN.
	FailRetry uint64 `json:"FailRetry,omitempty"`
	// The number of successful attempts to declare a server UP.
	SuccRetry uint64 `json:"SuccRetry,omitempty"`
	// Enable or disable the server and remove the existing sessions using disabled-with-fastage option.
	State SlbCurCfgRealServerTableState `json:"State,omitempty"`
	// The server type.  It participates in global server
	// 	 load balancing when it is configured as remote-server.
	Type SlbCurCfgRealServerTableType `json:"Type,omitempty"`
	// The name of the real server.
	Name string `json:"Name,omitempty"`
	// The URL Paths selected for URL load balancing for by the real
	// server.  The selected URL Paths are presented in a bitmap format.
	// in receiving order:
	// OCTET 1  OCTET 2  .....
	// xxxxxxxx xxxxxxxx .....
	// ||    || |_ URL Path 9
	// ||    ||
	// ||    ||___ URL Path 8
	// ||    |____ URL Path 7
	// ||      .    .   .
	// ||_________ URL Path 2
	// |__________ URL Path 1
	// where x : 1 - The represented URL Path is selected
	// 0 - The represented URL Path is not selected
	UrlBmap string `json:"UrlBmap,omitempty"`
	// The real server that will handle client requests that doesn't
	// contain an URL cookie if Cookie loadbalance is enabled.
	Cookie SlbCurCfgRealServerTableCookie `json:"Cookie,omitempty"`
	// The real server will handle requests that don't match the
	// loadbalance string if it is enabled.
	ExcludeStr SlbCurCfgRealServerTableExcludeStr `json:"ExcludeStr,omitempty"`
	// The real server config to enable/disable MAC SA substitution for
	// L4 traffic. If disabled (the default) we will NOT substitute the
	// MAC SA of client-to-server frames.  If enabled, we will substitute
	// the MAC SA.
	Submac SlbCurCfgRealServerTableSubmac `json:"Submac,omitempty"`
	// The real server config to enable/disable client proxy operation.
	Proxy SlbCurCfgRealServerTableProxy `json:"Proxy,omitempty"`
	// The real server config to enable/disable LDAP write server.
	Ldapwr SlbCurCfgRealServerTableLdapwr `json:"Ldapwr,omitempty"`
	// The OID to be sent in the SNMP get request packet.
	Oid string `json:"Oid,omitempty"`
	// The community string to be used in the SNMP get request packet.
	CommString string `json:"CommString,omitempty"`
	// The VLAN to be associated with the IDS server.
	Idsvlan uint32 `json:"Idsvlan,omitempty"`
	// The port to be connected to the IDS server.
	Idsport int32 `json:"Idsport,omitempty"`
	// The remote real server Global SLB availability.
	Avail uint64 `json:"Avail,omitempty"`
	// The real server config to enable/disable Fast Health Check Operation.
	FastHealthCheck SlbCurCfgRealServerTableFastHealthCheck `json:"FastHealthCheck,omitempty"`
	// The real server config to enable/disable MAC DA substitution for
	// L4 traffic. If disabled, we will NOT substitute the MAC DA of
	// client-to-server frames.  If enabled(default), we will substitute
	// the MAC DA.
	Subdmac SlbCurCfgRealServerTableSubdmac `json:"Subdmac,omitempty"`
	// The real server config to enable/disable Overflow. If enabled(default)
	// allows Backup server to kick in if real server reaches maximum
	// connections.
	Overflow SlbCurCfgRealServerTableOverflow `json:"Overflow,omitempty"`
	// The real server config to enable/disable backup preemption. If enabled
	// (default)allows to preempt the backup server when the primary server
	// comes up. If disabled the backup server will continue to serve the
	// traffic even when primary server comes up.
	BkpPreempt SlbCurCfgRealServerTableBkpPreempt `json:"BkpPreempt,omitempty"`
	// The type of IP address.
	IpVer SlbCurCfgRealServerTableIpVer `json:"IpVer,omitempty"`
	// IPV6 address of the real server identified by the instance of the slbRealServerIndex.
	// Address should be 4-byte hexadecimal colon notation. Valid IPv6 address should be in
	// any of the following forms xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx or xxxx::xxxx:xxxx:xxxx:xxxx or ::xxxx
	Ipv6Addr string `json:"Ipv6Addr,omitempty"`
	// The mode of the real server. By default it is set to physical.
	Mode SlbCurCfgRealServerTableMode `json:"Mode,omitempty"`
	// the real server Proxy IP mode
	ProxyIpMode SlbCurCfgRealServerTableProxyIpMode `json:"ProxyIpMode,omitempty"`
	// This object ID shows current configuration of real server proxy IP address .
	// Returns 0 when slbCurCfgRealServerProxyIpMode is not set to address.
	// Returns 0 if IP version is IPv6.
	ProxyIpAddr string `json:"ProxyIpAddr,omitempty"`
	// This object ID shows current configuration of real server proxy IP subnet mask.
	// Returns 0 when slbCurCfgRealServerProxyIpMode is not set to address.
	// Returns 0 if IP version is IPv6.
	ProxyIpMask string `json:"ProxyIpMask,omitempty"`
	// This object ID shows current configuration of real server proxy IPv6 address .
	// 		The IP version for addr must be the same as the real server IP version.
	// Returns 0 when slbCurCfgRealServerProxyIpMode is not set to address.
	// Returns 0 if IP version is IPv4.
	// 		When a subnet is configured, user has the ability to select PIP persistency mode.
	// 		Persistency is relevant only if the PIP is configured as subnet. If PIP is not configured or
	// 		not subnet, the persistency configuration value is disable.
	ProxyIpv6Addr string `json:"ProxyIpv6Addr,omitempty"`
	// This object ID allows configuration of real server proxy IPv6 Mask.
	// 		The IP version for addr must be the same as the real server IP version.
	// Returns 0 when slbCurCfgRealServerProxyIpMode is not set to address.
	// Returns 0 if IP version is IPv6.
	// 		When a subnet is configured user has the ability to select PIP persistency mode.
	// 		Persistency is relevant only if the PIP is configured as subnet. If PIP is not configured or
	// 		not subnet, the persistency configuration value is disable.
	ProxyIpv6Prefix uint32 `json:"ProxyIpv6Prefix,omitempty"`
	// This object ID shows current configuration of real server proxy IP persistency mode.
	// 		Persistency is relevant only if a PIP class is configured. If PIP is not configured the
	// 		persistency configuration is disable.
	ProxyIpPersistency SlbCurCfgRealServerTableProxyIpPersistency `json:"ProxyIpPersistency,omitempty"`
	// This  object ID allows configuration of real server proxy IP IPv4 or IPv6 Network Class as PIP.
	// 		Network Class PIP Persistency is relevant only if a PIP class is configured. If PIP is not configured the
	// 		persistency configuration is disable.
	// Returns 0 when slbCurCfgRealServerProxyIpMode is not set to nwclss(3).
	ProxyIpNWclass string `json:"ProxyIpNWclass,omitempty"`
	// This object ID allows configuration of real server Network Class PIP persistency mode.
	// 		Persistency is relevant only if a PIP class is configured. If PIP is not configured the
	// 		persistency configuration is disable.
	// Returns 0 when slbCurCfgRealServerProxyIpMode is not set to nwclss(3).
	ProxyIpNWclassPersistency SlbCurCfgRealServerTableProxyIpNWclassPersistency `json:"ProxyIpNWclassPersistency,omitempty"`
	// The Advanced HC ID.
	HealthID string `json:"HealthID,omitempty"`
}

func (p SlbCurCfgRealServerTableParams) iMABean() {}
