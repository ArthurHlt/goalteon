package beans

import (
	"fmt"
	"reflect"
)

// SlbNewAdvhcWtsTable Note:This mib is not supported on VX instance of Virtualization.
type SlbNewAdvhcWtsTable struct {
	// WTS Health check id.
	SlbNewAdvhcWtsID string
	Params           *SlbNewAdvhcWtsTableParams
}

func NewSlbNewAdvhcWtsTableList() *SlbNewAdvhcWtsTable {
	return &SlbNewAdvhcWtsTable{}
}

func NewSlbNewAdvhcWtsTable(
	slbNewAdvhcWtsID string,
	params *SlbNewAdvhcWtsTableParams,
) *SlbNewAdvhcWtsTable {
	return &SlbNewAdvhcWtsTable{
		SlbNewAdvhcWtsID: slbNewAdvhcWtsID,
		Params:           params,
	}
}

func (c *SlbNewAdvhcWtsTable) Name() string {
	return "SlbNewAdvhcWtsTable"
}

func (c *SlbNewAdvhcWtsTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbNewAdvhcWtsTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbNewAdvhcWtsTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbNewAdvhcWtsID).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbNewAdvhcWtsID)
}

type SlbNewAdvhcWtsTableIPVer int32

const (
	SlbNewAdvhcWtsTableIPVer_Ipv4 SlbNewAdvhcWtsTableIPVer = 1
	SlbNewAdvhcWtsTableIPVer_Ipv6 SlbNewAdvhcWtsTableIPVer = 2
	SlbNewAdvhcWtsTableIPVer_None SlbNewAdvhcWtsTableIPVer = 3
)

type SlbNewAdvhcWtsTableTransparent int32

const (
	SlbNewAdvhcWtsTableTransparent_Enabled     SlbNewAdvhcWtsTableTransparent = 1
	SlbNewAdvhcWtsTableTransparent_Disabled    SlbNewAdvhcWtsTableTransparent = 2
	SlbNewAdvhcWtsTableTransparent_Unsupported SlbNewAdvhcWtsTableTransparent = 2147483647
)

type SlbNewAdvhcWtsTableInvert int32

const (
	SlbNewAdvhcWtsTableInvert_Enabled     SlbNewAdvhcWtsTableInvert = 1
	SlbNewAdvhcWtsTableInvert_Disabled    SlbNewAdvhcWtsTableInvert = 2
	SlbNewAdvhcWtsTableInvert_Unsupported SlbNewAdvhcWtsTableInvert = 2147483647
)

type SlbNewAdvhcWtsTableDelete int32

const (
	SlbNewAdvhcWtsTableDelete_Other       SlbNewAdvhcWtsTableDelete = 1
	SlbNewAdvhcWtsTableDelete_Delete      SlbNewAdvhcWtsTableDelete = 2
	SlbNewAdvhcWtsTableDelete_Unsupported SlbNewAdvhcWtsTableDelete = 2147483647
)

type SlbNewAdvhcWtsTableSnat int32

const (
	SlbNewAdvhcWtsTableSnat_Enabled     SlbNewAdvhcWtsTableSnat = 1
	SlbNewAdvhcWtsTableSnat_Disabled    SlbNewAdvhcWtsTableSnat = 2
	SlbNewAdvhcWtsTableSnat_Unsupported SlbNewAdvhcWtsTableSnat = 2147483647
)

type SlbNewAdvhcWtsTableParams struct {
	// WTS Health check id.
	ID string `json:"ID,omitempty"`
	// WTS Health check name.
	Name string `json:"Name,omitempty"`
	// WTS Health check destination port.
	DPort uint64 `json:"DPort,omitempty"`
	// WTS Health check destination IP version.
	IPVer SlbNewAdvhcWtsTableIPVer `json:"IPVer,omitempty"`
	// WTS Health check destination hostname.
	HostName string `json:"HostName,omitempty"`
	// WTS Health check transparent flag.
	Transparent SlbNewAdvhcWtsTableTransparent `json:"Transparent,omitempty"`
	// WTS Health check interval.
	Interval uint64 `json:"Interval,omitempty"`
	// WTS Health check retries counter.
	Retries uint64 `json:"Retries,omitempty"`
	// WTS Health check retries when in the down state counter.
	RestoreRetries uint64 `json:"RestoreRetries,omitempty"`
	// WTS Health check timeout.
	Timeout uint64 `json:"Timeout,omitempty"`
	// WTS Health check timeout.
	Overflow uint64 `json:"Overflow,omitempty"`
	// WTS Health check interval when in the down state.
	DownInterval uint64 `json:"DownInterval,omitempty"`
	// WTS Health check invert flag.
	Invert SlbNewAdvhcWtsTableInvert `json:"Invert,omitempty"`
	// WTS Health check user name.
	UserName string `json:"UserName,omitempty"`
	// WTS Health check copy indicator.
	Copy DisplayString `json:"Copy,omitempty"`
	// When set to the value of 2 (delete), the entire row is deleted.
	// When read, other(1) is returned. Setting the value to anything
	// other than 2(delete) has no effect on the state of the row.
	Delete SlbNewAdvhcWtsTableDelete `json:"Delete,omitempty"`
	// WTS Health Check src NAT (PIP) flag.
	Snat SlbNewAdvhcWtsTableSnat `json:"Snat,omitempty"`
}

func (p SlbNewAdvhcWtsTableParams) iMABean() {}
