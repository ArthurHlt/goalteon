package beans

import (
	"fmt"
	"reflect"
)

// SlbNewCfgVirtServicesTable The table of virtual Services.
// Note:This mib is not supported for VX instance of Virtualization.
type SlbNewCfgVirtServicesTable struct {
	// The number of the virtual server.
	SlbNewCfgVirtServIndex int32
	// The service index. This has no external meaning
	SlbNewCfgVirtServiceIndex int32
	Params                    *SlbNewCfgVirtServicesTableParams
}

func NewSlbNewCfgVirtServicesTableList() *SlbNewCfgVirtServicesTable {
	return &SlbNewCfgVirtServicesTable{}
}

func NewSlbNewCfgVirtServicesTable(
	slbNewCfgVirtServIndex int32,
	slbNewCfgVirtServiceIndex int32,
	params *SlbNewCfgVirtServicesTableParams,
) *SlbNewCfgVirtServicesTable {
	return &SlbNewCfgVirtServicesTable{
		SlbNewCfgVirtServIndex:    slbNewCfgVirtServIndex,
		SlbNewCfgVirtServiceIndex: slbNewCfgVirtServiceIndex,
		Params:                    params,
	}
}

func (c *SlbNewCfgVirtServicesTable) Name() string {
	return "SlbNewCfgVirtServicesTable"
}

func (c *SlbNewCfgVirtServicesTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbNewCfgVirtServicesTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbNewCfgVirtServicesTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbNewCfgVirtServIndex).IsZero() &&
		reflect.ValueOf(c.SlbNewCfgVirtServiceIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbNewCfgVirtServIndex) + "/" + fmt.Sprint(c.SlbNewCfgVirtServiceIndex)
}

type SlbNewCfgVirtServicesTableUDPBalance int32

const (
	SlbNewCfgVirtServicesTableUDPBalance_Udp         SlbNewCfgVirtServicesTableUDPBalance = 2
	SlbNewCfgVirtServicesTableUDPBalance_Tcp         SlbNewCfgVirtServicesTableUDPBalance = 3
	SlbNewCfgVirtServicesTableUDPBalance_Stateless   SlbNewCfgVirtServicesTableUDPBalance = 4
	SlbNewCfgVirtServicesTableUDPBalance_TcpAndUdp   SlbNewCfgVirtServicesTableUDPBalance = 5
	SlbNewCfgVirtServicesTableUDPBalance_Sctp        SlbNewCfgVirtServicesTableUDPBalance = 6
	SlbNewCfgVirtServicesTableUDPBalance_Unsupported SlbNewCfgVirtServicesTableUDPBalance = 2147483647
)

type SlbNewCfgVirtServicesTableDirServerRtn int32

const (
	SlbNewCfgVirtServicesTableDirServerRtn_Enabled     SlbNewCfgVirtServicesTableDirServerRtn = 1
	SlbNewCfgVirtServicesTableDirServerRtn_Disabled    SlbNewCfgVirtServicesTableDirServerRtn = 2
	SlbNewCfgVirtServicesTableDirServerRtn_Unsupported SlbNewCfgVirtServicesTableDirServerRtn = 2147483647
)

type SlbNewCfgVirtServicesTableRtspUrlParse int32

const (
	SlbNewCfgVirtServicesTableRtspUrlParse_None         SlbNewCfgVirtServicesTableRtspUrlParse = 1
	SlbNewCfgVirtServicesTableRtspUrlParse_L4hash       SlbNewCfgVirtServicesTableRtspUrlParse = 2
	SlbNewCfgVirtServicesTableRtspUrlParse_Hash         SlbNewCfgVirtServicesTableRtspUrlParse = 3
	SlbNewCfgVirtServicesTableRtspUrlParse_PatternMatch SlbNewCfgVirtServicesTableRtspUrlParse = 4
	SlbNewCfgVirtServicesTableRtspUrlParse_Unsupported  SlbNewCfgVirtServicesTableRtspUrlParse = 2147483647
)

type SlbNewCfgVirtServicesTableDBind int32

const (
	SlbNewCfgVirtServicesTableDBind_Enabled     SlbNewCfgVirtServicesTableDBind = 1
	SlbNewCfgVirtServicesTableDBind_Disabled    SlbNewCfgVirtServicesTableDBind = 2
	SlbNewCfgVirtServicesTableDBind_Forceproxy  SlbNewCfgVirtServicesTableDBind = 3
	SlbNewCfgVirtServicesTableDBind_Unsupported SlbNewCfgVirtServicesTableDBind = 2147483647
)

