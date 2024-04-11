package beans

import (
	"fmt"
	"reflect"
)

// SlbCurAppwallLdapCfgTable Cur web Ldap servers table.
// Note:This mib is not supported on VX instance of Virtualization.
type SlbCurAppwallLdapCfgTable struct {
	// Secure web Ldap server ID
	SlbCurAppwallLdapId string
	Params              *SlbCurAppwallLdapCfgTableParams
}

func NewSlbCurAppwallLdapCfgTableList() *SlbCurAppwallLdapCfgTable {
	return &SlbCurAppwallLdapCfgTable{}
}

func NewSlbCurAppwallLdapCfgTable(
	slbCurAppwallLdapId string,
	params *SlbCurAppwallLdapCfgTableParams,
) *SlbCurAppwallLdapCfgTable {
	return &SlbCurAppwallLdapCfgTable{
		SlbCurAppwallLdapId: slbCurAppwallLdapId,
		Params:              params,
	}
}

func (c *SlbCurAppwallLdapCfgTable) Name() string {
	return "SlbCurAppwallLdapCfgTable"
}

func (c *SlbCurAppwallLdapCfgTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbCurAppwallLdapCfgTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbCurAppwallLdapCfgTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbCurAppwallLdapId).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbCurAppwallLdapId)
}

type SlbCurAppwallLdapCfgTableLdapSSLEnable int32

const (
	SlbCurAppwallLdapCfgTableLdapSSLEnable_Disable     SlbCurAppwallLdapCfgTableLdapSSLEnable = 0
	SlbCurAppwallLdapCfgTableLdapSSLEnable_Enable      SlbCurAppwallLdapCfgTableLdapSSLEnable = 1
	SlbCurAppwallLdapCfgTableLdapSSLEnable_Unsupported SlbCurAppwallLdapCfgTableLdapSSLEnable = 2147483647
)

type SlbCurAppwallLdapCfgTableLdapTLSEnable int32

const (
	SlbCurAppwallLdapCfgTableLdapTLSEnable_Disable     SlbCurAppwallLdapCfgTableLdapTLSEnable = 0
	SlbCurAppwallLdapCfgTableLdapTLSEnable_Enable      SlbCurAppwallLdapCfgTableLdapTLSEnable = 1
	SlbCurAppwallLdapCfgTableLdapTLSEnable_Unsupported SlbCurAppwallLdapCfgTableLdapTLSEnable = 2147483647
)

type SlbCurAppwallLdapCfgTableLdapSrvType int32

const (
	SlbCurAppwallLdapCfgTableLdapSrvType_MicrosoftActDirect SlbCurAppwallLdapCfgTableLdapSrvType = 0
	SlbCurAppwallLdapCfgTableLdapSrvType_RedhatDirect       SlbCurAppwallLdapCfgTableLdapSrvType = 1
	SlbCurAppwallLdapCfgTableLdapSrvType_AppleOpenDirect    SlbCurAppwallLdapCfgTableLdapSrvType = 2
	SlbCurAppwallLdapCfgTableLdapSrvType_OtherDirect        SlbCurAppwallLdapCfgTableLdapSrvType = 3
	SlbCurAppwallLdapCfgTableLdapSrvType_Unsupported        SlbCurAppwallLdapCfgTableLdapSrvType = 2147483647
)

type SlbCurAppwallLdapCfgTableLdapDel int32

const (
	SlbCurAppwallLdapCfgTableLdapDel_Other       SlbCurAppwallLdapCfgTableLdapDel = 1
	SlbCurAppwallLdapCfgTableLdapDel_Delete      SlbCurAppwallLdapCfgTableLdapDel = 2
	SlbCurAppwallLdapCfgTableLdapDel_Unsupported SlbCurAppwallLdapCfgTableLdapDel = 2147483647
)

type SlbCurAppwallLdapCfgTableLdapIgnoreEnable int32

const (
	SlbCurAppwallLdapCfgTableLdapIgnoreEnable_Disable     SlbCurAppwallLdapCfgTableLdapIgnoreEnable = 0
	SlbCurAppwallLdapCfgTableLdapIgnoreEnable_Enable      SlbCurAppwallLdapCfgTableLdapIgnoreEnable = 1
	SlbCurAppwallLdapCfgTableLdapIgnoreEnable_Unsupported SlbCurAppwallLdapCfgTableLdapIgnoreEnable = 2147483647
)

type SlbCurAppwallLdapCfgTableParams struct {
	// Secure web Ldap server ID
	LdapId string `json:"LdapId,omitempty"`
	// Secure web Ldap Primary IP
	LdapPrimaryIpAddress string `json:"LdapPrimaryIpAddress,omitempty"`
	// Secure web Ldap Primary Port
	LdapPrimaryPort uint64 `json:"LdapPrimaryPort,omitempty"`
	// Secure web Ldap Secondary IP address
	LdapSecondaryIpAddress string `json:"LdapSecondaryIpAddress,omitempty"`
	// Secure web Ldap Secondary Port , default 3268
	LdapSecondaryPort uint64 `json:"LdapSecondaryPort,omitempty"`
	// Secure web Ldap Request timeout in seconds, default 5
	LdapReqTimeoutSec uint32 `json:"LdapReqTimeoutSec,omitempty"`
	// Secure web Ldap Host timeout in seconds, default 5
	LdapHostTimeoutSec uint32 `json:"LdapHostTimeoutSec,omitempty"`
	// Secure web Ldap SSL En/Dis.
	LdapSSLEnable SlbCurAppwallLdapCfgTableLdapSSLEnable `json:"LdapSSLEnable,omitempty"`
	// Secure web Ldap TLS En/Dis.
	LdapTLSEnable SlbCurAppwallLdapCfgTableLdapTLSEnable `json:"LdapTLSEnable,omitempty"`
	// Secure web Ldap server type.default microsoftActDirect
	LdapSrvType SlbCurAppwallLdapCfgTableLdapSrvType `json:"LdapSrvType,omitempty"`
	// Secure web Ldap Base path
	LdapBase string `json:"LdapBase,omitempty"`
	// Secure web Ldap User Test
	LdapUserTest string `json:"LdapUserTest,omitempty"`
	// Secure web Ldap Password
	LdapPassTest string `json:"LdapPassTest,omitempty"`
	// By setting the value to delete(2), the entire row is deleted.On GET always other(1)
	LdapDel SlbCurAppwallLdapCfgTableLdapDel `json:"LdapDel,omitempty"`
	// Secure web Ldap Ignore En/Dis.
	LdapIgnoreEnable SlbCurAppwallLdapCfgTableLdapIgnoreEnable `json:"LdapIgnoreEnable,omitempty"`
}

func (p SlbCurAppwallLdapCfgTableParams) iMABean() {}
