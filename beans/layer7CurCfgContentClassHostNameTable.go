package beans

import (
	"fmt"
	"reflect"
)

// Layer7CurCfgContentClassHostNameTable The table for configuring Content Class Hostname.
// Note:This mib is not supported for VX instance of virtualization.
type Layer7CurCfgContentClassHostNameTable struct {
	// The content Class HostName ID(key id) as an index, length of the string should be 32 characters.
	Layer7CurCfgContentClassHostNameContentClassID string
	// The content Class ID(key id) as an index, length of the string should be 32 characters.
	Layer7CurCfgContentClassHostNameID string
	Params                             *Layer7CurCfgContentClassHostNameTableParams
}

func NewLayer7CurCfgContentClassHostNameTableList() *Layer7CurCfgContentClassHostNameTable {
	return &Layer7CurCfgContentClassHostNameTable{}
}

func NewLayer7CurCfgContentClassHostNameTable(
	layer7CurCfgContentClassHostNameContentClassID string,
	layer7CurCfgContentClassHostNameID string,
	params *Layer7CurCfgContentClassHostNameTableParams,
) *Layer7CurCfgContentClassHostNameTable {
	return &Layer7CurCfgContentClassHostNameTable{
		Layer7CurCfgContentClassHostNameContentClassID: layer7CurCfgContentClassHostNameContentClassID,
		Layer7CurCfgContentClassHostNameID:             layer7CurCfgContentClassHostNameID,
		Params:                                         params,
	}
}

func (c *Layer7CurCfgContentClassHostNameTable) Name() string {
	return "Layer7CurCfgContentClassHostNameTable"
}

func (c *Layer7CurCfgContentClassHostNameTable) GetParams() BeanType {
	return c.Params
}

func (c *Layer7CurCfgContentClassHostNameTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *Layer7CurCfgContentClassHostNameTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.Layer7CurCfgContentClassHostNameContentClassID).IsZero() &&
		reflect.ValueOf(c.Layer7CurCfgContentClassHostNameID).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.Layer7CurCfgContentClassHostNameContentClassID) + "/" + fmt.Sprint(c.Layer7CurCfgContentClassHostNameID)
}

type Layer7CurCfgContentClassHostNameTableMatchType int32

const (
	Layer7CurCfgContentClassHostNameTableMatchType_Sufx        Layer7CurCfgContentClassHostNameTableMatchType = 1
	Layer7CurCfgContentClassHostNameTableMatchType_Prefx       Layer7CurCfgContentClassHostNameTableMatchType = 2
	Layer7CurCfgContentClassHostNameTableMatchType_Equal       Layer7CurCfgContentClassHostNameTableMatchType = 3
	Layer7CurCfgContentClassHostNameTableMatchType_Include     Layer7CurCfgContentClassHostNameTableMatchType = 4
	Layer7CurCfgContentClassHostNameTableMatchType_Regex       Layer7CurCfgContentClassHostNameTableMatchType = 5
	Layer7CurCfgContentClassHostNameTableMatchType_Unsupported Layer7CurCfgContentClassHostNameTableMatchType = 2147483647
)

type Layer7CurCfgContentClassHostNameTableParams struct {
	// The content Class HostName ID(key id) as an index, length of the string should be 32 characters.
	ContentClassID string `json:"ContentClassID,omitempty"`
	// The content Class ID(key id) as an index, length of the string should be 32 characters.
	ID string `json:"ID,omitempty"`
	// Content Class Hostname to match, length of the string should be 32 characters.
	HostName string `json:"HostName,omitempty"`
	// Host Name Match type.
	MatchType Layer7CurCfgContentClassHostNameTableMatchType `json:"MatchType,omitempty"`
	// Id of a data class to associate Hostname element with.
	DataclassID string `json:"DataclassID,omitempty"`
}

func (p Layer7CurCfgContentClassHostNameTableParams) iMABean() {}
