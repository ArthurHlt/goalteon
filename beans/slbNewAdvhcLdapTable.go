package beans

import (
	"fmt"
	"reflect"
)

// SlbNewAdvhcLdapTable Note:This mib is not supported on VX instance of Virtualization.
type SlbNewAdvhcLdapTable struct {
	// LDAP Health check id.
	SlbNewAdvhcLdapID string
	Params            *SlbNewAdvhcLdapTableParams
}

func NewSlbNewAdvhcLdapTableList() *SlbNewAdvhcLdapTable {
	return &SlbNewAdvhcLdapTable{}
}

func NewSlbNewAdvhcLdapTable(
	slbNewAdvhcLdapID string,
	params *SlbNewAdvhcLdapTableParams,
) *SlbNewAdvhcLdapTable {
	return &SlbNewAdvhcLdapTable{
		SlbNewAdvhcLdapID: slbNewAdvhcLdapID,
		Params:            params,
	}
}

func (c *SlbNewAdvhcLdapTable) Name() string {
	return "SlbNewAdvhcLdapTable"
}

func (c *SlbNewAdvhcLdapTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbNewAdvhcLdapTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbNewAdvhcLdapTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbNewAdvhcLdapID).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbNewAdvhcLdapID)
}

type SlbNewAdvhcLdapTableIPVer int32

const (
	SlbNewAdvhcLdapTableIPVer_Ipv4 SlbNewAdvhcLdapTableIPVer = 1
	SlbNewAdvhcLdapTableIPVer_Ipv6 SlbNewAdvhcLdapTableIPVer = 2
	SlbNewAdvhcLdapTableIPVer_None SlbNewAdvhcLdapTableIPVer = 3
)

type SlbNewAdvhcLdapTableTransparent int32

const (
	SlbNewAdvhcLdapTableTransparent_Enabled     SlbNewAdvhcLdapTableTransparent = 1
	SlbNewAdvhcLdapTableTransparent_Disabled    SlbNewAdvhcLdapTableTransparent = 2
	SlbNewAdvhcLdapTableTransparent_Unsupported SlbNewAdvhcLdapTableTransparent = 2147483647
)

type SlbNewAdvhcLdapTableInvert int32

const (
	SlbNewAdvhcLdapTableInvert_Enabled     SlbNewAdvhcLdapTableInvert = 1
	SlbNewAdvhcLdapTableInvert_Disabled    SlbNewAdvhcLdapTableInvert = 2
	SlbNewAdvhcLdapTableInvert_Unsupported SlbNewAdvhcLdapTableInvert = 2147483647
)

type SlbNewAdvhcLdapTableLdaps int32

const (
	SlbNewAdvhcLdapTableLdaps_Enabled     SlbNewAdvhcLdapTableLdaps = 1
	SlbNewAdvhcLdapTableLdaps_Disabled    SlbNewAdvhcLdapTableLdaps = 2
	SlbNewAdvhcLdapTableLdaps_Unsupported SlbNewAdvhcLdapTableLdaps = 2147483647
)

type SlbNewAdvhcLdapTableDelete int32

const (
	SlbNewAdvhcLdapTableDelete_Other       SlbNewAdvhcLdapTableDelete = 1
	SlbNewAdvhcLdapTableDelete_Delete      SlbNewAdvhcLdapTableDelete = 2
	SlbNewAdvhcLdapTableDelete_Unsupported SlbNewAdvhcLdapTableDelete = 2147483647
)

type SlbNewAdvhcLdapTableBaseFmt int32

const (
	SlbNewAdvhcLdapTableBaseFmt_Fqdn        SlbNewAdvhcLdapTableBaseFmt = 0
	SlbNewAdvhcLdapTableBaseFmt_Ldap        SlbNewAdvhcLdapTableBaseFmt = 1
	SlbNewAdvhcLdapTableBaseFmt_Unsupported SlbNewAdvhcLdapTableBaseFmt = 2147483647
)

type SlbNewAdvhcLdapTableParams struct {
	// LDAP Health check id.
	ID string `json:"ID,omitempty"`
	// LDAP Health check name.
	Name string `json:"Name,omitempty"`
	// LDAP Health check destination port.
	DPort uint64 `json:"DPort,omitempty"`
	// LDAP Health check destination IP version.
	IPVer SlbNewAdvhcLdapTableIPVer `json:"IPVer,omitempty"`
	// LDAP Health check destination hostname.
	HostName string `json:"HostName,omitempty"`
	// LDAP Health check transparent flag.
	Transparent SlbNewAdvhcLdapTableTransparent `json:"Transparent,omitempty"`
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
	Invert SlbNewAdvhcLdapTableInvert `json:"Invert,omitempty"`
	// LDAP Health check e/d secure protocol.
	Ldaps SlbNewAdvhcLdapTableLdaps `json:"Ldaps,omitempty"`
	// LDAP Health check user name.
	UserName string `json:"UserName,omitempty"`
	// LDAP Health check password.
	Password string `json:"Password,omitempty"`
	// LDAP Health check base distinguish name.
	BaseObject string `json:"BaseObject,omitempty"`
	// LDAP Health check base distinguish name copy flag.
	Copy string `json:"Copy,omitempty"`
	// When set to the value of 2 (delete), the entire row is deleted.
	// When read, other(1) is returned. Setting the value to anything
	// other than 2(delete) has no effect on the state of the row.
	Delete SlbNewAdvhcLdapTableDelete `json:"Delete,omitempty"`
	// LDAP base DN format.
	BaseFmt SlbNewAdvhcLdapTableBaseFmt `json:"BaseFmt,omitempty"`
}

func (p SlbNewAdvhcLdapTableParams) iMABean() {}
