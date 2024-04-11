package beans

import (
	"fmt"
	"reflect"
)

// SlbCurCfgFiltsetTable The table of filter sets and their SLB states,
// Note:This mib is not supported for VX instance of Virtualization.
type SlbCurCfgFiltsetTable struct {
	// The filter set number for which the SLB information pertains.
	SlbCurCfgFiltsetIndex int32
	Params                *SlbCurCfgFiltsetTableParams
}

func NewSlbCurCfgFiltsetTableList() *SlbCurCfgFiltsetTable {
	return &SlbCurCfgFiltsetTable{}
}

func NewSlbCurCfgFiltsetTable(
	slbCurCfgFiltsetIndex int32,
	params *SlbCurCfgFiltsetTableParams,
) *SlbCurCfgFiltsetTable {
	return &SlbCurCfgFiltsetTable{
		SlbCurCfgFiltsetIndex: slbCurCfgFiltsetIndex,
		Params:                params,
	}
}

func (c *SlbCurCfgFiltsetTable) Name() string {
	return "SlbCurCfgFiltsetTable"
}

func (c *SlbCurCfgFiltsetTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbCurCfgFiltsetTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbCurCfgFiltsetTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbCurCfgFiltsetIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbCurCfgFiltsetIndex)
}

type SlbCurCfgFiltsetTableParams struct {
	// The filter set number for which the SLB information pertains.
	Index int32 `json:"Index,omitempty"`
	// The filtering rules applied to the filter set.  The filtering rules are
	// presented in bitmap format.
	// 	 in receiving order:
	// 	     OCTET 1  OCTET 2  .....
	// xxxxxxxx xxxxxxxx .....
	// ||    || |_ filter 9
	// ||    ||
	// ||    ||___ filter 8
	// ||    |____ filter 7
	// ||      .    .   .
	// ||_________ filter 2
	// |__________ filter 1 (as index to fltCurCfgTable)
	// where x : 1 - The represented filter rule applied to the filter set
	// 		   0 - The represented filter rule not applied to the filter set
	SlbFilterBmap string `json:"SlbFilterBmap,omitempty"`
}

func (p SlbCurCfgFiltsetTableParams) iMABean() {}
