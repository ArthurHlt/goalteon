package beans

import (
	"fmt"
	"reflect"
)

// SlbCurCfgPortTable The table of ports and their SLB states,
// Note:This mib is not supported for VX instance of Virtualization.
type SlbCurCfgPortTable struct {
	// The port number for which the SLB information pertains.
	SlbCurCfgPortIndex int32
	Params             *SlbCurCfgPortTableParams
}

func NewSlbCurCfgPortTableList() *SlbCurCfgPortTable {
	return &SlbCurCfgPortTable{}
}

func NewSlbCurCfgPortTable(
	slbCurCfgPortIndex int32,
	params *SlbCurCfgPortTableParams,
) *SlbCurCfgPortTable {
	return &SlbCurCfgPortTable{
		SlbCurCfgPortIndex: slbCurCfgPortIndex,
		Params:             params,
	}
}

func (c *SlbCurCfgPortTable) Name() string {
	return "SlbCurCfgPortTable"
}

func (c *SlbCurCfgPortTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbCurCfgPortTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbCurCfgPortTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbCurCfgPortIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbCurCfgPortIndex)
}

type SlbCurCfgPortTableSlbState int32

const (
	SlbCurCfgPortTableSlbState_None         SlbCurCfgPortTableSlbState = 1
	SlbCurCfgPortTableSlbState_Client       SlbCurCfgPortTableSlbState = 2
	SlbCurCfgPortTableSlbState_Server       SlbCurCfgPortTableSlbState = 3
	SlbCurCfgPortTableSlbState_ClientServer SlbCurCfgPortTableSlbState = 4
	SlbCurCfgPortTableSlbState_Unsupported  SlbCurCfgPortTableSlbState = 2147483647
)

type SlbCurCfgPortTableSlbHotStandby int32

const (
	SlbCurCfgPortTableSlbHotStandby_Enabled     SlbCurCfgPortTableSlbHotStandby = 1
	SlbCurCfgPortTableSlbHotStandby_Disabled    SlbCurCfgPortTableSlbHotStandby = 2
	SlbCurCfgPortTableSlbHotStandby_Unsupported SlbCurCfgPortTableSlbHotStandby = 2147483647
)

type SlbCurCfgPortTableSlbInterSwitch int32

const (
	SlbCurCfgPortTableSlbInterSwitch_Enabled     SlbCurCfgPortTableSlbInterSwitch = 1
	SlbCurCfgPortTableSlbInterSwitch_Disabled    SlbCurCfgPortTableSlbInterSwitch = 2
	SlbCurCfgPortTableSlbInterSwitch_Unsupported SlbCurCfgPortTableSlbInterSwitch = 2147483647
)

type SlbCurCfgPortTableSlbPipState int32

const (
	SlbCurCfgPortTableSlbPipState_Enabled     SlbCurCfgPortTableSlbPipState = 1
	SlbCurCfgPortTableSlbPipState_Disabled    SlbCurCfgPortTableSlbPipState = 2
	SlbCurCfgPortTableSlbPipState_Unsupported SlbCurCfgPortTableSlbPipState = 2147483647
)

type SlbCurCfgPortTableSlbRtsState int32

const (
	SlbCurCfgPortTableSlbRtsState_Enabled     SlbCurCfgPortTableSlbRtsState = 1
	SlbCurCfgPortTableSlbRtsState_Disabled    SlbCurCfgPortTableSlbRtsState = 2
	SlbCurCfgPortTableSlbRtsState_Unsupported SlbCurCfgPortTableSlbRtsState = 2147483647
)

type SlbCurCfgPortTableSlbIdslbState int32

const (
	SlbCurCfgPortTableSlbIdslbState_Enabled     SlbCurCfgPortTableSlbIdslbState = 1
	SlbCurCfgPortTableSlbIdslbState_Disabled    SlbCurCfgPortTableSlbIdslbState = 2
	SlbCurCfgPortTableSlbIdslbState_Unsupported SlbCurCfgPortTableSlbIdslbState = 2147483647
)

type SlbCurCfgPortTableSlbFilter int32

const (
	SlbCurCfgPortTableSlbFilter_Enabled     SlbCurCfgPortTableSlbFilter = 1
	SlbCurCfgPortTableSlbFilter_Disabled    SlbCurCfgPortTableSlbFilter = 2
	SlbCurCfgPortTableSlbFilter_Unsupported SlbCurCfgPortTableSlbFilter = 2147483647
)

type SlbCurCfgPortTableSlbServState int32

const (
	SlbCurCfgPortTableSlbServState_Enabled     SlbCurCfgPortTableSlbServState = 1
	SlbCurCfgPortTableSlbServState_Disabled    SlbCurCfgPortTableSlbServState = 2
	SlbCurCfgPortTableSlbServState_Unsupported SlbCurCfgPortTableSlbServState = 2147483647
)

type SlbCurCfgPortTableSlbClntState int32

