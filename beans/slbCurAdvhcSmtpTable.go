package beans

import (
	"fmt"
	"reflect"
)

// SlbCurAdvhcSmtpTable Note:This mib is not supported on VX instance of Virtualization.
type SlbCurAdvhcSmtpTable struct {
	// SMTP health check id.
	SlbCurAdvhcSmtpID string
	Params            *SlbCurAdvhcSmtpTableParams
}

func NewSlbCurAdvhcSmtpTableList() *SlbCurAdvhcSmtpTable {
	return &SlbCurAdvhcSmtpTable{}
}

func NewSlbCurAdvhcSmtpTable(
	slbCurAdvhcSmtpID string,
	params *SlbCurAdvhcSmtpTableParams,
) *SlbCurAdvhcSmtpTable {
	return &SlbCurAdvhcSmtpTable{
		SlbCurAdvhcSmtpID: slbCurAdvhcSmtpID,
		Params:            params,
	}
}

func (c *SlbCurAdvhcSmtpTable) Name() string {
	return "SlbCurAdvhcSmtpTable"
}

func (c *SlbCurAdvhcSmtpTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbCurAdvhcSmtpTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbCurAdvhcSmtpTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbCurAdvhcSmtpID).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbCurAdvhcSmtpID)
}

type SlbCurAdvhcSmtpTableIPVer int32

const (
	SlbCurAdvhcSmtpTableIPVer_Ipv4 SlbCurAdvhcSmtpTableIPVer = 1
	SlbCurAdvhcSmtpTableIPVer_Ipv6 SlbCurAdvhcSmtpTableIPVer = 2
	SlbCurAdvhcSmtpTableIPVer_None SlbCurAdvhcSmtpTableIPVer = 3
)

type SlbCurAdvhcSmtpTableTransparent int32

const (
	SlbCurAdvhcSmtpTableTransparent_Enabled     SlbCurAdvhcSmtpTableTransparent = 1
	SlbCurAdvhcSmtpTableTransparent_Disabled    SlbCurAdvhcSmtpTableTransparent = 2
	SlbCurAdvhcSmtpTableTransparent_Unsupported SlbCurAdvhcSmtpTableTransparent = 2147483647
)

type SlbCurAdvhcSmtpTableInvert int32

const (
	SlbCurAdvhcSmtpTableInvert_Enabled     SlbCurAdvhcSmtpTableInvert = 1
	SlbCurAdvhcSmtpTableInvert_Disabled    SlbCurAdvhcSmtpTableInvert = 2
	SlbCurAdvhcSmtpTableInvert_Unsupported SlbCurAdvhcSmtpTableInvert = 2147483647
)

type SlbCurAdvhcSmtpTableSnat int32

const (
	SlbCurAdvhcSmtpTableSnat_Enabled     SlbCurAdvhcSmtpTableSnat = 1
	SlbCurAdvhcSmtpTableSnat_Disabled    SlbCurAdvhcSmtpTableSnat = 2
	SlbCurAdvhcSmtpTableSnat_Unsupported SlbCurAdvhcSmtpTableSnat = 2147483647
)

type SlbCurAdvhcSmtpTableParams struct {
	// SMTP health check id.
	ID string `json:"ID,omitempty"`
	// SMTP health check name.
	Name string `json:"Name,omitempty"`
	// SMTP health check destination port.
	DPort uint64 `json:"DPort,omitempty"`
	// SMTP health check destination IP version.
	IPVer SlbCurAdvhcSmtpTableIPVer `json:"IPVer,omitempty"`
	// SMTP health check destination hostname.
	HostName string `json:"HostName,omitempty"`
	// SMTP health check transparent flag.
	Transparent SlbCurAdvhcSmtpTableTransparent `json:"Transparent,omitempty"`
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
	Invert SlbCurAdvhcSmtpTableInvert `json:"Invert,omitempty"`
	// SMTP health check username.
	UserName string `json:"UserName,omitempty"`
	// SMTP Health Check src NAT (PIP) flag.
	Snat SlbCurAdvhcSmtpTableSnat `json:"Snat,omitempty"`
}

func (p SlbCurAdvhcSmtpTableParams) iMABean() {}
