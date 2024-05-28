package beans

import (
	"fmt"
	"reflect"
)

// CachPerRuleStatsTable A table for cache statistics per rule.
// Note:This mib is not supported for VX instance of Virtualization.
type CachPerRuleStatsTable struct {
	// Cache rule rule-list index.
	CachRuleRuleListIndex int32
	// Cache rule number.
	CachRuleIndex int32
	Params        *CachPerRuleStatsTableParams
}

func NewCachPerRuleStatsTableList() *CachPerRuleStatsTable {
	return &CachPerRuleStatsTable{}
}

func NewCachPerRuleStatsTable(
	cachRuleRuleListIndex int32,
	cachRuleIndex int32,
	params *CachPerRuleStatsTableParams,
) *CachPerRuleStatsTable {
	return &CachPerRuleStatsTable{
		CachRuleRuleListIndex: cachRuleRuleListIndex,
		CachRuleIndex:         cachRuleIndex,
		Params:                params,
	}
}

func (c *CachPerRuleStatsTable) Name() string {
	return "CachPerRuleStatsTable"
}

func (c *CachPerRuleStatsTable) GetParams() BeanType {
	return c.Params
}

func (c *CachPerRuleStatsTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *CachPerRuleStatsTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.CachRuleRuleListIndex).IsZero() &&
		reflect.ValueOf(c.CachRuleIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.CachRuleRuleListIndex) + "/" + fmt.Sprint(c.CachRuleIndex)
}

type CachPerRuleStatsTableParams struct {
	// Cache rule rule-list index.
	RuleRuleListIndex int32 `json:"RuleRuleListIndex,omitempty"`
	// Cache rule number.
	RuleIndex int32 `json:"RuleIndex,omitempty"`
	// Cache rule rule-list identifier.
	RuleRuleListId string `json:"RuleRuleListId,omitempty"`
	// Number of objects cached by the rule.
	RuleNumOfObjCac int32 `json:"RuleNumOfObjCac,omitempty"`
	// Number of bytes cached by the rule.
	RuleNumOfBytesCac int32 `json:"RuleNumOfBytesCac,omitempty"`
}

func (p CachPerRuleStatsTableParams) iMABean() {}
