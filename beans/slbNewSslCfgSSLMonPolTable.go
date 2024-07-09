package beans

import (
	"fmt"
	"reflect"
)

// SlbNewSslCfgSSLMonPolTable The table for configuring ssl monitor  policy.
// Note:This mib is not supported for VX instance of Virtualization.
type SlbNewSslCfgSSLMonPolTable struct {
	// The SSL Monitor policy name(key id) as an index.
	SlbNewSslCfgSSLMonPolNameIdIndex string
	Params                           *SlbNewSslCfgSSLMonPolTableParams
}

func NewSlbNewSslCfgSSLMonPolTableList() *SlbNewSslCfgSSLMonPolTable {
	return &SlbNewSslCfgSSLMonPolTable{}
}

func NewSlbNewSslCfgSSLMonPolTable(
	slbNewSslCfgSSLMonPolNameIdIndex string,
	params *SlbNewSslCfgSSLMonPolTableParams,
) *SlbNewSslCfgSSLMonPolTable {
	return &SlbNewSslCfgSSLMonPolTable{
		SlbNewSslCfgSSLMonPolNameIdIndex: slbNewSslCfgSSLMonPolNameIdIndex,
		Params:                           params,
	}
}

func (c *SlbNewSslCfgSSLMonPolTable) Name() string {
	return "SlbNewSslCfgSSLMonPolTable"
}

func (c *SlbNewSslCfgSSLMonPolTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbNewSslCfgSSLMonPolTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbNewSslCfgSSLMonPolTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbNewSslCfgSSLMonPolNameIdIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbNewSslCfgSSLMonPolNameIdIndex)
}

type SlbNewSslCfgSSLMonPolTableHwOffload int32

const (
	SlbNewSslCfgSSLMonPolTableHwOffload_Enabled     SlbNewSslCfgSSLMonPolTableHwOffload = 1
	SlbNewSslCfgSSLMonPolTableHwOffload_Disabled    SlbNewSslCfgSSLMonPolTableHwOffload = 2
	SlbNewSslCfgSSLMonPolTableHwOffload_Unsupported SlbNewSslCfgSSLMonPolTableHwOffload = 2147483647
)

type SlbNewSslCfgSSLMonPolTableAdminStatus int32

const (
	SlbNewSslCfgSSLMonPolTableAdminStatus_Enabled     SlbNewSslCfgSSLMonPolTableAdminStatus = 1
	SlbNewSslCfgSSLMonPolTableAdminStatus_Disabled    SlbNewSslCfgSSLMonPolTableAdminStatus = 2
	SlbNewSslCfgSSLMonPolTableAdminStatus_Unsupported SlbNewSslCfgSSLMonPolTableAdminStatus = 2147483647
)

type SlbNewSslCfgSSLMonPolTableDel int32

const (
	SlbNewSslCfgSSLMonPolTableDel_Other       SlbNewSslCfgSSLMonPolTableDel = 1
	SlbNewSslCfgSSLMonPolTableDel_Delete      SlbNewSslCfgSSLMonPolTableDel = 2
	SlbNewSslCfgSSLMonPolTableDel_Unsupported SlbNewSslCfgSSLMonPolTableDel = 2147483647
)

type SlbNewSslCfgSSLMonPolTableParams struct {
	// The SSL Monitor policy name(key id) as an index.
	NameIdIndex string `json:"NameIdIndex,omitempty"`
	// SSL Monitor  policy name.
	Name string `json:"Name,omitempty"`
	// SSL Monitor policy cache size
	CacheSize uint32 `json:"CacheSize,omitempty"`
	// SSL Monitor policy cache timeout.
	CacheTimeOut uint32 `json:"CacheTimeOut,omitempty"`
	// Disable/Enable HW offload .
	HwOffload SlbNewSslCfgSSLMonPolTableHwOffload `json:"HwOffload,omitempty"`
	// Enable or disable ssl monitor policy.
	AdminStatus SlbNewSslCfgSSLMonPolTableAdminStatus `json:"AdminStatus,omitempty"`
	// Delete SSL Monitor  policy.
	Del SlbNewSslCfgSSLMonPolTableDel `json:"Del,omitempty"`
}

func (p SlbNewSslCfgSSLMonPolTableParams) iMABean() {}
