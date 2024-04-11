package beans

import (
	"fmt"
	"reflect"
)

// SlbCurAdvhcLogexpTable Note:This mib is not supported on VX instance of Virtualization.
type SlbCurAdvhcLogexpTable struct {
	// Logical expression health check id.
	SlbCurAdvhcLogexpID string
	Params              *SlbCurAdvhcLogexpTableParams
}

func NewSlbCurAdvhcLogexpTableList() *SlbCurAdvhcLogexpTable {
	return &SlbCurAdvhcLogexpTable{}
}

func NewSlbCurAdvhcLogexpTable(
	slbCurAdvhcLogexpID string,
	params *SlbCurAdvhcLogexpTableParams,
) *SlbCurAdvhcLogexpTable {
	return &SlbCurAdvhcLogexpTable{
		SlbCurAdvhcLogexpID: slbCurAdvhcLogexpID,
		Params:              params,
	}
}

func (c *SlbCurAdvhcLogexpTable) Name() string {
	return "SlbCurAdvhcLogexpTable"
}

func (c *SlbCurAdvhcLogexpTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbCurAdvhcLogexpTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbCurAdvhcLogexpTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbCurAdvhcLogexpID).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbCurAdvhcLogexpID)
}

type SlbCurAdvhcLogexpTableAlways int32

const (
	SlbCurAdvhcLogexpTableAlways_Enabled     SlbCurAdvhcLogexpTableAlways = 1
	SlbCurAdvhcLogexpTableAlways_Disabled    SlbCurAdvhcLogexpTableAlways = 2
	SlbCurAdvhcLogexpTableAlways_Unsupported SlbCurAdvhcLogexpTableAlways = 2147483647
)

type SlbCurAdvhcLogexpTableParams struct {
	// Logical expression health check id.
	ID string `json:"ID,omitempty"`
	// Logical expression health check name.
	Name string `json:"Name,omitempty"`
	// Logical expression health check text.
	Text string `json:"Text,omitempty"`
	// This flag determines whether HC is allowed for standalone real.
	Always SlbCurAdvhcLogexpTableAlways `json:"Always,omitempty"`
}

func (p SlbCurAdvhcLogexpTableParams) iMABean() {}
