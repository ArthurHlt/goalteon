package beans

import (
	"fmt"
	"reflect"
)

// SlbNewCfgEnhRealServerSecondPartTable The second table of Real Server configuration in the current_config.
// Note:This mib is not supported on VX instance of Virtualization.
type SlbNewCfgEnhRealServerSecondPartTable struct {
	// The real server number
	SlbNewCfgEnhRealServerSecondPartIndex string
	Params                                *SlbNewCfgEnhRealServerSecondPartTableParams
}

func NewSlbNewCfgEnhRealServerSecondPartTableList() *SlbNewCfgEnhRealServerSecondPartTable {
	return &SlbNewCfgEnhRealServerSecondPartTable{}
}

func NewSlbNewCfgEnhRealServerSecondPartTable(
	slbNewCfgEnhRealServerSecondPartIndex string,
	params *SlbNewCfgEnhRealServerSecondPartTableParams,
) *SlbNewCfgEnhRealServerSecondPartTable {
	return &SlbNewCfgEnhRealServerSecondPartTable{
		SlbNewCfgEnhRealServerSecondPartIndex: slbNewCfgEnhRealServerSecondPartIndex,
		Params:                                params,
	}
}

func (c *SlbNewCfgEnhRealServerSecondPartTable) Name() string {
	return "SlbNewCfgEnhRealServerSecondPartTable"
}

func (c *SlbNewCfgEnhRealServerSecondPartTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbNewCfgEnhRealServerSecondPartTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbNewCfgEnhRealServerSecondPartTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbNewCfgEnhRealServerSecondPartIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbNewCfgEnhRealServerSecondPartIndex)
}

type SlbNewCfgEnhRealServerSecondPartTableProxy int32

const (
	SlbNewCfgEnhRealServerSecondPartTableProxy_Enabled     SlbNewCfgEnhRealServerSecondPartTableProxy = 1
	SlbNewCfgEnhRealServerSecondPartTableProxy_Disabled    SlbNewCfgEnhRealServerSecondPartTableProxy = 2
	SlbNewCfgEnhRealServerSecondPartTableProxy_Unsupported SlbNewCfgEnhRealServerSecondPartTableProxy = 2147483647
)

type SlbNewCfgEnhRealServerSecondPartTableLdapwr int32

const (
	SlbNewCfgEnhRealServerSecondPartTableLdapwr_Enabled     SlbNewCfgEnhRealServerSecondPartTableLdapwr = 1
	SlbNewCfgEnhRealServerSecondPartTableLdapwr_Disabled    SlbNewCfgEnhRealServerSecondPartTableLdapwr = 2
	SlbNewCfgEnhRealServerSecondPartTableLdapwr_Unsupported SlbNewCfgEnhRealServerSecondPartTableLdapwr = 2147483647
)

type SlbNewCfgEnhRealServerSecondPartTableFastHealthCheck int32

const (
	SlbNewCfgEnhRealServerSecondPartTableFastHealthCheck_Enabled     SlbNewCfgEnhRealServerSecondPartTableFastHealthCheck = 1
	SlbNewCfgEnhRealServerSecondPartTableFastHealthCheck_Disabled    SlbNewCfgEnhRealServerSecondPartTableFastHealthCheck = 2
	SlbNewCfgEnhRealServerSecondPartTableFastHealthCheck_Unsupported SlbNewCfgEnhRealServerSecondPartTableFastHealthCheck = 2147483647
)

type SlbNewCfgEnhRealServerSecondPartTableSubdmac int32

const (
	SlbNewCfgEnhRealServerSecondPartTableSubdmac_Enabled     SlbNewCfgEnhRealServerSecondPartTableSubdmac = 1
	SlbNewCfgEnhRealServerSecondPartTableSubdmac_Disabled    SlbNewCfgEnhRealServerSecondPartTableSubdmac = 2
	SlbNewCfgEnhRealServerSecondPartTableSubdmac_Unsupported SlbNewCfgEnhRealServerSecondPartTableSubdmac = 2147483647
)

type SlbNewCfgEnhRealServerSecondPartTableOverflow int32

const (
	SlbNewCfgEnhRealServerSecondPartTableOverflow_Enabled     SlbNewCfgEnhRealServerSecondPartTableOverflow = 1
	SlbNewCfgEnhRealServerSecondPartTableOverflow_Disabled    SlbNewCfgEnhRealServerSecondPartTableOverflow = 2
	SlbNewCfgEnhRealServerSecondPartTableOverflow_Unsupported SlbNewCfgEnhRealServerSecondPartTableOverflow = 2147483647
)

