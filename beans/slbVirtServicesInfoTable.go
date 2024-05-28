package beans

import (
	"fmt"
	"reflect"
)

// SlbVirtServicesInfoTable The table of virtual server services run-time information.
// Note:This mib is not supported for VX instance of Virtualization.
type SlbVirtServicesInfoTable struct {
	// The virtual server index.
	SlbVirtServicesInfoVirtServIndex int32
	// The virtual service index.
	SlbVirtServicesInfoSvcIndex int32
	// The real server index.
	SlbVirtServicesInfoRealServIndex int32
	Params                           *SlbVirtServicesInfoTableParams
}

func NewSlbVirtServicesInfoTableList() *SlbVirtServicesInfoTable {
	return &SlbVirtServicesInfoTable{}
}

func NewSlbVirtServicesInfoTable(
	slbVirtServicesInfoVirtServIndex int32,
	slbVirtServicesInfoSvcIndex int32,
	slbVirtServicesInfoRealServIndex int32,
	params *SlbVirtServicesInfoTableParams,
) *SlbVirtServicesInfoTable {
	return &SlbVirtServicesInfoTable{
		SlbVirtServicesInfoVirtServIndex: slbVirtServicesInfoVirtServIndex,
		SlbVirtServicesInfoSvcIndex:      slbVirtServicesInfoSvcIndex,
		SlbVirtServicesInfoRealServIndex: slbVirtServicesInfoRealServIndex,
		Params:                           params,
	}
}

func (c *SlbVirtServicesInfoTable) Name() string {
	return "SlbVirtServicesInfoTable"
}

func (c *SlbVirtServicesInfoTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbVirtServicesInfoTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbVirtServicesInfoTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbVirtServicesInfoVirtServIndex).IsZero() &&
		reflect.ValueOf(c.SlbVirtServicesInfoSvcIndex).IsZero() &&
		reflect.ValueOf(c.SlbVirtServicesInfoRealServIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbVirtServicesInfoVirtServIndex) + "/" + fmt.Sprint(c.SlbVirtServicesInfoSvcIndex) + "/" + fmt.Sprint(c.SlbVirtServicesInfoRealServIndex)
}

type SlbVirtServicesInfoTableState int32

const (
	SlbVirtServicesInfoTableState_Blocked     SlbVirtServicesInfoTableState = 1
	SlbVirtServicesInfoTableState_Running     SlbVirtServicesInfoTableState = 2
	SlbVirtServicesInfoTableState_Failed      SlbVirtServicesInfoTableState = 3
	SlbVirtServicesInfoTableState_Disabled    SlbVirtServicesInfoTableState = 4
	SlbVirtServicesInfoTableState_Slowstart   SlbVirtServicesInfoTableState = 5
	SlbVirtServicesInfoTableState_Overflow    SlbVirtServicesInfoTableState = 6
	SlbVirtServicesInfoTableState_Noinstance  SlbVirtServicesInfoTableState = 7
	SlbVirtServicesInfoTableState_Unsupported SlbVirtServicesInfoTableState = 2147483647
)

type SlbVirtServicesInfoTableParams struct {
	// The virtual server index.
	VirtServIndex int32 `json:"VirtServIndex,omitempty"`
	// The virtual service index.
	SvcIndex int32 `json:"SvcIndex,omitempty"`
	// The real server index.
	RealServIndex int32 `json:"RealServIndex,omitempty"`
	// The layer 4 virtual port number for the service.
	Vport int32 `json:"Vport,omitempty"`
	// The layer 4 real port number for the service.
	Rport int32 `json:"Rport,omitempty"`
	// The state of the real server for this virtual service.
	State SlbVirtServicesInfoTableState `json:"State,omitempty"`
	// The health check response time for the real server for this
	// virtual service.
	ResponseTime int32 `json:"ResponseTime,omitempty"`
	// The real server weight.
	Weight int32 `json:"Weight,omitempty"`
	// The configured health check for this real server.
	CfgRealHealth string `json:"CfgRealHealth,omitempty"`
	// The runtime health check for ths real server.
	RtRealHealth string `json:"RtRealHealth,omitempty"`
	// The reason for health check fail.
	StateFailReason string `json:"StateFailReason,omitempty"`
	// The logical expression for this real server.
	RealLogexp string `json:"RealLogexp,omitempty"`
}

func (p SlbVirtServicesInfoTableParams) iMABean() {}
