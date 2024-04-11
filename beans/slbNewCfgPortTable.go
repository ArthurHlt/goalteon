package beans

import (
	"fmt"
	"reflect"
)

// SlbNewCfgPortTable The table of ports and their SLB states,
// Note:This mib is not supported for VX instance of Virtualization.
type SlbNewCfgPortTable struct {
	// The port number for which the SLB information pertains.
	SlbNewCfgPortIndex int32
	Params             *SlbNewCfgPortTableParams
}

func NewSlbNewCfgPortTableList() *SlbNewCfgPortTable {
	return &SlbNewCfgPortTable{}
}

func NewSlbNewCfgPortTable(
	slbNewCfgPortIndex int32,
	params *SlbNewCfgPortTableParams,
) *SlbNewCfgPortTable {
	return &SlbNewCfgPortTable{
		SlbNewCfgPortIndex: slbNewCfgPortIndex,
		Params:             params,
	}
}

func (c *SlbNewCfgPortTable) Name() string {
	return "SlbNewCfgPortTable"
}

func (c *SlbNewCfgPortTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbNewCfgPortTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbNewCfgPortTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbNewCfgPortIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbNewCfgPortIndex)
}

type SlbNewCfgPortTableSlbState int32

const (
	SlbNewCfgPortTableSlbState_None         SlbNewCfgPortTableSlbState = 1
	SlbNewCfgPortTableSlbState_Client       SlbNewCfgPortTableSlbState = 2
	SlbNewCfgPortTableSlbState_Server       SlbNewCfgPortTableSlbState = 3
	SlbNewCfgPortTableSlbState_ClientServer SlbNewCfgPortTableSlbState = 4
	SlbNewCfgPortTableSlbState_Unsupported  SlbNewCfgPortTableSlbState = 2147483647
)

type SlbNewCfgPortTableSlbHotStandby int32

const (
	SlbNewCfgPortTableSlbHotStandby_Enabled     SlbNewCfgPortTableSlbHotStandby = 1
	SlbNewCfgPortTableSlbHotStandby_Disabled    SlbNewCfgPortTableSlbHotStandby = 2
	SlbNewCfgPortTableSlbHotStandby_Unsupported SlbNewCfgPortTableSlbHotStandby = 2147483647
)

type SlbNewCfgPortTableSlbInterSwitch int32

const (
	SlbNewCfgPortTableSlbInterSwitch_Enabled     SlbNewCfgPortTableSlbInterSwitch = 1
	SlbNewCfgPortTableSlbInterSwitch_Disabled    SlbNewCfgPortTableSlbInterSwitch = 2
	SlbNewCfgPortTableSlbInterSwitch_Unsupported SlbNewCfgPortTableSlbInterSwitch = 2147483647
)

type SlbNewCfgPortTableSlbPipState int32

const (
	SlbNewCfgPortTableSlbPipState_Enabled     SlbNewCfgPortTableSlbPipState = 1
	SlbNewCfgPortTableSlbPipState_Disabled    SlbNewCfgPortTableSlbPipState = 2
	SlbNewCfgPortTableSlbPipState_Unsupported SlbNewCfgPortTableSlbPipState = 2147483647
)

type SlbNewCfgPortTableSlbRtsState int32

const (
	SlbNewCfgPortTableSlbRtsState_Enabled     SlbNewCfgPortTableSlbRtsState = 1
	SlbNewCfgPortTableSlbRtsState_Disabled    SlbNewCfgPortTableSlbRtsState = 2
	SlbNewCfgPortTableSlbRtsState_Unsupported SlbNewCfgPortTableSlbRtsState = 2147483647
)

type SlbNewCfgPortTableDelete int32

const (
	SlbNewCfgPortTableDelete_Other       SlbNewCfgPortTableDelete = 1
	SlbNewCfgPortTableDelete_Delete      SlbNewCfgPortTableDelete = 2
	SlbNewCfgPortTableDelete_Unsupported SlbNewCfgPortTableDelete = 2147483647
)

type SlbNewCfgPortTableSlbIdslbState int32

const (
	SlbNewCfgPortTableSlbIdslbState_Enabled     SlbNewCfgPortTableSlbIdslbState = 1
	SlbNewCfgPortTableSlbIdslbState_Disabled    SlbNewCfgPortTableSlbIdslbState = 2
	SlbNewCfgPortTableSlbIdslbState_Unsupported SlbNewCfgPortTableSlbIdslbState = 2147483647
)

type SlbNewCfgPortTableSlbFilter int32

const (
	SlbNewCfgPortTableSlbFilter_Enabled     SlbNewCfgPortTableSlbFilter = 1
	SlbNewCfgPortTableSlbFilter_Disabled    SlbNewCfgPortTableSlbFilter = 2
	SlbNewCfgPortTableSlbFilter_Unsupported SlbNewCfgPortTableSlbFilter = 2147483647
)

