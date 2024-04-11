package beans

import (
	"fmt"
	"reflect"
)

// SlbCurCfgRealServPortTable The table of real server service ports.
// Note:This mib is not supported on VX instance of Virtualization.
type SlbCurCfgRealServPortTable struct {
	// The number of the real server.
	SlbCurCfgRealServIndex int32
	// The port index. This has no external meaning
	SlbCurCfgRealServPortIndex int32
	Params                     *SlbCurCfgRealServPortTableParams
}

func NewSlbCurCfgRealServPortTableList() *SlbCurCfgRealServPortTable {
	return &SlbCurCfgRealServPortTable{}
}

func NewSlbCurCfgRealServPortTable(
	slbCurCfgRealServIndex int32,
	slbCurCfgRealServPortIndex int32,
	params *SlbCurCfgRealServPortTableParams,
) *SlbCurCfgRealServPortTable {
	return &SlbCurCfgRealServPortTable{
		SlbCurCfgRealServIndex:     slbCurCfgRealServIndex,
		SlbCurCfgRealServPortIndex: slbCurCfgRealServPortIndex,
		Params:                     params,
	}
}

func (c *SlbCurCfgRealServPortTable) Name() string {
	return "SlbCurCfgRealServPortTable"
}

func (c *SlbCurCfgRealServPortTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbCurCfgRealServPortTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbCurCfgRealServPortTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbCurCfgRealServIndex).IsZero() &&
		reflect.ValueOf(c.SlbCurCfgRealServPortIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbCurCfgRealServIndex) + "/" + fmt.Sprint(c.SlbCurCfgRealServPortIndex)
}

type SlbCurCfgRealServPortTableParams struct {
	// The number of the real server.
	Index int32 `json:"Index,omitempty"`
	// The port index. This has no external meaning
	PortIndex int32 `json:"PortIndex,omitempty"`
	// The layer4 real service port number.
	RealPort uint64 `json:"RealPort,omitempty"`
}

func (p SlbCurCfgRealServPortTableParams) iMABean() {}
