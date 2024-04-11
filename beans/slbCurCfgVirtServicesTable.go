package beans

import (
	"fmt"
	"reflect"
)

// SlbCurCfgVirtServicesTable The table of virtual services.
// Note:This mib is not supported for VX instance of Virtualization.
type SlbCurCfgVirtServicesTable struct {
	// The number of the virtual server.
	SlbCurCfgVirtServIndex int32
	// The service index. This has no external meaning
	SlbCurCfgVirtServiceIndex int32
	Params                    *SlbCurCfgVirtServicesTableParams
}

func NewSlbCurCfgVirtServicesTableList() *SlbCurCfgVirtServicesTable {
	return &SlbCurCfgVirtServicesTable{}
}

func NewSlbCurCfgVirtServicesTable(
	slbCurCfgVirtServIndex int32,
	slbCurCfgVirtServiceIndex int32,
	params *SlbCurCfgVirtServicesTableParams,
) *SlbCurCfgVirtServicesTable {
	return &SlbCurCfgVirtServicesTable{
		SlbCurCfgVirtServIndex:    slbCurCfgVirtServIndex,
		SlbCurCfgVirtServiceIndex: slbCurCfgVirtServiceIndex,
		Params:                    params,
	}
}

func (c *SlbCurCfgVirtServicesTable) Name() string {
	return "SlbCurCfgVirtServicesTable"
}

func (c *SlbCurCfgVirtServicesTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbCurCfgVirtServicesTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbCurCfgVirtServicesTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbCurCfgVirtServIndex).IsZero() &&
		reflect.ValueOf(c.SlbCurCfgVirtServiceIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbCurCfgVirtServIndex) + "/" + fmt.Sprint(c.SlbCurCfgVirtServiceIndex)
}

type SlbCurCfgVirtServicesTableUDPBalance int32

const (
	SlbCurCfgVirtServicesTableUDPBalance_Udp         SlbCurCfgVirtServicesTableUDPBalance = 2
	SlbCurCfgVirtServicesTableUDPBalance_Tcp         SlbCurCfgVirtServicesTableUDPBalance = 3
	SlbCurCfgVirtServicesTableUDPBalance_Stateless   SlbCurCfgVirtServicesTableUDPBalance = 4
	SlbCurCfgVirtServicesTableUDPBalance_TcpAndUdp   SlbCurCfgVirtServicesTableUDPBalance = 5
	SlbCurCfgVirtServicesTableUDPBalance_Sctp        SlbCurCfgVirtServicesTableUDPBalance = 6
	SlbCurCfgVirtServicesTableUDPBalance_Unsupported SlbCurCfgVirtServicesTableUDPBalance = 2147483647
)

type SlbCurCfgVirtServicesTableDirServerRtn int32

const (
	SlbCurCfgVirtServicesTableDirServerRtn_Enabled     SlbCurCfgVirtServicesTableDirServerRtn = 1
	SlbCurCfgVirtServicesTableDirServerRtn_Disabled    SlbCurCfgVirtServicesTableDirServerRtn = 2
	SlbCurCfgVirtServicesTableDirServerRtn_Unsupported SlbCurCfgVirtServicesTableDirServerRtn = 2147483647
)

type SlbCurCfgVirtServicesTableRtspUrlParse int32

const (
	SlbCurCfgVirtServicesTableRtspUrlParse_None         SlbCurCfgVirtServicesTableRtspUrlParse = 1
	SlbCurCfgVirtServicesTableRtspUrlParse_L4hash       SlbCurCfgVirtServicesTableRtspUrlParse = 2
	SlbCurCfgVirtServicesTableRtspUrlParse_Hash         SlbCurCfgVirtServicesTableRtspUrlParse = 3
	SlbCurCfgVirtServicesTableRtspUrlParse_PatternMatch SlbCurCfgVirtServicesTableRtspUrlParse = 4
	SlbCurCfgVirtServicesTableRtspUrlParse_Unsupported  SlbCurCfgVirtServicesTableRtspUrlParse = 2147483647
)

type SlbCurCfgVirtServicesTableDBind int32

