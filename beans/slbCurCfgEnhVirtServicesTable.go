package beans

import (
	"fmt"
	"reflect"
)

// SlbCurCfgEnhVirtServicesTable The table of virtual services.
// Note:This mib is not supported for VX instance of Virtualization.
type SlbCurCfgEnhVirtServicesTable struct {
	// The number of the virtual server.
	SlbCurCfgEnhVirtServIndex string
	// The service index. This has no external meaning
	SlbCurCfgEnhVirtServiceIndex int32
	Params                       *SlbCurCfgEnhVirtServicesTableParams
}

func NewSlbCurCfgEnhVirtServicesTableList() *SlbCurCfgEnhVirtServicesTable {
	return &SlbCurCfgEnhVirtServicesTable{}
}

func NewSlbCurCfgEnhVirtServicesTable(
	slbCurCfgEnhVirtServIndex string,
	slbCurCfgEnhVirtServiceIndex int32,
	params *SlbCurCfgEnhVirtServicesTableParams,
) *SlbCurCfgEnhVirtServicesTable {
	return &SlbCurCfgEnhVirtServicesTable{
		SlbCurCfgEnhVirtServIndex:    slbCurCfgEnhVirtServIndex,
		SlbCurCfgEnhVirtServiceIndex: slbCurCfgEnhVirtServiceIndex,
		Params:                       params,
	}
}

func (c *SlbCurCfgEnhVirtServicesTable) Name() string {
	return "SlbCurCfgEnhVirtServicesTable"
}

func (c *SlbCurCfgEnhVirtServicesTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbCurCfgEnhVirtServicesTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbCurCfgEnhVirtServicesTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbCurCfgEnhVirtServIndex).IsZero() &&
		reflect.ValueOf(c.SlbCurCfgEnhVirtServiceIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbCurCfgEnhVirtServIndex) + "/" + fmt.Sprint(c.SlbCurCfgEnhVirtServiceIndex)
}

type SlbCurCfgEnhVirtServicesTableUDPBalance int32

const (
	SlbCurCfgEnhVirtServicesTableUDPBalance_Udp         SlbCurCfgEnhVirtServicesTableUDPBalance = 2
	SlbCurCfgEnhVirtServicesTableUDPBalance_Tcp         SlbCurCfgEnhVirtServicesTableUDPBalance = 3
	SlbCurCfgEnhVirtServicesTableUDPBalance_Stateless   SlbCurCfgEnhVirtServicesTableUDPBalance = 4
	SlbCurCfgEnhVirtServicesTableUDPBalance_TcpAndUdp   SlbCurCfgEnhVirtServicesTableUDPBalance = 5
	SlbCurCfgEnhVirtServicesTableUDPBalance_Sctp        SlbCurCfgEnhVirtServicesTableUDPBalance = 6
	SlbCurCfgEnhVirtServicesTableUDPBalance_Unsupported SlbCurCfgEnhVirtServicesTableUDPBalance = 2147483647
)

type SlbCurCfgEnhVirtServicesTableDirServerRtn int32

const (
	SlbCurCfgEnhVirtServicesTableDirServerRtn_Enabled     SlbCurCfgEnhVirtServicesTableDirServerRtn = 1
	SlbCurCfgEnhVirtServicesTableDirServerRtn_Disabled    SlbCurCfgEnhVirtServicesTableDirServerRtn = 2
	SlbCurCfgEnhVirtServicesTableDirServerRtn_Unsupported SlbCurCfgEnhVirtServicesTableDirServerRtn = 2147483647
)

type SlbCurCfgEnhVirtServicesTableRtspUrlParse int32

const (
	SlbCurCfgEnhVirtServicesTableRtspUrlParse_None         SlbCurCfgEnhVirtServicesTableRtspUrlParse = 1
	SlbCurCfgEnhVirtServicesTableRtspUrlParse_L4hash       SlbCurCfgEnhVirtServicesTableRtspUrlParse = 2
	SlbCurCfgEnhVirtServicesTableRtspUrlParse_Hash         SlbCurCfgEnhVirtServicesTableRtspUrlParse = 3
	SlbCurCfgEnhVirtServicesTableRtspUrlParse_PatternMatch SlbCurCfgEnhVirtServicesTableRtspUrlParse = 4
	SlbCurCfgEnhVirtServicesTableRtspUrlParse_Unsupported  SlbCurCfgEnhVirtServicesTableRtspUrlParse = 2147483647
)

type SlbCurCfgEnhVirtServicesTableFtpParsing int32

