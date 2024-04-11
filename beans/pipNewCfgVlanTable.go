package beans

import (
	"fmt"
	"reflect"
)

// PipNewCfgVlanTable The PIP Vlan table New configuration block.
// Note:This mib is not supported for VX instance of Virtualization.
type PipNewCfgVlanTable struct {
	// PIP address by the instance of the pipCurCfgIndex.
	PipNewCfgVlanPip string
	Params           *PipNewCfgVlanTableParams
}

func NewPipNewCfgVlanTableList() *PipNewCfgVlanTable {
	return &PipNewCfgVlanTable{}
}

func NewPipNewCfgVlanTable(
	pipNewCfgVlanPip string,
	params *PipNewCfgVlanTableParams,
) *PipNewCfgVlanTable {
	return &PipNewCfgVlanTable{
		PipNewCfgVlanPip: pipNewCfgVlanPip,
		Params:           params,
	}
}

func (c *PipNewCfgVlanTable) Name() string {
	return "PipNewCfgVlanTable"
}

func (c *PipNewCfgVlanTable) GetParams() BeanType {
	return c.Params
}

func (c *PipNewCfgVlanTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *PipNewCfgVlanTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.PipNewCfgVlanPip).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.PipNewCfgVlanPip)
}

type PipNewCfgVlanTableDelete int32

const (
	PipNewCfgVlanTableDelete_Other       PipNewCfgVlanTableDelete = 1
	PipNewCfgVlanTableDelete_Delete      PipNewCfgVlanTableDelete = 2
	PipNewCfgVlanTableDelete_Unsupported PipNewCfgVlanTableDelete = 2147483647
)

type PipNewCfgVlanTableParams struct {
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
	// The VLAN to be associated with the PIP. When read, 0 is returned.
	Add int32 `json:"Add,omitempty"`
	// The VLAN to be disassociated from the PIP. When read,
	// 0 is returned.
	Remove int32 `json:"Remove,omitempty"`
	// By setting the value to delete(2), the entire row is deleted.
	Delete PipNewCfgVlanTableDelete `json:"Delete,omitempty"`
}

func (p PipNewCfgVlanTableParams) iMABean() {}
