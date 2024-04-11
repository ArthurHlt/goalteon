package beans

import (
	"fmt"
	"reflect"
)

// SlbNewCfgEnhVirtServicesTable The table of virtual Services.
// Note:This mib is not supported for VX instance of Virtualization.
type SlbNewCfgEnhVirtServicesTable struct {
	// The number of the virtual server.
	SlbNewCfgEnhVirtServIndex string
	// The service index. This has no external meaning
	SlbNewCfgEnhVirtServiceIndex int32
	Params                       *SlbNewCfgEnhVirtServicesTableParams
}

func NewSlbNewCfgEnhVirtServicesTableList() *SlbNewCfgEnhVirtServicesTable {
	return &SlbNewCfgEnhVirtServicesTable{}
}

func NewSlbNewCfgEnhVirtServicesTable(
	slbNewCfgEnhVirtServIndex string,
	slbNewCfgEnhVirtServiceIndex int32,
	params *SlbNewCfgEnhVirtServicesTableParams,
) *SlbNewCfgEnhVirtServicesTable {
	return &SlbNewCfgEnhVirtServicesTable{
		SlbNewCfgEnhVirtServIndex:    slbNewCfgEnhVirtServIndex,
		SlbNewCfgEnhVirtServiceIndex: slbNewCfgEnhVirtServiceIndex,
		Params:                       params,
	}
}

func (c *SlbNewCfgEnhVirtServicesTable) Name() string {
	return "SlbNewCfgEnhVirtServicesTable"
}

func (c *SlbNewCfgEnhVirtServicesTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbNewCfgEnhVirtServicesTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbNewCfgEnhVirtServicesTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbNewCfgEnhVirtServIndex).IsZero() &&
		reflect.ValueOf(c.SlbNewCfgEnhVirtServiceIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbNewCfgEnhVirtServIndex) + "/" + fmt.Sprint(c.SlbNewCfgEnhVirtServiceIndex)
}

type SlbNewCfgEnhVirtServicesTableUDPBalance int32

const (
	SlbNewCfgEnhVirtServicesTableUDPBalance_Udp         SlbNewCfgEnhVirtServicesTableUDPBalance = 2
	SlbNewCfgEnhVirtServicesTableUDPBalance_Tcp         SlbNewCfgEnhVirtServicesTableUDPBalance = 3
	SlbNewCfgEnhVirtServicesTableUDPBalance_Stateless   SlbNewCfgEnhVirtServicesTableUDPBalance = 4
	SlbNewCfgEnhVirtServicesTableUDPBalance_TcpAndUdp   SlbNewCfgEnhVirtServicesTableUDPBalance = 5
	SlbNewCfgEnhVirtServicesTableUDPBalance_Sctp        SlbNewCfgEnhVirtServicesTableUDPBalance = 6
	SlbNewCfgEnhVirtServicesTableUDPBalance_Unsupported SlbNewCfgEnhVirtServicesTableUDPBalance = 2147483647
)

type SlbNewCfgEnhVirtServicesTableDirServerRtn int32

const (
	SlbNewCfgEnhVirtServicesTableDirServerRtn_Enabled     SlbNewCfgEnhVirtServicesTableDirServerRtn = 1
	SlbNewCfgEnhVirtServicesTableDirServerRtn_Disabled    SlbNewCfgEnhVirtServicesTableDirServerRtn = 2
	SlbNewCfgEnhVirtServicesTableDirServerRtn_Unsupported SlbNewCfgEnhVirtServicesTableDirServerRtn = 2147483647
)

type SlbNewCfgEnhVirtServicesTableRtspUrlParse int32

const (
	SlbNewCfgEnhVirtServicesTableRtspUrlParse_None         SlbNewCfgEnhVirtServicesTableRtspUrlParse = 1
	SlbNewCfgEnhVirtServicesTableRtspUrlParse_L4hash       SlbNewCfgEnhVirtServicesTableRtspUrlParse = 2
	SlbNewCfgEnhVirtServicesTableRtspUrlParse_Hash         SlbNewCfgEnhVirtServicesTableRtspUrlParse = 3
	SlbNewCfgEnhVirtServicesTableRtspUrlParse_PatternMatch SlbNewCfgEnhVirtServicesTableRtspUrlParse = 4
	SlbNewCfgEnhVirtServicesTableRtspUrlParse_Unsupported  SlbNewCfgEnhVirtServicesTableRtspUrlParse = 2147483647
)

