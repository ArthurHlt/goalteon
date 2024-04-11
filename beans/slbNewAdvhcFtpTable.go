package beans

import (
	"fmt"
	"reflect"
)

// SlbNewAdvhcFtpTable Note:This mib is not supported on VX instance of Virtualization.
type SlbNewAdvhcFtpTable struct {
	// FTP Health check id.
	SlbNewAdvhcFtpID string
	Params           *SlbNewAdvhcFtpTableParams
}

func NewSlbNewAdvhcFtpTableList() *SlbNewAdvhcFtpTable {
	return &SlbNewAdvhcFtpTable{}
}

func NewSlbNewAdvhcFtpTable(
	slbNewAdvhcFtpID string,
	params *SlbNewAdvhcFtpTableParams,
) *SlbNewAdvhcFtpTable {
	return &SlbNewAdvhcFtpTable{
		SlbNewAdvhcFtpID: slbNewAdvhcFtpID,
		Params:           params,
	}
}

func (c *SlbNewAdvhcFtpTable) Name() string {
	return "SlbNewAdvhcFtpTable"
}

func (c *SlbNewAdvhcFtpTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbNewAdvhcFtpTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbNewAdvhcFtpTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbNewAdvhcFtpID).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbNewAdvhcFtpID)
}

type SlbNewAdvhcFtpTableIPVer int32

const (
	SlbNewAdvhcFtpTableIPVer_Ipv4 SlbNewAdvhcFtpTableIPVer = 1
	SlbNewAdvhcFtpTableIPVer_Ipv6 SlbNewAdvhcFtpTableIPVer = 2
	SlbNewAdvhcFtpTableIPVer_None SlbNewAdvhcFtpTableIPVer = 3
)

type SlbNewAdvhcFtpTableTransparent int32

const (
	SlbNewAdvhcFtpTableTransparent_Enabled     SlbNewAdvhcFtpTableTransparent = 1
	SlbNewAdvhcFtpTableTransparent_Disabled    SlbNewAdvhcFtpTableTransparent = 2
	SlbNewAdvhcFtpTableTransparent_Unsupported SlbNewAdvhcFtpTableTransparent = 2147483647
)

type SlbNewAdvhcFtpTableInvert int32

const (
	SlbNewAdvhcFtpTableInvert_Enabled     SlbNewAdvhcFtpTableInvert = 1
	SlbNewAdvhcFtpTableInvert_Disabled    SlbNewAdvhcFtpTableInvert = 2
	SlbNewAdvhcFtpTableInvert_Unsupported SlbNewAdvhcFtpTableInvert = 2147483647
)

type SlbNewAdvhcFtpTableDelete int32

const (
	SlbNewAdvhcFtpTableDelete_Other       SlbNewAdvhcFtpTableDelete = 1
	SlbNewAdvhcFtpTableDelete_Delete      SlbNewAdvhcFtpTableDelete = 2
	SlbNewAdvhcFtpTableDelete_Unsupported SlbNewAdvhcFtpTableDelete = 2147483647
)

type SlbNewAdvhcFtpTableParams struct {
	// FTP Health check id.
	ID string `json:"ID,omitempty"`
	// FTP Health check name.
	Name string `json:"Name,omitempty"`
	// FTP Health check destination port.
	DPort uint64 `json:"DPort,omitempty"`
	// FTP Health check destination IP version.
	IPVer SlbNewAdvhcFtpTableIPVer `json:"IPVer,omitempty"`
	// FTP Health check destination hostname.
	HostName string `json:"HostName,omitempty"`
	// FTP Health check transparent flag.
	Transparent SlbNewAdvhcFtpTableTransparent `json:"Transparent,omitempty"`
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
	Invert SlbNewAdvhcFtpTableInvert `json:"Invert,omitempty"`
	// FTP Health check connection user name.
	UserName string `json:"UserName,omitempty"`
	// FTP Health check connection password.
	Password string `json:"Password,omitempty"`
	// FTP Health check file path parameter.
	Path string `json:"Path,omitempty"`
	// FTP Health check copy flag.
	Copy string `json:"Copy,omitempty"`
	// When set to the value of 2 (delete), the entire row is deleted.
	// When read, other(1) is returned. Setting the value to anything
	// other than 2(delete) has no effect on the state of the row.
	Delete SlbNewAdvhcFtpTableDelete `json:"Delete,omitempty"`
}

func (p SlbNewAdvhcFtpTableParams) iMABean() {}
