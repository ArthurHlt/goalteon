package beans

import (
	"fmt"
	"reflect"
)

// SlbCurAdvhcAdvVirtualWireTable The table of Advanced Virtual Wire Health Check objects,
// Note:This mib is not supported for VX instance of Virtualization.
type SlbCurAdvhcAdvVirtualWireTable struct {
	// Describes Virtual Wire Health Check ID.
	SlbCurAdvhcAdvVirtualWireID string
	Params                      *SlbCurAdvhcAdvVirtualWireTableParams
}

func NewSlbCurAdvhcAdvVirtualWireTableList() *SlbCurAdvhcAdvVirtualWireTable {
	return &SlbCurAdvhcAdvVirtualWireTable{}
}

func NewSlbCurAdvhcAdvVirtualWireTable(
	slbCurAdvhcAdvVirtualWireID string,
	params *SlbCurAdvhcAdvVirtualWireTableParams,
) *SlbCurAdvhcAdvVirtualWireTable {
	return &SlbCurAdvhcAdvVirtualWireTable{
		SlbCurAdvhcAdvVirtualWireID: slbCurAdvhcAdvVirtualWireID,
		Params:                      params,
	}
}

func (c *SlbCurAdvhcAdvVirtualWireTable) Name() string {
	return "SlbCurAdvhcAdvVirtualWireTable"
}

func (c *SlbCurAdvhcAdvVirtualWireTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbCurAdvhcAdvVirtualWireTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbCurAdvhcAdvVirtualWireTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbCurAdvhcAdvVirtualWireID).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbCurAdvhcAdvVirtualWireID)
}

type SlbCurAdvhcAdvVirtualWireTableIPVer int32

const (
	SlbCurAdvhcAdvVirtualWireTableIPVer_Ipv4 SlbCurAdvhcAdvVirtualWireTableIPVer = 1
	SlbCurAdvhcAdvVirtualWireTableIPVer_Ipv6 SlbCurAdvhcAdvVirtualWireTableIPVer = 2
	SlbCurAdvhcAdvVirtualWireTableIPVer_None SlbCurAdvhcAdvVirtualWireTableIPVer = 3
)

type SlbCurAdvhcAdvVirtualWireTableTransparent int32

const (
	SlbCurAdvhcAdvVirtualWireTableTransparent_Enabled     SlbCurAdvhcAdvVirtualWireTableTransparent = 1
	SlbCurAdvhcAdvVirtualWireTableTransparent_Disabled    SlbCurAdvhcAdvVirtualWireTableTransparent = 2
	SlbCurAdvhcAdvVirtualWireTableTransparent_Unsupported SlbCurAdvhcAdvVirtualWireTableTransparent = 2147483647
)

type SlbCurAdvhcAdvVirtualWireTableInvert int32

const (
	SlbCurAdvhcAdvVirtualWireTableInvert_Enabled     SlbCurAdvhcAdvVirtualWireTableInvert = 1
	SlbCurAdvhcAdvVirtualWireTableInvert_Disabled    SlbCurAdvhcAdvVirtualWireTableInvert = 2
	SlbCurAdvhcAdvVirtualWireTableInvert_Unsupported SlbCurAdvhcAdvVirtualWireTableInvert = 2147483647
)

type SlbCurAdvhcAdvVirtualWireTableSnat int32

const (
	SlbCurAdvhcAdvVirtualWireTableSnat_Enabled     SlbCurAdvhcAdvVirtualWireTableSnat = 1
	SlbCurAdvhcAdvVirtualWireTableSnat_Disabled    SlbCurAdvhcAdvVirtualWireTableSnat = 2
	SlbCurAdvhcAdvVirtualWireTableSnat_Unsupported SlbCurAdvhcAdvVirtualWireTableSnat = 2147483647
)

type SlbCurAdvhcAdvVirtualWireTableParams struct {
	// Describes Virtual Wire Health Check ID.
	ID string `json:"ID,omitempty"`
	// Describes Virtual Wire Health Check Name.
	Name string `json:"Name,omitempty"`
	// Describes Virtual Wire Health Check destination port.
	DPort uint64 `json:"DPort,omitempty"`
	// Describes Virtual Wire Health Check IP version.
	IPVer SlbCurAdvhcAdvVirtualWireTableIPVer `json:"IPVer,omitempty"`
	// Describes Virtual Wire Health Check host name.
	HostName string `json:"HostName,omitempty"`
	// Describes Virtual Wire Health Check transparency mode.
	Transparent SlbCurAdvhcAdvVirtualWireTableTransparent `json:"Transparent,omitempty"`
	// Describes Virtual Wire Health Check interval.
	Interval uint64 `json:"Interval,omitempty"`
	// Describes Virtual Wire Health Check maximum retries number.
	Retries uint64 `json:"Retries,omitempty"`
	// Describes Virtual Wire Health Check minimum restore retries number.
	RestoreRetries uint64 `json:"RestoreRetries,omitempty"`
	// Describes Virtual Wire Health Check timeout.
	Timeout uint64 `json:"Timeout,omitempty"`
	// Describes Virtual Wire Health Check Overflow.
	Overflow uint64 `json:"Overflow,omitempty"`
	// Describes Virtual Wire Health Check down interval.
	DownInterval uint64 `json:"DownInterval,omitempty"`
	// Describes Virtual Wire Health Check invert state.
	Invert SlbCurAdvhcAdvVirtualWireTableInvert `json:"Invert,omitempty"`
	// Describes Virtual Wire Health Check src NAT (PIP) flag.
	Snat SlbCurAdvhcAdvVirtualWireTableSnat `json:"Snat,omitempty"`
}

func (p SlbCurAdvhcAdvVirtualWireTableParams) iMABean() {}
