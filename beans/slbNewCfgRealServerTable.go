package beans

import (
	"fmt"
	"reflect"
)

// SlbNewCfgRealServerTable The table of Real Server configuration in the new_config.
// Note:This mib is not supported on VX instance of Virtualization.
type SlbNewCfgRealServerTable struct {
	// The real server number
	SlbNewCfgRealServerIndex int32
	Params                   *SlbNewCfgRealServerTableParams
}

func NewSlbNewCfgRealServerTableList() *SlbNewCfgRealServerTable {
	return &SlbNewCfgRealServerTable{}
}

func NewSlbNewCfgRealServerTable(
	slbNewCfgRealServerIndex int32,
	params *SlbNewCfgRealServerTableParams,
) *SlbNewCfgRealServerTable {
	return &SlbNewCfgRealServerTable{
		SlbNewCfgRealServerIndex: slbNewCfgRealServerIndex,
		Params:                   params,
	}
}

func (c *SlbNewCfgRealServerTable) Name() string {
	return "SlbNewCfgRealServerTable"
}

func (c *SlbNewCfgRealServerTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbNewCfgRealServerTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbNewCfgRealServerTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbNewCfgRealServerIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbNewCfgRealServerIndex)
}

type SlbNewCfgRealServerTableState int32

const (
	SlbNewCfgRealServerTableState_Enabled             SlbNewCfgRealServerTableState = 2
	SlbNewCfgRealServerTableState_Disabled            SlbNewCfgRealServerTableState = 3
	SlbNewCfgRealServerTableState_DisabledWithFastage SlbNewCfgRealServerTableState = 4
	SlbNewCfgRealServerTableState_Unsupported         SlbNewCfgRealServerTableState = 2147483647
)

type SlbNewCfgRealServerTableDelete int32

const (
	SlbNewCfgRealServerTableDelete_Other       SlbNewCfgRealServerTableDelete = 1
	SlbNewCfgRealServerTableDelete_Delete      SlbNewCfgRealServerTableDelete = 2
	SlbNewCfgRealServerTableDelete_Unsupported SlbNewCfgRealServerTableDelete = 2147483647
)

type SlbNewCfgRealServerTableType int32

const (
	SlbNewCfgRealServerTableType_LocalServer  SlbNewCfgRealServerTableType = 1
	SlbNewCfgRealServerTableType_RemoteServer SlbNewCfgRealServerTableType = 2
	SlbNewCfgRealServerTableType_Unsupported  SlbNewCfgRealServerTableType = 2147483647
)

type SlbNewCfgRealServerTableCookie int32

const (
	SlbNewCfgRealServerTableCookie_Enabled     SlbNewCfgRealServerTableCookie = 1
	SlbNewCfgRealServerTableCookie_Disabled    SlbNewCfgRealServerTableCookie = 2
	SlbNewCfgRealServerTableCookie_Unsupported SlbNewCfgRealServerTableCookie = 2147483647
)

type SlbNewCfgRealServerTableExcludeStr int32

const (
	SlbNewCfgRealServerTableExcludeStr_Enabled     SlbNewCfgRealServerTableExcludeStr = 1
	SlbNewCfgRealServerTableExcludeStr_Disabled    SlbNewCfgRealServerTableExcludeStr = 2
	SlbNewCfgRealServerTableExcludeStr_Unsupported SlbNewCfgRealServerTableExcludeStr = 2147483647
)

type SlbNewCfgRealServerTableSubmac int32

const (
	SlbNewCfgRealServerTableSubmac_Enabled     SlbNewCfgRealServerTableSubmac = 1
	SlbNewCfgRealServerTableSubmac_Disabled    SlbNewCfgRealServerTableSubmac = 2
	SlbNewCfgRealServerTableSubmac_Unsupported SlbNewCfgRealServerTableSubmac = 2147483647
)

type SlbNewCfgRealServerTableProxy int32

const (
	SlbNewCfgRealServerTableProxy_Enabled     SlbNewCfgRealServerTableProxy = 1
	SlbNewCfgRealServerTableProxy_Disabled    SlbNewCfgRealServerTableProxy = 2
	SlbNewCfgRealServerTableProxy_Unsupported SlbNewCfgRealServerTableProxy = 2147483647
)