const (
	SlbCurCfgVirtServicesTableDBind_Enabled     SlbCurCfgVirtServicesTableDBind = 1
	SlbCurCfgVirtServicesTableDBind_Disabled    SlbCurCfgVirtServicesTableDBind = 2
	SlbCurCfgVirtServicesTableDBind_Forceproxy  SlbCurCfgVirtServicesTableDBind = 3
	SlbCurCfgVirtServicesTableDBind_Unsupported SlbCurCfgVirtServicesTableDBind = 2147483647
)

type SlbCurCfgVirtServicesTableFtpParsing int32

const (
	SlbCurCfgVirtServicesTableFtpParsing_Enabled     SlbCurCfgVirtServicesTableFtpParsing = 1
	SlbCurCfgVirtServicesTableFtpParsing_Disabled    SlbCurCfgVirtServicesTableFtpParsing = 2
	SlbCurCfgVirtServicesTableFtpParsing_Unsupported SlbCurCfgVirtServicesTableFtpParsing = 2147483647
)

type SlbCurCfgVirtServicesTableRemapUDPFrags int32

const (
	SlbCurCfgVirtServicesTableRemapUDPFrags_Enabled     SlbCurCfgVirtServicesTableRemapUDPFrags = 1
	SlbCurCfgVirtServicesTableRemapUDPFrags_Disabled    SlbCurCfgVirtServicesTableRemapUDPFrags = 2
	SlbCurCfgVirtServicesTableRemapUDPFrags_Unsupported SlbCurCfgVirtServicesTableRemapUDPFrags = 2147483647
)

type SlbCurCfgVirtServicesTableDnsSlb int32

const (
	SlbCurCfgVirtServicesTableDnsSlb_Enabled     SlbCurCfgVirtServicesTableDnsSlb = 1
	SlbCurCfgVirtServicesTableDnsSlb_Disabled    SlbCurCfgVirtServicesTableDnsSlb = 2
	SlbCurCfgVirtServicesTableDnsSlb_Unsupported SlbCurCfgVirtServicesTableDnsSlb = 2147483647
)

type SlbCurCfgVirtServicesTablePBind int32

const (
	SlbCurCfgVirtServicesTablePBind_Clientip    SlbCurCfgVirtServicesTablePBind = 2
	SlbCurCfgVirtServicesTablePBind_Disabled    SlbCurCfgVirtServicesTablePBind = 3
	SlbCurCfgVirtServicesTablePBind_Sslid       SlbCurCfgVirtServicesTablePBind = 4
	SlbCurCfgVirtServicesTablePBind_Cookie      SlbCurCfgVirtServicesTablePBind = 5
	SlbCurCfgVirtServicesTablePBind_Unsupported SlbCurCfgVirtServicesTablePBind = 2147483647
)

type SlbCurCfgVirtServicesTableUriCookie int32

const (
	SlbCurCfgVirtServicesTableUriCookie_Enabled     SlbCurCfgVirtServicesTableUriCookie = 1
	SlbCurCfgVirtServicesTableUriCookie_Disabled    SlbCurCfgVirtServicesTableUriCookie = 2
	SlbCurCfgVirtServicesTableUriCookie_Unsupported SlbCurCfgVirtServicesTableUriCookie = 2147483647
)

type SlbCurCfgVirtServicesTableCookieMode int32

const (
	SlbCurCfgVirtServicesTableCookieMode_Rewrite     SlbCurCfgVirtServicesTableCookieMode = 1
	SlbCurCfgVirtServicesTableCookieMode_Passive     SlbCurCfgVirtServicesTableCookieMode = 2
	SlbCurCfgVirtServicesTableCookieMode_Insert      SlbCurCfgVirtServicesTableCookieMode = 3
	SlbCurCfgVirtServicesTableCookieMode_Disabled    SlbCurCfgVirtServicesTableCookieMode = 4
	SlbCurCfgVirtServicesTableCookieMode_Unsupported SlbCurCfgVirtServicesTableCookieMode = 2147483647
)

type SlbCurCfgVirtServicesTableHttpSlb int32

