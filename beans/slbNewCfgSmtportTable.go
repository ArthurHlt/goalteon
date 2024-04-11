package beans

import (
	"fmt"
	"reflect"
)

// SlbNewCfgSmtportTable The table of real server service ports.
// Note:This mib is not supported for VX instance of Virtualization.
type SlbNewCfgSmtportTable struct {
	// The service mapping real port index
	SlbNewCfgSmtportIndex int32
	Params                *SlbNewCfgSmtportTableParams
}

func NewSlbNewCfgSmtportTableList() *SlbNewCfgSmtportTable {
	return &SlbNewCfgSmtportTable{}
}

func NewSlbNewCfgSmtportTable(
	slbNewCfgSmtportIndex int32,
	params *SlbNewCfgSmtportTableParams,
) *SlbNewCfgSmtportTable {
	return &SlbNewCfgSmtportTable{
		SlbNewCfgSmtportIndex: slbNewCfgSmtportIndex,
		Params:                params,
	}
}

func (c *SlbNewCfgSmtportTable) Name() string {
	return "SlbNewCfgSmtportTable"
}

func (c *SlbNewCfgSmtportTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbNewCfgSmtportTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbNewCfgSmtportTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbNewCfgSmtportIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbNewCfgSmtportIndex)
}

type SlbNewCfgSmtportTableDelete int32

const (
	SlbNewCfgSmtportTableDelete_Other       SlbNewCfgSmtportTableDelete = 1
	SlbNewCfgSmtportTableDelete_Delete      SlbNewCfgSmtportTableDelete = 2
	SlbNewCfgSmtportTableDelete_Unsupported SlbNewCfgSmtportTableDelete = 2147483647
)

type SlbNewCfgSmtportTableParams struct {
	// The service mapping real port index
	Index int32 `json:"Index,omitempty"`
	// The service mapping real port number.
	Num uint64 `json:"Num,omitempty"`
	// When set to the value of 2 (delete), the entire row is deleted.
	// When read, other(1) is returned. Setting the value to anything
	// other than 2(delete) has no effect on the state of the row.
	Delete SlbNewCfgSmtportTableDelete `json:"Delete,omitempty"`
}

func (p SlbNewCfgSmtportTableParams) iMABean() {}
