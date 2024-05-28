package beans

import (
	"fmt"
	"reflect"
)

// FltStatTable The filter statistics table.
// Note:This mib is not supported for VX instance of Virtualization.
type FltStatTable struct {
	// The number that identifies the filter.
	FltStatFltIndex int32
	Params          *FltStatTableParams
}

func NewFltStatTableList() *FltStatTable {
	return &FltStatTable{}
}

func NewFltStatTable(
	fltStatFltIndex int32,
	params *FltStatTableParams,
) *FltStatTable {
	return &FltStatTable{
		FltStatFltIndex: fltStatFltIndex,
		Params:          params,
	}
}

func (c *FltStatTable) Name() string {
	return "FltStatTable"
}

func (c *FltStatTable) GetParams() BeanType {
	return c.Params
}

func (c *FltStatTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *FltStatTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.FltStatFltIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.FltStatFltIndex)
}

type FltStatTableFltFiringsUnit int32

const (
	FltStatTableFltFiringsUnit_Sessions    FltStatTableFltFiringsUnit = 1
	FltStatTableFltFiringsUnit_Requests    FltStatTableFltFiringsUnit = 2
	FltStatTableFltFiringsUnit_Packets     FltStatTableFltFiringsUnit = 3
	FltStatTableFltFiringsUnit_Unsupported FltStatTableFltFiringsUnit = 2147483647
)

type FltStatTableParams struct {
	// The number that identifies the filter.
	FltIndex int32 `json:"FltIndex,omitempty"`
	// The number of instances that the received packet matches the
	// filter rule.
	FltFirings uint32 `json:"FltFirings,omitempty"`
	// The number of instances that the received packet matches the
	// filter rule.
	FltFiringsU64 uint64 `json:"FltFiringsU64,omitempty"`
	// Unit of total hits for filter as sessions or requests or packets.
	FltFiringsUnit FltStatTableFltFiringsUnit `json:"FltFiringsUnit,omitempty"`
	// The current number of sessions for this filter.
	FltCurSess uint64 `json:"FltCurSess,omitempty"`
	// The highest number of session for this filter.
	FltHighSess uint64 `json:"FltHighSess,omitempty"`
	// The total number of session for this filter.
	FltTotSess uint64 `json:"FltTotSess,omitempty"`
	// The current bandwidth for this filter in Kbps.
	FltCurBand uint64 `json:"FltCurBand,omitempty"`
	// The highest bandwidth for this filter in Kbps.
	FltHighBand uint64 `json:"FltHighBand,omitempty"`
	// The total bandwidth for this filter in Mb
	FltTotBand uint64 `json:"FltTotBand,omitempty"`
	// The connections per second for this filter
	FltConnPerSec uint64 `json:"FltConnPerSec,omitempty"`
	// The packets per second for this filter
	FltPktPerSec uint64 `json:"FltPktPerSec,omitempty"`
}

func (p FltStatTableParams) iMABean() {}