type SlbNewCfgEnhVirtServicesTableDBind int32

const (
	SlbNewCfgEnhVirtServicesTableDBind_Enabled     SlbNewCfgEnhVirtServicesTableDBind = 1
	SlbNewCfgEnhVirtServicesTableDBind_Disabled    SlbNewCfgEnhVirtServicesTableDBind = 2
	SlbNewCfgEnhVirtServicesTableDBind_Forceproxy  SlbNewCfgEnhVirtServicesTableDBind = 3
	SlbNewCfgEnhVirtServicesTableDBind_Unsupported SlbNewCfgEnhVirtServicesTableDBind = 2147483647
)

type SlbNewCfgEnhVirtServicesTableFtpParsing int32

const (
	SlbNewCfgEnhVirtServicesTableFtpParsing_Enabled     SlbNewCfgEnhVirtServicesTableFtpParsing = 1
	SlbNewCfgEnhVirtServicesTableFtpParsing_Disabled    SlbNewCfgEnhVirtServicesTableFtpParsing = 2
	SlbNewCfgEnhVirtServicesTableFtpParsing_Unsupported SlbNewCfgEnhVirtServicesTableFtpParsing = 2147483647
)

type SlbNewCfgEnhVirtServicesTableRemapUDPFrags int32

const (
	SlbNewCfgEnhVirtServicesTableRemapUDPFrags_Enabled     SlbNewCfgEnhVirtServicesTableRemapUDPFrags = 1
	SlbNewCfgEnhVirtServicesTableRemapUDPFrags_Disabled    SlbNewCfgEnhVirtServicesTableRemapUDPFrags = 2
	SlbNewCfgEnhVirtServicesTableRemapUDPFrags_Unsupported SlbNewCfgEnhVirtServicesTableRemapUDPFrags = 2147483647
)

type SlbNewCfgEnhVirtServicesTableDnsSlb int32

const (
	SlbNewCfgEnhVirtServicesTableDnsSlb_Enabled     SlbNewCfgEnhVirtServicesTableDnsSlb = 1
	SlbNewCfgEnhVirtServicesTableDnsSlb_Disabled    SlbNewCfgEnhVirtServicesTableDnsSlb = 2
	SlbNewCfgEnhVirtServicesTableDnsSlb_Unsupported SlbNewCfgEnhVirtServicesTableDnsSlb = 2147483647
)

type SlbNewCfgEnhVirtServicesTablePBind int32

const (
	SlbNewCfgEnhVirtServicesTablePBind_Clientip    SlbNewCfgEnhVirtServicesTablePBind = 2
	SlbNewCfgEnhVirtServicesTablePBind_Disabled    SlbNewCfgEnhVirtServicesTablePBind = 3
	SlbNewCfgEnhVirtServicesTablePBind_Sslid       SlbNewCfgEnhVirtServicesTablePBind = 4
	SlbNewCfgEnhVirtServicesTablePBind_Cookie      SlbNewCfgEnhVirtServicesTablePBind = 5
	SlbNewCfgEnhVirtServicesTablePBind_Unsupported SlbNewCfgEnhVirtServicesTablePBind = 2147483647
)

type SlbNewCfgEnhVirtServicesTableUriCookie int32

const (
	SlbNewCfgEnhVirtServicesTableUriCookie_Enabled     SlbNewCfgEnhVirtServicesTableUriCookie = 1
	SlbNewCfgEnhVirtServicesTableUriCookie_Disabled    SlbNewCfgEnhVirtServicesTableUriCookie = 2
	SlbNewCfgEnhVirtServicesTableUriCookie_Unsupported SlbNewCfgEnhVirtServicesTableUriCookie = 2147483647
)

type SlbNewCfgEnhVirtServicesTableCookieMode int32

const (
	SlbNewCfgEnhVirtServicesTableCookieMode_Rewrite     SlbNewCfgEnhVirtServicesTableCookieMode = 1
	SlbNewCfgEnhVirtServicesTableCookieMode_Passive     SlbNewCfgEnhVirtServicesTableCookieMode = 2
	SlbNewCfgEnhVirtServicesTableCookieMode_Insert      SlbNewCfgEnhVirtServicesTableCookieMode = 3
	SlbNewCfgEnhVirtServicesTableCookieMode_Unsupported SlbNewCfgEnhVirtServicesTableCookieMode = 2147483647
)

type SlbNewCfgEnhVirtServicesTableHttpSlb int32

