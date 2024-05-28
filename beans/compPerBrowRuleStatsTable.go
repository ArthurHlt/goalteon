package beans

import (
	"fmt"
	"reflect"
)

// CompPerBrowRuleStatsTable A table for compression statistics per browser rule.
// Note:This mib is not supported for VX instance of Virtualization.
type CompPerBrowRuleStatsTable struct {
	// Compression browser rule rule-list index.
	CompBrowRuleRuleListIndex int32
	// Compression browser rule index.
	CompBrowRuleIndex int32
	Params            *CompPerBrowRuleStatsTableParams
}

func NewCompPerBrowRuleStatsTableList() *CompPerBrowRuleStatsTable {
	return &CompPerBrowRuleStatsTable{}
}

func NewCompPerBrowRuleStatsTable(
	compBrowRuleRuleListIndex int32,
	compBrowRuleIndex int32,
	params *CompPerBrowRuleStatsTableParams,
) *CompPerBrowRuleStatsTable {
	return &CompPerBrowRuleStatsTable{
		CompBrowRuleRuleListIndex: compBrowRuleRuleListIndex,
		CompBrowRuleIndex:         compBrowRuleIndex,
		Params:                    params,
	}
}

func (c *CompPerBrowRuleStatsTable) Name() string {
	return "CompPerBrowRuleStatsTable"
}

func (c *CompPerBrowRuleStatsTable) GetParams() BeanType {
	return c.Params
}

func (c *CompPerBrowRuleStatsTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *CompPerBrowRuleStatsTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.CompBrowRuleRuleListIndex).IsZero() &&
		reflect.ValueOf(c.CompBrowRuleIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.CompBrowRuleRuleListIndex) + "/" + fmt.Sprint(c.CompBrowRuleIndex)
}

type CompPerBrowRuleStatsTableParams struct {
	// Compression browser rule rule-list index.
	BrowRuleRuleListIndex int32 `json:"BrowRuleRuleListIndex,omitempty"`
	// Compression browser rule index.
	BrowRuleIndex int32 `json:"BrowRuleIndex,omitempty"`
	// Compression browser rule rule-list identifier.
	BrowRuleRuleListId string `json:"BrowRuleRuleListId,omitempty"`
	// Number of objects matched by this browser rule during measuring period.
	BrowRuleNumOfObj int32 `json:"BrowRuleNumOfObj,omitempty"`
	// Total size of all matched objects for this browser rule before compression.
	BrowRuleSizeBefComp int32 `json:"BrowRuleSizeBefComp,omitempty"`
	// Total size of all matched objects for this browser rule after compression.
	BrowRuleSizeAftComp int32 `json:"BrowRuleSizeAftComp,omitempty"`
	// Compression ratio per browser rule.
	BrowRuleCompRatio int32 `json:"BrowRuleCompRatio,omitempty"`
}

func (p CompPerBrowRuleStatsTableParams) iMABean() {}
