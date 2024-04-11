package beans

import (
	"fmt"
	"reflect"
)

// SlbNewCfgEnhGroupRealServerTable The table of real server per group.
// Note:This mib is not supported for VX instance of Virtualization.
type SlbNewCfgEnhGroupRealServerTable struct {
	// The real server group index.
	SlbNewCfgEnhRealServGroupIndex string
	// The real server index.
	SlbNewCfgEnhGroupRealServIndex string
	Params                         *SlbNewCfgEnhGroupRealServerTableParams
}

func NewSlbNewCfgEnhGroupRealServerTableList() *SlbNewCfgEnhGroupRealServerTable {
	return &SlbNewCfgEnhGroupRealServerTable{}
}

func NewSlbNewCfgEnhGroupRealServerTable(
	slbNewCfgEnhRealServGroupIndex string,
	slbNewCfgEnhGroupRealServIndex string,
	params *SlbNewCfgEnhGroupRealServerTableParams,
) *SlbNewCfgEnhGroupRealServerTable {
	return &SlbNewCfgEnhGroupRealServerTable{
		SlbNewCfgEnhRealServGroupIndex: slbNewCfgEnhRealServGroupIndex,
		SlbNewCfgEnhGroupRealServIndex: slbNewCfgEnhGroupRealServIndex,
		Params:                         params,
	}
}

func (c *SlbNewCfgEnhGroupRealServerTable) Name() string {
	return "SlbNewCfgEnhGroupRealServerTable"
}

func (c *SlbNewCfgEnhGroupRealServerTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbNewCfgEnhGroupRealServerTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbNewCfgEnhGroupRealServerTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbNewCfgEnhRealServGroupIndex).IsZero() &&
		reflect.ValueOf(c.SlbNewCfgEnhGroupRealServIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbNewCfgEnhRealServGroupIndex) + "/" + fmt.Sprint(c.SlbNewCfgEnhGroupRealServIndex)
}

type SlbNewCfgEnhGroupRealServerTableState int32

const (
	SlbNewCfgEnhGroupRealServerTableState_Enabled                    SlbNewCfgEnhGroupRealServerTableState = 1
	SlbNewCfgEnhGroupRealServerTableState_Disabled                   SlbNewCfgEnhGroupRealServerTableState = 2
	SlbNewCfgEnhGroupRealServerTableState_ShutdownConnection         SlbNewCfgEnhGroupRealServerTableState = 3
	SlbNewCfgEnhGroupRealServerTableState_ShutdownPersistentSessions SlbNewCfgEnhGroupRealServerTableState = 4
	SlbNewCfgEnhGroupRealServerTableState_Unsupported                SlbNewCfgEnhGroupRealServerTableState = 2147483647
)

type SlbNewCfgEnhGroupRealServerTableParams struct {
	// The real server group index.
	RealServGroupIndex string `json:"RealServGroupIndex,omitempty"`
	// The real server index.
	ServIndex string `json:"ServIndex,omitempty"`
	// Enable/disable a real server gracefully on a per group basis.
	State SlbNewCfgEnhGroupRealServerTableState `json:"State,omitempty"`
}

func (p SlbNewCfgEnhGroupRealServerTableParams) iMABean() {}
