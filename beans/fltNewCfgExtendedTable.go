package beans

import (
	"fmt"
	"reflect"
)

// FltNewCfgExtendedTable The filtering table in the new configuration block.
// Note:This mib is not supported for VX instance of Virtualization.
type FltNewCfgExtendedTable struct {
	// The filtering table index.
	FltNewCfgExtendedIndx int32
	Params                *FltNewCfgExtendedTableParams
}

func NewFltNewCfgExtendedTableList() *FltNewCfgExtendedTable {
	return &FltNewCfgExtendedTable{}
}

func NewFltNewCfgExtendedTable(
	fltNewCfgExtendedIndx int32,
	params *FltNewCfgExtendedTableParams,
) *FltNewCfgExtendedTable {
	return &FltNewCfgExtendedTable{
		FltNewCfgExtendedIndx: fltNewCfgExtendedIndx,
		Params:                params,
	}
}

func (c *FltNewCfgExtendedTable) Name() string {
	return "FltNewCfgExtendedTable"
}

func (c *FltNewCfgExtendedTable) GetParams() BeanType {
	return c.Params
}

func (c *FltNewCfgExtendedTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *FltNewCfgExtendedTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.FltNewCfgExtendedIndx).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.FltNewCfgExtendedIndx)
}

type FltNewCfgExtendedTableLayer7DenyState int32

const (
	FltNewCfgExtendedTableLayer7DenyState_Enabled     FltNewCfgExtendedTableLayer7DenyState = 1
	FltNewCfgExtendedTableLayer7DenyState_Disabled    FltNewCfgExtendedTableLayer7DenyState = 2
	FltNewCfgExtendedTableLayer7DenyState_Unsupported FltNewCfgExtendedTableLayer7DenyState = 2147483647
)

type FltNewCfgExtendedTableRadiusWapPersist int32

const (
	FltNewCfgExtendedTableRadiusWapPersist_Enabled     FltNewCfgExtendedTableRadiusWapPersist = 1
	FltNewCfgExtendedTableRadiusWapPersist_Disabled    FltNewCfgExtendedTableRadiusWapPersist = 2
	FltNewCfgExtendedTableRadiusWapPersist_Unsupported FltNewCfgExtendedTableRadiusWapPersist = 2147483647
)

type FltNewCfgExtendedTablePbind int32

const (
	FltNewCfgExtendedTablePbind_Enabled     FltNewCfgExtendedTablePbind = 1
	FltNewCfgExtendedTablePbind_Disabled    FltNewCfgExtendedTablePbind = 2
	FltNewCfgExtendedTablePbind_Unsupported FltNewCfgExtendedTablePbind = 2147483647
)

type FltNewCfgExtendedTablePatternMatch int32

const (
	FltNewCfgExtendedTablePatternMatch_Enabled     FltNewCfgExtendedTablePatternMatch = 1
	FltNewCfgExtendedTablePatternMatch_Disabled    FltNewCfgExtendedTablePatternMatch = 2
	FltNewCfgExtendedTablePatternMatch_Unsupported FltNewCfgExtendedTablePatternMatch = 2147483647
)

type FltNewCfgExtendedTableLayer7DenyMatchAll int32

const (
	FltNewCfgExtendedTableLayer7DenyMatchAll_Enabled     FltNewCfgExtendedTableLayer7DenyMatchAll = 1
	FltNewCfgExtendedTableLayer7DenyMatchAll_Disabled    FltNewCfgExtendedTableLayer7DenyMatchAll = 2
	FltNewCfgExtendedTableLayer7DenyMatchAll_Unsupported FltNewCfgExtendedTableLayer7DenyMatchAll = 2147483647
)

type FltNewCfgExtendedTableLayer7ParseAll int32

const (
	FltNewCfgExtendedTableLayer7ParseAll_Enabled     FltNewCfgExtendedTableLayer7ParseAll = 1
	FltNewCfgExtendedTableLayer7ParseAll_Disabled    FltNewCfgExtendedTableLayer7ParseAll = 2
	FltNewCfgExtendedTableLayer7ParseAll_Unsupported FltNewCfgExtendedTableLayer7ParseAll = 2147483647
)

type FltNewCfgExtendedTableSecurityParseAll int32

