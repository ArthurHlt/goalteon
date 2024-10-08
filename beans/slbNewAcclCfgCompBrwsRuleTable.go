package beans

import (
	"fmt"
	"reflect"
)

// SlbNewAcclCfgCompBrwsRuleTable The table for configuring compression browser Rules.
// Note:This mib is not supported for VX instance of Virtualization.
type SlbNewAcclCfgCompBrwsRuleTable struct {
	// The compression browser list (key id) as an index.
	SlbNewAcclCfgCompBrwsRuleListIdIndex string
	// The compression browser Rule number as an index.
	SlbNewAcclCfgCompBrwsRuleIndex int32
	Params                         *SlbNewAcclCfgCompBrwsRuleTableParams
}

func NewSlbNewAcclCfgCompBrwsRuleTableList() *SlbNewAcclCfgCompBrwsRuleTable {
	return &SlbNewAcclCfgCompBrwsRuleTable{}
}

func NewSlbNewAcclCfgCompBrwsRuleTable(
	slbNewAcclCfgCompBrwsRuleListIdIndex string,
	slbNewAcclCfgCompBrwsRuleIndex int32,
	params *SlbNewAcclCfgCompBrwsRuleTableParams,
) *SlbNewAcclCfgCompBrwsRuleTable {
	return &SlbNewAcclCfgCompBrwsRuleTable{
		SlbNewAcclCfgCompBrwsRuleListIdIndex: slbNewAcclCfgCompBrwsRuleListIdIndex,
		SlbNewAcclCfgCompBrwsRuleIndex:       slbNewAcclCfgCompBrwsRuleIndex,
		Params:                               params,
	}
}

func (c *SlbNewAcclCfgCompBrwsRuleTable) Name() string {
	return "SlbNewAcclCfgCompBrwsRuleTable"
}

func (c *SlbNewAcclCfgCompBrwsRuleTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbNewAcclCfgCompBrwsRuleTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbNewAcclCfgCompBrwsRuleTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbNewAcclCfgCompBrwsRuleListIdIndex).IsZero() &&
		reflect.ValueOf(c.SlbNewAcclCfgCompBrwsRuleIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbNewAcclCfgCompBrwsRuleListIdIndex) + "/" + fmt.Sprint(c.SlbNewAcclCfgCompBrwsRuleIndex)
}

type SlbNewAcclCfgCompBrwsRuleTableAgentM int32

const (
	SlbNewAcclCfgCompBrwsRuleTableAgentM_Any         SlbNewAcclCfgCompBrwsRuleTableAgentM = 1
	SlbNewAcclCfgCompBrwsRuleTableAgentM_Text        SlbNewAcclCfgCompBrwsRuleTableAgentM = 2
	SlbNewAcclCfgCompBrwsRuleTableAgentM_Regex       SlbNewAcclCfgCompBrwsRuleTableAgentM = 3
	SlbNewAcclCfgCompBrwsRuleTableAgentM_Unsupported SlbNewAcclCfgCompBrwsRuleTableAgentM = 2147483647
)

type SlbNewAcclCfgCompBrwsRuleTableContentM int32

const (
	SlbNewAcclCfgCompBrwsRuleTableContentM_Any         SlbNewAcclCfgCompBrwsRuleTableContentM = 1
	SlbNewAcclCfgCompBrwsRuleTableContentM_Text        SlbNewAcclCfgCompBrwsRuleTableContentM = 2
	SlbNewAcclCfgCompBrwsRuleTableContentM_Regex       SlbNewAcclCfgCompBrwsRuleTableContentM = 3
	SlbNewAcclCfgCompBrwsRuleTableContentM_Unsupported SlbNewAcclCfgCompBrwsRuleTableContentM = 2147483647
)

type SlbNewAcclCfgCompBrwsRuleTableCompress int32

const (
	SlbNewAcclCfgCompBrwsRuleTableCompress_Enabled     SlbNewAcclCfgCompBrwsRuleTableCompress = 1
	SlbNewAcclCfgCompBrwsRuleTableCompress_Disabled    SlbNewAcclCfgCompBrwsRuleTableCompress = 2
	SlbNewAcclCfgCompBrwsRuleTableCompress_Unsupported SlbNewAcclCfgCompBrwsRuleTableCompress = 2147483647
)

type SlbNewAcclCfgCompBrwsRuleTableAdminStatus int32

const (
	SlbNewAcclCfgCompBrwsRuleTableAdminStatus_Enabled     SlbNewAcclCfgCompBrwsRuleTableAdminStatus = 1
	SlbNewAcclCfgCompBrwsRuleTableAdminStatus_Disabled    SlbNewAcclCfgCompBrwsRuleTableAdminStatus = 2
	SlbNewAcclCfgCompBrwsRuleTableAdminStatus_Unsupported SlbNewAcclCfgCompBrwsRuleTableAdminStatus = 2147483647
)

type SlbNewAcclCfgCompBrwsRuleTableDelete int32

const (
	SlbNewAcclCfgCompBrwsRuleTableDelete_Other       SlbNewAcclCfgCompBrwsRuleTableDelete = 1
	SlbNewAcclCfgCompBrwsRuleTableDelete_Delete      SlbNewAcclCfgCompBrwsRuleTableDelete = 2
	SlbNewAcclCfgCompBrwsRuleTableDelete_Unsupported SlbNewAcclCfgCompBrwsRuleTableDelete = 2147483647
)

type SlbNewAcclCfgCompBrwsRuleTableParams struct {
	// The compression browser list (key id) as an index.
	ListIdIndex string `json:"ListIdIndex,omitempty"`
	// The compression browser Rule number as an index.
	Index int32 `json:"Index,omitempty"`
	// The compression browser Rule name.
	Name string `json:"Name,omitempty"`
	// Defines whether user agent field should be evaluated as String, Regex or match, any (default any).
	AgentM SlbNewAcclCfgCompBrwsRuleTableAgentM `json:"AgentM,omitempty"`
	// Defines the User-Agent string for which this rule applies.
	Agent string `json:"Agent,omitempty"`
	// Defines whether content type should be evaluated as String, Regex or match, any (default any).
	ContentM SlbNewAcclCfgCompBrwsRuleTableContentM `json:"ContentM,omitempty"`
	// Defines the Content-type string for which this rule applies.
	Content string `json:"Content,omitempty"`
	// Enable/Disable the compression. Defines if the matched response should be compressed or not
	Compress SlbNewAcclCfgCompBrwsRuleTableCompress `json:"Compress,omitempty"`
	// Status (enable/disable) of compression browser Rule.
	AdminStatus SlbNewAcclCfgCompBrwsRuleTableAdminStatus `json:"AdminStatus,omitempty"`
	// By setting the value to delete(2), the entire row is deleted.
	Delete SlbNewAcclCfgCompBrwsRuleTableDelete `json:"Delete,omitempty"`
	// Copy rule to another index in the same rule-list.
	Copy DisplayString `json:"Copy,omitempty"`
}

func (p SlbNewAcclCfgCompBrwsRuleTableParams) iMABean() {}
