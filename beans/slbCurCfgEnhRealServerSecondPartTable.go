package beans

import (
	"fmt"
	"reflect"
)

// SlbCurCfgEnhRealServerSecondPartTable The second table of Real Server configuration in the current_config.
// Note:This mib is not supported on VX instance of Virtualization.
type SlbCurCfgEnhRealServerSecondPartTable struct {
	// The real server number
	SlbCurCfgEnhRealServerSecondPartIndex string
	Params                                *SlbCurCfgEnhRealServerSecondPartTableParams
}

func NewSlbCurCfgEnhRealServerSecondPartTableList() *SlbCurCfgEnhRealServerSecondPartTable {
	return &SlbCurCfgEnhRealServerSecondPartTable{}
}

func NewSlbCurCfgEnhRealServerSecondPartTable(
	slbCurCfgEnhRealServerSecondPartIndex string,
	params *SlbCurCfgEnhRealServerSecondPartTableParams,
) *SlbCurCfgEnhRealServerSecondPartTable {
	return &SlbCurCfgEnhRealServerSecondPartTable{
		SlbCurCfgEnhRealServerSecondPartIndex: slbCurCfgEnhRealServerSecondPartIndex,
		Params:                                params,
	}
}

func (c *SlbCurCfgEnhRealServerSecondPartTable) Name() string {
	return "SlbCurCfgEnhRealServerSecondPartTable"
}

func (c *SlbCurCfgEnhRealServerSecondPartTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbCurCfgEnhRealServerSecondPartTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbCurCfgEnhRealServerSecondPartTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbCurCfgEnhRealServerSecondPartIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbCurCfgEnhRealServerSecondPartIndex)
}

type SlbCurCfgEnhRealServerSecondPartTableProxy int32

const (
	SlbCurCfgEnhRealServerSecondPartTableProxy_Enabled     SlbCurCfgEnhRealServerSecondPartTableProxy = 1
	SlbCurCfgEnhRealServerSecondPartTableProxy_Disabled    SlbCurCfgEnhRealServerSecondPartTableProxy = 2
	SlbCurCfgEnhRealServerSecondPartTableProxy_Unsupported SlbCurCfgEnhRealServerSecondPartTableProxy = 2147483647
)

type SlbCurCfgEnhRealServerSecondPartTableLdapwr int32

const (
	SlbCurCfgEnhRealServerSecondPartTableLdapwr_Enabled     SlbCurCfgEnhRealServerSecondPartTableLdapwr = 1
	SlbCurCfgEnhRealServerSecondPartTableLdapwr_Disabled    SlbCurCfgEnhRealServerSecondPartTableLdapwr = 2
	SlbCurCfgEnhRealServerSecondPartTableLdapwr_Unsupported SlbCurCfgEnhRealServerSecondPartTableLdapwr = 2147483647
)

type SlbCurCfgEnhRealServerSecondPartTableFastHealthCheck int32

const (
	SlbCurCfgEnhRealServerSecondPartTableFastHealthCheck_Enabled     SlbCurCfgEnhRealServerSecondPartTableFastHealthCheck = 1
	SlbCurCfgEnhRealServerSecondPartTableFastHealthCheck_Disabled    SlbCurCfgEnhRealServerSecondPartTableFastHealthCheck = 2
	SlbCurCfgEnhRealServerSecondPartTableFastHealthCheck_Unsupported SlbCurCfgEnhRealServerSecondPartTableFastHealthCheck = 2147483647
)

type SlbCurCfgEnhRealServerSecondPartTableSubdmac int32

const (
	SlbCurCfgEnhRealServerSecondPartTableSubdmac_Enabled     SlbCurCfgEnhRealServerSecondPartTableSubdmac = 1
	SlbCurCfgEnhRealServerSecondPartTableSubdmac_Disabled    SlbCurCfgEnhRealServerSecondPartTableSubdmac = 2
	SlbCurCfgEnhRealServerSecondPartTableSubdmac_Unsupported SlbCurCfgEnhRealServerSecondPartTableSubdmac = 2147483647
)

type SlbCurCfgEnhRealServerSecondPartTableOverflow int32

const (
	SlbCurCfgEnhRealServerSecondPartTableOverflow_Enabled     SlbCurCfgEnhRealServerSecondPartTableOverflow = 1
	SlbCurCfgEnhRealServerSecondPartTableOverflow_Disabled    SlbCurCfgEnhRealServerSecondPartTableOverflow = 2
	SlbCurCfgEnhRealServerSecondPartTableOverflow_Unsupported SlbCurCfgEnhRealServerSecondPartTableOverflow = 2147483647
)

