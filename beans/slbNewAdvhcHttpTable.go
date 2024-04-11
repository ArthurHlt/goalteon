package beans

import (
	"fmt"
	"reflect"
)

// SlbNewAdvhcHttpTable Note:This mib is not supported on VX instance of Virtualization.
type SlbNewAdvhcHttpTable struct {
	// HTTP Health check id.
	SlbNewAdvhcHttpID string
	Params            *SlbNewAdvhcHttpTableParams
}

func NewSlbNewAdvhcHttpTableList() *SlbNewAdvhcHttpTable {
	return &SlbNewAdvhcHttpTable{}
}

func NewSlbNewAdvhcHttpTable(
	slbNewAdvhcHttpID string,
	params *SlbNewAdvhcHttpTableParams,
) *SlbNewAdvhcHttpTable {
	return &SlbNewAdvhcHttpTable{
		SlbNewAdvhcHttpID: slbNewAdvhcHttpID,
		Params:            params,
	}
}

func (c *SlbNewAdvhcHttpTable) Name() string {
	return "SlbNewAdvhcHttpTable"
}

func (c *SlbNewAdvhcHttpTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbNewAdvhcHttpTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbNewAdvhcHttpTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbNewAdvhcHttpID).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbNewAdvhcHttpID)
}

type SlbNewAdvhcHttpTableIPVer int32

const (
	SlbNewAdvhcHttpTableIPVer_Ipv4 SlbNewAdvhcHttpTableIPVer = 1
	SlbNewAdvhcHttpTableIPVer_Ipv6 SlbNewAdvhcHttpTableIPVer = 2
	SlbNewAdvhcHttpTableIPVer_None SlbNewAdvhcHttpTableIPVer = 3
)

type SlbNewAdvhcHttpTableTransparent int32

const (
	SlbNewAdvhcHttpTableTransparent_Enabled     SlbNewAdvhcHttpTableTransparent = 1
	SlbNewAdvhcHttpTableTransparent_Disabled    SlbNewAdvhcHttpTableTransparent = 2
	SlbNewAdvhcHttpTableTransparent_Unsupported SlbNewAdvhcHttpTableTransparent = 2147483647
)

type SlbNewAdvhcHttpTableInvert int32

const (
	SlbNewAdvhcHttpTableInvert_Enabled     SlbNewAdvhcHttpTableInvert = 1
	SlbNewAdvhcHttpTableInvert_Disabled    SlbNewAdvhcHttpTableInvert = 2
	SlbNewAdvhcHttpTableInvert_Unsupported SlbNewAdvhcHttpTableInvert = 2147483647
)

type SlbNewAdvhcHttpTableHttps int32

const (
	SlbNewAdvhcHttpTableHttps_Enabled     SlbNewAdvhcHttpTableHttps = 1
	SlbNewAdvhcHttpTableHttps_Disabled    SlbNewAdvhcHttpTableHttps = 2
	SlbNewAdvhcHttpTableHttps_Unsupported SlbNewAdvhcHttpTableHttps = 2147483647
)

type SlbNewAdvhcHttpTableMethod int32

const (
	SlbNewAdvhcHttpTableMethod_Get         SlbNewAdvhcHttpTableMethod = 1
	SlbNewAdvhcHttpTableMethod_Post        SlbNewAdvhcHttpTableMethod = 2
	SlbNewAdvhcHttpTableMethod_Head        SlbNewAdvhcHttpTableMethod = 3
	SlbNewAdvhcHttpTableMethod_Unsupported SlbNewAdvhcHttpTableMethod = 2147483647
)

type SlbNewAdvhcHttpTableAuthLevel int32

const (
	SlbNewAdvhcHttpTableAuthLevel_None        SlbNewAdvhcHttpTableAuthLevel = 1
	SlbNewAdvhcHttpTableAuthLevel_Basic       SlbNewAdvhcHttpTableAuthLevel = 2
	SlbNewAdvhcHttpTableAuthLevel_Ntlm2       SlbNewAdvhcHttpTableAuthLevel = 3
	SlbNewAdvhcHttpTableAuthLevel_Ntlmssp     SlbNewAdvhcHttpTableAuthLevel = 4
	SlbNewAdvhcHttpTableAuthLevel_Unsupported SlbNewAdvhcHttpTableAuthLevel = 2147483647
)

