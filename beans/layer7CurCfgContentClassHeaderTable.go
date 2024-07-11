package beans

import (
	"fmt"
	"reflect"
)

// Layer7CurCfgContentClassHeaderTable The table for configuring Content Class Header.
// Note:This mib is not supported for VX instance of virtualization.
type Layer7CurCfgContentClassHeaderTable struct {
	// The content Class Header ID(key id) as an index, length of the string should be 32 characters.
	Layer7CurCfgContentClassHeaderContentClassID string
	// The content Class ID(key id) as an index, length of the string should be 32 characters.
	Layer7CurCfgContentClassHeaderID string
	Params                           *Layer7CurCfgContentClassHeaderTableParams
}

func NewLayer7CurCfgContentClassHeaderTableList() *Layer7CurCfgContentClassHeaderTable {
	return &Layer7CurCfgContentClassHeaderTable{}
}

func NewLayer7CurCfgContentClassHeaderTable(
	layer7CurCfgContentClassHeaderContentClassID string,
	layer7CurCfgContentClassHeaderID string,
	params *Layer7CurCfgContentClassHeaderTableParams,
) *Layer7CurCfgContentClassHeaderTable {
	return &Layer7CurCfgContentClassHeaderTable{
		Layer7CurCfgContentClassHeaderContentClassID: layer7CurCfgContentClassHeaderContentClassID,
		Layer7CurCfgContentClassHeaderID:             layer7CurCfgContentClassHeaderID,
		Params:                                       params,
	}
}

func (c *Layer7CurCfgContentClassHeaderTable) Name() string {
	return "Layer7CurCfgContentClassHeaderTable"
}

func (c *Layer7CurCfgContentClassHeaderTable) GetParams() BeanType {
	return c.Params
}

func (c *Layer7CurCfgContentClassHeaderTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *Layer7CurCfgContentClassHeaderTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.Layer7CurCfgContentClassHeaderContentClassID).IsZero() &&
		reflect.ValueOf(c.Layer7CurCfgContentClassHeaderID).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.Layer7CurCfgContentClassHeaderContentClassID) + "/" + fmt.Sprint(c.Layer7CurCfgContentClassHeaderID)
}

type Layer7CurCfgContentClassHeaderTableMatchTypeName int32

const (
	Layer7CurCfgContentClassHeaderTableMatchTypeName_Equal       Layer7CurCfgContentClassHeaderTableMatchTypeName = 3
	Layer7CurCfgContentClassHeaderTableMatchTypeName_Include     Layer7CurCfgContentClassHeaderTableMatchTypeName = 4
	Layer7CurCfgContentClassHeaderTableMatchTypeName_Regex       Layer7CurCfgContentClassHeaderTableMatchTypeName = 5
	Layer7CurCfgContentClassHeaderTableMatchTypeName_Unsupported Layer7CurCfgContentClassHeaderTableMatchTypeName = 2147483647
)

type Layer7CurCfgContentClassHeaderTableMatchTypeVal int32

const (
	Layer7CurCfgContentClassHeaderTableMatchTypeVal_Equal       Layer7CurCfgContentClassHeaderTableMatchTypeVal = 3
	Layer7CurCfgContentClassHeaderTableMatchTypeVal_Include     Layer7CurCfgContentClassHeaderTableMatchTypeVal = 4
	Layer7CurCfgContentClassHeaderTableMatchTypeVal_Regex       Layer7CurCfgContentClassHeaderTableMatchTypeVal = 5
	Layer7CurCfgContentClassHeaderTableMatchTypeVal_Unsupported Layer7CurCfgContentClassHeaderTableMatchTypeVal = 2147483647
)

type Layer7CurCfgContentClassHeaderTableCase int32

const (
	Layer7CurCfgContentClassHeaderTableCase_Enabled     Layer7CurCfgContentClassHeaderTableCase = 1
	Layer7CurCfgContentClassHeaderTableCase_Disabled    Layer7CurCfgContentClassHeaderTableCase = 2
	Layer7CurCfgContentClassHeaderTableCase_Unsupported Layer7CurCfgContentClassHeaderTableCase = 2147483647
)

type Layer7CurCfgContentClassHeaderTableParams struct {
	// The content Class Header ID(key id) as an index, length of the string should be 32 characters.
	ContentClassID string `json:"ContentClassID,omitempty"`
	// The content Class ID(key id) as an index, length of the string should be 32 characters.
	ID string `json:"ID,omitempty"`
	// Content Class Header to match, length of the string should be 32 characters.
	Name string `json:"Name,omitempty"`
	// Content Class Header to match, length of the string should be 32 characters.
	Val string `json:"Val,omitempty"`
	// Header Match type.
	MatchTypeName Layer7CurCfgContentClassHeaderTableMatchTypeName `json:"MatchTypeName,omitempty"`
	// Header Match type.
	MatchTypeVal Layer7CurCfgContentClassHeaderTableMatchTypeVal `json:"MatchTypeVal,omitempty"`
	// Enable or Disable Case sensitive for String matching.
	Case Layer7CurCfgContentClassHeaderTableCase `json:"Case,omitempty"`
}

func (p Layer7CurCfgContentClassHeaderTableParams) iMABean() {}
