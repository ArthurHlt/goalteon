package beans

import (
	"fmt"
	"reflect"
)

// SlbNewAcclCfgOptExcListTable The table for configuring optimization LIST.
// Note:This mib is not supported for VX instance of Virtualization.
type SlbNewAcclCfgOptExcListTable struct {
	// The caching optimization LIST (key id) as an index.
	SlbNewAcclCfgOptExcListIdIndex string
	Params                         *SlbNewAcclCfgOptExcListTableParams
}

func NewSlbNewAcclCfgOptExcListTableList() *SlbNewAcclCfgOptExcListTable {
	return &SlbNewAcclCfgOptExcListTable{}
}

func NewSlbNewAcclCfgOptExcListTable(
	slbNewAcclCfgOptExcListIdIndex string,
	params *SlbNewAcclCfgOptExcListTableParams,
) *SlbNewAcclCfgOptExcListTable {
	return &SlbNewAcclCfgOptExcListTable{
		SlbNewAcclCfgOptExcListIdIndex: slbNewAcclCfgOptExcListIdIndex,
		Params:                         params,
	}
}

func (c *SlbNewAcclCfgOptExcListTable) Name() string {
	return "SlbNewAcclCfgOptExcListTable"
}

func (c *SlbNewAcclCfgOptExcListTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbNewAcclCfgOptExcListTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbNewAcclCfgOptExcListTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbNewAcclCfgOptExcListIdIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbNewAcclCfgOptExcListIdIndex)
}

type SlbNewAcclCfgOptExcListTableAdminStatus int32

const (
	SlbNewAcclCfgOptExcListTableAdminStatus_Enabled     SlbNewAcclCfgOptExcListTableAdminStatus = 1
	SlbNewAcclCfgOptExcListTableAdminStatus_Disabled    SlbNewAcclCfgOptExcListTableAdminStatus = 2
	SlbNewAcclCfgOptExcListTableAdminStatus_Unsupported SlbNewAcclCfgOptExcListTableAdminStatus = 2147483647
)

type SlbNewAcclCfgOptExcListTableDel int32

const (
	SlbNewAcclCfgOptExcListTableDel_Other       SlbNewAcclCfgOptExcListTableDel = 1
	SlbNewAcclCfgOptExcListTableDel_Delete      SlbNewAcclCfgOptExcListTableDel = 2
	SlbNewAcclCfgOptExcListTableDel_Unsupported SlbNewAcclCfgOptExcListTableDel = 2147483647
)

type SlbNewAcclCfgOptExcListTableParams struct {
	// The caching optimization LIST (key id) as an index.
	IdIndex string `json:"IdIndex,omitempty"`
	// Optimization LIST name.
	Name string `json:"Name,omitempty"`
	// Status (enable/disable) of optimization LIST.
	AdminStatus SlbNewAcclCfgOptExcListTableAdminStatus `json:"AdminStatus,omitempty"`
	// Delete optimization list.
	Del SlbNewAcclCfgOptExcListTableDel `json:"Del,omitempty"`
	// Duplicating an entire Rule-List by defining the destination Rule-list id.
	Copy string `json:"Copy,omitempty"`
}

func (p SlbNewAcclCfgOptExcListTableParams) iMABean() {}
