package beans

import (
	"fmt"
	"reflect"
)

// SlbCurCfgEnhRealServPortTable The table of real server service ports.
// Note:This mib is not supported on VX instance of Virtualization.
type SlbCurCfgEnhRealServPortTable struct {
	// The number of the real server.
	SlbCurCfgEnhRealServIndex string
	// The port index. This has no external meaning
	SlbCurCfgEnhRealServPortIndex int32
	Params                        *SlbCurCfgEnhRealServPortTableParams
}

func NewSlbCurCfgEnhRealServPortTableList() *SlbCurCfgEnhRealServPortTable {
	return &SlbCurCfgEnhRealServPortTable{}
}

func NewSlbCurCfgEnhRealServPortTable(
	slbCurCfgEnhRealServIndex string,
	slbCurCfgEnhRealServPortIndex int32,
	params *SlbCurCfgEnhRealServPortTableParams,
) *SlbCurCfgEnhRealServPortTable {
	return &SlbCurCfgEnhRealServPortTable{
		SlbCurCfgEnhRealServIndex:     slbCurCfgEnhRealServIndex,
		SlbCurCfgEnhRealServPortIndex: slbCurCfgEnhRealServPortIndex,
		Params:                        params,
	}
}

func (c *SlbCurCfgEnhRealServPortTable) Name() string {
	return "SlbCurCfgEnhRealServPortTable"
}

func (c *SlbCurCfgEnhRealServPortTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbCurCfgEnhRealServPortTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbCurCfgEnhRealServPortTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbCurCfgEnhRealServIndex).IsZero() &&
		reflect.ValueOf(c.SlbCurCfgEnhRealServPortIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbCurCfgEnhRealServIndex) + "/" + fmt.Sprint(c.SlbCurCfgEnhRealServPortIndex)
}

type SlbCurCfgEnhRealServPortTableParams struct {
	// The number of the real server.
	Index string `json:"Index,omitempty"`
	// The port index. This has no external meaning
	PortIndex int32 `json:"PortIndex,omitempty"`
	// The layer4 real service port number.
	RealPort uint64 `json:"RealPort,omitempty"`
}

func (p SlbCurCfgEnhRealServPortTableParams) iMABean() {}
