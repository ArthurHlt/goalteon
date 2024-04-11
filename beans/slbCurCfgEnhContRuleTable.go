package beans

import (
	"fmt"
	"reflect"
)

// SlbCurCfgEnhContRuleTable The table of Content based Services Rules.
// Note:This mib is not supported for VX instance of Virtualization.
type SlbCurCfgEnhContRuleTable struct {
	// The number of the virtual server.
	SlbCurCfgEnhContVirtServIndex string
	// The service index. This has no external meaning
	SlbCurCfgEnhContVirtServiceIndex int32
	// The content based service rule index.
	SlbCurCfgEnhContRuleIndex int32
	Params                    *SlbCurCfgEnhContRuleTableParams
}

func NewSlbCurCfgEnhContRuleTableList() *SlbCurCfgEnhContRuleTable {
	return &SlbCurCfgEnhContRuleTable{}
}

func NewSlbCurCfgEnhContRuleTable(
	slbCurCfgEnhContVirtServIndex string,
	slbCurCfgEnhContVirtServiceIndex int32,
	slbCurCfgEnhContRuleIndex int32,
	params *SlbCurCfgEnhContRuleTableParams,
) *SlbCurCfgEnhContRuleTable {
	return &SlbCurCfgEnhContRuleTable{
		SlbCurCfgEnhContVirtServIndex:    slbCurCfgEnhContVirtServIndex,
		SlbCurCfgEnhContVirtServiceIndex: slbCurCfgEnhContVirtServiceIndex,
		SlbCurCfgEnhContRuleIndex:        slbCurCfgEnhContRuleIndex,
		Params:                           params,
	}
}

func (c *SlbCurCfgEnhContRuleTable) Name() string {
	return "SlbCurCfgEnhContRuleTable"
}

func (c *SlbCurCfgEnhContRuleTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbCurCfgEnhContRuleTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbCurCfgEnhContRuleTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbCurCfgEnhContVirtServIndex).IsZero() &&
		reflect.ValueOf(c.SlbCurCfgEnhContVirtServiceIndex).IsZero() &&
		reflect.ValueOf(c.SlbCurCfgEnhContRuleIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbCurCfgEnhContVirtServIndex) + "/" + fmt.Sprint(c.SlbCurCfgEnhContVirtServiceIndex) + "/" + fmt.Sprint(c.SlbCurCfgEnhContRuleIndex)
}

type SlbCurCfgEnhContRuleTableAction int32

const (
	SlbCurCfgEnhContRuleTableAction_Group       SlbCurCfgEnhContRuleTableAction = 1
	SlbCurCfgEnhContRuleTableAction_Appredir    SlbCurCfgEnhContRuleTableAction = 2
	SlbCurCfgEnhContRuleTableAction_Discard     SlbCurCfgEnhContRuleTableAction = 3
	SlbCurCfgEnhContRuleTableAction_Goto        SlbCurCfgEnhContRuleTableAction = 4
	SlbCurCfgEnhContRuleTableAction_Unsupported SlbCurCfgEnhContRuleTableAction = 2147483647
)

type SlbCurCfgEnhContRuleTableState int32

const (
	SlbCurCfgEnhContRuleTableState_Enabled     SlbCurCfgEnhContRuleTableState = 1
	SlbCurCfgEnhContRuleTableState_Disabled    SlbCurCfgEnhContRuleTableState = 2
	SlbCurCfgEnhContRuleTableState_Unsupported SlbCurCfgEnhContRuleTableState = 2147483647
)

type SlbCurCfgEnhContRuleTableSidebandProccessing int32

const (
	SlbCurCfgEnhContRuleTableSidebandProccessing_Inherit     SlbCurCfgEnhContRuleTableSidebandProccessing = 1
	SlbCurCfgEnhContRuleTableSidebandProccessing_Specific    SlbCurCfgEnhContRuleTableSidebandProccessing = 2
	SlbCurCfgEnhContRuleTableSidebandProccessing_Disabled    SlbCurCfgEnhContRuleTableSidebandProccessing = 3
	SlbCurCfgEnhContRuleTableSidebandProccessing_Unsupported SlbCurCfgEnhContRuleTableSidebandProccessing = 2147483647
)

type SlbCurCfgEnhContRuleTableParams struct {
	// The number of the virtual server.
	VirtServIndex string `json:"VirtServIndex,omitempty"`
	// The service index. This has no external meaning
	VirtServiceIndex int32 `json:"VirtServiceIndex,omitempty"`
	// The content based service rule index.
	Index int32 `json:"Index,omitempty"`
	// The Content Rule Name.
	Name string `json:"Name,omitempty"`
	// The Content Class for the rule.
	ContClass string `json:"ContClass,omitempty"`
	// Action type of the rule.
	Action SlbCurCfgEnhContRuleTableAction `json:"Action,omitempty"`
	// The group number of real server.
	RealGrpNum string `json:"RealGrpNum,omitempty"`
	// The group number of real server.
	GotoRuleNum uint32 `json:"GotoRuleNum,omitempty"`
	// Application redirection for the rule.
	Redirection string `json:"Redirection,omitempty"`
	// Enable or disable Content Based Services Rule .
	State SlbCurCfgEnhContRuleTableState `json:"State,omitempty"`
	// Fastview web application name associated with this rule.
	FastWa string `json:"FastWa,omitempty"`
	// Set SecurePath Policy for this rule.
	SecurePathPolicy string `json:"SecurePathPolicy,omitempty"`
	// Set Sideband Processing for this rule.
	SidebandProccessing SlbCurCfgEnhContRuleTableSidebandProccessing `json:"SidebandProccessing,omitempty"`
	// Set sideband policy for this rule.
	SidebandID string `json:"SidebandID,omitempty"`
}

func (p SlbCurCfgEnhContRuleTableParams) iMABean() {}
