package beans

import (
	"fmt"
	"reflect"
)

// SlbCurCfgEnhVirtServicesSixthPartTable The table of virtual Services.
// Note:This mib is not supported for VX instance of Virtualization.
type SlbCurCfgEnhVirtServicesSixthPartTable struct {
	// The number of the virtual server.
	SlbCurCfgEnhVirtServSixthPartIndex string
	// The service index. This has no external meaning
	SlbCurCfgEnhVirtServiceSixthPartIndex int32
	Params                                *SlbCurCfgEnhVirtServicesSixthPartTableParams
}

func NewSlbCurCfgEnhVirtServicesSixthPartTableList() *SlbCurCfgEnhVirtServicesSixthPartTable {
	return &SlbCurCfgEnhVirtServicesSixthPartTable{}
}

func NewSlbCurCfgEnhVirtServicesSixthPartTable(
	slbCurCfgEnhVirtServSixthPartIndex string,
	slbCurCfgEnhVirtServiceSixthPartIndex int32,
	params *SlbCurCfgEnhVirtServicesSixthPartTableParams,
) *SlbCurCfgEnhVirtServicesSixthPartTable {
	return &SlbCurCfgEnhVirtServicesSixthPartTable{
		SlbCurCfgEnhVirtServSixthPartIndex:    slbCurCfgEnhVirtServSixthPartIndex,
		SlbCurCfgEnhVirtServiceSixthPartIndex: slbCurCfgEnhVirtServiceSixthPartIndex,
		Params:                                params,
	}
}

func (c *SlbCurCfgEnhVirtServicesSixthPartTable) Name() string {
	return "SlbCurCfgEnhVirtServicesSixthPartTable"
}

func (c *SlbCurCfgEnhVirtServicesSixthPartTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbCurCfgEnhVirtServicesSixthPartTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbCurCfgEnhVirtServicesSixthPartTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbCurCfgEnhVirtServSixthPartIndex).IsZero() &&
		reflect.ValueOf(c.SlbCurCfgEnhVirtServiceSixthPartIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbCurCfgEnhVirtServSixthPartIndex) + "/" + fmt.Sprint(c.SlbCurCfgEnhVirtServiceSixthPartIndex)
}

type SlbCurCfgEnhVirtServicesSixthPartTableDirect int32

const (
	SlbCurCfgEnhVirtServicesSixthPartTableDirect_Enabled     SlbCurCfgEnhVirtServicesSixthPartTableDirect = 1
	SlbCurCfgEnhVirtServicesSixthPartTableDirect_Disabled    SlbCurCfgEnhVirtServicesSixthPartTableDirect = 2
	SlbCurCfgEnhVirtServicesSixthPartTableDirect_Unsupported SlbCurCfgEnhVirtServicesSixthPartTableDirect = 2147483647
)

type SlbCurCfgEnhVirtServicesSixthPartTableThash int32

const (
	SlbCurCfgEnhVirtServicesSixthPartTableThash_Sip         SlbCurCfgEnhVirtServicesSixthPartTableThash = 1
	SlbCurCfgEnhVirtServicesSixthPartTableThash_SipSport    SlbCurCfgEnhVirtServicesSixthPartTableThash = 2
	SlbCurCfgEnhVirtServicesSixthPartTableThash_Unsupported SlbCurCfgEnhVirtServicesSixthPartTableThash = 2147483647
)

type SlbCurCfgEnhVirtServicesSixthPartTableLdapreset int32

const (
	SlbCurCfgEnhVirtServicesSixthPartTableLdapreset_Enabled     SlbCurCfgEnhVirtServicesSixthPartTableLdapreset = 1
	SlbCurCfgEnhVirtServicesSixthPartTableLdapreset_Disabled    SlbCurCfgEnhVirtServicesSixthPartTableLdapreset = 2
	SlbCurCfgEnhVirtServicesSixthPartTableLdapreset_Unsupported SlbCurCfgEnhVirtServicesSixthPartTableLdapreset = 2147483647
)

type SlbCurCfgEnhVirtServicesSixthPartTableLdapslb int32

const (
	SlbCurCfgEnhVirtServicesSixthPartTableLdapslb_Enabled     SlbCurCfgEnhVirtServicesSixthPartTableLdapslb = 1
	SlbCurCfgEnhVirtServicesSixthPartTableLdapslb_Disabled    SlbCurCfgEnhVirtServicesSixthPartTableLdapslb = 2
	SlbCurCfgEnhVirtServicesSixthPartTableLdapslb_Unsupported SlbCurCfgEnhVirtServicesSixthPartTableLdapslb = 2147483647
)