type SlbNewCfgRealServerTableLdapwr int32

const (
	SlbNewCfgRealServerTableLdapwr_Enabled     SlbNewCfgRealServerTableLdapwr = 1
	SlbNewCfgRealServerTableLdapwr_Disabled    SlbNewCfgRealServerTableLdapwr = 2
	SlbNewCfgRealServerTableLdapwr_Unsupported SlbNewCfgRealServerTableLdapwr = 2147483647
)

type SlbNewCfgRealServerTableFastHealthCheck int32

const (
	SlbNewCfgRealServerTableFastHealthCheck_Enabled     SlbNewCfgRealServerTableFastHealthCheck = 1
	SlbNewCfgRealServerTableFastHealthCheck_Disabled    SlbNewCfgRealServerTableFastHealthCheck = 2
	SlbNewCfgRealServerTableFastHealthCheck_Unsupported SlbNewCfgRealServerTableFastHealthCheck = 2147483647
)

type SlbNewCfgRealServerTableSubdmac int32

const (
	SlbNewCfgRealServerTableSubdmac_Enabled     SlbNewCfgRealServerTableSubdmac = 1
	SlbNewCfgRealServerTableSubdmac_Disabled    SlbNewCfgRealServerTableSubdmac = 2
	SlbNewCfgRealServerTableSubdmac_Unsupported SlbNewCfgRealServerTableSubdmac = 2147483647
)

type SlbNewCfgRealServerTableOverflow int32

const (
	SlbNewCfgRealServerTableOverflow_Enabled     SlbNewCfgRealServerTableOverflow = 1
	SlbNewCfgRealServerTableOverflow_Disabled    SlbNewCfgRealServerTableOverflow = 2
	SlbNewCfgRealServerTableOverflow_Unsupported SlbNewCfgRealServerTableOverflow = 2147483647
)

type SlbNewCfgRealServerTableBkpPreempt int32

const (
	SlbNewCfgRealServerTableBkpPreempt_Enabled     SlbNewCfgRealServerTableBkpPreempt = 1
	SlbNewCfgRealServerTableBkpPreempt_Disabled    SlbNewCfgRealServerTableBkpPreempt = 2
	SlbNewCfgRealServerTableBkpPreempt_Unsupported SlbNewCfgRealServerTableBkpPreempt = 2147483647
)

type SlbNewCfgRealServerTableIpVer int32

const (
	SlbNewCfgRealServerTableIpVer_Ipv4        SlbNewCfgRealServerTableIpVer = 1
	SlbNewCfgRealServerTableIpVer_Ipv6        SlbNewCfgRealServerTableIpVer = 2
	SlbNewCfgRealServerTableIpVer_Unsupported SlbNewCfgRealServerTableIpVer = 2147483647
)

type SlbNewCfgRealServerTableMode int32

const (
	SlbNewCfgRealServerTableMode_Physical    SlbNewCfgRealServerTableMode = 1
	SlbNewCfgRealServerTableMode_Logical     SlbNewCfgRealServerTableMode = 2
	SlbNewCfgRealServerTableMode_Unsupported SlbNewCfgRealServerTableMode = 2147483647
)

type SlbNewCfgRealServerTableUpdateAllRealServers int32

const (
	SlbNewCfgRealServerTableUpdateAllRealServers_None        SlbNewCfgRealServerTableUpdateAllRealServers = 0
	SlbNewCfgRealServerTableUpdateAllRealServers_Mode        SlbNewCfgRealServerTableUpdateAllRealServers = 1
	SlbNewCfgRealServerTableUpdateAllRealServers_Maxcon      SlbNewCfgRealServerTableUpdateAllRealServers = 2
	SlbNewCfgRealServerTableUpdateAllRealServers_Weight      SlbNewCfgRealServerTableUpdateAllRealServers = 3
	SlbNewCfgRealServerTableUpdateAllRealServers_All         SlbNewCfgRealServerTableUpdateAllRealServers = 4
	SlbNewCfgRealServerTableUpdateAllRealServers_Unsupported SlbNewCfgRealServerTableUpdateAllRealServers = 2147483647
)

