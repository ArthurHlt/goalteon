package beans

import (
	"fmt"
	"reflect"
)

// FltNewCfgTable The filtering table in the new configuration block.
// Note:This mib is not supported for VX instance of Virtualization.
type FltNewCfgTable struct {
	// The filtering table index.
	FltNewCfgIndx int32
	Params        *FltNewCfgTableParams
}

func NewFltNewCfgTableList() *FltNewCfgTable {
	return &FltNewCfgTable{}
}

func NewFltNewCfgTable(
	fltNewCfgIndx int32,
	params *FltNewCfgTableParams,
) *FltNewCfgTable {
	return &FltNewCfgTable{
		FltNewCfgIndx: fltNewCfgIndx,
		Params:        params,
	}
}

func (c *FltNewCfgTable) Name() string {
	return "FltNewCfgTable"
}

func (c *FltNewCfgTable) GetParams() BeanType {
	return c.Params
}

func (c *FltNewCfgTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *FltNewCfgTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.FltNewCfgIndx).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.FltNewCfgIndx)
}

type FltNewCfgTableAction int32

const (
	FltNewCfgTableAction_Allow       FltNewCfgTableAction = 1
	FltNewCfgTableAction_Deny        FltNewCfgTableAction = 2
	FltNewCfgTableAction_Redirect    FltNewCfgTableAction = 3
	FltNewCfgTableAction_Nat         FltNewCfgTableAction = 4
	FltNewCfgTableAction_Goto        FltNewCfgTableAction = 5
	FltNewCfgTableAction_OutboundLlb FltNewCfgTableAction = 6
	FltNewCfgTableAction_Monitor     FltNewCfgTableAction = 7
	FltNewCfgTableAction_Unsupported FltNewCfgTableAction = 2147483647
)

type FltNewCfgTableLog int32

const (
	FltNewCfgTableLog_Enabled     FltNewCfgTableLog = 1
	FltNewCfgTableLog_Disabled    FltNewCfgTableLog = 2
	FltNewCfgTableLog_Unsupported FltNewCfgTableLog = 2147483647
)

type FltNewCfgTableState int32

const (
	FltNewCfgTableState_Enabled     FltNewCfgTableState = 1
	FltNewCfgTableState_Disabled    FltNewCfgTableState = 2
	FltNewCfgTableState_Unsupported FltNewCfgTableState = 2147483647
)

type FltNewCfgTableDelete int32

const (
	FltNewCfgTableDelete_Other       FltNewCfgTableDelete = 1
	FltNewCfgTableDelete_Delete      FltNewCfgTableDelete = 2
	FltNewCfgTableDelete_Unsupported FltNewCfgTableDelete = 2147483647
)

type FltNewCfgTableNat int32

const (
	FltNewCfgTableNat_DestinationAddress FltNewCfgTableNat = 1
	FltNewCfgTableNat_SourceAddress      FltNewCfgTableNat = 2
	FltNewCfgTableNat_MulticastAddress   FltNewCfgTableNat = 3
	FltNewCfgTableNat_Unsupported        FltNewCfgTableNat = 2147483647
)

type FltNewCfgTableCache int32

const (
	FltNewCfgTableCache_Enabled     FltNewCfgTableCache = 1
	FltNewCfgTableCache_Disabled    FltNewCfgTableCache = 2
	FltNewCfgTableCache_Unsupported FltNewCfgTableCache = 2147483647
)

type FltNewCfgTableInvert int32

const (
	FltNewCfgTableInvert_Enabled     FltNewCfgTableInvert = 1
	FltNewCfgTableInvert_Disabled    FltNewCfgTableInvert = 2
	FltNewCfgTableInvert_Unsupported FltNewCfgTableInvert = 2147483647
)

type FltNewCfgTableClientProxy int32

const (
	FltNewCfgTableClientProxy_Enabled     FltNewCfgTableClientProxy = 1
	FltNewCfgTableClientProxy_Disabled    FltNewCfgTableClientProxy = 2
	FltNewCfgTableClientProxy_Unsupported FltNewCfgTableClientProxy = 2147483647
)

type FltNewCfgTableTcpAck int32

const (
	FltNewCfgTableTcpAck_Enabled     FltNewCfgTableTcpAck = 1
	FltNewCfgTableTcpAck_Disabled    FltNewCfgTableTcpAck = 2
	FltNewCfgTableTcpAck_Unsupported FltNewCfgTableTcpAck = 2147483647
)

type FltNewCfgTableFtpNatActive int32

const (
	FltNewCfgTableFtpNatActive_Enabled     FltNewCfgTableFtpNatActive = 1
	FltNewCfgTableFtpNatActive_Disabled    FltNewCfgTableFtpNatActive = 2
	FltNewCfgTableFtpNatActive_Unsupported FltNewCfgTableFtpNatActive = 2147483647
)

type FltNewCfgTableAclTcpUrg int32

const (
	FltNewCfgTableAclTcpUrg_Enable      FltNewCfgTableAclTcpUrg = 1
	FltNewCfgTableAclTcpUrg_Disable     FltNewCfgTableAclTcpUrg = 2
	FltNewCfgTableAclTcpUrg_Unsupported FltNewCfgTableAclTcpUrg = 2147483647
)

type FltNewCfgTableAclTcpAck int32

const (
	FltNewCfgTableAclTcpAck_Enable      FltNewCfgTableAclTcpAck = 1
	FltNewCfgTableAclTcpAck_Disable     FltNewCfgTableAclTcpAck = 2
	FltNewCfgTableAclTcpAck_Unsupported FltNewCfgTableAclTcpAck = 2147483647
)

