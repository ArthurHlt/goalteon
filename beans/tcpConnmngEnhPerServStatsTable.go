package beans

import (
	"fmt"
	"reflect"
)

// TcpConnmngEnhPerServStatsTable A table for tcp connection management statistics per virtual service.
// Note:This mib is not supported for VX instance of Virtualization.
type TcpConnmngEnhPerServStatsTable struct {
	// Virtual server index.
	TcpConnmngEnhPerServStatsVirtServIndex string
	// Virtual server service index.
	TcpConnmngEnhPerServStatsVirtServiceIndex int32
	Params                                    *TcpConnmngEnhPerServStatsTableParams
}

func NewTcpConnmngEnhPerServStatsTableList() *TcpConnmngEnhPerServStatsTable {
	return &TcpConnmngEnhPerServStatsTable{}
}

func NewTcpConnmngEnhPerServStatsTable(
	tcpConnmngEnhPerServStatsVirtServIndex string,
	tcpConnmngEnhPerServStatsVirtServiceIndex int32,
	params *TcpConnmngEnhPerServStatsTableParams,
) *TcpConnmngEnhPerServStatsTable {
	return &TcpConnmngEnhPerServStatsTable{
		TcpConnmngEnhPerServStatsVirtServIndex:    tcpConnmngEnhPerServStatsVirtServIndex,
		TcpConnmngEnhPerServStatsVirtServiceIndex: tcpConnmngEnhPerServStatsVirtServiceIndex,
		Params: params,
	}
}

func (c *TcpConnmngEnhPerServStatsTable) Name() string {
	return "TcpConnmngEnhPerServStatsTable"
}

func (c *TcpConnmngEnhPerServStatsTable) GetParams() BeanType {
	return c.Params
}

func (c *TcpConnmngEnhPerServStatsTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *TcpConnmngEnhPerServStatsTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.TcpConnmngEnhPerServStatsVirtServIndex).IsZero() &&
		reflect.ValueOf(c.TcpConnmngEnhPerServStatsVirtServiceIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.TcpConnmngEnhPerServStatsVirtServIndex) + "/" + fmt.Sprint(c.TcpConnmngEnhPerServStatsVirtServiceIndex)
}

type TcpConnmngEnhPerServStatsTableParams struct {
	// Virtual server index.
	VirtServIndex string `json:"VirtServIndex,omitempty"`
	// Virtual server service index.
	VirtServiceIndex int32 `json:"VirtServiceIndex,omitempty"`
	// Virtual server service port number.
	VirtServPort int32 `json:"VirtServPort,omitempty"`
	// Number of concurrent backend server connections for virtual service.
	ServConn int32 `json:"ServConn,omitempty"`
	// Number of reused backend server connections for virtual service.
	ServConnReuse int32 `json:"ServConnReuse,omitempty"`
	// Number of client transactions passed to AX for virtual service.
	CliTrans int32 `json:"CliTrans,omitempty"`
	// Connection multiplexing ratio for virtual service.
	MulRatio int32 `json:"MulRatio,omitempty"`
	// Total number of new backend server connections for virtual service.
	TotalServConn uint64 `json:"TotalServConn,omitempty"`
	// Total number of reused backend server connections for virtual service.
	TotalServConnReuse uint64 `json:"TotalServConnReuse,omitempty"`
	// Total number of client transactions passed to AX for virtual service.
	TotalCliTrans uint64 `json:"TotalCliTrans,omitempty"`
	// Connection multiplexing ratio on total numbers for virtual service.
	TotalMulRatio int32 `json:"TotalMulRatio,omitempty"`
}

func (p TcpConnmngEnhPerServStatsTableParams) iMABean() {}
