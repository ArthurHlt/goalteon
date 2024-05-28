package beans

import (
	"fmt"
	"reflect"
)

// SlbStatEnhGroupTable The real server group statistics table.
// Note:This mib is not supported for VX instance of Virtualization.
type SlbStatEnhGroupTable struct {
	// The real server group number that identifies the group.
	SlbStatEnhGroupIndex string
	Params               *SlbStatEnhGroupTableParams
}

func NewSlbStatEnhGroupTableList() *SlbStatEnhGroupTable {
	return &SlbStatEnhGroupTable{}
}

func NewSlbStatEnhGroupTable(
	slbStatEnhGroupIndex string,
	params *SlbStatEnhGroupTableParams,
) *SlbStatEnhGroupTable {
	return &SlbStatEnhGroupTable{
		SlbStatEnhGroupIndex: slbStatEnhGroupIndex,
		Params:               params,
	}
}

func (c *SlbStatEnhGroupTable) Name() string {
	return "SlbStatEnhGroupTable"
}

func (c *SlbStatEnhGroupTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbStatEnhGroupTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbStatEnhGroupTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbStatEnhGroupIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbStatEnhGroupIndex)
}

type SlbStatEnhGroupTableParams struct {
	// The real server group number that identifies the group.
	Index string `json:"Index,omitempty"`
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

func (p SlbStatEnhGroupTableParams) iMABean() {}
