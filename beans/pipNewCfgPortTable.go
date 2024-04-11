package beans

import (
	"fmt"
	"reflect"
)

// PipNewCfgPortTable The PIP Port table New configuration block.
// Note:This mib is not supported for VX instance of Virtualization.
type PipNewCfgPortTable struct {
	// PIP address by the instance of the pipNewCfgPortIndex.
	PipNewCfgPortPip string
	Params           *PipNewCfgPortTableParams
}

func NewPipNewCfgPortTableList() *PipNewCfgPortTable {
	return &PipNewCfgPortTable{}
}

func NewPipNewCfgPortTable(
	pipNewCfgPortPip string,
	params *PipNewCfgPortTableParams,
) *PipNewCfgPortTable {
	return &PipNewCfgPortTable{
		PipNewCfgPortPip: pipNewCfgPortPip,
		Params:           params,
	}
}

func (c *PipNewCfgPortTable) Name() string {
	return "PipNewCfgPortTable"
}

func (c *PipNewCfgPortTable) GetParams() BeanType {
	return c.Params
}

func (c *PipNewCfgPortTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *PipNewCfgPortTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.PipNewCfgPortPip).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.PipNewCfgPortPip)
}

type PipNewCfgPortTableDelete int32

const (
	PipNewCfgPortTableDelete_Other       PipNewCfgPortTableDelete = 1
	PipNewCfgPortTableDelete_Delete      PipNewCfgPortTableDelete = 2
	PipNewCfgPortTableDelete_Unsupported PipNewCfgPortTableDelete = 2147483647
)

type PipNewCfgPortTableParams struct {
	// PIP address by the instance of the pipNewCfgPortIndex.
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
	PipMap string `json:"PipMap,omitempty"`
	// The Port to be associated with the PIP. When read, 0 is returned.
	Add int32 `json:"Add,omitempty"`
	// The Port to be disassociated from the PIP. When read,
	// 0 is returned.
	Remove int32 `json:"Remove,omitempty"`
	// By setting the value to delete(2), the entire row is deleted.
	Delete PipNewCfgPortTableDelete `json:"Delete,omitempty"`
}

func (p PipNewCfgPortTableParams) iMABean() {}
