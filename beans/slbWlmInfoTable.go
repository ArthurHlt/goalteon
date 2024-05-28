package beans

import (
	"fmt"
	"reflect"
)

// SlbWlmInfoTable The table of workload manager run-time information.
// Note:This mib is not supported for VX instance of Virtualization.
type SlbWlmInfoTable struct {
	// The workload manager index.
	SlbWlmInfoIndex int32
	Params          *SlbWlmInfoTableParams
}

func NewSlbWlmInfoTableList() *SlbWlmInfoTable {
	return &SlbWlmInfoTable{}
}

func NewSlbWlmInfoTable(
	slbWlmInfoIndex int32,
	params *SlbWlmInfoTableParams,
) *SlbWlmInfoTable {
	return &SlbWlmInfoTable{
		SlbWlmInfoIndex: slbWlmInfoIndex,
		Params:          params,
	}
}

func (c *SlbWlmInfoTable) Name() string {
	return "SlbWlmInfoTable"
}

func (c *SlbWlmInfoTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbWlmInfoTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbWlmInfoTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbWlmInfoIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbWlmInfoIndex)
}

type SlbWlmInfoTableState int32

const (
	SlbWlmInfoTableState_Connected    SlbWlmInfoTableState = 1
	SlbWlmInfoTableState_Notconnected SlbWlmInfoTableState = 2
)

type SlbWlmInfoTableParams struct {
	// The workload manager index.
	Index int32 `json:"Index,omitempty"`
	// IP address of the workload manager.
	IpAddr string `json:"IpAddr,omitempty"`
	// The port number of the workload manager.
	Port uint64 `json:"Port,omitempty"`
	// The state of the workload manager.
	State SlbWlmInfoTableState `json:"State,omitempty"`
}

func (p SlbWlmInfoTableParams) iMABean() {}
