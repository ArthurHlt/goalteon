package beans

import (
	"fmt"
	"reflect"
)

// Layer7CurCfgHttpmodListTable The table for configuring HTTP content modification rule-lists.
// Note:This mib is not supported for VX instance of virtualization.
type Layer7CurCfgHttpmodListTable struct {
	// HTTP content modification rule-lists name(key id) as an index.
	// Note:This mib is not supported for VX instance of virtualization.
	Layer7CurCfgHttpmodListNameIdIndex string
	Params                             *Layer7CurCfgHttpmodListTableParams
}

func NewLayer7CurCfgHttpmodListTableList() *Layer7CurCfgHttpmodListTable {
	return &Layer7CurCfgHttpmodListTable{}
}

func NewLayer7CurCfgHttpmodListTable(
	layer7CurCfgHttpmodListNameIdIndex string,
	params *Layer7CurCfgHttpmodListTableParams,
) *Layer7CurCfgHttpmodListTable {
	return &Layer7CurCfgHttpmodListTable{
		Layer7CurCfgHttpmodListNameIdIndex: layer7CurCfgHttpmodListNameIdIndex,
		Params:                             params,
	}
}

func (c *Layer7CurCfgHttpmodListTable) Name() string {
	return "Layer7CurCfgHttpmodListTable"
}

func (c *Layer7CurCfgHttpmodListTable) GetParams() BeanType {
	return c.Params
}

func (c *Layer7CurCfgHttpmodListTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *Layer7CurCfgHttpmodListTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.Layer7CurCfgHttpmodListNameIdIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.Layer7CurCfgHttpmodListNameIdIndex)
}

type Layer7CurCfgHttpmodListTableAdminStatus int32

const (
	Layer7CurCfgHttpmodListTableAdminStatus_Enabled     Layer7CurCfgHttpmodListTableAdminStatus = 1
	Layer7CurCfgHttpmodListTableAdminStatus_Disabled    Layer7CurCfgHttpmodListTableAdminStatus = 2
	Layer7CurCfgHttpmodListTableAdminStatus_Unsupported Layer7CurCfgHttpmodListTableAdminStatus = 2147483647
)

type Layer7CurCfgHttpmodListTableParams struct {
	// HTTP content modification rule-lists name(key id) as an index.
	// Note:This mib is not supported for VX instance of virtualization.
	NameIdIndex string `json:"NameIdIndex,omitempty"`
	// HTTP content modification rule-lists name.
	// Note:This mib is not supported for VX instance of virtualization.
	Name string `json:"Name,omitempty"`
	// Status (enable/disable) of rule_list.
	// Note:This mib is not supported for VX instance of virtualization.
	AdminStatus Layer7CurCfgHttpmodListTableAdminStatus `json:"AdminStatus,omitempty"`
}

func (p Layer7CurCfgHttpmodListTableParams) iMABean() {}
