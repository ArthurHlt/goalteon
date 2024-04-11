package beans

import (
	"fmt"
	"reflect"
)

// SlbCurCfgDnsNsRecordTable The table of DNS NS records for each NS group.
// Note:This mib is not supported on VX instance of Virtualization.
type SlbCurCfgDnsNsRecordTable struct {
	// The DNS NS group id as index. Each group can have upto 10 NS records.
	SlbCurCfgDnsNsRecordGroupId string
	// Defines the NS record string in fqdn format.
	SlbCurCfgDnsNsRecordString string
	Params                     *SlbCurCfgDnsNsRecordTableParams
}

func NewSlbCurCfgDnsNsRecordTableList() *SlbCurCfgDnsNsRecordTable {
	return &SlbCurCfgDnsNsRecordTable{}
}

func NewSlbCurCfgDnsNsRecordTable(
	slbCurCfgDnsNsRecordGroupId string,
	slbCurCfgDnsNsRecordString string,
	params *SlbCurCfgDnsNsRecordTableParams,
) *SlbCurCfgDnsNsRecordTable {
	return &SlbCurCfgDnsNsRecordTable{
		SlbCurCfgDnsNsRecordGroupId: slbCurCfgDnsNsRecordGroupId,
		SlbCurCfgDnsNsRecordString:  slbCurCfgDnsNsRecordString,
		Params:                      params,
	}
}

func (c *SlbCurCfgDnsNsRecordTable) Name() string {
	return "SlbCurCfgDnsNsRecordTable"
}

func (c *SlbCurCfgDnsNsRecordTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbCurCfgDnsNsRecordTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbCurCfgDnsNsRecordTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbCurCfgDnsNsRecordGroupId).IsZero() &&
		reflect.ValueOf(c.SlbCurCfgDnsNsRecordString).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbCurCfgDnsNsRecordGroupId) + "/" + fmt.Sprint(c.SlbCurCfgDnsNsRecordString)
}

type SlbCurCfgDnsNsRecordTableParams struct {
	// The DNS NS group id as index. Each group can have upto 10 NS records.
	GroupId string `json:"GroupId,omitempty"`
	// Defines the NS record string in fqdn format.
	StringVal string `json:"StringVal,omitempty"`
	// Defines the IPv4 glue address for this NS record entry. This is optional.
	// Returns 0 when not set.
	Ipv4Addr string `json:"Ipv4Addr,omitempty"`
	// Defines the IPv6 glue address for this NS record entry. This is optional.
	// Returns 0 when not set.
	Ipv6Addr string `json:"Ipv6Addr,omitempty"`
}

func (p SlbCurCfgDnsNsRecordTableParams) iMABean() {}