type FltNewCfgTableAclTcpPsh int32

const (
	FltNewCfgTableAclTcpPsh_Enable      FltNewCfgTableAclTcpPsh = 1
	FltNewCfgTableAclTcpPsh_Disable     FltNewCfgTableAclTcpPsh = 2
	FltNewCfgTableAclTcpPsh_Unsupported FltNewCfgTableAclTcpPsh = 2147483647
)

type FltNewCfgTableAclTcpRst int32

const (
	FltNewCfgTableAclTcpRst_Enable      FltNewCfgTableAclTcpRst = 1
	FltNewCfgTableAclTcpRst_Disable     FltNewCfgTableAclTcpRst = 2
	FltNewCfgTableAclTcpRst_Unsupported FltNewCfgTableAclTcpRst = 2147483647
)

type FltNewCfgTableAclTcpSyn int32

const (
	FltNewCfgTableAclTcpSyn_Enable      FltNewCfgTableAclTcpSyn = 1
	FltNewCfgTableAclTcpSyn_Disable     FltNewCfgTableAclTcpSyn = 2
	FltNewCfgTableAclTcpSyn_Unsupported FltNewCfgTableAclTcpSyn = 2147483647
)

type FltNewCfgTableAclTcpFin int32

const (
	FltNewCfgTableAclTcpFin_Enable      FltNewCfgTableAclTcpFin = 1
	FltNewCfgTableAclTcpFin_Disable     FltNewCfgTableAclTcpFin = 2
	FltNewCfgTableAclTcpFin_Unsupported FltNewCfgTableAclTcpFin = 2147483647
)

type FltNewCfgTableAclIpOption int32

const (
	FltNewCfgTableAclIpOption_Enable      FltNewCfgTableAclIpOption = 1
	FltNewCfgTableAclIpOption_Disable     FltNewCfgTableAclIpOption = 2
	FltNewCfgTableAclIpOption_Unsupported FltNewCfgTableAclIpOption = 2147483647
)

type FltNewCfgTableFwlb int32

const (
	FltNewCfgTableFwlb_Enabled     FltNewCfgTableFwlb = 1
	FltNewCfgTableFwlb_Disabled    FltNewCfgTableFwlb = 2
	FltNewCfgTableFwlb_Unsupported FltNewCfgTableFwlb = 2147483647
)

type FltNewCfgTableLinklb int32

const (
	FltNewCfgTableLinklb_Enabled     FltNewCfgTableLinklb = 1
	FltNewCfgTableLinklb_Disabled    FltNewCfgTableLinklb = 2
	FltNewCfgTableLinklb_Unsupported FltNewCfgTableLinklb = 2147483647
)

type FltNewCfgTableWapRadiusSnoop int32

const (
	FltNewCfgTableWapRadiusSnoop_Enabled     FltNewCfgTableWapRadiusSnoop = 1
	FltNewCfgTableWapRadiusSnoop_Disabled    FltNewCfgTableWapRadiusSnoop = 2
	FltNewCfgTableWapRadiusSnoop_Unsupported FltNewCfgTableWapRadiusSnoop = 2147483647
)

type FltNewCfgTableSrcIpMac int32

const (
	FltNewCfgTableSrcIpMac_Ip          FltNewCfgTableSrcIpMac = 1
	FltNewCfgTableSrcIpMac_Mac         FltNewCfgTableSrcIpMac = 2
	FltNewCfgTableSrcIpMac_Unsupported FltNewCfgTableSrcIpMac = 2147483647
)

type FltNewCfgTableDstIpMac int32

const (
	FltNewCfgTableDstIpMac_Ip          FltNewCfgTableDstIpMac = 1
	FltNewCfgTableDstIpMac_Mac         FltNewCfgTableDstIpMac = 2
	FltNewCfgTableDstIpMac_Unsupported FltNewCfgTableDstIpMac = 2147483647
)

type FltNewCfgTableIdslbHash int32

const (
	FltNewCfgTableIdslbHash_Sip         FltNewCfgTableIdslbHash = 1
	FltNewCfgTableIdslbHash_Dip         FltNewCfgTableIdslbHash = 2
	FltNewCfgTableIdslbHash_Sipdip      FltNewCfgTableIdslbHash = 3
	FltNewCfgTableIdslbHash_Sipsport    FltNewCfgTableIdslbHash = 4
	FltNewCfgTableIdslbHash_Dipdport    FltNewCfgTableIdslbHash = 5
	FltNewCfgTableIdslbHash_All         FltNewCfgTableIdslbHash = 6
	FltNewCfgTableIdslbHash_Unsupported FltNewCfgTableIdslbHash = 2147483647
)

type FltNewCfgTableTcpRateLimit int32

const (
	FltNewCfgTableTcpRateLimit_Enabled     FltNewCfgTableTcpRateLimit = 1
	FltNewCfgTableTcpRateLimit_Disabled    FltNewCfgTableTcpRateLimit = 2
	FltNewCfgTableTcpRateLimit_Unsupported FltNewCfgTableTcpRateLimit = 2147483647
)

type FltNewCfgTableHash int32

