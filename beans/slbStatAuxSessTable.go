package beans

import (
	"fmt"
	"reflect"
)

// SlbStatAuxSessTable The auxiliary session table statistics.
// Note:This mib is not supported for VX instance of Virtualization.
type SlbStatAuxSessTable struct {
	// The auxiliary session table number.
	SlbStatAuxSessIndex int32
	Params              *SlbStatAuxSessTableParams
}

func NewSlbStatAuxSessTableList() *SlbStatAuxSessTable {
	return &SlbStatAuxSessTable{}
}

func NewSlbStatAuxSessTable(
	slbStatAuxSessIndex int32,
	params *SlbStatAuxSessTableParams,
) *SlbStatAuxSessTable {
	return &SlbStatAuxSessTable{
		SlbStatAuxSessIndex: slbStatAuxSessIndex,
		Params:              params,
	}
}

func (c *SlbStatAuxSessTable) Name() string {
	return "SlbStatAuxSessTable"
}

func (c *SlbStatAuxSessTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbStatAuxSessTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbStatAuxSessTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbStatAuxSessIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbStatAuxSessIndex)
}

type SlbStatAuxSessTableParams struct {
	// The auxiliary session table number.
	Index int32 `json:"Index,omitempty"`
	// The number of connections in this auxiliary sessions table.
	CurConn uint32 `json:"CurConn,omitempty"`
	// The maximum number of connections handled by this auxiliary session
	// table.
	MaxConn int32 `json:"MaxConn,omitempty"`
	// The total number of allocation failures for this auxiliary session
	// table.
	AllocFails uint32 `json:"AllocFails,omitempty"`
}

func (p SlbStatAuxSessTableParams) iMABean() {}
