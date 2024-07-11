package beans

import (
	"fmt"
	"reflect"
)

// SlbCurCfgUrlHttpMethodsTable The http methods table in layer7 processing engine.
// Note:This mib is not supported for VX instance of virtualization.
type SlbCurCfgUrlHttpMethodsTable struct {
	// The http method table index.
	SlbCurCfgUrlHttpMethodIndex int32
	Params                      *SlbCurCfgUrlHttpMethodsTableParams
}

func NewSlbCurCfgUrlHttpMethodsTableList() *SlbCurCfgUrlHttpMethodsTable {
	return &SlbCurCfgUrlHttpMethodsTable{}
}

func NewSlbCurCfgUrlHttpMethodsTable(
	slbCurCfgUrlHttpMethodIndex int32,
	params *SlbCurCfgUrlHttpMethodsTableParams,
) *SlbCurCfgUrlHttpMethodsTable {
	return &SlbCurCfgUrlHttpMethodsTable{
		SlbCurCfgUrlHttpMethodIndex: slbCurCfgUrlHttpMethodIndex,
		Params:                      params,
	}
}

func (c *SlbCurCfgUrlHttpMethodsTable) Name() string {
	return "SlbCurCfgUrlHttpMethodsTable"
}

func (c *SlbCurCfgUrlHttpMethodsTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbCurCfgUrlHttpMethodsTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbCurCfgUrlHttpMethodsTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbCurCfgUrlHttpMethodIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbCurCfgUrlHttpMethodIndex)
}

type SlbCurCfgUrlHttpMethodsTableParams struct {
	// The http method table index.
	Index int32 `json:"Index,omitempty"`
	// The http method string to process in layer7 engine.
	StringVal string `json:"StringVal,omitempty"`
}

func (p SlbCurCfgUrlHttpMethodsTableParams) iMABean() {}
