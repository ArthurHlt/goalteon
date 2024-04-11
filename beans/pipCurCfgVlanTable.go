package beans

import (
	"fmt"
	"reflect"
)

// PipCurCfgVlanTable The PIP Vlan table Current configuration block.
// Note:This mib is not supported for VX instance of Virtualization.
type PipCurCfgVlanTable struct {
	// PIP address by the instance of the pipCurCfgIndex.
	PipCurCfgVlanPip string
	Params           *PipCurCfgVlanTableParams
}

func NewPipCurCfgVlanTableList() *PipCurCfgVlanTable {
	return &PipCurCfgVlanTable{}
}

func NewPipCurCfgVlanTable(
	pipCurCfgVlanPip string,
	params *PipCurCfgVlanTableParams,
) *PipCurCfgVlanTable {
	return &PipCurCfgVlanTable{
		PipCurCfgVlanPip: pipCurCfgVlanPip,
		Params:           params,
	}
}

func (c *PipCurCfgVlanTable) Name() string {
	return "PipCurCfgVlanTable"
}

func (c *PipCurCfgVlanTable) GetParams() BeanType {
	return c.Params
}

func (c *PipCurCfgVlanTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *PipCurCfgVlanTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.PipCurCfgVlanPip).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.PipCurCfgVlanPip)
}

type PipCurCfgVlanTableParams struct {
	// PIP address by the instance of the pipCurCfgIndex.
	Pip string `json:"Pip,omitempty"`
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
	PipMap string `json:"PipMap,omitempty"`
}

func (p PipCurCfgVlanTableParams) iMABean() {}
