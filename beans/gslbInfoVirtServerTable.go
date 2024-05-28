package beans

import (
	"fmt"
	"reflect"
)

// GslbInfoVirtServerTable The Global SLB virtual server information table.
// Note:This mib is not supported for VX instance of Virtualization.
type GslbInfoVirtServerTable struct {
	// The virtual server number that identifies the virtual server.
	GslbInfoVirtServerIdx int32
	// The virtual server service number that identifies the virtual service.
	GslbInfoVirtServerServiceIdx int32
	// The real server number that identifies a remote real server.
	// An index of 0 indicates virtual server.
	GslbInfoVirtServerRserverIdx int32
	Params                       *GslbInfoVirtServerTableParams
}

func NewGslbInfoVirtServerTableList() *GslbInfoVirtServerTable {
	return &GslbInfoVirtServerTable{}
}

func NewGslbInfoVirtServerTable(
	gslbInfoVirtServerIdx int32,
	gslbInfoVirtServerServiceIdx int32,
	gslbInfoVirtServerRserverIdx int32,
	params *GslbInfoVirtServerTableParams,
) *GslbInfoVirtServerTable {
	return &GslbInfoVirtServerTable{
		GslbInfoVirtServerIdx:        gslbInfoVirtServerIdx,
		GslbInfoVirtServerServiceIdx: gslbInfoVirtServerServiceIdx,
		GslbInfoVirtServerRserverIdx: gslbInfoVirtServerRserverIdx,
		Params:                       params,
	}
}

func (c *GslbInfoVirtServerTable) Name() string {
	return "GslbInfoVirtServerTable"
}

func (c *GslbInfoVirtServerTable) GetParams() BeanType {
	return c.Params
}

func (c *GslbInfoVirtServerTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *GslbInfoVirtServerTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.GslbInfoVirtServerIdx).IsZero() &&
		reflect.ValueOf(c.GslbInfoVirtServerServiceIdx).IsZero() &&
		reflect.ValueOf(c.GslbInfoVirtServerRserverIdx).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.GslbInfoVirtServerIdx) + "/" + fmt.Sprint(c.GslbInfoVirtServerServiceIdx) + "/" + fmt.Sprint(c.GslbInfoVirtServerRserverIdx)
}

type GslbInfoVirtServerTableRegion int32

const (
	GslbInfoVirtServerTableRegion_Unknown            GslbInfoVirtServerTableRegion = 0
	GslbInfoVirtServerTableRegion_Northamerica       GslbInfoVirtServerTableRegion = 1
	GslbInfoVirtServerTableRegion_Southamerica       GslbInfoVirtServerTableRegion = 2
	GslbInfoVirtServerTableRegion_Europe             GslbInfoVirtServerTableRegion = 3
	GslbInfoVirtServerTableRegion_Caribbean          GslbInfoVirtServerTableRegion = 4
	GslbInfoVirtServerTableRegion_Pacificrim         GslbInfoVirtServerTableRegion = 5
	GslbInfoVirtServerTableRegion_Subsahara          GslbInfoVirtServerTableRegion = 6
	GslbInfoVirtServerTableRegion_Japan              GslbInfoVirtServerTableRegion = 7
	GslbInfoVirtServerTableRegion_Caribbeansubsahara GslbInfoVirtServerTableRegion = 8
	GslbInfoVirtServerTableRegion_Africa             GslbInfoVirtServerTableRegion = 9
	GslbInfoVirtServerTableRegion_Unsupported        GslbInfoVirtServerTableRegion = 2147483647
)

type GslbInfoVirtServerTableParams struct {
	// The virtual server number that identifies the virtual server.
	Idx int32 `json:"Idx,omitempty"`
	// The virtual server service number that identifies the virtual service.
	ServiceIdx int32 `json:"ServiceIdx,omitempty"`
	// The real server number that identifies a remote real server.
	// An index of 0 indicates virtual server.
	RserverIdx int32 `json:"RserverIdx,omitempty"`
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
	Region GslbInfoVirtServerTableRegion `json:"Region,omitempty"`
}

func (p GslbInfoVirtServerTableParams) iMABean() {}
