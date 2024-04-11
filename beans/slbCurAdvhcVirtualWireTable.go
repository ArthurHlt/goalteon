package beans

import (
	"fmt"
	"reflect"
)

// SlbCurAdvhcVirtualWireTable The table of Virtual Wire Health Check objects.
// Note:This mib is not supported on VX instance of Virtualization.
type SlbCurAdvhcVirtualWireTable struct {
	// Describes Virtual Wire Health Check ID.
	SlbCurAdvhcVirtualWireID string
	Params                   *SlbCurAdvhcVirtualWireTableParams
}

func NewSlbCurAdvhcVirtualWireTableList() *SlbCurAdvhcVirtualWireTable {
	return &SlbCurAdvhcVirtualWireTable{}
}

func NewSlbCurAdvhcVirtualWireTable(
	slbCurAdvhcVirtualWireID string,
	params *SlbCurAdvhcVirtualWireTableParams,
) *SlbCurAdvhcVirtualWireTable {
	return &SlbCurAdvhcVirtualWireTable{
		SlbCurAdvhcVirtualWireID: slbCurAdvhcVirtualWireID,
		Params:                   params,
	}
}

func (c *SlbCurAdvhcVirtualWireTable) Name() string {
	return "SlbCurAdvhcVirtualWireTable"
}

func (c *SlbCurAdvhcVirtualWireTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbCurAdvhcVirtualWireTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbCurAdvhcVirtualWireTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbCurAdvhcVirtualWireID).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbCurAdvhcVirtualWireID)
}

type SlbCurAdvhcVirtualWireTableIPVer int32

const (
	SlbCurAdvhcVirtualWireTableIPVer_Ipv4 SlbCurAdvhcVirtualWireTableIPVer = 1
	SlbCurAdvhcVirtualWireTableIPVer_Ipv6 SlbCurAdvhcVirtualWireTableIPVer = 2
	SlbCurAdvhcVirtualWireTableIPVer_None SlbCurAdvhcVirtualWireTableIPVer = 3
)

type SlbCurAdvhcVirtualWireTableTransparent int32

const (
	SlbCurAdvhcVirtualWireTableTransparent_Enabled     SlbCurAdvhcVirtualWireTableTransparent = 1
	SlbCurAdvhcVirtualWireTableTransparent_Disabled    SlbCurAdvhcVirtualWireTableTransparent = 2
	SlbCurAdvhcVirtualWireTableTransparent_Unsupported SlbCurAdvhcVirtualWireTableTransparent = 2147483647
)

type SlbCurAdvhcVirtualWireTableInvert int32

const (
	SlbCurAdvhcVirtualWireTableInvert_Enabled     SlbCurAdvhcVirtualWireTableInvert = 1
	SlbCurAdvhcVirtualWireTableInvert_Disabled    SlbCurAdvhcVirtualWireTableInvert = 2
	SlbCurAdvhcVirtualWireTableInvert_Unsupported SlbCurAdvhcVirtualWireTableInvert = 2147483647
)

type SlbCurAdvhcVirtualWireTableParams struct {
	// Describes Virtual Wire Health Check ID.
	ID string `json:"ID,omitempty"`
	// Describes Virtual Wire Health Check Name.
	Name string `json:"Name,omitempty"`
	// Describes Virtual Wire Health Check destination port.
	DPort uint64 `json:"DPort,omitempty"`
	// Describes Virtual Wire Health Check IP version.
	IPVer SlbCurAdvhcVirtualWireTableIPVer `json:"IPVer,omitempty"`
	// Describes Virtual Wire Health Check host name.
	HostName string `json:"HostName,omitempty"`
	// Describes Virtual Wire Health Check transparency mode.
	Transparent SlbCurAdvhcVirtualWireTableTransparent `json:"Transparent,omitempty"`
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
	Invert SlbCurAdvhcVirtualWireTableInvert `json:"Invert,omitempty"`
}

func (p SlbCurAdvhcVirtualWireTableParams) iMABean() {}
