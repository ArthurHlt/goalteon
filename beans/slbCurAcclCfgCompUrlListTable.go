package beans

import (
	"fmt"
	"reflect"
)

// SlbCurAcclCfgCompUrlListTable The table for configuring compression URL LIST.
// Note:This mib is not supported for VX instance of Virtualization.
type SlbCurAcclCfgCompUrlListTable struct {
	// The compression policy URL LIST (key id) as an index.
	SlbCurAcclCfgCompUrlListIdIndex string
	Params                          *SlbCurAcclCfgCompUrlListTableParams
}

func NewSlbCurAcclCfgCompUrlListTableList() *SlbCurAcclCfgCompUrlListTable {
	return &SlbCurAcclCfgCompUrlListTable{}
}

func NewSlbCurAcclCfgCompUrlListTable(
	slbCurAcclCfgCompUrlListIdIndex string,
	params *SlbCurAcclCfgCompUrlListTableParams,
) *SlbCurAcclCfgCompUrlListTable {
	return &SlbCurAcclCfgCompUrlListTable{
		SlbCurAcclCfgCompUrlListIdIndex: slbCurAcclCfgCompUrlListIdIndex,
		Params:                          params,
	}
}

func (c *SlbCurAcclCfgCompUrlListTable) Name() string {
	return "SlbCurAcclCfgCompUrlListTable"
}

func (c *SlbCurAcclCfgCompUrlListTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbCurAcclCfgCompUrlListTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbCurAcclCfgCompUrlListTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbCurAcclCfgCompUrlListIdIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbCurAcclCfgCompUrlListIdIndex)
}

type SlbCurAcclCfgCompUrlListTableAdminStatus int32

const (
	SlbCurAcclCfgCompUrlListTableAdminStatus_Enabled     SlbCurAcclCfgCompUrlListTableAdminStatus = 1
	SlbCurAcclCfgCompUrlListTableAdminStatus_Disabled    SlbCurAcclCfgCompUrlListTableAdminStatus = 2
	SlbCurAcclCfgCompUrlListTableAdminStatus_Unsupported SlbCurAcclCfgCompUrlListTableAdminStatus = 2147483647
)

type SlbCurAcclCfgCompUrlListTableParams struct {
	// The compression policy URL LIST (key id) as an index.
	IdIndex string `json:"IdIndex,omitempty"`
	// Compression URL LIST name.
	Name string `json:"Name,omitempty"`
	// Status (enable/disable) of compression URL LIST.
	AdminStatus SlbCurAcclCfgCompUrlListTableAdminStatus `json:"AdminStatus,omitempty"`
}

func (p SlbCurAcclCfgCompUrlListTableParams) iMABean() {}
