package beans

import (
	"fmt"
	"reflect"
)

// SlbStatSpMaintTable The table of SLB SP maintenance statistics.
// Note:This mib is not supported for VX instance of Virtualization.
type SlbStatSpMaintTable struct {
	// The SP number for which the statistics apply.
	SlbStatSpMaintSpIndex int32
	Params                *SlbStatSpMaintTableParams
}

func NewSlbStatSpMaintTableList() *SlbStatSpMaintTable {
	return &SlbStatSpMaintTable{}
}

func NewSlbStatSpMaintTable(
	slbStatSpMaintSpIndex int32,
	params *SlbStatSpMaintTableParams,
) *SlbStatSpMaintTable {
	return &SlbStatSpMaintTable{
		SlbStatSpMaintSpIndex: slbStatSpMaintSpIndex,
		Params:                params,
	}
}

func (c *SlbStatSpMaintTable) Name() string {
	return "SlbStatSpMaintTable"
}

func (c *SlbStatSpMaintTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbStatSpMaintTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbStatSpMaintTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbStatSpMaintSpIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbStatSpMaintSpIndex)
}

type SlbStatSpMaintTableClear int32

const (
	SlbStatSpMaintTableClear_Ok          SlbStatSpMaintTableClear = 1
	SlbStatSpMaintTableClear_Clear       SlbStatSpMaintTableClear = 2
	SlbStatSpMaintTableClear_Unsupported SlbStatSpMaintTableClear = 2147483647
)

