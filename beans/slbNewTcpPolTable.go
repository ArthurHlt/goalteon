package beans

import (
	"fmt"
	"reflect"
)

// SlbNewTcpPolTable The table for configuring TCP policy.
// Note:This mib is not supported on VX instance of Virtualization.
type SlbNewTcpPolTable struct {
	// The TCP policy name(key id) as an index.
	SlbNewTcpPolNameIdIndex string
	Params                  *SlbNewTcpPolTableParams
}

func NewSlbNewTcpPolTableList() *SlbNewTcpPolTable {
	return &SlbNewTcpPolTable{}
}

func NewSlbNewTcpPolTable(
	slbNewTcpPolNameIdIndex string,
	params *SlbNewTcpPolTableParams,
) *SlbNewTcpPolTable {
	return &SlbNewTcpPolTable{
		SlbNewTcpPolNameIdIndex: slbNewTcpPolNameIdIndex,
		Params:                  params,
	}
}

func (c *SlbNewTcpPolTable) Name() string {
	return "SlbNewTcpPolTable"
}

func (c *SlbNewTcpPolTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbNewTcpPolTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbNewTcpPolTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbNewTcpPolNameIdIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbNewTcpPolNameIdIndex)
}

type SlbNewTcpPolTableRcvWndSize int32

const (
	SlbNewTcpPolTableRcvWndSize_Sz128K      SlbNewTcpPolTableRcvWndSize = 0
	SlbNewTcpPolTableRcvWndSize_Sz256K      SlbNewTcpPolTableRcvWndSize = 1
	SlbNewTcpPolTableRcvWndSize_Sz512K      SlbNewTcpPolTableRcvWndSize = 2
	SlbNewTcpPolTableRcvWndSize_Sz1M        SlbNewTcpPolTableRcvWndSize = 3
	SlbNewTcpPolTableRcvWndSize_Sz2M        SlbNewTcpPolTableRcvWndSize = 4
	SlbNewTcpPolTableRcvWndSize_Sz3M        SlbNewTcpPolTableRcvWndSize = 5
	SlbNewTcpPolTableRcvWndSize_Sz4M        SlbNewTcpPolTableRcvWndSize = 6
	SlbNewTcpPolTableRcvWndSize_Unsupported SlbNewTcpPolTableRcvWndSize = 2147483647
)

type SlbNewTcpPolTableSndWndSize int32

const (
	SlbNewTcpPolTableSndWndSize_Sz128K      SlbNewTcpPolTableSndWndSize = 0
	SlbNewTcpPolTableSndWndSize_Sz256K      SlbNewTcpPolTableSndWndSize = 1
	SlbNewTcpPolTableSndWndSize_Sz512K      SlbNewTcpPolTableSndWndSize = 2
	SlbNewTcpPolTableSndWndSize_Sz1M        SlbNewTcpPolTableSndWndSize = 3
	SlbNewTcpPolTableSndWndSize_Sz2M        SlbNewTcpPolTableSndWndSize = 4
	SlbNewTcpPolTableSndWndSize_Sz3M        SlbNewTcpPolTableSndWndSize = 5
	SlbNewTcpPolTableSndWndSize_Sz4M        SlbNewTcpPolTableSndWndSize = 6
	SlbNewTcpPolTableSndWndSize_Unsupported SlbNewTcpPolTableSndWndSize = 2147483647
)

type SlbNewTcpPolTableMss int32

