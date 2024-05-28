package beans

import (
	"fmt"
	"reflect"
)

// SlbOverlayStatsMapTable A table for slb overlay mapping statistics .
// Note:This mib is not supported for VX instance of Virtualization.
type SlbOverlayStatsMapTable struct {
	// slb overlay mapping stats table index.
	SlbOverlayStatsMapIndex string
	Params                  *SlbOverlayStatsMapTableParams
}

func NewSlbOverlayStatsMapTableList() *SlbOverlayStatsMapTable {
	return &SlbOverlayStatsMapTable{}
}

func NewSlbOverlayStatsMapTable(
	slbOverlayStatsMapIndex string,
	params *SlbOverlayStatsMapTableParams,
) *SlbOverlayStatsMapTable {
	return &SlbOverlayStatsMapTable{
		SlbOverlayStatsMapIndex: slbOverlayStatsMapIndex,
		Params:                  params,
	}
}

func (c *SlbOverlayStatsMapTable) Name() string {
	return "SlbOverlayStatsMapTable"
}

func (c *SlbOverlayStatsMapTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbOverlayStatsMapTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbOverlayStatsMapTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbOverlayStatsMapIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbOverlayStatsMapIndex)
}

type SlbOverlayStatsMapTableParams struct {
	// slb overlay mapping stats table index.
	Index string `json:"Index,omitempty"`
	// Overlay mapping table hits per entry.
	Hits uint64 `json:"Hits,omitempty"`
}

func (p SlbOverlayStatsMapTableParams) iMABean() {}
