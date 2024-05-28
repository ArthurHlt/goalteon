package beans

import (
	"fmt"
	"reflect"
)

// SecPolPerServStatsTable A table for security policy statistics per virtual service.
// Note:This mib is not supported on VX instance of Virtualization.
type SecPolPerServStatsTable struct {
	// Virtual server number in alphanumeric.
	SecPolPerServVirtServIndex string
	// Virtual server service index.
	SecPolPerServVirtServiceIndex int32
	Params                        *SecPolPerServStatsTableParams
}

func NewSecPolPerServStatsTableList() *SecPolPerServStatsTable {
	return &SecPolPerServStatsTable{}
}

func NewSecPolPerServStatsTable(
	secPolPerServVirtServIndex string,
	secPolPerServVirtServiceIndex int32,
	params *SecPolPerServStatsTableParams,
) *SecPolPerServStatsTable {
	return &SecPolPerServStatsTable{
		SecPolPerServVirtServIndex:    secPolPerServVirtServIndex,
		SecPolPerServVirtServiceIndex: secPolPerServVirtServiceIndex,
		Params:                        params,
	}
}

func (c *SecPolPerServStatsTable) Name() string {
	return "SecPolPerServStatsTable"
}

func (c *SecPolPerServStatsTable) GetParams() BeanType {
	return c.Params
}

func (c *SecPolPerServStatsTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SecPolPerServStatsTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SecPolPerServVirtServIndex).IsZero() &&
		reflect.ValueOf(c.SecPolPerServVirtServiceIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SecPolPerServVirtServIndex) + "/" + fmt.Sprint(c.SecPolPerServVirtServiceIndex)
}

type SecPolPerServStatsTableParams struct {
	// Virtual server number in alphanumeric.
	VirtServIndex string `json:"VirtServIndex,omitempty"`
	// Virtual server service index.
	VirtServiceIndex int32 `json:"VirtServiceIndex,omitempty"`
	// Virtual server service port number.
	VirtServPort int32 `json:"VirtServPort,omitempty"`
	// Security policy identifier associated with the virtual service.
	SecPolId string `json:"SecPolId,omitempty"`
	// Virtual service security policy stats - Bandwidth current value.
	BwCurVal int32 `json:"BwCurVal,omitempty"`
	// Virtual service security policy stats - Bandwidth last period average.
	BwLastPeriodAvg int32 `json:"BwLastPeriodAvg,omitempty"`
	// Virtual service security policy stats - Bandwidth current period average.
	BwCurPeriodAvg int32 `json:"BwCurPeriodAvg,omitempty"`
	// Virtual service security policy stats - Bandwidth peak value.
	BwPeak int32 `json:"BwPeak,omitempty"`
	// Virtual service security policy stats - Bandwidth peak timestamp.
	BwPeakTimestamp string `json:"BwPeakTimestamp,omitempty"`
	// Virtual service security policy stats - PPS current value.
	PPSCurVal int32 `json:"PPSCurVal,omitempty"`
	// Virtual service security policy stats - PPS last period average.
	PPSLastPeriodAvg int32 `json:"PPSLastPeriodAvg,omitempty"`
	// Virtual service security policy stats - PPS current period average.
	PPSCurPeriodAvg int32 `json:"PPSCurPeriodAvg,omitempty"`
	// Virtual service security policy stats - PPS peak value.
	PPSPeak int32 `json:"PPSPeak,omitempty"`
	// Virtual service security policy stats - PPS peak timestamp.
	PPSPeakTimestamp string `json:"PPSPeakTimestamp,omitempty"`
	// Virtual service security policy stats - CPS current value.
	CPSCurVal int32 `json:"CPSCurVal,omitempty"`
	// Virtual service security policy stats - CPS last period average.
	CPSLastPeriodAvg int32 `json:"CPSLastPeriodAvg,omitempty"`
	// Virtual service security policy stats - CPS current period average.
	CPSCurPeriodAvg int32 `json:"CPSCurPeriodAvg,omitempty"`
	// Virtual service security policy stats - CPS peak value.
	CPSPeak int32 `json:"CPSPeak,omitempty"`
	// Virtual service security policy stats - CPS peak timestamp.
	CPSPeakTimestamp string `json:"CPSPeakTimestamp,omitempty"`
	// Virtual service security policy stats - CEC current value.
	CECCurVal int32 `json:"CECCurVal,omitempty"`
	// Virtual service security policy stats - CEC last period average.
	CECLastPeriodAvg int32 `json:"CECLastPeriodAvg,omitempty"`
	// Virtual service security policy stats - CEC current period average.
	CECCurPeriodAvg int32 `json:"CECCurPeriodAvg,omitempty"`
	// Virtual service security policy stats - CEC peak value.
	CECPeak int32 `json:"CECPeak,omitempty"`
	// Virtual service security policy stats - CEC peak timestamp.
	CECPeakTimestamp string `json:"CECPeakTimestamp,omitempty"`
	// Virtual service security policy stats - Latency current value.
	LatencyCurVal int32 `json:"LatencyCurVal,omitempty"`
	// Virtual service security policy stats - Latency last period average.
	LatencyLastPeriodAvg int32 `json:"LatencyLastPeriodAvg,omitempty"`
	// Virtual service security policy stats - Latency current period average.
	LatencyCurPeriodAvg int32 `json:"LatencyCurPeriodAvg,omitempty"`
	// Virtual service security policy stats - Latency peak value.
	LatencyPeak int32 `json:"LatencyPeak,omitempty"`
	// Virtual service security policy stats - Latency peak timestamp.
	LatencyPeakTimestamp string `json:"LatencyPeakTimestamp,omitempty"`
}

func (p SecPolPerServStatsTableParams) iMABean() {}
