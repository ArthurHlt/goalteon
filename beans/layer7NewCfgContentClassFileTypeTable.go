package beans

import (
	"fmt"
	"reflect"
)

// Layer7NewCfgContentClassFileTypeTable The table for configuring Content Class File Name.
// Note:This mib is not supported for VX instance of virtualization.
type Layer7NewCfgContentClassFileTypeTable struct {
	// The content Class Filetype ID(key id) as an index, length of the string should be 32 characters.
	Layer7NewCfgContentClassFileTypeContentClassID string
	// The content Class ID(key id) as an index, length of the string should be 32 characters.
	Layer7NewCfgContentClassFileTypeID string
	Params                             *Layer7NewCfgContentClassFileTypeTableParams
}

func NewLayer7NewCfgContentClassFileTypeTableList() *Layer7NewCfgContentClassFileTypeTable {
	return &Layer7NewCfgContentClassFileTypeTable{}
}

func NewLayer7NewCfgContentClassFileTypeTable(
	layer7NewCfgContentClassFileTypeContentClassID string,
	layer7NewCfgContentClassFileTypeID string,
	params *Layer7NewCfgContentClassFileTypeTableParams,
) *Layer7NewCfgContentClassFileTypeTable {
	return &Layer7NewCfgContentClassFileTypeTable{
		Layer7NewCfgContentClassFileTypeContentClassID: layer7NewCfgContentClassFileTypeContentClassID,
		Layer7NewCfgContentClassFileTypeID:             layer7NewCfgContentClassFileTypeID,
		Params:                                         params,
	}
}

func (c *Layer7NewCfgContentClassFileTypeTable) Name() string {
	return "Layer7NewCfgContentClassFileTypeTable"
}

func (c *Layer7NewCfgContentClassFileTypeTable) GetParams() BeanType {
	return c.Params
}

func (c *Layer7NewCfgContentClassFileTypeTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *Layer7NewCfgContentClassFileTypeTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.Layer7NewCfgContentClassFileTypeContentClassID).IsZero() &&
		reflect.ValueOf(c.Layer7NewCfgContentClassFileTypeID).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.Layer7NewCfgContentClassFileTypeContentClassID) + "/" + fmt.Sprint(c.Layer7NewCfgContentClassFileTypeID)
}

type Layer7NewCfgContentClassFileTypeTableMatchType int32

const (
	Layer7NewCfgContentClassFileTypeTableMatchType_Sufx        Layer7NewCfgContentClassFileTypeTableMatchType = 1
	Layer7NewCfgContentClassFileTypeTableMatchType_Prefx       Layer7NewCfgContentClassFileTypeTableMatchType = 2
	Layer7NewCfgContentClassFileTypeTableMatchType_Equal       Layer7NewCfgContentClassFileTypeTableMatchType = 3
	Layer7NewCfgContentClassFileTypeTableMatchType_Include     Layer7NewCfgContentClassFileTypeTableMatchType = 4
	Layer7NewCfgContentClassFileTypeTableMatchType_Regex       Layer7NewCfgContentClassFileTypeTableMatchType = 5
	Layer7NewCfgContentClassFileTypeTableMatchType_Unsupported Layer7NewCfgContentClassFileTypeTableMatchType = 2147483647
)

type Layer7NewCfgContentClassFileTypeTableCase int32

const (
	Layer7NewCfgContentClassFileTypeTableCase_Enabled     Layer7NewCfgContentClassFileTypeTableCase = 1
	Layer7NewCfgContentClassFileTypeTableCase_Disabled    Layer7NewCfgContentClassFileTypeTableCase = 2
	Layer7NewCfgContentClassFileTypeTableCase_Unsupported Layer7NewCfgContentClassFileTypeTableCase = 2147483647
)

type Layer7NewCfgContentClassFileTypeTableDelete int32

const (
	Layer7NewCfgContentClassFileTypeTableDelete_Other       Layer7NewCfgContentClassFileTypeTableDelete = 1
	Layer7NewCfgContentClassFileTypeTableDelete_Delete      Layer7NewCfgContentClassFileTypeTableDelete = 2
	Layer7NewCfgContentClassFileTypeTableDelete_Unsupported Layer7NewCfgContentClassFileTypeTableDelete = 2147483647
)

type Layer7NewCfgContentClassFileTypeTableParams struct {
	// The content Class Filetype ID(key id) as an index, length of the string should be 32 characters.
	ContentClassID string `json:"ContentClassID,omitempty"`
	// The content Class ID(key id) as an index, length of the string should be 32 characters.
	ID string `json:"ID,omitempty"`
	// Content Class Filetype to match, length of the string should be 32 characters.
	FileType string `json:"FileType,omitempty"`
	// Filetype Match type.
	MatchType Layer7NewCfgContentClassFileTypeTableMatchType `json:"MatchType,omitempty"`
	// Enable or Disable Case sensitive for String matching.
	Case Layer7NewCfgContentClassFileTypeTableCase `json:"Case,omitempty"`
	// By setting the value to delete(2), the entire row is deleted.
	Delete Layer7NewCfgContentClassFileTypeTableDelete `json:"Delete,omitempty"`
	// This is an action object.Enter the new Filetype
	// to which the curent Filetype has to be copied.
	Copy DisplayString `json:"Copy,omitempty"`
}

func (p Layer7NewCfgContentClassFileTypeTableParams) iMABean() {}