const (
	FltNewCfgExtendedTableSecurityParseAll_Enabled     FltNewCfgExtendedTableSecurityParseAll = 1
	FltNewCfgExtendedTableSecurityParseAll_Disabled    FltNewCfgExtendedTableSecurityParseAll = 2
	FltNewCfgExtendedTableSecurityParseAll_Unsupported FltNewCfgExtendedTableSecurityParseAll = 2147483647
)

type FltNewCfgExtendedTableExtended8021pBitsMatch int32

const (
	FltNewCfgExtendedTableExtended8021pBitsMatch_Enabled     FltNewCfgExtendedTableExtended8021pBitsMatch = 1
	FltNewCfgExtendedTableExtended8021pBitsMatch_Disabled    FltNewCfgExtendedTableExtended8021pBitsMatch = 2
	FltNewCfgExtendedTableExtended8021pBitsMatch_Unsupported FltNewCfgExtendedTableExtended8021pBitsMatch = 2147483647
)

type FltNewCfgExtendedTableEgressPip int32

const (
	FltNewCfgExtendedTableEgressPip_Enabled     FltNewCfgExtendedTableEgressPip = 1
	FltNewCfgExtendedTableEgressPip_Disabled    FltNewCfgExtendedTableEgressPip = 2
	FltNewCfgExtendedTableEgressPip_Unsupported FltNewCfgExtendedTableEgressPip = 2147483647
)

type FltNewCfgExtendedTableDbind int32

const (
	FltNewCfgExtendedTableDbind_Enabled     FltNewCfgExtendedTableDbind = 1
	FltNewCfgExtendedTableDbind_Disabled    FltNewCfgExtendedTableDbind = 2
	FltNewCfgExtendedTableDbind_Unsupported FltNewCfgExtendedTableDbind = 2147483647
)

type FltNewCfgExtendedTableReverse int32

const (
	FltNewCfgExtendedTableReverse_Enabled     FltNewCfgExtendedTableReverse = 1
	FltNewCfgExtendedTableReverse_Disabled    FltNewCfgExtendedTableReverse = 2
	FltNewCfgExtendedTableReverse_Unsupported FltNewCfgExtendedTableReverse = 2147483647
)

type FltNewCfgExtendedTableParseChn int32

const (
	FltNewCfgExtendedTableParseChn_Enabled     FltNewCfgExtendedTableParseChn = 1
	FltNewCfgExtendedTableParseChn_Disabled    FltNewCfgExtendedTableParseChn = 2
	FltNewCfgExtendedTableParseChn_Unsupported FltNewCfgExtendedTableParseChn = 2147483647
)

type FltNewCfgExtendedTableSipParsing int32

const (
	FltNewCfgExtendedTableSipParsing_Enabled     FltNewCfgExtendedTableSipParsing = 1
	FltNewCfgExtendedTableSipParsing_Disabled    FltNewCfgExtendedTableSipParsing = 2
	FltNewCfgExtendedTableSipParsing_Unsupported FltNewCfgExtendedTableSipParsing = 2147483647
)

type FltNewCfgExtendedTableSessionMirror int32

const (
	FltNewCfgExtendedTableSessionMirror_Enabled     FltNewCfgExtendedTableSessionMirror = 1
	FltNewCfgExtendedTableSessionMirror_Disabled    FltNewCfgExtendedTableSessionMirror = 2
	FltNewCfgExtendedTableSessionMirror_Unsupported FltNewCfgExtendedTableSessionMirror = 2147483647
)

type FltNewCfgExtendedTableIpVer int32

const (
	FltNewCfgExtendedTableIpVer_Ipv4        FltNewCfgExtendedTableIpVer = 1
	FltNewCfgExtendedTableIpVer_Ipv6        FltNewCfgExtendedTableIpVer = 2
	FltNewCfgExtendedTableIpVer_Unsupported FltNewCfgExtendedTableIpVer = 2147483647
)

type FltNewCfgExtendedTableHdrHash int32

const (
	FltNewCfgExtendedTableHdrHash_None        FltNewCfgExtendedTableHdrHash = 1
	FltNewCfgExtendedTableHdrHash_Headerhash  FltNewCfgExtendedTableHdrHash = 2
	FltNewCfgExtendedTableHdrHash_Unsupported FltNewCfgExtendedTableHdrHash = 2147483647
)

