package beans

import (
	"fmt"
	"reflect"
)

// SlbCurAcclCfgClusterTable The table for configuring fastview non affinity policy.
// Note:This mib is not supported for VX instance of Virtualization.
type SlbCurAcclCfgClusterTable struct {
	//
	SlbCurAcclCfgClusterVirtId string
	// The layer 4 virtual port number for the service.
	SlbCurAcclCfgClusterServiceVport int32
	Params                           *SlbCurAcclCfgClusterTableParams
}

func NewSlbCurAcclCfgClusterTableList() *SlbCurAcclCfgClusterTable {
	return &SlbCurAcclCfgClusterTable{}
}

func NewSlbCurAcclCfgClusterTable(
	slbCurAcclCfgClusterVirtId string,
	slbCurAcclCfgClusterServiceVport int32,
	params *SlbCurAcclCfgClusterTableParams,
) *SlbCurAcclCfgClusterTable {
	return &SlbCurAcclCfgClusterTable{
		SlbCurAcclCfgClusterVirtId:       slbCurAcclCfgClusterVirtId,
		SlbCurAcclCfgClusterServiceVport: slbCurAcclCfgClusterServiceVport,
		Params:                           params,
	}
}

func (c *SlbCurAcclCfgClusterTable) Name() string {
	return "SlbCurAcclCfgClusterTable"
}

func (c *SlbCurAcclCfgClusterTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbCurAcclCfgClusterTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbCurAcclCfgClusterTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbCurAcclCfgClusterVirtId).IsZero() &&
		reflect.ValueOf(c.SlbCurAcclCfgClusterServiceVport).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbCurAcclCfgClusterVirtId) + "/" + fmt.Sprint(c.SlbCurAcclCfgClusterServiceVport)
}

type SlbCurAcclCfgClusterTableAdminStatus int32

const (
	SlbCurAcclCfgClusterTableAdminStatus_Enabled     SlbCurAcclCfgClusterTableAdminStatus = 1
	SlbCurAcclCfgClusterTableAdminStatus_Disabled    SlbCurAcclCfgClusterTableAdminStatus = 2
	SlbCurAcclCfgClusterTableAdminStatus_Unsupported SlbCurAcclCfgClusterTableAdminStatus = 2147483647
)

type SlbCurAcclCfgClusterTableParams struct {
	//
	VirtId string `json:"VirtId,omitempty"`
	// The layer 4 virtual port number for the service.
	ServiceVport int32 `json:"ServiceVport,omitempty"`
	// Enable or disable fastview cluster.
	AdminStatus SlbCurAcclCfgClusterTableAdminStatus `json:"AdminStatus,omitempty"`
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
}

func (p SlbCurAcclCfgClusterTableParams) iMABean() {}
