package beans

import (
	"fmt"
	"reflect"
)

// SlbNewAppwallRadiusCfgTable New secure web Radius servers table.
// Note:This mib is not supported on VX instance of Virtualization.
type SlbNewAppwallRadiusCfgTable struct {
	// Secure web Radius server ID
	SlbNewAppwallRadiusId string
	Params                *SlbNewAppwallRadiusCfgTableParams
}

func NewSlbNewAppwallRadiusCfgTableList() *SlbNewAppwallRadiusCfgTable {
	return &SlbNewAppwallRadiusCfgTable{}
}

func NewSlbNewAppwallRadiusCfgTable(
	slbNewAppwallRadiusId string,
	params *SlbNewAppwallRadiusCfgTableParams,
) *SlbNewAppwallRadiusCfgTable {
	return &SlbNewAppwallRadiusCfgTable{
		SlbNewAppwallRadiusId: slbNewAppwallRadiusId,
		Params:                params,
	}
}

func (c *SlbNewAppwallRadiusCfgTable) Name() string {
	return "SlbNewAppwallRadiusCfgTable"
}

func (c *SlbNewAppwallRadiusCfgTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbNewAppwallRadiusCfgTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbNewAppwallRadiusCfgTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbNewAppwallRadiusId).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbNewAppwallRadiusId)
}

type SlbNewAppwallRadiusCfgTableRadiusDel int32

const (
	SlbNewAppwallRadiusCfgTableRadiusDel_Other       SlbNewAppwallRadiusCfgTableRadiusDel = 1
	SlbNewAppwallRadiusCfgTableRadiusDel_Delete      SlbNewAppwallRadiusCfgTableRadiusDel = 2
	SlbNewAppwallRadiusCfgTableRadiusDel_Unsupported SlbNewAppwallRadiusCfgTableRadiusDel = 2147483647
)

type SlbNewAppwallRadiusCfgTableParams struct {
	// Secure web Radius server ID
	RadiusId string `json:"RadiusId,omitempty"`
	// Secure web Radius Primary IP address
	RadiusPrimaryIpAddress string `json:"RadiusPrimaryIpAddress,omitempty"`
	// Secure web Radius Primary Port
	RadiusPrimaryPort uint64 `json:"RadiusPrimaryPort,omitempty"`
	// Secure web Radius Primary secret
	RadiusPrimarySecret string `json:"RadiusPrimarySecret,omitempty"`
	// Secure web Radius Secondary IP address
	RadiusSecondaryIpAddress string `json:"RadiusSecondaryIpAddress,omitempty"`
	// Secure web Radius Secondary Port
	RadiusSecondaryPort uint64 `json:"RadiusSecondaryPort,omitempty"`
	// Secure web Radius Secondary secret
	RadiusSecondarySecret string `json:"RadiusSecondarySecret,omitempty"`
	// Secure web Radius retries default 3
	RadiusRetries uint64 `json:"RadiusRetries,omitempty"`
	// Secure web Radius Timeout default 5 sec
	RadiusTimeout uint32 `json:"RadiusTimeout,omitempty"`
	// By setting the value to delete(2), the entire row is deleted.On GET always other(1)
	RadiusDel SlbNewAppwallRadiusCfgTableRadiusDel `json:"RadiusDel,omitempty"`
	// Secure web Radius Tertiary IP address
	RadiusTertiaryIpAddress string `json:"RadiusTertiaryIpAddress,omitempty"`
	// Secure web Radius Tertiary Port
	RadiusTertiaryPort uint64 `json:"RadiusTertiaryPort,omitempty"`
	// Secure web Radius Tertiary secret
	RadiusTertiarySecret string `json:"RadiusTertiarySecret,omitempty"`
}

func (p SlbNewAppwallRadiusCfgTableParams) iMABean() {}