const (
	SlbCurCfgEnhVirtServicesTableFtpParsing_Enabled     SlbCurCfgEnhVirtServicesTableFtpParsing = 1
	SlbCurCfgEnhVirtServicesTableFtpParsing_Disabled    SlbCurCfgEnhVirtServicesTableFtpParsing = 2
	SlbCurCfgEnhVirtServicesTableFtpParsing_Unsupported SlbCurCfgEnhVirtServicesTableFtpParsing = 2147483647
)

type SlbCurCfgEnhVirtServicesTableRemapUDPFrags int32

const (
	SlbCurCfgEnhVirtServicesTableRemapUDPFrags_Enabled     SlbCurCfgEnhVirtServicesTableRemapUDPFrags = 1
	SlbCurCfgEnhVirtServicesTableRemapUDPFrags_Disabled    SlbCurCfgEnhVirtServicesTableRemapUDPFrags = 2
	SlbCurCfgEnhVirtServicesTableRemapUDPFrags_Unsupported SlbCurCfgEnhVirtServicesTableRemapUDPFrags = 2147483647
)

type SlbCurCfgEnhVirtServicesTableDnsSlb int32

const (
	SlbCurCfgEnhVirtServicesTableDnsSlb_Enabled     SlbCurCfgEnhVirtServicesTableDnsSlb = 1
	SlbCurCfgEnhVirtServicesTableDnsSlb_Disabled    SlbCurCfgEnhVirtServicesTableDnsSlb = 2
	SlbCurCfgEnhVirtServicesTableDnsSlb_Unsupported SlbCurCfgEnhVirtServicesTableDnsSlb = 2147483647
)

type SlbCurCfgEnhVirtServicesTablePBind int32

const (
	SlbCurCfgEnhVirtServicesTablePBind_Clientip    SlbCurCfgEnhVirtServicesTablePBind = 2
	SlbCurCfgEnhVirtServicesTablePBind_Disabled    SlbCurCfgEnhVirtServicesTablePBind = 3
	SlbCurCfgEnhVirtServicesTablePBind_Sslid       SlbCurCfgEnhVirtServicesTablePBind = 4
	SlbCurCfgEnhVirtServicesTablePBind_Cookie      SlbCurCfgEnhVirtServicesTablePBind = 5
	SlbCurCfgEnhVirtServicesTablePBind_Unsupported SlbCurCfgEnhVirtServicesTablePBind = 2147483647
)

type SlbCurCfgEnhVirtServicesTableUriCookie int32

const (
	SlbCurCfgEnhVirtServicesTableUriCookie_Enabled     SlbCurCfgEnhVirtServicesTableUriCookie = 1
	SlbCurCfgEnhVirtServicesTableUriCookie_Disabled    SlbCurCfgEnhVirtServicesTableUriCookie = 2
	SlbCurCfgEnhVirtServicesTableUriCookie_Unsupported SlbCurCfgEnhVirtServicesTableUriCookie = 2147483647
)

type SlbCurCfgEnhVirtServicesTableCookieMode int32

const (
	SlbCurCfgEnhVirtServicesTableCookieMode_Rewrite     SlbCurCfgEnhVirtServicesTableCookieMode = 1
	SlbCurCfgEnhVirtServicesTableCookieMode_Passive     SlbCurCfgEnhVirtServicesTableCookieMode = 2
	SlbCurCfgEnhVirtServicesTableCookieMode_Insert      SlbCurCfgEnhVirtServicesTableCookieMode = 3
	SlbCurCfgEnhVirtServicesTableCookieMode_Disabled    SlbCurCfgEnhVirtServicesTableCookieMode = 4
	SlbCurCfgEnhVirtServicesTableCookieMode_Unsupported SlbCurCfgEnhVirtServicesTableCookieMode = 2147483647
)

type SlbCurCfgEnhVirtServicesTableDBind int32

const (
	SlbCurCfgEnhVirtServicesTableDBind_Enabled     SlbCurCfgEnhVirtServicesTableDBind = 1
	SlbCurCfgEnhVirtServicesTableDBind_Disabled    SlbCurCfgEnhVirtServicesTableDBind = 2
	SlbCurCfgEnhVirtServicesTableDBind_Forceproxy  SlbCurCfgEnhVirtServicesTableDBind = 3
	SlbCurCfgEnhVirtServicesTableDBind_Unsupported SlbCurCfgEnhVirtServicesTableDBind = 2147483647
)

type SlbCurCfgEnhVirtServicesTableHttpSlb int32

