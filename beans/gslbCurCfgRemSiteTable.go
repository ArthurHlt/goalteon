package beans

import (
	"fmt"
	"reflect"
)

// GslbCurCfgRemSiteTable The GSLB remote site table in the current configuration block.
// Note:This mib is not supported for VX instance of Virtualization.
type GslbCurCfgRemSiteTable struct {
	// The GSLB remote site table index.
	GslbCurCfgRemSiteIndx int32
	Params                *GslbCurCfgRemSiteTableParams
}

func NewGslbCurCfgRemSiteTableList() *GslbCurCfgRemSiteTable {
	return &GslbCurCfgRemSiteTable{}
}

func NewGslbCurCfgRemSiteTable(
	gslbCurCfgRemSiteIndx int32,
	params *GslbCurCfgRemSiteTableParams,
) *GslbCurCfgRemSiteTable {
	return &GslbCurCfgRemSiteTable{
		GslbCurCfgRemSiteIndx: gslbCurCfgRemSiteIndx,
		Params:                params,
	}
}

func (c *GslbCurCfgRemSiteTable) Name() string {
	return "GslbCurCfgRemSiteTable"
}

func (c *GslbCurCfgRemSiteTable) GetParams() BeanType {
	return c.Params
}

func (c *GslbCurCfgRemSiteTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *GslbCurCfgRemSiteTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.GslbCurCfgRemSiteIndx).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.GslbCurCfgRemSiteIndx)
}

type GslbCurCfgRemSiteTableState int32

const (
	GslbCurCfgRemSiteTableState_Enabled     GslbCurCfgRemSiteTableState = 1
	GslbCurCfgRemSiteTableState_Disabled    GslbCurCfgRemSiteTableState = 2
	GslbCurCfgRemSiteTableState_Unsupported GslbCurCfgRemSiteTableState = 2147483647
)

type GslbCurCfgRemSiteTableUpdate int32

const (
	GslbCurCfgRemSiteTableUpdate_Enabled     GslbCurCfgRemSiteTableUpdate = 1
	GslbCurCfgRemSiteTableUpdate_Disabled    GslbCurCfgRemSiteTableUpdate = 2
	GslbCurCfgRemSiteTableUpdate_Unsupported GslbCurCfgRemSiteTableUpdate = 2147483647
)

type GslbCurCfgRemSiteTablePrimaryIPVer int32

const (
	GslbCurCfgRemSiteTablePrimaryIPVer_Ipv4        GslbCurCfgRemSiteTablePrimaryIPVer = 1
	GslbCurCfgRemSiteTablePrimaryIPVer_Ipv6        GslbCurCfgRemSiteTablePrimaryIPVer = 2
	GslbCurCfgRemSiteTablePrimaryIPVer_Unsupported GslbCurCfgRemSiteTablePrimaryIPVer = 2147483647
)

type GslbCurCfgRemSiteTableSecondaryIPVer int32

const (
	GslbCurCfgRemSiteTableSecondaryIPVer_Ipv4        GslbCurCfgRemSiteTableSecondaryIPVer = 1
	GslbCurCfgRemSiteTableSecondaryIPVer_Ipv6        GslbCurCfgRemSiteTableSecondaryIPVer = 2
	GslbCurCfgRemSiteTableSecondaryIPVer_Unsupported GslbCurCfgRemSiteTableSecondaryIPVer = 2147483647
)

type GslbCurCfgRemSiteTablePeer int32

const (
	GslbCurCfgRemSiteTablePeer_Enabled     GslbCurCfgRemSiteTablePeer = 1
	GslbCurCfgRemSiteTablePeer_Disabled    GslbCurCfgRemSiteTablePeer = 2
	GslbCurCfgRemSiteTablePeer_Unsupported GslbCurCfgRemSiteTablePeer = 2147483647
)

type GslbCurCfgRemSiteTableParams struct {
	// The GSLB remote site table index.
	Indx int32 `json:"Indx,omitempty"`
	// The primary IP address of the remote site in the
	// current_configuration block.
	PrimaryIp string `json:"PrimaryIp,omitempty"`
	// The secondary IP address of the remote site in the
	// 	 current_configuration block.
	SecondaryIp string `json:"SecondaryIp,omitempty"`
	// Enable/Disable GSLB for the remote site in the
	// current_configuration block.
	State GslbCurCfgRemSiteTableState `json:"State,omitempty"`
	// Enable/Disable GSLB for the remote site status update in the
	// current_configuration block.
	Update GslbCurCfgRemSiteTableUpdate `json:"Update,omitempty"`
	// The name of the GSLB remote site.
	Name string `json:"Name,omitempty"`
	// The primary IP address version of the remote site.
	PrimaryIPVer GslbCurCfgRemSiteTablePrimaryIPVer `json:"PrimaryIPVer,omitempty"`
	// Primary IPV6 address of the remote site
	// Address should be 4-byte hexadecimal colon notation. Valid IPv6 address should be in
	// any of the following forms xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx or xxxx::xxxx:xxxx:xxxx:xxxx or ::xxxx
	PrimaryIp6 string `json:"PrimaryIp6,omitempty"`
	// The secondary IP address version of the remote site.
	SecondaryIPVer GslbCurCfgRemSiteTableSecondaryIPVer `json:"SecondaryIPVer,omitempty"`
	// Secondary IPV6 address of the remote site
	// Address should be 4-byte hexadecimal colon notation. Valid IPv6 address should be in
	// any of the following forms xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx or xxxx::xxxx:xxxx:xxxx:xxxx or ::xxxx
	SecondaryIp6 string `json:"SecondaryIp6,omitempty"`
	// Enable/Disable treatment of site as vrrp peer device.
	Peer GslbCurCfgRemSiteTablePeer `json:"Peer,omitempty"`
}

func (p GslbCurCfgRemSiteTableParams) iMABean() {}
