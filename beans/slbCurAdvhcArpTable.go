package beans

import (
	"fmt"
	"reflect"
)

// SlbCurAdvhcArpTable Note:This mib is not supported on VX instance of Virtualization.
type SlbCurAdvhcArpTable struct {
	// ARP Health check id.
	SlbCurAdvhcArpID string
	Params           *SlbCurAdvhcArpTableParams
}

func NewSlbCurAdvhcArpTableList() *SlbCurAdvhcArpTable {
	return &SlbCurAdvhcArpTable{}
}

func NewSlbCurAdvhcArpTable(
	slbCurAdvhcArpID string,
	params *SlbCurAdvhcArpTableParams,
) *SlbCurAdvhcArpTable {
	return &SlbCurAdvhcArpTable{
		SlbCurAdvhcArpID: slbCurAdvhcArpID,
		Params:           params,
	}
}

func (c *SlbCurAdvhcArpTable) Name() string {
	return "SlbCurAdvhcArpTable"
}

func (c *SlbCurAdvhcArpTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbCurAdvhcArpTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbCurAdvhcArpTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbCurAdvhcArpID).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbCurAdvhcArpID)
}

type SlbCurAdvhcArpTableIPVer int32

const (
	SlbCurAdvhcArpTableIPVer_Ipv4 SlbCurAdvhcArpTableIPVer = 1
	SlbCurAdvhcArpTableIPVer_Ipv6 SlbCurAdvhcArpTableIPVer = 2
	SlbCurAdvhcArpTableIPVer_None SlbCurAdvhcArpTableIPVer = 3
)

type SlbCurAdvhcArpTableTransparent int32

const (
	SlbCurAdvhcArpTableTransparent_Enabled     SlbCurAdvhcArpTableTransparent = 1
	SlbCurAdvhcArpTableTransparent_Disabled    SlbCurAdvhcArpTableTransparent = 2
	SlbCurAdvhcArpTableTransparent_Unsupported SlbCurAdvhcArpTableTransparent = 2147483647
)

type SlbCurAdvhcArpTableInvert int32

const (
	SlbCurAdvhcArpTableInvert_Enabled     SlbCurAdvhcArpTableInvert = 1
	SlbCurAdvhcArpTableInvert_Disabled    SlbCurAdvhcArpTableInvert = 2
	SlbCurAdvhcArpTableInvert_Unsupported SlbCurAdvhcArpTableInvert = 2147483647
)

type SlbCurAdvhcArpTableSnat int32

const (
	SlbCurAdvhcArpTableSnat_Enabled     SlbCurAdvhcArpTableSnat = 1
	SlbCurAdvhcArpTableSnat_Disabled    SlbCurAdvhcArpTableSnat = 2
	SlbCurAdvhcArpTableSnat_Unsupported SlbCurAdvhcArpTableSnat = 2147483647
)

type SlbCurAdvhcArpTableParams struct {
	// ARP Health check id.
	ID string `json:"ID,omitempty"`
	// ARP Health check name.
	Name string `json:"Name,omitempty"`
	// ARP Health check IP version.
	IPVer SlbCurAdvhcArpTableIPVer `json:"IPVer,omitempty"`
	// ARP health check hostname.
	HostName string `json:"HostName,omitempty"`
	// Is ARP health check transparent.
	Transparent SlbCurAdvhcArpTableTransparent `json:"Transparent,omitempty"`
	// ARP Health check interval.
	Interval uint64 `json:"Interval,omitempty"`
	// ARP Health check retries before failure counter.
	Retries uint64 `json:"Retries,omitempty"`
	// ARP Health check retries before marking health check up, counter.
	RestoreRetries uint64 `json:"RestoreRetries,omitempty"`
	// ARP Health checheck timeout parameter.
	Timeout uint64 `json:"Timeout,omitempty"`
	// ARP Health check overflow parameter.
	Overflow uint64 `json:"Overflow,omitempty"`
	// ARP Health check checks Interval on Down-Time.
	DownInterval uint64 `json:"DownInterval,omitempty"`
	// ARP Health check invert flag.
	Invert SlbCurAdvhcArpTableInvert `json:"Invert,omitempty"`
	// ARP Health Check src NAT (PIP) flag.
	Snat SlbCurAdvhcArpTableSnat `json:"Snat,omitempty"`
}

func (p SlbCurAdvhcArpTableParams) iMABean() {}
