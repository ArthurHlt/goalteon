package beans

import (
	"fmt"
	"reflect"
)

// Layer7CurCfgHttpmodRuleHdrTable The table for configuring HTTP content modification rules.
// Note:This mib is not supported for VX instance of virtualization.
type Layer7CurCfgHttpmodRuleHdrTable struct {
	// The HTTP Modification Rule List (key id) as an index.
	Layer7CurCfgHttpmodRuleHdrListIdIndex string
	// The HTTP Modification Rule number as an index.
	Layer7CurCfgHttpmodRuleHdrIndex int32
	Params                          *Layer7CurCfgHttpmodRuleHdrTableParams
}

func NewLayer7CurCfgHttpmodRuleHdrTableList() *Layer7CurCfgHttpmodRuleHdrTable {
	return &Layer7CurCfgHttpmodRuleHdrTable{}
}

func NewLayer7CurCfgHttpmodRuleHdrTable(
	layer7CurCfgHttpmodRuleHdrListIdIndex string,
	layer7CurCfgHttpmodRuleHdrIndex int32,
	params *Layer7CurCfgHttpmodRuleHdrTableParams,
) *Layer7CurCfgHttpmodRuleHdrTable {
	return &Layer7CurCfgHttpmodRuleHdrTable{
		Layer7CurCfgHttpmodRuleHdrListIdIndex: layer7CurCfgHttpmodRuleHdrListIdIndex,
		Layer7CurCfgHttpmodRuleHdrIndex:       layer7CurCfgHttpmodRuleHdrIndex,
		Params:                                params,
	}
}

func (c *Layer7CurCfgHttpmodRuleHdrTable) Name() string {
	return "Layer7CurCfgHttpmodRuleHdrTable"
}

func (c *Layer7CurCfgHttpmodRuleHdrTable) GetParams() BeanType {
	return c.Params
}

func (c *Layer7CurCfgHttpmodRuleHdrTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *Layer7CurCfgHttpmodRuleHdrTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.Layer7CurCfgHttpmodRuleHdrListIdIndex).IsZero() &&
		reflect.ValueOf(c.Layer7CurCfgHttpmodRuleHdrIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.Layer7CurCfgHttpmodRuleHdrListIdIndex) + "/" + fmt.Sprint(c.Layer7CurCfgHttpmodRuleHdrIndex)
}

type Layer7CurCfgHttpmodRuleHdrTableElmnt int32

const (
	Layer7CurCfgHttpmodRuleHdrTableElmnt_Url         Layer7CurCfgHttpmodRuleHdrTableElmnt = 1
	Layer7CurCfgHttpmodRuleHdrTableElmnt_Header      Layer7CurCfgHttpmodRuleHdrTableElmnt = 2
	Layer7CurCfgHttpmodRuleHdrTableElmnt_Cookie      Layer7CurCfgHttpmodRuleHdrTableElmnt = 3
	Layer7CurCfgHttpmodRuleHdrTableElmnt_Filetype    Layer7CurCfgHttpmodRuleHdrTableElmnt = 4
	Layer7CurCfgHttpmodRuleHdrTableElmnt_Statusline  Layer7CurCfgHttpmodRuleHdrTableElmnt = 5
	Layer7CurCfgHttpmodRuleHdrTableElmnt_Text        Layer7CurCfgHttpmodRuleHdrTableElmnt = 6
	Layer7CurCfgHttpmodRuleHdrTableElmnt_Regex       Layer7CurCfgHttpmodRuleHdrTableElmnt = 7
	Layer7CurCfgHttpmodRuleHdrTableElmnt_None        Layer7CurCfgHttpmodRuleHdrTableElmnt = 8
	Layer7CurCfgHttpmodRuleHdrTableElmnt_Unsupported Layer7CurCfgHttpmodRuleHdrTableElmnt = 2147483647
)

type Layer7CurCfgHttpmodRuleHdrTableParams struct {
	// The HTTP Modification Rule List (key id) as an index.
	ListIdIndex string `json:"ListIdIndex,omitempty"`
	// The HTTP Modification Rule number as an index.
	Index int32 `json:"Index,omitempty"`
	// Header Field to Insert.
	Insert string `json:"Insert,omitempty"`
	// Value to Insert.
	Value string `json:"Value,omitempty"`
	// Element to match.
	Elmnt Layer7CurCfgHttpmodRuleHdrTableElmnt `json:"Elmnt,omitempty"`
	// Host to match.
	ElmntUrlHost string `json:"ElmntUrlHost,omitempty"`
	// Path to match.
	ElmntUrlPath string `json:"ElmntUrlPath,omitempty"`
	// Header Field to Match.
	ElmntHdrField string `json:"ElmntHdrField,omitempty"`
	// Value to Match.
	ElmntHdrVal string `json:"ElmntHdrVal,omitempty"`
	// Cookie Key to Match.
	ElmntCookey string `json:"ElmntCookey,omitempty"`
	// Cookie Value to Match.
	ElmntCkieVal string `json:"ElmntCkieVal,omitempty"`
	// File Type To Match.
	ElmntFileTyp string `json:"ElmntFileTyp,omitempty"`
	// Status Code to Match.
	ElmntStatusCode int32 `json:"ElmntStatusCode,omitempty"`
	// Status Text to Match.
	ElmntStatusTxt string `json:"ElmntStatusTxt,omitempty"`
	// Text to Match.
	ElmntTxt string `json:"ElmntTxt,omitempty"`
	// Regex to Match.
	ElmntRegx string `json:"ElmntRegx,omitempty"`
	// Header Field to Replace.
	ReplacHdr string `json:"ReplacHdr,omitempty"`
	// Value to Replace.
	ReplacVal string `json:"ReplacVal,omitempty"`
	// New Header Field.
	ReplacNewHdr string `json:"ReplacNewHdr,omitempty"`
	// New Value.
	ReplacNewVal string `json:"ReplacNewVal,omitempty"`
	// Header Field to Remove.
	RemvHdr string `json:"RemvHdr,omitempty"`
	// Value to Remove.
	RemvVal string `json:"RemvVal,omitempty"`
}

func (p Layer7CurCfgHttpmodRuleHdrTableParams) iMABean() {}
