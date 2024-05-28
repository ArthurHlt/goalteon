package beans

import (
	"fmt"
	"reflect"
)

// SlbStatVirtServiceTable The virtual service statistics table.
// Note:This mib is not supported for VX instance of Virtualization.
type SlbStatVirtServiceTable struct {
	// The virtual server number that identifies the service.
	SlbStatVirtServerIndex int32
	// The virtual service number that identifies the service.
	SlbStatVirtServiceIndex int32
	// The real server index.
	SlbStatRealServerIndex int32
	Params                 *SlbStatVirtServiceTableParams
}

func NewSlbStatVirtServiceTableList() *SlbStatVirtServiceTable {
	return &SlbStatVirtServiceTable{}
}

func NewSlbStatVirtServiceTable(
	slbStatVirtServerIndex int32,
	slbStatVirtServiceIndex int32,
	slbStatRealServerIndex int32,
	params *SlbStatVirtServiceTableParams,
) *SlbStatVirtServiceTable {
	return &SlbStatVirtServiceTable{
		SlbStatVirtServerIndex:  slbStatVirtServerIndex,
		SlbStatVirtServiceIndex: slbStatVirtServiceIndex,
		SlbStatRealServerIndex:  slbStatRealServerIndex,
		Params:                  params,
	}
}

func (c *SlbStatVirtServiceTable) Name() string {
	return "SlbStatVirtServiceTable"
}

func (c *SlbStatVirtServiceTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbStatVirtServiceTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbStatVirtServiceTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbStatVirtServerIndex).IsZero() &&
		reflect.ValueOf(c.SlbStatVirtServiceIndex).IsZero() &&
		reflect.ValueOf(c.SlbStatRealServerIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbStatVirtServerIndex) + "/" + fmt.Sprint(c.SlbStatVirtServiceIndex) + "/" + fmt.Sprint(c.SlbStatRealServerIndex)
}

type SlbStatVirtServiceTableParams struct {
	// The virtual server number that identifies the service.
	ServerIndex int32 `json:"ServerIndex,omitempty"`
	// The virtual service number that identifies the service.
	Index int32 `json:"Index,omitempty"`
	// The real server index.
	RealServerIndex int32 `json:"RealServerIndex,omitempty"`
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
}

func (p SlbStatVirtServiceTableParams) iMABean() {}
