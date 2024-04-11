package beans

import (
	"fmt"
	"reflect"
)

// SlbCurAdvhcLdapTable Note:This mib is not supported on VX instance of Virtualization.
type SlbCurAdvhcLdapTable struct {
	// LDAP Health check id.
	SlbCurAdvhcLdapID string
	Params            *SlbCurAdvhcLdapTableParams
}

func NewSlbCurAdvhcLdapTableList() *SlbCurAdvhcLdapTable {
	return &SlbCurAdvhcLdapTable{}
}

func NewSlbCurAdvhcLdapTable(
	slbCurAdvhcLdapID string,
	params *SlbCurAdvhcLdapTableParams,
) *SlbCurAdvhcLdapTable {
	return &SlbCurAdvhcLdapTable{
		SlbCurAdvhcLdapID: slbCurAdvhcLdapID,
		Params:            params,
	}
}

func (c *SlbCurAdvhcLdapTable) Name() string {
	return "SlbCurAdvhcLdapTable"
}

func (c *SlbCurAdvhcLdapTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbCurAdvhcLdapTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbCurAdvhcLdapTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbCurAdvhcLdapID).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbCurAdvhcLdapID)
}

type SlbCurAdvhcLdapTableIPVer int32

const (
	SlbCurAdvhcLdapTableIPVer_Ipv4 SlbCurAdvhcLdapTableIPVer = 1
	SlbCurAdvhcLdapTableIPVer_Ipv6 SlbCurAdvhcLdapTableIPVer = 2
	SlbCurAdvhcLdapTableIPVer_None SlbCurAdvhcLdapTableIPVer = 3
)

type SlbCurAdvhcLdapTableTransparent int32

const (
	SlbCurAdvhcLdapTableTransparent_Enabled     SlbCurAdvhcLdapTableTransparent = 1
	SlbCurAdvhcLdapTableTransparent_Disabled    SlbCurAdvhcLdapTableTransparent = 2
	SlbCurAdvhcLdapTableTransparent_Unsupported SlbCurAdvhcLdapTableTransparent = 2147483647
)

type SlbCurAdvhcLdapTableInvert int32

const (
	SlbCurAdvhcLdapTableInvert_Enabled     SlbCurAdvhcLdapTableInvert = 1
	SlbCurAdvhcLdapTableInvert_Disabled    SlbCurAdvhcLdapTableInvert = 2
	SlbCurAdvhcLdapTableInvert_Unsupported SlbCurAdvhcLdapTableInvert = 2147483647
)

type SlbCurAdvhcLdapTableLdaps int32

const (
	SlbCurAdvhcLdapTableLdaps_Enabled     SlbCurAdvhcLdapTableLdaps = 1
	SlbCurAdvhcLdapTableLdaps_Disabled    SlbCurAdvhcLdapTableLdaps = 2
	SlbCurAdvhcLdapTableLdaps_Unsupported SlbCurAdvhcLdapTableLdaps = 2147483647
)

type SlbCurAdvhcLdapTableBaseFmt int32

const (
	SlbCurAdvhcLdapTableBaseFmt_Fqdn        SlbCurAdvhcLdapTableBaseFmt = 0
	SlbCurAdvhcLdapTableBaseFmt_Ldap        SlbCurAdvhcLdapTableBaseFmt = 1
	SlbCurAdvhcLdapTableBaseFmt_Unsupported SlbCurAdvhcLdapTableBaseFmt = 2147483647
)

type SlbCurAdvhcLdapTableParams struct {
	// LDAP Health check id.
	ID string `json:"ID,omitempty"`
	// LDAP Health check name.
	Name string `json:"Name,omitempty"`
	// LDAP Health check destination port.
	DPort uint64 `json:"DPort,omitempty"`
	// LDAP Health check destination IP version.
	IPVer SlbCurAdvhcLdapTableIPVer `json:"IPVer,omitempty"`
	// LDAP Health check destination hostname.
	HostName string `json:"HostName,omitempty"`
	// LDAP Health check transparent flag.
	Transparent SlbCurAdvhcLdapTableTransparent `json:"Transparent,omitempty"`
	// LDAP Health check interval.
	Interval uint64 `json:"Interval,omitempty"`
	// LDAP Health check retries counter.
	Retries uint64 `json:"Retries,omitempty"`
	// LDAP Health check restore from down state counter.
	RestoreRetries uint64 `json:"RestoreRetries,omitempty"`
	// LDAP Health check timeout.
	Timeout uint64 `json:"Timeout,omitempty"`
	// LDAP Health check overflow flag.
	Overflow uint64 `json:"Overflow,omitempty"`
	// LDAP Health check interval in down state.
	DownInterval uint64 `json:"DownInterval,omitempty"`
	// LDAP Health check invert flag.
	Invert SlbCurAdvhcLdapTableInvert `json:"Invert,omitempty"`
	// LDAP Health check e/d secure protocol.
	Ldaps SlbCurAdvhcLdapTableLdaps `json:"Ldaps,omitempty"`
	// LDAP Health check user name.
	UserName string `json:"UserName,omitempty"`
	// LDAP Health check password.
	Password string `json:"Password,omitempty"`
	// LDAP Health check base distinguish name.
	BaseObject string `json:"BaseObject,omitempty"`
	// LDAP base DN format.
	BaseFmt SlbCurAdvhcLdapTableBaseFmt `json:"BaseFmt,omitempty"`
}

func (p SlbCurAdvhcLdapTableParams) iMABean() {}
