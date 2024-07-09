package beans

import (
	"fmt"
	"reflect"
)

// SlbCurAcclCfgHttp3PolTable The table for configuring http3 policy.
// Note:This mib is not supported for VX instance of Virtualization.
type SlbCurAcclCfgHttp3PolTable struct {
	// The http3 policy ID(key id) as an index.
	SlbCurAcclCfgHttp3PolNameIdIndex string
	Params                           *SlbCurAcclCfgHttp3PolTableParams
}

func NewSlbCurAcclCfgHttp3PolTableList() *SlbCurAcclCfgHttp3PolTable {
	return &SlbCurAcclCfgHttp3PolTable{}
}

func NewSlbCurAcclCfgHttp3PolTable(
	slbCurAcclCfgHttp3PolNameIdIndex string,
	params *SlbCurAcclCfgHttp3PolTableParams,
) *SlbCurAcclCfgHttp3PolTable {
	return &SlbCurAcclCfgHttp3PolTable{
		SlbCurAcclCfgHttp3PolNameIdIndex: slbCurAcclCfgHttp3PolNameIdIndex,
		Params:                           params,
	}
}

func (c *SlbCurAcclCfgHttp3PolTable) Name() string {
	return "SlbCurAcclCfgHttp3PolTable"
}

func (c *SlbCurAcclCfgHttp3PolTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbCurAcclCfgHttp3PolTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbCurAcclCfgHttp3PolTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbCurAcclCfgHttp3PolNameIdIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbCurAcclCfgHttp3PolNameIdIndex)
}

type SlbCurAcclCfgHttp3PolTableAdminStatus int32

const (
	SlbCurAcclCfgHttp3PolTableAdminStatus_Enabled     SlbCurAcclCfgHttp3PolTableAdminStatus = 1
	SlbCurAcclCfgHttp3PolTableAdminStatus_Disabled    SlbCurAcclCfgHttp3PolTableAdminStatus = 2
	SlbCurAcclCfgHttp3PolTableAdminStatus_Unsupported SlbCurAcclCfgHttp3PolTableAdminStatus = 2147483647
)

type SlbCurAcclCfgHttp3PolTableParams struct {
	// The http3 policy ID(key id) as an index.
	NameIdIndex string `json:"NameIdIndex,omitempty"`
	// Http3 policy name.
	Name string `json:"Name,omitempty"`
	// Status (enable/disable) of http3 policy.
	AdminStatus SlbCurAcclCfgHttp3PolTableAdminStatus `json:"AdminStatus,omitempty"`
	// maximum concurrent streams per connection.
	Streams uint32 `json:"Streams,omitempty"`
	// HTTP3 policy QPACK table size (small / medium / large).
	QpackSize string `json:"QpackSize,omitempty"`
}

func (p SlbCurAcclCfgHttp3PolTableParams) iMABean() {}
