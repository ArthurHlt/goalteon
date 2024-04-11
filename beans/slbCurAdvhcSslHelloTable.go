package beans

import (
	"fmt"
	"reflect"
)

// SlbCurAdvhcSslHelloTable Note:This mib is not supported on VX instance of Virtualization.
type SlbCurAdvhcSslHelloTable struct {
	// SSL Hello Health check id.
	SlbCurAdvhcSslHelloID string
	Params                *SlbCurAdvhcSslHelloTableParams
}

func NewSlbCurAdvhcSslHelloTableList() *SlbCurAdvhcSslHelloTable {
	return &SlbCurAdvhcSslHelloTable{}
}

func NewSlbCurAdvhcSslHelloTable(
	slbCurAdvhcSslHelloID string,
	params *SlbCurAdvhcSslHelloTableParams,
) *SlbCurAdvhcSslHelloTable {
	return &SlbCurAdvhcSslHelloTable{
		SlbCurAdvhcSslHelloID: slbCurAdvhcSslHelloID,
		Params:                params,
	}
}

func (c *SlbCurAdvhcSslHelloTable) Name() string {
	return "SlbCurAdvhcSslHelloTable"
}

func (c *SlbCurAdvhcSslHelloTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbCurAdvhcSslHelloTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbCurAdvhcSslHelloTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbCurAdvhcSslHelloID).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbCurAdvhcSslHelloID)
}

type SlbCurAdvhcSslHelloTableIPVer int32

const (
	SlbCurAdvhcSslHelloTableIPVer_Ipv4 SlbCurAdvhcSslHelloTableIPVer = 1
	SlbCurAdvhcSslHelloTableIPVer_Ipv6 SlbCurAdvhcSslHelloTableIPVer = 2
	SlbCurAdvhcSslHelloTableIPVer_None SlbCurAdvhcSslHelloTableIPVer = 3
)

type SlbCurAdvhcSslHelloTableTransparent int32

const (
	SlbCurAdvhcSslHelloTableTransparent_Enabled     SlbCurAdvhcSslHelloTableTransparent = 1
	SlbCurAdvhcSslHelloTableTransparent_Disabled    SlbCurAdvhcSslHelloTableTransparent = 2
	SlbCurAdvhcSslHelloTableTransparent_Unsupported SlbCurAdvhcSslHelloTableTransparent = 2147483647
)

type SlbCurAdvhcSslHelloTableInvert int32

const (
	SlbCurAdvhcSslHelloTableInvert_Enabled     SlbCurAdvhcSslHelloTableInvert = 1
	SlbCurAdvhcSslHelloTableInvert_Disabled    SlbCurAdvhcSslHelloTableInvert = 2
	SlbCurAdvhcSslHelloTableInvert_Unsupported SlbCurAdvhcSslHelloTableInvert = 2147483647
)

type SlbCurAdvhcSslHelloTableSslVersion int32

const (
	SlbCurAdvhcSslHelloTableSslVersion_Ver2        SlbCurAdvhcSslHelloTableSslVersion = 2
	SlbCurAdvhcSslHelloTableSslVersion_Ver3        SlbCurAdvhcSslHelloTableSslVersion = 3
	SlbCurAdvhcSslHelloTableSslVersion_Tls         SlbCurAdvhcSslHelloTableSslVersion = 4
	SlbCurAdvhcSslHelloTableSslVersion_Unsupported SlbCurAdvhcSslHelloTableSslVersion = 2147483647
)

type SlbCurAdvhcSslHelloTableCipherName int32

const (
	SlbCurAdvhcSslHelloTableCipherName_UserDefined SlbCurAdvhcSslHelloTableCipherName = 1
	SlbCurAdvhcSslHelloTableCipherName_Low         SlbCurAdvhcSslHelloTableCipherName = 2
	SlbCurAdvhcSslHelloTableCipherName_Medium      SlbCurAdvhcSslHelloTableCipherName = 3
	SlbCurAdvhcSslHelloTableCipherName_High        SlbCurAdvhcSslHelloTableCipherName = 4
	SlbCurAdvhcSslHelloTableCipherName_Unsupported SlbCurAdvhcSslHelloTableCipherName = 2147483647
)

type SlbCurAdvhcSslHelloTableParams struct {
	// SSL Hello Health check id.
	ID string `json:"ID,omitempty"`
	// SSL Hello Health check name.
	Name string `json:"Name,omitempty"`
	// SSL Hello Health check destination port.
	DPort uint64 `json:"DPort,omitempty"`
	// SSL Hello Health check destination IP version.
	IPVer SlbCurAdvhcSslHelloTableIPVer `json:"IPVer,omitempty"`
	// SSL Hello Health check destination hostname.
	HostName string `json:"HostName,omitempty"`
	// SSL Hello Health check transparent flag.
	Transparent SlbCurAdvhcSslHelloTableTransparent `json:"Transparent,omitempty"`
	// SSL Hello Health check interval.
	Interval uint64 `json:"Interval,omitempty"`
	// SSL Hello Health check retries counter.
	Retries uint64 `json:"Retries,omitempty"`
	// SSL Hello Health check retries counter.
	RestoreRetries uint64 `json:"RestoreRetries,omitempty"`
	// SSL Hello Health check timeout.
	Timeout uint64 `json:"Timeout,omitempty"`
	// SSL Hello Health check overflow flag.
	Overflow uint64 `json:"Overflow,omitempty"`
	// SSL Hello Health check interval in the down state.
	DownInterval uint64 `json:"DownInterval,omitempty"`
	// SSL Hello Health check invert flag.
	Invert SlbCurAdvhcSslHelloTableInvert `json:"Invert,omitempty"`
	// SSL Hello Health check SSL version field.
	SslVersion SlbCurAdvhcSslHelloTableSslVersion `json:"SslVersion,omitempty"`
	// Cipher name for SSLHELLO HC Context.
	CipherName SlbCurAdvhcSslHelloTableCipherName `json:"CipherName,omitempty"`
	// Cipher-suite allowed for SSLHELLO HC Context.
	CipherUserdef string `json:"CipherUserdef,omitempty"`
}

func (p SlbCurAdvhcSslHelloTableParams) iMABean() {}
