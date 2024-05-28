package beans

import (
	"fmt"
	"reflect"
)

// SlbOperGroupRealServerTable The table of real server per group.
// Note:This mib is not supported for VX instance of Virtualization.
type SlbOperGroupRealServerTable struct {
	// The real server group index.
	SlbOperRealServGroupIndex int32
	// The real server index.
	SlbOperGroupRealServIndex int32
	Params                    *SlbOperGroupRealServerTableParams
}

func NewSlbOperGroupRealServerTableList() *SlbOperGroupRealServerTable {
	return &SlbOperGroupRealServerTable{}
}

func NewSlbOperGroupRealServerTable(
	slbOperRealServGroupIndex int32,
	slbOperGroupRealServIndex int32,
	params *SlbOperGroupRealServerTableParams,
) *SlbOperGroupRealServerTable {
	return &SlbOperGroupRealServerTable{
		SlbOperRealServGroupIndex: slbOperRealServGroupIndex,
		SlbOperGroupRealServIndex: slbOperGroupRealServIndex,
		Params:                    params,
	}
}

func (c *SlbOperGroupRealServerTable) Name() string {
	return "SlbOperGroupRealServerTable"
}

func (c *SlbOperGroupRealServerTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbOperGroupRealServerTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbOperGroupRealServerTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbOperRealServGroupIndex).IsZero() &&
		reflect.ValueOf(c.SlbOperGroupRealServIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbOperRealServGroupIndex) + "/" + fmt.Sprint(c.SlbOperGroupRealServIndex)
}

type SlbOperGroupRealServerTableState int32

const (
	SlbOperGroupRealServerTableState_Enable      SlbOperGroupRealServerTableState = 1
	SlbOperGroupRealServerTableState_Disable     SlbOperGroupRealServerTableState = 2
	SlbOperGroupRealServerTableState_Unsupported SlbOperGroupRealServerTableState = 2147483647
)

type SlbOperGroupRealServerTableStatus int32

const (
	SlbOperGroupRealServerTableStatus_Enable      SlbOperGroupRealServerTableStatus = 1
	SlbOperGroupRealServerTableStatus_Disable     SlbOperGroupRealServerTableStatus = 2
	SlbOperGroupRealServerTableStatus_Unsupported SlbOperGroupRealServerTableStatus = 2147483647
)

type SlbOperGroupRealServerTableParams struct {
	// The real server group index.
	RealServGroupIndex int32 `json:"RealServGroupIndex,omitempty"`
	// The real server index.
	ServIndex int32 `json:"ServIndex,omitempty"`
	// This an action object which is used to temporarily enable/disable a
	// real server in group. The real server will be returned to its configured
	// operational mode when the switch is reset.
	State SlbOperGroupRealServerTableState `json:"State,omitempty"`
	// This an action object which is used to temporarily enable/disable a
	// real server in group. The real server will be returned to its configured
	// operational mode when the switch is reset.
	Status SlbOperGroupRealServerTableStatus `json:"Status,omitempty"`
}

func (p SlbOperGroupRealServerTableParams) iMABean() {}
