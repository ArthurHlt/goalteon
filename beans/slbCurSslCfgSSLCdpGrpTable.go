package beans

import (
	"fmt"
	"reflect"
)

// SlbCurSslCfgSSLCdpGrpTable The table for configuring CDP group.
// Note:This mib is not supported for VX instance of Virtualization.
type SlbCurSslCfgSSLCdpGrpTable struct {
	// The CDP group (key id) as an index.
	SlbCurSslCfgSSLCdpGroupIdIndex string
	Params                         *SlbCurSslCfgSSLCdpGrpTableParams
}

func NewSlbCurSslCfgSSLCdpGrpTableList() *SlbCurSslCfgSSLCdpGrpTable {
	return &SlbCurSslCfgSSLCdpGrpTable{}
}

func NewSlbCurSslCfgSSLCdpGrpTable(
	slbCurSslCfgSSLCdpGroupIdIndex string,
	params *SlbCurSslCfgSSLCdpGrpTableParams,
) *SlbCurSslCfgSSLCdpGrpTable {
	return &SlbCurSslCfgSSLCdpGrpTable{
		SlbCurSslCfgSSLCdpGroupIdIndex: slbCurSslCfgSSLCdpGroupIdIndex,
		Params:                         params,
	}
}

func (c *SlbCurSslCfgSSLCdpGrpTable) Name() string {
	return "SlbCurSslCfgSSLCdpGrpTable"
}

func (c *SlbCurSslCfgSSLCdpGrpTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbCurSslCfgSSLCdpGrpTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbCurSslCfgSSLCdpGrpTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbCurSslCfgSSLCdpGroupIdIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbCurSslCfgSSLCdpGroupIdIndex)
}

type SlbCurSslCfgSSLCdpGrpTableParams struct {
	// The CDP group (key id) as an index.
	GroupIdIndex string `json:"GroupIdIndex,omitempty"`
}

func (p SlbCurSslCfgSSLCdpGrpTableParams) iMABean() {}
