package beans

import (
	"fmt"
	"reflect"
)

// SlbNewAcclCfgCachePolTable The table for configuring caching policy.
// Note:This mib is not supported for VX instance of Virtualization.
type SlbNewAcclCfgCachePolTable struct {
	// The caching policy ID(key id) as an index, length of the string should be 32 characters.
	SlbNewAcclCfgCachePolNameIdIndex string
	Params                           *SlbNewAcclCfgCachePolTableParams
}

func NewSlbNewAcclCfgCachePolTableList() *SlbNewAcclCfgCachePolTable {
	return &SlbNewAcclCfgCachePolTable{}
}

func NewSlbNewAcclCfgCachePolTable(
	slbNewAcclCfgCachePolNameIdIndex string,
	params *SlbNewAcclCfgCachePolTableParams,
) *SlbNewAcclCfgCachePolTable {
	return &SlbNewAcclCfgCachePolTable{
		SlbNewAcclCfgCachePolNameIdIndex: slbNewAcclCfgCachePolNameIdIndex,
		Params:                           params,
	}
}

func (c *SlbNewAcclCfgCachePolTable) Name() string {
	return "SlbNewAcclCfgCachePolTable"
}

func (c *SlbNewAcclCfgCachePolTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbNewAcclCfgCachePolTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbNewAcclCfgCachePolTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbNewAcclCfgCachePolNameIdIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbNewAcclCfgCachePolNameIdIndex)
}

type SlbNewAcclCfgCachePolTableStore int32

const (
	SlbNewAcclCfgCachePolTableStore_Srvrhdr     SlbNewAcclCfgCachePolTableStore = 1
	SlbNewAcclCfgCachePolTableStore_Cacheall    SlbNewAcclCfgCachePolTableStore = 2
	SlbNewAcclCfgCachePolTableStore_Unsupported SlbNewAcclCfgCachePolTableStore = 2147483647
)

type SlbNewAcclCfgCachePolTableServe int32

const (
	SlbNewAcclCfgCachePolTableServe_Clnthdr     SlbNewAcclCfgCachePolTableServe = 1
	SlbNewAcclCfgCachePolTableServe_Refresh     SlbNewAcclCfgCachePolTableServe = 2
	SlbNewAcclCfgCachePolTableServe_Cache       SlbNewAcclCfgCachePolTableServe = 3
	SlbNewAcclCfgCachePolTableServe_Unsupported SlbNewAcclCfgCachePolTableServe = 2147483647
)

type SlbNewAcclCfgCachePolTableQuery int32

const (
	SlbNewAcclCfgCachePolTableQuery_Consider    SlbNewAcclCfgCachePolTableQuery = 1
	SlbNewAcclCfgCachePolTableQuery_Ignore      SlbNewAcclCfgCachePolTableQuery = 2
	SlbNewAcclCfgCachePolTableQuery_Unsupported SlbNewAcclCfgCachePolTableQuery = 2147483647
)

type SlbNewAcclCfgCachePolTableBrowser int32

const (
	SlbNewAcclCfgCachePolTableBrowser_Enabled     SlbNewAcclCfgCachePolTableBrowser = 1
	SlbNewAcclCfgCachePolTableBrowser_Disabled    SlbNewAcclCfgCachePolTableBrowser = 2
	SlbNewAcclCfgCachePolTableBrowser_Unsupported SlbNewAcclCfgCachePolTableBrowser = 2147483647
)

type SlbNewAcclCfgCachePolTableAdminStatus int32

const (
	SlbNewAcclCfgCachePolTableAdminStatus_Enabled     SlbNewAcclCfgCachePolTableAdminStatus = 1
	SlbNewAcclCfgCachePolTableAdminStatus_Disabled    SlbNewAcclCfgCachePolTableAdminStatus = 2
	SlbNewAcclCfgCachePolTableAdminStatus_Unsupported SlbNewAcclCfgCachePolTableAdminStatus = 2147483647
)

type SlbNewAcclCfgCachePolTableDelete int32

