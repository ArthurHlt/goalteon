package beans

import (
	"fmt"
	"reflect"
)

// SlbNewAdvhcArpTable Note:This mib is not supported on VX instance of Virtualization.
type SlbNewAdvhcArpTable struct {
	// ARP Health check id.
	SlbNewAdvhcArpID string
	Params           *SlbNewAdvhcArpTableParams
}

func NewSlbNewAdvhcArpTableList() *SlbNewAdvhcArpTable {
	return &SlbNewAdvhcArpTable{}
}

func NewSlbNewAdvhcArpTable(
	slbNewAdvhcArpID string,
	params *SlbNewAdvhcArpTableParams,
) *SlbNewAdvhcArpTable {
	return &SlbNewAdvhcArpTable{
		SlbNewAdvhcArpID: slbNewAdvhcArpID,
		Params:           params,
	}
}

func (c *SlbNewAdvhcArpTable) Name() string {
	return "SlbNewAdvhcArpTable"
}

func (c *SlbNewAdvhcArpTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbNewAdvhcArpTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbNewAdvhcArpTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbNewAdvhcArpID).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbNewAdvhcArpID)
}

type SlbNewAdvhcArpTableIPVer int32

const (
	SlbNewAdvhcArpTableIPVer_Ipv4 SlbNewAdvhcArpTableIPVer = 1
	SlbNewAdvhcArpTableIPVer_Ipv6 SlbNewAdvhcArpTableIPVer = 2
	SlbNewAdvhcArpTableIPVer_None SlbNewAdvhcArpTableIPVer = 3
)

type SlbNewAdvhcArpTableTransparent int32

const (
	SlbNewAdvhcArpTableTransparent_Enabled     SlbNewAdvhcArpTableTransparent = 1
	SlbNewAdvhcArpTableTransparent_Disabled    SlbNewAdvhcArpTableTransparent = 2
	SlbNewAdvhcArpTableTransparent_Unsupported SlbNewAdvhcArpTableTransparent = 2147483647
)

type SlbNewAdvhcArpTableInvert int32

const (
	SlbNewAdvhcArpTableInvert_Enabled     SlbNewAdvhcArpTableInvert = 1
	SlbNewAdvhcArpTableInvert_Disabled    SlbNewAdvhcArpTableInvert = 2
	SlbNewAdvhcArpTableInvert_Unsupported SlbNewAdvhcArpTableInvert = 2147483647
)

type SlbNewAdvhcArpTableDelete int32

const (
	SlbNewAdvhcArpTableDelete_Other       SlbNewAdvhcArpTableDelete = 1
	SlbNewAdvhcArpTableDelete_Delete      SlbNewAdvhcArpTableDelete = 2
	SlbNewAdvhcArpTableDelete_Unsupported SlbNewAdvhcArpTableDelete = 2147483647
)

type SlbNewAdvhcArpTableSnat int32

const (
	SlbNewAdvhcArpTableSnat_Enabled     SlbNewAdvhcArpTableSnat = 1
	SlbNewAdvhcArpTableSnat_Disabled    SlbNewAdvhcArpTableSnat = 2
	SlbNewAdvhcArpTableSnat_Unsupported SlbNewAdvhcArpTableSnat = 2147483647
)

type SlbNewAdvhcArpTableParams struct {
	// ARP Health check id.
	ID string `json:"ID,omitempty"`
	// ARP Health check name.
	Name string `json:"Name,omitempty"`
	// ARP Health check IP version.
	IPVer SlbNewAdvhcArpTableIPVer `json:"IPVer,omitempty"`
	// ARP health check hostname.
	HostName string `json:"HostName,omitempty"`
	// Is ARP health check transparent.
	Transparent SlbNewAdvhcArpTableTransparent `json:"Transparent,omitempty"`
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
	Invert SlbNewAdvhcArpTableInvert `json:"Invert,omitempty"`
	// ARP Health check copy indicator.
	Copy string `json:"Copy,omitempty"`
	// When set to the value of 2 (delete), the entire row is deleted.
	// When read, other(1) is returned. Setting the value to anything
	// other than 2(delete) has no effect on the state of the row.
	Delete SlbNewAdvhcArpTableDelete `json:"Delete,omitempty"`
	// ARP Health Check src NAT (PIP) flag.
	Snat SlbNewAdvhcArpTableSnat `json:"Snat,omitempty"`
}

func (p SlbNewAdvhcArpTableParams) iMABean() {}
