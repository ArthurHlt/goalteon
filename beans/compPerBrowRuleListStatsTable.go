package beans

import (
	"fmt"
	"reflect"
)

// CompPerBrowRuleListStatsTable A table for compression statistics per rule list.
// Note:This mib is not supported for VX instance of Virtualization.
type CompPerBrowRuleListStatsTable struct {
	// Compression browsing rule list index.
	CompBrowRuleListIndex int32
	Params                *CompPerBrowRuleListStatsTableParams
}

func NewCompPerBrowRuleListStatsTableList() *CompPerBrowRuleListStatsTable {
	return &CompPerBrowRuleListStatsTable{}
}

func NewCompPerBrowRuleListStatsTable(
	compBrowRuleListIndex int32,
	params *CompPerBrowRuleListStatsTableParams,
) *CompPerBrowRuleListStatsTable {
	return &CompPerBrowRuleListStatsTable{
		CompBrowRuleListIndex: compBrowRuleListIndex,
		Params:                params,
	}
}

func (c *CompPerBrowRuleListStatsTable) Name() string {
	return "CompPerBrowRuleListStatsTable"
}

func (c *CompPerBrowRuleListStatsTable) GetParams() BeanType {
	return c.Params
}

func (c *CompPerBrowRuleListStatsTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *CompPerBrowRuleListStatsTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.CompBrowRuleListIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.CompBrowRuleListIndex)
}

type CompPerBrowRuleListStatsTableParams struct {
	// Compression browsing rule list index.
	BrowRuleListIndex int32 `json:"BrowRuleListIndex,omitempty"`
	// Compression browsing rule list identifier.
	BrowRuleListId string `json:"BrowRuleListId,omitempty"`
	// Number of objects matched by this rule-list during measuring period.
	BrowRuleListNumOfObj int32 `json:"BrowRuleListNumOfObj,omitempty"`
	// Total size of all matched objects by this browser rule-list before compression.
	BrowRuleListSizeBefComp int32 `json:"BrowRuleListSizeBefComp,omitempty"`
	// Total size of all matched objects by this browser rule-list after compression.
	BrowRuleListSizeAftComp int32 `json:"BrowRuleListSizeAftComp,omitempty"`
	// Compression ratio per rule list.
	BrowRuleListCompRatio int32 `json:"BrowRuleListCompRatio,omitempty"`
}

func (p CompPerBrowRuleListStatsTableParams) iMABean() {}
