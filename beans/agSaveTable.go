package beans

import (
	"fmt"
	"reflect"
)

// AgSaveTable The table of error messages for asynchrnous save operation.
type AgSaveTable struct {
	// The table index for save.
	AgSaveIndex int32
	Params      *AgSaveTableParams
}

func NewAgSaveTableList() *AgSaveTable {
	return &AgSaveTable{}
}

func NewAgSaveTable(
	agSaveIndex int32,
	params *AgSaveTableParams,
) *AgSaveTable {
	return &AgSaveTable{
		AgSaveIndex: agSaveIndex,
		Params:      params,
	}
}

func (c *AgSaveTable) Name() string {
	return "AgSaveTable"
}

func (c *AgSaveTable) GetParams() BeanType {
	return c.Params
}

func (c *AgSaveTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *AgSaveTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.AgSaveIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.AgSaveIndex)
}

type AgSaveTableParams struct {
	// The table index for save.
	Index int32 `json:"Index,omitempty"`
	// A string in the Save table.
	StringVal string `json:"StringVal,omitempty"`
}

func (p AgSaveTableParams) iMABean() {}