const (
	SlbCurCfgVirtServicesTableHttpSlb_Disabled    SlbCurCfgVirtServicesTableHttpSlb = 1
	SlbCurCfgVirtServicesTableHttpSlb_Urlslb      SlbCurCfgVirtServicesTableHttpSlb = 2
	SlbCurCfgVirtServicesTableHttpSlb_Urlhash     SlbCurCfgVirtServicesTableHttpSlb = 3
	SlbCurCfgVirtServicesTableHttpSlb_Cookie      SlbCurCfgVirtServicesTableHttpSlb = 4
	SlbCurCfgVirtServicesTableHttpSlb_Host        SlbCurCfgVirtServicesTableHttpSlb = 5
	SlbCurCfgVirtServicesTableHttpSlb_Browser     SlbCurCfgVirtServicesTableHttpSlb = 6
	SlbCurCfgVirtServicesTableHttpSlb_Others      SlbCurCfgVirtServicesTableHttpSlb = 7
	SlbCurCfgVirtServicesTableHttpSlb_Headerhash  SlbCurCfgVirtServicesTableHttpSlb = 8
	SlbCurCfgVirtServicesTableHttpSlb_Version     SlbCurCfgVirtServicesTableHttpSlb = 9
	SlbCurCfgVirtServicesTableHttpSlb_Unsupported SlbCurCfgVirtServicesTableHttpSlb = 2147483647
)

type SlbCurCfgVirtServicesTableHttpSlbOption int32

const (
	SlbCurCfgVirtServicesTableHttpSlbOption_And         SlbCurCfgVirtServicesTableHttpSlbOption = 1
	SlbCurCfgVirtServicesTableHttpSlbOption_Or          SlbCurCfgVirtServicesTableHttpSlbOption = 2
	SlbCurCfgVirtServicesTableHttpSlbOption_None        SlbCurCfgVirtServicesTableHttpSlbOption = 3
	SlbCurCfgVirtServicesTableHttpSlbOption_Unsupported SlbCurCfgVirtServicesTableHttpSlbOption = 2147483647
)

type SlbCurCfgVirtServicesTableHttpSlb2 int32

const (
	SlbCurCfgVirtServicesTableHttpSlb2_Disabled    SlbCurCfgVirtServicesTableHttpSlb2 = 1
	SlbCurCfgVirtServicesTableHttpSlb2_Urlslb      SlbCurCfgVirtServicesTableHttpSlb2 = 2
	SlbCurCfgVirtServicesTableHttpSlb2_Urlhash     SlbCurCfgVirtServicesTableHttpSlb2 = 3
	SlbCurCfgVirtServicesTableHttpSlb2_Cookie      SlbCurCfgVirtServicesTableHttpSlb2 = 4
	SlbCurCfgVirtServicesTableHttpSlb2_Host        SlbCurCfgVirtServicesTableHttpSlb2 = 5
	SlbCurCfgVirtServicesTableHttpSlb2_Browser     SlbCurCfgVirtServicesTableHttpSlb2 = 6
	SlbCurCfgVirtServicesTableHttpSlb2_Others      SlbCurCfgVirtServicesTableHttpSlb2 = 7
	SlbCurCfgVirtServicesTableHttpSlb2_Headerhash  SlbCurCfgVirtServicesTableHttpSlb2 = 8
	SlbCurCfgVirtServicesTableHttpSlb2_Version     SlbCurCfgVirtServicesTableHttpSlb2 = 9
	SlbCurCfgVirtServicesTableHttpSlb2_Unsupported SlbCurCfgVirtServicesTableHttpSlb2 = 2147483647
)

type SlbCurCfgVirtServicesTableDirect int32

const (
	SlbCurCfgVirtServicesTableDirect_Enabled     SlbCurCfgVirtServicesTableDirect = 1
	SlbCurCfgVirtServicesTableDirect_Disabled    SlbCurCfgVirtServicesTableDirect = 2
	SlbCurCfgVirtServicesTableDirect_Unsupported SlbCurCfgVirtServicesTableDirect = 2147483647
)

type SlbCurCfgVirtServicesTableThash int32

const (
	SlbCurCfgVirtServicesTableThash_Sip         SlbCurCfgVirtServicesTableThash = 1
	SlbCurCfgVirtServicesTableThash_SipSport    SlbCurCfgVirtServicesTableThash = 2
	SlbCurCfgVirtServicesTableThash_Unsupported SlbCurCfgVirtServicesTableThash = 2147483647
)

