package beans

import (
	"fmt"
	"reflect"
)

// SlbStatSpRealServerTable The sp-server statistics table.  This table shows the statistics
// of real servers for a particular SP.
// Note:This mib is not supported for VX instance of Virtualization.
type SlbStatSpRealServerTable struct {
	// The SP index.
	SlbStatSpRealServerSpIndex int32
	// The real server number that identifies the server.
	SlbStatSpRealServerServerIndex int32
	Params                         *SlbStatSpRealServerTableParams
}

func NewSlbStatSpRealServerTableList() *SlbStatSpRealServerTable {
	return &SlbStatSpRealServerTable{}
}

func NewSlbStatSpRealServerTable(
	slbStatSpRealServerSpIndex int32,
	slbStatSpRealServerServerIndex int32,
	params *SlbStatSpRealServerTableParams,
) *SlbStatSpRealServerTable {
	return &SlbStatSpRealServerTable{
		SlbStatSpRealServerSpIndex:     slbStatSpRealServerSpIndex,
		SlbStatSpRealServerServerIndex: slbStatSpRealServerServerIndex,
		Params:                         params,
	}
}

func (c *SlbStatSpRealServerTable) Name() string {
	return "SlbStatSpRealServerTable"
}

func (c *SlbStatSpRealServerTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbStatSpRealServerTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbStatSpRealServerTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbStatSpRealServerSpIndex).IsZero() &&
		reflect.ValueOf(c.SlbStatSpRealServerServerIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbStatSpRealServerSpIndex) + "/" + fmt.Sprint(c.SlbStatSpRealServerServerIndex)
}

type SlbStatSpRealServerTableParams struct {
	// The SP index.
	SpIndex int32 `json:"SpIndex,omitempty"`
	// The real server number that identifies the server.
	ServerIndex int32 `json:"ServerIndex,omitempty"`
	// The current sessions for the real server on a particular SP.
	CurrSessions uint32 `json:"CurrSessions,omitempty"`
	// The SP total number of sessions for the real server on a particular SP.
	TotalSessions uint32 `json:"TotalSessions,omitempty"`
	// The lower 32 bit value of the total octets received and transmitted
	// out of the real server on a particular SP.
	HCOctetsLow32 uint32 `json:"HCOctetsLow32,omitempty"`
	// The higher 32 bit value of the total octets received and transmitted
	// out of the real server on a particular SP.
	HCOctetsHigh32 uint32 `json:"HCOctetsHigh32,omitempty"`
	// The total octets received and transmitted by the real server on a
	// particular SP.
	HCOctets uint64 `json:"HCOctets,omitempty"`
}

func (p SlbStatSpRealServerTableParams) iMABean() {}
