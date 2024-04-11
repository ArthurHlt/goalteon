package beans

import (
	"fmt"
	"reflect"
)

// SlbStatSpVServerTable The virtual server statistics table. This table shows the statistics
// of virtual server for a particular SP.
// Note:This mib is not supported for VX instance of Virtualization.
type SlbStatSpVServerTable struct {
	// The SP index.
	SlbStatSpVServerSpIndex int32
	// The virtual server number that identifies the server.
	SlbStatSpVServerIndex string
	Params                *SlbStatSpVServerTableParams
}

func NewSlbStatSpVServerTableList() *SlbStatSpVServerTable {
	return &SlbStatSpVServerTable{}
}

func NewSlbStatSpVServerTable(
	slbStatSpVServerSpIndex int32,
	slbStatSpVServerIndex string,
	params *SlbStatSpVServerTableParams,
) *SlbStatSpVServerTable {
	return &SlbStatSpVServerTable{
		SlbStatSpVServerSpIndex: slbStatSpVServerSpIndex,
		SlbStatSpVServerIndex:   slbStatSpVServerIndex,
		Params:                  params,
	}
}

func (c *SlbStatSpVServerTable) Name() string {
	return "SlbStatSpVServerTable"
}

func (c *SlbStatSpVServerTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbStatSpVServerTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbStatSpVServerTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbStatSpVServerSpIndex).IsZero() &&
		reflect.ValueOf(c.SlbStatSpVServerIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbStatSpVServerSpIndex) + "/" + fmt.Sprint(c.SlbStatSpVServerIndex)
}

type SlbStatSpVServerTableParams struct {
	// The SP index.
	SpIndex int32 `json:"SpIndex,omitempty"`
	// The virtual server number that identifies the server.
	Index string `json:"Index,omitempty"`
	// The number of persistent sessions that are currently handled by the
	// virtual server on this SP.
	CurrPSessions uint64 `json:"CurrPSessions,omitempty"`
	// The total number of persistent sessions that are handled by the
	// virtual server on this SP.
	TotalPSessions uint64 `json:"TotalPSessions,omitempty"`
	// The highest number of persistent sessions that have been handled
	// by the virtual server on this SP.
	HighestPSessions uint64 `json:"HighestPSessions,omitempty"`
}

func (p SlbStatSpVServerTableParams) iMABean() {}
