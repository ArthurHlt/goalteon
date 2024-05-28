package beans

import (
	"fmt"
	"reflect"
)

// GslbStatEnhVirtServerTable The Global SLB virtual server statistics table.
// Note:This mib is not supported for VX instance of Virtualization.
type GslbStatEnhVirtServerTable struct {
	// The virtual server number that identifies the virtual server.
	GslbStatEnhVirtServerIdx string
	// The virtual server service number that identifies the virtual
	// service.
	GslbStatEnhVirtServerServiceIdx int32
	// The real server number that identifies a remote site. A index of
	// 0 indicates a local virtual server number.
	GslbStatEnhVirtServerRserverIdx string
	Params                          *GslbStatEnhVirtServerTableParams
}

func NewGslbStatEnhVirtServerTableList() *GslbStatEnhVirtServerTable {
	return &GslbStatEnhVirtServerTable{}
}

func NewGslbStatEnhVirtServerTable(
	gslbStatEnhVirtServerIdx string,
	gslbStatEnhVirtServerServiceIdx int32,
	gslbStatEnhVirtServerRserverIdx string,
	params *GslbStatEnhVirtServerTableParams,
) *GslbStatEnhVirtServerTable {
	return &GslbStatEnhVirtServerTable{
		GslbStatEnhVirtServerIdx:        gslbStatEnhVirtServerIdx,
		GslbStatEnhVirtServerServiceIdx: gslbStatEnhVirtServerServiceIdx,
		GslbStatEnhVirtServerRserverIdx: gslbStatEnhVirtServerRserverIdx,
		Params:                          params,
	}
}

func (c *GslbStatEnhVirtServerTable) Name() string {
	return "GslbStatEnhVirtServerTable"
}

func (c *GslbStatEnhVirtServerTable) GetParams() BeanType {
	return c.Params
}

func (c *GslbStatEnhVirtServerTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *GslbStatEnhVirtServerTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.GslbStatEnhVirtServerIdx).IsZero() &&
		reflect.ValueOf(c.GslbStatEnhVirtServerServiceIdx).IsZero() &&
		reflect.ValueOf(c.GslbStatEnhVirtServerRserverIdx).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.GslbStatEnhVirtServerIdx) + "/" + fmt.Sprint(c.GslbStatEnhVirtServerServiceIdx) + "/" + fmt.Sprint(c.GslbStatEnhVirtServerRserverIdx)
}

type GslbStatEnhVirtServerTableParams struct {
	// The virtual server number that identifies the virtual server.
	Idx string `json:"Idx,omitempty"`
	// The virtual server service number that identifies the virtual
	// service.
	ServiceIdx int32 `json:"ServiceIdx,omitempty"`
	// The real server number that identifies a remote site. A index of
	// 0 indicates a local virtual server number.
	RserverIdx string `json:"RserverIdx,omitempty"`
	// The virtual server service port.
	VirtPort int32 `json:"VirtPort,omitempty"`
	// The IP address of the virtual server or remote real server.
	IpAddress string `json:"IpAddress,omitempty"`
	// The average time (current weight) that each service takes to respond
	// to information exchanges with its peers The time is specified in
	// ticks of 65 milliseconds.
	ResponseTime int32 `json:"ResponseTime,omitempty"`
	// The current number of sessions available for serving client requests.
	// This number will change as client traffic loads change, or as real
	// servers under the virtual server or remote sites go in or out of
	// service.
	MinSessAvail uint32 `json:"MinSessAvail,omitempty"`
	// The domain name of the virtual server and remote real server.
	Dname string `json:"Dname,omitempty"`
	// The remote site of the remote real server. For virtual server, the remote site is 0.
	RemSite int32 `json:"RemSite,omitempty"`
	// The total number of DNS directs sent to the virtual server or remote real server.
	DnsDirect uint32 `json:"DnsDirect,omitempty"`
	// The IPv6 address of the virtual server or remote real server.
	Ipv6Address string `json:"Ipv6Address,omitempty"`
	// The number of HTTP redirections by the remote real server.
	HttpRedirs uint32 `json:"HttpRedirs,omitempty"`
}

func (p GslbStatEnhVirtServerTableParams) iMABean() {}
