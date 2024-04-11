package beans

import (
	"fmt"
	"reflect"
)

// GslbNewCfgRuleTable The rule table in the new configuration block.
// Note:This mib is not supported for VX instance of Virtualization.
type GslbNewCfgRuleTable struct {
	// The rule table index in the new configuration block.
	GslbNewCfgRuleIndx int32
	Params             *GslbNewCfgRuleTableParams
}

func NewGslbNewCfgRuleTableList() *GslbNewCfgRuleTable {
	return &GslbNewCfgRuleTable{}
}

func NewGslbNewCfgRuleTable(
	gslbNewCfgRuleIndx int32,
	params *GslbNewCfgRuleTableParams,
) *GslbNewCfgRuleTable {
	return &GslbNewCfgRuleTable{
		GslbNewCfgRuleIndx: gslbNewCfgRuleIndx,
		Params:             params,
	}
}

func (c *GslbNewCfgRuleTable) Name() string {
	return "GslbNewCfgRuleTable"
}

func (c *GslbNewCfgRuleTable) GetParams() BeanType {
	return c.Params
}

func (c *GslbNewCfgRuleTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *GslbNewCfgRuleTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.GslbNewCfgRuleIndx).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.GslbNewCfgRuleIndx)
}

type GslbNewCfgRuleTableState int32

const (
	GslbNewCfgRuleTableState_Enabled     GslbNewCfgRuleTableState = 1
	GslbNewCfgRuleTableState_Disabled    GslbNewCfgRuleTableState = 2
	GslbNewCfgRuleTableState_Unsupported GslbNewCfgRuleTableState = 2147483647
)

type GslbNewCfgRuleTableDelete int32

const (
	GslbNewCfgRuleTableDelete_Other       GslbNewCfgRuleTableDelete = 1
	GslbNewCfgRuleTableDelete_Delete      GslbNewCfgRuleTableDelete = 2
	GslbNewCfgRuleTableDelete_Unsupported GslbNewCfgRuleTableDelete = 2147483647
)

type GslbNewCfgRuleTableType int32

const (
	GslbNewCfgRuleTableType_Gslb        GslbNewCfgRuleTableType = 0
	GslbNewCfgRuleTableType_Inboundllb  GslbNewCfgRuleTableType = 1
	GslbNewCfgRuleTableType_Unsupported GslbNewCfgRuleTableType = 2147483647
)

type GslbNewCfgRuleTableEdnsPrst int32

const (
	GslbNewCfgRuleTableEdnsPrst_Enabled     GslbNewCfgRuleTableEdnsPrst = 1
	GslbNewCfgRuleTableEdnsPrst_Disabled    GslbNewCfgRuleTableEdnsPrst = 2
	GslbNewCfgRuleTableEdnsPrst_Unsupported GslbNewCfgRuleTableEdnsPrst = 2147483647
)

type GslbNewCfgRuleTablePersist int32

const (
	GslbNewCfgRuleTablePersist_Domain      GslbNewCfgRuleTablePersist = 1
	GslbNewCfgRuleTablePersist_Ip          GslbNewCfgRuleTablePersist = 2
	GslbNewCfgRuleTablePersist_Unsupported GslbNewCfgRuleTablePersist = 2147483647
)

type GslbNewCfgRuleTableNetworkFallback int32

const (
	GslbNewCfgRuleTableNetworkFallback_Enabled     GslbNewCfgRuleTableNetworkFallback = 1
	GslbNewCfgRuleTableNetworkFallback_Disabled    GslbNewCfgRuleTableNetworkFallback = 2
	GslbNewCfgRuleTableNetworkFallback_Unsupported GslbNewCfgRuleTableNetworkFallback = 2147483647
)

type GslbNewCfgRuleTableParams struct {
	// The rule table index in the new configuration block.
	Indx int32 `json:"Indx,omitempty"`
	// Enable/Disable the rule in
	// the new configuration block.
	State GslbNewCfgRuleTableState `json:"State,omitempty"`
	// Start hour in 24-hour format for the rule in
	// the new configuration block.
	StartHour uint32 `json:"StartHour,omitempty"`
	// Start minutes for the rule in
	// the new configuration block.
	StartMin uint64 `json:"StartMin,omitempty"`
	// End hour in 24-hour format for the rule in
	// the new configuration block.
	EndHour uint32 `json:"EndHour,omitempty"`
	// End minutes for the rule in
	// the new configuration block.
	EndMin uint64 `json:"EndMin,omitempty"`
	// Time to live in seconds of DNS resource records for the rule in
	// the new configuration block.
	TTL uint64 `json:"TTL,omitempty"`
	// DNS resource records in DNS response for the rule in
	// the new configuration block.
	RR uint32 `json:"RR,omitempty"`
	// By setting the value to delete(2), the entire row is deleted.
	Delete GslbNewCfgRuleTableDelete `json:"Delete,omitempty"`
	// Network preference domain name for rule
	Dname string `json:"Dname,omitempty"`
	// Source IP netmask for phash in
	// the new configuration block.
	IpNetMask string `json:"IpNetMask,omitempty"`
	// Timeout in minutes for phash in
	// the new configuration block.
	Timeout uint32 `json:"Timeout,omitempty"`
	// Source IPv6 prefix for phash in
	// the new configuration block.
	Ipv6Prefix uint32 `json:"Ipv6Prefix,omitempty"`
	// Rule type.
	Type GslbNewCfgRuleTableType `json:"Type,omitempty"`
	// Rule name
	Name string `json:"Name,omitempty"`
	// Rule port.
	Port uint64 `json:"Port,omitempty"`
	// Persist on EDNS value.
	EdnsPrst GslbNewCfgRuleTableEdnsPrst `json:"EdnsPrst,omitempty"`
	// Set persistence/hash parameter for rule.
	Persist GslbNewCfgRuleTablePersist `json:"Persist,omitempty"`
	// Set Rule Network metric fallback to enable/disable.
	NetworkFallback GslbNewCfgRuleTableNetworkFallback `json:"NetworkFallback,omitempty"`
	// DNS Nameserver group for rule
	Namesrvr string `json:"Namesrvr,omitempty"`
}

func (p GslbNewCfgRuleTableParams) iMABean() {}
