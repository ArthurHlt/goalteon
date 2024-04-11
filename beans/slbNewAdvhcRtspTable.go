package beans

import (
	"fmt"
	"reflect"
)

// SlbNewAdvhcRtspTable Note:This mib is not supported on VX instance of Virtualization.
type SlbNewAdvhcRtspTable struct {
	// RTSP health check id.
	SlbNewAdvhcRtspID string
	Params            *SlbNewAdvhcRtspTableParams
}

func NewSlbNewAdvhcRtspTableList() *SlbNewAdvhcRtspTable {
	return &SlbNewAdvhcRtspTable{}
}

func NewSlbNewAdvhcRtspTable(
	slbNewAdvhcRtspID string,
	params *SlbNewAdvhcRtspTableParams,
) *SlbNewAdvhcRtspTable {
	return &SlbNewAdvhcRtspTable{
		SlbNewAdvhcRtspID: slbNewAdvhcRtspID,
		Params:            params,
	}
}

func (c *SlbNewAdvhcRtspTable) Name() string {
	return "SlbNewAdvhcRtspTable"
}

func (c *SlbNewAdvhcRtspTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbNewAdvhcRtspTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbNewAdvhcRtspTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbNewAdvhcRtspID).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbNewAdvhcRtspID)
}

type SlbNewAdvhcRtspTableIPVer int32

const (
	SlbNewAdvhcRtspTableIPVer_Ipv4 SlbNewAdvhcRtspTableIPVer = 1
	SlbNewAdvhcRtspTableIPVer_Ipv6 SlbNewAdvhcRtspTableIPVer = 2
	SlbNewAdvhcRtspTableIPVer_None SlbNewAdvhcRtspTableIPVer = 3
)

type SlbNewAdvhcRtspTableTransparent int32

const (
	SlbNewAdvhcRtspTableTransparent_Enabled     SlbNewAdvhcRtspTableTransparent = 1
	SlbNewAdvhcRtspTableTransparent_Disabled    SlbNewAdvhcRtspTableTransparent = 2
	SlbNewAdvhcRtspTableTransparent_Unsupported SlbNewAdvhcRtspTableTransparent = 2147483647
)

type SlbNewAdvhcRtspTableInvert int32

const (
	SlbNewAdvhcRtspTableInvert_Enabled     SlbNewAdvhcRtspTableInvert = 1
	SlbNewAdvhcRtspTableInvert_Disabled    SlbNewAdvhcRtspTableInvert = 2
	SlbNewAdvhcRtspTableInvert_Unsupported SlbNewAdvhcRtspTableInvert = 2147483647
)

type SlbNewAdvhcRtspTableMethod int32

const (
	SlbNewAdvhcRtspTableMethod_Options     SlbNewAdvhcRtspTableMethod = 1
	SlbNewAdvhcRtspTableMethod_Describe    SlbNewAdvhcRtspTableMethod = 2
	SlbNewAdvhcRtspTableMethod_Inherit     SlbNewAdvhcRtspTableMethod = 3
	SlbNewAdvhcRtspTableMethod_Unsupported SlbNewAdvhcRtspTableMethod = 2147483647
)

type SlbNewAdvhcRtspTableDelete int32

const (
	SlbNewAdvhcRtspTableDelete_Other       SlbNewAdvhcRtspTableDelete = 1
	SlbNewAdvhcRtspTableDelete_Delete      SlbNewAdvhcRtspTableDelete = 2
	SlbNewAdvhcRtspTableDelete_Unsupported SlbNewAdvhcRtspTableDelete = 2147483647
)

type SlbNewAdvhcRtspTableParams struct {
	// RTSP health check id.
	ID string `json:"ID,omitempty"`
	// RTSP health check name.
	Name string `json:"Name,omitempty"`
	// RTSP health check destination port.
	DPort uint64 `json:"DPort,omitempty"`
	// RTSP health check destination IP version.
	IPVer SlbNewAdvhcRtspTableIPVer `json:"IPVer,omitempty"`
	// RTSP health check destination hostname.
	HostName string `json:"HostName,omitempty"`
	// RTSP health check transparent flag.
	Transparent SlbNewAdvhcRtspTableTransparent `json:"Transparent,omitempty"`
	// RTSP health check interval.
	Interval uint64 `json:"Interval,omitempty"`
	// RTSP health check retries counter.
	Retries uint64 `json:"Retries,omitempty"`
	// RTSP health check retries counter in the down state.
	RestoreRetries uint64 `json:"RestoreRetries,omitempty"`
	// RTSP health check timeout.
	Timeout uint64 `json:"Timeout,omitempty"`
	// RTSP health check overflow flag.
	Overflow uint64 `json:"Overflow,omitempty"`
	// RTSP health check interval when in the down state.
	DownInterval uint64 `json:"DownInterval,omitempty"`
	// RTSP health check invert flag.
	Invert SlbNewAdvhcRtspTableInvert `json:"Invert,omitempty"`
	// RTSP health check method.
	Method SlbNewAdvhcRtspTableMethod `json:"Method,omitempty"`
	// RTSP health check server hostname.
	Hostname string `json:"Hostname,omitempty"`
	// RTSP health check path.
	Path string `json:"Path,omitempty"`
	// RTSP health check response code.
	ResponseCodes string `json:"ResponseCodes,omitempty"`
	// RTSP health check copy flag.
	Copy string `json:"Copy,omitempty"`
	// When set to the value of 2 (delete), the entire row is deleted.
	// When read, other(1) is returned. Setting the value to anything
	// other than 2(delete) has no effect on the state of the row.
	Delete SlbNewAdvhcRtspTableDelete `json:"Delete,omitempty"`
}

func (p SlbNewAdvhcRtspTableParams) iMABean() {}