const (
	FltNewCfgTableHash_Auto        FltNewCfgTableHash = 1
	FltNewCfgTableHash_Sip         FltNewCfgTableHash = 2
	FltNewCfgTableHash_Dip         FltNewCfgTableHash = 3
	FltNewCfgTableHash_Both        FltNewCfgTableHash = 4
	FltNewCfgTableHash_Sipsport    FltNewCfgTableHash = 5
	FltNewCfgTableHash_Dip32       FltNewCfgTableHash = 6
	FltNewCfgTableHash_Unsupported FltNewCfgTableHash = 2147483647
)

type FltNewCfgTableLayer7DenyState int32

const (
	FltNewCfgTableLayer7DenyState_Enabled     FltNewCfgTableLayer7DenyState = 1
	FltNewCfgTableLayer7DenyState_Disabled    FltNewCfgTableLayer7DenyState = 2
	FltNewCfgTableLayer7DenyState_Unsupported FltNewCfgTableLayer7DenyState = 2147483647
)

type FltNewCfgTableRadiusWapPersist int32

const (
	FltNewCfgTableRadiusWapPersist_Enabled     FltNewCfgTableRadiusWapPersist = 1
	FltNewCfgTableRadiusWapPersist_Disabled    FltNewCfgTableRadiusWapPersist = 2
	FltNewCfgTableRadiusWapPersist_Unsupported FltNewCfgTableRadiusWapPersist = 2147483647
)

type FltNewCfgTablePbind int32

const (
	FltNewCfgTablePbind_Enabled     FltNewCfgTablePbind = 1
	FltNewCfgTablePbind_Disabled    FltNewCfgTablePbind = 2
	FltNewCfgTablePbind_Unsupported FltNewCfgTablePbind = 2147483647
)

type FltNewCfgTablePatternMatch int32

const (
	FltNewCfgTablePatternMatch_Enabled     FltNewCfgTablePatternMatch = 1
	FltNewCfgTablePatternMatch_Disabled    FltNewCfgTablePatternMatch = 2
	FltNewCfgTablePatternMatch_Unsupported FltNewCfgTablePatternMatch = 2147483647
)

type FltNewCfgTableLayer7DenyMatchAll int32

const (
	FltNewCfgTableLayer7DenyMatchAll_Enabled     FltNewCfgTableLayer7DenyMatchAll = 1
	FltNewCfgTableLayer7DenyMatchAll_Disabled    FltNewCfgTableLayer7DenyMatchAll = 2
	FltNewCfgTableLayer7DenyMatchAll_Unsupported FltNewCfgTableLayer7DenyMatchAll = 2147483647
)

type FltNewCfgTableLayer7ParseAll int32

const (
	FltNewCfgTableLayer7ParseAll_Enabled     FltNewCfgTableLayer7ParseAll = 1
	FltNewCfgTableLayer7ParseAll_Disabled    FltNewCfgTableLayer7ParseAll = 2
	FltNewCfgTableLayer7ParseAll_Unsupported FltNewCfgTableLayer7ParseAll = 2147483647
)

type FltNewCfgTableSecurityParseAll int32

const (
	FltNewCfgTableSecurityParseAll_Enabled     FltNewCfgTableSecurityParseAll = 1
	FltNewCfgTableSecurityParseAll_Disabled    FltNewCfgTableSecurityParseAll = 2
	FltNewCfgTableSecurityParseAll_Unsupported FltNewCfgTableSecurityParseAll = 2147483647
)

type FltNewCfgTableCfg8021pBitsMatch int32

const (
	FltNewCfgTableCfg8021pBitsMatch_Enabled     FltNewCfgTableCfg8021pBitsMatch = 1
	FltNewCfgTableCfg8021pBitsMatch_Disabled    FltNewCfgTableCfg8021pBitsMatch = 2
	FltNewCfgTableCfg8021pBitsMatch_Unsupported FltNewCfgTableCfg8021pBitsMatch = 2147483647
)

type FltNewCfgTableEgressPip int32

const (
	FltNewCfgTableEgressPip_Enabled     FltNewCfgTableEgressPip = 1
	FltNewCfgTableEgressPip_Disabled    FltNewCfgTableEgressPip = 2
	FltNewCfgTableEgressPip_Unsupported FltNewCfgTableEgressPip = 2147483647
)

type FltNewCfgTableDbind int32

const (
	FltNewCfgTableDbind_Enabled     FltNewCfgTableDbind = 1
	FltNewCfgTableDbind_Disabled    FltNewCfgTableDbind = 2
	FltNewCfgTableDbind_Forceproxy  FltNewCfgTableDbind = 3
	FltNewCfgTableDbind_Unsupported FltNewCfgTableDbind = 2147483647
)

type FltNewCfgTableReverse int32

const (
	FltNewCfgTableReverse_Enabled     FltNewCfgTableReverse = 1
	FltNewCfgTableReverse_Disabled    FltNewCfgTableReverse = 2
	FltNewCfgTableReverse_Unsupported FltNewCfgTableReverse = 2147483647
)

type FltNewCfgTableParseChn int32

const (
	FltNewCfgTableParseChn_Enabled     FltNewCfgTableParseChn = 1
	FltNewCfgTableParseChn_Disabled    FltNewCfgTableParseChn = 2
	FltNewCfgTableParseChn_Unsupported FltNewCfgTableParseChn = 2147483647
)

type FltNewCfgTableSipParsing int32

const (
	FltNewCfgTableSipParsing_Enabled     FltNewCfgTableSipParsing = 1
	FltNewCfgTableSipParsing_Disabled    FltNewCfgTableSipParsing = 2
	FltNewCfgTableSipParsing_Unsupported FltNewCfgTableSipParsing = 2147483647
)

