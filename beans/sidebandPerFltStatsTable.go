package beans

import (
	"fmt"
	"reflect"
)

// SidebandPerFltStatsTable A table for Sideband statistics per filter.
// Note:This mib is not supported for VX instance of Virtualization.
type SidebandPerFltStatsTable struct {
	// Virtual server service port number.
	SidebandPerFltStatsFltIndex int32
	Params                      *SidebandPerFltStatsTableParams
}

func NewSidebandPerFltStatsTableList() *SidebandPerFltStatsTable {
	return &SidebandPerFltStatsTable{}
}

func NewSidebandPerFltStatsTable(
	sidebandPerFltStatsFltIndex int32,
	params *SidebandPerFltStatsTableParams,
) *SidebandPerFltStatsTable {
	return &SidebandPerFltStatsTable{
		SidebandPerFltStatsFltIndex: sidebandPerFltStatsFltIndex,
		Params:                      params,
	}
}

func (c *SidebandPerFltStatsTable) Name() string {
	return "SidebandPerFltStatsTable"
}

func (c *SidebandPerFltStatsTable) GetParams() BeanType {
	return c.Params
}

func (c *SidebandPerFltStatsTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SidebandPerFltStatsTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SidebandPerFltStatsFltIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SidebandPerFltStatsFltIndex)
}

type SidebandPerFltStatsTableParams struct {
	// Virtual server service port number.
	FltIndex int32 `json:"FltIndex,omitempty"`
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

func (p SidebandPerFltStatsTableParams) iMABean() {}