type SlbCurCfgVirtServicesTableLdapreset int32

const (
	SlbCurCfgVirtServicesTableLdapreset_Enabled     SlbCurCfgVirtServicesTableLdapreset = 1
	SlbCurCfgVirtServicesTableLdapreset_Disabled    SlbCurCfgVirtServicesTableLdapreset = 2
	SlbCurCfgVirtServicesTableLdapreset_Unsupported SlbCurCfgVirtServicesTableLdapreset = 2147483647
)

type SlbCurCfgVirtServicesTableLdapslb int32

const (
	SlbCurCfgVirtServicesTableLdapslb_Enabled     SlbCurCfgVirtServicesTableLdapslb = 1
	SlbCurCfgVirtServicesTableLdapslb_Disabled    SlbCurCfgVirtServicesTableLdapslb = 2
	SlbCurCfgVirtServicesTableLdapslb_Unsupported SlbCurCfgVirtServicesTableLdapslb = 2147483647
)

type SlbCurCfgVirtServicesTableSip int32

const (
	SlbCurCfgVirtServicesTableSip_Enabled     SlbCurCfgVirtServicesTableSip = 1
	SlbCurCfgVirtServicesTableSip_Disabled    SlbCurCfgVirtServicesTableSip = 2
	SlbCurCfgVirtServicesTableSip_Unsupported SlbCurCfgVirtServicesTableSip = 2147483647
)

type SlbCurCfgVirtServicesTableXForwardedFor int32

const (
	SlbCurCfgVirtServicesTableXForwardedFor_Enabled     SlbCurCfgVirtServicesTableXForwardedFor = 1
	SlbCurCfgVirtServicesTableXForwardedFor_Disabled    SlbCurCfgVirtServicesTableXForwardedFor = 2
	SlbCurCfgVirtServicesTableXForwardedFor_Unsupported SlbCurCfgVirtServicesTableXForwardedFor = 2147483647
)

type SlbCurCfgVirtServicesTableHttpRedir int32

const (
	SlbCurCfgVirtServicesTableHttpRedir_Enabled     SlbCurCfgVirtServicesTableHttpRedir = 1
	SlbCurCfgVirtServicesTableHttpRedir_Disabled    SlbCurCfgVirtServicesTableHttpRedir = 2
	SlbCurCfgVirtServicesTableHttpRedir_Unsupported SlbCurCfgVirtServicesTableHttpRedir = 2147483647
)

type SlbCurCfgVirtServicesTablePbindRport int32

const (
	SlbCurCfgVirtServicesTablePbindRport_Enabled     SlbCurCfgVirtServicesTablePbindRport = 1
	SlbCurCfgVirtServicesTablePbindRport_Disabled    SlbCurCfgVirtServicesTablePbindRport = 2
	SlbCurCfgVirtServicesTablePbindRport_Unsupported SlbCurCfgVirtServicesTablePbindRport = 2147483647
)

type SlbCurCfgVirtServicesTableEgressPip int32

const (
	SlbCurCfgVirtServicesTableEgressPip_Enabled     SlbCurCfgVirtServicesTableEgressPip = 1
	SlbCurCfgVirtServicesTableEgressPip_Disabled    SlbCurCfgVirtServicesTableEgressPip = 2
	SlbCurCfgVirtServicesTableEgressPip_Unsupported SlbCurCfgVirtServicesTableEgressPip = 2147483647
)

type SlbCurCfgVirtServicesTableCookieDname int32

const (
	SlbCurCfgVirtServicesTableCookieDname_Enabled     SlbCurCfgVirtServicesTableCookieDname = 1
	SlbCurCfgVirtServicesTableCookieDname_Disabled    SlbCurCfgVirtServicesTableCookieDname = 2
	SlbCurCfgVirtServicesTableCookieDname_Unsupported SlbCurCfgVirtServicesTableCookieDname = 2147483647
)

type SlbCurCfgVirtServicesTableWts int32

const (
	SlbCurCfgVirtServicesTableWts_Enabled     SlbCurCfgVirtServicesTableWts = 1
	SlbCurCfgVirtServicesTableWts_Disabled    SlbCurCfgVirtServicesTableWts = 2
	SlbCurCfgVirtServicesTableWts_Unsupported SlbCurCfgVirtServicesTableWts = 2147483647
)