type SlbNewCfgPortTableSlbServState int32

const (
	SlbNewCfgPortTableSlbServState_Enabled     SlbNewCfgPortTableSlbServState = 1
	SlbNewCfgPortTableSlbServState_Disabled    SlbNewCfgPortTableSlbServState = 2
	SlbNewCfgPortTableSlbServState_Unsupported SlbNewCfgPortTableSlbServState = 2147483647
)

type SlbNewCfgPortTableSlbClntState int32

const (
	SlbNewCfgPortTableSlbClntState_Enabled     SlbNewCfgPortTableSlbClntState = 1
	SlbNewCfgPortTableSlbClntState_Disabled    SlbNewCfgPortTableSlbClntState = 2
	SlbNewCfgPortTableSlbClntState_Unsupported SlbNewCfgPortTableSlbClntState = 2147483647
)

type SlbNewCfgPortTableSlbL3Filter int32

const (
	SlbNewCfgPortTableSlbL3Filter_Enabled     SlbNewCfgPortTableSlbL3Filter = 1
	SlbNewCfgPortTableSlbL3Filter_Disabled    SlbNewCfgPortTableSlbL3Filter = 2
	SlbNewCfgPortTableSlbL3Filter_Unsupported SlbNewCfgPortTableSlbL3Filter = 2147483647
)

type SlbNewCfgPortTableSlbAwMonitorState int32

const (
	SlbNewCfgPortTableSlbAwMonitorState_Enabled     SlbNewCfgPortTableSlbAwMonitorState = 1
	SlbNewCfgPortTableSlbAwMonitorState_Disabled    SlbNewCfgPortTableSlbAwMonitorState = 2
	SlbNewCfgPortTableSlbAwMonitorState_Unsupported SlbNewCfgPortTableSlbAwMonitorState = 2147483647
)

type SlbNewCfgPortTableParams struct {
	// The port number for which the SLB information pertains.
	Index int32 `json:"Index,omitempty"`
	// The SLB state of the port.
	// 	 none(1)            - not SLB port
	// 	 client(2)          - SLB client port
	// 	 server(3)          - SLB server port
	// 	 client-server(4)   - SLB client and server port
	SlbState SlbNewCfgPortTableSlbState `json:"SlbState,omitempty"`
	// Enable or disable hot standby processing on the switch port.
	SlbHotStandby SlbNewCfgPortTableSlbHotStandby `json:"SlbHotStandby,omitempty"`
	// Enable or disable inter-switch processing on the switch port.
	SlbInterSwitch SlbNewCfgPortTableSlbInterSwitch `json:"SlbInterSwitch,omitempty"`
	// Enable or disable use of proxy IP address on the switch port.
	SlbPipState SlbNewCfgPortTableSlbPipState `json:"SlbPipState,omitempty"`
	// Enable or disable RTS processing on the switch port.
	SlbRtsState SlbNewCfgPortTableSlbRtsState `json:"SlbRtsState,omitempty"`
	// When set to the value of 2 (delete), the entire row is deleted.
	// When read, other(1) is returned. Setting the value to anything
	// other than 2(delete) has no effect on the state of the row.
	Delete SlbNewCfgPortTableDelete `json:"Delete,omitempty"`
	// Enable or disable Intrusion Detection server load balancing.
	SlbIdslbState SlbNewCfgPortTableSlbIdslbState `json:"SlbIdslbState,omitempty"`
	// Enable or disable Filtering.
	SlbFilter SlbNewCfgPortTableSlbFilter `json:"SlbFilter,omitempty"`
	// Specify the Filter Number to be added to this Port table.
	SlbAddFilter uint32 `json:"SlbAddFilter,omitempty"`
	// Specify the Filter Number to be deleted from this Port table.
	SlbRemFilter uint32 `json:"SlbRemFilter,omitempty"`
	// Enable or disable Server Processing.
	SlbServState SlbNewCfgPortTableSlbServState `json:"SlbServState,omitempty"`
	// Enable or disable Client Processing.
	SlbClntState SlbNewCfgPortTableSlbClntState `json:"SlbClntState,omitempty"`
	// Enable or disable Layer3 Filtering.
	SlbL3Filter SlbNewCfgPortTableSlbL3Filter `json:"SlbL3Filter,omitempty"`
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
	// Enable/disable AppWall Monitoring
	SlbAwMonitorState SlbNewCfgPortTableSlbAwMonitorState `json:"SlbAwMonitorState,omitempty"`
}

func (p SlbNewCfgPortTableParams) iMABean() {}
