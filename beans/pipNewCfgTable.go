package beans

import (
	"fmt"
	"reflect"
)

// PipNewCfgTable The PIP table in the New configuration block.
// Note:This mib is not supported for VX instance of Virtualization.
type PipNewCfgTable struct {
	// PIP address by the instance of the pipCurCfgIndex.
	PipNewCfgPip string
	Params       *PipNewCfgTableParams
}

func NewPipNewCfgTableList() *PipNewCfgTable {
	return &PipNewCfgTable{}
}

func NewPipNewCfgTable(
	pipNewCfgPip string,
	params *PipNewCfgTableParams,
) *PipNewCfgTable {
	return &PipNewCfgTable{
		PipNewCfgPip: pipNewCfgPip,
		Params:       params,
	}
}

func (c *PipNewCfgTable) Name() string {
	return "PipNewCfgTable"
}

func (c *PipNewCfgTable) GetParams() BeanType {
	return c.Params
}

func (c *PipNewCfgTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *PipNewCfgTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.PipNewCfgPip).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.PipNewCfgPip)
}

type PipNewCfgTableDelete int32

const (
	PipNewCfgTableDelete_Other       PipNewCfgTableDelete = 1
	PipNewCfgTableDelete_Delete      PipNewCfgTableDelete = 2
	PipNewCfgTableDelete_Unsupported PipNewCfgTableDelete = 2147483647
)

type PipNewCfgTableParams struct {
	// PIP address by the instance of the pipCurCfgIndex.
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
	Delete PipNewCfgTableDelete `json:"Delete,omitempty"`
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

func (p PipNewCfgTableParams) iMABean() {}
