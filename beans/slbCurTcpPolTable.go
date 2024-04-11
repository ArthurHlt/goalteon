package beans

import (
	"fmt"
	"reflect"
)

// SlbCurTcpPolTable The table for configuring TCP policy.
// Note:This mib is not supported on VX instance of Virtualization.
type SlbCurTcpPolTable struct {
	// The TCP policy name(key id) as an index.
	SlbCurTcpPolNameIdIndex string
	Params                  *SlbCurTcpPolTableParams
}

func NewSlbCurTcpPolTableList() *SlbCurTcpPolTable {
	return &SlbCurTcpPolTable{}
}

func NewSlbCurTcpPolTable(
	slbCurTcpPolNameIdIndex string,
	params *SlbCurTcpPolTableParams,
) *SlbCurTcpPolTable {
	return &SlbCurTcpPolTable{
		SlbCurTcpPolNameIdIndex: slbCurTcpPolNameIdIndex,
		Params:                  params,
	}
}

func (c *SlbCurTcpPolTable) Name() string {
	return "SlbCurTcpPolTable"
}

func (c *SlbCurTcpPolTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbCurTcpPolTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbCurTcpPolTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbCurTcpPolNameIdIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbCurTcpPolNameIdIndex)
}

type SlbCurTcpPolTableRcvWndSize int32

const (
	SlbCurTcpPolTableRcvWndSize_Sz128K      SlbCurTcpPolTableRcvWndSize = 0
	SlbCurTcpPolTableRcvWndSize_Sz256K      SlbCurTcpPolTableRcvWndSize = 1
	SlbCurTcpPolTableRcvWndSize_Sz512K      SlbCurTcpPolTableRcvWndSize = 2
	SlbCurTcpPolTableRcvWndSize_Sz1M        SlbCurTcpPolTableRcvWndSize = 3
	SlbCurTcpPolTableRcvWndSize_Sz2M        SlbCurTcpPolTableRcvWndSize = 4
	SlbCurTcpPolTableRcvWndSize_Sz3M        SlbCurTcpPolTableRcvWndSize = 5
	SlbCurTcpPolTableRcvWndSize_Sz4M        SlbCurTcpPolTableRcvWndSize = 6
	SlbCurTcpPolTableRcvWndSize_Unsupported SlbCurTcpPolTableRcvWndSize = 2147483647
)

type SlbCurTcpPolTableSndWndSize int32

const (
	SlbCurTcpPolTableSndWndSize_Sz128K      SlbCurTcpPolTableSndWndSize = 0
	SlbCurTcpPolTableSndWndSize_Sz256K      SlbCurTcpPolTableSndWndSize = 1
	SlbCurTcpPolTableSndWndSize_Sz512K      SlbCurTcpPolTableSndWndSize = 2
	SlbCurTcpPolTableSndWndSize_Sz1M        SlbCurTcpPolTableSndWndSize = 3
	SlbCurTcpPolTableSndWndSize_Sz2M        SlbCurTcpPolTableSndWndSize = 4
	SlbCurTcpPolTableSndWndSize_Sz3M        SlbCurTcpPolTableSndWndSize = 5
	SlbCurTcpPolTableSndWndSize_Sz4M        SlbCurTcpPolTableSndWndSize = 6
	SlbCurTcpPolTableSndWndSize_Unsupported SlbCurTcpPolTableSndWndSize = 2147483647
)

type SlbCurTcpPolTableMss int32

const (
	SlbCurTcpPolTableMss_Def         SlbCurTcpPolTableMss = 0
	SlbCurTcpPolTableMss_Sz1460      SlbCurTcpPolTableMss = 1
	SlbCurTcpPolTableMss_Sz1440      SlbCurTcpPolTableMss = 2
	SlbCurTcpPolTableMss_Sz1360      SlbCurTcpPolTableMss = 3
	SlbCurTcpPolTableMss_Sz1212      SlbCurTcpPolTableMss = 4
	SlbCurTcpPolTableMss_Sz956       SlbCurTcpPolTableMss = 5
	SlbCurTcpPolTableMss_Sz536       SlbCurTcpPolTableMss = 6
	SlbCurTcpPolTableMss_Sz384       SlbCurTcpPolTableMss = 7
	SlbCurTcpPolTableMss_Sz128       SlbCurTcpPolTableMss = 8
	SlbCurTcpPolTableMss_Sz1400      SlbCurTcpPolTableMss = 9
	SlbCurTcpPolTableMss_Sz1600      SlbCurTcpPolTableMss = 10
	SlbCurTcpPolTableMss_Sz9156      SlbCurTcpPolTableMss = 11
	SlbCurTcpPolTableMss_Sz9176      SlbCurTcpPolTableMss = 12
	SlbCurTcpPolTableMss_Sz8940      SlbCurTcpPolTableMss = 13
	SlbCurTcpPolTableMss_Sz8960      SlbCurTcpPolTableMss = 14
	SlbCurTcpPolTableMss_Unsupported SlbCurTcpPolTableMss = 2147483647
)

