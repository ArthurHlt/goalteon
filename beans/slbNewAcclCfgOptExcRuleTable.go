package beans

import (
	"fmt"
	"reflect"
)

// SlbNewAcclCfgOptExcRuleTable The table for configuring optimization Rules.
// Note:This mib is not supported for VX instance of Virtualization.
type SlbNewAcclCfgOptExcRuleTable struct {
	// The optimization rule list (key id) as an index.
	SlbNewAcclCfgOptExcRuleListIdIndex string
	// The optimization Rule number as an index.
	SlbNewAcclCfgOptExcRuleIndex int32
	Params                       *SlbNewAcclCfgOptExcRuleTableParams
}

func NewSlbNewAcclCfgOptExcRuleTableList() *SlbNewAcclCfgOptExcRuleTable {
	return &SlbNewAcclCfgOptExcRuleTable{}
}

func NewSlbNewAcclCfgOptExcRuleTable(
	slbNewAcclCfgOptExcRuleListIdIndex string,
	slbNewAcclCfgOptExcRuleIndex int32,
	params *SlbNewAcclCfgOptExcRuleTableParams,
) *SlbNewAcclCfgOptExcRuleTable {
	return &SlbNewAcclCfgOptExcRuleTable{
		SlbNewAcclCfgOptExcRuleListIdIndex: slbNewAcclCfgOptExcRuleListIdIndex,
		SlbNewAcclCfgOptExcRuleIndex:       slbNewAcclCfgOptExcRuleIndex,
		Params:                             params,
	}
}

func (c *SlbNewAcclCfgOptExcRuleTable) Name() string {
	return "SlbNewAcclCfgOptExcRuleTable"
}

func (c *SlbNewAcclCfgOptExcRuleTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbNewAcclCfgOptExcRuleTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbNewAcclCfgOptExcRuleTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbNewAcclCfgOptExcRuleListIdIndex).IsZero() &&
		reflect.ValueOf(c.SlbNewAcclCfgOptExcRuleIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbNewAcclCfgOptExcRuleListIdIndex) + "/" + fmt.Sprint(c.SlbNewAcclCfgOptExcRuleIndex)
}

type SlbNewAcclCfgOptExcRuleTableResType int32

const (
	SlbNewAcclCfgOptExcRuleTableResType_Html        SlbNewAcclCfgOptExcRuleTableResType = 1
	SlbNewAcclCfgOptExcRuleTableResType_Js          SlbNewAcclCfgOptExcRuleTableResType = 2
	SlbNewAcclCfgOptExcRuleTableResType_Css         SlbNewAcclCfgOptExcRuleTableResType = 3
	SlbNewAcclCfgOptExcRuleTableResType_Image       SlbNewAcclCfgOptExcRuleTableResType = 4
	SlbNewAcclCfgOptExcRuleTableResType_Other       SlbNewAcclCfgOptExcRuleTableResType = 5
	SlbNewAcclCfgOptExcRuleTableResType_Unsupported SlbNewAcclCfgOptExcRuleTableResType = 2147483647
)

type SlbNewAcclCfgOptExcRuleTableCombineCSS int32

const (
	SlbNewAcclCfgOptExcRuleTableCombineCSS_Enabled     SlbNewAcclCfgOptExcRuleTableCombineCSS = 1
	SlbNewAcclCfgOptExcRuleTableCombineCSS_Disabled    SlbNewAcclCfgOptExcRuleTableCombineCSS = 2
	SlbNewAcclCfgOptExcRuleTableCombineCSS_Unsupported SlbNewAcclCfgOptExcRuleTableCombineCSS = 2147483647
)

type SlbNewAcclCfgOptExcRuleTableCombineJS int32

