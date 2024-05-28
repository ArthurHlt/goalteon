package beans

import (
	"fmt"
	"reflect"
)

// SlbEnhVirtServicesInfoTable The table of virtual server services run-time information.
// Note:This mib is not supported for VX instance of Virtualization.
type SlbEnhVirtServicesInfoTable struct {
	// The virtual server index.
	SlbEnhVirtServicesInfoVirtServIndex string
	// The virtual service index.
	SlbEnhVirtServicesInfoSvcIndex int32
	// The real server index.
	SlbEnhVirtServicesInfoRealServIndex string
	Params                              *SlbEnhVirtServicesInfoTableParams
}

func NewSlbEnhVirtServicesInfoTableList() *SlbEnhVirtServicesInfoTable {
	return &SlbEnhVirtServicesInfoTable{}
}

func NewSlbEnhVirtServicesInfoTable(
	slbEnhVirtServicesInfoVirtServIndex string,
	slbEnhVirtServicesInfoSvcIndex int32,
	slbEnhVirtServicesInfoRealServIndex string,
	params *SlbEnhVirtServicesInfoTableParams,
) *SlbEnhVirtServicesInfoTable {
	return &SlbEnhVirtServicesInfoTable{
		SlbEnhVirtServicesInfoVirtServIndex: slbEnhVirtServicesInfoVirtServIndex,
		SlbEnhVirtServicesInfoSvcIndex:      slbEnhVirtServicesInfoSvcIndex,
		SlbEnhVirtServicesInfoRealServIndex: slbEnhVirtServicesInfoRealServIndex,
		Params:                              params,
	}
}

func (c *SlbEnhVirtServicesInfoTable) Name() string {
	return "SlbEnhVirtServicesInfoTable"
}

func (c *SlbEnhVirtServicesInfoTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbEnhVirtServicesInfoTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbEnhVirtServicesInfoTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbEnhVirtServicesInfoVirtServIndex).IsZero() &&
		reflect.ValueOf(c.SlbEnhVirtServicesInfoSvcIndex).IsZero() &&
		reflect.ValueOf(c.SlbEnhVirtServicesInfoRealServIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbEnhVirtServicesInfoVirtServIndex) + "/" + fmt.Sprint(c.SlbEnhVirtServicesInfoSvcIndex) + "/" + fmt.Sprint(c.SlbEnhVirtServicesInfoRealServIndex)
}

type SlbEnhVirtServicesInfoTableState int32

const (
	SlbEnhVirtServicesInfoTableState_Blocked     SlbEnhVirtServicesInfoTableState = 1
	SlbEnhVirtServicesInfoTableState_Running     SlbEnhVirtServicesInfoTableState = 2
	SlbEnhVirtServicesInfoTableState_Failed      SlbEnhVirtServicesInfoTableState = 3
	SlbEnhVirtServicesInfoTableState_Disabled    SlbEnhVirtServicesInfoTableState = 4
	SlbEnhVirtServicesInfoTableState_Slowstart   SlbEnhVirtServicesInfoTableState = 5
	SlbEnhVirtServicesInfoTableState_Overflow    SlbEnhVirtServicesInfoTableState = 6
	SlbEnhVirtServicesInfoTableState_Noinstance  SlbEnhVirtServicesInfoTableState = 7
	SlbEnhVirtServicesInfoTableState_Unsupported SlbEnhVirtServicesInfoTableState = 2147483647
)

type SlbEnhVirtServicesInfoTableParams struct {
	// The virtual server index.
	VirtServIndex string `json:"VirtServIndex,omitempty"`
	// The virtual service index.
	SvcIndex int32 `json:"SvcIndex,omitempty"`
	// The real server index.
	RealServIndex string `json:"RealServIndex,omitempty"`
	// The layer 4 virtual port number for the service.
	Vport int32 `json:"Vport,omitempty"`
	// The layer 4 real port number for the service.
	Rport int32 `json:"Rport,omitempty"`
	// The state of the real server for this virtual service.
	State SlbEnhVirtServicesInfoTableState `json:"State,omitempty"`
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

func (p SlbEnhVirtServicesInfoTableParams) iMABean() {}
