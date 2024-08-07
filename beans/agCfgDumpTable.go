package beans

import (
	"fmt"
	"reflect"
)

// AgCfgDumpTable The table of dump strings.
type AgCfgDumpTable struct {
	// The table index.
	AgCfgDumpIndex int32
	Params         *AgCfgDumpTableParams
}

func NewAgCfgDumpTableList() *AgCfgDumpTable {
	return &AgCfgDumpTable{}
}

func NewAgCfgDumpTable(
	agCfgDumpIndex int32,
	params *AgCfgDumpTableParams,
) *AgCfgDumpTable {
	return &AgCfgDumpTable{
		AgCfgDumpIndex: agCfgDumpIndex,
		Params:         params,
	}
}

func (c *AgCfgDumpTable) Name() string {
	return "AgCfgDumpTable"
}

func (c *AgCfgDumpTable) GetParams() BeanType {
	return c.Params
}

func (c *AgCfgDumpTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *AgCfgDumpTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.AgCfgDumpIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.AgCfgDumpIndex)
}

type AgCfgDumpTableParams struct {
	// The table index.
	Index int32 `json:"Index,omitempty"`
	// A string in the dump table.
	StringVal string `json:"StringVal,omitempty"`
}

func (p AgCfgDumpTableParams) iMABean() {}
