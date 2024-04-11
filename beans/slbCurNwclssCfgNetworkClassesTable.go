package beans

import (
	"fmt"
	"reflect"
)

// SlbCurNwclssCfgNetworkClassesTable Current network classes table.
// Note:This mib is not supported for VX instance of Virtualization.
type SlbCurNwclssCfgNetworkClassesTable struct {
	// Network class id.
	SlbCurNwclssCfgNetworkClassesId string
	Params                          *SlbCurNwclssCfgNetworkClassesTableParams
}

func NewSlbCurNwclssCfgNetworkClassesTableList() *SlbCurNwclssCfgNetworkClassesTable {
	return &SlbCurNwclssCfgNetworkClassesTable{}
}

func NewSlbCurNwclssCfgNetworkClassesTable(
	slbCurNwclssCfgNetworkClassesId string,
	params *SlbCurNwclssCfgNetworkClassesTableParams,
) *SlbCurNwclssCfgNetworkClassesTable {
	return &SlbCurNwclssCfgNetworkClassesTable{
		SlbCurNwclssCfgNetworkClassesId: slbCurNwclssCfgNetworkClassesId,
		Params:                          params,
	}
}

func (c *SlbCurNwclssCfgNetworkClassesTable) Name() string {
	return "SlbCurNwclssCfgNetworkClassesTable"
}

func (c *SlbCurNwclssCfgNetworkClassesTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbCurNwclssCfgNetworkClassesTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbCurNwclssCfgNetworkClassesTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbCurNwclssCfgNetworkClassesId).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbCurNwclssCfgNetworkClassesId)
}

type SlbCurNwclssCfgNetworkClassesTableIpVer int32

const (
	SlbCurNwclssCfgNetworkClassesTableIpVer_Ipv4        SlbCurNwclssCfgNetworkClassesTableIpVer = 1
	SlbCurNwclssCfgNetworkClassesTableIpVer_Ipv6        SlbCurNwclssCfgNetworkClassesTableIpVer = 2
	SlbCurNwclssCfgNetworkClassesTableIpVer_Unsupported SlbCurNwclssCfgNetworkClassesTableIpVer = 2147483647
)

type SlbCurNwclssCfgNetworkClassesTableType int32

const (
	SlbCurNwclssCfgNetworkClassesTableType_Address     SlbCurNwclssCfgNetworkClassesTableType = 1
	SlbCurNwclssCfgNetworkClassesTableType_Region      SlbCurNwclssCfgNetworkClassesTableType = 2
	SlbCurNwclssCfgNetworkClassesTableType_Unsupported SlbCurNwclssCfgNetworkClassesTableType = 2147483647
)

type SlbCurNwclssCfgNetworkClassesTableParams struct {
	// Network class id.
	Id string `json:"Id,omitempty"`
	// Network class name.
	Name string `json:"Name,omitempty"`
	// The type of IP address.
	IpVer SlbCurNwclssCfgNetworkClassesTableIpVer `json:"IpVer,omitempty"`
	// The network class type should be either 'address' or 'region'.
	Type SlbCurNwclssCfgNetworkClassesTableType `json:"Type,omitempty"`
}

func (p SlbCurNwclssCfgNetworkClassesTableParams) iMABean() {}
