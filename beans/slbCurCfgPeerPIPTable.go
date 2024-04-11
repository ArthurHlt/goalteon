package beans

import (
	"fmt"
	"reflect"
)

// SlbCurCfgPeerPIPTable The table of Peer PIP configuration in the current_config.
// Note:This mib is not supported for VX instance of Virtualization.
type SlbCurCfgPeerPIPTable struct {
	// The index for peer PIP.
	SlbCurCfgPeerPIPIndex int32
	Params                *SlbCurCfgPeerPIPTableParams
}

func NewSlbCurCfgPeerPIPTableList() *SlbCurCfgPeerPIPTable {
	return &SlbCurCfgPeerPIPTable{}
}

func NewSlbCurCfgPeerPIPTable(
	slbCurCfgPeerPIPIndex int32,
	params *SlbCurCfgPeerPIPTableParams,
) *SlbCurCfgPeerPIPTable {
	return &SlbCurCfgPeerPIPTable{
		SlbCurCfgPeerPIPIndex: slbCurCfgPeerPIPIndex,
		Params:                params,
	}
}

func (c *SlbCurCfgPeerPIPTable) Name() string {
	return "SlbCurCfgPeerPIPTable"
}

func (c *SlbCurCfgPeerPIPTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbCurCfgPeerPIPTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbCurCfgPeerPIPTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbCurCfgPeerPIPIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbCurCfgPeerPIPIndex)
}

type SlbCurCfgPeerPIPTablePIPVersion int32

const (
	SlbCurCfgPeerPIPTablePIPVersion_Unknown     SlbCurCfgPeerPIPTablePIPVersion = 0
	SlbCurCfgPeerPIPTablePIPVersion_Ipv4        SlbCurCfgPeerPIPTablePIPVersion = 4
	SlbCurCfgPeerPIPTablePIPVersion_Ipv6        SlbCurCfgPeerPIPTablePIPVersion = 6
	SlbCurCfgPeerPIPTablePIPVersion_Unsupported SlbCurCfgPeerPIPTablePIPVersion = 2147483647
)

type SlbCurCfgPeerPIPTableParams struct {
	// The index for peer PIP.
	PIPIndex int32 `json:"PIPIndex,omitempty"`
	// The IP address of the peer PIP.
	PIPAddr string `json:"PIPAddr,omitempty"`
	// The IPv6 address of the Peer PIP.
	// Address should be 16-byte hexadecimal colon notation.
	// Valid IPv6 address should be in any of the following forms
	// xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx or
	// xxxx::xxxx:xxxx:xxxx:xxxx or ::xxxx
	PIPv6Addr string `json:"PIPv6Addr,omitempty"`
	// The IP version of the peer PIP.
	PIPVersion SlbCurCfgPeerPIPTablePIPVersion `json:"PIPVersion,omitempty"`
}

func (p SlbCurCfgPeerPIPTableParams) iMABean() {}
