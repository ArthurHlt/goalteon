package beans

import (
	"fmt"
	"reflect"
)

// PipCurCfgPortTable The PIP Port table Current configuration block.
// Note:This mib is not supported for VX instance of Virtualization.
type PipCurCfgPortTable struct {
	// PIP address by the instance of the pipCurCfgPortIndex.
	PipCurCfgPortPip string
	Params           *PipCurCfgPortTableParams
}

func NewPipCurCfgPortTableList() *PipCurCfgPortTable {
	return &PipCurCfgPortTable{}
}

func NewPipCurCfgPortTable(
	pipCurCfgPortPip string,
	params *PipCurCfgPortTableParams,
) *PipCurCfgPortTable {
	return &PipCurCfgPortTable{
		PipCurCfgPortPip: pipCurCfgPortPip,
		Params:           params,
	}
}

func (c *PipCurCfgPortTable) Name() string {
	return "PipCurCfgPortTable"
}

func (c *PipCurCfgPortTable) GetParams() BeanType {
	return c.Params
}

func (c *PipCurCfgPortTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *PipCurCfgPortTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.PipCurCfgPortPip).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.PipCurCfgPortPip)
}

type PipCurCfgPortTableParams struct {
	// PIP address by the instance of the pipCurCfgPortIndex.
	Pip string `json:"Pip,omitempty"`
	// The bit map of port for PIP.in receiving order:
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
}

func (p PipCurCfgPortTableParams) iMABean() {}
