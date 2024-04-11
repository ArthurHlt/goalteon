package beans

import (
	"fmt"
	"reflect"
)

// SlbNewCfgClusterPeerTable The table of sync cluster peer switch configuration in the current_config.
// Note:This mib is not supported for VX instance of Virtualization.
type SlbNewCfgClusterPeerTable struct {
	// The index for cluster sync peer switchs.
	SlbNewCfgClusterPeerIndex int32
	Params                    *SlbNewCfgClusterPeerTableParams
}

func NewSlbNewCfgClusterPeerTableList() *SlbNewCfgClusterPeerTable {
	return &SlbNewCfgClusterPeerTable{}
}

func NewSlbNewCfgClusterPeerTable(
	slbNewCfgClusterPeerIndex int32,
	params *SlbNewCfgClusterPeerTableParams,
) *SlbNewCfgClusterPeerTable {
	return &SlbNewCfgClusterPeerTable{
		SlbNewCfgClusterPeerIndex: slbNewCfgClusterPeerIndex,
		Params:                    params,
	}
}

func (c *SlbNewCfgClusterPeerTable) Name() string {
	return "SlbNewCfgClusterPeerTable"
}

func (c *SlbNewCfgClusterPeerTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbNewCfgClusterPeerTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbNewCfgClusterPeerTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbNewCfgClusterPeerIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbNewCfgClusterPeerIndex)
}

type SlbNewCfgClusterPeerTableIpVer int32

const (
	SlbNewCfgClusterPeerTableIpVer_Ipv4        SlbNewCfgClusterPeerTableIpVer = 1
	SlbNewCfgClusterPeerTableIpVer_Ipv6        SlbNewCfgClusterPeerTableIpVer = 2
	SlbNewCfgClusterPeerTableIpVer_Unsupported SlbNewCfgClusterPeerTableIpVer = 2147483647
)

type SlbNewCfgClusterPeerTableDelete int32

const (
	SlbNewCfgClusterPeerTableDelete_Other       SlbNewCfgClusterPeerTableDelete = 1
	SlbNewCfgClusterPeerTableDelete_Delete      SlbNewCfgClusterPeerTableDelete = 2
	SlbNewCfgClusterPeerTableDelete_Unsupported SlbNewCfgClusterPeerTableDelete = 2147483647
)

type SlbNewCfgClusterPeerTableParams struct {
	// The index for cluster sync peer switchs.
	Index int32 `json:"Index,omitempty"`
	// The IP address of the cluster sync peer switch.
	IpAddr string `json:"IpAddr,omitempty"`
	// The type of cluster peer IP address version.
	IpVer SlbNewCfgClusterPeerTableIpVer `json:"IpVer,omitempty"`
	// The IPv6 address of the cluster Peer switch.
	// Address should be 16-byte hexadecimal colon notation.
	// Valid IPv6 address should be in any of the following forms
	// xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx or
	// xxxx::xxxx:xxxx:xxxx:xxxx or ::xxxx
	Ipv6Addr string `json:"Ipv6Addr,omitempty"`
	// By setting the value to delete(2), the entire row is
	// deleted.
	Delete SlbNewCfgClusterPeerTableDelete `json:"Delete,omitempty"`
}

func (p SlbNewCfgClusterPeerTableParams) iMABean() {}
