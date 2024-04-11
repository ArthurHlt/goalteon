package beans

import (
	"fmt"
	"reflect"
)

// SlbNewCfgEnhBuddyTable The table of buddy servers.
// Note:This mib is not supported on VX instance of Virtualization.
type SlbNewCfgEnhBuddyTable struct {
	// The number of the real server.
	SlbNewCfgEnhRealSerIndex string
	// The buddy index. This has no external meaning
	SlbNewCfgEnhBuddyIndex int32
	Params                 *SlbNewCfgEnhBuddyTableParams
}

func NewSlbNewCfgEnhBuddyTableList() *SlbNewCfgEnhBuddyTable {
	return &SlbNewCfgEnhBuddyTable{}
}

func NewSlbNewCfgEnhBuddyTable(
	slbNewCfgEnhRealSerIndex string,
	slbNewCfgEnhBuddyIndex int32,
	params *SlbNewCfgEnhBuddyTableParams,
) *SlbNewCfgEnhBuddyTable {
	return &SlbNewCfgEnhBuddyTable{
		SlbNewCfgEnhRealSerIndex: slbNewCfgEnhRealSerIndex,
		SlbNewCfgEnhBuddyIndex:   slbNewCfgEnhBuddyIndex,
		Params:                   params,
	}
}

func (c *SlbNewCfgEnhBuddyTable) Name() string {
	return "SlbNewCfgEnhBuddyTable"
}

func (c *SlbNewCfgEnhBuddyTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbNewCfgEnhBuddyTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbNewCfgEnhBuddyTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbNewCfgEnhRealSerIndex).IsZero() &&
		reflect.ValueOf(c.SlbNewCfgEnhBuddyIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbNewCfgEnhRealSerIndex) + "/" + fmt.Sprint(c.SlbNewCfgEnhBuddyIndex)
}

type SlbNewCfgEnhBuddyTableDelete int32

const (
	SlbNewCfgEnhBuddyTableDelete_Other       SlbNewCfgEnhBuddyTableDelete = 1
	SlbNewCfgEnhBuddyTableDelete_Delete      SlbNewCfgEnhBuddyTableDelete = 2
	SlbNewCfgEnhBuddyTableDelete_Unsupported SlbNewCfgEnhBuddyTableDelete = 2147483647
)

type SlbNewCfgEnhBuddyTableParams struct {
	// The number of the real server.
	RealSerIndex string `json:"RealSerIndex,omitempty"`
	// The buddy index. This has no external meaning
	Index int32 `json:"Index,omitempty"`
	// The buddy server's real server number.
	RealIndex string `json:"RealIndex,omitempty"`
	// The buddy server's real group number.
	GroupIndex string `json:"GroupIndex,omitempty"`
	// The buddy server's service.
	Service uint64 `json:"Service,omitempty"`
	// When set to the value of 2 (delete), the entire row is deleted.
	// When read, other(1) is returned. Setting the value to anything
	// other than 2(delete) has no effect on the state of the row.
	Delete SlbNewCfgEnhBuddyTableDelete `json:"Delete,omitempty"`
}

func (p SlbNewCfgEnhBuddyTableParams) iMABean() {}