type FltNewCfgTableSessionMirror int32

const (
	FltNewCfgTableSessionMirror_Enabled     FltNewCfgTableSessionMirror = 1
	FltNewCfgTableSessionMirror_Disabled    FltNewCfgTableSessionMirror = 2
	FltNewCfgTableSessionMirror_Unsupported FltNewCfgTableSessionMirror = 2147483647
)

type FltNewCfgTableIpVer int32

const (
	FltNewCfgTableIpVer_Ipv4        FltNewCfgTableIpVer = 1
	FltNewCfgTableIpVer_Ipv6        FltNewCfgTableIpVer = 2
	FltNewCfgTableIpVer_Unsupported FltNewCfgTableIpVer = 2147483647
)

type FltNewCfgTableHdrHash int32

const (
	FltNewCfgTableHdrHash_None        FltNewCfgTableHdrHash = 1
	FltNewCfgTableHdrHash_Headerhash  FltNewCfgTableHdrHash = 2
	FltNewCfgTableHdrHash_Unsupported FltNewCfgTableHdrHash = 2147483647
)

type FltNewCfgTableL3Filter int32

const (
	FltNewCfgTableL3Filter_Enabled     FltNewCfgTableL3Filter = 1
	FltNewCfgTableL3Filter_Disabled    FltNewCfgTableL3Filter = 2
	FltNewCfgTableL3Filter_Unsupported FltNewCfgTableL3Filter = 2147483647
)

type FltNewCfgTableL7SipFilt int32

const (
	FltNewCfgTableL7SipFilt_Enabled     FltNewCfgTableL7SipFilt = 1
	FltNewCfgTableL7SipFilt_Disabled    FltNewCfgTableL7SipFilt = 2
	FltNewCfgTableL7SipFilt_Unsupported FltNewCfgTableL7SipFilt = 2147483647
)

type FltNewCfgTableNbind int32

const (
	FltNewCfgTableNbind_Enabled     FltNewCfgTableNbind = 1
	FltNewCfgTableNbind_Disabled    FltNewCfgTableNbind = 2
	FltNewCfgTableNbind_Unsupported FltNewCfgTableNbind = 2147483647
)

type FltNewCfgTableLayer7InvertAction int32

const (
	FltNewCfgTableLayer7InvertAction_Enabled     FltNewCfgTableLayer7InvertAction = 1
	FltNewCfgTableLayer7InvertAction_Disabled    FltNewCfgTableLayer7InvertAction = 2
	FltNewCfgTableLayer7InvertAction_Unsupported FltNewCfgTableLayer7InvertAction = 2147483647
)

type FltNewCfgTableRtsrcmac int32

const (
	FltNewCfgTableRtsrcmac_Enabled     FltNewCfgTableRtsrcmac = 1
	FltNewCfgTableRtsrcmac_Disabled    FltNewCfgTableRtsrcmac = 2
	FltNewCfgTableRtsrcmac_Unsupported FltNewCfgTableRtsrcmac = 2147483647
)

type FltNewCfgTableRtproxy int32

const (
	FltNewCfgTableRtproxy_Enabled     FltNewCfgTableRtproxy = 1
	FltNewCfgTableRtproxy_Disabled    FltNewCfgTableRtproxy = 2
	FltNewCfgTableRtproxy_Unsupported FltNewCfgTableRtproxy = 2147483647
)

type FltNewCfgTableSesslog int32

const (
	FltNewCfgTableSesslog_Enabled     FltNewCfgTableSesslog = 1
	FltNewCfgTableSesslog_Disabled    FltNewCfgTableSesslog = 2
	FltNewCfgTableSesslog_Unsupported FltNewCfgTableSesslog = 2147483647
)

type FltNewCfgTableApplicType int32

const (
	FltNewCfgTableApplicType_None        FltNewCfgTableApplicType = 1
	FltNewCfgTableApplicType_Basic       FltNewCfgTableApplicType = 2
	FltNewCfgTableApplicType_Http        FltNewCfgTableApplicType = 3
	FltNewCfgTableApplicType_Sip         FltNewCfgTableApplicType = 4
	FltNewCfgTableApplicType_Dns         FltNewCfgTableApplicType = 5
	FltNewCfgTableApplicType_Smtp        FltNewCfgTableApplicType = 6
	FltNewCfgTableApplicType_Pop3        FltNewCfgTableApplicType = 7
	FltNewCfgTableApplicType_Imap        FltNewCfgTableApplicType = 8
	FltNewCfgTableApplicType_Ftp         FltNewCfgTableApplicType = 9
	FltNewCfgTableApplicType_Unsupported FltNewCfgTableApplicType = 2147483647
)

type FltNewCfgTableRtsport int32

const (
	FltNewCfgTableRtsport_Mod5060     FltNewCfgTableRtsport = 1
	FltNewCfgTableRtsport_Preserve    FltNewCfgTableRtsport = 2
	FltNewCfgTableRtsport_Disabled    FltNewCfgTableRtsport = 3
	FltNewCfgTableRtsport_Unsupported FltNewCfgTableRtsport = 2147483647
)

type FltNewCfgTableSrcAddrType int32

const (
	FltNewCfgTableSrcAddrType_IpAddress   FltNewCfgTableSrcAddrType = 1
	FltNewCfgTableSrcAddrType_Network     FltNewCfgTableSrcAddrType = 2
	FltNewCfgTableSrcAddrType_Unsupported FltNewCfgTableSrcAddrType = 2147483647
)

