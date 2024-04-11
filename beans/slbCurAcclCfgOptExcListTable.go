package beans

import (
	"fmt"
	"reflect"
)

// SlbCurAcclCfgOptExcListTable The table for configuring optimization LIST.
// Note:This mib is not supported for VX instance of Virtualization.
type SlbCurAcclCfgOptExcListTable struct {
	// The optimization LIST (key id) as an index.
	SlbCurAcclCfgOptExcListIdIndex string
	Params                         *SlbCurAcclCfgOptExcListTableParams
}

func NewSlbCurAcclCfgOptExcListTableList() *SlbCurAcclCfgOptExcListTable {
	return &SlbCurAcclCfgOptExcListTable{}
}

func NewSlbCurAcclCfgOptExcListTable(
	slbCurAcclCfgOptExcListIdIndex string,
	params *SlbCurAcclCfgOptExcListTableParams,
) *SlbCurAcclCfgOptExcListTable {
	return &SlbCurAcclCfgOptExcListTable{
		SlbCurAcclCfgOptExcListIdIndex: slbCurAcclCfgOptExcListIdIndex,
		Params:                         params,
	}
}

func (c *SlbCurAcclCfgOptExcListTable) Name() string {
	return "SlbCurAcclCfgOptExcListTable"
}

func (c *SlbCurAcclCfgOptExcListTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbCurAcclCfgOptExcListTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbCurAcclCfgOptExcListTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbCurAcclCfgOptExcListIdIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbCurAcclCfgOptExcListIdIndex)
}

type SlbCurAcclCfgOptExcListTableAdminStatus int32

const (
	SlbCurAcclCfgOptExcListTableAdminStatus_Enabled     SlbCurAcclCfgOptExcListTableAdminStatus = 1
	SlbCurAcclCfgOptExcListTableAdminStatus_Disabled    SlbCurAcclCfgOptExcListTableAdminStatus = 2
	SlbCurAcclCfgOptExcListTableAdminStatus_Unsupported SlbCurAcclCfgOptExcListTableAdminStatus = 2147483647
)

type SlbCurAcclCfgOptExcListTableParams struct {
	// The optimization LIST (key id) as an index.
	IdIndex string `json:"IdIndex,omitempty"`
	// Optimization LIST name.
	Name string `json:"Name,omitempty"`
	// Status (enable/disable) of optimization LIST.
	AdminStatus SlbCurAcclCfgOptExcListTableAdminStatus `json:"AdminStatus,omitempty"`
}

func (p SlbCurAcclCfgOptExcListTableParams) iMABean() {}