type SlbNewCfgVirtServicesTableFtpParsing int32

const (
	SlbNewCfgVirtServicesTableFtpParsing_Enabled     SlbNewCfgVirtServicesTableFtpParsing = 1
	SlbNewCfgVirtServicesTableFtpParsing_Disabled    SlbNewCfgVirtServicesTableFtpParsing = 2
	SlbNewCfgVirtServicesTableFtpParsing_Unsupported SlbNewCfgVirtServicesTableFtpParsing = 2147483647
)

type SlbNewCfgVirtServicesTableRemapUDPFrags int32

const (
	SlbNewCfgVirtServicesTableRemapUDPFrags_Enabled     SlbNewCfgVirtServicesTableRemapUDPFrags = 1
	SlbNewCfgVirtServicesTableRemapUDPFrags_Disabled    SlbNewCfgVirtServicesTableRemapUDPFrags = 2
	SlbNewCfgVirtServicesTableRemapUDPFrags_Unsupported SlbNewCfgVirtServicesTableRemapUDPFrags = 2147483647
)

type SlbNewCfgVirtServicesTableDnsSlb int32

const (
	SlbNewCfgVirtServicesTableDnsSlb_Enabled     SlbNewCfgVirtServicesTableDnsSlb = 1
	SlbNewCfgVirtServicesTableDnsSlb_Disabled    SlbNewCfgVirtServicesTableDnsSlb = 2
	SlbNewCfgVirtServicesTableDnsSlb_Unsupported SlbNewCfgVirtServicesTableDnsSlb = 2147483647
)

type SlbNewCfgVirtServicesTablePBind int32

const (
	SlbNewCfgVirtServicesTablePBind_Clientip    SlbNewCfgVirtServicesTablePBind = 2
	SlbNewCfgVirtServicesTablePBind_Disabled    SlbNewCfgVirtServicesTablePBind = 3
	SlbNewCfgVirtServicesTablePBind_Sslid       SlbNewCfgVirtServicesTablePBind = 4
	SlbNewCfgVirtServicesTablePBind_Cookie      SlbNewCfgVirtServicesTablePBind = 5
	SlbNewCfgVirtServicesTablePBind_Unsupported SlbNewCfgVirtServicesTablePBind = 2147483647
)

type SlbNewCfgVirtServicesTableUriCookie int32

const (
	SlbNewCfgVirtServicesTableUriCookie_Enabled     SlbNewCfgVirtServicesTableUriCookie = 1
	SlbNewCfgVirtServicesTableUriCookie_Disabled    SlbNewCfgVirtServicesTableUriCookie = 2
	SlbNewCfgVirtServicesTableUriCookie_Unsupported SlbNewCfgVirtServicesTableUriCookie = 2147483647
)

type SlbNewCfgVirtServicesTableCookieMode int32

const (
	SlbNewCfgVirtServicesTableCookieMode_Rewrite     SlbNewCfgVirtServicesTableCookieMode = 1
	SlbNewCfgVirtServicesTableCookieMode_Passive     SlbNewCfgVirtServicesTableCookieMode = 2
	SlbNewCfgVirtServicesTableCookieMode_Insert      SlbNewCfgVirtServicesTableCookieMode = 3
	SlbNewCfgVirtServicesTableCookieMode_Unsupported SlbNewCfgVirtServicesTableCookieMode = 2147483647
)

type SlbNewCfgVirtServicesTableHttpSlb int32

const (
	SlbNewCfgVirtServicesTableHttpSlb_Disabled    SlbNewCfgVirtServicesTableHttpSlb = 1
	SlbNewCfgVirtServicesTableHttpSlb_Urlslb      SlbNewCfgVirtServicesTableHttpSlb = 2
	SlbNewCfgVirtServicesTableHttpSlb_Urlhash     SlbNewCfgVirtServicesTableHttpSlb = 3
	SlbNewCfgVirtServicesTableHttpSlb_Cookie      SlbNewCfgVirtServicesTableHttpSlb = 4
	SlbNewCfgVirtServicesTableHttpSlb_Host        SlbNewCfgVirtServicesTableHttpSlb = 5
	SlbNewCfgVirtServicesTableHttpSlb_Browser     SlbNewCfgVirtServicesTableHttpSlb = 6
	SlbNewCfgVirtServicesTableHttpSlb_Others      SlbNewCfgVirtServicesTableHttpSlb = 7
	SlbNewCfgVirtServicesTableHttpSlb_Headerhash  SlbNewCfgVirtServicesTableHttpSlb = 8
	SlbNewCfgVirtServicesTableHttpSlb_Version     SlbNewCfgVirtServicesTableHttpSlb = 9
	SlbNewCfgVirtServicesTableHttpSlb_Unsupported SlbNewCfgVirtServicesTableHttpSlb = 2147483647
)

