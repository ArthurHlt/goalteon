package beans

import (
	"fmt"
	"reflect"
)

// CompPerUrlRuleListStatsTable A table for compression statistics per URL Rule list.
// Note:This mib is not supported for VX instance of Virtualization.
type CompPerUrlRuleListStatsTable struct {
	// Compression url rule list index.
	CompUrlRuleListIndex int32
	Params               *CompPerUrlRuleListStatsTableParams
}

func NewCompPerUrlRuleListStatsTableList() *CompPerUrlRuleListStatsTable {
	return &CompPerUrlRuleListStatsTable{}
}

func NewCompPerUrlRuleListStatsTable(
	compUrlRuleListIndex int32,
	params *CompPerUrlRuleListStatsTableParams,
) *CompPerUrlRuleListStatsTable {
	return &CompPerUrlRuleListStatsTable{
		CompUrlRuleListIndex: compUrlRuleListIndex,
		Params:               params,
	}
}

func (c *CompPerUrlRuleListStatsTable) Name() string {
	return "CompPerUrlRuleListStatsTable"
}

func (c *CompPerUrlRuleListStatsTable) GetParams() BeanType {
	return c.Params
}

func (c *CompPerUrlRuleListStatsTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *CompPerUrlRuleListStatsTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.CompUrlRuleListIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.CompUrlRuleListIndex)
}

type CompPerUrlRuleListStatsTableParams struct {
	// Compression url rule list index.
	UrlRuleListIndex int32 `json:"UrlRuleListIndex,omitempty"`
	// Compression url rule list identifier.
	UrlRuleListId string `json:"UrlRuleListId,omitempty"`
	// Number of objects matched by this url rule list during measuring period.
	UrlRuleListNumOfObj int32 `json:"UrlRuleListNumOfObj,omitempty"`
	// Total size of all matched objects for this url rule list before compression.
	UrlRuleListSizeBefComp int32 `json:"UrlRuleListSizeBefComp,omitempty"`
	// Total size of all matched objects this url rule list after compression.
	UrlRuleListSizeAftComp int32 `json:"UrlRuleListSizeAftComp,omitempty"`
	// Compression ratio for this url rule list.
	UrlRuleListCompRatio int32 `json:"UrlRuleListCompRatio,omitempty"`
}

func (p CompPerUrlRuleListStatsTableParams) iMABean() {}
