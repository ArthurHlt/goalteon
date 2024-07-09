package beans

import (
	"fmt"
	"reflect"
)

// SlbCurAcclCfgQuicPolTable The table for configuring quic policy.
// Note:This mib is not supported for VX instance of Virtualization.
type SlbCurAcclCfgQuicPolTable struct {
	// The quic policy ID(key id) as an index.
	SlbCurAcclCfgQuicPolNameIdIndex string
	Params                          *SlbCurAcclCfgQuicPolTableParams
}

func NewSlbCurAcclCfgQuicPolTableList() *SlbCurAcclCfgQuicPolTable {
	return &SlbCurAcclCfgQuicPolTable{}
}

func NewSlbCurAcclCfgQuicPolTable(
	slbCurAcclCfgQuicPolNameIdIndex string,
	params *SlbCurAcclCfgQuicPolTableParams,
) *SlbCurAcclCfgQuicPolTable {
	return &SlbCurAcclCfgQuicPolTable{
		SlbCurAcclCfgQuicPolNameIdIndex: slbCurAcclCfgQuicPolNameIdIndex,
		Params:                          params,
	}
}

func (c *SlbCurAcclCfgQuicPolTable) Name() string {
	return "SlbCurAcclCfgQuicPolTable"
}

func (c *SlbCurAcclCfgQuicPolTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbCurAcclCfgQuicPolTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbCurAcclCfgQuicPolTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbCurAcclCfgQuicPolNameIdIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbCurAcclCfgQuicPolNameIdIndex)
}

type SlbCurAcclCfgQuicPolTableAdminStatus int32

const (
	SlbCurAcclCfgQuicPolTableAdminStatus_Enabled     SlbCurAcclCfgQuicPolTableAdminStatus = 1
	SlbCurAcclCfgQuicPolTableAdminStatus_Disabled    SlbCurAcclCfgQuicPolTableAdminStatus = 2
	SlbCurAcclCfgQuicPolTableAdminStatus_Unsupported SlbCurAcclCfgQuicPolTableAdminStatus = 2147483647
)

type SlbCurAcclCfgQuicPolTableKeepAlive int32

const (
	SlbCurAcclCfgQuicPolTableKeepAlive_Enabled     SlbCurAcclCfgQuicPolTableKeepAlive = 1
	SlbCurAcclCfgQuicPolTableKeepAlive_Disabled    SlbCurAcclCfgQuicPolTableKeepAlive = 2
	SlbCurAcclCfgQuicPolTableKeepAlive_Unsupported SlbCurAcclCfgQuicPolTableKeepAlive = 2147483647
)

type SlbCurAcclCfgQuicPolTableParams struct {
	// The quic policy ID(key id) as an index.
	NameIdIndex string `json:"NameIdIndex,omitempty"`
	// Quic policy name.
	Name string `json:"Name,omitempty"`
	// Status (enable/disable) of quic policy.
	AdminStatus SlbCurAcclCfgQuicPolTableAdminStatus `json:"AdminStatus,omitempty"`
	// Congestion control mechanism (reno / cubic / bbr / brr2).
	Conctrl string `json:"Conctrl,omitempty"`
	// maximum udp size.
	MaxUdpSize uint64 `json:"MaxUdpSize,omitempty"`
	// maximum window size.
	MaxWinSize uint32 `json:"MaxWinSize,omitempty"`
	// init max data size.
	InitMaxDataSize uint64 `json:"InitMaxDataSize,omitempty"`
	// init max stream data size.
	InitMaxStreamSize uint64 `json:"InitMaxStreamSize,omitempty"`
	// Enable/disable keep alive mode.
	KeepAlive SlbCurAcclCfgQuicPolTableKeepAlive `json:"KeepAlive,omitempty"`
	// Set idle timeout.
	IdleTimeout uint64 `json:"IdleTimeout,omitempty"`
	// Ena/Dis Client IP migration.
	ClientMmgration uint32 `json:"ClientMmgration,omitempty"`
}

func (p SlbCurAcclCfgQuicPolTableParams) iMABean() {}
