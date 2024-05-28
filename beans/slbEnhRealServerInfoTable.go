package beans

import (
	"fmt"
	"reflect"
)

// SlbEnhRealServerInfoTable The table of real server run-time information.
// Note:This mib is not supported for VX instance of Virtualization.
type SlbEnhRealServerInfoTable struct {
	// The real server index
	SlbEnhRealServerInfoIndex string
	Params                    *SlbEnhRealServerInfoTableParams
}

func NewSlbEnhRealServerInfoTableList() *SlbEnhRealServerInfoTable {
	return &SlbEnhRealServerInfoTable{}
}

func NewSlbEnhRealServerInfoTable(
	slbEnhRealServerInfoIndex string,
	params *SlbEnhRealServerInfoTableParams,
) *SlbEnhRealServerInfoTable {
	return &SlbEnhRealServerInfoTable{
		SlbEnhRealServerInfoIndex: slbEnhRealServerInfoIndex,
		Params:                    params,
	}
}

func (c *SlbEnhRealServerInfoTable) Name() string {
	return "SlbEnhRealServerInfoTable"
}

func (c *SlbEnhRealServerInfoTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbEnhRealServerInfoTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbEnhRealServerInfoTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbEnhRealServerInfoIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbEnhRealServerInfoIndex)
}

type SlbEnhRealServerInfoTableHealthLayer int32

const (
	SlbEnhRealServerInfoTableHealthLayer_Layer1      SlbEnhRealServerInfoTableHealthLayer = 1
	SlbEnhRealServerInfoTableHealthLayer_Layer3      SlbEnhRealServerInfoTableHealthLayer = 3
	SlbEnhRealServerInfoTableHealthLayer_Layer4      SlbEnhRealServerInfoTableHealthLayer = 4
	SlbEnhRealServerInfoTableHealthLayer_Unsupported SlbEnhRealServerInfoTableHealthLayer = 2147483647
)

type SlbEnhRealServerInfoTableOverflow int32

const (
	SlbEnhRealServerInfoTableOverflow_Overflow    SlbEnhRealServerInfoTableOverflow = 1
	SlbEnhRealServerInfoTableOverflow_NoOverflow  SlbEnhRealServerInfoTableOverflow = 2
	SlbEnhRealServerInfoTableOverflow_Unsupported SlbEnhRealServerInfoTableOverflow = 2147483647
)

type SlbEnhRealServerInfoTableState int32

const (
	SlbEnhRealServerInfoTableState_Running     SlbEnhRealServerInfoTableState = 2
	SlbEnhRealServerInfoTableState_Failed      SlbEnhRealServerInfoTableState = 3
	SlbEnhRealServerInfoTableState_Disabled    SlbEnhRealServerInfoTableState = 4
	SlbEnhRealServerInfoTableState_Unsupported SlbEnhRealServerInfoTableState = 2147483647
)

type SlbEnhRealServerInfoTableParams struct {
	// The real server index
	Index string `json:"Index,omitempty"`
	// IP address of the real server.
	IpAddr string `json:"IpAddr,omitempty"`
	// The MAC address of the real server.
	MacAddr string `json:"MacAddr,omitempty"`
	// The switch port that the real server is connected to.
	SwitchPort int32 `json:"SwitchPort,omitempty"`
	// The OSI layer at which the real server functionality is verified.
	HealthLayer SlbEnhRealServerInfoTableHealthLayer `json:"HealthLayer,omitempty"`
	// The overflow state of the real server.
	Overflow SlbEnhRealServerInfoTableOverflow `json:"Overflow,omitempty"`
	// The state of the real server.
	State SlbEnhRealServerInfoTableState `json:"State,omitempty"`
	// The VLAN that the real server belongs to.
	Vlan int32 `json:"Vlan,omitempty"`
	// The health check id configured for the real server.
	Health string `json:"Health,omitempty"`
	// The total up time of the real server.
	UpTime string `json:"UpTime,omitempty"`
	// The total down time of the real server.
	DownTime string `json:"DownTime,omitempty"`
	// The time-stamp when the last failure of the real server occurred.
	LastFailureTime string `json:"LastFailureTime,omitempty"`
}

func (p SlbEnhRealServerInfoTableParams) iMABean() {}
