package beans

import (
	"fmt"
	"reflect"
)

// SlbCurCfgPeerTable The table of Synch Peer Switch configuration in the current_config.
// Note:This mib is not supported for VX instance of Virtualization.
type SlbCurCfgPeerTable struct {
	// The index for synch peer switchs.
	SlbCurCfgPeerIndex int32
	Params             *SlbCurCfgPeerTableParams
}

func NewSlbCurCfgPeerTableList() *SlbCurCfgPeerTable {
	return &SlbCurCfgPeerTable{}
}

func NewSlbCurCfgPeerTable(
	slbCurCfgPeerIndex int32,
	params *SlbCurCfgPeerTableParams,
) *SlbCurCfgPeerTable {
	return &SlbCurCfgPeerTable{
		SlbCurCfgPeerIndex: slbCurCfgPeerIndex,
		Params:             params,
	}
}

func (c *SlbCurCfgPeerTable) Name() string {
	return "SlbCurCfgPeerTable"
}

func (c *SlbCurCfgPeerTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbCurCfgPeerTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbCurCfgPeerTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbCurCfgPeerIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbCurCfgPeerIndex)
}

type SlbCurCfgPeerTableState int32

const (
	SlbCurCfgPeerTableState_Enable      SlbCurCfgPeerTableState = 1
	SlbCurCfgPeerTableState_Disable     SlbCurCfgPeerTableState = 2
	SlbCurCfgPeerTableState_Unsupported SlbCurCfgPeerTableState = 2147483647
)

type SlbCurCfgPeerTableIpVersion int32

const (
	SlbCurCfgPeerTableIpVersion_Ipv4        SlbCurCfgPeerTableIpVersion = 4
	SlbCurCfgPeerTableIpVersion_Ipv6        SlbCurCfgPeerTableIpVersion = 6
	SlbCurCfgPeerTableIpVersion_Unsupported SlbCurCfgPeerTableIpVersion = 2147483647
)

type SlbCurCfgPeerTableParams struct {
	// The index for synch peer switchs.
	Index int32 `json:"Index,omitempty"`
	// The IP address of the peer switch.
	IpAddr string `json:"IpAddr,omitempty"`
	// Enable or disable the peer switch.
	State SlbCurCfgPeerTableState `json:"State,omitempty"`
	// The IPv6 address of the Peer switch.
	// Address should be 16-byte hexadecimal colon notation.
	// Valid IPv6 address should be in any of the following forms
	// xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx or
	// xxxx::xxxx:xxxx:xxxx:xxxx or ::xxxx
	Ipv6Addr string `json:"Ipv6Addr,omitempty"`
	// Peer Ip Address Version.
	IpVersion SlbCurCfgPeerTableIpVersion `json:"IpVersion,omitempty"`
}

func (p SlbCurCfgPeerTableParams) iMABean() {}
