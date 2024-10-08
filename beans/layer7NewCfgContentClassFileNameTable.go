package beans

import (
	"fmt"
	"reflect"
)

// Layer7NewCfgContentClassFileNameTable The table for configuring Content Class File Name.
// Note:This mib is not supported for VX instance of virtualization.
type Layer7NewCfgContentClassFileNameTable struct {
	// The content Class FileName ID(key id) as an index, length of the string should be 32 characters.
	Layer7NewCfgContentClassFileNameContentClassID string
	// The content Class ID(key id) as an index, length of the string should be 32 characters.
	Layer7NewCfgContentClassFileNameID string
	Params                             *Layer7NewCfgContentClassFileNameTableParams
}

func NewLayer7NewCfgContentClassFileNameTableList() *Layer7NewCfgContentClassFileNameTable {
	return &Layer7NewCfgContentClassFileNameTable{}
}

func NewLayer7NewCfgContentClassFileNameTable(
	layer7NewCfgContentClassFileNameContentClassID string,
	layer7NewCfgContentClassFileNameID string,
	params *Layer7NewCfgContentClassFileNameTableParams,
) *Layer7NewCfgContentClassFileNameTable {
	return &Layer7NewCfgContentClassFileNameTable{
		Layer7NewCfgContentClassFileNameContentClassID: layer7NewCfgContentClassFileNameContentClassID,
		Layer7NewCfgContentClassFileNameID:             layer7NewCfgContentClassFileNameID,
		Params:                                         params,
	}
}

func (c *Layer7NewCfgContentClassFileNameTable) Name() string {
	return "Layer7NewCfgContentClassFileNameTable"
}

func (c *Layer7NewCfgContentClassFileNameTable) GetParams() BeanType {
	return c.Params
}

func (c *Layer7NewCfgContentClassFileNameTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *Layer7NewCfgContentClassFileNameTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.Layer7NewCfgContentClassFileNameContentClassID).IsZero() &&
		reflect.ValueOf(c.Layer7NewCfgContentClassFileNameID).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.Layer7NewCfgContentClassFileNameContentClassID) + "/" + fmt.Sprint(c.Layer7NewCfgContentClassFileNameID)
}

type Layer7NewCfgContentClassFileNameTableMatchType int32

const (
	Layer7NewCfgContentClassFileNameTableMatchType_Sufx        Layer7NewCfgContentClassFileNameTableMatchType = 1
	Layer7NewCfgContentClassFileNameTableMatchType_Prefx       Layer7NewCfgContentClassFileNameTableMatchType = 2
	Layer7NewCfgContentClassFileNameTableMatchType_Equal       Layer7NewCfgContentClassFileNameTableMatchType = 3
	Layer7NewCfgContentClassFileNameTableMatchType_Include     Layer7NewCfgContentClassFileNameTableMatchType = 4
	Layer7NewCfgContentClassFileNameTableMatchType_Regex       Layer7NewCfgContentClassFileNameTableMatchType = 5
	Layer7NewCfgContentClassFileNameTableMatchType_Unsupported Layer7NewCfgContentClassFileNameTableMatchType = 2147483647
)

type Layer7NewCfgContentClassFileNameTableCase int32

const (
	Layer7NewCfgContentClassFileNameTableCase_Enabled     Layer7NewCfgContentClassFileNameTableCase = 1
	Layer7NewCfgContentClassFileNameTableCase_Disabled    Layer7NewCfgContentClassFileNameTableCase = 2
	Layer7NewCfgContentClassFileNameTableCase_Unsupported Layer7NewCfgContentClassFileNameTableCase = 2147483647
)

type Layer7NewCfgContentClassFileNameTableDelete int32

const (
	Layer7NewCfgContentClassFileNameTableDelete_Other       Layer7NewCfgContentClassFileNameTableDelete = 1
	Layer7NewCfgContentClassFileNameTableDelete_Delete      Layer7NewCfgContentClassFileNameTableDelete = 2
	Layer7NewCfgContentClassFileNameTableDelete_Unsupported Layer7NewCfgContentClassFileNameTableDelete = 2147483647
)

type Layer7NewCfgContentClassFileNameTableParams struct {
	// The content Class FileName ID(key id) as an index, length of the string should be 32 characters.
	ContentClassID string `json:"ContentClassID,omitempty"`
	// The content Class ID(key id) as an index, length of the string should be 32 characters.
	ID string `json:"ID,omitempty"`
	// Content Class Filename to match, length of the string should be 32 characters.
	FileName string `json:"FileName,omitempty"`
	// File Name Match type.
	MatchType Layer7NewCfgContentClassFileNameTableMatchType `json:"MatchType,omitempty"`
	// Enable or Disable Case sensitive for String matching.
	Case Layer7NewCfgContentClassFileNameTableCase `json:"Case,omitempty"`
	// By setting the value to delete(2), the entire row is deleted.
	Delete Layer7NewCfgContentClassFileNameTableDelete `json:"Delete,omitempty"`
	// This is an action object.Enter the new Filename
	// to which the curent Filename has to be copied.
	Copy DisplayString `json:"Copy,omitempty"`
}

func (p Layer7NewCfgContentClassFileNameTableParams) iMABean() {}
