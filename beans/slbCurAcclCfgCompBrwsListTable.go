package beans

import (
	"fmt"
	"reflect"
)

// SlbCurAcclCfgCompBrwsListTable The table for configuring compression browser LIST.
// Note:This mib is not supported for VX instance of Virtualization.
type SlbCurAcclCfgCompBrwsListTable struct {
	// The compression policy browser LIST (key id) as an index.
	SlbCurAcclCfgCompBrwsListIdIndex string
	Params                           *SlbCurAcclCfgCompBrwsListTableParams
}

func NewSlbCurAcclCfgCompBrwsListTableList() *SlbCurAcclCfgCompBrwsListTable {
	return &SlbCurAcclCfgCompBrwsListTable{}
}

func NewSlbCurAcclCfgCompBrwsListTable(
	slbCurAcclCfgCompBrwsListIdIndex string,
	params *SlbCurAcclCfgCompBrwsListTableParams,
) *SlbCurAcclCfgCompBrwsListTable {
	return &SlbCurAcclCfgCompBrwsListTable{
		SlbCurAcclCfgCompBrwsListIdIndex: slbCurAcclCfgCompBrwsListIdIndex,
		Params:                           params,
	}
}

func (c *SlbCurAcclCfgCompBrwsListTable) Name() string {
	return "SlbCurAcclCfgCompBrwsListTable"
}

func (c *SlbCurAcclCfgCompBrwsListTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbCurAcclCfgCompBrwsListTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbCurAcclCfgCompBrwsListTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbCurAcclCfgCompBrwsListIdIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbCurAcclCfgCompBrwsListIdIndex)
}

type SlbCurAcclCfgCompBrwsListTableAdminStatus int32

const (
	SlbCurAcclCfgCompBrwsListTableAdminStatus_Enabled     SlbCurAcclCfgCompBrwsListTableAdminStatus = 1
	SlbCurAcclCfgCompBrwsListTableAdminStatus_Disabled    SlbCurAcclCfgCompBrwsListTableAdminStatus = 2
	SlbCurAcclCfgCompBrwsListTableAdminStatus_Unsupported SlbCurAcclCfgCompBrwsListTableAdminStatus = 2147483647
)

type SlbCurAcclCfgCompBrwsListTableParams struct {
	// The compression policy browser LIST (key id) as an index.
	IdIndex string `json:"IdIndex,omitempty"`
	// Compression browser list name.
	Name string `json:"Name,omitempty"`
	// Status (enable/disable) of compression browser list.
	AdminStatus SlbCurAcclCfgCompBrwsListTableAdminStatus `json:"AdminStatus,omitempty"`
}

func (p SlbCurAcclCfgCompBrwsListTableParams) iMABean() {}
