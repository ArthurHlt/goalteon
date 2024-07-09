package beans

import (
	"fmt"
	"reflect"
)

// SlbNewAcclCfgHttp3PolTable The table for configuring http3 policy.
// Note:This mib is not supported for VX instance of Virtualization.
type SlbNewAcclCfgHttp3PolTable struct {
	// The http3 policy ID(key id) as an index.
	SlbNewAcclCfgHttp3PolNameIdIndex string
	Params                           *SlbNewAcclCfgHttp3PolTableParams
}

func NewSlbNewAcclCfgHttp3PolTableList() *SlbNewAcclCfgHttp3PolTable {
	return &SlbNewAcclCfgHttp3PolTable{}
}

func NewSlbNewAcclCfgHttp3PolTable(
	slbNewAcclCfgHttp3PolNameIdIndex string,
	params *SlbNewAcclCfgHttp3PolTableParams,
) *SlbNewAcclCfgHttp3PolTable {
	return &SlbNewAcclCfgHttp3PolTable{
		SlbNewAcclCfgHttp3PolNameIdIndex: slbNewAcclCfgHttp3PolNameIdIndex,
		Params:                           params,
	}
}

func (c *SlbNewAcclCfgHttp3PolTable) Name() string {
	return "SlbNewAcclCfgHttp3PolTable"
}

func (c *SlbNewAcclCfgHttp3PolTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbNewAcclCfgHttp3PolTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbNewAcclCfgHttp3PolTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbNewAcclCfgHttp3PolNameIdIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbNewAcclCfgHttp3PolNameIdIndex)
}

type SlbNewAcclCfgHttp3PolTableAdminStatus int32

const (
	SlbNewAcclCfgHttp3PolTableAdminStatus_Enabled     SlbNewAcclCfgHttp3PolTableAdminStatus = 1
	SlbNewAcclCfgHttp3PolTableAdminStatus_Disabled    SlbNewAcclCfgHttp3PolTableAdminStatus = 2
	SlbNewAcclCfgHttp3PolTableAdminStatus_Unsupported SlbNewAcclCfgHttp3PolTableAdminStatus = 2147483647
)

type SlbNewAcclCfgHttp3PolTableDelete int32

const (
	SlbNewAcclCfgHttp3PolTableDelete_Other       SlbNewAcclCfgHttp3PolTableDelete = 1
	SlbNewAcclCfgHttp3PolTableDelete_Delete      SlbNewAcclCfgHttp3PolTableDelete = 2
	SlbNewAcclCfgHttp3PolTableDelete_Unsupported SlbNewAcclCfgHttp3PolTableDelete = 2147483647
)

type SlbNewAcclCfgHttp3PolTableParams struct {
	// The http3 policy ID(key id) as an index.
	NameIdIndex string `json:"NameIdIndex,omitempty"`
	// Http3 policy name.
	Name string `json:"Name,omitempty"`
	// Status (enable/disable) of http3 policy.
	AdminStatus SlbNewAcclCfgHttp3PolTableAdminStatus `json:"AdminStatus,omitempty"`
	// Defines the maximum concurrent streams per connection.
	Streams uint32 `json:"Streams,omitempty"`
	// Set HTTP3 policy QPACK table size (small / medium / large).
	QpackSize string `json:"QpackSize,omitempty"`
	// Delete Http3 policy.
	Delete SlbNewAcclCfgHttp3PolTableDelete `json:"Delete,omitempty"`
}

func (p SlbNewAcclCfgHttp3PolTableParams) iMABean() {}