const (
	SlbCurCfgEnhVirtServicesTableHttpSlb_Disabled    SlbCurCfgEnhVirtServicesTableHttpSlb = 1
	SlbCurCfgEnhVirtServicesTableHttpSlb_Urlslb      SlbCurCfgEnhVirtServicesTableHttpSlb = 2
	SlbCurCfgEnhVirtServicesTableHttpSlb_Urlhash     SlbCurCfgEnhVirtServicesTableHttpSlb = 3
	SlbCurCfgEnhVirtServicesTableHttpSlb_Cookie      SlbCurCfgEnhVirtServicesTableHttpSlb = 4
	SlbCurCfgEnhVirtServicesTableHttpSlb_Host        SlbCurCfgEnhVirtServicesTableHttpSlb = 5
	SlbCurCfgEnhVirtServicesTableHttpSlb_Browser     SlbCurCfgEnhVirtServicesTableHttpSlb = 6
	SlbCurCfgEnhVirtServicesTableHttpSlb_Others      SlbCurCfgEnhVirtServicesTableHttpSlb = 7
	SlbCurCfgEnhVirtServicesTableHttpSlb_Headerhash  SlbCurCfgEnhVirtServicesTableHttpSlb = 8
	SlbCurCfgEnhVirtServicesTableHttpSlb_Version     SlbCurCfgEnhVirtServicesTableHttpSlb = 9
	SlbCurCfgEnhVirtServicesTableHttpSlb_Unsupported SlbCurCfgEnhVirtServicesTableHttpSlb = 2147483647
)

type SlbCurCfgEnhVirtServicesTableHttpSlbOption int32

const (
	SlbCurCfgEnhVirtServicesTableHttpSlbOption_And         SlbCurCfgEnhVirtServicesTableHttpSlbOption = 1
	SlbCurCfgEnhVirtServicesTableHttpSlbOption_Or          SlbCurCfgEnhVirtServicesTableHttpSlbOption = 2
	SlbCurCfgEnhVirtServicesTableHttpSlbOption_None        SlbCurCfgEnhVirtServicesTableHttpSlbOption = 3
	SlbCurCfgEnhVirtServicesTableHttpSlbOption_Unsupported SlbCurCfgEnhVirtServicesTableHttpSlbOption = 2147483647
)

type SlbCurCfgEnhVirtServicesTableHttpSlb2 int32

const (
	SlbCurCfgEnhVirtServicesTableHttpSlb2_Disabled    SlbCurCfgEnhVirtServicesTableHttpSlb2 = 1
	SlbCurCfgEnhVirtServicesTableHttpSlb2_Urlslb      SlbCurCfgEnhVirtServicesTableHttpSlb2 = 2
	SlbCurCfgEnhVirtServicesTableHttpSlb2_Urlhash     SlbCurCfgEnhVirtServicesTableHttpSlb2 = 3
	SlbCurCfgEnhVirtServicesTableHttpSlb2_Cookie      SlbCurCfgEnhVirtServicesTableHttpSlb2 = 4
	SlbCurCfgEnhVirtServicesTableHttpSlb2_Host        SlbCurCfgEnhVirtServicesTableHttpSlb2 = 5
	SlbCurCfgEnhVirtServicesTableHttpSlb2_Browser     SlbCurCfgEnhVirtServicesTableHttpSlb2 = 6
	SlbCurCfgEnhVirtServicesTableHttpSlb2_Others      SlbCurCfgEnhVirtServicesTableHttpSlb2 = 7
	SlbCurCfgEnhVirtServicesTableHttpSlb2_Headerhash  SlbCurCfgEnhVirtServicesTableHttpSlb2 = 8
	SlbCurCfgEnhVirtServicesTableHttpSlb2_Version     SlbCurCfgEnhVirtServicesTableHttpSlb2 = 9
	SlbCurCfgEnhVirtServicesTableHttpSlb2_Unsupported SlbCurCfgEnhVirtServicesTableHttpSlb2 = 2147483647
)

type SlbCurCfgEnhVirtServicesTableApm int32

const (
	SlbCurCfgEnhVirtServicesTableApm_Enabled     SlbCurCfgEnhVirtServicesTableApm = 1
	SlbCurCfgEnhVirtServicesTableApm_Disabled    SlbCurCfgEnhVirtServicesTableApm = 2
	SlbCurCfgEnhVirtServicesTableApm_Unsupported SlbCurCfgEnhVirtServicesTableApm = 2147483647
)

type SlbCurCfgEnhVirtServicesTableNonHTTP int32

const (
	SlbCurCfgEnhVirtServicesTableNonHTTP_Enabled     SlbCurCfgEnhVirtServicesTableNonHTTP = 1
	SlbCurCfgEnhVirtServicesTableNonHTTP_Disabled    SlbCurCfgEnhVirtServicesTableNonHTTP = 2
	SlbCurCfgEnhVirtServicesTableNonHTTP_Unsupported SlbCurCfgEnhVirtServicesTableNonHTTP = 2147483647
)

