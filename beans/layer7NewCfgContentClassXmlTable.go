package beans

import (
	"fmt"
	"reflect"
)

// Layer7NewCfgContentClassXmlTable The table for configuring Content Class Xml Tag.
// Note:This mib is not supported for VX instance of virtualization.
type Layer7NewCfgContentClassXmlTable struct {
	// content Class Text ID(key id) as an index, length of the string should be 32 characters.
	Layer7NewCfgContentClassXmlTagContentClassID string
	// Content Class ID(key id) as an index, length of the string should be 32 characters.
	Layer7NewCfgContentClassXmlTagID string
	Params                           *Layer7NewCfgContentClassXmlTableParams
}

func NewLayer7NewCfgContentClassXmlTableList() *Layer7NewCfgContentClassXmlTable {
	return &Layer7NewCfgContentClassXmlTable{}
}

func NewLayer7NewCfgContentClassXmlTable(
	layer7NewCfgContentClassXmlTagContentClassID string,
	layer7NewCfgContentClassXmlTagID string,
	params *Layer7NewCfgContentClassXmlTableParams,
) *Layer7NewCfgContentClassXmlTable {
	return &Layer7NewCfgContentClassXmlTable{
		Layer7NewCfgContentClassXmlTagContentClassID: layer7NewCfgContentClassXmlTagContentClassID,
		Layer7NewCfgContentClassXmlTagID:             layer7NewCfgContentClassXmlTagID,
		Params:                                       params,
	}
}

func (c *Layer7NewCfgContentClassXmlTable) Name() string {
	return "Layer7NewCfgContentClassXmlTable"
}

func (c *Layer7NewCfgContentClassXmlTable) GetParams() BeanType {
	return c.Params
}

func (c *Layer7NewCfgContentClassXmlTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *Layer7NewCfgContentClassXmlTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.Layer7NewCfgContentClassXmlTagContentClassID).IsZero() &&
		reflect.ValueOf(c.Layer7NewCfgContentClassXmlTagID).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.Layer7NewCfgContentClassXmlTagContentClassID) + "/" + fmt.Sprint(c.Layer7NewCfgContentClassXmlTagID)
}

type Layer7NewCfgContentClassXmlTableTagMatchTypeName int32

const (
	Layer7NewCfgContentClassXmlTableTagMatchTypeName_Sufx        Layer7NewCfgContentClassXmlTableTagMatchTypeName = 1
	Layer7NewCfgContentClassXmlTableTagMatchTypeName_Equal       Layer7NewCfgContentClassXmlTableTagMatchTypeName = 3
	Layer7NewCfgContentClassXmlTableTagMatchTypeName_Unsupported Layer7NewCfgContentClassXmlTableTagMatchTypeName = 2147483647
)

type Layer7NewCfgContentClassXmlTableTagMatchTypeVal int32

const (
	Layer7NewCfgContentClassXmlTableTagMatchTypeVal_Sufx        Layer7NewCfgContentClassXmlTableTagMatchTypeVal = 1
	Layer7NewCfgContentClassXmlTableTagMatchTypeVal_Equal       Layer7NewCfgContentClassXmlTableTagMatchTypeVal = 3
	Layer7NewCfgContentClassXmlTableTagMatchTypeVal_Include     Layer7NewCfgContentClassXmlTableTagMatchTypeVal = 4
	Layer7NewCfgContentClassXmlTableTagMatchTypeVal_Unsupported Layer7NewCfgContentClassXmlTableTagMatchTypeVal = 2147483647
)

type Layer7NewCfgContentClassXmlTableTagCase int32

const (
	Layer7NewCfgContentClassXmlTableTagCase_Enabled     Layer7NewCfgContentClassXmlTableTagCase = 1
	Layer7NewCfgContentClassXmlTableTagCase_Disabled    Layer7NewCfgContentClassXmlTableTagCase = 2
	Layer7NewCfgContentClassXmlTableTagCase_Unsupported Layer7NewCfgContentClassXmlTableTagCase = 2147483647
)

type Layer7NewCfgContentClassXmlTableTagDelete int32

const (
	Layer7NewCfgContentClassXmlTableTagDelete_Other       Layer7NewCfgContentClassXmlTableTagDelete = 1
	Layer7NewCfgContentClassXmlTableTagDelete_Delete      Layer7NewCfgContentClassXmlTableTagDelete = 2
	Layer7NewCfgContentClassXmlTableTagDelete_Unsupported Layer7NewCfgContentClassXmlTableTagDelete = 2147483647
)

type Layer7NewCfgContentClassXmlTableParams struct {
	// content Class Text ID(key id) as an index, length of the string should be 32 characters.
	TagContentClassID string `json:"TagContentClassID,omitempty"`
	// Content Class ID(key id) as an index, length of the string should be 32 characters.
	TagID string `json:"TagID,omitempty"`
	// Content Class Xml to match, length of the string should be 32 characters.
	TagName string `json:"TagName,omitempty"`
	// Content Class Xml to match, length of the string should be 32 characters.
	TagVal string `json:"TagVal,omitempty"`
	// Xml Match type.
	TagMatchTypeName Layer7NewCfgContentClassXmlTableTagMatchTypeName `json:"TagMatchTypeName,omitempty"`
	// Xml Match type.
	TagMatchTypeVal Layer7NewCfgContentClassXmlTableTagMatchTypeVal `json:"TagMatchTypeVal,omitempty"`
	// Enable or Disable Case sensitive for String matching.
	TagCase Layer7NewCfgContentClassXmlTableTagCase `json:"TagCase,omitempty"`
	// By setting the value to delete(2), the entire row is deleted.
	TagDelete Layer7NewCfgContentClassXmlTableTagDelete `json:"TagDelete,omitempty"`
	// This is an action object.Enter the new Xmltag
	// to which the curent Xmltag has to be copied.
	TagCopy string `json:"TagCopy,omitempty"`
}

func (p Layer7NewCfgContentClassXmlTableParams) iMABean() {}
