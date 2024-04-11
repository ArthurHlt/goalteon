package beans

import (
	"fmt"
	"reflect"
)

// SlbCurCfgVirtServicesFifthPartTable The table of virtual Services.
// Note:This mib is not supported for VX instance of Virtualization.
type SlbCurCfgVirtServicesFifthPartTable struct {
	// The number of the virtual server.
	SlbCurCfgVirtServFifthPartIndex int32
	// The service index. This has no external meaning
	SlbCurCfgVirtServiceFifthPartIndex int32
	Params                             *SlbCurCfgVirtServicesFifthPartTableParams
}

func NewSlbCurCfgVirtServicesFifthPartTableList() *SlbCurCfgVirtServicesFifthPartTable {
	return &SlbCurCfgVirtServicesFifthPartTable{}
}

func NewSlbCurCfgVirtServicesFifthPartTable(
	slbCurCfgVirtServFifthPartIndex int32,
	slbCurCfgVirtServiceFifthPartIndex int32,
	params *SlbCurCfgVirtServicesFifthPartTableParams,
) *SlbCurCfgVirtServicesFifthPartTable {
	return &SlbCurCfgVirtServicesFifthPartTable{
		SlbCurCfgVirtServFifthPartIndex:    slbCurCfgVirtServFifthPartIndex,
		SlbCurCfgVirtServiceFifthPartIndex: slbCurCfgVirtServiceFifthPartIndex,
		Params:                             params,
	}
}

func (c *SlbCurCfgVirtServicesFifthPartTable) Name() string {
	return "SlbCurCfgVirtServicesFifthPartTable"
}

func (c *SlbCurCfgVirtServicesFifthPartTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbCurCfgVirtServicesFifthPartTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbCurCfgVirtServicesFifthPartTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbCurCfgVirtServFifthPartIndex).IsZero() &&
		reflect.ValueOf(c.SlbCurCfgVirtServiceFifthPartIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbCurCfgVirtServFifthPartIndex) + "/" + fmt.Sprint(c.SlbCurCfgVirtServiceFifthPartIndex)
}

type SlbCurCfgVirtServicesFifthPartTableServApplicationType int32

const (
	SlbCurCfgVirtServicesFifthPartTableServApplicationType_BasicSlb    SlbCurCfgVirtServicesFifthPartTableServApplicationType = 1
	SlbCurCfgVirtServicesFifthPartTableServApplicationType_Dns         SlbCurCfgVirtServicesFifthPartTableServApplicationType = 2
	SlbCurCfgVirtServicesFifthPartTableServApplicationType_Ftp         SlbCurCfgVirtServicesFifthPartTableServApplicationType = 3
	SlbCurCfgVirtServicesFifthPartTableServApplicationType_FtpData     SlbCurCfgVirtServicesFifthPartTableServApplicationType = 4
	SlbCurCfgVirtServicesFifthPartTableServApplicationType_Ldap        SlbCurCfgVirtServicesFifthPartTableServApplicationType = 5
	SlbCurCfgVirtServicesFifthPartTableServApplicationType_Http        SlbCurCfgVirtServicesFifthPartTableServApplicationType = 6
	SlbCurCfgVirtServicesFifthPartTableServApplicationType_Https       SlbCurCfgVirtServicesFifthPartTableServApplicationType = 7
	SlbCurCfgVirtServicesFifthPartTableServApplicationType_Ssl         SlbCurCfgVirtServicesFifthPartTableServApplicationType = 8
	SlbCurCfgVirtServicesFifthPartTableServApplicationType_Rtsp        SlbCurCfgVirtServicesFifthPartTableServApplicationType = 9
	SlbCurCfgVirtServicesFifthPartTableServApplicationType_Sip         SlbCurCfgVirtServicesFifthPartTableServApplicationType = 10
	SlbCurCfgVirtServicesFifthPartTableServApplicationType_Wts         SlbCurCfgVirtServicesFifthPartTableServApplicationType = 11
	SlbCurCfgVirtServicesFifthPartTableServApplicationType_Tftp        SlbCurCfgVirtServicesFifthPartTableServApplicationType = 12
	SlbCurCfgVirtServicesFifthPartTableServApplicationType_Smtp        SlbCurCfgVirtServicesFifthPartTableServApplicationType = 13
	SlbCurCfgVirtServicesFifthPartTableServApplicationType_Pop3        SlbCurCfgVirtServicesFifthPartTableServApplicationType = 14
	SlbCurCfgVirtServicesFifthPartTableServApplicationType_Ip          SlbCurCfgVirtServicesFifthPartTableServApplicationType = 15
	SlbCurCfgVirtServicesFifthPartTableServApplicationType_MhSctp      SlbCurCfgVirtServicesFifthPartTableServApplicationType = 16
	SlbCurCfgVirtServicesFifthPartTableServApplicationType_Unsupported SlbCurCfgVirtServicesFifthPartTableServApplicationType = 2147483647
)