type FltNewCfgExtendedTableL7SipFilt int32

const (
	FltNewCfgExtendedTableL7SipFilt_Enabled     FltNewCfgExtendedTableL7SipFilt = 1
	FltNewCfgExtendedTableL7SipFilt_Disabled    FltNewCfgExtendedTableL7SipFilt = 2
	FltNewCfgExtendedTableL7SipFilt_Unsupported FltNewCfgExtendedTableL7SipFilt = 2147483647
)

type FltNewCfgExtendedTableNbind int32

const (
	FltNewCfgExtendedTableNbind_Enabled     FltNewCfgExtendedTableNbind = 1
	FltNewCfgExtendedTableNbind_Disabled    FltNewCfgExtendedTableNbind = 2
	FltNewCfgExtendedTableNbind_Unsupported FltNewCfgExtendedTableNbind = 2147483647
)

type FltNewCfgExtendedTableL3State int32

const (
	FltNewCfgExtendedTableL3State_Enabled     FltNewCfgExtendedTableL3State = 1
	FltNewCfgExtendedTableL3State_Disabled    FltNewCfgExtendedTableL3State = 2
	FltNewCfgExtendedTableL3State_Unsupported FltNewCfgExtendedTableL3State = 2147483647
)

type FltNewCfgExtendedTableProx int32

const (
	FltNewCfgExtendedTableProx_Enabled     FltNewCfgExtendedTableProx = 1
	FltNewCfgExtendedTableProx_Disabled    FltNewCfgExtendedTableProx = 2
	FltNewCfgExtendedTableProx_Unsupported FltNewCfgExtendedTableProx = 2147483647
)

type FltNewCfgExtendedTableReportState int32

const (
	FltNewCfgExtendedTableReportState_Disabled    FltNewCfgExtendedTableReportState = 0
	FltNewCfgExtendedTableReportState_Enabled     FltNewCfgExtendedTableReportState = 1
	FltNewCfgExtendedTableReportState_Unsupported FltNewCfgExtendedTableReportState = 2147483647
)

type FltNewCfgExtendedTableReportLocation int32

const (
	FltNewCfgExtendedTableReportLocation_None        FltNewCfgExtendedTableReportLocation = 0
	FltNewCfgExtendedTableReportLocation_Clientside  FltNewCfgExtendedTableReportLocation = 1
	FltNewCfgExtendedTableReportLocation_Serverside  FltNewCfgExtendedTableReportLocation = 2
	FltNewCfgExtendedTableReportLocation_Unsupported FltNewCfgExtendedTableReportLocation = 2147483647
)

type FltNewCfgExtendedTableReportPurpose int32

const (
	FltNewCfgExtendedTableReportPurpose_None        FltNewCfgExtendedTableReportPurpose = 0
	FltNewCfgExtendedTableReportPurpose_Bypass      FltNewCfgExtendedTableReportPurpose = 1
	FltNewCfgExtendedTableReportPurpose_Inspect     FltNewCfgExtendedTableReportPurpose = 2
	FltNewCfgExtendedTableReportPurpose_Unsupported FltNewCfgExtendedTableReportPurpose = 2147483647
)

type FltNewCfgExtendedTableReportAppl int32

const (
	FltNewCfgExtendedTableReportAppl_None        FltNewCfgExtendedTableReportAppl = 0
	FltNewCfgExtendedTableReportAppl_Https       FltNewCfgExtendedTableReportAppl = 1
	FltNewCfgExtendedTableReportAppl_Http        FltNewCfgExtendedTableReportAppl = 2
	FltNewCfgExtendedTableReportAppl_Unsupported FltNewCfgExtendedTableReportAppl = 2147483647
)

type FltNewCfgExtendedTableReportDir int32

const (
	FltNewCfgExtendedTableReportDir_Outbound    FltNewCfgExtendedTableReportDir = 0
	FltNewCfgExtendedTableReportDir_Inbound     FltNewCfgExtendedTableReportDir = 1
	FltNewCfgExtendedTableReportDir_None        FltNewCfgExtendedTableReportDir = 2
	FltNewCfgExtendedTableReportDir_Unsupported FltNewCfgExtendedTableReportDir = 2147483647
)

