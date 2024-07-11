package beans

import (
	"fmt"
	"reflect"
)

// Layer7CurCfgContentClassXmlTable The table for configuring Content Class Xml Tag.
// Note:This mib is not supported for VX instance of virtualization.
type Layer7CurCfgContentClassXmlTable struct {
	// content Class Text ID(key id) as an index, length of the string should be 32 characters.
	Layer7CurCfgContentClassXmlTagContentClassID string
	// Content Class ID(key id) as an index, length of the string should be 32 characters.
	Layer7CurCfgContentClassXmlTagID string
	Params                           *Layer7CurCfgContentClassXmlTableParams
}

func NewLayer7CurCfgContentClassXmlTableList() *Layer7CurCfgContentClassXmlTable {
	return &Layer7CurCfgContentClassXmlTable{}
}

func NewLayer7CurCfgContentClassXmlTable(
	layer7CurCfgContentClassXmlTagContentClassID string,
	layer7CurCfgContentClassXmlTagID string,
	params *Layer7CurCfgContentClassXmlTableParams,
) *Layer7CurCfgContentClassXmlTable {
	return &Layer7CurCfgContentClassXmlTable{
		Layer7CurCfgContentClassXmlTagContentClassID: layer7CurCfgContentClassXmlTagContentClassID,
		Layer7CurCfgContentClassXmlTagID:             layer7CurCfgContentClassXmlTagID,
		Params:                                       params,
	}
}

func (c *Layer7CurCfgContentClassXmlTable) Name() string {
	return "Layer7CurCfgContentClassXmlTable"
}

func (c *Layer7CurCfgContentClassXmlTable) GetParams() BeanType {
	return c.Params
}

func (c *Layer7CurCfgContentClassXmlTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *Layer7CurCfgContentClassXmlTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.Layer7CurCfgContentClassXmlTagContentClassID).IsZero() &&
		reflect.ValueOf(c.Layer7CurCfgContentClassXmlTagID).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.Layer7CurCfgContentClassXmlTagContentClassID) + "/" + fmt.Sprint(c.Layer7CurCfgContentClassXmlTagID)
}

type Layer7CurCfgContentClassXmlTableTagMatchTypeName int32

const (
	Layer7CurCfgContentClassXmlTableTagMatchTypeName_Sufx        Layer7CurCfgContentClassXmlTableTagMatchTypeName = 1
	Layer7CurCfgContentClassXmlTableTagMatchTypeName_Equal       Layer7CurCfgContentClassXmlTableTagMatchTypeName = 3
	Layer7CurCfgContentClassXmlTableTagMatchTypeName_Unsupported Layer7CurCfgContentClassXmlTableTagMatchTypeName = 2147483647
)

type Layer7CurCfgContentClassXmlTableTagMatchTypeVal int32

const (
	Layer7CurCfgContentClassXmlTableTagMatchTypeVal_Sufx        Layer7CurCfgContentClassXmlTableTagMatchTypeVal = 1
	Layer7CurCfgContentClassXmlTableTagMatchTypeVal_Equal       Layer7CurCfgContentClassXmlTableTagMatchTypeVal = 3
	Layer7CurCfgContentClassXmlTableTagMatchTypeVal_Include     Layer7CurCfgContentClassXmlTableTagMatchTypeVal = 4
	Layer7CurCfgContentClassXmlTableTagMatchTypeVal_Unsupported Layer7CurCfgContentClassXmlTableTagMatchTypeVal = 2147483647
)

type Layer7CurCfgContentClassXmlTableTagCase int32

const (
	Layer7CurCfgContentClassXmlTableTagCase_Enabled     Layer7CurCfgContentClassXmlTableTagCase = 1
	Layer7CurCfgContentClassXmlTableTagCase_Disabled    Layer7CurCfgContentClassXmlTableTagCase = 2
	Layer7CurCfgContentClassXmlTableTagCase_Unsupported Layer7CurCfgContentClassXmlTableTagCase = 2147483647
)

type Layer7CurCfgContentClassXmlTableParams struct {
	// content Class Text ID(key id) as an index, length of the string should be 32 characters.
	TagContentClassID string `json:"TagContentClassID,omitempty"`
	// Content Class ID(key id) as an index, length of the string should be 32 characters.
	TagID string `json:"TagID,omitempty"`
	// Content Class Xml to match, length of the string should be 32 characters.
	TagName string `json:"TagName,omitempty"`
	// Content Class Xml to match, length of the string should be 32 characters.
	TagVal string `json:"TagVal,omitempty"`
	// Xml Match type.
	TagMatchTypeName Layer7CurCfgContentClassXmlTableTagMatchTypeName `json:"TagMatchTypeName,omitempty"`
	// Xml Match type.
	TagMatchTypeVal Layer7CurCfgContentClassXmlTableTagMatchTypeVal `json:"TagMatchTypeVal,omitempty"`
	// Enable or Disable Case sensitive for String matching.
	TagCase Layer7CurCfgContentClassXmlTableTagCase `json:"TagCase,omitempty"`
}

func (p Layer7CurCfgContentClassXmlTableParams) iMABean() {}
