package beans

import (
	"fmt"
	"reflect"
)

// SlbCurCfgEnhVirtServicesFifthPartTable The table of virtual Services.
// Note:This mib is not supported for VX instance of Virtualization.
type SlbCurCfgEnhVirtServicesFifthPartTable struct {
	// The number of the virtual server.
	SlbCurCfgEnhVirtServFifthPartIndex string
	// The service index. This has no external meaning
	SlbCurCfgEnhVirtServiceFifthPartIndex int32
	Params                                *SlbCurCfgEnhVirtServicesFifthPartTableParams
}

func NewSlbCurCfgEnhVirtServicesFifthPartTableList() *SlbCurCfgEnhVirtServicesFifthPartTable {
	return &SlbCurCfgEnhVirtServicesFifthPartTable{}
}

func NewSlbCurCfgEnhVirtServicesFifthPartTable(
	slbCurCfgEnhVirtServFifthPartIndex string,
	slbCurCfgEnhVirtServiceFifthPartIndex int32,
	params *SlbCurCfgEnhVirtServicesFifthPartTableParams,
) *SlbCurCfgEnhVirtServicesFifthPartTable {
	return &SlbCurCfgEnhVirtServicesFifthPartTable{
		SlbCurCfgEnhVirtServFifthPartIndex:    slbCurCfgEnhVirtServFifthPartIndex,
		SlbCurCfgEnhVirtServiceFifthPartIndex: slbCurCfgEnhVirtServiceFifthPartIndex,
		Params:                                params,
	}
}

func (c *SlbCurCfgEnhVirtServicesFifthPartTable) Name() string {
	return "SlbCurCfgEnhVirtServicesFifthPartTable"
}

func (c *SlbCurCfgEnhVirtServicesFifthPartTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbCurCfgEnhVirtServicesFifthPartTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbCurCfgEnhVirtServicesFifthPartTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbCurCfgEnhVirtServFifthPartIndex).IsZero() &&
		reflect.ValueOf(c.SlbCurCfgEnhVirtServiceFifthPartIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbCurCfgEnhVirtServFifthPartIndex) + "/" + fmt.Sprint(c.SlbCurCfgEnhVirtServiceFifthPartIndex)
}

type SlbCurCfgEnhVirtServicesFifthPartTableServApplicationType int32

const (
	SlbCurCfgEnhVirtServicesFifthPartTableServApplicationType_BasicSlb    SlbCurCfgEnhVirtServicesFifthPartTableServApplicationType = 1
	SlbCurCfgEnhVirtServicesFifthPartTableServApplicationType_Dns         SlbCurCfgEnhVirtServicesFifthPartTableServApplicationType = 2
	SlbCurCfgEnhVirtServicesFifthPartTableServApplicationType_Ftp         SlbCurCfgEnhVirtServicesFifthPartTableServApplicationType = 3
	SlbCurCfgEnhVirtServicesFifthPartTableServApplicationType_FtpData     SlbCurCfgEnhVirtServicesFifthPartTableServApplicationType = 4
	SlbCurCfgEnhVirtServicesFifthPartTableServApplicationType_Ldap        SlbCurCfgEnhVirtServicesFifthPartTableServApplicationType = 5
	SlbCurCfgEnhVirtServicesFifthPartTableServApplicationType_Http        SlbCurCfgEnhVirtServicesFifthPartTableServApplicationType = 6
	SlbCurCfgEnhVirtServicesFifthPartTableServApplicationType_Https       SlbCurCfgEnhVirtServicesFifthPartTableServApplicationType = 7
	SlbCurCfgEnhVirtServicesFifthPartTableServApplicationType_Ssl         SlbCurCfgEnhVirtServicesFifthPartTableServApplicationType = 8
	SlbCurCfgEnhVirtServicesFifthPartTableServApplicationType_Rtsp        SlbCurCfgEnhVirtServicesFifthPartTableServApplicationType = 9
	SlbCurCfgEnhVirtServicesFifthPartTableServApplicationType_Sip         SlbCurCfgEnhVirtServicesFifthPartTableServApplicationType = 10
	SlbCurCfgEnhVirtServicesFifthPartTableServApplicationType_Wts         SlbCurCfgEnhVirtServicesFifthPartTableServApplicationType = 11
	SlbCurCfgEnhVirtServicesFifthPartTableServApplicationType_Tftp        SlbCurCfgEnhVirtServicesFifthPartTableServApplicationType = 12
	SlbCurCfgEnhVirtServicesFifthPartTableServApplicationType_Smtp        SlbCurCfgEnhVirtServicesFifthPartTableServApplicationType = 13
	SlbCurCfgEnhVirtServicesFifthPartTableServApplicationType_Pop3        SlbCurCfgEnhVirtServicesFifthPartTableServApplicationType = 14
	SlbCurCfgEnhVirtServicesFifthPartTableServApplicationType_Ip          SlbCurCfgEnhVirtServicesFifthPartTableServApplicationType = 15
	SlbCurCfgEnhVirtServicesFifthPartTableServApplicationType_MhSctp      SlbCurCfgEnhVirtServicesFifthPartTableServApplicationType = 16
	SlbCurCfgEnhVirtServicesFifthPartTableServApplicationType_Unsupported SlbCurCfgEnhVirtServicesFifthPartTableServApplicationType = 2147483647
)

