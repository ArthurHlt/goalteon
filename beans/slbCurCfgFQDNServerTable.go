package beans

import (
	"fmt"
	"reflect"
)

// SlbCurCfgFQDNServerTable The table of FQDN Real Server configuration in the Cur_config.
// Note:This mib is not supported for VX instance of Virtualization.
type SlbCurCfgFQDNServerTable struct {
	// The Fqdn server index
	SlbCurCfgFQDNServerIdIndex string
	Params                     *SlbCurCfgFQDNServerTableParams
}

func NewSlbCurCfgFQDNServerTableList() *SlbCurCfgFQDNServerTable {
	return &SlbCurCfgFQDNServerTable{}
}

func NewSlbCurCfgFQDNServerTable(
	slbCurCfgFQDNServerIdIndex string,
	params *SlbCurCfgFQDNServerTableParams,
) *SlbCurCfgFQDNServerTable {
	return &SlbCurCfgFQDNServerTable{
		SlbCurCfgFQDNServerIdIndex: slbCurCfgFQDNServerIdIndex,
		Params:                     params,
	}
}

func (c *SlbCurCfgFQDNServerTable) Name() string {
	return "SlbCurCfgFQDNServerTable"
}

func (c *SlbCurCfgFQDNServerTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbCurCfgFQDNServerTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbCurCfgFQDNServerTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbCurCfgFQDNServerIdIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbCurCfgFQDNServerIdIndex)
}

type SlbCurCfgFQDNServerTableIpVers int32

const (
	SlbCurCfgFQDNServerTableIpVers_Ipv4        SlbCurCfgFQDNServerTableIpVers = 1
	SlbCurCfgFQDNServerTableIpVers_Ipv6        SlbCurCfgFQDNServerTableIpVers = 2
	SlbCurCfgFQDNServerTableIpVers_Unsupported SlbCurCfgFQDNServerTableIpVers = 2147483647
)

type SlbCurCfgFQDNServerTableState int32

const (
	SlbCurCfgFQDNServerTableState_Enabled     SlbCurCfgFQDNServerTableState = 2
	SlbCurCfgFQDNServerTableState_Disabled    SlbCurCfgFQDNServerTableState = 3
	SlbCurCfgFQDNServerTableState_Unsupported SlbCurCfgFQDNServerTableState = 2147483647
)

type SlbCurCfgFQDNServerTableMode int32

const (
	SlbCurCfgFQDNServerTableMode_Fqdn        SlbCurCfgFQDNServerTableMode = 1
	SlbCurCfgFQDNServerTableMode_Cscale      SlbCurCfgFQDNServerTableMode = 2
	SlbCurCfgFQDNServerTableMode_Unsupported SlbCurCfgFQDNServerTableMode = 2147483647
)

type SlbCurCfgFQDNServerTableParams struct {
	// The Fqdn server index
	IdIndex string `json:"IdIndex,omitempty"`
	// The Fqdn field of the server
	FQDN string `json:"FQDN,omitempty"`
	// The type of IP address.
	IpVers SlbCurCfgFQDNServerTableIpVers `json:"IpVers,omitempty"`
	// minimum time to live.
	TTL uint32 `json:"TTL,omitempty"`
	// FQDN Server Group Id.
	Group string `json:"Group,omitempty"`
	// Template Real Server Id.
	Templ string `json:"Templ,omitempty"`
	// Enable or disable the server.
	State SlbCurCfgFQDNServerTableState `json:"State,omitempty"`
	// Set the server mode to FQDN or Auto Scaling (Cloud platform).
	Mode SlbCurCfgFQDNServerTableMode `json:"Mode,omitempty"`
	// Resource Group Name (Cloud platform).
	RGName string `json:"RGName,omitempty"`
}

func (p SlbCurCfgFQDNServerTableParams) iMABean() {}
