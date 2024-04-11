package beans

import (
	"fmt"
	"reflect"
)

// SlbCurCfgContRuleTable The table of Content based Services Rules.
// Note:This mib is not supported for VX instance of Virtualization.
type SlbCurCfgContRuleTable struct {
	// The number of the virtual server.
	SlbCurCfgContVirtServIndex int32
	// The service index. This has no external meaning
	SlbCurCfgContVirtServiceIndex int32
	// The content based service rule index.
	SlbCurCfgContRuleIndex int32
	Params                 *SlbCurCfgContRuleTableParams
}

func NewSlbCurCfgContRuleTableList() *SlbCurCfgContRuleTable {
	return &SlbCurCfgContRuleTable{}
}

func NewSlbCurCfgContRuleTable(
	slbCurCfgContVirtServIndex int32,
	slbCurCfgContVirtServiceIndex int32,
	slbCurCfgContRuleIndex int32,
	params *SlbCurCfgContRuleTableParams,
) *SlbCurCfgContRuleTable {
	return &SlbCurCfgContRuleTable{
		SlbCurCfgContVirtServIndex:    slbCurCfgContVirtServIndex,
		SlbCurCfgContVirtServiceIndex: slbCurCfgContVirtServiceIndex,
		SlbCurCfgContRuleIndex:        slbCurCfgContRuleIndex,
		Params:                        params,
	}
}

func (c *SlbCurCfgContRuleTable) Name() string {
	return "SlbCurCfgContRuleTable"
}

func (c *SlbCurCfgContRuleTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbCurCfgContRuleTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbCurCfgContRuleTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbCurCfgContVirtServIndex).IsZero() &&
		reflect.ValueOf(c.SlbCurCfgContVirtServiceIndex).IsZero() &&
		reflect.ValueOf(c.SlbCurCfgContRuleIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbCurCfgContVirtServIndex) + "/" + fmt.Sprint(c.SlbCurCfgContVirtServiceIndex) + "/" + fmt.Sprint(c.SlbCurCfgContRuleIndex)
}

type SlbCurCfgContRuleTableAction int32

const (
	SlbCurCfgContRuleTableAction_Group       SlbCurCfgContRuleTableAction = 1
	SlbCurCfgContRuleTableAction_Appredir    SlbCurCfgContRuleTableAction = 2
	SlbCurCfgContRuleTableAction_Discard     SlbCurCfgContRuleTableAction = 3
	SlbCurCfgContRuleTableAction_Goto        SlbCurCfgContRuleTableAction = 4
	SlbCurCfgContRuleTableAction_Unsupported SlbCurCfgContRuleTableAction = 2147483647
)

type SlbCurCfgContRuleTableState int32

const (
	SlbCurCfgContRuleTableState_Enabled     SlbCurCfgContRuleTableState = 1
	SlbCurCfgContRuleTableState_Disabled    SlbCurCfgContRuleTableState = 2
	SlbCurCfgContRuleTableState_Unsupported SlbCurCfgContRuleTableState = 2147483647
)

type SlbCurCfgContRuleTableParams struct {
	// The number of the virtual server.
	VirtServIndex int32 `json:"VirtServIndex,omitempty"`
	// The service index. This has no external meaning
	VirtServiceIndex int32 `json:"VirtServiceIndex,omitempty"`
	// The content based service rule index.
	Index int32 `json:"Index,omitempty"`
	// The Content Rule Name.
	Name string `json:"Name,omitempty"`
	// The Content Class for the rule.
	ContClass string `json:"ContClass,omitempty"`
	// Action type of the rule.
	Action SlbCurCfgContRuleTableAction `json:"Action,omitempty"`
	// The group number of real server.
	RealGrpNum uint32 `json:"RealGrpNum,omitempty"`
	// The group number of real server.
	GotoRuleNum uint32 `json:"GotoRuleNum,omitempty"`
	// Application redirection for the rule.
	Redirection string `json:"Redirection,omitempty"`
	// Enable or disable Content Based Services Rule .
	State SlbCurCfgContRuleTableState `json:"State,omitempty"`
}

func (p SlbCurCfgContRuleTableParams) iMABean() {}