type SlbCurCfgEnhVirtServicesFifthPartTableAction int32

const (
	SlbCurCfgEnhVirtServicesFifthPartTableAction_Group       SlbCurCfgEnhVirtServicesFifthPartTableAction = 1
	SlbCurCfgEnhVirtServicesFifthPartTableAction_Redirect    SlbCurCfgEnhVirtServicesFifthPartTableAction = 2
	SlbCurCfgEnhVirtServicesFifthPartTableAction_Discard     SlbCurCfgEnhVirtServicesFifthPartTableAction = 3
	SlbCurCfgEnhVirtServicesFifthPartTableAction_Unsupported SlbCurCfgEnhVirtServicesFifthPartTableAction = 2147483647
)

type SlbCurCfgEnhVirtServicesFifthPartTableServCertGrpMark int32

const (
	SlbCurCfgEnhVirtServicesFifthPartTableServCertGrpMark_Cert        SlbCurCfgEnhVirtServicesFifthPartTableServCertGrpMark = 1
	SlbCurCfgEnhVirtServicesFifthPartTableServCertGrpMark_Group       SlbCurCfgEnhVirtServicesFifthPartTableServCertGrpMark = 2
	SlbCurCfgEnhVirtServicesFifthPartTableServCertGrpMark_Unsupported SlbCurCfgEnhVirtServicesFifthPartTableServCertGrpMark = 2147483647
)

type SlbCurCfgEnhVirtServicesFifthPartTableDnsType int32

const (
	SlbCurCfgEnhVirtServicesFifthPartTableDnsType_Dns         SlbCurCfgEnhVirtServicesFifthPartTableDnsType = 1
	SlbCurCfgEnhVirtServicesFifthPartTableDnsType_Dnssec      SlbCurCfgEnhVirtServicesFifthPartTableDnsType = 2
	SlbCurCfgEnhVirtServicesFifthPartTableDnsType_Both        SlbCurCfgEnhVirtServicesFifthPartTableDnsType = 3
	SlbCurCfgEnhVirtServicesFifthPartTableDnsType_Unsupported SlbCurCfgEnhVirtServicesFifthPartTableDnsType = 2147483647
)

type SlbCurCfgEnhVirtServicesFifthPartTableClntproxType int32

const (
	SlbCurCfgEnhVirtServicesFifthPartTableClntproxType_None        SlbCurCfgEnhVirtServicesFifthPartTableClntproxType = 1
	SlbCurCfgEnhVirtServicesFifthPartTableClntproxType_Http        SlbCurCfgEnhVirtServicesFifthPartTableClntproxType = 2
	SlbCurCfgEnhVirtServicesFifthPartTableClntproxType_Https       SlbCurCfgEnhVirtServicesFifthPartTableClntproxType = 3
	SlbCurCfgEnhVirtServicesFifthPartTableClntproxType_Unsupported SlbCurCfgEnhVirtServicesFifthPartTableClntproxType = 2147483647
)

type SlbCurCfgEnhVirtServicesFifthPartTableZerowinSize int32