type SlbCurCfgEnhRealServerSecondPartTableBkpPreempt int32

const (
	SlbCurCfgEnhRealServerSecondPartTableBkpPreempt_Enabled     SlbCurCfgEnhRealServerSecondPartTableBkpPreempt = 1
	SlbCurCfgEnhRealServerSecondPartTableBkpPreempt_Disabled    SlbCurCfgEnhRealServerSecondPartTableBkpPreempt = 2
	SlbCurCfgEnhRealServerSecondPartTableBkpPreempt_Unsupported SlbCurCfgEnhRealServerSecondPartTableBkpPreempt = 2147483647
)

type SlbCurCfgEnhRealServerSecondPartTableMode int32

const (
	SlbCurCfgEnhRealServerSecondPartTableMode_Physical    SlbCurCfgEnhRealServerSecondPartTableMode = 1
	SlbCurCfgEnhRealServerSecondPartTableMode_Logical     SlbCurCfgEnhRealServerSecondPartTableMode = 2
	SlbCurCfgEnhRealServerSecondPartTableMode_Unsupported SlbCurCfgEnhRealServerSecondPartTableMode = 2147483647
)

type SlbCurCfgEnhRealServerSecondPartTableProxyIpMode int32

const (
	SlbCurCfgEnhRealServerSecondPartTableProxyIpMode_Enable      SlbCurCfgEnhRealServerSecondPartTableProxyIpMode = 0
	SlbCurCfgEnhRealServerSecondPartTableProxyIpMode_Address     SlbCurCfgEnhRealServerSecondPartTableProxyIpMode = 2
	SlbCurCfgEnhRealServerSecondPartTableProxyIpMode_Nwclss      SlbCurCfgEnhRealServerSecondPartTableProxyIpMode = 3
	SlbCurCfgEnhRealServerSecondPartTableProxyIpMode_Disable     SlbCurCfgEnhRealServerSecondPartTableProxyIpMode = 4
	SlbCurCfgEnhRealServerSecondPartTableProxyIpMode_Unsupported SlbCurCfgEnhRealServerSecondPartTableProxyIpMode = 2147483647
)

type SlbCurCfgEnhRealServerSecondPartTableProxyIpPersistency int32

const (
	SlbCurCfgEnhRealServerSecondPartTableProxyIpPersistency_Disable     SlbCurCfgEnhRealServerSecondPartTableProxyIpPersistency = 0
	SlbCurCfgEnhRealServerSecondPartTableProxyIpPersistency_Client      SlbCurCfgEnhRealServerSecondPartTableProxyIpPersistency = 1
	SlbCurCfgEnhRealServerSecondPartTableProxyIpPersistency_Host        SlbCurCfgEnhRealServerSecondPartTableProxyIpPersistency = 2
	SlbCurCfgEnhRealServerSecondPartTableProxyIpPersistency_Unsupported SlbCurCfgEnhRealServerSecondPartTableProxyIpPersistency = 2147483647
)

type SlbCurCfgEnhRealServerSecondPartTableProxyIpNWclassPersistency int32

const (
	SlbCurCfgEnhRealServerSecondPartTableProxyIpNWclassPersistency_Disable     SlbCurCfgEnhRealServerSecondPartTableProxyIpNWclassPersistency = 0
	SlbCurCfgEnhRealServerSecondPartTableProxyIpNWclassPersistency_Client      SlbCurCfgEnhRealServerSecondPartTableProxyIpNWclassPersistency = 1
	SlbCurCfgEnhRealServerSecondPartTableProxyIpNWclassPersistency_Unsupported SlbCurCfgEnhRealServerSecondPartTableProxyIpNWclassPersistency = 2147483647
)

