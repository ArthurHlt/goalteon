package beans

import (
	"fmt"
	"reflect"
)

// SlbNewCfgGroupRealServerTable The table of real server per group.
// Note:This mib is not supported on VX instance of Virtualization.
type SlbNewCfgGroupRealServerTable struct {
	// The real server group index.
	SlbNewCfgRealServGroupIndex int32
	// The real server index.
	SlbNewCfgGroupRealServIndex int32
	Params                      *SlbNewCfgGroupRealServerTableParams
}

func NewSlbNewCfgGroupRealServerTableList() *SlbNewCfgGroupRealServerTable {
	return &SlbNewCfgGroupRealServerTable{}
}

func NewSlbNewCfgGroupRealServerTable(
	slbNewCfgRealServGroupIndex int32,
	slbNewCfgGroupRealServIndex int32,
	params *SlbNewCfgGroupRealServerTableParams,
) *SlbNewCfgGroupRealServerTable {
	return &SlbNewCfgGroupRealServerTable{
		SlbNewCfgRealServGroupIndex: slbNewCfgRealServGroupIndex,
		SlbNewCfgGroupRealServIndex: slbNewCfgGroupRealServIndex,
		Params:                      params,
	}
}

func (c *SlbNewCfgGroupRealServerTable) Name() string {
	return "SlbNewCfgGroupRealServerTable"
}

func (c *SlbNewCfgGroupRealServerTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbNewCfgGroupRealServerTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbNewCfgGroupRealServerTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbNewCfgRealServGroupIndex).IsZero() &&
		reflect.ValueOf(c.SlbNewCfgGroupRealServIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbNewCfgRealServGroupIndex) + "/" + fmt.Sprint(c.SlbNewCfgGroupRealServIndex)
}

type SlbNewCfgGroupRealServerTableState int32

const (
	SlbNewCfgGroupRealServerTableState_Enabled     SlbNewCfgGroupRealServerTableState = 1
	SlbNewCfgGroupRealServerTableState_Disabled    SlbNewCfgGroupRealServerTableState = 2
	SlbNewCfgGroupRealServerTableState_Unsupported SlbNewCfgGroupRealServerTableState = 2147483647
)

type SlbNewCfgGroupRealServerTableParams struct {
	// The real server group index.
	RealServGroupIndex int32 `json:"RealServGroupIndex,omitempty"`
	// The real server index.
	ServIndex int32 `json:"ServIndex,omitempty"`
	// Enable/disable a real server gracefully on a per group basis.
	State SlbNewCfgGroupRealServerTableState `json:"State,omitempty"`
}

func (p SlbNewCfgGroupRealServerTableParams) iMABean() {}
