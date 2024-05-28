package beans

import (
	"fmt"
	"reflect"
)

// SlbRealServerRportInfoTable The table of real server rport run-time information.
// Note:This mib is not supported for VX instance of Virtualization.
type SlbRealServerRportInfoTable struct {
	// The real server index
	SlbRealServerRportRealIndex int32
	// The real server service index
	SlbRealServerRportServIndex int32
	Params                      *SlbRealServerRportInfoTableParams
}

func NewSlbRealServerRportInfoTableList() *SlbRealServerRportInfoTable {
	return &SlbRealServerRportInfoTable{}
}

func NewSlbRealServerRportInfoTable(
	slbRealServerRportRealIndex int32,
	slbRealServerRportServIndex int32,
	params *SlbRealServerRportInfoTableParams,
) *SlbRealServerRportInfoTable {
	return &SlbRealServerRportInfoTable{
		SlbRealServerRportRealIndex: slbRealServerRportRealIndex,
		SlbRealServerRportServIndex: slbRealServerRportServIndex,
		Params:                      params,
	}
}

func (c *SlbRealServerRportInfoTable) Name() string {
	return "SlbRealServerRportInfoTable"
}

func (c *SlbRealServerRportInfoTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbRealServerRportInfoTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbRealServerRportInfoTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbRealServerRportRealIndex).IsZero() &&
		reflect.ValueOf(c.SlbRealServerRportServIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbRealServerRportRealIndex) + "/" + fmt.Sprint(c.SlbRealServerRportServIndex)
}

type SlbRealServerRportInfoTableState int32

const (
	SlbRealServerRportInfoTableState_Up          SlbRealServerRportInfoTableState = 1
	SlbRealServerRportInfoTableState_Down        SlbRealServerRportInfoTableState = 2
	SlbRealServerRportInfoTableState_Slowstart   SlbRealServerRportInfoTableState = 3
	SlbRealServerRportInfoTableState_Unsupported SlbRealServerRportInfoTableState = 2147483647
)

type SlbRealServerRportInfoTableParams struct {
	// The real server index
	RealIndex int32 `json:"RealIndex,omitempty"`
	// The real server service index
	ServIndex int32 `json:"ServIndex,omitempty"`
	// The real server real port number
	Rport int32 `json:"Rport,omitempty"`
	// The state of the real server port.
	State SlbRealServerRportInfoTableState `json:"State,omitempty"`
	// The real server group number.
	Group int32 `json:"Group,omitempty"`
	// The response time of the real server health check.
	RespTime int32 `json:"RespTime,omitempty"`
	// The reason for health check fail.
	FailReason string `json:"FailReason,omitempty"`
}

func (p SlbRealServerRportInfoTableParams) iMABean() {}
