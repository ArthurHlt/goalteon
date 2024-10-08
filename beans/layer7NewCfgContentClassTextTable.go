package beans

import (
	"fmt"
	"reflect"
)

// Layer7NewCfgContentClassTextTable The table for configuring Content Class Text.
// Note:This mib is not supported for VX instance of virtualization.
type Layer7NewCfgContentClassTextTable struct {
	// The content Class Text ID(key id) as an index, length of the string should be 32 characters.
	Layer7NewCfgContentClassTextContentClassID string
	// The content Class ID(key id) as an index, length of the string should be 32 characters.
	Layer7NewCfgContentClassTextID string
	Params                         *Layer7NewCfgContentClassTextTableParams
}

func NewLayer7NewCfgContentClassTextTableList() *Layer7NewCfgContentClassTextTable {
	return &Layer7NewCfgContentClassTextTable{}
}

func NewLayer7NewCfgContentClassTextTable(
	layer7NewCfgContentClassTextContentClassID string,
	layer7NewCfgContentClassTextID string,
	params *Layer7NewCfgContentClassTextTableParams,
) *Layer7NewCfgContentClassTextTable {
	return &Layer7NewCfgContentClassTextTable{
		Layer7NewCfgContentClassTextContentClassID: layer7NewCfgContentClassTextContentClassID,
		Layer7NewCfgContentClassTextID:             layer7NewCfgContentClassTextID,
		Params:                                     params,
	}
}

func (c *Layer7NewCfgContentClassTextTable) Name() string {
	return "Layer7NewCfgContentClassTextTable"
}

func (c *Layer7NewCfgContentClassTextTable) GetParams() BeanType {
	return c.Params
}

func (c *Layer7NewCfgContentClassTextTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *Layer7NewCfgContentClassTextTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.Layer7NewCfgContentClassTextContentClassID).IsZero() &&
		reflect.ValueOf(c.Layer7NewCfgContentClassTextID).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.Layer7NewCfgContentClassTextContentClassID) + "/" + fmt.Sprint(c.Layer7NewCfgContentClassTextID)
}

type Layer7NewCfgContentClassTextTableMatchType int32

const (
	Layer7NewCfgContentClassTextTableMatchType_Include     Layer7NewCfgContentClassTextTableMatchType = 4
	Layer7NewCfgContentClassTextTableMatchType_Regex       Layer7NewCfgContentClassTextTableMatchType = 5
	Layer7NewCfgContentClassTextTableMatchType_Unsupported Layer7NewCfgContentClassTextTableMatchType = 2147483647
)

type Layer7NewCfgContentClassTextTableLookupArea int32

const (
	Layer7NewCfgContentClassTextTableLookupArea_Header      Layer7NewCfgContentClassTextTableLookupArea = 1
	Layer7NewCfgContentClassTextTableLookupArea_Body        Layer7NewCfgContentClassTextTableLookupArea = 2
	Layer7NewCfgContentClassTextTableLookupArea_Both        Layer7NewCfgContentClassTextTableLookupArea = 3
	Layer7NewCfgContentClassTextTableLookupArea_Unsupported Layer7NewCfgContentClassTextTableLookupArea = 2147483647
)

type Layer7NewCfgContentClassTextTableCase int32

const (
	Layer7NewCfgContentClassTextTableCase_Enabled     Layer7NewCfgContentClassTextTableCase = 1
	Layer7NewCfgContentClassTextTableCase_Disabled    Layer7NewCfgContentClassTextTableCase = 2
	Layer7NewCfgContentClassTextTableCase_Unsupported Layer7NewCfgContentClassTextTableCase = 2147483647
)

type Layer7NewCfgContentClassTextTableDelete int32

const (
	Layer7NewCfgContentClassTextTableDelete_Other       Layer7NewCfgContentClassTextTableDelete = 1
	Layer7NewCfgContentClassTextTableDelete_Delete      Layer7NewCfgContentClassTextTableDelete = 2
	Layer7NewCfgContentClassTextTableDelete_Unsupported Layer7NewCfgContentClassTextTableDelete = 2147483647
)

type Layer7NewCfgContentClassTextTableParams struct {
	// The content Class Text ID(key id) as an index, length of the string should be 32 characters.
	ContentClassID string `json:"ContentClassID,omitempty"`
	// The content Class ID(key id) as an index, length of the string should be 32 characters.
	ID string `json:"ID,omitempty"`
	// Content Class Text to match, length of the string should be 32 characters.
	Text string `json:"Text,omitempty"`
	// Text Match type.
	MatchType Layer7NewCfgContentClassTextTableMatchType `json:"MatchType,omitempty"`
	// Text Match type.
	LookupArea Layer7NewCfgContentClassTextTableLookupArea `json:"LookupArea,omitempty"`
	// Enable or Disable Case sensitive for String matching.
	Case Layer7NewCfgContentClassTextTableCase `json:"Case,omitempty"`
	// By setting the value to delete(2), the entire row is deleted.
	Delete Layer7NewCfgContentClassTextTableDelete `json:"Delete,omitempty"`
	// This is an action object.Enter the new Text
	// to which the curent Text has to be copied.
	Copy DisplayString `json:"Copy,omitempty"`
}

func (p Layer7NewCfgContentClassTextTableParams) iMABean() {}
