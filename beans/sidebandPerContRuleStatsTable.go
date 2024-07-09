package beans

import (
	"fmt"
	"reflect"
)

// SidebandPerContRuleStatsTable A table for Sideband statistics per content rule.
// Note:This mib is not supported for VX instance of Virtualization.
type SidebandPerContRuleStatsTable struct {
	// Virtual server number.
	SidebandPerContRuleStatsVirtServIndex string
	// Virtual server service index.
	SidebandPerContRuleStatsServiceIndex int32
	// Virtual server service index.
	SidebandPerContRuleStatsContRuleIndex int32
	Params                                *SidebandPerContRuleStatsTableParams
}

func NewSidebandPerContRuleStatsTableList() *SidebandPerContRuleStatsTable {
	return &SidebandPerContRuleStatsTable{}
}

func NewSidebandPerContRuleStatsTable(
	sidebandPerContRuleStatsVirtServIndex string,
	sidebandPerContRuleStatsServiceIndex int32,
	sidebandPerContRuleStatsContRuleIndex int32,
	params *SidebandPerContRuleStatsTableParams,
) *SidebandPerContRuleStatsTable {
	return &SidebandPerContRuleStatsTable{
		SidebandPerContRuleStatsVirtServIndex: sidebandPerContRuleStatsVirtServIndex,
		SidebandPerContRuleStatsServiceIndex:  sidebandPerContRuleStatsServiceIndex,
		SidebandPerContRuleStatsContRuleIndex: sidebandPerContRuleStatsContRuleIndex,
		Params:                                params,
	}
}

func (c *SidebandPerContRuleStatsTable) Name() string {
	return "SidebandPerContRuleStatsTable"
}

func (c *SidebandPerContRuleStatsTable) GetParams() BeanType {
	return c.Params
}

func (c *SidebandPerContRuleStatsTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SidebandPerContRuleStatsTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SidebandPerContRuleStatsVirtServIndex).IsZero() &&
		reflect.ValueOf(c.SidebandPerContRuleStatsServiceIndex).IsZero() &&
		reflect.ValueOf(c.SidebandPerContRuleStatsContRuleIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SidebandPerContRuleStatsVirtServIndex) + "/" + fmt.Sprint(c.SidebandPerContRuleStatsServiceIndex) + "/" + fmt.Sprint(c.SidebandPerContRuleStatsContRuleIndex)
}

type SidebandPerContRuleStatsTableParams struct {
	// Virtual server number.
	VirtServIndex string `json:"VirtServIndex,omitempty"`
	// Virtual server service index.
	ServiceIndex int32 `json:"ServiceIndex,omitempty"`
	// Virtual server service index.
	ContRuleIndex int32 `json:"ContRuleIndex,omitempty"`
	// Virtual server service port number.
	VirtServPort int32 `json:"VirtServPort,omitempty"`
	// Current average latency of sideband processing.
	AvgLatency int32 `json:"AvgLatency,omitempty"`
	// Total average latency of sideband processing.
	AvgLatencyTotal uint64 `json:"AvgLatencyTotal,omitempty"`
	// Number of requests that sent for sideband processing.
	ReqSentToSideband int32 `json:"ReqSentToSideband,omitempty"`
	// Total number of requests that sent for sideband processing.
	ReqSentToSidebandTotal uint64 `json:"ReqSentToSidebandTotal,omitempty"`
	// Number of requests that bypassed sideband processing.
	ReqFilteredFromSideband int32 `json:"ReqFilteredFromSideband,omitempty"`
	// Total number of requests that bypassed sideband processing.
	ReqFilteredFromSidebandTotal uint64 `json:"ReqFilteredFromSidebandTotal,omitempty"`
	// Number of requests that the sideband didn't respond to in time.
	SidebandTimeouts int32 `json:"SidebandTimeouts,omitempty"`
	// Total Number of requests that the sideband didn't respond to in time.
	SidebandTimeoutsTotal uint64 `json:"SidebandTimeoutsTotal,omitempty"`
	// Number of requests we sent to server after sideband processing.
	SentToServer int32 `json:"SentToServer,omitempty"`
	// Total number of requests we sent to server after sideband processing.
	SentToServerTotal uint64 `json:"SentToServerTotal,omitempty"`
	// Number of requests that were responded by Alteon due to sideband processing.
	ResponseToClient int32 `json:"ResponseToClient,omitempty"`
	// Total Number of requests that were responded by Alteon due to sideband processing.
	ResponseToClientTotal uint64 `json:"ResponseToClientTotal,omitempty"`
	// Number of client connections closed due to sideband processing.
	ClientConnectionClosed int32 `json:"ClientConnectionClosed,omitempty"`
	// Total Number of client connections closed due to sideband processing.
	ClientConnectionClosedTotal uint64 `json:"ClientConnectionClosedTotal,omitempty"`
	// Number requests that resulted in sideband failure.
	SidebandFailure int32 `json:"SidebandFailure,omitempty"`
	// Total Number requests that resulted in sideband failure.
	SidebandFailureTotal uint64 `json:"SidebandFailureTotal,omitempty"`
}

func (p SidebandPerContRuleStatsTableParams) iMABean() {}
