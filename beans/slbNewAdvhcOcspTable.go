package beans

import (
	"fmt"
	"reflect"
)

// SlbNewAdvhcOcspTable Note:This mib is not supported on VX instance of Virtualization.
type SlbNewAdvhcOcspTable struct {
	// OCSP Health check id.
	SlbNewAdvhcOcspID string
	Params            *SlbNewAdvhcOcspTableParams
}

func NewSlbNewAdvhcOcspTableList() *SlbNewAdvhcOcspTable {
	return &SlbNewAdvhcOcspTable{}
}

func NewSlbNewAdvhcOcspTable(
	slbNewAdvhcOcspID string,
	params *SlbNewAdvhcOcspTableParams,
) *SlbNewAdvhcOcspTable {
	return &SlbNewAdvhcOcspTable{
		SlbNewAdvhcOcspID: slbNewAdvhcOcspID,
		Params:            params,
	}
}

func (c *SlbNewAdvhcOcspTable) Name() string {
	return "SlbNewAdvhcOcspTable"
}

func (c *SlbNewAdvhcOcspTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbNewAdvhcOcspTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbNewAdvhcOcspTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbNewAdvhcOcspID).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbNewAdvhcOcspID)
}

type SlbNewAdvhcOcspTableIPVer int32

const (
	SlbNewAdvhcOcspTableIPVer_Ipv4 SlbNewAdvhcOcspTableIPVer = 1
	SlbNewAdvhcOcspTableIPVer_Ipv6 SlbNewAdvhcOcspTableIPVer = 2
	SlbNewAdvhcOcspTableIPVer_None SlbNewAdvhcOcspTableIPVer = 3
)

type SlbNewAdvhcOcspTableTransparent int32

const (
	SlbNewAdvhcOcspTableTransparent_Enabled     SlbNewAdvhcOcspTableTransparent = 1
	SlbNewAdvhcOcspTableTransparent_Disabled    SlbNewAdvhcOcspTableTransparent = 2
	SlbNewAdvhcOcspTableTransparent_Unsupported SlbNewAdvhcOcspTableTransparent = 2147483647
)

type SlbNewAdvhcOcspTableInvert int32

const (
	SlbNewAdvhcOcspTableInvert_Enabled     SlbNewAdvhcOcspTableInvert = 1
	SlbNewAdvhcOcspTableInvert_Disabled    SlbNewAdvhcOcspTableInvert = 2
	SlbNewAdvhcOcspTableInvert_Unsupported SlbNewAdvhcOcspTableInvert = 2147483647
)

type SlbNewAdvhcOcspTableHttps int32

const (
	SlbNewAdvhcOcspTableHttps_Enabled     SlbNewAdvhcOcspTableHttps = 1
	SlbNewAdvhcOcspTableHttps_Disabled    SlbNewAdvhcOcspTableHttps = 2
	SlbNewAdvhcOcspTableHttps_Unsupported SlbNewAdvhcOcspTableHttps = 2147483647
)

type SlbNewAdvhcOcspTableGoodCert int32

const (
	SlbNewAdvhcOcspTableGoodCert_Enabled     SlbNewAdvhcOcspTableGoodCert = 1
	SlbNewAdvhcOcspTableGoodCert_Disabled    SlbNewAdvhcOcspTableGoodCert = 2
	SlbNewAdvhcOcspTableGoodCert_Unsupported SlbNewAdvhcOcspTableGoodCert = 2147483647
)

type SlbNewAdvhcOcspTableDelete int32

const (
	SlbNewAdvhcOcspTableDelete_Other       SlbNewAdvhcOcspTableDelete = 1
	SlbNewAdvhcOcspTableDelete_Delete      SlbNewAdvhcOcspTableDelete = 2
	SlbNewAdvhcOcspTableDelete_Unsupported SlbNewAdvhcOcspTableDelete = 2147483647
)

type SlbNewAdvhcOcspTableConnTerm int32

const (
	SlbNewAdvhcOcspTableConnTerm_Fin         SlbNewAdvhcOcspTableConnTerm = 1
	SlbNewAdvhcOcspTableConnTerm_Rst         SlbNewAdvhcOcspTableConnTerm = 2
	SlbNewAdvhcOcspTableConnTerm_Unsupported SlbNewAdvhcOcspTableConnTerm = 2147483647
)

type SlbNewAdvhcOcspTableAlways int32

const (
	SlbNewAdvhcOcspTableAlways_Enabled     SlbNewAdvhcOcspTableAlways = 1
	SlbNewAdvhcOcspTableAlways_Disabled    SlbNewAdvhcOcspTableAlways = 2
	SlbNewAdvhcOcspTableAlways_Unsupported SlbNewAdvhcOcspTableAlways = 2147483647
)

type SlbNewAdvhcOcspTableSnat int32

