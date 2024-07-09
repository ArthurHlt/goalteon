package beans

import (
	"fmt"
	"reflect"
)

// SlbCurSslCfgSSLMonPolTable The table for configuring ssl  monitor policy.
// Note:This mib is not supported for VX instance of Virtualization.
type SlbCurSslCfgSSLMonPolTable struct {
	// The SSL Monitor policy name(key id) as an index.
	SlbCurSslCfgSSLMonPolNameIdIndex string
	Params                           *SlbCurSslCfgSSLMonPolTableParams
}

func NewSlbCurSslCfgSSLMonPolTableList() *SlbCurSslCfgSSLMonPolTable {
	return &SlbCurSslCfgSSLMonPolTable{}
}

func NewSlbCurSslCfgSSLMonPolTable(
	slbCurSslCfgSSLMonPolNameIdIndex string,
	params *SlbCurSslCfgSSLMonPolTableParams,
) *SlbCurSslCfgSSLMonPolTable {
	return &SlbCurSslCfgSSLMonPolTable{
		SlbCurSslCfgSSLMonPolNameIdIndex: slbCurSslCfgSSLMonPolNameIdIndex,
		Params:                           params,
	}
}

func (c *SlbCurSslCfgSSLMonPolTable) Name() string {
	return "SlbCurSslCfgSSLMonPolTable"
}

func (c *SlbCurSslCfgSSLMonPolTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbCurSslCfgSSLMonPolTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbCurSslCfgSSLMonPolTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbCurSslCfgSSLMonPolNameIdIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbCurSslCfgSSLMonPolNameIdIndex)
}

type SlbCurSslCfgSSLMonPolTableHwOffload int32

const (
	SlbCurSslCfgSSLMonPolTableHwOffload_Enabled     SlbCurSslCfgSSLMonPolTableHwOffload = 1
	SlbCurSslCfgSSLMonPolTableHwOffload_Disabled    SlbCurSslCfgSSLMonPolTableHwOffload = 2
	SlbCurSslCfgSSLMonPolTableHwOffload_Unsupported SlbCurSslCfgSSLMonPolTableHwOffload = 2147483647
)

type SlbCurSslCfgSSLMonPolTableAdminStatus int32

const (
	SlbCurSslCfgSSLMonPolTableAdminStatus_Enabled     SlbCurSslCfgSSLMonPolTableAdminStatus = 1
	SlbCurSslCfgSSLMonPolTableAdminStatus_Disabled    SlbCurSslCfgSSLMonPolTableAdminStatus = 2
	SlbCurSslCfgSSLMonPolTableAdminStatus_Unsupported SlbCurSslCfgSSLMonPolTableAdminStatus = 2147483647
)

type SlbCurSslCfgSSLMonPolTableParams struct {
	// The SSL Monitor policy name(key id) as an index.
	NameIdIndex string `json:"NameIdIndex,omitempty"`
	// SSL Monitor  policy name.
	Name string `json:"Name,omitempty"`
	// SSL Monitor policy cache size
	CacheSize int32 `json:"CacheSize,omitempty"`
	// SSL Monitor policy cache timeout.
	CacheTimeOut int32 `json:"CacheTimeOut,omitempty"`
	// Disable/Enable HW offload .
	HwOffload SlbCurSslCfgSSLMonPolTableHwOffload `json:"HwOffload,omitempty"`
	// Enable or disable ssl monitor policy.
	AdminStatus SlbCurSslCfgSSLMonPolTableAdminStatus `json:"AdminStatus,omitempty"`
}

func (p SlbCurSslCfgSSLMonPolTableParams) iMABean() {}
