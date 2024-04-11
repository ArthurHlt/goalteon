package beans

import (
	"fmt"
	"reflect"
)

// SlbNewCfgVirtServicesFifthPartTable The table of virtual Services.
// Note:This mib is not supported for VX instance of Virtualization.
type SlbNewCfgVirtServicesFifthPartTable struct {
	// The number of the virtual server.
	SlbNewCfgVirtServFifthPartIndex int32
	// The service index. This has no external meaning
	SlbNewCfgVirtServiceFifthPartIndex int32
	Params                             *SlbNewCfgVirtServicesFifthPartTableParams
}

func NewSlbNewCfgVirtServicesFifthPartTableList() *SlbNewCfgVirtServicesFifthPartTable {
	return &SlbNewCfgVirtServicesFifthPartTable{}
}

func NewSlbNewCfgVirtServicesFifthPartTable(
	slbNewCfgVirtServFifthPartIndex int32,
	slbNewCfgVirtServiceFifthPartIndex int32,
	params *SlbNewCfgVirtServicesFifthPartTableParams,
) *SlbNewCfgVirtServicesFifthPartTable {
	return &SlbNewCfgVirtServicesFifthPartTable{
		SlbNewCfgVirtServFifthPartIndex:    slbNewCfgVirtServFifthPartIndex,
		SlbNewCfgVirtServiceFifthPartIndex: slbNewCfgVirtServiceFifthPartIndex,
		Params:                             params,
	}
}

func (c *SlbNewCfgVirtServicesFifthPartTable) Name() string {
	return "SlbNewCfgVirtServicesFifthPartTable"
}

func (c *SlbNewCfgVirtServicesFifthPartTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbNewCfgVirtServicesFifthPartTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbNewCfgVirtServicesFifthPartTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbNewCfgVirtServFifthPartIndex).IsZero() &&
		reflect.ValueOf(c.SlbNewCfgVirtServiceFifthPartIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbNewCfgVirtServFifthPartIndex) + "/" + fmt.Sprint(c.SlbNewCfgVirtServiceFifthPartIndex)
}

type SlbNewCfgVirtServicesFifthPartTableServApplicationType int32

const (
	SlbNewCfgVirtServicesFifthPartTableServApplicationType_BasicSlb SlbNewCfgVirtServicesFifthPartTableServApplicationType = 1
	SlbNewCfgVirtServicesFifthPartTableServApplicationType_Dns      SlbNewCfgVirtServicesFifthPartTableServApplicationType = 2
	SlbNewCfgVirtServicesFifthPartTableServApplicationType_Ftp      SlbNewCfgVirtServicesFifthPartTableServApplicationType = 3
	SlbNewCfgVirtServicesFifthPartTableServApplicationType_FtpData  SlbNewCfgVirtServicesFifthPartTableServApplicationType = 4
	SlbNewCfgVirtServicesFifthPartTableServApplicationType_Ldap     SlbNewCfgVirtServicesFifthPartTableServApplicationType = 5
	SlbNewCfgVirtServicesFifthPartTableServApplicationType_Http     SlbNewCfgVirtServicesFifthPartTableServApplicationType = 6
	SlbNewCfgVirtServicesFifthPartTableServApplicationType_Https    SlbNewCfgVirtServicesFifthPartTableServApplicationType = 7
	SlbNewCfgVirtServicesFifthPartTableServApplicationType_Ssl      SlbNewCfgVirtServicesFifthPartTableServApplicationType = 8
	SlbNewCfgVirtServicesFifthPartTableServApplicationType_Rtsp     SlbNewCfgVirtServicesFifthPartTableServApplicationType = 9
	SlbNewCfgVirtServicesFifthPartTableServApplicationType_Sip      SlbNewCfgVirtServicesFifthPartTableServApplicationType = 10
	SlbNewCfgVirtServicesFifthPartTableServApplicationType_Wts      SlbNewCfgVirtServicesFifthPartTableServApplicationType = 11
	SlbNewCfgVirtServicesFifthPartTableServApplicationType_Tftp     SlbNewCfgVirtServicesFifthPartTableServApplicationType = 12
	SlbNewCfgVirtServicesFifthPartTableServApplicationType_Smtp     SlbNewCfgVirtServicesFifthPartTableServApplicationType = 13
	SlbNewCfgVirtServicesFifthPartTableServApplicationType_Pop3     SlbNewCfgVirtServicesFifthPartTableServApplicationType = 14
	SlbNewCfgVirtServicesFifthPartTableServApplicationType_Ip       SlbNewCfgVirtServicesFifthPartTableServApplicationType = 15
	SlbNewCfgVirtServicesFifthPartTableServApplicationType_MhSctp   SlbNewCfgVirtServicesFifthPartTableServApplicationType = 16
)

