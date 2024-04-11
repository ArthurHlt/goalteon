package beans

import (
	"fmt"
	"reflect"
)

// GslbCurCfgRuleTable The rule table in the current configuration block.
// Note:This mib is not supported for VX instance of Virtualization.
type GslbCurCfgRuleTable struct {
	// The rule table index in the current configuration block.
	GslbCurCfgRuleIndx int32
	Params             *GslbCurCfgRuleTableParams
}

func NewGslbCurCfgRuleTableList() *GslbCurCfgRuleTable {
	return &GslbCurCfgRuleTable{}
}

func NewGslbCurCfgRuleTable(
	gslbCurCfgRuleIndx int32,
	params *GslbCurCfgRuleTableParams,
) *GslbCurCfgRuleTable {
	return &GslbCurCfgRuleTable{
		GslbCurCfgRuleIndx: gslbCurCfgRuleIndx,
		Params:             params,
	}
}

func (c *GslbCurCfgRuleTable) Name() string {
	return "GslbCurCfgRuleTable"
}

func (c *GslbCurCfgRuleTable) GetParams() BeanType {
	return c.Params
}

func (c *GslbCurCfgRuleTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *GslbCurCfgRuleTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.GslbCurCfgRuleIndx).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.GslbCurCfgRuleIndx)
}

type GslbCurCfgRuleTableState int32

const (
	GslbCurCfgRuleTableState_Enabled     GslbCurCfgRuleTableState = 1
	GslbCurCfgRuleTableState_Disabled    GslbCurCfgRuleTableState = 2
	GslbCurCfgRuleTableState_Unsupported GslbCurCfgRuleTableState = 2147483647
)

type GslbCurCfgRuleTableType int32

const (
	GslbCurCfgRuleTableType_Gslb        GslbCurCfgRuleTableType = 0
	GslbCurCfgRuleTableType_Inboundllb  GslbCurCfgRuleTableType = 1
	GslbCurCfgRuleTableType_Unsupported GslbCurCfgRuleTableType = 2147483647
)

type GslbCurCfgRuleTableEdnsPrst int32

const (
	GslbCurCfgRuleTableEdnsPrst_Enabled     GslbCurCfgRuleTableEdnsPrst = 1
	GslbCurCfgRuleTableEdnsPrst_Disabled    GslbCurCfgRuleTableEdnsPrst = 2
	GslbCurCfgRuleTableEdnsPrst_Unsupported GslbCurCfgRuleTableEdnsPrst = 2147483647
)

type GslbCurCfgRuleTablePersist int32

const (
	GslbCurCfgRuleTablePersist_Domain      GslbCurCfgRuleTablePersist = 1
	GslbCurCfgRuleTablePersist_Ip          GslbCurCfgRuleTablePersist = 2
	GslbCurCfgRuleTablePersist_Unsupported GslbCurCfgRuleTablePersist = 2147483647
)

type GslbCurCfgRuleTableNetworkFallback int32

const (
	GslbCurCfgRuleTableNetworkFallback_Enabled     GslbCurCfgRuleTableNetworkFallback = 1
	GslbCurCfgRuleTableNetworkFallback_Disabled    GslbCurCfgRuleTableNetworkFallback = 2
	GslbCurCfgRuleTableNetworkFallback_Unsupported GslbCurCfgRuleTableNetworkFallback = 2147483647
)

type GslbCurCfgRuleTableParams struct {
	// The rule table index in the current configuration block.
	Indx int32 `json:"Indx,omitempty"`
	// Enable/Disable the rule in
	// the current configuration block.
	State GslbCurCfgRuleTableState `json:"State,omitempty"`
	// Start hour in 24-hour format for the rule in
	// the current configuration block.
	StartHour uint32 `json:"StartHour,omitempty"`
	// Start minutes for the rule in
	// the current configuration block.
	StartMin uint64 `json:"StartMin,omitempty"`
	// End hour in 24-hour format for the rule in
	// the current configuration block.
	EndHour uint32 `json:"EndHour,omitempty"`
	// End minutes for the rule in
	// the current configuration block.
	EndMin uint64 `json:"EndMin,omitempty"`
	// Time to live in seconds of DNS resource records for the rule in
	// the current configuration block.
	TTL uint64 `json:"TTL,omitempty"`
	// DNS resource records in DNS response for the rule in
	// the current configuration block.
	RR uint32 `json:"RR,omitempty"`
	// Network preference domain name for rule.
	Dname string `json:"Dname,omitempty"`
	// Source IP netmask for phash in
	// the current configuration block.
	IpNetMask string `json:"IpNetMask,omitempty"`
	// Timeout in minutes for phash in
	// the current configuration block.
	Timeout uint32 `json:"Timeout,omitempty"`
	// Source IPv6 prefix for phash in
	// the current configuration block.
	Ipv6Prefix uint32 `json:"Ipv6Prefix,omitempty"`
	// Rule type.
	Type GslbCurCfgRuleTableType `json:"Type,omitempty"`
	// Rule name
	Name string `json:"Name,omitempty"`
	// Rule port.
	Port uint64 `json:"Port,omitempty"`
	// Persist on EDNS value.
	EdnsPrst GslbCurCfgRuleTableEdnsPrst `json:"EdnsPrst,omitempty"`
	// Persistence/hash parameter for rule.
	Persist GslbCurCfgRuleTablePersist `json:"Persist,omitempty"`
	// Current state of Rule Network metric fallback.
	NetworkFallback GslbCurCfgRuleTableNetworkFallback `json:"NetworkFallback,omitempty"`
	// DNS Nameserver group for rule.
	Namesrvr string `json:"Namesrvr,omitempty"`
}

func (p GslbCurCfgRuleTableParams) iMABean() {}