const (
	SlbNewTcpPolTableMss_Def         SlbNewTcpPolTableMss = 0
	SlbNewTcpPolTableMss_Sz1460      SlbNewTcpPolTableMss = 1
	SlbNewTcpPolTableMss_Sz1440      SlbNewTcpPolTableMss = 2
	SlbNewTcpPolTableMss_Sz1360      SlbNewTcpPolTableMss = 3
	SlbNewTcpPolTableMss_Sz1212      SlbNewTcpPolTableMss = 4
	SlbNewTcpPolTableMss_Sz956       SlbNewTcpPolTableMss = 5
	SlbNewTcpPolTableMss_Sz536       SlbNewTcpPolTableMss = 6
	SlbNewTcpPolTableMss_Sz384       SlbNewTcpPolTableMss = 7
	SlbNewTcpPolTableMss_Sz128       SlbNewTcpPolTableMss = 8
	SlbNewTcpPolTableMss_Sz1400      SlbNewTcpPolTableMss = 9
	SlbNewTcpPolTableMss_Sz1600      SlbNewTcpPolTableMss = 10
	SlbNewTcpPolTableMss_Sz9156      SlbNewTcpPolTableMss = 11
	SlbNewTcpPolTableMss_Sz9176      SlbNewTcpPolTableMss = 12
	SlbNewTcpPolTableMss_Sz8940      SlbNewTcpPolTableMss = 13
	SlbNewTcpPolTableMss_Sz8960      SlbNewTcpPolTableMss = 14
	SlbNewTcpPolTableMss_Unsupported SlbNewTcpPolTableMss = 2147483647
)

type SlbNewTcpPolTableCaAlgorithm int32

const (
	SlbNewTcpPolTableCaAlgorithm_Reno           SlbNewTcpPolTableCaAlgorithm = 0
	SlbNewTcpPolTableCaAlgorithm_Hybla          SlbNewTcpPolTableCaAlgorithm = 1
	SlbNewTcpPolTableCaAlgorithm_HyblaAndPacing SlbNewTcpPolTableCaAlgorithm = 2
	SlbNewTcpPolTableCaAlgorithm_Westwood       SlbNewTcpPolTableCaAlgorithm = 3
	SlbNewTcpPolTableCaAlgorithm_Unsupported    SlbNewTcpPolTableCaAlgorithm = 2147483647
)

type SlbNewTcpPolTableAdaptiveTuning int32

const (
	SlbNewTcpPolTableAdaptiveTuning_Enabled     SlbNewTcpPolTableAdaptiveTuning = 1
	SlbNewTcpPolTableAdaptiveTuning_Disabled    SlbNewTcpPolTableAdaptiveTuning = 2
	SlbNewTcpPolTableAdaptiveTuning_Unsupported SlbNewTcpPolTableAdaptiveTuning = 2147483647
)

type SlbNewTcpPolTableSelAck int32

const (
	SlbNewTcpPolTableSelAck_Enabled     SlbNewTcpPolTableSelAck = 1
	SlbNewTcpPolTableSelAck_Disabled    SlbNewTcpPolTableSelAck = 2
	SlbNewTcpPolTableSelAck_Unsupported SlbNewTcpPolTableSelAck = 2147483647
)

type SlbNewTcpPolTableState int32

const (
	SlbNewTcpPolTableState_Enabled     SlbNewTcpPolTableState = 1
	SlbNewTcpPolTableState_Disabled    SlbNewTcpPolTableState = 2
	SlbNewTcpPolTableState_Unsupported SlbNewTcpPolTableState = 2147483647
)

type SlbNewTcpPolTableDel int32

const (
	SlbNewTcpPolTableDel_Other       SlbNewTcpPolTableDel = 1
	SlbNewTcpPolTableDel_Delete      SlbNewTcpPolTableDel = 2
	SlbNewTcpPolTableDel_Unsupported SlbNewTcpPolTableDel = 2147483647
)

type SlbNewTcpPolTableImmediateAck int32

const (
	SlbNewTcpPolTableImmediateAck_Enabled     SlbNewTcpPolTableImmediateAck = 1
	SlbNewTcpPolTableImmediateAck_Disabled    SlbNewTcpPolTableImmediateAck = 2
	SlbNewTcpPolTableImmediateAck_Unsupported SlbNewTcpPolTableImmediateAck = 2147483647
)

type SlbNewTcpPolTableNagle int32

