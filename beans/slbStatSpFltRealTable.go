package beans

import (
	"fmt"
	"reflect"
)

// SlbStatSpFltRealTable The filter statistics table for real servers for a particular SP.
// Note:This mib is not supported for VX instance of Virtualization.
type SlbStatSpFltRealTable struct {
	// The SP index.
	SlbStatSpFltRealSpIndex int32
	// The number that identifies the filter.
	SlbStatSpFltRealFltIndex int32
	// The real server index.
	SlbStatSpFltRealRealIndex string
	Params                    *SlbStatSpFltRealTableParams
}

func NewSlbStatSpFltRealTableList() *SlbStatSpFltRealTable {
	return &SlbStatSpFltRealTable{}
}

func NewSlbStatSpFltRealTable(
	slbStatSpFltRealSpIndex int32,
	slbStatSpFltRealFltIndex int32,
	slbStatSpFltRealRealIndex string,
	params *SlbStatSpFltRealTableParams,
) *SlbStatSpFltRealTable {
	return &SlbStatSpFltRealTable{
		SlbStatSpFltRealSpIndex:   slbStatSpFltRealSpIndex,
		SlbStatSpFltRealFltIndex:  slbStatSpFltRealFltIndex,
		SlbStatSpFltRealRealIndex: slbStatSpFltRealRealIndex,
		Params:                    params,
	}
}

func (c *SlbStatSpFltRealTable) Name() string {
	return "SlbStatSpFltRealTable"
}

func (c *SlbStatSpFltRealTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbStatSpFltRealTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbStatSpFltRealTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbStatSpFltRealSpIndex).IsZero() &&
		reflect.ValueOf(c.SlbStatSpFltRealFltIndex).IsZero() &&
		reflect.ValueOf(c.SlbStatSpFltRealRealIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbStatSpFltRealSpIndex) + "/" + fmt.Sprint(c.SlbStatSpFltRealFltIndex) + "/" + fmt.Sprint(c.SlbStatSpFltRealRealIndex)
}

type SlbStatSpFltRealTableParams struct {
	// The SP index.
	SpIndex int32 `json:"SpIndex,omitempty"`
	// The number that identifies the filter.
	FltIndex int32 `json:"FltIndex,omitempty"`
	// The real server index.
	RealIndex string `json:"RealIndex,omitempty"`
	// IP address of the real server.
	IPAddr string `json:"IPAddr,omitempty"`
	// The current number of sessions for the real server on this SP.
	CurSess uint64 `json:"CurSess,omitempty"`
	// The highest number of session for the real server on this SP.
	HighSess uint64 `json:"HighSess,omitempty"`
	// The total number of session for the real server on this SP.
	TotSess uint64 `json:"TotSess,omitempty"`
	// The current bandwidth for the real server in Kbps on this SP.
	CurBand uint64 `json:"CurBand,omitempty"`
	// The highest bandwidth for the real server in Kbps on this SP.
	HighBand uint64 `json:"HighBand,omitempty"`
	// The total bandwidth for the real server in Mb on this SP.
	TotBand uint64 `json:"TotBand,omitempty"`
	// The connections per second for the real server on this SP.
	ConnPerSec uint64 `json:"ConnPerSec,omitempty"`
}

func (p SlbStatSpFltRealTableParams) iMABean() {}
