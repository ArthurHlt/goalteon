package beans

import (
	"fmt"
	"reflect"
)

// Layer7CurCfgContentClassCookieTable The table for configuring Content Class Cookie.
// Note:This mib is not supported for VX instance of virtualization.
type Layer7CurCfgContentClassCookieTable struct {
	// The content Class Cookie ID(key id) as an index, length of the string should be 32 characters.
	Layer7CurCfgContentClassCookieContentClassID string
	// The content Class ID(key id) as an index, length of the string should be 32 characters.
	Layer7CurCfgContentClassCookieID string
	Params                           *Layer7CurCfgContentClassCookieTableParams
}

func NewLayer7CurCfgContentClassCookieTableList() *Layer7CurCfgContentClassCookieTable {
	return &Layer7CurCfgContentClassCookieTable{}
}

func NewLayer7CurCfgContentClassCookieTable(
	layer7CurCfgContentClassCookieContentClassID string,
	layer7CurCfgContentClassCookieID string,
	params *Layer7CurCfgContentClassCookieTableParams,
) *Layer7CurCfgContentClassCookieTable {
	return &Layer7CurCfgContentClassCookieTable{
		Layer7CurCfgContentClassCookieContentClassID: layer7CurCfgContentClassCookieContentClassID,
		Layer7CurCfgContentClassCookieID:             layer7CurCfgContentClassCookieID,
		Params:                                       params,
	}
}

func (c *Layer7CurCfgContentClassCookieTable) Name() string {
	return "Layer7CurCfgContentClassCookieTable"
}

func (c *Layer7CurCfgContentClassCookieTable) GetParams() BeanType {
	return c.Params
}

func (c *Layer7CurCfgContentClassCookieTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *Layer7CurCfgContentClassCookieTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.Layer7CurCfgContentClassCookieContentClassID).IsZero() &&
		reflect.ValueOf(c.Layer7CurCfgContentClassCookieID).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.Layer7CurCfgContentClassCookieContentClassID) + "/" + fmt.Sprint(c.Layer7CurCfgContentClassCookieID)
}

type Layer7CurCfgContentClassCookieTableMatchTypeKey int32

const (
	Layer7CurCfgContentClassCookieTableMatchTypeKey_Equal       Layer7CurCfgContentClassCookieTableMatchTypeKey = 3
	Layer7CurCfgContentClassCookieTableMatchTypeKey_Include     Layer7CurCfgContentClassCookieTableMatchTypeKey = 4
	Layer7CurCfgContentClassCookieTableMatchTypeKey_Regex       Layer7CurCfgContentClassCookieTableMatchTypeKey = 5
	Layer7CurCfgContentClassCookieTableMatchTypeKey_Unsupported Layer7CurCfgContentClassCookieTableMatchTypeKey = 2147483647
)

type Layer7CurCfgContentClassCookieTableMatchTypeVal int32

const (
	Layer7CurCfgContentClassCookieTableMatchTypeVal_Equal       Layer7CurCfgContentClassCookieTableMatchTypeVal = 3
	Layer7CurCfgContentClassCookieTableMatchTypeVal_Include     Layer7CurCfgContentClassCookieTableMatchTypeVal = 4
	Layer7CurCfgContentClassCookieTableMatchTypeVal_Regex       Layer7CurCfgContentClassCookieTableMatchTypeVal = 5
	Layer7CurCfgContentClassCookieTableMatchTypeVal_Unsupported Layer7CurCfgContentClassCookieTableMatchTypeVal = 2147483647
)

type Layer7CurCfgContentClassCookieTableCase int32

const (
	Layer7CurCfgContentClassCookieTableCase_Enabled     Layer7CurCfgContentClassCookieTableCase = 1
	Layer7CurCfgContentClassCookieTableCase_Disabled    Layer7CurCfgContentClassCookieTableCase = 2
	Layer7CurCfgContentClassCookieTableCase_Unsupported Layer7CurCfgContentClassCookieTableCase = 2147483647
)

type Layer7CurCfgContentClassCookieTableParams struct {
	// The content Class Cookie ID(key id) as an index, length of the string should be 32 characters.
	ContentClassID string `json:"ContentClassID,omitempty"`
	// The content Class ID(key id) as an index, length of the string should be 32 characters.
	ID string `json:"ID,omitempty"`
	// Content Class Cookie to match, length of the string should be 32 characters.
	Key string `json:"Key,omitempty"`
	// Content Class Cookie to match, length of the string should be 32 characters.
	Val string `json:"Val,omitempty"`
	// Header Match type.
	MatchTypeKey Layer7CurCfgContentClassCookieTableMatchTypeKey `json:"MatchTypeKey,omitempty"`
	// Header Match type.
	MatchTypeVal Layer7CurCfgContentClassCookieTableMatchTypeVal `json:"MatchTypeVal,omitempty"`
	// Enable or Disable Case sensitive for String matching.
	Case Layer7CurCfgContentClassCookieTableCase `json:"Case,omitempty"`
}

func (p Layer7CurCfgContentClassCookieTableParams) iMABean() {}
