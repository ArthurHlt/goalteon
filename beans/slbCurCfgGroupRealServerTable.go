package beans

import (
	"fmt"
	"reflect"
)

// SlbCurCfgGroupRealServerTable The table of real servers per group.
// Note:This mib is not supported on VX instance of Virtualization.
type SlbCurCfgGroupRealServerTable struct {
	// The real server group index.
	SlbCurCfgRealServGroupIndex int32
	// The real server index.
	SlbCurCfgGroupRealServIndex int32
	Params                      *SlbCurCfgGroupRealServerTableParams
}

func NewSlbCurCfgGroupRealServerTableList() *SlbCurCfgGroupRealServerTable {
	return &SlbCurCfgGroupRealServerTable{}
}

func NewSlbCurCfgGroupRealServerTable(
	slbCurCfgRealServGroupIndex int32,
	slbCurCfgGroupRealServIndex int32,
	params *SlbCurCfgGroupRealServerTableParams,
) *SlbCurCfgGroupRealServerTable {
	return &SlbCurCfgGroupRealServerTable{
		SlbCurCfgRealServGroupIndex: slbCurCfgRealServGroupIndex,
		SlbCurCfgGroupRealServIndex: slbCurCfgGroupRealServIndex,
		Params:                      params,
	}
}

func (c *SlbCurCfgGroupRealServerTable) Name() string {
	return "SlbCurCfgGroupRealServerTable"
}

func (c *SlbCurCfgGroupRealServerTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbCurCfgGroupRealServerTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbCurCfgGroupRealServerTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbCurCfgRealServGroupIndex).IsZero() &&
		reflect.ValueOf(c.SlbCurCfgGroupRealServIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbCurCfgRealServGroupIndex) + "/" + fmt.Sprint(c.SlbCurCfgGroupRealServIndex)
}

type SlbCurCfgGroupRealServerTableState int32

const (
	SlbCurCfgGroupRealServerTableState_Enabled     SlbCurCfgGroupRealServerTableState = 1
	SlbCurCfgGroupRealServerTableState_Disabled    SlbCurCfgGroupRealServerTableState = 2
	SlbCurCfgGroupRealServerTableState_Unsupported SlbCurCfgGroupRealServerTableState = 2147483647
)

type SlbCurCfgGroupRealServerTableParams struct {
	// The real server group index.
	RealServGroupIndex int32 `json:"RealServGroupIndex,omitempty"`
	// The real server index.
	ServIndex int32 `json:"ServIndex,omitempty"`
	// Enable/disable a real server gracefully on a per group basis.
	State SlbCurCfgGroupRealServerTableState `json:"State,omitempty"`
}

func (p SlbCurCfgGroupRealServerTableParams) iMABean() {}
