package beans

import (
	"fmt"
	"reflect"
)

// SlbCurAdvhcOcspTable Note:This mib is not supported on VX instance of Virtualization.
type SlbCurAdvhcOcspTable struct {
	// OCSP Health check id.
	SlbCurAdvhcOcspID string
	Params            *SlbCurAdvhcOcspTableParams
}

func NewSlbCurAdvhcOcspTableList() *SlbCurAdvhcOcspTable {
	return &SlbCurAdvhcOcspTable{}
}

func NewSlbCurAdvhcOcspTable(
	slbCurAdvhcOcspID string,
	params *SlbCurAdvhcOcspTableParams,
) *SlbCurAdvhcOcspTable {
	return &SlbCurAdvhcOcspTable{
		SlbCurAdvhcOcspID: slbCurAdvhcOcspID,
		Params:            params,
	}
}

func (c *SlbCurAdvhcOcspTable) Name() string {
	return "SlbCurAdvhcOcspTable"
}

func (c *SlbCurAdvhcOcspTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbCurAdvhcOcspTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbCurAdvhcOcspTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbCurAdvhcOcspID).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbCurAdvhcOcspID)
}

type SlbCurAdvhcOcspTableIPVer int32

const (
	SlbCurAdvhcOcspTableIPVer_Ipv4 SlbCurAdvhcOcspTableIPVer = 1
	SlbCurAdvhcOcspTableIPVer_Ipv6 SlbCurAdvhcOcspTableIPVer = 2
	SlbCurAdvhcOcspTableIPVer_None SlbCurAdvhcOcspTableIPVer = 3
)

type SlbCurAdvhcOcspTableTransparent int32

const (
	SlbCurAdvhcOcspTableTransparent_Enabled     SlbCurAdvhcOcspTableTransparent = 1
	SlbCurAdvhcOcspTableTransparent_Disabled    SlbCurAdvhcOcspTableTransparent = 2
	SlbCurAdvhcOcspTableTransparent_Unsupported SlbCurAdvhcOcspTableTransparent = 2147483647
)

type SlbCurAdvhcOcspTableInvert int32

const (
	SlbCurAdvhcOcspTableInvert_Enabled     SlbCurAdvhcOcspTableInvert = 1
	SlbCurAdvhcOcspTableInvert_Disabled    SlbCurAdvhcOcspTableInvert = 2
	SlbCurAdvhcOcspTableInvert_Unsupported SlbCurAdvhcOcspTableInvert = 2147483647
)

type SlbCurAdvhcOcspTableHttps int32

const (
	SlbCurAdvhcOcspTableHttps_Enabled     SlbCurAdvhcOcspTableHttps = 1
	SlbCurAdvhcOcspTableHttps_Disabled    SlbCurAdvhcOcspTableHttps = 2
	SlbCurAdvhcOcspTableHttps_Unsupported SlbCurAdvhcOcspTableHttps = 2147483647
)

type SlbCurAdvhcOcspTableGoodCert int32

const (
	SlbCurAdvhcOcspTableGoodCert_Enabled     SlbCurAdvhcOcspTableGoodCert = 1
	SlbCurAdvhcOcspTableGoodCert_Disabled    SlbCurAdvhcOcspTableGoodCert = 2
	SlbCurAdvhcOcspTableGoodCert_Unsupported SlbCurAdvhcOcspTableGoodCert = 2147483647
)

type SlbCurAdvhcOcspTableConnTerm int32

const (
	SlbCurAdvhcOcspTableConnTerm_Fin         SlbCurAdvhcOcspTableConnTerm = 1
	SlbCurAdvhcOcspTableConnTerm_Rst         SlbCurAdvhcOcspTableConnTerm = 2
	SlbCurAdvhcOcspTableConnTerm_Unsupported SlbCurAdvhcOcspTableConnTerm = 2147483647
)

type SlbCurAdvhcOcspTableAlways int32

const (
	SlbCurAdvhcOcspTableAlways_Enabled     SlbCurAdvhcOcspTableAlways = 1
	SlbCurAdvhcOcspTableAlways_Disabled    SlbCurAdvhcOcspTableAlways = 2
	SlbCurAdvhcOcspTableAlways_Unsupported SlbCurAdvhcOcspTableAlways = 2147483647
)

