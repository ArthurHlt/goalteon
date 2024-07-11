package beans

import (
	"fmt"
	"reflect"
)

// UrlSpMaintStatsTable The table of URL SP maintenance statistics.
// Note:This mib is not supported for VX instance of virtualization.
type UrlSpMaintStatsTable struct {
	// The SP index.
	UrlSpMaintStatsSpIndex int32
	Params                 *UrlSpMaintStatsTableParams
}

func NewUrlSpMaintStatsTableList() *UrlSpMaintStatsTable {
	return &UrlSpMaintStatsTable{}
}

func NewUrlSpMaintStatsTable(
	urlSpMaintStatsSpIndex int32,
	params *UrlSpMaintStatsTableParams,
) *UrlSpMaintStatsTable {
	return &UrlSpMaintStatsTable{
		UrlSpMaintStatsSpIndex: urlSpMaintStatsSpIndex,
		Params:                 params,
	}
}

func (c *UrlSpMaintStatsTable) Name() string {
	return "UrlSpMaintStatsTable"
}

func (c *UrlSpMaintStatsTable) GetParams() BeanType {
	return c.Params
}

func (c *UrlSpMaintStatsTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *UrlSpMaintStatsTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.UrlSpMaintStatsSpIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.UrlSpMaintStatsSpIndex)
}

type UrlSpMaintStatsTableParams struct {
	// The SP index.
	SpIndex int32 `json:"SpIndex,omitempty"`
	// The number of memory units available.
	CurMemUnits uint32 `json:"CurMemUnits,omitempty"`
	// The lowest number of memory units available.
	LowestMemUnits uint32 `json:"LowestMemUnits,omitempty"`
}

func (p UrlSpMaintStatsTableParams) iMABean() {}
