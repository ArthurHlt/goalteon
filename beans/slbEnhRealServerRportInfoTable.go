package beans

import (
	"fmt"
	"reflect"
)

// SlbEnhRealServerRportInfoTable The table of real server rport run-time information.
// Note:This mib is not supported for VX instance of Virtualization.
type SlbEnhRealServerRportInfoTable struct {
	// The real server index
	SlbEnhRealServerRportRealIndex string
	// The real server service index
	SlbEnhRealServerRportServIndex int32
	Params                         *SlbEnhRealServerRportInfoTableParams
}

func NewSlbEnhRealServerRportInfoTableList() *SlbEnhRealServerRportInfoTable {
	return &SlbEnhRealServerRportInfoTable{}
}

func NewSlbEnhRealServerRportInfoTable(
	slbEnhRealServerRportRealIndex string,
	slbEnhRealServerRportServIndex int32,
	params *SlbEnhRealServerRportInfoTableParams,
) *SlbEnhRealServerRportInfoTable {
	return &SlbEnhRealServerRportInfoTable{
		SlbEnhRealServerRportRealIndex: slbEnhRealServerRportRealIndex,
		SlbEnhRealServerRportServIndex: slbEnhRealServerRportServIndex,
		Params:                         params,
	}
}

func (c *SlbEnhRealServerRportInfoTable) Name() string {
	return "SlbEnhRealServerRportInfoTable"
}

func (c *SlbEnhRealServerRportInfoTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbEnhRealServerRportInfoTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbEnhRealServerRportInfoTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbEnhRealServerRportRealIndex).IsZero() &&
		reflect.ValueOf(c.SlbEnhRealServerRportServIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbEnhRealServerRportRealIndex) + "/" + fmt.Sprint(c.SlbEnhRealServerRportServIndex)
}

type SlbEnhRealServerRportInfoTableState int32

const (
	SlbEnhRealServerRportInfoTableState_Up          SlbEnhRealServerRportInfoTableState = 1
	SlbEnhRealServerRportInfoTableState_Down        SlbEnhRealServerRportInfoTableState = 2
	SlbEnhRealServerRportInfoTableState_Slowstart   SlbEnhRealServerRportInfoTableState = 3
	SlbEnhRealServerRportInfoTableState_Unsupported SlbEnhRealServerRportInfoTableState = 2147483647
)

type SlbEnhRealServerRportInfoTableParams struct {
	// The real server index
	RealIndex string `json:"RealIndex,omitempty"`
	// The real server service index
	ServIndex int32 `json:"ServIndex,omitempty"`
	// The real server real port number
	Rport int32 `json:"Rport,omitempty"`
	// The state of the real server port.
	State SlbEnhRealServerRportInfoTableState `json:"State,omitempty"`
	// The real server group number.
	Group string `json:"Group,omitempty"`
	// The response time of the real server health check.
	RespTime int32 `json:"RespTime,omitempty"`
	// The reason for health check fail.
	FailReason string `json:"FailReason,omitempty"`
}

func (p SlbEnhRealServerRportInfoTableParams) iMABean() {}
