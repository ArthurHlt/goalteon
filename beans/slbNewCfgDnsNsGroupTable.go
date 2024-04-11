package beans

import (
	"fmt"
	"reflect"
)

// SlbNewCfgDnsNsGroupTable The table of DNS NS groups.
// Note:This mib is not supported on VX instance of Virtualization.
type SlbNewCfgDnsNsGroupTable struct {
	// The DNS NS group id as index.
	SlbNewCfgDnsNsGroupId string
	Params                *SlbNewCfgDnsNsGroupTableParams
}

func NewSlbNewCfgDnsNsGroupTableList() *SlbNewCfgDnsNsGroupTable {
	return &SlbNewCfgDnsNsGroupTable{}
}

func NewSlbNewCfgDnsNsGroupTable(
	slbNewCfgDnsNsGroupId string,
	params *SlbNewCfgDnsNsGroupTableParams,
) *SlbNewCfgDnsNsGroupTable {
	return &SlbNewCfgDnsNsGroupTable{
		SlbNewCfgDnsNsGroupId: slbNewCfgDnsNsGroupId,
		Params:                params,
	}
}

func (c *SlbNewCfgDnsNsGroupTable) Name() string {
	return "SlbNewCfgDnsNsGroupTable"
}

func (c *SlbNewCfgDnsNsGroupTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbNewCfgDnsNsGroupTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbNewCfgDnsNsGroupTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbNewCfgDnsNsGroupId).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbNewCfgDnsNsGroupId)
}

type SlbNewCfgDnsNsGroupTableDelete int32

const (
	SlbNewCfgDnsNsGroupTableDelete_Other  SlbNewCfgDnsNsGroupTableDelete = 1
	SlbNewCfgDnsNsGroupTableDelete_Delete SlbNewCfgDnsNsGroupTableDelete = 2
)

type SlbNewCfgDnsNsGroupTableParams struct {
	// The DNS NS group id as index.
	Id string `json:"Id,omitempty"`
	// The DNS NS group TTL value.
	Ttl uint64 `json:"Ttl,omitempty"`
	// When set to the value of 2(delete), the entire DNS NS group is deleted.
	// When read, other(1) is returned. Setting the value to anything
	// other than 2(delete) has no effect on the state of the row.
	Delete SlbNewCfgDnsNsGroupTableDelete `json:"Delete,omitempty"`
}

func (p SlbNewCfgDnsNsGroupTableParams) iMABean() {}
