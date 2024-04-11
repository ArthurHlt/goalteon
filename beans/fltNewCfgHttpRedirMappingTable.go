package beans

import (
	"fmt"
	"reflect"
)

// FltNewCfgHttpRedirMappingTable The table of HTTP redirection mapping table.
// Note:This mib is not supported for VX instance of Virtualization.
type FltNewCfgHttpRedirMappingTable struct {
	// The filter number.
	FltNewCfgHttpRedirMappingFilter int32
	// The string ID of Layer7 string table. If HTTP request
	// matches this string then switch issues a HTTP REDIRECT
	// to client with the fltNewCfgHttpRedirMappingToStr string.
	FltNewCfgHttpRedirMappingFromStr uint32
	Params                           *FltNewCfgHttpRedirMappingTableParams
}

func NewFltNewCfgHttpRedirMappingTableList() *FltNewCfgHttpRedirMappingTable {
	return &FltNewCfgHttpRedirMappingTable{}
}

func NewFltNewCfgHttpRedirMappingTable(
	fltNewCfgHttpRedirMappingFilter int32,
	fltNewCfgHttpRedirMappingFromStr uint32,
	params *FltNewCfgHttpRedirMappingTableParams,
) *FltNewCfgHttpRedirMappingTable {
	return &FltNewCfgHttpRedirMappingTable{
		FltNewCfgHttpRedirMappingFilter:  fltNewCfgHttpRedirMappingFilter,
		FltNewCfgHttpRedirMappingFromStr: fltNewCfgHttpRedirMappingFromStr,
		Params:                           params,
	}
}

func (c *FltNewCfgHttpRedirMappingTable) Name() string {
	return "FltNewCfgHttpRedirMappingTable"
}

func (c *FltNewCfgHttpRedirMappingTable) GetParams() BeanType {
	return c.Params
}

func (c *FltNewCfgHttpRedirMappingTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *FltNewCfgHttpRedirMappingTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.FltNewCfgHttpRedirMappingFilter).IsZero() &&
		reflect.ValueOf(c.FltNewCfgHttpRedirMappingFromStr).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.FltNewCfgHttpRedirMappingFilter) + "/" + fmt.Sprint(c.FltNewCfgHttpRedirMappingFromStr)
}

type FltNewCfgHttpRedirMappingTableDelete int32

const (
	FltNewCfgHttpRedirMappingTableDelete_Other       FltNewCfgHttpRedirMappingTableDelete = 1
	FltNewCfgHttpRedirMappingTableDelete_Delete      FltNewCfgHttpRedirMappingTableDelete = 2
	FltNewCfgHttpRedirMappingTableDelete_Unsupported FltNewCfgHttpRedirMappingTableDelete = 2147483647
)

type FltNewCfgHttpRedirMappingTableParams struct {
	// The filter number.
	Filter int32 `json:"Filter,omitempty"`
	// The string ID of Layer7 string table. If HTTP request
	// matches this string then switch issues a HTTP REDIRECT
	// to client with the fltNewCfgHttpRedirMappingToStr string.
	FromStr uint32 `json:"FromStr,omitempty"`
	// The string ID of Layer7 string table. The switch sends
	// HTTP Redirect back to client with this string when the
	// HTTP request matches fltNewCfgHttpRedirMappingFromStr
	// string.
	ToStr uint32 `json:"ToStr,omitempty"`
	// By setting the value to delete(2), the entire row is deleted.
	Delete FltNewCfgHttpRedirMappingTableDelete `json:"Delete,omitempty"`
}

func (p FltNewCfgHttpRedirMappingTableParams) iMABean() {}
