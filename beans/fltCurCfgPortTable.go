package beans

import (
	"fmt"
	"reflect"
)

// FltCurCfgPortTable The filtering port table in the current configuration block.
// Note:This mib is not supported for VX instance of Virtualization.
type FltCurCfgPortTable struct {
	// The port index.
	FltCurCfgPortIndx int32
	Params            *FltCurCfgPortTableParams
}

func NewFltCurCfgPortTableList() *FltCurCfgPortTable {
	return &FltCurCfgPortTable{}
}

func NewFltCurCfgPortTable(
	fltCurCfgPortIndx int32,
	params *FltCurCfgPortTableParams,
) *FltCurCfgPortTable {
	return &FltCurCfgPortTable{
		FltCurCfgPortIndx: fltCurCfgPortIndx,
		Params:            params,
	}
}

func (c *FltCurCfgPortTable) Name() string {
	return "FltCurCfgPortTable"
}

func (c *FltCurCfgPortTable) GetParams() BeanType {
	return c.Params
}

func (c *FltCurCfgPortTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *FltCurCfgPortTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.FltCurCfgPortIndx).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.FltCurCfgPortIndx)
}

type FltCurCfgPortTableState int32

const (
	FltCurCfgPortTableState_Enabled     FltCurCfgPortTableState = 1
	FltCurCfgPortTableState_Disabled    FltCurCfgPortTableState = 2
	FltCurCfgPortTableState_Unsupported FltCurCfgPortTableState = 2147483647
)

type FltCurCfgPortTableParams struct {
	// The port index.
	Indx int32 `json:"Indx,omitempty"`
	// Enable or disable filtering.
	State FltCurCfgPortTableState `json:"State,omitempty"`
	// The filtering rules applied to the port.  The filtering rules are
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
	// where x : 1 - The represented filter rule applied to the port
	// 		   0 - The represented filter rule not applied to the port
	FiltBmap string `json:"FiltBmap,omitempty"`
}

func (p FltCurCfgPortTableParams) iMABean() {}
