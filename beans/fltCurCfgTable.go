package beans

import (
	"fmt"
	"reflect"
)

// FltCurCfgTable The filtering table in the current configuration block.
// Note:This mib is not supported for VX instance of Virtualization.
type FltCurCfgTable struct {
	// The filtering table index.
	FltCurCfgIndx int32
	Params        *FltCurCfgTableParams
}

func NewFltCurCfgTableList() *FltCurCfgTable {
	return &FltCurCfgTable{}
}

func NewFltCurCfgTable(
	fltCurCfgIndx int32,
	params *FltCurCfgTableParams,
) *FltCurCfgTable {
	return &FltCurCfgTable{
		FltCurCfgIndx: fltCurCfgIndx,
		Params:        params,
	}
}

func (c *FltCurCfgTable) Name() string {
	return "FltCurCfgTable"
}

func (c *FltCurCfgTable) GetParams() BeanType {
	return c.Params
}

func (c *FltCurCfgTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *FltCurCfgTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.FltCurCfgIndx).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.FltCurCfgIndx)
}

type FltCurCfgTableAction int32

const (
	FltCurCfgTableAction_Allow       FltCurCfgTableAction = 1
	FltCurCfgTableAction_Deny        FltCurCfgTableAction = 2
	FltCurCfgTableAction_Redirect    FltCurCfgTableAction = 3
	FltCurCfgTableAction_Nat         FltCurCfgTableAction = 4
	FltCurCfgTableAction_Goto        FltCurCfgTableAction = 5
	FltCurCfgTableAction_OutboundLlb FltCurCfgTableAction = 6
	FltCurCfgTableAction_Monitor     FltCurCfgTableAction = 7
	FltCurCfgTableAction_Unsupported FltCurCfgTableAction = 2147483647
)

type FltCurCfgTableLog int32

const (
	FltCurCfgTableLog_Enabled     FltCurCfgTableLog = 1
	FltCurCfgTableLog_Disabled    FltCurCfgTableLog = 2
	FltCurCfgTableLog_Unsupported FltCurCfgTableLog = 2147483647
)

type FltCurCfgTableState int32

const (
	FltCurCfgTableState_Enabled     FltCurCfgTableState = 1
	FltCurCfgTableState_Disabled    FltCurCfgTableState = 2
	FltCurCfgTableState_Unsupported FltCurCfgTableState = 2147483647
)

type FltCurCfgTableNat int32

const (
	FltCurCfgTableNat_DestinationAddress FltCurCfgTableNat = 1
	FltCurCfgTableNat_SourceAddress      FltCurCfgTableNat = 2
	FltCurCfgTableNat_MulticastAddress   FltCurCfgTableNat = 3
	FltCurCfgTableNat_Unsupported        FltCurCfgTableNat = 2147483647
)

type FltCurCfgTableCache int32

const (
	FltCurCfgTableCache_Enabled     FltCurCfgTableCache = 1
	FltCurCfgTableCache_Disabled    FltCurCfgTableCache = 2
	FltCurCfgTableCache_Unsupported FltCurCfgTableCache = 2147483647
)

type FltCurCfgTableInvert int32

const (
	FltCurCfgTableInvert_Enabled     FltCurCfgTableInvert = 1
	FltCurCfgTableInvert_Disabled    FltCurCfgTableInvert = 2
	FltCurCfgTableInvert_Unsupported FltCurCfgTableInvert = 2147483647
)

type FltCurCfgTableClientProxy int32

const (
	FltCurCfgTableClientProxy_Enabled     FltCurCfgTableClientProxy = 1
	FltCurCfgTableClientProxy_Disabled    FltCurCfgTableClientProxy = 2
	FltCurCfgTableClientProxy_Unsupported FltCurCfgTableClientProxy = 2147483647
)

type FltCurCfgTableTcpAck int32

const (
	FltCurCfgTableTcpAck_Enabled     FltCurCfgTableTcpAck = 1
	FltCurCfgTableTcpAck_Disabled    FltCurCfgTableTcpAck = 2
	FltCurCfgTableTcpAck_Unsupported FltCurCfgTableTcpAck = 2147483647
)

type FltCurCfgTableFtpNatActive int32

const (
	FltCurCfgTableFtpNatActive_Enabled     FltCurCfgTableFtpNatActive = 1
	FltCurCfgTableFtpNatActive_Disabled    FltCurCfgTableFtpNatActive = 2
	FltCurCfgTableFtpNatActive_Unsupported FltCurCfgTableFtpNatActive = 2147483647
)

type FltCurCfgTableAclTcpUrg int32

const (
	FltCurCfgTableAclTcpUrg_Enable      FltCurCfgTableAclTcpUrg = 1
	FltCurCfgTableAclTcpUrg_Disable     FltCurCfgTableAclTcpUrg = 2
	FltCurCfgTableAclTcpUrg_Unsupported FltCurCfgTableAclTcpUrg = 2147483647
)

type FltCurCfgTableAclTcpAck int32

const (
	FltCurCfgTableAclTcpAck_Enable      FltCurCfgTableAclTcpAck = 1
	FltCurCfgTableAclTcpAck_Disable     FltCurCfgTableAclTcpAck = 2
	FltCurCfgTableAclTcpAck_Unsupported FltCurCfgTableAclTcpAck = 2147483647
)

type FltCurCfgTableAclTcpPsh int32

const (
	FltCurCfgTableAclTcpPsh_Enable      FltCurCfgTableAclTcpPsh = 1
	FltCurCfgTableAclTcpPsh_Disable     FltCurCfgTableAclTcpPsh = 2
	FltCurCfgTableAclTcpPsh_Unsupported FltCurCfgTableAclTcpPsh = 2147483647
)

type FltCurCfgTableAclTcpRst int32

