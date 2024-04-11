package beans

import (
	"fmt"
	"reflect"
)

// SlbCurAdvhcRtspTable Note:This mib is not supported on VX instance of Virtualization.
type SlbCurAdvhcRtspTable struct {
	// RTSP health check id.
	SlbCurAdvhcRtspID string
	Params            *SlbCurAdvhcRtspTableParams
}

func NewSlbCurAdvhcRtspTableList() *SlbCurAdvhcRtspTable {
	return &SlbCurAdvhcRtspTable{}
}

func NewSlbCurAdvhcRtspTable(
	slbCurAdvhcRtspID string,
	params *SlbCurAdvhcRtspTableParams,
) *SlbCurAdvhcRtspTable {
	return &SlbCurAdvhcRtspTable{
		SlbCurAdvhcRtspID: slbCurAdvhcRtspID,
		Params:            params,
	}
}

func (c *SlbCurAdvhcRtspTable) Name() string {
	return "SlbCurAdvhcRtspTable"
}

func (c *SlbCurAdvhcRtspTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbCurAdvhcRtspTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbCurAdvhcRtspTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbCurAdvhcRtspID).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbCurAdvhcRtspID)
}

type SlbCurAdvhcRtspTableIPVer int32

const (
	SlbCurAdvhcRtspTableIPVer_Ipv4 SlbCurAdvhcRtspTableIPVer = 1
	SlbCurAdvhcRtspTableIPVer_Ipv6 SlbCurAdvhcRtspTableIPVer = 2
	SlbCurAdvhcRtspTableIPVer_None SlbCurAdvhcRtspTableIPVer = 3
)

type SlbCurAdvhcRtspTableTransparent int32

const (
	SlbCurAdvhcRtspTableTransparent_Enabled     SlbCurAdvhcRtspTableTransparent = 1
	SlbCurAdvhcRtspTableTransparent_Disabled    SlbCurAdvhcRtspTableTransparent = 2
	SlbCurAdvhcRtspTableTransparent_Unsupported SlbCurAdvhcRtspTableTransparent = 2147483647
)

type SlbCurAdvhcRtspTableInvert int32

const (
	SlbCurAdvhcRtspTableInvert_Enabled     SlbCurAdvhcRtspTableInvert = 1
	SlbCurAdvhcRtspTableInvert_Disabled    SlbCurAdvhcRtspTableInvert = 2
	SlbCurAdvhcRtspTableInvert_Unsupported SlbCurAdvhcRtspTableInvert = 2147483647
)

type SlbCurAdvhcRtspTableMethod int32

const (
	SlbCurAdvhcRtspTableMethod_Options     SlbCurAdvhcRtspTableMethod = 1
	SlbCurAdvhcRtspTableMethod_Describe    SlbCurAdvhcRtspTableMethod = 2
	SlbCurAdvhcRtspTableMethod_Inherit     SlbCurAdvhcRtspTableMethod = 3
	SlbCurAdvhcRtspTableMethod_Unsupported SlbCurAdvhcRtspTableMethod = 2147483647
)

type SlbCurAdvhcRtspTableParams struct {
	// RTSP health check id.
	ID string `json:"ID,omitempty"`
	// RTSP health check name.
	Name string `json:"Name,omitempty"`
	// RTSP health check destination port.
	DPort uint64 `json:"DPort,omitempty"`
	// RTSP health check destination IP version.
	IPVer SlbCurAdvhcRtspTableIPVer `json:"IPVer,omitempty"`
	// RTSP health check destination hostname.
	HostName string `json:"HostName,omitempty"`
	// RTSP health check transparent flag.
	Transparent SlbCurAdvhcRtspTableTransparent `json:"Transparent,omitempty"`
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
	Invert SlbCurAdvhcRtspTableInvert `json:"Invert,omitempty"`
	// RTSP health check method.
	Method SlbCurAdvhcRtspTableMethod `json:"Method,omitempty"`
	// RTSP health check server hostname.
	Hostname string `json:"Hostname,omitempty"`
	// RTSP health check path.
	Path string `json:"Path,omitempty"`
	// RTSP health check response code.
	ResponseCodes string `json:"ResponseCodes,omitempty"`
}

func (p SlbCurAdvhcRtspTableParams) iMABean() {}