const (
	SlbNewAcclCfgCachePolTableDelete_Other       SlbNewAcclCfgCachePolTableDelete = 1
	SlbNewAcclCfgCachePolTableDelete_Delete      SlbNewAcclCfgCachePolTableDelete = 2
	SlbNewAcclCfgCachePolTableDelete_Unsupported SlbNewAcclCfgCachePolTableDelete = 2147483647
)

type SlbNewAcclCfgCachePolTableCombineCSS int32

const (
	SlbNewAcclCfgCachePolTableCombineCSS_Enabled     SlbNewAcclCfgCachePolTableCombineCSS = 1
	SlbNewAcclCfgCachePolTableCombineCSS_Disabled    SlbNewAcclCfgCachePolTableCombineCSS = 2
	SlbNewAcclCfgCachePolTableCombineCSS_Unsupported SlbNewAcclCfgCachePolTableCombineCSS = 2147483647
)

type SlbNewAcclCfgCachePolTableCombineJS int32

const (
	SlbNewAcclCfgCachePolTableCombineJS_Enabled     SlbNewAcclCfgCachePolTableCombineJS = 1
	SlbNewAcclCfgCachePolTableCombineJS_Disabled    SlbNewAcclCfgCachePolTableCombineJS = 2
	SlbNewAcclCfgCachePolTableCombineJS_Unsupported SlbNewAcclCfgCachePolTableCombineJS = 2147483647
)

type SlbNewAcclCfgCachePolTableDynamicCache int32

const (
	SlbNewAcclCfgCachePolTableDynamicCache_Enabled     SlbNewAcclCfgCachePolTableDynamicCache = 1
	SlbNewAcclCfgCachePolTableDynamicCache_Disabled    SlbNewAcclCfgCachePolTableDynamicCache = 2
	SlbNewAcclCfgCachePolTableDynamicCache_Unsupported SlbNewAcclCfgCachePolTableDynamicCache = 2147483647
)

type SlbNewAcclCfgCachePolTableInlineCSS int32

const (
	SlbNewAcclCfgCachePolTableInlineCSS_Enabled     SlbNewAcclCfgCachePolTableInlineCSS = 1
	SlbNewAcclCfgCachePolTableInlineCSS_Disabled    SlbNewAcclCfgCachePolTableInlineCSS = 2
	SlbNewAcclCfgCachePolTableInlineCSS_Unsupported SlbNewAcclCfgCachePolTableInlineCSS = 2147483647
)

type SlbNewAcclCfgCachePolTableInlineJS int32

const (
	SlbNewAcclCfgCachePolTableInlineJS_Enabled     SlbNewAcclCfgCachePolTableInlineJS = 1
	SlbNewAcclCfgCachePolTableInlineJS_Disabled    SlbNewAcclCfgCachePolTableInlineJS = 2
	SlbNewAcclCfgCachePolTableInlineJS_Unsupported SlbNewAcclCfgCachePolTableInlineJS = 2147483647
)

type SlbNewAcclCfgCachePolTableImageDim int32

const (
	SlbNewAcclCfgCachePolTableImageDim_Enabled     SlbNewAcclCfgCachePolTableImageDim = 1
	SlbNewAcclCfgCachePolTableImageDim_Disabled    SlbNewAcclCfgCachePolTableImageDim = 2
	SlbNewAcclCfgCachePolTableImageDim_Unsupported SlbNewAcclCfgCachePolTableImageDim = 2147483647
)

type SlbNewAcclCfgCachePolTableRemoveCmnt int32

const (
	SlbNewAcclCfgCachePolTableRemoveCmnt_Enabled     SlbNewAcclCfgCachePolTableRemoveCmnt = 1
	SlbNewAcclCfgCachePolTableRemoveCmnt_Disabled    SlbNewAcclCfgCachePolTableRemoveCmnt = 2
	SlbNewAcclCfgCachePolTableRemoveCmnt_Unsupported SlbNewAcclCfgCachePolTableRemoveCmnt = 2147483647
)

type SlbNewAcclCfgCachePolTableRemoveWS int32

