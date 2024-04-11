package beans

import (
	"fmt"
	"reflect"
)

// SlbCurAppwallRadiusCfgTable Current secure web Radius servers table.
// Note:This mib is not supported on VX instance of Virtualization.
type SlbCurAppwallRadiusCfgTable struct {
	// Secure web Radius server ID
	SlbCurAppwallRadiusId string
	Params                *SlbCurAppwallRadiusCfgTableParams
}

func NewSlbCurAppwallRadiusCfgTableList() *SlbCurAppwallRadiusCfgTable {
	return &SlbCurAppwallRadiusCfgTable{}
}

func NewSlbCurAppwallRadiusCfgTable(
	slbCurAppwallRadiusId string,
	params *SlbCurAppwallRadiusCfgTableParams,
) *SlbCurAppwallRadiusCfgTable {
	return &SlbCurAppwallRadiusCfgTable{
		SlbCurAppwallRadiusId: slbCurAppwallRadiusId,
		Params:                params,
	}
}

func (c *SlbCurAppwallRadiusCfgTable) Name() string {
	return "SlbCurAppwallRadiusCfgTable"
}

func (c *SlbCurAppwallRadiusCfgTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbCurAppwallRadiusCfgTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbCurAppwallRadiusCfgTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbCurAppwallRadiusId).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbCurAppwallRadiusId)
}

type SlbCurAppwallRadiusCfgTableRadiusDel int32

const (
	SlbCurAppwallRadiusCfgTableRadiusDel_Other       SlbCurAppwallRadiusCfgTableRadiusDel = 1
	SlbCurAppwallRadiusCfgTableRadiusDel_Delete      SlbCurAppwallRadiusCfgTableRadiusDel = 2
	SlbCurAppwallRadiusCfgTableRadiusDel_Unsupported SlbCurAppwallRadiusCfgTableRadiusDel = 2147483647
)

type SlbCurAppwallRadiusCfgTableParams struct {
	// Secure web Radius server ID
	RadiusId string `json:"RadiusId,omitempty"`
	// Secure web Radius Primary IP
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
	// Secure web Radius retries, default 3
	RadiusRetries uint64 `json:"RadiusRetries,omitempty"`
	// Secure web Radius Timeout default 5 sec
	RadiusTimeout uint32 `json:"RadiusTimeout,omitempty"`
	// By setting the value to delete(2), the entire row is deleted. On GET always other(1)
	RadiusDel SlbCurAppwallRadiusCfgTableRadiusDel `json:"RadiusDel,omitempty"`
	// Secure web Radius Tertiary IP address
	RadiusTertiaryIpAddress string `json:"RadiusTertiaryIpAddress,omitempty"`
	// Secure web Radius Tertiary Port
	RadiusTertiaryPort uint64 `json:"RadiusTertiaryPort,omitempty"`
	// Secure web Radius Tertiary secret
	RadiusTertiarySecret string `json:"RadiusTertiarySecret,omitempty"`
}

func (p SlbCurAppwallRadiusCfgTableParams) iMABean() {}
