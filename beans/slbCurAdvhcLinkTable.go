package beans

import (
	"fmt"
	"reflect"
)

// SlbCurAdvhcLinkTable Note:This mib is not supported on VX instance of Virtualization.
type SlbCurAdvhcLinkTable struct {
	// Link Health check id.
	SlbCurAdvhcLinkID string
	Params            *SlbCurAdvhcLinkTableParams
}

func NewSlbCurAdvhcLinkTableList() *SlbCurAdvhcLinkTable {
	return &SlbCurAdvhcLinkTable{}
}

func NewSlbCurAdvhcLinkTable(
	slbCurAdvhcLinkID string,
	params *SlbCurAdvhcLinkTableParams,
) *SlbCurAdvhcLinkTable {
	return &SlbCurAdvhcLinkTable{
		SlbCurAdvhcLinkID: slbCurAdvhcLinkID,
		Params:            params,
	}
}

func (c *SlbCurAdvhcLinkTable) Name() string {
	return "SlbCurAdvhcLinkTable"
}

func (c *SlbCurAdvhcLinkTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbCurAdvhcLinkTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbCurAdvhcLinkTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbCurAdvhcLinkID).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbCurAdvhcLinkID)
}

type SlbCurAdvhcLinkTableIPVer int32

const (
	SlbCurAdvhcLinkTableIPVer_Ipv4 SlbCurAdvhcLinkTableIPVer = 1
	SlbCurAdvhcLinkTableIPVer_Ipv6 SlbCurAdvhcLinkTableIPVer = 2
	SlbCurAdvhcLinkTableIPVer_None SlbCurAdvhcLinkTableIPVer = 3
)

type SlbCurAdvhcLinkTableTransparent int32

const (
	SlbCurAdvhcLinkTableTransparent_Enabled     SlbCurAdvhcLinkTableTransparent = 1
	SlbCurAdvhcLinkTableTransparent_Disabled    SlbCurAdvhcLinkTableTransparent = 2
	SlbCurAdvhcLinkTableTransparent_Unsupported SlbCurAdvhcLinkTableTransparent = 2147483647
)

type SlbCurAdvhcLinkTableInvert int32

const (
	SlbCurAdvhcLinkTableInvert_Enabled     SlbCurAdvhcLinkTableInvert = 1
	SlbCurAdvhcLinkTableInvert_Disabled    SlbCurAdvhcLinkTableInvert = 2
	SlbCurAdvhcLinkTableInvert_Unsupported SlbCurAdvhcLinkTableInvert = 2147483647
)

type SlbCurAdvhcLinkTableSnat int32

const (
	SlbCurAdvhcLinkTableSnat_Enabled     SlbCurAdvhcLinkTableSnat = 1
	SlbCurAdvhcLinkTableSnat_Disabled    SlbCurAdvhcLinkTableSnat = 2
	SlbCurAdvhcLinkTableSnat_Unsupported SlbCurAdvhcLinkTableSnat = 2147483647
)

type SlbCurAdvhcLinkTableParams struct {
	// Link Health check id.
	ID string `json:"ID,omitempty"`
	// Link Health check name.
	Name string `json:"Name,omitempty"`
	// Link Health check destination port.
	DPort uint64 `json:"DPort,omitempty"`
	// Link Health check destination IP version.
	IPVer SlbCurAdvhcLinkTableIPVer `json:"IPVer,omitempty"`
	// Link Health check destination hostname.
	HostName string `json:"HostName,omitempty"`
	// Link Health check transparent flag.
	Transparent SlbCurAdvhcLinkTableTransparent `json:"Transparent,omitempty"`
	// Link Health check interval.
	Interval uint64 `json:"Interval,omitempty"`
	// Link Health check retries counter.
	Retries uint64 `json:"Retries,omitempty"`
	// Link Health check retries counter in down state.
	RestoreRetries uint64 `json:"RestoreRetries,omitempty"`
	// Link Health check timeout.
	Timeout uint64 `json:"Timeout,omitempty"`
	// Link Health check overflow flag.
	Overflow uint64 `json:"Overflow,omitempty"`
	// Link Health check interval in down state.
	DownInterval uint64 `json:"DownInterval,omitempty"`
	// Link Health check invert flag.
	Invert SlbCurAdvhcLinkTableInvert `json:"Invert,omitempty"`
	// Link Health Check src NAT (PIP) flag.
	Snat SlbCurAdvhcLinkTableSnat `json:"Snat,omitempty"`
}

func (p SlbCurAdvhcLinkTableParams) iMABean() {}
