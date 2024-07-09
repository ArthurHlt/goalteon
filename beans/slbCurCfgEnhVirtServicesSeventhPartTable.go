package beans

import (
	"fmt"
	"reflect"
)

// SlbCurCfgEnhVirtServicesSeventhPartTable The table of virtual Services.
// Note:This mib is not supported for VX instance of Virtualization.
type SlbCurCfgEnhVirtServicesSeventhPartTable struct {
	// The number of the virtual server.
	SlbCurCfgEnhVirtServSeventhPartIndex string
	// The service index. This has no external meaning
	SlbCurCfgEnhVirtServiceSeventhPartIndex int32
	Params                                  *SlbCurCfgEnhVirtServicesSeventhPartTableParams
}

func NewSlbCurCfgEnhVirtServicesSeventhPartTableList() *SlbCurCfgEnhVirtServicesSeventhPartTable {
	return &SlbCurCfgEnhVirtServicesSeventhPartTable{}
}

func NewSlbCurCfgEnhVirtServicesSeventhPartTable(
	slbCurCfgEnhVirtServSeventhPartIndex string,
	slbCurCfgEnhVirtServiceSeventhPartIndex int32,
	params *SlbCurCfgEnhVirtServicesSeventhPartTableParams,
) *SlbCurCfgEnhVirtServicesSeventhPartTable {
	return &SlbCurCfgEnhVirtServicesSeventhPartTable{
		SlbCurCfgEnhVirtServSeventhPartIndex:    slbCurCfgEnhVirtServSeventhPartIndex,
		SlbCurCfgEnhVirtServiceSeventhPartIndex: slbCurCfgEnhVirtServiceSeventhPartIndex,
		Params:                                  params,
	}
}

func (c *SlbCurCfgEnhVirtServicesSeventhPartTable) Name() string {
	return "SlbCurCfgEnhVirtServicesSeventhPartTable"
}

func (c *SlbCurCfgEnhVirtServicesSeventhPartTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbCurCfgEnhVirtServicesSeventhPartTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbCurCfgEnhVirtServicesSeventhPartTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbCurCfgEnhVirtServSeventhPartIndex).IsZero() &&
		reflect.ValueOf(c.SlbCurCfgEnhVirtServiceSeventhPartIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbCurCfgEnhVirtServSeventhPartIndex) + "/" + fmt.Sprint(c.SlbCurCfgEnhVirtServiceSeventhPartIndex)
}

type SlbCurCfgEnhVirtServicesSeventhPartTableSessionMirror int32

const (
	SlbCurCfgEnhVirtServicesSeventhPartTableSessionMirror_Enabled     SlbCurCfgEnhVirtServicesSeventhPartTableSessionMirror = 1
	SlbCurCfgEnhVirtServicesSeventhPartTableSessionMirror_Disabled    SlbCurCfgEnhVirtServicesSeventhPartTableSessionMirror = 2
	SlbCurCfgEnhVirtServicesSeventhPartTableSessionMirror_Unsupported SlbCurCfgEnhVirtServicesSeventhPartTableSessionMirror = 2147483647
)

type SlbCurCfgEnhVirtServicesSeventhPartTableSoftGrid int32

const (
	SlbCurCfgEnhVirtServicesSeventhPartTableSoftGrid_Enabled     SlbCurCfgEnhVirtServicesSeventhPartTableSoftGrid = 1
	SlbCurCfgEnhVirtServicesSeventhPartTableSoftGrid_Disabled    SlbCurCfgEnhVirtServicesSeventhPartTableSoftGrid = 2
	SlbCurCfgEnhVirtServicesSeventhPartTableSoftGrid_Unsupported SlbCurCfgEnhVirtServicesSeventhPartTableSoftGrid = 2147483647
)

type SlbCurCfgEnhVirtServicesSeventhPartTableConnPooling int32

const (
	SlbCurCfgEnhVirtServicesSeventhPartTableConnPooling_Enabled     SlbCurCfgEnhVirtServicesSeventhPartTableConnPooling = 1
	SlbCurCfgEnhVirtServicesSeventhPartTableConnPooling_Disabled    SlbCurCfgEnhVirtServicesSeventhPartTableConnPooling = 2
	SlbCurCfgEnhVirtServicesSeventhPartTableConnPooling_Unsupported SlbCurCfgEnhVirtServicesSeventhPartTableConnPooling = 2147483647
)

