package beans

import (
	"fmt"
	"reflect"
)

// SlbCurCfgBuddyTable The table of buddy information.
// Note:This mib is not supported on VX instance of Virtualization.
type SlbCurCfgBuddyTable struct {
	// The number of the real server.
	SlbCurCfgRealSerIndex int32
	// The buddy index. This has no external meaning
	SlbCurCfgBuddyIndex int32
	Params              *SlbCurCfgBuddyTableParams
}

func NewSlbCurCfgBuddyTableList() *SlbCurCfgBuddyTable {
	return &SlbCurCfgBuddyTable{}
}

func NewSlbCurCfgBuddyTable(
	slbCurCfgRealSerIndex int32,
	slbCurCfgBuddyIndex int32,
	params *SlbCurCfgBuddyTableParams,
) *SlbCurCfgBuddyTable {
	return &SlbCurCfgBuddyTable{
		SlbCurCfgRealSerIndex: slbCurCfgRealSerIndex,
		SlbCurCfgBuddyIndex:   slbCurCfgBuddyIndex,
		Params:                params,
	}
}

func (c *SlbCurCfgBuddyTable) Name() string {
	return "SlbCurCfgBuddyTable"
}

func (c *SlbCurCfgBuddyTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbCurCfgBuddyTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbCurCfgBuddyTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbCurCfgRealSerIndex).IsZero() &&
		reflect.ValueOf(c.SlbCurCfgBuddyIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbCurCfgRealSerIndex) + "/" + fmt.Sprint(c.SlbCurCfgBuddyIndex)
}

type SlbCurCfgBuddyTableParams struct {
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
}

func (p SlbCurCfgBuddyTableParams) iMABean() {}