const (
	SlbNewTcpPolTableNagle_Enabled     SlbNewTcpPolTableNagle = 1
	SlbNewTcpPolTableNagle_Disabled    SlbNewTcpPolTableNagle = 2
	SlbNewTcpPolTableNagle_Unsupported SlbNewTcpPolTableNagle = 2147483647
)

type SlbNewTcpPolTableTimestamp int32

const (
	SlbNewTcpPolTableTimestamp_Enabled     SlbNewTcpPolTableTimestamp = 1
	SlbNewTcpPolTableTimestamp_Disabled    SlbNewTcpPolTableTimestamp = 2
	SlbNewTcpPolTableTimestamp_Unsupported SlbNewTcpPolTableTimestamp = 2147483647
)

type SlbNewTcpPolTableTcpKeepalive int32

const (
	SlbNewTcpPolTableTcpKeepalive_Enabled     SlbNewTcpPolTableTcpKeepalive = 1
	SlbNewTcpPolTableTcpKeepalive_Disabled    SlbNewTcpPolTableTcpKeepalive = 2
	SlbNewTcpPolTableTcpKeepalive_Unsupported SlbNewTcpPolTableTcpKeepalive = 2147483647
)

type SlbNewTcpPolTableParams struct {
	// The TCP policy name(key id) as an index.
	NameIdIndex string `json:"NameIdIndex,omitempty"`
	// TCP policy name,length of the string should be 32 characters.
	Name string `json:"Name,omitempty"`
	// TCP read buffer size.
	RcvWndSize SlbNewTcpPolTableRcvWndSize `json:"RcvWndSize,omitempty"`
	// TCP write buffer size.
	SndWndSize SlbNewTcpPolTableSndWndSize `json:"SndWndSize,omitempty"`
	// TCP maximum segment size.
	Mss SlbNewTcpPolTableMss `json:"Mss,omitempty"`
	// TCP congestion avoidance algorithm.
	CaAlgorithm SlbNewTcpPolTableCaAlgorithm `json:"CaAlgorithm,omitempty"`
	// Enable/Disable TCP Adaptive Tuning.
	AdaptiveTuning SlbNewTcpPolTableAdaptiveTuning `json:"AdaptiveTuning,omitempty"`
	// Enable/Disable selective ACK.
	SelAck SlbNewTcpPolTableSelAck `json:"SelAck,omitempty"`
	// Enable/Disable TCP policy.
	State SlbNewTcpPolTableState `json:"State,omitempty"`
	// Delete TCP policy.
	Del SlbNewTcpPolTableDel `json:"Del,omitempty"`
	// Read buffer size.
	ReadBufferSize uint64 `json:"ReadBufferSize,omitempty"`
	// Congestion Scale.
	CaScale uint64 `json:"CaScale,omitempty"`
	// Congestion Decrease policy.
	CaDecrease uint32 `json:"CaDecrease,omitempty"`
	// Enable/Disable Immediate ACK on push.
	ImmediateAck SlbNewTcpPolTableImmediateAck `json:"ImmediateAck,omitempty"`
	// Enable/Disable Nagle.
	Nagle SlbNewTcpPolTableNagle `json:"Nagle,omitempty"`
	// Connection closing aging in seconds on FIN receipt (0 - no aging).
	Finage uint64 `json:"Finage,omitempty"`
	// Enable/Disable TCP timestamp.
	Timestamp SlbNewTcpPolTableTimestamp `json:"Timestamp,omitempty"`
	// Enable/Disable TCP keepalive.
	TcpKeepalive SlbNewTcpPolTableTcpKeepalive `json:"TcpKeepalive,omitempty"`
	// TCP Keepalive idle
	TcpkaIdle uint64 `json:"TcpkaIdle,omitempty"`
	// TCP Keepalive count
	TcpkaCount uint32 `json:"TcpkaCount,omitempty"`
	// TCP Keepalive interval
	TcpkaInterval uint32 `json:"TcpkaInterval,omitempty"`
}

func (p SlbNewTcpPolTableParams) iMABean() {}