const (
	SlbCurCfgEnhVirtServicesFifthPartTableZerowinSize_Enabled     SlbCurCfgEnhVirtServicesFifthPartTableZerowinSize = 1
	SlbCurCfgEnhVirtServicesFifthPartTableZerowinSize_Disabled    SlbCurCfgEnhVirtServicesFifthPartTableZerowinSize = 2
	SlbCurCfgEnhVirtServicesFifthPartTableZerowinSize_Unsupported SlbCurCfgEnhVirtServicesFifthPartTableZerowinSize = 2147483647
)

type SlbCurCfgEnhVirtServicesFifthPartTableCookieSecure int32

const (
	SlbCurCfgEnhVirtServicesFifthPartTableCookieSecure_No          SlbCurCfgEnhVirtServicesFifthPartTableCookieSecure = 1
	SlbCurCfgEnhVirtServicesFifthPartTableCookieSecure_Yes         SlbCurCfgEnhVirtServicesFifthPartTableCookieSecure = 2
	SlbCurCfgEnhVirtServicesFifthPartTableCookieSecure_Unsupported SlbCurCfgEnhVirtServicesFifthPartTableCookieSecure = 2147483647
)

type SlbCurCfgEnhVirtServicesFifthPartTableNoRtsp int32

const (
	SlbCurCfgEnhVirtServicesFifthPartTableNoRtsp_Enabled     SlbCurCfgEnhVirtServicesFifthPartTableNoRtsp = 1
	SlbCurCfgEnhVirtServicesFifthPartTableNoRtsp_Disabled    SlbCurCfgEnhVirtServicesFifthPartTableNoRtsp = 2
	SlbCurCfgEnhVirtServicesFifthPartTableNoRtsp_Unsupported SlbCurCfgEnhVirtServicesFifthPartTableNoRtsp = 2147483647
)

type SlbCurCfgEnhVirtServicesFifthPartTableCkRebind int32

const (
	SlbCurCfgEnhVirtServicesFifthPartTableCkRebind_Enabled     SlbCurCfgEnhVirtServicesFifthPartTableCkRebind = 1
	SlbCurCfgEnhVirtServicesFifthPartTableCkRebind_Disabled    SlbCurCfgEnhVirtServicesFifthPartTableCkRebind = 2
	SlbCurCfgEnhVirtServicesFifthPartTableCkRebind_Unsupported SlbCurCfgEnhVirtServicesFifthPartTableCkRebind = 2147483647
)

type SlbCurCfgEnhVirtServicesFifthPartTableParseLimit int32

const (
	SlbCurCfgEnhVirtServicesFifthPartTableParseLimit_Enabled     SlbCurCfgEnhVirtServicesFifthPartTableParseLimit = 1
	SlbCurCfgEnhVirtServicesFifthPartTableParseLimit_Disabled    SlbCurCfgEnhVirtServicesFifthPartTableParseLimit = 2
	SlbCurCfgEnhVirtServicesFifthPartTableParseLimit_Unsupported SlbCurCfgEnhVirtServicesFifthPartTableParseLimit = 2147483647
)

type SlbCurCfgEnhVirtServicesFifthPartTableUriNorm int32

const (
	SlbCurCfgEnhVirtServicesFifthPartTableUriNorm_Enabled     SlbCurCfgEnhVirtServicesFifthPartTableUriNorm = 1
	SlbCurCfgEnhVirtServicesFifthPartTableUriNorm_Disabled    SlbCurCfgEnhVirtServicesFifthPartTableUriNorm = 2
	SlbCurCfgEnhVirtServicesFifthPartTableUriNorm_Unsupported SlbCurCfgEnhVirtServicesFifthPartTableUriNorm = 2147483647
)

type SlbCurCfgEnhVirtServicesFifthPartTableGranularity int32

const (
	SlbCurCfgEnhVirtServicesFifthPartTableGranularity_Service     SlbCurCfgEnhVirtServicesFifthPartTableGranularity = 0
	SlbCurCfgEnhVirtServicesFifthPartTableGranularity_Real        SlbCurCfgEnhVirtServicesFifthPartTableGranularity = 2
	SlbCurCfgEnhVirtServicesFifthPartTableGranularity_Unsupported SlbCurCfgEnhVirtServicesFifthPartTableGranularity = 2147483647
)

type SlbCurCfgEnhVirtServicesFifthPartTableSessLog int32

