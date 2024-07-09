package beans

import (
	"fmt"
	"reflect"
)

// SlbCurAdvhcSipTable Note:This mib is not supported on VX instance of Virtualization.
type SlbCurAdvhcSipTable struct {
	// SIP Health check id.
	SlbCurAdvhcSipID string
	Params           *SlbCurAdvhcSipTableParams
}

func NewSlbCurAdvhcSipTableList() *SlbCurAdvhcSipTable {
	return &SlbCurAdvhcSipTable{}
}

func NewSlbCurAdvhcSipTable(
	slbCurAdvhcSipID string,
	params *SlbCurAdvhcSipTableParams,
) *SlbCurAdvhcSipTable {
	return &SlbCurAdvhcSipTable{
		SlbCurAdvhcSipID: slbCurAdvhcSipID,
		Params:           params,
	}
}

func (c *SlbCurAdvhcSipTable) Name() string {
	return "SlbCurAdvhcSipTable"
}

func (c *SlbCurAdvhcSipTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbCurAdvhcSipTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbCurAdvhcSipTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbCurAdvhcSipID).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbCurAdvhcSipID)
}

type SlbCurAdvhcSipTableIPVer int32

const (
	SlbCurAdvhcSipTableIPVer_Ipv4 SlbCurAdvhcSipTableIPVer = 1
	SlbCurAdvhcSipTableIPVer_Ipv6 SlbCurAdvhcSipTableIPVer = 2
	SlbCurAdvhcSipTableIPVer_None SlbCurAdvhcSipTableIPVer = 3
)

type SlbCurAdvhcSipTableTransparent int32

const (
	SlbCurAdvhcSipTableTransparent_Enabled     SlbCurAdvhcSipTableTransparent = 1
	SlbCurAdvhcSipTableTransparent_Disabled    SlbCurAdvhcSipTableTransparent = 2
	SlbCurAdvhcSipTableTransparent_Unsupported SlbCurAdvhcSipTableTransparent = 2147483647
)

type SlbCurAdvhcSipTableInvert int32

const (
	SlbCurAdvhcSipTableInvert_Enabled     SlbCurAdvhcSipTableInvert = 1
	SlbCurAdvhcSipTableInvert_Disabled    SlbCurAdvhcSipTableInvert = 2
	SlbCurAdvhcSipTableInvert_Unsupported SlbCurAdvhcSipTableInvert = 2147483647
)

type SlbCurAdvhcSipTableMethod int32

const (
	SlbCurAdvhcSipTableMethod_Options     SlbCurAdvhcSipTableMethod = 1
	SlbCurAdvhcSipTableMethod_Ping        SlbCurAdvhcSipTableMethod = 2
	SlbCurAdvhcSipTableMethod_Unsupported SlbCurAdvhcSipTableMethod = 2147483647
)

type SlbCurAdvhcSipTableTransport int32

const (
	SlbCurAdvhcSipTableTransport_Udp         SlbCurAdvhcSipTableTransport = 1
	SlbCurAdvhcSipTableTransport_Tcp         SlbCurAdvhcSipTableTransport = 2
	SlbCurAdvhcSipTableTransport_Unsupported SlbCurAdvhcSipTableTransport = 2147483647
)

type SlbCurAdvhcSipTableSnat int32

const (
	SlbCurAdvhcSipTableSnat_Enabled     SlbCurAdvhcSipTableSnat = 1
	SlbCurAdvhcSipTableSnat_Disabled    SlbCurAdvhcSipTableSnat = 2
	SlbCurAdvhcSipTableSnat_Unsupported SlbCurAdvhcSipTableSnat = 2147483647
)

type SlbCurAdvhcSipTableSips int32

const (
	SlbCurAdvhcSipTableSips_Enabled     SlbCurAdvhcSipTableSips = 1
	SlbCurAdvhcSipTableSips_Disabled    SlbCurAdvhcSipTableSips = 2
	SlbCurAdvhcSipTableSips_Unsupported SlbCurAdvhcSipTableSips = 2147483647
)

type SlbCurAdvhcSipTableParams struct {
	// SIP Health check id.
	ID string `json:"ID,omitempty"`
	// SIP Health check name.
	Name string `json:"Name,omitempty"`
	// SIP Health check destination port.
	DPort uint64 `json:"DPort,omitempty"`
	// SIP Health check destination IP version.
	IPVer SlbCurAdvhcSipTableIPVer `json:"IPVer,omitempty"`
	// SIP Health check destination hostname.
	HostName string `json:"HostName,omitempty"`
	// SIP Health check transparent flag.
	Transparent SlbCurAdvhcSipTableTransparent `json:"Transparent,omitempty"`
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
	Invert SlbCurAdvhcSipTableInvert `json:"Invert,omitempty"`
	// SIP Health check method.
	Method SlbCurAdvhcSipTableMethod `json:"Method,omitempty"`
	// SIP Health check transport.
	Transport SlbCurAdvhcSipTableTransport `json:"Transport,omitempty"`
	// SIP Health check URI.
	RequestUri string `json:"RequestUri,omitempty"`
	// SIP Health check 'from' string.
	From string `json:"From,omitempty"`
	// SIP Health check expected response code.
	ResponseCodes string `json:"ResponseCodes,omitempty"`
	// SIP Health Check src NAT (PIP) flag.
	Snat SlbCurAdvhcSipTableSnat `json:"Snat,omitempty"`
	// SIP Health check SIPS enable/disable flag.
	Sips SlbCurAdvhcSipTableSips `json:"Sips,omitempty"`
	// SIP Health Check Cipher-suite allowed for SSL for SIPS Context.
	SipsCipherUserdef string `json:"SipsCipherUserdef,omitempty"`
}

func (p SlbCurAdvhcSipTableParams) iMABean() {}
