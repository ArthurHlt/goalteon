package beans

import (
	"fmt"
	"reflect"
)

// SlbNewAcclCfgCompUrlRuleTable The table for configuring compression URL Rules.
// Note:This mib is not supported for VX instance of Virtualization.
type SlbNewAcclCfgCompUrlRuleTable struct {
	// The compression URL LIST (key id) as an index.
	SlbNewAcclCfgCompUrlRuleListIdIndex string
	// The compression URL Rule number as an index.
	SlbNewAcclCfgCompUrlRuleIndex int32
	Params                        *SlbNewAcclCfgCompUrlRuleTableParams
}

func NewSlbNewAcclCfgCompUrlRuleTableList() *SlbNewAcclCfgCompUrlRuleTable {
	return &SlbNewAcclCfgCompUrlRuleTable{}
}

func NewSlbNewAcclCfgCompUrlRuleTable(
	slbNewAcclCfgCompUrlRuleListIdIndex string,
	slbNewAcclCfgCompUrlRuleIndex int32,
	params *SlbNewAcclCfgCompUrlRuleTableParams,
) *SlbNewAcclCfgCompUrlRuleTable {
	return &SlbNewAcclCfgCompUrlRuleTable{
		SlbNewAcclCfgCompUrlRuleListIdIndex: slbNewAcclCfgCompUrlRuleListIdIndex,
		SlbNewAcclCfgCompUrlRuleIndex:       slbNewAcclCfgCompUrlRuleIndex,
		Params:                              params,
	}
}

func (c *SlbNewAcclCfgCompUrlRuleTable) Name() string {
	return "SlbNewAcclCfgCompUrlRuleTable"
}

func (c *SlbNewAcclCfgCompUrlRuleTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbNewAcclCfgCompUrlRuleTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbNewAcclCfgCompUrlRuleTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbNewAcclCfgCompUrlRuleListIdIndex).IsZero() &&
		reflect.ValueOf(c.SlbNewAcclCfgCompUrlRuleIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbNewAcclCfgCompUrlRuleListIdIndex) + "/" + fmt.Sprint(c.SlbNewAcclCfgCompUrlRuleIndex)
}

type SlbNewAcclCfgCompUrlRuleTableDomainM int32

const (
	SlbNewAcclCfgCompUrlRuleTableDomainM_Any         SlbNewAcclCfgCompUrlRuleTableDomainM = 1
	SlbNewAcclCfgCompUrlRuleTableDomainM_Text        SlbNewAcclCfgCompUrlRuleTableDomainM = 2
	SlbNewAcclCfgCompUrlRuleTableDomainM_Regex       SlbNewAcclCfgCompUrlRuleTableDomainM = 3
	SlbNewAcclCfgCompUrlRuleTableDomainM_Unsupported SlbNewAcclCfgCompUrlRuleTableDomainM = 2147483647
)

type SlbNewAcclCfgCompUrlRuleTableURLm int32

const (
	SlbNewAcclCfgCompUrlRuleTableURLm_Any         SlbNewAcclCfgCompUrlRuleTableURLm = 1
	SlbNewAcclCfgCompUrlRuleTableURLm_Text        SlbNewAcclCfgCompUrlRuleTableURLm = 2
	SlbNewAcclCfgCompUrlRuleTableURLm_Regex       SlbNewAcclCfgCompUrlRuleTableURLm = 3
	SlbNewAcclCfgCompUrlRuleTableURLm_Unsupported SlbNewAcclCfgCompUrlRuleTableURLm = 2147483647
)

type SlbNewAcclCfgCompUrlRuleTableCompress int32

const (
	SlbNewAcclCfgCompUrlRuleTableCompress_Enabled     SlbNewAcclCfgCompUrlRuleTableCompress = 1
	SlbNewAcclCfgCompUrlRuleTableCompress_Disabled    SlbNewAcclCfgCompUrlRuleTableCompress = 2
	SlbNewAcclCfgCompUrlRuleTableCompress_Unsupported SlbNewAcclCfgCompUrlRuleTableCompress = 2147483647
)

type SlbNewAcclCfgCompUrlRuleTableAdminStatus int32

const (
	SlbNewAcclCfgCompUrlRuleTableAdminStatus_Enabled     SlbNewAcclCfgCompUrlRuleTableAdminStatus = 1
	SlbNewAcclCfgCompUrlRuleTableAdminStatus_Disabled    SlbNewAcclCfgCompUrlRuleTableAdminStatus = 2
	SlbNewAcclCfgCompUrlRuleTableAdminStatus_Unsupported SlbNewAcclCfgCompUrlRuleTableAdminStatus = 2147483647
)

type SlbNewAcclCfgCompUrlRuleTableDelete int32

const (
	SlbNewAcclCfgCompUrlRuleTableDelete_Other       SlbNewAcclCfgCompUrlRuleTableDelete = 1
	SlbNewAcclCfgCompUrlRuleTableDelete_Delete      SlbNewAcclCfgCompUrlRuleTableDelete = 2
	SlbNewAcclCfgCompUrlRuleTableDelete_Unsupported SlbNewAcclCfgCompUrlRuleTableDelete = 2147483647
)

type SlbNewAcclCfgCompUrlRuleTableParams struct {
	// The compression URL LIST (key id) as an index.
	ListIdIndex string `json:"ListIdIndex,omitempty"`
	// The compression URL Rule number as an index.
	Index int32 `json:"Index,omitempty"`
	// The compression URL Rule name.
	Name string `json:"Name,omitempty"`
	// Defines whether Domain matching should be evaluated as String, Regex or match, any (default any).
	DomainM SlbNewAcclCfgCompUrlRuleTableDomainM `json:"DomainM,omitempty"`
	// Defines the Virtual Host for which this rule applies.
	Domain string `json:"Domain,omitempty"`
	// Defines whether URL matching should be evaluated as String, Regex or match, any (default any).
	URLm SlbNewAcclCfgCompUrlRuleTableURLm `json:"URLm,omitempty"`
	// Defines URL of specific object (file/folder) to be matched by this rule.
	URL string `json:"URL,omitempty"`
	// Enable/Disable the compression.
	Compress SlbNewAcclCfgCompUrlRuleTableCompress `json:"Compress,omitempty"`
	// Status (enable/disable) of compression URL Rule.
	AdminStatus SlbNewAcclCfgCompUrlRuleTableAdminStatus `json:"AdminStatus,omitempty"`
	// By setting the value to delete(2), the entire row is deleted.
	Delete SlbNewAcclCfgCompUrlRuleTableDelete `json:"Delete,omitempty"`
	// Copy rule to another index in the same rule-list.
	Copy DisplayString `json:"Copy,omitempty"`
}

func (p SlbNewAcclCfgCompUrlRuleTableParams) iMABean() {}