type SlbNewCfgEnhRealServerSecondPartTableBkpPreempt int32

const (
	SlbNewCfgEnhRealServerSecondPartTableBkpPreempt_Enabled     SlbNewCfgEnhRealServerSecondPartTableBkpPreempt = 1
	SlbNewCfgEnhRealServerSecondPartTableBkpPreempt_Disabled    SlbNewCfgEnhRealServerSecondPartTableBkpPreempt = 2
	SlbNewCfgEnhRealServerSecondPartTableBkpPreempt_Unsupported SlbNewCfgEnhRealServerSecondPartTableBkpPreempt = 2147483647
)

type SlbNewCfgEnhRealServerSecondPartTableMode int32

const (
	SlbNewCfgEnhRealServerSecondPartTableMode_Physical    SlbNewCfgEnhRealServerSecondPartTableMode = 1
	SlbNewCfgEnhRealServerSecondPartTableMode_Logical     SlbNewCfgEnhRealServerSecondPartTableMode = 2
	SlbNewCfgEnhRealServerSecondPartTableMode_Unsupported SlbNewCfgEnhRealServerSecondPartTableMode = 2147483647
)

type SlbNewCfgEnhRealServerSecondPartTableUpdateAllRealServers int32

const (
	SlbNewCfgEnhRealServerSecondPartTableUpdateAllRealServers_None        SlbNewCfgEnhRealServerSecondPartTableUpdateAllRealServers = 0
	SlbNewCfgEnhRealServerSecondPartTableUpdateAllRealServers_Mode        SlbNewCfgEnhRealServerSecondPartTableUpdateAllRealServers = 1
	SlbNewCfgEnhRealServerSecondPartTableUpdateAllRealServers_Maxcon      SlbNewCfgEnhRealServerSecondPartTableUpdateAllRealServers = 2
	SlbNewCfgEnhRealServerSecondPartTableUpdateAllRealServers_Weight      SlbNewCfgEnhRealServerSecondPartTableUpdateAllRealServers = 3
	SlbNewCfgEnhRealServerSecondPartTableUpdateAllRealServers_All         SlbNewCfgEnhRealServerSecondPartTableUpdateAllRealServers = 4
	SlbNewCfgEnhRealServerSecondPartTableUpdateAllRealServers_Unsupported SlbNewCfgEnhRealServerSecondPartTableUpdateAllRealServers = 2147483647
)

type SlbNewCfgEnhRealServerSecondPartTableProxyIpMode int32

const (
	SlbNewCfgEnhRealServerSecondPartTableProxyIpMode_Enable      SlbNewCfgEnhRealServerSecondPartTableProxyIpMode = 0
	SlbNewCfgEnhRealServerSecondPartTableProxyIpMode_Address     SlbNewCfgEnhRealServerSecondPartTableProxyIpMode = 2
	SlbNewCfgEnhRealServerSecondPartTableProxyIpMode_Nwclss      SlbNewCfgEnhRealServerSecondPartTableProxyIpMode = 3
	SlbNewCfgEnhRealServerSecondPartTableProxyIpMode_Disable     SlbNewCfgEnhRealServerSecondPartTableProxyIpMode = 4
	SlbNewCfgEnhRealServerSecondPartTableProxyIpMode_Unsupported SlbNewCfgEnhRealServerSecondPartTableProxyIpMode = 2147483647
)

type SlbNewCfgEnhRealServerSecondPartTableProxyIpPersistency int32

const (
	SlbNewCfgEnhRealServerSecondPartTableProxyIpPersistency_Disable     SlbNewCfgEnhRealServerSecondPartTableProxyIpPersistency = 0
	SlbNewCfgEnhRealServerSecondPartTableProxyIpPersistency_Client      SlbNewCfgEnhRealServerSecondPartTableProxyIpPersistency = 1
	SlbNewCfgEnhRealServerSecondPartTableProxyIpPersistency_Host        SlbNewCfgEnhRealServerSecondPartTableProxyIpPersistency = 2
	SlbNewCfgEnhRealServerSecondPartTableProxyIpPersistency_Unsupported SlbNewCfgEnhRealServerSecondPartTableProxyIpPersistency = 2147483647
)

