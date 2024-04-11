package beans

import (
	"fmt"
	"reflect"
)

// SlbNewAcclCfgCompBrwsListTable The table for configuring compression browser list.
// Note:This mib is not supported for VX instance of Virtualization.
type SlbNewAcclCfgCompBrwsListTable struct {
	// The compression policy browser list (key id) as an index.
	SlbNewAcclCfgCompBrwsListIdIndex string
	Params                           *SlbNewAcclCfgCompBrwsListTableParams
}

func NewSlbNewAcclCfgCompBrwsListTableList() *SlbNewAcclCfgCompBrwsListTable {
	return &SlbNewAcclCfgCompBrwsListTable{}
}

func NewSlbNewAcclCfgCompBrwsListTable(
	slbNewAcclCfgCompBrwsListIdIndex string,
	params *SlbNewAcclCfgCompBrwsListTableParams,
) *SlbNewAcclCfgCompBrwsListTable {
	return &SlbNewAcclCfgCompBrwsListTable{
		SlbNewAcclCfgCompBrwsListIdIndex: slbNewAcclCfgCompBrwsListIdIndex,
		Params:                           params,
	}
}

func (c *SlbNewAcclCfgCompBrwsListTable) Name() string {
	return "SlbNewAcclCfgCompBrwsListTable"
}

func (c *SlbNewAcclCfgCompBrwsListTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbNewAcclCfgCompBrwsListTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbNewAcclCfgCompBrwsListTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbNewAcclCfgCompBrwsListIdIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbNewAcclCfgCompBrwsListIdIndex)
}

type SlbNewAcclCfgCompBrwsListTableAdminStatus int32

const (
	SlbNewAcclCfgCompBrwsListTableAdminStatus_Enabled     SlbNewAcclCfgCompBrwsListTableAdminStatus = 1
	SlbNewAcclCfgCompBrwsListTableAdminStatus_Disabled    SlbNewAcclCfgCompBrwsListTableAdminStatus = 2
	SlbNewAcclCfgCompBrwsListTableAdminStatus_Unsupported SlbNewAcclCfgCompBrwsListTableAdminStatus = 2147483647
)

type SlbNewAcclCfgCompBrwsListTableDel int32

const (
	SlbNewAcclCfgCompBrwsListTableDel_Other       SlbNewAcclCfgCompBrwsListTableDel = 1
	SlbNewAcclCfgCompBrwsListTableDel_Delete      SlbNewAcclCfgCompBrwsListTableDel = 2
	SlbNewAcclCfgCompBrwsListTableDel_Unsupported SlbNewAcclCfgCompBrwsListTableDel = 2147483647
)

type SlbNewAcclCfgCompBrwsListTableParams struct {
	// The compression policy browser list (key id) as an index.
	IdIndex string `json:"IdIndex,omitempty"`
	// Compression browser list name.
	Name string `json:"Name,omitempty"`
	// Status (enable/disable) of compression browser list.
	AdminStatus SlbNewAcclCfgCompBrwsListTableAdminStatus `json:"AdminStatus,omitempty"`
	// Delete Compression Browser list.
	Del SlbNewAcclCfgCompBrwsListTableDel `json:"Del,omitempty"`
	// Duplicating an entire browser list by defining the destination browser list id.
	Copy string `json:"Copy,omitempty"`
}

func (p SlbNewAcclCfgCompBrwsListTableParams) iMABean() {}
