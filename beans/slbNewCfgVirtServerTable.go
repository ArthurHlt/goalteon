package beans

import (
	"fmt"
	"reflect"
)

// SlbNewCfgVirtServerTable The table of virtual Servers.
// Note:This mib is not supported for VX instance of Virtualization.
type SlbNewCfgVirtServerTable struct {
	// The Virtual Server Number
	SlbNewCfgVirtServerIndex int32
	Params                   *SlbNewCfgVirtServerTableParams
}

func NewSlbNewCfgVirtServerTableList() *SlbNewCfgVirtServerTable {
	return &SlbNewCfgVirtServerTable{}
}

func NewSlbNewCfgVirtServerTable(
	slbNewCfgVirtServerIndex int32,
	params *SlbNewCfgVirtServerTableParams,
) *SlbNewCfgVirtServerTable {
	return &SlbNewCfgVirtServerTable{
		SlbNewCfgVirtServerIndex: slbNewCfgVirtServerIndex,
		Params:                   params,
	}
}

func (c *SlbNewCfgVirtServerTable) Name() string {
	return "SlbNewCfgVirtServerTable"
}

func (c *SlbNewCfgVirtServerTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbNewCfgVirtServerTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbNewCfgVirtServerTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbNewCfgVirtServerIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbNewCfgVirtServerIndex)
}

type SlbNewCfgVirtServerTableVirtServerLayer3Only int32

const (
	SlbNewCfgVirtServerTableVirtServerLayer3Only_Layer3Only  SlbNewCfgVirtServerTableVirtServerLayer3Only = 1
	SlbNewCfgVirtServerTableVirtServerLayer3Only_Disabled    SlbNewCfgVirtServerTableVirtServerLayer3Only = 2
	SlbNewCfgVirtServerTableVirtServerLayer3Only_Unsupported SlbNewCfgVirtServerTableVirtServerLayer3Only = 2147483647
)

type SlbNewCfgVirtServerTableVirtServerState int32

const (
	SlbNewCfgVirtServerTableVirtServerState_Enabled     SlbNewCfgVirtServerTableVirtServerState = 2
	SlbNewCfgVirtServerTableVirtServerState_Disabled    SlbNewCfgVirtServerTableVirtServerState = 3
	SlbNewCfgVirtServerTableVirtServerState_Unsupported SlbNewCfgVirtServerTableVirtServerState = 2147483647
)

type SlbNewCfgVirtServerTableVirtServerDelete int32

const (
	SlbNewCfgVirtServerTableVirtServerDelete_Other       SlbNewCfgVirtServerTableVirtServerDelete = 1
	SlbNewCfgVirtServerTableVirtServerDelete_Delete      SlbNewCfgVirtServerTableVirtServerDelete = 2
	SlbNewCfgVirtServerTableVirtServerDelete_Unsupported SlbNewCfgVirtServerTableVirtServerDelete = 2147483647
)

type SlbNewCfgVirtServerTableVirtServerIpVer int32

const (
	SlbNewCfgVirtServerTableVirtServerIpVer_Ipv4        SlbNewCfgVirtServerTableVirtServerIpVer = 1
	SlbNewCfgVirtServerTableVirtServerIpVer_Ipv6        SlbNewCfgVirtServerTableVirtServerIpVer = 2
	SlbNewCfgVirtServerTableVirtServerIpVer_Unsupported SlbNewCfgVirtServerTableVirtServerIpVer = 2147483647
)

type SlbNewCfgVirtServerTableVirtServerCReset int32

const (
	SlbNewCfgVirtServerTableVirtServerCReset_Enabled     SlbNewCfgVirtServerTableVirtServerCReset = 1
	SlbNewCfgVirtServerTableVirtServerCReset_Disabled    SlbNewCfgVirtServerTableVirtServerCReset = 2
	SlbNewCfgVirtServerTableVirtServerCReset_Unsupported SlbNewCfgVirtServerTableVirtServerCReset = 2147483647
)

type SlbNewCfgVirtServerTableVirtServerIsDnsSecVip int32

const (
	SlbNewCfgVirtServerTableVirtServerIsDnsSecVip_No          SlbNewCfgVirtServerTableVirtServerIsDnsSecVip = 0
	SlbNewCfgVirtServerTableVirtServerIsDnsSecVip_Yes         SlbNewCfgVirtServerTableVirtServerIsDnsSecVip = 1
	SlbNewCfgVirtServerTableVirtServerIsDnsSecVip_Unsupported SlbNewCfgVirtServerTableVirtServerIsDnsSecVip = 2147483647
)

type SlbNewCfgVirtServerTableParams struct {
	// The Virtual Server Number
	VirtServerIndex int32 `json:"VirtServerIndex,omitempty"`
	// IP address of the virtual server.
	VirtServerIpAddress string `json:"VirtServerIpAddress,omitempty"`
	// Enable or disable layer3 only balancing.
	VirtServerLayer3Only SlbNewCfgVirtServerTableVirtServerLayer3Only `json:"VirtServerLayer3Only,omitempty"`
	// Enable or disable the virtual server.
	VirtServerState SlbNewCfgVirtServerTableVirtServerState `json:"VirtServerState,omitempty"`
	// The domain name of the virtual server.
	VirtServerDname string `json:"VirtServerDname,omitempty"`
	// The default BW contract number of the virtual server.
	VirtServerBwmContract int32 `json:"VirtServerBwmContract,omitempty"`
	// By setting the value to delete(2), the entire row is deleted.
	VirtServerDelete SlbNewCfgVirtServerTableVirtServerDelete `json:"VirtServerDelete,omitempty"`
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
	// The rule to be added to the domain. When read, 0 is returned.
	VirtServerAddRule int32 `json:"VirtServerAddRule,omitempty"`
	// The rule to be removed from the domain. When read, 0 is returned.
	VirtServerRemoveRule int32 `json:"VirtServerRemoveRule,omitempty"`
	// The name of the virtual server.
	VirtServerVname string `json:"VirtServerVname,omitempty"`
	// The type of IP address.
	VirtServerIpVer SlbNewCfgVirtServerTableVirtServerIpVer `json:"VirtServerIpVer,omitempty"`
	// The IPv6 address of the virtual server. Address should be 4-byte
	// hexadecimal colon notation.
	// Valid IPv6 address should be in any of the following forms
	// xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx or
	// xxxx::xxxx:xxxx:xxxx:xxxx or ::xxxx
	VirtServerIpv6Addr string `json:"VirtServerIpv6Addr,omitempty"`
	// The first free service index number of the virtual server.
	// Value 0 will be returned when all 8 virtual services are
	// configured for a virtual server.
	VirtServerFreeServiceIdx uint64 `json:"VirtServerFreeServiceIdx,omitempty"`
	// Enable/disable client connection reset for invalid VPORT.
	VirtServerCReset SlbNewCfgVirtServerTableVirtServerCReset `json:"VirtServerCReset,omitempty"`
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
	// This mib returns Yes(1) if virtual server is a DNS Responder VIP, else returns no(0)
	VirtServerIsDnsSecVip SlbNewCfgVirtServerTableVirtServerIsDnsSecVip `json:"VirtServerIsDnsSecVip,omitempty"`
}

func (p SlbNewCfgVirtServerTableParams) iMABean() {}
