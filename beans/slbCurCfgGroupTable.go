package beans

import (
	"fmt"
	"reflect"
)

// SlbCurCfgGroupTable The table of groups,
// Note:This mib is not supported on VX instance of Virtualization.
type SlbCurCfgGroupTable struct {
	// The group number for which the information pertains.
	SlbCurCfgGroupIndex int32
	Params              *SlbCurCfgGroupTableParams
}

func NewSlbCurCfgGroupTableList() *SlbCurCfgGroupTable {
	return &SlbCurCfgGroupTable{}
}

func NewSlbCurCfgGroupTable(
	slbCurCfgGroupIndex int32,
	params *SlbCurCfgGroupTableParams,
) *SlbCurCfgGroupTable {
	return &SlbCurCfgGroupTable{
		SlbCurCfgGroupIndex: slbCurCfgGroupIndex,
		Params:              params,
	}
}

func (c *SlbCurCfgGroupTable) Name() string {
	return "SlbCurCfgGroupTable"
}

func (c *SlbCurCfgGroupTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbCurCfgGroupTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbCurCfgGroupTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbCurCfgGroupIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbCurCfgGroupIndex)
}

type SlbCurCfgGroupTableMetric int32

const (
	SlbCurCfgGroupTableMetric_RoundRobin       SlbCurCfgGroupTableMetric = 1
	SlbCurCfgGroupTableMetric_LeastConnections SlbCurCfgGroupTableMetric = 2
	SlbCurCfgGroupTableMetric_MinMisses        SlbCurCfgGroupTableMetric = 3
	SlbCurCfgGroupTableMetric_Hash             SlbCurCfgGroupTableMetric = 4
	SlbCurCfgGroupTableMetric_Response         SlbCurCfgGroupTableMetric = 5
	SlbCurCfgGroupTableMetric_Bandwidth        SlbCurCfgGroupTableMetric = 6
	SlbCurCfgGroupTableMetric_Phash            SlbCurCfgGroupTableMetric = 7
	SlbCurCfgGroupTableMetric_SvcLeast         SlbCurCfgGroupTableMetric = 8
	SlbCurCfgGroupTableMetric_Unsupported      SlbCurCfgGroupTableMetric = 2147483647
)

type SlbCurCfgGroupTableHealthCheckLayer int32