type SlbNewCfgVirtServicesFifthPartTableAction int32

const (
	SlbNewCfgVirtServicesFifthPartTableAction_Group       SlbNewCfgVirtServicesFifthPartTableAction = 1
	SlbNewCfgVirtServicesFifthPartTableAction_Redirect    SlbNewCfgVirtServicesFifthPartTableAction = 2
	SlbNewCfgVirtServicesFifthPartTableAction_Discard     SlbNewCfgVirtServicesFifthPartTableAction = 3
	SlbNewCfgVirtServicesFifthPartTableAction_Unsupported SlbNewCfgVirtServicesFifthPartTableAction = 2147483647
)

type SlbNewCfgVirtServicesFifthPartTableServCertGrpMark int32

const (
	SlbNewCfgVirtServicesFifthPartTableServCertGrpMark_Cert        SlbNewCfgVirtServicesFifthPartTableServCertGrpMark = 1
	SlbNewCfgVirtServicesFifthPartTableServCertGrpMark_Group       SlbNewCfgVirtServicesFifthPartTableServCertGrpMark = 2
	SlbNewCfgVirtServicesFifthPartTableServCertGrpMark_Unsupported SlbNewCfgVirtServicesFifthPartTableServCertGrpMark = 2147483647
)

type SlbNewCfgVirtServicesFifthPartTableDnsType int32

const (
	SlbNewCfgVirtServicesFifthPartTableDnsType_Dns         SlbNewCfgVirtServicesFifthPartTableDnsType = 1
	SlbNewCfgVirtServicesFifthPartTableDnsType_Dnssec      SlbNewCfgVirtServicesFifthPartTableDnsType = 2
	SlbNewCfgVirtServicesFifthPartTableDnsType_Both        SlbNewCfgVirtServicesFifthPartTableDnsType = 3
	SlbNewCfgVirtServicesFifthPartTableDnsType_Unsupported SlbNewCfgVirtServicesFifthPartTableDnsType = 2147483647
)

type SlbNewCfgVirtServicesFifthPartTableClntproxType int32

const (
	SlbNewCfgVirtServicesFifthPartTableClntproxType_None        SlbNewCfgVirtServicesFifthPartTableClntproxType = 1
	SlbNewCfgVirtServicesFifthPartTableClntproxType_Http        SlbNewCfgVirtServicesFifthPartTableClntproxType = 2
	SlbNewCfgVirtServicesFifthPartTableClntproxType_Https       SlbNewCfgVirtServicesFifthPartTableClntproxType = 3
	SlbNewCfgVirtServicesFifthPartTableClntproxType_Unsupported SlbNewCfgVirtServicesFifthPartTableClntproxType = 2147483647
)

type SlbNewCfgVirtServicesFifthPartTableZerowinSize int32

const (
	SlbNewCfgVirtServicesFifthPartTableZerowinSize_Enabled     SlbNewCfgVirtServicesFifthPartTableZerowinSize = 1
	SlbNewCfgVirtServicesFifthPartTableZerowinSize_Disabled    SlbNewCfgVirtServicesFifthPartTableZerowinSize = 2
	SlbNewCfgVirtServicesFifthPartTableZerowinSize_Unsupported SlbNewCfgVirtServicesFifthPartTableZerowinSize = 2147483647
)

