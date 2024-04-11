package beans

import (
	"fmt"
	"reflect"
)

// SlbCurAcclCfgCachePolTable The table for configuring caching policy.
// Note:This mib is not supported for VX instance of Virtualization.
type SlbCurAcclCfgCachePolTable struct {
	// The caching policy ID(key id) as an index.
	SlbCurAcclCfgCachePolNameIdIndex string
	Params                           *SlbCurAcclCfgCachePolTableParams
}

func NewSlbCurAcclCfgCachePolTableList() *SlbCurAcclCfgCachePolTable {
	return &SlbCurAcclCfgCachePolTable{}
}

func NewSlbCurAcclCfgCachePolTable(
	slbCurAcclCfgCachePolNameIdIndex string,
	params *SlbCurAcclCfgCachePolTableParams,
) *SlbCurAcclCfgCachePolTable {
	return &SlbCurAcclCfgCachePolTable{
		SlbCurAcclCfgCachePolNameIdIndex: slbCurAcclCfgCachePolNameIdIndex,
		Params:                           params,
	}
}

func (c *SlbCurAcclCfgCachePolTable) Name() string {
	return "SlbCurAcclCfgCachePolTable"
}

func (c *SlbCurAcclCfgCachePolTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbCurAcclCfgCachePolTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbCurAcclCfgCachePolTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbCurAcclCfgCachePolNameIdIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbCurAcclCfgCachePolNameIdIndex)
}

type SlbCurAcclCfgCachePolTableStore int32

const (
	SlbCurAcclCfgCachePolTableStore_Srvrhdr     SlbCurAcclCfgCachePolTableStore = 1
	SlbCurAcclCfgCachePolTableStore_Cacheall    SlbCurAcclCfgCachePolTableStore = 2
	SlbCurAcclCfgCachePolTableStore_Unsupported SlbCurAcclCfgCachePolTableStore = 2147483647
)

type SlbCurAcclCfgCachePolTableServe int32

const (
	SlbCurAcclCfgCachePolTableServe_Clnthdr     SlbCurAcclCfgCachePolTableServe = 1
	SlbCurAcclCfgCachePolTableServe_Refresh     SlbCurAcclCfgCachePolTableServe = 2
	SlbCurAcclCfgCachePolTableServe_Cache       SlbCurAcclCfgCachePolTableServe = 3
	SlbCurAcclCfgCachePolTableServe_Unsupported SlbCurAcclCfgCachePolTableServe = 2147483647
)

type SlbCurAcclCfgCachePolTableQuery int32

const (
	SlbCurAcclCfgCachePolTableQuery_Consider    SlbCurAcclCfgCachePolTableQuery = 1
	SlbCurAcclCfgCachePolTableQuery_Ignore      SlbCurAcclCfgCachePolTableQuery = 2
	SlbCurAcclCfgCachePolTableQuery_Unsupported SlbCurAcclCfgCachePolTableQuery = 2147483647
)

type SlbCurAcclCfgCachePolTableBrowser int32

const (
	SlbCurAcclCfgCachePolTableBrowser_Enabled     SlbCurAcclCfgCachePolTableBrowser = 1
	SlbCurAcclCfgCachePolTableBrowser_Disabled    SlbCurAcclCfgCachePolTableBrowser = 2
	SlbCurAcclCfgCachePolTableBrowser_Unsupported SlbCurAcclCfgCachePolTableBrowser = 2147483647
)

type SlbCurAcclCfgCachePolTableAdminStatus int32

const (
	SlbCurAcclCfgCachePolTableAdminStatus_Enabled     SlbCurAcclCfgCachePolTableAdminStatus = 1
	SlbCurAcclCfgCachePolTableAdminStatus_Disabled    SlbCurAcclCfgCachePolTableAdminStatus = 2
	SlbCurAcclCfgCachePolTableAdminStatus_Unsupported SlbCurAcclCfgCachePolTableAdminStatus = 2147483647
)

type SlbCurAcclCfgCachePolTableDummy int32

