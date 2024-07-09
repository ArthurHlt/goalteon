package beans

import (
	"fmt"
	"reflect"
)

// SlbCurAdvhcRadiusTable Note:This mib is not supported on VX instance of Virtualization.
type SlbCurAdvhcRadiusTable struct {
	// RADIUS health check id.
	SlbCurAdvhcRadiusID string
	Params              *SlbCurAdvhcRadiusTableParams
}

func NewSlbCurAdvhcRadiusTableList() *SlbCurAdvhcRadiusTable {
	return &SlbCurAdvhcRadiusTable{}
}

func NewSlbCurAdvhcRadiusTable(
	slbCurAdvhcRadiusID string,
	params *SlbCurAdvhcRadiusTableParams,
) *SlbCurAdvhcRadiusTable {
	return &SlbCurAdvhcRadiusTable{
		SlbCurAdvhcRadiusID: slbCurAdvhcRadiusID,
		Params:              params,
	}
}

func (c *SlbCurAdvhcRadiusTable) Name() string {
	return "SlbCurAdvhcRadiusTable"
}

func (c *SlbCurAdvhcRadiusTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbCurAdvhcRadiusTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbCurAdvhcRadiusTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbCurAdvhcRadiusID).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbCurAdvhcRadiusID)
}

type SlbCurAdvhcRadiusTableIPVer int32

const (
	SlbCurAdvhcRadiusTableIPVer_Ipv4 SlbCurAdvhcRadiusTableIPVer = 1
	SlbCurAdvhcRadiusTableIPVer_Ipv6 SlbCurAdvhcRadiusTableIPVer = 2
	SlbCurAdvhcRadiusTableIPVer_None SlbCurAdvhcRadiusTableIPVer = 3
)

type SlbCurAdvhcRadiusTableTransparent int32

const (
	SlbCurAdvhcRadiusTableTransparent_Enabled     SlbCurAdvhcRadiusTableTransparent = 1
	SlbCurAdvhcRadiusTableTransparent_Disabled    SlbCurAdvhcRadiusTableTransparent = 2
	SlbCurAdvhcRadiusTableTransparent_Unsupported SlbCurAdvhcRadiusTableTransparent = 2147483647
)

type SlbCurAdvhcRadiusTableInvert int32

const (
	SlbCurAdvhcRadiusTableInvert_Enabled     SlbCurAdvhcRadiusTableInvert = 1
	SlbCurAdvhcRadiusTableInvert_Disabled    SlbCurAdvhcRadiusTableInvert = 2
	SlbCurAdvhcRadiusTableInvert_Unsupported SlbCurAdvhcRadiusTableInvert = 2147483647
)

type SlbCurAdvhcRadiusTableDownType int32

const (
	SlbCurAdvhcRadiusTableDownType_Accounting     SlbCurAdvhcRadiusTableDownType = 1
	SlbCurAdvhcRadiusTableDownType_Authentication SlbCurAdvhcRadiusTableDownType = 2
	SlbCurAdvhcRadiusTableDownType_Any            SlbCurAdvhcRadiusTableDownType = 3
	SlbCurAdvhcRadiusTableDownType_Unsupported    SlbCurAdvhcRadiusTableDownType = 2147483647
)

type SlbCurAdvhcRadiusTableSnat int32

const (
	SlbCurAdvhcRadiusTableSnat_Enabled     SlbCurAdvhcRadiusTableSnat = 1
	SlbCurAdvhcRadiusTableSnat_Disabled    SlbCurAdvhcRadiusTableSnat = 2
	SlbCurAdvhcRadiusTableSnat_Unsupported SlbCurAdvhcRadiusTableSnat = 2147483647
)

type SlbCurAdvhcRadiusTableParams struct {
	// RADIUS health check id.
	ID string `json:"ID,omitempty"`
	// RADIUS health check name.
	Name string `json:"Name,omitempty"`
	// RADIUS health check destination port.
	DPort uint64 `json:"DPort,omitempty"`
	// RADIUS health check destination IP version.
	IPVer SlbCurAdvhcRadiusTableIPVer `json:"IPVer,omitempty"`
	// RADIUS health check destination hostname.
	HostName string `json:"HostName,omitempty"`
	// RADIUS health check transparent flag.
	Transparent SlbCurAdvhcRadiusTableTransparent `json:"Transparent,omitempty"`
	// RADIUS health check interval.
	Interval uint64 `json:"Interval,omitempty"`
	// RADIUS health check retries counter.
	Retries uint64 `json:"Retries,omitempty"`
	// RADIUS health check retries in the down state counter.
	RestoreRetries uint64 `json:"RestoreRetries,omitempty"`
	// RADIUS health check timeout parameter.
	Timeout uint64 `json:"Timeout,omitempty"`
	// RADIUS health check overflow flag.
	Overflow uint64 `json:"Overflow,omitempty"`
	// RADIUS health check interval in down state.
	DownInterval uint64 `json:"DownInterval,omitempty"`
	// RADIUS health check invert flag.
	Invert SlbCurAdvhcRadiusTableInvert `json:"Invert,omitempty"`
	// RADIUS health check type.
	DownType SlbCurAdvhcRadiusTableDownType `json:"DownType,omitempty"`
	// RADIUS health check user name.
	UserName string `json:"UserName,omitempty"`
	// RADIUS health check password.
	Password string `json:"Password,omitempty"`
	// RADIUS health check secret.
	Secret string `json:"Secret,omitempty"`
	// RADIUS Health Check src NAT (PIP) flag.
	Snat SlbCurAdvhcRadiusTableSnat `json:"Snat,omitempty"`
}

func (p SlbCurAdvhcRadiusTableParams) iMABean() {}
