package beans

import (
	"fmt"
	"reflect"
)

// SlbCurCfgEnhSerAppShapeTable The table of App Rules added to virtual services.
// Note:This mib is not supported for VX instance of Virtualization.
type SlbCurCfgEnhSerAppShapeTable struct {
	// The number of the virtual server.
	SlbCurCfgEnhSerAppShapeVirtServIndex string
	// The service index. This has no external meaning
	SlbCurCfgEnhSerAppShapeVirtServiceIndex int32
	// The number of the virtual server.
	SlbCurCfgEnhSerAppShapePriority int32
	Params                          *SlbCurCfgEnhSerAppShapeTableParams
}

func NewSlbCurCfgEnhSerAppShapeTableList() *SlbCurCfgEnhSerAppShapeTable {
	return &SlbCurCfgEnhSerAppShapeTable{}
}

func NewSlbCurCfgEnhSerAppShapeTable(
	slbCurCfgEnhSerAppShapeVirtServIndex string,
	slbCurCfgEnhSerAppShapeVirtServiceIndex int32,
	slbCurCfgEnhSerAppShapePriority int32,
	params *SlbCurCfgEnhSerAppShapeTableParams,
) *SlbCurCfgEnhSerAppShapeTable {
	return &SlbCurCfgEnhSerAppShapeTable{
		SlbCurCfgEnhSerAppShapeVirtServIndex:    slbCurCfgEnhSerAppShapeVirtServIndex,
		SlbCurCfgEnhSerAppShapeVirtServiceIndex: slbCurCfgEnhSerAppShapeVirtServiceIndex,
		SlbCurCfgEnhSerAppShapePriority:         slbCurCfgEnhSerAppShapePriority,
		Params:                                  params,
	}
}

func (c *SlbCurCfgEnhSerAppShapeTable) Name() string {
	return "SlbCurCfgEnhSerAppShapeTable"
}

func (c *SlbCurCfgEnhSerAppShapeTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbCurCfgEnhSerAppShapeTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbCurCfgEnhSerAppShapeTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbCurCfgEnhSerAppShapeVirtServIndex).IsZero() &&
		reflect.ValueOf(c.SlbCurCfgEnhSerAppShapeVirtServiceIndex).IsZero() &&
		reflect.ValueOf(c.SlbCurCfgEnhSerAppShapePriority).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbCurCfgEnhSerAppShapeVirtServIndex) + "/" + fmt.Sprint(c.SlbCurCfgEnhSerAppShapeVirtServiceIndex) + "/" + fmt.Sprint(c.SlbCurCfgEnhSerAppShapePriority)
}

type SlbCurCfgEnhSerAppShapeTableParams struct {
	// The number of the virtual server.
	VirtServIndex string `json:"VirtServIndex,omitempty"`
	// The service index. This has no external meaning
	VirtServiceIndex int32 `json:"VirtServiceIndex,omitempty"`
	// The number of the virtual server.
	Priority int32 `json:"Priority,omitempty"`
	// The AppShape ID added to the virtual service
	Index string `json:"Index,omitempty"`
}

func (p SlbCurCfgEnhSerAppShapeTableParams) iMABean() {}
