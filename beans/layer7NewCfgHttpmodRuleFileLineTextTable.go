package beans

import (
	"fmt"
	"reflect"
)

// Layer7NewCfgHttpmodRuleFileLineTextTable The table for configuring HTTP content modification rules.
// Note:This mib is not supported for VX instance of virtualization.
type Layer7NewCfgHttpmodRuleFileLineTextTable struct {
	// The HTTP Modification Rule List (key id) as an index.
	Layer7NewCfgHttpmodRuleFileLineTextListIdIndex string
	// The HTTP Modification Rule number as an index.
	Layer7NewCfgHttpmodRuleFileLineTextIndex int32
	Params                                   *Layer7NewCfgHttpmodRuleFileLineTextTableParams
}

func NewLayer7NewCfgHttpmodRuleFileLineTextTableList() *Layer7NewCfgHttpmodRuleFileLineTextTable {
	return &Layer7NewCfgHttpmodRuleFileLineTextTable{}
}

func NewLayer7NewCfgHttpmodRuleFileLineTextTable(
	layer7NewCfgHttpmodRuleFileLineTextListIdIndex string,
	layer7NewCfgHttpmodRuleFileLineTextIndex int32,
	params *Layer7NewCfgHttpmodRuleFileLineTextTableParams,
) *Layer7NewCfgHttpmodRuleFileLineTextTable {
	return &Layer7NewCfgHttpmodRuleFileLineTextTable{
		Layer7NewCfgHttpmodRuleFileLineTextListIdIndex: layer7NewCfgHttpmodRuleFileLineTextListIdIndex,
		Layer7NewCfgHttpmodRuleFileLineTextIndex:       layer7NewCfgHttpmodRuleFileLineTextIndex,
		Params:                                         params,
	}
}

func (c *Layer7NewCfgHttpmodRuleFileLineTextTable) Name() string {
	return "Layer7NewCfgHttpmodRuleFileLineTextTable"
}

func (c *Layer7NewCfgHttpmodRuleFileLineTextTable) GetParams() BeanType {
	return c.Params
}

func (c *Layer7NewCfgHttpmodRuleFileLineTextTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *Layer7NewCfgHttpmodRuleFileLineTextTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.Layer7NewCfgHttpmodRuleFileLineTextListIdIndex).IsZero() &&
		reflect.ValueOf(c.Layer7NewCfgHttpmodRuleFileLineTextIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.Layer7NewCfgHttpmodRuleFileLineTextListIdIndex) + "/" + fmt.Sprint(c.Layer7NewCfgHttpmodRuleFileLineTextIndex)
}

type Layer7NewCfgHttpmodRuleFileLineTextTableParams struct {
	// The HTTP Modification Rule List (key id) as an index.
	ListIdIndex string `json:"ListIdIndex,omitempty"`
	// The HTTP Modification Rule number as an index.
	Index int32 `json:"Index,omitempty"`
	// File Type to Replace.
	TypRep string `json:"TypRep,omitempty"`
	// New File Type.
	TypNew string `json:"TypNew,omitempty"`
	// Status Code to Replace.
	StatlineCode int32 `json:"StatlineCode,omitempty"`
	// Status Text to Replace.
	StatlineTxt string `json:"StatlineTxt,omitempty"`
	// New Status Code.
	StatlineNewCode int32 `json:"StatlineNewCode,omitempty"`
	// New Status Text.
	StatlineNewTxt string `json:"StatlineNewTxt,omitempty"`
	// Text to Replace.
	TextReplace string `json:"TextReplace,omitempty"`
	// New Text.
	TextNewText string `json:"TextNewText,omitempty"`
	// Text to Remove.
	TextRemove string `json:"TextRemove,omitempty"`
}

func (p Layer7NewCfgHttpmodRuleFileLineTextTableParams) iMABean() {}
