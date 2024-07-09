package beans

import (
	"fmt"
	"reflect"
)

// SlbNewAdvhcTftpTable TFTP Health check path.
type SlbNewAdvhcTftpTable struct {
	// TFTP Health check id.
	SlbNewAdvhcTftpID string
	Params            *SlbNewAdvhcTftpTableParams
}

func NewSlbNewAdvhcTftpTableList() *SlbNewAdvhcTftpTable {
	return &SlbNewAdvhcTftpTable{}
}

func NewSlbNewAdvhcTftpTable(
	slbNewAdvhcTftpID string,
	params *SlbNewAdvhcTftpTableParams,
) *SlbNewAdvhcTftpTable {
	return &SlbNewAdvhcTftpTable{
		SlbNewAdvhcTftpID: slbNewAdvhcTftpID,
		Params:            params,
	}
}

func (c *SlbNewAdvhcTftpTable) Name() string {
	return "SlbNewAdvhcTftpTable"
}

func (c *SlbNewAdvhcTftpTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbNewAdvhcTftpTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbNewAdvhcTftpTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbNewAdvhcTftpID).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbNewAdvhcTftpID)
}

type SlbNewAdvhcTftpTableIPVer int32

const (
	SlbNewAdvhcTftpTableIPVer_Ipv4 SlbNewAdvhcTftpTableIPVer = 1
	SlbNewAdvhcTftpTableIPVer_Ipv6 SlbNewAdvhcTftpTableIPVer = 2
	SlbNewAdvhcTftpTableIPVer_None SlbNewAdvhcTftpTableIPVer = 3
)

type SlbNewAdvhcTftpTableTransparent int32

const (
	SlbNewAdvhcTftpTableTransparent_Enabled     SlbNewAdvhcTftpTableTransparent = 1
	SlbNewAdvhcTftpTableTransparent_Disabled    SlbNewAdvhcTftpTableTransparent = 2
	SlbNewAdvhcTftpTableTransparent_Unsupported SlbNewAdvhcTftpTableTransparent = 2147483647
)

type SlbNewAdvhcTftpTableInvert int32

const (
	SlbNewAdvhcTftpTableInvert_Enabled     SlbNewAdvhcTftpTableInvert = 1
	SlbNewAdvhcTftpTableInvert_Disabled    SlbNewAdvhcTftpTableInvert = 2
	SlbNewAdvhcTftpTableInvert_Unsupported SlbNewAdvhcTftpTableInvert = 2147483647
)

type SlbNewAdvhcTftpTableDelete int32

const (
	SlbNewAdvhcTftpTableDelete_Other       SlbNewAdvhcTftpTableDelete = 1
	SlbNewAdvhcTftpTableDelete_Delete      SlbNewAdvhcTftpTableDelete = 2
	SlbNewAdvhcTftpTableDelete_Unsupported SlbNewAdvhcTftpTableDelete = 2147483647
)

type SlbNewAdvhcTftpTableSnat int32

const (
	SlbNewAdvhcTftpTableSnat_Enabled     SlbNewAdvhcTftpTableSnat = 1
	SlbNewAdvhcTftpTableSnat_Disabled    SlbNewAdvhcTftpTableSnat = 2
	SlbNewAdvhcTftpTableSnat_Unsupported SlbNewAdvhcTftpTableSnat = 2147483647
)

type SlbNewAdvhcTftpTableParams struct {
	// TFTP Health check id.
	ID string `json:"ID,omitempty"`
	// TFTP Health check name.
	Name string `json:"Name,omitempty"`
	// TFTP Health check destination port.
	DPort uint64 `json:"DPort,omitempty"`
	// TFTP Health check destination IP version.
	IPVer SlbNewAdvhcTftpTableIPVer `json:"IPVer,omitempty"`
	// TFTP Health check host name.
	HostName string `json:"HostName,omitempty"`
	// TFTP Health check transparent flag.
	Transparent SlbNewAdvhcTftpTableTransparent `json:"Transparent,omitempty"`
	// TFTP Health check interval.
	Interval uint64 `json:"Interval,omitempty"`
	// TFTP Health check retries counter.
	Retries uint64 `json:"Retries,omitempty"`
	// TFTP Health check retries counter in the down state.
	RestoreRetries uint64 `json:"RestoreRetries,omitempty"`
	// TFTP Health check timeout.
	Timeout uint64 `json:"Timeout,omitempty"`
	// TFTP Health check overflow flag.
	Overflow uint64 `json:"Overflow,omitempty"`
	// TFTP Health check interval in the down state.
	DownInterval uint64 `json:"DownInterval,omitempty"`
	// TFTP Health check invert flag.
	Invert SlbNewAdvhcTftpTableInvert `json:"Invert,omitempty"`
	// TFTP Health check path.
	TftpfileFullPath string `json:"TftpfileFullPath,omitempty"`
	// TFTP Health check copy indicator.
	Copy string `json:"Copy,omitempty"`
	// When set to the value of 2 (delete), the entire row is deleted.
	// When read, other(1) is returned. Setting the value to anything
	// other than 2(delete) has no effect on the state of the row.
	Delete SlbNewAdvhcTftpTableDelete `json:"Delete,omitempty"`
	// TFTP Health Check src NAT (PIP) flag.
	Snat SlbNewAdvhcTftpTableSnat `json:"Snat,omitempty"`
}

func (p SlbNewAdvhcTftpTableParams) iMABean() {}
