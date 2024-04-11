package beans

import (
	"fmt"
	"reflect"
)

// FltNewCfgPortTable The filtering port table in the new configuration block.
// Note:This mib is not supported for VX instance of Virtualization.
type FltNewCfgPortTable struct {
	// The port index.
	FltNewCfgPortIndx int32
	Params            *FltNewCfgPortTableParams
}

func NewFltNewCfgPortTableList() *FltNewCfgPortTable {
	return &FltNewCfgPortTable{}
}

func NewFltNewCfgPortTable(
	fltNewCfgPortIndx int32,
	params *FltNewCfgPortTableParams,
) *FltNewCfgPortTable {
	return &FltNewCfgPortTable{
		FltNewCfgPortIndx: fltNewCfgPortIndx,
		Params:            params,
	}
}

func (c *FltNewCfgPortTable) Name() string {
	return "FltNewCfgPortTable"
}

func (c *FltNewCfgPortTable) GetParams() BeanType {
	return c.Params
}

func (c *FltNewCfgPortTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *FltNewCfgPortTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.FltNewCfgPortIndx).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.FltNewCfgPortIndx)
}

type FltNewCfgPortTableState int32

const (
	FltNewCfgPortTableState_Enabled     FltNewCfgPortTableState = 1
	FltNewCfgPortTableState_Disabled    FltNewCfgPortTableState = 2
	FltNewCfgPortTableState_Unsupported FltNewCfgPortTableState = 2147483647
)

type FltNewCfgPortTableParams struct {
	// The port index.
	Indx int32 `json:"Indx,omitempty"`
	// Enable or disable filtering.
	State FltNewCfgPortTableState `json:"State,omitempty"`
	// The filtering rules applied to the port.  The filtering rules
	// are presented in bitmap format.
	// in receiving order:
	// 	     OCTET 1  OCTET 2  .....
	// xxxxxxxx xxxxxxxx .....
	// ||    || |_ filter 9
	// ||    ||
	// ||    ||___ filter 8
	// ||    |____ filter 7
	// ||      .    .   .
	// ||_________ filter 2
	// |__________ filter 1 (as index to fltNewCfgTable)
	// where x : 1 - The represented filter rule applied to the port
	// 		   0 - The represented filter rule not applied to the port
	FiltBmap string `json:"FiltBmap,omitempty"`
	// This is an action object to add filtering rule to a port. The value
	// specified with this object is the index to the fltNewCfgTable for
	// which filtering rule to be added to the port. The range of the
	// valid index is between 1 and fltCurCfgTableMaxSize. When read, the
	// value '0' is returned always.
	AddFiltRule int32 `json:"AddFiltRule,omitempty"`
	// This is an action object to remove filtering rule from a port.	The
	// value specified with this object is the index to the fltNewCfgTable
	// for which filtering rule to be removed from the port. The range of
	// the valid index is between 1 and fltCurCfgTableMaxSize.  When read,
	// the value '0' is returned always.
	RemFiltRule int32 `json:"RemFiltRule,omitempty"`
}

func (p FltNewCfgPortTableParams) iMABean() {}
