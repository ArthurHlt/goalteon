package beans

import (
	"fmt"
	"reflect"
)

// SlbStatSpFltTable The filter statistics table.  This table shows the statistics
// of filters for a particular SP.
// Note:This mib is not supported for VX instance of Virtualization.
type SlbStatSpFltTable struct {
	// The SP index.
	SlbStatSpFltSpIndex int32
	// The filter number.
	SlbStatSpFltIndex int32
	Params            *SlbStatSpFltTableParams
}

func NewSlbStatSpFltTableList() *SlbStatSpFltTable {
	return &SlbStatSpFltTable{}
}

func NewSlbStatSpFltTable(
	slbStatSpFltSpIndex int32,
	slbStatSpFltIndex int32,
	params *SlbStatSpFltTableParams,
) *SlbStatSpFltTable {
	return &SlbStatSpFltTable{
		SlbStatSpFltSpIndex: slbStatSpFltSpIndex,
		SlbStatSpFltIndex:   slbStatSpFltIndex,
		Params:              params,
	}
}

func (c *SlbStatSpFltTable) Name() string {
	return "SlbStatSpFltTable"
}

func (c *SlbStatSpFltTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbStatSpFltTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbStatSpFltTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbStatSpFltSpIndex).IsZero() &&
		reflect.ValueOf(c.SlbStatSpFltIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbStatSpFltSpIndex) + "/" + fmt.Sprint(c.SlbStatSpFltIndex)
}

type SlbStatSpFltTableFiringsUnit int32

const (
	SlbStatSpFltTableFiringsUnit_Sessions    SlbStatSpFltTableFiringsUnit = 1
	SlbStatSpFltTableFiringsUnit_Requests    SlbStatSpFltTableFiringsUnit = 2
	SlbStatSpFltTableFiringsUnit_Packets     SlbStatSpFltTableFiringsUnit = 3
	SlbStatSpFltTableFiringsUnit_Unsupported SlbStatSpFltTableFiringsUnit = 2147483647
)

type SlbStatSpFltTableParams struct {
	// The SP index.
	SpIndex int32 `json:"SpIndex,omitempty"`
	// The filter number.
	Index int32 `json:"Index,omitempty"`
	// The number of instances that the received packet matches the
	// filter rule on this SP.
	Firings uint32 `json:"Firings,omitempty"`
	// Unit of total hits for filter as sessions or requests or packets.
	FiringsUnit SlbStatSpFltTableFiringsUnit `json:"FiringsUnit,omitempty"`
	// The current number of sessions for this filter on this SP.
	CurSess uint64 `json:"CurSess,omitempty"`
	// The highest number of session for this filter on this SP.
	HighSess uint64 `json:"HighSess,omitempty"`
	// The total number of session for this filter on this SP.
	TotSess uint64 `json:"TotSess,omitempty"`
	// The current bandwidth for this filter in Kbps on this SP.
	CurBand uint64 `json:"CurBand,omitempty"`
	// The highest bandwidth for this filter in Kbps on this SP.
	HighBand uint64 `json:"HighBand,omitempty"`
	// The total bandwidth for this filter in Mb on this SP
	TotBand uint64 `json:"TotBand,omitempty"`
	// The connections per second for this filter on this SP
	ConnPerSec uint64 `json:"ConnPerSec,omitempty"`
}

func (p SlbStatSpFltTableParams) iMABean() {}
