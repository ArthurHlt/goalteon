package beans

import (
	"fmt"
	"reflect"
)

// Pip6NewCfgPortTable The PIP Port table New configuration block.
// Note:This mib is not supported for VX instance of Virtualization.
type Pip6NewCfgPortTable struct {
	// PIP address by the instance of the pip6NewCfgPortIndex.
	Pip6NewCfgPortPip string
	Params            *Pip6NewCfgPortTableParams
}

func NewPip6NewCfgPortTableList() *Pip6NewCfgPortTable {
	return &Pip6NewCfgPortTable{}
}

func NewPip6NewCfgPortTable(
	pip6NewCfgPortPip string,
	params *Pip6NewCfgPortTableParams,
) *Pip6NewCfgPortTable {
	return &Pip6NewCfgPortTable{
		Pip6NewCfgPortPip: pip6NewCfgPortPip,
		Params:            params,
	}
}

func (c *Pip6NewCfgPortTable) Name() string {
	return "Pip6NewCfgPortTable"
}

func (c *Pip6NewCfgPortTable) GetParams() BeanType {
	return c.Params
}

func (c *Pip6NewCfgPortTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *Pip6NewCfgPortTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.Pip6NewCfgPortPip).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.Pip6NewCfgPortPip)
}

type Pip6NewCfgPortTableDelete int32

const (
	Pip6NewCfgPortTableDelete_Other       Pip6NewCfgPortTableDelete = 1
	Pip6NewCfgPortTableDelete_Delete      Pip6NewCfgPortTableDelete = 2
	Pip6NewCfgPortTableDelete_Unsupported Pip6NewCfgPortTableDelete = 2147483647
)

type Pip6NewCfgPortTableParams struct {
	// PIP address by the instance of the pip6NewCfgPortIndex.
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
	Delete Pip6NewCfgPortTableDelete `json:"Delete,omitempty"`
}

func (p Pip6NewCfgPortTableParams) iMABean() {}