type SlbCurCfgVirtServicesTableUhash int32

const (
	SlbCurCfgVirtServicesTableUhash_Enabled     SlbCurCfgVirtServicesTableUhash = 1
	SlbCurCfgVirtServicesTableUhash_Disabled    SlbCurCfgVirtServicesTableUhash = 2
	SlbCurCfgVirtServicesTableUhash_Unsupported SlbCurCfgVirtServicesTableUhash = 2147483647
)

type SlbCurCfgVirtServicesTableSdpNat int32

const (
	SlbCurCfgVirtServicesTableSdpNat_Enabled     SlbCurCfgVirtServicesTableSdpNat = 1
	SlbCurCfgVirtServicesTableSdpNat_Disabled    SlbCurCfgVirtServicesTableSdpNat = 2
	SlbCurCfgVirtServicesTableSdpNat_Unsupported SlbCurCfgVirtServicesTableSdpNat = 2147483647
)

type SlbCurCfgVirtServicesTableSessionMirror int32

const (
	SlbCurCfgVirtServicesTableSessionMirror_Enabled     SlbCurCfgVirtServicesTableSessionMirror = 1
	SlbCurCfgVirtServicesTableSessionMirror_Disabled    SlbCurCfgVirtServicesTableSessionMirror = 2
	SlbCurCfgVirtServicesTableSessionMirror_Unsupported SlbCurCfgVirtServicesTableSessionMirror = 2147483647
)

type SlbCurCfgVirtServicesTableSoftGrid int32

const (
	SlbCurCfgVirtServicesTableSoftGrid_Enabled     SlbCurCfgVirtServicesTableSoftGrid = 1
	SlbCurCfgVirtServicesTableSoftGrid_Disabled    SlbCurCfgVirtServicesTableSoftGrid = 2
	SlbCurCfgVirtServicesTableSoftGrid_Unsupported SlbCurCfgVirtServicesTableSoftGrid = 2147483647
)

type SlbCurCfgVirtServicesTableConnPooling int32

const (
	SlbCurCfgVirtServicesTableConnPooling_Enabled     SlbCurCfgVirtServicesTableConnPooling = 1
	SlbCurCfgVirtServicesTableConnPooling_Disabled    SlbCurCfgVirtServicesTableConnPooling = 2
	SlbCurCfgVirtServicesTableConnPooling_Unsupported SlbCurCfgVirtServicesTableConnPooling = 2147483647
)

type SlbCurCfgVirtServicesTableProxyIpMode int32

const (
	SlbCurCfgVirtServicesTableProxyIpMode_Ingress     SlbCurCfgVirtServicesTableProxyIpMode = 0
	SlbCurCfgVirtServicesTableProxyIpMode_Egress      SlbCurCfgVirtServicesTableProxyIpMode = 1
	SlbCurCfgVirtServicesTableProxyIpMode_Address     SlbCurCfgVirtServicesTableProxyIpMode = 2
	SlbCurCfgVirtServicesTableProxyIpMode_Nwclss      SlbCurCfgVirtServicesTableProxyIpMode = 3
	SlbCurCfgVirtServicesTableProxyIpMode_Disable     SlbCurCfgVirtServicesTableProxyIpMode = 4
	SlbCurCfgVirtServicesTableProxyIpMode_Unsupported SlbCurCfgVirtServicesTableProxyIpMode = 2147483647
)

type SlbCurCfgVirtServicesTableProxyIpPersistency int32

const (
	SlbCurCfgVirtServicesTableProxyIpPersistency_Disable     SlbCurCfgVirtServicesTableProxyIpPersistency = 0
	SlbCurCfgVirtServicesTableProxyIpPersistency_Client      SlbCurCfgVirtServicesTableProxyIpPersistency = 1
	SlbCurCfgVirtServicesTableProxyIpPersistency_Host        SlbCurCfgVirtServicesTableProxyIpPersistency = 2
	SlbCurCfgVirtServicesTableProxyIpPersistency_Unsupported SlbCurCfgVirtServicesTableProxyIpPersistency = 2147483647
)

type SlbCurCfgVirtServicesTableProxyIpNWclassPersistency int32

