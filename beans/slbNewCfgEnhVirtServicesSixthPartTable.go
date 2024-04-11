package beans

import (
	"fmt"
	"reflect"
)

// SlbNewCfgEnhVirtServicesSixthPartTable The table of virtual Services.
// Note:This mib is not supported for VX instance of Virtualization.
type SlbNewCfgEnhVirtServicesSixthPartTable struct {
	// The number of the virtual server.
	SlbNewCfgEnhVirtServSixthPartIndex string
	// The service index. This has no external meaning
	SlbNewCfgEnhVirtServiceSixthPartIndex int32
	Params                                *SlbNewCfgEnhVirtServicesSixthPartTableParams
}

func NewSlbNewCfgEnhVirtServicesSixthPartTableList() *SlbNewCfgEnhVirtServicesSixthPartTable {
	return &SlbNewCfgEnhVirtServicesSixthPartTable{}
}

func NewSlbNewCfgEnhVirtServicesSixthPartTable(
	slbNewCfgEnhVirtServSixthPartIndex string,
	slbNewCfgEnhVirtServiceSixthPartIndex int32,
	params *SlbNewCfgEnhVirtServicesSixthPartTableParams,
) *SlbNewCfgEnhVirtServicesSixthPartTable {
	return &SlbNewCfgEnhVirtServicesSixthPartTable{
		SlbNewCfgEnhVirtServSixthPartIndex:    slbNewCfgEnhVirtServSixthPartIndex,
		SlbNewCfgEnhVirtServiceSixthPartIndex: slbNewCfgEnhVirtServiceSixthPartIndex,
		Params:                                params,
	}
}

func (c *SlbNewCfgEnhVirtServicesSixthPartTable) Name() string {
	return "SlbNewCfgEnhVirtServicesSixthPartTable"
}

func (c *SlbNewCfgEnhVirtServicesSixthPartTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbNewCfgEnhVirtServicesSixthPartTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbNewCfgEnhVirtServicesSixthPartTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbNewCfgEnhVirtServSixthPartIndex).IsZero() &&
		reflect.ValueOf(c.SlbNewCfgEnhVirtServiceSixthPartIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbNewCfgEnhVirtServSixthPartIndex) + "/" + fmt.Sprint(c.SlbNewCfgEnhVirtServiceSixthPartIndex)
}

type SlbNewCfgEnhVirtServicesSixthPartTableDummyDelete int32

const (
	SlbNewCfgEnhVirtServicesSixthPartTableDummyDelete_Other       SlbNewCfgEnhVirtServicesSixthPartTableDummyDelete = 1
	SlbNewCfgEnhVirtServicesSixthPartTableDummyDelete_Delete      SlbNewCfgEnhVirtServicesSixthPartTableDummyDelete = 2
	SlbNewCfgEnhVirtServicesSixthPartTableDummyDelete_Unsupported SlbNewCfgEnhVirtServicesSixthPartTableDummyDelete = 2147483647
)

type SlbNewCfgEnhVirtServicesSixthPartTableDirect int32

const (
	SlbNewCfgEnhVirtServicesSixthPartTableDirect_Enabled     SlbNewCfgEnhVirtServicesSixthPartTableDirect = 1
	SlbNewCfgEnhVirtServicesSixthPartTableDirect_Disabled    SlbNewCfgEnhVirtServicesSixthPartTableDirect = 2
	SlbNewCfgEnhVirtServicesSixthPartTableDirect_Unsupported SlbNewCfgEnhVirtServicesSixthPartTableDirect = 2147483647
)

type SlbNewCfgEnhVirtServicesSixthPartTableThash int32

const (
	SlbNewCfgEnhVirtServicesSixthPartTableThash_Sip         SlbNewCfgEnhVirtServicesSixthPartTableThash = 1
	SlbNewCfgEnhVirtServicesSixthPartTableThash_SipSport    SlbNewCfgEnhVirtServicesSixthPartTableThash = 2
	SlbNewCfgEnhVirtServicesSixthPartTableThash_Unsupported SlbNewCfgEnhVirtServicesSixthPartTableThash = 2147483647
)

type SlbNewCfgEnhVirtServicesSixthPartTableLdapreset int32

const (
	SlbNewCfgEnhVirtServicesSixthPartTableLdapreset_Enabled     SlbNewCfgEnhVirtServicesSixthPartTableLdapreset = 1
	SlbNewCfgEnhVirtServicesSixthPartTableLdapreset_Disabled    SlbNewCfgEnhVirtServicesSixthPartTableLdapreset = 2
	SlbNewCfgEnhVirtServicesSixthPartTableLdapreset_Unsupported SlbNewCfgEnhVirtServicesSixthPartTableLdapreset = 2147483647
)