const (
	SlbCurCfgEnhVirtServicesFifthPartTableSessLog_Enabled     SlbCurCfgEnhVirtServicesFifthPartTableSessLog = 1
	SlbCurCfgEnhVirtServicesFifthPartTableSessLog_Disabled    SlbCurCfgEnhVirtServicesFifthPartTableSessLog = 2
	SlbCurCfgEnhVirtServicesFifthPartTableSessLog_Unsupported SlbCurCfgEnhVirtServicesFifthPartTableSessLog = 2147483647
)

type SlbCurCfgEnhVirtServicesFifthPartTableUdpAge int32

const (
	SlbCurCfgEnhVirtServicesFifthPartTableUdpAge_Enabled     SlbCurCfgEnhVirtServicesFifthPartTableUdpAge = 1
	SlbCurCfgEnhVirtServicesFifthPartTableUdpAge_Disabled    SlbCurCfgEnhVirtServicesFifthPartTableUdpAge = 2
	SlbCurCfgEnhVirtServicesFifthPartTableUdpAge_Unsupported SlbCurCfgEnhVirtServicesFifthPartTableUdpAge = 2147483647
)

type SlbCurCfgEnhVirtServicesFifthPartTableSessEntryMode int32

const (
	SlbCurCfgEnhVirtServicesFifthPartTableSessEntryMode_Single      SlbCurCfgEnhVirtServicesFifthPartTableSessEntryMode = 1
	SlbCurCfgEnhVirtServicesFifthPartTableSessEntryMode_Multiple    SlbCurCfgEnhVirtServicesFifthPartTableSessEntryMode = 2
	SlbCurCfgEnhVirtServicesFifthPartTableSessEntryMode_Unsupported SlbCurCfgEnhVirtServicesFifthPartTableSessEntryMode = 2147483647
)

type SlbCurCfgEnhVirtServicesFifthPartTableAlwaysOn int32

const (
	SlbCurCfgEnhVirtServicesFifthPartTableAlwaysOn_Enabled     SlbCurCfgEnhVirtServicesFifthPartTableAlwaysOn = 1
	SlbCurCfgEnhVirtServicesFifthPartTableAlwaysOn_Disabled    SlbCurCfgEnhVirtServicesFifthPartTableAlwaysOn = 2
	SlbCurCfgEnhVirtServicesFifthPartTableAlwaysOn_Unsupported SlbCurCfgEnhVirtServicesFifthPartTableAlwaysOn = 2147483647
)

type SlbCurCfgEnhVirtServicesFifthPartTableSendRST int32

const (
	SlbCurCfgEnhVirtServicesFifthPartTableSendRST_Reset       SlbCurCfgEnhVirtServicesFifthPartTableSendRST = 1
	SlbCurCfgEnhVirtServicesFifthPartTableSendRST_Drop        SlbCurCfgEnhVirtServicesFifthPartTableSendRST = 2
	SlbCurCfgEnhVirtServicesFifthPartTableSendRST_UdpIcmpErr  SlbCurCfgEnhVirtServicesFifthPartTableSendRST = 3
	SlbCurCfgEnhVirtServicesFifthPartTableSendRST_Unsupported SlbCurCfgEnhVirtServicesFifthPartTableSendRST = 2147483647
)

type SlbCurCfgEnhVirtServicesFifthPartTableClsOnSlowage int32

const (
	SlbCurCfgEnhVirtServicesFifthPartTableClsOnSlowage_None        SlbCurCfgEnhVirtServicesFifthPartTableClsOnSlowage = 1
	SlbCurCfgEnhVirtServicesFifthPartTableClsOnSlowage_Client      SlbCurCfgEnhVirtServicesFifthPartTableClsOnSlowage = 2
	SlbCurCfgEnhVirtServicesFifthPartTableClsOnSlowage_Server      SlbCurCfgEnhVirtServicesFifthPartTableClsOnSlowage = 3
	SlbCurCfgEnhVirtServicesFifthPartTableClsOnSlowage_Both        SlbCurCfgEnhVirtServicesFifthPartTableClsOnSlowage = 4
	SlbCurCfgEnhVirtServicesFifthPartTableClsOnSlowage_Unsupported SlbCurCfgEnhVirtServicesFifthPartTableClsOnSlowage = 2147483647
)

type SlbCurCfgEnhVirtServicesFifthPartTableCookieSameSite int32

