package beans

import (
	"fmt"
	"reflect"
)

// GslbCurDnsProxyDomainTable The table of domain entries in current config.
// Note:This mib is not supported for VX instance of Virtualization.
type GslbCurDnsProxyDomainTable struct {
	// DNS Proxy table index in the current configuration block.
	GslbCurDnsProxyDomainId int32
	Params                  *GslbCurDnsProxyDomainTableParams
}

func NewGslbCurDnsProxyDomainTableList() *GslbCurDnsProxyDomainTable {
	return &GslbCurDnsProxyDomainTable{}
}

func NewGslbCurDnsProxyDomainTable(
	gslbCurDnsProxyDomainId int32,
	params *GslbCurDnsProxyDomainTableParams,
) *GslbCurDnsProxyDomainTable {
	return &GslbCurDnsProxyDomainTable{
		GslbCurDnsProxyDomainId: gslbCurDnsProxyDomainId,
		Params:                  params,
	}
}

func (c *GslbCurDnsProxyDomainTable) Name() string {
	return "GslbCurDnsProxyDomainTable"
}

func (c *GslbCurDnsProxyDomainTable) GetParams() BeanType {
	return c.Params
}

func (c *GslbCurDnsProxyDomainTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *GslbCurDnsProxyDomainTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.GslbCurDnsProxyDomainId).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.GslbCurDnsProxyDomainId)
}

type GslbCurDnsProxyDomainTableState int32

const (
	GslbCurDnsProxyDomainTableState_Enabled     GslbCurDnsProxyDomainTableState = 1
	GslbCurDnsProxyDomainTableState_Disabled    GslbCurDnsProxyDomainTableState = 2
	GslbCurDnsProxyDomainTableState_Unsupported GslbCurDnsProxyDomainTableState = 2147483647
)

type GslbCurDnsProxyDomainTableParams struct {
	// DNS Proxy table index in the current configuration block.
	Id int32 `json:"Id,omitempty"`
	// DNS Proxy Domain name.
	Name string `json:"Name,omitempty"`
	// DNS Proxy Domain group.
	Group string `json:"Group,omitempty"`
	// Enable/Disable the proxy entry.
	State GslbCurDnsProxyDomainTableState `json:"State,omitempty"`
}

func (p GslbCurDnsProxyDomainTableParams) iMABean() {}
