package beans

import (
	"fmt"
	"reflect"
)

// SlbNewAdvhcWapTable Note:This mib is not supported on VX instance of Virtualization.
type SlbNewAdvhcWapTable struct {
	// WAP Health check id.
	SlbNewAdvhcWapID string
	Params           *SlbNewAdvhcWapTableParams
}

func NewSlbNewAdvhcWapTableList() *SlbNewAdvhcWapTable {
	return &SlbNewAdvhcWapTable{}
}

func NewSlbNewAdvhcWapTable(
	slbNewAdvhcWapID string,
	params *SlbNewAdvhcWapTableParams,
) *SlbNewAdvhcWapTable {
	return &SlbNewAdvhcWapTable{
		SlbNewAdvhcWapID: slbNewAdvhcWapID,
		Params:           params,
	}
}

func (c *SlbNewAdvhcWapTable) Name() string {
	return "SlbNewAdvhcWapTable"
}

func (c *SlbNewAdvhcWapTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbNewAdvhcWapTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbNewAdvhcWapTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbNewAdvhcWapID).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbNewAdvhcWapID)
}

type SlbNewAdvhcWapTableIPVer int32

const (
	SlbNewAdvhcWapTableIPVer_Ipv4 SlbNewAdvhcWapTableIPVer = 1
	SlbNewAdvhcWapTableIPVer_Ipv6 SlbNewAdvhcWapTableIPVer = 2
	SlbNewAdvhcWapTableIPVer_None SlbNewAdvhcWapTableIPVer = 3
)

type SlbNewAdvhcWapTableTransparent int32

const (
	SlbNewAdvhcWapTableTransparent_Enabled     SlbNewAdvhcWapTableTransparent = 1
	SlbNewAdvhcWapTableTransparent_Disabled    SlbNewAdvhcWapTableTransparent = 2
	SlbNewAdvhcWapTableTransparent_Unsupported SlbNewAdvhcWapTableTransparent = 2147483647
)

type SlbNewAdvhcWapTableInvert int32

const (
	SlbNewAdvhcWapTableInvert_Enabled     SlbNewAdvhcWapTableInvert = 1
	SlbNewAdvhcWapTableInvert_Disabled    SlbNewAdvhcWapTableInvert = 2
	SlbNewAdvhcWapTableInvert_Unsupported SlbNewAdvhcWapTableInvert = 2147483647
)

type SlbNewAdvhcWapTableType int32

const (
	SlbNewAdvhcWapTableType_Wsp         SlbNewAdvhcWapTableType = 1
	SlbNewAdvhcWapTableType_Wtp         SlbNewAdvhcWapTableType = 2
	SlbNewAdvhcWapTableType_Wtlswsp     SlbNewAdvhcWapTableType = 3
	SlbNewAdvhcWapTableType_Wtlswtp     SlbNewAdvhcWapTableType = 4
	SlbNewAdvhcWapTableType_Wtls        SlbNewAdvhcWapTableType = 5
	SlbNewAdvhcWapTableType_Unsupported SlbNewAdvhcWapTableType = 2147483647
)

type SlbNewAdvhcWapTableCouple int32

const (
	SlbNewAdvhcWapTableCouple_Enabled     SlbNewAdvhcWapTableCouple = 1
	SlbNewAdvhcWapTableCouple_Disabled    SlbNewAdvhcWapTableCouple = 2
	SlbNewAdvhcWapTableCouple_Unsupported SlbNewAdvhcWapTableCouple = 2147483647
)

type SlbNewAdvhcWapTableDelete int32

const (
	SlbNewAdvhcWapTableDelete_Other       SlbNewAdvhcWapTableDelete = 1
	SlbNewAdvhcWapTableDelete_Delete      SlbNewAdvhcWapTableDelete = 2
	SlbNewAdvhcWapTableDelete_Unsupported SlbNewAdvhcWapTableDelete = 2147483647
)

type SlbNewAdvhcWapTableSnat int32

const (
	SlbNewAdvhcWapTableSnat_Enabled     SlbNewAdvhcWapTableSnat = 1
	SlbNewAdvhcWapTableSnat_Disabled    SlbNewAdvhcWapTableSnat = 2
	SlbNewAdvhcWapTableSnat_Unsupported SlbNewAdvhcWapTableSnat = 2147483647
)

type SlbNewAdvhcWapTableParams struct {
	// WAP Health check id.
	ID string `json:"ID,omitempty"`
	// WAP Health check name.
	Name string `json:"Name,omitempty"`
	// WAP Health check destination port.
	DPort uint64 `json:"DPort,omitempty"`
	// WAP Health check destination IP version.
	IPVer SlbNewAdvhcWapTableIPVer `json:"IPVer,omitempty"`
	// WAP Health check destination host name.
	HostName string `json:"HostName,omitempty"`
	// WAP Health check transparent flag.
	Transparent SlbNewAdvhcWapTableTransparent `json:"Transparent,omitempty"`
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
	Invert SlbNewAdvhcWapTableInvert `json:"Invert,omitempty"`
	// WAP Health check type.
	// wtls mode allowed only for upgrade OOTB wtls object
	Type SlbNewAdvhcWapTableType `json:"Type,omitempty"`
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
	Couple SlbNewAdvhcWapTableCouple `json:"Couple,omitempty"`
	// WAP Health check copy indicator.
	Copy string `json:"Copy,omitempty"`
	// When set to the value of 2 (delete), the entire row is deleted.
	// When read, other(1) is returned. Setting the value to anything
	// other than 2(delete) has no effect on the state of the row.
	Delete SlbNewAdvhcWapTableDelete `json:"Delete,omitempty"`
	// WAP Health Check src NAT (PIP) flag.
	Snat SlbNewAdvhcWapTableSnat `json:"Snat,omitempty"`
}

func (p SlbNewAdvhcWapTableParams) iMABean() {}
