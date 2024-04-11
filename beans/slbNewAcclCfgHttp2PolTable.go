package beans

import (
	"fmt"
	"reflect"
)

// SlbNewAcclCfgHttp2PolTable The table for configuring http2 policy.
// Note:This mib is not supported for VX instance of Virtualization.
type SlbNewAcclCfgHttp2PolTable struct {
	// The http2 policy ID(key id) as an index.
	SlbNewAcclCfgHttp2PolNameIdIndex string
	Params                           *SlbNewAcclCfgHttp2PolTableParams
}

func NewSlbNewAcclCfgHttp2PolTableList() *SlbNewAcclCfgHttp2PolTable {
	return &SlbNewAcclCfgHttp2PolTable{}
}

func NewSlbNewAcclCfgHttp2PolTable(
	slbNewAcclCfgHttp2PolNameIdIndex string,
	params *SlbNewAcclCfgHttp2PolTableParams,
) *SlbNewAcclCfgHttp2PolTable {
	return &SlbNewAcclCfgHttp2PolTable{
		SlbNewAcclCfgHttp2PolNameIdIndex: slbNewAcclCfgHttp2PolNameIdIndex,
		Params:                           params,
	}
}

func (c *SlbNewAcclCfgHttp2PolTable) Name() string {
	return "SlbNewAcclCfgHttp2PolTable"
}

func (c *SlbNewAcclCfgHttp2PolTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbNewAcclCfgHttp2PolTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbNewAcclCfgHttp2PolTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbNewAcclCfgHttp2PolNameIdIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbNewAcclCfgHttp2PolNameIdIndex)
}

type SlbNewAcclCfgHttp2PolTableAdminStatus int32

const (
	SlbNewAcclCfgHttp2PolTableAdminStatus_Enabled     SlbNewAcclCfgHttp2PolTableAdminStatus = 1
	SlbNewAcclCfgHttp2PolTableAdminStatus_Disabled    SlbNewAcclCfgHttp2PolTableAdminStatus = 2
	SlbNewAcclCfgHttp2PolTableAdminStatus_Unsupported SlbNewAcclCfgHttp2PolTableAdminStatus = 2147483647
)

type SlbNewAcclCfgHttp2PolTableEnaInsert int32

const (
	SlbNewAcclCfgHttp2PolTableEnaInsert_Enabled     SlbNewAcclCfgHttp2PolTableEnaInsert = 1
	SlbNewAcclCfgHttp2PolTableEnaInsert_Disabled    SlbNewAcclCfgHttp2PolTableEnaInsert = 2
	SlbNewAcclCfgHttp2PolTableEnaInsert_Unsupported SlbNewAcclCfgHttp2PolTableEnaInsert = 2147483647
)

type SlbNewAcclCfgHttp2PolTableEnaServerPush int32

const (
	SlbNewAcclCfgHttp2PolTableEnaServerPush_Enabled     SlbNewAcclCfgHttp2PolTableEnaServerPush = 1
	SlbNewAcclCfgHttp2PolTableEnaServerPush_Disabled    SlbNewAcclCfgHttp2PolTableEnaServerPush = 2
	SlbNewAcclCfgHttp2PolTableEnaServerPush_Unsupported SlbNewAcclCfgHttp2PolTableEnaServerPush = 2147483647
)

type SlbNewAcclCfgHttp2PolTableDelete int32

const (
	SlbNewAcclCfgHttp2PolTableDelete_Other       SlbNewAcclCfgHttp2PolTableDelete = 1
	SlbNewAcclCfgHttp2PolTableDelete_Delete      SlbNewAcclCfgHttp2PolTableDelete = 2
	SlbNewAcclCfgHttp2PolTableDelete_Unsupported SlbNewAcclCfgHttp2PolTableDelete = 2147483647
)

type SlbNewAcclCfgHttp2PolTableBackendStatus int32

const (
	SlbNewAcclCfgHttp2PolTableBackendStatus_Enabled     SlbNewAcclCfgHttp2PolTableBackendStatus = 1
	SlbNewAcclCfgHttp2PolTableBackendStatus_Disabled    SlbNewAcclCfgHttp2PolTableBackendStatus = 2
	SlbNewAcclCfgHttp2PolTableBackendStatus_Unsupported SlbNewAcclCfgHttp2PolTableBackendStatus = 2147483647
)

type SlbNewAcclCfgHttp2PolTableParams struct {
	// The http2 policy ID(key id) as an index.
	NameIdIndex string `json:"NameIdIndex,omitempty"`
	// Http2 policy name.
	Name string `json:"Name,omitempty"`
	// Status (enable/disable) of http2 policy.
	AdminStatus SlbNewAcclCfgHttp2PolTableAdminStatus `json:"AdminStatus,omitempty"`
	// Defines the maximum concurrent streams per connection.
	Streams uint32 `json:"Streams,omitempty"`
	// Defines the number of seconds an HTTP/2 connection is left open idly before it is closed.
	Idle uint64 `json:"Idle,omitempty"`
	// Define (enable/disable) of insert private http2 header.
	EnaInsert SlbNewAcclCfgHttp2PolTableEnaInsert `json:"EnaInsert,omitempty"`
	// Http2 policy header string.
	Header string `json:"Header,omitempty"`
	// Define (enable/disable) of http2 server push.
	EnaServerPush SlbNewAcclCfgHttp2PolTableEnaServerPush `json:"EnaServerPush,omitempty"`
	// Set HTTP2 policy HPACK table size (small / medium / large).
	HpackSize string `json:"HpackSize,omitempty"`
	// Delete Http2 policy.
	Delete SlbNewAcclCfgHttp2PolTableDelete `json:"Delete,omitempty"`
	// Backend Status (enable/disable) of backend http2.
	BackendStatus SlbNewAcclCfgHttp2PolTableBackendStatus `json:"BackendStatus,omitempty"`
	// maximum concurrent streams per backend connection.
	BackendStreams uint32 `json:"BackendStreams,omitempty"`
	// HTTP2 policy backend HPACK table size (small / medium / large).
	BackendHpackSize string `json:"BackendHpackSize,omitempty"`
	// maximum concurrent server push streams per backend connection.
	BackendServerPush uint32 `json:"BackendServerPush,omitempty"`
}

func (p SlbNewAcclCfgHttp2PolTableParams) iMABean() {}
