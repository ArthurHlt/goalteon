package beans

import (
	"fmt"
	"reflect"
)

// HwFanInfoTable The table of fan slot information.
type HwFanInfoTable struct {
	// fan slot index.
	HwFanInfoIndex int32
	Params         *HwFanInfoTableParams
}

func NewHwFanInfoTableList() *HwFanInfoTable {
	return &HwFanInfoTable{}
}

func NewHwFanInfoTable(
	hwFanInfoIndex int32,
	params *HwFanInfoTableParams,
) *HwFanInfoTable {
	return &HwFanInfoTable{
		HwFanInfoIndex: hwFanInfoIndex,
		Params:         params,
	}
}

func (c *HwFanInfoTable) Name() string {
	return "HwFanInfoTable"
}

func (c *HwFanInfoTable) GetParams() BeanType {
	return c.Params
}

func (c *HwFanInfoTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *HwFanInfoTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.HwFanInfoIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.HwFanInfoIndex)
}

type HwFanInfoTableStatus int32

const (
	HwFanInfoTableStatus_NotRelevant HwFanInfoTableStatus = 0
	HwFanInfoTableStatus_Ok          HwFanInfoTableStatus = 1
	HwFanInfoTableStatus_Failed      HwFanInfoTableStatus = 2
	HwFanInfoTableStatus_Unplugged   HwFanInfoTableStatus = 3
	HwFanInfoTableStatus_Empty       HwFanInfoTableStatus = 4
	HwFanInfoTableStatus_Unsupported HwFanInfoTableStatus = 2147483647
)

type HwFanInfoTableIsCritical int32

const (
	HwFanInfoTableIsCritical_No          HwFanInfoTableIsCritical = 0
	HwFanInfoTableIsCritical_Yes         HwFanInfoTableIsCritical = 1
	HwFanInfoTableIsCritical_NotRelevant HwFanInfoTableIsCritical = 2
	HwFanInfoTableIsCritical_Unsupported HwFanInfoTableIsCritical = 2147483647
)

type HwFanInfoTableParams struct {
	// fan slot index.
	Index int32 `json:"Index,omitempty"`
	// The status of the fan slot.
	Status HwFanInfoTableStatus `json:"Status,omitempty"`
	// The total count of fans per slot.
	TotalCount int32 `json:"TotalCount,omitempty"`
	// The total count of failed fans.
	CountFailed int32 `json:"CountFailed,omitempty"`
	// A flag indicates whether the fan is critical or not.
	IsCritical HwFanInfoTableIsCritical `json:"IsCritical,omitempty"`
}

func (p HwFanInfoTableParams) iMABean() {}
