package beans

import (
	"fmt"
	"reflect"
)

// AgPortOperTable The table of port operations.
type AgPortOperTable struct {
	// The port index.
	PortOperIdx int32
	Params      *AgPortOperTableParams
}

func NewAgPortOperTableList() *AgPortOperTable {
	return &AgPortOperTable{}
}

func NewAgPortOperTable(
	portOperIdx int32,
	params *AgPortOperTableParams,
) *AgPortOperTable {
	return &AgPortOperTable{
		PortOperIdx: portOperIdx,
		Params:      params,
	}
}

func (c *AgPortOperTable) Name() string {
	return "AgPortOperTable"
}

func (c *AgPortOperTable) GetParams() BeanType {
	return c.Params
}

func (c *AgPortOperTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *AgPortOperTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.PortOperIdx).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.PortOperIdx)
}

type AgPortOperTablePortOperState int32

const (
	AgPortOperTablePortOperState_Enabled     AgPortOperTablePortOperState = 1
	AgPortOperTablePortOperState_Disabled    AgPortOperTablePortOperState = 2
	AgPortOperTablePortOperState_Unsupported AgPortOperTablePortOperState = 2147483647
)

type AgPortOperTablePortOperRmon int32

const (
	AgPortOperTablePortOperRmon_Enabled     AgPortOperTablePortOperRmon = 1
	AgPortOperTablePortOperRmon_Disabled    AgPortOperTablePortOperRmon = 2
	AgPortOperTablePortOperRmon_Unsupported AgPortOperTablePortOperRmon = 2147483647
)

type AgPortOperTableParams struct {
	// The port index.
	PortOperIdx int32 `json:"portOperIdx,omitempty"`
	// Enable/Disable port.
	PortOperState AgPortOperTablePortOperState `json:"portOperState,omitempty"`
	// Enable/Disable RMON for port.
	PortOperRmon AgPortOperTablePortOperRmon `json:"portOperRmon,omitempty"`
}

func (p AgPortOperTableParams) iMABean() {}