const (
	SlbNewAdvhcOcspTableSnat_Enabled     SlbNewAdvhcOcspTableSnat = 1
	SlbNewAdvhcOcspTableSnat_Disabled    SlbNewAdvhcOcspTableSnat = 2
	SlbNewAdvhcOcspTableSnat_Unsupported SlbNewAdvhcOcspTableSnat = 2147483647
)

type SlbNewAdvhcOcspTableMethod int32

const (
	SlbNewAdvhcOcspTableMethod_Get         SlbNewAdvhcOcspTableMethod = 1
	SlbNewAdvhcOcspTableMethod_Post        SlbNewAdvhcOcspTableMethod = 2
	SlbNewAdvhcOcspTableMethod_Unsupported SlbNewAdvhcOcspTableMethod = 2147483647
)

type SlbNewAdvhcOcspTableHttpsCipherName int32

const (
	SlbNewAdvhcOcspTableHttpsCipherName_UserDefined SlbNewAdvhcOcspTableHttpsCipherName = 1
	SlbNewAdvhcOcspTableHttpsCipherName_Low         SlbNewAdvhcOcspTableHttpsCipherName = 2
	SlbNewAdvhcOcspTableHttpsCipherName_Medium      SlbNewAdvhcOcspTableHttpsCipherName = 3
	SlbNewAdvhcOcspTableHttpsCipherName_High        SlbNewAdvhcOcspTableHttpsCipherName = 4
	SlbNewAdvhcOcspTableHttpsCipherName_Unsupported SlbNewAdvhcOcspTableHttpsCipherName = 2147483647
)

type SlbNewAdvhcOcspTableParams struct {
	// OCSP Health check id.
	ID string `json:"ID,omitempty"`
	// OCSP Health check name.
	Name string `json:"Name,omitempty"`
	// OCSP Health check destination port.
	DPort uint64 `json:"DPort,omitempty"`
	// OCSP Health check destination IP version.
	IPVer SlbNewAdvhcOcspTableIPVer `json:"IPVer,omitempty"`
	// OCSP Health check destination hostname.
	HostName string `json:"HostName,omitempty"`
	// OCSP Health check transparent flag.
	Transparent SlbNewAdvhcOcspTableTransparent `json:"Transparent,omitempty"`
	// OCSP Health check interval.
	Interval uint64 `json:"Interval,omitempty"`
	// OCSP Health check retries counter.
	Retries uint64 `json:"Retries,omitempty"`
	// OCSP Health check retries in down state counter.
	RestoreRetries uint64 `json:"RestoreRetries,omitempty"`
	// OCSP Health check timeout.
	Timeout uint64 `json:"Timeout,omitempty"`
	// OCSP Health check overflow flag.
	Overflow uint64 `json:"Overflow,omitempty"`
	// OCSP Health check interval in down state.
	DownInterval uint64 `json:"DownInterval,omitempty"`
	// OCSP Health check invert flag.
	Invert SlbNewAdvhcOcspTableInvert `json:"Invert,omitempty"`
	// OCSP Health check HTTPS enable/disable flag.
	Https SlbNewAdvhcOcspTableHttps `json:"Https,omitempty"`
	// OCSP Health check good certificate check enable/disable flag.
	GoodCert SlbNewAdvhcOcspTableGoodCert `json:"GoodCert,omitempty"`
	// OCSP Health check certificate id.
	Cert string `json:"Cert,omitempty"`
	// OCSP Health check Tusted CA certificate id.
	TrustedCA string `json:"TrustedCA,omitempty"`
	// OCSP Health check copy action trigger.
	Copy DisplayString `json:"Copy,omitempty"`
	// When set to the value of 2 (delete), the entire row is deleted.
	// When read, other(1) is returned. Setting the value to anything
	// other than 2(delete) has no effect on the state of the row.
	Delete SlbNewAdvhcOcspTableDelete `json:"Delete,omitempty"`
	// OCSP Health check Connection termination type.
	ConnTerm SlbNewAdvhcOcspTableConnTerm `json:"ConnTerm,omitempty"`
	// OCSP Health check is allowed for standalone real enable/disable flag.
	Always SlbNewAdvhcOcspTableAlways `json:"Always,omitempty"`
	// OCSP Health Check src NAT (PIP) flag.
	Snat SlbNewAdvhcOcspTableSnat `json:"Snat,omitempty"`
	// OCSP Health check host field.
	Host string `json:"Host,omitempty"`
	// OCSP Health check path field.
	Path string `json:"Path,omitempty"`
	// OCSP Health check HTTP method.
	Method SlbNewAdvhcOcspTableMethod `json:"Method,omitempty"`
	// OCSP Health check Cipher name for SSL for HTTPS Context.
	HttpsCipherName SlbNewAdvhcOcspTableHttpsCipherName `json:"HttpsCipherName,omitempty"`
	// OCSP Health check Cipher-suite allowed for SSL for HTTPS Context.
	HttpsCipherUserdef string `json:"HttpsCipherUserdef,omitempty"`
}

func (p SlbNewAdvhcOcspTableParams) iMABean() {}