type FltNewCfgTableDstAddrType int32

const (
	FltNewCfgTableDstAddrType_IpAddress   FltNewCfgTableDstAddrType = 1
	FltNewCfgTableDstAddrType_Network     FltNewCfgTableDstAddrType = 2
	FltNewCfgTableDstAddrType_Unsupported FltNewCfgTableDstAddrType = 2147483647
)

type FltNewCfgTableUdpAge int32

const (
	FltNewCfgTableUdpAge_Onlydns     FltNewCfgTableUdpAge = 1
	FltNewCfgTableUdpAge_Disabled    FltNewCfgTableUdpAge = 2
	FltNewCfgTableUdpAge_Enabled     FltNewCfgTableUdpAge = 3
	FltNewCfgTableUdpAge_Unsupported FltNewCfgTableUdpAge = 2147483647
)

type FltNewCfgTableSslInspectionEna int32

const (
	FltNewCfgTableSslInspectionEna_Enabled     FltNewCfgTableSslInspectionEna = 1
	FltNewCfgTableSslInspectionEna_Disabled    FltNewCfgTableSslInspectionEna = 2
	FltNewCfgTableSslInspectionEna_Unsupported FltNewCfgTableSslInspectionEna = 2147483647
)

type FltNewCfgTableSrvCertGroup int32

const (
	FltNewCfgTableSrvCertGroup_Group       FltNewCfgTableSrvCertGroup = 1
	FltNewCfgTableSrvCertGroup_Cert        FltNewCfgTableSrvCertGroup = 2
	FltNewCfgTableSrvCertGroup_None        FltNewCfgTableSrvCertGroup = 3
	FltNewCfgTableSrvCertGroup_Unsupported FltNewCfgTableSrvCertGroup = 2147483647
)

type FltNewCfgTableMatchDev int32

const (
	FltNewCfgTableMatchDev_Allexclif   FltNewCfgTableMatchDev = 1
	FltNewCfgTableMatchDev_All         FltNewCfgTableMatchDev = 2
	FltNewCfgTableMatchDev_None        FltNewCfgTableMatchDev = 3
	FltNewCfgTableMatchDev_Unsupported FltNewCfgTableMatchDev = 2147483647
)

type FltNewCfgTableSslL7Action int32

const (
	FltNewCfgTableSslL7Action_None        FltNewCfgTableSslL7Action = 1
	FltNewCfgTableSslL7Action_Bypass      FltNewCfgTableSslL7Action = 2
	FltNewCfgTableSslL7Action_Inspect     FltNewCfgTableSslL7Action = 3
	FltNewCfgTableSslL7Action_Unsupported FltNewCfgTableSslL7Action = 2147483647
)

type FltNewCfgTableFallback int32

const (
	FltNewCfgTableFallback_Allow       FltNewCfgTableFallback = 1
	FltNewCfgTableFallback_Deny        FltNewCfgTableFallback = 2
	FltNewCfgTableFallback_Goto        FltNewCfgTableFallback = 5
	FltNewCfgTableFallback_ContFlow    FltNewCfgTableFallback = 6
	FltNewCfgTableFallback_Unsupported FltNewCfgTableFallback = 2147483647
)

type FltNewCfgTableVpnFlood int32

const (
	FltNewCfgTableVpnFlood_Enabled     FltNewCfgTableVpnFlood = 1
	FltNewCfgTableVpnFlood_Disabled    FltNewCfgTableVpnFlood = 2
	FltNewCfgTableVpnFlood_Unsupported FltNewCfgTableVpnFlood = 2147483647
)

type FltNewCfgTableMacToMe int32

const (
	FltNewCfgTableMacToMe_Enabled     FltNewCfgTableMacToMe = 1
	FltNewCfgTableMacToMe_Disabled    FltNewCfgTableMacToMe = 2
	FltNewCfgTableMacToMe_Unsupported FltNewCfgTableMacToMe = 2147483647
)

type FltNewCfgTableAwinflow int32

const (
	FltNewCfgTableAwinflow_After       FltNewCfgTableAwinflow = 1
	FltNewCfgTableAwinflow_Before      FltNewCfgTableAwinflow = 2
	FltNewCfgTableAwinflow_Unsupported FltNewCfgTableAwinflow = 2147483647
)

type FltNewCfgTableJsinject int32

const (
	FltNewCfgTableJsinject_Enabled     FltNewCfgTableJsinject = 1
	FltNewCfgTableJsinject_Disabled    FltNewCfgTableJsinject = 2
	FltNewCfgTableJsinject_Unsupported FltNewCfgTableJsinject = 2147483647
)

