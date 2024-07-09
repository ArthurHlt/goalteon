package beans

import (
	"fmt"
	"reflect"
)

// SidebandPerServStatsTable A table for Sideband statistics per virtual service.
// Note:This mib is not supported for VX instance of Virtualization.
type SidebandPerServStatsTable struct {
	// Virtual server number.
	SidebandPerServStatsVirtServIndex string
	// Virtual server service index.
	SidebandPerServStatsVirtServiceIndex int32
	Params                               *SidebandPerServStatsTableParams
}

func NewSidebandPerServStatsTableList() *SidebandPerServStatsTable {
	return &SidebandPerServStatsTable{}
}

func NewSidebandPerServStatsTable(
	sidebandPerServStatsVirtServIndex string,
	sidebandPerServStatsVirtServiceIndex int32,
	params *SidebandPerServStatsTableParams,
) *SidebandPerServStatsTable {
	return &SidebandPerServStatsTable{
		SidebandPerServStatsVirtServIndex:    sidebandPerServStatsVirtServIndex,
		SidebandPerServStatsVirtServiceIndex: sidebandPerServStatsVirtServiceIndex,
		Params:                               params,
	}
}

func (c *SidebandPerServStatsTable) Name() string {
	return "SidebandPerServStatsTable"
}

func (c *SidebandPerServStatsTable) GetParams() BeanType {
	return c.Params
}

func (c *SidebandPerServStatsTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SidebandPerServStatsTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SidebandPerServStatsVirtServIndex).IsZero() &&
		reflect.ValueOf(c.SidebandPerServStatsVirtServiceIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SidebandPerServStatsVirtServIndex) + "/" + fmt.Sprint(c.SidebandPerServStatsVirtServiceIndex)
}

type SidebandPerServStatsTableParams struct {
	// Virtual server number.
	VirtServIndex string `json:"VirtServIndex,omitempty"`
	// Virtual server service index.
	VirtServiceIndex int32 `json:"VirtServiceIndex,omitempty"`
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

func (p SidebandPerServStatsTableParams) iMABean() {}
