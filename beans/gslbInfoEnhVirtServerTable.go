package beans

import (
	"fmt"
	"reflect"
)

// GslbInfoEnhVirtServerTable The Global SLB virtual server information table.
// Note:This mib is not supported for VX instance of Virtualization.
type GslbInfoEnhVirtServerTable struct {
	// The virtual server number that identifies the virtual server.
	GslbInfoEnhVirtServerIdx string
	// The virtual server service number that identifies the virtual service.
	GslbInfoEnhVirtServerServiceIdx int32
	// The real server number that identifies a remote real server.
	// An index of 0 indicates virtual server.
	GslbInfoEnhVirtServerRserverIdx string
	Params                          *GslbInfoEnhVirtServerTableParams
}

func NewGslbInfoEnhVirtServerTableList() *GslbInfoEnhVirtServerTable {
	return &GslbInfoEnhVirtServerTable{}
}

func NewGslbInfoEnhVirtServerTable(
	gslbInfoEnhVirtServerIdx string,
	gslbInfoEnhVirtServerServiceIdx int32,
	gslbInfoEnhVirtServerRserverIdx string,
	params *GslbInfoEnhVirtServerTableParams,
) *GslbInfoEnhVirtServerTable {
	return &GslbInfoEnhVirtServerTable{
		GslbInfoEnhVirtServerIdx:        gslbInfoEnhVirtServerIdx,
		GslbInfoEnhVirtServerServiceIdx: gslbInfoEnhVirtServerServiceIdx,
		GslbInfoEnhVirtServerRserverIdx: gslbInfoEnhVirtServerRserverIdx,
		Params:                          params,
	}
}

func (c *GslbInfoEnhVirtServerTable) Name() string {
	return "GslbInfoEnhVirtServerTable"
}

func (c *GslbInfoEnhVirtServerTable) GetParams() BeanType {
	return c.Params
}

func (c *GslbInfoEnhVirtServerTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *GslbInfoEnhVirtServerTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.GslbInfoEnhVirtServerIdx).IsZero() &&
		reflect.ValueOf(c.GslbInfoEnhVirtServerServiceIdx).IsZero() &&
		reflect.ValueOf(c.GslbInfoEnhVirtServerRserverIdx).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.GslbInfoEnhVirtServerIdx) + "/" + fmt.Sprint(c.GslbInfoEnhVirtServerServiceIdx) + "/" + fmt.Sprint(c.GslbInfoEnhVirtServerRserverIdx)
}

type GslbInfoEnhVirtServerTableRegion int32

const (
	GslbInfoEnhVirtServerTableRegion_Unknown            GslbInfoEnhVirtServerTableRegion = 0
	GslbInfoEnhVirtServerTableRegion_Northamerica       GslbInfoEnhVirtServerTableRegion = 1
	GslbInfoEnhVirtServerTableRegion_Southamerica       GslbInfoEnhVirtServerTableRegion = 2
	GslbInfoEnhVirtServerTableRegion_Europe             GslbInfoEnhVirtServerTableRegion = 3
	GslbInfoEnhVirtServerTableRegion_Caribbean          GslbInfoEnhVirtServerTableRegion = 4
	GslbInfoEnhVirtServerTableRegion_Pacificrim         GslbInfoEnhVirtServerTableRegion = 5
	GslbInfoEnhVirtServerTableRegion_Subsahara          GslbInfoEnhVirtServerTableRegion = 6
	GslbInfoEnhVirtServerTableRegion_Japan              GslbInfoEnhVirtServerTableRegion = 7
	GslbInfoEnhVirtServerTableRegion_Caribbeansubsahara GslbInfoEnhVirtServerTableRegion = 8
	GslbInfoEnhVirtServerTableRegion_Africa             GslbInfoEnhVirtServerTableRegion = 9
	GslbInfoEnhVirtServerTableRegion_Unsupported        GslbInfoEnhVirtServerTableRegion = 2147483647
)

type GslbInfoEnhVirtServerTableParams struct {
	// The virtual server number that identifies the virtual server.
	Idx string `json:"Idx,omitempty"`
	// The virtual server service number that identifies the virtual service.
	ServiceIdx int32 `json:"ServiceIdx,omitempty"`
	// The real server number that identifies a remote real server.
	// An index of 0 indicates virtual server.
	RserverIdx string `json:"RserverIdx,omitempty"`
	// The domain name of the virtual server and remote real server.
	Dname string `json:"Dname,omitempty"`
	// The service port of the virtual server and remote real server.
	VirtPort int32 `json:"VirtPort,omitempty"`
	// The IP address of the virtual server or remote real server.
	IpAddress string `json:"IpAddress,omitempty"`
	// The health check response time of the virtual server or remote real server.
	Response int32 `json:"Response,omitempty"`
	// The available sessions of the virtual server or remote real server.
	SessAvail int32 `json:"SessAvail,omitempty"`
	// The current sessions of the virtual server or remote real server.
	SessCur int32 `json:"SessCur,omitempty"`
	// The maximum sessions supported by the virtual server or remote real server.
	SessMax int32 `json:"SessMax,omitempty"`
	// The sessions utilization (current/max) of the virtual server or remote real server.
	SessUtil uint32 `json:"SessUtil,omitempty"`
	// The CPU utilization of the virtual server or remote real server.
	CpuUtil uint32 `json:"CpuUtil,omitempty"`
	// The remote site of the remote real server. For virtual server, the remote site is 0.
	RemSite int32 `json:"RemSite,omitempty"`
	// The weight of the virtual server or remote real server.
	Weight uint64 `json:"Weight,omitempty"`
	// The availability of the virtual server or remote real server.
	Avail uint64 `json:"Avail,omitempty"`
	// The region of the virtual server or remote real server.
	Region GslbInfoEnhVirtServerTableRegion `json:"Region,omitempty"`
}

func (p GslbInfoEnhVirtServerTableParams) iMABean() {}
