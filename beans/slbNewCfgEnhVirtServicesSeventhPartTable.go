package beans

import (
	"fmt"
	"reflect"
)

// SlbNewCfgEnhVirtServicesSeventhPartTable The table of virtual Services.
// Note:This mib is not supported for VX instance of Virtualization.
type SlbNewCfgEnhVirtServicesSeventhPartTable struct {
	// The number of the virtual server.
	SlbNewCfgEnhVirtServSeventhPartIndex string
	// The service index. This has no external meaning
	SlbNewCfgEnhVirtServiceSeventhPartIndex int32
	Params                                  *SlbNewCfgEnhVirtServicesSeventhPartTableParams
}

func NewSlbNewCfgEnhVirtServicesSeventhPartTableList() *SlbNewCfgEnhVirtServicesSeventhPartTable {
	return &SlbNewCfgEnhVirtServicesSeventhPartTable{}
}

func NewSlbNewCfgEnhVirtServicesSeventhPartTable(
	slbNewCfgEnhVirtServSeventhPartIndex string,
	slbNewCfgEnhVirtServiceSeventhPartIndex int32,
	params *SlbNewCfgEnhVirtServicesSeventhPartTableParams,
) *SlbNewCfgEnhVirtServicesSeventhPartTable {
	return &SlbNewCfgEnhVirtServicesSeventhPartTable{
		SlbNewCfgEnhVirtServSeventhPartIndex:    slbNewCfgEnhVirtServSeventhPartIndex,
		SlbNewCfgEnhVirtServiceSeventhPartIndex: slbNewCfgEnhVirtServiceSeventhPartIndex,
		Params:                                  params,
	}
}

func (c *SlbNewCfgEnhVirtServicesSeventhPartTable) Name() string {
	return "SlbNewCfgEnhVirtServicesSeventhPartTable"
}

func (c *SlbNewCfgEnhVirtServicesSeventhPartTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbNewCfgEnhVirtServicesSeventhPartTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbNewCfgEnhVirtServicesSeventhPartTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbNewCfgEnhVirtServSeventhPartIndex).IsZero() &&
		reflect.ValueOf(c.SlbNewCfgEnhVirtServiceSeventhPartIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbNewCfgEnhVirtServSeventhPartIndex) + "/" + fmt.Sprint(c.SlbNewCfgEnhVirtServiceSeventhPartIndex)
}

type SlbNewCfgEnhVirtServicesSeventhPartTableSessionMirror int32

const (
	SlbNewCfgEnhVirtServicesSeventhPartTableSessionMirror_Enabled     SlbNewCfgEnhVirtServicesSeventhPartTableSessionMirror = 1
	SlbNewCfgEnhVirtServicesSeventhPartTableSessionMirror_Disabled    SlbNewCfgEnhVirtServicesSeventhPartTableSessionMirror = 2
	SlbNewCfgEnhVirtServicesSeventhPartTableSessionMirror_Unsupported SlbNewCfgEnhVirtServicesSeventhPartTableSessionMirror = 2147483647
)

type SlbNewCfgEnhVirtServicesSeventhPartTableSoftGrid int32

const (
	SlbNewCfgEnhVirtServicesSeventhPartTableSoftGrid_Enabled     SlbNewCfgEnhVirtServicesSeventhPartTableSoftGrid = 1
	SlbNewCfgEnhVirtServicesSeventhPartTableSoftGrid_Disabled    SlbNewCfgEnhVirtServicesSeventhPartTableSoftGrid = 2
	SlbNewCfgEnhVirtServicesSeventhPartTableSoftGrid_Unsupported SlbNewCfgEnhVirtServicesSeventhPartTableSoftGrid = 2147483647
)

type SlbNewCfgEnhVirtServicesSeventhPartTableConnPooling int32

const (
	SlbNewCfgEnhVirtServicesSeventhPartTableConnPooling_Enabled     SlbNewCfgEnhVirtServicesSeventhPartTableConnPooling = 1
	SlbNewCfgEnhVirtServicesSeventhPartTableConnPooling_Disabled    SlbNewCfgEnhVirtServicesSeventhPartTableConnPooling = 2
	SlbNewCfgEnhVirtServicesSeventhPartTableConnPooling_Unsupported SlbNewCfgEnhVirtServicesSeventhPartTableConnPooling = 2147483647
)

type SlbNewCfgEnhVirtServicesSeventhPartTableProxyIpMode int32

