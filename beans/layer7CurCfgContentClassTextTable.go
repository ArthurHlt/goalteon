package beans

import (
	"fmt"
	"reflect"
)

// Layer7CurCfgContentClassTextTable The table for configuring Content Class Text.
// Note:This mib is not supported for VX instance of virtualization.
type Layer7CurCfgContentClassTextTable struct {
	// The content Class Text ID(key id) as an index, length of the string should be 32 characters.
	Layer7CurCfgContentClassTextContentClassID string
	// The content Class ID(key id) as an index, length of the string should be 32 characters.
	Layer7CurCfgContentClassTextID string
	Params                         *Layer7CurCfgContentClassTextTableParams
}

func NewLayer7CurCfgContentClassTextTableList() *Layer7CurCfgContentClassTextTable {
	return &Layer7CurCfgContentClassTextTable{}
}

func NewLayer7CurCfgContentClassTextTable(
	layer7CurCfgContentClassTextContentClassID string,
	layer7CurCfgContentClassTextID string,
	params *Layer7CurCfgContentClassTextTableParams,
) *Layer7CurCfgContentClassTextTable {
	return &Layer7CurCfgContentClassTextTable{
		Layer7CurCfgContentClassTextContentClassID: layer7CurCfgContentClassTextContentClassID,
		Layer7CurCfgContentClassTextID:             layer7CurCfgContentClassTextID,
		Params:                                     params,
	}
}

func (c *Layer7CurCfgContentClassTextTable) Name() string {
	return "Layer7CurCfgContentClassTextTable"
}

func (c *Layer7CurCfgContentClassTextTable) GetParams() BeanType {
	return c.Params
}

func (c *Layer7CurCfgContentClassTextTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *Layer7CurCfgContentClassTextTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.Layer7CurCfgContentClassTextContentClassID).IsZero() &&
		reflect.ValueOf(c.Layer7CurCfgContentClassTextID).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.Layer7CurCfgContentClassTextContentClassID) + "/" + fmt.Sprint(c.Layer7CurCfgContentClassTextID)
}

type Layer7CurCfgContentClassTextTableMatchType int32

const (
	Layer7CurCfgContentClassTextTableMatchType_Include     Layer7CurCfgContentClassTextTableMatchType = 4
	Layer7CurCfgContentClassTextTableMatchType_Regex       Layer7CurCfgContentClassTextTableMatchType = 5
	Layer7CurCfgContentClassTextTableMatchType_Unsupported Layer7CurCfgContentClassTextTableMatchType = 2147483647
)

type Layer7CurCfgContentClassTextTableLookupArea int32

const (
	Layer7CurCfgContentClassTextTableLookupArea_Header      Layer7CurCfgContentClassTextTableLookupArea = 1
	Layer7CurCfgContentClassTextTableLookupArea_Body        Layer7CurCfgContentClassTextTableLookupArea = 2
	Layer7CurCfgContentClassTextTableLookupArea_Both        Layer7CurCfgContentClassTextTableLookupArea = 3
	Layer7CurCfgContentClassTextTableLookupArea_Unsupported Layer7CurCfgContentClassTextTableLookupArea = 2147483647
)

type Layer7CurCfgContentClassTextTableCase int32

const (
	Layer7CurCfgContentClassTextTableCase_Enabled     Layer7CurCfgContentClassTextTableCase = 1
	Layer7CurCfgContentClassTextTableCase_Disabled    Layer7CurCfgContentClassTextTableCase = 2
	Layer7CurCfgContentClassTextTableCase_Unsupported Layer7CurCfgContentClassTextTableCase = 2147483647
)

type Layer7CurCfgContentClassTextTableParams struct {
	// The content Class Text ID(key id) as an index, length of the string should be 32 characters.
	ContentClassID string `json:"ContentClassID,omitempty"`
	// The content Class ID(key id) as an index, length of the string should be 32 characters.
	ID string `json:"ID,omitempty"`
	// Content Class Text to match, length of the string should be 32 characters.
	Text string `json:"Text,omitempty"`
	// Text Match type.
	MatchType Layer7CurCfgContentClassTextTableMatchType `json:"MatchType,omitempty"`
	// Text Match type.
	LookupArea Layer7CurCfgContentClassTextTableLookupArea `json:"LookupArea,omitempty"`
	// Enable or Disable Case sensitive for String matching.
	Case Layer7CurCfgContentClassTextTableCase `json:"Case,omitempty"`
}

func (p Layer7CurCfgContentClassTextTableParams) iMABean() {}
