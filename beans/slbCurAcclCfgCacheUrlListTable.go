package beans

import (
	"fmt"
	"reflect"
)

// SlbCurAcclCfgCacheUrlListTable The table for configuring caching URL LIST.
// Note:This mib is not supported for VX instance of Virtualization.
type SlbCurAcclCfgCacheUrlListTable struct {
	// The caching policy URL LIST (key id) as an index.
	SlbCurAcclCfgCacheUrlListIdIndex string
	Params                           *SlbCurAcclCfgCacheUrlListTableParams
}

func NewSlbCurAcclCfgCacheUrlListTableList() *SlbCurAcclCfgCacheUrlListTable {
	return &SlbCurAcclCfgCacheUrlListTable{}
}

func NewSlbCurAcclCfgCacheUrlListTable(
	slbCurAcclCfgCacheUrlListIdIndex string,
	params *SlbCurAcclCfgCacheUrlListTableParams,
) *SlbCurAcclCfgCacheUrlListTable {
	return &SlbCurAcclCfgCacheUrlListTable{
		SlbCurAcclCfgCacheUrlListIdIndex: slbCurAcclCfgCacheUrlListIdIndex,
		Params:                           params,
	}
}

func (c *SlbCurAcclCfgCacheUrlListTable) Name() string {
	return "SlbCurAcclCfgCacheUrlListTable"
}

func (c *SlbCurAcclCfgCacheUrlListTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbCurAcclCfgCacheUrlListTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbCurAcclCfgCacheUrlListTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbCurAcclCfgCacheUrlListIdIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbCurAcclCfgCacheUrlListIdIndex)
}

type SlbCurAcclCfgCacheUrlListTableAdminStatus int32

const (
	SlbCurAcclCfgCacheUrlListTableAdminStatus_Enabled     SlbCurAcclCfgCacheUrlListTableAdminStatus = 1
	SlbCurAcclCfgCacheUrlListTableAdminStatus_Disabled    SlbCurAcclCfgCacheUrlListTableAdminStatus = 2
	SlbCurAcclCfgCacheUrlListTableAdminStatus_Unsupported SlbCurAcclCfgCacheUrlListTableAdminStatus = 2147483647
)

type SlbCurAcclCfgCacheUrlListTableParams struct {
	// The caching policy URL LIST (key id) as an index.
	IdIndex string `json:"IdIndex,omitempty"`
	// Cache URL LIST name.
	Name string `json:"Name,omitempty"`
	// Status (enable/disable) of caching URL LIST.
	AdminStatus SlbCurAcclCfgCacheUrlListTableAdminStatus `json:"AdminStatus,omitempty"`
}

func (p SlbCurAcclCfgCacheUrlListTableParams) iMABean() {}