const (
	FltCurCfgTableAclTcpRst_Enable      FltCurCfgTableAclTcpRst = 1
	FltCurCfgTableAclTcpRst_Disable     FltCurCfgTableAclTcpRst = 2
	FltCurCfgTableAclTcpRst_Unsupported FltCurCfgTableAclTcpRst = 2147483647
)

type FltCurCfgTableAclTcpSyn int32

const (
	FltCurCfgTableAclTcpSyn_Enable      FltCurCfgTableAclTcpSyn = 1
	FltCurCfgTableAclTcpSyn_Disable     FltCurCfgTableAclTcpSyn = 2
	FltCurCfgTableAclTcpSyn_Unsupported FltCurCfgTableAclTcpSyn = 2147483647
)

type FltCurCfgTableAclTcpFin int32

const (
	FltCurCfgTableAclTcpFin_Enable      FltCurCfgTableAclTcpFin = 1
	FltCurCfgTableAclTcpFin_Disable     FltCurCfgTableAclTcpFin = 2
	FltCurCfgTableAclTcpFin_Unsupported FltCurCfgTableAclTcpFin = 2147483647
)

type FltCurCfgTableAclIpOption int32

const (
	FltCurCfgTableAclIpOption_Enable      FltCurCfgTableAclIpOption = 1
	FltCurCfgTableAclIpOption_Disable     FltCurCfgTableAclIpOption = 2
	FltCurCfgTableAclIpOption_Unsupported FltCurCfgTableAclIpOption = 2147483647
)

type FltCurCfgTableFwlb int32

const (
	FltCurCfgTableFwlb_Enabled     FltCurCfgTableFwlb = 1
	FltCurCfgTableFwlb_Disabled    FltCurCfgTableFwlb = 2
	FltCurCfgTableFwlb_Unsupported FltCurCfgTableFwlb = 2147483647
)

type FltCurCfgTableLinklb int32

const (
	FltCurCfgTableLinklb_Enabled     FltCurCfgTableLinklb = 1
	FltCurCfgTableLinklb_Disabled    FltCurCfgTableLinklb = 2
	FltCurCfgTableLinklb_Unsupported FltCurCfgTableLinklb = 2147483647
)

type FltCurCfgTableWapRadiusSnoop int32

const (
	FltCurCfgTableWapRadiusSnoop_Enabled     FltCurCfgTableWapRadiusSnoop = 1
	FltCurCfgTableWapRadiusSnoop_Disabled    FltCurCfgTableWapRadiusSnoop = 2
	FltCurCfgTableWapRadiusSnoop_Unsupported FltCurCfgTableWapRadiusSnoop = 2147483647
)

type FltCurCfgTableSrcIpMac int32

const (
	FltCurCfgTableSrcIpMac_Ip          FltCurCfgTableSrcIpMac = 1
	FltCurCfgTableSrcIpMac_Mac         FltCurCfgTableSrcIpMac = 2
	FltCurCfgTableSrcIpMac_Unsupported FltCurCfgTableSrcIpMac = 2147483647
)

type FltCurCfgTableDstIpMac int32

const (
	FltCurCfgTableDstIpMac_Ip          FltCurCfgTableDstIpMac = 1
	FltCurCfgTableDstIpMac_Mac         FltCurCfgTableDstIpMac = 2
	FltCurCfgTableDstIpMac_Unsupported FltCurCfgTableDstIpMac = 2147483647
)

type FltCurCfgTableIdslbHash int32

const (
	FltCurCfgTableIdslbHash_Sip         FltCurCfgTableIdslbHash = 1
	FltCurCfgTableIdslbHash_Dip         FltCurCfgTableIdslbHash = 2
	FltCurCfgTableIdslbHash_Sipdip      FltCurCfgTableIdslbHash = 3
	FltCurCfgTableIdslbHash_Sipsport    FltCurCfgTableIdslbHash = 4
	FltCurCfgTableIdslbHash_Dipdport    FltCurCfgTableIdslbHash = 5
	FltCurCfgTableIdslbHash_All         FltCurCfgTableIdslbHash = 6
	FltCurCfgTableIdslbHash_Unsupported FltCurCfgTableIdslbHash = 2147483647
)

type FltCurCfgTableTcpRateLimit int32

const (
	FltCurCfgTableTcpRateLimit_Enabled     FltCurCfgTableTcpRateLimit = 1
	FltCurCfgTableTcpRateLimit_Disabled    FltCurCfgTableTcpRateLimit = 2
	FltCurCfgTableTcpRateLimit_Unsupported FltCurCfgTableTcpRateLimit = 2147483647
)

type FltCurCfgTableHash int32

const (
	FltCurCfgTableHash_Auto        FltCurCfgTableHash = 1
	FltCurCfgTableHash_Sip         FltCurCfgTableHash = 2
	FltCurCfgTableHash_Dip         FltCurCfgTableHash = 3
	FltCurCfgTableHash_Both        FltCurCfgTableHash = 4
	FltCurCfgTableHash_Sipsport    FltCurCfgTableHash = 5
	FltCurCfgTableHash_Dip32       FltCurCfgTableHash = 6
	FltCurCfgTableHash_Unsupported FltCurCfgTableHash = 2147483647
)

type FltCurCfgTableLayer7DenyState int32

const (
	FltCurCfgTableLayer7DenyState_Enabled     FltCurCfgTableLayer7DenyState = 1
	FltCurCfgTableLayer7DenyState_Disabled    FltCurCfgTableLayer7DenyState = 2
	FltCurCfgTableLayer7DenyState_Unsupported FltCurCfgTableLayer7DenyState = 2147483647
)

type FltCurCfgTableRadiusWapPersist int32

const (
	FltCurCfgTableRadiusWapPersist_Enabled     FltCurCfgTableRadiusWapPersist = 1
	FltCurCfgTableRadiusWapPersist_Disabled    FltCurCfgTableRadiusWapPersist = 2
	FltCurCfgTableRadiusWapPersist_Unsupported FltCurCfgTableRadiusWapPersist = 2147483647
)

type FltCurCfgTablePbind int32

