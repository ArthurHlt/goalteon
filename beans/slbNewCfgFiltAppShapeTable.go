package beans

import (
	"fmt"
	"reflect"
)

// SlbNewCfgFiltAppShapeTable The table of AppShape based Filter Rules.
// Note:This mib is not supported for VX instance of Virtualization.
type SlbNewCfgFiltAppShapeTable struct {
	// The number of the filter.
	SlbNewCfgFiltAppShapeFiltIndex int32
	// The AppShape priority.
	SlbNewCfgFiltAppShapePriority int32
	Params                        *SlbNewCfgFiltAppShapeTableParams
}

func NewSlbNewCfgFiltAppShapeTableList() *SlbNewCfgFiltAppShapeTable {
	return &SlbNewCfgFiltAppShapeTable{}
}

func NewSlbNewCfgFiltAppShapeTable(
	slbNewCfgFiltAppShapeFiltIndex int32,
	slbNewCfgFiltAppShapePriority int32,
	params *SlbNewCfgFiltAppShapeTableParams,
) *SlbNewCfgFiltAppShapeTable {
	return &SlbNewCfgFiltAppShapeTable{
		SlbNewCfgFiltAppShapeFiltIndex: slbNewCfgFiltAppShapeFiltIndex,
		SlbNewCfgFiltAppShapePriority:  slbNewCfgFiltAppShapePriority,
		Params:                         params,
	}
}

func (c *SlbNewCfgFiltAppShapeTable) Name() string {
	return "SlbNewCfgFiltAppShapeTable"
}

func (c *SlbNewCfgFiltAppShapeTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbNewCfgFiltAppShapeTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbNewCfgFiltAppShapeTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbNewCfgFiltAppShapeFiltIndex).IsZero() &&
		reflect.ValueOf(c.SlbNewCfgFiltAppShapePriority).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbNewCfgFiltAppShapeFiltIndex) + "/" + fmt.Sprint(c.SlbNewCfgFiltAppShapePriority)
}

type SlbNewCfgFiltAppShapeTableDelete int32

const (
	SlbNewCfgFiltAppShapeTableDelete_Other       SlbNewCfgFiltAppShapeTableDelete = 1
	SlbNewCfgFiltAppShapeTableDelete_Delete      SlbNewCfgFiltAppShapeTableDelete = 2
	SlbNewCfgFiltAppShapeTableDelete_Unsupported SlbNewCfgFiltAppShapeTableDelete = 2147483647
)

type SlbNewCfgFiltAppShapeTableParams struct {
	// The number of the filter.
	FiltIndex int32 `json:"FiltIndex,omitempty"`
	// The AppShape priority.
	Priority int32 `json:"Priority,omitempty"`
	// The AppShape ID added to the filter
	Index string `json:"Index,omitempty"`
	// When set to the value of 2 (delete), the entire row is deleted.
	// When read, other(1) is returned. Setting the value to anything
	// other than 2(delete) has no effect on the state of the row.
	Delete SlbNewCfgFiltAppShapeTableDelete `json:"Delete,omitempty"`
}

func (p SlbNewCfgFiltAppShapeTableParams) iMABean() {}
