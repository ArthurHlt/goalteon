package beans

import (
	"fmt"
	"reflect"
)

// SlbNewAdvhcLinkTable Note:This mib is not supported on VX instance of Virtualization.
type SlbNewAdvhcLinkTable struct {
	// Link Health check id.
	SlbNewAdvhcLinkID string
	Params            *SlbNewAdvhcLinkTableParams
}

func NewSlbNewAdvhcLinkTableList() *SlbNewAdvhcLinkTable {
	return &SlbNewAdvhcLinkTable{}
}

func NewSlbNewAdvhcLinkTable(
	slbNewAdvhcLinkID string,
	params *SlbNewAdvhcLinkTableParams,
) *SlbNewAdvhcLinkTable {
	return &SlbNewAdvhcLinkTable{
		SlbNewAdvhcLinkID: slbNewAdvhcLinkID,
		Params:            params,
	}
}

func (c *SlbNewAdvhcLinkTable) Name() string {
	return "SlbNewAdvhcLinkTable"
}

func (c *SlbNewAdvhcLinkTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbNewAdvhcLinkTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbNewAdvhcLinkTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbNewAdvhcLinkID).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbNewAdvhcLinkID)
}

type SlbNewAdvhcLinkTableIPVer int32

const (
	SlbNewAdvhcLinkTableIPVer_Ipv4 SlbNewAdvhcLinkTableIPVer = 1
	SlbNewAdvhcLinkTableIPVer_Ipv6 SlbNewAdvhcLinkTableIPVer = 2
	SlbNewAdvhcLinkTableIPVer_None SlbNewAdvhcLinkTableIPVer = 3
)

type SlbNewAdvhcLinkTableTransparent int32

const (
	SlbNewAdvhcLinkTableTransparent_Enabled     SlbNewAdvhcLinkTableTransparent = 1
	SlbNewAdvhcLinkTableTransparent_Disabled    SlbNewAdvhcLinkTableTransparent = 2
	SlbNewAdvhcLinkTableTransparent_Unsupported SlbNewAdvhcLinkTableTransparent = 2147483647
)

type SlbNewAdvhcLinkTableInvert int32

const (
	SlbNewAdvhcLinkTableInvert_Enabled     SlbNewAdvhcLinkTableInvert = 1
	SlbNewAdvhcLinkTableInvert_Disabled    SlbNewAdvhcLinkTableInvert = 2
	SlbNewAdvhcLinkTableInvert_Unsupported SlbNewAdvhcLinkTableInvert = 2147483647
)

type SlbNewAdvhcLinkTableDelete int32

const (
	SlbNewAdvhcLinkTableDelete_Other       SlbNewAdvhcLinkTableDelete = 1
	SlbNewAdvhcLinkTableDelete_Delete      SlbNewAdvhcLinkTableDelete = 2
	SlbNewAdvhcLinkTableDelete_Unsupported SlbNewAdvhcLinkTableDelete = 2147483647
)

type SlbNewAdvhcLinkTableSnat int32

const (
	SlbNewAdvhcLinkTableSnat_Enabled     SlbNewAdvhcLinkTableSnat = 1
	SlbNewAdvhcLinkTableSnat_Disabled    SlbNewAdvhcLinkTableSnat = 2
	SlbNewAdvhcLinkTableSnat_Unsupported SlbNewAdvhcLinkTableSnat = 2147483647
)

type SlbNewAdvhcLinkTableParams struct {
	// Link Health check id.
	ID string `json:"ID,omitempty"`
	// Link Health check name.
	Name string `json:"Name,omitempty"`
	// Link Health check destination port.
	DPort uint64 `json:"DPort,omitempty"`
	// Link Health check destination IP version.
	IPVer SlbNewAdvhcLinkTableIPVer `json:"IPVer,omitempty"`
	// Link Health check destination hostname.
	HostName string `json:"HostName,omitempty"`
	// Link Health check transparent flag.
	Transparent SlbNewAdvhcLinkTableTransparent `json:"Transparent,omitempty"`
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
	Invert SlbNewAdvhcLinkTableInvert `json:"Invert,omitempty"`
	// Link Health check copy indicator.
	Copy DisplayString `json:"Copy,omitempty"`
	// When set to the value of 2 (delete), the entire row is deleted.
	// When read, other(1) is returned. Setting the value to anything
	// other than 2(delete) has no effect on the state of the row.
	Delete SlbNewAdvhcLinkTableDelete `json:"Delete,omitempty"`
	// Link Health Check src NAT (PIP) flag.
	Snat SlbNewAdvhcLinkTableSnat `json:"Snat,omitempty"`
}

func (p SlbNewAdvhcLinkTableParams) iMABean() {}
