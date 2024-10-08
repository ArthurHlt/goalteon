package beans

import (
	"fmt"
	"reflect"
)

// SlbNewAdvhcNntpTable Note:This mib is not supported on VX instance of Virtualization.
type SlbNewAdvhcNntpTable struct {
	// NNTP Health check id.
	SlbNewAdvhcNntpID string
	Params            *SlbNewAdvhcNntpTableParams
}

func NewSlbNewAdvhcNntpTableList() *SlbNewAdvhcNntpTable {
	return &SlbNewAdvhcNntpTable{}
}

func NewSlbNewAdvhcNntpTable(
	slbNewAdvhcNntpID string,
	params *SlbNewAdvhcNntpTableParams,
) *SlbNewAdvhcNntpTable {
	return &SlbNewAdvhcNntpTable{
		SlbNewAdvhcNntpID: slbNewAdvhcNntpID,
		Params:            params,
	}
}

func (c *SlbNewAdvhcNntpTable) Name() string {
	return "SlbNewAdvhcNntpTable"
}

func (c *SlbNewAdvhcNntpTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbNewAdvhcNntpTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbNewAdvhcNntpTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbNewAdvhcNntpID).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbNewAdvhcNntpID)
}

type SlbNewAdvhcNntpTableIPVer int32

const (
	SlbNewAdvhcNntpTableIPVer_Ipv4 SlbNewAdvhcNntpTableIPVer = 1
	SlbNewAdvhcNntpTableIPVer_Ipv6 SlbNewAdvhcNntpTableIPVer = 2
	SlbNewAdvhcNntpTableIPVer_None SlbNewAdvhcNntpTableIPVer = 3
)

type SlbNewAdvhcNntpTableTransparent int32

const (
	SlbNewAdvhcNntpTableTransparent_Enabled     SlbNewAdvhcNntpTableTransparent = 1
	SlbNewAdvhcNntpTableTransparent_Disabled    SlbNewAdvhcNntpTableTransparent = 2
	SlbNewAdvhcNntpTableTransparent_Unsupported SlbNewAdvhcNntpTableTransparent = 2147483647
)

type SlbNewAdvhcNntpTableInvert int32

const (
	SlbNewAdvhcNntpTableInvert_Enabled     SlbNewAdvhcNntpTableInvert = 1
	SlbNewAdvhcNntpTableInvert_Disabled    SlbNewAdvhcNntpTableInvert = 2
	SlbNewAdvhcNntpTableInvert_Unsupported SlbNewAdvhcNntpTableInvert = 2147483647
)

type SlbNewAdvhcNntpTableDelete int32

const (
	SlbNewAdvhcNntpTableDelete_Other       SlbNewAdvhcNntpTableDelete = 1
	SlbNewAdvhcNntpTableDelete_Delete      SlbNewAdvhcNntpTableDelete = 2
	SlbNewAdvhcNntpTableDelete_Unsupported SlbNewAdvhcNntpTableDelete = 2147483647
)

type SlbNewAdvhcNntpTableSnat int32

const (
	SlbNewAdvhcNntpTableSnat_Enabled     SlbNewAdvhcNntpTableSnat = 1
	SlbNewAdvhcNntpTableSnat_Disabled    SlbNewAdvhcNntpTableSnat = 2
	SlbNewAdvhcNntpTableSnat_Unsupported SlbNewAdvhcNntpTableSnat = 2147483647
)

type SlbNewAdvhcNntpTableParams struct {
	// NNTP Health check id.
	ID string `json:"ID,omitempty"`
	// NNTP Health check name.
	Name string `json:"Name,omitempty"`
	// NNTP Health check destination port.
	DPort uint64 `json:"DPort,omitempty"`
	// NNTP Health check destination IP version.
	IPVer SlbNewAdvhcNntpTableIPVer `json:"IPVer,omitempty"`
	// NNTP Health check destination hostname.
	HostName string `json:"HostName,omitempty"`
	// NNTP Health check transparent flag.
	Transparent SlbNewAdvhcNntpTableTransparent `json:"Transparent,omitempty"`
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
	Invert SlbNewAdvhcNntpTableInvert `json:"Invert,omitempty"`
	// NNTP Health check newsgroup name.
	NewsgroupName string `json:"NewsgroupName,omitempty"`
	// NNTP Health check copy flag.
	Copy DisplayString `json:"Copy,omitempty"`
	// When set to the value of 2 (delete), the entire row is deleted.
	// When read, other(1) is returned. Setting the value to anything
	// other than 2(delete) has no effect on the state of the row.
	Delete SlbNewAdvhcNntpTableDelete `json:"Delete,omitempty"`
	// NNTP Health Check src NAT (PIP) flag.
	Snat SlbNewAdvhcNntpTableSnat `json:"Snat,omitempty"`
}

func (p SlbNewAdvhcNntpTableParams) iMABean() {}
