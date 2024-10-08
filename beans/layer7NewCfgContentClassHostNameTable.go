package beans

import (
	"fmt"
	"reflect"
)

// Layer7NewCfgContentClassHostNameTable The table for configuring Content Class Hostname.
// Note:This mib is not supported for VX instance of virtualization.
type Layer7NewCfgContentClassHostNameTable struct {
	// The content Class HostName ID(key id) as an index, length of the string should be 32 characters.
	Layer7NewCfgContentClassHostNameContentClassID string
	// The content Class ID(key id) as an index, length of the string should be 32 characters.
	Layer7NewCfgContentClassHostNameID string
	Params                             *Layer7NewCfgContentClassHostNameTableParams
}

func NewLayer7NewCfgContentClassHostNameTableList() *Layer7NewCfgContentClassHostNameTable {
	return &Layer7NewCfgContentClassHostNameTable{}
}

func NewLayer7NewCfgContentClassHostNameTable(
	layer7NewCfgContentClassHostNameContentClassID string,
	layer7NewCfgContentClassHostNameID string,
	params *Layer7NewCfgContentClassHostNameTableParams,
) *Layer7NewCfgContentClassHostNameTable {
	return &Layer7NewCfgContentClassHostNameTable{
		Layer7NewCfgContentClassHostNameContentClassID: layer7NewCfgContentClassHostNameContentClassID,
		Layer7NewCfgContentClassHostNameID:             layer7NewCfgContentClassHostNameID,
		Params:                                         params,
	}
}

func (c *Layer7NewCfgContentClassHostNameTable) Name() string {
	return "Layer7NewCfgContentClassHostNameTable"
}

func (c *Layer7NewCfgContentClassHostNameTable) GetParams() BeanType {
	return c.Params
}

func (c *Layer7NewCfgContentClassHostNameTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *Layer7NewCfgContentClassHostNameTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.Layer7NewCfgContentClassHostNameContentClassID).IsZero() &&
		reflect.ValueOf(c.Layer7NewCfgContentClassHostNameID).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.Layer7NewCfgContentClassHostNameContentClassID) + "/" + fmt.Sprint(c.Layer7NewCfgContentClassHostNameID)
}

type Layer7NewCfgContentClassHostNameTableMatchType int32

const (
	Layer7NewCfgContentClassHostNameTableMatchType_Sufx        Layer7NewCfgContentClassHostNameTableMatchType = 1
	Layer7NewCfgContentClassHostNameTableMatchType_Prefx       Layer7NewCfgContentClassHostNameTableMatchType = 2
	Layer7NewCfgContentClassHostNameTableMatchType_Equal       Layer7NewCfgContentClassHostNameTableMatchType = 3
	Layer7NewCfgContentClassHostNameTableMatchType_Include     Layer7NewCfgContentClassHostNameTableMatchType = 4
	Layer7NewCfgContentClassHostNameTableMatchType_Regex       Layer7NewCfgContentClassHostNameTableMatchType = 5
	Layer7NewCfgContentClassHostNameTableMatchType_Unsupported Layer7NewCfgContentClassHostNameTableMatchType = 2147483647
)

type Layer7NewCfgContentClassHostNameTableDelete int32

const (
	Layer7NewCfgContentClassHostNameTableDelete_Other       Layer7NewCfgContentClassHostNameTableDelete = 1
	Layer7NewCfgContentClassHostNameTableDelete_Delete      Layer7NewCfgContentClassHostNameTableDelete = 2
	Layer7NewCfgContentClassHostNameTableDelete_Unsupported Layer7NewCfgContentClassHostNameTableDelete = 2147483647
)

type Layer7NewCfgContentClassHostNameTableParams struct {
	// The content Class HostName ID(key id) as an index, length of the string should be 32 characters.
	ContentClassID string `json:"ContentClassID,omitempty"`
	// The content Class ID(key id) as an index, length of the string should be 32 characters.
	ID string `json:"ID,omitempty"`
	// Content Class Hostname to match, length of the string should be 32 characters.
	HostName string `json:"HostName,omitempty"`
	// Host Name Match type.
	MatchType Layer7NewCfgContentClassHostNameTableMatchType `json:"MatchType,omitempty"`
	// By setting the value to delete(2), the entire row is deleted.
	Delete Layer7NewCfgContentClassHostNameTableDelete `json:"Delete,omitempty"`
	// Id of a data class to associate Hostname element with.
	DataclassID string `json:"DataclassID,omitempty"`
	// This is an action object.Enter the new HostName
	// to which the curent HostName has to be copied.
	Copy DisplayString `json:"Copy,omitempty"`
}

func (p Layer7NewCfgContentClassHostNameTableParams) iMABean() {}