type SlbNewCfgEnhVirtServicesSixthPartTableLdapslb int32

const (
	SlbNewCfgEnhVirtServicesSixthPartTableLdapslb_Enabled     SlbNewCfgEnhVirtServicesSixthPartTableLdapslb = 1
	SlbNewCfgEnhVirtServicesSixthPartTableLdapslb_Disabled    SlbNewCfgEnhVirtServicesSixthPartTableLdapslb = 2
	SlbNewCfgEnhVirtServicesSixthPartTableLdapslb_Unsupported SlbNewCfgEnhVirtServicesSixthPartTableLdapslb = 2147483647
)

type SlbNewCfgEnhVirtServicesSixthPartTableSip int32

const (
	SlbNewCfgEnhVirtServicesSixthPartTableSip_Enabled     SlbNewCfgEnhVirtServicesSixthPartTableSip = 1
	SlbNewCfgEnhVirtServicesSixthPartTableSip_Disabled    SlbNewCfgEnhVirtServicesSixthPartTableSip = 2
	SlbNewCfgEnhVirtServicesSixthPartTableSip_Unsupported SlbNewCfgEnhVirtServicesSixthPartTableSip = 2147483647
)

type SlbNewCfgEnhVirtServicesSixthPartTableXForwardedFor int32

const (
	SlbNewCfgEnhVirtServicesSixthPartTableXForwardedFor_Enabled     SlbNewCfgEnhVirtServicesSixthPartTableXForwardedFor = 1
	SlbNewCfgEnhVirtServicesSixthPartTableXForwardedFor_Disabled    SlbNewCfgEnhVirtServicesSixthPartTableXForwardedFor = 2
	SlbNewCfgEnhVirtServicesSixthPartTableXForwardedFor_Unsupported SlbNewCfgEnhVirtServicesSixthPartTableXForwardedFor = 2147483647
)

type SlbNewCfgEnhVirtServicesSixthPartTableHttpRedir int32

const (
	SlbNewCfgEnhVirtServicesSixthPartTableHttpRedir_Enabled     SlbNewCfgEnhVirtServicesSixthPartTableHttpRedir = 1
	SlbNewCfgEnhVirtServicesSixthPartTableHttpRedir_Disabled    SlbNewCfgEnhVirtServicesSixthPartTableHttpRedir = 2
	SlbNewCfgEnhVirtServicesSixthPartTableHttpRedir_Unsupported SlbNewCfgEnhVirtServicesSixthPartTableHttpRedir = 2147483647
)

type SlbNewCfgEnhVirtServicesSixthPartTablePbindRport int32

const (
	SlbNewCfgEnhVirtServicesSixthPartTablePbindRport_Enabled     SlbNewCfgEnhVirtServicesSixthPartTablePbindRport = 1
	SlbNewCfgEnhVirtServicesSixthPartTablePbindRport_Disabled    SlbNewCfgEnhVirtServicesSixthPartTablePbindRport = 2
	SlbNewCfgEnhVirtServicesSixthPartTablePbindRport_Unsupported SlbNewCfgEnhVirtServicesSixthPartTablePbindRport = 2147483647
)

type SlbNewCfgEnhVirtServicesSixthPartTableEgressPip int32

const (
	SlbNewCfgEnhVirtServicesSixthPartTableEgressPip_Enabled     SlbNewCfgEnhVirtServicesSixthPartTableEgressPip = 1
	SlbNewCfgEnhVirtServicesSixthPartTableEgressPip_Disabled    SlbNewCfgEnhVirtServicesSixthPartTableEgressPip = 2
	SlbNewCfgEnhVirtServicesSixthPartTableEgressPip_Unsupported SlbNewCfgEnhVirtServicesSixthPartTableEgressPip = 2147483647
)

type SlbNewCfgEnhVirtServicesSixthPartTableCookieDname int32

const (
	SlbNewCfgEnhVirtServicesSixthPartTableCookieDname_Enabled     SlbNewCfgEnhVirtServicesSixthPartTableCookieDname = 1
	SlbNewCfgEnhVirtServicesSixthPartTableCookieDname_Disabled    SlbNewCfgEnhVirtServicesSixthPartTableCookieDname = 2
	SlbNewCfgEnhVirtServicesSixthPartTableCookieDname_Unsupported SlbNewCfgEnhVirtServicesSixthPartTableCookieDname = 2147483647
)

type SlbNewCfgEnhVirtServicesSixthPartTableWts int32

