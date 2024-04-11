package beans

import (
	"fmt"
	"reflect"
)

// SlbCurSidebandTable Sideband current configuration table.
// Note:This mib is not supported for VX instance of Virtualization.
type SlbCurSidebandTable struct {
	// Sideband Index.
	SlbCurSidebandID string
	Params           *SlbCurSidebandTableParams
}

func NewSlbCurSidebandTableList() *SlbCurSidebandTable {
	return &SlbCurSidebandTable{}
}

func NewSlbCurSidebandTable(
	slbCurSidebandID string,
	params *SlbCurSidebandTableParams,
) *SlbCurSidebandTable {
	return &SlbCurSidebandTable{
		SlbCurSidebandID: slbCurSidebandID,
		Params:           params,
	}
}

func (c *SlbCurSidebandTable) Name() string {
	return "SlbCurSidebandTable"
}

func (c *SlbCurSidebandTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbCurSidebandTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbCurSidebandTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbCurSidebandID).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbCurSidebandID)
}

type SlbCurSidebandTableEnaDis int32

const (
	SlbCurSidebandTableEnaDis_Enabled     SlbCurSidebandTableEnaDis = 1
	SlbCurSidebandTableEnaDis_Disabled    SlbCurSidebandTableEnaDis = 2
	SlbCurSidebandTableEnaDis_Unsupported SlbCurSidebandTableEnaDis = 2147483647
)

type SlbCurSidebandTableApplic int32

const (
	SlbCurSidebandTableApplic_Http        SlbCurSidebandTableApplic = 1
	SlbCurSidebandTableApplic_Dns         SlbCurSidebandTableApplic = 2
	SlbCurSidebandTableApplic_Unsupported SlbCurSidebandTableApplic = 2147483647
)

type SlbCurSidebandTableProxyIpMode int32

const (
	SlbCurSidebandTableProxyIpMode_Egress      SlbCurSidebandTableProxyIpMode = 0
	SlbCurSidebandTableProxyIpMode_Address     SlbCurSidebandTableProxyIpMode = 1
	SlbCurSidebandTableProxyIpMode_Unsupported SlbCurSidebandTableProxyIpMode = 2147483647
)

type SlbCurSidebandTableFallback int32

const (
	SlbCurSidebandTableFallback_FallbackClosed SlbCurSidebandTableFallback = 1
	SlbCurSidebandTableFallback_FallbackOpen   SlbCurSidebandTableFallback = 2
	SlbCurSidebandTableFallback_Unsupported    SlbCurSidebandTableFallback = 2147483647
)

type SlbCurSidebandTableClnsnat int32

const (
	SlbCurSidebandTableClnsnat_Enabled     SlbCurSidebandTableClnsnat = 1
	SlbCurSidebandTableClnsnat_Disabled    SlbCurSidebandTableClnsnat = 2
	SlbCurSidebandTableClnsnat_Unsupported SlbCurSidebandTableClnsnat = 2147483647
)

type SlbCurSidebandTableParams struct {
	// Sideband Index.
	ID string `json:"ID,omitempty"`
	// Sideband descriptive name.
	Name string `json:"Name,omitempty"`
	// Sideband server port.
	Port uint64 `json:"Port,omitempty"`
	// Sideband group.
	Group string `json:"Group,omitempty"`
	// Sideband SSL policy.
	SslPol string `json:"SslPol,omitempty"`
	// Enable/Disable sideband server.
	EnaDis SlbCurSidebandTableEnaDis `json:"EnaDis,omitempty"`
	// Sideband response timeout.
	Timeout uint64 `json:"Timeout,omitempty"`
	// Sideband applic type.
	Applic SlbCurSidebandTableApplic `json:"Applic,omitempty"`
	// Set the Proxy IP mode for sideband policy
	// Changing from address(1) to egress will clear the configured IPv4/IPv6 address & prefix.
	ProxyIpMode SlbCurSidebandTableProxyIpMode `json:"ProxyIpMode,omitempty"`
	// This object ID allows configuring IPv4 PIP address.
	// This object ID can be set only if slbNewSidebandProxyIpMode is address else return failure.
	// Returns 0 when slbNewSidebandProxyIpMode is not set to address.
	ProxyIpAddr string `json:"ProxyIpAddr,omitempty"`
	// This object ID allows configuring IPv4 PIP Mask.
	// Returns 0 when slbNewSidebandProxyIpMode is not set to address.
	ProxyIpMask string `json:"ProxyIpMask,omitempty"`
	// This object ID allows configuring IPv6 PIP address.
	// Returns 0 when slbNewSidebandProxyIpMode is not set to address.
	// Address should be 4-byte hexadecimal colon notation.
	// Valid IPv6 address should be in any of the following forms
	// xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx or
	// xxxx::xxxx:xxxx:xxxx:xxxx or ::xxxx.
	ProxyIpv6Addr string `json:"ProxyIpv6Addr,omitempty"`
	// This object ID allows configuring IPv6 PIP Mask.
	// Returns 0 when slbNewSidebandProxyIpMode is not set to address.
	ProxyIpv6Prefix uint32 `json:"ProxyIpv6Prefix,omitempty"`
	// Fallback action in case of no-response or failure.
	Fallback SlbCurSidebandTableFallback `json:"Fallback,omitempty"`
	// Get client snat.
	Clnsnat SlbCurSidebandTableClnsnat `json:"Clnsnat,omitempty"`
}

func (p SlbCurSidebandTableParams) iMABean() {}
