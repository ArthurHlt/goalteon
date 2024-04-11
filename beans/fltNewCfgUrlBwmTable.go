package beans

import (
	"fmt"
	"reflect"
)

// FltNewCfgUrlBwmTable The table of URL based BWM for filtering.
// Note:This mib is not supported for VX instance of Virtualization.
type FltNewCfgUrlBwmTable struct {
	// The number of the filter.
	FltNewCfgUrlBwmFltIndex int32
	// The URL Path Identifier.
	FltNewCfgUrlBwmUrlId int32
	Params               *FltNewCfgUrlBwmTableParams
}

func NewFltNewCfgUrlBwmTableList() *FltNewCfgUrlBwmTable {
	return &FltNewCfgUrlBwmTable{}
}

func NewFltNewCfgUrlBwmTable(
	fltNewCfgUrlBwmFltIndex int32,
	fltNewCfgUrlBwmUrlId int32,
	params *FltNewCfgUrlBwmTableParams,
) *FltNewCfgUrlBwmTable {
	return &FltNewCfgUrlBwmTable{
		FltNewCfgUrlBwmFltIndex: fltNewCfgUrlBwmFltIndex,
		FltNewCfgUrlBwmUrlId:    fltNewCfgUrlBwmUrlId,
		Params:                  params,
	}
}

func (c *FltNewCfgUrlBwmTable) Name() string {
	return "FltNewCfgUrlBwmTable"
}

func (c *FltNewCfgUrlBwmTable) GetParams() BeanType {
	return c.Params
}

func (c *FltNewCfgUrlBwmTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *FltNewCfgUrlBwmTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.FltNewCfgUrlBwmFltIndex).IsZero() &&
		reflect.ValueOf(c.FltNewCfgUrlBwmUrlId).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.FltNewCfgUrlBwmFltIndex) + "/" + fmt.Sprint(c.FltNewCfgUrlBwmUrlId)
}

type FltNewCfgUrlBwmTableDelete int32

const (
	FltNewCfgUrlBwmTableDelete_Other       FltNewCfgUrlBwmTableDelete = 1
	FltNewCfgUrlBwmTableDelete_Delete      FltNewCfgUrlBwmTableDelete = 2
	FltNewCfgUrlBwmTableDelete_Unsupported FltNewCfgUrlBwmTableDelete = 2147483647
)

type FltNewCfgUrlBwmTableParams struct {
	// The number of the filter.
	FltIndex int32 `json:"FltIndex,omitempty"`
	// The URL Path Identifier.
	UrlId int32 `json:"UrlId,omitempty"`
	// The BW traffic contract.
	Contract int32 `json:"Contract,omitempty"`
	// When set to the value of 2 (delete), the entire row is deleted.
	// When read, other(1) is returned. Setting the value to anything
	// other than 2(delete) has no effect on the state of the row.
	Delete FltNewCfgUrlBwmTableDelete `json:"Delete,omitempty"`
	// The BWM contract for reverse traffic.
	ReverseBwmContract int32 `json:"ReverseBwmContract,omitempty"`
}

func (p FltNewCfgUrlBwmTableParams) iMABean() {}