const (
	SlbNewCfgEnhVirtServicesSixthPartTableWts_Enabled     SlbNewCfgEnhVirtServicesSixthPartTableWts = 1
	SlbNewCfgEnhVirtServicesSixthPartTableWts_Disabled    SlbNewCfgEnhVirtServicesSixthPartTableWts = 2
	SlbNewCfgEnhVirtServicesSixthPartTableWts_Unsupported SlbNewCfgEnhVirtServicesSixthPartTableWts = 2147483647
)

type SlbNewCfgEnhVirtServicesSixthPartTableUhash int32

const (
	SlbNewCfgEnhVirtServicesSixthPartTableUhash_Enabled     SlbNewCfgEnhVirtServicesSixthPartTableUhash = 1
	SlbNewCfgEnhVirtServicesSixthPartTableUhash_Disabled    SlbNewCfgEnhVirtServicesSixthPartTableUhash = 2
	SlbNewCfgEnhVirtServicesSixthPartTableUhash_Unsupported SlbNewCfgEnhVirtServicesSixthPartTableUhash = 2147483647
)

type SlbNewCfgEnhVirtServicesSixthPartTableSdpNat int32

const (
	SlbNewCfgEnhVirtServicesSixthPartTableSdpNat_Enabled     SlbNewCfgEnhVirtServicesSixthPartTableSdpNat = 1
	SlbNewCfgEnhVirtServicesSixthPartTableSdpNat_Disabled    SlbNewCfgEnhVirtServicesSixthPartTableSdpNat = 2
	SlbNewCfgEnhVirtServicesSixthPartTableSdpNat_Unsupported SlbNewCfgEnhVirtServicesSixthPartTableSdpNat = 2147483647
)

type SlbNewCfgEnhVirtServicesSixthPartTableParams struct {
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
	// This mib is dummy,the main Delete mib is in slbNewCfgEnhVirtServicesTable
	// When read, other(1) is returned.
	DummyDelete SlbNewCfgEnhVirtServicesSixthPartTableDummyDelete `json:"DummyDelete,omitempty"`
	// Enable or disable DAM for this service.
	Direct SlbNewCfgEnhVirtServicesSixthPartTableDirect `json:"Direct,omitempty"`
	// Set hash parameter.
	Thash SlbNewCfgEnhVirtServicesSixthPartTableThash `json:"Thash,omitempty"`
	// Enable or disable LDAP Server Reset
	Ldapreset SlbNewCfgEnhVirtServicesSixthPartTableLdapreset `json:"Ldapreset,omitempty"`
	// Enable or disable LDAP Server load balancing
	Ldapslb SlbNewCfgEnhVirtServicesSixthPartTableLdapslb `json:"Ldapslb,omitempty"`
	// Enable/disable SIP load balancing.
	Sip SlbNewCfgEnhVirtServicesSixthPartTableSip `json:"Sip,omitempty"`
	// Enable/disable X-Forwarded-For for proxy mode.
	XForwardedFor SlbNewCfgEnhVirtServicesSixthPartTableXForwardedFor `json:"XForwardedFor,omitempty"`
	// Enable/disable HTTP/HTTPS redirect for GSLB.
	HttpRedir SlbNewCfgEnhVirtServicesSixthPartTableHttpRedir `json:"HttpRedir,omitempty"`
	// Enable or disable use of rport in the session lookup for a persistent
	// session.
	PbindRport SlbNewCfgEnhVirtServicesSixthPartTablePbindRport `json:"PbindRport,omitempty"`
	// Enable/disable pip selection based on egress port/vlan.
	EgressPip SlbNewCfgEnhVirtServicesSixthPartTableEgressPip `json:"EgressPip,omitempty"`
	// Select dname for insert cookie persistence mode.
	CookieDname SlbNewCfgEnhVirtServicesSixthPartTableCookieDname `json:"CookieDname,omitempty"`
	// Enable or disable WTS loadbalancing and persistence.
	Wts SlbNewCfgEnhVirtServicesSixthPartTableWts `json:"Wts,omitempty"`
	// Enable when there is no Session Directory server.
	Uhash SlbNewCfgEnhVirtServicesSixthPartTableUhash `json:"Uhash,omitempty"`
	// The maximum number of minutes an inactive connection remains open.
	TimeOut uint32 `json:"TimeOut,omitempty"`
	// Enable/disable SIP Media portal NAT.
	SdpNat SlbNewCfgEnhVirtServicesSixthPartTableSdpNat `json:"SdpNat,omitempty"`
	// Set ip address header.
	IpHeader string `json:"IpHeader,omitempty"`
	// Set ip address header set by the user.
	UserDefinedIpHeader string `json:"UserDefinedIpHeader,omitempty"`
}

func (p SlbNewCfgEnhVirtServicesSixthPartTableParams) iMABean() {}
