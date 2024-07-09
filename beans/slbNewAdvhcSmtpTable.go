package beans

import (
	"fmt"
	"reflect"
)

// SlbNewAdvhcSmtpTable Note:This mib is not supported on VX instance of Virtualization.
type SlbNewAdvhcSmtpTable struct {
	// SMTP health check id.
	SlbNewAdvhcSmtpID string
	Params            *SlbNewAdvhcSmtpTableParams
}

func NewSlbNewAdvhcSmtpTableList() *SlbNewAdvhcSmtpTable {
	return &SlbNewAdvhcSmtpTable{}
}

func NewSlbNewAdvhcSmtpTable(
	slbNewAdvhcSmtpID string,
	params *SlbNewAdvhcSmtpTableParams,
) *SlbNewAdvhcSmtpTable {
	return &SlbNewAdvhcSmtpTable{
		SlbNewAdvhcSmtpID: slbNewAdvhcSmtpID,
		Params:            params,
	}
}

func (c *SlbNewAdvhcSmtpTable) Name() string {
	return "SlbNewAdvhcSmtpTable"
}

func (c *SlbNewAdvhcSmtpTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbNewAdvhcSmtpTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbNewAdvhcSmtpTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbNewAdvhcSmtpID).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbNewAdvhcSmtpID)
}

type SlbNewAdvhcSmtpTableIPVer int32

const (
	SlbNewAdvhcSmtpTableIPVer_Ipv4 SlbNewAdvhcSmtpTableIPVer = 1
	SlbNewAdvhcSmtpTableIPVer_Ipv6 SlbNewAdvhcSmtpTableIPVer = 2
	SlbNewAdvhcSmtpTableIPVer_None SlbNewAdvhcSmtpTableIPVer = 3
)

type SlbNewAdvhcSmtpTableTransparent int32

const (
	SlbNewAdvhcSmtpTableTransparent_Enabled     SlbNewAdvhcSmtpTableTransparent = 1
	SlbNewAdvhcSmtpTableTransparent_Disabled    SlbNewAdvhcSmtpTableTransparent = 2
	SlbNewAdvhcSmtpTableTransparent_Unsupported SlbNewAdvhcSmtpTableTransparent = 2147483647
)

type SlbNewAdvhcSmtpTableInvert int32

const (
	SlbNewAdvhcSmtpTableInvert_Enabled     SlbNewAdvhcSmtpTableInvert = 1
	SlbNewAdvhcSmtpTableInvert_Disabled    SlbNewAdvhcSmtpTableInvert = 2
	SlbNewAdvhcSmtpTableInvert_Unsupported SlbNewAdvhcSmtpTableInvert = 2147483647
)

type SlbNewAdvhcSmtpTableDelete int32

const (
	SlbNewAdvhcSmtpTableDelete_Other       SlbNewAdvhcSmtpTableDelete = 1
	SlbNewAdvhcSmtpTableDelete_Delete      SlbNewAdvhcSmtpTableDelete = 2
	SlbNewAdvhcSmtpTableDelete_Unsupported SlbNewAdvhcSmtpTableDelete = 2147483647
)

type SlbNewAdvhcSmtpTableSnat int32

const (
	SlbNewAdvhcSmtpTableSnat_Enabled     SlbNewAdvhcSmtpTableSnat = 1
	SlbNewAdvhcSmtpTableSnat_Disabled    SlbNewAdvhcSmtpTableSnat = 2
	SlbNewAdvhcSmtpTableSnat_Unsupported SlbNewAdvhcSmtpTableSnat = 2147483647
)

type SlbNewAdvhcSmtpTableParams struct {
	// SMTP health check id.
	ID string `json:"ID,omitempty"`
	// SMTP health check name.
	Name string `json:"Name,omitempty"`
	// SMTP health check destination port.
	DPort uint64 `json:"DPort,omitempty"`
	// SMTP health check destination IP version.
	IPVer SlbNewAdvhcSmtpTableIPVer `json:"IPVer,omitempty"`
	// SMTP health check destination hostname.
	HostName string `json:"HostName,omitempty"`
	// SMTP health check transparent flag.
	Transparent SlbNewAdvhcSmtpTableTransparent `json:"Transparent,omitempty"`
	// SMTP health check interval.
	Interval uint64 `json:"Interval,omitempty"`
	// SMTP health check retries counter.
	Retries uint64 `json:"Retries,omitempty"`
	// SMTP health check retries when in down state counter.
	RestoreRetries uint64 `json:"RestoreRetries,omitempty"`
	// SMTP health check timeout.
	Timeout uint64 `json:"Timeout,omitempty"`
	// SMTP health check overflow flag.
	Overflow uint64 `json:"Overflow,omitempty"`
	// SMTP health check interval when in the down state.
	DownInterval uint64 `json:"DownInterval,omitempty"`
	// SMTP health check invert flag.
	Invert SlbNewAdvhcSmtpTableInvert `json:"Invert,omitempty"`
	// SMTP health check username.
	UserName string `json:"UserName,omitempty"`
	// SMTP health check copy flag.
	Copy string `json:"Copy,omitempty"`
	// When set to the value of 2 (delete), the entire row is deleted.
	// When read, other(1) is returned. Setting the value to anything
	// other than 2(delete) has no effect on the state of the row.
	Delete SlbNewAdvhcSmtpTableDelete `json:"Delete,omitempty"`
	// SMTP Health Check src NAT (PIP) flag.
	Snat SlbNewAdvhcSmtpTableSnat `json:"Snat,omitempty"`
}

func (p SlbNewAdvhcSmtpTableParams) iMABean() {}
