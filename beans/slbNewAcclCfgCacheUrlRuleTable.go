package beans

import (
	"fmt"
	"reflect"
)

// SlbNewAcclCfgCacheUrlRuleTable The table for configuring caching URL Rules.
// Note:This mib is not supported for VX instance of Virtualization.
type SlbNewAcclCfgCacheUrlRuleTable struct {
	// The caching policy URL LIST (key id) as an index.
	SlbNewAcclCfgCacheUrlRuleListIdIndex string
	// The caching URL Rule number as an index.
	SlbNewAcclCfgCacheUrlRuleIndex int32
	Params                         *SlbNewAcclCfgCacheUrlRuleTableParams
}

func NewSlbNewAcclCfgCacheUrlRuleTableList() *SlbNewAcclCfgCacheUrlRuleTable {
	return &SlbNewAcclCfgCacheUrlRuleTable{}
}

func NewSlbNewAcclCfgCacheUrlRuleTable(
	slbNewAcclCfgCacheUrlRuleListIdIndex string,
	slbNewAcclCfgCacheUrlRuleIndex int32,
	params *SlbNewAcclCfgCacheUrlRuleTableParams,
) *SlbNewAcclCfgCacheUrlRuleTable {
	return &SlbNewAcclCfgCacheUrlRuleTable{
		SlbNewAcclCfgCacheUrlRuleListIdIndex: slbNewAcclCfgCacheUrlRuleListIdIndex,
		SlbNewAcclCfgCacheUrlRuleIndex:       slbNewAcclCfgCacheUrlRuleIndex,
		Params:                               params,
	}
}

func (c *SlbNewAcclCfgCacheUrlRuleTable) Name() string {
	return "SlbNewAcclCfgCacheUrlRuleTable"
}

func (c *SlbNewAcclCfgCacheUrlRuleTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbNewAcclCfgCacheUrlRuleTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbNewAcclCfgCacheUrlRuleTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbNewAcclCfgCacheUrlRuleListIdIndex).IsZero() &&
		reflect.ValueOf(c.SlbNewAcclCfgCacheUrlRuleIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbNewAcclCfgCacheUrlRuleListIdIndex) + "/" + fmt.Sprint(c.SlbNewAcclCfgCacheUrlRuleIndex)
}

type SlbNewAcclCfgCacheUrlRuleTableDomainM int32

const (
	SlbNewAcclCfgCacheUrlRuleTableDomainM_Any         SlbNewAcclCfgCacheUrlRuleTableDomainM = 1
	SlbNewAcclCfgCacheUrlRuleTableDomainM_Text        SlbNewAcclCfgCacheUrlRuleTableDomainM = 2
	SlbNewAcclCfgCacheUrlRuleTableDomainM_Regex       SlbNewAcclCfgCacheUrlRuleTableDomainM = 3
	SlbNewAcclCfgCacheUrlRuleTableDomainM_Unsupported SlbNewAcclCfgCacheUrlRuleTableDomainM = 2147483647
)

type SlbNewAcclCfgCacheUrlRuleTableURLm int32

const (
	SlbNewAcclCfgCacheUrlRuleTableURLm_Any         SlbNewAcclCfgCacheUrlRuleTableURLm = 1
	SlbNewAcclCfgCacheUrlRuleTableURLm_Text        SlbNewAcclCfgCacheUrlRuleTableURLm = 2
	SlbNewAcclCfgCacheUrlRuleTableURLm_Regex       SlbNewAcclCfgCacheUrlRuleTableURLm = 3
	SlbNewAcclCfgCacheUrlRuleTableURLm_Unsupported SlbNewAcclCfgCacheUrlRuleTableURLm = 2147483647
)

type SlbNewAcclCfgCacheUrlRuleTableCache int32

const (
	SlbNewAcclCfgCacheUrlRuleTableCache_Enabled     SlbNewAcclCfgCacheUrlRuleTableCache = 1
	SlbNewAcclCfgCacheUrlRuleTableCache_Disabled    SlbNewAcclCfgCacheUrlRuleTableCache = 2
	SlbNewAcclCfgCacheUrlRuleTableCache_Unsupported SlbNewAcclCfgCacheUrlRuleTableCache = 2147483647
)

type SlbNewAcclCfgCacheUrlRuleTableAdminStatus int32

const (
	SlbNewAcclCfgCacheUrlRuleTableAdminStatus_Enabled     SlbNewAcclCfgCacheUrlRuleTableAdminStatus = 1
	SlbNewAcclCfgCacheUrlRuleTableAdminStatus_Disabled    SlbNewAcclCfgCacheUrlRuleTableAdminStatus = 2
	SlbNewAcclCfgCacheUrlRuleTableAdminStatus_Unsupported SlbNewAcclCfgCacheUrlRuleTableAdminStatus = 2147483647
)

type SlbNewAcclCfgCacheUrlRuleTableDelete int32

const (
	SlbNewAcclCfgCacheUrlRuleTableDelete_Other       SlbNewAcclCfgCacheUrlRuleTableDelete = 1
	SlbNewAcclCfgCacheUrlRuleTableDelete_Delete      SlbNewAcclCfgCacheUrlRuleTableDelete = 2
	SlbNewAcclCfgCacheUrlRuleTableDelete_Unsupported SlbNewAcclCfgCacheUrlRuleTableDelete = 2147483647
)

type SlbNewAcclCfgCacheUrlRuleTableParams struct {
	// The caching policy URL LIST (key id) as an index.
	ListIdIndex string `json:"ListIdIndex,omitempty"`
	// The caching URL Rule number as an index.
	Index int32 `json:"Index,omitempty"`
	// The caching URL Rule name.
	Name string `json:"Name,omitempty"`
	// Defines whether Domain matching should be evaluated as String, Regex or match, any (default any).
	DomainM SlbNewAcclCfgCacheUrlRuleTableDomainM `json:"DomainM,omitempty"`
	// Defines the Virtual Host for which this rule applies.
	Domain string `json:"Domain,omitempty"`
	// Defines whether URL matching should be evaluated as String, Regex or match, any (default any).
	URLm SlbNewAcclCfgCacheUrlRuleTableURLm `json:"URLm,omitempty"`
	// Defines URL of specific object (file/folder) to be matched by this rule.
	URL string `json:"URL,omitempty"`
	// The Maximum time for the cache object to served from cache (in seconds).
	Expire int32 `json:"Expire,omitempty"`
	// Defines if the matched response should be cached or not.
	Cache SlbNewAcclCfgCacheUrlRuleTableCache `json:"Cache,omitempty"`
	// Status (enable/disable) of caching URL Rule.
	AdminStatus SlbNewAcclCfgCacheUrlRuleTableAdminStatus `json:"AdminStatus,omitempty"`
	// By setting the value to delete(2), the entire row is deleted.
	Delete SlbNewAcclCfgCacheUrlRuleTableDelete `json:"Delete,omitempty"`
	// Copy rule to another index in the same rule-list. When read,
	// zero is returned.
	Copy string `json:"Copy,omitempty"`
}

func (p SlbNewAcclCfgCacheUrlRuleTableParams) iMABean() {}