const (
	FltCurCfgTablePbind_Enabled     FltCurCfgTablePbind = 1
	FltCurCfgTablePbind_Disabled    FltCurCfgTablePbind = 2
	FltCurCfgTablePbind_Unsupported FltCurCfgTablePbind = 2147483647
)

type FltCurCfgTablePatternMatch int32

const (
	FltCurCfgTablePatternMatch_Enabled     FltCurCfgTablePatternMatch = 1
	FltCurCfgTablePatternMatch_Disabled    FltCurCfgTablePatternMatch = 2
	FltCurCfgTablePatternMatch_Unsupported FltCurCfgTablePatternMatch = 2147483647
)

type FltCurCfgTableLayer7DenyMatchAll int32

const (
	FltCurCfgTableLayer7DenyMatchAll_Enabled     FltCurCfgTableLayer7DenyMatchAll = 1
	FltCurCfgTableLayer7DenyMatchAll_Disabled    FltCurCfgTableLayer7DenyMatchAll = 2
	FltCurCfgTableLayer7DenyMatchAll_Unsupported FltCurCfgTableLayer7DenyMatchAll = 2147483647
)

type FltCurCfgTableLayer7ParseAll int32

const (
	FltCurCfgTableLayer7ParseAll_Enabled     FltCurCfgTableLayer7ParseAll = 1
	FltCurCfgTableLayer7ParseAll_Disabled    FltCurCfgTableLayer7ParseAll = 2
	FltCurCfgTableLayer7ParseAll_Unsupported FltCurCfgTableLayer7ParseAll = 2147483647
)

type FltCurCfgTableSecurityParseAll int32

const (
	FltCurCfgTableSecurityParseAll_Enabled     FltCurCfgTableSecurityParseAll = 1
	FltCurCfgTableSecurityParseAll_Disabled    FltCurCfgTableSecurityParseAll = 2
	FltCurCfgTableSecurityParseAll_Unsupported FltCurCfgTableSecurityParseAll = 2147483647
)

type FltCurCfgTableCfg8021pBitsMatch int32

const (
	FltCurCfgTableCfg8021pBitsMatch_Enabled     FltCurCfgTableCfg8021pBitsMatch = 1
	FltCurCfgTableCfg8021pBitsMatch_Disabled    FltCurCfgTableCfg8021pBitsMatch = 2
	FltCurCfgTableCfg8021pBitsMatch_Unsupported FltCurCfgTableCfg8021pBitsMatch = 2147483647
)

type FltCurCfgTableEgressPip int32

const (
	FltCurCfgTableEgressPip_Enabled     FltCurCfgTableEgressPip = 1
	FltCurCfgTableEgressPip_Disabled    FltCurCfgTableEgressPip = 2
	FltCurCfgTableEgressPip_Unsupported FltCurCfgTableEgressPip = 2147483647
)

type FltCurCfgTableDbind int32

const (
	FltCurCfgTableDbind_Enabled     FltCurCfgTableDbind = 1
	FltCurCfgTableDbind_Disabled    FltCurCfgTableDbind = 2
	FltCurCfgTableDbind_Forceproxy  FltCurCfgTableDbind = 3
	FltCurCfgTableDbind_Unsupported FltCurCfgTableDbind = 2147483647
)

type FltCurCfgTableReverse int32

const (
	FltCurCfgTableReverse_Enabled     FltCurCfgTableReverse = 1
	FltCurCfgTableReverse_Disabled    FltCurCfgTableReverse = 2
	FltCurCfgTableReverse_Unsupported FltCurCfgTableReverse = 2147483647
)

type FltCurCfgTableParseChn int32

const (
	FltCurCfgTableParseChn_Enabled     FltCurCfgTableParseChn = 1
	FltCurCfgTableParseChn_Disabled    FltCurCfgTableParseChn = 2
	FltCurCfgTableParseChn_Unsupported FltCurCfgTableParseChn = 2147483647
)

type FltCurCfgTableSipParsing int32

const (
	FltCurCfgTableSipParsing_Enabled     FltCurCfgTableSipParsing = 1
	FltCurCfgTableSipParsing_Disabled    FltCurCfgTableSipParsing = 2
	FltCurCfgTableSipParsing_Unsupported FltCurCfgTableSipParsing = 2147483647
)

type FltCurCfgTableSessionMirror int32

const (
	FltCurCfgTableSessionMirror_Enabled     FltCurCfgTableSessionMirror = 1
	FltCurCfgTableSessionMirror_Disabled    FltCurCfgTableSessionMirror = 2
	FltCurCfgTableSessionMirror_Unsupported FltCurCfgTableSessionMirror = 2147483647
)

type FltCurCfgTableIpVer int32

const (
	FltCurCfgTableIpVer_Ipv4        FltCurCfgTableIpVer = 1
	FltCurCfgTableIpVer_Ipv6        FltCurCfgTableIpVer = 2
	FltCurCfgTableIpVer_Unsupported FltCurCfgTableIpVer = 2147483647
)

type FltCurCfgTableHdrHash int32

const (
	FltCurCfgTableHdrHash_None        FltCurCfgTableHdrHash = 1
	FltCurCfgTableHdrHash_Headerhash  FltCurCfgTableHdrHash = 2
	FltCurCfgTableHdrHash_Unsupported FltCurCfgTableHdrHash = 2147483647
)

type FltCurCfgTableL3Filter int32

const (
	FltCurCfgTableL3Filter_Enabled     FltCurCfgTableL3Filter = 1
	FltCurCfgTableL3Filter_Disabled    FltCurCfgTableL3Filter = 2
	FltCurCfgTableL3Filter_Unsupported FltCurCfgTableL3Filter = 2147483647
)

type FltCurCfgTableL7SipFilt int32

