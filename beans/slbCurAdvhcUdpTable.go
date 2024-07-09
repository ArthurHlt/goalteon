package beans

import (
	"fmt"
	"reflect"
)

// SlbCurAdvhcUdpTable Note:This mib is not supported on VX instance of Virtualization.
type SlbCurAdvhcUdpTable struct {
	// UDP Health check id.
	SlbCurAdvhcUdpID string
	Params           *SlbCurAdvhcUdpTableParams
}

func NewSlbCurAdvhcUdpTableList() *SlbCurAdvhcUdpTable {
	return &SlbCurAdvhcUdpTable{}
}

func NewSlbCurAdvhcUdpTable(
	slbCurAdvhcUdpID string,
	params *SlbCurAdvhcUdpTableParams,
) *SlbCurAdvhcUdpTable {
	return &SlbCurAdvhcUdpTable{
		SlbCurAdvhcUdpID: slbCurAdvhcUdpID,
		Params:           params,
	}
}

func (c *SlbCurAdvhcUdpTable) Name() string {
	return "SlbCurAdvhcUdpTable"
}

func (c *SlbCurAdvhcUdpTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbCurAdvhcUdpTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbCurAdvhcUdpTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbCurAdvhcUdpID).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbCurAdvhcUdpID)
}

type SlbCurAdvhcUdpTableIPVer int32

const (
	SlbCurAdvhcUdpTableIPVer_Ipv4 SlbCurAdvhcUdpTableIPVer = 1
	SlbCurAdvhcUdpTableIPVer_Ipv6 SlbCurAdvhcUdpTableIPVer = 2
	SlbCurAdvhcUdpTableIPVer_None SlbCurAdvhcUdpTableIPVer = 3
)

type SlbCurAdvhcUdpTableTransparent int32

const (
	SlbCurAdvhcUdpTableTransparent_Enabled     SlbCurAdvhcUdpTableTransparent = 1
	SlbCurAdvhcUdpTableTransparent_Disabled    SlbCurAdvhcUdpTableTransparent = 2
	SlbCurAdvhcUdpTableTransparent_Unsupported SlbCurAdvhcUdpTableTransparent = 2147483647
)

type SlbCurAdvhcUdpTableInvert int32

const (
	SlbCurAdvhcUdpTableInvert_Enabled     SlbCurAdvhcUdpTableInvert = 1
	SlbCurAdvhcUdpTableInvert_Disabled    SlbCurAdvhcUdpTableInvert = 2
	SlbCurAdvhcUdpTableInvert_Unsupported SlbCurAdvhcUdpTableInvert = 2147483647
)

type SlbCurAdvhcUdpTablePadding int32

const (
	SlbCurAdvhcUdpTablePadding_Enabled     SlbCurAdvhcUdpTablePadding = 1
	SlbCurAdvhcUdpTablePadding_Disabled    SlbCurAdvhcUdpTablePadding = 2
	SlbCurAdvhcUdpTablePadding_Unsupported SlbCurAdvhcUdpTablePadding = 2147483647
)

type SlbCurAdvhcUdpTableSnat int32

const (
	SlbCurAdvhcUdpTableSnat_Enabled     SlbCurAdvhcUdpTableSnat = 1
	SlbCurAdvhcUdpTableSnat_Disabled    SlbCurAdvhcUdpTableSnat = 2
	SlbCurAdvhcUdpTableSnat_Unsupported SlbCurAdvhcUdpTableSnat = 2147483647
)

type SlbCurAdvhcUdpTableParams struct {
	// UDP Health check id.
	ID string `json:"ID,omitempty"`
	// UDP Health check name.
	Name string `json:"Name,omitempty"`
	// UDP Health check destination port.
	DPort uint64 `json:"DPort,omitempty"`
	// UDP Health check destination IP version.
	IPVer SlbCurAdvhcUdpTableIPVer `json:"IPVer,omitempty"`
	// UDP Health check destination hostname.
	HostName string `json:"HostName,omitempty"`
	// UDP Health check transparent flag.
	Transparent SlbCurAdvhcUdpTableTransparent `json:"Transparent,omitempty"`
	// UDP Health check interval.
	Interval uint64 `json:"Interval,omitempty"`
	// UDP Health check retries counter.
	Retries uint64 `json:"Retries,omitempty"`
	// UDP Health check retries counter in the down state.
	RestoreRetries uint64 `json:"RestoreRetries,omitempty"`
	// UDP Health check timeout.
	Timeout uint64 `json:"Timeout,omitempty"`
	// UDP Health check overflow flag.
	Overflow uint64 `json:"Overflow,omitempty"`
	// UDP Health check interval in the down state.
	DownInterval uint64 `json:"DownInterval,omitempty"`
	// UDP Health check invert flag.
	Invert SlbCurAdvhcUdpTableInvert `json:"Invert,omitempty"`
	// Current status of 'padding' setting for the UDP health check.
	// This setting specifies if padding to 64 bytes should be done
	// for the UDP health check packet.
	Padding SlbCurAdvhcUdpTablePadding `json:"Padding,omitempty"`
	// UDP Health Check src NAT (PIP) flag.
	Snat SlbCurAdvhcUdpTableSnat `json:"Snat,omitempty"`
}

func (p SlbCurAdvhcUdpTableParams) iMABean() {}