const (
	SlbCurCfgVirtServicesTableProxyIpNWclassPersistency_Disable     SlbCurCfgVirtServicesTableProxyIpNWclassPersistency = 0
	SlbCurCfgVirtServicesTableProxyIpNWclassPersistency_Client      SlbCurCfgVirtServicesTableProxyIpNWclassPersistency = 1
	SlbCurCfgVirtServicesTableProxyIpNWclassPersistency_Unsupported SlbCurCfgVirtServicesTableProxyIpNWclassPersistency = 2147483647
)

type SlbCurCfgVirtServicesTableApm int32

const (
	SlbCurCfgVirtServicesTableApm_Enabled     SlbCurCfgVirtServicesTableApm = 1
	SlbCurCfgVirtServicesTableApm_Disabled    SlbCurCfgVirtServicesTableApm = 2
	SlbCurCfgVirtServicesTableApm_Unsupported SlbCurCfgVirtServicesTableApm = 2147483647
)

type SlbCurCfgVirtServicesTableClsRST int32

const (
	SlbCurCfgVirtServicesTableClsRST_Enabled     SlbCurCfgVirtServicesTableClsRST = 1
	SlbCurCfgVirtServicesTableClsRST_Disabled    SlbCurCfgVirtServicesTableClsRST = 2
	SlbCurCfgVirtServicesTableClsRST_Unsupported SlbCurCfgVirtServicesTableClsRST = 2147483647
)

type SlbCurCfgVirtServicesTableNonHTTP int32

const (
	SlbCurCfgVirtServicesTableNonHTTP_Enabled     SlbCurCfgVirtServicesTableNonHTTP = 1
	SlbCurCfgVirtServicesTableNonHTTP_Disabled    SlbCurCfgVirtServicesTableNonHTTP = 2
	SlbCurCfgVirtServicesTableNonHTTP_Unsupported SlbCurCfgVirtServicesTableNonHTTP = 2147483647
)

