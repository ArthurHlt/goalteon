package beans

import (
	"fmt"
	"reflect"
)

// SlbNewCfgWlmTable The table of Workload Managers configuration.
// Note:This mib is not supported for VX instance of Virtualization.
type SlbNewCfgWlmTable struct {
	// The Workload Manager number
	SlbNewCfgWlmIndex int32
	Params            *SlbNewCfgWlmTableParams
}

func NewSlbNewCfgWlmTableList() *SlbNewCfgWlmTable {
	return &SlbNewCfgWlmTable{}
}

func NewSlbNewCfgWlmTable(
	slbNewCfgWlmIndex int32,
	params *SlbNewCfgWlmTableParams,
) *SlbNewCfgWlmTable {
	return &SlbNewCfgWlmTable{
		SlbNewCfgWlmIndex: slbNewCfgWlmIndex,
		Params:            params,
	}
}

func (c *SlbNewCfgWlmTable) Name() string {
	return "SlbNewCfgWlmTable"
}

func (c *SlbNewCfgWlmTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbNewCfgWlmTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbNewCfgWlmTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbNewCfgWlmIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbNewCfgWlmIndex)
}

type SlbNewCfgWlmTableDelete int32

const (
	SlbNewCfgWlmTableDelete_Other       SlbNewCfgWlmTableDelete = 1
	SlbNewCfgWlmTableDelete_Delete      SlbNewCfgWlmTableDelete = 2
	SlbNewCfgWlmTableDelete_Unsupported SlbNewCfgWlmTableDelete = 2147483647
)

type SlbNewCfgWlmTableState int32

const (
	SlbNewCfgWlmTableState_Enabled     SlbNewCfgWlmTableState = 1
	SlbNewCfgWlmTableState_Disabled    SlbNewCfgWlmTableState = 2
	SlbNewCfgWlmTableState_Unsupported SlbNewCfgWlmTableState = 2147483647
)

type SlbNewCfgWlmTableParams struct {
	// The Workload Manager number
	Index int32 `json:"Index,omitempty"`
	// IP address of the Workload Manager.
	IpAddr string `json:"IpAddr,omitempty"`
	// The port number for the Workload Manager.
	Port uint64 `json:"Port,omitempty"`
	// By setting the value to delete(2), the entire row is deleted.
	Delete SlbNewCfgWlmTableDelete `json:"Delete,omitempty"`
	// The group number for the Workload Manager.
	Groups string `json:"Groups,omitempty"`
	// Connection state of the Workload Manager.
	State SlbNewCfgWlmTableState `json:"State,omitempty"`
}

func (p SlbNewCfgWlmTableParams) iMABean() {}
