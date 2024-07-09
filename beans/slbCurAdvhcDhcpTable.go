package beans

import (
	"fmt"
	"reflect"
)

// SlbCurAdvhcDhcpTable Note:This mib is not supported on VX instance of Virtualization.
type SlbCurAdvhcDhcpTable struct {
	// DHCP Health check id.
	SlbCurAdvhcDhcpID string
	Params            *SlbCurAdvhcDhcpTableParams
}

func NewSlbCurAdvhcDhcpTableList() *SlbCurAdvhcDhcpTable {
	return &SlbCurAdvhcDhcpTable{}
}

func NewSlbCurAdvhcDhcpTable(
	slbCurAdvhcDhcpID string,
	params *SlbCurAdvhcDhcpTableParams,
) *SlbCurAdvhcDhcpTable {
	return &SlbCurAdvhcDhcpTable{
		SlbCurAdvhcDhcpID: slbCurAdvhcDhcpID,
		Params:            params,
	}
}

func (c *SlbCurAdvhcDhcpTable) Name() string {
	return "SlbCurAdvhcDhcpTable"
}

func (c *SlbCurAdvhcDhcpTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbCurAdvhcDhcpTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbCurAdvhcDhcpTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbCurAdvhcDhcpID).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbCurAdvhcDhcpID)
}

type SlbCurAdvhcDhcpTableIPVer int32

const (
	SlbCurAdvhcDhcpTableIPVer_Ipv4 SlbCurAdvhcDhcpTableIPVer = 1
	SlbCurAdvhcDhcpTableIPVer_Ipv6 SlbCurAdvhcDhcpTableIPVer = 2
	SlbCurAdvhcDhcpTableIPVer_None SlbCurAdvhcDhcpTableIPVer = 3
)

type SlbCurAdvhcDhcpTableTransparent int32

const (
	SlbCurAdvhcDhcpTableTransparent_Enabled     SlbCurAdvhcDhcpTableTransparent = 1
	SlbCurAdvhcDhcpTableTransparent_Disabled    SlbCurAdvhcDhcpTableTransparent = 2
	SlbCurAdvhcDhcpTableTransparent_Unsupported SlbCurAdvhcDhcpTableTransparent = 2147483647
)

type SlbCurAdvhcDhcpTableInvert int32

const (
	SlbCurAdvhcDhcpTableInvert_Enabled     SlbCurAdvhcDhcpTableInvert = 1
	SlbCurAdvhcDhcpTableInvert_Disabled    SlbCurAdvhcDhcpTableInvert = 2
	SlbCurAdvhcDhcpTableInvert_Unsupported SlbCurAdvhcDhcpTableInvert = 2147483647
)

type SlbCurAdvhcDhcpTableDhcpType int32

const (
	SlbCurAdvhcDhcpTableDhcpType_Inherit     SlbCurAdvhcDhcpTableDhcpType = 1
	SlbCurAdvhcDhcpTableDhcpType_Inform      SlbCurAdvhcDhcpTableDhcpType = 2
	SlbCurAdvhcDhcpTableDhcpType_Request     SlbCurAdvhcDhcpTableDhcpType = 3
	SlbCurAdvhcDhcpTableDhcpType_Relay       SlbCurAdvhcDhcpTableDhcpType = 4
	SlbCurAdvhcDhcpTableDhcpType_Unsupported SlbCurAdvhcDhcpTableDhcpType = 2147483647
)

type SlbCurAdvhcDhcpTableSourcePortType int32

const (
	SlbCurAdvhcDhcpTableSourcePortType_Inherit     SlbCurAdvhcDhcpTableSourcePortType = 1
	SlbCurAdvhcDhcpTableSourcePortType_Random      SlbCurAdvhcDhcpTableSourcePortType = 2
	SlbCurAdvhcDhcpTableSourcePortType_Strict      SlbCurAdvhcDhcpTableSourcePortType = 3
	SlbCurAdvhcDhcpTableSourcePortType_Unsupported SlbCurAdvhcDhcpTableSourcePortType = 2147483647
)

type SlbCurAdvhcDhcpTableDhcpDuidType int32

const (
	SlbCurAdvhcDhcpTableDhcpDuidType_Linklayertime SlbCurAdvhcDhcpTableDhcpDuidType = 1
	SlbCurAdvhcDhcpTableDhcpDuidType_Linklayer     SlbCurAdvhcDhcpTableDhcpDuidType = 3
	SlbCurAdvhcDhcpTableDhcpDuidType_Unsupported   SlbCurAdvhcDhcpTableDhcpDuidType = 2147483647
)

type SlbCurAdvhcDhcpTableSnat int32

const (
	SlbCurAdvhcDhcpTableSnat_Enabled     SlbCurAdvhcDhcpTableSnat = 1
	SlbCurAdvhcDhcpTableSnat_Disabled    SlbCurAdvhcDhcpTableSnat = 2
	SlbCurAdvhcDhcpTableSnat_Unsupported SlbCurAdvhcDhcpTableSnat = 2147483647
)

type SlbCurAdvhcDhcpTableParams struct {
	// DHCP Health check id.
	ID string `json:"ID,omitempty"`
	// DHCP Health check name.
	Name string `json:"Name,omitempty"`
	// DHCP Health check port.
	DPort uint64 `json:"DPort,omitempty"`
	// DHCP Health check IP version.
	IPVer SlbCurAdvhcDhcpTableIPVer `json:"IPVer,omitempty"`
	// DHCP Health check destination hostname.
	HostName string `json:"HostName,omitempty"`
	// DHCP Health check transparent flag.
	Transparent SlbCurAdvhcDhcpTableTransparent `json:"Transparent,omitempty"`
	// DHCP Health check interval.
	Interval uint64 `json:"Interval,omitempty"`
	// DHCP Health check retries counter.
	Retries uint64 `json:"Retries,omitempty"`
	// DHCP Health check restore from down state counter.
	RestoreRetries uint64 `json:"RestoreRetries,omitempty"`
	// DHCP Health check timeout period.
	Timeout uint64 `json:"Timeout,omitempty"`
	// DHCP Health check overflow flag.
	Overflow uint64 `json:"Overflow,omitempty"`
	// DHCP Health check interval in down state.
	DownInterval uint64 `json:"DownInterval,omitempty"`
	// DHCP Health check invert flag.
	Invert SlbCurAdvhcDhcpTableInvert `json:"Invert,omitempty"`
	// DHCP Health check packet type.
	DhcpType SlbCurAdvhcDhcpTableDhcpType `json:"DhcpType,omitempty"`
	// DHCP Health check source port selection method.
	SourcePortType SlbCurAdvhcDhcpTableSourcePortType `json:"SourcePortType,omitempty"`
	// DHCP server's DUID Type. Only 1 and 3 are supported.
	DhcpDuidType SlbCurAdvhcDhcpTableDhcpDuidType `json:"DhcpDuidType,omitempty"`
	// DHCP server's Link-Layer address.
	DhcpDuidLLA string `json:"DhcpDuidLLA,omitempty"`
	// DHCP server's DUID time, e.g. 500058203.
	DhcpDuidTime uint32 `json:"DhcpDuidTime,omitempty"`
	// DHCP Health Check src NAT (PIP) flag.
	Snat SlbCurAdvhcDhcpTableSnat `json:"Snat,omitempty"`
}

func (p SlbCurAdvhcDhcpTableParams) iMABean() {}
