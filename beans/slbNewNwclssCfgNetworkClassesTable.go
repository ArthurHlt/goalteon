package beans

import (
	"fmt"
	"reflect"
)

// SlbNewNwclssCfgNetworkClassesTable New network classes table.
// Note:This mib is not supported for VX instance of Virtualization.
type SlbNewNwclssCfgNetworkClassesTable struct {
	// Network class id.
	SlbNewNwclssCfgNetworkClassesId string
	Params                          *SlbNewNwclssCfgNetworkClassesTableParams
}

func NewSlbNewNwclssCfgNetworkClassesTableList() *SlbNewNwclssCfgNetworkClassesTable {
	return &SlbNewNwclssCfgNetworkClassesTable{}
}

func NewSlbNewNwclssCfgNetworkClassesTable(
	slbNewNwclssCfgNetworkClassesId string,
	params *SlbNewNwclssCfgNetworkClassesTableParams,
) *SlbNewNwclssCfgNetworkClassesTable {
	return &SlbNewNwclssCfgNetworkClassesTable{
		SlbNewNwclssCfgNetworkClassesId: slbNewNwclssCfgNetworkClassesId,
		Params:                          params,
	}
}

func (c *SlbNewNwclssCfgNetworkClassesTable) Name() string {
	return "SlbNewNwclssCfgNetworkClassesTable"
}

func (c *SlbNewNwclssCfgNetworkClassesTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbNewNwclssCfgNetworkClassesTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbNewNwclssCfgNetworkClassesTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbNewNwclssCfgNetworkClassesId).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbNewNwclssCfgNetworkClassesId)
}

type SlbNewNwclssCfgNetworkClassesTableIpVer int32

const (
	SlbNewNwclssCfgNetworkClassesTableIpVer_Ipv4        SlbNewNwclssCfgNetworkClassesTableIpVer = 1
	SlbNewNwclssCfgNetworkClassesTableIpVer_Ipv6        SlbNewNwclssCfgNetworkClassesTableIpVer = 2
	SlbNewNwclssCfgNetworkClassesTableIpVer_Unsupported SlbNewNwclssCfgNetworkClassesTableIpVer = 2147483647
)

type SlbNewNwclssCfgNetworkClassesTableDel int32

const (
	SlbNewNwclssCfgNetworkClassesTableDel_Other       SlbNewNwclssCfgNetworkClassesTableDel = 1
	SlbNewNwclssCfgNetworkClassesTableDel_Delete      SlbNewNwclssCfgNetworkClassesTableDel = 2
	SlbNewNwclssCfgNetworkClassesTableDel_Unsupported SlbNewNwclssCfgNetworkClassesTableDel = 2147483647
)

type SlbNewNwclssCfgNetworkClassesTableType int32

const (
	SlbNewNwclssCfgNetworkClassesTableType_Address     SlbNewNwclssCfgNetworkClassesTableType = 1
	SlbNewNwclssCfgNetworkClassesTableType_Region      SlbNewNwclssCfgNetworkClassesTableType = 2
	SlbNewNwclssCfgNetworkClassesTableType_Unsupported SlbNewNwclssCfgNetworkClassesTableType = 2147483647
)

type SlbNewNwclssCfgNetworkClassesTableParams struct {
	// Network class id.
	Id string `json:"Id,omitempty"`
	// Network class name.
	Name string `json:"Name,omitempty"`
	// The type of IP address.
	IpVer SlbNewNwclssCfgNetworkClassesTableIpVer `json:"IpVer,omitempty"`
	// By setting the value to delete(2), the entire row is deleted.
	Del SlbNewNwclssCfgNetworkClassesTableDel `json:"Del,omitempty"`
	// Duplicating an entire Network class by defining the destination Network class id.
	Copy DisplayString `json:"Copy,omitempty"`
	// The network class type should be either 'address' or 'region' to be added to the network.
	Type SlbNewNwclssCfgNetworkClassesTableType `json:"Type,omitempty"`
}

func (p SlbNewNwclssCfgNetworkClassesTableParams) iMABean() {}
