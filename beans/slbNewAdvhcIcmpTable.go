package beans

import (
	"fmt"
	"reflect"
)

// SlbNewAdvhcIcmpTable Note:This mib is not supported on VX instance of Virtualization.
type SlbNewAdvhcIcmpTable struct {
	// ICMP health check id.
	SlbNewAdvhcIcmpID string
	Params            *SlbNewAdvhcIcmpTableParams
}

func NewSlbNewAdvhcIcmpTableList() *SlbNewAdvhcIcmpTable {
	return &SlbNewAdvhcIcmpTable{}
}

func NewSlbNewAdvhcIcmpTable(
	slbNewAdvhcIcmpID string,
	params *SlbNewAdvhcIcmpTableParams,
) *SlbNewAdvhcIcmpTable {
	return &SlbNewAdvhcIcmpTable{
		SlbNewAdvhcIcmpID: slbNewAdvhcIcmpID,
		Params:            params,
	}
}

func (c *SlbNewAdvhcIcmpTable) Name() string {
	return "SlbNewAdvhcIcmpTable"
}

func (c *SlbNewAdvhcIcmpTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbNewAdvhcIcmpTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbNewAdvhcIcmpTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbNewAdvhcIcmpID).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbNewAdvhcIcmpID)
}

type SlbNewAdvhcIcmpTableIPVer int32

const (
	SlbNewAdvhcIcmpTableIPVer_Ipv4 SlbNewAdvhcIcmpTableIPVer = 1
	SlbNewAdvhcIcmpTableIPVer_Ipv6 SlbNewAdvhcIcmpTableIPVer = 2
	SlbNewAdvhcIcmpTableIPVer_None SlbNewAdvhcIcmpTableIPVer = 3
)

type SlbNewAdvhcIcmpTableTransparent int32

const (
	SlbNewAdvhcIcmpTableTransparent_Enabled     SlbNewAdvhcIcmpTableTransparent = 1
	SlbNewAdvhcIcmpTableTransparent_Disabled    SlbNewAdvhcIcmpTableTransparent = 2
	SlbNewAdvhcIcmpTableTransparent_Unsupported SlbNewAdvhcIcmpTableTransparent = 2147483647
)

type SlbNewAdvhcIcmpTableInvert int32

const (
	SlbNewAdvhcIcmpTableInvert_Enabled     SlbNewAdvhcIcmpTableInvert = 1
	SlbNewAdvhcIcmpTableInvert_Disabled    SlbNewAdvhcIcmpTableInvert = 2
	SlbNewAdvhcIcmpTableInvert_Unsupported SlbNewAdvhcIcmpTableInvert = 2147483647
)

type SlbNewAdvhcIcmpTableDelete int32

const (
	SlbNewAdvhcIcmpTableDelete_Other       SlbNewAdvhcIcmpTableDelete = 1
	SlbNewAdvhcIcmpTableDelete_Delete      SlbNewAdvhcIcmpTableDelete = 2
	SlbNewAdvhcIcmpTableDelete_Unsupported SlbNewAdvhcIcmpTableDelete = 2147483647
)

type SlbNewAdvhcIcmpTableSnat int32

const (
	SlbNewAdvhcIcmpTableSnat_Enabled     SlbNewAdvhcIcmpTableSnat = 1
	SlbNewAdvhcIcmpTableSnat_Disabled    SlbNewAdvhcIcmpTableSnat = 2
	SlbNewAdvhcIcmpTableSnat_Unsupported SlbNewAdvhcIcmpTableSnat = 2147483647
)

type SlbNewAdvhcIcmpTableParams struct {
	// ICMP health check id.
	ID string `json:"ID,omitempty"`
	// ICMP health check name.
	Name string `json:"Name,omitempty"`
	// ICMP health check destination port.
	DPort uint64 `json:"DPort,omitempty"`
	// ICMP health check destination IP version.
	IPVer SlbNewAdvhcIcmpTableIPVer `json:"IPVer,omitempty"`
	// ICMP health check destination hostname.
	HostName string `json:"HostName,omitempty"`
	// ICMP health check transparent flag.
	Transparent SlbNewAdvhcIcmpTableTransparent `json:"Transparent,omitempty"`
	// ICMP health check interval.
	Interval uint64 `json:"Interval,omitempty"`
	// ICMP health check retries counter.
	Retries uint64 `json:"Retries,omitempty"`
	// ICMP health check retries counter to restore from the down state.
	RestoreRetries uint64 `json:"RestoreRetries,omitempty"`
	// ICMP health check timeout.
	Timeout uint64 `json:"Timeout,omitempty"`
	// ICMP health check overflow flag.
	Overflow uint64 `json:"Overflow,omitempty"`
	// ICMP health check interval in the down state.
	DownInterval uint64 `json:"DownInterval,omitempty"`
	// ICMP health check invert flag.
	Invert SlbNewAdvhcIcmpTableInvert `json:"Invert,omitempty"`
	// ICMP health check copy indicator.
	Copy DisplayString `json:"Copy,omitempty"`
	// When set to the value of 2 (delete), the entire row is deleted.
	// When read, other(1) is returned. Setting the value to anything
	// other than 2(delete) has no effect on the state of the row.
	Delete SlbNewAdvhcIcmpTableDelete `json:"Delete,omitempty"`
	// ICMP Health Check src NAT (PIP) flag.
	Snat SlbNewAdvhcIcmpTableSnat `json:"Snat,omitempty"`
}

func (p SlbNewAdvhcIcmpTableParams) iMABean() {}
