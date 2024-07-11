package beans

import (
	"fmt"
	"reflect"
)

// Layer7CurCfgHttpmodRuleFileLineTextTable The table for configuring HTTP content modification rules.
// Note:This mib is not supported for VX instance of virtualization.
type Layer7CurCfgHttpmodRuleFileLineTextTable struct {
	// The HTTP Modification Rule List (key id) as an index.
	Layer7CurCfgHttpmodRuleFileLineTextListIdIndex string
	// The HTTP Modification Rule number as an index.
	Layer7CurCfgHttpmodRuleFileLineTextIndex int32
	Params                                   *Layer7CurCfgHttpmodRuleFileLineTextTableParams
}

func NewLayer7CurCfgHttpmodRuleFileLineTextTableList() *Layer7CurCfgHttpmodRuleFileLineTextTable {
	return &Layer7CurCfgHttpmodRuleFileLineTextTable{}
}

func NewLayer7CurCfgHttpmodRuleFileLineTextTable(
	layer7CurCfgHttpmodRuleFileLineTextListIdIndex string,
	layer7CurCfgHttpmodRuleFileLineTextIndex int32,
	params *Layer7CurCfgHttpmodRuleFileLineTextTableParams,
) *Layer7CurCfgHttpmodRuleFileLineTextTable {
	return &Layer7CurCfgHttpmodRuleFileLineTextTable{
		Layer7CurCfgHttpmodRuleFileLineTextListIdIndex: layer7CurCfgHttpmodRuleFileLineTextListIdIndex,
		Layer7CurCfgHttpmodRuleFileLineTextIndex:       layer7CurCfgHttpmodRuleFileLineTextIndex,
		Params:                                         params,
	}
}

func (c *Layer7CurCfgHttpmodRuleFileLineTextTable) Name() string {
	return "Layer7CurCfgHttpmodRuleFileLineTextTable"
}

func (c *Layer7CurCfgHttpmodRuleFileLineTextTable) GetParams() BeanType {
	return c.Params
}

func (c *Layer7CurCfgHttpmodRuleFileLineTextTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *Layer7CurCfgHttpmodRuleFileLineTextTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.Layer7CurCfgHttpmodRuleFileLineTextListIdIndex).IsZero() &&
		reflect.ValueOf(c.Layer7CurCfgHttpmodRuleFileLineTextIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.Layer7CurCfgHttpmodRuleFileLineTextListIdIndex) + "/" + fmt.Sprint(c.Layer7CurCfgHttpmodRuleFileLineTextIndex)
}

type Layer7CurCfgHttpmodRuleFileLineTextTableParams struct {
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

func (p Layer7CurCfgHttpmodRuleFileLineTextTableParams) iMABean() {}
