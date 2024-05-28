package beans

import (
	"fmt"
	"reflect"
)

// SlbClusterInfoPeerTable The information of slb cluster.
// Note:This mib is not supported for VX instance of Virtualization.
type SlbClusterInfoPeerTable struct {
	// The index for cluster sync peer switchs.
	SlbClusterInfoPeerIndex int32
	Params                  *SlbClusterInfoPeerTableParams
}

func NewSlbClusterInfoPeerTableList() *SlbClusterInfoPeerTable {
	return &SlbClusterInfoPeerTable{}
}

func NewSlbClusterInfoPeerTable(
	slbClusterInfoPeerIndex int32,
	params *SlbClusterInfoPeerTableParams,
) *SlbClusterInfoPeerTable {
	return &SlbClusterInfoPeerTable{
		SlbClusterInfoPeerIndex: slbClusterInfoPeerIndex,
		Params:                  params,
	}
}

func (c *SlbClusterInfoPeerTable) Name() string {
	return "SlbClusterInfoPeerTable"
}

func (c *SlbClusterInfoPeerTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbClusterInfoPeerTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbClusterInfoPeerTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbClusterInfoPeerIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbClusterInfoPeerIndex)
}

type SlbClusterInfoPeerTableParams struct {
	// The index for cluster sync peer switchs.
	Index int32 `json:"Index,omitempty"`
	// The IP address of the cluster sync peer switch.
	IpAddr string `json:"IpAddr,omitempty"`
	// Total Keep Alives Sent
	KeepAliveSent int32 `json:"KeepAliveSent,omitempty"`
	// Total Keep Alives Sent
	KeepAliveReceived int32 `json:"KeepAliveReceived,omitempty"`
	// Total Keep Alives Sent
	KeepAliveLastReceived int32 `json:"KeepAliveLastReceived,omitempty"`
}

func (p SlbClusterInfoPeerTableParams) iMABean() {}