type SlbCurCfgVirtServicesFifthPartTableAction int32

const (
	SlbCurCfgVirtServicesFifthPartTableAction_Group       SlbCurCfgVirtServicesFifthPartTableAction = 1
	SlbCurCfgVirtServicesFifthPartTableAction_Redirect    SlbCurCfgVirtServicesFifthPartTableAction = 2
	SlbCurCfgVirtServicesFifthPartTableAction_Discard     SlbCurCfgVirtServicesFifthPartTableAction = 3
	SlbCurCfgVirtServicesFifthPartTableAction_Unsupported SlbCurCfgVirtServicesFifthPartTableAction = 2147483647
)

type SlbCurCfgVirtServicesFifthPartTableServCertGrpMark int32

const (
	SlbCurCfgVirtServicesFifthPartTableServCertGrpMark_Cert        SlbCurCfgVirtServicesFifthPartTableServCertGrpMark = 1
	SlbCurCfgVirtServicesFifthPartTableServCertGrpMark_Group       SlbCurCfgVirtServicesFifthPartTableServCertGrpMark = 2
	SlbCurCfgVirtServicesFifthPartTableServCertGrpMark_Unsupported SlbCurCfgVirtServicesFifthPartTableServCertGrpMark = 2147483647
)

type SlbCurCfgVirtServicesFifthPartTableDnsType int32

const (
	SlbCurCfgVirtServicesFifthPartTableDnsType_Dns         SlbCurCfgVirtServicesFifthPartTableDnsType = 1
	SlbCurCfgVirtServicesFifthPartTableDnsType_Dnssec      SlbCurCfgVirtServicesFifthPartTableDnsType = 2
	SlbCurCfgVirtServicesFifthPartTableDnsType_Both        SlbCurCfgVirtServicesFifthPartTableDnsType = 3
	SlbCurCfgVirtServicesFifthPartTableDnsType_Unsupported SlbCurCfgVirtServicesFifthPartTableDnsType = 2147483647
)

type SlbCurCfgVirtServicesFifthPartTableClntproxType int32

const (
	SlbCurCfgVirtServicesFifthPartTableClntproxType_None        SlbCurCfgVirtServicesFifthPartTableClntproxType = 1
	SlbCurCfgVirtServicesFifthPartTableClntproxType_Http        SlbCurCfgVirtServicesFifthPartTableClntproxType = 2
	SlbCurCfgVirtServicesFifthPartTableClntproxType_Https       SlbCurCfgVirtServicesFifthPartTableClntproxType = 3
	SlbCurCfgVirtServicesFifthPartTableClntproxType_Unsupported SlbCurCfgVirtServicesFifthPartTableClntproxType = 2147483647
)

type SlbCurCfgVirtServicesFifthPartTableZerowinSize int32

const (
	SlbCurCfgVirtServicesFifthPartTableZerowinSize_Enabled     SlbCurCfgVirtServicesFifthPartTableZerowinSize = 1
	SlbCurCfgVirtServicesFifthPartTableZerowinSize_Disabled    SlbCurCfgVirtServicesFifthPartTableZerowinSize = 2
	SlbCurCfgVirtServicesFifthPartTableZerowinSize_Unsupported SlbCurCfgVirtServicesFifthPartTableZerowinSize = 2147483647
)