type SlbNewCfgVirtServicesFifthPartTableCookieSecure int32

const (
	SlbNewCfgVirtServicesFifthPartTableCookieSecure_No          SlbNewCfgVirtServicesFifthPartTableCookieSecure = 1
	SlbNewCfgVirtServicesFifthPartTableCookieSecure_Yes         SlbNewCfgVirtServicesFifthPartTableCookieSecure = 2
	SlbNewCfgVirtServicesFifthPartTableCookieSecure_Unsupported SlbNewCfgVirtServicesFifthPartTableCookieSecure = 2147483647
)

type SlbNewCfgVirtServicesFifthPartTableNoRtsp int32

const (
	SlbNewCfgVirtServicesFifthPartTableNoRtsp_Enabled     SlbNewCfgVirtServicesFifthPartTableNoRtsp = 1
	SlbNewCfgVirtServicesFifthPartTableNoRtsp_Disabled    SlbNewCfgVirtServicesFifthPartTableNoRtsp = 2
	SlbNewCfgVirtServicesFifthPartTableNoRtsp_Unsupported SlbNewCfgVirtServicesFifthPartTableNoRtsp = 2147483647
)

type SlbNewCfgVirtServicesFifthPartTableCkRebind int32

const (
	SlbNewCfgVirtServicesFifthPartTableCkRebind_Enabled     SlbNewCfgVirtServicesFifthPartTableCkRebind = 1
	SlbNewCfgVirtServicesFifthPartTableCkRebind_Disabled    SlbNewCfgVirtServicesFifthPartTableCkRebind = 2
	SlbNewCfgVirtServicesFifthPartTableCkRebind_Unsupported SlbNewCfgVirtServicesFifthPartTableCkRebind = 2147483647
)

type SlbNewCfgVirtServicesFifthPartTableParseLimit int32

const (
	SlbNewCfgVirtServicesFifthPartTableParseLimit_Enabled     SlbNewCfgVirtServicesFifthPartTableParseLimit = 1
	SlbNewCfgVirtServicesFifthPartTableParseLimit_Disabled    SlbNewCfgVirtServicesFifthPartTableParseLimit = 2
	SlbNewCfgVirtServicesFifthPartTableParseLimit_Unsupported SlbNewCfgVirtServicesFifthPartTableParseLimit = 2147483647
)

type SlbNewCfgVirtServicesFifthPartTableUriNorm int32

const (
	SlbNewCfgVirtServicesFifthPartTableUriNorm_Enabled     SlbNewCfgVirtServicesFifthPartTableUriNorm = 1
	SlbNewCfgVirtServicesFifthPartTableUriNorm_Disabled    SlbNewCfgVirtServicesFifthPartTableUriNorm = 2
	SlbNewCfgVirtServicesFifthPartTableUriNorm_Unsupported SlbNewCfgVirtServicesFifthPartTableUriNorm = 2147483647
)

type SlbNewCfgVirtServicesFifthPartTableGranularity int32

const (
	SlbNewCfgVirtServicesFifthPartTableGranularity_Service     SlbNewCfgVirtServicesFifthPartTableGranularity = 0
	SlbNewCfgVirtServicesFifthPartTableGranularity_Real        SlbNewCfgVirtServicesFifthPartTableGranularity = 2
	SlbNewCfgVirtServicesFifthPartTableGranularity_Unsupported SlbNewCfgVirtServicesFifthPartTableGranularity = 2147483647
)

type SlbNewCfgVirtServicesFifthPartTableSessLog int32

const (
	SlbNewCfgVirtServicesFifthPartTableSessLog_Enabled     SlbNewCfgVirtServicesFifthPartTableSessLog = 1
	SlbNewCfgVirtServicesFifthPartTableSessLog_Disabled    SlbNewCfgVirtServicesFifthPartTableSessLog = 2
	SlbNewCfgVirtServicesFifthPartTableSessLog_Unsupported SlbNewCfgVirtServicesFifthPartTableSessLog = 2147483647
)