const (
	SlbNewAcclCfgOptExcRuleTableCombineJS_Enabled     SlbNewAcclCfgOptExcRuleTableCombineJS = 1
	SlbNewAcclCfgOptExcRuleTableCombineJS_Disabled    SlbNewAcclCfgOptExcRuleTableCombineJS = 2
	SlbNewAcclCfgOptExcRuleTableCombineJS_Unsupported SlbNewAcclCfgOptExcRuleTableCombineJS = 2147483647
)

type SlbNewAcclCfgOptExcRuleTableDynamicCache int32

const (
	SlbNewAcclCfgOptExcRuleTableDynamicCache_Enabled     SlbNewAcclCfgOptExcRuleTableDynamicCache = 1
	SlbNewAcclCfgOptExcRuleTableDynamicCache_Disabled    SlbNewAcclCfgOptExcRuleTableDynamicCache = 2
	SlbNewAcclCfgOptExcRuleTableDynamicCache_Unsupported SlbNewAcclCfgOptExcRuleTableDynamicCache = 2147483647
)

type SlbNewAcclCfgOptExcRuleTableInlineCSS int32

const (
	SlbNewAcclCfgOptExcRuleTableInlineCSS_Enabled     SlbNewAcclCfgOptExcRuleTableInlineCSS = 1
	SlbNewAcclCfgOptExcRuleTableInlineCSS_Disabled    SlbNewAcclCfgOptExcRuleTableInlineCSS = 2
	SlbNewAcclCfgOptExcRuleTableInlineCSS_Unsupported SlbNewAcclCfgOptExcRuleTableInlineCSS = 2147483647
)

type SlbNewAcclCfgOptExcRuleTableInlineJS int32

const (
	SlbNewAcclCfgOptExcRuleTableInlineJS_Enabled     SlbNewAcclCfgOptExcRuleTableInlineJS = 1
	SlbNewAcclCfgOptExcRuleTableInlineJS_Disabled    SlbNewAcclCfgOptExcRuleTableInlineJS = 2
	SlbNewAcclCfgOptExcRuleTableInlineJS_Unsupported SlbNewAcclCfgOptExcRuleTableInlineJS = 2147483647
)

type SlbNewAcclCfgOptExcRuleTableImageDim int32

const (
	SlbNewAcclCfgOptExcRuleTableImageDim_Enabled     SlbNewAcclCfgOptExcRuleTableImageDim = 1
	SlbNewAcclCfgOptExcRuleTableImageDim_Disabled    SlbNewAcclCfgOptExcRuleTableImageDim = 2
	SlbNewAcclCfgOptExcRuleTableImageDim_Unsupported SlbNewAcclCfgOptExcRuleTableImageDim = 2147483647
)

type SlbNewAcclCfgOptExcRuleTableRemoveComments int32

const (
	SlbNewAcclCfgOptExcRuleTableRemoveComments_Enabled     SlbNewAcclCfgOptExcRuleTableRemoveComments = 1
	SlbNewAcclCfgOptExcRuleTableRemoveComments_Disabled    SlbNewAcclCfgOptExcRuleTableRemoveComments = 2
	SlbNewAcclCfgOptExcRuleTableRemoveComments_Unsupported SlbNewAcclCfgOptExcRuleTableRemoveComments = 2147483647
)

type SlbNewAcclCfgOptExcRuleTableRemoveWS int32

const (
	SlbNewAcclCfgOptExcRuleTableRemoveWS_Enabled     SlbNewAcclCfgOptExcRuleTableRemoveWS = 1
	SlbNewAcclCfgOptExcRuleTableRemoveWS_Disabled    SlbNewAcclCfgOptExcRuleTableRemoveWS = 2
	SlbNewAcclCfgOptExcRuleTableRemoveWS_Unsupported SlbNewAcclCfgOptExcRuleTableRemoveWS = 2147483647
)

type SlbNewAcclCfgOptExcRuleTableTrimURL int32

