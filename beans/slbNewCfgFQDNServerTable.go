package beans

import (
	"fmt"
	"reflect"
)

// SlbNewCfgFQDNServerTable The table of FQDN Real Server configuration in the new_config.
// Note:This mib is not supported for VX instance of Virtualization.
type SlbNewCfgFQDNServerTable struct {
	// The Fqdn server index
	SlbNewCfgFQDNServerIdIndex string
	Params                     *SlbNewCfgFQDNServerTableParams
}

func NewSlbNewCfgFQDNServerTableList() *SlbNewCfgFQDNServerTable {
	return &SlbNewCfgFQDNServerTable{}
}

func NewSlbNewCfgFQDNServerTable(
	slbNewCfgFQDNServerIdIndex string,
	params *SlbNewCfgFQDNServerTableParams,
) *SlbNewCfgFQDNServerTable {
	return &SlbNewCfgFQDNServerTable{
		SlbNewCfgFQDNServerIdIndex: slbNewCfgFQDNServerIdIndex,
		Params:                     params,
	}
}

func (c *SlbNewCfgFQDNServerTable) Name() string {
	return "SlbNewCfgFQDNServerTable"
}

func (c *SlbNewCfgFQDNServerTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbNewCfgFQDNServerTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbNewCfgFQDNServerTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbNewCfgFQDNServerIdIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbNewCfgFQDNServerIdIndex)
}

type SlbNewCfgFQDNServerTableIpVers int32

const (
	SlbNewCfgFQDNServerTableIpVers_Ipv4        SlbNewCfgFQDNServerTableIpVers = 1
	SlbNewCfgFQDNServerTableIpVers_Ipv6        SlbNewCfgFQDNServerTableIpVers = 2
	SlbNewCfgFQDNServerTableIpVers_Unsupported SlbNewCfgFQDNServerTableIpVers = 2147483647
)

type SlbNewCfgFQDNServerTableState int32

const (
	SlbNewCfgFQDNServerTableState_Enabled     SlbNewCfgFQDNServerTableState = 2
	SlbNewCfgFQDNServerTableState_Disabled    SlbNewCfgFQDNServerTableState = 3
	SlbNewCfgFQDNServerTableState_Unsupported SlbNewCfgFQDNServerTableState = 2147483647
)

type SlbNewCfgFQDNServerTableDelete int32

const (
	SlbNewCfgFQDNServerTableDelete_Other       SlbNewCfgFQDNServerTableDelete = 1
	SlbNewCfgFQDNServerTableDelete_Delete      SlbNewCfgFQDNServerTableDelete = 2
	SlbNewCfgFQDNServerTableDelete_Unsupported SlbNewCfgFQDNServerTableDelete = 2147483647
)

type SlbNewCfgFQDNServerTableMode int32

const (
	SlbNewCfgFQDNServerTableMode_Fqdn        SlbNewCfgFQDNServerTableMode = 1
	SlbNewCfgFQDNServerTableMode_Cscale      SlbNewCfgFQDNServerTableMode = 2
	SlbNewCfgFQDNServerTableMode_Unsupported SlbNewCfgFQDNServerTableMode = 2147483647
)

type SlbNewCfgFQDNServerTableParams struct {
	// The Fqdn server index
	IdIndex string `json:"IdIndex,omitempty"`
	// The Fqdn field of the server
	FQDN string `json:"FQDN,omitempty"`
	// The type of IP address.
	IpVers SlbNewCfgFQDNServerTableIpVers `json:"IpVers,omitempty"`
	// minimum time to live.
	TTL uint32 `json:"TTL,omitempty"`
	// FQDN Server Group Id.
	Group string `json:"Group,omitempty"`
	// Template Real Server Id.
	Templ string `json:"Templ,omitempty"`
	// Enable or disable the server.
	State SlbNewCfgFQDNServerTableState `json:"State,omitempty"`
	// By setting the value to delete(2), the entire row is deleted.
	Delete SlbNewCfgFQDNServerTableDelete `json:"Delete,omitempty"`
	// Set the server mode to FQDN or Auto Scaling (Cloud platform).
	Mode SlbNewCfgFQDNServerTableMode `json:"Mode,omitempty"`
	// Resource Group Name (Cloud platform).
	RGName string `json:"RGName,omitempty"`
}

func (p SlbNewCfgFQDNServerTableParams) iMABean() {}