const (
	SlbNewCfgEnhVirtServicesSeventhPartTableProxyIpMode_Ingress     SlbNewCfgEnhVirtServicesSeventhPartTableProxyIpMode = 0
	SlbNewCfgEnhVirtServicesSeventhPartTableProxyIpMode_Egress      SlbNewCfgEnhVirtServicesSeventhPartTableProxyIpMode = 1
	SlbNewCfgEnhVirtServicesSeventhPartTableProxyIpMode_Address     SlbNewCfgEnhVirtServicesSeventhPartTableProxyIpMode = 2
	SlbNewCfgEnhVirtServicesSeventhPartTableProxyIpMode_Nwclss      SlbNewCfgEnhVirtServicesSeventhPartTableProxyIpMode = 3
	SlbNewCfgEnhVirtServicesSeventhPartTableProxyIpMode_Disable     SlbNewCfgEnhVirtServicesSeventhPartTableProxyIpMode = 4
	SlbNewCfgEnhVirtServicesSeventhPartTableProxyIpMode_Unsupported SlbNewCfgEnhVirtServicesSeventhPartTableProxyIpMode = 2147483647
)

type SlbNewCfgEnhVirtServicesSeventhPartTableProxyIpPersistency int32

const (
	SlbNewCfgEnhVirtServicesSeventhPartTableProxyIpPersistency_Disable     SlbNewCfgEnhVirtServicesSeventhPartTableProxyIpPersistency = 0
	SlbNewCfgEnhVirtServicesSeventhPartTableProxyIpPersistency_Client      SlbNewCfgEnhVirtServicesSeventhPartTableProxyIpPersistency = 1
	SlbNewCfgEnhVirtServicesSeventhPartTableProxyIpPersistency_Host        SlbNewCfgEnhVirtServicesSeventhPartTableProxyIpPersistency = 2
	SlbNewCfgEnhVirtServicesSeventhPartTableProxyIpPersistency_Unsupported SlbNewCfgEnhVirtServicesSeventhPartTableProxyIpPersistency = 2147483647
)

type SlbNewCfgEnhVirtServicesSeventhPartTableProxyIpNWclassPersistency int32

const (
	SlbNewCfgEnhVirtServicesSeventhPartTableProxyIpNWclassPersistency_Disable     SlbNewCfgEnhVirtServicesSeventhPartTableProxyIpNWclassPersistency = 0
	SlbNewCfgEnhVirtServicesSeventhPartTableProxyIpNWclassPersistency_Client      SlbNewCfgEnhVirtServicesSeventhPartTableProxyIpNWclassPersistency = 1
	SlbNewCfgEnhVirtServicesSeventhPartTableProxyIpNWclassPersistency_Unsupported SlbNewCfgEnhVirtServicesSeventhPartTableProxyIpNWclassPersistency = 2147483647
)

type SlbNewCfgEnhVirtServicesSeventhPartTableClsRST int32

const (
	SlbNewCfgEnhVirtServicesSeventhPartTableClsRST_Enabled     SlbNewCfgEnhVirtServicesSeventhPartTableClsRST = 1
	SlbNewCfgEnhVirtServicesSeventhPartTableClsRST_Disabled    SlbNewCfgEnhVirtServicesSeventhPartTableClsRST = 2
	SlbNewCfgEnhVirtServicesSeventhPartTableClsRST_Unsupported SlbNewCfgEnhVirtServicesSeventhPartTableClsRST = 2147483647
)

type SlbNewCfgEnhVirtServicesSeventhPartTableCluster int32

const (
	SlbNewCfgEnhVirtServicesSeventhPartTableCluster_Enabled     SlbNewCfgEnhVirtServicesSeventhPartTableCluster = 1
	SlbNewCfgEnhVirtServicesSeventhPartTableCluster_Disabled    SlbNewCfgEnhVirtServicesSeventhPartTableCluster = 2
	SlbNewCfgEnhVirtServicesSeventhPartTableCluster_Unsupported SlbNewCfgEnhVirtServicesSeventhPartTableCluster = 2147483647
)

type SlbNewCfgEnhVirtServicesSeventhPartTableReport int32

const (
	SlbNewCfgEnhVirtServicesSeventhPartTableReport_Enabled     SlbNewCfgEnhVirtServicesSeventhPartTableReport = 1
	SlbNewCfgEnhVirtServicesSeventhPartTableReport_Disabled    SlbNewCfgEnhVirtServicesSeventhPartTableReport = 2
	SlbNewCfgEnhVirtServicesSeventhPartTableReport_Unsupported SlbNewCfgEnhVirtServicesSeventhPartTableReport = 2147483647
)

