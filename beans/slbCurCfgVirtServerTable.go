package beans

import (
	"fmt"
	"reflect"
)

// SlbCurCfgVirtServerTable The table of virtual Servers.
// Note:This mib is not supported for VX instance of Virtualization.
type SlbCurCfgVirtServerTable struct {
	// Virtual Server Number
	SlbCurCfgVirtServerIndex int32
	Params                   *SlbCurCfgVirtServerTableParams
}

func NewSlbCurCfgVirtServerTableList() *SlbCurCfgVirtServerTable {
	return &SlbCurCfgVirtServerTable{}
}

func NewSlbCurCfgVirtServerTable(
	slbCurCfgVirtServerIndex int32,
	params *SlbCurCfgVirtServerTableParams,
) *SlbCurCfgVirtServerTable {
	return &SlbCurCfgVirtServerTable{
		SlbCurCfgVirtServerIndex: slbCurCfgVirtServerIndex,
		Params:                   params,
	}
}

func (c *SlbCurCfgVirtServerTable) Name() string {
	return "SlbCurCfgVirtServerTable"
}

func (c *SlbCurCfgVirtServerTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbCurCfgVirtServerTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbCurCfgVirtServerTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbCurCfgVirtServerIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbCurCfgVirtServerIndex)
}

type SlbCurCfgVirtServerTableVirtServerLayer3Only int32

const (
	SlbCurCfgVirtServerTableVirtServerLayer3Only_Layer3Only  SlbCurCfgVirtServerTableVirtServerLayer3Only = 1
	SlbCurCfgVirtServerTableVirtServerLayer3Only_Disabled    SlbCurCfgVirtServerTableVirtServerLayer3Only = 2
	SlbCurCfgVirtServerTableVirtServerLayer3Only_Unsupported SlbCurCfgVirtServerTableVirtServerLayer3Only = 2147483647
)

type SlbCurCfgVirtServerTableVirtServerState int32

const (
	SlbCurCfgVirtServerTableVirtServerState_Enabled     SlbCurCfgVirtServerTableVirtServerState = 2
	SlbCurCfgVirtServerTableVirtServerState_Disabled    SlbCurCfgVirtServerTableVirtServerState = 3
	SlbCurCfgVirtServerTableVirtServerState_Unsupported SlbCurCfgVirtServerTableVirtServerState = 2147483647
)

type SlbCurCfgVirtServerTableVirtServerIpVer int32

const (
	SlbCurCfgVirtServerTableVirtServerIpVer_Ipv4        SlbCurCfgVirtServerTableVirtServerIpVer = 1
	SlbCurCfgVirtServerTableVirtServerIpVer_Ipv6        SlbCurCfgVirtServerTableVirtServerIpVer = 2
	SlbCurCfgVirtServerTableVirtServerIpVer_Unsupported SlbCurCfgVirtServerTableVirtServerIpVer = 2147483647
)

type SlbCurCfgVirtServerTableVirtServerCReset int32

const (
	SlbCurCfgVirtServerTableVirtServerCReset_Enabled     SlbCurCfgVirtServerTableVirtServerCReset = 1
	SlbCurCfgVirtServerTableVirtServerCReset_Disabled    SlbCurCfgVirtServerTableVirtServerCReset = 2
	SlbCurCfgVirtServerTableVirtServerCReset_Unsupported SlbCurCfgVirtServerTableVirtServerCReset = 2147483647
)

type SlbCurCfgVirtServerTableVirtServerIsDnsSecVip int32

const (
	SlbCurCfgVirtServerTableVirtServerIsDnsSecVip_No          SlbCurCfgVirtServerTableVirtServerIsDnsSecVip = 0
	SlbCurCfgVirtServerTableVirtServerIsDnsSecVip_Yes         SlbCurCfgVirtServerTableVirtServerIsDnsSecVip = 1
	SlbCurCfgVirtServerTableVirtServerIsDnsSecVip_Unsupported SlbCurCfgVirtServerTableVirtServerIsDnsSecVip = 2147483647
)

type SlbCurCfgVirtServerTableParams struct {
	// Virtual Server Number
	VirtServerIndex int32 `json:"VirtServerIndex,omitempty"`
	// IP address of the virtual server.
	VirtServerIpAddress string `json:"VirtServerIpAddress,omitempty"`
	// Enable or disable layer3 only balancing.
	VirtServerLayer3Only SlbCurCfgVirtServerTableVirtServerLayer3Only `json:"VirtServerLayer3Only,omitempty"`
	// Enable or disable the virtual server.
	VirtServerState SlbCurCfgVirtServerTableVirtServerState `json:"VirtServerState,omitempty"`
	// The domain name of the virtual server.
	VirtServerDname string `json:"VirtServerDname,omitempty"`
	// The default BW contract of virtual server.
	VirtServerBwmContract int32 `json:"VirtServerBwmContract,omitempty"`
	// The virtual server Global SLB weight.
	VirtServerWeight uint64 `json:"VirtServerWeight,omitempty"`
	// The virtual server Global SLB availability.
	VirtServerAvail uint64 `json:"VirtServerAvail,omitempty"`
	// The Global SLB rules for the domain.
	// The rules are presented in bitmap format.
	// in receiving order:
	// OCTET 1  OCTET 2  .....
	// xxxxxxxx xxxxxxxx .....
	// |     || |_ server 9
	// |     ||
	// |     ||___ server 8
	// |     |____ server 7
	// |       .    .   .
	// |__________ server 1
	// where x : 1 - The represented rule belongs to the domain
	// 0 - The represented rule does not belong to the domain
	VirtServerRule string `json:"VirtServerRule,omitempty"`
	// The name of the virtual server.
	VirtServerVname string `json:"VirtServerVname,omitempty"`
	// The type of IP address.
	VirtServerIpVer SlbCurCfgVirtServerTableVirtServerIpVer `json:"VirtServerIpVer,omitempty"`
	// The IPv6 address of the virtual server. Address should be 4-byte
	// hexadecimal colon notation.
	// Valid IPv6 address should be in any of the following forms
	// xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx or
	// xxxx::xxxx:xxxx:xxxx:xxxx or ::xxxx
	VirtServerIpv6Addr string `json:"VirtServerIpv6Addr,omitempty"`
	// Enable/disable client connection reset for invalid VPORT.
	VirtServerCReset SlbCurCfgVirtServerTableVirtServerCReset `json:"VirtServerCReset,omitempty"`
	// Source Network Classifier of the virtual server.
	VirtServerSrcNetwork string `json:"VirtServerSrcNetwork,omitempty"`
	// IP address of the NAT.
	VirtServerNat string `json:"VirtServerNat,omitempty"`
	// The IPv6 address of the NAT. Address should be 4-byte
	// hexadecimal colon notation.
	// Valid IPv6 address should be in any of the following forms
	// xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx or
	// xxxx::xxxx:xxxx:xxxx:xxxx or ::xxxx
	VirtServerNat6 string `json:"VirtServerNat6,omitempty"`
	// It returns Yes(1) if virtual server is a DNS Responder VIP, else returns no(0).
	// This mib is added for Vision purpose.
	VirtServerIsDnsSecVip SlbCurCfgVirtServerTableVirtServerIsDnsSecVip `json:"VirtServerIsDnsSecVip,omitempty"`
}

func (p SlbCurCfgVirtServerTableParams) iMABean() {}