type SlbNewCfgRealServerTableProxyIpMode int32

const (
	SlbNewCfgRealServerTableProxyIpMode_Enable      SlbNewCfgRealServerTableProxyIpMode = 0
	SlbNewCfgRealServerTableProxyIpMode_Address     SlbNewCfgRealServerTableProxyIpMode = 2
	SlbNewCfgRealServerTableProxyIpMode_Nwclss      SlbNewCfgRealServerTableProxyIpMode = 3
	SlbNewCfgRealServerTableProxyIpMode_Disable     SlbNewCfgRealServerTableProxyIpMode = 4
	SlbNewCfgRealServerTableProxyIpMode_Unsupported SlbNewCfgRealServerTableProxyIpMode = 2147483647
)

type SlbNewCfgRealServerTableProxyIpPersistency int32

const (
	SlbNewCfgRealServerTableProxyIpPersistency_Disable     SlbNewCfgRealServerTableProxyIpPersistency = 0
	SlbNewCfgRealServerTableProxyIpPersistency_Client      SlbNewCfgRealServerTableProxyIpPersistency = 1
	SlbNewCfgRealServerTableProxyIpPersistency_Host        SlbNewCfgRealServerTableProxyIpPersistency = 2
	SlbNewCfgRealServerTableProxyIpPersistency_Unsupported SlbNewCfgRealServerTableProxyIpPersistency = 2147483647
)

type SlbNewCfgRealServerTableProxyIpNWclassPersistency int32

const (
	SlbNewCfgRealServerTableProxyIpNWclassPersistency_Disable     SlbNewCfgRealServerTableProxyIpNWclassPersistency = 0
	SlbNewCfgRealServerTableProxyIpNWclassPersistency_Client      SlbNewCfgRealServerTableProxyIpNWclassPersistency = 1
	SlbNewCfgRealServerTableProxyIpNWclassPersistency_Unsupported SlbNewCfgRealServerTableProxyIpNWclassPersistency = 2147483647
)

