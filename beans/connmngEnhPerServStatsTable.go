package beans

import (
	"fmt"
	"reflect"
)

// ConnmngEnhPerServStatsTable A table for connection management statistics per virtual service.
// Note:This mib is not supported for VX instance of Virtualization.
type ConnmngEnhPerServStatsTable struct {
	// Virtual server index.
	ConnmngEnhPerServStatsVirtServIndex string
	// Virtual server service index.
	ConnmngEnhPerServStatsVirtServiceIndex int32
	Params                                 *ConnmngEnhPerServStatsTableParams
}

func NewConnmngEnhPerServStatsTableList() *ConnmngEnhPerServStatsTable {
	return &ConnmngEnhPerServStatsTable{}
}

func NewConnmngEnhPerServStatsTable(
	connmngEnhPerServStatsVirtServIndex string,
	connmngEnhPerServStatsVirtServiceIndex int32,
	params *ConnmngEnhPerServStatsTableParams,
) *ConnmngEnhPerServStatsTable {
	return &ConnmngEnhPerServStatsTable{
		ConnmngEnhPerServStatsVirtServIndex:    connmngEnhPerServStatsVirtServIndex,
		ConnmngEnhPerServStatsVirtServiceIndex: connmngEnhPerServStatsVirtServiceIndex,
		Params:                                 params,
	}
}

func (c *ConnmngEnhPerServStatsTable) Name() string {
	return "ConnmngEnhPerServStatsTable"
}

func (c *ConnmngEnhPerServStatsTable) GetParams() BeanType {
	return c.Params
}

func (c *ConnmngEnhPerServStatsTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *ConnmngEnhPerServStatsTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.ConnmngEnhPerServStatsVirtServIndex).IsZero() &&
		reflect.ValueOf(c.ConnmngEnhPerServStatsVirtServiceIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.ConnmngEnhPerServStatsVirtServIndex) + "/" + fmt.Sprint(c.ConnmngEnhPerServStatsVirtServiceIndex)
}

type ConnmngEnhPerServStatsTableParams struct {
	// Virtual server index.
	VirtServIndex string `json:"VirtServIndex,omitempty"`
	// Virtual server service index.
	VirtServiceIndex int32 `json:"VirtServiceIndex,omitempty"`
	// Virtual server service port number.
	VirtServPort int32 `json:"VirtServPort,omitempty"`
	// Number of concurrent backend server connections for virtual service.
	ServConn int32 `json:"ServConn,omitempty"`
	// Number of client requests passed to AX for virtual service.
	CliReq int32 `json:"CliReq,omitempty"`
	// Connection multiplexing ratio for virtual service.
	MulRatio int32 `json:"MulRatio,omitempty"`
}

func (p ConnmngEnhPerServStatsTableParams) iMABean() {}