const (
	SlbCurCfgGroupTableHealthCheckLayer_Icmp        SlbCurCfgGroupTableHealthCheckLayer = 1
	SlbCurCfgGroupTableHealthCheckLayer_Tcp         SlbCurCfgGroupTableHealthCheckLayer = 2
	SlbCurCfgGroupTableHealthCheckLayer_Http        SlbCurCfgGroupTableHealthCheckLayer = 3
	SlbCurCfgGroupTableHealthCheckLayer_Dns         SlbCurCfgGroupTableHealthCheckLayer = 4
	SlbCurCfgGroupTableHealthCheckLayer_Smtp        SlbCurCfgGroupTableHealthCheckLayer = 5
	SlbCurCfgGroupTableHealthCheckLayer_Pop3        SlbCurCfgGroupTableHealthCheckLayer = 6
	SlbCurCfgGroupTableHealthCheckLayer_Nntp        SlbCurCfgGroupTableHealthCheckLayer = 7
	SlbCurCfgGroupTableHealthCheckLayer_Ftp         SlbCurCfgGroupTableHealthCheckLayer = 8
	SlbCurCfgGroupTableHealthCheckLayer_Imap        SlbCurCfgGroupTableHealthCheckLayer = 9
	SlbCurCfgGroupTableHealthCheckLayer_Radius      SlbCurCfgGroupTableHealthCheckLayer = 10
	SlbCurCfgGroupTableHealthCheckLayer_Sslh        SlbCurCfgGroupTableHealthCheckLayer = 11
	SlbCurCfgGroupTableHealthCheckLayer_Script1     SlbCurCfgGroupTableHealthCheckLayer = 12
	SlbCurCfgGroupTableHealthCheckLayer_Script2     SlbCurCfgGroupTableHealthCheckLayer = 13
	SlbCurCfgGroupTableHealthCheckLayer_Script3     SlbCurCfgGroupTableHealthCheckLayer = 14
	SlbCurCfgGroupTableHealthCheckLayer_Script4     SlbCurCfgGroupTableHealthCheckLayer = 15
	SlbCurCfgGroupTableHealthCheckLayer_Script5     SlbCurCfgGroupTableHealthCheckLayer = 16
	SlbCurCfgGroupTableHealthCheckLayer_Script6     SlbCurCfgGroupTableHealthCheckLayer = 17
	SlbCurCfgGroupTableHealthCheckLayer_Script7     SlbCurCfgGroupTableHealthCheckLayer = 18
	SlbCurCfgGroupTableHealthCheckLayer_Script8     SlbCurCfgGroupTableHealthCheckLayer = 19
	SlbCurCfgGroupTableHealthCheckLayer_Script9     SlbCurCfgGroupTableHealthCheckLayer = 20
	SlbCurCfgGroupTableHealthCheckLayer_Script10    SlbCurCfgGroupTableHealthCheckLayer = 21
	SlbCurCfgGroupTableHealthCheckLayer_Script11    SlbCurCfgGroupTableHealthCheckLayer = 22
	SlbCurCfgGroupTableHealthCheckLayer_Script12    SlbCurCfgGroupTableHealthCheckLayer = 23
	SlbCurCfgGroupTableHealthCheckLayer_Script13    SlbCurCfgGroupTableHealthCheckLayer = 24
	SlbCurCfgGroupTableHealthCheckLayer_Script14    SlbCurCfgGroupTableHealthCheckLayer = 25
	SlbCurCfgGroupTableHealthCheckLayer_Script15    SlbCurCfgGroupTableHealthCheckLayer = 26
	SlbCurCfgGroupTableHealthCheckLayer_Script16    SlbCurCfgGroupTableHealthCheckLayer = 27
	SlbCurCfgGroupTableHealthCheckLayer_Link        SlbCurCfgGroupTableHealthCheckLayer = 28
	SlbCurCfgGroupTableHealthCheckLayer_Wsp         SlbCurCfgGroupTableHealthCheckLayer = 29
	SlbCurCfgGroupTableHealthCheckLayer_Wtls        SlbCurCfgGroupTableHealthCheckLayer = 30
	SlbCurCfgGroupTableHealthCheckLayer_Ldap        SlbCurCfgGroupTableHealthCheckLayer = 31
	SlbCurCfgGroupTableHealthCheckLayer_Udpdns      SlbCurCfgGroupTableHealthCheckLayer = 32
	SlbCurCfgGroupTableHealthCheckLayer_Arp         SlbCurCfgGroupTableHealthCheckLayer = 33
	SlbCurCfgGroupTableHealthCheckLayer_Snmp1       SlbCurCfgGroupTableHealthCheckLayer = 34
	SlbCurCfgGroupTableHealthCheckLayer_Snmp2       SlbCurCfgGroupTableHealthCheckLayer = 35
	SlbCurCfgGroupTableHealthCheckLayer_Snmp3       SlbCurCfgGroupTableHealthCheckLayer = 36
	SlbCurCfgGroupTableHealthCheckLayer_Snmp4       SlbCurCfgGroupTableHealthCheckLayer = 37
	SlbCurCfgGroupTableHealthCheckLayer_Snmp5       SlbCurCfgGroupTableHealthCheckLayer = 38
	SlbCurCfgGroupTableHealthCheckLayer_Radiusacs   SlbCurCfgGroupTableHealthCheckLayer = 39
	SlbCurCfgGroupTableHealthCheckLayer_Tftp        SlbCurCfgGroupTableHealthCheckLayer = 40
	SlbCurCfgGroupTableHealthCheckLayer_Wtp         SlbCurCfgGroupTableHealthCheckLayer = 41
	SlbCurCfgGroupTableHealthCheckLayer_Rtsp        SlbCurCfgGroupTableHealthCheckLayer = 42
	SlbCurCfgGroupTableHealthCheckLayer_Sipping     SlbCurCfgGroupTableHealthCheckLayer = 43
	SlbCurCfgGroupTableHealthCheckLayer_Httphead    SlbCurCfgGroupTableHealthCheckLayer = 44
	SlbCurCfgGroupTableHealthCheckLayer_Sipoptions  SlbCurCfgGroupTableHealthCheckLayer = 45
	SlbCurCfgGroupTableHealthCheckLayer_Wts         SlbCurCfgGroupTableHealthCheckLayer = 46
	SlbCurCfgGroupTableHealthCheckLayer_Dhcp        SlbCurCfgGroupTableHealthCheckLayer = 47
	SlbCurCfgGroupTableHealthCheckLayer_Radiusaa    SlbCurCfgGroupTableHealthCheckLayer = 48
	SlbCurCfgGroupTableHealthCheckLayer_Sslv3       SlbCurCfgGroupTableHealthCheckLayer = 49
	SlbCurCfgGroupTableHealthCheckLayer_Script17    SlbCurCfgGroupTableHealthCheckLayer = 116
	SlbCurCfgGroupTableHealthCheckLayer_Script18    SlbCurCfgGroupTableHealthCheckLayer = 117
	SlbCurCfgGroupTableHealthCheckLayer_Script19    SlbCurCfgGroupTableHealthCheckLayer = 118
	SlbCurCfgGroupTableHealthCheckLayer_Script20    SlbCurCfgGroupTableHealthCheckLayer = 119
	SlbCurCfgGroupTableHealthCheckLayer_Script21    SlbCurCfgGroupTableHealthCheckLayer = 120
	SlbCurCfgGroupTableHealthCheckLayer_Script22    SlbCurCfgGroupTableHealthCheckLayer = 121
	SlbCurCfgGroupTableHealthCheckLayer_Script23    SlbCurCfgGroupTableHealthCheckLayer = 122
	SlbCurCfgGroupTableHealthCheckLayer_Script24    SlbCurCfgGroupTableHealthCheckLayer = 123
	SlbCurCfgGroupTableHealthCheckLayer_Script25    SlbCurCfgGroupTableHealthCheckLayer = 124
	SlbCurCfgGroupTableHealthCheckLayer_Script26    SlbCurCfgGroupTableHealthCheckLayer = 125
	SlbCurCfgGroupTableHealthCheckLayer_Script27    SlbCurCfgGroupTableHealthCheckLayer = 126
	SlbCurCfgGroupTableHealthCheckLayer_Script28    SlbCurCfgGroupTableHealthCheckLayer = 127
	SlbCurCfgGroupTableHealthCheckLayer_Script29    SlbCurCfgGroupTableHealthCheckLayer = 128
	SlbCurCfgGroupTableHealthCheckLayer_Script30    SlbCurCfgGroupTableHealthCheckLayer = 129
	SlbCurCfgGroupTableHealthCheckLayer_Script31    SlbCurCfgGroupTableHealthCheckLayer = 130
	SlbCurCfgGroupTableHealthCheckLayer_Script32    SlbCurCfgGroupTableHealthCheckLayer = 131
	SlbCurCfgGroupTableHealthCheckLayer_Script33    SlbCurCfgGroupTableHealthCheckLayer = 132
	SlbCurCfgGroupTableHealthCheckLayer_Script34    SlbCurCfgGroupTableHealthCheckLayer = 133
	SlbCurCfgGroupTableHealthCheckLayer_Script35    SlbCurCfgGroupTableHealthCheckLayer = 134
	SlbCurCfgGroupTableHealthCheckLayer_Script36    SlbCurCfgGroupTableHealthCheckLayer = 135
	SlbCurCfgGroupTableHealthCheckLayer_Script37    SlbCurCfgGroupTableHealthCheckLayer = 136
	SlbCurCfgGroupTableHealthCheckLayer_Script38    SlbCurCfgGroupTableHealthCheckLayer = 137
	SlbCurCfgGroupTableHealthCheckLayer_Script39    SlbCurCfgGroupTableHealthCheckLayer = 138
	SlbCurCfgGroupTableHealthCheckLayer_Script40    SlbCurCfgGroupTableHealthCheckLayer = 139
	SlbCurCfgGroupTableHealthCheckLayer_Script41    SlbCurCfgGroupTableHealthCheckLayer = 140
	SlbCurCfgGroupTableHealthCheckLayer_Script42    SlbCurCfgGroupTableHealthCheckLayer = 141
	SlbCurCfgGroupTableHealthCheckLayer_Script43    SlbCurCfgGroupTableHealthCheckLayer = 142
	SlbCurCfgGroupTableHealthCheckLayer_Script44    SlbCurCfgGroupTableHealthCheckLayer = 143
	SlbCurCfgGroupTableHealthCheckLayer_Script45    SlbCurCfgGroupTableHealthCheckLayer = 144
	SlbCurCfgGroupTableHealthCheckLayer_Script46    SlbCurCfgGroupTableHealthCheckLayer = 145
	SlbCurCfgGroupTableHealthCheckLayer_Script47    SlbCurCfgGroupTableHealthCheckLayer = 146
	SlbCurCfgGroupTableHealthCheckLayer_Script48    SlbCurCfgGroupTableHealthCheckLayer = 147
	SlbCurCfgGroupTableHealthCheckLayer_Script49    SlbCurCfgGroupTableHealthCheckLayer = 148
	SlbCurCfgGroupTableHealthCheckLayer_Script50    SlbCurCfgGroupTableHealthCheckLayer = 149
	SlbCurCfgGroupTableHealthCheckLayer_Script51    SlbCurCfgGroupTableHealthCheckLayer = 150
	SlbCurCfgGroupTableHealthCheckLayer_Script52    SlbCurCfgGroupTableHealthCheckLayer = 151
	SlbCurCfgGroupTableHealthCheckLayer_Script53    SlbCurCfgGroupTableHealthCheckLayer = 152
	SlbCurCfgGroupTableHealthCheckLayer_Script54    SlbCurCfgGroupTableHealthCheckLayer = 153
	SlbCurCfgGroupTableHealthCheckLayer_Script55    SlbCurCfgGroupTableHealthCheckLayer = 154
	SlbCurCfgGroupTableHealthCheckLayer_Script56    SlbCurCfgGroupTableHealthCheckLayer = 155
	SlbCurCfgGroupTableHealthCheckLayer_Script57    SlbCurCfgGroupTableHealthCheckLayer = 156
	SlbCurCfgGroupTableHealthCheckLayer_Script58    SlbCurCfgGroupTableHealthCheckLayer = 157
	SlbCurCfgGroupTableHealthCheckLayer_Script59    SlbCurCfgGroupTableHealthCheckLayer = 158
	SlbCurCfgGroupTableHealthCheckLayer_Script60    SlbCurCfgGroupTableHealthCheckLayer = 159
	SlbCurCfgGroupTableHealthCheckLayer_Script61    SlbCurCfgGroupTableHealthCheckLayer = 160
	SlbCurCfgGroupTableHealthCheckLayer_Script62    SlbCurCfgGroupTableHealthCheckLayer = 161
	SlbCurCfgGroupTableHealthCheckLayer_Script63    SlbCurCfgGroupTableHealthCheckLayer = 162
	SlbCurCfgGroupTableHealthCheckLayer_Script64    SlbCurCfgGroupTableHealthCheckLayer = 163
	SlbCurCfgGroupTableHealthCheckLayer_None        SlbCurCfgGroupTableHealthCheckLayer = 164
	SlbCurCfgGroupTableHealthCheckLayer_Unknown     SlbCurCfgGroupTableHealthCheckLayer = 165
	SlbCurCfgGroupTableHealthCheckLayer_Unsupported SlbCurCfgGroupTableHealthCheckLayer = 2147483647
)