const (
	SlbCurCfgPortTableSlbClntState_Enabled     SlbCurCfgPortTableSlbClntState = 1
	SlbCurCfgPortTableSlbClntState_Disabled    SlbCurCfgPortTableSlbClntState = 2
	SlbCurCfgPortTableSlbClntState_Unsupported SlbCurCfgPortTableSlbClntState = 2147483647
)

type SlbCurCfgPortTableSlbL3Filter int32

const (
	SlbCurCfgPortTableSlbL3Filter_Enabled     SlbCurCfgPortTableSlbL3Filter = 1
	SlbCurCfgPortTableSlbL3Filter_Disabled    SlbCurCfgPortTableSlbL3Filter = 2
	SlbCurCfgPortTableSlbL3Filter_Unsupported SlbCurCfgPortTableSlbL3Filter = 2147483647
)

type SlbCurCfgPortTableSlbAwMonitorState int32

const (
	SlbCurCfgPortTableSlbAwMonitorState_Enabled     SlbCurCfgPortTableSlbAwMonitorState = 1
	SlbCurCfgPortTableSlbAwMonitorState_Disabled    SlbCurCfgPortTableSlbAwMonitorState = 2
	SlbCurCfgPortTableSlbAwMonitorState_Unsupported SlbCurCfgPortTableSlbAwMonitorState = 2147483647
)

type SlbCurCfgPortTableParams struct {
	// The port number for which the SLB information pertains.
	Index int32 `json:"Index,omitempty"`
	// The SLB state of the port.
	// 	 none(1)            - not SLB port
	// 	 client(2)          - SLB client port
	// 	 server(3)          - SLB server port
	// 	 client-server(4)   - SLB client and server port
	SlbState SlbCurCfgPortTableSlbState `json:"SlbState,omitempty"`
	// Enable or disable hot standby processing on the switch port.
	SlbHotStandby SlbCurCfgPortTableSlbHotStandby `json:"SlbHotStandby,omitempty"`
	// Enable or disable inter-switch processing on the switch port.
	SlbInterSwitch SlbCurCfgPortTableSlbInterSwitch `json:"SlbInterSwitch,omitempty"`
	// Enable or disable use of proxy IP address on the switch port.
	SlbPipState SlbCurCfgPortTableSlbPipState `json:"SlbPipState,omitempty"`
	// Enable or disable RTS processing on the switch port.
	SlbRtsState SlbCurCfgPortTableSlbRtsState `json:"SlbRtsState,omitempty"`
	// Enable or disable Intrusion Detection server load balancing.
	SlbIdslbState SlbCurCfgPortTableSlbIdslbState `json:"SlbIdslbState,omitempty"`
	// Enable or disable Filtering.
	SlbFilter SlbCurCfgPortTableSlbFilter `json:"SlbFilter,omitempty"`
	// Enable or disable server processing.
	SlbServState SlbCurCfgPortTableSlbServState `json:"SlbServState,omitempty"`
	// Enable or disable client processing.
	SlbClntState SlbCurCfgPortTableSlbClntState `json:"SlbClntState,omitempty"`
	// Enable or disable Layer3 Filtering.
	SlbL3Filter SlbCurCfgPortTableSlbL3Filter `json:"SlbL3Filter,omitempty"`
	// The filtering rules applied to the port.  The filtering rules are
	// presented in bitmap format.
	// 	 in receiving order:
	// 	     OCTET 1  OCTET 2  .....
	// xxxxxxxx xxxxxxxx .....
	// ||    || |_ filter 9
	// ||    ||
	// ||    ||___ filter 8
	// ||    |____ filter 7
	// ||      .    .   .
	// ||_________ filter 2
	// |__________ filter 1 (as index to fltCurCfgTable)
	// where x : 1 - The represented filter rule applied to the port
	// 		   0 - The represented filter rule not applied to the port
	SlbFilterBmap string `json:"SlbFilterBmap,omitempty"`
	// VLAN for inter-switch processing.
	InterSwitchVlan uint32 `json:"InterSwitchVlan,omitempty"`
	// VLANs associated with this port.
	// The VLANs are presented in a bitmap format.
	// OCTET 1  OCTET 2  .....
	// xxxxxxxx xxxxxxxx .....
	// ||    || |_ VLAN 9
	// ||    ||
	// ||    ||___ VLAN 8
	// ||    |____ VLAN 7
	// ||      .    .   .
	// ||_________ VLAN 2
	// |__________ VLAN 1
	// where x : 1 - VLAN is associated
	// 0 - VLAN is not associated
	VlanBmap string `json:"VlanBmap,omitempty"`
	// Enable/disable AppWall Monitoring.
	SlbAwMonitorState SlbCurCfgPortTableSlbAwMonitorState `json:"SlbAwMonitorState,omitempty"`
}

func (p SlbCurCfgPortTableParams) iMABean() {}
