package beans

import (
	"fmt"
	"reflect"
)

// HwTemperatureInfoTable The table of temperature slot information.
type HwTemperatureInfoTable struct {
	// The slot index
	HwTempInfoSlotIndex int32
	// The sensor index
	HwTempInfoSensorIndex int32
	Params                *HwTemperatureInfoTableParams
}

func NewHwTemperatureInfoTableList() *HwTemperatureInfoTable {
	return &HwTemperatureInfoTable{}
}

func NewHwTemperatureInfoTable(
	hwTempInfoSlotIndex int32,
	hwTempInfoSensorIndex int32,
	params *HwTemperatureInfoTableParams,
) *HwTemperatureInfoTable {
	return &HwTemperatureInfoTable{
		HwTempInfoSlotIndex:   hwTempInfoSlotIndex,
		HwTempInfoSensorIndex: hwTempInfoSensorIndex,
		Params:                params,
	}
}

func (c *HwTemperatureInfoTable) Name() string {
	return "HwTemperatureInfoTable"
}

func (c *HwTemperatureInfoTable) GetParams() BeanType {
	return c.Params
}

func (c *HwTemperatureInfoTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *HwTemperatureInfoTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.HwTempInfoSlotIndex).IsZero() &&
		reflect.ValueOf(c.HwTempInfoSensorIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.HwTempInfoSlotIndex) + "/" + fmt.Sprint(c.HwTempInfoSensorIndex)
}

type HwTemperatureInfoTableTempInfoSensorStatus int32

const (
	HwTemperatureInfoTableTempInfoSensorStatus_Low         HwTemperatureInfoTableTempInfoSensorStatus = 1
	HwTemperatureInfoTableTempInfoSensorStatus_Normal      HwTemperatureInfoTableTempInfoSensorStatus = 2
	HwTemperatureInfoTableTempInfoSensorStatus_High        HwTemperatureInfoTableTempInfoSensorStatus = 3
	HwTemperatureInfoTableTempInfoSensorStatus_Critical    HwTemperatureInfoTableTempInfoSensorStatus = 4
	HwTemperatureInfoTableTempInfoSensorStatus_Unsupported HwTemperatureInfoTableTempInfoSensorStatus = 2147483647
)

type HwTemperatureInfoTableParams struct {
	// The slot index
	TempInfoSlotIndex int32 `json:"TempInfoSlotIndex,omitempty"`
	// The sensor index
	TempInfoSensorIndex int32 `json:"TempInfoSensorIndex,omitempty"`
	// The sensor status
	TempInfoSensorStatus HwTemperatureInfoTableTempInfoSensorStatus `json:"TempInfoSensorStatus,omitempty"`
	// The sensor temperature
	TempInfoSensorTemp int32 `json:"TempInfoSensorTemp,omitempty"`
	// The sensor name
	TempInfoSensorStr string `json:"TempInfoSensorStr,omitempty"`
}

func (p HwTemperatureInfoTableParams) iMABean() {}
