package beans

import (
	"fmt"
	"reflect"
)

// SlbNewAdvhcExtScriptTable Note:This mib is not supported on VX instance of Virtualization.
type SlbNewAdvhcExtScriptTable struct {
	// Script Health check id.
	SlbNewAdvhcExtScriptID string
	Params                 *SlbNewAdvhcExtScriptTableParams
}

func NewSlbNewAdvhcExtScriptTableList() *SlbNewAdvhcExtScriptTable {
	return &SlbNewAdvhcExtScriptTable{}
}

func NewSlbNewAdvhcExtScriptTable(
	slbNewAdvhcExtScriptID string,
	params *SlbNewAdvhcExtScriptTableParams,
) *SlbNewAdvhcExtScriptTable {
	return &SlbNewAdvhcExtScriptTable{
		SlbNewAdvhcExtScriptID: slbNewAdvhcExtScriptID,
		Params:                 params,
	}
}

func (c *SlbNewAdvhcExtScriptTable) Name() string {
	return "SlbNewAdvhcExtScriptTable"
}

func (c *SlbNewAdvhcExtScriptTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbNewAdvhcExtScriptTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbNewAdvhcExtScriptTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbNewAdvhcExtScriptID).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbNewAdvhcExtScriptID)
}

type SlbNewAdvhcExtScriptTableDelete int32

const (
	SlbNewAdvhcExtScriptTableDelete_Other       SlbNewAdvhcExtScriptTableDelete = 1
	SlbNewAdvhcExtScriptTableDelete_Delete      SlbNewAdvhcExtScriptTableDelete = 2
	SlbNewAdvhcExtScriptTableDelete_Unsupported SlbNewAdvhcExtScriptTableDelete = 2147483647
)

type SlbNewAdvhcExtScriptTableIPVer int32

const (
	SlbNewAdvhcExtScriptTableIPVer_Ipv4 SlbNewAdvhcExtScriptTableIPVer = 1
	SlbNewAdvhcExtScriptTableIPVer_Ipv6 SlbNewAdvhcExtScriptTableIPVer = 2
	SlbNewAdvhcExtScriptTableIPVer_None SlbNewAdvhcExtScriptTableIPVer = 3
)

type SlbNewAdvhcExtScriptTableInvert int32

const (
	SlbNewAdvhcExtScriptTableInvert_Enabled     SlbNewAdvhcExtScriptTableInvert = 1
	SlbNewAdvhcExtScriptTableInvert_Disabled    SlbNewAdvhcExtScriptTableInvert = 2
	SlbNewAdvhcExtScriptTableInvert_Unsupported SlbNewAdvhcExtScriptTableInvert = 2147483647
)

type SlbNewAdvhcExtScriptTableParams struct {
	// Script Health check id.
	ID string `json:"ID,omitempty"`
	// Script Health check name.
	Name string `json:"Name,omitempty"`
	// Script ID of the External Health check file associated.
	StringVal string `json:"StringVal,omitempty"`
	// Script Health check copy indicator.
	Copy DisplayString `json:"Copy,omitempty"`
	// When set to the value of 2 (delete), the entire row is deleted.
	// When read, other(1) is returned. Setting the value to anything
	// other than 2(delete) has no effect on the state of the row.
	Delete SlbNewAdvhcExtScriptTableDelete `json:"Delete,omitempty"`
	// Script Health check port.
	DPort uint64 `json:"DPort,omitempty"`
	// Script Health check Ip Version. Set ipv4(1) and ipv6(2).
	IPVer SlbNewAdvhcExtScriptTableIPVer `json:"IPVer,omitempty"`
	// Script Health check ip address / Hostname.
	HostName string `json:"HostName,omitempty"`
	// Invert Result
	Invert SlbNewAdvhcExtScriptTableInvert `json:"Invert,omitempty"`
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

func (p SlbNewAdvhcExtScriptTableParams) iMABean() {}