const (
	FltCurCfgTableL7SipFilt_Enabled     FltCurCfgTableL7SipFilt = 1
	FltCurCfgTableL7SipFilt_Disabled    FltCurCfgTableL7SipFilt = 2
	FltCurCfgTableL7SipFilt_Unsupported FltCurCfgTableL7SipFilt = 2147483647
)

type FltCurCfgTableNbind int32

const (
	FltCurCfgTableNbind_Enabled     FltCurCfgTableNbind = 1
	FltCurCfgTableNbind_Disabled    FltCurCfgTableNbind = 2
	FltCurCfgTableNbind_Unsupported FltCurCfgTableNbind = 2147483647
)

type FltCurCfgTableLayer7InvertAction int32

const (
	FltCurCfgTableLayer7InvertAction_Enabled     FltCurCfgTableLayer7InvertAction = 1
	FltCurCfgTableLayer7InvertAction_Disabled    FltCurCfgTableLayer7InvertAction = 2
	FltCurCfgTableLayer7InvertAction_Unsupported FltCurCfgTableLayer7InvertAction = 2147483647
)

type FltCurCfgTableRtsrcmac int32

const (
	FltCurCfgTableRtsrcmac_Enabled     FltCurCfgTableRtsrcmac = 1
	FltCurCfgTableRtsrcmac_Disabled    FltCurCfgTableRtsrcmac = 2
	FltCurCfgTableRtsrcmac_Unsupported FltCurCfgTableRtsrcmac = 2147483647
)

type FltCurCfgTableRtproxy int32

const (
	FltCurCfgTableRtproxy_Enabled     FltCurCfgTableRtproxy = 1
	FltCurCfgTableRtproxy_Disabled    FltCurCfgTableRtproxy = 2
	FltCurCfgTableRtproxy_Unsupported FltCurCfgTableRtproxy = 2147483647
)

type FltCurCfgTableSesslog int32

const (
	FltCurCfgTableSesslog_Enabled     FltCurCfgTableSesslog = 1
	FltCurCfgTableSesslog_Disabled    FltCurCfgTableSesslog = 2
	FltCurCfgTableSesslog_Unsupported FltCurCfgTableSesslog = 2147483647
)

type FltCurCfgTableApplicType int32

const (
	FltCurCfgTableApplicType_None        FltCurCfgTableApplicType = 1
	FltCurCfgTableApplicType_Basic       FltCurCfgTableApplicType = 2
	FltCurCfgTableApplicType_Http        FltCurCfgTableApplicType = 3
	FltCurCfgTableApplicType_Sip         FltCurCfgTableApplicType = 4
	FltCurCfgTableApplicType_Dns         FltCurCfgTableApplicType = 5
	FltCurCfgTableApplicType_Smtp        FltCurCfgTableApplicType = 6
	FltCurCfgTableApplicType_Pop3        FltCurCfgTableApplicType = 7
	FltCurCfgTableApplicType_Imap        FltCurCfgTableApplicType = 8
	FltCurCfgTableApplicType_Ftp         FltCurCfgTableApplicType = 9
	FltCurCfgTableApplicType_Unsupported FltCurCfgTableApplicType = 2147483647
)

type FltCurCfgTableRtsport int32

const (
	FltCurCfgTableRtsport_Mod5060     FltCurCfgTableRtsport = 1
	FltCurCfgTableRtsport_Preserve    FltCurCfgTableRtsport = 2
	FltCurCfgTableRtsport_Disabled    FltCurCfgTableRtsport = 3
	FltCurCfgTableRtsport_Unsupported FltCurCfgTableRtsport = 2147483647
)

type FltCurCfgTableSrcAddrType int32

const (
	FltCurCfgTableSrcAddrType_IpAddress   FltCurCfgTableSrcAddrType = 1
	FltCurCfgTableSrcAddrType_Network     FltCurCfgTableSrcAddrType = 2
	FltCurCfgTableSrcAddrType_Unsupported FltCurCfgTableSrcAddrType = 2147483647
)

type FltCurCfgTableDstAddrType int32

const (
	FltCurCfgTableDstAddrType_IpAddress   FltCurCfgTableDstAddrType = 1
	FltCurCfgTableDstAddrType_Network     FltCurCfgTableDstAddrType = 2
	FltCurCfgTableDstAddrType_Unsupported FltCurCfgTableDstAddrType = 2147483647
)

type FltCurCfgTableUdpAge int32

const (
	FltCurCfgTableUdpAge_Onlydns     FltCurCfgTableUdpAge = 1
	FltCurCfgTableUdpAge_Disabled    FltCurCfgTableUdpAge = 2
	FltCurCfgTableUdpAge_Enabled     FltCurCfgTableUdpAge = 3
	FltCurCfgTableUdpAge_Unsupported FltCurCfgTableUdpAge = 2147483647
)

type FltCurCfgTableSslInspectionEna int32

const (
	FltCurCfgTableSslInspectionEna_Enabled     FltCurCfgTableSslInspectionEna = 1
	FltCurCfgTableSslInspectionEna_Disabled    FltCurCfgTableSslInspectionEna = 2
	FltCurCfgTableSslInspectionEna_Unsupported FltCurCfgTableSslInspectionEna = 2147483647
)

type FltCurCfgTableSrvCertGroup int32

const (
	FltCurCfgTableSrvCertGroup_Group       FltCurCfgTableSrvCertGroup = 1
	FltCurCfgTableSrvCertGroup_Cert        FltCurCfgTableSrvCertGroup = 2
	FltCurCfgTableSrvCertGroup_None        FltCurCfgTableSrvCertGroup = 3
	FltCurCfgTableSrvCertGroup_Unsupported FltCurCfgTableSrvCertGroup = 2147483647
)

type FltCurCfgTableMatchDev int32

const (
	FltCurCfgTableMatchDev_Allexclif   FltCurCfgTableMatchDev = 1
	FltCurCfgTableMatchDev_All         FltCurCfgTableMatchDev = 2
	FltCurCfgTableMatchDev_None        FltCurCfgTableMatchDev = 3
	FltCurCfgTableMatchDev_Unsupported FltCurCfgTableMatchDev = 2147483647
)