const (
	SlbNewAcclCfgCachePolTableRemoveWS_Enabled     SlbNewAcclCfgCachePolTableRemoveWS = 1
	SlbNewAcclCfgCachePolTableRemoveWS_Disabled    SlbNewAcclCfgCachePolTableRemoveWS = 2
	SlbNewAcclCfgCachePolTableRemoveWS_Unsupported SlbNewAcclCfgCachePolTableRemoveWS = 2147483647
)

type SlbNewAcclCfgCachePolTableTrimURL int32

const (
	SlbNewAcclCfgCachePolTableTrimURL_Enabled     SlbNewAcclCfgCachePolTableTrimURL = 1
	SlbNewAcclCfgCachePolTableTrimURL_Disabled    SlbNewAcclCfgCachePolTableTrimURL = 2
	SlbNewAcclCfgCachePolTableTrimURL_Unsupported SlbNewAcclCfgCachePolTableTrimURL = 2147483647
)

type SlbNewAcclCfgCachePolTableParams struct {
	// The caching policy ID(key id) as an index, length of the string should be 32 characters.
	NameIdIndex string `json:"NameIdIndex,omitempty"`
	// Cache policy name, length of the string should be 32 characters.
	Name string `json:"Name,omitempty"`
	// Set maximum expiration time, Range is 60- 43,200,000 seconds (~500 days).
	ExpireTime int32 `json:"ExpireTime,omitempty"`
	// Minimum Object size to be stored [Bytes], Values: 1-65536.
	MinSize int32 `json:"MinSize,omitempty"`
	// Maximum Object size to be stored [Bytes], Values: 1-512000000.
	MaxSize int32 `json:"MaxSize,omitempty"`
	// URL list to be associated with this policy, Values: 1-32 character string.
	URLList string `json:"URLList,omitempty"`
	// Define caching behavior for storing New objects in cache.
	Store SlbNewAcclCfgCachePolTableStore `json:"Store,omitempty"`
	// Define cache behavior when serving clients with objects.
	Serve SlbNewAcclCfgCachePolTableServe `json:"Serve,omitempty"`
	// Specifies whether query parameters added to object URL should be considered when storing/serving object from cache.
	Query SlbNewAcclCfgCachePolTableQuery `json:"Query,omitempty"`
	// Optimize browser cache.
	Browser SlbNewAcclCfgCachePolTableBrowser `json:"Browser,omitempty"`
	// Enable or disable caching policy.
	AdminStatus SlbNewAcclCfgCachePolTableAdminStatus `json:"AdminStatus,omitempty"`
	// Delete Caching policy.
	Delete SlbNewAcclCfgCachePolTableDelete `json:"Delete,omitempty"`
	// Combine CSS files.
	CombineCSS SlbNewAcclCfgCachePolTableCombineCSS `json:"CombineCSS,omitempty"`
	// Combine JavaScript files.
	CombineJS SlbNewAcclCfgCachePolTableCombineJS `json:"CombineJS,omitempty"`
	// Enhance broswer cache.
	DynamicCache SlbNewAcclCfgCachePolTableDynamicCache `json:"DynamicCache,omitempty"`
	// Inline small CSS files.
	InlineCSS SlbNewAcclCfgCachePolTableInlineCSS `json:"InlineCSS,omitempty"`
	// Inline small JavaScript files.
	InlineJS SlbNewAcclCfgCachePolTableInlineJS `json:"InlineJS,omitempty"`
	// Update image dimensions in html.
	ImageDim SlbNewAcclCfgCachePolTableImageDim `json:"ImageDim,omitempty"`
	// Remove comments from html.
	RemoveCmnt SlbNewAcclCfgCachePolTableRemoveCmnt `json:"RemoveCmnt,omitempty"`
	// Remove unnecessary whitespace from html code.
	RemoveWS SlbNewAcclCfgCachePolTableRemoveWS `json:"RemoveWS,omitempty"`
	// Trim absolute URLs in html.
	TrimURL SlbNewAcclCfgCachePolTableTrimURL `json:"TrimURL,omitempty"`
	// Optimization exception list to be associated with this policy, Values: 1-32 character string.
	OptList string `json:"OptList,omitempty"`
}

func (p SlbNewAcclCfgCachePolTableParams) iMABean() {}
