package beans

import (
	"fmt"
	"reflect"
)

// SlbStatRServerTable The real server statistics table.
// Note:This mib is not supported for VX instance of Virtualization.
type SlbStatRServerTable struct {
	// The real server number that identifies the server.
	SlbStatRServerIndex int32
	Params              *SlbStatRServerTableParams
}

func NewSlbStatRServerTableList() *SlbStatRServerTable {
	return &SlbStatRServerTable{}
}

func NewSlbStatRServerTable(
	slbStatRServerIndex int32,
	params *SlbStatRServerTableParams,
) *SlbStatRServerTable {
	return &SlbStatRServerTable{
		SlbStatRServerIndex: slbStatRServerIndex,
		Params:              params,
	}
}

func (c *SlbStatRServerTable) Name() string {
	return "SlbStatRServerTable"
}

func (c *SlbStatRServerTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbStatRServerTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbStatRServerTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbStatRServerIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbStatRServerIndex)
}

type SlbStatRServerTableParams struct {
	// The real server number that identifies the server.
	Index int32 `json:"Index,omitempty"`
	// The number of sessions that are currently handled by the real server.
	CurrSessions uint32 `json:"CurrSessions,omitempty"`
	// The total number of sessions that are handled by the real server.
	TotalSessions uint32 `json:"TotalSessions,omitempty"`
	// The total number of times that the real server is claimed down.
	Failures uint32 `json:"Failures,omitempty"`
	// The highest sessions that have been handled by the real server.
	HighestSessions uint32 `json:"HighestSessions,omitempty"`
	// The lower 32 bit value of octets received and transmitted out of the
	// real server.
	HCOctetsLow32 uint32 `json:"HCOctetsLow32,omitempty"`
	// The higher 32 bit value of octets received and transmitted out
	// of the real server.
	HCOctetsHigh32 uint32 `json:"HCOctetsHigh32,omitempty"`
	// The total number of octets received and transmitted by the real
	// server.
	HCOctets uint64 `json:"HCOctets,omitempty"`
	// The total up time of the real server.
	UpTime string `json:"UpTime,omitempty"`
	// The total down time of the real server.
	DownTime string `json:"DownTime,omitempty"`
	// The time-stamp when the last failure of the real server occurred.
	LastFailureTime string `json:"LastFailureTime,omitempty"`
}

func (p SlbStatRServerTableParams) iMABean() {}