type FltCurCfgTableSslL7Action int32

const (
	FltCurCfgTableSslL7Action_None        FltCurCfgTableSslL7Action = 1
	FltCurCfgTableSslL7Action_Bypass      FltCurCfgTableSslL7Action = 2
	FltCurCfgTableSslL7Action_Inspect     FltCurCfgTableSslL7Action = 3
	FltCurCfgTableSslL7Action_Unsupported FltCurCfgTableSslL7Action = 2147483647
)

type FltCurCfgTableFallback int32

const (
	FltCurCfgTableFallback_Allow       FltCurCfgTableFallback = 1
	FltCurCfgTableFallback_Deny        FltCurCfgTableFallback = 2
	FltCurCfgTableFallback_Goto        FltCurCfgTableFallback = 5
	FltCurCfgTableFallback_ContFlow    FltCurCfgTableFallback = 6
	FltCurCfgTableFallback_Unsupported FltCurCfgTableFallback = 2147483647
)

type FltCurCfgTableVpnFlood int32

const (
	FltCurCfgTableVpnFlood_Enabled     FltCurCfgTableVpnFlood = 1
	FltCurCfgTableVpnFlood_Disabled    FltCurCfgTableVpnFlood = 2
	FltCurCfgTableVpnFlood_Unsupported FltCurCfgTableVpnFlood = 2147483647
)

type FltCurCfgTableReportState int32

const (
	FltCurCfgTableReportState_Disabled    FltCurCfgTableReportState = 0
	FltCurCfgTableReportState_Enabled     FltCurCfgTableReportState = 1
	FltCurCfgTableReportState_Unsupported FltCurCfgTableReportState = 2147483647
)

type FltCurCfgTableReportLocation int32

const (
	FltCurCfgTableReportLocation_None        FltCurCfgTableReportLocation = 0
	FltCurCfgTableReportLocation_Clientside  FltCurCfgTableReportLocation = 1
	FltCurCfgTableReportLocation_Serverside  FltCurCfgTableReportLocation = 2
	FltCurCfgTableReportLocation_Unsupported FltCurCfgTableReportLocation = 2147483647
)

type FltCurCfgTableReportPurpose int32

const (
	FltCurCfgTableReportPurpose_None        FltCurCfgTableReportPurpose = 0
	FltCurCfgTableReportPurpose_Bypass      FltCurCfgTableReportPurpose = 1
	FltCurCfgTableReportPurpose_Inspect     FltCurCfgTableReportPurpose = 2
	FltCurCfgTableReportPurpose_Unsupported FltCurCfgTableReportPurpose = 2147483647
)

type FltCurCfgTableReportAppl int32

const (
	FltCurCfgTableReportAppl_None        FltCurCfgTableReportAppl = 0
	FltCurCfgTableReportAppl_Https       FltCurCfgTableReportAppl = 1
	FltCurCfgTableReportAppl_Http        FltCurCfgTableReportAppl = 2
	FltCurCfgTableReportAppl_Unsupported FltCurCfgTableReportAppl = 2147483647
)

type FltCurCfgTableReportDir int32

const (
	FltCurCfgTableReportDir_Outbound    FltCurCfgTableReportDir = 0
	FltCurCfgTableReportDir_Inbound     FltCurCfgTableReportDir = 1
	FltCurCfgTableReportDir_None        FltCurCfgTableReportDir = 2
	FltCurCfgTableReportDir_Unsupported FltCurCfgTableReportDir = 2147483647
)

type FltCurCfgTableDpmReportState int32

const (
	FltCurCfgTableDpmReportState_Disabled    FltCurCfgTableDpmReportState = 0
	FltCurCfgTableDpmReportState_Enabled     FltCurCfgTableDpmReportState = 1
	FltCurCfgTableDpmReportState_Unsupported FltCurCfgTableDpmReportState = 2147483647
)

type FltCurCfgTableMacToMe int32

const (
	FltCurCfgTableMacToMe_Enabled     FltCurCfgTableMacToMe = 1
	FltCurCfgTableMacToMe_Disabled    FltCurCfgTableMacToMe = 2
	FltCurCfgTableMacToMe_Unsupported FltCurCfgTableMacToMe = 2147483647
)

type FltCurCfgTableRtSrcTnl int32

const (
	FltCurCfgTableRtSrcTnl_Enabled     FltCurCfgTableRtSrcTnl = 1
	FltCurCfgTableRtSrcTnl_Disabled    FltCurCfgTableRtSrcTnl = 2
	FltCurCfgTableRtSrcTnl_Unsupported FltCurCfgTableRtSrcTnl = 2147483647
)

type FltCurCfgTableAwinflow int32

const (
	FltCurCfgTableAwinflow_After       FltCurCfgTableAwinflow = 1
	FltCurCfgTableAwinflow_Before      FltCurCfgTableAwinflow = 2
	FltCurCfgTableAwinflow_Unsupported FltCurCfgTableAwinflow = 2147483647
)

type FltCurCfgTableJsinject int32

const (
	FltCurCfgTableJsinject_Enabled     FltCurCfgTableJsinject = 1
	FltCurCfgTableJsinject_Disabled    FltCurCfgTableJsinject = 2
	FltCurCfgTableJsinject_Unsupported FltCurCfgTableJsinject = 2147483647
)

type FltCurCfgTableForceBind int32

const (
	FltCurCfgTableForceBind_Enabled     FltCurCfgTableForceBind = 1
	FltCurCfgTableForceBind_Disabled    FltCurCfgTableForceBind = 2
	FltCurCfgTableForceBind_Unsupported FltCurCfgTableForceBind = 2147483647
)