const (
	SlbCurCfgEnhVirtServicesFifthPartTableCookieSameSite_Exclude     SlbCurCfgEnhVirtServicesFifthPartTableCookieSameSite = 0
	SlbCurCfgEnhVirtServicesFifthPartTableCookieSameSite_None        SlbCurCfgEnhVirtServicesFifthPartTableCookieSameSite = 1
	SlbCurCfgEnhVirtServicesFifthPartTableCookieSameSite_Lax         SlbCurCfgEnhVirtServicesFifthPartTableCookieSameSite = 2
	SlbCurCfgEnhVirtServicesFifthPartTableCookieSameSite_Strict      SlbCurCfgEnhVirtServicesFifthPartTableCookieSameSite = 3
	SlbCurCfgEnhVirtServicesFifthPartTableCookieSameSite_Unsupported SlbCurCfgEnhVirtServicesFifthPartTableCookieSameSite = 2147483647
)

type SlbCurCfgEnhVirtServicesFifthPartTableClsOnFastage int32

const (
	SlbCurCfgEnhVirtServicesFifthPartTableClsOnFastage_None        SlbCurCfgEnhVirtServicesFifthPartTableClsOnFastage = 1
	SlbCurCfgEnhVirtServicesFifthPartTableClsOnFastage_Client      SlbCurCfgEnhVirtServicesFifthPartTableClsOnFastage = 2
	SlbCurCfgEnhVirtServicesFifthPartTableClsOnFastage_Server      SlbCurCfgEnhVirtServicesFifthPartTableClsOnFastage = 3
	SlbCurCfgEnhVirtServicesFifthPartTableClsOnFastage_Both        SlbCurCfgEnhVirtServicesFifthPartTableClsOnFastage = 4
	SlbCurCfgEnhVirtServicesFifthPartTableClsOnFastage_Unsupported SlbCurCfgEnhVirtServicesFifthPartTableClsOnFastage = 2147483647
)

type SlbCurCfgEnhVirtServicesFifthPartTableCookieHttponly int32

const (
	SlbCurCfgEnhVirtServicesFifthPartTableCookieHttponly_No          SlbCurCfgEnhVirtServicesFifthPartTableCookieHttponly = 1
	SlbCurCfgEnhVirtServicesFifthPartTableCookieHttponly_Yes         SlbCurCfgEnhVirtServicesFifthPartTableCookieHttponly = 2
	SlbCurCfgEnhVirtServicesFifthPartTableCookieHttponly_Unsupported SlbCurCfgEnhVirtServicesFifthPartTableCookieHttponly = 2147483647
)

