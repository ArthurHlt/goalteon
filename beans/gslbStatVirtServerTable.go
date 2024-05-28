package beans

import (
	"fmt"
	"reflect"
)

// GslbStatVirtServerTable The Global SLB virtual server statistics table.
// Note:This mib is not supported for VX instance of Virtualization.
type GslbStatVirtServerTable struct {
	// The virtual server number that identifies the virtual server.
	GslbStatVirtServerIdx int32
	// The virtual server service number that identifies the virtual
	// service.
	GslbStatVirtServerServiceIdx int32
	// The real server number that identifies a remote site. A index of
	// 0 indicates a local virtual server number.
	GslbStatVirtServerRserverIdx int32
	Params                       *GslbStatVirtServerTableParams
}

func NewGslbStatVirtServerTableList() *GslbStatVirtServerTable {
	return &GslbStatVirtServerTable{}
}

func NewGslbStatVirtServerTable(
	gslbStatVirtServerIdx int32,
	gslbStatVirtServerServiceIdx int32,
	gslbStatVirtServerRserverIdx int32,
	params *GslbStatVirtServerTableParams,
) *GslbStatVirtServerTable {
	return &GslbStatVirtServerTable{
		GslbStatVirtServerIdx:        gslbStatVirtServerIdx,
		GslbStatVirtServerServiceIdx: gslbStatVirtServerServiceIdx,
		GslbStatVirtServerRserverIdx: gslbStatVirtServerRserverIdx,
		Params:                       params,
	}
}

func (c *GslbStatVirtServerTable) Name() string {
	return "GslbStatVirtServerTable"
}

func (c *GslbStatVirtServerTable) GetParams() BeanType {
	return c.Params
}

func (c *GslbStatVirtServerTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *GslbStatVirtServerTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.GslbStatVirtServerIdx).IsZero() &&
		reflect.ValueOf(c.GslbStatVirtServerServiceIdx).IsZero() &&
		reflect.ValueOf(c.GslbStatVirtServerRserverIdx).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.GslbStatVirtServerIdx) + "/" + fmt.Sprint(c.GslbStatVirtServerServiceIdx) + "/" + fmt.Sprint(c.GslbStatVirtServerRserverIdx)
}

type GslbStatVirtServerTableParams struct {
	// The virtual server number that identifies the virtual server.
	Idx int32 `json:"Idx,omitempty"`
	// The virtual server service number that identifies the virtual
	// service.
	ServiceIdx int32 `json:"ServiceIdx,omitempty"`
	// The real server number that identifies a remote site. A index of
	// 0 indicates a local virtual server number.
	RserverIdx int32 `json:"RserverIdx,omitempty"`
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
}

func (p GslbStatVirtServerTableParams) iMABean() {}
