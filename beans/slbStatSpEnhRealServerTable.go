package beans

import (
	"fmt"
	"reflect"
)

// SlbStatSpEnhRealServerTable The sp-server statistics table.  This table shows the statistics
// of real servers for a particular SP.
// Note:This mib is not supported for VX instance of Virtualization.
type SlbStatSpEnhRealServerTable struct {
	// The SP index.
	SlbStatSpEnhRealServerSpIndex int32
	// The real server number that identifies the server.
	SlbStatSpEnhRealServerServerIndex string
	Params                            *SlbStatSpEnhRealServerTableParams
}

func NewSlbStatSpEnhRealServerTableList() *SlbStatSpEnhRealServerTable {
	return &SlbStatSpEnhRealServerTable{}
}

func NewSlbStatSpEnhRealServerTable(
	slbStatSpEnhRealServerSpIndex int32,
	slbStatSpEnhRealServerServerIndex string,
	params *SlbStatSpEnhRealServerTableParams,
) *SlbStatSpEnhRealServerTable {
	return &SlbStatSpEnhRealServerTable{
		SlbStatSpEnhRealServerSpIndex:     slbStatSpEnhRealServerSpIndex,
		SlbStatSpEnhRealServerServerIndex: slbStatSpEnhRealServerServerIndex,
		Params:                            params,
	}
}

func (c *SlbStatSpEnhRealServerTable) Name() string {
	return "SlbStatSpEnhRealServerTable"
}

func (c *SlbStatSpEnhRealServerTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbStatSpEnhRealServerTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbStatSpEnhRealServerTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbStatSpEnhRealServerSpIndex).IsZero() &&
		reflect.ValueOf(c.SlbStatSpEnhRealServerServerIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbStatSpEnhRealServerSpIndex) + "/" + fmt.Sprint(c.SlbStatSpEnhRealServerServerIndex)
}

type SlbStatSpEnhRealServerTableParams struct {
	// The SP index.
	SpIndex int32 `json:"SpIndex,omitempty"`
	// The real server number that identifies the server.
	ServerIndex string `json:"ServerIndex,omitempty"`
	// The current sessions for the real server.
	CurrSessions uint32 `json:"CurrSessions,omitempty"`
	// The total number of sessions for the real server.
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

func (p SlbStatSpEnhRealServerTableParams) iMABean() {}
