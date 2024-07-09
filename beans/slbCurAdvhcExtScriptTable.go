package beans

import (
	"fmt"
	"reflect"
)

// SlbCurAdvhcExtScriptTable Note:This mib is not supported on VX instance of Virtualization.
type SlbCurAdvhcExtScriptTable struct {
	// Script Health check id.
	SlbCurAdvhcExtScriptID string
	Params                 *SlbCurAdvhcExtScriptTableParams
}

func NewSlbCurAdvhcExtScriptTableList() *SlbCurAdvhcExtScriptTable {
	return &SlbCurAdvhcExtScriptTable{}
}

func NewSlbCurAdvhcExtScriptTable(
	slbCurAdvhcExtScriptID string,
	params *SlbCurAdvhcExtScriptTableParams,
) *SlbCurAdvhcExtScriptTable {
	return &SlbCurAdvhcExtScriptTable{
		SlbCurAdvhcExtScriptID: slbCurAdvhcExtScriptID,
		Params:                 params,
	}
}

func (c *SlbCurAdvhcExtScriptTable) Name() string {
	return "SlbCurAdvhcExtScriptTable"
}

func (c *SlbCurAdvhcExtScriptTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbCurAdvhcExtScriptTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbCurAdvhcExtScriptTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbCurAdvhcExtScriptID).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbCurAdvhcExtScriptID)
}

type SlbCurAdvhcExtScriptTableIPVer int32

const (
	SlbCurAdvhcExtScriptTableIPVer_Ipv4 SlbCurAdvhcExtScriptTableIPVer = 1
	SlbCurAdvhcExtScriptTableIPVer_Ipv6 SlbCurAdvhcExtScriptTableIPVer = 2
	SlbCurAdvhcExtScriptTableIPVer_None SlbCurAdvhcExtScriptTableIPVer = 3
)

type SlbCurAdvhcExtScriptTableInvert int32

const (
	SlbCurAdvhcExtScriptTableInvert_Enabled     SlbCurAdvhcExtScriptTableInvert = 1
	SlbCurAdvhcExtScriptTableInvert_Disabled    SlbCurAdvhcExtScriptTableInvert = 2
	SlbCurAdvhcExtScriptTableInvert_Unsupported SlbCurAdvhcExtScriptTableInvert = 2147483647
)

type SlbCurAdvhcExtScriptTableParams struct {
	// Script Health check id.
	ID string `json:"ID,omitempty"`
	// Script Health check name.
	Name string `json:"Name,omitempty"`
	// Script ID of the External Health check file associated.
	StringVal string `json:"StringVal,omitempty"`
	// Script Health check port.
	DPort uint64 `json:"DPort,omitempty"`
	// Script Health check Ip Version. (1) ipv4, (2) ipv6
	IPVer SlbCurAdvhcExtScriptTableIPVer `json:"IPVer,omitempty"`
	// Script Health check ip address / Hostname.
	HostName string `json:"HostName,omitempty"`
	// Invert Result
	Invert SlbCurAdvhcExtScriptTableInvert `json:"Invert,omitempty"`
	// Script Health check retries counter.
	Retries uint64 `json:"Retries,omitempty"`
	// Script Health check retries in the down state counter.
	RestoreRetries uint64 `json:"RestoreRetries,omitempty"`
	// Script Health check timeout.
	Timeout uint64 `json:"Timeout,omitempty"`
	// Script Health check interval in the down state.
	DownInterval uint64 `json:"DownInterval,omitempty"`
	// Script Health check interval.
	Interval uint64 `json:"Interval,omitempty"`
}

func (p SlbCurAdvhcExtScriptTableParams) iMABean() {}
