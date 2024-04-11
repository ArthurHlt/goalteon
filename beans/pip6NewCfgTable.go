package beans

import (
	"fmt"
	"reflect"
)

// Pip6NewCfgTable The IPv6 PIP table in the New configuration block.
// Note:This mib is not supported for VX instance of Virtualization.
type Pip6NewCfgTable struct {
	// New configured IPv6 Proxy IP address.
	// it should be of the form:
	// <OID>.4.16.x.x.x.x.x.x.x.x.x.x.x.x.x.x.x.x
	// Here, 4 stands for IPv6 type
	// 16 - stands for number of octets.
	// x - the IPv6 address octets
	// example: for 1111::1111 address, it should be
	// pip6NewCfgPip.4.16.17.17.0.0.0.0.0.0.0.0.0.0.0.0.17.17
	Pip6NewCfgPip string
	Params        *Pip6NewCfgTableParams
}

func NewPip6NewCfgTableList() *Pip6NewCfgTable {
	return &Pip6NewCfgTable{}
}

func NewPip6NewCfgTable(
	pip6NewCfgPip string,
	params *Pip6NewCfgTableParams,
) *Pip6NewCfgTable {
	return &Pip6NewCfgTable{
		Pip6NewCfgPip: pip6NewCfgPip,
		Params:        params,
	}
}

func (c *Pip6NewCfgTable) Name() string {
	return "Pip6NewCfgTable"
}

func (c *Pip6NewCfgTable) GetParams() BeanType {
	return c.Params
}

func (c *Pip6NewCfgTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *Pip6NewCfgTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.Pip6NewCfgPip).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.Pip6NewCfgPip)
}

type Pip6NewCfgTableDelete int32

const (
	Pip6NewCfgTableDelete_Other       Pip6NewCfgTableDelete = 1
	Pip6NewCfgTableDelete_Delete      Pip6NewCfgTableDelete = 2
	Pip6NewCfgTableDelete_Unsupported Pip6NewCfgTableDelete = 2147483647
)

type Pip6NewCfgTableParams struct {
	// New configured IPv6 Proxy IP address.
	// it should be of the form:
	// <OID>.4.16.x.x.x.x.x.x.x.x.x.x.x.x.x.x.x.x
	// Here, 4 stands for IPv6 type
	// 16 - stands for number of octets.
	// x - the IPv6 address octets
	// example: for 1111::1111 address, it should be
	// pip6NewCfgPip.4.16.17.17.0.0.0.0.0.0.0.0.0.0.0.0.17.17
	Pip string `json:"Pip,omitempty"`
	// The bit map of port for PIP.
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
	// The bit map of VLAN for PIP.
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
	// By setting the value to delete(2), the entire row is deleted.
	Delete Pip6NewCfgTableDelete `json:"Delete,omitempty"`
	// The Port or VLAN to be associated with the PIP. When read, 0 is returned.
	AddPortVlan int32 `json:"AddPortVlan,omitempty"`
	// The Port or VLAN to be disassociated from the PIP. When read,
	// 0 is returned.
	RemovePortVlan int32 `json:"RemovePortVlan,omitempty"`
	// The Port to be associated with the PIP. When read, 0 is returned.
	AddPort int32 `json:"AddPort,omitempty"`
	// The VLAN to be associated with the PIP. When read, 0 is returned.
	AddVlan int32 `json:"AddVlan,omitempty"`
	// The Port to be disassociated from the PIP. When read,
	// 0 is returned.
	RemovePort int32 `json:"RemovePort,omitempty"`
	// The VLAN to be disassociated from the PIP. When read,
	// 0 is returned.
	RemoveVlan int32 `json:"RemoveVlan,omitempty"`
}

func (p Pip6NewCfgTableParams) iMABean() {}
