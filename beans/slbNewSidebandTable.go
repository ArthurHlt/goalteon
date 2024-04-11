package beans

import (
	"fmt"
	"reflect"
)

// SlbNewSidebandTable Sideband new configuration table.
// Note:This mib is not supported for VX instance of Virtualization.
type SlbNewSidebandTable struct {
	// Sideband Index.
	SlbNewSidebandID string
	Params           *SlbNewSidebandTableParams
}

func NewSlbNewSidebandTableList() *SlbNewSidebandTable {
	return &SlbNewSidebandTable{}
}

func NewSlbNewSidebandTable(
	slbNewSidebandID string,
	params *SlbNewSidebandTableParams,
) *SlbNewSidebandTable {
	return &SlbNewSidebandTable{
		SlbNewSidebandID: slbNewSidebandID,
		Params:           params,
	}
}

func (c *SlbNewSidebandTable) Name() string {
	return "SlbNewSidebandTable"
}

func (c *SlbNewSidebandTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbNewSidebandTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbNewSidebandTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbNewSidebandID).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbNewSidebandID)
}

type SlbNewSidebandTableEnaDis int32

const (
	SlbNewSidebandTableEnaDis_Enabled     SlbNewSidebandTableEnaDis = 1
	SlbNewSidebandTableEnaDis_Disabled    SlbNewSidebandTableEnaDis = 2
	SlbNewSidebandTableEnaDis_Unsupported SlbNewSidebandTableEnaDis = 2147483647
)

type SlbNewSidebandTableDel int32

const (
	SlbNewSidebandTableDel_Other       SlbNewSidebandTableDel = 1
	SlbNewSidebandTableDel_Delete      SlbNewSidebandTableDel = 2
	SlbNewSidebandTableDel_Unsupported SlbNewSidebandTableDel = 2147483647
)

type SlbNewSidebandTableApplic int32

const (
	SlbNewSidebandTableApplic_Http        SlbNewSidebandTableApplic = 1
	SlbNewSidebandTableApplic_Dns         SlbNewSidebandTableApplic = 2
	SlbNewSidebandTableApplic_Unsupported SlbNewSidebandTableApplic = 2147483647
)

type SlbNewSidebandTableProxyIpMode int32

const (
	SlbNewSidebandTableProxyIpMode_Egress      SlbNewSidebandTableProxyIpMode = 1
	SlbNewSidebandTableProxyIpMode_Address     SlbNewSidebandTableProxyIpMode = 2
	SlbNewSidebandTableProxyIpMode_Unsupported SlbNewSidebandTableProxyIpMode = 2147483647
)

type SlbNewSidebandTableFallback int32

const (
	SlbNewSidebandTableFallback_FallbackClosed SlbNewSidebandTableFallback = 1
	SlbNewSidebandTableFallback_FallbackOpen   SlbNewSidebandTableFallback = 2
	SlbNewSidebandTableFallback_Unsupported    SlbNewSidebandTableFallback = 2147483647
)

type SlbNewSidebandTableClnsnat int32

const (
	SlbNewSidebandTableClnsnat_Enabled     SlbNewSidebandTableClnsnat = 1
	SlbNewSidebandTableClnsnat_Disabled    SlbNewSidebandTableClnsnat = 2
	SlbNewSidebandTableClnsnat_Unsupported SlbNewSidebandTableClnsnat = 2147483647
)

type SlbNewSidebandTableParams struct {
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
	EnaDis SlbNewSidebandTableEnaDis `json:"EnaDis,omitempty"`
	// Sideband response timeout.
	Timeout uint64 `json:"Timeout,omitempty"`
	// When set to the value of 2 (delete), the entire row is deleted.
	// When read, other(1) is returned. Setting the value to anything
	// other than 2(delete) has no effect on the state of the row.
	Del SlbNewSidebandTableDel `json:"Del,omitempty"`
	// Sideband applic type.
	Applic SlbNewSidebandTableApplic `json:"Applic,omitempty"`
	// Set the Proxy IP mode for sideband policy
	// Changing from address(1) to egress will clear the configured IPv4/IPv6 address & prefix.
	ProxyIpMode SlbNewSidebandTableProxyIpMode `json:"ProxyIpMode,omitempty"`
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
	// Set the fallback action in case of no-response or failure.
	Fallback SlbNewSidebandTableFallback `json:"Fallback,omitempty"`
	// Set client snat.
	Clnsnat SlbNewSidebandTableClnsnat `json:"Clnsnat,omitempty"`
}

func (p SlbNewSidebandTableParams) iMABean() {}
