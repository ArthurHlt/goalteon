package beans

import (
	"fmt"
	"reflect"
)

// SlbStatSpAuxSessTable The auxiliary session table statistics for a particular SP.
// Note:This mib is not supported for VX instance of Virtualization.
type SlbStatSpAuxSessTable struct {
	// The SP number.
	SlbStatSpAuxSessSpIndex int32
	// The auxiliary session table number.
	SlbStatSpAuxSessIndex int32
	Params                *SlbStatSpAuxSessTableParams
}

func NewSlbStatSpAuxSessTableList() *SlbStatSpAuxSessTable {
	return &SlbStatSpAuxSessTable{}
}

func NewSlbStatSpAuxSessTable(
	slbStatSpAuxSessSpIndex int32,
	slbStatSpAuxSessIndex int32,
	params *SlbStatSpAuxSessTableParams,
) *SlbStatSpAuxSessTable {
	return &SlbStatSpAuxSessTable{
		SlbStatSpAuxSessSpIndex: slbStatSpAuxSessSpIndex,
		SlbStatSpAuxSessIndex:   slbStatSpAuxSessIndex,
		Params:                  params,
	}
}

func (c *SlbStatSpAuxSessTable) Name() string {
	return "SlbStatSpAuxSessTable"
}

func (c *SlbStatSpAuxSessTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbStatSpAuxSessTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbStatSpAuxSessTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbStatSpAuxSessSpIndex).IsZero() &&
		reflect.ValueOf(c.SlbStatSpAuxSessIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbStatSpAuxSessSpIndex) + "/" + fmt.Sprint(c.SlbStatSpAuxSessIndex)
}

type SlbStatSpAuxSessTableParams struct {
	// The SP number.
	SpIndex int32 `json:"SpIndex,omitempty"`
	// The auxiliary session table number.
	Index int32 `json:"Index,omitempty"`
	// The number of connections in this auxiliary sessions table on the SP.
	CurConn uint32 `json:"CurConn,omitempty"`
	// The maximum number of connections handled by this auxiliary session
	// table on the SP.
	MaxConn int32 `json:"MaxConn,omitempty"`
	// The total number of allocation failures for this auxiliary session
	// table on the SP.
	AllocFails uint32 `json:"AllocFails,omitempty"`
}

func (p SlbStatSpAuxSessTableParams) iMABean() {}
