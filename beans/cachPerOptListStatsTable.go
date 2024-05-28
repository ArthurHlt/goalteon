package beans

import (
	"fmt"
	"reflect"
)

// CachPerOptListStatsTable A table for cache statistics per optimization rule list.
// Note:This mib is not supported for VX instance of Virtualization.
type CachPerOptListStatsTable struct {
	// Cache opt rule list identifier.
	CachOptListId string
	Params        *CachPerOptListStatsTableParams
}

func NewCachPerOptListStatsTableList() *CachPerOptListStatsTable {
	return &CachPerOptListStatsTable{}
}

func NewCachPerOptListStatsTable(
	cachOptListId string,
	params *CachPerOptListStatsTableParams,
) *CachPerOptListStatsTable {
	return &CachPerOptListStatsTable{
		CachOptListId: cachOptListId,
		Params:        params,
	}
}

func (c *CachPerOptListStatsTable) Name() string {
	return "CachPerOptListStatsTable"
}

func (c *CachPerOptListStatsTable) GetParams() BeanType {
	return c.Params
}

func (c *CachPerOptListStatsTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *CachPerOptListStatsTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.CachOptListId).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.CachOptListId)
}

type CachPerOptListStatsTableParams struct {
	// Cache opt rule list index.
	OptListIndex int32 `json:"OptListIndex,omitempty"`
	// Cache opt rule list identifier.
	OptListId string `json:"OptListId,omitempty"`
	// Number of hits on this rule list.
	OptListNumOfHits int32 `json:"OptListNumOfHits,omitempty"`
}

func (p CachPerOptListStatsTableParams) iMABean() {}
