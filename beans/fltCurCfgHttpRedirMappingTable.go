package beans

import (
	"fmt"
	"reflect"
)

// FltCurCfgHttpRedirMappingTable The table of HTTP redirection mapping table.
// Note:This mib is not supported for VX instance of Virtualization.
type FltCurCfgHttpRedirMappingTable struct {
	// The filter number.
	FltCurCfgHttpRedirMappingFilter int32
	// This index is the string ID of Layer7 string table.
	// If HTTP request matches this string, then switch issues
	// a HTTP REDIRECT to client with the
	// fltCurCfgHttpRedirMappingToStr string.
	FltCurCfgHttpRedirMappingFromStr uint32
	Params                           *FltCurCfgHttpRedirMappingTableParams
}

func NewFltCurCfgHttpRedirMappingTableList() *FltCurCfgHttpRedirMappingTable {
	return &FltCurCfgHttpRedirMappingTable{}
}

func NewFltCurCfgHttpRedirMappingTable(
	fltCurCfgHttpRedirMappingFilter int32,
	fltCurCfgHttpRedirMappingFromStr uint32,
	params *FltCurCfgHttpRedirMappingTableParams,
) *FltCurCfgHttpRedirMappingTable {
	return &FltCurCfgHttpRedirMappingTable{
		FltCurCfgHttpRedirMappingFilter:  fltCurCfgHttpRedirMappingFilter,
		FltCurCfgHttpRedirMappingFromStr: fltCurCfgHttpRedirMappingFromStr,
		Params:                           params,
	}
}

func (c *FltCurCfgHttpRedirMappingTable) Name() string {
	return "FltCurCfgHttpRedirMappingTable"
}

func (c *FltCurCfgHttpRedirMappingTable) GetParams() BeanType {
	return c.Params
}

func (c *FltCurCfgHttpRedirMappingTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *FltCurCfgHttpRedirMappingTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.FltCurCfgHttpRedirMappingFilter).IsZero() &&
		reflect.ValueOf(c.FltCurCfgHttpRedirMappingFromStr).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.FltCurCfgHttpRedirMappingFilter) + "/" + fmt.Sprint(c.FltCurCfgHttpRedirMappingFromStr)
}

type FltCurCfgHttpRedirMappingTableParams struct {
	// The filter number.
	Filter int32 `json:"Filter,omitempty"`
	// This index is the string ID of Layer7 string table.
	// If HTTP request matches this string, then switch issues
	// a HTTP REDIRECT to client with the
	// fltCurCfgHttpRedirMappingToStr string.
	FromStr uint32 `json:"FromStr,omitempty"`
	// The string ID of Layer7 string table. The switch sends
	// HTTP Redirect back to client with this string when HTTP
	// request matches the fltCurCfgHttpRedirMappingFromStr string
	ToStr uint32 `json:"ToStr,omitempty"`
}

func (p FltCurCfgHttpRedirMappingTableParams) iMABean() {}