type FltNewCfgTableParams struct {
	// The filtering table index.
	Indx int32 `json:"Indx,omitempty"`
	// The source IP address to be filtered.
	SrcIp string `json:"SrcIp,omitempty"`
	// The source IP sub-net mask for filtering.
	SrcIpMask string `json:"SrcIpMask,omitempty"`
	// The destination IP address to be filtered.
	DstIp string `json:"DstIp,omitempty"`
	// The destination IP sub-net mask for filtering.
	DstIpMask string `json:"DstIpMask,omitempty"`
	// The protocol to be filtered.
	Protocol uint32 `json:"Protocol,omitempty"`
	// The higher source TCP/UDP port number to be filtered. It applies
	// only when protocol defined in fltNewCfgProtocol is UDP or TCP.
	// '0' means no filtering.
	RangeHighSrcPort uint64 `json:"RangeHighSrcPort,omitempty"`
	// The lower source TCP/UDP port number to be filtered. It applies
	// only when protocol defined in fltNewCfgProtocol is UDP or TCP.
	// '0' means no filtering.
	RangeLowSrcPort uint64 `json:"RangeLowSrcPort,omitempty"`
	// The lower destination TCP/UDP port number to be filtered.
	// It applies only when protocol defined in fltNewCfgProtocol
	// 	 is UDP or TCP.  '0' means no filtering.
	RangeLowDstPort uint64 `json:"RangeLowDstPort,omitempty"`
	// The higher destination TCP/UDP port number to be filtered.
	// It applies only when protocol defined in fltNewCfgProtocol
	// is UDP or TCP.  '0' means no filtering.
	RangeHighDstPort uint64 `json:"RangeHighDstPort,omitempty"`
	// The action for the filtering rule.
	Action FltNewCfgTableAction `json:"Action,omitempty"`
	// The real server port number used for redirection.
	RedirPort uint64 `json:"RedirPort,omitempty"`
	// The real server group to be redirected to.
	RedirGroup int32 `json:"RedirGroup,omitempty"`
	// Enable or disable logging.
	Log FltNewCfgTableLog `json:"Log,omitempty"`
	// The state of this filtering rule.
	State FltNewCfgTableState `json:"State,omitempty"`
	// When set to the value of 2 (delete), the entire row is deleted.
	// When read, other(1) is returned. Setting the value to anything
	// other than 2(delete) has no effect on the state of the row.
	Delete FltNewCfgTableDelete `json:"Delete,omitempty"`
	// The selection of destination or source or multicast for
	// network address translation.
	Nat FltNewCfgTableNat `json:"Nat,omitempty"`
	// Enable or disable caching sessions that match filter.
	Cache FltNewCfgTableCache `json:"Cache,omitempty"`
	// Turn the invert logic on or off for the filter entry.
	Invert FltNewCfgTableInvert `json:"Invert,omitempty"`
	// Enable or disable client proxy.
	ClientProxy FltNewCfgTableClientProxy `json:"ClientProxy,omitempty"`
	// Enable or disable filtering on matching TCP ACK and RST flag.
	TcpAck FltNewCfgTableTcpAck `json:"TcpAck,omitempty"`
	// The source MAC address to be filtered.
	SrcMac string `json:"SrcMac,omitempty"`
	// The Destination MAC address to be filtered.
	DstMac string `json:"DstMac,omitempty"`
	// Enable or disable FTP NAT for active ftp.
	FtpNatActive FltNewCfgTableFtpNatActive `json:"FtpNatActive,omitempty"`
	// Enable or disable TCP URG packet.
	AclTcpUrg FltNewCfgTableAclTcpUrg `json:"AclTcpUrg,omitempty"`
	// Enable or disable TCP ACK packet.
	AclTcpAck FltNewCfgTableAclTcpAck `json:"AclTcpAck,omitempty"`
	// Enable or disable TCP PSH packet.
	AclTcpPsh FltNewCfgTableAclTcpPsh `json:"AclTcpPsh,omitempty"`
	// Enable or disable TCP RST packet.
	AclTcpRst FltNewCfgTableAclTcpRst `json:"AclTcpRst,omitempty"`
	// Enable or disable TCP SYN packet.
	AclTcpSyn FltNewCfgTableAclTcpSyn `json:"AclTcpSyn,omitempty"`
	// Enable or disable TCP FIN packet.
	AclTcpFin FltNewCfgTableAclTcpFin `json:"AclTcpFin,omitempty"`
	// ICMP type to be filtered. A ICMP type of 255 indicates 'any'
	AclIcmp uint32 `json:"AclIcmp,omitempty"`
	// Enable or disable IP option matching.
	AclIpOption FltNewCfgTableAclIpOption `json:"AclIpOption,omitempty"`
	// Filt default BW contract number.
	BwmContract int32 `json:"BwmContract,omitempty"`
	// The IP TOS value to be filtered.
	AclIpTos uint32 `json:"AclIpTos,omitempty"`
	// The IP TOS mask for filtering.
	AclIpTosMask uint32 `json:"AclIpTosMask,omitempty"`
	// The new IP TOS value to over-write when filtering fired.
	AclIpTosNew uint32 `json:"AclIpTosNew,omitempty"`
	// Enable or disable filtering on firewall redirect hash method.
	Fwlb FltNewCfgTableFwlb `json:"Fwlb,omitempty"`
	// The NAT session timeout. The timeout value should be an even
	// number between 4 and 32768.
	NatTimeout uint32 `json:"NatTimeout,omitempty"`
	// Enable or disable WAN link load balancing.
	Linklb FltNewCfgTableLinklb `json:"Linklb,omitempty"`
	// Enable or disable WAP RADIUS snooping.
	WapRadiusSnoop FltNewCfgTableWapRadiusSnoop `json:"WapRadiusSnoop,omitempty"`
	// Set a flag indicating whether filtering should be based on the
	// source IP address or the source MAC address field.
	SrcIpMac FltNewCfgTableSrcIpMac `json:"SrcIpMac,omitempty"`
	// Set a flag indicating whether filtering should be based on the
	// destination IP address or the destination MAC address field.
	DstIpMac FltNewCfgTableDstIpMac `json:"DstIpMac,omitempty"`
	// Set hash parameter for intrusion detection server load balancing.
	IdslbHash FltNewCfgTableIdslbHash `json:"IdslbHash,omitempty"`
	// Set the VLAN assoicated with the filter.
	Vlan uint32 `json:"Vlan,omitempty"`
	// The name of the filter.
	Name string `json:"Name,omitempty"`
	// Enable or disable protocol rate limiting.
	TcpRateLimit FltNewCfgTableTcpRateLimit `json:"TcpRateLimit,omitempty"`
	// Set maximum connections (number of connections in units of 10)
	// for TCP connection rate limiting. In the case of ICMP and UDP,
	// 	 this is the maximum packets (number of packets in units of 10).
	TcpRateMaxConn uint32 `json:"TcpRateMaxConn,omitempty"`
	// Set hash parameter for the filter.
	Hash FltNewCfgTableHash `json:"Hash,omitempty"`
	// Enable or disable layer 7 deny filtering.
	Layer7DenyState FltNewCfgTableLayer7DenyState `json:"Layer7DenyState,omitempty"`
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
	RadiusWapPersist FltNewCfgTableRadiusWapPersist `json:"RadiusWapPersist,omitempty"`
	// Enable or disable filter persistent binding.
	Pbind FltNewCfgTablePbind `json:"Pbind,omitempty"`
	// The time window for protocol rate limiting (in seconds).
	TimeWindow uint64 `json:"TimeWindow,omitempty"`
	// The hold down duration for protocol rate limiting (in minutes).
	HoldDuration uint64 `json:"HoldDuration,omitempty"`
	// Enable or disable binary pattern matching.
	PatternMatch FltNewCfgTablePatternMatch `json:"PatternMatch,omitempty"`
	// Enable or disable match-all criteria for L7 deny string matching.
	Layer7DenyMatchAll FltNewCfgTableLayer7DenyMatchAll `json:"Layer7DenyMatchAll,omitempty"`
	// The client proxy IP address for NAT and REDIR filter.
	ProxyIp string `json:"ProxyIp,omitempty"`
	// Enable or disable layer 7 lookup (parsing) of all packets.
	Layer7ParseAll FltNewCfgTableLayer7ParseAll `json:"Layer7ParseAll,omitempty"`
	// Enable or disable pattern string lookup (parsing) of all packets.
	SecurityParseAll FltNewCfgTableSecurityParseAll `json:"SecurityParseAll,omitempty"`
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
	Cfg8021pBitsValue uint64 `json:"Cfg8021pBitsValue,omitempty"`
	// Enable or disable matching on 802.1p bits in the packets.
	Cfg8021pBitsMatch FltNewCfgTableCfg8021pBitsMatch `json:"Cfg8021pBitsMatch,omitempty"`
	// Set the IP maximum packet length in bytes. A value can be either 0
	// which indicates 'any' length or between 64 and 65535.
	AclIpLength uint64 `json:"AclIpLength,omitempty"`
	// The real server group for IDS load balancing. A value of 0 indicates
	// 	 'none'.
	IdsGroup int32 `json:"IdsGroup,omitempty"`
	// Enable or disable pip selection based on egress port/vlan.
	EgressPip FltNewCfgTableEgressPip `json:"EgressPip,omitempty"`
	// Enable or disable filter delayed binding.
	Dbind FltNewCfgTableDbind `json:"Dbind,omitempty"`
	// Filt reverse session BWM contract number.
	RevBwmContract int32 `json:"RevBwmContract,omitempty"`
	// Enable or disable creating session for reverse side traffic.
	Reverse FltNewCfgTableReverse `json:"Reverse,omitempty"`
	// Enable or disable chained pgroup match criteria for l7 filtering.
	ParseChn FltNewCfgTableParseChn `json:"ParseChn,omitempty"`
	// BWM contract for SIP RTP traffic.
	RtpBwmContract int32 `json:"RtpBwmContract,omitempty"`
	// Enable or disable SIP NAT.
	SipParsing FltNewCfgTableSipParsing `json:"SipParsing,omitempty"`
	// Enable or disable session mirroring.
	SessionMirror FltNewCfgTableSessionMirror `json:"SessionMirror,omitempty"`
	// The type of IP address.
	IpVer FltNewCfgTableIpVer `json:"IpVer,omitempty"`
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
	HdrHash FltNewCfgTableHdrHash `json:"HdrHash,omitempty"`
	// The header name of the filter. Headerhash should be
	// enabled before header name is configured.
	HdrName string `json:"HdrName,omitempty"`
	// The header hash length of the filter. Headerhash should be
	// enabled before hash length is configured.
	HdrHashLen uint32 `json:"HdrHashLen,omitempty"`
	// The L3 filter processing state for this filter.
	L3Filter FltNewCfgTableL3Filter `json:"L3Filter,omitempty"`
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
	L7SipFilt FltNewCfgTableL7SipFilt `json:"L7SipFilt,omitempty"`
	// Multicast VLAN.
	NatMcastVlan uint32 `json:"NatMcastVlan,omitempty"`
	// Enable or disable subnet binding for redirection.
	Nbind FltNewCfgTableNbind `json:"Nbind,omitempty"`
	// Enable or disable invert action for layer 7 string matching.
	Layer7InvertAction FltNewCfgTableLayer7InvertAction `json:"Layer7InvertAction,omitempty"`
	// The source network class ID to be filtered.
	SrcClassId string `json:"SrcClassId,omitempty"`
	// The destination network class ID to be filtered.
	DstClassId string `json:"DstClassId,omitempty"`
	// Enable or disable return to source mac addr.
	Rtsrcmac FltNewCfgTableRtsrcmac `json:"Rtsrcmac,omitempty"`
	// Enable or disable redirect to proxy server.
	Rtproxy FltNewCfgTableRtproxy `json:"Rtproxy,omitempty"`
	// Enable or disable session logging.
	Sesslog FltNewCfgTableSesslog `json:"Sesslog,omitempty"`
	// The cntclass of the filter.
	Cntclass string `json:"Cntclass,omitempty"`
	// Application type for this filter.
	ApplicType FltNewCfgTableApplicType `json:"ApplicType,omitempty"`
	// Source port modification for SIP.
	// the mod5060 indicates modification of source port to 5060,
	// preserve indicates to retain the source port.
	Rtsport FltNewCfgTableRtsport `json:"Rtsport,omitempty"`
	// Source address type from address and network for NAT filter.
	SrcAddrType FltNewCfgTableSrcAddrType `json:"SrcAddrType,omitempty"`
	// Destination address type from address and network for NAT filter.
	DstAddrType FltNewCfgTableDstAddrType `json:"DstAddrType,omitempty"`
	// The class of service of the filter.
	// 	Empty string - ignore cos
	// 	'Any' string - ignore string use only the source IP
	// 	Specific string - match only if the string matches
	CosStr string `json:"CosStr,omitempty"`
	// Udp Aging type for this filter.
	UdpAge FltNewCfgTableUdpAge `json:"UdpAge,omitempty"`
	// frontend TCP optimization policy
	FeTcpPolId string `json:"FeTcpPolId,omitempty"`
	// backend TCP optimization policy
	BeTcpPolId string `json:"BeTcpPolId,omitempty"`
	// Compression policy name associated with this filter.
	Comppol string `json:"Comppol,omitempty"`
	// Enable or disable SSL session inspection.
	SslInspectionEna FltNewCfgTableSslInspectionEna `json:"SslInspectionEna,omitempty"`
	// Set Certificate group or cert.
	SrvCertGroup FltNewCfgTableSrvCertGroup `json:"SrvCertGroup,omitempty"`
	// Server certificate / group name associated with this filter.
	SrvCert string `json:"SrvCert,omitempty"`
	// SSL policy name associated with this filter.
	SslPolicy string `json:"SslPolicy,omitempty"`
	// Match device IP addresses in this filter:
	// allexclif : match all device IPs, but exclude interface IPs
	// all       : match all device IPs
	// none      : match no device IPs
	MatchDev FltNewCfgTableMatchDev `json:"MatchDev,omitempty"`
	// Set SSL inspection action on L7 match.
	SslL7Action FltNewCfgTableSslL7Action `json:"SslL7Action,omitempty"`
	// Fallback action (group down).
	Fallback FltNewCfgTableFallback `json:"Fallback,omitempty"`
	// The filter ID for GOTO fallback action in the new config.
	Fbgoto int32 `json:"Fbgoto,omitempty"`
	// Enable or disable Two Way VPN LB.
	VpnFlood FltNewCfgTableVpnFlood `json:"VpnFlood,omitempty"`
	// The url filter policy of the filter.
	UrlFltPol string `json:"UrlFltPol,omitempty"`
	// The port list in the filter. The ports are presented in bitmap format.
	// 	 in receiving order:
	// 	     OCTET 1  OCTET 2  .....
	// xxxxxxxx xxxxxxxx .....
	// ||    || |_ port 8
	// ||    ||
	// ||    ||___ port 7
	// ||    |____ port 6
	// ||      .    .   .
	// ||_________ port 1
	// |__________ reserved
	// where x : 1 - The represented port belongs to the filter
	// 		   0 - The represented port does not belong to the filter
	Ports string `json:"Ports,omitempty"`
	// The port to be added to the specified filter.  A '0' value is
	// returned when read.
	AddPort int32 `json:"AddPort,omitempty"`
	// The port to be removed from the specified filter.  A '0'
	// value is returned when read.
	RemovePort int32 `json:"RemovePort,omitempty"`
	// The ingress port for fallback action continueFlow in the new config.
	Fbport int32 `json:"Fbport,omitempty"`
	// Group Index of filters on same tunnel.
	SetIndex int32 `json:"SetIndex,omitempty"`
	// Enable/disable matching traffic to device mac addresses.
	MacToMe FltNewCfgTableMacToMe `json:"MacToMe,omitempty"`
	// The ingress VLAN ID for fallback action continueFlow in the new config.
	Fbvlan uint32 `json:"Fbvlan,omitempty"`
	// URL filter classification mode.
	Urlfmode uint32 `json:"Urlfmode,omitempty"`
	// Set Bot Manager Policy.
	Botpol string `json:"Botpol,omitempty"`
	// Set if AW processing comes before or after Alteon HTTP parsing.
	Awinflow FltNewCfgTableAwinflow `json:"Awinflow,omitempty"`
	// Server GmSSL encryption certificate / group name associated with this filter.
	SrvCertEnc string `json:"SrvCertEnc,omitempty"`
	// Server GmSSL signing certificate / group name associated with this filter.
	SrvCertSign string `json:"SrvCertSign,omitempty"`
	// Set Sideband policy.
	SidebandID string `json:"SidebandID,omitempty"`
	// Set secure path Policy.
	SecurePathPolicy string `json:"SecurePathPolicy,omitempty"`
	// Set JS inject mode .
	Jsinject FltNewCfgTableJsinject `json:"Jsinject,omitempty"`
}

func (p FltNewCfgTableParams) iMABean() {}
