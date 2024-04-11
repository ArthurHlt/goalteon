package beans

import (
	"fmt"
	"reflect"
)

// SlbCurCfgSerAppShapeTable The table of App Rules added to virtual services.
// Note:This mib is not supported for VX instance of Virtualization.
type SlbCurCfgSerAppShapeTable struct {
	// The number of the virtual server.
	SlbCurCfgSerAppShapeVirtServIndex int32
	// The service index. This has no external meaning
	SlbCurCfgSerAppShapeVirtServiceIndex int32
	// The number of the virtual server.
	SlbCurCfgSerAppShapePriority int32
	Params                       *SlbCurCfgSerAppShapeTableParams
}

func NewSlbCurCfgSerAppShapeTableList() *SlbCurCfgSerAppShapeTable {
	return &SlbCurCfgSerAppShapeTable{}
}

func NewSlbCurCfgSerAppShapeTable(
	slbCurCfgSerAppShapeVirtServIndex int32,
	slbCurCfgSerAppShapeVirtServiceIndex int32,
	slbCurCfgSerAppShapePriority int32,
	params *SlbCurCfgSerAppShapeTableParams,
) *SlbCurCfgSerAppShapeTable {
	return &SlbCurCfgSerAppShapeTable{
		SlbCurCfgSerAppShapeVirtServIndex:    slbCurCfgSerAppShapeVirtServIndex,
		SlbCurCfgSerAppShapeVirtServiceIndex: slbCurCfgSerAppShapeVirtServiceIndex,
		SlbCurCfgSerAppShapePriority:         slbCurCfgSerAppShapePriority,
		Params:                               params,
	}
}

func (c *SlbCurCfgSerAppShapeTable) Name() string {
	return "SlbCurCfgSerAppShapeTable"
}

func (c *SlbCurCfgSerAppShapeTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbCurCfgSerAppShapeTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbCurCfgSerAppShapeTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbCurCfgSerAppShapeVirtServIndex).IsZero() &&
		reflect.ValueOf(c.SlbCurCfgSerAppShapeVirtServiceIndex).IsZero() &&
		reflect.ValueOf(c.SlbCurCfgSerAppShapePriority).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbCurCfgSerAppShapeVirtServIndex) + "/" + fmt.Sprint(c.SlbCurCfgSerAppShapeVirtServiceIndex) + "/" + fmt.Sprint(c.SlbCurCfgSerAppShapePriority)
}

type SlbCurCfgSerAppShapeTableParams struct {
	// The number of the virtual server.
	VirtServIndex int32 `json:"VirtServIndex,omitempty"`
	// The service index. This has no external meaning
	VirtServiceIndex int32 `json:"VirtServiceIndex,omitempty"`
	// The number of the virtual server.
	Priority int32 `json:"Priority,omitempty"`
	// The AppShape ID added to the virtual service
	Index string `json:"Index,omitempty"`
}

func (p SlbCurCfgSerAppShapeTableParams) iMABean() {}
