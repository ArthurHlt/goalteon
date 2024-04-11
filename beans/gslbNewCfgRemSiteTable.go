package beans

import (
	"fmt"
	"reflect"
)

// GslbNewCfgRemSiteTable The GSLB remote site table in the current configuration block.
// Note:This mib is not supported for VX instance of Virtualization.
type GslbNewCfgRemSiteTable struct {
	// The GSLB remote site table index.
	GslbNewCfgRemSiteIndx int32
	Params                *GslbNewCfgRemSiteTableParams
}

func NewGslbNewCfgRemSiteTableList() *GslbNewCfgRemSiteTable {
	return &GslbNewCfgRemSiteTable{}
}

func NewGslbNewCfgRemSiteTable(
	gslbNewCfgRemSiteIndx int32,
	params *GslbNewCfgRemSiteTableParams,
) *GslbNewCfgRemSiteTable {
	return &GslbNewCfgRemSiteTable{
		GslbNewCfgRemSiteIndx: gslbNewCfgRemSiteIndx,
		Params:                params,
	}
}

func (c *GslbNewCfgRemSiteTable) Name() string {
	return "GslbNewCfgRemSiteTable"
}

func (c *GslbNewCfgRemSiteTable) GetParams() BeanType {
	return c.Params
}

func (c *GslbNewCfgRemSiteTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *GslbNewCfgRemSiteTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.GslbNewCfgRemSiteIndx).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.GslbNewCfgRemSiteIndx)
}

type GslbNewCfgRemSiteTableState int32

const (
	GslbNewCfgRemSiteTableState_Enabled     GslbNewCfgRemSiteTableState = 1
	GslbNewCfgRemSiteTableState_Disabled    GslbNewCfgRemSiteTableState = 2
	GslbNewCfgRemSiteTableState_Unsupported GslbNewCfgRemSiteTableState = 2147483647
)

type GslbNewCfgRemSiteTableUpdate int32

const (
	GslbNewCfgRemSiteTableUpdate_Enabled     GslbNewCfgRemSiteTableUpdate = 1
	GslbNewCfgRemSiteTableUpdate_Disabled    GslbNewCfgRemSiteTableUpdate = 2
	GslbNewCfgRemSiteTableUpdate_Unsupported GslbNewCfgRemSiteTableUpdate = 2147483647
)

type GslbNewCfgRemSiteTableDelete int32

const (
	GslbNewCfgRemSiteTableDelete_Other       GslbNewCfgRemSiteTableDelete = 1
	GslbNewCfgRemSiteTableDelete_Delete      GslbNewCfgRemSiteTableDelete = 2
	GslbNewCfgRemSiteTableDelete_Unsupported GslbNewCfgRemSiteTableDelete = 2147483647
)

type GslbNewCfgRemSiteTablePrimaryIPVer int32

const (
	GslbNewCfgRemSiteTablePrimaryIPVer_Ipv4        GslbNewCfgRemSiteTablePrimaryIPVer = 1
	GslbNewCfgRemSiteTablePrimaryIPVer_Ipv6        GslbNewCfgRemSiteTablePrimaryIPVer = 2
	GslbNewCfgRemSiteTablePrimaryIPVer_Unsupported GslbNewCfgRemSiteTablePrimaryIPVer = 2147483647
)

type GslbNewCfgRemSiteTableSecondaryIPVer int32

const (
	GslbNewCfgRemSiteTableSecondaryIPVer_Ipv4        GslbNewCfgRemSiteTableSecondaryIPVer = 1
	GslbNewCfgRemSiteTableSecondaryIPVer_Ipv6        GslbNewCfgRemSiteTableSecondaryIPVer = 2
	GslbNewCfgRemSiteTableSecondaryIPVer_Unsupported GslbNewCfgRemSiteTableSecondaryIPVer = 2147483647
)

type GslbNewCfgRemSiteTablePeer int32

const (
	GslbNewCfgRemSiteTablePeer_Enabled     GslbNewCfgRemSiteTablePeer = 1
	GslbNewCfgRemSiteTablePeer_Disabled    GslbNewCfgRemSiteTablePeer = 2
	GslbNewCfgRemSiteTablePeer_Unsupported GslbNewCfgRemSiteTablePeer = 2147483647
)

type GslbNewCfgRemSiteTableParams struct {
	// The GSLB remote site table index.
	Indx int32 `json:"Indx,omitempty"`
	// The primary IP address of the remote site in the new_configuration
	// block.
	PrimaryIp string `json:"PrimaryIp,omitempty"`
	// The secondary IP address of the remote site in the new_configuration
	// block.
	SecondaryIp string `json:"SecondaryIp,omitempty"`
	// Enable/Disable GSLB for the remote site in the new_configuration
	// block.
	State GslbNewCfgRemSiteTableState `json:"State,omitempty"`
	// Enable/Disable GSLB for the remote site status update in the
	// new_configuration block.
	Update GslbNewCfgRemSiteTableUpdate `json:"Update,omitempty"`
	// By setting the value to delete(2), the entire row is deleted.
	Delete GslbNewCfgRemSiteTableDelete `json:"Delete,omitempty"`
	// The name of the GSLB remote site.
	Name string `json:"Name,omitempty"`
	// The primary IP address version of the remote site.
	PrimaryIPVer GslbNewCfgRemSiteTablePrimaryIPVer `json:"PrimaryIPVer,omitempty"`
	// Primary IPV6 address of the remote site
	// Address should be 4-byte hexadecimal colon notation. Valid IPv6 address should be in
	// any of the following forms xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx or xxxx::xxxx:xxxx:xxxx:xxxx or ::xxxx
	PrimaryIp6 string `json:"PrimaryIp6,omitempty"`
	// The secondary IP address version of the remote site.
	SecondaryIPVer GslbNewCfgRemSiteTableSecondaryIPVer `json:"SecondaryIPVer,omitempty"`
	// Secondary IPV6 address of the remote site
	// Address should be 4-byte hexadecimal colon notation. Valid IPv6 address should be in
	// any of the following forms xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx or xxxx::xxxx:xxxx:xxxx:xxxx or ::xxxx
	SecondaryIp6 string `json:"SecondaryIp6,omitempty"`
	// Enable/Disable treatment of site as vrrp peer device.
	Peer GslbNewCfgRemSiteTablePeer `json:"Peer,omitempty"`
}

func (p GslbNewCfgRemSiteTableParams) iMABean() {}