type SlbNewCfgVirtServicesTableHttpSlbOption int32

const (
	SlbNewCfgVirtServicesTableHttpSlbOption_And         SlbNewCfgVirtServicesTableHttpSlbOption = 1
	SlbNewCfgVirtServicesTableHttpSlbOption_Or          SlbNewCfgVirtServicesTableHttpSlbOption = 2
	SlbNewCfgVirtServicesTableHttpSlbOption_None        SlbNewCfgVirtServicesTableHttpSlbOption = 3
	SlbNewCfgVirtServicesTableHttpSlbOption_Unsupported SlbNewCfgVirtServicesTableHttpSlbOption = 2147483647
)

type SlbNewCfgVirtServicesTableHttpSlb2 int32

const (
	SlbNewCfgVirtServicesTableHttpSlb2_Disabled    SlbNewCfgVirtServicesTableHttpSlb2 = 1
	SlbNewCfgVirtServicesTableHttpSlb2_Urlslb      SlbNewCfgVirtServicesTableHttpSlb2 = 2
	SlbNewCfgVirtServicesTableHttpSlb2_Urlhash     SlbNewCfgVirtServicesTableHttpSlb2 = 3
	SlbNewCfgVirtServicesTableHttpSlb2_Cookie      SlbNewCfgVirtServicesTableHttpSlb2 = 4
	SlbNewCfgVirtServicesTableHttpSlb2_Host        SlbNewCfgVirtServicesTableHttpSlb2 = 5
	SlbNewCfgVirtServicesTableHttpSlb2_Browser     SlbNewCfgVirtServicesTableHttpSlb2 = 6
	SlbNewCfgVirtServicesTableHttpSlb2_Others      SlbNewCfgVirtServicesTableHttpSlb2 = 7
	SlbNewCfgVirtServicesTableHttpSlb2_Headerhash  SlbNewCfgVirtServicesTableHttpSlb2 = 8
	SlbNewCfgVirtServicesTableHttpSlb2_Version     SlbNewCfgVirtServicesTableHttpSlb2 = 9
	SlbNewCfgVirtServicesTableHttpSlb2_Unsupported SlbNewCfgVirtServicesTableHttpSlb2 = 2147483647
)

type SlbNewCfgVirtServicesTableDelete int32

const (
	SlbNewCfgVirtServicesTableDelete_Other       SlbNewCfgVirtServicesTableDelete = 1
	SlbNewCfgVirtServicesTableDelete_Delete      SlbNewCfgVirtServicesTableDelete = 2
	SlbNewCfgVirtServicesTableDelete_Unsupported SlbNewCfgVirtServicesTableDelete = 2147483647
)

type SlbNewCfgVirtServicesTableDirect int32

const (
	SlbNewCfgVirtServicesTableDirect_Enabled     SlbNewCfgVirtServicesTableDirect = 1
	SlbNewCfgVirtServicesTableDirect_Disabled    SlbNewCfgVirtServicesTableDirect = 2
	SlbNewCfgVirtServicesTableDirect_Unsupported SlbNewCfgVirtServicesTableDirect = 2147483647
)

type SlbNewCfgVirtServicesTableThash int32

const (
	SlbNewCfgVirtServicesTableThash_Sip         SlbNewCfgVirtServicesTableThash = 1
	SlbNewCfgVirtServicesTableThash_SipSport    SlbNewCfgVirtServicesTableThash = 2
	SlbNewCfgVirtServicesTableThash_Unsupported SlbNewCfgVirtServicesTableThash = 2147483647
)

type SlbNewCfgVirtServicesTableLdapreset int32

const (
	SlbNewCfgVirtServicesTableLdapreset_Enabled     SlbNewCfgVirtServicesTableLdapreset = 1
	SlbNewCfgVirtServicesTableLdapreset_Disabled    SlbNewCfgVirtServicesTableLdapreset = 2
	SlbNewCfgVirtServicesTableLdapreset_Unsupported SlbNewCfgVirtServicesTableLdapreset = 2147483647
)

type SlbNewCfgVirtServicesTableLdapslb int32

