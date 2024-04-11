package beans

import (
	"fmt"
	"reflect"
)

// PipCurCfgTable The PIP table in the current configuration block.
// Note:This mib is not supported for VX instance of Virtualization.
type PipCurCfgTable struct {
	// Proxy IP address by the instance of the pipCurCfgIndex.
	PipCurCfgPip string
	Params       *PipCurCfgTableParams
}

func NewPipCurCfgTableList() *PipCurCfgTable {
	return &PipCurCfgTable{}
}

func NewPipCurCfgTable(
	pipCurCfgPip string,
	params *PipCurCfgTableParams,
) *PipCurCfgTable {
	return &PipCurCfgTable{
		PipCurCfgPip: pipCurCfgPip,
		Params:       params,
	}
}

func (c *PipCurCfgTable) Name() string {
	return "PipCurCfgTable"
}

func (c *PipCurCfgTable) GetParams() BeanType {
	return c.Params
}

func (c *PipCurCfgTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *PipCurCfgTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.PipCurCfgPip).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.PipCurCfgPip)
}

type PipCurCfgTableParams struct {
	// Proxy IP address by the instance of the pipCurCfgIndex.
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

func (p PipCurCfgTableParams) iMABean() {}
