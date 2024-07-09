package beans

import (
	"fmt"
	"reflect"
)

// SlbCurAdvhcIcmpTable Note:This mib is not supported on VX instance of Virtualization.
type SlbCurAdvhcIcmpTable struct {
	// ICMP health check id.
	SlbCurAdvhcIcmpID string
	Params            *SlbCurAdvhcIcmpTableParams
}

func NewSlbCurAdvhcIcmpTableList() *SlbCurAdvhcIcmpTable {
	return &SlbCurAdvhcIcmpTable{}
}

func NewSlbCurAdvhcIcmpTable(
	slbCurAdvhcIcmpID string,
	params *SlbCurAdvhcIcmpTableParams,
) *SlbCurAdvhcIcmpTable {
	return &SlbCurAdvhcIcmpTable{
		SlbCurAdvhcIcmpID: slbCurAdvhcIcmpID,
		Params:            params,
	}
}

func (c *SlbCurAdvhcIcmpTable) Name() string {
	return "SlbCurAdvhcIcmpTable"
}

func (c *SlbCurAdvhcIcmpTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbCurAdvhcIcmpTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbCurAdvhcIcmpTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbCurAdvhcIcmpID).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbCurAdvhcIcmpID)
}

type SlbCurAdvhcIcmpTableIPVer int32

const (
	SlbCurAdvhcIcmpTableIPVer_Ipv4 SlbCurAdvhcIcmpTableIPVer = 1
	SlbCurAdvhcIcmpTableIPVer_Ipv6 SlbCurAdvhcIcmpTableIPVer = 2
	SlbCurAdvhcIcmpTableIPVer_None SlbCurAdvhcIcmpTableIPVer = 3
)

type SlbCurAdvhcIcmpTableTransparent int32

const (
	SlbCurAdvhcIcmpTableTransparent_Enabled     SlbCurAdvhcIcmpTableTransparent = 1
	SlbCurAdvhcIcmpTableTransparent_Disabled    SlbCurAdvhcIcmpTableTransparent = 2
	SlbCurAdvhcIcmpTableTransparent_Unsupported SlbCurAdvhcIcmpTableTransparent = 2147483647
)

type SlbCurAdvhcIcmpTableInvert int32

const (
	SlbCurAdvhcIcmpTableInvert_Enabled     SlbCurAdvhcIcmpTableInvert = 1
	SlbCurAdvhcIcmpTableInvert_Disabled    SlbCurAdvhcIcmpTableInvert = 2
	SlbCurAdvhcIcmpTableInvert_Unsupported SlbCurAdvhcIcmpTableInvert = 2147483647
)

type SlbCurAdvhcIcmpTableSnat int32

const (
	SlbCurAdvhcIcmpTableSnat_Enabled     SlbCurAdvhcIcmpTableSnat = 1
	SlbCurAdvhcIcmpTableSnat_Disabled    SlbCurAdvhcIcmpTableSnat = 2
	SlbCurAdvhcIcmpTableSnat_Unsupported SlbCurAdvhcIcmpTableSnat = 2147483647
)

type SlbCurAdvhcIcmpTableParams struct {
	// ICMP health check id.
	ID string `json:"ID,omitempty"`
	// ICMP health check name.
	Name string `json:"Name,omitempty"`
	// ICMP health check destination port.
	DPort uint64 `json:"DPort,omitempty"`
	// ICMP health check destination IP version.
	IPVer SlbCurAdvhcIcmpTableIPVer `json:"IPVer,omitempty"`
	// ICMP health check destination hostname.
	HostName string `json:"HostName,omitempty"`
	// ICMP health check transparent flag.
	Transparent SlbCurAdvhcIcmpTableTransparent `json:"Transparent,omitempty"`
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
	Invert SlbCurAdvhcIcmpTableInvert `json:"Invert,omitempty"`
	// ICMP hHealth Check src NAT (PIP) flag.
	Snat SlbCurAdvhcIcmpTableSnat `json:"Snat,omitempty"`
}

func (p SlbCurAdvhcIcmpTableParams) iMABean() {}
