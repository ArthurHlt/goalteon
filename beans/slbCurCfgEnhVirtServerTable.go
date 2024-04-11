package beans

import (
	"fmt"
	"reflect"
)

// SlbCurCfgEnhVirtServerTable The table of virtual Servers.
// Note:This mib is not supported for VX instance of Virtualization.
type SlbCurCfgEnhVirtServerTable struct {
	// Virtual Server Number in Alphanumeric
	SlbCurCfgEnhVirtServerIndex string
	Params                      *SlbCurCfgEnhVirtServerTableParams
}

func NewSlbCurCfgEnhVirtServerTableList() *SlbCurCfgEnhVirtServerTable {
	return &SlbCurCfgEnhVirtServerTable{}
}

func NewSlbCurCfgEnhVirtServerTable(
	slbCurCfgEnhVirtServerIndex string,
	params *SlbCurCfgEnhVirtServerTableParams,
) *SlbCurCfgEnhVirtServerTable {
	return &SlbCurCfgEnhVirtServerTable{
		SlbCurCfgEnhVirtServerIndex: slbCurCfgEnhVirtServerIndex,
		Params:                      params,
	}
}

func (c *SlbCurCfgEnhVirtServerTable) Name() string {
	return "SlbCurCfgEnhVirtServerTable"
}

func (c *SlbCurCfgEnhVirtServerTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbCurCfgEnhVirtServerTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbCurCfgEnhVirtServerTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbCurCfgEnhVirtServerIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbCurCfgEnhVirtServerIndex)
}

type SlbCurCfgEnhVirtServerTableVirtServerLayer3Only int32

const (
	SlbCurCfgEnhVirtServerTableVirtServerLayer3Only_Layer3Only  SlbCurCfgEnhVirtServerTableVirtServerLayer3Only = 1
	SlbCurCfgEnhVirtServerTableVirtServerLayer3Only_Disabled    SlbCurCfgEnhVirtServerTableVirtServerLayer3Only = 2
	SlbCurCfgEnhVirtServerTableVirtServerLayer3Only_Unsupported SlbCurCfgEnhVirtServerTableVirtServerLayer3Only = 2147483647
)

type SlbCurCfgEnhVirtServerTableVirtServerState int32

const (
	SlbCurCfgEnhVirtServerTableVirtServerState_Enabled     SlbCurCfgEnhVirtServerTableVirtServerState = 2
	SlbCurCfgEnhVirtServerTableVirtServerState_Disabled    SlbCurCfgEnhVirtServerTableVirtServerState = 3
	SlbCurCfgEnhVirtServerTableVirtServerState_Unsupported SlbCurCfgEnhVirtServerTableVirtServerState = 2147483647
)

type SlbCurCfgEnhVirtServerTableVirtServerIpVer int32

const (
	SlbCurCfgEnhVirtServerTableVirtServerIpVer_Ipv4        SlbCurCfgEnhVirtServerTableVirtServerIpVer = 1
	SlbCurCfgEnhVirtServerTableVirtServerIpVer_Ipv6        SlbCurCfgEnhVirtServerTableVirtServerIpVer = 2
	SlbCurCfgEnhVirtServerTableVirtServerIpVer_Unsupported SlbCurCfgEnhVirtServerTableVirtServerIpVer = 2147483647
)

type SlbCurCfgEnhVirtServerTableVirtServerCReset int32

const (
	SlbCurCfgEnhVirtServerTableVirtServerCReset_Enabled     SlbCurCfgEnhVirtServerTableVirtServerCReset = 1
	SlbCurCfgEnhVirtServerTableVirtServerCReset_Disabled    SlbCurCfgEnhVirtServerTableVirtServerCReset = 2
	SlbCurCfgEnhVirtServerTableVirtServerCReset_Unsupported SlbCurCfgEnhVirtServerTableVirtServerCReset = 2147483647
)

type SlbCurCfgEnhVirtServerTableVirtServerIsDnsSecVip int32

const (
	SlbCurCfgEnhVirtServerTableVirtServerIsDnsSecVip_No          SlbCurCfgEnhVirtServerTableVirtServerIsDnsSecVip = 0
	SlbCurCfgEnhVirtServerTableVirtServerIsDnsSecVip_Yes         SlbCurCfgEnhVirtServerTableVirtServerIsDnsSecVip = 1
	SlbCurCfgEnhVirtServerTableVirtServerIsDnsSecVip_Unsupported SlbCurCfgEnhVirtServerTableVirtServerIsDnsSecVip = 2147483647
)

type SlbCurCfgEnhVirtServerTableVirtServerAvailPersist int32

const (
	SlbCurCfgEnhVirtServerTableVirtServerAvailPersist_Enabled     SlbCurCfgEnhVirtServerTableVirtServerAvailPersist = 1
	SlbCurCfgEnhVirtServerTableVirtServerAvailPersist_Disabled    SlbCurCfgEnhVirtServerTableVirtServerAvailPersist = 2
	SlbCurCfgEnhVirtServerTableVirtServerAvailPersist_Unsupported SlbCurCfgEnhVirtServerTableVirtServerAvailPersist = 2147483647
)