type SlbNewCfgRealServerTableParams struct {
	// The real server number
	Index int32 `json:"Index,omitempty"`
	// IP address of the real server identified by the instance of
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
	// The interval between keep-alive (ping) attempts in number of
	// seconds. Zero means disabling ping attempt.
	PingInterval uint64 `json:"PingInterval,omitempty"`
	// The number of failed attempts to declare this server DOWN.
	FailRetry uint64 `json:"FailRetry,omitempty"`
	// The number of successful attempts to declare a server UP.
	SuccRetry uint64 `json:"SuccRetry,omitempty"`
	// Enable or disable the server and remove the existing sessions using disabled-with-fastage option.
	State SlbNewCfgRealServerTableState `json:"State,omitempty"`
	// By setting the value to delete(2), the entire row is deleted.
	Delete SlbNewCfgRealServerTableDelete `json:"Delete,omitempty"`
	// The server type.  It participates in global server load balancing
	// when it is configured as remote-server.
	Type SlbNewCfgRealServerTableType `json:"Type,omitempty"`
	// The name of the real server.
	Name string `json:"Name,omitempty"`
	// The URL Paths selected for URL load balancing for by the
	// 	 real server.  The selected URL Paths are presented in
	// 	 a bitmap format.
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
	// The URL Path (slbCurCfgUrlLbPathIndex) to be added to the
	// real server. A zero is returned when read.
	AddUrl int32 `json:"AddUrl,omitempty"`
	// The URL Path (slbCurCfgUrlLbPathIndex) to be removed from
	// the real server. A zero is returned when read.
	RemUrl int32 `json:"RemUrl,omitempty"`
	// Enable or disable real server to handle client requests
	// that don't contain a cookie if cookie loadbalance is enabled.
	Cookie SlbNewCfgRealServerTableCookie `json:"Cookie,omitempty"`
	// Enable or disable exclusionary matching string on real server.
	ExcludeStr SlbNewCfgRealServerTableExcludeStr `json:"ExcludeStr,omitempty"`
	// The real server config to enable/disable MAC SA substitution for
	// L4 traffic. If disabled (the default) we will NOT substitute the
	// MAC SA of client-to-server frames.  If enabled, we will substitute
	// the MAC SA.
	Submac SlbNewCfgRealServerTableSubmac `json:"Submac,omitempty"`
	// The real server config to enable/disable client proxy operation.
	Proxy SlbNewCfgRealServerTableProxy `json:"Proxy,omitempty"`
	// The real server config to enable/disable LDAP write server.
	Ldapwr SlbNewCfgRealServerTableLdapwr `json:"Ldapwr,omitempty"`
	// The OID to be sent in the SNMP get request packet.
	Oid string `json:"Oid,omitempty"`
	// The community string to be used in the SNMP get request packet.
	CommString string `json:"CommString,omitempty"`
	// The VLAN to be associated with IDS server.
	Idsvlan uint32 `json:"Idsvlan,omitempty"`
	// The port to be connected to IDS server.
	Idsport int32 `json:"Idsport,omitempty"`
	// The remote real server Global SLB availability.
	Avail uint64 `json:"Avail,omitempty"`
	// The real server config to enable/disable Fast Health Check Operation.
	FastHealthCheck SlbNewCfgRealServerTableFastHealthCheck `json:"FastHealthCheck,omitempty"`
	// The real server config to enable/disable MAC DA substitution for
	// L4 traffic. If disabled, we will NOT substitute the MAC DA of
	// client-to-server frames.  If enabled(default), we will substitute
	// the MAC DA.
	Subdmac SlbNewCfgRealServerTableSubdmac `json:"Subdmac,omitempty"`
	// The real server config to enable/disable Overflow. If enabled(default)
	// allows Backup server to kick in if real server reaches maximum
	// connections.
	Overflow SlbNewCfgRealServerTableOverflow `json:"Overflow,omitempty"`
	// The real server config to enable/disable backup preemption. If enabled
	// (default)allows to preempt the backup server when the primary server
	// comes up. If disabled the backup server will continue to serve the
	// traffic even when primary server comes up.
	BkpPreempt SlbNewCfgRealServerTableBkpPreempt `json:"BkpPreempt,omitempty"`
	// The type of IP address.
	IpVer SlbNewCfgRealServerTableIpVer `json:"IpVer,omitempty"`
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
	// Set the mode of the real server. By default the mode is set to physical.
	Mode SlbNewCfgRealServerTableMode `json:"Mode,omitempty"`
	// Set all the real servers having the same RIP address as this real server
	// 	with the mode and/or maximum connection value (if mode is physical) that is
	// 	set in this real server. When read, none(0) is returned.
	// 	mode   - set only the mode to all real servers with same RIP address.
	// 	maxcon - set only the max connections to all real servers with same RIP address.
	// 	weight - set only the weight to all real servers with same RIP address.
	// 	all    - set mode, max connections and weight to all real servers with same RIP address.
	UpdateAllRealServers SlbNewCfgRealServerTableUpdateAllRealServers `json:"UpdateAllRealServers,omitempty"`
	// Set the real server Proxy IP mode.
	// Changing from address(2) to any other mode will clear the configured IPv4/IPv6 address,prefix and persistancy.
	// Changing from  nwclass(3) to any other mode will clear the configured NWclass and NWpersistancy.
	ProxyIpMode SlbNewCfgRealServerTableProxyIpMode `json:"ProxyIpMode,omitempty"`
	// This object ID allows configuration of real server proxy IP v4 address .
	// 		The IP version for addr must be the same as the real server IP version.
	// 		This object ID can be set only if slbNewCfgRealServerIpVer is set to ipv4 & slbNewCfgRealServerProxyIpMode is address else return failure.
	// 		Returns 0 if IP version is IPv6 or slbNewCfgRealServerProxyIpMode is not set to address.
	// 		When a subnet is configured, user has the ability to select PIP persistency mode.
	// 		Persistency is relevant only if the PIP is configured as subnet. If PIP is not configured or
	// 		not subnet, the persistency configuration value is disable.
	ProxyIpAddr string `json:"ProxyIpAddr,omitempty"`
	// This object ID allows configuration of real server proxy IP Mask.
	// 	        The IP version for addr must be the same as the real server IP version.
	// 	 This object ID can be set only if slbNewCfgRealServerIpVer is set to ipv4  & slbNewCfgRealServerProxyIpMode is address else return failure.
	// 		 Returns 0 if IP version is IPv6 or slbNewCfgRealServerProxyIpMode is not set to address.
	// 		 When a subnet is configured user has the ability to select PIP persistency mode.
	// 	 	 Persistency is relevant only if the PIP is configured as subnet. If PIP is not configured or
	// 		 not subnet, the persistency configuration value is disable.
	ProxyIpMask string `json:"ProxyIpMask,omitempty"`
	// This object ID allows configuration of real server proxy IPv6 address .
	// 		The IP version for addr must be the same as the real server IP version.
	// 		Returns emply if IP version is IPv4 or slbNewCfgRealServerProxyIpMode is not set to address.
	// 		This object ID can be set only if slbNewCfgRealServerIpVer is set to ipv6  & slbNewCfgRealServerProxyIpMode is address else return failure.
	// 		When a subnet is configured, user has the ability to select PIP persistency mode.
	// 		Persistency is relevant only if the PIP is configured as subnet. If PIP is not configured or
	// 		not subnet, the persistency configuration value is disable.
	// 		Address should be 4-byte hexadecimal colon notation.
	// 		Valid IPv6 address should be in any of the following forms
	// xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx or
	// xxxx::xxxx:xxxx:xxxx:xxxx or ::xxxx.
	ProxyIpv6Addr string `json:"ProxyIpv6Addr,omitempty"`
	// This object ID allows configuration of real server proxy IPv6 Mask.
	// 		The IP version for addr must be the same as the real server IP version.
	// 		This object ID can be set only if slbNewCfgRealServerIpVer is set to ipv6  & slbNewCfgRealServerProxyIpMode is address else return failure.
	// 		Returns 0 if IP version is IPv4 or slbNewCfgRealServerProxyIpMode is set other than address.
	// 		When a subnet is configured user has the ability to select PIP persistency mode.
	// 		Persistency is relevant only if the PIP is configured as subnet. If PIP is not configured or
	// 		not subnet, the persistency configuration value is disable.
	ProxyIpv6Prefix uint32 `json:"ProxyIpv6Prefix,omitempty"`
	// When a subnet is configured user has the ability to select PIP persistency mode.
	// 		This object ID can be set only if slbNewCfgRealServerProxyIpMode is address else return failure.
	// 		Persistency is relevant only if a PIP class is configured. If PIP is not configured the
	// 		persistency configuration is disable.
	ProxyIpPersistency SlbNewCfgRealServerTableProxyIpPersistency `json:"ProxyIpPersistency,omitempty"`
	// This  object ID allows configuration of real server proxy IP IPv4 or IPv6 Network Class as PIP.
	// 		This object ID can be set only if slbNewCfgRealServerProxyIpMode is nwclss else return failure.
	// 		Persistency is relevant only if a PIP class is configured. If PIP is not configured the
	// 		persistency configuration is disable.
	// Returns empty string when slbNewCfgRealServerProxyIpMode is not set to nwclss(3).
	ProxyIpNWclass string `json:"ProxyIpNWclass,omitempty"`
	// This object ID allows configuration of real server Network Class PIP persistency mode.
	// 		This object ID can be set only if slbNewCfgRealServerProxyIpMode is nwclss else return failure.
	// 		Persistency is relevant only if a PIP class is configured. If PIP is not configured the
	// 		persistency configuration is disable.
	// Returns 0 when slbNewCfgRealServerProxyIpMode is not set to nwclss(3).
	ProxyIpNWclassPersistency SlbNewCfgRealServerTableProxyIpNWclassPersistency `json:"ProxyIpNWclassPersistency,omitempty"`
	// The Advanced HC ID.
	HealthID string `json:"HealthID,omitempty"`
}

func (p SlbNewCfgRealServerTableParams) iMABean() {}
