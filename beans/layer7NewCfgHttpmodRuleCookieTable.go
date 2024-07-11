package beans

import (
	"fmt"
	"reflect"
)

// Layer7NewCfgHttpmodRuleCookieTable The table for configuring HTTP content modification rules.
// Note:This mib is not supported for VX instance of virtualization.
type Layer7NewCfgHttpmodRuleCookieTable struct {
	// The HTTP Modification Rule List (key id) as an index.
	Layer7NewCfgHttpmodRuleCookieListIdIndex string
	// The HTTP Modification Rule number as an index.
	Layer7NewCfgHttpmodRuleCookieIndex int32
	Params                             *Layer7NewCfgHttpmodRuleCookieTableParams
}

func NewLayer7NewCfgHttpmodRuleCookieTableList() *Layer7NewCfgHttpmodRuleCookieTable {
	return &Layer7NewCfgHttpmodRuleCookieTable{}
}

func NewLayer7NewCfgHttpmodRuleCookieTable(
	layer7NewCfgHttpmodRuleCookieListIdIndex string,
	layer7NewCfgHttpmodRuleCookieIndex int32,
	params *Layer7NewCfgHttpmodRuleCookieTableParams,
) *Layer7NewCfgHttpmodRuleCookieTable {
	return &Layer7NewCfgHttpmodRuleCookieTable{
		Layer7NewCfgHttpmodRuleCookieListIdIndex: layer7NewCfgHttpmodRuleCookieListIdIndex,
		Layer7NewCfgHttpmodRuleCookieIndex:       layer7NewCfgHttpmodRuleCookieIndex,
		Params:                                   params,
	}
}

func (c *Layer7NewCfgHttpmodRuleCookieTable) Name() string {
	return "Layer7NewCfgHttpmodRuleCookieTable"
}

func (c *Layer7NewCfgHttpmodRuleCookieTable) GetParams() BeanType {
	return c.Params
}

func (c *Layer7NewCfgHttpmodRuleCookieTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *Layer7NewCfgHttpmodRuleCookieTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.Layer7NewCfgHttpmodRuleCookieListIdIndex).IsZero() &&
		reflect.ValueOf(c.Layer7NewCfgHttpmodRuleCookieIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.Layer7NewCfgHttpmodRuleCookieListIdIndex) + "/" + fmt.Sprint(c.Layer7NewCfgHttpmodRuleCookieIndex)
}

type Layer7NewCfgHttpmodRuleCookieTableInsrtElem int32

const (
	Layer7NewCfgHttpmodRuleCookieTableInsrtElem_Url         Layer7NewCfgHttpmodRuleCookieTableInsrtElem = 1
	Layer7NewCfgHttpmodRuleCookieTableInsrtElem_Header      Layer7NewCfgHttpmodRuleCookieTableInsrtElem = 2
	Layer7NewCfgHttpmodRuleCookieTableInsrtElem_Cookie      Layer7NewCfgHttpmodRuleCookieTableInsrtElem = 3
	Layer7NewCfgHttpmodRuleCookieTableInsrtElem_Filetype    Layer7NewCfgHttpmodRuleCookieTableInsrtElem = 4
	Layer7NewCfgHttpmodRuleCookieTableInsrtElem_Statusline  Layer7NewCfgHttpmodRuleCookieTableInsrtElem = 5
	Layer7NewCfgHttpmodRuleCookieTableInsrtElem_Text        Layer7NewCfgHttpmodRuleCookieTableInsrtElem = 6
	Layer7NewCfgHttpmodRuleCookieTableInsrtElem_Regex       Layer7NewCfgHttpmodRuleCookieTableInsrtElem = 7
	Layer7NewCfgHttpmodRuleCookieTableInsrtElem_None        Layer7NewCfgHttpmodRuleCookieTableInsrtElem = 8
	Layer7NewCfgHttpmodRuleCookieTableInsrtElem_Unsupported Layer7NewCfgHttpmodRuleCookieTableInsrtElem = 2147483647
)

type Layer7NewCfgHttpmodRuleCookieTableParams struct {
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
	InsrtElem Layer7NewCfgHttpmodRuleCookieTableInsrtElem `json:"InsrtElem,omitempty"`
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

func (p Layer7NewCfgHttpmodRuleCookieTableParams) iMABean() {}
