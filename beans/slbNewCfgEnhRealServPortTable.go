package beans

import (
	"fmt"
	"reflect"
)

// SlbNewCfgEnhRealServPortTable The table of real server service ports.
// Note:This mib is not supported on VX instance of Virtualization.
type SlbNewCfgEnhRealServPortTable struct {
	// The number of the real server.
	SlbNewCfgEnhRealServIndex string
	// The port index. This has no external meaning
	SlbNewCfgEnhRealServPortIndex int32
	Params                        *SlbNewCfgEnhRealServPortTableParams
}

func NewSlbNewCfgEnhRealServPortTableList() *SlbNewCfgEnhRealServPortTable {
	return &SlbNewCfgEnhRealServPortTable{}
}

func NewSlbNewCfgEnhRealServPortTable(
	slbNewCfgEnhRealServIndex string,
	slbNewCfgEnhRealServPortIndex int32,
	params *SlbNewCfgEnhRealServPortTableParams,
) *SlbNewCfgEnhRealServPortTable {
	return &SlbNewCfgEnhRealServPortTable{
		SlbNewCfgEnhRealServIndex:     slbNewCfgEnhRealServIndex,
		SlbNewCfgEnhRealServPortIndex: slbNewCfgEnhRealServPortIndex,
		Params:                        params,
	}
}

func (c *SlbNewCfgEnhRealServPortTable) Name() string {
	return "SlbNewCfgEnhRealServPortTable"
}

func (c *SlbNewCfgEnhRealServPortTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbNewCfgEnhRealServPortTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbNewCfgEnhRealServPortTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbNewCfgEnhRealServIndex).IsZero() &&
		reflect.ValueOf(c.SlbNewCfgEnhRealServPortIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbNewCfgEnhRealServIndex) + "/" + fmt.Sprint(c.SlbNewCfgEnhRealServPortIndex)
}

type SlbNewCfgEnhRealServPortTableDelete int32

const (
	SlbNewCfgEnhRealServPortTableDelete_Other  SlbNewCfgEnhRealServPortTableDelete = 1
	SlbNewCfgEnhRealServPortTableDelete_Delete SlbNewCfgEnhRealServPortTableDelete = 2
)

type SlbNewCfgEnhRealServPortTableParams struct {
	// The number of the real server.
	Index string `json:"Index,omitempty"`
	// The port index. This has no external meaning
	PortIndex int32 `json:"PortIndex,omitempty"`
	// The layer4 real service port number.
	RealPort uint64 `json:"RealPort,omitempty"`
	// When set to the value of 2 (delete), the entire row is deleted.
	// When read, other(1) is returned. Setting the value to anything
	// other than 2(delete) has no effect on the state of the row.
	Delete SlbNewCfgEnhRealServPortTableDelete `json:"Delete,omitempty"`
	// The first free port index number of the real server.
	// Value 0 will be returned when all 64 ports are
	// configured for a real server.
	RealPortFreeIdx uint64 `json:"RealPortFreeIdx,omitempty"`
}

func (p SlbNewCfgEnhRealServPortTableParams) iMABean() {}
