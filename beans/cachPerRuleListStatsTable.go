package beans

import (
	"fmt"
	"reflect"
)

// CachPerRuleListStatsTable A table for cache statistics per rule list.
// Note:This mib is not supported for VX instance of Virtualization.
type CachPerRuleListStatsTable struct {
	// Cache rule list identifier.
	CachRuleListId string
	Params         *CachPerRuleListStatsTableParams
}

func NewCachPerRuleListStatsTableList() *CachPerRuleListStatsTable {
	return &CachPerRuleListStatsTable{}
}

func NewCachPerRuleListStatsTable(
	cachRuleListId string,
	params *CachPerRuleListStatsTableParams,
) *CachPerRuleListStatsTable {
	return &CachPerRuleListStatsTable{
		CachRuleListId: cachRuleListId,
		Params:         params,
	}
}

func (c *CachPerRuleListStatsTable) Name() string {
	return "CachPerRuleListStatsTable"
}

func (c *CachPerRuleListStatsTable) GetParams() BeanType {
	return c.Params
}

func (c *CachPerRuleListStatsTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *CachPerRuleListStatsTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.CachRuleListId).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.CachRuleListId)
}

type CachPerRuleListStatsTableParams struct {
	// Cache rule list index.
	RuleListIndex int32 `json:"RuleListIndex,omitempty"`
	// Cache rule list identifier.
	RuleListId string `json:"RuleListId,omitempty"`
	// Number of objects cached by this rule list.
	RuleListNumOfObjCac int32 `json:"RuleListNumOfObjCac,omitempty"`
	// Number of bytes cached by this rule list.
	RuleListNumOfBytesCac int32 `json:"RuleListNumOfBytesCac,omitempty"`
}

func (p CachPerRuleListStatsTableParams) iMABean() {}