type FltCurCfgTableParams struct {
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
	// only when protocol defined in fltCurCfgProtocol is UDP or TCP.
	// '0' means no filtering.
	RangeHighSrcPort uint64 `json:"RangeHighSrcPort,omitempty"`
	// The lower source TCP/UDP port number to be filtered. It applies
	// only when protocol defined in fltCurCfgProtocol is UDP or TCP.
	// '0' means no filtering.
	RangeLowSrcPort uint64 `json:"RangeLowSrcPort,omitempty"`
	// The lower destination TCP/UDP port number to be filtered. It applies
	// only when protocol defined in fltCurCfgProtocol is UDP or TCP.
	// '0' means no filtering.
	RangeLowDstPort uint64 `json:"RangeLowDstPort,omitempty"`
	// The higher destination TCP/UDP port number to be filtered. It applies
	// only when protocol defined in fltCurCfgProtocol is UDP or TCP.
	// '0' means no filtering.
	RangeHighDstPort uint64 `json:"RangeHighDstPort,omitempty"`
	// The action for the filtering rule.
	Action FltCurCfgTableAction `json:"Action,omitempty"`
	// The real server port number used for redirection.
	RedirPort uint64 `json:"RedirPort,omitempty"`
	// The real server group to be redirected to.
	RedirGroup int32 `json:"RedirGroup,omitempty"`
	// Enable or disable logging.
	Log FltCurCfgTableLog `json:"Log,omitempty"`
	// The state of this filtering rule.
	State FltCurCfgTableState `json:"State,omitempty"`
	// The selection of destination or source or multicast for
	// network address translation.
	Nat FltCurCfgTableNat `json:"Nat,omitempty"`
	// Enable or disable caching sessions that match filter.
	Cache FltCurCfgTableCache `json:"Cache,omitempty"`
	// Turn the invert logic on or off for the filter entry.
	Invert FltCurCfgTableInvert `json:"Invert,omitempty"`
	// Enable or disable client proxy.
	ClientProxy FltCurCfgTableClientProxy `json:"ClientProxy,omitempty"`
	// Enable or disable filtering on matching TCP ACK and RST flag.
	TcpAck FltCurCfgTableTcpAck `json:"TcpAck,omitempty"`
	// The source MAC address to be filtered.
	SrcMac string `json:"SrcMac,omitempty"`
	// The Destination MAC address to be filtered.
	DstMac string `json:"DstMac,omitempty"`
	// Enable or disable FTP NAT for active ftp only.
	FtpNatActive FltCurCfgTableFtpNatActive `json:"FtpNatActive,omitempty"`
	// Enable or disable TCP URG packet.
	AclTcpUrg FltCurCfgTableAclTcpUrg `json:"AclTcpUrg,omitempty"`
	// Enable or disable TCP ACK packet.
	AclTcpAck FltCurCfgTableAclTcpAck `json:"AclTcpAck,omitempty"`
	// Enable or disable TCP PSH packet.
	AclTcpPsh FltCurCfgTableAclTcpPsh `json:"AclTcpPsh,omitempty"`
	// Enable or disable TCP RST packet.
	AclTcpRst FltCurCfgTableAclTcpRst `json:"AclTcpRst,omitempty"`
	// Enable or disable TCP SYN packet.
	AclTcpSyn FltCurCfgTableAclTcpSyn `json:"AclTcpSyn,omitempty"`
	// Enable or disable TCP FIN packet.
	AclTcpFin FltCurCfgTableAclTcpFin `json:"AclTcpFin,omitempty"`
	// ICMP type to be filtered. A ICMP type of 255 indicates 'any'
	AclIcmp uint32 `json:"AclIcmp,omitempty"`
	// Enable or disable IP option matching.
	AclIpOption FltCurCfgTableAclIpOption `json:"AclIpOption,omitempty"`
	// Filt default BW contract number.
	BwmContract int32 `json:"BwmContract,omitempty"`
	// The IP TOS value to be filtered.
	AclIpTos uint32 `json:"AclIpTos,omitempty"`
	// The IP TOS mask for filtering.
	AclIpTosMask uint32 `json:"AclIpTosMask,omitempty"`
	// The new IP TOS value to over-write when filtering fired.
	AclIpTosNew uint32 `json:"AclIpTosNew,omitempty"`
	// Enable or disable filtering on firewall redirect hash method.
	Fwlb FltCurCfgTableFwlb `json:"Fwlb,omitempty"`
	// The NAT session timeout. The timeout value should be an even number
	// between 4 and 32768.
	NatTimeout uint32 `json:"NatTimeout,omitempty"`
	// Enable or disable WAN link load balancing.
	Linklb FltCurCfgTableLinklb `json:"Linklb,omitempty"`
	// Enable or disable WAP RADIUS snooping.
	WapRadiusSnoop FltCurCfgTableWapRadiusSnoop `json:"WapRadiusSnoop,omitempty"`
	// Set a flag indicating whether filtering should be based on the
	// source IP address or the source MAC address field.
	SrcIpMac FltCurCfgTableSrcIpMac `json:"SrcIpMac,omitempty"`
	// Set a flag indicating whether filtering should be based on the
	// destination IP address or the destination MAC address field.
	DstIpMac FltCurCfgTableDstIpMac `json:"DstIpMac,omitempty"`
	// Set hash parameter for intrusion detection server load balancing.
	IdslbHash FltCurCfgTableIdslbHash `json:"IdslbHash,omitempty"`
	// Set the VLAN assoicated with the filter.
	Vlan uint32 `json:"Vlan,omitempty"`
	// The name of the filter.
	Name string `json:"Name,omitempty"`
	// Enable or disable protocol rate limiting.
	TcpRateLimit FltCurCfgTableTcpRateLimit `json:"TcpRateLimit,omitempty"`
	// Set maximum connections (number of connections in units of 10)
	// for TCP connection rate limiting.  In the case of ICMP and UDP,
	// 	 this is the maximum packets (number of packets in units of 10).
	TcpRateMaxConn uint32 `json:"TcpRateMaxConn,omitempty"`
	// Set hash parameter for the filter.
	Hash FltCurCfgTableHash `json:"Hash,omitempty"`
	// Enable or disable layer 7 deny filtering.
	Layer7DenyState FltCurCfgTableLayer7DenyState `json:"Layer7DenyState,omitempty"`
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
	// The filter ID for GOTO action in the current config.
	GotoFilter int32 `json:"GotoFilter,omitempty"`
	// Enable or disable Radius/WAP persistence.
	RadiusWapPersist FltCurCfgTableRadiusWapPersist `json:"RadiusWapPersist,omitempty"`
	// Enable or disable filter persistent binding.
	Pbind FltCurCfgTablePbind `json:"Pbind,omitempty"`
	// The time window for protocol rate limiting (in seconds).
	TimeWindow uint64 `json:"TimeWindow,omitempty"`
	// The hold down duration for protocol rate limiting (in minutes).
	HoldDuration uint64 `json:"HoldDuration,omitempty"`
	// Enable or disable binary pattern matching.
	PatternMatch FltCurCfgTablePatternMatch `json:"PatternMatch,omitempty"`
	// Enable or disable match-all criteria for L7 deny string matching.
	Layer7DenyMatchAll FltCurCfgTableLayer7DenyMatchAll `json:"Layer7DenyMatchAll,omitempty"`
	// The client proxy IP address for NAT and REDIR filter.
	ProxyIp string `json:"ProxyIp,omitempty"`
	// Enable or disable layer 7 lookup (parsing) of all packets.
	Layer7ParseAll FltCurCfgTableLayer7ParseAll `json:"Layer7ParseAll,omitempty"`
	// Enable or disable pattern string lookup (parsing) of all packets.
	SecurityParseAll FltCurCfgTableSecurityParseAll `json:"SecurityParseAll,omitempty"`
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
	// The 802.1p bits value to match.
	Cfg8021pBitsValue uint64 `json:"Cfg8021pBitsValue,omitempty"`
	// Enable or disable matching on 802.1p bits in the packets.
	Cfg8021pBitsMatch FltCurCfgTableCfg8021pBitsMatch `json:"Cfg8021pBitsMatch,omitempty"`
	// Set the IP maximum packet length in bytes. A value can be either 0
	// which indicates 'any' length or between 64 and 65535.
	AclIpLength uint64 `json:"AclIpLength,omitempty"`
	// The real server group for IDS load balancing. A value of 0 indicates
	// 	 'none'.
	IdsGroup int32 `json:"IdsGroup,omitempty"`
	// Enable or disable pip selection based on egress port/vlan.
	EgressPip FltCurCfgTableEgressPip `json:"EgressPip,omitempty"`
	// Enable or disable filter delayed binding.
	Dbind FltCurCfgTableDbind `json:"Dbind,omitempty"`
	// Filt reverse session BWM contract number.
	RevBwmContract int32 `json:"RevBwmContract,omitempty"`
	// Enable or disable creating session for reverse side traffic.
	Reverse FltCurCfgTableReverse `json:"Reverse,omitempty"`
	// Enable or disable chained pgroup match criteria for l7 filtering.
	ParseChn FltCurCfgTableParseChn `json:"ParseChn,omitempty"`
	// BWM contract for SIP RTP traffic.
	RtpBwmContract int32 `json:"RtpBwmContract,omitempty"`
	// Enable or disable SIP NAT.
	SipParsing FltCurCfgTableSipParsing `json:"SipParsing,omitempty"`
	// Enable or disable session mirroring.
	SessionMirror FltCurCfgTableSessionMirror `json:"SessionMirror,omitempty"`
	// The type of IP address.
	IpVer FltCurCfgTableIpVer `json:"IpVer,omitempty"`
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
	HdrHash FltCurCfgTableHdrHash `json:"HdrHash,omitempty"`
	// The header name of the filter.
	HdrName string `json:"HdrName,omitempty"`
	// The header hash length of the filter.
	HdrHashLen uint32 `json:"HdrHashLen,omitempty"`
	// The L3 filter processing state for this filter.
	L3Filter FltCurCfgTableL3Filter `json:"L3Filter,omitempty"`
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
	L7SipFilt FltCurCfgTableL7SipFilt `json:"L7SipFilt,omitempty"`
	// Multicast VLAN.
	NatMcastVlan uint32 `json:"NatMcastVlan,omitempty"`
	// Enable or disable subnet binding for redirection.
	Nbind FltCurCfgTableNbind `json:"Nbind,omitempty"`
	// Enable or disable invert action for layer 7 string matching .
	Layer7InvertAction FltCurCfgTableLayer7InvertAction `json:"Layer7InvertAction,omitempty"`
	// The source network class ID to be filtered.
	SrcClassId string `json:"SrcClassId,omitempty"`
	// The destination network class ID to be filtered.
	DstClassId string `json:"DstClassId,omitempty"`
	// Enable or disable return to source mac addr.
	Rtsrcmac FltCurCfgTableRtsrcmac `json:"Rtsrcmac,omitempty"`
	// Enable or disable redirect to proxy server.
	Rtproxy FltCurCfgTableRtproxy `json:"Rtproxy,omitempty"`
	// Enable or disable session logging.
	Sesslog FltCurCfgTableSesslog `json:"Sesslog,omitempty"`
	// The cntclass of the filter.
	Cntclass string `json:"Cntclass,omitempty"`
	// Application type for this filter.
	ApplicType FltCurCfgTableApplicType `json:"ApplicType,omitempty"`
	// The real server group to be redirected to.
	EnhRedirGroup string `json:"EnhRedirGroup,omitempty"`
	// The real server group for IDS load balancing. A NULL string indicates
	// 'none'.
	EnhIdsGroup string `json:"EnhIdsGroup,omitempty"`
	// Source port modification for SIP.
	// the mod5060 indicates modification of source port to 5060,
	// preserve indicates to retain the source port.
	Rtsport FltCurCfgTableRtsport `json:"Rtsport,omitempty"`
	// Source address type from address and network for NAT filter.
	SrcAddrType FltCurCfgTableSrcAddrType `json:"SrcAddrType,omitempty"`
	// Destination address type from address and network for NAT filter.
	DstAddrType FltCurCfgTableDstAddrType `json:"DstAddrType,omitempty"`
	// Filter source Ip. It can be IPv4/IPv6.
	SourceIp string `json:"SourceIp,omitempty"`
	// Filter Destination Ip. It can be IPv4/IPv6.
	DestIp string `json:"DestIp,omitempty"`
	// Filter Mask or Prefix.
	SourceMask string `json:"SourceMask,omitempty"`
	// Filter Mask or Prefix.
	DestMask string `json:"DestMask,omitempty"`
	// The class of service of the filter.
	// 	Empty string - ignore cos
	// 	'Any' string - ignore string use only the source IP
	// 	Specific string - match only if the string matches
	CosStr string `json:"CosStr,omitempty"`
	// Fast aging of UDP sessions.
	UdpAge FltCurCfgTableUdpAge `json:"UdpAge,omitempty"`
	// frontend TCP optimization policy
	FeTcpPolId string `json:"FeTcpPolId,omitempty"`
	// backend TCP optimization policy
	BeTcpPolId string `json:"BeTcpPolId,omitempty"`
	// Compression policy name associated with this filter.
	Comppol string `json:"Comppol,omitempty"`
	// Enable or disable SSL inspection inspection.
	SslInspectionEna FltCurCfgTableSslInspectionEna `json:"SslInspectionEna,omitempty"`
	// Server Certificate or Group.
	SrvCertGroup FltCurCfgTableSrvCertGroup `json:"SrvCertGroup,omitempty"`
	// Server certificate / group name associated with this filter.
	SrvCert string `json:"SrvCert,omitempty"`
	// SSL policy name associated with this filter.
	SslPolicy string `json:"SslPolicy,omitempty"`
	// Match device IP addresses in this filter:
	// allexclif : match all device IPs, but exclude interface IPs
	// all       : match all device IPs
	// none      : match no device IPs
	MatchDev FltCurCfgTableMatchDev `json:"MatchDev,omitempty"`
	// Set SSL inspection action on L7 match.
	SslL7Action FltCurCfgTableSslL7Action `json:"SslL7Action,omitempty"`
	// Fallback action (group down).
	Fallback FltCurCfgTableFallback `json:"Fallback,omitempty"`
	// The filter ID for GOTO fallback action in the current config.
	Fbgoto int32 `json:"Fbgoto,omitempty"`
	// Enable or disable Two Way VPN LB.
	VpnFlood FltCurCfgTableVpnFlood `json:"VpnFlood,omitempty"`
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
	// The ingress port for fallback action continueFlow in the current config.
	Fbport int32 `json:"Fbport,omitempty"`
	// Group Index of filters on same tunnel.
	SetIndex int32 `json:"SetIndex,omitempty"`
	// Enable or disable filter inspect reporting.
	ReportState FltCurCfgTableReportState `json:"ReportState,omitempty"`
	// get filter location none/client/server.
	ReportLocation FltCurCfgTableReportLocation `json:"ReportLocation,omitempty"`
	// get filter purpose none/bypass/inspect.
	ReportPurpose FltCurCfgTableReportPurpose `json:"ReportPurpose,omitempty"`
	// get filter application none/https/http.
	ReportAppl FltCurCfgTableReportAppl `json:"ReportAppl,omitempty"`
	// get filter direction inbound/outbound.
	ReportDir FltCurCfgTableReportDir `json:"ReportDir,omitempty"`
	// The icap policy of the filter.
	IcapPol string `json:"IcapPol,omitempty"`
	// The traffic event policy of the filter.
	TrafficEventPol string `json:"TrafficEventPol,omitempty"`
	// Enable or disable filter dpm reporting.
	DpmReportState FltCurCfgTableDpmReportState `json:"DpmReportState,omitempty"`
	// Enable/disable matching traffic to device mac addresses.
	MacToMe FltCurCfgTableMacToMe `json:"MacToMe,omitempty"`
	// Secured web application associated with this filter.
	Secwa string `json:"Secwa,omitempty"`
	// The ingress VLAN ID for fallback action continueFlow in the current config.
	Fbvlan uint32 `json:"Fbvlan,omitempty"`
	// URL filter classification mode.
	Urlfmode uint32 `json:"Urlfmode,omitempty"`
	// Enable or disable Return to Source Tunnel.
	RtSrcTnl FltCurCfgTableRtSrcTnl `json:"RtSrcTnl,omitempty"`
	// Aw monitor private key name.
	SslawMonPriKey string `json:"SslawMonPriKey,omitempty"`
	// Aw monitor policy name.
	SslawMonPolicy string `json:"SslawMonPolicy,omitempty"`
	// Get Bot Manager Policy.
	Botpol string `json:"Botpol,omitempty"`
	// Set if AW processing comes before or after Alteon HTTP parsing.
	Awinflow FltCurCfgTableAwinflow `json:"Awinflow,omitempty"`
	// Server GmSSL signing certificate / group name associated with this filter.
	SrvCertSign string `json:"SrvCertSign,omitempty"`
	// Server GmSSL encryption certificate / group name associated with this filter.
	SrvCertEnc string `json:"SrvCertEnc,omitempty"`
	// Get Sideband policy.
	SidebandID string `json:"SidebandID,omitempty"`
	// Get secure path Policy.
	SecurePathPolicy string `json:"SecurePathPolicy,omitempty"`
	// Get JS inject mode .
	Jsinject FltCurCfgTableJsinject `json:"Jsinject,omitempty"`
	// Enable or disable Force bind for proxy.
	ForceBind FltCurCfgTableForceBind `json:"ForceBind,omitempty"`
}

func (p FltCurCfgTableParams) iMABean() {}
