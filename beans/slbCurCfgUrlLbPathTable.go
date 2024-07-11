package beans

import (
	"fmt"
	"reflect"
)

// SlbCurCfgUrlLbPathTable The table of URL path for URL load balancing in the current_config.
// Note:This mib is not supported for VX instance of virtualization.
type SlbCurCfgUrlLbPathTable struct {
	// The URL path table index.
	SlbCurCfgUrlLbPathIndex int32
	Params                  *SlbCurCfgUrlLbPathTableParams
}

func NewSlbCurCfgUrlLbPathTableList() *SlbCurCfgUrlLbPathTable {
	return &SlbCurCfgUrlLbPathTable{}
}

func NewSlbCurCfgUrlLbPathTable(
	slbCurCfgUrlLbPathIndex int32,
	params *SlbCurCfgUrlLbPathTableParams,
) *SlbCurCfgUrlLbPathTable {
	return &SlbCurCfgUrlLbPathTable{
		SlbCurCfgUrlLbPathIndex: slbCurCfgUrlLbPathIndex,
		Params:                  params,
	}
}

func (c *SlbCurCfgUrlLbPathTable) Name() string {
	return "SlbCurCfgUrlLbPathTable"
}

func (c *SlbCurCfgUrlLbPathTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbCurCfgUrlLbPathTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbCurCfgUrlLbPathTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbCurCfgUrlLbPathIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbCurCfgUrlLbPathIndex)
}

type SlbCurCfgUrlLbPathTablePatternStringType int32

const (
	SlbCurCfgUrlLbPathTablePatternStringType_Ascii       SlbCurCfgUrlLbPathTablePatternStringType = 1
	SlbCurCfgUrlLbPathTablePatternStringType_Binary      SlbCurCfgUrlLbPathTablePatternStringType = 2
	SlbCurCfgUrlLbPathTablePatternStringType_None        SlbCurCfgUrlLbPathTablePatternStringType = 3
	SlbCurCfgUrlLbPathTablePatternStringType_Unsupported SlbCurCfgUrlLbPathTablePatternStringType = 2147483647
)

type SlbCurCfgUrlLbPathTableOper int32

const (
	SlbCurCfgUrlLbPathTableOper_Eq          SlbCurCfgUrlLbPathTableOper = 1
	SlbCurCfgUrlLbPathTableOper_Gt          SlbCurCfgUrlLbPathTableOper = 2
	SlbCurCfgUrlLbPathTableOper_Lt          SlbCurCfgUrlLbPathTableOper = 3
	SlbCurCfgUrlLbPathTableOper_None        SlbCurCfgUrlLbPathTableOper = 4
	SlbCurCfgUrlLbPathTableOper_Unsupported SlbCurCfgUrlLbPathTableOper = 2147483647
)

type SlbCurCfgUrlLbPathTableAllowRegExp int32

const (
	SlbCurCfgUrlLbPathTableAllowRegExp_Yes         SlbCurCfgUrlLbPathTableAllowRegExp = 1
	SlbCurCfgUrlLbPathTableAllowRegExp_No          SlbCurCfgUrlLbPathTableAllowRegExp = 2
	SlbCurCfgUrlLbPathTableAllowRegExp_Unsupported SlbCurCfgUrlLbPathTableAllowRegExp = 2147483647
)

type SlbCurCfgUrlLbPathTableDnsType int32

const (
	SlbCurCfgUrlLbPathTableDnsType_Dns         SlbCurCfgUrlLbPathTableDnsType = 1
	SlbCurCfgUrlLbPathTableDnsType_Dnssec      SlbCurCfgUrlLbPathTableDnsType = 2
	SlbCurCfgUrlLbPathTableDnsType_Any         SlbCurCfgUrlLbPathTableDnsType = 3
	SlbCurCfgUrlLbPathTableDnsType_Unsupported SlbCurCfgUrlLbPathTableDnsType = 2147483647
)

type SlbCurCfgUrlLbPathTableApplication int32

const (
	SlbCurCfgUrlLbPathTableApplication_Http        SlbCurCfgUrlLbPathTableApplication = 1
	SlbCurCfgUrlLbPathTableApplication_Dns         SlbCurCfgUrlLbPathTableApplication = 2
	SlbCurCfgUrlLbPathTableApplication_Other       SlbCurCfgUrlLbPathTableApplication = 3
	SlbCurCfgUrlLbPathTableApplication_Unsupported SlbCurCfgUrlLbPathTableApplication = 2147483647
)

type SlbCurCfgUrlLbPathTableParams struct {
	// The URL path table index.
	Index int32 `json:"Index,omitempty"`
	// The SLB string or ASCII/BINARY string for pattern matching .
	StringVal string `json:"StringVal,omitempty"`
	// The BW contract for the path.
	BwmContract int32 `json:"BwmContract,omitempty"`
	// The HTTP header for URL load balancing.
	HTTPHeader string `json:"HTTPHeader,omitempty"`
	// The HTTP header value for URL load balancing.
	HTTPHeaderValue string `json:"HTTPHeaderValue,omitempty"`
	// Type of pattern string (ASCII or binary or none).
	PatternStringType SlbCurCfgUrlLbPathTablePatternStringType `json:"PatternStringType,omitempty"`
	// Offset from beginning of IP packet to start matching the
	// 	 pattern string.
	Offset uint32 `json:"Offset,omitempty"`
	// Depth of IP packet to search and match the pattern string.
	Depth uint32 `json:"Depth,omitempty"`
	// Operation to be performed on the pattern match string.
	// 	 For ASCII pattern strings, only the equal (eq) operation
	// 	 is valid.
	Oper SlbCurCfgUrlLbPathTableOper `json:"Oper,omitempty"`
	// DNS Query type(s) (by number).
	DnsQueryTypes string `json:"DnsQueryTypes,omitempty"`
	// The complete SLB or ASCII/BINARY string for pattern matching .
	CompleteString string `json:"CompleteString,omitempty"`
	// Allow regular expression for layer7 SLB path string.
	AllowRegExp SlbCurCfgUrlLbPathTableAllowRegExp `json:"AllowRegExp,omitempty"`
	// DNS type (dns, dnssec, any).
	DnsType SlbCurCfgUrlLbPathTableDnsType `json:"DnsType,omitempty"`
	// Type of appliction (http, dns, other).
	Application SlbCurCfgUrlLbPathTableApplication `json:"Application,omitempty"`
}

func (p SlbCurCfgUrlLbPathTableParams) iMABean() {}
