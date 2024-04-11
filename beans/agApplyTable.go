package beans

import (
	"fmt"
	"reflect"
)

// AgApplyTable The table of URL path for URL load balancing in the current_config.
type AgApplyTable struct {
	// The table index.
	AgApplyIndex int32
	Params       *AgApplyTableParams
}

func NewAgApplyTableList() *AgApplyTable {
	return &AgApplyTable{}
}

func NewAgApplyTable(
	agApplyIndex int32,
	params *AgApplyTableParams,
) *AgApplyTable {
	return &AgApplyTable{
		AgApplyIndex: agApplyIndex,
		Params:       params,
	}
}

func (c *AgApplyTable) Name() string {
	return "AgApplyTable"
}

func (c *AgApplyTable) GetParams() BeanType {
	return c.Params
}

func (c *AgApplyTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *AgApplyTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.AgApplyIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.AgApplyIndex)
}

type AgApplyTableParams struct {
	// The table index.
	Index int32 `json:"Index,omitempty"`
	// A string in the apply table.
	StringVal string `json:"StringVal,omitempty"`
}

func (p AgApplyTableParams) iMABean() {}