const (
	SlbNewAcclCfgOptExcRuleTableTrimURL_Enabled     SlbNewAcclCfgOptExcRuleTableTrimURL = 1
	SlbNewAcclCfgOptExcRuleTableTrimURL_Disabled    SlbNewAcclCfgOptExcRuleTableTrimURL = 2
	SlbNewAcclCfgOptExcRuleTableTrimURL_Unsupported SlbNewAcclCfgOptExcRuleTableTrimURL = 2147483647
)

type SlbNewAcclCfgOptExcRuleTableAdminStatus int32

const (
	SlbNewAcclCfgOptExcRuleTableAdminStatus_Enabled     SlbNewAcclCfgOptExcRuleTableAdminStatus = 1
	SlbNewAcclCfgOptExcRuleTableAdminStatus_Disabled    SlbNewAcclCfgOptExcRuleTableAdminStatus = 2
	SlbNewAcclCfgOptExcRuleTableAdminStatus_Unsupported SlbNewAcclCfgOptExcRuleTableAdminStatus = 2147483647
)

type SlbNewAcclCfgOptExcRuleTableDelete int32

const (
	SlbNewAcclCfgOptExcRuleTableDelete_Other       SlbNewAcclCfgOptExcRuleTableDelete = 1
	SlbNewAcclCfgOptExcRuleTableDelete_Delete      SlbNewAcclCfgOptExcRuleTableDelete = 2
	SlbNewAcclCfgOptExcRuleTableDelete_Unsupported SlbNewAcclCfgOptExcRuleTableDelete = 2147483647
)

type SlbNewAcclCfgOptExcRuleTableParams struct {
	// The optimization rule list (key id) as an index.
	ListIdIndex string `json:"ListIdIndex,omitempty"`
	// The optimization Rule number as an index.
	Index int32 `json:"Index,omitempty"`
	// The optimization Rule name.
	Name string `json:"Name,omitempty"`
	// Resource type for the rule.
	ResType SlbNewAcclCfgOptExcRuleTableResType `json:"ResType,omitempty"`
	// URI of the resource.
	URI string `json:"URI,omitempty"`
	// Combine CSS optimization.
	CombineCSS SlbNewAcclCfgOptExcRuleTableCombineCSS `json:"CombineCSS,omitempty"`
	// Combine JavaScript optimization.
	CombineJS SlbNewAcclCfgOptExcRuleTableCombineJS `json:"CombineJS,omitempty"`
	// Dynamic cache optimization.
	DynamicCache SlbNewAcclCfgOptExcRuleTableDynamicCache `json:"DynamicCache,omitempty"`
	// Inline CSS optimization.
	InlineCSS SlbNewAcclCfgOptExcRuleTableInlineCSS `json:"InlineCSS,omitempty"`
	// Inline JavaScript optimization.
	InlineJS SlbNewAcclCfgOptExcRuleTableInlineJS `json:"InlineJS,omitempty"`
	// Image dimensions setting optimization.
	ImageDim SlbNewAcclCfgOptExcRuleTableImageDim `json:"ImageDim,omitempty"`
	// Remove comments optimization.
	RemoveComments SlbNewAcclCfgOptExcRuleTableRemoveComments `json:"RemoveComments,omitempty"`
	// Remove whitespace optimization.
	RemoveWS SlbNewAcclCfgOptExcRuleTableRemoveWS `json:"RemoveWS,omitempty"`
	// Trim URLs optimization.
	TrimURL SlbNewAcclCfgOptExcRuleTableTrimURL `json:"TrimURL,omitempty"`
	// Status (enable/disable) of optimization Rule.
	AdminStatus SlbNewAcclCfgOptExcRuleTableAdminStatus `json:"AdminStatus,omitempty"`
	// By setting the value to delete(2), the entire row is deleted.
	Delete SlbNewAcclCfgOptExcRuleTableDelete `json:"Delete,omitempty"`
	// Copy rule to another index in the same rule-list. When read,
	// zero is returned.
	Copy DisplayString `json:"Copy,omitempty"`
}

func (p SlbNewAcclCfgOptExcRuleTableParams) iMABean() {}