const (
	SlbCurAcclCfgCachePolTableDummy_Enabled     SlbCurAcclCfgCachePolTableDummy = 1
	SlbCurAcclCfgCachePolTableDummy_Disabled    SlbCurAcclCfgCachePolTableDummy = 2
	SlbCurAcclCfgCachePolTableDummy_Unsupported SlbCurAcclCfgCachePolTableDummy = 2147483647
)

type SlbCurAcclCfgCachePolTableCombineCSS int32

const (
	SlbCurAcclCfgCachePolTableCombineCSS_Enabled     SlbCurAcclCfgCachePolTableCombineCSS = 1
	SlbCurAcclCfgCachePolTableCombineCSS_Disabled    SlbCurAcclCfgCachePolTableCombineCSS = 2
	SlbCurAcclCfgCachePolTableCombineCSS_Unsupported SlbCurAcclCfgCachePolTableCombineCSS = 2147483647
)

type SlbCurAcclCfgCachePolTableCombineJS int32

const (
	SlbCurAcclCfgCachePolTableCombineJS_Enabled     SlbCurAcclCfgCachePolTableCombineJS = 1
	SlbCurAcclCfgCachePolTableCombineJS_Disabled    SlbCurAcclCfgCachePolTableCombineJS = 2
	SlbCurAcclCfgCachePolTableCombineJS_Unsupported SlbCurAcclCfgCachePolTableCombineJS = 2147483647
)

type SlbCurAcclCfgCachePolTableDynamicCache int32

const (
	SlbCurAcclCfgCachePolTableDynamicCache_Enabled     SlbCurAcclCfgCachePolTableDynamicCache = 1
	SlbCurAcclCfgCachePolTableDynamicCache_Disabled    SlbCurAcclCfgCachePolTableDynamicCache = 2
	SlbCurAcclCfgCachePolTableDynamicCache_Unsupported SlbCurAcclCfgCachePolTableDynamicCache = 2147483647
)

type SlbCurAcclCfgCachePolTableInlineCSS int32

const (
	SlbCurAcclCfgCachePolTableInlineCSS_Enabled     SlbCurAcclCfgCachePolTableInlineCSS = 1
	SlbCurAcclCfgCachePolTableInlineCSS_Disabled    SlbCurAcclCfgCachePolTableInlineCSS = 2
	SlbCurAcclCfgCachePolTableInlineCSS_Unsupported SlbCurAcclCfgCachePolTableInlineCSS = 2147483647
)

type SlbCurAcclCfgCachePolTableInlineJS int32

const (
	SlbCurAcclCfgCachePolTableInlineJS_Enabled     SlbCurAcclCfgCachePolTableInlineJS = 1
	SlbCurAcclCfgCachePolTableInlineJS_Disabled    SlbCurAcclCfgCachePolTableInlineJS = 2
	SlbCurAcclCfgCachePolTableInlineJS_Unsupported SlbCurAcclCfgCachePolTableInlineJS = 2147483647
)

type SlbCurAcclCfgCachePolTableImageDim int32

const (
	SlbCurAcclCfgCachePolTableImageDim_Enabled     SlbCurAcclCfgCachePolTableImageDim = 1
	SlbCurAcclCfgCachePolTableImageDim_Disabled    SlbCurAcclCfgCachePolTableImageDim = 2
	SlbCurAcclCfgCachePolTableImageDim_Unsupported SlbCurAcclCfgCachePolTableImageDim = 2147483647
)

type SlbCurAcclCfgCachePolTableRemoveCmnt int32

const (
	SlbCurAcclCfgCachePolTableRemoveCmnt_Enabled     SlbCurAcclCfgCachePolTableRemoveCmnt = 1
	SlbCurAcclCfgCachePolTableRemoveCmnt_Disabled    SlbCurAcclCfgCachePolTableRemoveCmnt = 2
	SlbCurAcclCfgCachePolTableRemoveCmnt_Unsupported SlbCurAcclCfgCachePolTableRemoveCmnt = 2147483647
)

type SlbCurAcclCfgCachePolTableRemoveWS int32

