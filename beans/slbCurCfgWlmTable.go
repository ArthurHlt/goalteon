package beans

import (
	"fmt"
	"reflect"
)

// SlbCurCfgWlmTable The table of Workload Managers configuration.
// Note:This mib is not supported for VX instance of Virtualization.
type SlbCurCfgWlmTable struct {
	// The Workload Manager number
	SlbCurCfgWlmIndex int32
	Params            *SlbCurCfgWlmTableParams
}

func NewSlbCurCfgWlmTableList() *SlbCurCfgWlmTable {
	return &SlbCurCfgWlmTable{}
}

func NewSlbCurCfgWlmTable(
	slbCurCfgWlmIndex int32,
	params *SlbCurCfgWlmTableParams,
) *SlbCurCfgWlmTable {
	return &SlbCurCfgWlmTable{
		SlbCurCfgWlmIndex: slbCurCfgWlmIndex,
		Params:            params,
	}
}

func (c *SlbCurCfgWlmTable) Name() string {
	return "SlbCurCfgWlmTable"
}

func (c *SlbCurCfgWlmTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbCurCfgWlmTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbCurCfgWlmTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbCurCfgWlmIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbCurCfgWlmIndex)
}

type SlbCurCfgWlmTableState int32

const (
	SlbCurCfgWlmTableState_Enabled     SlbCurCfgWlmTableState = 1
	SlbCurCfgWlmTableState_Disabled    SlbCurCfgWlmTableState = 2
	SlbCurCfgWlmTableState_Unsupported SlbCurCfgWlmTableState = 2147483647
)

type SlbCurCfgWlmTableParams struct {
	// The Workload Manager number
	Index int32 `json:"Index,omitempty"`
	// IP address of the Workload Manager.
	IpAddr string `json:"IpAddr,omitempty"`
	// The port number for the Workload Manager.
	Port uint64 `json:"Port,omitempty"`
	// The group number for the Workload Manager.
	Groups string `json:"Groups,omitempty"`
	// Connection state of the Workload Manager.
	State SlbCurCfgWlmTableState `json:"State,omitempty"`
}

func (p SlbCurCfgWlmTableParams) iMABean() {}
