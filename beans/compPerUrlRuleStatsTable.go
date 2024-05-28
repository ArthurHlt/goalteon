package beans

import (
	"fmt"
	"reflect"
)

// CompPerUrlRuleStatsTable A table for compression statistics per URL Rule.
// Note:This mib is not supported for VX instance of Virtualization.
type CompPerUrlRuleStatsTable struct {
	// Compression url rule rule-list index.
	CompUrlRuleRuleListIndex int32
	// Compression url rule index.
	CompUrlRuleIndex int32
	Params           *CompPerUrlRuleStatsTableParams
}

func NewCompPerUrlRuleStatsTableList() *CompPerUrlRuleStatsTable {
	return &CompPerUrlRuleStatsTable{}
}

func NewCompPerUrlRuleStatsTable(
	compUrlRuleRuleListIndex int32,
	compUrlRuleIndex int32,
	params *CompPerUrlRuleStatsTableParams,
) *CompPerUrlRuleStatsTable {
	return &CompPerUrlRuleStatsTable{
		CompUrlRuleRuleListIndex: compUrlRuleRuleListIndex,
		CompUrlRuleIndex:         compUrlRuleIndex,
		Params:                   params,
	}
}

func (c *CompPerUrlRuleStatsTable) Name() string {
	return "CompPerUrlRuleStatsTable"
}

func (c *CompPerUrlRuleStatsTable) GetParams() BeanType {
	return c.Params
}

func (c *CompPerUrlRuleStatsTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *CompPerUrlRuleStatsTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.CompUrlRuleRuleListIndex).IsZero() &&
		reflect.ValueOf(c.CompUrlRuleIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.CompUrlRuleRuleListIndex) + "/" + fmt.Sprint(c.CompUrlRuleIndex)
}

type CompPerUrlRuleStatsTableParams struct {
	// Compression url rule rule-list index.
	UrlRuleRuleListIndex int32 `json:"UrlRuleRuleListIndex,omitempty"`
	// Compression url rule index.
	UrlRuleIndex int32 `json:"UrlRuleIndex,omitempty"`
	// Compression url rule rule-list identifier.
	UrlRuleRuleListId string `json:"UrlRuleRuleListId,omitempty"`
	// Number of objects matched by this url rule during measuring period.
	UrlRuleNumOfObj int32 `json:"UrlRuleNumOfObj,omitempty"`
	// Total size of all matched objects for this url rule before compression.
	UrlRuleSizeBefComp int32 `json:"UrlRuleSizeBefComp,omitempty"`
	// Total size of all matched objects for this url rule after compression.
	UrlRuleSizeAftComp int32 `json:"UrlRuleSizeAftComp,omitempty"`
	// Compression ratio for this url rule.
	UrlRuleCompRatio int32 `json:"UrlRuleCompRatio,omitempty"`
}

func (p CompPerUrlRuleStatsTableParams) iMABean() {}
