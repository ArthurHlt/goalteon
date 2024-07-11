package beans

import (
	"fmt"
	"reflect"
)

// SlbStatEnhContRuleActionGroupTable The statistics table of instances that the content rule matched.
// Note:This mib is not supported for VX instance of virtualization.
type SlbStatEnhContRuleActionGroupTable struct {
	// The virtual server id.
	SlbStatEnhContActionVirtServIndex string
	// Virtual service index.
	SlbStatEnhContActionVirtServiceIndex int32
	// The content rule index.
	SlbStatEnhContRuleActionIndex int32
	// The real server id.
	SlbEnhContActionRealServerIndex string
	Params                          *SlbStatEnhContRuleActionGroupTableParams
}

func NewSlbStatEnhContRuleActionGroupTableList() *SlbStatEnhContRuleActionGroupTable {
	return &SlbStatEnhContRuleActionGroupTable{}
}

func NewSlbStatEnhContRuleActionGroupTable(
	slbStatEnhContActionVirtServIndex string,
	slbStatEnhContActionVirtServiceIndex int32,
	slbStatEnhContRuleActionIndex int32,
	slbEnhContActionRealServerIndex string,
	params *SlbStatEnhContRuleActionGroupTableParams,
) *SlbStatEnhContRuleActionGroupTable {
	return &SlbStatEnhContRuleActionGroupTable{
		SlbStatEnhContActionVirtServIndex:    slbStatEnhContActionVirtServIndex,
		SlbStatEnhContActionVirtServiceIndex: slbStatEnhContActionVirtServiceIndex,
		SlbStatEnhContRuleActionIndex:        slbStatEnhContRuleActionIndex,
		SlbEnhContActionRealServerIndex:      slbEnhContActionRealServerIndex,
		Params:                               params,
	}
}

func (c *SlbStatEnhContRuleActionGroupTable) Name() string {
	return "SlbStatEnhContRuleActionGroupTable"
}

func (c *SlbStatEnhContRuleActionGroupTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbStatEnhContRuleActionGroupTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbStatEnhContRuleActionGroupTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbStatEnhContActionVirtServIndex).IsZero() &&
		reflect.ValueOf(c.SlbStatEnhContActionVirtServiceIndex).IsZero() &&
		reflect.ValueOf(c.SlbStatEnhContRuleActionIndex).IsZero() &&
		reflect.ValueOf(c.SlbEnhContActionRealServerIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbStatEnhContActionVirtServIndex) + "/" + fmt.Sprint(c.SlbStatEnhContActionVirtServiceIndex) + "/" + fmt.Sprint(c.SlbStatEnhContRuleActionIndex) + "/" + fmt.Sprint(c.SlbEnhContActionRealServerIndex)
}

type SlbStatEnhContRuleActionGroupTableRealStatus int32

const (
	SlbStatEnhContRuleActionGroupTableRealStatus_Up          SlbStatEnhContRuleActionGroupTableRealStatus = 0
	SlbStatEnhContRuleActionGroupTableRealStatus_Down        SlbStatEnhContRuleActionGroupTableRealStatus = 1
	SlbStatEnhContRuleActionGroupTableRealStatus_Admindown   SlbStatEnhContRuleActionGroupTableRealStatus = 2
	SlbStatEnhContRuleActionGroupTableRealStatus_Warning     SlbStatEnhContRuleActionGroupTableRealStatus = 3
	SlbStatEnhContRuleActionGroupTableRealStatus_Shutdown    SlbStatEnhContRuleActionGroupTableRealStatus = 4
	SlbStatEnhContRuleActionGroupTableRealStatus_Error       SlbStatEnhContRuleActionGroupTableRealStatus = 5
	SlbStatEnhContRuleActionGroupTableRealStatus_Unsupported SlbStatEnhContRuleActionGroupTableRealStatus = 2147483647
)

type SlbStatEnhContRuleActionGroupTableParams struct {
	// The virtual server id.
	ActionVirtServIndex string `json:"ActionVirtServIndex,omitempty"`
	// Virtual service index.
	ActionVirtServiceIndex int32 `json:"ActionVirtServiceIndex,omitempty"`
	// The content rule index.
	Index int32 `json:"Index,omitempty"`
	// The real server id.
	EnhContActionRealServerIndex string `json:"EnhContActionRealServerIndex,omitempty"`
	// The content rule current sessions per real server.
	CurrSess int32 `json:"CurrSess,omitempty"`
	// The content rule total sessions per real server.
	TotSess int32 `json:"TotSess,omitempty"`
	// The content rule highest sessions per real server.
	HighSess int32 `json:"HighSess,omitempty"`
	// The content rule total octets per real server.
	TotOcts uint64 `json:"TotOcts,omitempty"`
	// The average ClientRTT in ms during the past collection interval.
	ClientRtt string `json:"ClientRtt,omitempty"`
	// The average ServerRTT in ms during the past collection interval.
	ServerRtt string `json:"ServerRtt,omitempty"`
	// The real server status. (0)Up, (1)Down, (2)Admin-Down,
	// 				 (3)Warning, (4)Shutdown, (5)Error
	RealStatus SlbStatEnhContRuleActionGroupTableRealStatus `json:"RealStatus,omitempty"`
	// Service health check failure reason.
	HcReason string `json:"HcReason,omitempty"`
	// The content rule packets per second on real server.
	Pps int32 `json:"Pps,omitempty"`
}

func (p SlbStatEnhContRuleActionGroupTableParams) iMABean() {}