const (
	SlbNewCfgEnhVirtServicesTableHttpSlb_Disabled    SlbNewCfgEnhVirtServicesTableHttpSlb = 1
	SlbNewCfgEnhVirtServicesTableHttpSlb_Urlslb      SlbNewCfgEnhVirtServicesTableHttpSlb = 2
	SlbNewCfgEnhVirtServicesTableHttpSlb_Urlhash     SlbNewCfgEnhVirtServicesTableHttpSlb = 3
	SlbNewCfgEnhVirtServicesTableHttpSlb_Cookie      SlbNewCfgEnhVirtServicesTableHttpSlb = 4
	SlbNewCfgEnhVirtServicesTableHttpSlb_Host        SlbNewCfgEnhVirtServicesTableHttpSlb = 5
	SlbNewCfgEnhVirtServicesTableHttpSlb_Browser     SlbNewCfgEnhVirtServicesTableHttpSlb = 6
	SlbNewCfgEnhVirtServicesTableHttpSlb_Others      SlbNewCfgEnhVirtServicesTableHttpSlb = 7
	SlbNewCfgEnhVirtServicesTableHttpSlb_Headerhash  SlbNewCfgEnhVirtServicesTableHttpSlb = 8
	SlbNewCfgEnhVirtServicesTableHttpSlb_Version     SlbNewCfgEnhVirtServicesTableHttpSlb = 9
	SlbNewCfgEnhVirtServicesTableHttpSlb_Unsupported SlbNewCfgEnhVirtServicesTableHttpSlb = 2147483647
)

type SlbNewCfgEnhVirtServicesTableHttpSlbOption int32

const (
	SlbNewCfgEnhVirtServicesTableHttpSlbOption_And         SlbNewCfgEnhVirtServicesTableHttpSlbOption = 1
	SlbNewCfgEnhVirtServicesTableHttpSlbOption_Or          SlbNewCfgEnhVirtServicesTableHttpSlbOption = 2
	SlbNewCfgEnhVirtServicesTableHttpSlbOption_None        SlbNewCfgEnhVirtServicesTableHttpSlbOption = 3
	SlbNewCfgEnhVirtServicesTableHttpSlbOption_Unsupported SlbNewCfgEnhVirtServicesTableHttpSlbOption = 2147483647
)

type SlbNewCfgEnhVirtServicesTableHttpSlb2 int32

const (
	SlbNewCfgEnhVirtServicesTableHttpSlb2_Disabled    SlbNewCfgEnhVirtServicesTableHttpSlb2 = 1
	SlbNewCfgEnhVirtServicesTableHttpSlb2_Urlslb      SlbNewCfgEnhVirtServicesTableHttpSlb2 = 2
	SlbNewCfgEnhVirtServicesTableHttpSlb2_Urlhash     SlbNewCfgEnhVirtServicesTableHttpSlb2 = 3
	SlbNewCfgEnhVirtServicesTableHttpSlb2_Cookie      SlbNewCfgEnhVirtServicesTableHttpSlb2 = 4
	SlbNewCfgEnhVirtServicesTableHttpSlb2_Host        SlbNewCfgEnhVirtServicesTableHttpSlb2 = 5
	SlbNewCfgEnhVirtServicesTableHttpSlb2_Browser     SlbNewCfgEnhVirtServicesTableHttpSlb2 = 6
	SlbNewCfgEnhVirtServicesTableHttpSlb2_Others      SlbNewCfgEnhVirtServicesTableHttpSlb2 = 7
	SlbNewCfgEnhVirtServicesTableHttpSlb2_Headerhash  SlbNewCfgEnhVirtServicesTableHttpSlb2 = 8
	SlbNewCfgEnhVirtServicesTableHttpSlb2_Version     SlbNewCfgEnhVirtServicesTableHttpSlb2 = 9
	SlbNewCfgEnhVirtServicesTableHttpSlb2_Unsupported SlbNewCfgEnhVirtServicesTableHttpSlb2 = 2147483647
)

type SlbNewCfgEnhVirtServicesTableDelete int32

const (
	SlbNewCfgEnhVirtServicesTableDelete_Other       SlbNewCfgEnhVirtServicesTableDelete = 1
	SlbNewCfgEnhVirtServicesTableDelete_Delete      SlbNewCfgEnhVirtServicesTableDelete = 2
	SlbNewCfgEnhVirtServicesTableDelete_Unsupported SlbNewCfgEnhVirtServicesTableDelete = 2147483647
)