type SlbCurCfgGroupTableVipHealthCheck int32

const (
	SlbCurCfgGroupTableVipHealthCheck_Enabled     SlbCurCfgGroupTableVipHealthCheck = 1
	SlbCurCfgGroupTableVipHealthCheck_Disabled    SlbCurCfgGroupTableVipHealthCheck = 2
	SlbCurCfgGroupTableVipHealthCheck_Unsupported SlbCurCfgGroupTableVipHealthCheck = 2147483647
)

type SlbCurCfgGroupTableIdsState int32

const (
	SlbCurCfgGroupTableIdsState_Enabled     SlbCurCfgGroupTableIdsState = 1
	SlbCurCfgGroupTableIdsState_Disabled    SlbCurCfgGroupTableIdsState = 2
	SlbCurCfgGroupTableIdsState_Unsupported SlbCurCfgGroupTableIdsState = 2147483647
)

type SlbCurCfgGroupTableIdsFlood int32

const (
	SlbCurCfgGroupTableIdsFlood_Enabled     SlbCurCfgGroupTableIdsFlood = 1
	SlbCurCfgGroupTableIdsFlood_Disabled    SlbCurCfgGroupTableIdsFlood = 2
	SlbCurCfgGroupTableIdsFlood_Unsupported SlbCurCfgGroupTableIdsFlood = 2147483647
)

