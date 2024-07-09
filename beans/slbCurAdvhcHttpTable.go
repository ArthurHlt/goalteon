package beans

import (
	"fmt"
	"reflect"
)

// SlbCurAdvhcHttpTable Note:This mib is not supported on VX instance of Virtualization.
type SlbCurAdvhcHttpTable struct {
	// HTTP Health check id.
	SlbCurAdvhcHttpID string
	Params            *SlbCurAdvhcHttpTableParams
}

func NewSlbCurAdvhcHttpTableList() *SlbCurAdvhcHttpTable {
	return &SlbCurAdvhcHttpTable{}
}

func NewSlbCurAdvhcHttpTable(
	slbCurAdvhcHttpID string,
	params *SlbCurAdvhcHttpTableParams,
) *SlbCurAdvhcHttpTable {
	return &SlbCurAdvhcHttpTable{
		SlbCurAdvhcHttpID: slbCurAdvhcHttpID,
		Params:            params,
	}
}

func (c *SlbCurAdvhcHttpTable) Name() string {
	return "SlbCurAdvhcHttpTable"
}

func (c *SlbCurAdvhcHttpTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbCurAdvhcHttpTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbCurAdvhcHttpTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbCurAdvhcHttpID).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbCurAdvhcHttpID)
}

type SlbCurAdvhcHttpTableIPVer int32

const (
	SlbCurAdvhcHttpTableIPVer_Ipv4 SlbCurAdvhcHttpTableIPVer = 1
	SlbCurAdvhcHttpTableIPVer_Ipv6 SlbCurAdvhcHttpTableIPVer = 2
	SlbCurAdvhcHttpTableIPVer_None SlbCurAdvhcHttpTableIPVer = 3
)

type SlbCurAdvhcHttpTableTransparent int32

const (
	SlbCurAdvhcHttpTableTransparent_Enabled     SlbCurAdvhcHttpTableTransparent = 1
	SlbCurAdvhcHttpTableTransparent_Disabled    SlbCurAdvhcHttpTableTransparent = 2
	SlbCurAdvhcHttpTableTransparent_Unsupported SlbCurAdvhcHttpTableTransparent = 2147483647
)

type SlbCurAdvhcHttpTableInvert int32

const (
	SlbCurAdvhcHttpTableInvert_Enabled     SlbCurAdvhcHttpTableInvert = 1
	SlbCurAdvhcHttpTableInvert_Disabled    SlbCurAdvhcHttpTableInvert = 2
	SlbCurAdvhcHttpTableInvert_Unsupported SlbCurAdvhcHttpTableInvert = 2147483647
)

type SlbCurAdvhcHttpTableHttps int32

const (
	SlbCurAdvhcHttpTableHttps_Enabled     SlbCurAdvhcHttpTableHttps = 1
	SlbCurAdvhcHttpTableHttps_Disabled    SlbCurAdvhcHttpTableHttps = 2
	SlbCurAdvhcHttpTableHttps_Unsupported SlbCurAdvhcHttpTableHttps = 2147483647
)

type SlbCurAdvhcHttpTableMethod int32

const (
	SlbCurAdvhcHttpTableMethod_Get         SlbCurAdvhcHttpTableMethod = 1
	SlbCurAdvhcHttpTableMethod_Post        SlbCurAdvhcHttpTableMethod = 2
	SlbCurAdvhcHttpTableMethod_Head        SlbCurAdvhcHttpTableMethod = 3
	SlbCurAdvhcHttpTableMethod_Unsupported SlbCurAdvhcHttpTableMethod = 2147483647
)

type SlbCurAdvhcHttpTableAuthLevel int32

const (
	SlbCurAdvhcHttpTableAuthLevel_None        SlbCurAdvhcHttpTableAuthLevel = 1
	SlbCurAdvhcHttpTableAuthLevel_Basic       SlbCurAdvhcHttpTableAuthLevel = 2
	SlbCurAdvhcHttpTableAuthLevel_Ntlm2       SlbCurAdvhcHttpTableAuthLevel = 3
	SlbCurAdvhcHttpTableAuthLevel_Ntlmssp     SlbCurAdvhcHttpTableAuthLevel = 4
	SlbCurAdvhcHttpTableAuthLevel_Unsupported SlbCurAdvhcHttpTableAuthLevel = 2147483647
)

type SlbCurAdvhcHttpTableResponseType int32

