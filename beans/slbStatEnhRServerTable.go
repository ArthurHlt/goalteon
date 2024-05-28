package beans

import (
	"fmt"
	"reflect"
)

// SlbStatEnhRServerTable The real server statistics table.
// Note:This mib is not supported for VX instance of Virtualization.
type SlbStatEnhRServerTable struct {
	// The real server index that identifies the server.
	SlbStatEnhRServerIndex string
	Params                 *SlbStatEnhRServerTableParams
}

func NewSlbStatEnhRServerTableList() *SlbStatEnhRServerTable {
	return &SlbStatEnhRServerTable{}
}

func NewSlbStatEnhRServerTable(
	slbStatEnhRServerIndex string,
	params *SlbStatEnhRServerTableParams,
) *SlbStatEnhRServerTable {
	return &SlbStatEnhRServerTable{
		SlbStatEnhRServerIndex: slbStatEnhRServerIndex,
		Params:                 params,
	}
}

func (c *SlbStatEnhRServerTable) Name() string {
	return "SlbStatEnhRServerTable"
}

func (c *SlbStatEnhRServerTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbStatEnhRServerTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbStatEnhRServerTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbStatEnhRServerIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbStatEnhRServerIndex)
}

type SlbStatEnhRServerTableParams struct {
	// The real server index that identifies the server.
	Index string `json:"Index,omitempty"`
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
}

func (p SlbStatEnhRServerTableParams) iMABean() {}
