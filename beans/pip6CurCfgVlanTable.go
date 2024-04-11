package beans

import (
	"fmt"
	"reflect"
)

// Pip6CurCfgVlanTable The PIP Vlan table Cur configuration block.
// Note:This mib is not supported for VX instance of Virtualization.
type Pip6CurCfgVlanTable struct {
	// PIP address by the instance of the pip6CurCfgIndex.
	Pip6CurCfgVlanPip string
	Params            *Pip6CurCfgVlanTableParams
}

func NewPip6CurCfgVlanTableList() *Pip6CurCfgVlanTable {
	return &Pip6CurCfgVlanTable{}
}

func NewPip6CurCfgVlanTable(
	pip6CurCfgVlanPip string,
	params *Pip6CurCfgVlanTableParams,
) *Pip6CurCfgVlanTable {
	return &Pip6CurCfgVlanTable{
		Pip6CurCfgVlanPip: pip6CurCfgVlanPip,
		Params:            params,
	}
}

func (c *Pip6CurCfgVlanTable) Name() string {
	return "Pip6CurCfgVlanTable"
}

func (c *Pip6CurCfgVlanTable) GetParams() BeanType {
	return c.Params
}

func (c *Pip6CurCfgVlanTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *Pip6CurCfgVlanTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.Pip6CurCfgVlanPip).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.Pip6CurCfgVlanPip)
}

type Pip6CurCfgVlanTableParams struct {
	// PIP address by the instance of the pip6CurCfgIndex.
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

func (p Pip6CurCfgVlanTableParams) iMABean() {}