const (
	SlbCurAdvhcHttpTableResponseType_None        SlbCurAdvhcHttpTableResponseType = 1
	SlbCurAdvhcHttpTableResponseType_Incl        SlbCurAdvhcHttpTableResponseType = 2
	SlbCurAdvhcHttpTableResponseType_Excl        SlbCurAdvhcHttpTableResponseType = 4
	SlbCurAdvhcHttpTableResponseType_Unsupported SlbCurAdvhcHttpTableResponseType = 2147483647
)

type SlbCurAdvhcHttpTableOverloadType int32

const (
	SlbCurAdvhcHttpTableOverloadType_None        SlbCurAdvhcHttpTableOverloadType = 1
	SlbCurAdvhcHttpTableOverloadType_Incl        SlbCurAdvhcHttpTableOverloadType = 2
	SlbCurAdvhcHttpTableOverloadType_Unsupported SlbCurAdvhcHttpTableOverloadType = 2147483647
)

type SlbCurAdvhcHttpTableProxy int32

const (
	SlbCurAdvhcHttpTableProxy_Enabled     SlbCurAdvhcHttpTableProxy = 1
	SlbCurAdvhcHttpTableProxy_Disabled    SlbCurAdvhcHttpTableProxy = 2
	SlbCurAdvhcHttpTableProxy_Unsupported SlbCurAdvhcHttpTableProxy = 2147483647
)

type SlbCurAdvhcHttpTableConnTerm int32

const (
	SlbCurAdvhcHttpTableConnTerm_Fin         SlbCurAdvhcHttpTableConnTerm = 1
	SlbCurAdvhcHttpTableConnTerm_Rst         SlbCurAdvhcHttpTableConnTerm = 2
	SlbCurAdvhcHttpTableConnTerm_Unsupported SlbCurAdvhcHttpTableConnTerm = 2147483647
)

type SlbCurAdvhcHttpTableHttpsCipherName int32

const (
	SlbCurAdvhcHttpTableHttpsCipherName_UserDefined SlbCurAdvhcHttpTableHttpsCipherName = 1
	SlbCurAdvhcHttpTableHttpsCipherName_Low         SlbCurAdvhcHttpTableHttpsCipherName = 2
	SlbCurAdvhcHttpTableHttpsCipherName_Medium      SlbCurAdvhcHttpTableHttpsCipherName = 3
	SlbCurAdvhcHttpTableHttpsCipherName_High        SlbCurAdvhcHttpTableHttpsCipherName = 4
	SlbCurAdvhcHttpTableHttpsCipherName_Unsupported SlbCurAdvhcHttpTableHttpsCipherName = 2147483647
)

type SlbCurAdvhcHttpTableHttp2 int32

const (
	SlbCurAdvhcHttpTableHttp2_Enabled     SlbCurAdvhcHttpTableHttp2 = 1
	SlbCurAdvhcHttpTableHttp2_Disabled    SlbCurAdvhcHttpTableHttp2 = 2
	SlbCurAdvhcHttpTableHttp2_Unsupported SlbCurAdvhcHttpTableHttp2 = 2147483647
)

type SlbCurAdvhcHttpTableAlways int32

const (
	SlbCurAdvhcHttpTableAlways_Enabled     SlbCurAdvhcHttpTableAlways = 1
	SlbCurAdvhcHttpTableAlways_Disabled    SlbCurAdvhcHttpTableAlways = 2
	SlbCurAdvhcHttpTableAlways_Unsupported SlbCurAdvhcHttpTableAlways = 2147483647
)

type SlbCurAdvhcHttpTableSnat int32

const (
	SlbCurAdvhcHttpTableSnat_Enabled     SlbCurAdvhcHttpTableSnat = 1
	SlbCurAdvhcHttpTableSnat_Disabled    SlbCurAdvhcHttpTableSnat = 2
	SlbCurAdvhcHttpTableSnat_Unsupported SlbCurAdvhcHttpTableSnat = 2147483647
)

type SlbCurAdvhcHttpTableConnTout int32

const (
	SlbCurAdvhcHttpTableConnTout_Fin         SlbCurAdvhcHttpTableConnTout = 1
	SlbCurAdvhcHttpTableConnTout_Rst         SlbCurAdvhcHttpTableConnTout = 2
	SlbCurAdvhcHttpTableConnTout_Unsupported SlbCurAdvhcHttpTableConnTout = 2147483647
)

