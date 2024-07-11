package beans

import (
	"fmt"
	"reflect"
)

// Layer7NewCfgHttpmodRuleHdrTable The table for configuring HTTP content modification rules.
// Note:This mib is not supported for VX instance of virtualization.
type Layer7NewCfgHttpmodRuleHdrTable struct {
	// The HTTP Modification Rule List (key id) as an index.
	Layer7NewCfgHttpmodRuleHdrListIdIndex string
	// The HTTP Modification Rule number as an index.
	Layer7NewCfgHttpmodRuleHdrIndex int32
	Params                          *Layer7NewCfgHttpmodRuleHdrTableParams
}

func NewLayer7NewCfgHttpmodRuleHdrTableList() *Layer7NewCfgHttpmodRuleHdrTable {
	return &Layer7NewCfgHttpmodRuleHdrTable{}
}

func NewLayer7NewCfgHttpmodRuleHdrTable(
	layer7NewCfgHttpmodRuleHdrListIdIndex string,
	layer7NewCfgHttpmodRuleHdrIndex int32,
	params *Layer7NewCfgHttpmodRuleHdrTableParams,
) *Layer7NewCfgHttpmodRuleHdrTable {
	return &Layer7NewCfgHttpmodRuleHdrTable{
		Layer7NewCfgHttpmodRuleHdrListIdIndex: layer7NewCfgHttpmodRuleHdrListIdIndex,
		Layer7NewCfgHttpmodRuleHdrIndex:       layer7NewCfgHttpmodRuleHdrIndex,
		Params:                                params,
	}
}

func (c *Layer7NewCfgHttpmodRuleHdrTable) Name() string {
	return "Layer7NewCfgHttpmodRuleHdrTable"
}

func (c *Layer7NewCfgHttpmodRuleHdrTable) GetParams() BeanType {
	return c.Params
}

func (c *Layer7NewCfgHttpmodRuleHdrTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *Layer7NewCfgHttpmodRuleHdrTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.Layer7NewCfgHttpmodRuleHdrListIdIndex).IsZero() &&
		reflect.ValueOf(c.Layer7NewCfgHttpmodRuleHdrIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.Layer7NewCfgHttpmodRuleHdrListIdIndex) + "/" + fmt.Sprint(c.Layer7NewCfgHttpmodRuleHdrIndex)
}

type Layer7NewCfgHttpmodRuleHdrTableElmnt int32

const (
	Layer7NewCfgHttpmodRuleHdrTableElmnt_Url         Layer7NewCfgHttpmodRuleHdrTableElmnt = 1
	Layer7NewCfgHttpmodRuleHdrTableElmnt_Header      Layer7NewCfgHttpmodRuleHdrTableElmnt = 2
	Layer7NewCfgHttpmodRuleHdrTableElmnt_Cookie      Layer7NewCfgHttpmodRuleHdrTableElmnt = 3
	Layer7NewCfgHttpmodRuleHdrTableElmnt_Filetype    Layer7NewCfgHttpmodRuleHdrTableElmnt = 4
	Layer7NewCfgHttpmodRuleHdrTableElmnt_Statusline  Layer7NewCfgHttpmodRuleHdrTableElmnt = 5
	Layer7NewCfgHttpmodRuleHdrTableElmnt_Text        Layer7NewCfgHttpmodRuleHdrTableElmnt = 6
	Layer7NewCfgHttpmodRuleHdrTableElmnt_Regex       Layer7NewCfgHttpmodRuleHdrTableElmnt = 7
	Layer7NewCfgHttpmodRuleHdrTableElmnt_None        Layer7NewCfgHttpmodRuleHdrTableElmnt = 8
	Layer7NewCfgHttpmodRuleHdrTableElmnt_Unsupported Layer7NewCfgHttpmodRuleHdrTableElmnt = 2147483647
)

type Layer7NewCfgHttpmodRuleHdrTableParams struct {
	// The HTTP Modification Rule List (key id) as an index.
	ListIdIndex string `json:"ListIdIndex,omitempty"`
	// The HTTP Modification Rule number as an index.
	Index int32 `json:"Index,omitempty"`
	// Header Field to Insert.
	Insert string `json:"Insert,omitempty"`
	// Value to Insert.
	Value string `json:"Value,omitempty"`
	// Element to match.
	Elmnt Layer7NewCfgHttpmodRuleHdrTableElmnt `json:"Elmnt,omitempty"`
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

func (p Layer7NewCfgHttpmodRuleHdrTableParams) iMABean() {}