type SlbCurTcpPolTableCaAlgorithm int32

const (
	SlbCurTcpPolTableCaAlgorithm_Reno           SlbCurTcpPolTableCaAlgorithm = 0
	SlbCurTcpPolTableCaAlgorithm_Hybla          SlbCurTcpPolTableCaAlgorithm = 1
	SlbCurTcpPolTableCaAlgorithm_HyblaAndPacing SlbCurTcpPolTableCaAlgorithm = 2
	SlbCurTcpPolTableCaAlgorithm_Westwood       SlbCurTcpPolTableCaAlgorithm = 3
	SlbCurTcpPolTableCaAlgorithm_Unsupported    SlbCurTcpPolTableCaAlgorithm = 2147483647
)

type SlbCurTcpPolTableAdaptiveTuning int32

const (
	SlbCurTcpPolTableAdaptiveTuning_Enabled     SlbCurTcpPolTableAdaptiveTuning = 1
	SlbCurTcpPolTableAdaptiveTuning_Disabled    SlbCurTcpPolTableAdaptiveTuning = 2
	SlbCurTcpPolTableAdaptiveTuning_Unsupported SlbCurTcpPolTableAdaptiveTuning = 2147483647
)

type SlbCurTcpPolTableSelAck int32

const (
	SlbCurTcpPolTableSelAck_Enabled     SlbCurTcpPolTableSelAck = 1
	SlbCurTcpPolTableSelAck_Disabled    SlbCurTcpPolTableSelAck = 2
	SlbCurTcpPolTableSelAck_Unsupported SlbCurTcpPolTableSelAck = 2147483647
)

type SlbCurTcpPolTableState int32

const (
	SlbCurTcpPolTableState_Enabled     SlbCurTcpPolTableState = 1
	SlbCurTcpPolTableState_Disabled    SlbCurTcpPolTableState = 2
	SlbCurTcpPolTableState_Unsupported SlbCurTcpPolTableState = 2147483647
)

type SlbCurTcpPolTableImmediateAck int32

const (
	SlbCurTcpPolTableImmediateAck_Enabled     SlbCurTcpPolTableImmediateAck = 1
	SlbCurTcpPolTableImmediateAck_Disabled    SlbCurTcpPolTableImmediateAck = 2
	SlbCurTcpPolTableImmediateAck_Unsupported SlbCurTcpPolTableImmediateAck = 2147483647
)

type SlbCurTcpPolTableNagle int32

const (
	SlbCurTcpPolTableNagle_Enabled     SlbCurTcpPolTableNagle = 1
	SlbCurTcpPolTableNagle_Disabled    SlbCurTcpPolTableNagle = 2
	SlbCurTcpPolTableNagle_Unsupported SlbCurTcpPolTableNagle = 2147483647
)

type SlbCurTcpPolTableParams struct {
	// The TCP policy name(key id) as an index.
	NameIdIndex string `json:"NameIdIndex,omitempty"`
	// TCP policy name.
	Name string `json:"Name,omitempty"`
	// TCP read buffer size.
	RcvWndSize SlbCurTcpPolTableRcvWndSize `json:"RcvWndSize,omitempty"`
	// TCP write buffer size.
	SndWndSize SlbCurTcpPolTableSndWndSize `json:"SndWndSize,omitempty"`
	// TCP maximum segment size.
	Mss SlbCurTcpPolTableMss `json:"Mss,omitempty"`
	// TCP congestion avoidance algorithm.
	CaAlgorithm SlbCurTcpPolTableCaAlgorithm `json:"CaAlgorithm,omitempty"`
	// Enable/Disable TCP Adaptive Tuning.
	AdaptiveTuning SlbCurTcpPolTableAdaptiveTuning `json:"AdaptiveTuning,omitempty"`
	// Enable/Disable selective ACK.
	SelAck SlbCurTcpPolTableSelAck `json:"SelAck,omitempty"`
	// Enable/Disable TCP policy.
	State SlbCurTcpPolTableState `json:"State,omitempty"`
	// Read buffer size.
	ReadBufferSize uint64 `json:"ReadBufferSize,omitempty"`
	// Congestion Scale.
	CaScale uint64 `json:"CaScale,omitempty"`
	// Congestion Decrease policy.
	CaDecrease uint32 `json:"CaDecrease,omitempty"`
	// Enable/Disable Immediate ACK on push.
	ImmediateAck SlbCurTcpPolTableImmediateAck `json:"ImmediateAck,omitempty"`
	// Enable/Disable Nagle Algorithm.
	Nagle SlbCurTcpPolTableNagle `json:"Nagle,omitempty"`
	// Connection closing aging in seconds on FIN receipt (0 - no aging).
	Finage uint64 `json:"Finage,omitempty"`
}

func (p SlbCurTcpPolTableParams) iMABean() {}
