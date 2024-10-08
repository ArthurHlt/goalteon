package beans

import (
	"fmt"
	"reflect"
)

// SlbNewAdvhcImapTable Note:This mib is not supported on VX instance of Virtualization.
type SlbNewAdvhcImapTable struct {
	// IMAP Health check id.
	SlbNewAdvhcImapID string
	Params            *SlbNewAdvhcImapTableParams
}

func NewSlbNewAdvhcImapTableList() *SlbNewAdvhcImapTable {
	return &SlbNewAdvhcImapTable{}
}

func NewSlbNewAdvhcImapTable(
	slbNewAdvhcImapID string,
	params *SlbNewAdvhcImapTableParams,
) *SlbNewAdvhcImapTable {
	return &SlbNewAdvhcImapTable{
		SlbNewAdvhcImapID: slbNewAdvhcImapID,
		Params:            params,
	}
}

func (c *SlbNewAdvhcImapTable) Name() string {
	return "SlbNewAdvhcImapTable"
}

func (c *SlbNewAdvhcImapTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbNewAdvhcImapTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbNewAdvhcImapTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbNewAdvhcImapID).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbNewAdvhcImapID)
}

type SlbNewAdvhcImapTableIPVer int32

const (
	SlbNewAdvhcImapTableIPVer_Ipv4 SlbNewAdvhcImapTableIPVer = 1
	SlbNewAdvhcImapTableIPVer_Ipv6 SlbNewAdvhcImapTableIPVer = 2
	SlbNewAdvhcImapTableIPVer_None SlbNewAdvhcImapTableIPVer = 3
)

type SlbNewAdvhcImapTableTransparent int32

const (
	SlbNewAdvhcImapTableTransparent_Enabled     SlbNewAdvhcImapTableTransparent = 1
	SlbNewAdvhcImapTableTransparent_Disabled    SlbNewAdvhcImapTableTransparent = 2
	SlbNewAdvhcImapTableTransparent_Unsupported SlbNewAdvhcImapTableTransparent = 2147483647
)

type SlbNewAdvhcImapTableInvert int32

const (
	SlbNewAdvhcImapTableInvert_Enabled     SlbNewAdvhcImapTableInvert = 1
	SlbNewAdvhcImapTableInvert_Disabled    SlbNewAdvhcImapTableInvert = 2
	SlbNewAdvhcImapTableInvert_Unsupported SlbNewAdvhcImapTableInvert = 2147483647
)

type SlbNewAdvhcImapTableDelete int32

const (
	SlbNewAdvhcImapTableDelete_Other       SlbNewAdvhcImapTableDelete = 1
	SlbNewAdvhcImapTableDelete_Delete      SlbNewAdvhcImapTableDelete = 2
	SlbNewAdvhcImapTableDelete_Unsupported SlbNewAdvhcImapTableDelete = 2147483647
)

type SlbNewAdvhcImapTableSnat int32

const (
	SlbNewAdvhcImapTableSnat_Enabled     SlbNewAdvhcImapTableSnat = 1
	SlbNewAdvhcImapTableSnat_Disabled    SlbNewAdvhcImapTableSnat = 2
	SlbNewAdvhcImapTableSnat_Unsupported SlbNewAdvhcImapTableSnat = 2147483647
)

type SlbNewAdvhcImapTableParams struct {
	// IMAP Health check id.
	ID string `json:"ID,omitempty"`
	// IMAP Health check name.
	Name string `json:"Name,omitempty"`
	// IMAP Health check destination port.
	DPort uint64 `json:"DPort,omitempty"`
	// IMAP Health check destination IP version.
	IPVer SlbNewAdvhcImapTableIPVer `json:"IPVer,omitempty"`
	// IMAP Health check destination hostname.
	HostName string `json:"HostName,omitempty"`
	// IMAP Health check transparent flag.
	Transparent SlbNewAdvhcImapTableTransparent `json:"Transparent,omitempty"`
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
	Invert SlbNewAdvhcImapTableInvert `json:"Invert,omitempty"`
	// IMAP Health check username.
	UserName string `json:"UserName,omitempty"`
	// IMAP Health check password.
	Password string `json:"Password,omitempty"`
	// IMAP Health check copy flag.
	Copy DisplayString `json:"Copy,omitempty"`
	// When set to the value of 2 (delete), the entire row is deleted.
	// When read, other(1) is returned. Setting the value to anything
	// other than 2(delete) has no effect on the state of the row.
	Delete SlbNewAdvhcImapTableDelete `json:"Delete,omitempty"`
	// IMAP Health Check src NAT (PIP) flag.
	Snat SlbNewAdvhcImapTableSnat `json:"Snat,omitempty"`
}

func (p SlbNewAdvhcImapTableParams) iMABean() {}