type SlbNewCfgVirtServicesFifthPartTableParams struct {
	// The number of the virtual server.
	ServFifthPartIndex int32 `json:"ServFifthPartIndex,omitempty"`
	// The service index. This has no external meaning
	FifthPartIndex int32 `json:"FifthPartIndex,omitempty"`
	// Enter text to be replaced.
	ServTextrepMatchText string `json:"ServTextrepMatchText,omitempty"`
	// Enter new text.
	ServTextrepReplacTxt string `json:"ServTextrepReplacTxt,omitempty"`
	// Application Type for virtual service.
	ServApplicationType SlbNewCfgVirtServicesFifthPartTableServApplicationType `json:"ServApplicationType,omitempty"`
	// The name of the virtual service.
	Name string `json:"Name,omitempty"`
	// Action type of the service.If we set value as group(1)
	// it will Load balance the traffic between the servers
	// defined in the group field after performing all other
	// services actions.when set to a value of redirect(2)
	// for http/https services, an http/s redirection will be
	// performed with the values set in the application
	// redirection.If we set value as discard(3) it will drop the session.
	Action SlbNewCfgVirtServicesFifthPartTableAction `json:"Action,omitempty"`
	// application redirection location.We need to provide this
	// value When action type is set to redirect(2).
	Redirect string `json:"Redirect,omitempty"`
	// Group Mark for Server Certificate. If we set value as cert(0)
	// It will denote that the server certificate (name) associated with
	// this virtual service, represents a certificate. Otherwise, a value of group(1), denotes that the server
	// certificate (name) represents a group.
	ServCertGrpMark SlbNewCfgVirtServicesFifthPartTableServCertGrpMark `json:"ServCertGrpMark,omitempty"`
	// Set DNS type for this service (DNS, DNSSEC).
	DnsType SlbNewCfgVirtServicesFifthPartTableDnsType `json:"DnsType,omitempty"`
	// Set client proximity type for this service.
	ClntproxType SlbNewCfgVirtServicesFifthPartTableClntproxType `json:"ClntproxType,omitempty"`
	// Enable or disable zero window size in SYN+ACK for this service.
	ZerowinSize SlbNewCfgVirtServicesFifthPartTableZerowinSize `json:"ZerowinSize,omitempty"`
	// The cookie path name of the virtual server used for cookie load balance.
	CookiePath string `json:"CookiePath,omitempty"`
	// Is cookie secure [yes/no] [Default: no].
	CookieSecure SlbNewCfgVirtServicesFifthPartTableCookieSecure `json:"CookieSecure,omitempty"`
	// Enable or disable only rtsp SLB for this service.
	NoRtsp SlbNewCfgVirtServicesFifthPartTableNoRtsp `json:"NoRtsp,omitempty"`
	// Enable or disable server rebalancing when cookie is absent.
	// When enabled, server load balancing will happen
	// for subsequent request comes without cookie.
	CkRebind SlbNewCfgVirtServicesFifthPartTableCkRebind `json:"CkRebind,omitempty"`
	// Enable or disable buffer limit for content based selection.
	ParseLimit SlbNewCfgVirtServicesFifthPartTableParseLimit `json:"ParseLimit,omitempty"`
	// The buffer length for content based selection.
	ParseLength uint32 `json:"ParseLength,omitempty"`
	// Enable or disable URI normalization for HTTP content matching.
	UriNorm SlbNewCfgVirtServicesFifthPartTableUriNorm `json:"UriNorm,omitempty"`
	// Sets the Granularity for this service, for statistics report-protocol information.
	// 	Group(1) - for group level, or GroupAndServers(2) - for server level.
	Granularity SlbNewCfgVirtServicesFifthPartTableGranularity `json:"Granularity,omitempty"`
	// Enable or disable Session Logging.
	SessLog SlbNewCfgVirtServicesFifthPartTableSessLog `json:"SessLog,omitempty"`
	// This object ID allows configuring the web security ID
	AppwallWebappId string `json:"AppwallWebappId,omitempty"`
}

func (p SlbNewCfgVirtServicesFifthPartTableParams) iMABean() {}
