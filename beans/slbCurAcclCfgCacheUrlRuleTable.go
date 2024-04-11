package beans

import (
	"fmt"
	"reflect"
)

// SlbCurAcclCfgCacheUrlRuleTable The table for configuring caching URL Rules.
// Note:This mib is not supported for VX instance of Virtualization.
type SlbCurAcclCfgCacheUrlRuleTable struct {
	// The caching policy URL LIST (key id) as an index.
	SlbCurAcclCfgCacheUrlRuleListIdIndex string
	// The caching URL Rule number as an index.
	SlbCurAcclCfgCacheUrlRuleIndex int32
	Params                         *SlbCurAcclCfgCacheUrlRuleTableParams
}

func NewSlbCurAcclCfgCacheUrlRuleTableList() *SlbCurAcclCfgCacheUrlRuleTable {
	return &SlbCurAcclCfgCacheUrlRuleTable{}
}

func NewSlbCurAcclCfgCacheUrlRuleTable(
	slbCurAcclCfgCacheUrlRuleListIdIndex string,
	slbCurAcclCfgCacheUrlRuleIndex int32,
	params *SlbCurAcclCfgCacheUrlRuleTableParams,
) *SlbCurAcclCfgCacheUrlRuleTable {
	return &SlbCurAcclCfgCacheUrlRuleTable{
		SlbCurAcclCfgCacheUrlRuleListIdIndex: slbCurAcclCfgCacheUrlRuleListIdIndex,
		SlbCurAcclCfgCacheUrlRuleIndex:       slbCurAcclCfgCacheUrlRuleIndex,
		Params:                               params,
	}
}

func (c *SlbCurAcclCfgCacheUrlRuleTable) Name() string {
	return "SlbCurAcclCfgCacheUrlRuleTable"
}

func (c *SlbCurAcclCfgCacheUrlRuleTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbCurAcclCfgCacheUrlRuleTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbCurAcclCfgCacheUrlRuleTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbCurAcclCfgCacheUrlRuleListIdIndex).IsZero() &&
		reflect.ValueOf(c.SlbCurAcclCfgCacheUrlRuleIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbCurAcclCfgCacheUrlRuleListIdIndex) + "/" + fmt.Sprint(c.SlbCurAcclCfgCacheUrlRuleIndex)
}

type SlbCurAcclCfgCacheUrlRuleTableDomainM int32

const (
	SlbCurAcclCfgCacheUrlRuleTableDomainM_Any         SlbCurAcclCfgCacheUrlRuleTableDomainM = 1
	SlbCurAcclCfgCacheUrlRuleTableDomainM_Text        SlbCurAcclCfgCacheUrlRuleTableDomainM = 2
	SlbCurAcclCfgCacheUrlRuleTableDomainM_Regex       SlbCurAcclCfgCacheUrlRuleTableDomainM = 3
	SlbCurAcclCfgCacheUrlRuleTableDomainM_Unsupported SlbCurAcclCfgCacheUrlRuleTableDomainM = 2147483647
)

type SlbCurAcclCfgCacheUrlRuleTableURLm int32

const (
	SlbCurAcclCfgCacheUrlRuleTableURLm_Any         SlbCurAcclCfgCacheUrlRuleTableURLm = 1
	SlbCurAcclCfgCacheUrlRuleTableURLm_Text        SlbCurAcclCfgCacheUrlRuleTableURLm = 2
	SlbCurAcclCfgCacheUrlRuleTableURLm_Regex       SlbCurAcclCfgCacheUrlRuleTableURLm = 3
	SlbCurAcclCfgCacheUrlRuleTableURLm_Unsupported SlbCurAcclCfgCacheUrlRuleTableURLm = 2147483647
)

type SlbCurAcclCfgCacheUrlRuleTableCache int32

const (
	SlbCurAcclCfgCacheUrlRuleTableCache_Enabled     SlbCurAcclCfgCacheUrlRuleTableCache = 1
	SlbCurAcclCfgCacheUrlRuleTableCache_Disabled    SlbCurAcclCfgCacheUrlRuleTableCache = 2
	SlbCurAcclCfgCacheUrlRuleTableCache_Unsupported SlbCurAcclCfgCacheUrlRuleTableCache = 2147483647
)

type SlbCurAcclCfgCacheUrlRuleTableAdminStatus int32

const (
	SlbCurAcclCfgCacheUrlRuleTableAdminStatus_Enabled     SlbCurAcclCfgCacheUrlRuleTableAdminStatus = 1
	SlbCurAcclCfgCacheUrlRuleTableAdminStatus_Disabled    SlbCurAcclCfgCacheUrlRuleTableAdminStatus = 2
	SlbCurAcclCfgCacheUrlRuleTableAdminStatus_Unsupported SlbCurAcclCfgCacheUrlRuleTableAdminStatus = 2147483647
)

type SlbCurAcclCfgCacheUrlRuleTableParams struct {
	// The caching policy URL LIST (key id) as an index.
	ListIdIndex string `json:"ListIdIndex,omitempty"`
	// The caching URL Rule number as an index.
	Index int32 `json:"Index,omitempty"`
	// The caching URL Rule name.
	Name string `json:"Name,omitempty"`
	// Defines whether Domain matching should be evaluated as String, Regex or match, any (default any).
	DomainM SlbCurAcclCfgCacheUrlRuleTableDomainM `json:"DomainM,omitempty"`
	// Defines the Virtual Host for which this rule applies.
	Domain string `json:"Domain,omitempty"`
	// Defines whether URL matching should be evaluated as String, Regex or match, any (default any).
	URLm SlbCurAcclCfgCacheUrlRuleTableURLm `json:"URLm,omitempty"`
	// Defines URL of specific object (file/folder) to be matched by this rule.
	URL string `json:"URL,omitempty"`
	// The Maximum time for the cache object to served from cache (in seconds).
	Expire int32 `json:"Expire,omitempty"`
	// Defines if the matched response should be cached or not.
	Cache SlbCurAcclCfgCacheUrlRuleTableCache `json:"Cache,omitempty"`
	// Status (enable/disable) of caching URL Rule.
	AdminStatus SlbCurAcclCfgCacheUrlRuleTableAdminStatus `json:"AdminStatus,omitempty"`
}

func (p SlbCurAcclCfgCacheUrlRuleTableParams) iMABean() {}
