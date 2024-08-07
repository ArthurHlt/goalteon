package beans

import (
	"fmt"
	"reflect"
)

// AgSecIpAclOperDstTable The table of Operations Source IP ACL. Note:This mib is not supported for VX instance of virtualization.
type AgSecIpAclOperDstTable struct {
	// The IP ACL Oper Destination IP Group index.
	// Note:This mib is not supported for VX instance of virtualization.
	AgSecIpAclOperDstIndx int32
	Params                *AgSecIpAclOperDstTableParams
}

func NewAgSecIpAclOperDstTableList() *AgSecIpAclOperDstTable {
	return &AgSecIpAclOperDstTable{}
}

func NewAgSecIpAclOperDstTable(
	agSecIpAclOperDstIndx int32,
	params *AgSecIpAclOperDstTableParams,
) *AgSecIpAclOperDstTable {
	return &AgSecIpAclOperDstTable{
		AgSecIpAclOperDstIndx: agSecIpAclOperDstIndx,
		Params:                params,
	}
}

func (c *AgSecIpAclOperDstTable) Name() string {
	return "AgSecIpAclOperDstTable"
}

func (c *AgSecIpAclOperDstTable) GetParams() BeanType {
	return c.Params
}

func (c *AgSecIpAclOperDstTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *AgSecIpAclOperDstTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.AgSecIpAclOperDstIndx).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.AgSecIpAclOperDstIndx)
}

type AgSecIpAclOperDstTableParams struct {
	// The IP ACL Oper Destination IP Group index.
	// Note:This mib is not supported for VX instance of virtualization.
	Indx int32 `json:"Indx,omitempty"`
	// IP ACL Operations Destination IP Address.
	Addr string `json:"Addr,omitempty"`
	// IP ACL Operations Destination NetMask.
	// Note:This mib is not supported for VX instance of virtualization.
	Mask string `json:"Mask,omitempty"`
	// IP ACL Operations Destination Ip Timeout.
	// Note:This mib is not supported for VX instance of virtualization.
	Timeout int32 `json:"Timeout,omitempty"`
}

func (p AgSecIpAclOperDstTableParams) iMABean() {}
