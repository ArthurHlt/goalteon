package beans

import (
	"fmt"
	"reflect"
)

// SlbPortInfoTable The table of port information.
// Note:This mib is not supported for VX instance of Virtualization.
type SlbPortInfoTable struct {
	// Index of slb port.
	SlbPortInfoIndex int32
	Params           *SlbPortInfoTableParams
}

func NewSlbPortInfoTableList() *SlbPortInfoTable {
	return &SlbPortInfoTable{}
}

func NewSlbPortInfoTable(
	slbPortInfoIndex int32,
	params *SlbPortInfoTableParams,
) *SlbPortInfoTable {
	return &SlbPortInfoTable{
		SlbPortInfoIndex: slbPortInfoIndex,
		Params:           params,
	}
}

func (c *SlbPortInfoTable) Name() string {
	return "SlbPortInfoTable"
}

func (c *SlbPortInfoTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbPortInfoTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbPortInfoTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbPortInfoIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbPortInfoIndex)
}

type SlbPortInfoTableClientState int32

const (
	SlbPortInfoTableClientState_Enabled     SlbPortInfoTableClientState = 1
	SlbPortInfoTableClientState_Disabled    SlbPortInfoTableClientState = 2
	SlbPortInfoTableClientState_Unsupported SlbPortInfoTableClientState = 2147483647
)

type SlbPortInfoTableSerState int32

const (
	SlbPortInfoTableSerState_Enabled     SlbPortInfoTableSerState = 1
	SlbPortInfoTableSerState_Disabled    SlbPortInfoTableSerState = 2
	SlbPortInfoTableSerState_Unsupported SlbPortInfoTableSerState = 2147483647
)

type SlbPortInfoTableFltState int32

const (
	SlbPortInfoTableFltState_Enabled     SlbPortInfoTableFltState = 1
	SlbPortInfoTableFltState_Disabled    SlbPortInfoTableFltState = 2
	SlbPortInfoTableFltState_Unsupported SlbPortInfoTableFltState = 2147483647
)

type SlbPortInfoTableRTSState int32

const (
	SlbPortInfoTableRTSState_Enabled     SlbPortInfoTableRTSState = 1
	SlbPortInfoTableRTSState_Disabled    SlbPortInfoTableRTSState = 2
	SlbPortInfoTableRTSState_Unsupported SlbPortInfoTableRTSState = 2147483647
)

type SlbPortInfoTableHotStandbyState int32

const (
	SlbPortInfoTableHotStandbyState_Enabled     SlbPortInfoTableHotStandbyState = 1
	SlbPortInfoTableHotStandbyState_Disabled    SlbPortInfoTableHotStandbyState = 2
	SlbPortInfoTableHotStandbyState_Unsupported SlbPortInfoTableHotStandbyState = 2147483647
)

type SlbPortInfoTableInterSWState int32

const (
	SlbPortInfoTableInterSWState_Enabled     SlbPortInfoTableInterSWState = 1
	SlbPortInfoTableInterSWState_Disabled    SlbPortInfoTableInterSWState = 2
	SlbPortInfoTableInterSWState_Unsupported SlbPortInfoTableInterSWState = 2147483647
)

type SlbPortInfoTableProxyState int32

const (
	SlbPortInfoTableProxyState_Enabled     SlbPortInfoTableProxyState = 1
	SlbPortInfoTableProxyState_Disabled    SlbPortInfoTableProxyState = 2
	SlbPortInfoTableProxyState_Unsupported SlbPortInfoTableProxyState = 2147483647
)

type SlbPortInfoTableIdSlbState int32

const (
	SlbPortInfoTableIdSlbState_Enabled     SlbPortInfoTableIdSlbState = 1
	SlbPortInfoTableIdSlbState_Disabled    SlbPortInfoTableIdSlbState = 2
	SlbPortInfoTableIdSlbState_Unsupported SlbPortInfoTableIdSlbState = 2147483647
)

type SlbPortInfoTableSymantecState int32

const (
	SlbPortInfoTableSymantecState_Enabled     SlbPortInfoTableSymantecState = 1
	SlbPortInfoTableSymantecState_Disabled    SlbPortInfoTableSymantecState = 2
	SlbPortInfoTableSymantecState_Unsupported SlbPortInfoTableSymantecState = 2147483647
)

type SlbPortInfoTableAwMonitorState int32

const (
	SlbPortInfoTableAwMonitorState_Enabled     SlbPortInfoTableAwMonitorState = 1
	SlbPortInfoTableAwMonitorState_Disabled    SlbPortInfoTableAwMonitorState = 2
	SlbPortInfoTableAwMonitorState_Unsupported SlbPortInfoTableAwMonitorState = 2147483647
)

type SlbPortInfoTableParams struct {
	// Index of slb port.
	Index int32 `json:"Index,omitempty"`
	// Client state on the slb port.
	ClientState SlbPortInfoTableClientState `json:"ClientState,omitempty"`
	// Server state on the slb port.
	SerState SlbPortInfoTableSerState `json:"SerState,omitempty"`
	// Filter state on the slb port .
	FltState SlbPortInfoTableFltState `json:"FltState,omitempty"`
	// RTS processing state on the slb port.
	RTSState SlbPortInfoTableRTSState `json:"RTSState,omitempty"`
	// Hot standby state on the slb port.
	HotStandbyState SlbPortInfoTableHotStandbyState `json:"HotStandbyState,omitempty"`
	// Inter-switch processing state on the slb port.
	InterSWState SlbPortInfoTableInterSWState `json:"InterSWState,omitempty"`
	// Proxy state on the slb port.
	ProxyState SlbPortInfoTableProxyState `json:"ProxyState,omitempty"`
	// Intrusion detection server load balancing state on the slb port.
	IdSlbState SlbPortInfoTableIdSlbState `json:"IdSlbState,omitempty"`
	// Symantec Processing state on the slb port.
	SymantecState SlbPortInfoTableSymantecState `json:"SymantecState,omitempty"`
	// The filtering rules applied to the port.  The filtering rules are
	// presented in bitmap format.
	// in receiving order:
	// OCTET 1  OCTET 2  .....
	// xxxxxxxx xxxxxxxx .....
	// ||    || |_ filter 9
	// ||    ||
	// ||    ||___ filter 8
	// ||    |____ filter 7
	// ||      .    .   .
	// ||_________ filter 2
	// |__________ filter 1
	// where x : 1 - The represented filter rule applied to the port
	// 0 - The represented filter rule not applied to the port.
	FitersAdded string `json:"FitersAdded,omitempty"`
	// Enable/disable AppWall Monitoring.
	AwMonitorState SlbPortInfoTableAwMonitorState `json:"AwMonitorState,omitempty"`
}

func (p SlbPortInfoTableParams) iMABean() {}
