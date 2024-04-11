package beans

import (
	"fmt"
	"reflect"
)

// SlbCurCfgSidebandAppShapeTable The table of App Rules added to virtual services.
// Note:This mib is not supported for VX instance of Virtualization.
type SlbCurCfgSidebandAppShapeTable struct {
	// The sideband index. This has no external meaning
	SlbCurCfgSidebandAppShapeSidebandIndex string
	// The priority of the AppShape.
	SlbCurCfgSidebandAppShapePriority int32
	Params                            *SlbCurCfgSidebandAppShapeTableParams
}

func NewSlbCurCfgSidebandAppShapeTableList() *SlbCurCfgSidebandAppShapeTable {
	return &SlbCurCfgSidebandAppShapeTable{}
}

func NewSlbCurCfgSidebandAppShapeTable(
	slbCurCfgSidebandAppShapeSidebandIndex string,
	slbCurCfgSidebandAppShapePriority int32,
	params *SlbCurCfgSidebandAppShapeTableParams,
) *SlbCurCfgSidebandAppShapeTable {
	return &SlbCurCfgSidebandAppShapeTable{
		SlbCurCfgSidebandAppShapeSidebandIndex: slbCurCfgSidebandAppShapeSidebandIndex,
		SlbCurCfgSidebandAppShapePriority:      slbCurCfgSidebandAppShapePriority,
		Params:                                 params,
	}
}

func (c *SlbCurCfgSidebandAppShapeTable) Name() string {
	return "SlbCurCfgSidebandAppShapeTable"
}

func (c *SlbCurCfgSidebandAppShapeTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbCurCfgSidebandAppShapeTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbCurCfgSidebandAppShapeTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbCurCfgSidebandAppShapeSidebandIndex).IsZero() &&
		reflect.ValueOf(c.SlbCurCfgSidebandAppShapePriority).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbCurCfgSidebandAppShapeSidebandIndex) + "/" + fmt.Sprint(c.SlbCurCfgSidebandAppShapePriority)
}

type SlbCurCfgSidebandAppShapeTableParams struct {
	// The sideband index. This has no external meaning
	SidebandIndex string `json:"SidebandIndex,omitempty"`
	// The priority of the AppShape.
	Priority int32 `json:"Priority,omitempty"`
	// The AppShape ID added to the sideband policy
	Index string `json:"Index,omitempty"`
}

func (p SlbCurCfgSidebandAppShapeTableParams) iMABean() {}
