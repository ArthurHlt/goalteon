package beans

import (
	"fmt"
	"reflect"
)

// SlbCurAcclCfgCompUrlRuleTable The table for configuring compression URL Rules.
// Note:This mib is not supported for VX instance of Virtualization.
type SlbCurAcclCfgCompUrlRuleTable struct {
	// The compression URL LIST (key id) as an index.
	SlbCurAcclCfgCompUrlRuleListIdIndex string
	// The compression URL Rule number as an index.
	SlbCurAcclCfgCompUrlRuleIndex int32
	Params                        *SlbCurAcclCfgCompUrlRuleTableParams
}

func NewSlbCurAcclCfgCompUrlRuleTableList() *SlbCurAcclCfgCompUrlRuleTable {
	return &SlbCurAcclCfgCompUrlRuleTable{}
}

func NewSlbCurAcclCfgCompUrlRuleTable(
	slbCurAcclCfgCompUrlRuleListIdIndex string,
	slbCurAcclCfgCompUrlRuleIndex int32,
	params *SlbCurAcclCfgCompUrlRuleTableParams,
) *SlbCurAcclCfgCompUrlRuleTable {
	return &SlbCurAcclCfgCompUrlRuleTable{
		SlbCurAcclCfgCompUrlRuleListIdIndex: slbCurAcclCfgCompUrlRuleListIdIndex,
		SlbCurAcclCfgCompUrlRuleIndex:       slbCurAcclCfgCompUrlRuleIndex,
		Params:                              params,
	}
}

func (c *SlbCurAcclCfgCompUrlRuleTable) Name() string {
	return "SlbCurAcclCfgCompUrlRuleTable"
}

func (c *SlbCurAcclCfgCompUrlRuleTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbCurAcclCfgCompUrlRuleTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbCurAcclCfgCompUrlRuleTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbCurAcclCfgCompUrlRuleListIdIndex).IsZero() &&
		reflect.ValueOf(c.SlbCurAcclCfgCompUrlRuleIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbCurAcclCfgCompUrlRuleListIdIndex) + "/" + fmt.Sprint(c.SlbCurAcclCfgCompUrlRuleIndex)
}

type SlbCurAcclCfgCompUrlRuleTableDomainM int32

const (
	SlbCurAcclCfgCompUrlRuleTableDomainM_Any         SlbCurAcclCfgCompUrlRuleTableDomainM = 1
	SlbCurAcclCfgCompUrlRuleTableDomainM_Text        SlbCurAcclCfgCompUrlRuleTableDomainM = 2
	SlbCurAcclCfgCompUrlRuleTableDomainM_Regex       SlbCurAcclCfgCompUrlRuleTableDomainM = 3
	SlbCurAcclCfgCompUrlRuleTableDomainM_Unsupported SlbCurAcclCfgCompUrlRuleTableDomainM = 2147483647
)

type SlbCurAcclCfgCompUrlRuleTableURLm int32

const (
	SlbCurAcclCfgCompUrlRuleTableURLm_Any         SlbCurAcclCfgCompUrlRuleTableURLm = 1
	SlbCurAcclCfgCompUrlRuleTableURLm_Text        SlbCurAcclCfgCompUrlRuleTableURLm = 2
	SlbCurAcclCfgCompUrlRuleTableURLm_Regex       SlbCurAcclCfgCompUrlRuleTableURLm = 3
	SlbCurAcclCfgCompUrlRuleTableURLm_Unsupported SlbCurAcclCfgCompUrlRuleTableURLm = 2147483647
)

type SlbCurAcclCfgCompUrlRuleTableCompress int32

const (
	SlbCurAcclCfgCompUrlRuleTableCompress_Enabled     SlbCurAcclCfgCompUrlRuleTableCompress = 1
	SlbCurAcclCfgCompUrlRuleTableCompress_Disabled    SlbCurAcclCfgCompUrlRuleTableCompress = 2
	SlbCurAcclCfgCompUrlRuleTableCompress_Unsupported SlbCurAcclCfgCompUrlRuleTableCompress = 2147483647
)

type SlbCurAcclCfgCompUrlRuleTableAdminStatus int32

const (
	SlbCurAcclCfgCompUrlRuleTableAdminStatus_Enabled     SlbCurAcclCfgCompUrlRuleTableAdminStatus = 1
	SlbCurAcclCfgCompUrlRuleTableAdminStatus_Disabled    SlbCurAcclCfgCompUrlRuleTableAdminStatus = 2
	SlbCurAcclCfgCompUrlRuleTableAdminStatus_Unsupported SlbCurAcclCfgCompUrlRuleTableAdminStatus = 2147483647
)

type SlbCurAcclCfgCompUrlRuleTableParams struct {
	// The compression URL LIST (key id) as an index.
	ListIdIndex string `json:"ListIdIndex,omitempty"`
	// The compression URL Rule number as an index.
	Index int32 `json:"Index,omitempty"`
	// The compression URL Rule name.
	Name string `json:"Name,omitempty"`
	// Defines whether Domain matching should be evaluated as String, Regex or match, any (default any).
	DomainM SlbCurAcclCfgCompUrlRuleTableDomainM `json:"DomainM,omitempty"`
	// Defines the Virtual Host for which this rule applies.
	Domain string `json:"Domain,omitempty"`
	// Defines whether URL matching should be evaluated as String, Regex or match, any (default any).
	URLm SlbCurAcclCfgCompUrlRuleTableURLm `json:"URLm,omitempty"`
	// Defines URL of specific object (file/folder) to be matched by this rule.
	URL string `json:"URL,omitempty"`
	// Enable/Disable the compression.
	Compress SlbCurAcclCfgCompUrlRuleTableCompress `json:"Compress,omitempty"`
	// Status (enable/disable) of compression URL Rule.
	AdminStatus SlbCurAcclCfgCompUrlRuleTableAdminStatus `json:"AdminStatus,omitempty"`
}

func (p SlbCurAcclCfgCompUrlRuleTableParams) iMABean() {}
