package beans

import (
	"fmt"
	"reflect"
)

// SlbCurAdvhcTftpTable Note:This mib is not supported on VX instance of Virtualization.
type SlbCurAdvhcTftpTable struct {
	// TFTP Health check id.
	SlbCurAdvhcTftpID string
	Params            *SlbCurAdvhcTftpTableParams
}

func NewSlbCurAdvhcTftpTableList() *SlbCurAdvhcTftpTable {
	return &SlbCurAdvhcTftpTable{}
}

func NewSlbCurAdvhcTftpTable(
	slbCurAdvhcTftpID string,
	params *SlbCurAdvhcTftpTableParams,
) *SlbCurAdvhcTftpTable {
	return &SlbCurAdvhcTftpTable{
		SlbCurAdvhcTftpID: slbCurAdvhcTftpID,
		Params:            params,
	}
}

func (c *SlbCurAdvhcTftpTable) Name() string {
	return "SlbCurAdvhcTftpTable"
}

func (c *SlbCurAdvhcTftpTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbCurAdvhcTftpTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbCurAdvhcTftpTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbCurAdvhcTftpID).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbCurAdvhcTftpID)
}

type SlbCurAdvhcTftpTableIPVer int32

const (
	SlbCurAdvhcTftpTableIPVer_Ipv4 SlbCurAdvhcTftpTableIPVer = 1
	SlbCurAdvhcTftpTableIPVer_Ipv6 SlbCurAdvhcTftpTableIPVer = 2
	SlbCurAdvhcTftpTableIPVer_None SlbCurAdvhcTftpTableIPVer = 3
)

type SlbCurAdvhcTftpTableTransparent int32

const (
	SlbCurAdvhcTftpTableTransparent_Enabled     SlbCurAdvhcTftpTableTransparent = 1
	SlbCurAdvhcTftpTableTransparent_Disabled    SlbCurAdvhcTftpTableTransparent = 2
	SlbCurAdvhcTftpTableTransparent_Unsupported SlbCurAdvhcTftpTableTransparent = 2147483647
)

type SlbCurAdvhcTftpTableInvert int32

const (
	SlbCurAdvhcTftpTableInvert_Enabled     SlbCurAdvhcTftpTableInvert = 1
	SlbCurAdvhcTftpTableInvert_Disabled    SlbCurAdvhcTftpTableInvert = 2
	SlbCurAdvhcTftpTableInvert_Unsupported SlbCurAdvhcTftpTableInvert = 2147483647
)

type SlbCurAdvhcTftpTableParams struct {
	// TFTP Health check id.
	ID string `json:"ID,omitempty"`
	// TFTP Health check name.
	Name string `json:"Name,omitempty"`
	// TFTP Health check destination port.
	DPort uint64 `json:"DPort,omitempty"`
	// TFTP Health check destination IP version.
	IPVer SlbCurAdvhcTftpTableIPVer `json:"IPVer,omitempty"`
	// TFTP Health check host name.
	HostName string `json:"HostName,omitempty"`
	// TFTP Health check transparent flag.
	Transparent SlbCurAdvhcTftpTableTransparent `json:"Transparent,omitempty"`
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
	Invert SlbCurAdvhcTftpTableInvert `json:"Invert,omitempty"`
	// .
	TftpfileFullPath string `json:"TftpfileFullPath,omitempty"`
}

func (p SlbCurAdvhcTftpTableParams) iMABean() {}
