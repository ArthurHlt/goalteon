package beans

import (
	"fmt"
	"reflect"
)

// SlbNewAdvhcAdvVirtualWireTable The table of Advanced Virtual Wire Health Check objects,
// Note:This mib is not supported for VX instance of Virtualization.
type SlbNewAdvhcAdvVirtualWireTable struct {
	// Set Virtual Wire Health Check ID.
	SlbNewAdvhcAdvVirtualWireID string
	Params                      *SlbNewAdvhcAdvVirtualWireTableParams
}

func NewSlbNewAdvhcAdvVirtualWireTableList() *SlbNewAdvhcAdvVirtualWireTable {
	return &SlbNewAdvhcAdvVirtualWireTable{}
}

func NewSlbNewAdvhcAdvVirtualWireTable(
	slbNewAdvhcAdvVirtualWireID string,
	params *SlbNewAdvhcAdvVirtualWireTableParams,
) *SlbNewAdvhcAdvVirtualWireTable {
	return &SlbNewAdvhcAdvVirtualWireTable{
		SlbNewAdvhcAdvVirtualWireID: slbNewAdvhcAdvVirtualWireID,
		Params:                      params,
	}
}

func (c *SlbNewAdvhcAdvVirtualWireTable) Name() string {
	return "SlbNewAdvhcAdvVirtualWireTable"
}

func (c *SlbNewAdvhcAdvVirtualWireTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbNewAdvhcAdvVirtualWireTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbNewAdvhcAdvVirtualWireTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbNewAdvhcAdvVirtualWireID).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbNewAdvhcAdvVirtualWireID)
}

type SlbNewAdvhcAdvVirtualWireTableIPVer int32

const (
	SlbNewAdvhcAdvVirtualWireTableIPVer_Ipv4 SlbNewAdvhcAdvVirtualWireTableIPVer = 1
	SlbNewAdvhcAdvVirtualWireTableIPVer_Ipv6 SlbNewAdvhcAdvVirtualWireTableIPVer = 2
	SlbNewAdvhcAdvVirtualWireTableIPVer_None SlbNewAdvhcAdvVirtualWireTableIPVer = 3
)

type SlbNewAdvhcAdvVirtualWireTableTransparent int32

const (
	SlbNewAdvhcAdvVirtualWireTableTransparent_Enabled     SlbNewAdvhcAdvVirtualWireTableTransparent = 1
	SlbNewAdvhcAdvVirtualWireTableTransparent_Disabled    SlbNewAdvhcAdvVirtualWireTableTransparent = 2
	SlbNewAdvhcAdvVirtualWireTableTransparent_Unsupported SlbNewAdvhcAdvVirtualWireTableTransparent = 2147483647
)

type SlbNewAdvhcAdvVirtualWireTableInvert int32

const (
	SlbNewAdvhcAdvVirtualWireTableInvert_Enabled     SlbNewAdvhcAdvVirtualWireTableInvert = 1
	SlbNewAdvhcAdvVirtualWireTableInvert_Disabled    SlbNewAdvhcAdvVirtualWireTableInvert = 2
	SlbNewAdvhcAdvVirtualWireTableInvert_Unsupported SlbNewAdvhcAdvVirtualWireTableInvert = 2147483647
)

type SlbNewAdvhcAdvVirtualWireTableDelete int32

const (
	SlbNewAdvhcAdvVirtualWireTableDelete_Other       SlbNewAdvhcAdvVirtualWireTableDelete = 1
	SlbNewAdvhcAdvVirtualWireTableDelete_Delete      SlbNewAdvhcAdvVirtualWireTableDelete = 2
	SlbNewAdvhcAdvVirtualWireTableDelete_Unsupported SlbNewAdvhcAdvVirtualWireTableDelete = 2147483647
)

type SlbNewAdvhcAdvVirtualWireTableParams struct {
	// Set Virtual Wire Health Check ID.
	ID string `json:"ID,omitempty"`
	// Set Virtual Wire Health Check Name.
	Name string `json:"Name,omitempty"`
	// Set Virtual Wire Health Check destination port.
	DPort uint64 `json:"DPort,omitempty"`
	// Set Virtual Wire Health Check IP version.
	IPVer SlbNewAdvhcAdvVirtualWireTableIPVer `json:"IPVer,omitempty"`
	// Set Virtual Wire Health Check host name.
	HostName string `json:"HostName,omitempty"`
	// Set Virtual Wire Health Check transparency mode.
	Transparent SlbNewAdvhcAdvVirtualWireTableTransparent `json:"Transparent,omitempty"`
	// Set Virtual Wire Health Check interval.
	Interval uint64 `json:"Interval,omitempty"`
	// Set Virtual Wire Health Check maximum retries number.
	Retries uint64 `json:"Retries,omitempty"`
	// Set Virtual Wire Health Check minimum restore retries number.
	RestoreRetries uint64 `json:"RestoreRetries,omitempty"`
	// Set Virtual Wire Health Check timeout.
	Timeout uint64 `json:"Timeout,omitempty"`
	// Set Virtual Wire Health Check Overflow.
	Overflow uint64 `json:"Overflow,omitempty"`
	// Set Virtual Wire Health Check down interval.
	DownInterval uint64 `json:"DownInterval,omitempty"`
	// Set Virtual Wire Health Check invert state.
	Invert SlbNewAdvhcAdvVirtualWireTableInvert `json:"Invert,omitempty"`
	// Copy Virtual Wire Health Check.
	Copy string `json:"Copy,omitempty"`
	// When set to the value of 2 (delete), the entire row is deleted.
	// When read, other(1) is returned. Setting the value to anything
	// other than 2(delete) has no effect on the state of the row.
	Delete SlbNewAdvhcAdvVirtualWireTableDelete `json:"Delete,omitempty"`
}

func (p SlbNewAdvhcAdvVirtualWireTableParams) iMABean() {}
