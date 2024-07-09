package beans

import (
	"fmt"
	"reflect"
)

// SlbCurAdvhcNntpTable Note:This mib is not supported on VX instance of Virtualization.
type SlbCurAdvhcNntpTable struct {
	// NNTP Health check id.
	SlbCurAdvhcNntpID string
	Params            *SlbCurAdvhcNntpTableParams
}

func NewSlbCurAdvhcNntpTableList() *SlbCurAdvhcNntpTable {
	return &SlbCurAdvhcNntpTable{}
}

func NewSlbCurAdvhcNntpTable(
	slbCurAdvhcNntpID string,
	params *SlbCurAdvhcNntpTableParams,
) *SlbCurAdvhcNntpTable {
	return &SlbCurAdvhcNntpTable{
		SlbCurAdvhcNntpID: slbCurAdvhcNntpID,
		Params:            params,
	}
}

func (c *SlbCurAdvhcNntpTable) Name() string {
	return "SlbCurAdvhcNntpTable"
}

func (c *SlbCurAdvhcNntpTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbCurAdvhcNntpTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbCurAdvhcNntpTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbCurAdvhcNntpID).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbCurAdvhcNntpID)
}

type SlbCurAdvhcNntpTableIPVer int32

const (
	SlbCurAdvhcNntpTableIPVer_Ipv4 SlbCurAdvhcNntpTableIPVer = 1
	SlbCurAdvhcNntpTableIPVer_Ipv6 SlbCurAdvhcNntpTableIPVer = 2
	SlbCurAdvhcNntpTableIPVer_None SlbCurAdvhcNntpTableIPVer = 3
)

type SlbCurAdvhcNntpTableTransparent int32

const (
	SlbCurAdvhcNntpTableTransparent_Enabled     SlbCurAdvhcNntpTableTransparent = 1
	SlbCurAdvhcNntpTableTransparent_Disabled    SlbCurAdvhcNntpTableTransparent = 2
	SlbCurAdvhcNntpTableTransparent_Unsupported SlbCurAdvhcNntpTableTransparent = 2147483647
)

type SlbCurAdvhcNntpTableInvert int32

const (
	SlbCurAdvhcNntpTableInvert_Enabled     SlbCurAdvhcNntpTableInvert = 1
	SlbCurAdvhcNntpTableInvert_Disabled    SlbCurAdvhcNntpTableInvert = 2
	SlbCurAdvhcNntpTableInvert_Unsupported SlbCurAdvhcNntpTableInvert = 2147483647
)

type SlbCurAdvhcNntpTableSnat int32

const (
	SlbCurAdvhcNntpTableSnat_Enabled     SlbCurAdvhcNntpTableSnat = 1
	SlbCurAdvhcNntpTableSnat_Disabled    SlbCurAdvhcNntpTableSnat = 2
	SlbCurAdvhcNntpTableSnat_Unsupported SlbCurAdvhcNntpTableSnat = 2147483647
)

type SlbCurAdvhcNntpTableParams struct {
	// NNTP Health check id.
	ID string `json:"ID,omitempty"`
	// NNTP Health check name.
	Name string `json:"Name,omitempty"`
	// NNTP Health check destination port.
	DPort uint64 `json:"DPort,omitempty"`
	// NNTP Health check destination IP version.
	IPVer SlbCurAdvhcNntpTableIPVer `json:"IPVer,omitempty"`
	// NNTP Health check destination hostname.
	HostName string `json:"HostName,omitempty"`
	// NNTP Health check transparent flag.
	Transparent SlbCurAdvhcNntpTableTransparent `json:"Transparent,omitempty"`
	// NNTP Health check interval.
	Interval uint64 `json:"Interval,omitempty"`
	// NNTP Health check retries counter.
	Retries uint64 `json:"Retries,omitempty"`
	// NNTP Health check restore from down state retries counter.
	RestoreRetries uint64 `json:"RestoreRetries,omitempty"`
	// NNTP Health check timeout.
	Timeout uint64 `json:"Timeout,omitempty"`
	// NNTP Health check overflow flag.
	Overflow uint64 `json:"Overflow,omitempty"`
	// NNTP Health check interval when in down state.
	DownInterval uint64 `json:"DownInterval,omitempty"`
	// NNTP Health check invert flag.
	Invert SlbCurAdvhcNntpTableInvert `json:"Invert,omitempty"`
	// NNTP Health check newsgroup name.
	NewsgroupName string `json:"NewsgroupName,omitempty"`
	// NNTP Health Check src NAT (PIP) flag.
	Snat SlbCurAdvhcNntpTableSnat `json:"Snat,omitempty"`
}

func (p SlbCurAdvhcNntpTableParams) iMABean() {}
