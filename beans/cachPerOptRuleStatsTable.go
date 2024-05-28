package beans

import (
	"fmt"
	"reflect"
)

// CachPerOptRuleStatsTable A table for optimization statistics per rule.
// Note:This mib is not supported for VX instance of Virtualization.
type CachPerOptRuleStatsTable struct {
	// Optimization rule rule-list index.
	CachOptRuleListIndex int32
	// Optimization rule number.
	CachOptRuleIndex int32
	Params           *CachPerOptRuleStatsTableParams
}

func NewCachPerOptRuleStatsTableList() *CachPerOptRuleStatsTable {
	return &CachPerOptRuleStatsTable{}
}

func NewCachPerOptRuleStatsTable(
	cachOptRuleListIndex int32,
	cachOptRuleIndex int32,
	params *CachPerOptRuleStatsTableParams,
) *CachPerOptRuleStatsTable {
	return &CachPerOptRuleStatsTable{
		CachOptRuleListIndex: cachOptRuleListIndex,
		CachOptRuleIndex:     cachOptRuleIndex,
		Params:               params,
	}
}

func (c *CachPerOptRuleStatsTable) Name() string {
	return "CachPerOptRuleStatsTable"
}

func (c *CachPerOptRuleStatsTable) GetParams() BeanType {
	return c.Params
}

func (c *CachPerOptRuleStatsTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *CachPerOptRuleStatsTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.CachOptRuleListIndex).IsZero() &&
		reflect.ValueOf(c.CachOptRuleIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.CachOptRuleListIndex) + "/" + fmt.Sprint(c.CachOptRuleIndex)
}

type CachPerOptRuleStatsTableParams struct {
	// Optimization rule rule-list index.
	OptRuleListIndex int32 `json:"OptRuleListIndex,omitempty"`
	// Optimization rule number.
	OptRuleIndex int32 `json:"OptRuleIndex,omitempty"`
	// Optimization rule rule-list identifier.
	OptRuleListId string `json:"OptRuleListId,omitempty"`
	// Number of hits on the rule.
	OptRuleNumOfHits int32 `json:"OptRuleNumOfHits,omitempty"`
}

func (p CachPerOptRuleStatsTableParams) iMABean() {}