const (
	SlbNewCfgVirtServicesTableLdapslb_Enabled     SlbNewCfgVirtServicesTableLdapslb = 1
	SlbNewCfgVirtServicesTableLdapslb_Disabled    SlbNewCfgVirtServicesTableLdapslb = 2
	SlbNewCfgVirtServicesTableLdapslb_Unsupported SlbNewCfgVirtServicesTableLdapslb = 2147483647
)

type SlbNewCfgVirtServicesTableSip int32

const (
	SlbNewCfgVirtServicesTableSip_Enabled     SlbNewCfgVirtServicesTableSip = 1
	SlbNewCfgVirtServicesTableSip_Disabled    SlbNewCfgVirtServicesTableSip = 2
	SlbNewCfgVirtServicesTableSip_Unsupported SlbNewCfgVirtServicesTableSip = 2147483647
)

type SlbNewCfgVirtServicesTableXForwardedFor int32

const (
	SlbNewCfgVirtServicesTableXForwardedFor_Enabled     SlbNewCfgVirtServicesTableXForwardedFor = 1
	SlbNewCfgVirtServicesTableXForwardedFor_Disabled    SlbNewCfgVirtServicesTableXForwardedFor = 2
	SlbNewCfgVirtServicesTableXForwardedFor_Unsupported SlbNewCfgVirtServicesTableXForwardedFor = 2147483647
)

type SlbNewCfgVirtServicesTableHttpRedir int32

const (
	SlbNewCfgVirtServicesTableHttpRedir_Enabled     SlbNewCfgVirtServicesTableHttpRedir = 1
	SlbNewCfgVirtServicesTableHttpRedir_Disabled    SlbNewCfgVirtServicesTableHttpRedir = 2
	SlbNewCfgVirtServicesTableHttpRedir_Unsupported SlbNewCfgVirtServicesTableHttpRedir = 2147483647
)

type SlbNewCfgVirtServicesTablePbindRport int32

const (
	SlbNewCfgVirtServicesTablePbindRport_Enabled     SlbNewCfgVirtServicesTablePbindRport = 1
	SlbNewCfgVirtServicesTablePbindRport_Disabled    SlbNewCfgVirtServicesTablePbindRport = 2
	SlbNewCfgVirtServicesTablePbindRport_Unsupported SlbNewCfgVirtServicesTablePbindRport = 2147483647
)

type SlbNewCfgVirtServicesTableEgressPip int32

const (
	SlbNewCfgVirtServicesTableEgressPip_Enabled     SlbNewCfgVirtServicesTableEgressPip = 1
	SlbNewCfgVirtServicesTableEgressPip_Disabled    SlbNewCfgVirtServicesTableEgressPip = 2
	SlbNewCfgVirtServicesTableEgressPip_Unsupported SlbNewCfgVirtServicesTableEgressPip = 2147483647
)

type SlbNewCfgVirtServicesTableCookieDname int32

const (
	SlbNewCfgVirtServicesTableCookieDname_Enabled     SlbNewCfgVirtServicesTableCookieDname = 1
	SlbNewCfgVirtServicesTableCookieDname_Disabled    SlbNewCfgVirtServicesTableCookieDname = 2
	SlbNewCfgVirtServicesTableCookieDname_Unsupported SlbNewCfgVirtServicesTableCookieDname = 2147483647
)

type SlbNewCfgVirtServicesTableWts int32

const (
	SlbNewCfgVirtServicesTableWts_Enabled     SlbNewCfgVirtServicesTableWts = 1
	SlbNewCfgVirtServicesTableWts_Disabled    SlbNewCfgVirtServicesTableWts = 2
	SlbNewCfgVirtServicesTableWts_Unsupported SlbNewCfgVirtServicesTableWts = 2147483647
)

type SlbNewCfgVirtServicesTableUhash int32

const (
	SlbNewCfgVirtServicesTableUhash_Enabled     SlbNewCfgVirtServicesTableUhash = 1
	SlbNewCfgVirtServicesTableUhash_Disabled    SlbNewCfgVirtServicesTableUhash = 2
	SlbNewCfgVirtServicesTableUhash_Unsupported SlbNewCfgVirtServicesTableUhash = 2147483647
)

type SlbNewCfgVirtServicesTableSdpNat int32

const (
	SlbNewCfgVirtServicesTableSdpNat_Enabled     SlbNewCfgVirtServicesTableSdpNat = 1
	SlbNewCfgVirtServicesTableSdpNat_Disabled    SlbNewCfgVirtServicesTableSdpNat = 2
	SlbNewCfgVirtServicesTableSdpNat_Unsupported SlbNewCfgVirtServicesTableSdpNat = 2147483647
)

