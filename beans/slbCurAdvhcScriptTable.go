package beans

import (
	"fmt"
	"reflect"
)

// SlbCurAdvhcScriptTable Note:This mib is not supported on VX instance of Virtualization.
type SlbCurAdvhcScriptTable struct {
	// Script Health check id.
	SlbCurAdvhcScriptID string
	Params              *SlbCurAdvhcScriptTableParams
}

func NewSlbCurAdvhcScriptTableList() *SlbCurAdvhcScriptTable {
	return &SlbCurAdvhcScriptTable{}
}

func NewSlbCurAdvhcScriptTable(
	slbCurAdvhcScriptID string,
	params *SlbCurAdvhcScriptTableParams,
) *SlbCurAdvhcScriptTable {
	return &SlbCurAdvhcScriptTable{
		SlbCurAdvhcScriptID: slbCurAdvhcScriptID,
		Params:              params,
	}
}

func (c *SlbCurAdvhcScriptTable) Name() string {
	return "SlbCurAdvhcScriptTable"
}

func (c *SlbCurAdvhcScriptTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbCurAdvhcScriptTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbCurAdvhcScriptTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbCurAdvhcScriptID).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbCurAdvhcScriptID)
}

type SlbCurAdvhcScriptTableIPVer int32

const (
	SlbCurAdvhcScriptTableIPVer_Ipv4 SlbCurAdvhcScriptTableIPVer = 1
	SlbCurAdvhcScriptTableIPVer_Ipv6 SlbCurAdvhcScriptTableIPVer = 2
	SlbCurAdvhcScriptTableIPVer_None SlbCurAdvhcScriptTableIPVer = 3
)

type SlbCurAdvhcScriptTableInvert int32

const (
	SlbCurAdvhcScriptTableInvert_Enabled     SlbCurAdvhcScriptTableInvert = 1
	SlbCurAdvhcScriptTableInvert_Disabled    SlbCurAdvhcScriptTableInvert = 2
	SlbCurAdvhcScriptTableInvert_Unsupported SlbCurAdvhcScriptTableInvert = 2147483647
)

type SlbCurAdvhcScriptTableTransparent int32

const (
	SlbCurAdvhcScriptTableTransparent_Enabled     SlbCurAdvhcScriptTableTransparent = 1
	SlbCurAdvhcScriptTableTransparent_Disabled    SlbCurAdvhcScriptTableTransparent = 2
	SlbCurAdvhcScriptTableTransparent_Unsupported SlbCurAdvhcScriptTableTransparent = 2147483647
)

type SlbCurAdvhcScriptTableAlways int32

const (
	SlbCurAdvhcScriptTableAlways_Enabled     SlbCurAdvhcScriptTableAlways = 1
	SlbCurAdvhcScriptTableAlways_Disabled    SlbCurAdvhcScriptTableAlways = 2
	SlbCurAdvhcScriptTableAlways_Unsupported SlbCurAdvhcScriptTableAlways = 2147483647
)

type SlbCurAdvhcScriptTableParams struct {
	// Script Health check id.
	ID string `json:"ID,omitempty"`
	// Script Health check name.
	Name string `json:"Name,omitempty"`
	// Script Health check string.
	StringVal string `json:"StringVal,omitempty"`
	// Script Health check Ip Version.
	IPVer SlbCurAdvhcScriptTableIPVer `json:"IPVer,omitempty"`
	// Script Healtch check ip address / Hostname.
	HostName string `json:"HostName,omitempty"`
	// Invert Result
	Invert SlbCurAdvhcScriptTableInvert `json:"Invert,omitempty"`
	// Script Health check retries counter.
	Retries uint64 `json:"Retries,omitempty"`
	// Script Health check retries in the down state counter.
	RestoreRetries uint64 `json:"RestoreRetries,omitempty"`
	// Script Health check timeout.
	Timeout uint64 `json:"Timeout,omitempty"`
	// Script Health check overflow flag.
	Overflow uint64 `json:"Overflow,omitempty"`
	// Script Health check interval in the down state.
	DownInterval uint64 `json:"DownInterval,omitempty"`
	// Script Health check transparent flag.
	Transparent SlbCurAdvhcScriptTableTransparent `json:"Transparent,omitempty"`
	// Script Health check interval.
	Interval uint64 `json:"Interval,omitempty"`
	// This flag determines whether HC is allowed for standalone real.
	Always SlbCurAdvhcScriptTableAlways `json:"Always,omitempty"`
}

func (p SlbCurAdvhcScriptTableParams) iMABean() {}
