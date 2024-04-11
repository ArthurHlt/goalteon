package beans

import (
	"fmt"
	"reflect"
)

// SlbNewCfgSerAppShapeTable The table of Content based Services Rules.
// Note:This mib is not supported for VX instance of Virtualization.
type SlbNewCfgSerAppShapeTable struct {
	// The number of the virtual server.
	SlbNewCfgSerAppShapeVirtServIndex int32
	// The service index. This has no external meaning
	SlbNewCfgSerAppShapeVirtServiceIndex int32
	// The number of the virtual server.
	SlbNewCfgSerAppShapePriority int32
	Params                       *SlbNewCfgSerAppShapeTableParams
}

func NewSlbNewCfgSerAppShapeTableList() *SlbNewCfgSerAppShapeTable {
	return &SlbNewCfgSerAppShapeTable{}
}

func NewSlbNewCfgSerAppShapeTable(
	slbNewCfgSerAppShapeVirtServIndex int32,
	slbNewCfgSerAppShapeVirtServiceIndex int32,
	slbNewCfgSerAppShapePriority int32,
	params *SlbNewCfgSerAppShapeTableParams,
) *SlbNewCfgSerAppShapeTable {
	return &SlbNewCfgSerAppShapeTable{
		SlbNewCfgSerAppShapeVirtServIndex:    slbNewCfgSerAppShapeVirtServIndex,
		SlbNewCfgSerAppShapeVirtServiceIndex: slbNewCfgSerAppShapeVirtServiceIndex,
		SlbNewCfgSerAppShapePriority:         slbNewCfgSerAppShapePriority,
		Params:                               params,
	}
}

func (c *SlbNewCfgSerAppShapeTable) Name() string {
	return "SlbNewCfgSerAppShapeTable"
}

func (c *SlbNewCfgSerAppShapeTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbNewCfgSerAppShapeTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbNewCfgSerAppShapeTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbNewCfgSerAppShapeVirtServIndex).IsZero() &&
		reflect.ValueOf(c.SlbNewCfgSerAppShapeVirtServiceIndex).IsZero() &&
		reflect.ValueOf(c.SlbNewCfgSerAppShapePriority).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbNewCfgSerAppShapeVirtServIndex) + "/" + fmt.Sprint(c.SlbNewCfgSerAppShapeVirtServiceIndex) + "/" + fmt.Sprint(c.SlbNewCfgSerAppShapePriority)
}

type SlbNewCfgSerAppShapeTableDelete int32

const (
	SlbNewCfgSerAppShapeTableDelete_Other       SlbNewCfgSerAppShapeTableDelete = 1
	SlbNewCfgSerAppShapeTableDelete_Delete      SlbNewCfgSerAppShapeTableDelete = 2
	SlbNewCfgSerAppShapeTableDelete_Unsupported SlbNewCfgSerAppShapeTableDelete = 2147483647
)

type SlbNewCfgSerAppShapeTableParams struct {
	// The number of the virtual server.
	VirtServIndex int32 `json:"VirtServIndex,omitempty"`
	// The service index. This has no external meaning
	VirtServiceIndex int32 `json:"VirtServiceIndex,omitempty"`
	// The number of the virtual server.
	Priority int32 `json:"Priority,omitempty"`
	// The AppShape ID added to the virtual service
	Index string `json:"Index,omitempty"`
	// When set to the value of 2 (delete), the entire row is deleted.
	// When read, other(1) is returned. Setting the value to anything
	// other than 2(delete) has no effect on the state of the row.
	Delete SlbNewCfgSerAppShapeTableDelete `json:"Delete,omitempty"`
}

func (p SlbNewCfgSerAppShapeTableParams) iMABean() {}
