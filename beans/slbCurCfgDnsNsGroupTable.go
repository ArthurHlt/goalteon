package beans

import (
	"fmt"
	"reflect"
)

// SlbCurCfgDnsNsGroupTable The table of DNS NS groups.
// Note:This mib is not supported on VX instance of Virtualization.
type SlbCurCfgDnsNsGroupTable struct {
	// The DNS NS group id as index.
	SlbCurCfgDnsNsGroupId string
	Params                *SlbCurCfgDnsNsGroupTableParams
}

func NewSlbCurCfgDnsNsGroupTableList() *SlbCurCfgDnsNsGroupTable {
	return &SlbCurCfgDnsNsGroupTable{}
}

func NewSlbCurCfgDnsNsGroupTable(
	slbCurCfgDnsNsGroupId string,
	params *SlbCurCfgDnsNsGroupTableParams,
) *SlbCurCfgDnsNsGroupTable {
	return &SlbCurCfgDnsNsGroupTable{
		SlbCurCfgDnsNsGroupId: slbCurCfgDnsNsGroupId,
		Params:                params,
	}
}

func (c *SlbCurCfgDnsNsGroupTable) Name() string {
	return "SlbCurCfgDnsNsGroupTable"
}

func (c *SlbCurCfgDnsNsGroupTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbCurCfgDnsNsGroupTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbCurCfgDnsNsGroupTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbCurCfgDnsNsGroupId).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbCurCfgDnsNsGroupId)
}

type SlbCurCfgDnsNsGroupTableParams struct {
	// The DNS NS group id as index.
	Id string `json:"Id,omitempty"`
	// The DNS NS group TTL value.
	Ttl uint64 `json:"Ttl,omitempty"`
}

func (p SlbCurCfgDnsNsGroupTableParams) iMABean() {}
