package beans

import (
	"fmt"
	"reflect"
)

// SlbNewAdvhcSipTable Note:This mib is not supported on VX instance of Virtualization.
type SlbNewAdvhcSipTable struct {
	// SIP Health check id .
	SlbNewAdvhcSipID string
	Params           *SlbNewAdvhcSipTableParams
}

func NewSlbNewAdvhcSipTableList() *SlbNewAdvhcSipTable {
	return &SlbNewAdvhcSipTable{}
}

func NewSlbNewAdvhcSipTable(
	slbNewAdvhcSipID string,
	params *SlbNewAdvhcSipTableParams,
) *SlbNewAdvhcSipTable {
	return &SlbNewAdvhcSipTable{
		SlbNewAdvhcSipID: slbNewAdvhcSipID,
		Params:           params,
	}
}

func (c *SlbNewAdvhcSipTable) Name() string {
	return "SlbNewAdvhcSipTable"
}

func (c *SlbNewAdvhcSipTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbNewAdvhcSipTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbNewAdvhcSipTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbNewAdvhcSipID).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbNewAdvhcSipID)
}

type SlbNewAdvhcSipTableIPVer int32

const (
	SlbNewAdvhcSipTableIPVer_Ipv4 SlbNewAdvhcSipTableIPVer = 1
	SlbNewAdvhcSipTableIPVer_Ipv6 SlbNewAdvhcSipTableIPVer = 2
	SlbNewAdvhcSipTableIPVer_None SlbNewAdvhcSipTableIPVer = 3
)

type SlbNewAdvhcSipTableTransparent int32

const (
	SlbNewAdvhcSipTableTransparent_Enabled     SlbNewAdvhcSipTableTransparent = 1
	SlbNewAdvhcSipTableTransparent_Disabled    SlbNewAdvhcSipTableTransparent = 2
	SlbNewAdvhcSipTableTransparent_Unsupported SlbNewAdvhcSipTableTransparent = 2147483647
)

type SlbNewAdvhcSipTableInvert int32

const (
	SlbNewAdvhcSipTableInvert_Enabled     SlbNewAdvhcSipTableInvert = 1
	SlbNewAdvhcSipTableInvert_Disabled    SlbNewAdvhcSipTableInvert = 2
	SlbNewAdvhcSipTableInvert_Unsupported SlbNewAdvhcSipTableInvert = 2147483647
)

type SlbNewAdvhcSipTableMethod int32

const (
	SlbNewAdvhcSipTableMethod_Options     SlbNewAdvhcSipTableMethod = 1
	SlbNewAdvhcSipTableMethod_Ping        SlbNewAdvhcSipTableMethod = 2
	SlbNewAdvhcSipTableMethod_Unsupported SlbNewAdvhcSipTableMethod = 2147483647
)

type SlbNewAdvhcSipTableTransport int32

const (
	SlbNewAdvhcSipTableTransport_Udp         SlbNewAdvhcSipTableTransport = 1
	SlbNewAdvhcSipTableTransport_Tcp         SlbNewAdvhcSipTableTransport = 2
	SlbNewAdvhcSipTableTransport_Unsupported SlbNewAdvhcSipTableTransport = 2147483647
)

type SlbNewAdvhcSipTableDelete int32

const (
	SlbNewAdvhcSipTableDelete_Other       SlbNewAdvhcSipTableDelete = 1
	SlbNewAdvhcSipTableDelete_Delete      SlbNewAdvhcSipTableDelete = 2
	SlbNewAdvhcSipTableDelete_Unsupported SlbNewAdvhcSipTableDelete = 2147483647
)

type SlbNewAdvhcSipTableSnat int32

const (
	SlbNewAdvhcSipTableSnat_Enabled     SlbNewAdvhcSipTableSnat = 1
	SlbNewAdvhcSipTableSnat_Disabled    SlbNewAdvhcSipTableSnat = 2
	SlbNewAdvhcSipTableSnat_Unsupported SlbNewAdvhcSipTableSnat = 2147483647
)

type SlbNewAdvhcSipTableSips int32

const (
	SlbNewAdvhcSipTableSips_Enabled     SlbNewAdvhcSipTableSips = 1
	SlbNewAdvhcSipTableSips_Disabled    SlbNewAdvhcSipTableSips = 2
	SlbNewAdvhcSipTableSips_Unsupported SlbNewAdvhcSipTableSips = 2147483647
)

type SlbNewAdvhcSipTableParams struct {
	// SIP Health check id .
	ID string `json:"ID,omitempty"`
	// SIP Health check name.
	Name string `json:"Name,omitempty"`
	// SIP Health check destination port.
	DPort uint64 `json:"DPort,omitempty"`
	// SIP Health check destination IP version.
	IPVer SlbNewAdvhcSipTableIPVer `json:"IPVer,omitempty"`
	// SIP Health check destination hostname.
	HostName string `json:"HostName,omitempty"`
	// SIP Health check transparent flag.
	Transparent SlbNewAdvhcSipTableTransparent `json:"Transparent,omitempty"`
	// SIP Health check interval.
	Interval uint64 `json:"Interval,omitempty"`
	// SIP Health check retries counter.
	Retries uint64 `json:"Retries,omitempty"`
	// SIP Health check retries when in the down state.
	RestoreRetries uint64 `json:"RestoreRetries,omitempty"`
	// SIP Health check timeout.
	Timeout uint64 `json:"Timeout,omitempty"`
	// SIP Health check overflow flag.
	Overflow uint64 `json:"Overflow,omitempty"`
	// SIP Health check interval when in down state.
	DownInterval uint64 `json:"DownInterval,omitempty"`
	// SIP Health check invert flag.
	Invert SlbNewAdvhcSipTableInvert `json:"Invert,omitempty"`
	// SIP Health check method.
	Method SlbNewAdvhcSipTableMethod `json:"Method,omitempty"`
	// SIP Health check transport.
	Transport SlbNewAdvhcSipTableTransport `json:"Transport,omitempty"`
	// SIP Health check URI.
	RequestUri string `json:"RequestUri,omitempty"`
	// SIP Health check 'from' string.
	From string `json:"From,omitempty"`
	// SIP Health check expected response code.
	ResponseCodes string `json:"ResponseCodes,omitempty"`
	// SIP Health check copy flag.
	Copy DisplayString `json:"Copy,omitempty"`
	// When set to the value of 2 (delete), the entire row is deleted.
	// When read, other(1) is returned. Setting the value to anything
	// other than 2(delete) has no effect on the state of the row.
	Delete SlbNewAdvhcSipTableDelete `json:"Delete,omitempty"`
	// SIP Health Check src NAT (PIP) flag.
	Snat SlbNewAdvhcSipTableSnat `json:"Snat,omitempty"`
	// SIP Health check SIPS enable/disable flag.
	Sips SlbNewAdvhcSipTableSips `json:"Sips,omitempty"`
	// SIP Health check Cipher-suite allowed for SSL for SIPS Context.
	SipsCipherUserdef string `json:"SipsCipherUserdef,omitempty"`
}

func (p SlbNewAdvhcSipTableParams) iMABean() {}
