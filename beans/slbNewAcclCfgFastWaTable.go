package beans

import (
	"fmt"
	"reflect"
)

// SlbNewAcclCfgFastWaTable The table for configuring fastview web application.
// Note:This mib is not supported for VX instance of Virtualization.
type SlbNewAcclCfgFastWaTable struct {
	// The fastview web application ID(key id) as an index, length of the string should be 32 characters.
	SlbNewAcclCfgFastWaNameIdIndex string
	Params                         *SlbNewAcclCfgFastWaTableParams
}

func NewSlbNewAcclCfgFastWaTableList() *SlbNewAcclCfgFastWaTable {
	return &SlbNewAcclCfgFastWaTable{}
}

func NewSlbNewAcclCfgFastWaTable(
	slbNewAcclCfgFastWaNameIdIndex string,
	params *SlbNewAcclCfgFastWaTableParams,
) *SlbNewAcclCfgFastWaTable {
	return &SlbNewAcclCfgFastWaTable{
		SlbNewAcclCfgFastWaNameIdIndex: slbNewAcclCfgFastWaNameIdIndex,
		Params:                         params,
	}
}

func (c *SlbNewAcclCfgFastWaTable) Name() string {
	return "SlbNewAcclCfgFastWaTable"
}

func (c *SlbNewAcclCfgFastWaTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbNewAcclCfgFastWaTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbNewAcclCfgFastWaTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbNewAcclCfgFastWaNameIdIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbNewAcclCfgFastWaNameIdIndex)
}

type SlbNewAcclCfgFastWaTableAdminStatus int32

const (
	SlbNewAcclCfgFastWaTableAdminStatus_Enabled     SlbNewAcclCfgFastWaTableAdminStatus = 1
	SlbNewAcclCfgFastWaTableAdminStatus_Disabled    SlbNewAcclCfgFastWaTableAdminStatus = 2
	SlbNewAcclCfgFastWaTableAdminStatus_Unsupported SlbNewAcclCfgFastWaTableAdminStatus = 2147483647
)

type SlbNewAcclCfgFastWaTableDelete int32

const (
	SlbNewAcclCfgFastWaTableDelete_Other       SlbNewAcclCfgFastWaTableDelete = 1
	SlbNewAcclCfgFastWaTableDelete_Delete      SlbNewAcclCfgFastWaTableDelete = 2
	SlbNewAcclCfgFastWaTableDelete_Unsupported SlbNewAcclCfgFastWaTableDelete = 2147483647
)

type SlbNewAcclCfgFastWaTableParams struct {
	// The fastview web application ID(key id) as an index, length of the string should be 32 characters.
	NameIdIndex string `json:"NameIdIndex,omitempty"`
	// Enable or disable fastview web application.
	AdminStatus SlbNewAcclCfgFastWaTableAdminStatus `json:"AdminStatus,omitempty"`
	// Delete fastview web application.
	Delete SlbNewAcclCfgFastWaTableDelete `json:"Delete,omitempty"`
	// Web application name.
	Name string `json:"Name,omitempty"`
	// Web application name.
	Copy string `json:"Copy,omitempty"`
}

func (p SlbNewAcclCfgFastWaTableParams) iMABean() {}