type SlbCurCfgVirtServicesFifthPartTableCookieSecure int32

const (
	SlbCurCfgVirtServicesFifthPartTableCookieSecure_No          SlbCurCfgVirtServicesFifthPartTableCookieSecure = 1
	SlbCurCfgVirtServicesFifthPartTableCookieSecure_Yes         SlbCurCfgVirtServicesFifthPartTableCookieSecure = 2
	SlbCurCfgVirtServicesFifthPartTableCookieSecure_Unsupported SlbCurCfgVirtServicesFifthPartTableCookieSecure = 2147483647
)

type SlbCurCfgVirtServicesFifthPartTableNoRtsp int32

const (
	SlbCurCfgVirtServicesFifthPartTableNoRtsp_Enabled     SlbCurCfgVirtServicesFifthPartTableNoRtsp = 1
	SlbCurCfgVirtServicesFifthPartTableNoRtsp_Disabled    SlbCurCfgVirtServicesFifthPartTableNoRtsp = 2
	SlbCurCfgVirtServicesFifthPartTableNoRtsp_Unsupported SlbCurCfgVirtServicesFifthPartTableNoRtsp = 2147483647
)

type SlbCurCfgVirtServicesFifthPartTableCkRebind int32

const (
	SlbCurCfgVirtServicesFifthPartTableCkRebind_Enabled     SlbCurCfgVirtServicesFifthPartTableCkRebind = 1
	SlbCurCfgVirtServicesFifthPartTableCkRebind_Disabled    SlbCurCfgVirtServicesFifthPartTableCkRebind = 2
	SlbCurCfgVirtServicesFifthPartTableCkRebind_Unsupported SlbCurCfgVirtServicesFifthPartTableCkRebind = 2147483647
)

type SlbCurCfgVirtServicesFifthPartTableParseLimit int32

const (
	SlbCurCfgVirtServicesFifthPartTableParseLimit_Enabled     SlbCurCfgVirtServicesFifthPartTableParseLimit = 1
	SlbCurCfgVirtServicesFifthPartTableParseLimit_Disabled    SlbCurCfgVirtServicesFifthPartTableParseLimit = 2
	SlbCurCfgVirtServicesFifthPartTableParseLimit_Unsupported SlbCurCfgVirtServicesFifthPartTableParseLimit = 2147483647
)

type SlbCurCfgVirtServicesFifthPartTableUriNorm int32

const (
	SlbCurCfgVirtServicesFifthPartTableUriNorm_Enabled     SlbCurCfgVirtServicesFifthPartTableUriNorm = 1
	SlbCurCfgVirtServicesFifthPartTableUriNorm_Disabled    SlbCurCfgVirtServicesFifthPartTableUriNorm = 2
	SlbCurCfgVirtServicesFifthPartTableUriNorm_Unsupported SlbCurCfgVirtServicesFifthPartTableUriNorm = 2147483647
)

type SlbCurCfgVirtServicesFifthPartTableGranularity int32

const (
	SlbCurCfgVirtServicesFifthPartTableGranularity_Service     SlbCurCfgVirtServicesFifthPartTableGranularity = 0
	SlbCurCfgVirtServicesFifthPartTableGranularity_Real        SlbCurCfgVirtServicesFifthPartTableGranularity = 2
	SlbCurCfgVirtServicesFifthPartTableGranularity_Unsupported SlbCurCfgVirtServicesFifthPartTableGranularity = 2147483647
)

type SlbCurCfgVirtServicesFifthPartTableSessLog int32

const (
	SlbCurCfgVirtServicesFifthPartTableSessLog_Enabled     SlbCurCfgVirtServicesFifthPartTableSessLog = 1
	SlbCurCfgVirtServicesFifthPartTableSessLog_Disabled    SlbCurCfgVirtServicesFifthPartTableSessLog = 2
	SlbCurCfgVirtServicesFifthPartTableSessLog_Unsupported SlbCurCfgVirtServicesFifthPartTableSessLog = 2147483647
)