type SlbNewCfgEnhVirtServicesTableApm int32

const (
	SlbNewCfgEnhVirtServicesTableApm_Enabled     SlbNewCfgEnhVirtServicesTableApm = 1
	SlbNewCfgEnhVirtServicesTableApm_Disabled    SlbNewCfgEnhVirtServicesTableApm = 2
	SlbNewCfgEnhVirtServicesTableApm_Unsupported SlbNewCfgEnhVirtServicesTableApm = 2147483647
)

type SlbNewCfgEnhVirtServicesTableNonHTTP int32

const (
	SlbNewCfgEnhVirtServicesTableNonHTTP_Enabled     SlbNewCfgEnhVirtServicesTableNonHTTP = 1
	SlbNewCfgEnhVirtServicesTableNonHTTP_Disabled    SlbNewCfgEnhVirtServicesTableNonHTTP = 2
	SlbNewCfgEnhVirtServicesTableNonHTTP_Unsupported SlbNewCfgEnhVirtServicesTableNonHTTP = 2147483647
)

type SlbNewCfgEnhVirtServicesTableIpRep int32

const (
	SlbNewCfgEnhVirtServicesTableIpRep_Enabled     SlbNewCfgEnhVirtServicesTableIpRep = 1
	SlbNewCfgEnhVirtServicesTableIpRep_Disabled    SlbNewCfgEnhVirtServicesTableIpRep = 2
	SlbNewCfgEnhVirtServicesTableIpRep_Unsupported SlbNewCfgEnhVirtServicesTableIpRep = 2147483647
)

type SlbNewCfgEnhVirtServicesTableCdnProxy int32

const (
	SlbNewCfgEnhVirtServicesTableCdnProxy_Enabled     SlbNewCfgEnhVirtServicesTableCdnProxy = 1
	SlbNewCfgEnhVirtServicesTableCdnProxy_Disabled    SlbNewCfgEnhVirtServicesTableCdnProxy = 2
	SlbNewCfgEnhVirtServicesTableCdnProxy_Unsupported SlbNewCfgEnhVirtServicesTableCdnProxy = 2147483647
)

type SlbNewCfgEnhVirtServicesTableStatus int32

const (
	SlbNewCfgEnhVirtServicesTableStatus_Up          SlbNewCfgEnhVirtServicesTableStatus = 1
	SlbNewCfgEnhVirtServicesTableStatus_Down        SlbNewCfgEnhVirtServicesTableStatus = 2
	SlbNewCfgEnhVirtServicesTableStatus_AdminDown   SlbNewCfgEnhVirtServicesTableStatus = 3
	SlbNewCfgEnhVirtServicesTableStatus_Warning     SlbNewCfgEnhVirtServicesTableStatus = 4
	SlbNewCfgEnhVirtServicesTableStatus_Shutdown    SlbNewCfgEnhVirtServicesTableStatus = 5
	SlbNewCfgEnhVirtServicesTableStatus_Error       SlbNewCfgEnhVirtServicesTableStatus = 6
	SlbNewCfgEnhVirtServicesTableStatus_Unsupported SlbNewCfgEnhVirtServicesTableStatus = 2147483647
)

type SlbNewCfgEnhVirtServicesTableRtSrcTnl int32

const (
	SlbNewCfgEnhVirtServicesTableRtSrcTnl_Enabled     SlbNewCfgEnhVirtServicesTableRtSrcTnl = 1
	SlbNewCfgEnhVirtServicesTableRtSrcTnl_Disabled    SlbNewCfgEnhVirtServicesTableRtSrcTnl = 2
	SlbNewCfgEnhVirtServicesTableRtSrcTnl_Unsupported SlbNewCfgEnhVirtServicesTableRtSrcTnl = 2147483647
)

