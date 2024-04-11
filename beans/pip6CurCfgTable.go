package beans

import (
	"fmt"
	"reflect"
)

// Pip6CurCfgTable The IPv6 PIP table in the current configuration block.
// Note:This mib is not supported for VX instance of Virtualization.
type Pip6CurCfgTable struct {
	// Current configured IPv6 Proxy IP address.
	// 	it should be of the form:
	// 	<OID>.4.16.x.x.x.x.x.x.x.x.x.x.x.x.x.x.x.x
	// 	Here, 4 stands for IPv6 type
	// 	16 - stands for number of octets.
	// 	x - the IPv6 address octets
	// 	example: for 1111::1111 address, it should be
	// 	pip6CurCfgPip.4.16.17.17.0.0.0.0.0.0.0.0.0.0.0.0.17.17
	Pip6CurCfgPip string
	Params        *Pip6CurCfgTableParams
}

func NewPip6CurCfgTableList() *Pip6CurCfgTable {
	return &Pip6CurCfgTable{}
}

func NewPip6CurCfgTable(
	pip6CurCfgPip string,
	params *Pip6CurCfgTableParams,
) *Pip6CurCfgTable {
	return &Pip6CurCfgTable{
		Pip6CurCfgPip: pip6CurCfgPip,
		Params:        params,
	}
}

func (c *Pip6CurCfgTable) Name() string {
	return "Pip6CurCfgTable"
}

func (c *Pip6CurCfgTable) GetParams() BeanType {
	return c.Params
}

func (c *Pip6CurCfgTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *Pip6CurCfgTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.Pip6CurCfgPip).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.Pip6CurCfgPip)
}

type Pip6CurCfgTableParams struct {
	// Current configured IPv6 Proxy IP address.
	// 	it should be of the form:
	// 	<OID>.4.16.x.x.x.x.x.x.x.x.x.x.x.x.x.x.x.x
	// 	Here, 4 stands for IPv6 type
	// 	16 - stands for number of octets.
	// 	x - the IPv6 address octets
	// 	example: for 1111::1111 address, it should be
	// 	pip6CurCfgPip.4.16.17.17.0.0.0.0.0.0.0.0.0.0.0.0.17.17
	Pip string `json:"Pip,omitempty"`
	// This is a bit map of port.
	// in receiving order:
	// OCTET 1  OCTET 2  .....
	// xxxxxxxx xxxxxxxx .....
	// ||    || |_ port 8
	// ||    ||
	// ||    ||___ port 7
	// ||    |____ port 6
	// ||_________ port 1
	// |__________ unused
	// where x : 1 - PIP is used for this port.
	// x : 0 - PIP Not used for this port.
	PortMap string `json:"PortMap,omitempty"`
	// This is a bit map of VLAN.
	// in receiving order:
	// OCTET 1  OCTET 2  .....
	// xxxxxxxx xxxxxxxx .....
	// ||    || |_ VLAN 8
	// ||    ||
	// ||    ||___ VLAN 7
	// ||    |____ VLAN 6
	// ||_________ VLAN 1
	// |__________ unused
	// where x : 1 - PIP is used for this VLAN.
	// x : 0 - PIP Not used for this VLAN.
	VlanMap string `json:"VlanMap,omitempty"`
}

func (p Pip6CurCfgTableParams) iMABean() {}
