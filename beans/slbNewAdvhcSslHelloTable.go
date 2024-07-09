package beans

import (
	"fmt"
	"reflect"
)

// SlbNewAdvhcSslHelloTable Note:This mib is not supported on VX instance of Virtualization.
type SlbNewAdvhcSslHelloTable struct {
	// SSL Hello Health check id.
	SlbNewAdvhcSslHelloID string
	Params                *SlbNewAdvhcSslHelloTableParams
}

func NewSlbNewAdvhcSslHelloTableList() *SlbNewAdvhcSslHelloTable {
	return &SlbNewAdvhcSslHelloTable{}
}

func NewSlbNewAdvhcSslHelloTable(
	slbNewAdvhcSslHelloID string,
	params *SlbNewAdvhcSslHelloTableParams,
) *SlbNewAdvhcSslHelloTable {
	return &SlbNewAdvhcSslHelloTable{
		SlbNewAdvhcSslHelloID: slbNewAdvhcSslHelloID,
		Params:                params,
	}
}

func (c *SlbNewAdvhcSslHelloTable) Name() string {
	return "SlbNewAdvhcSslHelloTable"
}

func (c *SlbNewAdvhcSslHelloTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbNewAdvhcSslHelloTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbNewAdvhcSslHelloTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbNewAdvhcSslHelloID).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbNewAdvhcSslHelloID)
}

type SlbNewAdvhcSslHelloTableIPVer int32

const (
	SlbNewAdvhcSslHelloTableIPVer_Ipv4 SlbNewAdvhcSslHelloTableIPVer = 1
	SlbNewAdvhcSslHelloTableIPVer_Ipv6 SlbNewAdvhcSslHelloTableIPVer = 2
	SlbNewAdvhcSslHelloTableIPVer_None SlbNewAdvhcSslHelloTableIPVer = 3
)

type SlbNewAdvhcSslHelloTableTransparent int32

const (
	SlbNewAdvhcSslHelloTableTransparent_Enabled     SlbNewAdvhcSslHelloTableTransparent = 1
	SlbNewAdvhcSslHelloTableTransparent_Disabled    SlbNewAdvhcSslHelloTableTransparent = 2
	SlbNewAdvhcSslHelloTableTransparent_Unsupported SlbNewAdvhcSslHelloTableTransparent = 2147483647
)

type SlbNewAdvhcSslHelloTableInvert int32

const (
	SlbNewAdvhcSslHelloTableInvert_Enabled     SlbNewAdvhcSslHelloTableInvert = 1
	SlbNewAdvhcSslHelloTableInvert_Disabled    SlbNewAdvhcSslHelloTableInvert = 2
	SlbNewAdvhcSslHelloTableInvert_Unsupported SlbNewAdvhcSslHelloTableInvert = 2147483647
)

type SlbNewAdvhcSslHelloTableSslVersion int32

const (
	SlbNewAdvhcSslHelloTableSslVersion_Ver2        SlbNewAdvhcSslHelloTableSslVersion = 2
	SlbNewAdvhcSslHelloTableSslVersion_Ver3        SlbNewAdvhcSslHelloTableSslVersion = 3
	SlbNewAdvhcSslHelloTableSslVersion_Tls         SlbNewAdvhcSslHelloTableSslVersion = 4
	SlbNewAdvhcSslHelloTableSslVersion_Unsupported SlbNewAdvhcSslHelloTableSslVersion = 2147483647
)

type SlbNewAdvhcSslHelloTableDelete int32

const (
	SlbNewAdvhcSslHelloTableDelete_Other       SlbNewAdvhcSslHelloTableDelete = 1
	SlbNewAdvhcSslHelloTableDelete_Delete      SlbNewAdvhcSslHelloTableDelete = 2
	SlbNewAdvhcSslHelloTableDelete_Unsupported SlbNewAdvhcSslHelloTableDelete = 2147483647
)

type SlbNewAdvhcSslHelloTableCipherName int32

const (
	SlbNewAdvhcSslHelloTableCipherName_UserDefined SlbNewAdvhcSslHelloTableCipherName = 1
	SlbNewAdvhcSslHelloTableCipherName_Low         SlbNewAdvhcSslHelloTableCipherName = 2
	SlbNewAdvhcSslHelloTableCipherName_Medium      SlbNewAdvhcSslHelloTableCipherName = 3
	SlbNewAdvhcSslHelloTableCipherName_High        SlbNewAdvhcSslHelloTableCipherName = 4
	SlbNewAdvhcSslHelloTableCipherName_Unsupported SlbNewAdvhcSslHelloTableCipherName = 2147483647
)

type SlbNewAdvhcSslHelloTableSnat int32

const (
	SlbNewAdvhcSslHelloTableSnat_Enabled     SlbNewAdvhcSslHelloTableSnat = 1
	SlbNewAdvhcSslHelloTableSnat_Disabled    SlbNewAdvhcSslHelloTableSnat = 2
	SlbNewAdvhcSslHelloTableSnat_Unsupported SlbNewAdvhcSslHelloTableSnat = 2147483647
)

type SlbNewAdvhcSslHelloTableParams struct {
	// SSL Hello Health check id.
	ID string `json:"ID,omitempty"`
	// SSL Hello Health check name.
	Name string `json:"Name,omitempty"`
	// SSL Hello Health check destination port.
	DPort uint64 `json:"DPort,omitempty"`
	// SSL Hello Health check destination IP version.
	IPVer SlbNewAdvhcSslHelloTableIPVer `json:"IPVer,omitempty"`
	// SSL Hello Health check destination hostname.
	HostName string `json:"HostName,omitempty"`
	// SSL Hello Health check transparent flag.
	Transparent SlbNewAdvhcSslHelloTableTransparent `json:"Transparent,omitempty"`
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
	Invert SlbNewAdvhcSslHelloTableInvert `json:"Invert,omitempty"`
	// SSL Hello Health check SSL version field.
	SslVersion SlbNewAdvhcSslHelloTableSslVersion `json:"SslVersion,omitempty"`
	// SSL Hello Health check copy indicator.
	Copy string `json:"Copy,omitempty"`
	// When set to the value of 2 (delete), the entire row is deleted.
	// When read, other(1) is returned. Setting the value to anything
	// other than 2(delete) has no effect on the state of the row.
	Delete SlbNewAdvhcSslHelloTableDelete `json:"Delete,omitempty"`
	// Cipher name for SSLHELLO HC Context.
	CipherName SlbNewAdvhcSslHelloTableCipherName `json:"CipherName,omitempty"`
	// Cipher-suite allowed for SSLHELLO HC Context.
	CipherUserdef string `json:"CipherUserdef,omitempty"`
	// SSL Hello Health Check src NAT (PIP) flag.
	Snat SlbNewAdvhcSslHelloTableSnat `json:"Snat,omitempty"`
}

func (p SlbNewAdvhcSslHelloTableParams) iMABean() {}
