package beans

import (
	"fmt"
	"reflect"
)

// SlbCurAdvhcTcpTable Note:This mib is not supported on VX instance of Virtualization.
type SlbCurAdvhcTcpTable struct {
	// TCP Health check id.
	SlbCurAdvhcTcpID string
	Params           *SlbCurAdvhcTcpTableParams
}

func NewSlbCurAdvhcTcpTableList() *SlbCurAdvhcTcpTable {
	return &SlbCurAdvhcTcpTable{}
}

func NewSlbCurAdvhcTcpTable(
	slbCurAdvhcTcpID string,
	params *SlbCurAdvhcTcpTableParams,
) *SlbCurAdvhcTcpTable {
	return &SlbCurAdvhcTcpTable{
		SlbCurAdvhcTcpID: slbCurAdvhcTcpID,
		Params:           params,
	}
}

func (c *SlbCurAdvhcTcpTable) Name() string {
	return "SlbCurAdvhcTcpTable"
}

func (c *SlbCurAdvhcTcpTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbCurAdvhcTcpTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbCurAdvhcTcpTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbCurAdvhcTcpID).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbCurAdvhcTcpID)
}

type SlbCurAdvhcTcpTableIPVer int32

const (
	SlbCurAdvhcTcpTableIPVer_Ipv4 SlbCurAdvhcTcpTableIPVer = 1
	SlbCurAdvhcTcpTableIPVer_Ipv6 SlbCurAdvhcTcpTableIPVer = 2
	SlbCurAdvhcTcpTableIPVer_None SlbCurAdvhcTcpTableIPVer = 3
)

type SlbCurAdvhcTcpTableTransparent int32

const (
	SlbCurAdvhcTcpTableTransparent_Enabled     SlbCurAdvhcTcpTableTransparent = 1
	SlbCurAdvhcTcpTableTransparent_Disabled    SlbCurAdvhcTcpTableTransparent = 2
	SlbCurAdvhcTcpTableTransparent_Unsupported SlbCurAdvhcTcpTableTransparent = 2147483647
)

type SlbCurAdvhcTcpTableInvert int32

const (
	SlbCurAdvhcTcpTableInvert_Enabled     SlbCurAdvhcTcpTableInvert = 1
	SlbCurAdvhcTcpTableInvert_Disabled    SlbCurAdvhcTcpTableInvert = 2
	SlbCurAdvhcTcpTableInvert_Unsupported SlbCurAdvhcTcpTableInvert = 2147483647
)

type SlbCurAdvhcTcpTableConnTerm int32

const (
	SlbCurAdvhcTcpTableConnTerm_Fin         SlbCurAdvhcTcpTableConnTerm = 1
	SlbCurAdvhcTcpTableConnTerm_Rst         SlbCurAdvhcTcpTableConnTerm = 2
	SlbCurAdvhcTcpTableConnTerm_Unsupported SlbCurAdvhcTcpTableConnTerm = 2147483647
)

type SlbCurAdvhcTcpTableAlways int32

const (
	SlbCurAdvhcTcpTableAlways_Enabled     SlbCurAdvhcTcpTableAlways = 1
	SlbCurAdvhcTcpTableAlways_Disabled    SlbCurAdvhcTcpTableAlways = 2
	SlbCurAdvhcTcpTableAlways_Unsupported SlbCurAdvhcTcpTableAlways = 2147483647
)

type SlbCurAdvhcTcpTableParams struct {
	// TCP Health check id.
	ID string `json:"ID,omitempty"`
	// TCP Health check name.
	Name string `json:"Name,omitempty"`
	// TCP Health check destination port.
	DPort uint64 `json:"DPort,omitempty"`
	// TCP Health check destination IP version.
	IPVer SlbCurAdvhcTcpTableIPVer `json:"IPVer,omitempty"`
	// TCP Health check destination host name.
	HostName string `json:"HostName,omitempty"`
	// TCP Health check transparent flag.
	Transparent SlbCurAdvhcTcpTableTransparent `json:"Transparent,omitempty"`
	// TCP Health check interval.
	Interval uint64 `json:"Interval,omitempty"`
	// TCP Health check retries counter.
	Retries uint64 `json:"Retries,omitempty"`
	// TCP Health check retries counter in the down state.
	RestoreRetries uint64 `json:"RestoreRetries,omitempty"`
	// TCP Health check timeout.
	Timeout uint64 `json:"Timeout,omitempty"`
	// TCP Health check overflow flag.
	Overflow uint64 `json:"Overflow,omitempty"`
	// TCP Health check interval in the down state.
	DownInterval uint64 `json:"DownInterval,omitempty"`
	// TCP Health check invert flag.
	Invert SlbCurAdvhcTcpTableInvert `json:"Invert,omitempty"`
	// Connection termination type
	ConnTerm SlbCurAdvhcTcpTableConnTerm `json:"ConnTerm,omitempty"`
	// This flag determines whether HC is allowed for standalone real.
	Always SlbCurAdvhcTcpTableAlways `json:"Always,omitempty"`
}

func (p SlbCurAdvhcTcpTableParams) iMABean() {}