type SlbNewCfgEnhVirtServicesTableParams struct {
	// The number of the virtual server.
	ServIndex string `json:"ServIndex,omitempty"`
	// The service index. This has no external meaning
	Index int32 `json:"Index,omitempty"`
	// The layer4 virtual port number of the service. it can be either 1 for ip or between 9 to 65534, virt port no. 2 to 9 are invalid
	VirtPort uint64 `json:"VirtPort,omitempty"`
	// The layer4 real port number of the service, it can be either 0 for multiple real ports or 1 for ip service or between 5 to 65534. (2 to 5 are invalid)
	RealPort uint64 `json:"RealPort,omitempty"`
	// Set protocol for the virtual service to
	// UDP or TCP or SCTP or tcpAndUdp or stateless.
	// tcpAndUdp is applicable only to ip service.
	UDPBalance SlbNewCfgEnhVirtServicesTableUDPBalance `json:"UDPBalance,omitempty"`
	// The BWM contract number for this service.
	BwmContract int32 `json:"BwmContract,omitempty"`
	// Enable or disable direct server return feature. To translate only
	// MAC addresses in performing server load balancing when enabled.
	// This allow servers to return directly to client since IP addresses
	// have not been changed.
	DirServerRtn SlbNewCfgEnhVirtServicesTableDirServerRtn `json:"DirServerRtn,omitempty"`
	// Select RTSP URL load balancing type.
	RtspUrlParse SlbNewCfgEnhVirtServicesTableRtspUrlParse `json:"RtspUrlParse,omitempty"`
	// Enable/disable/forceproxy delayed binding.
	DBind SlbNewCfgEnhVirtServicesTableDBind `json:"DBind,omitempty"`
	// Enable or Disable the ftp parsing for the virtual service.
	FtpParsing SlbNewCfgEnhVirtServicesTableFtpParsing `json:"FtpParsing,omitempty"`
	// Enable or disable remapping UDP server fragments
	RemapUDPFrags SlbNewCfgEnhVirtServicesTableRemapUDPFrags `json:"RemapUDPFrags,omitempty"`
	// Enable or disable DNS query load balancing.
	DnsSlb SlbNewCfgEnhVirtServicesTableDnsSlb `json:"DnsSlb,omitempty"`
	// The number of cookie search response count.
	ResponseCount uint32 `json:"ResponseCount,omitempty"`
	// Enable or disable persistent bindings for the virtual port.
	PBind SlbNewCfgEnhVirtServicesTablePBind `json:"PBind,omitempty"`
	// The starting byte offset of the cookie value.
	Coffset uint64 `json:"Coffset,omitempty"`
	// The number of bytes to extract from the cookie value.
	Clength uint64 `json:"Clength,omitempty"`
	// Enable or disable cookie search in URI
	UriCookie SlbNewCfgEnhVirtServicesTableUriCookie `json:"UriCookie,omitempty"`
	// Select cookie persistence mode. Mode disabled(4) not supported on Alteon
	CookieMode SlbNewCfgEnhVirtServicesTableCookieMode `json:"CookieMode,omitempty"`
	// Select HTTP server loadbalancing for the virtual port.
	HttpSlb SlbNewCfgEnhVirtServicesTableHttpSlb `json:"HttpSlb,omitempty"`
	// Select HTTP server loadbalancing for the virtual port.
	HttpSlbOption SlbNewCfgEnhVirtServicesTableHttpSlbOption `json:"HttpSlbOption,omitempty"`
	// Select HTTP server loadbalancing for the virtual port.
	HttpSlb2 SlbNewCfgEnhVirtServicesTableHttpSlb2 `json:"HttpSlb2,omitempty"`
	// When set to the value of 2 (delete), the entire row is deleted.
	// When read, other(1) is returned. Setting the value to anything
	// other than 2(delete) has no effect on the state of the row.
	Delete SlbNewCfgEnhVirtServicesTableDelete `json:"Delete,omitempty"`
	// Enable/disable apm.
	Apm SlbNewCfgEnhVirtServicesTableApm `json:"Apm,omitempty"`
	// Enable/Disable to send non-HTTP traffic.
	NonHTTP SlbNewCfgEnhVirtServicesTableNonHTTP `json:"NonHTTP,omitempty"`
	// Enable/Disable IP reputation.
	IpRep SlbNewCfgEnhVirtServicesTableIpRep `json:"IpRep,omitempty"`
	// Enable/Disable CDN proxy.
	CdnProxy SlbNewCfgEnhVirtServicesTableCdnProxy `json:"CdnProxy,omitempty"`
	// The service status: up, down, admin down, warning, shutdown, error.
	Status SlbNewCfgEnhVirtServicesTableStatus `json:"Status,omitempty"`
	// Enable/Disable Return to Source Tunnel.
	RtSrcTnl SlbNewCfgEnhVirtServicesTableRtSrcTnl `json:"RtSrcTnl,omitempty"`
	// Set Sideband policy.
	Sideband string `json:"Sideband,omitempty"`
}

func (p SlbNewCfgEnhVirtServicesTableParams) iMABean() {}
