package beans

import (
	"fmt"
	"reflect"
)

// GslbStatDnsSecTable The Global SLB DNS statistics table.
// Note:This mib is not supported for VX instance of Virtualization.
type GslbStatDnsSecTable struct {
	// The DNS ID that identifies the DNS.
	GslbStatDnsIdx string
	Params         *GslbStatDnsSecTableParams
}

func NewGslbStatDnsSecTableList() *GslbStatDnsSecTable {
	return &GslbStatDnsSecTable{}
}

func NewGslbStatDnsSecTable(
	gslbStatDnsIdx string,
	params *GslbStatDnsSecTableParams,
) *GslbStatDnsSecTable {
	return &GslbStatDnsSecTable{
		GslbStatDnsIdx: gslbStatDnsIdx,
		Params:         params,
	}
}

func (c *GslbStatDnsSecTable) Name() string {
	return "GslbStatDnsSecTable"
}

func (c *GslbStatDnsSecTable) GetParams() BeanType {
	return c.Params
}

func (c *GslbStatDnsSecTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *GslbStatDnsSecTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.GslbStatDnsIdx).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.GslbStatDnsIdx)
}

type GslbStatDnsSecTableParams struct {
	// The DNS ID that identifies the DNS.
	Idx string `json:"Idx,omitempty"`
	// The number of total DNS requests since boot time.
	TotalRequest uint32 `json:"TotalRequest,omitempty"`
	// The number of total DNSSEC requests since boot time.
	DnssecTotalRequest uint32 `json:"DnssecTotalRequest,omitempty"`
	// The number of total DNSSEC out of Total DNS in percent.
	DnssecRequestPercent uint32 `json:"DnssecRequestPercent,omitempty"`
	// The number of DNS request per seconds (CPS).
	RequestPerSec uint32 `json:"RequestPerSec,omitempty"`
	// The number of total DNSSEC request per seconds (CPS).
	DnssecRequestPerSec uint32 `json:"DnssecRequestPerSec,omitempty"`
	// The number of total TCP DNS requests.
	TcpRequest uint32 `json:"TcpRequest,omitempty"`
	// The number of total UDP DNS requests.
	UdpRequest uint32 `json:"UdpRequest,omitempty"`
	// The number of invalid DNS requests.
	InvalidRequest uint32 `json:"InvalidRequest,omitempty"`
	// The number of total NSEC records answer since boot time.
	NsecRecordAns uint32 `json:"NsecRecordAns,omitempty"`
}

func (p GslbStatDnsSecTableParams) iMABean() {}
