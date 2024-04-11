package beans

import (
	"fmt"
	"reflect"
)

// SlbNewCfgEnhVirtServerTable The table of virtual Servers.
// Note:This mib is not supported for VX instance of Virtualization.
type SlbNewCfgEnhVirtServerTable struct {
	// The Virtual Server Number
	SlbNewCfgEnhVirtServerIndex string
	Params                      *SlbNewCfgEnhVirtServerTableParams
}

func NewSlbNewCfgEnhVirtServerTableList() *SlbNewCfgEnhVirtServerTable {
	return &SlbNewCfgEnhVirtServerTable{}
}

func NewSlbNewCfgEnhVirtServerTable(
	slbNewCfgEnhVirtServerIndex string,
	params *SlbNewCfgEnhVirtServerTableParams,
) *SlbNewCfgEnhVirtServerTable {
	return &SlbNewCfgEnhVirtServerTable{
		SlbNewCfgEnhVirtServerIndex: slbNewCfgEnhVirtServerIndex,
		Params:                      params,
	}
}

func (c *SlbNewCfgEnhVirtServerTable) Name() string {
	return "SlbNewCfgEnhVirtServerTable"
}

func (c *SlbNewCfgEnhVirtServerTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbNewCfgEnhVirtServerTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbNewCfgEnhVirtServerTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbNewCfgEnhVirtServerIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbNewCfgEnhVirtServerIndex)
}

type SlbNewCfgEnhVirtServerTableVirtServerLayer3Only int32

const (
	SlbNewCfgEnhVirtServerTableVirtServerLayer3Only_Layer3Only  SlbNewCfgEnhVirtServerTableVirtServerLayer3Only = 1
	SlbNewCfgEnhVirtServerTableVirtServerLayer3Only_Disabled    SlbNewCfgEnhVirtServerTableVirtServerLayer3Only = 2
	SlbNewCfgEnhVirtServerTableVirtServerLayer3Only_Unsupported SlbNewCfgEnhVirtServerTableVirtServerLayer3Only = 2147483647
)

type SlbNewCfgEnhVirtServerTableVirtServerState int32

const (
	SlbNewCfgEnhVirtServerTableVirtServerState_Enabled     SlbNewCfgEnhVirtServerTableVirtServerState = 2
	SlbNewCfgEnhVirtServerTableVirtServerState_Disabled    SlbNewCfgEnhVirtServerTableVirtServerState = 3
	SlbNewCfgEnhVirtServerTableVirtServerState_Unsupported SlbNewCfgEnhVirtServerTableVirtServerState = 2147483647
)

type SlbNewCfgEnhVirtServerTableVirtServerDelete int32

const (
	SlbNewCfgEnhVirtServerTableVirtServerDelete_Other       SlbNewCfgEnhVirtServerTableVirtServerDelete = 1
	SlbNewCfgEnhVirtServerTableVirtServerDelete_Delete      SlbNewCfgEnhVirtServerTableVirtServerDelete = 2
	SlbNewCfgEnhVirtServerTableVirtServerDelete_Unsupported SlbNewCfgEnhVirtServerTableVirtServerDelete = 2147483647
)

type SlbNewCfgEnhVirtServerTableVirtServerIpVer int32

const (
	SlbNewCfgEnhVirtServerTableVirtServerIpVer_Ipv4        SlbNewCfgEnhVirtServerTableVirtServerIpVer = 1
	SlbNewCfgEnhVirtServerTableVirtServerIpVer_Ipv6        SlbNewCfgEnhVirtServerTableVirtServerIpVer = 2
	SlbNewCfgEnhVirtServerTableVirtServerIpVer_Unsupported SlbNewCfgEnhVirtServerTableVirtServerIpVer = 2147483647
)

type SlbNewCfgEnhVirtServerTableVirtServerCReset int32

const (
	SlbNewCfgEnhVirtServerTableVirtServerCReset_Enabled     SlbNewCfgEnhVirtServerTableVirtServerCReset = 1
	SlbNewCfgEnhVirtServerTableVirtServerCReset_Disabled    SlbNewCfgEnhVirtServerTableVirtServerCReset = 2
	SlbNewCfgEnhVirtServerTableVirtServerCReset_Unsupported SlbNewCfgEnhVirtServerTableVirtServerCReset = 2147483647
)

type SlbNewCfgEnhVirtServerTableVirtServerIsDnsSecVip int32

const (
	SlbNewCfgEnhVirtServerTableVirtServerIsDnsSecVip_No          SlbNewCfgEnhVirtServerTableVirtServerIsDnsSecVip = 0
	SlbNewCfgEnhVirtServerTableVirtServerIsDnsSecVip_Yes         SlbNewCfgEnhVirtServerTableVirtServerIsDnsSecVip = 1
	SlbNewCfgEnhVirtServerTableVirtServerIsDnsSecVip_Unsupported SlbNewCfgEnhVirtServerTableVirtServerIsDnsSecVip = 2147483647
)