type SlbNewAdvhcHttpTableResponseType int32

const (
	SlbNewAdvhcHttpTableResponseType_None        SlbNewAdvhcHttpTableResponseType = 1
	SlbNewAdvhcHttpTableResponseType_Incl        SlbNewAdvhcHttpTableResponseType = 2
	SlbNewAdvhcHttpTableResponseType_Excl        SlbNewAdvhcHttpTableResponseType = 4
	SlbNewAdvhcHttpTableResponseType_Unsupported SlbNewAdvhcHttpTableResponseType = 2147483647
)

type SlbNewAdvhcHttpTableOverloadType int32

const (
	SlbNewAdvhcHttpTableOverloadType_None        SlbNewAdvhcHttpTableOverloadType = 1
	SlbNewAdvhcHttpTableOverloadType_Incl        SlbNewAdvhcHttpTableOverloadType = 2
	SlbNewAdvhcHttpTableOverloadType_Unsupported SlbNewAdvhcHttpTableOverloadType = 2147483647
)

type SlbNewAdvhcHttpTableDelete int32

const (
	SlbNewAdvhcHttpTableDelete_Other       SlbNewAdvhcHttpTableDelete = 1
	SlbNewAdvhcHttpTableDelete_Delete      SlbNewAdvhcHttpTableDelete = 2
	SlbNewAdvhcHttpTableDelete_Unsupported SlbNewAdvhcHttpTableDelete = 2147483647
)

type SlbNewAdvhcHttpTableProxy int32

const (
	SlbNewAdvhcHttpTableProxy_Enabled     SlbNewAdvhcHttpTableProxy = 1
	SlbNewAdvhcHttpTableProxy_Disabled    SlbNewAdvhcHttpTableProxy = 2
	SlbNewAdvhcHttpTableProxy_Unsupported SlbNewAdvhcHttpTableProxy = 2147483647
)

type SlbNewAdvhcHttpTableConnTerm int32

const (
	SlbNewAdvhcHttpTableConnTerm_Fin         SlbNewAdvhcHttpTableConnTerm = 1
	SlbNewAdvhcHttpTableConnTerm_Rst         SlbNewAdvhcHttpTableConnTerm = 2
	SlbNewAdvhcHttpTableConnTerm_Unsupported SlbNewAdvhcHttpTableConnTerm = 2147483647
)

type SlbNewAdvhcHttpTableHttpsCipherName int32

const (
	SlbNewAdvhcHttpTableHttpsCipherName_UserDefined SlbNewAdvhcHttpTableHttpsCipherName = 1
	SlbNewAdvhcHttpTableHttpsCipherName_Low         SlbNewAdvhcHttpTableHttpsCipherName = 2
	SlbNewAdvhcHttpTableHttpsCipherName_Medium      SlbNewAdvhcHttpTableHttpsCipherName = 3
	SlbNewAdvhcHttpTableHttpsCipherName_High        SlbNewAdvhcHttpTableHttpsCipherName = 4
	SlbNewAdvhcHttpTableHttpsCipherName_Unsupported SlbNewAdvhcHttpTableHttpsCipherName = 2147483647
)

type SlbNewAdvhcHttpTableHttp2 int32

const (
	SlbNewAdvhcHttpTableHttp2_Enabled     SlbNewAdvhcHttpTableHttp2 = 1
	SlbNewAdvhcHttpTableHttp2_Disabled    SlbNewAdvhcHttpTableHttp2 = 2
	SlbNewAdvhcHttpTableHttp2_Unsupported SlbNewAdvhcHttpTableHttp2 = 2147483647
)

type SlbNewAdvhcHttpTableAlways int32

const (
	SlbNewAdvhcHttpTableAlways_Enabled     SlbNewAdvhcHttpTableAlways = 1
	SlbNewAdvhcHttpTableAlways_Disabled    SlbNewAdvhcHttpTableAlways = 2
	SlbNewAdvhcHttpTableAlways_Unsupported SlbNewAdvhcHttpTableAlways = 2147483647
)

type SlbNewAdvhcHttpTableConnTout int32

