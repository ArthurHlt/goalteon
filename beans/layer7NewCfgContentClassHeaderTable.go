package beans

import (
	"fmt"
	"reflect"
)

// Layer7NewCfgContentClassHeaderTable The table for configuring Content Class Header.
// Note:This mib is not supported for VX instance of virtualization.
type Layer7NewCfgContentClassHeaderTable struct {
	// The content Class Header ID(key id) as an index, length of the string should be 32 characters.
	Layer7NewCfgContentClassHeaderContentClassID string
	// The content Class ID(key id) as an index, length of the string should be 32 characters.
	Layer7NewCfgContentClassHeaderID string
	Params                           *Layer7NewCfgContentClassHeaderTableParams
}

func NewLayer7NewCfgContentClassHeaderTableList() *Layer7NewCfgContentClassHeaderTable {
	return &Layer7NewCfgContentClassHeaderTable{}
}

func NewLayer7NewCfgContentClassHeaderTable(
	layer7NewCfgContentClassHeaderContentClassID string,
	layer7NewCfgContentClassHeaderID string,
	params *Layer7NewCfgContentClassHeaderTableParams,
) *Layer7NewCfgContentClassHeaderTable {
	return &Layer7NewCfgContentClassHeaderTable{
		Layer7NewCfgContentClassHeaderContentClassID: layer7NewCfgContentClassHeaderContentClassID,
		Layer7NewCfgContentClassHeaderID:             layer7NewCfgContentClassHeaderID,
		Params:                                       params,
	}
}

func (c *Layer7NewCfgContentClassHeaderTable) Name() string {
	return "Layer7NewCfgContentClassHeaderTable"
}

func (c *Layer7NewCfgContentClassHeaderTable) GetParams() BeanType {
	return c.Params
}

func (c *Layer7NewCfgContentClassHeaderTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *Layer7NewCfgContentClassHeaderTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.Layer7NewCfgContentClassHeaderContentClassID).IsZero() &&
		reflect.ValueOf(c.Layer7NewCfgContentClassHeaderID).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.Layer7NewCfgContentClassHeaderContentClassID) + "/" + fmt.Sprint(c.Layer7NewCfgContentClassHeaderID)
}

type Layer7NewCfgContentClassHeaderTableMatchTypeName int32

const (
	Layer7NewCfgContentClassHeaderTableMatchTypeName_Equal       Layer7NewCfgContentClassHeaderTableMatchTypeName = 3
	Layer7NewCfgContentClassHeaderTableMatchTypeName_Include     Layer7NewCfgContentClassHeaderTableMatchTypeName = 4
	Layer7NewCfgContentClassHeaderTableMatchTypeName_Regex       Layer7NewCfgContentClassHeaderTableMatchTypeName = 5
	Layer7NewCfgContentClassHeaderTableMatchTypeName_Unsupported Layer7NewCfgContentClassHeaderTableMatchTypeName = 2147483647
)

type Layer7NewCfgContentClassHeaderTableMatchTypeVal int32

const (
	Layer7NewCfgContentClassHeaderTableMatchTypeVal_Equal       Layer7NewCfgContentClassHeaderTableMatchTypeVal = 3
	Layer7NewCfgContentClassHeaderTableMatchTypeVal_Include     Layer7NewCfgContentClassHeaderTableMatchTypeVal = 4
	Layer7NewCfgContentClassHeaderTableMatchTypeVal_Regex       Layer7NewCfgContentClassHeaderTableMatchTypeVal = 5
	Layer7NewCfgContentClassHeaderTableMatchTypeVal_Unsupported Layer7NewCfgContentClassHeaderTableMatchTypeVal = 2147483647
)

type Layer7NewCfgContentClassHeaderTableCase int32

const (
	Layer7NewCfgContentClassHeaderTableCase_Enabled     Layer7NewCfgContentClassHeaderTableCase = 1
	Layer7NewCfgContentClassHeaderTableCase_Disabled    Layer7NewCfgContentClassHeaderTableCase = 2
	Layer7NewCfgContentClassHeaderTableCase_Unsupported Layer7NewCfgContentClassHeaderTableCase = 2147483647
)

type Layer7NewCfgContentClassHeaderTableDelete int32

const (
	Layer7NewCfgContentClassHeaderTableDelete_Other       Layer7NewCfgContentClassHeaderTableDelete = 1
	Layer7NewCfgContentClassHeaderTableDelete_Delete      Layer7NewCfgContentClassHeaderTableDelete = 2
	Layer7NewCfgContentClassHeaderTableDelete_Unsupported Layer7NewCfgContentClassHeaderTableDelete = 2147483647
)

type Layer7NewCfgContentClassHeaderTableParams struct {
	// The content Class Header ID(key id) as an index, length of the string should be 32 characters.
	ContentClassID string `json:"ContentClassID,omitempty"`
	// The content Class ID(key id) as an index, length of the string should be 32 characters.
	ID string `json:"ID,omitempty"`
	// Content Class Header to match, length of the string should be 32 characters.
	Name string `json:"Name,omitempty"`
	// Content Class Header to match, length of the string should be 32 characters.
	Val string `json:"Val,omitempty"`
	// Header Match type.
	MatchTypeName Layer7NewCfgContentClassHeaderTableMatchTypeName `json:"MatchTypeName,omitempty"`
	// Header Match type.
	MatchTypeVal Layer7NewCfgContentClassHeaderTableMatchTypeVal `json:"MatchTypeVal,omitempty"`
	// Enable or Disable Case sensitive for String matching.
	Case Layer7NewCfgContentClassHeaderTableCase `json:"Case,omitempty"`
	// By setting the value to delete(2), the entire row is deleted.
	Delete Layer7NewCfgContentClassHeaderTableDelete `json:"Delete,omitempty"`
	// This is an action object.Enter the new Header
	// to which the curent Header has to be copied.
	Copy string `json:"Copy,omitempty"`
}

func (p Layer7NewCfgContentClassHeaderTableParams) iMABean() {}
