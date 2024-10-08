package beans

import (
	"fmt"
	"reflect"
)

// SlbNewCfgContRuleTable The table of Content based Services Rules.
// Note:This mib is not supported for VX instance of Virtualization.
type SlbNewCfgContRuleTable struct {
	// The number of the virtual server.
	SlbNewCfgContVirtServIndex int32
	// The service index. This has no external meaning
	SlbNewCfgContVirtServiceIndex int32
	// The content based service index.
	SlbNewCfgContRuleIndex int32
	Params                 *SlbNewCfgContRuleTableParams
}

func NewSlbNewCfgContRuleTableList() *SlbNewCfgContRuleTable {
	return &SlbNewCfgContRuleTable{}
}

func NewSlbNewCfgContRuleTable(
	slbNewCfgContVirtServIndex int32,
	slbNewCfgContVirtServiceIndex int32,
	slbNewCfgContRuleIndex int32,
	params *SlbNewCfgContRuleTableParams,
) *SlbNewCfgContRuleTable {
	return &SlbNewCfgContRuleTable{
		SlbNewCfgContVirtServIndex:    slbNewCfgContVirtServIndex,
		SlbNewCfgContVirtServiceIndex: slbNewCfgContVirtServiceIndex,
		SlbNewCfgContRuleIndex:        slbNewCfgContRuleIndex,
		Params:                        params,
	}
}

func (c *SlbNewCfgContRuleTable) Name() string {
	return "SlbNewCfgContRuleTable"
}

func (c *SlbNewCfgContRuleTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbNewCfgContRuleTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbNewCfgContRuleTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbNewCfgContVirtServIndex).IsZero() &&
		reflect.ValueOf(c.SlbNewCfgContVirtServiceIndex).IsZero() &&
		reflect.ValueOf(c.SlbNewCfgContRuleIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbNewCfgContVirtServIndex) + "/" + fmt.Sprint(c.SlbNewCfgContVirtServiceIndex) + "/" + fmt.Sprint(c.SlbNewCfgContRuleIndex)
}

type SlbNewCfgContRuleTableAction int32

const (
	SlbNewCfgContRuleTableAction_Group       SlbNewCfgContRuleTableAction = 1
	SlbNewCfgContRuleTableAction_Appredir    SlbNewCfgContRuleTableAction = 2
	SlbNewCfgContRuleTableAction_Discard     SlbNewCfgContRuleTableAction = 3
	SlbNewCfgContRuleTableAction_Goto        SlbNewCfgContRuleTableAction = 4
	SlbNewCfgContRuleTableAction_Unsupported SlbNewCfgContRuleTableAction = 2147483647
)

type SlbNewCfgContRuleTableState int32

const (
	SlbNewCfgContRuleTableState_Enabled     SlbNewCfgContRuleTableState = 1
	SlbNewCfgContRuleTableState_Disabled    SlbNewCfgContRuleTableState = 2
	SlbNewCfgContRuleTableState_Unsupported SlbNewCfgContRuleTableState = 2147483647
)

type SlbNewCfgContRuleTableDelete int32

const (
	SlbNewCfgContRuleTableDelete_Other       SlbNewCfgContRuleTableDelete = 1
	SlbNewCfgContRuleTableDelete_Delete      SlbNewCfgContRuleTableDelete = 2
	SlbNewCfgContRuleTableDelete_Unsupported SlbNewCfgContRuleTableDelete = 2147483647
)

type SlbNewCfgContRuleTableParams struct {
	// The number of the virtual server.
	VirtServIndex int32 `json:"VirtServIndex,omitempty"`
	// The service index. This has no external meaning
	VirtServiceIndex int32 `json:"VirtServiceIndex,omitempty"`
	// The content based service index.
	Index int32 `json:"Index,omitempty"`
	// The Content Rule Name.
	Name string `json:"Name,omitempty"`
	// The Content Class for the rule.
	ContClass string `json:"ContClass,omitempty"`
	// Action type of the rule.
	Action SlbNewCfgContRuleTableAction `json:"Action,omitempty"`
	// The group number of real server.
	RealGrpNum uint32 `json:"RealGrpNum,omitempty"`
	// The group number of real server.
	GotoRuleNum uint32 `json:"GotoRuleNum,omitempty"`
	// Application redirection for the rule.
	Redirection string `json:"Redirection,omitempty"`
	// The Number to which current rule is to be copied.
	Copy DisplayString `json:"Copy,omitempty"`
	// Enable or disable Content Based Services Rule .
	State SlbNewCfgContRuleTableState `json:"State,omitempty"`
	// When set to the value of 2 (delete), the entire row is deleted.
	// When read, other(1) is returned. Setting the value to anything
	// other than 2(delete) has no effect on the state of the row.
	Delete SlbNewCfgContRuleTableDelete `json:"Delete,omitempty"`
}

func (p SlbNewCfgContRuleTableParams) iMABean() {}