type SlbCurCfgEnhRealServerSecondPartTableParams struct {
	// The real server number
	Index string `json:"Index,omitempty"`
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
	// The real server config to enable/disable client proxy operation.
	Proxy SlbCurCfgEnhRealServerSecondPartTableProxy `json:"Proxy,omitempty"`
	// The real server config to enable/disable LDAP write server.
	Ldapwr SlbCurCfgEnhRealServerSecondPartTableLdapwr `json:"Ldapwr,omitempty"`
	// The VLAN to be associated with the IDS server.
	Idsvlan uint32 `json:"Idsvlan,omitempty"`
	// The remote real server Global SLB availability.
	Avail uint64 `json:"Avail,omitempty"`
	// The real server config to enable/disable Fast Health Check Operation.
	FastHealthCheck SlbCurCfgEnhRealServerSecondPartTableFastHealthCheck `json:"FastHealthCheck,omitempty"`
	// The real server config to enable/disable MAC DA substitution for
	// L4 traffic. If disabled, we will NOT substitute the MAC DA of
	// client-to-server frames.  If enabled(default), we will substitute
	// the MAC DA.
	Subdmac SlbCurCfgEnhRealServerSecondPartTableSubdmac `json:"Subdmac,omitempty"`
	// The real server config to enable/disable Overflow. If enabled(default)
	// allows Backup server to kick in if real server reaches maximum
	// connections.
	Overflow SlbCurCfgEnhRealServerSecondPartTableOverflow `json:"Overflow,omitempty"`
	// The real server config to enable/disable backup preemption. If enabled
	// (default)allows to preempt the backup server when the primary server
	// comes up. If disabled the backup server will continue to serve the
	// traffic even when primary server comes up.
	BkpPreempt SlbCurCfgEnhRealServerSecondPartTableBkpPreempt `json:"BkpPreempt,omitempty"`
	// The mode of the real server. By default it is set to physical.
	Mode SlbCurCfgEnhRealServerSecondPartTableMode `json:"Mode,omitempty"`
	// the real server Proxy IP mode
	ProxyIpMode SlbCurCfgEnhRealServerSecondPartTableProxyIpMode `json:"ProxyIpMode,omitempty"`
	// This object ID shows current configuration of real server proxy IP address .
	// Returns 0 when slbCurCfgEnhRealServerProxyIpMode is not set to address.
	// Returns 0 if IP version is IPv6.
	ProxyIpAddr string `json:"ProxyIpAddr,omitempty"`
	// This object ID shows current configuration of real server proxy IP subnet mask.
	// Returns 0 when slbCurCfgEnhRealServerProxyIpMode is not set to address.
	// Returns 0 if IP version is IPv6.
	ProxyIpMask string `json:"ProxyIpMask,omitempty"`
	// This object ID shows current configuration of real server proxy IPv6 address .
	// The IP version for addr must be the same as the real server IP version.
	// Returns 0 when slbCurCfgEnhRealServerProxyIpMode is not set to address.
	// Returns 0 if IP version is IPv4.
	// When a subnet is configured, user has the ability to select PIP persistency mode.
	// Persistency is relevant only if the PIP is configured as subnet. If PIP is not configured or
	// not subnet, the persistency configuration value is disable.
	ProxyIpv6Addr string `json:"ProxyIpv6Addr,omitempty"`
	// This object ID allows configuration of real server proxy IPv6 Mask.
	// The IP version for addr must be the same as the real server IP version.
	// Returns 0 when slbCurCfgEnhRealServerProxyIpMode is not set to address.
	// Returns 0 if IP version is IPv6.
	// When a subnet is configured user has the ability to select PIP persistency mode.
	// Persistency is relevant only if the PIP is configured as subnet. If PIP is not configured or
	// not subnet, the persistency configuration value is disable.
	ProxyIpv6Prefix uint32 `json:"ProxyIpv6Prefix,omitempty"`
	// This object ID shows current configuration of real server proxy IP persistency mode.
	// Persistency is relevant only if a PIP class is configured. If PIP is not configured the
	// persistency configuration is disable.
	ProxyIpPersistency SlbCurCfgEnhRealServerSecondPartTableProxyIpPersistency `json:"ProxyIpPersistency,omitempty"`
	// This  object ID allows configuration of real server proxy IP IPv4 or IPv6 Network Class as PIP.
	// Network Class PIP Persistency is relevant only if a PIP class is configured. If PIP is not configured the
	// persistency configuration is disable.
	// Returns 0 when slbCurCfgEnhRealServerProxyIpMode is not set to nwclss(3).
	ProxyIpNWclass string `json:"ProxyIpNWclass,omitempty"`
	// This object ID allows configuration of real server Network Class PIP persistency mode.
	// Persistency is relevant only if a PIP class is configured. If PIP is not configured the
	// persistency configuration is disable.
	// Returns 0 when slbCurCfgEnhRealServerProxyIpMode is not set to nwclss(3).
	ProxyIpNWclassPersistency SlbCurCfgEnhRealServerSecondPartTableProxyIpNWclassPersistency `json:"ProxyIpNWclassPersistency,omitempty"`
	// The ingress VLAN to be associated with the IDS server.
	Ingvlan uint32 `json:"Ingvlan,omitempty"`
}

func (p SlbCurCfgEnhRealServerSecondPartTableParams) iMABean() {}
