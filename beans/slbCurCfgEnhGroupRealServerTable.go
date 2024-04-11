package beans

import (
	"fmt"
	"reflect"
)

// SlbCurCfgEnhGroupRealServerTable The table of real servers per group.
// Note:This mib is not supported for VX instance of Virtualization.
type SlbCurCfgEnhGroupRealServerTable struct {
	// The real server group index.
	SlbCurCfgEnhRealServGroupIndex string
	// The real server index.
	SlbCurCfgEnhGroupRealServIndex string
	Params                         *SlbCurCfgEnhGroupRealServerTableParams
}

func NewSlbCurCfgEnhGroupRealServerTableList() *SlbCurCfgEnhGroupRealServerTable {
	return &SlbCurCfgEnhGroupRealServerTable{}
}

func NewSlbCurCfgEnhGroupRealServerTable(
	slbCurCfgEnhRealServGroupIndex string,
	slbCurCfgEnhGroupRealServIndex string,
	params *SlbCurCfgEnhGroupRealServerTableParams,
) *SlbCurCfgEnhGroupRealServerTable {
	return &SlbCurCfgEnhGroupRealServerTable{
		SlbCurCfgEnhRealServGroupIndex: slbCurCfgEnhRealServGroupIndex,
		SlbCurCfgEnhGroupRealServIndex: slbCurCfgEnhGroupRealServIndex,
		Params:                         params,
	}
}

func (c *SlbCurCfgEnhGroupRealServerTable) Name() string {
	return "SlbCurCfgEnhGroupRealServerTable"
}

func (c *SlbCurCfgEnhGroupRealServerTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbCurCfgEnhGroupRealServerTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbCurCfgEnhGroupRealServerTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbCurCfgEnhRealServGroupIndex).IsZero() &&
		reflect.ValueOf(c.SlbCurCfgEnhGroupRealServIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbCurCfgEnhRealServGroupIndex) + "/" + fmt.Sprint(c.SlbCurCfgEnhGroupRealServIndex)
}

type SlbCurCfgEnhGroupRealServerTableState int32

const (
	SlbCurCfgEnhGroupRealServerTableState_Enabled                    SlbCurCfgEnhGroupRealServerTableState = 1
	SlbCurCfgEnhGroupRealServerTableState_Disabled                   SlbCurCfgEnhGroupRealServerTableState = 2
	SlbCurCfgEnhGroupRealServerTableState_ShutdownConnection         SlbCurCfgEnhGroupRealServerTableState = 3
	SlbCurCfgEnhGroupRealServerTableState_ShutdownPersistentSessions SlbCurCfgEnhGroupRealServerTableState = 4
	SlbCurCfgEnhGroupRealServerTableState_Unsupported                SlbCurCfgEnhGroupRealServerTableState = 2147483647
)

type SlbCurCfgEnhGroupRealServerTableParams struct {
	// The real server group index.
	RealServGroupIndex string `json:"RealServGroupIndex,omitempty"`
	// The real server index.
	ServIndex string `json:"ServIndex,omitempty"`
	// Enable/disable a real server gracefully on a per group basis.
	State SlbCurCfgEnhGroupRealServerTableState `json:"State,omitempty"`
}

func (p SlbCurCfgEnhGroupRealServerTableParams) iMABean() {}
