package beans

import (
	"fmt"
	"reflect"
)

// Layer7NewCfgHttpmodRuleTable The table for configuring HTTP content modification rules.
// Note:This mib is not supported for VX instance of virtualization.
type Layer7NewCfgHttpmodRuleTable struct {
	// The HTTP Modification Rule List (key id) as an index.
	Layer7NewCfgHttpmodRuleListIdIndex string
	// The HTTP Modification Rule number as an index.
	Layer7NewCfgHttpmodRuleIndex int32
	Params                       *Layer7NewCfgHttpmodRuleTableParams
}

func NewLayer7NewCfgHttpmodRuleTableList() *Layer7NewCfgHttpmodRuleTable {
	return &Layer7NewCfgHttpmodRuleTable{}
}

func NewLayer7NewCfgHttpmodRuleTable(
	layer7NewCfgHttpmodRuleListIdIndex string,
	layer7NewCfgHttpmodRuleIndex int32,
	params *Layer7NewCfgHttpmodRuleTableParams,
) *Layer7NewCfgHttpmodRuleTable {
	return &Layer7NewCfgHttpmodRuleTable{
		Layer7NewCfgHttpmodRuleListIdIndex: layer7NewCfgHttpmodRuleListIdIndex,
		Layer7NewCfgHttpmodRuleIndex:       layer7NewCfgHttpmodRuleIndex,
		Params:                             params,
	}
}

func (c *Layer7NewCfgHttpmodRuleTable) Name() string {
	return "Layer7NewCfgHttpmodRuleTable"
}

func (c *Layer7NewCfgHttpmodRuleTable) GetParams() BeanType {
	return c.Params
}

func (c *Layer7NewCfgHttpmodRuleTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *Layer7NewCfgHttpmodRuleTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.Layer7NewCfgHttpmodRuleListIdIndex).IsZero() &&
		reflect.ValueOf(c.Layer7NewCfgHttpmodRuleIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.Layer7NewCfgHttpmodRuleListIdIndex) + "/" + fmt.Sprint(c.Layer7NewCfgHttpmodRuleIndex)
}

type Layer7NewCfgHttpmodRuleTableDirectn int32

const (
	Layer7NewCfgHttpmodRuleTableDirectn_Request       Layer7NewCfgHttpmodRuleTableDirectn = 1
	Layer7NewCfgHttpmodRuleTableDirectn_Response      Layer7NewCfgHttpmodRuleTableDirectn = 2
	Layer7NewCfgHttpmodRuleTableDirectn_Bidirectional Layer7NewCfgHttpmodRuleTableDirectn = 3
	Layer7NewCfgHttpmodRuleTableDirectn_Unsupported   Layer7NewCfgHttpmodRuleTableDirectn = 2147483647
)

type Layer7NewCfgHttpmodRuleTableAction int32

const (
	Layer7NewCfgHttpmodRuleTableAction_Insert      Layer7NewCfgHttpmodRuleTableAction = 1
	Layer7NewCfgHttpmodRuleTableAction_Replace     Layer7NewCfgHttpmodRuleTableAction = 2
	Layer7NewCfgHttpmodRuleTableAction_Remove      Layer7NewCfgHttpmodRuleTableAction = 3
	Layer7NewCfgHttpmodRuleTableAction_None        Layer7NewCfgHttpmodRuleTableAction = 4
	Layer7NewCfgHttpmodRuleTableAction_Unsupported Layer7NewCfgHttpmodRuleTableAction = 2147483647
)

type Layer7NewCfgHttpmodRuleTableAdminStatus int32

const (
	Layer7NewCfgHttpmodRuleTableAdminStatus_Enable      Layer7NewCfgHttpmodRuleTableAdminStatus = 1
	Layer7NewCfgHttpmodRuleTableAdminStatus_Disable     Layer7NewCfgHttpmodRuleTableAdminStatus = 2
	Layer7NewCfgHttpmodRuleTableAdminStatus_Unsupported Layer7NewCfgHttpmodRuleTableAdminStatus = 2147483647
)

type Layer7NewCfgHttpmodRuleTableDelete int32

const (
	Layer7NewCfgHttpmodRuleTableDelete_Other       Layer7NewCfgHttpmodRuleTableDelete = 1
	Layer7NewCfgHttpmodRuleTableDelete_Delete      Layer7NewCfgHttpmodRuleTableDelete = 2
	Layer7NewCfgHttpmodRuleTableDelete_Unsupported Layer7NewCfgHttpmodRuleTableDelete = 2147483647
)

type Layer7NewCfgHttpmodRuleTableElement int32

const (
	Layer7NewCfgHttpmodRuleTableElement_Url         Layer7NewCfgHttpmodRuleTableElement = 1
	Layer7NewCfgHttpmodRuleTableElement_Header      Layer7NewCfgHttpmodRuleTableElement = 2
	Layer7NewCfgHttpmodRuleTableElement_Cookie      Layer7NewCfgHttpmodRuleTableElement = 3
	Layer7NewCfgHttpmodRuleTableElement_Filetype    Layer7NewCfgHttpmodRuleTableElement = 4
	Layer7NewCfgHttpmodRuleTableElement_Statusline  Layer7NewCfgHttpmodRuleTableElement = 5
	Layer7NewCfgHttpmodRuleTableElement_Text        Layer7NewCfgHttpmodRuleTableElement = 6
	Layer7NewCfgHttpmodRuleTableElement_Unsupported Layer7NewCfgHttpmodRuleTableElement = 2147483647
)

type Layer7NewCfgHttpmodRuleTableHttpBody int32

const (
	Layer7NewCfgHttpmodRuleTableHttpBody_Include     Layer7NewCfgHttpmodRuleTableHttpBody = 1
	Layer7NewCfgHttpmodRuleTableHttpBody_Exclude     Layer7NewCfgHttpmodRuleTableHttpBody = 2
	Layer7NewCfgHttpmodRuleTableHttpBody_Unsupported Layer7NewCfgHttpmodRuleTableHttpBody = 2147483647
)

type Layer7NewCfgHttpmodRuleTableParams struct {
	// The HTTP Modification Rule List (key id) as an index.
	ListIdIndex string `json:"ListIdIndex,omitempty"`
	// The HTTP Modification Rule number as an index.
	Index int32 `json:"Index,omitempty"`
	// The HTTP Modification Rule name.
	Name string `json:"Name,omitempty"`
	// Defines whether rule modification direction should be evaluated as Request or Response (default Response). SET layer7NewCfgHttpmodRuleElementbefore SETTING the direction.
	Directn Layer7NewCfgHttpmodRuleTableDirectn `json:"Directn,omitempty"`
	// Defines rule action (default none).
	Action Layer7NewCfgHttpmodRuleTableAction `json:"Action,omitempty"`
	// Status (enable/disable) of rule.
	AdminStatus Layer7NewCfgHttpmodRuleTableAdminStatus `json:"AdminStatus,omitempty"`
	// Copy rules.
	Copy DisplayString `json:"Copy,omitempty"`
	// By setting the value to delete(2), the entire row is deleted.
	Delete Layer7NewCfgHttpmodRuleTableDelete `json:"Delete,omitempty"`
	// Element to be modified.
	Element Layer7NewCfgHttpmodRuleTableElement `json:"Element,omitempty"`
	// Modifications to also be done in the HTTP body.
	HttpBody Layer7NewCfgHttpmodRuleTableHttpBody `json:"HttpBody,omitempty"`
}

func (p Layer7NewCfgHttpmodRuleTableParams) iMABean() {}