type FltNewCfgExtendedTableDpmReportState int32

const (
	FltNewCfgExtendedTableDpmReportState_Disabled    FltNewCfgExtendedTableDpmReportState = 0
	FltNewCfgExtendedTableDpmReportState_Enabled     FltNewCfgExtendedTableDpmReportState = 1
	FltNewCfgExtendedTableDpmReportState_Unsupported FltNewCfgExtendedTableDpmReportState = 2147483647
)

type FltNewCfgExtendedTableRtSrcTnl int32

const (
	FltNewCfgExtendedTableRtSrcTnl_Enabled     FltNewCfgExtendedTableRtSrcTnl = 1
	FltNewCfgExtendedTableRtSrcTnl_Disabled    FltNewCfgExtendedTableRtSrcTnl = 2
	FltNewCfgExtendedTableRtSrcTnl_Unsupported FltNewCfgExtendedTableRtSrcTnl = 2147483647
)

type FltNewCfgExtendedTableForceBind int32

const (
	FltNewCfgExtendedTableForceBind_Enabled     FltNewCfgExtendedTableForceBind = 1
	FltNewCfgExtendedTableForceBind_Disabled    FltNewCfgExtendedTableForceBind = 2
	FltNewCfgExtendedTableForceBind_Unsupported FltNewCfgExtendedTableForceBind = 2147483647
)

