package beans

import (
	"fmt"
	"reflect"
)

// Pip6NewCfgVlanTable The PIP Vlan table New configuration block.
// Note:This mib is not supported for VX instance of Virtualization.
type Pip6NewCfgVlanTable struct {
	// PIP address by the instance of the pip6CurCfgIndex.
	Pip6NewCfgVlanPip string
	Params            *Pip6NewCfgVlanTableParams
}

func NewPip6NewCfgVlanTableList() *Pip6NewCfgVlanTable {
	return &Pip6NewCfgVlanTable{}
}

func NewPip6NewCfgVlanTable(
	pip6NewCfgVlanPip string,
	params *Pip6NewCfgVlanTableParams,
) *Pip6NewCfgVlanTable {
	return &Pip6NewCfgVlanTable{
		Pip6NewCfgVlanPip: pip6NewCfgVlanPip,
		Params:            params,
	}
}

func (c *Pip6NewCfgVlanTable) Name() string {
	return "Pip6NewCfgVlanTable"
}

func (c *Pip6NewCfgVlanTable) GetParams() BeanType {
	return c.Params
}

func (c *Pip6NewCfgVlanTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *Pip6NewCfgVlanTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.Pip6NewCfgVlanPip).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.Pip6NewCfgVlanPip)
}

type Pip6NewCfgVlanTableDelete int32

const (
	Pip6NewCfgVlanTableDelete_Other       Pip6NewCfgVlanTableDelete = 1
	Pip6NewCfgVlanTableDelete_Delete      Pip6NewCfgVlanTableDelete = 2
	Pip6NewCfgVlanTableDelete_Unsupported Pip6NewCfgVlanTableDelete = 2147483647
)

type Pip6NewCfgVlanTableParams struct {
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
	// The VLAN to be associated with the PIP. When read, 0 is returned.
	Add int32 `json:"Add,omitempty"`
	// The VLAN to be disassociated from the PIP. When read,
	// 0 is returned.
	Remove int32 `json:"Remove,omitempty"`
	// By setting the value to delete(2), the entire row is deleted.
	Delete Pip6NewCfgVlanTableDelete `json:"Delete,omitempty"`
}

func (p Pip6NewCfgVlanTableParams) iMABean() {}
