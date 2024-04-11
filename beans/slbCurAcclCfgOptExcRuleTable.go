package beans

import (
	"fmt"
	"reflect"
)

// SlbCurAcclCfgOptExcRuleTable The table for configuring optimization Rules.
// Note:This mib is not supported for VX instance of Virtualization.
type SlbCurAcclCfgOptExcRuleTable struct {
	// The optimization rule list (key id) as an index.
	SlbCurAcclCfgOptExcRuleListIdIndex string
	// The optimization Rule number as an index.
	SlbCurAcclCfgOptExcRuleIndex int32
	Params                       *SlbCurAcclCfgOptExcRuleTableParams
}

func NewSlbCurAcclCfgOptExcRuleTableList() *SlbCurAcclCfgOptExcRuleTable {
	return &SlbCurAcclCfgOptExcRuleTable{}
}

func NewSlbCurAcclCfgOptExcRuleTable(
	slbCurAcclCfgOptExcRuleListIdIndex string,
	slbCurAcclCfgOptExcRuleIndex int32,
	params *SlbCurAcclCfgOptExcRuleTableParams,
) *SlbCurAcclCfgOptExcRuleTable {
	return &SlbCurAcclCfgOptExcRuleTable{
		SlbCurAcclCfgOptExcRuleListIdIndex: slbCurAcclCfgOptExcRuleListIdIndex,
		SlbCurAcclCfgOptExcRuleIndex:       slbCurAcclCfgOptExcRuleIndex,
		Params:                             params,
	}
}

func (c *SlbCurAcclCfgOptExcRuleTable) Name() string {
	return "SlbCurAcclCfgOptExcRuleTable"
}

func (c *SlbCurAcclCfgOptExcRuleTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbCurAcclCfgOptExcRuleTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbCurAcclCfgOptExcRuleTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbCurAcclCfgOptExcRuleListIdIndex).IsZero() &&
		reflect.ValueOf(c.SlbCurAcclCfgOptExcRuleIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbCurAcclCfgOptExcRuleListIdIndex) + "/" + fmt.Sprint(c.SlbCurAcclCfgOptExcRuleIndex)
}

type SlbCurAcclCfgOptExcRuleTableResType int32

const (
	SlbCurAcclCfgOptExcRuleTableResType_Html        SlbCurAcclCfgOptExcRuleTableResType = 1
	SlbCurAcclCfgOptExcRuleTableResType_Js          SlbCurAcclCfgOptExcRuleTableResType = 2
	SlbCurAcclCfgOptExcRuleTableResType_Css         SlbCurAcclCfgOptExcRuleTableResType = 3
	SlbCurAcclCfgOptExcRuleTableResType_Image       SlbCurAcclCfgOptExcRuleTableResType = 4
	SlbCurAcclCfgOptExcRuleTableResType_Other       SlbCurAcclCfgOptExcRuleTableResType = 5
	SlbCurAcclCfgOptExcRuleTableResType_Unsupported SlbCurAcclCfgOptExcRuleTableResType = 2147483647
)

type SlbCurAcclCfgOptExcRuleTableCombineCSS int32

const (
	SlbCurAcclCfgOptExcRuleTableCombineCSS_Enabled     SlbCurAcclCfgOptExcRuleTableCombineCSS = 1
	SlbCurAcclCfgOptExcRuleTableCombineCSS_Disabled    SlbCurAcclCfgOptExcRuleTableCombineCSS = 2
	SlbCurAcclCfgOptExcRuleTableCombineCSS_Unsupported SlbCurAcclCfgOptExcRuleTableCombineCSS = 2147483647
)

type SlbCurAcclCfgOptExcRuleTableCombineJS int32

const (
	SlbCurAcclCfgOptExcRuleTableCombineJS_Enabled     SlbCurAcclCfgOptExcRuleTableCombineJS = 1
	SlbCurAcclCfgOptExcRuleTableCombineJS_Disabled    SlbCurAcclCfgOptExcRuleTableCombineJS = 2
	SlbCurAcclCfgOptExcRuleTableCombineJS_Unsupported SlbCurAcclCfgOptExcRuleTableCombineJS = 2147483647
)

type SlbCurAcclCfgOptExcRuleTableDynamicCache int32

const (
	SlbCurAcclCfgOptExcRuleTableDynamicCache_Enabled     SlbCurAcclCfgOptExcRuleTableDynamicCache = 1
	SlbCurAcclCfgOptExcRuleTableDynamicCache_Disabled    SlbCurAcclCfgOptExcRuleTableDynamicCache = 2
	SlbCurAcclCfgOptExcRuleTableDynamicCache_Unsupported SlbCurAcclCfgOptExcRuleTableDynamicCache = 2147483647
)

type SlbCurAcclCfgOptExcRuleTableInlineCSS int32

const (
	SlbCurAcclCfgOptExcRuleTableInlineCSS_Enabled     SlbCurAcclCfgOptExcRuleTableInlineCSS = 1
	SlbCurAcclCfgOptExcRuleTableInlineCSS_Disabled    SlbCurAcclCfgOptExcRuleTableInlineCSS = 2
	SlbCurAcclCfgOptExcRuleTableInlineCSS_Unsupported SlbCurAcclCfgOptExcRuleTableInlineCSS = 2147483647
)

type SlbCurAcclCfgOptExcRuleTableInlineJS int32

