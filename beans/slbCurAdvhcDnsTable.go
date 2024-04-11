package beans

import (
	"fmt"
	"reflect"
)

// SlbCurAdvhcDnsTable Note:This mib is not supported on VX instance of Virtualization.
type SlbCurAdvhcDnsTable struct {
	// DNS Health check id.
	SlbCurAdvhcDnsID string
	Params           *SlbCurAdvhcDnsTableParams
}

func NewSlbCurAdvhcDnsTableList() *SlbCurAdvhcDnsTable {
	return &SlbCurAdvhcDnsTable{}
}

func NewSlbCurAdvhcDnsTable(
	slbCurAdvhcDnsID string,
	params *SlbCurAdvhcDnsTableParams,
) *SlbCurAdvhcDnsTable {
	return &SlbCurAdvhcDnsTable{
		SlbCurAdvhcDnsID: slbCurAdvhcDnsID,
		Params:           params,
	}
}

func (c *SlbCurAdvhcDnsTable) Name() string {
	return "SlbCurAdvhcDnsTable"
}

func (c *SlbCurAdvhcDnsTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbCurAdvhcDnsTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbCurAdvhcDnsTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbCurAdvhcDnsID).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbCurAdvhcDnsID)
}

type SlbCurAdvhcDnsTableIPVer int32

const (
	SlbCurAdvhcDnsTableIPVer_Ipv4 SlbCurAdvhcDnsTableIPVer = 1
	SlbCurAdvhcDnsTableIPVer_Ipv6 SlbCurAdvhcDnsTableIPVer = 2
	SlbCurAdvhcDnsTableIPVer_None SlbCurAdvhcDnsTableIPVer = 3
)

type SlbCurAdvhcDnsTableTransparent int32

const (
	SlbCurAdvhcDnsTableTransparent_Enabled     SlbCurAdvhcDnsTableTransparent = 1
	SlbCurAdvhcDnsTableTransparent_Disabled    SlbCurAdvhcDnsTableTransparent = 2
	SlbCurAdvhcDnsTableTransparent_Unsupported SlbCurAdvhcDnsTableTransparent = 2147483647
)

type SlbCurAdvhcDnsTableInvert int32

const (
	SlbCurAdvhcDnsTableInvert_Enabled     SlbCurAdvhcDnsTableInvert = 1
	SlbCurAdvhcDnsTableInvert_Disabled    SlbCurAdvhcDnsTableInvert = 2
	SlbCurAdvhcDnsTableInvert_Unsupported SlbCurAdvhcDnsTableInvert = 2147483647
)

type SlbCurAdvhcDnsTableTransport int32

const (
	SlbCurAdvhcDnsTableTransport_Tcp         SlbCurAdvhcDnsTableTransport = 1
	SlbCurAdvhcDnsTableTransport_Udp         SlbCurAdvhcDnsTableTransport = 2
	SlbCurAdvhcDnsTableTransport_Unsupported SlbCurAdvhcDnsTableTransport = 2147483647
)

type SlbCurAdvhcDnsTableParams struct {
	// DNS Health check id.
	ID string `json:"ID,omitempty"`
	// DNS Health check name.
	Name string `json:"Name,omitempty"`
	// DNS Health check destination port.
	DPort uint64 `json:"DPort,omitempty"`
	// DNS Health check IP version.
	IPVer SlbCurAdvhcDnsTableIPVer `json:"IPVer,omitempty"`
	// DNS Health check destination host name.
	HostName string `json:"HostName,omitempty"`
	// DNS Health check transparent flag.
	Transparent SlbCurAdvhcDnsTableTransparent `json:"Transparent,omitempty"`
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
	Invert SlbCurAdvhcDnsTableInvert `json:"Invert,omitempty"`
	// DNS Health check domain name parameter.
	DomainName string `json:"DomainName,omitempty"`
	// DNS Health check transport layer.
	Transport SlbCurAdvhcDnsTableTransport `json:"Transport,omitempty"`
}

func (p SlbCurAdvhcDnsTableParams) iMABean() {}
