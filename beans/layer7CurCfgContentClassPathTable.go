package beans

import (
	"fmt"
	"reflect"
)

// Layer7CurCfgContentClassPathTable The table for configuring Content Class Path.
// Note:This mib is not supported for VX instance of virtualization.
type Layer7CurCfgContentClassPathTable struct {
	// The content Class Path ID(key id) as an index, length of the string should be 32 characters.
	Layer7CurCfgContentClassPathContentClassID string
	// The content Class ID(key id) as an index, length of the string should be 32 characters.
	Layer7CurCfgContentClassPathID string
	Params                         *Layer7CurCfgContentClassPathTableParams
}

func NewLayer7CurCfgContentClassPathTableList() *Layer7CurCfgContentClassPathTable {
	return &Layer7CurCfgContentClassPathTable{}
}

func NewLayer7CurCfgContentClassPathTable(
	layer7CurCfgContentClassPathContentClassID string,
	layer7CurCfgContentClassPathID string,
	params *Layer7CurCfgContentClassPathTableParams,
) *Layer7CurCfgContentClassPathTable {
	return &Layer7CurCfgContentClassPathTable{
		Layer7CurCfgContentClassPathContentClassID: layer7CurCfgContentClassPathContentClassID,
		Layer7CurCfgContentClassPathID:             layer7CurCfgContentClassPathID,
		Params:                                     params,
	}
}

func (c *Layer7CurCfgContentClassPathTable) Name() string {
	return "Layer7CurCfgContentClassPathTable"
}

func (c *Layer7CurCfgContentClassPathTable) GetParams() BeanType {
	return c.Params
}

func (c *Layer7CurCfgContentClassPathTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *Layer7CurCfgContentClassPathTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.Layer7CurCfgContentClassPathContentClassID).IsZero() &&
		reflect.ValueOf(c.Layer7CurCfgContentClassPathID).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.Layer7CurCfgContentClassPathContentClassID) + "/" + fmt.Sprint(c.Layer7CurCfgContentClassPathID)
}

type Layer7CurCfgContentClassPathTableMatchType int32

const (
	Layer7CurCfgContentClassPathTableMatchType_Sufx        Layer7CurCfgContentClassPathTableMatchType = 1
	Layer7CurCfgContentClassPathTableMatchType_Prefx       Layer7CurCfgContentClassPathTableMatchType = 2
	Layer7CurCfgContentClassPathTableMatchType_Equal       Layer7CurCfgContentClassPathTableMatchType = 3
	Layer7CurCfgContentClassPathTableMatchType_Include     Layer7CurCfgContentClassPathTableMatchType = 4
	Layer7CurCfgContentClassPathTableMatchType_Regex       Layer7CurCfgContentClassPathTableMatchType = 5
	Layer7CurCfgContentClassPathTableMatchType_Unsupported Layer7CurCfgContentClassPathTableMatchType = 2147483647
)

type Layer7CurCfgContentClassPathTableCase int32

const (
	Layer7CurCfgContentClassPathTableCase_Enabled     Layer7CurCfgContentClassPathTableCase = 1
	Layer7CurCfgContentClassPathTableCase_Disabled    Layer7CurCfgContentClassPathTableCase = 2
	Layer7CurCfgContentClassPathTableCase_Unsupported Layer7CurCfgContentClassPathTableCase = 2147483647
)

type Layer7CurCfgContentClassPathTableParams struct {
	// The content Class Path ID(key id) as an index, length of the string should be 32 characters.
	ContentClassID string `json:"ContentClassID,omitempty"`
	// The content Class ID(key id) as an index, length of the string should be 32 characters.
	ID string `json:"ID,omitempty"`
	// Content Class Path to match, length of the string should be 256 characters.
	FilePath string `json:"FilePath,omitempty"`
	// Path Match type.
	MatchType Layer7CurCfgContentClassPathTableMatchType `json:"MatchType,omitempty"`
	// Enable or Disable Case sensitive for String matching.
	Case Layer7CurCfgContentClassPathTableCase `json:"Case,omitempty"`
	// Id of a data class to associate Path element with.
	DataclassID string `json:"DataclassID,omitempty"`
}

func (p Layer7CurCfgContentClassPathTableParams) iMABean() {}
