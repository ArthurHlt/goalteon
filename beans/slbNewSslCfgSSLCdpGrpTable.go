package beans

import (
	"fmt"
	"reflect"
)

// SlbNewSslCfgSSLCdpGrpTable The table for configuring Cdp Group.
// Note:This mib is not supported for VX instance of Virtualization.
type SlbNewSslCfgSSLCdpGrpTable struct {
	// The CDP group (key id) as an index.
	SlbNewSslCfgCdpGrpIdIndex string
	Params                    *SlbNewSslCfgSSLCdpGrpTableParams
}

func NewSlbNewSslCfgSSLCdpGrpTableList() *SlbNewSslCfgSSLCdpGrpTable {
	return &SlbNewSslCfgSSLCdpGrpTable{}
}

func NewSlbNewSslCfgSSLCdpGrpTable(
	slbNewSslCfgCdpGrpIdIndex string,
	params *SlbNewSslCfgSSLCdpGrpTableParams,
) *SlbNewSslCfgSSLCdpGrpTable {
	return &SlbNewSslCfgSSLCdpGrpTable{
		SlbNewSslCfgCdpGrpIdIndex: slbNewSslCfgCdpGrpIdIndex,
		Params:                    params,
	}
}

func (c *SlbNewSslCfgSSLCdpGrpTable) Name() string {
	return "SlbNewSslCfgSSLCdpGrpTable"
}

func (c *SlbNewSslCfgSSLCdpGrpTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbNewSslCfgSSLCdpGrpTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbNewSslCfgSSLCdpGrpTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbNewSslCfgCdpGrpIdIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbNewSslCfgCdpGrpIdIndex)
}

type SlbNewSslCfgSSLCdpGrpTableCdpGrpDel int32

const (
	SlbNewSslCfgSSLCdpGrpTableCdpGrpDel_Other       SlbNewSslCfgSSLCdpGrpTableCdpGrpDel = 1
	SlbNewSslCfgSSLCdpGrpTableCdpGrpDel_Delete      SlbNewSslCfgSSLCdpGrpTableCdpGrpDel = 2
	SlbNewSslCfgSSLCdpGrpTableCdpGrpDel_Unsupported SlbNewSslCfgSSLCdpGrpTableCdpGrpDel = 2147483647
)

type SlbNewSslCfgSSLCdpGrpTableParams struct {
	// The CDP group (key id) as an index.
	CdpGrpIdIndex string `json:"CdpGrpIdIndex,omitempty"`
	// Delete CDP Group.
	CdpGrpDel SlbNewSslCfgSSLCdpGrpTableCdpGrpDel `json:"CdpGrpDel,omitempty"`
}

func (p SlbNewSslCfgSSLCdpGrpTableParams) iMABean() {}
