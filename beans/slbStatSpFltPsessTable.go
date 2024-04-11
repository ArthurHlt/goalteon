package beans

import (
	"fmt"
	"reflect"
)

// SlbStatSpFltPsessTable The filter persistent session statistics table.  This table shows the statistics
// of filters for a particular SP.
// Note:This mib is not supported for VX instance of Virtualization.
type SlbStatSpFltPsessTable struct {
	// The SP index.
	SlbStatSpFltSpIndex int32
	Params              *SlbStatSpFltPsessTableParams
}

func NewSlbStatSpFltPsessTableList() *SlbStatSpFltPsessTable {
	return &SlbStatSpFltPsessTable{}
}

func NewSlbStatSpFltPsessTable(
	slbStatSpFltSpIndex int32,
	params *SlbStatSpFltPsessTableParams,
) *SlbStatSpFltPsessTable {
	return &SlbStatSpFltPsessTable{
		SlbStatSpFltSpIndex: slbStatSpFltSpIndex,
		Params:              params,
	}
}

func (c *SlbStatSpFltPsessTable) Name() string {
	return "SlbStatSpFltPsessTable"
}

func (c *SlbStatSpFltPsessTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbStatSpFltPsessTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbStatSpFltPsessTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbStatSpFltSpIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbStatSpFltSpIndex)
}

type SlbStatSpFltPsessTableParams struct {
	// The SP index.
	SpIndex int32 `json:"SpIndex,omitempty"`
	// The current number of persistent session for filters on this SP.
	Cur uint64 `json:"Cur,omitempty"`
	// The hightest number of persistent session for filters on this SP.
	High uint64 `json:"High,omitempty"`
	// The total number of persistent session for filters on this SP.
	Total uint64 `json:"Total,omitempty"`
}

func (p SlbStatSpFltPsessTableParams) iMABean() {}