type SlbCurCfgEnhVirtServicesSeventhPartTableProxyIpMode int32

const (
	SlbCurCfgEnhVirtServicesSeventhPartTableProxyIpMode_Ingress     SlbCurCfgEnhVirtServicesSeventhPartTableProxyIpMode = 0
	SlbCurCfgEnhVirtServicesSeventhPartTableProxyIpMode_Egress      SlbCurCfgEnhVirtServicesSeventhPartTableProxyIpMode = 1
	SlbCurCfgEnhVirtServicesSeventhPartTableProxyIpMode_Address     SlbCurCfgEnhVirtServicesSeventhPartTableProxyIpMode = 2
	SlbCurCfgEnhVirtServicesSeventhPartTableProxyIpMode_Nwclss      SlbCurCfgEnhVirtServicesSeventhPartTableProxyIpMode = 3
	SlbCurCfgEnhVirtServicesSeventhPartTableProxyIpMode_Disable     SlbCurCfgEnhVirtServicesSeventhPartTableProxyIpMode = 4
	SlbCurCfgEnhVirtServicesSeventhPartTableProxyIpMode_Unsupported SlbCurCfgEnhVirtServicesSeventhPartTableProxyIpMode = 2147483647
)

type SlbCurCfgEnhVirtServicesSeventhPartTableProxyIpPersistency int32

const (
	SlbCurCfgEnhVirtServicesSeventhPartTableProxyIpPersistency_Disable     SlbCurCfgEnhVirtServicesSeventhPartTableProxyIpPersistency = 0
	SlbCurCfgEnhVirtServicesSeventhPartTableProxyIpPersistency_Client      SlbCurCfgEnhVirtServicesSeventhPartTableProxyIpPersistency = 1
	SlbCurCfgEnhVirtServicesSeventhPartTableProxyIpPersistency_Host        SlbCurCfgEnhVirtServicesSeventhPartTableProxyIpPersistency = 2
	SlbCurCfgEnhVirtServicesSeventhPartTableProxyIpPersistency_Unsupported SlbCurCfgEnhVirtServicesSeventhPartTableProxyIpPersistency = 2147483647
)

type SlbCurCfgEnhVirtServicesSeventhPartTableProxyIpNWclassPersistency int32

const (
	SlbCurCfgEnhVirtServicesSeventhPartTableProxyIpNWclassPersistency_Disable     SlbCurCfgEnhVirtServicesSeventhPartTableProxyIpNWclassPersistency = 0
	SlbCurCfgEnhVirtServicesSeventhPartTableProxyIpNWclassPersistency_Client      SlbCurCfgEnhVirtServicesSeventhPartTableProxyIpNWclassPersistency = 1
	SlbCurCfgEnhVirtServicesSeventhPartTableProxyIpNWclassPersistency_Unsupported SlbCurCfgEnhVirtServicesSeventhPartTableProxyIpNWclassPersistency = 2147483647
)

type SlbCurCfgEnhVirtServicesSeventhPartTableClsRST int32

const (
	SlbCurCfgEnhVirtServicesSeventhPartTableClsRST_Enabled     SlbCurCfgEnhVirtServicesSeventhPartTableClsRST = 1
	SlbCurCfgEnhVirtServicesSeventhPartTableClsRST_Disabled    SlbCurCfgEnhVirtServicesSeventhPartTableClsRST = 2
	SlbCurCfgEnhVirtServicesSeventhPartTableClsRST_Unsupported SlbCurCfgEnhVirtServicesSeventhPartTableClsRST = 2147483647
)

type SlbCurCfgEnhVirtServicesSeventhPartTableCluster int32

const (
	SlbCurCfgEnhVirtServicesSeventhPartTableCluster_Enabled     SlbCurCfgEnhVirtServicesSeventhPartTableCluster = 1
	SlbCurCfgEnhVirtServicesSeventhPartTableCluster_Disabled    SlbCurCfgEnhVirtServicesSeventhPartTableCluster = 2
	SlbCurCfgEnhVirtServicesSeventhPartTableCluster_Unsupported SlbCurCfgEnhVirtServicesSeventhPartTableCluster = 2147483647
)

