package beans

import (
	"fmt"
	"reflect"
)

// SwitchCapSlbPortInfoTable The table of slb port information.
// Note:This mib is not supported for VX instance of virtualization.
type SwitchCapSlbPortInfoTable struct {
	// Index of switch port. Note:This mib is not supported for VX instance of virtualization.
	SwitchCapSlbPortInfoIndx int32
	Params                   *SwitchCapSlbPortInfoTableParams
}

func NewSwitchCapSlbPortInfoTableList() *SwitchCapSlbPortInfoTable {
	return &SwitchCapSlbPortInfoTable{}
}

func NewSwitchCapSlbPortInfoTable(
	switchCapSlbPortInfoIndx int32,
	params *SwitchCapSlbPortInfoTableParams,
) *SwitchCapSlbPortInfoTable {
	return &SwitchCapSlbPortInfoTable{
		SwitchCapSlbPortInfoIndx: switchCapSlbPortInfoIndx,
		Params:                   params,
	}
}

func (c *SwitchCapSlbPortInfoTable) Name() string {
	return "SwitchCapSlbPortInfoTable"
}

func (c *SwitchCapSlbPortInfoTable) GetParams() BeanType {
	return c.Params
}

func (c *SwitchCapSlbPortInfoTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SwitchCapSlbPortInfoTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SwitchCapSlbPortInfoIndx).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SwitchCapSlbPortInfoIndx)
}

type SwitchCapSlbPortInfoTableClientState int32

const (
	SwitchCapSlbPortInfoTableClientState_Enabled     SwitchCapSlbPortInfoTableClientState = 1
	SwitchCapSlbPortInfoTableClientState_Disabled    SwitchCapSlbPortInfoTableClientState = 2
	SwitchCapSlbPortInfoTableClientState_Unsupported SwitchCapSlbPortInfoTableClientState = 2147483647
)

type SwitchCapSlbPortInfoTableSerState int32

const (
	SwitchCapSlbPortInfoTableSerState_Enabled     SwitchCapSlbPortInfoTableSerState = 1
	SwitchCapSlbPortInfoTableSerState_Disabled    SwitchCapSlbPortInfoTableSerState = 2
	SwitchCapSlbPortInfoTableSerState_Unsupported SwitchCapSlbPortInfoTableSerState = 2147483647
)

type SwitchCapSlbPortInfoTableRTSState int32

const (
	SwitchCapSlbPortInfoTableRTSState_Enabled     SwitchCapSlbPortInfoTableRTSState = 1
	SwitchCapSlbPortInfoTableRTSState_Disabled    SwitchCapSlbPortInfoTableRTSState = 2
	SwitchCapSlbPortInfoTableRTSState_Unsupported SwitchCapSlbPortInfoTableRTSState = 2147483647
)

type SwitchCapSlbPortInfoTableParams struct {
	// Index of switch port. Note:This mib is not supported for VX instance of virtualization.
	Indx int32 `json:"Indx,omitempty"`
	// Client state on the switch port.
	// Note:This mib is not supported for VX instance of virtualization.
	ClientState SwitchCapSlbPortInfoTableClientState `json:"ClientState,omitempty"`
	// Server state on the switch port.
	// Note:This mib is not supported for VX instance of virtualization.
	SerState SwitchCapSlbPortInfoTableSerState `json:"SerState,omitempty"`
	// It is in the form E(x) where filter is enabled on the port and
	// x is the number of filters configured for port. Otherwise disable.
	// Note:This mib is not supported for VX instance of virtualization.
	FltState string `json:"FltState,omitempty"`
	// RTS processing state on the switch port.
	// Note:This mib is not supported for VX instance of virtualization.
	RTSState SwitchCapSlbPortInfoTableRTSState `json:"RTSState,omitempty"`
}

func (p SwitchCapSlbPortInfoTableParams) iMABean() {}
