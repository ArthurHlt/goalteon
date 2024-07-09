package beans

import (
	"fmt"
	"reflect"
)

// SlbNewAcclCfgQuicPolTable The table for configuring quic policy.
// Note:This mib is not supported for VX instance of Virtualization.
type SlbNewAcclCfgQuicPolTable struct {
	// The quic policy ID(key id) as an index.
	SlbNewAcclCfgQuicPolNameIdIndex string
	Params                          *SlbNewAcclCfgQuicPolTableParams
}

func NewSlbNewAcclCfgQuicPolTableList() *SlbNewAcclCfgQuicPolTable {
	return &SlbNewAcclCfgQuicPolTable{}
}

func NewSlbNewAcclCfgQuicPolTable(
	slbNewAcclCfgQuicPolNameIdIndex string,
	params *SlbNewAcclCfgQuicPolTableParams,
) *SlbNewAcclCfgQuicPolTable {
	return &SlbNewAcclCfgQuicPolTable{
		SlbNewAcclCfgQuicPolNameIdIndex: slbNewAcclCfgQuicPolNameIdIndex,
		Params:                          params,
	}
}

func (c *SlbNewAcclCfgQuicPolTable) Name() string {
	return "SlbNewAcclCfgQuicPolTable"
}

func (c *SlbNewAcclCfgQuicPolTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbNewAcclCfgQuicPolTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbNewAcclCfgQuicPolTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbNewAcclCfgQuicPolNameIdIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbNewAcclCfgQuicPolNameIdIndex)
}

type SlbNewAcclCfgQuicPolTableAdminStatus int32

const (
	SlbNewAcclCfgQuicPolTableAdminStatus_Enabled     SlbNewAcclCfgQuicPolTableAdminStatus = 1
	SlbNewAcclCfgQuicPolTableAdminStatus_Disabled    SlbNewAcclCfgQuicPolTableAdminStatus = 2
	SlbNewAcclCfgQuicPolTableAdminStatus_Unsupported SlbNewAcclCfgQuicPolTableAdminStatus = 2147483647
)

type SlbNewAcclCfgQuicPolTableDelete int32

const (
	SlbNewAcclCfgQuicPolTableDelete_Other       SlbNewAcclCfgQuicPolTableDelete = 1
	SlbNewAcclCfgQuicPolTableDelete_Delete      SlbNewAcclCfgQuicPolTableDelete = 2
	SlbNewAcclCfgQuicPolTableDelete_Unsupported SlbNewAcclCfgQuicPolTableDelete = 2147483647
)

type SlbNewAcclCfgQuicPolTableKeepAlive int32

const (
	SlbNewAcclCfgQuicPolTableKeepAlive_Enabled     SlbNewAcclCfgQuicPolTableKeepAlive = 1
	SlbNewAcclCfgQuicPolTableKeepAlive_Disabled    SlbNewAcclCfgQuicPolTableKeepAlive = 2
	SlbNewAcclCfgQuicPolTableKeepAlive_Unsupported SlbNewAcclCfgQuicPolTableKeepAlive = 2147483647
)

type SlbNewAcclCfgQuicPolTableParams struct {
	// The quic policy ID(key id) as an index.
	NameIdIndex string `json:"NameIdIndex,omitempty"`
	// Quic policy name.
	Name string `json:"Name,omitempty"`
	// Status (enable/disable) of quic policy.
	AdminStatus SlbNewAcclCfgQuicPolTableAdminStatus `json:"AdminStatus,omitempty"`
	// Congestion control mechanism (reno / cubic / bbr / bbr2).
	Conctrl string `json:"Conctrl,omitempty"`
	// Max udp size.
	MaxUdpSize uint64 `json:"MaxUdpSize,omitempty"`
	// Quic max window size in Mega
	MaxWinSize uint32 `json:"MaxWinSize,omitempty"`
	// Delete Quic policy.
	Delete SlbNewAcclCfgQuicPolTableDelete `json:"Delete,omitempty"`
	// initial Max data size.
	InitMaxDataSize uint64 `json:"InitMaxDataSize,omitempty"`
	// initial Max stream size.
	InitMaxStreamSize uint64 `json:"InitMaxStreamSize,omitempty"`
	// Enable/disable keep alive mode.
	KeepAlive SlbNewAcclCfgQuicPolTableKeepAlive `json:"KeepAlive,omitempty"`
	// Set idle timeout.
	IdleTimeout uint64 `json:"IdleTimeout,omitempty"`
	// Ena/Dis Client IP migration.
	ClientMmgration uint32 `json:"ClientMmgration,omitempty"`
}

func (p SlbNewAcclCfgQuicPolTableParams) iMABean() {}
