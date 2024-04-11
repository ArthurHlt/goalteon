package beans

import (
	"fmt"
	"reflect"
)

// SlbNewAcclCfgCompUrlListTable The table for configuring compression URL LIST.
// Note:This mib is not supported for VX instance of Virtualization.
type SlbNewAcclCfgCompUrlListTable struct {
	// The compression policy URL LIST (key id) as an index.
	SlbNewAcclCfgCompUrlListIdIndex string
	Params                          *SlbNewAcclCfgCompUrlListTableParams
}

func NewSlbNewAcclCfgCompUrlListTableList() *SlbNewAcclCfgCompUrlListTable {
	return &SlbNewAcclCfgCompUrlListTable{}
}

func NewSlbNewAcclCfgCompUrlListTable(
	slbNewAcclCfgCompUrlListIdIndex string,
	params *SlbNewAcclCfgCompUrlListTableParams,
) *SlbNewAcclCfgCompUrlListTable {
	return &SlbNewAcclCfgCompUrlListTable{
		SlbNewAcclCfgCompUrlListIdIndex: slbNewAcclCfgCompUrlListIdIndex,
		Params:                          params,
	}
}

func (c *SlbNewAcclCfgCompUrlListTable) Name() string {
	return "SlbNewAcclCfgCompUrlListTable"
}

func (c *SlbNewAcclCfgCompUrlListTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbNewAcclCfgCompUrlListTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbNewAcclCfgCompUrlListTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbNewAcclCfgCompUrlListIdIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbNewAcclCfgCompUrlListIdIndex)
}

type SlbNewAcclCfgCompUrlListTableAdminStatus int32

const (
	SlbNewAcclCfgCompUrlListTableAdminStatus_Enabled     SlbNewAcclCfgCompUrlListTableAdminStatus = 1
	SlbNewAcclCfgCompUrlListTableAdminStatus_Disabled    SlbNewAcclCfgCompUrlListTableAdminStatus = 2
	SlbNewAcclCfgCompUrlListTableAdminStatus_Unsupported SlbNewAcclCfgCompUrlListTableAdminStatus = 2147483647
)

type SlbNewAcclCfgCompUrlListTableDel int32

const (
	SlbNewAcclCfgCompUrlListTableDel_Other       SlbNewAcclCfgCompUrlListTableDel = 1
	SlbNewAcclCfgCompUrlListTableDel_Delete      SlbNewAcclCfgCompUrlListTableDel = 2
	SlbNewAcclCfgCompUrlListTableDel_Unsupported SlbNewAcclCfgCompUrlListTableDel = 2147483647
)

type SlbNewAcclCfgCompUrlListTableParams struct {
	// The compression policy URL LIST (key id) as an index.
	IdIndex string `json:"IdIndex,omitempty"`
	// Compression URL LIST name.
	Name string `json:"Name,omitempty"`
	// Status (enable/disable) of compression URL LIST.
	AdminStatus SlbNewAcclCfgCompUrlListTableAdminStatus `json:"AdminStatus,omitempty"`
	// Delete Compression URL list.
	Del SlbNewAcclCfgCompUrlListTableDel `json:"Del,omitempty"`
	// Duplicating an entire Rule-List by defining the destination Rule-list id.
	Copy string `json:"Copy,omitempty"`
}

func (p SlbNewAcclCfgCompUrlListTableParams) iMABean() {}
