package beans

import (
	"fmt"
	"reflect"
)

// SlbNewCfgDnsNsRecordTable The table of DNS NS records for each NS group.
// Note:This mib is not supported on VX instance of Virtualization.
type SlbNewCfgDnsNsRecordTable struct {
	// The DNS NS group id as index. Each group can have upto 10 records.
	SlbNewCfgDnsNsRecordGroupId string
	// Defines the NS record string in fqdn format.
	SlbNewCfgDnsNsRecordString string
	Params                     *SlbNewCfgDnsNsRecordTableParams
}

func NewSlbNewCfgDnsNsRecordTableList() *SlbNewCfgDnsNsRecordTable {
	return &SlbNewCfgDnsNsRecordTable{}
}

func NewSlbNewCfgDnsNsRecordTable(
	slbNewCfgDnsNsRecordGroupId string,
	slbNewCfgDnsNsRecordString string,
	params *SlbNewCfgDnsNsRecordTableParams,
) *SlbNewCfgDnsNsRecordTable {
	return &SlbNewCfgDnsNsRecordTable{
		SlbNewCfgDnsNsRecordGroupId: slbNewCfgDnsNsRecordGroupId,
		SlbNewCfgDnsNsRecordString:  slbNewCfgDnsNsRecordString,
		Params:                      params,
	}
}

func (c *SlbNewCfgDnsNsRecordTable) Name() string {
	return "SlbNewCfgDnsNsRecordTable"
}

func (c *SlbNewCfgDnsNsRecordTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbNewCfgDnsNsRecordTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbNewCfgDnsNsRecordTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbNewCfgDnsNsRecordGroupId).IsZero() &&
		reflect.ValueOf(c.SlbNewCfgDnsNsRecordString).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbNewCfgDnsNsRecordGroupId) + "/" + fmt.Sprint(c.SlbNewCfgDnsNsRecordString)
}

type SlbNewCfgDnsNsRecordTableDelete int32

const (
	SlbNewCfgDnsNsRecordTableDelete_Other  SlbNewCfgDnsNsRecordTableDelete = 1
	SlbNewCfgDnsNsRecordTableDelete_Delete SlbNewCfgDnsNsRecordTableDelete = 2
)

type SlbNewCfgDnsNsRecordTableParams struct {
	// The DNS NS group id as index. Each group can have upto 10 records.
	GroupId string `json:"GroupId,omitempty"`
	// Defines the NS record string in fqdn format.
	StringVal string `json:"StringVal,omitempty"`
	// Defines the IPv4 glue address for this NS record entry. This is optional.
	// Returns 0 when not set.
	Ipv4Addr string `json:"Ipv4Addr,omitempty"`
	// Defines the IPv6 glue address for this NS record entry. This is optional.
	// Returns 0 when not set.
	Ipv6Addr string `json:"Ipv6Addr,omitempty"`
	// When set to the value of 2(delete), the entire NS record is deleted.
	// When read, other(1) is returned. Setting the value to anything
	// other than 2(delete) has no effect on the state of the row.
	Delete SlbNewCfgDnsNsRecordTableDelete `json:"Delete,omitempty"`
}

func (p SlbNewCfgDnsNsRecordTableParams) iMABean() {}