type SlbNewCfgEnhVirtServicesSeventhPartTableIsDnsSecVip int32

const (
	SlbNewCfgEnhVirtServicesSeventhPartTableIsDnsSecVip_No          SlbNewCfgEnhVirtServicesSeventhPartTableIsDnsSecVip = 0
	SlbNewCfgEnhVirtServicesSeventhPartTableIsDnsSecVip_Yes         SlbNewCfgEnhVirtServicesSeventhPartTableIsDnsSecVip = 1
	SlbNewCfgEnhVirtServicesSeventhPartTableIsDnsSecVip_Unsupported SlbNewCfgEnhVirtServicesSeventhPartTableIsDnsSecVip = 2147483647
)

type SlbNewCfgEnhVirtServicesSeventhPartTableParams struct {
	// The number of the virtual server.
	ServSeventhPartIndex string `json:"ServSeventhPartIndex,omitempty"`
	// The service index. This has no external meaning
	SeventhPartIndex int32 `json:"SeventhPartIndex,omitempty"`
	// The real server group number for this service.
	RealGroup string `json:"RealGroup,omitempty"`
	// Enable/disable session mirroring.
	SessionMirror SlbNewCfgEnhVirtServicesSeventhPartTableSessionMirror `json:"SessionMirror,omitempty"`
	// Enable/disable softgrid load balancing.
	SoftGrid SlbNewCfgEnhVirtServicesSeventhPartTableSoftGrid `json:"SoftGrid,omitempty"`
	// Enable/disable connection pooling for HTTP traffic.
	ConnPooling SlbNewCfgEnhVirtServicesSeventhPartTableConnPooling `json:"ConnPooling,omitempty"`
	// The maximum number of minutes a persistent session should exist.
	PersistentTimeOut uint32 `json:"PersistentTimeOut,omitempty"`
	// Set the Proxy IP mode, default is ingress(0).
	// Changing from address(2) to any other mode will clear the configured IPv4/IPv6 address,prefix and persistancy.
	// Changing from  nwclass(3) to any other mode will clear the configured NWclass and NWpersistancy.
	ProxyIpMode SlbNewCfgEnhVirtServicesSeventhPartTableProxyIpMode `json:"ProxyIpMode,omitempty"`
	// This object ID allows configuring IPv4 PIP address.
	// This object ID can be set only if slbNewCfgVirtServiceProxyIpMode is address else return failure.
	// Returns 0 when slbNewCfgVirtServiceProxyIpMode is not set to address.
	// When a subnet is configured user has the ability to select PIP persistency mode.
	// Persistency is relevant only if either IPv4 PIP or IPv6 PIP (or both) are configured as
	// subnet.
	// If neither of the addresses (v4 & v6) are configured or are subnets, the persistency
	// value is disable.
	ProxyIpAddr string `json:"ProxyIpAddr,omitempty"`
	// This object ID allows configuring IPv4 PIP Mask.
	// Returns 0 when slbNewCfgVirtServiceProxyIpMode is not set to address.
	// When a subnet is configured user has the ability to select PIP persistency mode.
	// Persistency is relevant only if either IPv4 PIP or IPv6 PIP (or both) are configured as
	// subnet.
	// If neither of the addresses (v4 & v6) are configured or are subnets, the persistency
	// value is disable.
	ProxyIpMask string `json:"ProxyIpMask,omitempty"`
	// This object ID allows configuring IPv6 PIP address.
	// Returns 0 when slbNewCfgVirtServiceProxyIpMode is not set to address.
	// When a subnet is configured user has the ability to select PIP persistency mode.
	// Persistency is relevant only if either IPv4 PIP or IPv6 PIP (or both) are configured as
	// subnet.
	// Address should be 4-byte hexadecimal colon notation.
	// Valid IPv6 address should be in any of the following forms
	// xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx or
	// xxxx::xxxx:xxxx:xxxx:xxxx or ::xxxx.
	// If neither of the addresses (v4 & v6) are configured or are subnets, the persistency
	// value is disable.
	ProxyIpv6Addr string `json:"ProxyIpv6Addr,omitempty"`
	// This object ID allows configuring IPv6 PIP Mask.
	// Returns 0 when slbNewCfgVirtServiceProxyIpMode is not set to address.
	// When a subnet is configured user has the ability to select PIP persistency mode.
	// Persistency is relevant only if either IPv4 PIP or IPv6 PIP (or both) are configured as
	// subnet.
	// If neither of the addresses (v4 & v6) are configured or are subnets, the persistency
	// value is disable.
	ProxyIpv6Prefix uint32 `json:"ProxyIpv6Prefix,omitempty"`
	// When a subnet is configured user has the ability to select PIP persistency mode.
	// Returns 0 when slbNewCfgVirtServiceProxyIpMode is not set to address.
	// Persistency is relevant only if either IPv4 PIP or IPv6 PIP (or both) are configured as
	// subnet.
	// If neither of the addresses (v4 & v6) are configured or are subnets, the persistency
	// value is disable.
	ProxyIpPersistency SlbNewCfgEnhVirtServicesSeventhPartTableProxyIpPersistency `json:"ProxyIpPersistency,omitempty"`
	// This object ID allows configuring IPv4 Network Class as PIP and PIP persistency mode.
	// Returns empty string when slbNewCfgVirtServiceProxyIpMode is not set to nwclass.
	// Persistency is relevant only if either IPv4 or IPv6 class (or both) are configured.
	// If neither of the classes (v4 & v6) are configured, the persistency configuration value is disable.
	ProxyIpNWclass string `json:"ProxyIpNWclass,omitempty"`
	// This object ID allows configuring IPv6 Network Class as PIP and PIP persistency mode.
	// Returns empty string when slbNewCfgVirtServiceProxyIpMode is not set to nwclass.
	// Persistency is relevant only if either IPv4 or IPv6 class (or both) are configured.
	// If neither of the classes (v4 & v6) are configured, the persistency configuration value is disable.
	ProxyIpv6NWclass string `json:"ProxyIpv6NWclass,omitempty"`
	// This object ID allows configuring Network Class PIP persistency mode.
	// Returns 0 when slbNewCfgVirtServiceProxyIpMode is not set to nwclass.
	// Persistency is relevant only if either IPv4 PIP or IPv6 PIP (or both) are configured as
	// subnet.
	// If neither of the addresses (v4 & v6) are configured or are subnets, the persistency
	// value is disable.
	ProxyIpNWclassPersistency SlbNewCfgEnhVirtServicesSeventhPartTableProxyIpNWclassPersistency `json:"ProxyIpNWclassPersistency,omitempty"`
	// Set length for slb service sip hashing (4- 256 bytes)
	HashLen uint32 `json:"HashLen,omitempty"`
	// Enable or disable Send RST on connection close.
	ClsRST SlbNewCfgEnhVirtServicesSeventhPartTableClsRST `json:"ClsRST,omitempty"`
	// The HTTP header name of the virtual server.
	HttpHdrName string `json:"HttpHdrName,omitempty"`
	// Fastview web application name associated with this virtual service.Set none to delete entry
	ServFastWa string `json:"ServFastWa,omitempty"`
	// This object ID allows configuring the web security ID
	AppwallWebappId string `json:"AppwallWebappId,omitempty"`
	// Http2 policy name associated with this virtual service.
	Http2 string `json:"Http2,omitempty"`
	// Enable or disable Cluster Updates for the service.
	Cluster SlbNewCfgEnhVirtServicesSeventhPartTableCluster `json:"Cluster,omitempty"`
	// The ftp control service data port
	DataPort uint64 `json:"DataPort,omitempty"`
	// Correlate several services into one application at the visualization.
	ApplicName string `json:"ApplicName,omitempty"`
	// Enable/disable counter based reporting for service.
	Report SlbNewCfgEnhVirtServicesSeventhPartTableReport `json:"Report,omitempty"`
	// Set Traffic Event Log Policy.
	Trevpol string `json:"Trevpol,omitempty"`
	// Application satisfied response time threshold, inherits the value from the global satisfied value or set with different value 1-999999 ms.
	Satisrt uint64 `json:"Satisrt,omitempty"`
	// Set Bot Manager Policy.
	Botpol string `json:"Botpol,omitempty"`
	// Set DNS Nameserver group.
	Namesrvr string `json:"Namesrvr,omitempty"`
	// It returns Yes(1) if virtual service is configure auto with a DNS Responder VIP, else returns no(0).
	IsDnsSecVip SlbNewCfgEnhVirtServicesSeventhPartTableIsDnsSecVip `json:"IsDnsSecVip,omitempty"`
	// Set SecurePath Policy.
	SecurePathpol string `json:"SecurePathpol,omitempty"`
}

func (p SlbNewCfgEnhVirtServicesSeventhPartTableParams) iMABean() {}
