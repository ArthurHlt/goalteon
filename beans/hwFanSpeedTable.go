package beans

import (
	"fmt"
	"reflect"
)

// HwFanSpeedTable The table of fans speed.
type HwFanSpeedTable struct {
	// The fan slot index.
	HwFanSpeedSlotIndex int32
	// The fan sensor index.
	HwFanSpeedSensorIndex int32
	Params                *HwFanSpeedTableParams
}

func NewHwFanSpeedTableList() *HwFanSpeedTable {
	return &HwFanSpeedTable{}
}

func NewHwFanSpeedTable(
	hwFanSpeedSlotIndex int32,
	hwFanSpeedSensorIndex int32,
	params *HwFanSpeedTableParams,
) *HwFanSpeedTable {
	return &HwFanSpeedTable{
		HwFanSpeedSlotIndex:   hwFanSpeedSlotIndex,
		HwFanSpeedSensorIndex: hwFanSpeedSensorIndex,
		Params:                params,
	}
}

func (c *HwFanSpeedTable) Name() string {
	return "HwFanSpeedTable"
}

func (c *HwFanSpeedTable) GetParams() BeanType {
	return c.Params
}

func (c *HwFanSpeedTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *HwFanSpeedTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.HwFanSpeedSlotIndex).IsZero() &&
		reflect.ValueOf(c.HwFanSpeedSensorIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.HwFanSpeedSlotIndex) + "/" + fmt.Sprint(c.HwFanSpeedSensorIndex)
}

type HwFanSpeedTableParams struct {
	// The fan slot index.
	SlotIndex int32 `json:"SlotIndex,omitempty"`
	// The fan sensor index.
	SensorIndex int32 `json:"SensorIndex,omitempty"`
	// The fan speed.
	Val int32 `json:"Val,omitempty"`
}

func (p HwFanSpeedTableParams) iMABean() {}