type SlbCurAdvhcOcspTableSnat int32

const (
	SlbCurAdvhcOcspTableSnat_Enabled     SlbCurAdvhcOcspTableSnat = 1
	SlbCurAdvhcOcspTableSnat_Disabled    SlbCurAdvhcOcspTableSnat = 2
	SlbCurAdvhcOcspTableSnat_Unsupported SlbCurAdvhcOcspTableSnat = 2147483647
)

type SlbCurAdvhcOcspTableMethod int32

const (
	SlbCurAdvhcOcspTableMethod_Get         SlbCurAdvhcOcspTableMethod = 1
	SlbCurAdvhcOcspTableMethod_Post        SlbCurAdvhcOcspTableMethod = 2
	SlbCurAdvhcOcspTableMethod_Unsupported SlbCurAdvhcOcspTableMethod = 2147483647
)

type SlbCurAdvhcOcspTableHttpsCipherName int32

const (
	SlbCurAdvhcOcspTableHttpsCipherName_UserDefined SlbCurAdvhcOcspTableHttpsCipherName = 1
	SlbCurAdvhcOcspTableHttpsCipherName_Low         SlbCurAdvhcOcspTableHttpsCipherName = 2
	SlbCurAdvhcOcspTableHttpsCipherName_Medium      SlbCurAdvhcOcspTableHttpsCipherName = 3
	SlbCurAdvhcOcspTableHttpsCipherName_High        SlbCurAdvhcOcspTableHttpsCipherName = 4
	SlbCurAdvhcOcspTableHttpsCipherName_Unsupported SlbCurAdvhcOcspTableHttpsCipherName = 2147483647
)

type SlbCurAdvhcOcspTableParams struct {
	// OCSP Health check id.
	ID string `json:"ID,omitempty"`
	// OCSP Health check name.
	Name string `json:"Name,omitempty"`
	// OCSP Health check destination port.
	DPort uint64 `json:"DPort,omitempty"`
	// OCSP Health check destination IP version.
	IPVer SlbCurAdvhcOcspTableIPVer `json:"IPVer,omitempty"`
	// OCSP Health check destination hostname.
	HostName string `json:"HostName,omitempty"`
	// OCSP Health check transparent flag.
	Transparent SlbCurAdvhcOcspTableTransparent `json:"Transparent,omitempty"`
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
	Invert SlbCurAdvhcOcspTableInvert `json:"Invert,omitempty"`
	// OCSP Health check HTTPS enable/disable flag.
	Https SlbCurAdvhcOcspTableHttps `json:"Https,omitempty"`
	// OCSP Health check good certificate check enable/disable flag.
	GoodCert SlbCurAdvhcOcspTableGoodCert `json:"GoodCert,omitempty"`
	// OCSP Health check certificate ID.
	Cert string `json:"Cert,omitempty"`
	// OCSP Health check Trusted CA certificate ID.
	TrustedCA string `json:"TrustedCA,omitempty"`
	// OCSP Health check Connection termination type.
	ConnTerm SlbCurAdvhcOcspTableConnTerm `json:"ConnTerm,omitempty"`
	// OCSP Health check is allowed for standalone real enable/disable flag.
	Always SlbCurAdvhcOcspTableAlways `json:"Always,omitempty"`
	// OCSP Health Check src NAT (PIP) flag.
	Snat SlbCurAdvhcOcspTableSnat `json:"Snat,omitempty"`
	// HTTP Health check host field.
	Host string `json:"Host,omitempty"`
	// OCSP Health check path field.
	Path string `json:"Path,omitempty"`
	// OCSP Health check HTTP method.
	Method SlbCurAdvhcOcspTableMethod `json:"Method,omitempty"`
	// OCSP Health Check Cipher name for SSL for HTTPS Context.
	HttpsCipherName SlbCurAdvhcOcspTableHttpsCipherName `json:"HttpsCipherName,omitempty"`
	// OCSP Health Check Cipher-suite allowed for SSL for HTTPS Context.
	HttpsCipherUserdef string `json:"HttpsCipherUserdef,omitempty"`
}

func (p SlbCurAdvhcOcspTableParams) iMABean() {}
