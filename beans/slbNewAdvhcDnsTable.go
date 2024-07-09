package beans

import (
	"fmt"
	"reflect"
)

// SlbNewAdvhcDnsTable Note:This mib is not supported on VX instance of Virtualization.
type SlbNewAdvhcDnsTable struct {
	// DNS Health check id.
	SlbNewAdvhcDnsID string
	Params           *SlbNewAdvhcDnsTableParams
}

func NewSlbNewAdvhcDnsTableList() *SlbNewAdvhcDnsTable {
	return &SlbNewAdvhcDnsTable{}
}

func NewSlbNewAdvhcDnsTable(
	slbNewAdvhcDnsID string,
	params *SlbNewAdvhcDnsTableParams,
) *SlbNewAdvhcDnsTable {
	return &SlbNewAdvhcDnsTable{
		SlbNewAdvhcDnsID: slbNewAdvhcDnsID,
		Params:           params,
	}
}

func (c *SlbNewAdvhcDnsTable) Name() string {
	return "SlbNewAdvhcDnsTable"
}

func (c *SlbNewAdvhcDnsTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbNewAdvhcDnsTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbNewAdvhcDnsTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbNewAdvhcDnsID).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbNewAdvhcDnsID)
}

type SlbNewAdvhcDnsTableIPVer int32

const (
	SlbNewAdvhcDnsTableIPVer_Ipv4 SlbNewAdvhcDnsTableIPVer = 1
	SlbNewAdvhcDnsTableIPVer_Ipv6 SlbNewAdvhcDnsTableIPVer = 2
	SlbNewAdvhcDnsTableIPVer_None SlbNewAdvhcDnsTableIPVer = 3
)

type SlbNewAdvhcDnsTableTransparent int32

const (
	SlbNewAdvhcDnsTableTransparent_Enabled     SlbNewAdvhcDnsTableTransparent = 1
	SlbNewAdvhcDnsTableTransparent_Disabled    SlbNewAdvhcDnsTableTransparent = 2
	SlbNewAdvhcDnsTableTransparent_Unsupported SlbNewAdvhcDnsTableTransparent = 2147483647
)

type SlbNewAdvhcDnsTableInvert int32

const (
	SlbNewAdvhcDnsTableInvert_Enabled     SlbNewAdvhcDnsTableInvert = 1
	SlbNewAdvhcDnsTableInvert_Disabled    SlbNewAdvhcDnsTableInvert = 2
	SlbNewAdvhcDnsTableInvert_Unsupported SlbNewAdvhcDnsTableInvert = 2147483647
)

type SlbNewAdvhcDnsTableTransport int32

const (
	SlbNewAdvhcDnsTableTransport_Tcp         SlbNewAdvhcDnsTableTransport = 1
	SlbNewAdvhcDnsTableTransport_Udp         SlbNewAdvhcDnsTableTransport = 2
	SlbNewAdvhcDnsTableTransport_Unsupported SlbNewAdvhcDnsTableTransport = 2147483647
)

type SlbNewAdvhcDnsTableDelete int32

const (
	SlbNewAdvhcDnsTableDelete_Other       SlbNewAdvhcDnsTableDelete = 1
	SlbNewAdvhcDnsTableDelete_Delete      SlbNewAdvhcDnsTableDelete = 2
	SlbNewAdvhcDnsTableDelete_Unsupported SlbNewAdvhcDnsTableDelete = 2147483647
)

type SlbNewAdvhcDnsTableSnat int32

const (
	SlbNewAdvhcDnsTableSnat_Enabled     SlbNewAdvhcDnsTableSnat = 1
	SlbNewAdvhcDnsTableSnat_Disabled    SlbNewAdvhcDnsTableSnat = 2
	SlbNewAdvhcDnsTableSnat_Unsupported SlbNewAdvhcDnsTableSnat = 2147483647
)

type SlbNewAdvhcDnsTableParams struct {
	// DNS Health check id.
	ID string `json:"ID,omitempty"`
	// DNS Health check name.
	Name string `json:"Name,omitempty"`
	// DNS Health check destination port.
	DPort uint64 `json:"DPort,omitempty"`
	// DNS Health check IP version.
	IPVer SlbNewAdvhcDnsTableIPVer `json:"IPVer,omitempty"`
	// DNS Health check destination host name.
	HostName string `json:"HostName,omitempty"`
	// DNS Health check transparent flag.
	Transparent SlbNewAdvhcDnsTableTransparent `json:"Transparent,omitempty"`
	// DNS Health check interval.
	Interval uint64 `json:"Interval,omitempty"`
	// DNS Health check retries counter.
	Retries uint64 `json:"Retries,omitempty"`
	// DNS Health check restore from down state counter.
	RestoreRetries uint64 `json:"RestoreRetries,omitempty"`
	// DNS Health check timeout.
	Timeout uint64 `json:"Timeout,omitempty"`
	// DNS Health check overflow flag.
	Overflow uint64 `json:"Overflow,omitempty"`
	// DNS Health check recovery from down state interval.
	DownInterval uint64 `json:"DownInterval,omitempty"`
	// DNS Health check invert flag.
	Invert SlbNewAdvhcDnsTableInvert `json:"Invert,omitempty"`
	// DNS Health check domain name parameter.
	DomainName string `json:"DomainName,omitempty"`
	// DNS Health check transport layer.
	Transport SlbNewAdvhcDnsTableTransport `json:"Transport,omitempty"`
	// DNS Health check field to copy object.
	Copy string `json:"Copy,omitempty"`
	// When set to the value of 2 (delete), the entire row is deleted.
	// When read, other(1) is returned. Setting the value to anything
	// other than 2(delete) has no effect on the state of the row.
	Delete SlbNewAdvhcDnsTableDelete `json:"Delete,omitempty"`
	// DNS Health Check src NAT (PIP) flag.
	Snat SlbNewAdvhcDnsTableSnat `json:"Snat,omitempty"`
}

func (p SlbNewAdvhcDnsTableParams) iMABean() {}
