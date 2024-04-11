package beans

import (
	"fmt"
	"reflect"
)

// SlbCurCfgFiltAppShapeTable The table of App Rules added to filters.
// Note:This mib is not supported for VX instance of Virtualization.
type SlbCurCfgFiltAppShapeTable struct {
	// The number of the filter.
	SlbCurCfgFiltAppShapeFiltIndex int32
	// The AppShape priority.
	SlbCurCfgFiltAppShapePriority int32
	Params                        *SlbCurCfgFiltAppShapeTableParams
}

func NewSlbCurCfgFiltAppShapeTableList() *SlbCurCfgFiltAppShapeTable {
	return &SlbCurCfgFiltAppShapeTable{}
}

func NewSlbCurCfgFiltAppShapeTable(
	slbCurCfgFiltAppShapeFiltIndex int32,
	slbCurCfgFiltAppShapePriority int32,
	params *SlbCurCfgFiltAppShapeTableParams,
) *SlbCurCfgFiltAppShapeTable {
	return &SlbCurCfgFiltAppShapeTable{
		SlbCurCfgFiltAppShapeFiltIndex: slbCurCfgFiltAppShapeFiltIndex,
		SlbCurCfgFiltAppShapePriority:  slbCurCfgFiltAppShapePriority,
		Params:                         params,
	}
}

func (c *SlbCurCfgFiltAppShapeTable) Name() string {
	return "SlbCurCfgFiltAppShapeTable"
}

func (c *SlbCurCfgFiltAppShapeTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbCurCfgFiltAppShapeTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbCurCfgFiltAppShapeTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbCurCfgFiltAppShapeFiltIndex).IsZero() &&
		reflect.ValueOf(c.SlbCurCfgFiltAppShapePriority).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbCurCfgFiltAppShapeFiltIndex) + "/" + fmt.Sprint(c.SlbCurCfgFiltAppShapePriority)
}

type SlbCurCfgFiltAppShapeTableParams struct {
	// The number of the filter.
	FiltIndex int32 `json:"FiltIndex,omitempty"`
	// The AppShape priority.
	Priority int32 `json:"Priority,omitempty"`
	// The AppShape ID added to the filter
	Index string `json:"Index,omitempty"`
}

func (p SlbCurCfgFiltAppShapeTableParams) iMABean() {}
