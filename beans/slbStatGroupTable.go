package beans

import (
	"fmt"
	"reflect"
)

// SlbStatGroupTable The real server group statistics table.
// Note:This mib is not supported for VX instance of Virtualization.
type SlbStatGroupTable struct {
	// The real server group number that identifies the group.
	SlbStatGroupIndex int32
	Params            *SlbStatGroupTableParams
}

func NewSlbStatGroupTableList() *SlbStatGroupTable {
	return &SlbStatGroupTable{}
}

func NewSlbStatGroupTable(
	slbStatGroupIndex int32,
	params *SlbStatGroupTableParams,
) *SlbStatGroupTable {
	return &SlbStatGroupTable{
		SlbStatGroupIndex: slbStatGroupIndex,
		Params:            params,
	}
}

func (c *SlbStatGroupTable) Name() string {
	return "SlbStatGroupTable"
}

func (c *SlbStatGroupTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbStatGroupTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbStatGroupTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbStatGroupIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbStatGroupIndex)
}

type SlbStatGroupTableParams struct {
	// The real server group number that identifies the group.
	Index int32 `json:"Index,omitempty"`
	// The number of sessions that are currently handled by the real server
	// group.
	CurrSessions uint32 `json:"CurrSessions,omitempty"`
	// The total number of sessions that are handled by the real
	// server group.
	TotalSessions uint32 `json:"TotalSessions,omitempty"`
	// The highest sessions that have been handled by the real server group.
	HighestSessions uint32 `json:"HighestSessions,omitempty"`
	// The lower 32 bit value of octets received and transmitted out of the
	// real server group.
	HCOctetsLow32 uint32 `json:"HCOctetsLow32,omitempty"`
	// The higher 32 bit value of octets received and transmitted out
	// of the real server group.
	HCOctetsHigh32 uint32 `json:"HCOctetsHigh32,omitempty"`
	// The total octets received and transmitted by the real server group.
	HCOctets uint64 `json:"HCOctets,omitempty"`
	// Total weight updates from work load manager.
	WlmUpdates uint32 `json:"WlmUpdates,omitempty"`
}

func (p SlbStatGroupTableParams) iMABean() {}