type SlbCurCfgGroupTableMinmissHash int32

const (
	SlbCurCfgGroupTableMinmissHash_Minmiss24   SlbCurCfgGroupTableMinmissHash = 1
	SlbCurCfgGroupTableMinmissHash_Minmiss32   SlbCurCfgGroupTableMinmissHash = 2
	SlbCurCfgGroupTableMinmissHash_Unsupported SlbCurCfgGroupTableMinmissHash = 2147483647
)

type SlbCurCfgGroupTableRmetric int32

const (
	SlbCurCfgGroupTableRmetric_RoundRobin       SlbCurCfgGroupTableRmetric = 1
	SlbCurCfgGroupTableRmetric_Hash             SlbCurCfgGroupTableRmetric = 2
	SlbCurCfgGroupTableRmetric_LeastConnections SlbCurCfgGroupTableRmetric = 3
	SlbCurCfgGroupTableRmetric_Unsupported      SlbCurCfgGroupTableRmetric = 2147483647
)

type SlbCurCfgGroupTableOperatorAccess int32

const (
	SlbCurCfgGroupTableOperatorAccess_Enabled     SlbCurCfgGroupTableOperatorAccess = 1
	SlbCurCfgGroupTableOperatorAccess_Disabled    SlbCurCfgGroupTableOperatorAccess = 2
	SlbCurCfgGroupTableOperatorAccess_Unsupported SlbCurCfgGroupTableOperatorAccess = 2147483647
)

