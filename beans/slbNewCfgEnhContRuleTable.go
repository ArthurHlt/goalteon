package beans

import (
	"fmt"
	"reflect"
)

// SlbNewCfgEnhContRuleTable The table of Content based Services Rules.
// Note:This mib is not supported for VX instance of Virtualization.
type SlbNewCfgEnhContRuleTable struct {
	// The number of the virtual server.
	SlbNewCfgEnhContVirtServIndex string
	// The service index. This has no external meaning
	SlbNewCfgEnhContVirtServiceIndex int32
	// The content based service index.
	SlbNewCfgEnhContRuleIndex int32
	Params                    *SlbNewCfgEnhContRuleTableParams
}

func NewSlbNewCfgEnhContRuleTableList() *SlbNewCfgEnhContRuleTable {
	return &SlbNewCfgEnhContRuleTable{}
}

func NewSlbNewCfgEnhContRuleTable(
	slbNewCfgEnhContVirtServIndex string,
	slbNewCfgEnhContVirtServiceIndex int32,
	slbNewCfgEnhContRuleIndex int32,
	params *SlbNewCfgEnhContRuleTableParams,
) *SlbNewCfgEnhContRuleTable {
	return &SlbNewCfgEnhContRuleTable{
		SlbNewCfgEnhContVirtServIndex:    slbNewCfgEnhContVirtServIndex,
		SlbNewCfgEnhContVirtServiceIndex: slbNewCfgEnhContVirtServiceIndex,
		SlbNewCfgEnhContRuleIndex:        slbNewCfgEnhContRuleIndex,
		Params:                           params,
	}
}

func (c *SlbNewCfgEnhContRuleTable) Name() string {
	return "SlbNewCfgEnhContRuleTable"
}

func (c *SlbNewCfgEnhContRuleTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbNewCfgEnhContRuleTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbNewCfgEnhContRuleTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbNewCfgEnhContVirtServIndex).IsZero() &&
		reflect.ValueOf(c.SlbNewCfgEnhContVirtServiceIndex).IsZero() &&
		reflect.ValueOf(c.SlbNewCfgEnhContRuleIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbNewCfgEnhContVirtServIndex) + "/" + fmt.Sprint(c.SlbNewCfgEnhContVirtServiceIndex) + "/" + fmt.Sprint(c.SlbNewCfgEnhContRuleIndex)
}

type SlbNewCfgEnhContRuleTableAction int32

const (
	SlbNewCfgEnhContRuleTableAction_Group       SlbNewCfgEnhContRuleTableAction = 1
	SlbNewCfgEnhContRuleTableAction_Appredir    SlbNewCfgEnhContRuleTableAction = 2
	SlbNewCfgEnhContRuleTableAction_Discard     SlbNewCfgEnhContRuleTableAction = 3
	SlbNewCfgEnhContRuleTableAction_Goto        SlbNewCfgEnhContRuleTableAction = 4
	SlbNewCfgEnhContRuleTableAction_Unsupported SlbNewCfgEnhContRuleTableAction = 2147483647
)

type SlbNewCfgEnhContRuleTableState int32

const (
	SlbNewCfgEnhContRuleTableState_Enabled     SlbNewCfgEnhContRuleTableState = 1
	SlbNewCfgEnhContRuleTableState_Disabled    SlbNewCfgEnhContRuleTableState = 2
	SlbNewCfgEnhContRuleTableState_Unsupported SlbNewCfgEnhContRuleTableState = 2147483647
)

type SlbNewCfgEnhContRuleTableDelete int32

const (
	SlbNewCfgEnhContRuleTableDelete_Other       SlbNewCfgEnhContRuleTableDelete = 1
	SlbNewCfgEnhContRuleTableDelete_Delete      SlbNewCfgEnhContRuleTableDelete = 2
	SlbNewCfgEnhContRuleTableDelete_Unsupported SlbNewCfgEnhContRuleTableDelete = 2147483647
)

type SlbNewCfgEnhContRuleTableSidebandProccessing int32

const (
	SlbNewCfgEnhContRuleTableSidebandProccessing_Inherit     SlbNewCfgEnhContRuleTableSidebandProccessing = 1
	SlbNewCfgEnhContRuleTableSidebandProccessing_Specific    SlbNewCfgEnhContRuleTableSidebandProccessing = 2
	SlbNewCfgEnhContRuleTableSidebandProccessing_Disabled    SlbNewCfgEnhContRuleTableSidebandProccessing = 3
	SlbNewCfgEnhContRuleTableSidebandProccessing_Unsupported SlbNewCfgEnhContRuleTableSidebandProccessing = 2147483647
)

type SlbNewCfgEnhContRuleTableParams struct {
	// The number of the virtual server.
	VirtServIndex string `json:"VirtServIndex,omitempty"`
	// The service index. This has no external meaning
	VirtServiceIndex int32 `json:"VirtServiceIndex,omitempty"`
	// The content based service index.
	Index int32 `json:"Index,omitempty"`
	// The Content Rule Name.
	Name string `json:"Name,omitempty"`
	// The Content Class for the rule.
	ContClass string `json:"ContClass,omitempty"`
	// Action type of the rule.
	Action SlbNewCfgEnhContRuleTableAction `json:"Action,omitempty"`
	// The group number of real server.
	RealGrpNum string `json:"RealGrpNum,omitempty"`
	// The group number of real server.
	GotoRuleNum uint32 `json:"GotoRuleNum,omitempty"`
	// Application redirection for the rule.
	Redirection string `json:"Redirection,omitempty"`
	// The Number to which current rule is to be copied.
	Copy uint32 `json:"Copy,omitempty"`
	// Enable or disable Content Based Services Rule .
	State SlbNewCfgEnhContRuleTableState `json:"State,omitempty"`
	// When set to the value of 2 (delete), the entire row is deleted.
	// When read, other(1) is returned. Setting the value to anything
	// other than 2(delete) has no effect on the state of the row.
	Delete SlbNewCfgEnhContRuleTableDelete `json:"Delete,omitempty"`
	// Fastview web application name associated with this rule.
	FastWa string `json:"FastWa,omitempty"`
	// Set SecurePath Policy for this rule.
	SecurePathPolicy string `json:"SecurePathPolicy,omitempty"`
	// Set Sideband Processing for this rule.
	SidebandProccessing SlbNewCfgEnhContRuleTableSidebandProccessing `json:"SidebandProccessing,omitempty"`
	// Set sideband policy for this rule.
	SidebandID string `json:"SidebandID,omitempty"`
}

func (p SlbNewCfgEnhContRuleTableParams) iMABean() {}
