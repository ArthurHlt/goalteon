package beans

import (
	"fmt"
	"reflect"
)

// SlbCurAcclCfgFastWaTable The table for configuring fastview web application.
// Note:This mib is not supported for VX instance of Virtualization.
type SlbCurAcclCfgFastWaTable struct {
	// The fastview web application ID(key id) as an index.
	SlbCurAcclCfgFastWaNameIdIndex string
	Params                         *SlbCurAcclCfgFastWaTableParams
}

func NewSlbCurAcclCfgFastWaTableList() *SlbCurAcclCfgFastWaTable {
	return &SlbCurAcclCfgFastWaTable{}
}

func NewSlbCurAcclCfgFastWaTable(
	slbCurAcclCfgFastWaNameIdIndex string,
	params *SlbCurAcclCfgFastWaTableParams,
) *SlbCurAcclCfgFastWaTable {
	return &SlbCurAcclCfgFastWaTable{
		SlbCurAcclCfgFastWaNameIdIndex: slbCurAcclCfgFastWaNameIdIndex,
		Params:                         params,
	}
}

func (c *SlbCurAcclCfgFastWaTable) Name() string {
	return "SlbCurAcclCfgFastWaTable"
}

func (c *SlbCurAcclCfgFastWaTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbCurAcclCfgFastWaTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbCurAcclCfgFastWaTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbCurAcclCfgFastWaNameIdIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbCurAcclCfgFastWaNameIdIndex)
}

type SlbCurAcclCfgFastWaTableAdminStatus int32

const (
	SlbCurAcclCfgFastWaTableAdminStatus_Enabled     SlbCurAcclCfgFastWaTableAdminStatus = 1
	SlbCurAcclCfgFastWaTableAdminStatus_Disabled    SlbCurAcclCfgFastWaTableAdminStatus = 2
	SlbCurAcclCfgFastWaTableAdminStatus_Unsupported SlbCurAcclCfgFastWaTableAdminStatus = 2147483647
)

type SlbCurAcclCfgFastWaTableDummy int32

const (
	SlbCurAcclCfgFastWaTableDummy_Enabled     SlbCurAcclCfgFastWaTableDummy = 1
	SlbCurAcclCfgFastWaTableDummy_Disabled    SlbCurAcclCfgFastWaTableDummy = 2
	SlbCurAcclCfgFastWaTableDummy_Unsupported SlbCurAcclCfgFastWaTableDummy = 2147483647
)

type SlbCurAcclCfgFastWaTableParams struct {
	// The fastview web application ID(key id) as an index.
	NameIdIndex string `json:"NameIdIndex,omitempty"`
	// Status (enable/disable) of fastview web application.
	AdminStatus SlbCurAcclCfgFastWaTableAdminStatus `json:"AdminStatus,omitempty"`
	// Dummy field to maintain offsets between cur and new.
	Dummy SlbCurAcclCfgFastWaTableDummy `json:"Dummy,omitempty"`
	// Web application name.
	Name string `json:"Name,omitempty"`
}

func (p SlbCurAcclCfgFastWaTableParams) iMABean() {}
