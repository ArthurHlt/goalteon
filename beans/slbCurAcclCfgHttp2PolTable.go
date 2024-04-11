package beans

import (
	"fmt"
	"reflect"
)

// SlbCurAcclCfgHttp2PolTable The table for configuring http2 policy.
// Note:This mib is not supported for VX instance of Virtualization.
type SlbCurAcclCfgHttp2PolTable struct {
	// The http2 policy ID(key id) as an index.
	SlbCurAcclCfgHttp2PolNameIdIndex string
	Params                           *SlbCurAcclCfgHttp2PolTableParams
}

func NewSlbCurAcclCfgHttp2PolTableList() *SlbCurAcclCfgHttp2PolTable {
	return &SlbCurAcclCfgHttp2PolTable{}
}

func NewSlbCurAcclCfgHttp2PolTable(
	slbCurAcclCfgHttp2PolNameIdIndex string,
	params *SlbCurAcclCfgHttp2PolTableParams,
) *SlbCurAcclCfgHttp2PolTable {
	return &SlbCurAcclCfgHttp2PolTable{
		SlbCurAcclCfgHttp2PolNameIdIndex: slbCurAcclCfgHttp2PolNameIdIndex,
		Params:                           params,
	}
}

func (c *SlbCurAcclCfgHttp2PolTable) Name() string {
	return "SlbCurAcclCfgHttp2PolTable"
}

func (c *SlbCurAcclCfgHttp2PolTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbCurAcclCfgHttp2PolTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbCurAcclCfgHttp2PolTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbCurAcclCfgHttp2PolNameIdIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbCurAcclCfgHttp2PolNameIdIndex)
}

type SlbCurAcclCfgHttp2PolTableAdminStatus int32

const (
	SlbCurAcclCfgHttp2PolTableAdminStatus_Enabled     SlbCurAcclCfgHttp2PolTableAdminStatus = 1
	SlbCurAcclCfgHttp2PolTableAdminStatus_Disabled    SlbCurAcclCfgHttp2PolTableAdminStatus = 2
	SlbCurAcclCfgHttp2PolTableAdminStatus_Unsupported SlbCurAcclCfgHttp2PolTableAdminStatus = 2147483647
)

type SlbCurAcclCfgHttp2PolTableEnaInsert int32

const (
	SlbCurAcclCfgHttp2PolTableEnaInsert_Enabled     SlbCurAcclCfgHttp2PolTableEnaInsert = 1
	SlbCurAcclCfgHttp2PolTableEnaInsert_Disabled    SlbCurAcclCfgHttp2PolTableEnaInsert = 2
	SlbCurAcclCfgHttp2PolTableEnaInsert_Unsupported SlbCurAcclCfgHttp2PolTableEnaInsert = 2147483647
)

type SlbCurAcclCfgHttp2PolTableEnaServerPush int32

const (
	SlbCurAcclCfgHttp2PolTableEnaServerPush_Enabled     SlbCurAcclCfgHttp2PolTableEnaServerPush = 1
	SlbCurAcclCfgHttp2PolTableEnaServerPush_Disabled    SlbCurAcclCfgHttp2PolTableEnaServerPush = 2
	SlbCurAcclCfgHttp2PolTableEnaServerPush_Unsupported SlbCurAcclCfgHttp2PolTableEnaServerPush = 2147483647
)

type SlbCurAcclCfgHttp2PolTableBackendStatus int32

const (
	SlbCurAcclCfgHttp2PolTableBackendStatus_Enabled     SlbCurAcclCfgHttp2PolTableBackendStatus = 1
	SlbCurAcclCfgHttp2PolTableBackendStatus_Disabled    SlbCurAcclCfgHttp2PolTableBackendStatus = 2
	SlbCurAcclCfgHttp2PolTableBackendStatus_Unsupported SlbCurAcclCfgHttp2PolTableBackendStatus = 2147483647
)

type SlbCurAcclCfgHttp2PolTableParams struct {
	// The http2 policy ID(key id) as an index.
	NameIdIndex string `json:"NameIdIndex,omitempty"`
	// Http2 policy name.
	Name string `json:"Name,omitempty"`
	// Status (enable/disable) of http2 policy.
	AdminStatus SlbCurAcclCfgHttp2PolTableAdminStatus `json:"AdminStatus,omitempty"`
	// maximum concurrent streams per connection.
	Streams uint32 `json:"Streams,omitempty"`
	// Display the number of seconds an HTTP/2 connection is left open idly before it is closed.
	Idle uint64 `json:"Idle,omitempty"`
	// Status (enable/disable) of insert private http2 header.
	EnaInsert SlbCurAcclCfgHttp2PolTableEnaInsert `json:"EnaInsert,omitempty"`
	// Http2 header string.
	Header string `json:"Header,omitempty"`
	// Status (enable/disable) of insert http2 server push.
	EnaServerPush SlbCurAcclCfgHttp2PolTableEnaServerPush `json:"EnaServerPush,omitempty"`
	// HTTP2 policy HPACK table size (small / medium / large).
	HpackSize string `json:"HpackSize,omitempty"`
	// Backend Status (enable/disable) of backend http2.
	BackendStatus SlbCurAcclCfgHttp2PolTableBackendStatus `json:"BackendStatus,omitempty"`
	// maximum concurrent streams per backend connection.
	BackendStreams uint32 `json:"BackendStreams,omitempty"`
	// HTTP2 policy backend HPACK table size (small / medium / large).
	BackendHpackSize string `json:"BackendHpackSize,omitempty"`
	// maximum concurrent server push streams per backend connection.
	BackendServerPush uint32 `json:"BackendServerPush,omitempty"`
}

func (p SlbCurAcclCfgHttp2PolTableParams) iMABean() {}