const (
	SlbNewAdvhcHttpTableConnTout_Fin         SlbNewAdvhcHttpTableConnTout = 1
	SlbNewAdvhcHttpTableConnTout_Rst         SlbNewAdvhcHttpTableConnTout = 2
	SlbNewAdvhcHttpTableConnTout_Unsupported SlbNewAdvhcHttpTableConnTout = 2147483647
)

type SlbNewAdvhcHttpTableParams struct {
	// HTTP Health check id.
	ID string `json:"ID,omitempty"`
	// HTTP Health check name.
	Name string `json:"Name,omitempty"`
	// HTTP Health check destination port.
	DPort uint64 `json:"DPort,omitempty"`
	// HTTP Health check destination IP version.
	IPVer SlbNewAdvhcHttpTableIPVer `json:"IPVer,omitempty"`
	// HTTP Health check destination hostname.
	HostName string `json:"HostName,omitempty"`
	// HTTP Health check transparent flag.
	Transparent SlbNewAdvhcHttpTableTransparent `json:"Transparent,omitempty"`
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
	Invert SlbNewAdvhcHttpTableInvert `json:"Invert,omitempty"`
	// HTTP Health check HTTPS enable/disable flag.
	Https SlbNewAdvhcHttpTableHttps `json:"Https,omitempty"`
	// HTTP Health check host field.
	Host string `json:"Host,omitempty"`
	// HTTP Health check path field.
	Path string `json:"Path,omitempty"`
	// HTTP Health check HTTP method.
	Method SlbNewAdvhcHttpTableMethod `json:"Method,omitempty"`
	// HTTP Health check headers list.
	Headers string `json:"Headers,omitempty"`
	// HTTP Health check body field.
	Body string `json:"Body,omitempty"`
	// HTTP Health check authentication method.
	AuthLevel SlbNewAdvhcHttpTableAuthLevel `json:"AuthLevel,omitempty"`
	// HTTP Health check user name.
	UserName string `json:"UserName,omitempty"`
	// HTTP Health check password.
	Password string `json:"Password,omitempty"`
	// HTTP Health check response handling method.
	ResponseType SlbNewAdvhcHttpTableResponseType `json:"ResponseType,omitempty"`
	// overload string is included or not included.
	OverloadType SlbNewAdvhcHttpTableOverloadType `json:"OverloadType,omitempty"`
	// HTTP Health check expected response code.
	ResponseCode string `json:"ResponseCode,omitempty"`
	// HTTP Health check expected response string.
	ReceiveString string `json:"ReceiveString,omitempty"`
	// HTTP Health check expected code for overflow state.
	ResponseCodeOverload string `json:"ResponseCodeOverload,omitempty"`
	// expected response for server overload.
	OverloadString string `json:"OverloadString,omitempty"`
	// HTTP Health check copy action trigger.
	Copy string `json:"Copy,omitempty"`
	// When set to the value of 2 (delete), the entire row is deleted.
	// When read, other(1) is returned. Setting the value to anything
	// other than 2(delete) has no effect on the state of the row.
	Delete SlbNewAdvhcHttpTableDelete `json:"Delete,omitempty"`
	// Enable/disable HTTP health check proxy request.
	Proxy SlbNewAdvhcHttpTableProxy `json:"Proxy,omitempty"`
	// Connection termination type.
	ConnTerm SlbNewAdvhcHttpTableConnTerm `json:"ConnTerm,omitempty"`
	// Cipher name for SSL for HTTPS HC Context.
	HttpsCipherName SlbNewAdvhcHttpTableHttpsCipherName `json:"HttpsCipherName,omitempty"`
	// Cipher-suite allowed for SSL for HTTPS HC Context.
	HttpsCipherUserdef string `json:"HttpsCipherUserdef,omitempty"`
	// HTTP Health check HTTP2 flag.
	Http2 SlbNewAdvhcHttpTableHttp2 `json:"Http2,omitempty"`
	// This flag determines whether HC is allowed for standalone real.
	Always SlbNewAdvhcHttpTableAlways `json:"Always,omitempty"`
	// Connection termination on timeout type.
	ConnTout SlbNewAdvhcHttpTableConnTout `json:"ConnTout,omitempty"`
}

func (p SlbNewAdvhcHttpTableParams) iMABean() {}
