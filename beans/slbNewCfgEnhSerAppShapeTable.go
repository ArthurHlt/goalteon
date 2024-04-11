package beans

import (
	"fmt"
	"reflect"
)

// SlbNewCfgEnhSerAppShapeTable The table of Content based Services Rules.
// Note:This mib is not supported for VX instance of Virtualization.
type SlbNewCfgEnhSerAppShapeTable struct {
	// The number of the virtual server.
	SlbNewCfgEnhSerAppShapeVirtServIndex string
	// The service index. This has no external meaning
	SlbNewCfgEnhSerAppShapeVirtServiceIndex int32
	// The number of the virtual server.
	SlbNewCfgEnhSerAppShapePriority int32
	Params                          *SlbNewCfgEnhSerAppShapeTableParams
}

func NewSlbNewCfgEnhSerAppShapeTableList() *SlbNewCfgEnhSerAppShapeTable {
	return &SlbNewCfgEnhSerAppShapeTable{}
}

func NewSlbNewCfgEnhSerAppShapeTable(
	slbNewCfgEnhSerAppShapeVirtServIndex string,
	slbNewCfgEnhSerAppShapeVirtServiceIndex int32,
	slbNewCfgEnhSerAppShapePriority int32,
	params *SlbNewCfgEnhSerAppShapeTableParams,
) *SlbNewCfgEnhSerAppShapeTable {
	return &SlbNewCfgEnhSerAppShapeTable{
		SlbNewCfgEnhSerAppShapeVirtServIndex:    slbNewCfgEnhSerAppShapeVirtServIndex,
		SlbNewCfgEnhSerAppShapeVirtServiceIndex: slbNewCfgEnhSerAppShapeVirtServiceIndex,
		SlbNewCfgEnhSerAppShapePriority:         slbNewCfgEnhSerAppShapePriority,
		Params:                                  params,
	}
}

func (c *SlbNewCfgEnhSerAppShapeTable) Name() string {
	return "SlbNewCfgEnhSerAppShapeTable"
}

func (c *SlbNewCfgEnhSerAppShapeTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbNewCfgEnhSerAppShapeTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbNewCfgEnhSerAppShapeTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbNewCfgEnhSerAppShapeVirtServIndex).IsZero() &&
		reflect.ValueOf(c.SlbNewCfgEnhSerAppShapeVirtServiceIndex).IsZero() &&
		reflect.ValueOf(c.SlbNewCfgEnhSerAppShapePriority).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbNewCfgEnhSerAppShapeVirtServIndex) + "/" + fmt.Sprint(c.SlbNewCfgEnhSerAppShapeVirtServiceIndex) + "/" + fmt.Sprint(c.SlbNewCfgEnhSerAppShapePriority)
}

type SlbNewCfgEnhSerAppShapeTableDelete int32

const (
	SlbNewCfgEnhSerAppShapeTableDelete_Other       SlbNewCfgEnhSerAppShapeTableDelete = 1
	SlbNewCfgEnhSerAppShapeTableDelete_Delete      SlbNewCfgEnhSerAppShapeTableDelete = 2
	SlbNewCfgEnhSerAppShapeTableDelete_Unsupported SlbNewCfgEnhSerAppShapeTableDelete = 2147483647
)

type SlbNewCfgEnhSerAppShapeTableParams struct {
	// The number of the virtual server.
	VirtServIndex string `json:"VirtServIndex,omitempty"`
	// The service index. This has no external meaning
	VirtServiceIndex int32 `json:"VirtServiceIndex,omitempty"`
	// The number of the virtual server.
	Priority int32 `json:"Priority,omitempty"`
	// The AppShape ID added to the virtual service
	Index string `json:"Index,omitempty"`
	// When set to the value of 2 (delete), the entire row is deleted.
	// When read, other(1) is returned. Setting the value to anything
	// other than 2(delete) has no effect on the state of the row.
	Delete SlbNewCfgEnhSerAppShapeTableDelete `json:"Delete,omitempty"`
}

func (p SlbNewCfgEnhSerAppShapeTableParams) iMABean() {}
