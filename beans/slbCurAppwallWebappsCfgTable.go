package beans

import (
	"fmt"
	"reflect"
)

// SlbCurAppwallWebappsCfgTable Current secure web applications table.
// Note:This mib is not supported on VX instance of Virtualization.
type SlbCurAppwallWebappsCfgTable struct {
	// Secure web application ID
	SlbCurAppwallWebappId string
	Params                *SlbCurAppwallWebappsCfgTableParams
}

func NewSlbCurAppwallWebappsCfgTableList() *SlbCurAppwallWebappsCfgTable {
	return &SlbCurAppwallWebappsCfgTable{}
}

func NewSlbCurAppwallWebappsCfgTable(
	slbCurAppwallWebappId string,
	params *SlbCurAppwallWebappsCfgTableParams,
) *SlbCurAppwallWebappsCfgTable {
	return &SlbCurAppwallWebappsCfgTable{
		SlbCurAppwallWebappId: slbCurAppwallWebappId,
		Params:                params,
	}
}

func (c *SlbCurAppwallWebappsCfgTable) Name() string {
	return "SlbCurAppwallWebappsCfgTable"
}

func (c *SlbCurAppwallWebappsCfgTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbCurAppwallWebappsCfgTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbCurAppwallWebappsCfgTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbCurAppwallWebappId).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbCurAppwallWebappId)
}

type SlbCurAppwallWebappsCfgTableWebappEnable int32

const (
	SlbCurAppwallWebappsCfgTableWebappEnable_Disable     SlbCurAppwallWebappsCfgTableWebappEnable = 0
	SlbCurAppwallWebappsCfgTableWebappEnable_Enable      SlbCurAppwallWebappsCfgTableWebappEnable = 1
	SlbCurAppwallWebappsCfgTableWebappEnable_Unsupported SlbCurAppwallWebappsCfgTableWebappEnable = 2147483647
)

type SlbCurAppwallWebappsCfgTableWebappEnableWaf int32

const (
	SlbCurAppwallWebappsCfgTableWebappEnableWaf_Disable     SlbCurAppwallWebappsCfgTableWebappEnableWaf = 0
	SlbCurAppwallWebappsCfgTableWebappEnableWaf_Enable      SlbCurAppwallWebappsCfgTableWebappEnableWaf = 1
	SlbCurAppwallWebappsCfgTableWebappEnableWaf_Unsupported SlbCurAppwallWebappsCfgTableWebappEnableWaf = 2147483647
)

type SlbCurAppwallWebappsCfgTableWebappEnableAuthSso int32

const (
	SlbCurAppwallWebappsCfgTableWebappEnableAuthSso_Disable     SlbCurAppwallWebappsCfgTableWebappEnableAuthSso = 0
	SlbCurAppwallWebappsCfgTableWebappEnableAuthSso_Enable      SlbCurAppwallWebappsCfgTableWebappEnableAuthSso = 1
	SlbCurAppwallWebappsCfgTableWebappEnableAuthSso_Unsupported SlbCurAppwallWebappsCfgTableWebappEnableAuthSso = 2147483647
)

type SlbCurAppwallWebappsCfgTableWebappMode int32

const (
	SlbCurAppwallWebappsCfgTableWebappMode_Oop         SlbCurAppwallWebappsCfgTableWebappMode = 0
	SlbCurAppwallWebappsCfgTableWebappMode_Inline      SlbCurAppwallWebappsCfgTableWebappMode = 1
	SlbCurAppwallWebappsCfgTableWebappMode_Unsupported SlbCurAppwallWebappsCfgTableWebappMode = 2147483647
)

type SlbCurAppwallWebappsCfgTableParams struct {
	// Secure web application ID
	WebappId string `json:"WebappId,omitempty"`
	// Secure web application name.
	WebappName string `json:"WebappName,omitempty"`
	// Secure web application Ena \ Dis.
	WebappEnable SlbCurAppwallWebappsCfgTableWebappEnable `json:"WebappEnable,omitempty"`
	// Secure web application WAF Ena \ Dis.
	WebappEnableWaf SlbCurAppwallWebappsCfgTableWebappEnableWaf `json:"WebappEnableWaf,omitempty"`
	// Secure web application Authentication SSO Ena \ Dis.
	WebappEnableAuthSso SlbCurAppwallWebappsCfgTableWebappEnableAuthSso `json:"WebappEnableAuthSso,omitempty"`
	// Secure web application mode.
	WebappMode SlbCurAppwallWebappsCfgTableWebappMode `json:"WebappMode,omitempty"`
}

func (p SlbCurAppwallWebappsCfgTableParams) iMABean() {}