type SlbCurCfgEnhVirtServicesSixthPartTableSip int32

const (
	SlbCurCfgEnhVirtServicesSixthPartTableSip_Enabled     SlbCurCfgEnhVirtServicesSixthPartTableSip = 1
	SlbCurCfgEnhVirtServicesSixthPartTableSip_Disabled    SlbCurCfgEnhVirtServicesSixthPartTableSip = 2
	SlbCurCfgEnhVirtServicesSixthPartTableSip_Unsupported SlbCurCfgEnhVirtServicesSixthPartTableSip = 2147483647
)

type SlbCurCfgEnhVirtServicesSixthPartTableXForwardedFor int32

const (
	SlbCurCfgEnhVirtServicesSixthPartTableXForwardedFor_Enabled     SlbCurCfgEnhVirtServicesSixthPartTableXForwardedFor = 1
	SlbCurCfgEnhVirtServicesSixthPartTableXForwardedFor_Disabled    SlbCurCfgEnhVirtServicesSixthPartTableXForwardedFor = 2
	SlbCurCfgEnhVirtServicesSixthPartTableXForwardedFor_Unsupported SlbCurCfgEnhVirtServicesSixthPartTableXForwardedFor = 2147483647
)

type SlbCurCfgEnhVirtServicesSixthPartTableHttpRedir int32

const (
	SlbCurCfgEnhVirtServicesSixthPartTableHttpRedir_Enabled     SlbCurCfgEnhVirtServicesSixthPartTableHttpRedir = 1
	SlbCurCfgEnhVirtServicesSixthPartTableHttpRedir_Disabled    SlbCurCfgEnhVirtServicesSixthPartTableHttpRedir = 2
	SlbCurCfgEnhVirtServicesSixthPartTableHttpRedir_Unsupported SlbCurCfgEnhVirtServicesSixthPartTableHttpRedir = 2147483647
)

type SlbCurCfgEnhVirtServicesSixthPartTablePbindRport int32

const (
	SlbCurCfgEnhVirtServicesSixthPartTablePbindRport_Enabled     SlbCurCfgEnhVirtServicesSixthPartTablePbindRport = 1
	SlbCurCfgEnhVirtServicesSixthPartTablePbindRport_Disabled    SlbCurCfgEnhVirtServicesSixthPartTablePbindRport = 2
	SlbCurCfgEnhVirtServicesSixthPartTablePbindRport_Unsupported SlbCurCfgEnhVirtServicesSixthPartTablePbindRport = 2147483647
)

type SlbCurCfgEnhVirtServicesSixthPartTableEgressPip int32

const (
	SlbCurCfgEnhVirtServicesSixthPartTableEgressPip_Enabled     SlbCurCfgEnhVirtServicesSixthPartTableEgressPip = 1
	SlbCurCfgEnhVirtServicesSixthPartTableEgressPip_Disabled    SlbCurCfgEnhVirtServicesSixthPartTableEgressPip = 2
	SlbCurCfgEnhVirtServicesSixthPartTableEgressPip_Unsupported SlbCurCfgEnhVirtServicesSixthPartTableEgressPip = 2147483647
)

type SlbCurCfgEnhVirtServicesSixthPartTableCookieDname int32

const (
	SlbCurCfgEnhVirtServicesSixthPartTableCookieDname_Enabled     SlbCurCfgEnhVirtServicesSixthPartTableCookieDname = 1
	SlbCurCfgEnhVirtServicesSixthPartTableCookieDname_Disabled    SlbCurCfgEnhVirtServicesSixthPartTableCookieDname = 2
	SlbCurCfgEnhVirtServicesSixthPartTableCookieDname_Unsupported SlbCurCfgEnhVirtServicesSixthPartTableCookieDname = 2147483647
)

type SlbCurCfgEnhVirtServicesSixthPartTableWts int32

const (
	SlbCurCfgEnhVirtServicesSixthPartTableWts_Enabled     SlbCurCfgEnhVirtServicesSixthPartTableWts = 1
	SlbCurCfgEnhVirtServicesSixthPartTableWts_Disabled    SlbCurCfgEnhVirtServicesSixthPartTableWts = 2
	SlbCurCfgEnhVirtServicesSixthPartTableWts_Unsupported SlbCurCfgEnhVirtServicesSixthPartTableWts = 2147483647
)

type SlbCurCfgEnhVirtServicesSixthPartTableUhash int32

