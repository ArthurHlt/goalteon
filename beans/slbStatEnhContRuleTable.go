package beans

import (
	"fmt"
	"reflect"
)

// SlbStatEnhContRuleTable The statistics table of instances that the content rule matched.
// Note:This mib is not supported for VX instance of virtualization.
type SlbStatEnhContRuleTable struct {
	// The virtual server id.
	SlbStatEnhContVirtServIndex string
	// Virtual service index.
	SlbStatEnhContVirtServiceIndex int32
	// The content rule index.
	SlbStatEnhContRuleIndex int32
	Params                  *SlbStatEnhContRuleTableParams
}

func NewSlbStatEnhContRuleTableList() *SlbStatEnhContRuleTable {
	return &SlbStatEnhContRuleTable{}
}

func NewSlbStatEnhContRuleTable(
	slbStatEnhContVirtServIndex string,
	slbStatEnhContVirtServiceIndex int32,
	slbStatEnhContRuleIndex int32,
	params *SlbStatEnhContRuleTableParams,
) *SlbStatEnhContRuleTable {
	return &SlbStatEnhContRuleTable{
		SlbStatEnhContVirtServIndex:    slbStatEnhContVirtServIndex,
		SlbStatEnhContVirtServiceIndex: slbStatEnhContVirtServiceIndex,
		SlbStatEnhContRuleIndex:        slbStatEnhContRuleIndex,
		Params:                         params,
	}
}

func (c *SlbStatEnhContRuleTable) Name() string {
	return "SlbStatEnhContRuleTable"
}

func (c *SlbStatEnhContRuleTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbStatEnhContRuleTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbStatEnhContRuleTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbStatEnhContVirtServIndex).IsZero() &&
		reflect.ValueOf(c.SlbStatEnhContVirtServiceIndex).IsZero() &&
		reflect.ValueOf(c.SlbStatEnhContRuleIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbStatEnhContVirtServIndex) + "/" + fmt.Sprint(c.SlbStatEnhContVirtServiceIndex) + "/" + fmt.Sprint(c.SlbStatEnhContRuleIndex)
}

type SlbStatEnhContRuleTableAction int32

const (
	SlbStatEnhContRuleTableAction_Group       SlbStatEnhContRuleTableAction = 1
	SlbStatEnhContRuleTableAction_Appredir    SlbStatEnhContRuleTableAction = 2
	SlbStatEnhContRuleTableAction_Discard     SlbStatEnhContRuleTableAction = 3
	SlbStatEnhContRuleTableAction_Goto        SlbStatEnhContRuleTableAction = 4
	SlbStatEnhContRuleTableAction_Unsupported SlbStatEnhContRuleTableAction = 2147483647
)

type SlbStatEnhContRuleTableParams struct {
	// The virtual server id.
	VirtServIndex string `json:"VirtServIndex,omitempty"`
	// Virtual service index.
	VirtServiceIndex int32 `json:"VirtServiceIndex,omitempty"`
	// The content rule index.
	Index int32 `json:"Index,omitempty"`
	// The content rule name.
	Name string `json:"Name,omitempty"`
	// The content rule content class.
	ContClass string `json:"ContClass,omitempty"`
	// The content rule action.
	Action SlbStatEnhContRuleTableAction `json:"Action,omitempty"`
	// The content rule current sessions.
	CurrSess int32 `json:"CurrSess,omitempty"`
	// The content rule total sessions.
	TotSess int32 `json:"TotSess,omitempty"`
	// The content rule highest sessions.
	HighSess int32 `json:"HighSess,omitempty"`
	// The content rule total octets.
	TotOcts uint64 `json:"TotOcts,omitempty"`
	// The average ClientRTT in ms during the past collection interval.
	ClientRtt string `json:"ClientRtt,omitempty"`
	// The average ServerRTT in ms during the past collection interval.
	ServerRtt string `json:"ServerRtt,omitempty"`
	// PPS stats
	Pps int32 `json:"Pps,omitempty"`
}

func (p SlbStatEnhContRuleTableParams) iMABean() {}
