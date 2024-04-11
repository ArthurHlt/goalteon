package beans

import (
	"fmt"
	"reflect"
)

// SlbNewCfgRealServPortTable The table of real server service ports.
// Note:This mib is not supported on VX instance of Virtualization.
type SlbNewCfgRealServPortTable struct {
	// The number of the real server.
	SlbNewCfgRealServIndex int32
	// The port index. This has no external meaning
	SlbNewCfgRealServPortIndex int32
	Params                     *SlbNewCfgRealServPortTableParams
}

func NewSlbNewCfgRealServPortTableList() *SlbNewCfgRealServPortTable {
	return &SlbNewCfgRealServPortTable{}
}

func NewSlbNewCfgRealServPortTable(
	slbNewCfgRealServIndex int32,
	slbNewCfgRealServPortIndex int32,
	params *SlbNewCfgRealServPortTableParams,
) *SlbNewCfgRealServPortTable {
	return &SlbNewCfgRealServPortTable{
		SlbNewCfgRealServIndex:     slbNewCfgRealServIndex,
		SlbNewCfgRealServPortIndex: slbNewCfgRealServPortIndex,
		Params:                     params,
	}
}

func (c *SlbNewCfgRealServPortTable) Name() string {
	return "SlbNewCfgRealServPortTable"
}

func (c *SlbNewCfgRealServPortTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbNewCfgRealServPortTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbNewCfgRealServPortTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbNewCfgRealServIndex).IsZero() &&
		reflect.ValueOf(c.SlbNewCfgRealServPortIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbNewCfgRealServIndex) + "/" + fmt.Sprint(c.SlbNewCfgRealServPortIndex)
}

type SlbNewCfgRealServPortTableDelete int32

const (
	SlbNewCfgRealServPortTableDelete_Other       SlbNewCfgRealServPortTableDelete = 1
	SlbNewCfgRealServPortTableDelete_Delete      SlbNewCfgRealServPortTableDelete = 2
	SlbNewCfgRealServPortTableDelete_Unsupported SlbNewCfgRealServPortTableDelete = 2147483647
)

type SlbNewCfgRealServPortTableParams struct {
	// The number of the real server.
	Index int32 `json:"Index,omitempty"`
	// The port index. This has no external meaning
	PortIndex int32 `json:"PortIndex,omitempty"`
	// The layer4 real service port number.
	RealPort uint64 `json:"RealPort,omitempty"`
	// When set to the value of 2 (delete), the entire row is deleted.
	// When read, other(1) is returned. Setting the value to anything
	// other than 2(delete) has no effect on the state of the row.
	Delete SlbNewCfgRealServPortTableDelete `json:"Delete,omitempty"`
	// The first free port index number of the real server.
	// Value 0 will be returned when all 64 ports are
	// configured for a real server.
	RealPortFreeIdx uint64 `json:"RealPortFreeIdx,omitempty"`
}

func (p SlbNewCfgRealServPortTableParams) iMABean() {}