type SlbCurCfgEnhVirtServicesSeventhPartTableReport int32

const (
	SlbCurCfgEnhVirtServicesSeventhPartTableReport_Enabled     SlbCurCfgEnhVirtServicesSeventhPartTableReport = 1
	SlbCurCfgEnhVirtServicesSeventhPartTableReport_Disabled    SlbCurCfgEnhVirtServicesSeventhPartTableReport = 2
	SlbCurCfgEnhVirtServicesSeventhPartTableReport_Unsupported SlbCurCfgEnhVirtServicesSeventhPartTableReport = 2147483647
)

type SlbCurCfgEnhVirtServicesSeventhPartTableIsDnsSecVip int32

const (
	SlbCurCfgEnhVirtServicesSeventhPartTableIsDnsSecVip_No          SlbCurCfgEnhVirtServicesSeventhPartTableIsDnsSecVip = 0
	SlbCurCfgEnhVirtServicesSeventhPartTableIsDnsSecVip_Yes         SlbCurCfgEnhVirtServicesSeventhPartTableIsDnsSecVip = 1
	SlbCurCfgEnhVirtServicesSeventhPartTableIsDnsSecVip_Unsupported SlbCurCfgEnhVirtServicesSeventhPartTableIsDnsSecVip = 2147483647
)

type SlbCurCfgEnhVirtServicesSeventhPartTableAwinflow int32

const (
	SlbCurCfgEnhVirtServicesSeventhPartTableAwinflow_After       SlbCurCfgEnhVirtServicesSeventhPartTableAwinflow = 1
	SlbCurCfgEnhVirtServicesSeventhPartTableAwinflow_Before      SlbCurCfgEnhVirtServicesSeventhPartTableAwinflow = 2
	SlbCurCfgEnhVirtServicesSeventhPartTableAwinflow_Unsupported SlbCurCfgEnhVirtServicesSeventhPartTableAwinflow = 2147483647
)

type SlbCurCfgEnhVirtServicesSeventhPartTableFallbackUseAW int32

const (
	SlbCurCfgEnhVirtServicesSeventhPartTableFallbackUseAW_Enabled     SlbCurCfgEnhVirtServicesSeventhPartTableFallbackUseAW = 1
	SlbCurCfgEnhVirtServicesSeventhPartTableFallbackUseAW_Disabled    SlbCurCfgEnhVirtServicesSeventhPartTableFallbackUseAW = 2
	SlbCurCfgEnhVirtServicesSeventhPartTableFallbackUseAW_Unsupported SlbCurCfgEnhVirtServicesSeventhPartTableFallbackUseAW = 2147483647
)

type SlbCurCfgEnhVirtServicesSeventhPartTableJsinject int32

const (
	SlbCurCfgEnhVirtServicesSeventhPartTableJsinject_Enabled     SlbCurCfgEnhVirtServicesSeventhPartTableJsinject = 1
	SlbCurCfgEnhVirtServicesSeventhPartTableJsinject_Disabled    SlbCurCfgEnhVirtServicesSeventhPartTableJsinject = 2
	SlbCurCfgEnhVirtServicesSeventhPartTableJsinject_Unsupported SlbCurCfgEnhVirtServicesSeventhPartTableJsinject = 2147483647
)

type SlbCurCfgEnhVirtServicesSeventhPartTableDohType int32

const (
	SlbCurCfgEnhVirtServicesSeventhPartTableDohType_Disabled       SlbCurCfgEnhVirtServicesSeventhPartTableDohType = 1
	SlbCurCfgEnhVirtServicesSeventhPartTableDohType_DohUdp         SlbCurCfgEnhVirtServicesSeventhPartTableDohType = 2
	SlbCurCfgEnhVirtServicesSeventhPartTableDohType_DohUdpFallback SlbCurCfgEnhVirtServicesSeventhPartTableDohType = 3
	SlbCurCfgEnhVirtServicesSeventhPartTableDohType_DohTcp         SlbCurCfgEnhVirtServicesSeventhPartTableDohType = 4
	SlbCurCfgEnhVirtServicesSeventhPartTableDohType_Unsupported    SlbCurCfgEnhVirtServicesSeventhPartTableDohType = 2147483647
)