type SlbCurCfgEnhVirtServicesTableIpRep int32

const (
	SlbCurCfgEnhVirtServicesTableIpRep_Enabled     SlbCurCfgEnhVirtServicesTableIpRep = 1
	SlbCurCfgEnhVirtServicesTableIpRep_Disabled    SlbCurCfgEnhVirtServicesTableIpRep = 2
	SlbCurCfgEnhVirtServicesTableIpRep_Unsupported SlbCurCfgEnhVirtServicesTableIpRep = 2147483647
)

type SlbCurCfgEnhVirtServicesTableCdnProxy int32

const (
	SlbCurCfgEnhVirtServicesTableCdnProxy_Enabled     SlbCurCfgEnhVirtServicesTableCdnProxy = 1
	SlbCurCfgEnhVirtServicesTableCdnProxy_Disabled    SlbCurCfgEnhVirtServicesTableCdnProxy = 2
	SlbCurCfgEnhVirtServicesTableCdnProxy_Unsupported SlbCurCfgEnhVirtServicesTableCdnProxy = 2147483647
)

type SlbCurCfgEnhVirtServicesTableStatus int32

const (
	SlbCurCfgEnhVirtServicesTableStatus_Up          SlbCurCfgEnhVirtServicesTableStatus = 1
	SlbCurCfgEnhVirtServicesTableStatus_Down        SlbCurCfgEnhVirtServicesTableStatus = 2
	SlbCurCfgEnhVirtServicesTableStatus_AdminDown   SlbCurCfgEnhVirtServicesTableStatus = 3
	SlbCurCfgEnhVirtServicesTableStatus_Warning     SlbCurCfgEnhVirtServicesTableStatus = 4
	SlbCurCfgEnhVirtServicesTableStatus_Shutdown    SlbCurCfgEnhVirtServicesTableStatus = 5
	SlbCurCfgEnhVirtServicesTableStatus_Error       SlbCurCfgEnhVirtServicesTableStatus = 6
	SlbCurCfgEnhVirtServicesTableStatus_Unsupported SlbCurCfgEnhVirtServicesTableStatus = 2147483647
)

type SlbCurCfgEnhVirtServicesTableRtSrcTnl int32

const (
	SlbCurCfgEnhVirtServicesTableRtSrcTnl_Enabled     SlbCurCfgEnhVirtServicesTableRtSrcTnl = 1
	SlbCurCfgEnhVirtServicesTableRtSrcTnl_Disabled    SlbCurCfgEnhVirtServicesTableRtSrcTnl = 2
	SlbCurCfgEnhVirtServicesTableRtSrcTnl_Unsupported SlbCurCfgEnhVirtServicesTableRtSrcTnl = 2147483647
)

