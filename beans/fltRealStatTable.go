package beans

import (
	"fmt"
	"reflect"
)

// FltRealStatTable The filter statistics table for real servers.
// Note:This mib is not supported for VX instance of Virtualization.
type FltRealStatTable struct {
	// The number that identifies the filter.
	FltRealStatFltIndex int32
	// The real server index.
	SlbRealStatRealIndex string
	Params               *FltRealStatTableParams
}

func NewFltRealStatTableList() *FltRealStatTable {
	return &FltRealStatTable{}
}

func NewFltRealStatTable(
	fltRealStatFltIndex int32,
	slbRealStatRealIndex string,
	params *FltRealStatTableParams,
) *FltRealStatTable {
	return &FltRealStatTable{
		FltRealStatFltIndex:  fltRealStatFltIndex,
		SlbRealStatRealIndex: slbRealStatRealIndex,
		Params:               params,
	}
}

func (c *FltRealStatTable) Name() string {
	return "FltRealStatTable"
}

func (c *FltRealStatTable) GetParams() BeanType {
	return c.Params
}

func (c *FltRealStatTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *FltRealStatTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.FltRealStatFltIndex).IsZero() &&
		reflect.ValueOf(c.SlbRealStatRealIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.FltRealStatFltIndex) + "/" + fmt.Sprint(c.SlbRealStatRealIndex)
}

type FltRealStatTableParams struct {
	// The number that identifies the filter.
	FltIndex int32 `json:"FltIndex,omitempty"`
	// The real server index.
	SlbRealStatRealIndex string `json:"slbRealStatRealIndex,omitempty"`
	// IP address of the real server.
	SlbRealStatRealIPAddr string `json:"slbRealStatRealIPAddr,omitempty"`
	// The current number of sessions for the real server.
	RealCurSess uint64 `json:"RealCurSess,omitempty"`
	// The highest number of session for the real server.
	RealHighSess uint64 `json:"RealHighSess,omitempty"`
	// The total number of session for the real server.
	RealTotSess uint64 `json:"RealTotSess,omitempty"`
	// The current bandwidth for the real server in Kbps.
	RealCurBand uint64 `json:"RealCurBand,omitempty"`
	// The highest bandwidth for the real server in Kbps.
	RealHighBand uint64 `json:"RealHighBand,omitempty"`
	// The total bandwidth for the real server in Mb.
	RealTotBand uint64 `json:"RealTotBand,omitempty"`
	// The connections per second for the real server.
	RealConnPerSec uint64 `json:"RealConnPerSec,omitempty"`
	// The filter HC failure reason description.
	RealHcReason string `json:"RealHcReason,omitempty"`
	// The packets per second for the real server.
	RealPktPerSec uint64 `json:"RealPktPerSec,omitempty"`
}

func (p FltRealStatTableParams) iMABean() {}
