package beans

import (
	"fmt"
	"reflect"
)

// SlbCurAdvhcWapTable Note:This mib is not supported on VX instance of Virtualization.
type SlbCurAdvhcWapTable struct {
	// WAP Health check id.
	SlbCurAdvhcWapID string
	Params           *SlbCurAdvhcWapTableParams
}

func NewSlbCurAdvhcWapTableList() *SlbCurAdvhcWapTable {
	return &SlbCurAdvhcWapTable{}
}

func NewSlbCurAdvhcWapTable(
	slbCurAdvhcWapID string,
	params *SlbCurAdvhcWapTableParams,
) *SlbCurAdvhcWapTable {
	return &SlbCurAdvhcWapTable{
		SlbCurAdvhcWapID: slbCurAdvhcWapID,
		Params:           params,
	}
}

func (c *SlbCurAdvhcWapTable) Name() string {
	return "SlbCurAdvhcWapTable"
}

func (c *SlbCurAdvhcWapTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbCurAdvhcWapTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbCurAdvhcWapTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbCurAdvhcWapID).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbCurAdvhcWapID)
}

type SlbCurAdvhcWapTableIPVer int32

const (
	SlbCurAdvhcWapTableIPVer_Ipv4 SlbCurAdvhcWapTableIPVer = 1
	SlbCurAdvhcWapTableIPVer_Ipv6 SlbCurAdvhcWapTableIPVer = 2
	SlbCurAdvhcWapTableIPVer_None SlbCurAdvhcWapTableIPVer = 3
)

type SlbCurAdvhcWapTableTransparent int32

const (
	SlbCurAdvhcWapTableTransparent_Enabled     SlbCurAdvhcWapTableTransparent = 1
	SlbCurAdvhcWapTableTransparent_Disabled    SlbCurAdvhcWapTableTransparent = 2
	SlbCurAdvhcWapTableTransparent_Unsupported SlbCurAdvhcWapTableTransparent = 2147483647
)

type SlbCurAdvhcWapTableInvert int32

const (
	SlbCurAdvhcWapTableInvert_Enabled     SlbCurAdvhcWapTableInvert = 1
	SlbCurAdvhcWapTableInvert_Disabled    SlbCurAdvhcWapTableInvert = 2
	SlbCurAdvhcWapTableInvert_Unsupported SlbCurAdvhcWapTableInvert = 2147483647
)

type SlbCurAdvhcWapTableType int32

const (
	SlbCurAdvhcWapTableType_Wsp         SlbCurAdvhcWapTableType = 1
	SlbCurAdvhcWapTableType_Wtp         SlbCurAdvhcWapTableType = 2
	SlbCurAdvhcWapTableType_Wtlswsp     SlbCurAdvhcWapTableType = 3
	SlbCurAdvhcWapTableType_Wtlswtp     SlbCurAdvhcWapTableType = 4
	SlbCurAdvhcWapTableType_Wtls        SlbCurAdvhcWapTableType = 5
	SlbCurAdvhcWapTableType_Unsupported SlbCurAdvhcWapTableType = 2147483647
)

type SlbCurAdvhcWapTableCouple int32

const (
	SlbCurAdvhcWapTableCouple_Enabled     SlbCurAdvhcWapTableCouple = 1
	SlbCurAdvhcWapTableCouple_Disabled    SlbCurAdvhcWapTableCouple = 2
	SlbCurAdvhcWapTableCouple_Unsupported SlbCurAdvhcWapTableCouple = 2147483647
)

type SlbCurAdvhcWapTableSnat int32

const (
	SlbCurAdvhcWapTableSnat_Enabled     SlbCurAdvhcWapTableSnat = 1
	SlbCurAdvhcWapTableSnat_Disabled    SlbCurAdvhcWapTableSnat = 2
	SlbCurAdvhcWapTableSnat_Unsupported SlbCurAdvhcWapTableSnat = 2147483647
)

type SlbCurAdvhcWapTableParams struct {
	// WAP Health check id.
	ID string `json:"ID,omitempty"`
	// WAP Health check name.
	Name string `json:"Name,omitempty"`
	// WAP Health check destination port.
	DPort uint64 `json:"DPort,omitempty"`
	// WAP Health check destination IP version.
	IPVer SlbCurAdvhcWapTableIPVer `json:"IPVer,omitempty"`
	// WAP Health check destination host name.
	HostName string `json:"HostName,omitempty"`
	// WAP Health check transparent flag.
	Transparent SlbCurAdvhcWapTableTransparent `json:"Transparent,omitempty"`
	// WAP Health check inteval.
	Interval uint64 `json:"Interval,omitempty"`
	// WAP Health check retries counter.
	Retries uint64 `json:"Retries,omitempty"`
	// WAP Health check retries counter when in the down state.
	RestoreRetries uint64 `json:"RestoreRetries,omitempty"`
	// WAP Health check timeout.
	Timeout uint64 `json:"Timeout,omitempty"`
	// WAP Health check overflow flag.
	Overflow uint64 `json:"Overflow,omitempty"`
	// WAP Health check interval when in the down state.
	DownInterval uint64 `json:"DownInterval,omitempty"`
	// WAP Health check invert flag.
	Invert SlbCurAdvhcWapTableInvert `json:"Invert,omitempty"`
	// WAP Health check type.wtls mode allowed only for upgrade OOTB wtls object.
	Type SlbCurAdvhcWapTableType `json:"Type,omitempty"`
	// WAP Health check string to send.
	WspSendString string `json:"WspSendString,omitempty"`
	// WAP Health check expected string.
	WspReceiveString string `json:"WspReceiveString,omitempty"`
	// WAP Health check wsp offeset.
	Wspoffset int32 `json:"Wspoffset,omitempty"`
	// WAP Health check connection header.
	ConnectHeaders string `json:"ConnectHeaders,omitempty"`
	// WAP Health check wtp string to send.
	WtpSendString string `json:"WtpSendString,omitempty"`
	// WAP Health check wtp expected string.
	WtpReceiveString string `json:"WtpReceiveString,omitempty"`
	// WAP Health check wtp offeset.
	Wtpoffset int32 `json:"Wtpoffset,omitempty"`
	// WAP Health check couple flag.
	Couple SlbCurAdvhcWapTableCouple `json:"Couple,omitempty"`
	// WAP Health Check src NAT (PIP) flag.
	Snat SlbCurAdvhcWapTableSnat `json:"Snat,omitempty"`
}

func (p SlbCurAdvhcWapTableParams) iMABean() {}
