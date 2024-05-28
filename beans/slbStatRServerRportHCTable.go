package beans

import (
	"fmt"
	"reflect"
)

// SlbStatRServerRportHCTable The table of real server rport run-time health check statistics.
// Note:This mib is not supported for VX instance of Virtualization.
type SlbStatRServerRportHCTable struct {
	// The real server index
	SlbStatRServerRportHCRealIndex int32
	// The real server service index
	SlbStatRServerRportHCServIndex int32
	Params                         *SlbStatRServerRportHCTableParams
}

func NewSlbStatRServerRportHCTableList() *SlbStatRServerRportHCTable {
	return &SlbStatRServerRportHCTable{}
}

func NewSlbStatRServerRportHCTable(
	slbStatRServerRportHCRealIndex int32,
	slbStatRServerRportHCServIndex int32,
	params *SlbStatRServerRportHCTableParams,
) *SlbStatRServerRportHCTable {
	return &SlbStatRServerRportHCTable{
		SlbStatRServerRportHCRealIndex: slbStatRServerRportHCRealIndex,
		SlbStatRServerRportHCServIndex: slbStatRServerRportHCServIndex,
		Params:                         params,
	}
}

func (c *SlbStatRServerRportHCTable) Name() string {
	return "SlbStatRServerRportHCTable"
}

func (c *SlbStatRServerRportHCTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbStatRServerRportHCTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbStatRServerRportHCTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbStatRServerRportHCRealIndex).IsZero() &&
		reflect.ValueOf(c.SlbStatRServerRportHCServIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbStatRServerRportHCRealIndex) + "/" + fmt.Sprint(c.SlbStatRServerRportHCServIndex)
}

type SlbStatRServerRportHCTableHCStatus int32

const (
	SlbStatRServerRportHCTableHCStatus_Up          SlbStatRServerRportHCTableHCStatus = 1
	SlbStatRServerRportHCTableHCStatus_Failed      SlbStatRServerRportHCTableHCStatus = 2
	SlbStatRServerRportHCTableHCStatus_Blocked     SlbStatRServerRportHCTableHCStatus = 3
	SlbStatRServerRportHCTableHCStatus_Overflow    SlbStatRServerRportHCTableHCStatus = 4
	SlbStatRServerRportHCTableHCStatus_Disabled    SlbStatRServerRportHCTableHCStatus = 5
	SlbStatRServerRportHCTableHCStatus_Unsupported SlbStatRServerRportHCTableHCStatus = 2147483647
)

type SlbStatRServerRportHCTableParams struct {
	// The real server index
	HCRealIndex int32 `json:"HCRealIndex,omitempty"`
	// The real server service index
	HCServIndex int32 `json:"HCServIndex,omitempty"`
	// The real server runtime instance health check ID
	HCId string `json:"HCId,omitempty"`
	// The real server runtime instance health check type
	HCType string `json:"HCType,omitempty"`
	// The real server runtime instance health check status
	HCStatus SlbStatRServerRportHCTableHCStatus `json:"HCStatus,omitempty"`
	// The total up time of the real server health check instance.
	HCUptime string `json:"HCUptime,omitempty"`
	// The total down time of the real server health check instance.
	HCDowntime string `json:"HCDowntime,omitempty"`
	// The timestamp when the last valid response of the real server
	// health check instance was received.
	HCLastValidTime string `json:"HCLastValidTime,omitempty"`
	// The average response time of the health check for the real server instance.
	HCAvgRespTime int32 `json:"HCAvgRespTime,omitempty"`
	// The peak response time of the health check for the real server instance.
	HCPeakRespTime int32 `json:"HCPeakRespTime,omitempty"`
	// The last valid response time of the health check for the real server instance.
	HCLastValidRespTime int32 `json:"HCLastValidRespTime,omitempty"`
	// The total number of times the real server health check instance state changed to FAILED.
	HCFailureCount int32 `json:"HCFailureCount,omitempty"`
}

func (p SlbStatRServerRportHCTableParams) iMABean() {}