type SlbCurCfgEnhVirtServicesTableParams struct {
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
	UDPBalance SlbCurCfgEnhVirtServicesTableUDPBalance `json:"UDPBalance,omitempty"`
	// The BWM contract number for this service.
	BwmContract int32 `json:"BwmContract,omitempty"`
	// Enable or disable direct server return feature. To translate only
	// MAC addresses in performing server load balancing when enabled.
	// This allow servers to return directly to client since IP addresses
	// have not been changed.
	DirServerRtn SlbCurCfgEnhVirtServicesTableDirServerRtn `json:"DirServerRtn,omitempty"`
	// Select RTSP URL load balancing type.
	RtspUrlParse SlbCurCfgEnhVirtServicesTableRtspUrlParse `json:"RtspUrlParse,omitempty"`
	// Enable or Disable the ftp parsing for the virtual service.
	FtpParsing SlbCurCfgEnhVirtServicesTableFtpParsing `json:"FtpParsing,omitempty"`
	// Enable or disable remapping UDP server fragments
	RemapUDPFrags SlbCurCfgEnhVirtServicesTableRemapUDPFrags `json:"RemapUDPFrags,omitempty"`
	// Enable or disable DNS query load balancing.
	DnsSlb SlbCurCfgEnhVirtServicesTableDnsSlb `json:"DnsSlb,omitempty"`
	// The number of cookie search response count.
	ResponseCount uint32 `json:"ResponseCount,omitempty"`
	// Enable or disable persistent bindings for the virtual port.
	PBind SlbCurCfgEnhVirtServicesTablePBind `json:"PBind,omitempty"`
	// The starting byte offset of the cookie value.
	Coffset uint64 `json:"Coffset,omitempty"`
	// The number of bytes to extract from the cookie value.
	Clength uint64 `json:"Clength,omitempty"`
	// Enable or disable cookie search in URI
	UriCookie SlbCurCfgEnhVirtServicesTableUriCookie `json:"UriCookie,omitempty"`
	// Select cookie persistence mode.
	CookieMode SlbCurCfgEnhVirtServicesTableCookieMode `json:"CookieMode,omitempty"`
	// Enable/disable/forceproxy delayed binding.
	DBind SlbCurCfgEnhVirtServicesTableDBind `json:"DBind,omitempty"`
	// Select HTTP server loadbalancing for the virtual port.
	HttpSlb SlbCurCfgEnhVirtServicesTableHttpSlb `json:"HttpSlb,omitempty"`
	// Select HTTP server loadbalancing for the virtual port.
	HttpSlbOption SlbCurCfgEnhVirtServicesTableHttpSlbOption `json:"HttpSlbOption,omitempty"`
	// Select HTTP server loadbalancing for the virtual port.
	HttpSlb2 SlbCurCfgEnhVirtServicesTableHttpSlb2 `json:"HttpSlb2,omitempty"`
	// Enable/disable apm.
	Apm SlbCurCfgEnhVirtServicesTableApm `json:"Apm,omitempty"`
	// Enable or disable Send non-HTTP traffic.
	NonHTTP SlbCurCfgEnhVirtServicesTableNonHTTP `json:"NonHTTP,omitempty"`
	// The domain name of the virtual server and remote real server.
	Dname string `json:"Dname,omitempty"`
	// Enable or disable IP reputation.
	IpRep SlbCurCfgEnhVirtServicesTableIpRep `json:"IpRep,omitempty"`
	// Enable or disable CDN proxy.
	CdnProxy SlbCurCfgEnhVirtServicesTableCdnProxy `json:"CdnProxy,omitempty"`
	// Total number of sessions that are currently handled by the virtual service.
	CurrSessions uint32 `json:"CurrSessions,omitempty"`
	// Total number of sessions that are handled by the virtual service.
	TotalSessions uint32 `json:"TotalSessions,omitempty"`
	// Highest number of sessions that have been handled by the virtual service.
	HighestSessions uint32 `json:"HighestSessions,omitempty"`
	// The lower 32 bit value of total octets received and transmitted out of the virtual service.
	HCOctetsLow32 uint32 `json:"HCOctetsLow32,omitempty"`
	// The higher 32 bit value of total octets received and transmitted out of the virtual service.
	HCOctetsHigh32 uint32 `json:"HCOctetsHigh32,omitempty"`
	// Total octets received and transmitted by the virtual service.
	HCOctets uint64 `json:"HCOctets,omitempty"`
	// Number of sessions that are currently handled per second by the virtual service .
	SessPerSec uint32 `json:"SessPerSec,omitempty"`
	// Number of octets received and transmitted per second by the virtual service.
	OctetsPerSec string `json:"OctetsPerSec,omitempty"`
	// The average ClientRTT in ms during the past collection interval.
	ClientRtt string `json:"ClientRtt,omitempty"`
	// The average ServerRTT in ms during the past collection interval.
	ServerRtt string `json:"ServerRtt,omitempty"`
	// The average application response from AX in ms during the past collection interval.
	AppResponse string `json:"AppResponse,omitempty"`
	// The average response transfer from AX in ms during the past collection interval.
	RespTransfer string `json:"RespTransfer,omitempty"`
	// The sum of clientRtt, serverRtt, application response, response transfer.
	EndToEndTotal string `json:"EndToEndTotal,omitempty"`
	// The service status: up, down, admin down, warning, shutdown, error.
	Status SlbCurCfgEnhVirtServicesTableStatus `json:"Status,omitempty"`
	// Enable or disable Return to Source Tunnel.
	RtSrcTnl SlbCurCfgEnhVirtServicesTableRtSrcTnl `json:"RtSrcTnl,omitempty"`
	// The number of the virtual server.
	Sideband string `json:"Sideband,omitempty"`
	// The sideband processing time in ms during the past collection interval.
	SidebandTime string `json:"SidebandTime,omitempty"`
	// Number of packets that are currently handled per second by the virtual service .
	PktPerSec uint32 `json:"PktPerSec,omitempty"`
	// The average request transfer to AX in ms during the past collection interval.
	ReqTransfer string `json:"ReqTransfer,omitempty"`
	// The average server response time in ms during the past collection interval.
	ServerResponse string `json:"ServerResponse,omitempty"`
}

func (p SlbCurCfgEnhVirtServicesTableParams) iMABean() {}
