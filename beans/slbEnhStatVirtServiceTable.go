package beans

import (
	"fmt"
	"reflect"
)

// SlbEnhStatVirtServiceTable The virtual service statistics table.
// Note:This mib is not supported for VX instance of Virtualization.
type SlbEnhStatVirtServiceTable struct {
	// The virtual server number that identifies the service.
	SlbEnhStatVirtServerIndex string
	// The virtual service number that identifies the service.
	SlbEnhStatVirtServiceIndex int32
	// The real server index.
	SlbEnhStatRealServerIndex string
	Params                    *SlbEnhStatVirtServiceTableParams
}

func NewSlbEnhStatVirtServiceTableList() *SlbEnhStatVirtServiceTable {
	return &SlbEnhStatVirtServiceTable{}
}

func NewSlbEnhStatVirtServiceTable(
	slbEnhStatVirtServerIndex string,
	slbEnhStatVirtServiceIndex int32,
	slbEnhStatRealServerIndex string,
	params *SlbEnhStatVirtServiceTableParams,
) *SlbEnhStatVirtServiceTable {
	return &SlbEnhStatVirtServiceTable{
		SlbEnhStatVirtServerIndex:  slbEnhStatVirtServerIndex,
		SlbEnhStatVirtServiceIndex: slbEnhStatVirtServiceIndex,
		SlbEnhStatRealServerIndex:  slbEnhStatRealServerIndex,
		Params:                     params,
	}
}

func (c *SlbEnhStatVirtServiceTable) Name() string {
	return "SlbEnhStatVirtServiceTable"
}

func (c *SlbEnhStatVirtServiceTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbEnhStatVirtServiceTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbEnhStatVirtServiceTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbEnhStatVirtServerIndex).IsZero() &&
		reflect.ValueOf(c.SlbEnhStatVirtServiceIndex).IsZero() &&
		reflect.ValueOf(c.SlbEnhStatRealServerIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbEnhStatVirtServerIndex) + "/" + fmt.Sprint(c.SlbEnhStatVirtServiceIndex) + "/" + fmt.Sprint(c.SlbEnhStatRealServerIndex)
}

type SlbEnhStatVirtServiceTableParams struct {
	// The virtual server number that identifies the service.
	ServerIndex string `json:"ServerIndex,omitempty"`
	// The virtual service number that identifies the service.
	Index int32 `json:"Index,omitempty"`
	// The real server index.
	RealServerIndex string `json:"RealServerIndex,omitempty"`
	// The number of sessions that are currently handled by the
	// virtual service.
	CurrSessions uint32 `json:"CurrSessions,omitempty"`
	// The total number of sessions that are handled by the virtual service.
	TotalSessions uint32 `json:"TotalSessions,omitempty"`
	// The highest sessions that have been handled by the virtual service.
	HighestSessions uint32 `json:"HighestSessions,omitempty"`
	// The lower 32 bit value of octets received and transmitted out
	// of the virtual service.
	HCOctetsLow32 uint32 `json:"HCOctetsLow32,omitempty"`
	// The higher 32 bit value of octets received and transmitted out
	// of the virtual service.
	HCOctetsHigh32 uint32 `json:"HCOctetsHigh32,omitempty"`
	// The total octets received and transmitted by the virtual service.
	HCOctets uint64 `json:"HCOctets,omitempty"`
	// The real server status. (0)Up, (1)Down, (2)Admin-Down,
	// 				 (3)Warning, (4)Shutdown, (5)Error
	RealStatus int32 `json:"RealStatus,omitempty"`
	// Service health check failure reason.
	HcReason string `json:"HcReason,omitempty"`
	// The total number of RST sessions that are handled by the virtual service.
	RSTSessions uint32 `json:"RSTSessions,omitempty"`
	// The average ClientRTT in ms during the past collection interval.
	ClientRtt string `json:"ClientRtt,omitempty"`
	// The average ServerRTT in ms during the past collection interval.
	ServerRtt string `json:"ServerRtt,omitempty"`
	// Throughput currently handled by the virtual service.
	Thruput uint64 `json:"Thruput,omitempty"`
	// Total bandwidth handled by the virtual service.
	TotalBw uint64 `json:"TotalBw,omitempty"`
	// Number of packets that are currently handled per second by the virtual service.
	PktPerSec uint32 `json:"PktPerSec,omitempty"`
}

func (p SlbEnhStatVirtServiceTableParams) iMABean() {}