type SlbCurCfgVirtServicesFifthPartTableParams struct {
	// The number of the virtual server.
	ServFifthPartIndex int32 `json:"ServFifthPartIndex,omitempty"`
	// The service index. This has no external meaning
	FifthPartIndex int32 `json:"FifthPartIndex,omitempty"`
	// Enter text to be replaced.
	ServTextrepMatchText string `json:"ServTextrepMatchText,omitempty"`
	// Enter Cur text.
	ServTextrepReplacTxt string `json:"ServTextrepReplacTxt,omitempty"`
	// Application Type for virtual service.
	ServApplicationType SlbCurCfgVirtServicesFifthPartTableServApplicationType `json:"ServApplicationType,omitempty"`
	// The name of the virtual service.
	Name string `json:"Name,omitempty"`
	// Action type of the service.If we set value as group(1)
	// it will Load balance the traffic between the servers
	// defined in the group field after performing all other
	// services actions.when set to a value of redirect(2)
	// for http/https services, an http/s redirection will be
	// performed with the values set in the application
	// redirection.If we set value as discard(3) it will drop the session.
	Action SlbCurCfgVirtServicesFifthPartTableAction `json:"Action,omitempty"`
	// application redirection location.We need to provide this
	// value When action type is set to redirect(2).
	Redirect string `json:"Redirect,omitempty"`
	// Group Mark for Server Certificate. If we set value as cert(0)
	// It will denote that the server certificate (name) associated with
	// this virtual service, represents a certificate. Otherwise, a value of group(1), denotes that the server
	// certificate (name) represents a group.
	ServCertGrpMark SlbCurCfgVirtServicesFifthPartTableServCertGrpMark `json:"ServCertGrpMark,omitempty"`
	// Set DNS type for this service (DNS, DNSSEC).
	DnsType SlbCurCfgVirtServicesFifthPartTableDnsType `json:"DnsType,omitempty"`
	// Set client proximity type for this service.
	ClntproxType SlbCurCfgVirtServicesFifthPartTableClntproxType `json:"ClntproxType,omitempty"`
	// Enable or disable zero window size in SYN+ACK for this service.
	ZerowinSize SlbCurCfgVirtServicesFifthPartTableZerowinSize `json:"ZerowinSize,omitempty"`
	// The cookie path name of the virtual server used for cookie load balance.
	CookiePath string `json:"CookiePath,omitempty"`
	// Is cookie secure [yes/no] [Default: no].
	CookieSecure SlbCurCfgVirtServicesFifthPartTableCookieSecure `json:"CookieSecure,omitempty"`
	// Enable or disable only rtsp SLB for this service.
	NoRtsp SlbCurCfgVirtServicesFifthPartTableNoRtsp `json:"NoRtsp,omitempty"`
	// Server rebalancing when cookie is absent.
	// When enabled, server load balancing will happen
	// for subsequent request comes without cookie.
	CkRebind SlbCurCfgVirtServicesFifthPartTableCkRebind `json:"CkRebind,omitempty"`
	// Enable or disable buffer limit for content based selection.
	ParseLimit SlbCurCfgVirtServicesFifthPartTableParseLimit `json:"ParseLimit,omitempty"`
	// The buffer length for content based selection.
	ParseLength uint32 `json:"ParseLength,omitempty"`
	// Enable or disable URI normalization for HTTP content matching.
	UriNorm SlbCurCfgVirtServicesFifthPartTableUriNorm `json:"UriNorm,omitempty"`
	// Get the current Granularity of this service, for statistics report-protocol information.
	// 	 Group(1) - for group level, or GroupAndServers(2) - for server level.
	Granularity SlbCurCfgVirtServicesFifthPartTableGranularity `json:"Granularity,omitempty"`
	// Enable or disable Session Logging.
	SessLog SlbCurCfgVirtServicesFifthPartTableSessLog `json:"SessLog,omitempty"`
	// This object ID shows current configured web security ID.
	AppwallWebappId string `json:"AppwallWebappId,omitempty"`
}

func (p SlbCurCfgVirtServicesFifthPartTableParams) iMABean() {}