type SlbCurCfgGroupTableIpVer int32

const (
	SlbCurCfgGroupTableIpVer_Ipv4        SlbCurCfgGroupTableIpVer = 1
	SlbCurCfgGroupTableIpVer_Ipv6        SlbCurCfgGroupTableIpVer = 2
	SlbCurCfgGroupTableIpVer_Mixed       SlbCurCfgGroupTableIpVer = 3
	SlbCurCfgGroupTableIpVer_Unsupported SlbCurCfgGroupTableIpVer = 2147483647
)

type SlbCurCfgGroupTableIdsChain int32

const (
	SlbCurCfgGroupTableIdsChain_Enabled     SlbCurCfgGroupTableIdsChain = 1
	SlbCurCfgGroupTableIdsChain_Disabled    SlbCurCfgGroupTableIdsChain = 2
	SlbCurCfgGroupTableIdsChain_Unsupported SlbCurCfgGroupTableIdsChain = 2147483647
)

type SlbCurCfgGroupTableMaxConEx int32

const (
	SlbCurCfgGroupTableMaxConEx_Enabled     SlbCurCfgGroupTableMaxConEx = 1
	SlbCurCfgGroupTableMaxConEx_Disabled    SlbCurCfgGroupTableMaxConEx = 2
	SlbCurCfgGroupTableMaxConEx_Unsupported SlbCurCfgGroupTableMaxConEx = 2147483647
)

