package beans

import (
	"fmt"
	"reflect"
)

// ConnmngPerServStatsTable A table for connection management statistics per virtual service.
// Note:This mib is not supported for VX instance of Virtualization.
type ConnmngPerServStatsTable struct {
	// Virtual server number.
	ConnmngPerServStatsVirtServIndex int32
	// Virtual server service index.
	ConnmngPerServStatsVirtServiceIndex int32
	Params                              *ConnmngPerServStatsTableParams
}

func NewConnmngPerServStatsTableList() *ConnmngPerServStatsTable {
	return &ConnmngPerServStatsTable{}
}

func NewConnmngPerServStatsTable(
	connmngPerServStatsVirtServIndex int32,
	connmngPerServStatsVirtServiceIndex int32,
	params *ConnmngPerServStatsTableParams,
) *ConnmngPerServStatsTable {
	return &ConnmngPerServStatsTable{
		ConnmngPerServStatsVirtServIndex:    connmngPerServStatsVirtServIndex,
		ConnmngPerServStatsVirtServiceIndex: connmngPerServStatsVirtServiceIndex,
		Params:                              params,
	}
}

func (c *ConnmngPerServStatsTable) Name() string {
	return "ConnmngPerServStatsTable"
}

func (c *ConnmngPerServStatsTable) GetParams() BeanType {
	return c.Params
}

func (c *ConnmngPerServStatsTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *ConnmngPerServStatsTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.ConnmngPerServStatsVirtServIndex).IsZero() &&
		reflect.ValueOf(c.ConnmngPerServStatsVirtServiceIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.ConnmngPerServStatsVirtServIndex) + "/" + fmt.Sprint(c.ConnmngPerServStatsVirtServiceIndex)
}

type ConnmngPerServStatsTableParams struct {
	// Virtual server number.
	VirtServIndex int32 `json:"VirtServIndex,omitempty"`
	// Virtual server service index.
	VirtServiceIndex int32 `json:"VirtServiceIndex,omitempty"`
	// Virtual server service port number.
	VirtServPort int32 `json:"VirtServPort,omitempty"`
	// Number of concurrent backend server connections for virtual service.
	ServConn int32 `json:"ServConn,omitempty"`
	// Number of client requests passed to AE for virtual service.
	CliReq int32 `json:"CliReq,omitempty"`
	// Connection multiplexing ratio for virtual service.
	MulRatio int32 `json:"MulRatio,omitempty"`
}

func (p ConnmngPerServStatsTableParams) iMABean() {}