const (
	SlbCurAcclCfgOptExcRuleTableInlineJS_Enabled     SlbCurAcclCfgOptExcRuleTableInlineJS = 1
	SlbCurAcclCfgOptExcRuleTableInlineJS_Disabled    SlbCurAcclCfgOptExcRuleTableInlineJS = 2
	SlbCurAcclCfgOptExcRuleTableInlineJS_Unsupported SlbCurAcclCfgOptExcRuleTableInlineJS = 2147483647
)

type SlbCurAcclCfgOptExcRuleTableImageDim int32

const (
	SlbCurAcclCfgOptExcRuleTableImageDim_Enabled     SlbCurAcclCfgOptExcRuleTableImageDim = 1
	SlbCurAcclCfgOptExcRuleTableImageDim_Disabled    SlbCurAcclCfgOptExcRuleTableImageDim = 2
	SlbCurAcclCfgOptExcRuleTableImageDim_Unsupported SlbCurAcclCfgOptExcRuleTableImageDim = 2147483647
)

type SlbCurAcclCfgOptExcRuleTableRemoveComments int32

const (
	SlbCurAcclCfgOptExcRuleTableRemoveComments_Enabled     SlbCurAcclCfgOptExcRuleTableRemoveComments = 1
	SlbCurAcclCfgOptExcRuleTableRemoveComments_Disabled    SlbCurAcclCfgOptExcRuleTableRemoveComments = 2
	SlbCurAcclCfgOptExcRuleTableRemoveComments_Unsupported SlbCurAcclCfgOptExcRuleTableRemoveComments = 2147483647
)

type SlbCurAcclCfgOptExcRuleTableRemoveWS int32

const (
	SlbCurAcclCfgOptExcRuleTableRemoveWS_Enabled     SlbCurAcclCfgOptExcRuleTableRemoveWS = 1
	SlbCurAcclCfgOptExcRuleTableRemoveWS_Disabled    SlbCurAcclCfgOptExcRuleTableRemoveWS = 2
	SlbCurAcclCfgOptExcRuleTableRemoveWS_Unsupported SlbCurAcclCfgOptExcRuleTableRemoveWS = 2147483647
)

type SlbCurAcclCfgOptExcRuleTableTrimURL int32

const (
	SlbCurAcclCfgOptExcRuleTableTrimURL_Enabled     SlbCurAcclCfgOptExcRuleTableTrimURL = 1
	SlbCurAcclCfgOptExcRuleTableTrimURL_Disabled    SlbCurAcclCfgOptExcRuleTableTrimURL = 2
	SlbCurAcclCfgOptExcRuleTableTrimURL_Unsupported SlbCurAcclCfgOptExcRuleTableTrimURL = 2147483647
)

type SlbCurAcclCfgOptExcRuleTableAdminStatus int32

const (
	SlbCurAcclCfgOptExcRuleTableAdminStatus_Enabled     SlbCurAcclCfgOptExcRuleTableAdminStatus = 1
	SlbCurAcclCfgOptExcRuleTableAdminStatus_Disabled    SlbCurAcclCfgOptExcRuleTableAdminStatus = 2
	SlbCurAcclCfgOptExcRuleTableAdminStatus_Unsupported SlbCurAcclCfgOptExcRuleTableAdminStatus = 2147483647
)

type SlbCurAcclCfgOptExcRuleTableParams struct {
	// The optimization rule list (key id) as an index.
	ListIdIndex string `json:"ListIdIndex,omitempty"`
	// The optimization Rule number as an index.
	Index int32 `json:"Index,omitempty"`
	// The optimization Rule name.
	Name string `json:"Name,omitempty"`
	// Resource type.
	ResType SlbCurAcclCfgOptExcRuleTableResType `json:"ResType,omitempty"`
	// URI of the resource for which this rule applies.
	URI string `json:"URI,omitempty"`
	// Combine CSS optimization.
	CombineCSS SlbCurAcclCfgOptExcRuleTableCombineCSS `json:"CombineCSS,omitempty"`
	// Combine JavaScript optimization.
	CombineJS SlbCurAcclCfgOptExcRuleTableCombineJS `json:"CombineJS,omitempty"`
	// Dynamic cache optimization.
	DynamicCache SlbCurAcclCfgOptExcRuleTableDynamicCache `json:"DynamicCache,omitempty"`
	// Inline CSS optimization.
	InlineCSS SlbCurAcclCfgOptExcRuleTableInlineCSS `json:"InlineCSS,omitempty"`
	// Inline JavaScript optimization.
	InlineJS SlbCurAcclCfgOptExcRuleTableInlineJS `json:"InlineJS,omitempty"`
	// Image dimensions setting optimization.
	ImageDim SlbCurAcclCfgOptExcRuleTableImageDim `json:"ImageDim,omitempty"`
	// Remove comments optimization.
	RemoveComments SlbCurAcclCfgOptExcRuleTableRemoveComments `json:"RemoveComments,omitempty"`
	// Remove whitespace optimization.
	RemoveWS SlbCurAcclCfgOptExcRuleTableRemoveWS `json:"RemoveWS,omitempty"`
	// Trim URLs optimization.
	TrimURL SlbCurAcclCfgOptExcRuleTableTrimURL `json:"TrimURL,omitempty"`
	// Status (enable/disable) of optimization Rule.
	AdminStatus SlbCurAcclCfgOptExcRuleTableAdminStatus `json:"AdminStatus,omitempty"`
}

func (p SlbCurAcclCfgOptExcRuleTableParams) iMABean() {}
