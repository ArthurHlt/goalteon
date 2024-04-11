package beans

import (
	"fmt"
	"reflect"
)

// SlbNewAppwallLdapCfgTable New web Ldap servers table.
// Note:This mib is not supported on VX instance of Virtualization.
type SlbNewAppwallLdapCfgTable struct {
	// Secure web Ldap server ID
	SlbNewAppwallLdapId string
	Params              *SlbNewAppwallLdapCfgTableParams
}

func NewSlbNewAppwallLdapCfgTableList() *SlbNewAppwallLdapCfgTable {
	return &SlbNewAppwallLdapCfgTable{}
}

func NewSlbNewAppwallLdapCfgTable(
	slbNewAppwallLdapId string,
	params *SlbNewAppwallLdapCfgTableParams,
) *SlbNewAppwallLdapCfgTable {
	return &SlbNewAppwallLdapCfgTable{
		SlbNewAppwallLdapId: slbNewAppwallLdapId,
		Params:              params,
	}
}

func (c *SlbNewAppwallLdapCfgTable) Name() string {
	return "SlbNewAppwallLdapCfgTable"
}

func (c *SlbNewAppwallLdapCfgTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbNewAppwallLdapCfgTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbNewAppwallLdapCfgTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbNewAppwallLdapId).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbNewAppwallLdapId)
}

type SlbNewAppwallLdapCfgTableLdapSSLEnable int32

const (
	SlbNewAppwallLdapCfgTableLdapSSLEnable_Disable     SlbNewAppwallLdapCfgTableLdapSSLEnable = 0
	SlbNewAppwallLdapCfgTableLdapSSLEnable_Enable      SlbNewAppwallLdapCfgTableLdapSSLEnable = 1
	SlbNewAppwallLdapCfgTableLdapSSLEnable_Unsupported SlbNewAppwallLdapCfgTableLdapSSLEnable = 2147483647
)

type SlbNewAppwallLdapCfgTableLdapTLSEnable int32

const (
	SlbNewAppwallLdapCfgTableLdapTLSEnable_Disable     SlbNewAppwallLdapCfgTableLdapTLSEnable = 0
	SlbNewAppwallLdapCfgTableLdapTLSEnable_Enable      SlbNewAppwallLdapCfgTableLdapTLSEnable = 1
	SlbNewAppwallLdapCfgTableLdapTLSEnable_Unsupported SlbNewAppwallLdapCfgTableLdapTLSEnable = 2147483647
)

type SlbNewAppwallLdapCfgTableLdapSrvType int32

const (
	SlbNewAppwallLdapCfgTableLdapSrvType_MicrosoftActDirect SlbNewAppwallLdapCfgTableLdapSrvType = 0
	SlbNewAppwallLdapCfgTableLdapSrvType_RedhatDirect       SlbNewAppwallLdapCfgTableLdapSrvType = 1
	SlbNewAppwallLdapCfgTableLdapSrvType_AppleOpenDirect    SlbNewAppwallLdapCfgTableLdapSrvType = 2
	SlbNewAppwallLdapCfgTableLdapSrvType_OtherDirect        SlbNewAppwallLdapCfgTableLdapSrvType = 3
	SlbNewAppwallLdapCfgTableLdapSrvType_Unsupported        SlbNewAppwallLdapCfgTableLdapSrvType = 2147483647
)

type SlbNewAppwallLdapCfgTableLdapDel int32

const (
	SlbNewAppwallLdapCfgTableLdapDel_Other       SlbNewAppwallLdapCfgTableLdapDel = 1
	SlbNewAppwallLdapCfgTableLdapDel_Delete      SlbNewAppwallLdapCfgTableLdapDel = 2
	SlbNewAppwallLdapCfgTableLdapDel_Unsupported SlbNewAppwallLdapCfgTableLdapDel = 2147483647
)

type SlbNewAppwallLdapCfgTableLdapIgnoreEnable int32

const (
	SlbNewAppwallLdapCfgTableLdapIgnoreEnable_Disable     SlbNewAppwallLdapCfgTableLdapIgnoreEnable = 0
	SlbNewAppwallLdapCfgTableLdapIgnoreEnable_Enable      SlbNewAppwallLdapCfgTableLdapIgnoreEnable = 1
	SlbNewAppwallLdapCfgTableLdapIgnoreEnable_Unsupported SlbNewAppwallLdapCfgTableLdapIgnoreEnable = 2147483647
)

type SlbNewAppwallLdapCfgTableParams struct {
	// Secure web Ldap server ID
	LdapId string `json:"LdapId,omitempty"`
	// Secure web Ldap Primary IP
	LdapPrimaryIpAddress string `json:"LdapPrimaryIpAddress,omitempty"`
	// Secure web Ldap Primary Port
	LdapPrimaryPort uint64 `json:"LdapPrimaryPort,omitempty"`
	// Secure web Ldap Secondary IP address
	LdapSecondaryIpAddress string `json:"LdapSecondaryIpAddress,omitempty"`
	// Secure web Ldap Secondary Port
	LdapSecondaryPort uint64 `json:"LdapSecondaryPort,omitempty"`
	// Secure web Ldap Request timeout in seconds
	LdapReqTimeoutSec uint32 `json:"LdapReqTimeoutSec,omitempty"`
	// Secure web Ldap Host timeout in seconds
	LdapHostTimeoutSec uint32 `json:"LdapHostTimeoutSec,omitempty"`
	// Secure web Ldap SSL En/Dis.
	LdapSSLEnable SlbNewAppwallLdapCfgTableLdapSSLEnable `json:"LdapSSLEnable,omitempty"`
	// Secure web Ldap TLS En/Dis.
	LdapTLSEnable SlbNewAppwallLdapCfgTableLdapTLSEnable `json:"LdapTLSEnable,omitempty"`
	// Secure web Ldap server type.
	LdapSrvType SlbNewAppwallLdapCfgTableLdapSrvType `json:"LdapSrvType,omitempty"`
	// Secure web Ldap Base path
	LdapBase string `json:"LdapBase,omitempty"`
	// Secure web Ldap User Test
	LdapUserTest string `json:"LdapUserTest,omitempty"`
	// Secure web Ldap Password
	LdapPassTest string `json:"LdapPassTest,omitempty"`
	// By setting the value to delete(2), the entire row is deleted.On GET always other(1)
	LdapDel SlbNewAppwallLdapCfgTableLdapDel `json:"LdapDel,omitempty"`
	// Secure web Ldap Ignore En/Dis.
	LdapIgnoreEnable SlbNewAppwallLdapCfgTableLdapIgnoreEnable `json:"LdapIgnoreEnable,omitempty"`
}

func (p SlbNewAppwallLdapCfgTableParams) iMABean() {}
