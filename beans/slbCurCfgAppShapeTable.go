package beans

import (
	"fmt"
	"reflect"
)

// SlbCurCfgAppShapeTable The table of AppShape,
// Note:This mib is not supported for VX instance of Virtualization.
type SlbCurCfgAppShapeTable struct {
	// The AppShape number
	SlbCurCfgAppShapeIndex string
	Params                 *SlbCurCfgAppShapeTableParams
}

func NewSlbCurCfgAppShapeTableList() *SlbCurCfgAppShapeTable {
	return &SlbCurCfgAppShapeTable{}
}

func NewSlbCurCfgAppShapeTable(
	slbCurCfgAppShapeIndex string,
	params *SlbCurCfgAppShapeTableParams,
) *SlbCurCfgAppShapeTable {
	return &SlbCurCfgAppShapeTable{
		SlbCurCfgAppShapeIndex: slbCurCfgAppShapeIndex,
		Params:                 params,
	}
}

func (c *SlbCurCfgAppShapeTable) Name() string {
	return "SlbCurCfgAppShapeTable"
}

func (c *SlbCurCfgAppShapeTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbCurCfgAppShapeTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbCurCfgAppShapeTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbCurCfgAppShapeIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbCurCfgAppShapeIndex)
}

type SlbCurCfgAppShapeTableState int32

const (
	SlbCurCfgAppShapeTableState_Enabled     SlbCurCfgAppShapeTableState = 1
	SlbCurCfgAppShapeTableState_Disabled    SlbCurCfgAppShapeTableState = 2
	SlbCurCfgAppShapeTableState_Unsupported SlbCurCfgAppShapeTableState = 2147483647
)

type SlbCurCfgAppShapeTableParams struct {
	// The AppShape number
	Index string `json:"Index,omitempty"`
	// The name of the apprule.
	Name string `json:"Name,omitempty"`
	// Enable or disable AppShape.
	State SlbCurCfgAppShapeTableState `json:"State,omitempty"`
}

func (p SlbCurCfgAppShapeTableParams) iMABean() {}
