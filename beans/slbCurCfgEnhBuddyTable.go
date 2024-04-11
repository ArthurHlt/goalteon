package beans

import (
	"fmt"
	"reflect"
)

// SlbCurCfgEnhBuddyTable The table of buddy information.
// Note:This mib is not supported on VX instance of Virtualization.
type SlbCurCfgEnhBuddyTable struct {
	// The number of the real server.
	SlbCurCfgEnhRealSerIndex string
	// The buddy index. This has no external meaning
	SlbCurCfgEnhBuddyIndex int32
	Params                 *SlbCurCfgEnhBuddyTableParams
}

func NewSlbCurCfgEnhBuddyTableList() *SlbCurCfgEnhBuddyTable {
	return &SlbCurCfgEnhBuddyTable{}
}

func NewSlbCurCfgEnhBuddyTable(
	slbCurCfgEnhRealSerIndex string,
	slbCurCfgEnhBuddyIndex int32,
	params *SlbCurCfgEnhBuddyTableParams,
) *SlbCurCfgEnhBuddyTable {
	return &SlbCurCfgEnhBuddyTable{
		SlbCurCfgEnhRealSerIndex: slbCurCfgEnhRealSerIndex,
		SlbCurCfgEnhBuddyIndex:   slbCurCfgEnhBuddyIndex,
		Params:                   params,
	}
}

func (c *SlbCurCfgEnhBuddyTable) Name() string {
	return "SlbCurCfgEnhBuddyTable"
}

func (c *SlbCurCfgEnhBuddyTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbCurCfgEnhBuddyTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbCurCfgEnhBuddyTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbCurCfgEnhRealSerIndex).IsZero() &&
		reflect.ValueOf(c.SlbCurCfgEnhBuddyIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbCurCfgEnhRealSerIndex) + "/" + fmt.Sprint(c.SlbCurCfgEnhBuddyIndex)
}

type SlbCurCfgEnhBuddyTableParams struct {
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
}

func (p SlbCurCfgEnhBuddyTableParams) iMABean() {}
