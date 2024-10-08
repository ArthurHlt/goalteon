package beans

import (
	"fmt"
	"reflect"
)

// SlbNewAdvhcVirtualWireTable The table of Virtual Wire Health Check objects.
// Note:This mib is not supported on VX instance of Virtualization.
type SlbNewAdvhcVirtualWireTable struct {
	// Set Virtual Wire Health Check ID.
	SlbNewAdvhcVirtualWireID string
	Params                   *SlbNewAdvhcVirtualWireTableParams
}

func NewSlbNewAdvhcVirtualWireTableList() *SlbNewAdvhcVirtualWireTable {
	return &SlbNewAdvhcVirtualWireTable{}
}

func NewSlbNewAdvhcVirtualWireTable(
	slbNewAdvhcVirtualWireID string,
	params *SlbNewAdvhcVirtualWireTableParams,
) *SlbNewAdvhcVirtualWireTable {
	return &SlbNewAdvhcVirtualWireTable{
		SlbNewAdvhcVirtualWireID: slbNewAdvhcVirtualWireID,
		Params:                   params,
	}
}

func (c *SlbNewAdvhcVirtualWireTable) Name() string {
	return "SlbNewAdvhcVirtualWireTable"
}

func (c *SlbNewAdvhcVirtualWireTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbNewAdvhcVirtualWireTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbNewAdvhcVirtualWireTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbNewAdvhcVirtualWireID).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbNewAdvhcVirtualWireID)
}

type SlbNewAdvhcVirtualWireTableIPVer int32

const (
	SlbNewAdvhcVirtualWireTableIPVer_Ipv4 SlbNewAdvhcVirtualWireTableIPVer = 1
	SlbNewAdvhcVirtualWireTableIPVer_Ipv6 SlbNewAdvhcVirtualWireTableIPVer = 2
	SlbNewAdvhcVirtualWireTableIPVer_None SlbNewAdvhcVirtualWireTableIPVer = 3
)

type SlbNewAdvhcVirtualWireTableTransparent int32

const (
	SlbNewAdvhcVirtualWireTableTransparent_Enabled     SlbNewAdvhcVirtualWireTableTransparent = 1
	SlbNewAdvhcVirtualWireTableTransparent_Disabled    SlbNewAdvhcVirtualWireTableTransparent = 2
	SlbNewAdvhcVirtualWireTableTransparent_Unsupported SlbNewAdvhcVirtualWireTableTransparent = 2147483647
)

type SlbNewAdvhcVirtualWireTableInvert int32

const (
	SlbNewAdvhcVirtualWireTableInvert_Enabled     SlbNewAdvhcVirtualWireTableInvert = 1
	SlbNewAdvhcVirtualWireTableInvert_Disabled    SlbNewAdvhcVirtualWireTableInvert = 2
	SlbNewAdvhcVirtualWireTableInvert_Unsupported SlbNewAdvhcVirtualWireTableInvert = 2147483647
)

type SlbNewAdvhcVirtualWireTableDelete int32

const (
	SlbNewAdvhcVirtualWireTableDelete_Other       SlbNewAdvhcVirtualWireTableDelete = 1
	SlbNewAdvhcVirtualWireTableDelete_Delete      SlbNewAdvhcVirtualWireTableDelete = 2
	SlbNewAdvhcVirtualWireTableDelete_Unsupported SlbNewAdvhcVirtualWireTableDelete = 2147483647
)

type SlbNewAdvhcVirtualWireTableSnat int32

const (
	SlbNewAdvhcVirtualWireTableSnat_Enabled     SlbNewAdvhcVirtualWireTableSnat = 1
	SlbNewAdvhcVirtualWireTableSnat_Disabled    SlbNewAdvhcVirtualWireTableSnat = 2
	SlbNewAdvhcVirtualWireTableSnat_Unsupported SlbNewAdvhcVirtualWireTableSnat = 2147483647
)

type SlbNewAdvhcVirtualWireTableParams struct {
	// Set Virtual Wire Health Check ID.
	ID string `json:"ID,omitempty"`
	// Set Virtual Wire Health Check Name.
	Name string `json:"Name,omitempty"`
	// Set Virtual Wire Health Check destination port.
	DPort uint64 `json:"DPort,omitempty"`
	// Set Virtual Wire Health Check IP version.
	IPVer SlbNewAdvhcVirtualWireTableIPVer `json:"IPVer,omitempty"`
	// Set Virtual Wire Health Check host name.
	HostName string `json:"HostName,omitempty"`
	// Set Virtual Wire Health Check transparency mode.
	Transparent SlbNewAdvhcVirtualWireTableTransparent `json:"Transparent,omitempty"`
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
	Invert SlbNewAdvhcVirtualWireTableInvert `json:"Invert,omitempty"`
	// Copy Virtual Wire Health Check.
	Copy DisplayString `json:"Copy,omitempty"`
	// When set to the value of 2 (delete), the entire row is deleted.
	// When read, other(1) is returned. Setting the value to anything
	// other than 2(delete) has no effect on the state of the row.
	Delete SlbNewAdvhcVirtualWireTableDelete `json:"Delete,omitempty"`
	// Set Virtual Wire Health Check src NAT (PIP) flag.
	Snat SlbNewAdvhcVirtualWireTableSnat `json:"Snat,omitempty"`
}

func (p SlbNewAdvhcVirtualWireTableParams) iMABean() {}
