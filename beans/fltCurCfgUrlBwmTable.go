package beans

import (
	"fmt"
	"reflect"
)

// FltCurCfgUrlBwmTable The table of URL based BWM for filtering.
// Note:This mib is not supported for VX instance of Virtualization.
type FltCurCfgUrlBwmTable struct {
	// The number of the filter.
	FltCurCfgUrlBwmFltIndex int32
	// The URL Path Identifier.
	FltCurCfgUrlBwmUrlId int32
	Params               *FltCurCfgUrlBwmTableParams
}

func NewFltCurCfgUrlBwmTableList() *FltCurCfgUrlBwmTable {
	return &FltCurCfgUrlBwmTable{}
}

func NewFltCurCfgUrlBwmTable(
	fltCurCfgUrlBwmFltIndex int32,
	fltCurCfgUrlBwmUrlId int32,
	params *FltCurCfgUrlBwmTableParams,
) *FltCurCfgUrlBwmTable {
	return &FltCurCfgUrlBwmTable{
		FltCurCfgUrlBwmFltIndex: fltCurCfgUrlBwmFltIndex,
		FltCurCfgUrlBwmUrlId:    fltCurCfgUrlBwmUrlId,
		Params:                  params,
	}
}

func (c *FltCurCfgUrlBwmTable) Name() string {
	return "FltCurCfgUrlBwmTable"
}

func (c *FltCurCfgUrlBwmTable) GetParams() BeanType {
	return c.Params
}

func (c *FltCurCfgUrlBwmTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *FltCurCfgUrlBwmTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.FltCurCfgUrlBwmFltIndex).IsZero() &&
		reflect.ValueOf(c.FltCurCfgUrlBwmUrlId).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.FltCurCfgUrlBwmFltIndex) + "/" + fmt.Sprint(c.FltCurCfgUrlBwmUrlId)
}

type FltCurCfgUrlBwmTableParams struct {
	// The number of the filter.
	FltIndex int32 `json:"FltIndex,omitempty"`
	// The URL Path Identifier.
	UrlId int32 `json:"UrlId,omitempty"`
	// The BW contract.
	Contract int32 `json:"Contract,omitempty"`
	// The BWM contract for reverse traffic.
	ReverseBwmContract int32 `json:"ReverseBwmContract,omitempty"`
}

func (p FltCurCfgUrlBwmTableParams) iMABean() {}
