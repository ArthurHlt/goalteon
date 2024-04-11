package beans

import (
	"fmt"
	"reflect"
)

// SlbCurSecCfgSecPolTable The table for configuring security policy.
// Note:This mib is not supported on VX instance of Virtualization.
type SlbCurSecCfgSecPolTable struct {
	// The Security policy name(key id) as an index.
	SlbCurSecCfgSecPolNameIdIndex string
	Params                        *SlbCurSecCfgSecPolTableParams
}

func NewSlbCurSecCfgSecPolTableList() *SlbCurSecCfgSecPolTable {
	return &SlbCurSecCfgSecPolTable{}
}

func NewSlbCurSecCfgSecPolTable(
	slbCurSecCfgSecPolNameIdIndex string,
	params *SlbCurSecCfgSecPolTableParams,
) *SlbCurSecCfgSecPolTable {
	return &SlbCurSecCfgSecPolTable{
		SlbCurSecCfgSecPolNameIdIndex: slbCurSecCfgSecPolNameIdIndex,
		Params:                        params,
	}
}

func (c *SlbCurSecCfgSecPolTable) Name() string {
	return "SlbCurSecCfgSecPolTable"
}

func (c *SlbCurSecCfgSecPolTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbCurSecCfgSecPolTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbCurSecCfgSecPolTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbCurSecCfgSecPolNameIdIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbCurSecCfgSecPolNameIdIndex)
}

type SlbCurSecCfgSecPolTableAdminStatus int32

const (
	SlbCurSecCfgSecPolTableAdminStatus_Enabled     SlbCurSecCfgSecPolTableAdminStatus = 1
	SlbCurSecCfgSecPolTableAdminStatus_Disabled    SlbCurSecCfgSecPolTableAdminStatus = 2
	SlbCurSecCfgSecPolTableAdminStatus_Unsupported SlbCurSecCfgSecPolTableAdminStatus = 2147483647
)

type SlbCurSecCfgSecPolTableParams struct {
	// The Security policy name(key id) as an index.
	NameIdIndex string `json:"NameIdIndex,omitempty"`
	// Security policy name.
	Name string `json:"Name,omitempty"`
	// Security policy BW max threshold value.
	BW uint32 `json:"BW,omitempty"`
	// Security policy BW min threshold value.
	BWmin uint32 `json:"BWmin,omitempty"`
	// Security policy BW ratio value.
	BWin uint32 `json:"BWin,omitempty"`
	// Security policy PPS max threshold value.
	PPS uint32 `json:"PPS,omitempty"`
	// Security policy PPS min threshold value.
	PPSmin uint32 `json:"PPSmin,omitempty"`
	// Security policy PPS ratio value.
	PPSin uint32 `json:"PPSin,omitempty"`
	// Security policy CPS max threshold value.
	CPS uint64 `json:"CPS,omitempty"`
	// Security policy CPS min threshold value.
	CPSmin uint64 `json:"CPSmin,omitempty"`
	// Security policy CPS ratio value.
	CPSin uint32 `json:"CPSin,omitempty"`
	// Security policy CEC max threshold value.
	CEC uint32 `json:"CEC,omitempty"`
	// Security policy CEC min threshold value.
	CECmin uint32 `json:"CECmin,omitempty"`
	// Security policy CEC ratio value.
	CECin uint32 `json:"CECin,omitempty"`
	// Security policy learning period (in hours).
	LearnPeriod uint32 `json:"LearnPeriod,omitempty"`
	// Enable or disable security policy.
	AdminStatus SlbCurSecCfgSecPolTableAdminStatus `json:"AdminStatus,omitempty"`
	// Security policy Latency alert threshold (in microseconds).
	LatencyAbs uint32 `json:"LatencyAbs,omitempty"`
	// Security policy Latency increased ratio (current/history).
	LatencyPercent uint32 `json:"LatencyPercent,omitempty"`
	// Security policy Latency minimum alert threshold for detection.
	LatencyMin uint32 `json:"LatencyMin,omitempty"`
}

func (p SlbCurSecCfgSecPolTableParams) iMABean() {}