type SlbNewCfgEnhVirtServerTableVirtServerAvailPersist int32

const (
	SlbNewCfgEnhVirtServerTableVirtServerAvailPersist_Enabled     SlbNewCfgEnhVirtServerTableVirtServerAvailPersist = 1
	SlbNewCfgEnhVirtServerTableVirtServerAvailPersist_Disabled    SlbNewCfgEnhVirtServerTableVirtServerAvailPersist = 2
	SlbNewCfgEnhVirtServerTableVirtServerAvailPersist_Unsupported SlbNewCfgEnhVirtServerTableVirtServerAvailPersist = 2147483647
)

type SlbNewCfgEnhVirtServerTableVirtServerRtSrcMac int32

const (
	SlbNewCfgEnhVirtServerTableVirtServerRtSrcMac_Enabled     SlbNewCfgEnhVirtServerTableVirtServerRtSrcMac = 1
	SlbNewCfgEnhVirtServerTableVirtServerRtSrcMac_Disabled    SlbNewCfgEnhVirtServerTableVirtServerRtSrcMac = 2
	SlbNewCfgEnhVirtServerTableVirtServerRtSrcMac_Unsupported SlbNewCfgEnhVirtServerTableVirtServerRtSrcMac = 2147483647
)

type SlbNewCfgEnhVirtServerTableVirtServerCreationType int32

const (
	SlbNewCfgEnhVirtServerTableVirtServerCreationType_General     SlbNewCfgEnhVirtServerTableVirtServerCreationType = 0
	SlbNewCfgEnhVirtServerTableVirtServerCreationType_Wizard      SlbNewCfgEnhVirtServerTableVirtServerCreationType = 1
	SlbNewCfgEnhVirtServerTableVirtServerCreationType_Unsupported SlbNewCfgEnhVirtServerTableVirtServerCreationType = 2147483647
)

type SlbNewCfgEnhVirtServerTableVirtServerdad int32

const (
	SlbNewCfgEnhVirtServerTableVirtServerdad_Enabled     SlbNewCfgEnhVirtServerTableVirtServerdad = 1
	SlbNewCfgEnhVirtServerTableVirtServerdad_Disabled    SlbNewCfgEnhVirtServerTableVirtServerdad = 2
	SlbNewCfgEnhVirtServerTableVirtServerdad_Unsupported SlbNewCfgEnhVirtServerTableVirtServerdad = 2147483647
)

type SlbNewCfgEnhVirtServerTableParams struct {
	// The Virtual Server Number
	VirtServerIndex string `json:"VirtServerIndex,omitempty"`
	// IP address of the virtual server.
	VirtServerIpAddress string `json:"VirtServerIpAddress,omitempty"`
	// Enable or disable layer3 only balancing.
	VirtServerLayer3Only SlbNewCfgEnhVirtServerTableVirtServerLayer3Only `json:"VirtServerLayer3Only,omitempty"`
	// Enable or disable the virtual server.
	VirtServerState SlbNewCfgEnhVirtServerTableVirtServerState `json:"VirtServerState,omitempty"`
	// The domain name of the virtual server.
	VirtServerDname string `json:"VirtServerDname,omitempty"`
	// The default BW contract number of the virtual server.
	VirtServerBwmContract int32 `json:"VirtServerBwmContract,omitempty"`
	// By setting the value to delete(2), the entire row is deleted.
	VirtServerDelete SlbNewCfgEnhVirtServerTableVirtServerDelete `json:"VirtServerDelete,omitempty"`
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
	VirtServerIpVer SlbNewCfgEnhVirtServerTableVirtServerIpVer `json:"VirtServerIpVer,omitempty"`
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
	VirtServerCReset SlbNewCfgEnhVirtServerTableVirtServerCReset `json:"VirtServerCReset,omitempty"`
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
	VirtServerIsDnsSecVip SlbNewCfgEnhVirtServerTableVirtServerIsDnsSecVip `json:"VirtServerIsDnsSecVip,omitempty"`
	// Enable/disable GSLB availability persistence.
	VirtServerAvailPersist SlbNewCfgEnhVirtServerTableVirtServerAvailPersist `json:"VirtServerAvailPersist,omitempty"`
	// Associate a real wanlink server to virtual server.
	VirtServerWanlink string `json:"VirtServerWanlink,omitempty"`
	// Enable/disable return to source mac address
	VirtServerRtSrcMac SlbNewCfgEnhVirtServerTableVirtServerRtSrcMac `json:"VirtServerRtSrcMac,omitempty"`
	// Virtual Server Creation Type
	VirtServerCreationType SlbNewCfgEnhVirtServerTableVirtServerCreationType `json:"VirtServerCreationType,omitempty"`
	// Enable/disable duplicate address detection
	VirtServerdad SlbNewCfgEnhVirtServerTableVirtServerdad `json:"VirtServerdad,omitempty"`
}

func (p SlbNewCfgEnhVirtServerTableParams) iMABean() {}
