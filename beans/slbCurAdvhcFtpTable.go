package beans

import (
	"fmt"
	"reflect"
)

// SlbCurAdvhcFtpTable Note:This mib is not supported on VX instance of Virtualization.
type SlbCurAdvhcFtpTable struct {
	// FTP Health check id.
	SlbCurAdvhcFtpID string
	Params           *SlbCurAdvhcFtpTableParams
}

func NewSlbCurAdvhcFtpTableList() *SlbCurAdvhcFtpTable {
	return &SlbCurAdvhcFtpTable{}
}

func NewSlbCurAdvhcFtpTable(
	slbCurAdvhcFtpID string,
	params *SlbCurAdvhcFtpTableParams,
) *SlbCurAdvhcFtpTable {
	return &SlbCurAdvhcFtpTable{
		SlbCurAdvhcFtpID: slbCurAdvhcFtpID,
		Params:           params,
	}
}

func (c *SlbCurAdvhcFtpTable) Name() string {
	return "SlbCurAdvhcFtpTable"
}

func (c *SlbCurAdvhcFtpTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbCurAdvhcFtpTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbCurAdvhcFtpTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbCurAdvhcFtpID).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbCurAdvhcFtpID)
}

type SlbCurAdvhcFtpTableIPVer int32

const (
	SlbCurAdvhcFtpTableIPVer_Ipv4 SlbCurAdvhcFtpTableIPVer = 1
	SlbCurAdvhcFtpTableIPVer_Ipv6 SlbCurAdvhcFtpTableIPVer = 2
	SlbCurAdvhcFtpTableIPVer_None SlbCurAdvhcFtpTableIPVer = 3
)

type SlbCurAdvhcFtpTableTransparent int32

const (
	SlbCurAdvhcFtpTableTransparent_Enabled     SlbCurAdvhcFtpTableTransparent = 1
	SlbCurAdvhcFtpTableTransparent_Disabled    SlbCurAdvhcFtpTableTransparent = 2
	SlbCurAdvhcFtpTableTransparent_Unsupported SlbCurAdvhcFtpTableTransparent = 2147483647
)

type SlbCurAdvhcFtpTableInvert int32

const (
	SlbCurAdvhcFtpTableInvert_Enabled     SlbCurAdvhcFtpTableInvert = 1
	SlbCurAdvhcFtpTableInvert_Disabled    SlbCurAdvhcFtpTableInvert = 2
	SlbCurAdvhcFtpTableInvert_Unsupported SlbCurAdvhcFtpTableInvert = 2147483647
)

type SlbCurAdvhcFtpTableParams struct {
	// FTP Health check id.
	ID string `json:"ID,omitempty"`
	// FTP Health check name.
	Name string `json:"Name,omitempty"`
	// FTP Health check destination port.
	DPort uint64 `json:"DPort,omitempty"`
	// FTP Health check destination IP version.
	IPVer SlbCurAdvhcFtpTableIPVer `json:"IPVer,omitempty"`
	// FTP Health check destination hostname.
	HostName string `json:"HostName,omitempty"`
	// FTP Health check transparent flag.
	Transparent SlbCurAdvhcFtpTableTransparent `json:"Transparent,omitempty"`
	// FTP Health check interval.
	Interval uint64 `json:"Interval,omitempty"`
	// FTP Health check retries counter.
	Retries uint64 `json:"Retries,omitempty"`
	// FTP Health check restore from down state counter.
	RestoreRetries uint64 `json:"RestoreRetries,omitempty"`
	// FTP Health check timeout period.
	Timeout uint64 `json:"Timeout,omitempty"`
	// FTP Health check overflow flag.
	Overflow uint64 `json:"Overflow,omitempty"`
	// FTP Health check interval in down state.
	DownInterval uint64 `json:"DownInterval,omitempty"`
	// FTP Health check invert flag.
	Invert SlbCurAdvhcFtpTableInvert `json:"Invert,omitempty"`
	// FTP Health check connection user name.
	UserName string `json:"UserName,omitempty"`
	// FTP Health check connection password.
	Password string `json:"Password,omitempty"`
	// FTP Health check file path parameter.
	Path string `json:"Path,omitempty"`
}

func (p SlbCurAdvhcFtpTableParams) iMABean() {}