type SlbNewCfgEnhRealServerSecondPartTableProxyIpNWclassPersistency int32

const (
	SlbNewCfgEnhRealServerSecondPartTableProxyIpNWclassPersistency_Disable     SlbNewCfgEnhRealServerSecondPartTableProxyIpNWclassPersistency = 0
	SlbNewCfgEnhRealServerSecondPartTableProxyIpNWclassPersistency_Client      SlbNewCfgEnhRealServerSecondPartTableProxyIpNWclassPersistency = 1
	SlbNewCfgEnhRealServerSecondPartTableProxyIpNWclassPersistency_Unsupported SlbNewCfgEnhRealServerSecondPartTableProxyIpNWclassPersistency = 2147483647
)

type SlbNewCfgEnhRealServerSecondPartTableParams struct {
	// The real server number
	Index string `json:"Index,omitempty"`
	// The URL Paths selected for URL load balancing for by the
	// real server.  The selected URL Paths are presented in
	// a bitmap format.
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
	// The real server config to enable/disable client proxy operation.
	Proxy SlbNewCfgEnhRealServerSecondPartTableProxy `json:"Proxy,omitempty"`
	// The real server config to enable/disable LDAP write server.
	Ldapwr SlbNewCfgEnhRealServerSecondPartTableLdapwr `json:"Ldapwr,omitempty"`
	// The VLAN to be associated with IDS server.
	Idsvlan uint32 `json:"Idsvlan,omitempty"`
	// The remote real server Global SLB availability.
	Avail uint64 `json:"Avail,omitempty"`
	// The real server config to enable/disable Fast Health Check Operation.
	FastHealthCheck SlbNewCfgEnhRealServerSecondPartTableFastHealthCheck `json:"FastHealthCheck,omitempty"`
	// The real server config to enable/disable MAC DA substitution for
	// L4 traffic. If disabled, we will NOT substitute the MAC DA of
	// client-to-server frames.  If enabled(default), we will substitute
	// the MAC DA.
	Subdmac SlbNewCfgEnhRealServerSecondPartTableSubdmac `json:"Subdmac,omitempty"`
	// The real server config to enable/disable Overflow. If enabled(default)
	// allows Backup server to kick in if real server reaches maximum
	// connections.
	Overflow SlbNewCfgEnhRealServerSecondPartTableOverflow `json:"Overflow,omitempty"`
	// The real server config to enable/disable backup preemption. If enabled
	// (default)allows to preempt the backup server when the primary server
	// comes up. If disabled the backup server will continue to serve the
	// traffic even when primary server comes up.
	BkpPreempt SlbNewCfgEnhRealServerSecondPartTableBkpPreempt `json:"BkpPreempt,omitempty"`
	// Set the mode of the real server. By default the mode is set to physical.
	Mode SlbNewCfgEnhRealServerSecondPartTableMode `json:"Mode,omitempty"`
	// Set all the real servers having the same RIP address as this real server
	// with the mode and/or maximum connection value (if mode is physical) that is
	// set in this real server. When read, none(0) is returned.
	// mode   - set only the mode to all real servers with same RIP address.
	// maxcon - set only the max connections to all real servers with same RIP address.
	// weight - set only the weight to all real servers with same RIP address.
	// all    - set mode, max connections and weight to all real servers with same RIP address.
	UpdateAllRealServers SlbNewCfgEnhRealServerSecondPartTableUpdateAllRealServers `json:"UpdateAllRealServers,omitempty"`
	// Set the real server Proxy IP mode.
	// Changing from address(2) to any other mode will clear the configured IPv4/IPv6 address,prefix and persistancy.
	// Changing from  nwclass(3) to any other mode will clear the configured NWclass and NWpersistancy.
	ProxyIpMode SlbNewCfgEnhRealServerSecondPartTableProxyIpMode `json:"ProxyIpMode,omitempty"`
	// This object ID allows configuration of real server proxy IP v4 address .
	// The IP version for addr must be the same as the real server IP version.
	// This object ID can be set only if slbNewCfgEnhRealServerIpVer is set to ipv4 &
	// slbNewCfgEnhRealServerProxyIpMode is address else return failure.
	// Returns 0 if IP version is IPv6 or slbNewCfgEnhRealServerProxyIpMode is not set to address.
	// When a subnet is configured, user has the ability to select PIP persistency mode.
	// Persistency is relevant only if the PIP is configured as subnet. If PIP is not configured or
	// not subnet, the persistency configuration value is disable.
	ProxyIpAddr string `json:"ProxyIpAddr,omitempty"`
	// This object ID allows configuration of real server proxy IP Mask.
	// The IP version for addr must be the same as the real server IP version.
	// This object ID can be set only if slbNewCfgEnhRealServerIpVer is set to ipv4
	// & slbNewCfgEnhRealServerProxyIpMode is address else return failure.
	// Returns 0 if IP version is IPv6 or slbNewCfgEnhRealServerProxyIpMode is not set to address.
	// When a subnet is configured user has the ability to select PIP persistency mode.
	// Persistency is relevant only if the PIP is configured as subnet. If PIP is not configured or
	// not subnet, the persistency configuration value is disable.
	ProxyIpMask string `json:"ProxyIpMask,omitempty"`
	// This object ID allows configuration of real server proxy IPv6 address .
	// The IP version for addr must be the same as the real server IP version.
	// Returns emply if IP version is IPv4 or slbNewCfgEnhRealServerProxyIpMode is not set to address.
	// This object ID can be set only if slbNewCfgEnhRealServerIpVer is set to ipv6
	// & slbNewCfgEnhRealServerProxyIpMode is address else return failure.
	// When a subnet is configured, user has the ability to select PIP persistency mode.
	// Persistency is relevant only if the PIP is configured as subnet. If PIP is not configured or
	// not subnet, the persistency configuration value is disable.
	// Address should be 4-byte hexadecimal colon notation.
	// Valid IPv6 address should be in any of the following forms
	// xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx or
	// xxxx::xxxx:xxxx:xxxx:xxxx or ::xxxx.
	ProxyIpv6Addr string `json:"ProxyIpv6Addr,omitempty"`
	// This object ID allows configuration of real server proxy IPv6 Mask.
	// The IP version for addr must be the same as the real server IP version.
	// This object ID can be set only if slbNewCfgEnhRealServerIpVer is set to ipv6
	// & slbNewCfgEnhRealServerProxyIpMode is address else return failure.
	// Returns 0 if IP version is IPv4 or slbNewCfgEnhRealServerProxyIpMode is set other than address.
	// When a subnet is configured user has the ability to select PIP persistency mode.
	// Persistency is relevant only if the PIP is configured as subnet. If PIP is not configured or
	// not subnet, the persistency configuration value is disable.
	ProxyIpv6Prefix uint32 `json:"ProxyIpv6Prefix,omitempty"`
	// When a subnet is configured user has the ability to select PIP persistency mode.
	// This object ID can be set only if slbNewCfgEnhRealServerProxyIpMode is address else return failure.
	// Persistency is relevant only if a PIP class is configured. If PIP is not configured the
	// persistency configuration is disable.
	ProxyIpPersistency SlbNewCfgEnhRealServerSecondPartTableProxyIpPersistency `json:"ProxyIpPersistency,omitempty"`
	// This  object ID allows configuration of real server proxy IP IPv4 or IPv6 Network Class as PIP.
	// This object ID can be set only if slbNewCfgEnhRealServerProxyIpMode is nwclss else return failure.
	// Persistency is relevant only if a PIP class is configured. If PIP is not configured the
	// persistency configuration is disable.
	// Returns empty string when slbNewCfgEnhRealServerProxyIpMode is not set to nwclss(3).
	ProxyIpNWclass string `json:"ProxyIpNWclass,omitempty"`
	// This object ID allows configuration of real server Network Class PIP persistency mode.
	// This object ID can be set only if slbNewCfgEnhRealServerProxyIpMode is nwclss else return failure.
	// Persistency is relevant only if a PIP class is configured. If PIP is not configured the
	// persistency configuration is disable.
	// Returns 0 when slbNewCfgEnhRealServerProxyIpMode is not set to nwclss(3).
	ProxyIpNWclassPersistency SlbNewCfgEnhRealServerSecondPartTableProxyIpNWclassPersistency `json:"ProxyIpNWclassPersistency,omitempty"`
	// The ingress VLAN to be associated with IDS server.
	Ingvlan uint32 `json:"Ingvlan,omitempty"`
}

func (p SlbNewCfgEnhRealServerSecondPartTableParams) iMABean() {}
