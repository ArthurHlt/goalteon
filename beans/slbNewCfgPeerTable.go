package beans

import (
	"fmt"
	"reflect"
)

// SlbNewCfgPeerTable The table of Synch Peer Switch configuration in the new_config.
// Note:This mib is not supported for VX instance of Virtualization.
type SlbNewCfgPeerTable struct {
	// The index number for synch peer switchs.
	SlbNewCfgPeerIndex int32
	Params             *SlbNewCfgPeerTableParams
}

func NewSlbNewCfgPeerTableList() *SlbNewCfgPeerTable {
	return &SlbNewCfgPeerTable{}
}

func NewSlbNewCfgPeerTable(
	slbNewCfgPeerIndex int32,
	params *SlbNewCfgPeerTableParams,
) *SlbNewCfgPeerTable {
	return &SlbNewCfgPeerTable{
		SlbNewCfgPeerIndex: slbNewCfgPeerIndex,
		Params:             params,
	}
}

func (c *SlbNewCfgPeerTable) Name() string {
	return "SlbNewCfgPeerTable"
}

func (c *SlbNewCfgPeerTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbNewCfgPeerTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbNewCfgPeerTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbNewCfgPeerIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbNewCfgPeerIndex)
}

type SlbNewCfgPeerTableState int32

const (
	SlbNewCfgPeerTableState_Enable      SlbNewCfgPeerTableState = 1
	SlbNewCfgPeerTableState_Disable     SlbNewCfgPeerTableState = 2
	SlbNewCfgPeerTableState_Unsupported SlbNewCfgPeerTableState = 2147483647
)

type SlbNewCfgPeerTableDelete int32

const (
	SlbNewCfgPeerTableDelete_Other       SlbNewCfgPeerTableDelete = 1
	SlbNewCfgPeerTableDelete_Delete      SlbNewCfgPeerTableDelete = 2
	SlbNewCfgPeerTableDelete_Unsupported SlbNewCfgPeerTableDelete = 2147483647
)

type SlbNewCfgPeerTableIpVersion int32

const (
	SlbNewCfgPeerTableIpVersion_Ipv4        SlbNewCfgPeerTableIpVersion = 4
	SlbNewCfgPeerTableIpVersion_Ipv6        SlbNewCfgPeerTableIpVersion = 6
	SlbNewCfgPeerTableIpVersion_Unsupported SlbNewCfgPeerTableIpVersion = 2147483647
)

type SlbNewCfgPeerTableParams struct {
	// The index number for synch peer switchs.
	Index int32 `json:"Index,omitempty"`
	// The IP address of the peer switch.
	IpAddr string `json:"IpAddr,omitempty"`
	// Enable or disable the peer switch.
	State SlbNewCfgPeerTableState `json:"State,omitempty"`
	// By setting the value to delete(2), the entire row is
	// deleted.
	Delete SlbNewCfgPeerTableDelete `json:"Delete,omitempty"`
	// The IPv6 address of the Peer switch.
	// Address should be 16-byte hexadecimal colon notation.
	// Valid IPv6 address should be in any of the following forms
	// xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx or
	// xxxx::xxxx:xxxx:xxxx:xxxx or ::xxxx
	Ipv6Addr string `json:"Ipv6Addr,omitempty"`
	// Peer Ip Address Version.
	IpVersion SlbNewCfgPeerTableIpVersion `json:"IpVersion,omitempty"`
}

func (p SlbNewCfgPeerTableParams) iMABean() {}
