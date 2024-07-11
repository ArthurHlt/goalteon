package beans

import (
	"fmt"
	"reflect"
)

// Layer7CurCfgContentClassFileTypeTable The table for configuring Content Class File Name.
// Note:This mib is not supported for VX instance of virtualization.
type Layer7CurCfgContentClassFileTypeTable struct {
	// The content Class Filetype ID(key id) as an index, length of the string should be 32 characters.
	Layer7CurCfgContentClassFileTypeContentClassID string
	// The content Class ID(key id) as an index, length of the string should be 32 characters.
	Layer7CurCfgContentClassFileTypeID string
	Params                             *Layer7CurCfgContentClassFileTypeTableParams
}

func NewLayer7CurCfgContentClassFileTypeTableList() *Layer7CurCfgContentClassFileTypeTable {
	return &Layer7CurCfgContentClassFileTypeTable{}
}

func NewLayer7CurCfgContentClassFileTypeTable(
	layer7CurCfgContentClassFileTypeContentClassID string,
	layer7CurCfgContentClassFileTypeID string,
	params *Layer7CurCfgContentClassFileTypeTableParams,
) *Layer7CurCfgContentClassFileTypeTable {
	return &Layer7CurCfgContentClassFileTypeTable{
		Layer7CurCfgContentClassFileTypeContentClassID: layer7CurCfgContentClassFileTypeContentClassID,
		Layer7CurCfgContentClassFileTypeID:             layer7CurCfgContentClassFileTypeID,
		Params:                                         params,
	}
}

func (c *Layer7CurCfgContentClassFileTypeTable) Name() string {
	return "Layer7CurCfgContentClassFileTypeTable"
}

func (c *Layer7CurCfgContentClassFileTypeTable) GetParams() BeanType {
	return c.Params
}

func (c *Layer7CurCfgContentClassFileTypeTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *Layer7CurCfgContentClassFileTypeTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.Layer7CurCfgContentClassFileTypeContentClassID).IsZero() &&
		reflect.ValueOf(c.Layer7CurCfgContentClassFileTypeID).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.Layer7CurCfgContentClassFileTypeContentClassID) + "/" + fmt.Sprint(c.Layer7CurCfgContentClassFileTypeID)
}

type Layer7CurCfgContentClassFileTypeTableMatchType int32

const (
	Layer7CurCfgContentClassFileTypeTableMatchType_Sufx        Layer7CurCfgContentClassFileTypeTableMatchType = 1
	Layer7CurCfgContentClassFileTypeTableMatchType_Prefx       Layer7CurCfgContentClassFileTypeTableMatchType = 2
	Layer7CurCfgContentClassFileTypeTableMatchType_Equal       Layer7CurCfgContentClassFileTypeTableMatchType = 3
	Layer7CurCfgContentClassFileTypeTableMatchType_Include     Layer7CurCfgContentClassFileTypeTableMatchType = 4
	Layer7CurCfgContentClassFileTypeTableMatchType_Regex       Layer7CurCfgContentClassFileTypeTableMatchType = 5
	Layer7CurCfgContentClassFileTypeTableMatchType_Unsupported Layer7CurCfgContentClassFileTypeTableMatchType = 2147483647
)

type Layer7CurCfgContentClassFileTypeTableCase int32

const (
	Layer7CurCfgContentClassFileTypeTableCase_Enabled     Layer7CurCfgContentClassFileTypeTableCase = 1
	Layer7CurCfgContentClassFileTypeTableCase_Disabled    Layer7CurCfgContentClassFileTypeTableCase = 2
	Layer7CurCfgContentClassFileTypeTableCase_Unsupported Layer7CurCfgContentClassFileTypeTableCase = 2147483647
)

type Layer7CurCfgContentClassFileTypeTableParams struct {
	// The content Class Filetype ID(key id) as an index, length of the string should be 32 characters.
	ContentClassID string `json:"ContentClassID,omitempty"`
	// The content Class ID(key id) as an index, length of the string should be 32 characters.
	ID string `json:"ID,omitempty"`
	// Content Class Filetype to match, length of the string should be 32 characters.
	FileType string `json:"FileType,omitempty"`
	// Filetype Match type.
	MatchType Layer7CurCfgContentClassFileTypeTableMatchType `json:"MatchType,omitempty"`
	// Enable or Disable Case sensitive for String matching.
	Case Layer7CurCfgContentClassFileTypeTableCase `json:"Case,omitempty"`
}

func (p Layer7CurCfgContentClassFileTypeTableParams) iMABean() {}
