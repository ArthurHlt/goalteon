package beans

import (
	"fmt"
	"reflect"
)

// SlbNewSecCfgSecPolTable The table for configuring security policy.
// Note:This mib is not supported on VX instance of Virtualization.
type SlbNewSecCfgSecPolTable struct {
	// The Security policy name(key id) as an index.
	SlbNewSecCfgSecPolNameIdIndex string
	Params                        *SlbNewSecCfgSecPolTableParams
}

func NewSlbNewSecCfgSecPolTableList() *SlbNewSecCfgSecPolTable {
	return &SlbNewSecCfgSecPolTable{}
}

func NewSlbNewSecCfgSecPolTable(
	slbNewSecCfgSecPolNameIdIndex string,
	params *SlbNewSecCfgSecPolTableParams,
) *SlbNewSecCfgSecPolTable {
	return &SlbNewSecCfgSecPolTable{
		SlbNewSecCfgSecPolNameIdIndex: slbNewSecCfgSecPolNameIdIndex,
		Params:                        params,
	}
}

func (c *SlbNewSecCfgSecPolTable) Name() string {
	return "SlbNewSecCfgSecPolTable"
}

func (c *SlbNewSecCfgSecPolTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbNewSecCfgSecPolTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbNewSecCfgSecPolTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbNewSecCfgSecPolNameIdIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbNewSecCfgSecPolNameIdIndex)
}

type SlbNewSecCfgSecPolTableAdminStatus int32

const (
	SlbNewSecCfgSecPolTableAdminStatus_Enabled     SlbNewSecCfgSecPolTableAdminStatus = 1
	SlbNewSecCfgSecPolTableAdminStatus_Disabled    SlbNewSecCfgSecPolTableAdminStatus = 2
	SlbNewSecCfgSecPolTableAdminStatus_Unsupported SlbNewSecCfgSecPolTableAdminStatus = 2147483647
)

type SlbNewSecCfgSecPolTableDel int32

const (
	SlbNewSecCfgSecPolTableDel_Other               SlbNewSecCfgSecPolTableDel = 1
	SlbNewSecCfgSecPolTableDel_Delete              SlbNewSecCfgSecPolTableDel = 2
	SlbNewSecCfgSecPolTableDel_DelWithAssociations SlbNewSecCfgSecPolTableDel = 3
	SlbNewSecCfgSecPolTableDel_Unsupported         SlbNewSecCfgSecPolTableDel = 2147483647
)

type SlbNewSecCfgSecPolTableParams struct {
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
	AdminStatus SlbNewSecCfgSecPolTableAdminStatus `json:"AdminStatus,omitempty"`
	// Delete Security policy.
	Del SlbNewSecCfgSecPolTableDel `json:"Del,omitempty"`
	// Security policy Latency alert threshold (in microseconds).
	LatencyAbs uint32 `json:"LatencyAbs,omitempty"`
	// Security policy Latency increased ratio (current/history).
	LatencyPercent uint32 `json:"LatencyPercent,omitempty"`
	// Security policy Latency minimum alert threshold for detection.
	LatencyMin uint32 `json:"LatencyMin,omitempty"`
}

func (p SlbNewSecCfgSecPolTableParams) iMABean() {}
