package beans

import (
	"fmt"
	"reflect"
)

// SlbNewCfgAppShapeTable The table of AppShape,
// Note:This mib is not supported for VX instance of Virtualization.
type SlbNewCfgAppShapeTable struct {
	// The AppShape number
	SlbNewCfgAppShapeIndex string
	Params                 *SlbNewCfgAppShapeTableParams
}

func NewSlbNewCfgAppShapeTableList() *SlbNewCfgAppShapeTable {
	return &SlbNewCfgAppShapeTable{}
}

func NewSlbNewCfgAppShapeTable(
	slbNewCfgAppShapeIndex string,
	params *SlbNewCfgAppShapeTableParams,
) *SlbNewCfgAppShapeTable {
	return &SlbNewCfgAppShapeTable{
		SlbNewCfgAppShapeIndex: slbNewCfgAppShapeIndex,
		Params:                 params,
	}
}

func (c *SlbNewCfgAppShapeTable) Name() string {
	return "SlbNewCfgAppShapeTable"
}

func (c *SlbNewCfgAppShapeTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbNewCfgAppShapeTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbNewCfgAppShapeTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbNewCfgAppShapeIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbNewCfgAppShapeIndex)
}

type SlbNewCfgAppShapeTableState int32

const (
	SlbNewCfgAppShapeTableState_Enabled     SlbNewCfgAppShapeTableState = 1
	SlbNewCfgAppShapeTableState_Disabled    SlbNewCfgAppShapeTableState = 2
	SlbNewCfgAppShapeTableState_Unsupported SlbNewCfgAppShapeTableState = 2147483647
)

type SlbNewCfgAppShapeTableDelete int32

const (
	SlbNewCfgAppShapeTableDelete_Other       SlbNewCfgAppShapeTableDelete = 1
	SlbNewCfgAppShapeTableDelete_Delete      SlbNewCfgAppShapeTableDelete = 2
	SlbNewCfgAppShapeTableDelete_Unsupported SlbNewCfgAppShapeTableDelete = 2147483647
)

type SlbNewCfgAppShapeTableDefault int32

const (
	SlbNewCfgAppShapeTableDefault_Other       SlbNewCfgAppShapeTableDefault = 1
	SlbNewCfgAppShapeTableDefault_Default     SlbNewCfgAppShapeTableDefault = 2
	SlbNewCfgAppShapeTableDefault_Unsupported SlbNewCfgAppShapeTableDefault = 2147483647
)

type SlbNewCfgAppShapeTableParams struct {
	// The AppShape number
	Index string `json:"Index,omitempty"`
	// The name of the Apprule.
	Name string `json:"Name,omitempty"`
	// Enable or disable AppShape.
	State SlbNewCfgAppShapeTableState `json:"State,omitempty"`
	// When set to the value of 2 (delete), the entire row is deleted.
	// When read, other(1) is returned. Setting the value to anything
	// other than 2(delete) has no effect on the state of the row.
	Delete SlbNewCfgAppShapeTableDelete `json:"Delete,omitempty"`
	// Value of 2(default),will restore the default APM AppShape++ script.
	// When read, other(1) is returned. Setting the value to anything
	// other than 2(default) has no effect. Note that default is only for
	// APM_script not for any other script
	Default SlbNewCfgAppShapeTableDefault `json:"Default,omitempty"`
}

func (p SlbNewCfgAppShapeTableParams) iMABean() {}