type SlbNewCfgVirtServicesTableSessionMirror int32

const (
	SlbNewCfgVirtServicesTableSessionMirror_Enabled     SlbNewCfgVirtServicesTableSessionMirror = 1
	SlbNewCfgVirtServicesTableSessionMirror_Disabled    SlbNewCfgVirtServicesTableSessionMirror = 2
	SlbNewCfgVirtServicesTableSessionMirror_Unsupported SlbNewCfgVirtServicesTableSessionMirror = 2147483647
)

type SlbNewCfgVirtServicesTableSoftGrid int32

const (
	SlbNewCfgVirtServicesTableSoftGrid_Enabled     SlbNewCfgVirtServicesTableSoftGrid = 1
	SlbNewCfgVirtServicesTableSoftGrid_Disabled    SlbNewCfgVirtServicesTableSoftGrid = 2
	SlbNewCfgVirtServicesTableSoftGrid_Unsupported SlbNewCfgVirtServicesTableSoftGrid = 2147483647
)

type SlbNewCfgVirtServicesTableConnPooling int32

const (
	SlbNewCfgVirtServicesTableConnPooling_Enabled     SlbNewCfgVirtServicesTableConnPooling = 1
	SlbNewCfgVirtServicesTableConnPooling_Disabled    SlbNewCfgVirtServicesTableConnPooling = 2
	SlbNewCfgVirtServicesTableConnPooling_Unsupported SlbNewCfgVirtServicesTableConnPooling = 2147483647
)

type SlbNewCfgVirtServicesTableProxyIpMode int32

const (
	SlbNewCfgVirtServicesTableProxyIpMode_Ingress     SlbNewCfgVirtServicesTableProxyIpMode = 0
	SlbNewCfgVirtServicesTableProxyIpMode_Egress      SlbNewCfgVirtServicesTableProxyIpMode = 1
	SlbNewCfgVirtServicesTableProxyIpMode_Address     SlbNewCfgVirtServicesTableProxyIpMode = 2
	SlbNewCfgVirtServicesTableProxyIpMode_Nwclss      SlbNewCfgVirtServicesTableProxyIpMode = 3
	SlbNewCfgVirtServicesTableProxyIpMode_Disable     SlbNewCfgVirtServicesTableProxyIpMode = 4
	SlbNewCfgVirtServicesTableProxyIpMode_Unsupported SlbNewCfgVirtServicesTableProxyIpMode = 2147483647
)

type SlbNewCfgVirtServicesTableProxyIpPersistency int32

const (
	SlbNewCfgVirtServicesTableProxyIpPersistency_Disable     SlbNewCfgVirtServicesTableProxyIpPersistency = 0
	SlbNewCfgVirtServicesTableProxyIpPersistency_Client      SlbNewCfgVirtServicesTableProxyIpPersistency = 1
	SlbNewCfgVirtServicesTableProxyIpPersistency_Host        SlbNewCfgVirtServicesTableProxyIpPersistency = 2
	SlbNewCfgVirtServicesTableProxyIpPersistency_Unsupported SlbNewCfgVirtServicesTableProxyIpPersistency = 2147483647
)

type SlbNewCfgVirtServicesTableProxyIpNWclassPersistency int32

const (
	SlbNewCfgVirtServicesTableProxyIpNWclassPersistency_Disable     SlbNewCfgVirtServicesTableProxyIpNWclassPersistency = 0
	SlbNewCfgVirtServicesTableProxyIpNWclassPersistency_Client      SlbNewCfgVirtServicesTableProxyIpNWclassPersistency = 1
	SlbNewCfgVirtServicesTableProxyIpNWclassPersistency_Unsupported SlbNewCfgVirtServicesTableProxyIpNWclassPersistency = 2147483647
)

type SlbNewCfgVirtServicesTableApm int32

const (
	SlbNewCfgVirtServicesTableApm_Enabled     SlbNewCfgVirtServicesTableApm = 1
	SlbNewCfgVirtServicesTableApm_Disabled    SlbNewCfgVirtServicesTableApm = 2
	SlbNewCfgVirtServicesTableApm_Unsupported SlbNewCfgVirtServicesTableApm = 2147483647
)

type SlbNewCfgVirtServicesTableClsRST int32

const (
	SlbNewCfgVirtServicesTableClsRST_Enabled     SlbNewCfgVirtServicesTableClsRST = 1
	SlbNewCfgVirtServicesTableClsRST_Disabled    SlbNewCfgVirtServicesTableClsRST = 2
	SlbNewCfgVirtServicesTableClsRST_Unsupported SlbNewCfgVirtServicesTableClsRST = 2147483647
)

