package beans

import (
	"fmt"
	"reflect"
)

// SlbNewAdvhcDhcpTable Note:This mib is not supported on VX instance of Virtualization.
type SlbNewAdvhcDhcpTable struct {
	// DHCP Health check id.
	SlbNewAdvhcDhcpID string
	Params            *SlbNewAdvhcDhcpTableParams
}

func NewSlbNewAdvhcDhcpTableList() *SlbNewAdvhcDhcpTable {
	return &SlbNewAdvhcDhcpTable{}
}

func NewSlbNewAdvhcDhcpTable(
	slbNewAdvhcDhcpID string,
	params *SlbNewAdvhcDhcpTableParams,
) *SlbNewAdvhcDhcpTable {
	return &SlbNewAdvhcDhcpTable{
		SlbNewAdvhcDhcpID: slbNewAdvhcDhcpID,
		Params:            params,
	}
}

func (c *SlbNewAdvhcDhcpTable) Name() string {
	return "SlbNewAdvhcDhcpTable"
}

func (c *SlbNewAdvhcDhcpTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbNewAdvhcDhcpTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbNewAdvhcDhcpTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbNewAdvhcDhcpID).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbNewAdvhcDhcpID)
}

type SlbNewAdvhcDhcpTableIPVer int32

const (
	SlbNewAdvhcDhcpTableIPVer_Ipv4 SlbNewAdvhcDhcpTableIPVer = 1
	SlbNewAdvhcDhcpTableIPVer_Ipv6 SlbNewAdvhcDhcpTableIPVer = 2
	SlbNewAdvhcDhcpTableIPVer_None SlbNewAdvhcDhcpTableIPVer = 3
)

type SlbNewAdvhcDhcpTableTransparent int32

const (
	SlbNewAdvhcDhcpTableTransparent_Enabled     SlbNewAdvhcDhcpTableTransparent = 1
	SlbNewAdvhcDhcpTableTransparent_Disabled    SlbNewAdvhcDhcpTableTransparent = 2
	SlbNewAdvhcDhcpTableTransparent_Unsupported SlbNewAdvhcDhcpTableTransparent = 2147483647
)

type SlbNewAdvhcDhcpTableInvert int32

const (
	SlbNewAdvhcDhcpTableInvert_Enabled     SlbNewAdvhcDhcpTableInvert = 1
	SlbNewAdvhcDhcpTableInvert_Disabled    SlbNewAdvhcDhcpTableInvert = 2
	SlbNewAdvhcDhcpTableInvert_Unsupported SlbNewAdvhcDhcpTableInvert = 2147483647
)

type SlbNewAdvhcDhcpTableDhcpType int32

const (
	SlbNewAdvhcDhcpTableDhcpType_Inherit     SlbNewAdvhcDhcpTableDhcpType = 1
	SlbNewAdvhcDhcpTableDhcpType_Inform      SlbNewAdvhcDhcpTableDhcpType = 2
	SlbNewAdvhcDhcpTableDhcpType_Request     SlbNewAdvhcDhcpTableDhcpType = 3
	SlbNewAdvhcDhcpTableDhcpType_Relay       SlbNewAdvhcDhcpTableDhcpType = 4
	SlbNewAdvhcDhcpTableDhcpType_Unsupported SlbNewAdvhcDhcpTableDhcpType = 2147483647
)

type SlbNewAdvhcDhcpTableSourcePortType int32

const (
	SlbNewAdvhcDhcpTableSourcePortType_Inherit     SlbNewAdvhcDhcpTableSourcePortType = 1
	SlbNewAdvhcDhcpTableSourcePortType_Random      SlbNewAdvhcDhcpTableSourcePortType = 2
	SlbNewAdvhcDhcpTableSourcePortType_Strict      SlbNewAdvhcDhcpTableSourcePortType = 3
	SlbNewAdvhcDhcpTableSourcePortType_Unsupported SlbNewAdvhcDhcpTableSourcePortType = 2147483647
)

type SlbNewAdvhcDhcpTableDelete int32

const (
	SlbNewAdvhcDhcpTableDelete_Other       SlbNewAdvhcDhcpTableDelete = 1
	SlbNewAdvhcDhcpTableDelete_Delete      SlbNewAdvhcDhcpTableDelete = 2
	SlbNewAdvhcDhcpTableDelete_Unsupported SlbNewAdvhcDhcpTableDelete = 2147483647
)

type SlbNewAdvhcDhcpTableDhcpDuidType int32

const (
	SlbNewAdvhcDhcpTableDhcpDuidType_Linklayertime SlbNewAdvhcDhcpTableDhcpDuidType = 1
	SlbNewAdvhcDhcpTableDhcpDuidType_Linklayer     SlbNewAdvhcDhcpTableDhcpDuidType = 3
	SlbNewAdvhcDhcpTableDhcpDuidType_Unsupported   SlbNewAdvhcDhcpTableDhcpDuidType = 2147483647
)

type SlbNewAdvhcDhcpTableSnat int32

const (
	SlbNewAdvhcDhcpTableSnat_Enabled     SlbNewAdvhcDhcpTableSnat = 1
	SlbNewAdvhcDhcpTableSnat_Disabled    SlbNewAdvhcDhcpTableSnat = 2
	SlbNewAdvhcDhcpTableSnat_Unsupported SlbNewAdvhcDhcpTableSnat = 2147483647
)

type SlbNewAdvhcDhcpTableParams struct {
	// DHCP Health check id.
	ID string `json:"ID,omitempty"`
	// DHCP Health check name.
	Name string `json:"Name,omitempty"`
	// DHCP Health check port.
	DPort uint64 `json:"DPort,omitempty"`
	// DHCP Health check IP version.
	IPVer SlbNewAdvhcDhcpTableIPVer `json:"IPVer,omitempty"`
	// DHCP Health check destination hostname.
	HostName string `json:"HostName,omitempty"`
	// DHCP Health check transparent flag.
	Transparent SlbNewAdvhcDhcpTableTransparent `json:"Transparent,omitempty"`
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
	Invert SlbNewAdvhcDhcpTableInvert `json:"Invert,omitempty"`
	// DHCP Health check packet type.
	DhcpType SlbNewAdvhcDhcpTableDhcpType `json:"DhcpType,omitempty"`
	// DHCP Health check source port selection method.
	SourcePortType SlbNewAdvhcDhcpTableSourcePortType `json:"SourcePortType,omitempty"`
	// Copy DHCP Health check.
	Copy string `json:"Copy,omitempty"`
	// When set to the value of 2 (delete), the entire row is deleted.
	// When read, other(1) is returned. Setting the value to anything
	// other than 2(delete) has no effect on the state of the row.
	Delete SlbNewAdvhcDhcpTableDelete `json:"Delete,omitempty"`
	// DHCP server's DUID Type. Only 1 and 3 are supported.
	DhcpDuidType SlbNewAdvhcDhcpTableDhcpDuidType `json:"DhcpDuidType,omitempty"`
	// DHCP server's Link-Layer address, e.g. e0:69:95:c8:ca:9d or e06995c8ca9d
	DhcpDuidLLA string `json:"DhcpDuidLLA,omitempty"`
	// DHCP server's DUID time, e.g. 500058203.
	DhcpDuidTime uint32 `json:"DhcpDuidTime,omitempty"`
	// DHCP Health Check src NAT (PIP) flag.
	Snat SlbNewAdvhcDhcpTableSnat `json:"Snat,omitempty"`
}

func (p SlbNewAdvhcDhcpTableParams) iMABean() {}