type SlbStatSpMaintTableParams struct {
	// The SP number for which the statistics apply.
	SpIndex int32 `json:"SpIndex,omitempty"`
	// The maximun number of entries per SP in the session table.
	MaximumSessions int32 `json:"MaximumSessions,omitempty"`
	// The current number of sessions on the SP.
	CurBindings uint32 `json:"CurBindings,omitempty"`
	// The 4 second average current number of sessions on the SP.
	CurBindings4Seconds uint32 `json:"CurBindings4Seconds,omitempty"`
	// The 64 second average current number of sessions on the SP.
	CurBindings64Seconds uint32 `json:"CurBindings64Seconds,omitempty"`
	// The count of the number of sessions closed because the server failed on the SP.
	TerminatedSessions uint32 `json:"TerminatedSessions,omitempty"`
	// The total number of binding failures on the SP. A binding failure
	// occurs when a SP runs out of session table entries.
	BindingFails uint32 `json:"BindingFails,omitempty"`
	// The total number of non-TCP/IP frames dropped on the SP.
	NonTcpFrames uint32 `json:"NonTcpFrames,omitempty"`
	// The total number of TCP fragments dropped on the SP.
	TcpFragments uint32 `json:"TcpFragments,omitempty"`
	// The total number of UDP datagrams dropped on the SP.
	UdpDatagrams uint32 `json:"UdpDatagrams,omitempty"`
	// The total number of frames with incorrect VIPs that are dropped on
	// the SP.
	IncorrectVIPs uint32 `json:"IncorrectVIPs,omitempty"`
	// The total number of frames with incorrect Virtual Port that are
	// dropped on the SP.
	IncorrectVports uint32 `json:"IncorrectVports,omitempty"`
	// The total number of frames that are dropped on the SP
	// 	 because no real server is available.
	RealServerNoAvails uint32 `json:"RealServerNoAvails,omitempty"`
	// The total number of frames that are denied on the SP by the filter.
	FilteredDeniedFrames uint32 `json:"FilteredDeniedFrames,omitempty"`
	// The total number of attacks on the SP that frames contain the same source
	// 	 and destination IP addresses.
	LandAttacks uint32 `json:"LandAttacks,omitempty"`
	// The total number of fragment sessions processed on the SP.
	IpFragTotalSessions uint32 `json:"IpFragTotalSessions,omitempty"`
	// The number of sessions in the IP fragment binding table on the SP.
	IpFragCurSessions uint32 `json:"IpFragCurSessions,omitempty"`
	// The total number of fragments discarded on the SP because there
	// is no entry in the IP fragment binding table.
	IpFragDiscards uint32 `json:"IpFragDiscards,omitempty"`
	// The total number of times IP fragment binding table is full on the
	// SP.
	IpFragTableFull uint32 `json:"IpFragTableFull,omitempty"`
	// This is an action object to clear the non-operational SLB statistics
	// on the particular SP.
	// ok(1) is returned when read.
	Clear SlbStatSpMaintTableClear `json:"Clear,omitempty"`
	// The total number of Out of state FIN Packets received on the SP.
	OOSFinPktDrops uint32 `json:"OOSFinPktDrops,omitempty"`
	// Number of Symantec sessions.
	SymSessions uint32 `json:"SymSessions,omitempty"`
	// Number of Symantec valid segments.
	SymValidSegments uint32 `json:"SymValidSegments,omitempty"`
	// Number of Symantec fragment sessions.
	SymFragSessions uint32 `json:"SymFragSessions,omitempty"`
	// Number of Symantec segment allocation fails.
	SymSegAllocFails uint32 `json:"SymSegAllocFails,omitempty"`
	// Number of Symantec buffers allocation fails.
	SymBufferAllocFails uint32 `json:"SymBufferAllocFails,omitempty"`
	// Number of Symantec Connection allocation fails.
	SymConnAllocFails uint32 `json:"SymConnAllocFails,omitempty"`
	// Number of Symantec Invalid buffers.
	SymInvalidBuffers uint32 `json:"SymInvalidBuffers,omitempty"`
	// Number of Symantec Segment reallocation fails.
	SymSegReallocFails uint32 `json:"SymSegReallocFails,omitempty"`
	// Symantec Deuce Inspection Stats - No. of packets in.
	SymPacketsIn uint32 `json:"SymPacketsIn,omitempty"`
	// Symantec Deuce Inspection Stats - No. of packets with no data.
	SymPacketsWithNoData uint32 `json:"SymPacketsWithNoData,omitempty"`
	// Symantec Deuce Inspection Stats - No. of TCP packets
	SymTcpPackets uint32 `json:"SymTcpPackets,omitempty"`
	// Symantec Deuce Inspection Stats - No. of  UDP Packets
	SymUdpPackets uint32 `json:"SymUdpPackets,omitempty"`
	// Symantec Deuce Inspection Stats - No. of  ICMP Packets
	SymIcmpPackets uint32 `json:"SymIcmpPackets,omitempty"`
	// Symantec Deuce Inspection Stats - No. of  Packets other than
	// 	 TCP UDP or ICMP
	SymOtherPackets uint32 `json:"SymOtherPackets,omitempty"`
	// Symantec Deuce Inspection Stats - Match Count
	SymMatchCount uint32 `json:"SymMatchCount,omitempty"`
	// Symantec Deuce Inspection Stats - No. of Fetch Errors
	SymFetchErrors uint32 `json:"SymFetchErrors,omitempty"`
	// Symantec Deuce Inspection Stats - No. of Truncated payload to MP
	SymTruncPayloadToMp uint32 `json:"SymTruncPayloadToMp,omitempty"`
	// Symantec Deuce Inspection Stats - No. of Packets in fast path
	SymPacketsInFastPath uint32 `json:"SymPacketsInFastPath,omitempty"`
	// The peak number of sessions on the SP.
	PeakBindings uint32 `json:"PeakBindings,omitempty"`
	// The current number of Ax sessions on the SP.
	CurAxBindings uint32 `json:"CurAxBindings,omitempty"`
	// The 4 second average current number of Ax sessions on the SP.
	CurAxBindings4Seconds uint32 `json:"CurAxBindings4Seconds,omitempty"`
	// The 64 second average current number of Ax sessions on the SP.
	CurAxBindings64Seconds uint32 `json:"CurAxBindings64Seconds,omitempty"`
	// The peak number of Ax sessions on the SP.
	PeakAxBindings uint32 `json:"PeakAxBindings,omitempty"`
}

func (p SlbStatSpMaintTableParams) iMABean() {}
