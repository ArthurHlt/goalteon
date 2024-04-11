package beans

import (
	"fmt"
	"reflect"
)

// SlbCurCfgSmtportTable The service mapping real port table.
// Note:This mib is not supported for VX instance of Virtualization.
type SlbCurCfgSmtportTable struct {
	// The service mapping real port index.
	SlbCurCfgSmtportIndex int32
	Params                *SlbCurCfgSmtportTableParams
}

func NewSlbCurCfgSmtportTableList() *SlbCurCfgSmtportTable {
	return &SlbCurCfgSmtportTable{}
}

func NewSlbCurCfgSmtportTable(
	slbCurCfgSmtportIndex int32,
	params *SlbCurCfgSmtportTableParams,
) *SlbCurCfgSmtportTable {
	return &SlbCurCfgSmtportTable{
		SlbCurCfgSmtportIndex: slbCurCfgSmtportIndex,
		Params:                params,
	}
}

func (c *SlbCurCfgSmtportTable) Name() string {
	return "SlbCurCfgSmtportTable"
}

func (c *SlbCurCfgSmtportTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbCurCfgSmtportTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbCurCfgSmtportTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbCurCfgSmtportIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbCurCfgSmtportIndex)
}

type SlbCurCfgSmtportTableParams struct {
	// The service mapping real port index.
	Index int32 `json:"Index,omitempty"`
	// The service mapping real port number.
	Num uint64 `json:"Num,omitempty"`
}

func (p SlbCurCfgSmtportTableParams) iMABean() {}
