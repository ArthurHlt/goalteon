package beans

import (
	"fmt"
	"reflect"
)

// SlbNewAcclCfgCacheUrlListTable The table for configuring caching URL LIST.
// Note:This mib is not supported for VX instance of Virtualization.
type SlbNewAcclCfgCacheUrlListTable struct {
	// The caching policy URL LIST (key id) as an index.
	SlbNewAcclCfgCacheUrlListIdIndex string
	Params                           *SlbNewAcclCfgCacheUrlListTableParams
}

func NewSlbNewAcclCfgCacheUrlListTableList() *SlbNewAcclCfgCacheUrlListTable {
	return &SlbNewAcclCfgCacheUrlListTable{}
}

func NewSlbNewAcclCfgCacheUrlListTable(
	slbNewAcclCfgCacheUrlListIdIndex string,
	params *SlbNewAcclCfgCacheUrlListTableParams,
) *SlbNewAcclCfgCacheUrlListTable {
	return &SlbNewAcclCfgCacheUrlListTable{
		SlbNewAcclCfgCacheUrlListIdIndex: slbNewAcclCfgCacheUrlListIdIndex,
		Params:                           params,
	}
}

func (c *SlbNewAcclCfgCacheUrlListTable) Name() string {
	return "SlbNewAcclCfgCacheUrlListTable"
}

func (c *SlbNewAcclCfgCacheUrlListTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbNewAcclCfgCacheUrlListTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbNewAcclCfgCacheUrlListTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbNewAcclCfgCacheUrlListIdIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbNewAcclCfgCacheUrlListIdIndex)
}

type SlbNewAcclCfgCacheUrlListTableAdminStatus int32

const (
	SlbNewAcclCfgCacheUrlListTableAdminStatus_Enabled     SlbNewAcclCfgCacheUrlListTableAdminStatus = 1
	SlbNewAcclCfgCacheUrlListTableAdminStatus_Disabled    SlbNewAcclCfgCacheUrlListTableAdminStatus = 2
	SlbNewAcclCfgCacheUrlListTableAdminStatus_Unsupported SlbNewAcclCfgCacheUrlListTableAdminStatus = 2147483647
)

type SlbNewAcclCfgCacheUrlListTableDel int32

const (
	SlbNewAcclCfgCacheUrlListTableDel_Other       SlbNewAcclCfgCacheUrlListTableDel = 1
	SlbNewAcclCfgCacheUrlListTableDel_Delete      SlbNewAcclCfgCacheUrlListTableDel = 2
	SlbNewAcclCfgCacheUrlListTableDel_Unsupported SlbNewAcclCfgCacheUrlListTableDel = 2147483647
)

type SlbNewAcclCfgCacheUrlListTableParams struct {
	// The caching policy URL LIST (key id) as an index.
	IdIndex string `json:"IdIndex,omitempty"`
	// Cache URL LIST name.
	Name string `json:"Name,omitempty"`
	// Status (enable/disable) of caching URL LIST.
	AdminStatus SlbNewAcclCfgCacheUrlListTableAdminStatus `json:"AdminStatus,omitempty"`
	// Delete Caching URL list.
	Del SlbNewAcclCfgCacheUrlListTableDel `json:"Del,omitempty"`
	// Duplicating an entire Rule-List by defining the destination Rule-list id.
	Copy DisplayString `json:"Copy,omitempty"`
}

func (p SlbNewAcclCfgCacheUrlListTableParams) iMABean() {}
