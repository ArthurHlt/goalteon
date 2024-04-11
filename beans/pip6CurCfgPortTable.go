package beans

import (
	"fmt"
	"reflect"
)

// Pip6CurCfgPortTable The PIP Port table Current configuration block.
// Note:This mib is not supported for VX instance of Virtualization.
type Pip6CurCfgPortTable struct {
	// PIP address by the instance of the pip6CurCfgPortIndex.
	Pip6CurCfgPortPip string
	Params            *Pip6CurCfgPortTableParams
}

func NewPip6CurCfgPortTableList() *Pip6CurCfgPortTable {
	return &Pip6CurCfgPortTable{}
}

func NewPip6CurCfgPortTable(
	pip6CurCfgPortPip string,
	params *Pip6CurCfgPortTableParams,
) *Pip6CurCfgPortTable {
	return &Pip6CurCfgPortTable{
		Pip6CurCfgPortPip: pip6CurCfgPortPip,
		Params:            params,
	}
}

func (c *Pip6CurCfgPortTable) Name() string {
	return "Pip6CurCfgPortTable"
}

func (c *Pip6CurCfgPortTable) GetParams() BeanType {
	return c.Params
}

func (c *Pip6CurCfgPortTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *Pip6CurCfgPortTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.Pip6CurCfgPortPip).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.Pip6CurCfgPortPip)
}

type Pip6CurCfgPortTableParams struct {
	// PIP address by the instance of the pip6CurCfgPortIndex.
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
}

func (p Pip6CurCfgPortTableParams) iMABean() {}
