package beans

import (
	"fmt"
	"reflect"
)

// SlbNewAdvhcUdpTable Note:This mib is not supported on VX instance of Virtualization.
type SlbNewAdvhcUdpTable struct {
	// UDP Health check id.
	SlbNewAdvhcUdpID string
	Params           *SlbNewAdvhcUdpTableParams
}

func NewSlbNewAdvhcUdpTableList() *SlbNewAdvhcUdpTable {
	return &SlbNewAdvhcUdpTable{}
}

func NewSlbNewAdvhcUdpTable(
	slbNewAdvhcUdpID string,
	params *SlbNewAdvhcUdpTableParams,
) *SlbNewAdvhcUdpTable {
	return &SlbNewAdvhcUdpTable{
		SlbNewAdvhcUdpID: slbNewAdvhcUdpID,
		Params:           params,
	}
}

func (c *SlbNewAdvhcUdpTable) Name() string {
	return "SlbNewAdvhcUdpTable"
}

func (c *SlbNewAdvhcUdpTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbNewAdvhcUdpTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbNewAdvhcUdpTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbNewAdvhcUdpID).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbNewAdvhcUdpID)
}

type SlbNewAdvhcUdpTableIPVer int32

const (
	SlbNewAdvhcUdpTableIPVer_Ipv4 SlbNewAdvhcUdpTableIPVer = 1
	SlbNewAdvhcUdpTableIPVer_Ipv6 SlbNewAdvhcUdpTableIPVer = 2
	SlbNewAdvhcUdpTableIPVer_None SlbNewAdvhcUdpTableIPVer = 3
)

type SlbNewAdvhcUdpTableTransparent int32

const (
	SlbNewAdvhcUdpTableTransparent_Enabled     SlbNewAdvhcUdpTableTransparent = 1
	SlbNewAdvhcUdpTableTransparent_Disabled    SlbNewAdvhcUdpTableTransparent = 2
	SlbNewAdvhcUdpTableTransparent_Unsupported SlbNewAdvhcUdpTableTransparent = 2147483647
)

type SlbNewAdvhcUdpTableInvert int32

const (
	SlbNewAdvhcUdpTableInvert_Enabled     SlbNewAdvhcUdpTableInvert = 1
	SlbNewAdvhcUdpTableInvert_Disabled    SlbNewAdvhcUdpTableInvert = 2
	SlbNewAdvhcUdpTableInvert_Unsupported SlbNewAdvhcUdpTableInvert = 2147483647
)

type SlbNewAdvhcUdpTableDelete int32

const (
	SlbNewAdvhcUdpTableDelete_Other       SlbNewAdvhcUdpTableDelete = 1
	SlbNewAdvhcUdpTableDelete_Delete      SlbNewAdvhcUdpTableDelete = 2
	SlbNewAdvhcUdpTableDelete_Unsupported SlbNewAdvhcUdpTableDelete = 2147483647
)

type SlbNewAdvhcUdpTablePadding int32

const (
	SlbNewAdvhcUdpTablePadding_Enabled     SlbNewAdvhcUdpTablePadding = 1
	SlbNewAdvhcUdpTablePadding_Disabled    SlbNewAdvhcUdpTablePadding = 2
	SlbNewAdvhcUdpTablePadding_Unsupported SlbNewAdvhcUdpTablePadding = 2147483647
)

type SlbNewAdvhcUdpTableSnat int32

const (
	SlbNewAdvhcUdpTableSnat_Enabled     SlbNewAdvhcUdpTableSnat = 1
	SlbNewAdvhcUdpTableSnat_Disabled    SlbNewAdvhcUdpTableSnat = 2
	SlbNewAdvhcUdpTableSnat_Unsupported SlbNewAdvhcUdpTableSnat = 2147483647
)

type SlbNewAdvhcUdpTableParams struct {
	// UDP Health check id.
	ID string `json:"ID,omitempty"`
	// UDP Health check name.
	Name string `json:"Name,omitempty"`
	// UDP Health check destination port.
	DPort uint64 `json:"DPort,omitempty"`
	// UDP Health check destination IP version.
	IPVer SlbNewAdvhcUdpTableIPVer `json:"IPVer,omitempty"`
	// UDP Health check destination hostname.
	HostName string `json:"HostName,omitempty"`
	// UDP Health check transparent flag.
	Transparent SlbNewAdvhcUdpTableTransparent `json:"Transparent,omitempty"`
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
	Invert SlbNewAdvhcUdpTableInvert `json:"Invert,omitempty"`
	// UDP Health check copy indicator.
	Copy string `json:"Copy,omitempty"`
	// When set to the value of 2 (delete), the entire row is deleted.
	// When read, other(1) is returned. Setting the value to anything
	// other than 2(delete) has no effect on the state of the row.
	Delete SlbNewAdvhcUdpTableDelete `json:"Delete,omitempty"`
	// Enabled/disable padding to 64 bytes for the UDP health check.
	// When set to enable(1) UDP health check packets will be padded
	// to 64 bytes. And when set to disable(2) the UDP health check
	// packets will not be padded to 64 bytes.
	Padding SlbNewAdvhcUdpTablePadding `json:"Padding,omitempty"`
	// UDP Health Check src NAT (PIP) flag.
	Snat SlbNewAdvhcUdpTableSnat `json:"Snat,omitempty"`
}

func (p SlbNewAdvhcUdpTableParams) iMABean() {}
