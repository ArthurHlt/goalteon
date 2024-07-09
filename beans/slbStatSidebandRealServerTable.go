package beans

import (
	"fmt"
	"reflect"
)

// SlbStatSidebandRealServerTable Sideband real servers statistics table.
// Note:This mib is not supported for VX instance of Virtualization.
type SlbStatSidebandRealServerTable struct {
	// The sideband policy name.
	SlbStatSidebandRealServerSidebandID string
	// The real server index.
	SlbStatSidebandRealServerRealServerID string
	Params                                *SlbStatSidebandRealServerTableParams
}

func NewSlbStatSidebandRealServerTableList() *SlbStatSidebandRealServerTable {
	return &SlbStatSidebandRealServerTable{}
}

func NewSlbStatSidebandRealServerTable(
	slbStatSidebandRealServerSidebandID string,
	slbStatSidebandRealServerRealServerID string,
	params *SlbStatSidebandRealServerTableParams,
) *SlbStatSidebandRealServerTable {
	return &SlbStatSidebandRealServerTable{
		SlbStatSidebandRealServerSidebandID:   slbStatSidebandRealServerSidebandID,
		SlbStatSidebandRealServerRealServerID: slbStatSidebandRealServerRealServerID,
		Params:                                params,
	}
}

func (c *SlbStatSidebandRealServerTable) Name() string {
	return "SlbStatSidebandRealServerTable"
}

func (c *SlbStatSidebandRealServerTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbStatSidebandRealServerTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbStatSidebandRealServerTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbStatSidebandRealServerSidebandID).IsZero() &&
		reflect.ValueOf(c.SlbStatSidebandRealServerRealServerID).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbStatSidebandRealServerSidebandID) + "/" + fmt.Sprint(c.SlbStatSidebandRealServerRealServerID)
}

type SlbStatSidebandRealServerTableParams struct {
	// The sideband policy name.
	SidebandID string `json:"SidebandID,omitempty"`
	// The real server index.
	RealServerID string `json:"RealServerID,omitempty"`
	// The number of sessions that are currently handled by the
	// sideband.
	CurrSessions uint32 `json:"CurrSessions,omitempty"`
	// The total number of sessions that are handled by the sideband.
	TotalSessions uint32 `json:"TotalSessions,omitempty"`
	// The highest sessions that have been handled by the sideband.
	HighestSessions uint32 `json:"HighestSessions,omitempty"`
	// The total octets received and transmitted by the sideband.
	Octets string `json:"Octets,omitempty"`
	// The real server status. (0)Up, (1)Down, (2)Admin-Down,
	// (3)Warning, (4)Shutdown, (5)Error
	RealStatus int32 `json:"RealStatus,omitempty"`
	// Sideband health check failure reason.
	HcReason string `json:"HcReason,omitempty"`
	// The average ServerRTT in ms during the past collection interval.
	ServerRtt string `json:"ServerRtt,omitempty"`
	// sideband real server connection per seconds.
	SessionsPerSec uint64 `json:"SessionsPerSec,omitempty"`
	// sideband real server throughput.
	OctetsPerSec uint64 `json:"OctetsPerSec,omitempty"`
}

func (p SlbStatSidebandRealServerTableParams) iMABean() {}
