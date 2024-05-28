package beans

import (
	"fmt"
	"reflect"
)

// SlbOperEnhGroupRealServerTable The table of real server per group.
// Note:This mib is not supported for VX instance of Virtualization.
type SlbOperEnhGroupRealServerTable struct {
	// The real server group index.
	SlbOperEnhRealServGroupIndex string
	// The real server index.
	SlbOperEnhGroupRealServIndex string
	Params                       *SlbOperEnhGroupRealServerTableParams
}

func NewSlbOperEnhGroupRealServerTableList() *SlbOperEnhGroupRealServerTable {
	return &SlbOperEnhGroupRealServerTable{}
}

func NewSlbOperEnhGroupRealServerTable(
	slbOperEnhRealServGroupIndex string,
	slbOperEnhGroupRealServIndex string,
	params *SlbOperEnhGroupRealServerTableParams,
) *SlbOperEnhGroupRealServerTable {
	return &SlbOperEnhGroupRealServerTable{
		SlbOperEnhRealServGroupIndex: slbOperEnhRealServGroupIndex,
		SlbOperEnhGroupRealServIndex: slbOperEnhGroupRealServIndex,
		Params:                       params,
	}
}

func (c *SlbOperEnhGroupRealServerTable) Name() string {
	return "SlbOperEnhGroupRealServerTable"
}

func (c *SlbOperEnhGroupRealServerTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbOperEnhGroupRealServerTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbOperEnhGroupRealServerTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbOperEnhRealServGroupIndex).IsZero() &&
		reflect.ValueOf(c.SlbOperEnhGroupRealServIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbOperEnhRealServGroupIndex) + "/" + fmt.Sprint(c.SlbOperEnhGroupRealServIndex)
}

type SlbOperEnhGroupRealServerTableState int32

const (
	SlbOperEnhGroupRealServerTableState_Enable                     SlbOperEnhGroupRealServerTableState = 1
	SlbOperEnhGroupRealServerTableState_Disable                    SlbOperEnhGroupRealServerTableState = 2
	SlbOperEnhGroupRealServerTableState_ShutdownConnection         SlbOperEnhGroupRealServerTableState = 3
	SlbOperEnhGroupRealServerTableState_ShutdownPersistentSessions SlbOperEnhGroupRealServerTableState = 4
	SlbOperEnhGroupRealServerTableState_Unsupported                SlbOperEnhGroupRealServerTableState = 2147483647
)

type SlbOperEnhGroupRealServerTableStatus int32

const (
	SlbOperEnhGroupRealServerTableStatus_Enable                     SlbOperEnhGroupRealServerTableStatus = 1
	SlbOperEnhGroupRealServerTableStatus_Disable                    SlbOperEnhGroupRealServerTableStatus = 2
	SlbOperEnhGroupRealServerTableStatus_ShutdownConnection         SlbOperEnhGroupRealServerTableStatus = 3
	SlbOperEnhGroupRealServerTableStatus_ShutdownPersistentSessions SlbOperEnhGroupRealServerTableStatus = 4
	SlbOperEnhGroupRealServerTableStatus_Unsupported                SlbOperEnhGroupRealServerTableStatus = 2147483647
)

type SlbOperEnhGroupRealServerTableRuntimeStatus int32

const (
	SlbOperEnhGroupRealServerTableRuntimeStatus_Running     SlbOperEnhGroupRealServerTableRuntimeStatus = 1
	SlbOperEnhGroupRealServerTableRuntimeStatus_Failed      SlbOperEnhGroupRealServerTableRuntimeStatus = 2
	SlbOperEnhGroupRealServerTableRuntimeStatus_Disabled    SlbOperEnhGroupRealServerTableRuntimeStatus = 3
	SlbOperEnhGroupRealServerTableRuntimeStatus_Overloaded  SlbOperEnhGroupRealServerTableRuntimeStatus = 4
	SlbOperEnhGroupRealServerTableRuntimeStatus_Unsupported SlbOperEnhGroupRealServerTableRuntimeStatus = 2147483647
)

type SlbOperEnhGroupRealServerTableParams struct {
	// The real server group index.
	RealServGroupIndex string `json:"RealServGroupIndex,omitempty"`
	// The real server index.
	ServIndex string `json:"ServIndex,omitempty"`
	// This an action object which is used to temporarily enable/disable a
	// real server in group. The real server will be returned to its configured
	// operational mode when the switch is reset.
	State SlbOperEnhGroupRealServerTableState `json:"State,omitempty"`
	// This an action object which is used to temporarily enable/disable a
	// real server in group. The real server will be returned to its configured
	// operational mode when the switch is reset.
	Status SlbOperEnhGroupRealServerTableStatus `json:"Status,omitempty"`
	// IP address of the real server identified by the instance of the
	// slbRealServerIndex.
	IP string `json:"IP,omitempty"`
	// Description of the real server identified by the instance of the
	// slbRealServerIndex.
	Descr string `json:"Descr,omitempty"`
	// The current status of a real server in the group.
	RuntimeStatus SlbOperEnhGroupRealServerTableRuntimeStatus `json:"RuntimeStatus,omitempty"`
}

func (p SlbOperEnhGroupRealServerTableParams) iMABean() {}
