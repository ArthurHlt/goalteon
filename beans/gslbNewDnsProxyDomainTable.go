package beans

import (
	"fmt"
	"reflect"
)

// GslbNewDnsProxyDomainTable The table of domain entries in new config.
// Note:This mib is not supported for VX instance of Virtualization.
type GslbNewDnsProxyDomainTable struct {
	// DNS Proxy table index in the new configuration block.
	GslbNewDnsProxyDomainId int32
	Params                  *GslbNewDnsProxyDomainTableParams
}

func NewGslbNewDnsProxyDomainTableList() *GslbNewDnsProxyDomainTable {
	return &GslbNewDnsProxyDomainTable{}
}

func NewGslbNewDnsProxyDomainTable(
	gslbNewDnsProxyDomainId int32,
	params *GslbNewDnsProxyDomainTableParams,
) *GslbNewDnsProxyDomainTable {
	return &GslbNewDnsProxyDomainTable{
		GslbNewDnsProxyDomainId: gslbNewDnsProxyDomainId,
		Params:                  params,
	}
}

func (c *GslbNewDnsProxyDomainTable) Name() string {
	return "GslbNewDnsProxyDomainTable"
}

func (c *GslbNewDnsProxyDomainTable) GetParams() BeanType {
	return c.Params
}

func (c *GslbNewDnsProxyDomainTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *GslbNewDnsProxyDomainTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.GslbNewDnsProxyDomainId).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.GslbNewDnsProxyDomainId)
}

type GslbNewDnsProxyDomainTableDelete int32

const (
	GslbNewDnsProxyDomainTableDelete_Other       GslbNewDnsProxyDomainTableDelete = 1
	GslbNewDnsProxyDomainTableDelete_Delete      GslbNewDnsProxyDomainTableDelete = 2
	GslbNewDnsProxyDomainTableDelete_Unsupported GslbNewDnsProxyDomainTableDelete = 2147483647
)

type GslbNewDnsProxyDomainTableState int32

const (
	GslbNewDnsProxyDomainTableState_Enabled     GslbNewDnsProxyDomainTableState = 1
	GslbNewDnsProxyDomainTableState_Disabled    GslbNewDnsProxyDomainTableState = 2
	GslbNewDnsProxyDomainTableState_Unsupported GslbNewDnsProxyDomainTableState = 2147483647
)

type GslbNewDnsProxyDomainTableParams struct {
	// DNS Proxy table index in the new configuration block.
	Id int32 `json:"Id,omitempty"`
	// DNS Proxy Domain name.
	Name string `json:"Name,omitempty"`
	// DNS Proxy Domain group.
	Group string `json:"Group,omitempty"`
	// By setting the value to delete(2), the domain entry is deleted.
	Delete GslbNewDnsProxyDomainTableDelete `json:"Delete,omitempty"`
	// Enable/Disable the proxy entry.
	State GslbNewDnsProxyDomainTableState `json:"State,omitempty"`
}

func (p GslbNewDnsProxyDomainTableParams) iMABean() {}
