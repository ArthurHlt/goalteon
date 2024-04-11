package beans

import (
	"fmt"
	"reflect"
)

// SlbNewCfgBuddyTable The table of buddy servers.
// Note:This mib is not supported on VX instance of Virtualization.
type SlbNewCfgBuddyTable struct {
	// The number of the real server.
	SlbNewCfgRealSerIndex int32
	// The buddy index. This has no external meaning
	SlbNewCfgBuddyIndex int32
	Params              *SlbNewCfgBuddyTableParams
}

func NewSlbNewCfgBuddyTableList() *SlbNewCfgBuddyTable {
	return &SlbNewCfgBuddyTable{}
}

func NewSlbNewCfgBuddyTable(
	slbNewCfgRealSerIndex int32,
	slbNewCfgBuddyIndex int32,
	params *SlbNewCfgBuddyTableParams,
) *SlbNewCfgBuddyTable {
	return &SlbNewCfgBuddyTable{
		SlbNewCfgRealSerIndex: slbNewCfgRealSerIndex,
		SlbNewCfgBuddyIndex:   slbNewCfgBuddyIndex,
		Params:                params,
	}
}

func (c *SlbNewCfgBuddyTable) Name() string {
	return "SlbNewCfgBuddyTable"
}

func (c *SlbNewCfgBuddyTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbNewCfgBuddyTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbNewCfgBuddyTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbNewCfgRealSerIndex).IsZero() &&
		reflect.ValueOf(c.SlbNewCfgBuddyIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbNewCfgRealSerIndex) + "/" + fmt.Sprint(c.SlbNewCfgBuddyIndex)
}

type SlbNewCfgBuddyTableDelete int32

const (
	SlbNewCfgBuddyTableDelete_Other       SlbNewCfgBuddyTableDelete = 1
	SlbNewCfgBuddyTableDelete_Delete      SlbNewCfgBuddyTableDelete = 2
	SlbNewCfgBuddyTableDelete_Unsupported SlbNewCfgBuddyTableDelete = 2147483647
)

type SlbNewCfgBuddyTableParams struct {
	// The number of the real server.
	RealSerIndex int32 `json:"RealSerIndex,omitempty"`
	// The buddy index. This has no external meaning
	Index int32 `json:"Index,omitempty"`
	// The buddy server's real server number.
	RealIndex int32 `json:"RealIndex,omitempty"`
	// The buddy server's real group number.
	GroupIndex int32 `json:"GroupIndex,omitempty"`
	// The buddy server's service.
	Service uint64 `json:"Service,omitempty"`
	// When set to the value of 2 (delete), the entire row is deleted.
	// When read, other(1) is returned. Setting the value to anything
	// other than 2(delete) has no effect on the state of the row.
	Delete SlbNewCfgBuddyTableDelete `json:"Delete,omitempty"`
}

func (p SlbNewCfgBuddyTableParams) iMABean() {}