type SlbCurCfgGroupTableParams struct {
	// The group number for which the information pertains.
	Index int32 `json:"Index,omitempty"`
	// The Real servers in the group. The servers are presented in
	// bitmap format.
	// in receiving order:
	// OCTET 1  OCTET 2  .....
	// xxxxxxxx xxxxxxxx .....
	// |     || |_ server 9
	// |     ||
	// |     ||___ server 8
	// |     |____ server 7
	// |       .    .   .
	// |__________ server 1
	// where x : 1 - The represented server belongs to the group
	// 0 - The represented server does not belong to the group
	RealServers string `json:"RealServers,omitempty"`
	// The metric used to select next server in group.
	Metric SlbCurCfgGroupTableMetric `json:"Metric,omitempty"`
	// The backup real server for this group.
	BackupServer int32 `json:"BackupServer,omitempty"`
	// The backup real server group for this group.
	BackupGroup int32 `json:"BackupGroup,omitempty"`
	// The specific content which is examined during health checks.
	// The content depends on the type of health check.
	HealthCheckUrl string `json:"HealthCheckUrl,omitempty"`
	// The OSI layer at which servers are health checked.
	// 		 From version 29.0.0.0 the following values are not supported:
	// 		 snmp2-snmp5, script1-script64.
	HealthCheckLayer SlbCurCfgGroupTableHealthCheckLayer `json:"HealthCheckLayer,omitempty"`
	// The name of the real server group.
	Name string `json:"Name,omitempty"`
	// The minimum number of real servers available. If at any time, the
	// number reaches this minimum limit, a SYSLOG ALERT message is send to
	// to the configured syslog servers stating that the real server
	// threshold has been reached for the concerned group.
	RealThreshold uint32 `json:"RealThreshold,omitempty"`
	// Enable or disable VIP health checking in DSR mode.
	VipHealthCheck SlbCurCfgGroupTableVipHealthCheck `json:"VipHealthCheck,omitempty"`
	// Enable or disable intrusion detection.
	IdsState SlbCurCfgGroupTableIdsState `json:"IdsState,omitempty"`
	// The intrusion detection port. A value of 1 is invalid.
	IdsPort uint64 `json:"IdsPort,omitempty"`
	// Enable or disable intrusion detection group flood.
	IdsFlood SlbCurCfgGroupTableIdsFlood `json:"IdsFlood,omitempty"`
	// 24|32 number of ip bits used for minmisses hash in the
	// current_configuration block.
	MinmissHash SlbCurCfgGroupTableMinmissHash `json:"MinmissHash,omitempty"`
	// IP address mask used by the persistent hash metric.
	PhashMask string `json:"PhashMask,omitempty"`
	// The metric used to select next rport in server.
	Rmetric SlbCurCfgGroupTableRmetric `json:"Rmetric,omitempty"`
	// The formula used to state the actual health of a virtual service.
	// It allows user to use the symbols of '(', ')', '|', '&' to
	// construct a formula to state the health of the server group.
	// This string can take the following formats :
	// '(1&2|3..)', '128' or 'none'
	// Example: Consider a group with 4 reals 1, 2, 3, 4.  An example of
	// a formula to mark the server group failed for the virtual server
	// if 1 and 2 FAILED or 3 and 4 FAILED is (1|2)&(3|4)
	// This formula will mark the virtual server UP, if 1 or 2 and 3 or 4
	// are health checked sucessfully.
	HealthCheckFormula string `json:"HealthCheckFormula,omitempty"`
	// Enable or disable access to this group for operator.
	OperatorAccess SlbCurCfgGroupTableOperatorAccess `json:"OperatorAccess,omitempty"`
	// The Workload Manager for this Group.
	Wlm int32 `json:"Wlm,omitempty"`
	// The group RADIUS authentication string. The string is used for
	// generating encrypted authentication string while doing RADIUS
	// health check for this group radius servers.
	RadiusAuthenString string `json:"RadiusAuthenString,omitempty"`
	// The Secondary backup real server group for this group.
	SecBackupGroup int32 `json:"SecBackupGroup,omitempty"`
	// The slow-start time for this group.
	Slowstart int32 `json:"Slowstart,omitempty"`
	// The minimum threshold value for this group.
	MinThreshold int32 `json:"MinThreshold,omitempty"`
	// The maximum threshold value for this group.
	MaxThreshold int32 `json:"MaxThreshold,omitempty"`
	// The type of real server group IP address.
	IpVer SlbCurCfgGroupTableIpVer `json:"IpVer,omitempty"`
	// The backup real group or real server for this group.
	Backup string `json:"Backup,omitempty"`
	// The Advanced HC ID.
	HealthID string `json:"HealthID,omitempty"`
	// Prefix length used by the persistent hash metric.
	PhashPrefixLength uint32 `json:"PhashPrefixLength,omitempty"`
	// Enable or disable IDS group participation in inspection chain.
	IdsChain SlbCurCfgGroupTableIdsChain `json:"IdsChain,omitempty"`
	// Enable or Disable override maximum connections limit.
	MaxConEx SlbCurCfgGroupTableMaxConEx `json:"MaxConEx,omitempty"`
}

func (p SlbCurCfgGroupTableParams) iMABean() {}
