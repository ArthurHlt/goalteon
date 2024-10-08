package beans

import (
	"fmt"
	"reflect"
)

// SlbNewAdvhcLogexpTable Note:This mib is not supported on VX instance of Virtualization.
type SlbNewAdvhcLogexpTable struct {
	// Logical expression health check id.
	SlbNewAdvhcLogexpID string
	Params              *SlbNewAdvhcLogexpTableParams
}

func NewSlbNewAdvhcLogexpTableList() *SlbNewAdvhcLogexpTable {
	return &SlbNewAdvhcLogexpTable{}
}

func NewSlbNewAdvhcLogexpTable(
	slbNewAdvhcLogexpID string,
	params *SlbNewAdvhcLogexpTableParams,
) *SlbNewAdvhcLogexpTable {
	return &SlbNewAdvhcLogexpTable{
		SlbNewAdvhcLogexpID: slbNewAdvhcLogexpID,
		Params:              params,
	}
}

func (c *SlbNewAdvhcLogexpTable) Name() string {
	return "SlbNewAdvhcLogexpTable"
}

func (c *SlbNewAdvhcLogexpTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbNewAdvhcLogexpTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbNewAdvhcLogexpTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbNewAdvhcLogexpID).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbNewAdvhcLogexpID)
}

type SlbNewAdvhcLogexpTableDelete int32

const (
	SlbNewAdvhcLogexpTableDelete_Other       SlbNewAdvhcLogexpTableDelete = 1
	SlbNewAdvhcLogexpTableDelete_Delete      SlbNewAdvhcLogexpTableDelete = 2
	SlbNewAdvhcLogexpTableDelete_Unsupported SlbNewAdvhcLogexpTableDelete = 2147483647
)

type SlbNewAdvhcLogexpTableAlways int32

const (
	SlbNewAdvhcLogexpTableAlways_Enabled     SlbNewAdvhcLogexpTableAlways = 1
	SlbNewAdvhcLogexpTableAlways_Disabled    SlbNewAdvhcLogexpTableAlways = 2
	SlbNewAdvhcLogexpTableAlways_Unsupported SlbNewAdvhcLogexpTableAlways = 2147483647
)

type SlbNewAdvhcLogexpTableParams struct {
	// Logical expression health check id.
	ID string `json:"ID,omitempty"`
	// Logical expression health check name.
	Name string `json:"Name,omitempty"`
	// Logical expression health check text.
	Text string `json:"Text,omitempty"`
	// Logical expression health check copy indicator.
	Copy DisplayString `json:"Copy,omitempty"`
	// When set to the value of 2 (delete), the entire row is deleted.
	// When read, other(1) is returned. Setting the value to anything
	// other than 2(delete) has no effect on the state of the row.
	Delete SlbNewAdvhcLogexpTableDelete `json:"Delete,omitempty"`
	// This flag determines whether HC is allowed for standalone real.
	Always SlbNewAdvhcLogexpTableAlways `json:"Always,omitempty"`
}

func (p SlbNewAdvhcLogexpTableParams) iMABean() {}