type SlbCurCfgEnhVirtServicesSeventhPartTableParams struct {
	// The number of the virtual server.
	ServSeventhPartIndex string `json:"ServSeventhPartIndex,omitempty"`
	// The service index. This has no external meaning
	SeventhPartIndex int32 `json:"SeventhPartIndex,omitempty"`
	// The real server group number for this service.
	RealGroup string `json:"RealGroup,omitempty"`
	// Enable/disable session mirroring.
	SessionMirror SlbCurCfgEnhVirtServicesSeventhPartTableSessionMirror `json:"SessionMirror,omitempty"`
	// Enable/disable softgrid load balancing.
	SoftGrid SlbCurCfgEnhVirtServicesSeventhPartTableSoftGrid `json:"SoftGrid,omitempty"`
	// Enable/disable connection pooling for HTTP traffic.
	ConnPooling SlbCurCfgEnhVirtServicesSeventhPartTableConnPooling `json:"ConnPooling,omitempty"`
	// The maximum number of minutes a persistent session should exist.
	PersistentTimeOut uint32 `json:"PersistentTimeOut,omitempty"`
	// Set the Proxy IP mode, default is ingress(0).
	// Changing from address(2) to any other mode will clear the configured IPv4/IPv6 address,prefix and persistancy.
	// Changing from  nwclass(3) to any other mode will clear the configured NWclass and NWpersistancy.
	ProxyIpMode SlbCurCfgEnhVirtServicesSeventhPartTableProxyIpMode `json:"ProxyIpMode,omitempty"`
	// This object ID allows configuring IPv4 PIP address.
	// This object ID can be set only if slbCurCfgVirtServiceProxyIpMode is address else return failure.
	// Returns 0 when slbCurCfgVirtServiceProxyIpMode is not set to address.
	// When a subnet is configured user has the ability to select PIP persistency mode.
	// Persistency is relevant only if either IPv4 PIP or IPv6 PIP (or both) are configured as
	// subnet.
	// If neither of the addresses (v4 & v6) are configured or are subnets, the persistency
	// value is disable.
	ProxyIpAddr string `json:"ProxyIpAddr,omitempty"`
	// This object ID allows configuring IPv4 PIP Mask.
	// Returns 0 when slbCurCfgVirtServiceProxyIpMode is not set to address.
	// When a subnet is configured user has the ability to select PIP persistency mode.
	// Persistency is relevant only if either IPv4 PIP or IPv6 PIP (or both) are configured as
	// subnet.
	// If neither of the addresses (v4 & v6) are configured or are subnets, the persistency
	// value is disable.
	ProxyIpMask string `json:"ProxyIpMask,omitempty"`
	// This object ID allows configuring IPv6 PIP address.
	// Returns 0 when slbCurCfgVirtServiceProxyIpMode is not set to address.
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
	// Returns 0 when slbCurCfgVirtServiceProxyIpMode is not set to address.
	// When a subnet is configured user has the ability to select PIP persistency mode.
	// Persistency is relevant only if either IPv4 PIP or IPv6 PIP (or both) are configured as
	// subnet.
	// If neither of the addresses (v4 & v6) are configured or are subnets, the persistency
	// value is disable.
	ProxyIpv6Prefix uint32 `json:"ProxyIpv6Prefix,omitempty"`
	// When a subnet is configured user has the ability to select PIP persistency mode.
	// Returns 0 when slbCurCfgVirtServiceProxyIpMode is not set to address.
	// Persistency is relevant only if either IPv4 PIP or IPv6 PIP (or both) are configured as
	// subnet.
	// If neither of the addresses (v4 & v6) are configured or are subnets, the persistency
	// value is disable.
	ProxyIpPersistency SlbCurCfgEnhVirtServicesSeventhPartTableProxyIpPersistency `json:"ProxyIpPersistency,omitempty"`
	// This object ID allows configuring IPv4 Network Class as PIP and PIP persistency mode.
	// Returns empty string when slbCurCfgVirtServiceProxyIpMode is not set to nwclass.
	// Persistency is relevant only if either IPv4 or IPv6 class (or both) are configured.
	// If neither of the classes (v4 & v6) are configured, the persistency configuration value is disable.
	ProxyIpNWclass string `json:"ProxyIpNWclass,omitempty"`
	// This object ID allows configuring IPv6 Network Class as PIP and PIP persistency mode.
	// Returns empty string when slbCurCfgVirtServiceProxyIpMode is not set to nwclass.
	// Persistency is relevant only if either IPv4 or IPv6 class (or both) are configured.
	// If neither of the classes (v4 & v6) are configured, the persistency configuration value is disable.
	ProxyIpv6NWclass string `json:"ProxyIpv6NWclass,omitempty"`
	// This object ID allows configuring Network Class PIP persistency mode.
	// Returns 0 when slbCurCfgVirtServiceProxyIpMode is not set to nwclass.
	// Persistency is relevant only if either IPv4 PIP or IPv6 PIP (or both) are configured as
	// subnet.
	// If neither of the addresses (v4 & v6) are configured or are subnets, the persistency
	// value is disable.
	ProxyIpNWclassPersistency SlbCurCfgEnhVirtServicesSeventhPartTableProxyIpNWclassPersistency `json:"ProxyIpNWclassPersistency,omitempty"`
	// Set length for slb service sip hashing (4- 256 bytes)
	HashLen uint32 `json:"HashLen,omitempty"`
	// Enable or disable Send RST on connection close.
	ClsRST SlbCurCfgEnhVirtServicesSeventhPartTableClsRST `json:"ClsRST,omitempty"`
	// The HTTP header name of the virtual server.
	HttpHdrName string `json:"HttpHdrName,omitempty"`
	// Fastview web application name associated with this virtual service.
	ServFastWa string `json:"ServFastWa,omitempty"`
	// This object ID shows current configured web security ID.
	AppwallWebappId string `json:"AppwallWebappId,omitempty"`
	// Http2 policy name associated with this virtual service.
	Http2 string `json:"Http2,omitempty"`
	// State of Cluster upadates for the service (Ena or Dis).
	Cluster SlbCurCfgEnhVirtServicesSeventhPartTableCluster `json:"Cluster,omitempty"`
	// The ftp control service data port
	DataPort uint64 `json:"DataPort,omitempty"`
	// Correlate several services into one application at the visualization.
	ApplicName string `json:"ApplicName,omitempty"`
	// Enable/disable counter based reporting for service.
	Report SlbCurCfgEnhVirtServicesSeventhPartTableReport `json:"Report,omitempty"`
	// Current Traffic Event Log Policy.
	Trevpol string `json:"Trevpol,omitempty"`
	// Application satisfied response time threshold, inherits the value from the global satisfied value or set with different value 1-999999 ms.
	Satisrt uint64 `json:"Satisrt,omitempty"`
	// Set Bot Manager Policy.
	Botpol string `json:"Botpol,omitempty"`
	// Current DNS Nameserver group.
	Namesrvr string `json:"Namesrvr,omitempty"`
	// It returns Yes(1) if virtual service is configure auto with a DNS Responder VIP, else returns no(0).
	IsDnsSecVip SlbCurCfgEnhVirtServicesSeventhPartTableIsDnsSecVip `json:"IsDnsSecVip,omitempty"`
	// Http3 policy name associated with this virtual service.
	Http3 string `json:"Http3,omitempty"`
	// Quic policy name associated with this virtual service.
	Quic string `json:"Quic,omitempty"`
	// Set if AW processing comes before or after Alteon HTTP parsing.
	Awinflow SlbCurCfgEnhVirtServicesSeventhPartTableAwinflow `json:"Awinflow,omitempty"`
	// Enable/disable security web application processing when no content rule matches.
	FallbackUseAW SlbCurCfgEnhVirtServicesSeventhPartTableFallbackUseAW `json:"FallbackUseAW,omitempty"`
	// Set HTTP3 port for this virtual service
	Http3port uint64 `json:"Http3port,omitempty"`
	// Set SecurePath Policy.
	SecurePathpol string `json:"SecurePathpol,omitempty"`
	// Set JS inject mode .
	Jsinject SlbCurCfgEnhVirtServicesSeventhPartTableJsinject `json:"Jsinject,omitempty"`
	// Set DOH mode .
	DohType SlbCurCfgEnhVirtServicesSeventhPartTableDohType `json:"DohType,omitempty"`
}

func (p SlbCurCfgEnhVirtServicesSeventhPartTableParams) iMABean() {}
