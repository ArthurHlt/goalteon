package beans

import (
	"fmt"
	"reflect"
)

// Layer7NewCfgContentClassCookieTable The table for configuring Content Class Cookie.
// Note:This mib is not supported for VX instance of virtualization.
type Layer7NewCfgContentClassCookieTable struct {
	// The content Class Cookie ID(key id) as an index, length of the string should be 32 characters.
	Layer7NewCfgContentClassCookieContentClassID string
	// The content Class ID(key id) as an index, length of the string should be 32 characters.
	Layer7NewCfgContentClassCookieID string
	Params                           *Layer7NewCfgContentClassCookieTableParams
}

func NewLayer7NewCfgContentClassCookieTableList() *Layer7NewCfgContentClassCookieTable {
	return &Layer7NewCfgContentClassCookieTable{}
}

func NewLayer7NewCfgContentClassCookieTable(
	layer7NewCfgContentClassCookieContentClassID string,
	layer7NewCfgContentClassCookieID string,
	params *Layer7NewCfgContentClassCookieTableParams,
) *Layer7NewCfgContentClassCookieTable {
	return &Layer7NewCfgContentClassCookieTable{
		Layer7NewCfgContentClassCookieContentClassID: layer7NewCfgContentClassCookieContentClassID,
		Layer7NewCfgContentClassCookieID:             layer7NewCfgContentClassCookieID,
		Params:                                       params,
	}
}

func (c *Layer7NewCfgContentClassCookieTable) Name() string {
	return "Layer7NewCfgContentClassCookieTable"
}

func (c *Layer7NewCfgContentClassCookieTable) GetParams() BeanType {
	return c.Params
}

func (c *Layer7NewCfgContentClassCookieTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *Layer7NewCfgContentClassCookieTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.Layer7NewCfgContentClassCookieContentClassID).IsZero() &&
		reflect.ValueOf(c.Layer7NewCfgContentClassCookieID).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.Layer7NewCfgContentClassCookieContentClassID) + "/" + fmt.Sprint(c.Layer7NewCfgContentClassCookieID)
}

type Layer7NewCfgContentClassCookieTableMatchTypeKey int32

const (
	Layer7NewCfgContentClassCookieTableMatchTypeKey_Equal       Layer7NewCfgContentClassCookieTableMatchTypeKey = 3
	Layer7NewCfgContentClassCookieTableMatchTypeKey_Include     Layer7NewCfgContentClassCookieTableMatchTypeKey = 4
	Layer7NewCfgContentClassCookieTableMatchTypeKey_Regex       Layer7NewCfgContentClassCookieTableMatchTypeKey = 5
	Layer7NewCfgContentClassCookieTableMatchTypeKey_Unsupported Layer7NewCfgContentClassCookieTableMatchTypeKey = 2147483647
)

type Layer7NewCfgContentClassCookieTableMatchTypeVal int32

const (
	Layer7NewCfgContentClassCookieTableMatchTypeVal_Equal       Layer7NewCfgContentClassCookieTableMatchTypeVal = 3
	Layer7NewCfgContentClassCookieTableMatchTypeVal_Include     Layer7NewCfgContentClassCookieTableMatchTypeVal = 4
	Layer7NewCfgContentClassCookieTableMatchTypeVal_Regex       Layer7NewCfgContentClassCookieTableMatchTypeVal = 5
	Layer7NewCfgContentClassCookieTableMatchTypeVal_Unsupported Layer7NewCfgContentClassCookieTableMatchTypeVal = 2147483647
)

type Layer7NewCfgContentClassCookieTableCase int32

const (
	Layer7NewCfgContentClassCookieTableCase_Enabled     Layer7NewCfgContentClassCookieTableCase = 1
	Layer7NewCfgContentClassCookieTableCase_Disabled    Layer7NewCfgContentClassCookieTableCase = 2
	Layer7NewCfgContentClassCookieTableCase_Unsupported Layer7NewCfgContentClassCookieTableCase = 2147483647
)

type Layer7NewCfgContentClassCookieTableDelete int32

const (
	Layer7NewCfgContentClassCookieTableDelete_Other       Layer7NewCfgContentClassCookieTableDelete = 1
	Layer7NewCfgContentClassCookieTableDelete_Delete      Layer7NewCfgContentClassCookieTableDelete = 2
	Layer7NewCfgContentClassCookieTableDelete_Unsupported Layer7NewCfgContentClassCookieTableDelete = 2147483647
)

type Layer7NewCfgContentClassCookieTableParams struct {
	// The content Class Cookie ID(key id) as an index, length of the string should be 32 characters.
	ContentClassID string `json:"ContentClassID,omitempty"`
	// The content Class ID(key id) as an index, length of the string should be 32 characters.
	ID string `json:"ID,omitempty"`
	// Content Class Cookie to match, length of the string should be 32 characters.
	Key string `json:"Key,omitempty"`
	// Content Class Cookie to match, length of the string should be 32 characters.
	Val string `json:"Val,omitempty"`
	// Header Match type.
	MatchTypeKey Layer7NewCfgContentClassCookieTableMatchTypeKey `json:"MatchTypeKey,omitempty"`
	// Header Match type.
	MatchTypeVal Layer7NewCfgContentClassCookieTableMatchTypeVal `json:"MatchTypeVal,omitempty"`
	// Enable or Disable Case sensitive for String matching.
	Case Layer7NewCfgContentClassCookieTableCase `json:"Case,omitempty"`
	// By setting the value to delete(2), the entire row is deleted.
	Delete Layer7NewCfgContentClassCookieTableDelete `json:"Delete,omitempty"`
	// This is an action object.Enter the new Cookie
	// to which the curent Cookie has to be copied.
	Copy DisplayString `json:"Copy,omitempty"`
}

func (p Layer7NewCfgContentClassCookieTableParams) iMABean() {}
