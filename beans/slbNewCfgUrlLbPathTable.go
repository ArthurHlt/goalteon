package beans

import (
	"fmt"
	"reflect"
)

// SlbNewCfgUrlLbPathTable The table of URL path for URL load balancing in the new_config.
// Note:This mib is not supported for VX instance of virtualization.
type SlbNewCfgUrlLbPathTable struct {
	// The URL path table index.
	SlbNewCfgUrlLbPathIndex int32
	Params                  *SlbNewCfgUrlLbPathTableParams
}

func NewSlbNewCfgUrlLbPathTableList() *SlbNewCfgUrlLbPathTable {
	return &SlbNewCfgUrlLbPathTable{}
}

func NewSlbNewCfgUrlLbPathTable(
	slbNewCfgUrlLbPathIndex int32,
	params *SlbNewCfgUrlLbPathTableParams,
) *SlbNewCfgUrlLbPathTable {
	return &SlbNewCfgUrlLbPathTable{
		SlbNewCfgUrlLbPathIndex: slbNewCfgUrlLbPathIndex,
		Params:                  params,
	}
}

func (c *SlbNewCfgUrlLbPathTable) Name() string {
	return "SlbNewCfgUrlLbPathTable"
}

func (c *SlbNewCfgUrlLbPathTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbNewCfgUrlLbPathTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbNewCfgUrlLbPathTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbNewCfgUrlLbPathIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbNewCfgUrlLbPathIndex)
}

type SlbNewCfgUrlLbPathTableDelete int32

const (
	SlbNewCfgUrlLbPathTableDelete_Other       SlbNewCfgUrlLbPathTableDelete = 1
	SlbNewCfgUrlLbPathTableDelete_Delete      SlbNewCfgUrlLbPathTableDelete = 2
	SlbNewCfgUrlLbPathTableDelete_Unsupported SlbNewCfgUrlLbPathTableDelete = 2147483647
)

type SlbNewCfgUrlLbPathTablePatternStringType int32

const (
	SlbNewCfgUrlLbPathTablePatternStringType_Ascii       SlbNewCfgUrlLbPathTablePatternStringType = 1
	SlbNewCfgUrlLbPathTablePatternStringType_Binary      SlbNewCfgUrlLbPathTablePatternStringType = 2
	SlbNewCfgUrlLbPathTablePatternStringType_None        SlbNewCfgUrlLbPathTablePatternStringType = 3
	SlbNewCfgUrlLbPathTablePatternStringType_Unsupported SlbNewCfgUrlLbPathTablePatternStringType = 2147483647
)

type SlbNewCfgUrlLbPathTableOper int32

const (
	SlbNewCfgUrlLbPathTableOper_Eq          SlbNewCfgUrlLbPathTableOper = 1
	SlbNewCfgUrlLbPathTableOper_Gt          SlbNewCfgUrlLbPathTableOper = 2
	SlbNewCfgUrlLbPathTableOper_Lt          SlbNewCfgUrlLbPathTableOper = 3
	SlbNewCfgUrlLbPathTableOper_None        SlbNewCfgUrlLbPathTableOper = 4
	SlbNewCfgUrlLbPathTableOper_Unsupported SlbNewCfgUrlLbPathTableOper = 2147483647
)

type SlbNewCfgUrlLbPathTableAllowRegExp int32

const (
	SlbNewCfgUrlLbPathTableAllowRegExp_Yes         SlbNewCfgUrlLbPathTableAllowRegExp = 1
	SlbNewCfgUrlLbPathTableAllowRegExp_No          SlbNewCfgUrlLbPathTableAllowRegExp = 2
	SlbNewCfgUrlLbPathTableAllowRegExp_Unsupported SlbNewCfgUrlLbPathTableAllowRegExp = 2147483647
)

type SlbNewCfgUrlLbPathTableDnsType int32

const (
	SlbNewCfgUrlLbPathTableDnsType_Dns         SlbNewCfgUrlLbPathTableDnsType = 1
	SlbNewCfgUrlLbPathTableDnsType_Dnssec      SlbNewCfgUrlLbPathTableDnsType = 2
	SlbNewCfgUrlLbPathTableDnsType_Any         SlbNewCfgUrlLbPathTableDnsType = 3
	SlbNewCfgUrlLbPathTableDnsType_Unsupported SlbNewCfgUrlLbPathTableDnsType = 2147483647
)

type SlbNewCfgUrlLbPathTableApplication int32

const (
	SlbNewCfgUrlLbPathTableApplication_Http        SlbNewCfgUrlLbPathTableApplication = 1
	SlbNewCfgUrlLbPathTableApplication_Dns         SlbNewCfgUrlLbPathTableApplication = 2
	SlbNewCfgUrlLbPathTableApplication_Other       SlbNewCfgUrlLbPathTableApplication = 3
	SlbNewCfgUrlLbPathTableApplication_Unsupported SlbNewCfgUrlLbPathTableApplication = 2147483647
)

type SlbNewCfgUrlLbPathTableParams struct {
	// The URL path table index.
	Index int32 `json:"Index,omitempty"`
	// The SLB string or ASCII/BINARY string for pattern matching .
	StringVal string `json:"StringVal,omitempty"`
	// Action object to delete an URL Path.  When set to the value
	// of 2 (delete), the entire row is deleted. When read, other(1) is
	// returned. Setting the value to anything other than delete(2)
	// has no effect on the state of the row.
	Delete SlbNewCfgUrlLbPathTableDelete `json:"Delete,omitempty"`
	// The BW contract for the load path.
	BwmContract int32 `json:"BwmContract,omitempty"`
	// The HTTP header for URL load balancing.
	HTTPHeader string `json:"HTTPHeader,omitempty"`
	// The HTTP header value for URL load balancing.
	HTTPHeaderValue string `json:"HTTPHeaderValue,omitempty"`
	// Type of pattern string (ASCII or binary).
	PatternStringType SlbNewCfgUrlLbPathTablePatternStringType `json:"PatternStringType,omitempty"`
	// Offset from beginning of IP packet to start matching the
	// 	 pattern string.
	Offset uint32 `json:"Offset,omitempty"`
	// Depth of IP packet to search and match the pattern string.
	Depth uint32 `json:"Depth,omitempty"`
	// Operation to be performed on the pattern match string.
	// 	 For ASCII pattern strings, only the equal (eq) operation
	// 	 is valid.
	Oper SlbNewCfgUrlLbPathTableOper `json:"Oper,omitempty"`
	// DNS Query type(s) (by number).
	DnsQueryTypes string `json:"DnsQueryTypes,omitempty"`
	// The complete SLB or ASCII/BINARY string for pattern matching .
	CompleteString string `json:"CompleteString,omitempty"`
	// Allow regular expression for layer7 SLB path string.
	AllowRegExp SlbNewCfgUrlLbPathTableAllowRegExp `json:"AllowRegExp,omitempty"`
	// DNS type (dns, dnssec, any).
	DnsType SlbNewCfgUrlLbPathTableDnsType `json:"DnsType,omitempty"`
	// Type of application (http, dns, other).
	Application SlbNewCfgUrlLbPathTableApplication `json:"Application,omitempty"`
}

func (p SlbNewCfgUrlLbPathTableParams) iMABean() {}