const (
	SlbCurAcclCfgCachePolTableRemoveWS_Enabled     SlbCurAcclCfgCachePolTableRemoveWS = 1
	SlbCurAcclCfgCachePolTableRemoveWS_Disabled    SlbCurAcclCfgCachePolTableRemoveWS = 2
	SlbCurAcclCfgCachePolTableRemoveWS_Unsupported SlbCurAcclCfgCachePolTableRemoveWS = 2147483647
)

type SlbCurAcclCfgCachePolTableTrimURL int32

const (
	SlbCurAcclCfgCachePolTableTrimURL_Enabled     SlbCurAcclCfgCachePolTableTrimURL = 1
	SlbCurAcclCfgCachePolTableTrimURL_Disabled    SlbCurAcclCfgCachePolTableTrimURL = 2
	SlbCurAcclCfgCachePolTableTrimURL_Unsupported SlbCurAcclCfgCachePolTableTrimURL = 2147483647
)

type SlbCurAcclCfgCachePolTableParams struct {
	// The caching policy ID(key id) as an index.
	NameIdIndex string `json:"NameIdIndex,omitempty"`
	// Cache policy name.
	Name string `json:"Name,omitempty"`
	// Maximum expiration time in seconds.
	ExpireTime int32 `json:"ExpireTime,omitempty"`
	// Minimum Object size to be stored in bytes.
	MinSize int32 `json:"MinSize,omitempty"`
	// Maximum Object size to be stored in bytes.
	MaxSize int32 `json:"MaxSize,omitempty"`
	// URL list which is associated with this policy.
	URLList string `json:"URLList,omitempty"`
	// Caching behavior for storing new objects in cache.
	Store SlbCurAcclCfgCachePolTableStore `json:"Store,omitempty"`
	// Cache behavior when serving clients with objects.
	Serve SlbCurAcclCfgCachePolTableServe `json:"Serve,omitempty"`
	// Specifies whether query parameters added to object URL should be considered when storing/serving object from cache.
	Query SlbCurAcclCfgCachePolTableQuery `json:"Query,omitempty"`
	// Optimize browser cache.
	Browser SlbCurAcclCfgCachePolTableBrowser `json:"Browser,omitempty"`
	// Status (enable/disable) of caching policy.
	AdminStatus SlbCurAcclCfgCachePolTableAdminStatus `json:"AdminStatus,omitempty"`
	// Dummy field to maintain offsets between cur and new.
	Dummy SlbCurAcclCfgCachePolTableDummy `json:"Dummy,omitempty"`
	// Combine CSS files.
	CombineCSS SlbCurAcclCfgCachePolTableCombineCSS `json:"CombineCSS,omitempty"`
	// Combine JavaScript files.
	CombineJS SlbCurAcclCfgCachePolTableCombineJS `json:"CombineJS,omitempty"`
	// Enhance broswer cache.
	DynamicCache SlbCurAcclCfgCachePolTableDynamicCache `json:"DynamicCache,omitempty"`
	// Inline small CSS files.
	InlineCSS SlbCurAcclCfgCachePolTableInlineCSS `json:"InlineCSS,omitempty"`
	// Inline small JavaScript files.
	InlineJS SlbCurAcclCfgCachePolTableInlineJS `json:"InlineJS,omitempty"`
	// Update image dimensions in html.
	ImageDim SlbCurAcclCfgCachePolTableImageDim `json:"ImageDim,omitempty"`
	// Remove comments from html.
	RemoveCmnt SlbCurAcclCfgCachePolTableRemoveCmnt `json:"RemoveCmnt,omitempty"`
	// Remove unnecessary whitespace from html code.
	RemoveWS SlbCurAcclCfgCachePolTableRemoveWS `json:"RemoveWS,omitempty"`
	// Trim absolute URLs in html.
	TrimURL SlbCurAcclCfgCachePolTableTrimURL `json:"TrimURL,omitempty"`
	// Optimization exception list which is associated with this policy.
	OptList string `json:"OptList,omitempty"`
}

func (p SlbCurAcclCfgCachePolTableParams) iMABean() {}