type SlbCurCfgVirtServicesTableParams struct {
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
	UDPBalance SlbCurCfgVirtServicesTableUDPBalance `json:"UDPBalance,omitempty"`
	// The host name of the virtual service.
	Hname string `json:"Hname,omitempty"`
	// The BWM contract number for this service.
	BwmContract int32 `json:"BwmContract,omitempty"`
	// Enable or disable direct server return feature. To translate only
	// MAC addresses in performing server load balancing when enabled.
	// This allow servers to return directly to client since IP addresses
	// have not been changed.
	DirServerRtn SlbCurCfgVirtServicesTableDirServerRtn `json:"DirServerRtn,omitempty"`
	// Select RTSP URL load balancing type.
	RtspUrlParse SlbCurCfgVirtServicesTableRtspUrlParse `json:"RtspUrlParse,omitempty"`
	// Enable/disable/forceproxy delayed binding.
	DBind SlbCurCfgVirtServicesTableDBind `json:"DBind,omitempty"`
	// Enable or Disable the ftp parsing for the virtual service.
	FtpParsing SlbCurCfgVirtServicesTableFtpParsing `json:"FtpParsing,omitempty"`
	// Enable or disable remapping UDP server fragments
	RemapUDPFrags SlbCurCfgVirtServicesTableRemapUDPFrags `json:"RemapUDPFrags,omitempty"`
	// Enable or disable DNS query load balancing.
	DnsSlb SlbCurCfgVirtServicesTableDnsSlb `json:"DnsSlb,omitempty"`
	// The number of cookie search response count.
	ResponseCount uint32 `json:"ResponseCount,omitempty"`
	// Enable or disable persistent bindings for the virtual port.
	PBind SlbCurCfgVirtServicesTablePBind `json:"PBind,omitempty"`
	// The cookie name of the virtual server used for cookie load balance.
	Cname string `json:"Cname,omitempty"`
	// The starting byte offset of the cookie value.
	Coffset uint64 `json:"Coffset,omitempty"`
	// The number of bytes to extract from the cookie value.
	Clength uint64 `json:"Clength,omitempty"`
	// Enable or disable cookie search in URI
	UriCookie SlbCurCfgVirtServicesTableUriCookie `json:"UriCookie,omitempty"`
	// The cookie expire of the virtual server used for insert cookie load
	// balance depending on the mode it has the following format
	// <MM/dd/yy[@hh:mm]> absolute mode or <days[:hours[:minutes]]>
	// for relative mode.
	CExpire string `json:"CExpire,omitempty"`
	// Select cookie persistence mode.
	CookieMode SlbCurCfgVirtServicesTableCookieMode `json:"CookieMode,omitempty"`
	// Select HTTP server loadbalancing for the virtual port.
	HttpSlb SlbCurCfgVirtServicesTableHttpSlb `json:"HttpSlb,omitempty"`
	// Select HTTP server loadbalancing for the virtual port.
	HttpSlbOption SlbCurCfgVirtServicesTableHttpSlbOption `json:"HttpSlbOption,omitempty"`
	// Select HTTP server loadbalancing for the virtual port.
	HttpSlb2 SlbCurCfgVirtServicesTableHttpSlb2 `json:"HttpSlb2,omitempty"`
	// The HTTP header name of the virtual server.
	HttpHdrName string `json:"HttpHdrName,omitempty"`
	// The number of bytes used to hash onto server, A zero means
	// URL hashing disabled.
	UrlHashLen uint32 `json:"UrlHashLen,omitempty"`
	// Enable or disable DAM for this service.
	Direct SlbCurCfgVirtServicesTableDirect `json:"Direct,omitempty"`
	// Set hash parameter.
	Thash SlbCurCfgVirtServicesTableThash `json:"Thash,omitempty"`
	// Enable or disable LDAP Server Reset
	Ldapreset SlbCurCfgVirtServicesTableLdapreset `json:"Ldapreset,omitempty"`
	// Enable or disable LDAP Server load balancing
	Ldapslb SlbCurCfgVirtServicesTableLdapslb `json:"Ldapslb,omitempty"`
	// Enable/disable SIP load balancing.
	Sip SlbCurCfgVirtServicesTableSip `json:"Sip,omitempty"`
	// Enable/disable X-Forwarded-For for proxy mode.
	XForwardedFor SlbCurCfgVirtServicesTableXForwardedFor `json:"XForwardedFor,omitempty"`
	// Enable/disable HTTP/HTTPS redirect for GSLB.
	HttpRedir SlbCurCfgVirtServicesTableHttpRedir `json:"HttpRedir,omitempty"`
	// Enable or disable use of rport in the session lookup for a persistent
	// session.
	PbindRport SlbCurCfgVirtServicesTablePbindRport `json:"PbindRport,omitempty"`
	// Enable/disable pip selection based on egress port/vlan.
	EgressPip SlbCurCfgVirtServicesTableEgressPip `json:"EgressPip,omitempty"`
	// Select dname for insert cookie persistence mode.
	CookieDname SlbCurCfgVirtServicesTableCookieDname `json:"CookieDname,omitempty"`
	// Enable or disable WTS loadbalancing and persistence.
	Wts SlbCurCfgVirtServicesTableWts `json:"Wts,omitempty"`
	// Enable when there is no Session Directory server.
	Uhash SlbCurCfgVirtServicesTableUhash `json:"Uhash,omitempty"`
	// The maximum number of minutes an inactive connection remains open.
	TimeOut uint32 `json:"TimeOut,omitempty"`
	// Enable/disable SIP Media portal NAT.
	SdpNat SlbCurCfgVirtServicesTableSdpNat `json:"SdpNat,omitempty"`
	// Enable/disable session mirroring.
	SessionMirror SlbCurCfgVirtServicesTableSessionMirror `json:"SessionMirror,omitempty"`
	// Enable/disable softgrid load balancing.
	SoftGrid SlbCurCfgVirtServicesTableSoftGrid `json:"SoftGrid,omitempty"`
	// Enable/disable connection pooling for HTTP traffic.
	ConnPooling SlbCurCfgVirtServicesTableConnPooling `json:"ConnPooling,omitempty"`
	// The maximum number of minutes a persistent session should exist.
	PersistentTimeOut uint32 `json:"PersistentTimeOut,omitempty"`
	// the configured Proxy IP mode, default is ingress
	ProxyIpMode SlbCurCfgVirtServicesTableProxyIpMode `json:"ProxyIpMode,omitempty"`
	// This object ID shows current configured IPv4 PIP address.
	// 		Returns 0 when slbCurCfgVirtServiceProxyIpMode is not set to address.
	// 		If neither of the addresses (v4 & v6) are configured or are subnets, the persistency
	// 		value is disable.
	ProxyIpAddr string `json:"ProxyIpAddr,omitempty"`
	// This object ID  shows current configured  IPv4 PIP Mask.
	// 		Returns 0 when slbCurCfgVirtServiceProxyIpMode is not set to address.
	// 		If neither of the addresses (v4 & v6) are configured or are subnets, the persistency
	// 		value is disable.
	ProxyIpMask string `json:"ProxyIpMask,omitempty"`
	// This object ID shows current configured IPv6 PIP address.
	// 		Address should be 4-byte hexadecimal colon notation.
	// 		Valid IPv6 address should be in any of the following forms
	// 		       xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx or
	// 		       xxxx::xxxx:xxxx:xxxx:xxxx or ::xxxx.
	// 		Returns empty string when slbCurCfgVirtServiceProxyIpMode is not set to address.
	// 		If neither of the addresses (v4 & v6) are configured or are subnets, the persistency
	// 		value is disable.
	ProxyIpv6Addr string `json:"ProxyIpv6Addr,omitempty"`
	// This object ID shows current configured IPv6 PIP Mask.
	// 		Returns 0 when slbCurCfgVirtServiceProxyIpMode is not set to address.
	// 		If neither of the addresses (v4 & v6) are configured or are subnets, the persistency
	// 		value is disable.
	ProxyIpv6Prefix uint32 `json:"ProxyIpv6Prefix,omitempty"`
	// This object ID shows current configured virtual service proxy IP persistency.
	// 		Returns 0 when slbCurCfgVirtServiceProxyIpMode is not set to address.
	// 		If neither of the addresses (v4 & v6) are configured or are subnets, the persistency
	// 		value is disable.
	ProxyIpPersistency SlbCurCfgVirtServicesTableProxyIpPersistency `json:"ProxyIpPersistency,omitempty"`
	// This object ID ows current configured IPv4 Network Class as PIP .
	// 		Returns empty string when slbCurCfgVirtServiceProxyIpMode is not set to nwclass.
	// 		Persistency is relevant only if either IPv4 or IPv6 class (or both) are configured.
	// 		If neither of the classes (v4 & v6) are configured, the persistency configuration value is disable.
	ProxyIpNWclass string `json:"ProxyIpNWclass,omitempty"`
	// This object ID shows current configured IPv6 Network Class as PIP .
	// 		Returns empty string when slbCurCfgVirtServiceProxyIpMode is not set to nwclass.
	// 		Persistency is relevant only if either IPv4 or IPv6 class (or both) are configured.
	// 		If neither of the classes (v4 & v6) are configured, the persistency configuration value is disable.
	ProxyIpv6NWclass string `json:"ProxyIpv6NWclass,omitempty"`
	// This object ID shows current configured Network Class PIP persistency mode.
	// 		Returns 0 when slbCurCfgVirtServiceProxyIpMode is not set to nwclass.
	// 		Persistency is relevant only if either IPv4 PIP or IPv6 PIP (or both) are configured as
	// 		subnet.
	// 		If neither of the classes (v4 & v6) are configured, the persistency configuration value is disable.
	ProxyIpNWclassPersistency SlbCurCfgVirtServicesTableProxyIpNWclassPersistency `json:"ProxyIpNWclassPersistency,omitempty"`
	// Set length for slb service sip hashing (4- 256 bytes)
	HashLen uint32 `json:"HashLen,omitempty"`
	// Enable/disable apm.
	Apm SlbCurCfgVirtServicesTableApm `json:"Apm,omitempty"`
	// Enable or disable Send RST on connection close.
	ClsRST SlbCurCfgVirtServicesTableClsRST `json:"ClsRST,omitempty"`
	// Enable or disable Send non-HTTP traffic.
	NonHTTP SlbCurCfgVirtServicesTableNonHTTP `json:"NonHTTP,omitempty"`
}

func (p SlbCurCfgVirtServicesTableParams) iMABean() {}
