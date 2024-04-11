package beans

import (
	"fmt"
	"reflect"
)

// SlbNewCfgPeerPIPTable The table of Peer PIP configuration in the new_config. Either v4 or v6 IP can be configured
// in a single row.
// Note:This mib is not supported for VX instance of Virtualization.
type SlbNewCfgPeerPIPTable struct {
	// The index for peer PIP.
	SlbNewCfgPeerPIPIndex int32
	Params                *SlbNewCfgPeerPIPTableParams
}

func NewSlbNewCfgPeerPIPTableList() *SlbNewCfgPeerPIPTable {
	return &SlbNewCfgPeerPIPTable{}
}

func NewSlbNewCfgPeerPIPTable(
	slbNewCfgPeerPIPIndex int32,
	params *SlbNewCfgPeerPIPTableParams,
) *SlbNewCfgPeerPIPTable {
	return &SlbNewCfgPeerPIPTable{
		SlbNewCfgPeerPIPIndex: slbNewCfgPeerPIPIndex,
		Params:                params,
	}
}

func (c *SlbNewCfgPeerPIPTable) Name() string {
	return "SlbNewCfgPeerPIPTable"
}

func (c *SlbNewCfgPeerPIPTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbNewCfgPeerPIPTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbNewCfgPeerPIPTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbNewCfgPeerPIPIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbNewCfgPeerPIPIndex)
}

type SlbNewCfgPeerPIPTablePIPDelete int32

const (
	SlbNewCfgPeerPIPTablePIPDelete_Other       SlbNewCfgPeerPIPTablePIPDelete = 1
	SlbNewCfgPeerPIPTablePIPDelete_Delete      SlbNewCfgPeerPIPTablePIPDelete = 2
	SlbNewCfgPeerPIPTablePIPDelete_Unsupported SlbNewCfgPeerPIPTablePIPDelete = 2147483647
)

type SlbNewCfgPeerPIPTablePIPVersion int32

const (
	SlbNewCfgPeerPIPTablePIPVersion_Unknown     SlbNewCfgPeerPIPTablePIPVersion = 0
	SlbNewCfgPeerPIPTablePIPVersion_Ipv4        SlbNewCfgPeerPIPTablePIPVersion = 4
	SlbNewCfgPeerPIPTablePIPVersion_Ipv6        SlbNewCfgPeerPIPTablePIPVersion = 6
	SlbNewCfgPeerPIPTablePIPVersion_Unsupported SlbNewCfgPeerPIPTablePIPVersion = 2147483647
)

type SlbNewCfgPeerPIPTableParams struct {
	// The index for peer PIP.
	PIPIndex int32 `json:"PIPIndex,omitempty"`
	// The IP address of the peer PIP. Only new entry can be created,cannot modify existing ip.
	PIPAddr string `json:"PIPAddr,omitempty"`
	// The IPv6 address of the Peer PIP.
	// Address should be 16-byte hexadecimal colon notation.
	// Valid IPv6 address should be in any of the following forms
	// xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx or
	// xxxx::xxxx:xxxx:xxxx:xxxx or ::xxxx
	// 	 Only new entry can be created,cannot modify existing ip.
	PIPv6Addr string `json:"PIPv6Addr,omitempty"`
	// By setting the value to delete(2), the entire row is
	// deleted.
	PIPDelete SlbNewCfgPeerPIPTablePIPDelete `json:"PIPDelete,omitempty"`
	// The IP version of the peer PIP.
	PIPVersion SlbNewCfgPeerPIPTablePIPVersion `json:"PIPVersion,omitempty"`
}

func (p SlbNewCfgPeerPIPTableParams) iMABean() {}
