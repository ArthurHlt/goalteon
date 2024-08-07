package beans

import (
	"fmt"
	"reflect"
)

// AgDiffTable The table of diff strings.
type AgDiffTable struct {
	// The table index.
	AgDiffIndex int32
	Params      *AgDiffTableParams
}

func NewAgDiffTableList() *AgDiffTable {
	return &AgDiffTable{}
}

func NewAgDiffTable(
	agDiffIndex int32,
	params *AgDiffTableParams,
) *AgDiffTable {
	return &AgDiffTable{
		AgDiffIndex: agDiffIndex,
		Params:      params,
	}
}

func (c *AgDiffTable) Name() string {
	return "AgDiffTable"
}

func (c *AgDiffTable) GetParams() BeanType {
	return c.Params
}

func (c *AgDiffTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *AgDiffTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.AgDiffIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.AgDiffIndex)
}

type AgDiffTableParams struct {
	// The table index.
	Index int32 `json:"Index,omitempty"`
	// A string in the diff table.
	StringVal string `json:"StringVal,omitempty"`
}

func (p AgDiffTableParams) iMABean() {}
