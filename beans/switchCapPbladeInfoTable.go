package beans

import (
	"fmt"
	"reflect"
)

// SwitchCapPbladeInfoTable Pblade Capacity information table
type SwitchCapPbladeInfoTable struct {
	// Pblade slot number
	SwitchCapPbladeInfoPbladeId uint64
	Params                      *SwitchCapPbladeInfoTableParams
}

func NewSwitchCapPbladeInfoTableList() *SwitchCapPbladeInfoTable {
	return &SwitchCapPbladeInfoTable{}
}

func NewSwitchCapPbladeInfoTable(
	switchCapPbladeInfoPbladeId uint64,
	params *SwitchCapPbladeInfoTableParams,
) *SwitchCapPbladeInfoTable {
	return &SwitchCapPbladeInfoTable{
		SwitchCapPbladeInfoPbladeId: switchCapPbladeInfoPbladeId,
		Params:                      params,
	}
}

func (c *SwitchCapPbladeInfoTable) Name() string {
	return "SwitchCapPbladeInfoTable"
}

func (c *SwitchCapPbladeInfoTable) GetParams() BeanType {
	return c.Params
}

func (c *SwitchCapPbladeInfoTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SwitchCapPbladeInfoTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SwitchCapPbladeInfoPbladeId).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SwitchCapPbladeInfoPbladeId)
}

type SwitchCapPbladeInfoTableParams struct {
	// Pblade slot number
	PbladeId uint64 `json:"PbladeId,omitempty"`
	// Pblade total RAM size (GB)
	RamMax int32 `json:"RamMax,omitempty"`
	// Pblade used RAM size (GB)
	RamCur int32 `json:"RamCur,omitempty"`
}

func (p SwitchCapPbladeInfoTableParams) iMABean() {}
