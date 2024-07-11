package beans

import (
	"fmt"
	"reflect"
)

// Layer7CurCfgHttpmodRuleCookieTable The table for configuring HTTP content modification rules.
// Note:This mib is not supported for VX instance of virtualization.
type Layer7CurCfgHttpmodRuleCookieTable struct {
	// The HTTP Modification Rule List (key id) as an index.
	Layer7CurCfgHttpmodRuleCookieListIdIndex string
	// The HTTP Modification Rule number as an index.
	Layer7CurCfgHttpmodRuleCookieIndex int32
	Params                             *Layer7CurCfgHttpmodRuleCookieTableParams
}

func NewLayer7CurCfgHttpmodRuleCookieTableList() *Layer7CurCfgHttpmodRuleCookieTable {
	return &Layer7CurCfgHttpmodRuleCookieTable{}
}

func NewLayer7CurCfgHttpmodRuleCookieTable(
	layer7CurCfgHttpmodRuleCookieListIdIndex string,
	layer7CurCfgHttpmodRuleCookieIndex int32,
	params *Layer7CurCfgHttpmodRuleCookieTableParams,
) *Layer7CurCfgHttpmodRuleCookieTable {
	return &Layer7CurCfgHttpmodRuleCookieTable{
		Layer7CurCfgHttpmodRuleCookieListIdIndex: layer7CurCfgHttpmodRuleCookieListIdIndex,
		Layer7CurCfgHttpmodRuleCookieIndex:       layer7CurCfgHttpmodRuleCookieIndex,
		Params:                                   params,
	}
}

func (c *Layer7CurCfgHttpmodRuleCookieTable) Name() string {
	return "Layer7CurCfgHttpmodRuleCookieTable"
}

func (c *Layer7CurCfgHttpmodRuleCookieTable) GetParams() BeanType {
	return c.Params
}

func (c *Layer7CurCfgHttpmodRuleCookieTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *Layer7CurCfgHttpmodRuleCookieTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.Layer7CurCfgHttpmodRuleCookieListIdIndex).IsZero() &&
		reflect.ValueOf(c.Layer7CurCfgHttpmodRuleCookieIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.Layer7CurCfgHttpmodRuleCookieListIdIndex) + "/" + fmt.Sprint(c.Layer7CurCfgHttpmodRuleCookieIndex)
}

type Layer7CurCfgHttpmodRuleCookieTableInsrtElem int32

const (
	Layer7CurCfgHttpmodRuleCookieTableInsrtElem_Url         Layer7CurCfgHttpmodRuleCookieTableInsrtElem = 1
	Layer7CurCfgHttpmodRuleCookieTableInsrtElem_Header      Layer7CurCfgHttpmodRuleCookieTableInsrtElem = 2
	Layer7CurCfgHttpmodRuleCookieTableInsrtElem_Cookie      Layer7CurCfgHttpmodRuleCookieTableInsrtElem = 3
	Layer7CurCfgHttpmodRuleCookieTableInsrtElem_Filetype    Layer7CurCfgHttpmodRuleCookieTableInsrtElem = 4
	Layer7CurCfgHttpmodRuleCookieTableInsrtElem_Statusline  Layer7CurCfgHttpmodRuleCookieTableInsrtElem = 5
	Layer7CurCfgHttpmodRuleCookieTableInsrtElem_Text        Layer7CurCfgHttpmodRuleCookieTableInsrtElem = 6
	Layer7CurCfgHttpmodRuleCookieTableInsrtElem_Regex       Layer7CurCfgHttpmodRuleCookieTableInsrtElem = 7
	Layer7CurCfgHttpmodRuleCookieTableInsrtElem_None        Layer7CurCfgHttpmodRuleCookieTableInsrtElem = 8
	Layer7CurCfgHttpmodRuleCookieTableInsrtElem_Unsupported Layer7CurCfgHttpmodRuleCookieTableInsrtElem = 2147483647
)

type Layer7CurCfgHttpmodRuleCookieTableParams struct {
	// The HTTP Modification Rule List (key id) as an index.
	ListIdIndex string `json:"ListIdIndex,omitempty"`
	// The HTTP Modification Rule number as an index.
	Index int32 `json:"Index,omitempty"`
	// Cookie Key to Insert.
	InsrtKey string `json:"InsrtKey,omitempty"`
	// Cookie Value to Insert.
	InsrtVal string `json:"InsrtVal,omitempty"`
	// Cookie Path to Insert.
	InsrtPath string `json:"InsrtPath,omitempty"`
	// Cookie Domain Name to Insert.
	InsrtDomn string `json:"InsrtDomn,omitempty"`
	// Cookie Expiration Time to Insert.
	InsrtExp string `json:"InsrtExp,omitempty"`
	// Element to match.
	InsrtElem Layer7CurCfgHttpmodRuleCookieTableInsrtElem `json:"InsrtElem,omitempty"`
	// Host to Match.
	InsrtUrlHost string `json:"InsrtUrlHost,omitempty"`
	// Path to Match.
	InsrtUrlPath string `json:"InsrtUrlPath,omitempty"`
	// Header Field to Match.
	InsrtHdrFld string `json:"InsrtHdrFld,omitempty"`
	// Value to Match.
	InsrtHdrVal string `json:"InsrtHdrVal,omitempty"`
	// Cookie Key to Match.
	InsrtCookey string `json:"InsrtCookey,omitempty"`
	// Cookie Value to Match.
	InsrtCookieVal string `json:"InsrtCookieVal,omitempty"`
	// File Type To Match.
	InsrtFiletyp string `json:"InsrtFiletyp,omitempty"`
	// Status Code to Match.
	InsrtStatsCode int32 `json:"InsrtStatsCode,omitempty"`
	// Status Text to Match.
	InsrtStatsTxt string `json:"InsrtStatsTxt,omitempty"`
	// Text to Match.
	InsrtTxt string `json:"InsrtTxt,omitempty"`
	// Regex to Match.
	InsrtRegx string `json:"InsrtRegx,omitempty"`
	// Cookie Key to Replace.
	ReplcCookey string `json:"ReplcCookey,omitempty"`
	// Cookie Value to Replace.
	ReplcVal string `json:"ReplcVal,omitempty"`
	// New Cookie Key.
	ReplcNewKey string `json:"ReplcNewKey,omitempty"`
	// New Cookie Value.
	ReplcNewVal string `json:"ReplcNewVal,omitempty"`
	// Cookie Key to Remove.
	RemvCookey string `json:"RemvCookey,omitempty"`
	// Cookie Value to Remove.
	RemvCookieVal string `json:"RemvCookieVal,omitempty"`
}

func (p Layer7CurCfgHttpmodRuleCookieTableParams) iMABean() {}