type SlbNewCfgVirtServicesTableNonHTTP int32

const (
	SlbNewCfgVirtServicesTableNonHTTP_Enabled     SlbNewCfgVirtServicesTableNonHTTP = 1
	SlbNewCfgVirtServicesTableNonHTTP_Disabled    SlbNewCfgVirtServicesTableNonHTTP = 2
	SlbNewCfgVirtServicesTableNonHTTP_Unsupported SlbNewCfgVirtServicesTableNonHTTP = 2147483647
)

type SlbNewCfgVirtServicesTableParams struct {
	// The number of the virtual server.
	ServIndex int32 `json:"ServIndex,omitempty"`
	// The service index. This has no external meaning
	Index int32 `json:"Index,omitempty"`
	// The layer4 virtual port number of the service.
	VirtPort uint64 `json:"VirtPort,omitempty"`
	// The real server group number for this service.
	RealGroup int32 `json:"RealGroup,omitempty"`
	// The layer4 real port number of the service.
	RealPort uint64 `json:"RealPort,omitempty"`
	// Set protocol for the virtual service to
	// UDP or TCP or SCTP or tcpAndUdp or stateless.
	// tcpAndUdp is applicable only to ip service.
	UDPBalance SlbNewCfgVirtServicesTableUDPBalance `json:"UDPBalance,omitempty"`
	// The host name of the virtual service.
	Hname string `json:"Hname,omitempty"`
	// The BWM contract number for this service.
	BwmContract int32 `json:"BwmContract,omitempty"`
	// Enable or disable direct server return feature. To translate only
	// MAC addresses in performing server load balancing when enabled.
	// This allow servers to return directly to client since IP addresses
	// have not been changed.
	DirServerRtn SlbNewCfgVirtServicesTableDirServerRtn `json:"DirServerRtn,omitempty"`
	// Select RTSP URL load balancing type.
	RtspUrlParse SlbNewCfgVirtServicesTableRtspUrlParse `json:"RtspUrlParse,omitempty"`
	// Enable/disable/forceproxy delayed binding.
	DBind SlbNewCfgVirtServicesTableDBind `json:"DBind,omitempty"`
	// Enable or Disable the ftp parsing for the virtual service.
	FtpParsing SlbNewCfgVirtServicesTableFtpParsing `json:"FtpParsing,omitempty"`
	// Enable or disable remapping UDP server fragments
	RemapUDPFrags SlbNewCfgVirtServicesTableRemapUDPFrags `json:"RemapUDPFrags,omitempty"`
	// Enable or disable DNS query load balancing.
	DnsSlb SlbNewCfgVirtServicesTableDnsSlb `json:"DnsSlb,omitempty"`
	// The number of cookie search response count.
	ResponseCount uint32 `json:"ResponseCount,omitempty"`
	// Enable or disable persistent bindings for the virtual port.
	PBind SlbNewCfgVirtServicesTablePBind `json:"PBind,omitempty"`
	// The cookie name of the virtual server used for cookie load balance.
	Cname string `json:"Cname,omitempty"`
	// The starting byte offset of the cookie value.
	Coffset uint64 `json:"Coffset,omitempty"`
	// The number of bytes to extract from the cookie value.
	Clength uint64 `json:"Clength,omitempty"`
	// Enable or disable cookie search in URI
	UriCookie SlbNewCfgVirtServicesTableUriCookie `json:"UriCookie,omitempty"`
	// The cookie expire of the virtual server used for insert cookie load
	// balance depending on the mode it has the following format
	// <MM/dd/yy[@hh:mm]> absolute mode or <days[:hours[:minutes]]>
	// for relative mode.
	CExpire string `json:"CExpire,omitempty"`
	// Select cookie persistence mode. Mode disabled(4) not supported on Alteon
	CookieMode SlbNewCfgVirtServicesTableCookieMode `json:"CookieMode,omitempty"`
	// Select HTTP server loadbalancing for the virtual port.
	HttpSlb SlbNewCfgVirtServicesTableHttpSlb `json:"HttpSlb,omitempty"`
	// Select HTTP server loadbalancing for the virtual port.
	HttpSlbOption SlbNewCfgVirtServicesTableHttpSlbOption `json:"HttpSlbOption,omitempty"`
	// Select HTTP server loadbalancing for the virtual port.
	HttpSlb2 SlbNewCfgVirtServicesTableHttpSlb2 `json:"HttpSlb2,omitempty"`
	// The HTTP header name of the virtual server.
	HttpHdrName string `json:"HttpHdrName,omitempty"`
	// The number of bytes used to hash onto server, A zero means
	// URL hashing disabled.
	UrlHashLen uint32 `json:"UrlHashLen,omitempty"`
	// When set to the value of 2 (delete), the entire row is deleted.
	// When read, other(1) is returned. Setting the value to anything
	// other than 2(delete) has no effect on the state of the row.
	Delete SlbNewCfgVirtServicesTableDelete `json:"Delete,omitempty"`
	// Enable or disable DAM for this service.
	Direct SlbNewCfgVirtServicesTableDirect `json:"Direct,omitempty"`
	// Set hash parameter.
	Thash SlbNewCfgVirtServicesTableThash `json:"Thash,omitempty"`
	// Enable or disable LDAP Server Reset
	Ldapreset SlbNewCfgVirtServicesTableLdapreset `json:"Ldapreset,omitempty"`
	// Enable or disable LDAP Server load balancing
	Ldapslb SlbNewCfgVirtServicesTableLdapslb `json:"Ldapslb,omitempty"`
	// Enable/disable SIP load balancing.
	Sip SlbNewCfgVirtServicesTableSip `json:"Sip,omitempty"`
	// Enable/disable X-Forwarded-For for proxy mode.
	XForwardedFor SlbNewCfgVirtServicesTableXForwardedFor `json:"XForwardedFor,omitempty"`
	// Enable/disable HTTP/HTTPS redirect for GSLB.
	HttpRedir SlbNewCfgVirtServicesTableHttpRedir `json:"HttpRedir,omitempty"`
	// Enable or disable use of rport in the session lookup for a persistent
	// session.
	PbindRport SlbNewCfgVirtServicesTablePbindRport `json:"PbindRport,omitempty"`
	// Enable/disable pip selection based on egress port/vlan.
	EgressPip SlbNewCfgVirtServicesTableEgressPip `json:"EgressPip,omitempty"`
	// Select dname for insert cookie persistence mode.
	CookieDname SlbNewCfgVirtServicesTableCookieDname `json:"CookieDname,omitempty"`
	// Enable or disable WTS loadbalancing and persistence.
	Wts SlbNewCfgVirtServicesTableWts `json:"Wts,omitempty"`
	// Enable when there is no Session Directory server.
	Uhash SlbNewCfgVirtServicesTableUhash `json:"Uhash,omitempty"`
	// The maximum number of minutes an inactive connection remains open.
	TimeOut uint32 `json:"TimeOut,omitempty"`
	// Enable/disable SIP Media portal NAT.
	SdpNat SlbNewCfgVirtServicesTableSdpNat `json:"SdpNat,omitempty"`
	// Enable/disable session mirroring.
	SessionMirror SlbNewCfgVirtServicesTableSessionMirror `json:"SessionMirror,omitempty"`
	// Enable/disable softgrid load balancing.
	SoftGrid SlbNewCfgVirtServicesTableSoftGrid `json:"SoftGrid,omitempty"`
	// Enable/disable connection pooling for HTTP traffic.
	ConnPooling SlbNewCfgVirtServicesTableConnPooling `json:"ConnPooling,omitempty"`
	// The maximum number of minutes a persistent session should exist.
	PersistentTimeOut uint32 `json:"PersistentTimeOut,omitempty"`
	// Set the Proxy IP mode, default is ingress(0).
	// 		Changing from address(2) to any other mode will clear the configured IPv4/IPv6 address,prefix and persistancy.
	// 		Changing from  nwclass(3) to any other mode will clear the configured NWclass and NWpersistancy.
	ProxyIpMode SlbNewCfgVirtServicesTableProxyIpMode `json:"ProxyIpMode,omitempty"`
	// This object ID allows configuring IPv4 PIP address.
	// 		This object ID can be set only if slbNewCfgVirtServiceProxyIpMode is address else return failure.
	// 		Returns 0 when slbNewCfgVirtServiceProxyIpMode is not set to address.
	// 		When a subnet is configured user has the ability to select PIP persistency mode.
	// 		Persistency is relevant only if either IPv4 PIP or IPv6 PIP (or both) are configured as
	// 		subnet.
	// 		If neither of the addresses (v4 & v6) are configured or are subnets, the persistency
	// 		value is disable.
	ProxyIpAddr string `json:"ProxyIpAddr,omitempty"`
	// This object ID allows configuring IPv4 PIP Mask.
	// 		Returns 0 when slbNewCfgVirtServiceProxyIpMode is not set to address.
	// 		When a subnet is configured user has the ability to select PIP persistency mode.
	// 		Persistency is relevant only if either IPv4 PIP or IPv6 PIP (or both) are configured as
	// 		subnet.
	// 		If neither of the addresses (v4 & v6) are configured or are subnets, the persistency
	// 		value is disable.
	ProxyIpMask string `json:"ProxyIpMask,omitempty"`
	// This object ID allows configuring IPv6 PIP address.
	// 		Returns 0 when slbNewCfgVirtServiceProxyIpMode is not set to address.
	// 		When a subnet is configured user has the ability to select PIP persistency mode.
	// 		Persistency is relevant only if either IPv4 PIP or IPv6 PIP (or both) are configured as
	// 		subnet.
	// 		Address should be 4-byte hexadecimal colon notation.
	// 		Valid IPv6 address should be in any of the following forms
	// 		       xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx or
	// 		       xxxx::xxxx:xxxx:xxxx:xxxx or ::xxxx.
	// 		If neither of the addresses (v4 & v6) are configured or are subnets, the persistency
	// 		value is disable.
	ProxyIpv6Addr string `json:"ProxyIpv6Addr,omitempty"`
	// This object ID allows configuring IPv6 PIP Mask.
	// 		Returns 0 when slbNewCfgVirtServiceProxyIpMode is not set to address.
	// 		When a subnet is configured user has the ability to select PIP persistency mode.
	// 		Persistency is relevant only if either IPv4 PIP or IPv6 PIP (or both) are configured as
	// 		subnet.
	// 		If neither of the addresses (v4 & v6) are configured or are subnets, the persistency
	// 		value is disable.
	ProxyIpv6Prefix uint32 `json:"ProxyIpv6Prefix,omitempty"`
	// When a subnet is configured user has the ability to select PIP persistency mode.
	// 		Returns 0 when slbNewCfgVirtServiceProxyIpMode is not set to address.
	// 		Persistency is relevant only if either IPv4 PIP or IPv6 PIP (or both) are configured as
	// 		subnet.
	// 		If neither of the addresses (v4 & v6) are configured or are subnets, the persistency
	// 		value is disable.
	ProxyIpPersistency SlbNewCfgVirtServicesTableProxyIpPersistency `json:"ProxyIpPersistency,omitempty"`
	// This object ID allows configuring IPv4 Network Class as PIP and PIP persistency mode.
	// 		Returns empty string when slbNewCfgVirtServiceProxyIpMode is not set to nwclass.
	// 		Persistency is relevant only if either IPv4 or IPv6 class (or both) are configured.
	// 		If neither of the classes (v4 & v6) are configured, the persistency configuration value is disable.
	ProxyIpNWclass string `json:"ProxyIpNWclass,omitempty"`
	// This object ID allows configuring IPv6 Network Class as PIP and PIP persistency mode.
	// 		Returns empty string when slbNewCfgVirtServiceProxyIpMode is not set to nwclass.
	// 		Persistency is relevant only if either IPv4 or IPv6 class (or both) are configured.
	// 		If neither of the classes (v4 & v6) are configured, the persistency configuration value is disable.
	ProxyIpv6NWclass string `json:"ProxyIpv6NWclass,omitempty"`
	// This object ID allows configuring Network Class PIP persistency mode.
	// 		Returns 0 when slbNewCfgVirtServiceProxyIpMode is not set to nwclass.
	// 		Persistency is relevant only if either IPv4 PIP or IPv6 PIP (or both) are configured as
	// 		subnet.
	// 		If neither of the addresses (v4 & v6) are configured or are subnets, the persistency
	// 		value is disable.
	ProxyIpNWclassPersistency SlbNewCfgVirtServicesTableProxyIpNWclassPersistency `json:"ProxyIpNWclassPersistency,omitempty"`
	// Set length for slb service sip hashing (4- 256 bytes)
	HashLen uint32 `json:"HashLen,omitempty"`
	// Enable/disable apm.
	Apm SlbNewCfgVirtServicesTableApm `json:"Apm,omitempty"`
	// Enable or disable Send RST on connection close.
	ClsRST SlbNewCfgVirtServicesTableClsRST `json:"ClsRST,omitempty"`
	// Enable or disable Send non-HTTP traffic.
	NonHTTP SlbNewCfgVirtServicesTableNonHTTP `json:"NonHTTP,omitempty"`
}

func (p SlbNewCfgVirtServicesTableParams) iMABean() {}
