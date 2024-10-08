package beans

import (
	"fmt"
	"reflect"
)

// Layer7NewCfgHttpmodListTable The table for configuring HTTP content modification rule-lists.
// Note:This mib is not supported for VX instance of virtualization.
type Layer7NewCfgHttpmodListTable struct {
	// HTTP content modification rule-lists name(key id) as an index, length of the string should be 32 characters.
	Layer7NewCfgHttpmodListNameIdIndex string
	Params                             *Layer7NewCfgHttpmodListTableParams
}

func NewLayer7NewCfgHttpmodListTableList() *Layer7NewCfgHttpmodListTable {
	return &Layer7NewCfgHttpmodListTable{}
}

func NewLayer7NewCfgHttpmodListTable(
	layer7NewCfgHttpmodListNameIdIndex string,
	params *Layer7NewCfgHttpmodListTableParams,
) *Layer7NewCfgHttpmodListTable {
	return &Layer7NewCfgHttpmodListTable{
		Layer7NewCfgHttpmodListNameIdIndex: layer7NewCfgHttpmodListNameIdIndex,
		Params:                             params,
	}
}

func (c *Layer7NewCfgHttpmodListTable) Name() string {
	return "Layer7NewCfgHttpmodListTable"
}

func (c *Layer7NewCfgHttpmodListTable) GetParams() BeanType {
	return c.Params
}

func (c *Layer7NewCfgHttpmodListTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *Layer7NewCfgHttpmodListTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.Layer7NewCfgHttpmodListNameIdIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.Layer7NewCfgHttpmodListNameIdIndex)
}

type Layer7NewCfgHttpmodListTableAdminStatus int32

const (
	Layer7NewCfgHttpmodListTableAdminStatus_Enabled     Layer7NewCfgHttpmodListTableAdminStatus = 1
	Layer7NewCfgHttpmodListTableAdminStatus_Disabled    Layer7NewCfgHttpmodListTableAdminStatus = 2
	Layer7NewCfgHttpmodListTableAdminStatus_Unsupported Layer7NewCfgHttpmodListTableAdminStatus = 2147483647
)

type Layer7NewCfgHttpmodListTableDelete int32

const (
	Layer7NewCfgHttpmodListTableDelete_Other       Layer7NewCfgHttpmodListTableDelete = 1
	Layer7NewCfgHttpmodListTableDelete_Delete      Layer7NewCfgHttpmodListTableDelete = 2
	Layer7NewCfgHttpmodListTableDelete_Unsupported Layer7NewCfgHttpmodListTableDelete = 2147483647
)

type Layer7NewCfgHttpmodListTableParams struct {
	// HTTP content modification rule-lists name(key id) as an index, length of the string should be 32 characters.
	NameIdIndex string `json:"NameIdIndex,omitempty"`
	// HTTP content modification rule-list name, length of the string should be 32 characters.
	Name string `json:"Name,omitempty"`
	// Status (enable/disable) of rule_list.
	AdminStatus Layer7NewCfgHttpmodListTableAdminStatus `json:"AdminStatus,omitempty"`
	// Copy rule-list.
	Copy DisplayString `json:"Copy,omitempty"`
	// Delete HTTP modifications rule-list.
	Delete Layer7NewCfgHttpmodListTableDelete `json:"Delete,omitempty"`
}

func (p Layer7NewCfgHttpmodListTableParams) iMABean() {}