const (
	SlbCurCfgEnhVirtServicesSixthPartTableUhash_Enabled     SlbCurCfgEnhVirtServicesSixthPartTableUhash = 1
	SlbCurCfgEnhVirtServicesSixthPartTableUhash_Disabled    SlbCurCfgEnhVirtServicesSixthPartTableUhash = 2
	SlbCurCfgEnhVirtServicesSixthPartTableUhash_Unsupported SlbCurCfgEnhVirtServicesSixthPartTableUhash = 2147483647
)

type SlbCurCfgEnhVirtServicesSixthPartTableSdpNat int32

const (
	SlbCurCfgEnhVirtServicesSixthPartTableSdpNat_Enabled     SlbCurCfgEnhVirtServicesSixthPartTableSdpNat = 1
	SlbCurCfgEnhVirtServicesSixthPartTableSdpNat_Disabled    SlbCurCfgEnhVirtServicesSixthPartTableSdpNat = 2
	SlbCurCfgEnhVirtServicesSixthPartTableSdpNat_Unsupported SlbCurCfgEnhVirtServicesSixthPartTableSdpNat = 2147483647
)

type SlbCurCfgEnhVirtServicesSixthPartTableParams struct {
	// The number of the virtual server.
	ServSixthPartIndex string `json:"ServSixthPartIndex,omitempty"`
	// The service index. This has no external meaning
	SixthPartIndex int32 `json:"SixthPartIndex,omitempty"`
	// The host name of the virtual service.
	Hname string `json:"Hname,omitempty"`
	// The cookie name of the virtual server used for cookie load balance.
	Cname string `json:"Cname,omitempty"`
	// The cookie expire of the virtual server used for insert cookie load
	// balance depending on the mode it has the following format
	// <MM/dd/yy[@hh:mm]> absolute mode or <days[:hours[:minutes]]>
	// for relative mode.
	CExpire string `json:"CExpire,omitempty"`
	// The number of bytes used to hash onto server, A zero means
	// URL hashing disabled.
	UrlHashLen uint32 `json:"UrlHashLen,omitempty"`
	// Enable or disable DAM for this service.
	Direct SlbCurCfgEnhVirtServicesSixthPartTableDirect `json:"Direct,omitempty"`
	// Set hash parameter.
	Thash SlbCurCfgEnhVirtServicesSixthPartTableThash `json:"Thash,omitempty"`
	// Enable or disable LDAP Server Reset
	Ldapreset SlbCurCfgEnhVirtServicesSixthPartTableLdapreset `json:"Ldapreset,omitempty"`
	// Enable or disable LDAP Server load balancing
	Ldapslb SlbCurCfgEnhVirtServicesSixthPartTableLdapslb `json:"Ldapslb,omitempty"`
	// Enable/disable SIP load balancing.
	Sip SlbCurCfgEnhVirtServicesSixthPartTableSip `json:"Sip,omitempty"`
	// Enable/disable X-Forwarded-For for proxy mode.
	XForwardedFor SlbCurCfgEnhVirtServicesSixthPartTableXForwardedFor `json:"XForwardedFor,omitempty"`
	// Enable/disable HTTP/HTTPS redirect for GSLB.
	HttpRedir SlbCurCfgEnhVirtServicesSixthPartTableHttpRedir `json:"HttpRedir,omitempty"`
	// Enable or disable use of rport in the session lookup for a persistent
	// session.
	PbindRport SlbCurCfgEnhVirtServicesSixthPartTablePbindRport `json:"PbindRport,omitempty"`
	// Enable/disable pip selection based on egress port/vlan.
	EgressPip SlbCurCfgEnhVirtServicesSixthPartTableEgressPip `json:"EgressPip,omitempty"`
	// Select dname for insert cookie persistence mode.
	CookieDname SlbCurCfgEnhVirtServicesSixthPartTableCookieDname `json:"CookieDname,omitempty"`
	// Enable or disable WTS loadbalancing and persistence.
	Wts SlbCurCfgEnhVirtServicesSixthPartTableWts `json:"Wts,omitempty"`
	// Enable when there is no Session Directory server.
	Uhash SlbCurCfgEnhVirtServicesSixthPartTableUhash `json:"Uhash,omitempty"`
	// The maximum number of minutes an inactive connection remains open.
	TimeOut uint32 `json:"TimeOut,omitempty"`
	// Enable/disable SIP Media portal NAT.
	SdpNat SlbCurCfgEnhVirtServicesSixthPartTableSdpNat `json:"SdpNat,omitempty"`
	// Set ip address header.
	IpHeader string `json:"IpHeader,omitempty"`
	// Set ip address header.
	UserDefinedIpHeader string `json:"UserDefinedIpHeader,omitempty"`
}

func (p SlbCurCfgEnhVirtServicesSixthPartTableParams) iMABean() {}
