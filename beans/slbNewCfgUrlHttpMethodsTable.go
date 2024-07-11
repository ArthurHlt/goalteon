package beans

import (
	"fmt"
	"reflect"
)

// SlbNewCfgUrlHttpMethodsTable The http methods table in layer7 processing engine.
// Note:This mib is not supported for VX instance of virtualization.
type SlbNewCfgUrlHttpMethodsTable struct {
	// The http method table index.
	SlbNewCfgUrlHttpMethodIndex int32
	Params                      *SlbNewCfgUrlHttpMethodsTableParams
}

func NewSlbNewCfgUrlHttpMethodsTableList() *SlbNewCfgUrlHttpMethodsTable {
	return &SlbNewCfgUrlHttpMethodsTable{}
}

func NewSlbNewCfgUrlHttpMethodsTable(
	slbNewCfgUrlHttpMethodIndex int32,
	params *SlbNewCfgUrlHttpMethodsTableParams,
) *SlbNewCfgUrlHttpMethodsTable {
	return &SlbNewCfgUrlHttpMethodsTable{
		SlbNewCfgUrlHttpMethodIndex: slbNewCfgUrlHttpMethodIndex,
		Params:                      params,
	}
}

func (c *SlbNewCfgUrlHttpMethodsTable) Name() string {
	return "SlbNewCfgUrlHttpMethodsTable"
}

func (c *SlbNewCfgUrlHttpMethodsTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbNewCfgUrlHttpMethodsTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbNewCfgUrlHttpMethodsTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbNewCfgUrlHttpMethodIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbNewCfgUrlHttpMethodIndex)
}

type SlbNewCfgUrlHttpMethodsTableDelete int32

const (
	SlbNewCfgUrlHttpMethodsTableDelete_Other       SlbNewCfgUrlHttpMethodsTableDelete = 1
	SlbNewCfgUrlHttpMethodsTableDelete_Delete      SlbNewCfgUrlHttpMethodsTableDelete = 2
	SlbNewCfgUrlHttpMethodsTableDelete_Unsupported SlbNewCfgUrlHttpMethodsTableDelete = 2147483647
)

type SlbNewCfgUrlHttpMethodsTableParams struct {
	// The http method table index.
	Index int32 `json:"Index,omitempty"`
	// The http method string to process in layer7 engine.
	StringVal string `json:"StringVal,omitempty"`
	// When set to the value of 2 (delete), the entire row is deleted.
	// When read, other(1) is returned. Setting the value to anything
	// other than 2(delete) has no effect on the state of the row.
	Delete SlbNewCfgUrlHttpMethodsTableDelete `json:"Delete,omitempty"`
}

func (p SlbNewCfgUrlHttpMethodsTableParams) iMABean() {}
