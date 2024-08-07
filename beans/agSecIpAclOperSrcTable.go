package beans

import (
	"fmt"
	"reflect"
)

// AgSecIpAclOperSrcTable The table of IP ACL Oper Config Source Ip.
// Note:This mib is not supported for VX instance of virtualization.
type AgSecIpAclOperSrcTable struct {
	// The IP ACL Oper Config Source Group index.
	// Note:This mib is not supported for VX instance of virtualization.
	AgSecIpAclOperListSrcIndx int32
	Params                    *AgSecIpAclOperSrcTableParams
}

func NewAgSecIpAclOperSrcTableList() *AgSecIpAclOperSrcTable {
	return &AgSecIpAclOperSrcTable{}
}

func NewAgSecIpAclOperSrcTable(
	agSecIpAclOperListSrcIndx int32,
	params *AgSecIpAclOperSrcTableParams,
) *AgSecIpAclOperSrcTable {
	return &AgSecIpAclOperSrcTable{
		AgSecIpAclOperListSrcIndx: agSecIpAclOperListSrcIndx,
		Params:                    params,
	}
}

func (c *AgSecIpAclOperSrcTable) Name() string {
	return "AgSecIpAclOperSrcTable"
}

func (c *AgSecIpAclOperSrcTable) GetParams() BeanType {
	return c.Params
}

func (c *AgSecIpAclOperSrcTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *AgSecIpAclOperSrcTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.AgSecIpAclOperListSrcIndx).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.AgSecIpAclOperListSrcIndx)
}

type AgSecIpAclOperSrcTableParams struct {
	// The IP ACL Oper Config Source Group index.
	// Note:This mib is not supported for VX instance of virtualization.
	ListSrcIndx int32 `json:"ListSrcIndx,omitempty"`
	// Oper IP ACL Config Source IP Address.
	// Note:This mib is not supported for VX instance of virtualization.
	ListSrcAddr string `json:"ListSrcAddr,omitempty"`
	// Oper IP ACL Config Source NetMask.
	// Note:This mib is not supported for VX instance of virtualization.
	ListSrcMask string `json:"ListSrcMask,omitempty"`
	// Config Source Ip Timeout. Note:This mib is not supported for VX instance of virtualization.
	ListSrcTimeOut int32 `json:"ListSrcTimeOut,omitempty"`
}

func (p AgSecIpAclOperSrcTableParams) iMABean() {}
