package beans

import (
	"fmt"
	"reflect"
)

// SlbCurAdvhcImapTable Note:This mib is not supported on VX instance of Virtualization.
type SlbCurAdvhcImapTable struct {
	// IMAP Health check id.
	SlbCurAdvhcImapID string
	Params            *SlbCurAdvhcImapTableParams
}

func NewSlbCurAdvhcImapTableList() *SlbCurAdvhcImapTable {
	return &SlbCurAdvhcImapTable{}
}

func NewSlbCurAdvhcImapTable(
	slbCurAdvhcImapID string,
	params *SlbCurAdvhcImapTableParams,
) *SlbCurAdvhcImapTable {
	return &SlbCurAdvhcImapTable{
		SlbCurAdvhcImapID: slbCurAdvhcImapID,
		Params:            params,
	}
}

func (c *SlbCurAdvhcImapTable) Name() string {
	return "SlbCurAdvhcImapTable"
}

func (c *SlbCurAdvhcImapTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbCurAdvhcImapTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbCurAdvhcImapTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbCurAdvhcImapID).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbCurAdvhcImapID)
}

type SlbCurAdvhcImapTableIPVer int32

const (
	SlbCurAdvhcImapTableIPVer_Ipv4 SlbCurAdvhcImapTableIPVer = 1
	SlbCurAdvhcImapTableIPVer_Ipv6 SlbCurAdvhcImapTableIPVer = 2
	SlbCurAdvhcImapTableIPVer_None SlbCurAdvhcImapTableIPVer = 3
)

type SlbCurAdvhcImapTableTransparent int32

const (
	SlbCurAdvhcImapTableTransparent_Enabled     SlbCurAdvhcImapTableTransparent = 1
	SlbCurAdvhcImapTableTransparent_Disabled    SlbCurAdvhcImapTableTransparent = 2
	SlbCurAdvhcImapTableTransparent_Unsupported SlbCurAdvhcImapTableTransparent = 2147483647
)

type SlbCurAdvhcImapTableInvert int32

const (
	SlbCurAdvhcImapTableInvert_Enabled     SlbCurAdvhcImapTableInvert = 1
	SlbCurAdvhcImapTableInvert_Disabled    SlbCurAdvhcImapTableInvert = 2
	SlbCurAdvhcImapTableInvert_Unsupported SlbCurAdvhcImapTableInvert = 2147483647
)

type SlbCurAdvhcImapTableSnat int32

const (
	SlbCurAdvhcImapTableSnat_Enabled     SlbCurAdvhcImapTableSnat = 1
	SlbCurAdvhcImapTableSnat_Disabled    SlbCurAdvhcImapTableSnat = 2
	SlbCurAdvhcImapTableSnat_Unsupported SlbCurAdvhcImapTableSnat = 2147483647
)

type SlbCurAdvhcImapTableParams struct {
	// IMAP Health check id.
	ID string `json:"ID,omitempty"`
	// IMAP Health check name.
	Name string `json:"Name,omitempty"`
	// IMAP Health check destination port.
	DPort uint64 `json:"DPort,omitempty"`
	// IMAP Health check destination IP version.
	IPVer SlbCurAdvhcImapTableIPVer `json:"IPVer,omitempty"`
	// IMAP Health check destination hostname.
	HostName string `json:"HostName,omitempty"`
	// IMAP Health check transparent flag.
	Transparent SlbCurAdvhcImapTableTransparent `json:"Transparent,omitempty"`
	// IMAP Health check interval.
	Interval uint64 `json:"Interval,omitempty"`
	// IMAP Health check retries counter.
	Retries uint64 `json:"Retries,omitempty"`
	// IMAP Health check retries from the down state counter.
	RestoreRetries uint64 `json:"RestoreRetries,omitempty"`
	// IMAP Health check timeout.
	Timeout uint64 `json:"Timeout,omitempty"`
	// IMAP Health check overflow flag.
	Overflow uint64 `json:"Overflow,omitempty"`
	// IMAP Health check interval when in the down state.
	DownInterval uint64 `json:"DownInterval,omitempty"`
	// IMAP Health check invert flag.
	Invert SlbCurAdvhcImapTableInvert `json:"Invert,omitempty"`
	// IMAP Health check username.
	UserName string `json:"UserName,omitempty"`
	// IMAP Health check password.
	Password string `json:"Password,omitempty"`
	// IMAP Health Check src NAT (PIP) flag.
	Snat SlbCurAdvhcImapTableSnat `json:"Snat,omitempty"`
}

func (p SlbCurAdvhcImapTableParams) iMABean() {}
