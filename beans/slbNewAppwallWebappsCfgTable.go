package beans

import (
	"fmt"
	"reflect"
)

// SlbNewAppwallWebappsCfgTable New secure web applications table.
// Note:This mib is not supported on VX instance of Virtualization.
type SlbNewAppwallWebappsCfgTable struct {
	// Secure web application ID
	SlbNewAppwallWebappId string
	Params                *SlbNewAppwallWebappsCfgTableParams
}

func NewSlbNewAppwallWebappsCfgTableList() *SlbNewAppwallWebappsCfgTable {
	return &SlbNewAppwallWebappsCfgTable{}
}

func NewSlbNewAppwallWebappsCfgTable(
	slbNewAppwallWebappId string,
	params *SlbNewAppwallWebappsCfgTableParams,
) *SlbNewAppwallWebappsCfgTable {
	return &SlbNewAppwallWebappsCfgTable{
		SlbNewAppwallWebappId: slbNewAppwallWebappId,
		Params:                params,
	}
}

func (c *SlbNewAppwallWebappsCfgTable) Name() string {
	return "SlbNewAppwallWebappsCfgTable"
}

func (c *SlbNewAppwallWebappsCfgTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbNewAppwallWebappsCfgTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbNewAppwallWebappsCfgTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbNewAppwallWebappId).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbNewAppwallWebappId)
}

type SlbNewAppwallWebappsCfgTableWebappEnable int32

const (
	SlbNewAppwallWebappsCfgTableWebappEnable_Disable     SlbNewAppwallWebappsCfgTableWebappEnable = 0
	SlbNewAppwallWebappsCfgTableWebappEnable_Enable      SlbNewAppwallWebappsCfgTableWebappEnable = 1
	SlbNewAppwallWebappsCfgTableWebappEnable_Unsupported SlbNewAppwallWebappsCfgTableWebappEnable = 2147483647
)

type SlbNewAppwallWebappsCfgTableWebappEnableWaf int32

const (
	SlbNewAppwallWebappsCfgTableWebappEnableWaf_Disable     SlbNewAppwallWebappsCfgTableWebappEnableWaf = 0
	SlbNewAppwallWebappsCfgTableWebappEnableWaf_Enable      SlbNewAppwallWebappsCfgTableWebappEnableWaf = 1
	SlbNewAppwallWebappsCfgTableWebappEnableWaf_Unsupported SlbNewAppwallWebappsCfgTableWebappEnableWaf = 2147483647
)

type SlbNewAppwallWebappsCfgTableWebappEnableAuthSso int32

const (
	SlbNewAppwallWebappsCfgTableWebappEnableAuthSso_Disable     SlbNewAppwallWebappsCfgTableWebappEnableAuthSso = 0
	SlbNewAppwallWebappsCfgTableWebappEnableAuthSso_Enable      SlbNewAppwallWebappsCfgTableWebappEnableAuthSso = 1
	SlbNewAppwallWebappsCfgTableWebappEnableAuthSso_Unsupported SlbNewAppwallWebappsCfgTableWebappEnableAuthSso = 2147483647
)

type SlbNewAppwallWebappsCfgTableWebappMode int32

const (
	SlbNewAppwallWebappsCfgTableWebappMode_Oop         SlbNewAppwallWebappsCfgTableWebappMode = 0
	SlbNewAppwallWebappsCfgTableWebappMode_Inline      SlbNewAppwallWebappsCfgTableWebappMode = 1
	SlbNewAppwallWebappsCfgTableWebappMode_Unsupported SlbNewAppwallWebappsCfgTableWebappMode = 2147483647
)

type SlbNewAppwallWebappsCfgTableWebappDel int32

const (
	SlbNewAppwallWebappsCfgTableWebappDel_Other       SlbNewAppwallWebappsCfgTableWebappDel = 1
	SlbNewAppwallWebappsCfgTableWebappDel_Delete      SlbNewAppwallWebappsCfgTableWebappDel = 2
	SlbNewAppwallWebappsCfgTableWebappDel_Unsupported SlbNewAppwallWebappsCfgTableWebappDel = 2147483647
)

type SlbNewAppwallWebappsCfgTableWebappCldirection int32

const (
	SlbNewAppwallWebappsCfgTableWebappCldirection_Request       SlbNewAppwallWebappsCfgTableWebappCldirection = 0
	SlbNewAppwallWebappsCfgTableWebappCldirection_Response      SlbNewAppwallWebappsCfgTableWebappCldirection = 1
	SlbNewAppwallWebappsCfgTableWebappCldirection_Bidirectional SlbNewAppwallWebappsCfgTableWebappCldirection = 2
	SlbNewAppwallWebappsCfgTableWebappCldirection_Unsupported   SlbNewAppwallWebappsCfgTableWebappCldirection = 2147483647
)

type SlbNewAppwallWebappsCfgTableParams struct {
	// Secure web application ID
	WebappId string `json:"WebappId,omitempty"`
	// Secure web application name.
	WebappName string `json:"WebappName,omitempty"`
	// Secure web application Ena \ Dis.
	WebappEnable SlbNewAppwallWebappsCfgTableWebappEnable `json:"WebappEnable,omitempty"`
	// Secure web application WAF Ena \ Dis.
	WebappEnableWaf SlbNewAppwallWebappsCfgTableWebappEnableWaf `json:"WebappEnableWaf,omitempty"`
	// Secure web application Authentication SSO Ena \ Dis.
	WebappEnableAuthSso SlbNewAppwallWebappsCfgTableWebappEnableAuthSso `json:"WebappEnableAuthSso,omitempty"`
	// Secure web application mode.
	WebappMode SlbNewAppwallWebappsCfgTableWebappMode `json:"WebappMode,omitempty"`
	// By setting the value to delete(2), the entire row is deleted.
	WebappDel SlbNewAppwallWebappsCfgTableWebappDel `json:"WebappDel,omitempty"`
	// Set timeout.
	WebappCltimeout uint32 `json:"WebappCltimeout,omitempty"`
	// Set direction.
	WebappCldirection SlbNewAppwallWebappsCfgTableWebappCldirection `json:"WebappCldirection,omitempty"`
}

func (p SlbNewAppwallWebappsCfgTableParams) iMABean() {}
