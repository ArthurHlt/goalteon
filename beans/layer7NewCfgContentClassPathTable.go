package beans

import (
	"fmt"
	"reflect"
)

// Layer7NewCfgContentClassPathTable The table for configuring Content Class Path.
// Note:This mib is not supported for VX instance of virtualization.
type Layer7NewCfgContentClassPathTable struct {
	// The content Class Path ID(key id) as an index, length of the string should be 32 characters.
	Layer7NewCfgContentClassPathContentClassID string
	// The content Class ID(key id) as an index, length of the string should be 32 characters.
	Layer7NewCfgContentClassPathID string
	Params                         *Layer7NewCfgContentClassPathTableParams
}

func NewLayer7NewCfgContentClassPathTableList() *Layer7NewCfgContentClassPathTable {
	return &Layer7NewCfgContentClassPathTable{}
}

func NewLayer7NewCfgContentClassPathTable(
	layer7NewCfgContentClassPathContentClassID string,
	layer7NewCfgContentClassPathID string,
	params *Layer7NewCfgContentClassPathTableParams,
) *Layer7NewCfgContentClassPathTable {
	return &Layer7NewCfgContentClassPathTable{
		Layer7NewCfgContentClassPathContentClassID: layer7NewCfgContentClassPathContentClassID,
		Layer7NewCfgContentClassPathID:             layer7NewCfgContentClassPathID,
		Params:                                     params,
	}
}

func (c *Layer7NewCfgContentClassPathTable) Name() string {
	return "Layer7NewCfgContentClassPathTable"
}

func (c *Layer7NewCfgContentClassPathTable) GetParams() BeanType {
	return c.Params
}

func (c *Layer7NewCfgContentClassPathTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *Layer7NewCfgContentClassPathTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.Layer7NewCfgContentClassPathContentClassID).IsZero() &&
		reflect.ValueOf(c.Layer7NewCfgContentClassPathID).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.Layer7NewCfgContentClassPathContentClassID) + "/" + fmt.Sprint(c.Layer7NewCfgContentClassPathID)
}

type Layer7NewCfgContentClassPathTableMatchType int32

const (
	Layer7NewCfgContentClassPathTableMatchType_Sufx        Layer7NewCfgContentClassPathTableMatchType = 1
	Layer7NewCfgContentClassPathTableMatchType_Prefx       Layer7NewCfgContentClassPathTableMatchType = 2
	Layer7NewCfgContentClassPathTableMatchType_Equal       Layer7NewCfgContentClassPathTableMatchType = 3
	Layer7NewCfgContentClassPathTableMatchType_Include     Layer7NewCfgContentClassPathTableMatchType = 4
	Layer7NewCfgContentClassPathTableMatchType_Regex       Layer7NewCfgContentClassPathTableMatchType = 5
	Layer7NewCfgContentClassPathTableMatchType_Unsupported Layer7NewCfgContentClassPathTableMatchType = 2147483647
)

type Layer7NewCfgContentClassPathTableCase int32

const (
	Layer7NewCfgContentClassPathTableCase_Enabled     Layer7NewCfgContentClassPathTableCase = 1
	Layer7NewCfgContentClassPathTableCase_Disabled    Layer7NewCfgContentClassPathTableCase = 2
	Layer7NewCfgContentClassPathTableCase_Unsupported Layer7NewCfgContentClassPathTableCase = 2147483647
)

type Layer7NewCfgContentClassPathTableDelete int32

const (
	Layer7NewCfgContentClassPathTableDelete_Other       Layer7NewCfgContentClassPathTableDelete = 1
	Layer7NewCfgContentClassPathTableDelete_Delete      Layer7NewCfgContentClassPathTableDelete = 2
	Layer7NewCfgContentClassPathTableDelete_Unsupported Layer7NewCfgContentClassPathTableDelete = 2147483647
)

type Layer7NewCfgContentClassPathTableParams struct {
	// The content Class Path ID(key id) as an index, length of the string should be 32 characters.
	ContentClassID string `json:"ContentClassID,omitempty"`
	// The content Class ID(key id) as an index, length of the string should be 32 characters.
	ID string `json:"ID,omitempty"`
	// Content Class Path to match, length of the string should be 256 characters.
	FilePath string `json:"FilePath,omitempty"`
	// Path Match type.
	MatchType Layer7NewCfgContentClassPathTableMatchType `json:"MatchType,omitempty"`
	// Enable or Disable Case sensitive for String matching.
	Case Layer7NewCfgContentClassPathTableCase `json:"Case,omitempty"`
	// By setting the value to delete(2), the entire row is deleted.
	Delete Layer7NewCfgContentClassPathTableDelete `json:"Delete,omitempty"`
	// Id of a data class to associate Path element with.
	DataclassID string `json:"DataclassID,omitempty"`
	// This is an action object.Enter the new Classpath
	// to which the curent Classpath has to be copied.
	Copy DisplayString `json:"Copy,omitempty"`
}

func (p Layer7NewCfgContentClassPathTableParams) iMABean() {}
