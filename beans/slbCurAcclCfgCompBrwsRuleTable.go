package beans

import (
	"fmt"
	"reflect"
)

// SlbCurAcclCfgCompBrwsRuleTable The table for configuring compression browser Rules.
// Note:This mib is not supported for VX instance of Virtualization.
type SlbCurAcclCfgCompBrwsRuleTable struct {
	// The compression browser list (key id) as an index.
	SlbCurAcclCfgCompBrwsRuleListIdIndex string
	// The compression browser Rule number as an index.
	SlbCurAcclCfgCompBrwsRuleIndex int32
	Params                         *SlbCurAcclCfgCompBrwsRuleTableParams
}

func NewSlbCurAcclCfgCompBrwsRuleTableList() *SlbCurAcclCfgCompBrwsRuleTable {
	return &SlbCurAcclCfgCompBrwsRuleTable{}
}

func NewSlbCurAcclCfgCompBrwsRuleTable(
	slbCurAcclCfgCompBrwsRuleListIdIndex string,
	slbCurAcclCfgCompBrwsRuleIndex int32,
	params *SlbCurAcclCfgCompBrwsRuleTableParams,
) *SlbCurAcclCfgCompBrwsRuleTable {
	return &SlbCurAcclCfgCompBrwsRuleTable{
		SlbCurAcclCfgCompBrwsRuleListIdIndex: slbCurAcclCfgCompBrwsRuleListIdIndex,
		SlbCurAcclCfgCompBrwsRuleIndex:       slbCurAcclCfgCompBrwsRuleIndex,
		Params:                               params,
	}
}

func (c *SlbCurAcclCfgCompBrwsRuleTable) Name() string {
	return "SlbCurAcclCfgCompBrwsRuleTable"
}

func (c *SlbCurAcclCfgCompBrwsRuleTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbCurAcclCfgCompBrwsRuleTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbCurAcclCfgCompBrwsRuleTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbCurAcclCfgCompBrwsRuleListIdIndex).IsZero() &&
		reflect.ValueOf(c.SlbCurAcclCfgCompBrwsRuleIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbCurAcclCfgCompBrwsRuleListIdIndex) + "/" + fmt.Sprint(c.SlbCurAcclCfgCompBrwsRuleIndex)
}

type SlbCurAcclCfgCompBrwsRuleTableAgentM int32

const (
	SlbCurAcclCfgCompBrwsRuleTableAgentM_Any         SlbCurAcclCfgCompBrwsRuleTableAgentM = 1
	SlbCurAcclCfgCompBrwsRuleTableAgentM_Text        SlbCurAcclCfgCompBrwsRuleTableAgentM = 2
	SlbCurAcclCfgCompBrwsRuleTableAgentM_Regex       SlbCurAcclCfgCompBrwsRuleTableAgentM = 3
	SlbCurAcclCfgCompBrwsRuleTableAgentM_Unsupported SlbCurAcclCfgCompBrwsRuleTableAgentM = 2147483647
)

type SlbCurAcclCfgCompBrwsRuleTableContentM int32

const (
	SlbCurAcclCfgCompBrwsRuleTableContentM_Any         SlbCurAcclCfgCompBrwsRuleTableContentM = 1
	SlbCurAcclCfgCompBrwsRuleTableContentM_Text        SlbCurAcclCfgCompBrwsRuleTableContentM = 2
	SlbCurAcclCfgCompBrwsRuleTableContentM_Regex       SlbCurAcclCfgCompBrwsRuleTableContentM = 3
	SlbCurAcclCfgCompBrwsRuleTableContentM_Unsupported SlbCurAcclCfgCompBrwsRuleTableContentM = 2147483647
)

type SlbCurAcclCfgCompBrwsRuleTableCompress int32

const (
	SlbCurAcclCfgCompBrwsRuleTableCompress_Enabled     SlbCurAcclCfgCompBrwsRuleTableCompress = 1
	SlbCurAcclCfgCompBrwsRuleTableCompress_Disabled    SlbCurAcclCfgCompBrwsRuleTableCompress = 2
	SlbCurAcclCfgCompBrwsRuleTableCompress_Unsupported SlbCurAcclCfgCompBrwsRuleTableCompress = 2147483647
)

type SlbCurAcclCfgCompBrwsRuleTableAdminStatus int32

const (
	SlbCurAcclCfgCompBrwsRuleTableAdminStatus_Enabled     SlbCurAcclCfgCompBrwsRuleTableAdminStatus = 1
	SlbCurAcclCfgCompBrwsRuleTableAdminStatus_Disabled    SlbCurAcclCfgCompBrwsRuleTableAdminStatus = 2
	SlbCurAcclCfgCompBrwsRuleTableAdminStatus_Unsupported SlbCurAcclCfgCompBrwsRuleTableAdminStatus = 2147483647
)

type SlbCurAcclCfgCompBrwsRuleTableParams struct {
	// The compression browser list (key id) as an index.
	ListIdIndex string `json:"ListIdIndex,omitempty"`
	// The compression browser Rule number as an index.
	Index int32 `json:"Index,omitempty"`
	// The compression browser Rule name.
	Name string `json:"Name,omitempty"`
	// Defines whether user agent field should be evaluated as String, Regex or match, any (default any).
	AgentM SlbCurAcclCfgCompBrwsRuleTableAgentM `json:"AgentM,omitempty"`
	// Defines the User-Agent string for which this rule applies.
	Agent string `json:"Agent,omitempty"`
	// Defines whether content type should be evaluated as String, Regex or match, any (default any).
	ContentM SlbCurAcclCfgCompBrwsRuleTableContentM `json:"ContentM,omitempty"`
	// Defines the Content-type string for which this rule applies.
	Content string `json:"Content,omitempty"`
	// Enable/Disable the compression. Defines if the matched response should be compressed or not
	Compress SlbCurAcclCfgCompBrwsRuleTableCompress `json:"Compress,omitempty"`
	// Status (enable/disable) of compression browser Rule.
	AdminStatus SlbCurAcclCfgCompBrwsRuleTableAdminStatus `json:"AdminStatus,omitempty"`
}

func (p SlbCurAcclCfgCompBrwsRuleTableParams) iMABean() {}
