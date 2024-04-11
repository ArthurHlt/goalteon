package beans

import (
	"fmt"
	"reflect"
)

// SlbNewAcclCfgClusterTable The table for configuring fastview cluster policy.
// Note:This mib is not supported for VX instance of Virtualization.
type SlbNewAcclCfgClusterTable struct {
	//
	SlbNewAcclCfgClusterVirtId string
	// The layer 4 virtual port number for the service.
	SlbNewAcclCfgClusterServiceVport int32
	Params                           *SlbNewAcclCfgClusterTableParams
}

func NewSlbNewAcclCfgClusterTableList() *SlbNewAcclCfgClusterTable {
	return &SlbNewAcclCfgClusterTable{}
}

func NewSlbNewAcclCfgClusterTable(
	slbNewAcclCfgClusterVirtId string,
	slbNewAcclCfgClusterServiceVport int32,
	params *SlbNewAcclCfgClusterTableParams,
) *SlbNewAcclCfgClusterTable {
	return &SlbNewAcclCfgClusterTable{
		SlbNewAcclCfgClusterVirtId:       slbNewAcclCfgClusterVirtId,
		SlbNewAcclCfgClusterServiceVport: slbNewAcclCfgClusterServiceVport,
		Params:                           params,
	}
}

func (c *SlbNewAcclCfgClusterTable) Name() string {
	return "SlbNewAcclCfgClusterTable"
}

func (c *SlbNewAcclCfgClusterTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbNewAcclCfgClusterTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbNewAcclCfgClusterTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbNewAcclCfgClusterVirtId).IsZero() &&
		reflect.ValueOf(c.SlbNewAcclCfgClusterServiceVport).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbNewAcclCfgClusterVirtId) + "/" + fmt.Sprint(c.SlbNewAcclCfgClusterServiceVport)
}

type SlbNewAcclCfgClusterTableAdminStatus int32

const (
	SlbNewAcclCfgClusterTableAdminStatus_Enabled     SlbNewAcclCfgClusterTableAdminStatus = 1
	SlbNewAcclCfgClusterTableAdminStatus_Disabled    SlbNewAcclCfgClusterTableAdminStatus = 2
	SlbNewAcclCfgClusterTableAdminStatus_Unsupported SlbNewAcclCfgClusterTableAdminStatus = 2147483647
)

type SlbNewAcclCfgClusterTableDelete int32

const (
	SlbNewAcclCfgClusterTableDelete_Other       SlbNewAcclCfgClusterTableDelete = 1
	SlbNewAcclCfgClusterTableDelete_Delete      SlbNewAcclCfgClusterTableDelete = 2
	SlbNewAcclCfgClusterTableDelete_Unsupported SlbNewAcclCfgClusterTableDelete = 2147483647
)

type SlbNewAcclCfgClusterTableParams struct {
	//
	VirtId string `json:"VirtId,omitempty"`
	// The layer 4 virtual port number for the service.
	ServiceVport int32 `json:"ServiceVport,omitempty"`
	// Enable or disable fastview cluster.
	AdminStatus SlbNewAcclCfgClusterTableAdminStatus `json:"AdminStatus,omitempty"`
	// IP address of the peer vip 1.
	PeerVIP1 string `json:"PeerVIP1,omitempty"`
	// IP address of the peer vip 2.
	PeerVIP2 string `json:"PeerVIP2,omitempty"`
	// IP address of the peer vip 3.
	PeerVIP3 string `json:"PeerVIP3,omitempty"`
	// IP address of the peer vip 4.
	PeerVIP4 string `json:"PeerVIP4,omitempty"`
	// IP address of the peer vip 5.
	PeerVIP5 string `json:"PeerVIP5,omitempty"`
	// Delete fastview web application.
	Delete SlbNewAcclCfgClusterTableDelete `json:"Delete,omitempty"`
}

func (p SlbNewAcclCfgClusterTableParams) iMABean() {}
