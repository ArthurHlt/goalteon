package beans

import (
	"fmt"
	"reflect"
)

// Layer7CurCfgContentClassFileNameTable The table for configuring Content Class File Name.
// Note:This mib is not supported for VX instance of virtualization.
type Layer7CurCfgContentClassFileNameTable struct {
	// The content Class FileName ID(key id) as an index, length of the string should be 32 characters.
	Layer7CurCfgContentClassFileNameContentClassID string
	// The content Class ID(key id) as an index, length of the string should be 32 characters.
	Layer7CurCfgContentClassFileNameID string
	Params                             *Layer7CurCfgContentClassFileNameTableParams
}

func NewLayer7CurCfgContentClassFileNameTableList() *Layer7CurCfgContentClassFileNameTable {
	return &Layer7CurCfgContentClassFileNameTable{}
}

func NewLayer7CurCfgContentClassFileNameTable(
	layer7CurCfgContentClassFileNameContentClassID string,
	layer7CurCfgContentClassFileNameID string,
	params *Layer7CurCfgContentClassFileNameTableParams,
) *Layer7CurCfgContentClassFileNameTable {
	return &Layer7CurCfgContentClassFileNameTable{
		Layer7CurCfgContentClassFileNameContentClassID: layer7CurCfgContentClassFileNameContentClassID,
		Layer7CurCfgContentClassFileNameID:             layer7CurCfgContentClassFileNameID,
		Params:                                         params,
	}
}

func (c *Layer7CurCfgContentClassFileNameTable) Name() string {
	return "Layer7CurCfgContentClassFileNameTable"
}

func (c *Layer7CurCfgContentClassFileNameTable) GetParams() BeanType {
	return c.Params
}

func (c *Layer7CurCfgContentClassFileNameTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *Layer7CurCfgContentClassFileNameTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.Layer7CurCfgContentClassFileNameContentClassID).IsZero() &&
		reflect.ValueOf(c.Layer7CurCfgContentClassFileNameID).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.Layer7CurCfgContentClassFileNameContentClassID) + "/" + fmt.Sprint(c.Layer7CurCfgContentClassFileNameID)
}

type Layer7CurCfgContentClassFileNameTableMatchType int32

const (
	Layer7CurCfgContentClassFileNameTableMatchType_Sufx        Layer7CurCfgContentClassFileNameTableMatchType = 1
	Layer7CurCfgContentClassFileNameTableMatchType_Prefx       Layer7CurCfgContentClassFileNameTableMatchType = 2
	Layer7CurCfgContentClassFileNameTableMatchType_Equal       Layer7CurCfgContentClassFileNameTableMatchType = 3
	Layer7CurCfgContentClassFileNameTableMatchType_Include     Layer7CurCfgContentClassFileNameTableMatchType = 4
	Layer7CurCfgContentClassFileNameTableMatchType_Regex       Layer7CurCfgContentClassFileNameTableMatchType = 5
	Layer7CurCfgContentClassFileNameTableMatchType_Unsupported Layer7CurCfgContentClassFileNameTableMatchType = 2147483647
)

type Layer7CurCfgContentClassFileNameTableCase int32

const (
	Layer7CurCfgContentClassFileNameTableCase_Enabled     Layer7CurCfgContentClassFileNameTableCase = 1
	Layer7CurCfgContentClassFileNameTableCase_Disabled    Layer7CurCfgContentClassFileNameTableCase = 2
	Layer7CurCfgContentClassFileNameTableCase_Unsupported Layer7CurCfgContentClassFileNameTableCase = 2147483647
)

type Layer7CurCfgContentClassFileNameTableParams struct {
	// The content Class FileName ID(key id) as an index, length of the string should be 32 characters.
	ContentClassID string `json:"ContentClassID,omitempty"`
	// The content Class ID(key id) as an index, length of the string should be 32 characters.
	ID string `json:"ID,omitempty"`
	// Content Class Filename to match, length of the string should be 32 characters.
	FileName string `json:"FileName,omitempty"`
	// File Name Match type.
	MatchType Layer7CurCfgContentClassFileNameTableMatchType `json:"MatchType,omitempty"`
	// Enable or Disable Case sensitive for String matching.
	Case Layer7CurCfgContentClassFileNameTableCase `json:"Case,omitempty"`
}

func (p Layer7CurCfgContentClassFileNameTableParams) iMABean() {}