type SlbCurCfgEnhVirtServicesFifthPartTableParams struct {
	// The number of the virtual server.
	ServFifthPartIndex string `json:"ServFifthPartIndex,omitempty"`
	// The service index. This has no external meaning
	FifthPartIndex int32 `json:"FifthPartIndex,omitempty"`
	// Enter text to be replaced.
	ServTextrepMatchText string `json:"ServTextrepMatchText,omitempty"`
	// Enter Cur text.
	ServTextrepReplacTxt string `json:"ServTextrepReplacTxt,omitempty"`
	// Application Type for virtual service.
	ServApplicationType SlbCurCfgEnhVirtServicesFifthPartTableServApplicationType `json:"ServApplicationType,omitempty"`
	// The name of the virtual service.
	Name string `json:"Name,omitempty"`
	// Action type of the service.If we set value as group(1)
	// it will Load balance the traffic between the servers
	// defined in the group field after performing all other
	// services actions.when set to a value of redirect(2)
	// for http/https services, an http/s redirection will be
	// performed with the values set in the application
	// redirection.If we set value as discard(3) it will drop the session.
	Action SlbCurCfgEnhVirtServicesFifthPartTableAction `json:"Action,omitempty"`
	// application redirection location.We need to provide this
	// value When action type is set to redirect(2).
	Redirect string `json:"Redirect,omitempty"`
	// Group Mark for Server Certificate. If we set value as cert(0)
	// It will denote that the server certificate (name) associated with
	// this virtual service, represents a certificate. Otherwise, a value of group(1), denotes that the server
	// certificate (name) represents a group.
	ServCertGrpMark SlbCurCfgEnhVirtServicesFifthPartTableServCertGrpMark `json:"ServCertGrpMark,omitempty"`
	// Set DNS type for this service (DNS, DNSSEC).
	DnsType SlbCurCfgEnhVirtServicesFifthPartTableDnsType `json:"DnsType,omitempty"`
	// Set client proximity type for this service.
	ClntproxType SlbCurCfgEnhVirtServicesFifthPartTableClntproxType `json:"ClntproxType,omitempty"`
	// Enable or disable zero window size in SYN+ACK for this service.
	ZerowinSize SlbCurCfgEnhVirtServicesFifthPartTableZerowinSize `json:"ZerowinSize,omitempty"`
	// The cookie path name of the virtual server used for cookie load balance.
	CookiePath string `json:"CookiePath,omitempty"`
	// Is cookie secure [yes/no] [Default: no].
	CookieSecure SlbCurCfgEnhVirtServicesFifthPartTableCookieSecure `json:"CookieSecure,omitempty"`
	// Enable or disable only rtsp SLB for this service.
	NoRtsp SlbCurCfgEnhVirtServicesFifthPartTableNoRtsp `json:"NoRtsp,omitempty"`
	// Server rebalancing when cookie is absent.
	// When enabled, server load balancing will happen
	// for subsequent request comes without cookie.
	CkRebind SlbCurCfgEnhVirtServicesFifthPartTableCkRebind `json:"CkRebind,omitempty"`
	// Enable or disable buffer limit for content based selection.
	ParseLimit SlbCurCfgEnhVirtServicesFifthPartTableParseLimit `json:"ParseLimit,omitempty"`
	// The buffer length for content based selection.
	ParseLength uint32 `json:"ParseLength,omitempty"`
	// Enable or disable URI normalization for HTTP content matching.
	UriNorm SlbCurCfgEnhVirtServicesFifthPartTableUriNorm `json:"UriNorm,omitempty"`
	// Get the current Granularity of this service, for statistics report-protocol information.
	// Group(1) - for group level, or GroupAndServers(2) - for server level.
	Granularity SlbCurCfgEnhVirtServicesFifthPartTableGranularity `json:"Granularity,omitempty"`
	// Enable or disable Session Logging.
	SessLog SlbCurCfgEnhVirtServicesFifthPartTableSessLog `json:"SessLog,omitempty"`
	// Fast aging of UDP sessions.
	UdpAge SlbCurCfgEnhVirtServicesFifthPartTableUdpAge `json:"UdpAge,omitempty"`
	// Session entry mode.
	SessEntryMode SlbCurCfgEnhVirtServicesFifthPartTableSessEntryMode `json:"SessEntryMode,omitempty"`
	// Security policy name associated with this virtual service.
	SecPol string `json:"SecPol,omitempty"`
	// service always on when AS++ script attached.
	AlwaysOn SlbCurCfgEnhVirtServicesFifthPartTableAlwaysOn `json:"AlwaysOn,omitempty"`
	// Enable/Disable sending reset/icmp-err when the service is down.
	SendRST SlbCurCfgEnhVirtServicesFifthPartTableSendRST `json:"SendRST,omitempty"`
	// Set close connection on aging treatment.
	ClsOnSlowage SlbCurCfgEnhVirtServicesFifthPartTableClsOnSlowage `json:"ClsOnSlowage,omitempty"`
	// The cookie samesite attribute of the virtual server used for cookie load balance.
	CookieSameSite SlbCurCfgEnhVirtServicesFifthPartTableCookieSameSite `json:"CookieSameSite,omitempty"`
	// Set close connection on fastaging treatment.
	ClsOnFastage SlbCurCfgEnhVirtServicesFifthPartTableClsOnFastage `json:"ClsOnFastage,omitempty"`
	// GM SSL Server encryption Certificate name associated with this virtual service.
	ServCertEnc string `json:"ServCertEnc,omitempty"`
	// GM SSL Server sign Certificate name associated with this virtual service.
	ServCertSign string `json:"ServCertSign,omitempty"`
	// Is cookie http only [yes/no] [Default: no].
	CookieHttponly SlbCurCfgEnhVirtServicesFifthPartTableCookieHttponly `json:"CookieHttponly,omitempty"`
}

func (p SlbCurCfgEnhVirtServicesFifthPartTableParams) iMABean() {}