type SlbCurAdvhcHttpTableParams struct {
	// HTTP Health check id.
	ID string `json:"ID,omitempty"`
	// HTTP Health check name.
	Name string `json:"Name,omitempty"`
	// HTTP Health check destination port.
	DPort uint64 `json:"DPort,omitempty"`
	// HTTP Health check destination IP version.
	IPVer SlbCurAdvhcHttpTableIPVer `json:"IPVer,omitempty"`
	// HTTP Health check destination hostname.
	HostName string `json:"HostName,omitempty"`
	// HTTP Health check transparent flag.
	Transparent SlbCurAdvhcHttpTableTransparent `json:"Transparent,omitempty"`
	// HTTP Health check interval.
	Interval uint64 `json:"Interval,omitempty"`
	// HTTP Health check retries counter.
	Retries uint64 `json:"Retries,omitempty"`
	// HTTP Health check retries in down state counter.
	RestoreRetries uint64 `json:"RestoreRetries,omitempty"`
	// HTTP Health check timeout.
	Timeout uint64 `json:"Timeout,omitempty"`
	// HTTP Health check overflow flag.
	Overflow uint64 `json:"Overflow,omitempty"`
	// HTTP Health check interval in down state.
	DownInterval uint64 `json:"DownInterval,omitempty"`
	// HTTP Health check invert flag.
	Invert SlbCurAdvhcHttpTableInvert `json:"Invert,omitempty"`
	// HTTP Health check HTTPS enable/disable flag.
	Https SlbCurAdvhcHttpTableHttps `json:"Https,omitempty"`
	// HTTP Health check host field.
	Host string `json:"Host,omitempty"`
	// HTTP Health check path field.
	Path string `json:"Path,omitempty"`
	// HTTP Health check HTTP method.
	Method SlbCurAdvhcHttpTableMethod `json:"Method,omitempty"`
	// HTTP Health check headers list.
	Headers string `json:"Headers,omitempty"`
	// HTTP Health check body field.
	Body string `json:"Body,omitempty"`
	// HTTP Health check authentication method.
	AuthLevel SlbCurAdvhcHttpTableAuthLevel `json:"AuthLevel,omitempty"`
	// HTTP Health check user name.
	UserName string `json:"UserName,omitempty"`
	// HTTP Health check password.
	Password string `json:"Password,omitempty"`
	// HTTP Health check response handling method.
	ResponseType SlbCurAdvhcHttpTableResponseType `json:"ResponseType,omitempty"`
	// overload string is included or not included.
	OverloadType SlbCurAdvhcHttpTableOverloadType `json:"OverloadType,omitempty"`
	// HTTP Health check expected response code.
	ResponseCode string `json:"ResponseCode,omitempty"`
	// HTTP Health check expected response string.
	ReceiveString string `json:"ReceiveString,omitempty"`
	// HTTP Health check expected code for overflow state.
	ResponseCodeOverload string `json:"ResponseCodeOverload,omitempty"`
	// expected response for server overload.
	OverloadString string `json:"OverloadString,omitempty"`
	// Enable/disable HTTP health check proxy request.
	Proxy SlbCurAdvhcHttpTableProxy `json:"Proxy,omitempty"`
	// Connection termination type.
	ConnTerm SlbCurAdvhcHttpTableConnTerm `json:"ConnTerm,omitempty"`
	// Cipher name for SSL for HTTPS HC Context.
	HttpsCipherName SlbCurAdvhcHttpTableHttpsCipherName `json:"HttpsCipherName,omitempty"`
	// Cipher-suite allowed for SSL for HTTPS HC Context.
	HttpsCipherUserdef string `json:"HttpsCipherUserdef,omitempty"`
	// HTTP Health check enable/disable HTTP2.
	Http2 SlbCurAdvhcHttpTableHttp2 `json:"Http2,omitempty"`
	// This flag determines whether HC is allowed for standalone real.
	Always SlbCurAdvhcHttpTableAlways `json:"Always,omitempty"`
	// HTTP Health Check src NAT (PIP) flag.
	Snat SlbCurAdvhcHttpTableSnat `json:"Snat,omitempty"`
	// Connection termination on timeout type.
	ConnTout SlbCurAdvhcHttpTableConnTout `json:"ConnTout,omitempty"`
}

func (p SlbCurAdvhcHttpTableParams) iMABean() {}
