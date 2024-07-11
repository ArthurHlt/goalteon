package beans

import (
	"fmt"
	"reflect"
)

// Layer7CurCfgHttpmodRuleTable The table for configuring HTTP content modification rules.
// Note:This mib is not supported for VX instance of virtualization.
type Layer7CurCfgHttpmodRuleTable struct {
	// The HTTP Modification Rule List (key id) as an index.
	Layer7CurCfgHttpmodRuleListIdIndex string
	// The HTTP Modification Rule number as an index.
	Layer7CurCfgHttpmodRuleIndex int32
	Params                       *Layer7CurCfgHttpmodRuleTableParams
}

func NewLayer7CurCfgHttpmodRuleTableList() *Layer7CurCfgHttpmodRuleTable {
	return &Layer7CurCfgHttpmodRuleTable{}
}

func NewLayer7CurCfgHttpmodRuleTable(
	layer7CurCfgHttpmodRuleListIdIndex string,
	layer7CurCfgHttpmodRuleIndex int32,
	params *Layer7CurCfgHttpmodRuleTableParams,
) *Layer7CurCfgHttpmodRuleTable {
	return &Layer7CurCfgHttpmodRuleTable{
		Layer7CurCfgHttpmodRuleListIdIndex: layer7CurCfgHttpmodRuleListIdIndex,
		Layer7CurCfgHttpmodRuleIndex:       layer7CurCfgHttpmodRuleIndex,
		Params:                             params,
	}
}

func (c *Layer7CurCfgHttpmodRuleTable) Name() string {
	return "Layer7CurCfgHttpmodRuleTable"
}

func (c *Layer7CurCfgHttpmodRuleTable) GetParams() BeanType {
	return c.Params
}

func (c *Layer7CurCfgHttpmodRuleTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *Layer7CurCfgHttpmodRuleTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.Layer7CurCfgHttpmodRuleListIdIndex).IsZero() &&
		reflect.ValueOf(c.Layer7CurCfgHttpmodRuleIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.Layer7CurCfgHttpmodRuleListIdIndex) + "/" + fmt.Sprint(c.Layer7CurCfgHttpmodRuleIndex)
}

type Layer7CurCfgHttpmodRuleTableDirectn int32

const (
	Layer7CurCfgHttpmodRuleTableDirectn_Request       Layer7CurCfgHttpmodRuleTableDirectn = 1
	Layer7CurCfgHttpmodRuleTableDirectn_Response      Layer7CurCfgHttpmodRuleTableDirectn = 2
	Layer7CurCfgHttpmodRuleTableDirectn_Bidirectional Layer7CurCfgHttpmodRuleTableDirectn = 3
	Layer7CurCfgHttpmodRuleTableDirectn_Unsupported   Layer7CurCfgHttpmodRuleTableDirectn = 2147483647
)

type Layer7CurCfgHttpmodRuleTableAction int32

const (
	Layer7CurCfgHttpmodRuleTableAction_Insert      Layer7CurCfgHttpmodRuleTableAction = 1
	Layer7CurCfgHttpmodRuleTableAction_Replace     Layer7CurCfgHttpmodRuleTableAction = 2
	Layer7CurCfgHttpmodRuleTableAction_Remove      Layer7CurCfgHttpmodRuleTableAction = 3
	Layer7CurCfgHttpmodRuleTableAction_None        Layer7CurCfgHttpmodRuleTableAction = 4
	Layer7CurCfgHttpmodRuleTableAction_Unsupported Layer7CurCfgHttpmodRuleTableAction = 2147483647
)

type Layer7CurCfgHttpmodRuleTableAdminStatus int32

const (
	Layer7CurCfgHttpmodRuleTableAdminStatus_Enable      Layer7CurCfgHttpmodRuleTableAdminStatus = 1
	Layer7CurCfgHttpmodRuleTableAdminStatus_Disable     Layer7CurCfgHttpmodRuleTableAdminStatus = 2
	Layer7CurCfgHttpmodRuleTableAdminStatus_Unsupported Layer7CurCfgHttpmodRuleTableAdminStatus = 2147483647
)

type Layer7CurCfgHttpmodRuleTableElement int32

const (
	Layer7CurCfgHttpmodRuleTableElement_Url         Layer7CurCfgHttpmodRuleTableElement = 1
	Layer7CurCfgHttpmodRuleTableElement_Header      Layer7CurCfgHttpmodRuleTableElement = 2
	Layer7CurCfgHttpmodRuleTableElement_Cookie      Layer7CurCfgHttpmodRuleTableElement = 3
	Layer7CurCfgHttpmodRuleTableElement_Filetype    Layer7CurCfgHttpmodRuleTableElement = 4
	Layer7CurCfgHttpmodRuleTableElement_Statusline  Layer7CurCfgHttpmodRuleTableElement = 5
	Layer7CurCfgHttpmodRuleTableElement_Text        Layer7CurCfgHttpmodRuleTableElement = 6
	Layer7CurCfgHttpmodRuleTableElement_Unsupported Layer7CurCfgHttpmodRuleTableElement = 2147483647
)

type Layer7CurCfgHttpmodRuleTableHttpBody int32

const (
	Layer7CurCfgHttpmodRuleTableHttpBody_Include     Layer7CurCfgHttpmodRuleTableHttpBody = 1
	Layer7CurCfgHttpmodRuleTableHttpBody_Exclude     Layer7CurCfgHttpmodRuleTableHttpBody = 2
	Layer7CurCfgHttpmodRuleTableHttpBody_Unsupported Layer7CurCfgHttpmodRuleTableHttpBody = 2147483647
)

type Layer7CurCfgHttpmodRuleTableParams struct {
	// The HTTP Modification Rule List (key id) as an index.
	ListIdIndex string `json:"ListIdIndex,omitempty"`
	// The HTTP Modification Rule number as an index.
	Index int32 `json:"Index,omitempty"`
	// The HTTP Modification Rule name.
	Name string `json:"Name,omitempty"`
	// Defines whether rule modification direction should be evaluated as Request or Response (default Request).
	Directn Layer7CurCfgHttpmodRuleTableDirectn `json:"Directn,omitempty"`
	// Defines rule action (default none).
	Action Layer7CurCfgHttpmodRuleTableAction `json:"Action,omitempty"`
	// Status (enable/disable) of rule.
	AdminStatus Layer7CurCfgHttpmodRuleTableAdminStatus `json:"AdminStatus,omitempty"`
	// Element to be modified.
	Element Layer7CurCfgHttpmodRuleTableElement `json:"Element,omitempty"`
	// Modifications to also be done in the HTTP body.
	HttpBody Layer7CurCfgHttpmodRuleTableHttpBody `json:"HttpBody,omitempty"`
}

func (p Layer7CurCfgHttpmodRuleTableParams) iMABean() {}
