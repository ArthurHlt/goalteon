package beans

import (
	"fmt"
	"reflect"
)

// SlbStatEnhRServerRportHCTable The table of real server rport run-time health check statistics.
// Note:This mib is not supported for VX instance of Virtualization.
type SlbStatEnhRServerRportHCTable struct {
	// The real server index
	SlbStatEnhRServerRportHCRealIndex string
	// The real server service index
	SlbStatEnhRServerRportHCServIndex int32
	Params                            *SlbStatEnhRServerRportHCTableParams
}

func NewSlbStatEnhRServerRportHCTableList() *SlbStatEnhRServerRportHCTable {
	return &SlbStatEnhRServerRportHCTable{}
}

func NewSlbStatEnhRServerRportHCTable(
	slbStatEnhRServerRportHCRealIndex string,
	slbStatEnhRServerRportHCServIndex int32,
	params *SlbStatEnhRServerRportHCTableParams,
) *SlbStatEnhRServerRportHCTable {
	return &SlbStatEnhRServerRportHCTable{
		SlbStatEnhRServerRportHCRealIndex: slbStatEnhRServerRportHCRealIndex,
		SlbStatEnhRServerRportHCServIndex: slbStatEnhRServerRportHCServIndex,
		Params:                            params,
	}
}

func (c *SlbStatEnhRServerRportHCTable) Name() string {
	return "SlbStatEnhRServerRportHCTable"
}

func (c *SlbStatEnhRServerRportHCTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbStatEnhRServerRportHCTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbStatEnhRServerRportHCTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbStatEnhRServerRportHCRealIndex).IsZero() &&
		reflect.ValueOf(c.SlbStatEnhRServerRportHCServIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbStatEnhRServerRportHCRealIndex) + "/" + fmt.Sprint(c.SlbStatEnhRServerRportHCServIndex)
}

type SlbStatEnhRServerRportHCTableHCStatus int32

const (
	SlbStatEnhRServerRportHCTableHCStatus_Up          SlbStatEnhRServerRportHCTableHCStatus = 1
	SlbStatEnhRServerRportHCTableHCStatus_Failed      SlbStatEnhRServerRportHCTableHCStatus = 2
	SlbStatEnhRServerRportHCTableHCStatus_Blocked     SlbStatEnhRServerRportHCTableHCStatus = 3
	SlbStatEnhRServerRportHCTableHCStatus_Overflow    SlbStatEnhRServerRportHCTableHCStatus = 4
	SlbStatEnhRServerRportHCTableHCStatus_Disabled    SlbStatEnhRServerRportHCTableHCStatus = 5
	SlbStatEnhRServerRportHCTableHCStatus_Unsupported SlbStatEnhRServerRportHCTableHCStatus = 2147483647
)

type SlbStatEnhRServerRportHCTableParams struct {
	// The real server index
	HCRealIndex string `json:"HCRealIndex,omitempty"`
	// The real server service index
	HCServIndex int32 `json:"HCServIndex,omitempty"`
	// The real server runtime instance health check ID
	HCId string `json:"HCId,omitempty"`
	// The real server runtime instance health check type
	HCType string `json:"HCType,omitempty"`
	// The real server runtime instance health check status
	HCStatus SlbStatEnhRServerRportHCTableHCStatus `json:"HCStatus,omitempty"`
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

func (p SlbStatEnhRServerRportHCTableParams) iMABean() {}
