package beans

import (
	"fmt"
	"reflect"
)

// SlbNewAdvhcRadiusTable Note:This mib is not supported on VX instance of Virtualization.
type SlbNewAdvhcRadiusTable struct {
	// RADIUS health check id.
	SlbNewAdvhcRadiusID string
	Params              *SlbNewAdvhcRadiusTableParams
}

func NewSlbNewAdvhcRadiusTableList() *SlbNewAdvhcRadiusTable {
	return &SlbNewAdvhcRadiusTable{}
}

func NewSlbNewAdvhcRadiusTable(
	slbNewAdvhcRadiusID string,
	params *SlbNewAdvhcRadiusTableParams,
) *SlbNewAdvhcRadiusTable {
	return &SlbNewAdvhcRadiusTable{
		SlbNewAdvhcRadiusID: slbNewAdvhcRadiusID,
		Params:              params,
	}
}

func (c *SlbNewAdvhcRadiusTable) Name() string {
	return "SlbNewAdvhcRadiusTable"
}

func (c *SlbNewAdvhcRadiusTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbNewAdvhcRadiusTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbNewAdvhcRadiusTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbNewAdvhcRadiusID).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbNewAdvhcRadiusID)
}

type SlbNewAdvhcRadiusTableIPVer int32

const (
	SlbNewAdvhcRadiusTableIPVer_Ipv4 SlbNewAdvhcRadiusTableIPVer = 1
	SlbNewAdvhcRadiusTableIPVer_Ipv6 SlbNewAdvhcRadiusTableIPVer = 2
	SlbNewAdvhcRadiusTableIPVer_None SlbNewAdvhcRadiusTableIPVer = 3
)

type SlbNewAdvhcRadiusTableTransparent int32

const (
	SlbNewAdvhcRadiusTableTransparent_Enabled     SlbNewAdvhcRadiusTableTransparent = 1
	SlbNewAdvhcRadiusTableTransparent_Disabled    SlbNewAdvhcRadiusTableTransparent = 2
	SlbNewAdvhcRadiusTableTransparent_Unsupported SlbNewAdvhcRadiusTableTransparent = 2147483647
)

type SlbNewAdvhcRadiusTableInvert int32

const (
	SlbNewAdvhcRadiusTableInvert_Enabled     SlbNewAdvhcRadiusTableInvert = 1
	SlbNewAdvhcRadiusTableInvert_Disabled    SlbNewAdvhcRadiusTableInvert = 2
	SlbNewAdvhcRadiusTableInvert_Unsupported SlbNewAdvhcRadiusTableInvert = 2147483647
)

type SlbNewAdvhcRadiusTableDownType int32

const (
	SlbNewAdvhcRadiusTableDownType_Accounting     SlbNewAdvhcRadiusTableDownType = 1
	SlbNewAdvhcRadiusTableDownType_Authentication SlbNewAdvhcRadiusTableDownType = 2
	SlbNewAdvhcRadiusTableDownType_Any            SlbNewAdvhcRadiusTableDownType = 3
	SlbNewAdvhcRadiusTableDownType_Unsupported    SlbNewAdvhcRadiusTableDownType = 2147483647
)

type SlbNewAdvhcRadiusTableDelete int32

const (
	SlbNewAdvhcRadiusTableDelete_Other       SlbNewAdvhcRadiusTableDelete = 1
	SlbNewAdvhcRadiusTableDelete_Delete      SlbNewAdvhcRadiusTableDelete = 2
	SlbNewAdvhcRadiusTableDelete_Unsupported SlbNewAdvhcRadiusTableDelete = 2147483647
)

type SlbNewAdvhcRadiusTableSnat int32

const (
	SlbNewAdvhcRadiusTableSnat_Enabled     SlbNewAdvhcRadiusTableSnat = 1
	SlbNewAdvhcRadiusTableSnat_Disabled    SlbNewAdvhcRadiusTableSnat = 2
	SlbNewAdvhcRadiusTableSnat_Unsupported SlbNewAdvhcRadiusTableSnat = 2147483647
)

type SlbNewAdvhcRadiusTableParams struct {
	// RADIUS health check id.
	ID string `json:"ID,omitempty"`
	// RADIUS health check name.
	Name string `json:"Name,omitempty"`
	// RADIUS health check destination port.
	DPort uint64 `json:"DPort,omitempty"`
	// RADIUS health check destination IP version.
	IPVer SlbNewAdvhcRadiusTableIPVer `json:"IPVer,omitempty"`
	// RADIUS health check destination hostname.
	HostName string `json:"HostName,omitempty"`
	// RADIUS health check transparent flag.
	Transparent SlbNewAdvhcRadiusTableTransparent `json:"Transparent,omitempty"`
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
	Invert SlbNewAdvhcRadiusTableInvert `json:"Invert,omitempty"`
	// RADIUS health check type.
	DownType SlbNewAdvhcRadiusTableDownType `json:"DownType,omitempty"`
	// RADIUS health check user name.
	UserName string `json:"UserName,omitempty"`
	// RADIUS health check password.
	Password string `json:"Password,omitempty"`
	// RADIUS health check secret.
	Secret string `json:"Secret,omitempty"`
	// RADIUS health check copy indicator.
	Copy string `json:"Copy,omitempty"`
	// When set to the value of 2 (delete), the entire row is deleted.
	// When read, other(1) is returned. Setting the value to anything
	// other than 2(delete) has no effect on the state of the row.
	Delete SlbNewAdvhcRadiusTableDelete `json:"Delete,omitempty"`
	// RADIUS Health Check src NAT (PIP) flag.
	Snat SlbNewAdvhcRadiusTableSnat `json:"Snat,omitempty"`
}

func (p SlbNewAdvhcRadiusTableParams) iMABean() {}
