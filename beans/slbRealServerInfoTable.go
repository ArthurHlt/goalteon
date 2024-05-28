package beans

import (
	"fmt"
	"reflect"
)

// SlbRealServerInfoTable The table of real server run-time information.
// Note:This mib is not supported for VX instance of Virtualization.
type SlbRealServerInfoTable struct {
	// The real server index
	SlbRealServerInfoIndex int32
	Params                 *SlbRealServerInfoTableParams
}

func NewSlbRealServerInfoTableList() *SlbRealServerInfoTable {
	return &SlbRealServerInfoTable{}
}

func NewSlbRealServerInfoTable(
	slbRealServerInfoIndex int32,
	params *SlbRealServerInfoTableParams,
) *SlbRealServerInfoTable {
	return &SlbRealServerInfoTable{
		SlbRealServerInfoIndex: slbRealServerInfoIndex,
		Params:                 params,
	}
}

func (c *SlbRealServerInfoTable) Name() string {
	return "SlbRealServerInfoTable"
}

func (c *SlbRealServerInfoTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbRealServerInfoTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbRealServerInfoTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbRealServerInfoIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbRealServerInfoIndex)
}

type SlbRealServerInfoTableHealthLayer int32

const (
	SlbRealServerInfoTableHealthLayer_Layer1      SlbRealServerInfoTableHealthLayer = 1
	SlbRealServerInfoTableHealthLayer_Layer3      SlbRealServerInfoTableHealthLayer = 3
	SlbRealServerInfoTableHealthLayer_Layer4      SlbRealServerInfoTableHealthLayer = 4
	SlbRealServerInfoTableHealthLayer_Unsupported SlbRealServerInfoTableHealthLayer = 2147483647
)

type SlbRealServerInfoTableOverflow int32

const (
	SlbRealServerInfoTableOverflow_Overflow    SlbRealServerInfoTableOverflow = 1
	SlbRealServerInfoTableOverflow_NoOverflow  SlbRealServerInfoTableOverflow = 2
	SlbRealServerInfoTableOverflow_Unsupported SlbRealServerInfoTableOverflow = 2147483647
)

type SlbRealServerInfoTableState int32

const (
	SlbRealServerInfoTableState_Running     SlbRealServerInfoTableState = 2
	SlbRealServerInfoTableState_Failed      SlbRealServerInfoTableState = 3
	SlbRealServerInfoTableState_Disabled    SlbRealServerInfoTableState = 4
	SlbRealServerInfoTableState_Unsupported SlbRealServerInfoTableState = 2147483647
)

type SlbRealServerInfoTableParams struct {
	// The real server index
	Index int32 `json:"Index,omitempty"`
	// IP address of the real server.
	IpAddr string `json:"IpAddr,omitempty"`
	// The MAC address of the real server.
	MacAddr string `json:"MacAddr,omitempty"`
	// The switch port that the real server is connected to.
	SwitchPort int32 `json:"SwitchPort,omitempty"`
	// The OSI layer at which the real server functionality is verified.
	HealthLayer SlbRealServerInfoTableHealthLayer `json:"HealthLayer,omitempty"`
	// The overflow state of the real server.
	Overflow SlbRealServerInfoTableOverflow `json:"Overflow,omitempty"`
	// The state of the real server.
	State SlbRealServerInfoTableState `json:"State,omitempty"`
	// The VLAN that the real server belongs to.
	Vlan int32 `json:"Vlan,omitempty"`
	// The health check id configured for the real server.
	Health string `json:"Health,omitempty"`
}

func (p SlbRealServerInfoTableParams) iMABean() {}
