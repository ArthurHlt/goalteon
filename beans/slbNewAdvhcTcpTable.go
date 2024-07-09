package beans

import (
	"fmt"
	"reflect"
)

// SlbNewAdvhcTcpTable Note:This mib is not supported on VX instance of Virtualization.
type SlbNewAdvhcTcpTable struct {
	// TCP Health check id.
	SlbNewAdvhcTcpID string
	Params           *SlbNewAdvhcTcpTableParams
}

func NewSlbNewAdvhcTcpTableList() *SlbNewAdvhcTcpTable {
	return &SlbNewAdvhcTcpTable{}
}

func NewSlbNewAdvhcTcpTable(
	slbNewAdvhcTcpID string,
	params *SlbNewAdvhcTcpTableParams,
) *SlbNewAdvhcTcpTable {
	return &SlbNewAdvhcTcpTable{
		SlbNewAdvhcTcpID: slbNewAdvhcTcpID,
		Params:           params,
	}
}

func (c *SlbNewAdvhcTcpTable) Name() string {
	return "SlbNewAdvhcTcpTable"
}

func (c *SlbNewAdvhcTcpTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbNewAdvhcTcpTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbNewAdvhcTcpTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbNewAdvhcTcpID).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbNewAdvhcTcpID)
}

type SlbNewAdvhcTcpTableIPVer int32

const (
	SlbNewAdvhcTcpTableIPVer_Ipv4 SlbNewAdvhcTcpTableIPVer = 1
	SlbNewAdvhcTcpTableIPVer_Ipv6 SlbNewAdvhcTcpTableIPVer = 2
	SlbNewAdvhcTcpTableIPVer_None SlbNewAdvhcTcpTableIPVer = 3
)

type SlbNewAdvhcTcpTableTransparent int32

const (
	SlbNewAdvhcTcpTableTransparent_Enabled     SlbNewAdvhcTcpTableTransparent = 1
	SlbNewAdvhcTcpTableTransparent_Disabled    SlbNewAdvhcTcpTableTransparent = 2
	SlbNewAdvhcTcpTableTransparent_Unsupported SlbNewAdvhcTcpTableTransparent = 2147483647
)

type SlbNewAdvhcTcpTableInvert int32

const (
	SlbNewAdvhcTcpTableInvert_Enabled     SlbNewAdvhcTcpTableInvert = 1
	SlbNewAdvhcTcpTableInvert_Disabled    SlbNewAdvhcTcpTableInvert = 2
	SlbNewAdvhcTcpTableInvert_Unsupported SlbNewAdvhcTcpTableInvert = 2147483647
)

type SlbNewAdvhcTcpTableDelete int32

const (
	SlbNewAdvhcTcpTableDelete_Other       SlbNewAdvhcTcpTableDelete = 1
	SlbNewAdvhcTcpTableDelete_Delete      SlbNewAdvhcTcpTableDelete = 2
	SlbNewAdvhcTcpTableDelete_Unsupported SlbNewAdvhcTcpTableDelete = 2147483647
)

type SlbNewAdvhcTcpTableConnTerm int32

const (
	SlbNewAdvhcTcpTableConnTerm_Fin         SlbNewAdvhcTcpTableConnTerm = 1
	SlbNewAdvhcTcpTableConnTerm_Rst         SlbNewAdvhcTcpTableConnTerm = 2
	SlbNewAdvhcTcpTableConnTerm_Unsupported SlbNewAdvhcTcpTableConnTerm = 2147483647
)

type SlbNewAdvhcTcpTableAlways int32

const (
	SlbNewAdvhcTcpTableAlways_Enabled     SlbNewAdvhcTcpTableAlways = 1
	SlbNewAdvhcTcpTableAlways_Disabled    SlbNewAdvhcTcpTableAlways = 2
	SlbNewAdvhcTcpTableAlways_Unsupported SlbNewAdvhcTcpTableAlways = 2147483647
)

type SlbNewAdvhcTcpTableSnat int32

const (
	SlbNewAdvhcTcpTableSnat_Enabled     SlbNewAdvhcTcpTableSnat = 1
	SlbNewAdvhcTcpTableSnat_Disabled    SlbNewAdvhcTcpTableSnat = 2
	SlbNewAdvhcTcpTableSnat_Unsupported SlbNewAdvhcTcpTableSnat = 2147483647
)

type SlbNewAdvhcTcpTableParams struct {
	// TCP Health check id.
	ID string `json:"ID,omitempty"`
	// TCP Health check name.
	Name string `json:"Name,omitempty"`
	// TCP Health check destination port.
	DPort uint64 `json:"DPort,omitempty"`
	// TCP Health check destination IP version.
	IPVer SlbNewAdvhcTcpTableIPVer `json:"IPVer,omitempty"`
	// TCP Health check destination host name.
	HostName string `json:"HostName,omitempty"`
	// TCP Health check transparent flag.
	Transparent SlbNewAdvhcTcpTableTransparent `json:"Transparent,omitempty"`
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
	Invert SlbNewAdvhcTcpTableInvert `json:"Invert,omitempty"`
	// TCP Health check copy indicator.
	Copy string `json:"Copy,omitempty"`
	// When set to the value of 2 (delete), the entire row is deleted.
	// When read, other(1) is returned. Setting the value to anything
	// other than 2(delete) has no effect on the state of the row.
	Delete SlbNewAdvhcTcpTableDelete `json:"Delete,omitempty"`
	// Connection termination type.
	ConnTerm SlbNewAdvhcTcpTableConnTerm `json:"ConnTerm,omitempty"`
	// This flag determines whether HC is allowed for standalone real.
	Always SlbNewAdvhcTcpTableAlways `json:"Always,omitempty"`
	// TCP Health Check src NAT (PIP) flag.
	Snat SlbNewAdvhcTcpTableSnat `json:"Snat,omitempty"`
}

func (p SlbNewAdvhcTcpTableParams) iMABean() {}