type FltNewCfgExtendedTableParams struct {
	// The filtering table index.
	Indx int32 `json:"Indx,omitempty"`
	// Filter security web application.
	Secwa string `json:"Secwa,omitempty"`
	// Enable or disable layer 7 deny filtering.
	Layer7DenyState FltNewCfgExtendedTableLayer7DenyState `json:"Layer7DenyState,omitempty"`
	// The URL strings selected for Layer 7 deny filters.
	// The selected URL strings are presented in a bitmap format.
	// in receiving order:
	// OCTET 1  OCTET 2  .....
	// xxxxxxxx xxxxxxxx .....
	// ||    || |_ URL Path 9
	// ||    ||
	// ||    ||___ URL Path 8
	// ||    |____ URL Path 7
	// ||      .    .   .
	// ||_________ URL Path 2
	// |__________ URL Path 1
	// where x : 1 - The represented URL string is selected
	// 0 - The represented URL string is not selected
	Layer7DenyUrlBmap string `json:"Layer7DenyUrlBmap,omitempty"`
	// The URL Path (slbCurCfgUrlLbPathIndex) to be added to the
	// Layer 7 deny filter. A zero is returned when read.
	Layer7DenyAddUrl int32 `json:"Layer7DenyAddUrl,omitempty"`
	// The URL Path (slbCurCfgUrlLbPathIndex) to be removed from
	// the Layer 7 deny filter. A zero is returned when read.
	Layer7DenyRemUrl int32 `json:"Layer7DenyRemUrl,omitempty"`
	// The filter ID for GOTO action in the new config.
	GotoFilter int32 `json:"GotoFilter,omitempty"`
	// Enable or disable Radius/WAP persistence.
	RadiusWapPersist FltNewCfgExtendedTableRadiusWapPersist `json:"RadiusWapPersist,omitempty"`
	// Enable or disable filter persistent binding.
	Pbind FltNewCfgExtendedTablePbind `json:"Pbind,omitempty"`
	// The time window for protocol rate limiting (in seconds).
	TimeWindow uint64 `json:"TimeWindow,omitempty"`
	// The hold down duration for protocol rate limiting (in minutes).
	HoldDuration uint64 `json:"HoldDuration,omitempty"`
	// Enable or disable binary pattern matching.
	PatternMatch FltNewCfgExtendedTablePatternMatch `json:"PatternMatch,omitempty"`
	// Enable or disable match-all criteria for L7 deny string matching.
	Layer7DenyMatchAll FltNewCfgExtendedTableLayer7DenyMatchAll `json:"Layer7DenyMatchAll,omitempty"`
	// The client proxy IP address for NAT and REDIR filter.
	ProxyIp string `json:"ProxyIp,omitempty"`
	// Enable or disable layer 7 lookup (parsing) of all packets.
	Layer7ParseAll FltNewCfgExtendedTableLayer7ParseAll `json:"Layer7ParseAll,omitempty"`
	// Enable or disable pattern string lookup (parsing) of all packets.
	SecurityParseAll FltNewCfgExtendedTableSecurityParseAll `json:"SecurityParseAll,omitempty"`
	// The pattern match group.
	// The pattern match groups are presented in a bitmap format.
	// in receiving order:
	// OCTET 1  OCTET 2  .....
	// xxxxxxxx xxxxxxxx .....
	// ||    || |_ pattern match group
	// ||    ||
	// ||    ||___ pattern match group 8
	// ||    |____ pattern match group 7
	// ||      .    .   .
	// ||_________ pattern match group 2
	// |__________ pattern match group 1
	// where x : 1 - The represented pattern match group is selected
	// 0 - The represented pattern match group is not selected
	PatternMatchGroupBmap string `json:"PatternMatchGroupBmap,omitempty"`
	// The pattern match group to be added to the
	// security filter. A zero is returned when read.
	AddPatternMatchGroup int32 `json:"AddPatternMatchGroup,omitempty"`
	// The pattern match group to be to be removed from
	// the security filter. A zero is returned when read.
	RemPatternMatchGroup int32 `json:"RemPatternMatchGroup,omitempty"`
	// The 802.1p bits value to match.
	Extended8021pBitsValue uint64 `json:"Extended8021pBitsValue,omitempty"`
	// Enable or disable matching on 802.1p bits in the packets.
	Extended8021pBitsMatch FltNewCfgExtendedTableExtended8021pBitsMatch `json:"Extended8021pBitsMatch,omitempty"`
	// Set the IP maximum packet length in bytes. A value can be either 0
	// which indicates 'any' length or between 64 and 65535.
	AclIpLength uint64 `json:"AclIpLength,omitempty"`
	// The real server group for IDS load balancing. A value of 0 indicates
	// 	 'none'.
	IdsGroup int32 `json:"IdsGroup,omitempty"`
	// Enable or disable pip selection based on egress port/vlan.
	EgressPip FltNewCfgExtendedTableEgressPip `json:"EgressPip,omitempty"`
	// Enable or disable filter delayed binding.
	Dbind FltNewCfgExtendedTableDbind `json:"Dbind,omitempty"`
	// Filt reverse session BWM contract number.
	RevBwmContract int32 `json:"RevBwmContract,omitempty"`
	// Enable or disable creating session for reverse side traffic.
	Reverse FltNewCfgExtendedTableReverse `json:"Reverse,omitempty"`
	// Enable or disable chained pgroup match criteria for l7 filtering.
	ParseChn FltNewCfgExtendedTableParseChn `json:"ParseChn,omitempty"`
	// BWM contract for SIP RTP traffic.
	RtpBwmContract int32 `json:"RtpBwmContract,omitempty"`
	// Enable or disable SIP NAT.
	SipParsing FltNewCfgExtendedTableSipParsing `json:"SipParsing,omitempty"`
	// Enable or disable session mirroring.
	SessionMirror FltNewCfgExtendedTableSessionMirror `json:"SessionMirror,omitempty"`
	// The type of IP address.
	IpVer FltNewCfgExtendedTableIpVer `json:"IpVer,omitempty"`
	// The source IPv6 address to be filtered. Address should be 4-byte
	// hexadecimal colon notation.
	// Valid IPv6 address should be in any of the following forms
	// xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx or
	// xxxx::xxxx:xxxx:xxxx:xxxx or ::xxxx
	Ipv6Sip string `json:"Ipv6Sip,omitempty"`
	// The prefix length associated with source IP address .
	Ipv6Sprefix uint32 `json:"Ipv6Sprefix,omitempty"`
	// The destination  IPv6 address to be filtered. Address should be 4-byte
	// hexadecimal colon notation.
	// Valid IPv6 address should be in any of the following forms
	// xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx or
	// xxxx::xxxx:xxxx:xxxx:xxxx or ::xxxx
	Ipv6Dip string `json:"Ipv6Dip,omitempty"`
	// The prefix length associated with destination IP address .
	Ipv6Dprefix uint32 `json:"Ipv6Dprefix,omitempty"`
	// The header hash filter.
	HdrHash FltNewCfgExtendedTableHdrHash `json:"HdrHash,omitempty"`
	// The header name of the filter. Headerhash should be
	// enabled before header name is configured.
	HdrName string `json:"HdrName,omitempty"`
	// The header hash length of the filter. Headerhash should be
	// enabled before hash length is configured.
	HdrHashLen uint32 `json:"HdrHashLen,omitempty"`
	// The nat IP address to be filtered.
	NatIp string `json:"NatIp,omitempty"`
	// The nat IPv6 address to be filtered. Address should be 4-byte
	// hexadecimal colon notation.
	// Valid IPv6 address should be in any of the following forms
	// xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx or
	// xxxx::xxxx:xxxx:xxxx:xxxx or ::xxxx
	Ipv6Nip string `json:"Ipv6Nip,omitempty"`
	// The proxy IPv6 address to be filtered. Address should be 4-byte
	// hexadecimal colon notation.
	// Valid IPv6 address should be in any of the following forms
	// xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx or
	// xxxx::xxxx:xxxx:xxxx:xxxx or ::xxxx
	Ipv6Proxy string `json:"Ipv6Proxy,omitempty"`
	// Enable/Disable L7 application SIP UDP filtering.
	L7SipFilt FltNewCfgExtendedTableL7SipFilt `json:"L7SipFilt,omitempty"`
	// Multicast VLAN.
	NatMcastVlan uint32 `json:"NatMcastVlan,omitempty"`
	// Enable or disable subnet binding for redirection.
	Nbind FltNewCfgExtendedTableNbind `json:"Nbind,omitempty"`
	// The L3 filter processing state for this filter.
	L3State FltNewCfgExtendedTableL3State `json:"L3State,omitempty"`
	// The real server group to be redirected to.
	RedirGroup string `json:"RedirGroup,omitempty"`
	// The real server group for IDS load balancing. A NULL string indicates
	// 'none'.
	IdsGroupEnh string `json:"IdsGroupEnh,omitempty"`
	// Filter source Ip. It can be IPv4/IPv6.
	SourceIp string `json:"SourceIp,omitempty"`
	// Filter Destination Ip. It can be IPv4/IPv6.
	DestIp string `json:"DestIp,omitempty"`
	// Filter Mask or Prefix.
	SourceMask string `json:"SourceMask,omitempty"`
	// Filter Mask or Prefix.
	DestMask string `json:"DestMask,omitempty"`
	// Enable or disable Outbound LLB based on proximity.
	Prox FltNewCfgExtendedTableProx `json:"Prox,omitempty"`
	// Enable or disable filter inspect reporting.
	ReportState FltNewCfgExtendedTableReportState `json:"ReportState,omitempty"`
	// Set filter location client/server or none as default value.
	ReportLocation FltNewCfgExtendedTableReportLocation `json:"ReportLocation,omitempty"`
	// Set filter purpose bypass/inspect or none as default value.
	ReportPurpose FltNewCfgExtendedTableReportPurpose `json:"ReportPurpose,omitempty"`
	// Set filter application https/http or none as default value.
	ReportAppl FltNewCfgExtendedTableReportAppl `json:"ReportAppl,omitempty"`
	// Set filter direction none/inbound/outbound.
	ReportDir FltNewCfgExtendedTableReportDir `json:"ReportDir,omitempty"`
	// Filter IcapPol.
	IcapPol string `json:"IcapPol,omitempty"`
	// The traffic event policy of the filter.
	TrafficEventPol string `json:"TrafficEventPol,omitempty"`
	// Enable or disable filter DPm reporting.
	DpmReportState FltNewCfgExtendedTableDpmReportState `json:"DpmReportState,omitempty"`
	// Enable or disable Return to Source Tunnel.
	RtSrcTnl FltNewCfgExtendedTableRtSrcTnl `json:"RtSrcTnl,omitempty"`
	// The aw monitor private key.
	SslawMonPriKey string `json:"SslawMonPriKey,omitempty"`
	// Enable or disable Force bind for proxy.
	ForceBind FltNewCfgExtendedTableForceBind `json:"ForceBind,omitempty"`
}

func (p FltNewCfgExtendedTableParams) iMABean() {}
