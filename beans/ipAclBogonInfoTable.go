package beans

import (
	"fmt"
	"reflect"
)

// IpAclBogonInfoTable The table of bogons.
// Note: This MIB table is not supported for VX instance of Virtualization.
type IpAclBogonInfoTable struct {
	// The table index.
	IpAclBogonInfoIndex int32
	Params              *IpAclBogonInfoTableParams
}

func NewIpAclBogonInfoTableList() *IpAclBogonInfoTable {
	return &IpAclBogonInfoTable{}
}

func NewIpAclBogonInfoTable(
	ipAclBogonInfoIndex int32,
	params *IpAclBogonInfoTableParams,
) *IpAclBogonInfoTable {
	return &IpAclBogonInfoTable{
		IpAclBogonInfoIndex: ipAclBogonInfoIndex,
		Params:              params,
	}
}

func (c *IpAclBogonInfoTable) Name() string {
	return "IpAclBogonInfoTable"
}

func (c *IpAclBogonInfoTable) GetParams() BeanType {
	return c.Params
}

func (c *IpAclBogonInfoTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *IpAclBogonInfoTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.IpAclBogonInfoIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.IpAclBogonInfoIndex)
}

type IpAclBogonInfoTableParams struct {
	// The table index.
	Index int32 `json:"Index,omitempty"`
	// Bogon IP Address.
	Ip string `json:"Ip,omitempty"`
	// Mask for the bogon IP.
	Mask string `json:"Mask,omitempty"`
}

func (p IpAclBogonInfoTableParams) iMABean() {}
