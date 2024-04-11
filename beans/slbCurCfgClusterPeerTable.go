package beans

import (
	"fmt"
	"reflect"
)

// SlbCurCfgClusterPeerTable The sync cluster peer wwitch configuration in the current_config.
// Note:This mib is not supported for VX instance of Virtualization.
type SlbCurCfgClusterPeerTable struct {
	// The index for cluster sync peer switchs.
	SlbCurCfgClusterPeerIndex int32
	Params                    *SlbCurCfgClusterPeerTableParams
}

func NewSlbCurCfgClusterPeerTableList() *SlbCurCfgClusterPeerTable {
	return &SlbCurCfgClusterPeerTable{}
}

func NewSlbCurCfgClusterPeerTable(
	slbCurCfgClusterPeerIndex int32,
	params *SlbCurCfgClusterPeerTableParams,
) *SlbCurCfgClusterPeerTable {
	return &SlbCurCfgClusterPeerTable{
		SlbCurCfgClusterPeerIndex: slbCurCfgClusterPeerIndex,
		Params:                    params,
	}
}

func (c *SlbCurCfgClusterPeerTable) Name() string {
	return "SlbCurCfgClusterPeerTable"
}

func (c *SlbCurCfgClusterPeerTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbCurCfgClusterPeerTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbCurCfgClusterPeerTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbCurCfgClusterPeerIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbCurCfgClusterPeerIndex)
}

type SlbCurCfgClusterPeerTableIpVer int32

const (
	SlbCurCfgClusterPeerTableIpVer_Ipv4        SlbCurCfgClusterPeerTableIpVer = 1
	SlbCurCfgClusterPeerTableIpVer_Ipv6        SlbCurCfgClusterPeerTableIpVer = 2
	SlbCurCfgClusterPeerTableIpVer_Unsupported SlbCurCfgClusterPeerTableIpVer = 2147483647
)

type SlbCurCfgClusterPeerTableParams struct {
	// The index for cluster sync peer switchs.
	Index int32 `json:"Index,omitempty"`
	// The IP address of the cluster sync peer switch.
	IpAddr string `json:"IpAddr,omitempty"`
	// The type of cluster sync peer IP address.
	IpVer SlbCurCfgClusterPeerTableIpVer `json:"IpVer,omitempty"`
	// The IPv6 address of the cluster Peer switch.
	// Address should be 16-byte hexadecimal colon notation.
	// Valid IPv6 address should be in any of the following forms
	// xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx or
	// xxxx::xxxx:xxxx:xxxx:xxxx or ::xxxx
	Ipv6Addr string `json:"Ipv6Addr,omitempty"`
}

func (p SlbCurCfgClusterPeerTableParams) iMABean() {}