type SlbCurCfgEnhVirtServerTableVirtServerRtSrcMac int32

const (
	SlbCurCfgEnhVirtServerTableVirtServerRtSrcMac_Enabled     SlbCurCfgEnhVirtServerTableVirtServerRtSrcMac = 1
	SlbCurCfgEnhVirtServerTableVirtServerRtSrcMac_Disabled    SlbCurCfgEnhVirtServerTableVirtServerRtSrcMac = 2
	SlbCurCfgEnhVirtServerTableVirtServerRtSrcMac_Unsupported SlbCurCfgEnhVirtServerTableVirtServerRtSrcMac = 2147483647
)

type SlbCurCfgEnhVirtServerTableVirtServerCreationType int32

const (
	SlbCurCfgEnhVirtServerTableVirtServerCreationType_General     SlbCurCfgEnhVirtServerTableVirtServerCreationType = 0
	SlbCurCfgEnhVirtServerTableVirtServerCreationType_Wizard      SlbCurCfgEnhVirtServerTableVirtServerCreationType = 1
	SlbCurCfgEnhVirtServerTableVirtServerCreationType_Unsupported SlbCurCfgEnhVirtServerTableVirtServerCreationType = 2147483647
)

type SlbCurCfgEnhVirtServerTableVirtServerdad int32

const (
	SlbCurCfgEnhVirtServerTableVirtServerdad_Enabled     SlbCurCfgEnhVirtServerTableVirtServerdad = 1
	SlbCurCfgEnhVirtServerTableVirtServerdad_Disabled    SlbCurCfgEnhVirtServerTableVirtServerdad = 2
	SlbCurCfgEnhVirtServerTableVirtServerdad_Unsupported SlbCurCfgEnhVirtServerTableVirtServerdad = 2147483647
)

type SlbCurCfgEnhVirtServerTableParams struct {
	// Virtual Server Number in Alphanumeric
	VirtServerIndex string `json:"VirtServerIndex,omitempty"`
	// IP address of the virtual server.
	VirtServerIpAddress string `json:"VirtServerIpAddress,omitempty"`
	// Enable or disable layer3 only balancing.
	VirtServerLayer3Only SlbCurCfgEnhVirtServerTableVirtServerLayer3Only `json:"VirtServerLayer3Only,omitempty"`
	// Enable or disable the virtual server.
	VirtServerState SlbCurCfgEnhVirtServerTableVirtServerState `json:"VirtServerState,omitempty"`
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
	VirtServerIpVer SlbCurCfgEnhVirtServerTableVirtServerIpVer `json:"VirtServerIpVer,omitempty"`
	// The IPv6 address of the virtual server. Address should be 4-byte
	// hexadecimal colon notation.
	// Valid IPv6 address should be in any of the following forms
	// xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx or
	// xxxx::xxxx:xxxx:xxxx:xxxx or ::xxxx
	VirtServerIpv6Addr string `json:"VirtServerIpv6Addr,omitempty"`
	// Enable/disable client connection reset for invalid VPORT.
	VirtServerCReset SlbCurCfgEnhVirtServerTableVirtServerCReset `json:"VirtServerCReset,omitempty"`
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
	VirtServerIsDnsSecVip SlbCurCfgEnhVirtServerTableVirtServerIsDnsSecVip `json:"VirtServerIsDnsSecVip,omitempty"`
	// Enable/disable GSLB availability persistence.
	VirtServerAvailPersist SlbCurCfgEnhVirtServerTableVirtServerAvailPersist `json:"VirtServerAvailPersist,omitempty"`
	// Associate a real wanlink server to virtual server.
	VirtServerWanlink string `json:"VirtServerWanlink,omitempty"`
	// Enable/disable return to source mac address
	VirtServerRtSrcMac SlbCurCfgEnhVirtServerTableVirtServerRtSrcMac `json:"VirtServerRtSrcMac,omitempty"`
	// Virtual Server Creation Type
	VirtServerCreationType SlbCurCfgEnhVirtServerTableVirtServerCreationType `json:"VirtServerCreationType,omitempty"`
	// The total number of DNS directs sent to the virtual server or remote real server.
	VirtServerDnsDirect uint32 `json:"VirtServerDnsDirect,omitempty"`
	// The total number of hits exceeded threshold for virtual server or remote real server.
	VirtServerThresholdExceeded uint32 `json:"VirtServerThresholdExceeded,omitempty"`
	// Enable/disable duplicate address detection
	VirtServerdad SlbCurCfgEnhVirtServerTableVirtServerdad `json:"VirtServerdad,omitempty"`
}

func (p SlbCurCfgEnhVirtServerTableParams) iMABean() {}
