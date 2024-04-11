package beans

import (
	"fmt"
	"reflect"
)

// SlbNewCfgSidebandAppShapeTable The table of Content based Services Rules.
// Note:This mib is not supported for VX instance of Virtualization.
type SlbNewCfgSidebandAppShapeTable struct {
	// The sideband index. This has no external meaning
	SlbNewCfgSidebandAppShapeSidebandIndex string
	// The priority of the AppShape.
	SlbNewCfgSidebandAppShapePriority int32
	Params                            *SlbNewCfgSidebandAppShapeTableParams
}

func NewSlbNewCfgSidebandAppShapeTableList() *SlbNewCfgSidebandAppShapeTable {
	return &SlbNewCfgSidebandAppShapeTable{}
}

func NewSlbNewCfgSidebandAppShapeTable(
	slbNewCfgSidebandAppShapeSidebandIndex string,
	slbNewCfgSidebandAppShapePriority int32,
	params *SlbNewCfgSidebandAppShapeTableParams,
) *SlbNewCfgSidebandAppShapeTable {
	return &SlbNewCfgSidebandAppShapeTable{
		SlbNewCfgSidebandAppShapeSidebandIndex: slbNewCfgSidebandAppShapeSidebandIndex,
		SlbNewCfgSidebandAppShapePriority:      slbNewCfgSidebandAppShapePriority,
		Params:                                 params,
	}
}

func (c *SlbNewCfgSidebandAppShapeTable) Name() string {
	return "SlbNewCfgSidebandAppShapeTable"
}

func (c *SlbNewCfgSidebandAppShapeTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbNewCfgSidebandAppShapeTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbNewCfgSidebandAppShapeTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbNewCfgSidebandAppShapeSidebandIndex).IsZero() &&
		reflect.ValueOf(c.SlbNewCfgSidebandAppShapePriority).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbNewCfgSidebandAppShapeSidebandIndex) + "/" + fmt.Sprint(c.SlbNewCfgSidebandAppShapePriority)
}

type SlbNewCfgSidebandAppShapeTableDelete int32

const (
	SlbNewCfgSidebandAppShapeTableDelete_Other       SlbNewCfgSidebandAppShapeTableDelete = 1
	SlbNewCfgSidebandAppShapeTableDelete_Delete      SlbNewCfgSidebandAppShapeTableDelete = 2
	SlbNewCfgSidebandAppShapeTableDelete_Unsupported SlbNewCfgSidebandAppShapeTableDelete = 2147483647
)

type SlbNewCfgSidebandAppShapeTableParams struct {
	// The sideband index. This has no external meaning
	SidebandIndex string `json:"SidebandIndex,omitempty"`
	// The priority of the AppShape.
	Priority int32 `json:"Priority,omitempty"`
	// The AppShape ID added to the sideband policy
	Index string `json:"Index,omitempty"`
	// When set to the value of 2 (delete), the entire row is deleted.
	// When read, other(1) is returned. Setting the value to anything
	// other than 2(delete) has no effect on the state of the row.
	Delete SlbNewCfgSidebandAppShapeTableDelete `json:"Delete,omitempty"`
}

func (p SlbNewCfgSidebandAppShapeTableParams) iMABean() {}
