package beans

import (
	"fmt"
	"reflect"
)

// SlbCurAdvhcWtsTable Note:This mib is not supported on VX instance of Virtualization.
type SlbCurAdvhcWtsTable struct {
	// WTS Health check id.
	SlbCurAdvhcWtsID string
	Params           *SlbCurAdvhcWtsTableParams
}

func NewSlbCurAdvhcWtsTableList() *SlbCurAdvhcWtsTable {
	return &SlbCurAdvhcWtsTable{}
}

func NewSlbCurAdvhcWtsTable(
	slbCurAdvhcWtsID string,
	params *SlbCurAdvhcWtsTableParams,
) *SlbCurAdvhcWtsTable {
	return &SlbCurAdvhcWtsTable{
		SlbCurAdvhcWtsID: slbCurAdvhcWtsID,
		Params:           params,
	}
}

func (c *SlbCurAdvhcWtsTable) Name() string {
	return "SlbCurAdvhcWtsTable"
}

func (c *SlbCurAdvhcWtsTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbCurAdvhcWtsTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbCurAdvhcWtsTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbCurAdvhcWtsID).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbCurAdvhcWtsID)
}

type SlbCurAdvhcWtsTableIPVer int32

const (
	SlbCurAdvhcWtsTableIPVer_Ipv4 SlbCurAdvhcWtsTableIPVer = 1
	SlbCurAdvhcWtsTableIPVer_Ipv6 SlbCurAdvhcWtsTableIPVer = 2
	SlbCurAdvhcWtsTableIPVer_None SlbCurAdvhcWtsTableIPVer = 3
)

type SlbCurAdvhcWtsTableTransparent int32

const (
	SlbCurAdvhcWtsTableTransparent_Enabled     SlbCurAdvhcWtsTableTransparent = 1
	SlbCurAdvhcWtsTableTransparent_Disabled    SlbCurAdvhcWtsTableTransparent = 2
	SlbCurAdvhcWtsTableTransparent_Unsupported SlbCurAdvhcWtsTableTransparent = 2147483647
)

type SlbCurAdvhcWtsTableInvert int32

const (
	SlbCurAdvhcWtsTableInvert_Enabled     SlbCurAdvhcWtsTableInvert = 1
	SlbCurAdvhcWtsTableInvert_Disabled    SlbCurAdvhcWtsTableInvert = 2
	SlbCurAdvhcWtsTableInvert_Unsupported SlbCurAdvhcWtsTableInvert = 2147483647
)

type SlbCurAdvhcWtsTableParams struct {
	// WTS Health check id.
	ID string `json:"ID,omitempty"`
	// WTS Health check name.
	Name string `json:"Name,omitempty"`
	// WTS Health check destination port.
	DPort uint64 `json:"DPort,omitempty"`
	// WTS Health check destination IP version.
	IPVer SlbCurAdvhcWtsTableIPVer `json:"IPVer,omitempty"`
	// WTS Health check destination hostname.
	HostName string `json:"HostName,omitempty"`
	// WTS Health check transparent flag.
	Transparent SlbCurAdvhcWtsTableTransparent `json:"Transparent,omitempty"`
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
	Invert SlbCurAdvhcWtsTableInvert `json:"Invert,omitempty"`
	// WTS Health check user name.
	UserName string `json:"UserName,omitempty"`
}

func (p SlbCurAdvhcWtsTableParams) iMABean() {}
