package beans

import (
	"fmt"
	"reflect"
)

// SlbCurCfgEnhGroupTable The table of groups,
// Note:This mib is not supported on VX instance of Virtualization.
type SlbCurCfgEnhGroupTable struct {
	// The group alphanumeric index for which the information pertains.
	SlbCurCfgEnhGroupIndex string
	Params                 *SlbCurCfgEnhGroupTableParams
}

func NewSlbCurCfgEnhGroupTableList() *SlbCurCfgEnhGroupTable {
	return &SlbCurCfgEnhGroupTable{}
}

func NewSlbCurCfgEnhGroupTable(
	slbCurCfgEnhGroupIndex string,
	params *SlbCurCfgEnhGroupTableParams,
) *SlbCurCfgEnhGroupTable {
	return &SlbCurCfgEnhGroupTable{
		SlbCurCfgEnhGroupIndex: slbCurCfgEnhGroupIndex,
		Params:                 params,
	}
}

func (c *SlbCurCfgEnhGroupTable) Name() string {
	return "SlbCurCfgEnhGroupTable"
}

func (c *SlbCurCfgEnhGroupTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbCurCfgEnhGroupTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbCurCfgEnhGroupTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbCurCfgEnhGroupIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbCurCfgEnhGroupIndex)
}

type SlbCurCfgEnhGroupTableMetric int32

const (
	SlbCurCfgEnhGroupTableMetric_RoundRobin       SlbCurCfgEnhGroupTableMetric = 1
	SlbCurCfgEnhGroupTableMetric_LeastConnections SlbCurCfgEnhGroupTableMetric = 2
	SlbCurCfgEnhGroupTableMetric_MinMisses        SlbCurCfgEnhGroupTableMetric = 3
	SlbCurCfgEnhGroupTableMetric_Hash             SlbCurCfgEnhGroupTableMetric = 4
	SlbCurCfgEnhGroupTableMetric_Response         SlbCurCfgEnhGroupTableMetric = 5
	SlbCurCfgEnhGroupTableMetric_Bandwidth        SlbCurCfgEnhGroupTableMetric = 6
	SlbCurCfgEnhGroupTableMetric_Phash            SlbCurCfgEnhGroupTableMetric = 7
	SlbCurCfgEnhGroupTableMetric_SvcLeast         SlbCurCfgEnhGroupTableMetric = 8
	SlbCurCfgEnhGroupTableMetric_Hrw              SlbCurCfgEnhGroupTableMetric = 9
	SlbCurCfgEnhGroupTableMetric_Unsupported      SlbCurCfgEnhGroupTableMetric = 2147483647
)

type SlbCurCfgEnhGroupTableHealthCheckLayer int32

const (
	SlbCurCfgEnhGroupTableHealthCheckLayer_Icmp        SlbCurCfgEnhGroupTableHealthCheckLayer = 1
	SlbCurCfgEnhGroupTableHealthCheckLayer_Tcp         SlbCurCfgEnhGroupTableHealthCheckLayer = 2
	SlbCurCfgEnhGroupTableHealthCheckLayer_Http        SlbCurCfgEnhGroupTableHealthCheckLayer = 3
	SlbCurCfgEnhGroupTableHealthCheckLayer_Dns         SlbCurCfgEnhGroupTableHealthCheckLayer = 4
	SlbCurCfgEnhGroupTableHealthCheckLayer_Smtp        SlbCurCfgEnhGroupTableHealthCheckLayer = 5
	SlbCurCfgEnhGroupTableHealthCheckLayer_Pop3        SlbCurCfgEnhGroupTableHealthCheckLayer = 6
	SlbCurCfgEnhGroupTableHealthCheckLayer_Nntp        SlbCurCfgEnhGroupTableHealthCheckLayer = 7
	SlbCurCfgEnhGroupTableHealthCheckLayer_Ftp         SlbCurCfgEnhGroupTableHealthCheckLayer = 8
	SlbCurCfgEnhGroupTableHealthCheckLayer_Imap        SlbCurCfgEnhGroupTableHealthCheckLayer = 9
	SlbCurCfgEnhGroupTableHealthCheckLayer_Radius      SlbCurCfgEnhGroupTableHealthCheckLayer = 10
	SlbCurCfgEnhGroupTableHealthCheckLayer_Sslh        SlbCurCfgEnhGroupTableHealthCheckLayer = 11
	SlbCurCfgEnhGroupTableHealthCheckLayer_Script1     SlbCurCfgEnhGroupTableHealthCheckLayer = 12
	SlbCurCfgEnhGroupTableHealthCheckLayer_Script2     SlbCurCfgEnhGroupTableHealthCheckLayer = 13
	SlbCurCfgEnhGroupTableHealthCheckLayer_Script3     SlbCurCfgEnhGroupTableHealthCheckLayer = 14
	SlbCurCfgEnhGroupTableHealthCheckLayer_Script4     SlbCurCfgEnhGroupTableHealthCheckLayer = 15
	SlbCurCfgEnhGroupTableHealthCheckLayer_Script5     SlbCurCfgEnhGroupTableHealthCheckLayer = 16
	SlbCurCfgEnhGroupTableHealthCheckLayer_Script6     SlbCurCfgEnhGroupTableHealthCheckLayer = 17
	SlbCurCfgEnhGroupTableHealthCheckLayer_Script7     SlbCurCfgEnhGroupTableHealthCheckLayer = 18
	SlbCurCfgEnhGroupTableHealthCheckLayer_Script8     SlbCurCfgEnhGroupTableHealthCheckLayer = 19
	SlbCurCfgEnhGroupTableHealthCheckLayer_Script9     SlbCurCfgEnhGroupTableHealthCheckLayer = 20
	SlbCurCfgEnhGroupTableHealthCheckLayer_Script10    SlbCurCfgEnhGroupTableHealthCheckLayer = 21
	SlbCurCfgEnhGroupTableHealthCheckLayer_Script11    SlbCurCfgEnhGroupTableHealthCheckLayer = 22
	SlbCurCfgEnhGroupTableHealthCheckLayer_Script12    SlbCurCfgEnhGroupTableHealthCheckLayer = 23
	SlbCurCfgEnhGroupTableHealthCheckLayer_Script13    SlbCurCfgEnhGroupTableHealthCheckLayer = 24
	SlbCurCfgEnhGroupTableHealthCheckLayer_Script14    SlbCurCfgEnhGroupTableHealthCheckLayer = 25
	SlbCurCfgEnhGroupTableHealthCheckLayer_Script15    SlbCurCfgEnhGroupTableHealthCheckLayer = 26
	SlbCurCfgEnhGroupTableHealthCheckLayer_Script16    SlbCurCfgEnhGroupTableHealthCheckLayer = 27
	SlbCurCfgEnhGroupTableHealthCheckLayer_Link        SlbCurCfgEnhGroupTableHealthCheckLayer = 28
	SlbCurCfgEnhGroupTableHealthCheckLayer_Wsp         SlbCurCfgEnhGroupTableHealthCheckLayer = 29
	SlbCurCfgEnhGroupTableHealthCheckLayer_Wtls        SlbCurCfgEnhGroupTableHealthCheckLayer = 30
	SlbCurCfgEnhGroupTableHealthCheckLayer_Ldap        SlbCurCfgEnhGroupTableHealthCheckLayer = 31
	SlbCurCfgEnhGroupTableHealthCheckLayer_Udpdns      SlbCurCfgEnhGroupTableHealthCheckLayer = 32
	SlbCurCfgEnhGroupTableHealthCheckLayer_Arp         SlbCurCfgEnhGroupTableHealthCheckLayer = 33
	SlbCurCfgEnhGroupTableHealthCheckLayer_Snmp1       SlbCurCfgEnhGroupTableHealthCheckLayer = 34
	SlbCurCfgEnhGroupTableHealthCheckLayer_Snmp2       SlbCurCfgEnhGroupTableHealthCheckLayer = 35
	SlbCurCfgEnhGroupTableHealthCheckLayer_Snmp3       SlbCurCfgEnhGroupTableHealthCheckLayer = 36
	SlbCurCfgEnhGroupTableHealthCheckLayer_Snmp4       SlbCurCfgEnhGroupTableHealthCheckLayer = 37
	SlbCurCfgEnhGroupTableHealthCheckLayer_Snmp5       SlbCurCfgEnhGroupTableHealthCheckLayer = 38
	SlbCurCfgEnhGroupTableHealthCheckLayer_Radiusacs   SlbCurCfgEnhGroupTableHealthCheckLayer = 39
	SlbCurCfgEnhGroupTableHealthCheckLayer_Tftp        SlbCurCfgEnhGroupTableHealthCheckLayer = 40
	SlbCurCfgEnhGroupTableHealthCheckLayer_Wtp         SlbCurCfgEnhGroupTableHealthCheckLayer = 41
	SlbCurCfgEnhGroupTableHealthCheckLayer_Rtsp        SlbCurCfgEnhGroupTableHealthCheckLayer = 42
	SlbCurCfgEnhGroupTableHealthCheckLayer_Sipping     SlbCurCfgEnhGroupTableHealthCheckLayer = 43
	SlbCurCfgEnhGroupTableHealthCheckLayer_Httphead    SlbCurCfgEnhGroupTableHealthCheckLayer = 44
	SlbCurCfgEnhGroupTableHealthCheckLayer_Sipoptions  SlbCurCfgEnhGroupTableHealthCheckLayer = 45
	SlbCurCfgEnhGroupTableHealthCheckLayer_Wts         SlbCurCfgEnhGroupTableHealthCheckLayer = 46
	SlbCurCfgEnhGroupTableHealthCheckLayer_Dhcp        SlbCurCfgEnhGroupTableHealthCheckLayer = 47
	SlbCurCfgEnhGroupTableHealthCheckLayer_Radiusaa    SlbCurCfgEnhGroupTableHealthCheckLayer = 48
	SlbCurCfgEnhGroupTableHealthCheckLayer_Sslv3       SlbCurCfgEnhGroupTableHealthCheckLayer = 49
	SlbCurCfgEnhGroupTableHealthCheckLayer_Script17    SlbCurCfgEnhGroupTableHealthCheckLayer = 116
	SlbCurCfgEnhGroupTableHealthCheckLayer_Script18    SlbCurCfgEnhGroupTableHealthCheckLayer = 117
	SlbCurCfgEnhGroupTableHealthCheckLayer_Script19    SlbCurCfgEnhGroupTableHealthCheckLayer = 118
	SlbCurCfgEnhGroupTableHealthCheckLayer_Script20    SlbCurCfgEnhGroupTableHealthCheckLayer = 119
	SlbCurCfgEnhGroupTableHealthCheckLayer_Script21    SlbCurCfgEnhGroupTableHealthCheckLayer = 120
	SlbCurCfgEnhGroupTableHealthCheckLayer_Script22    SlbCurCfgEnhGroupTableHealthCheckLayer = 121
	SlbCurCfgEnhGroupTableHealthCheckLayer_Script23    SlbCurCfgEnhGroupTableHealthCheckLayer = 122
	SlbCurCfgEnhGroupTableHealthCheckLayer_Script24    SlbCurCfgEnhGroupTableHealthCheckLayer = 123
	SlbCurCfgEnhGroupTableHealthCheckLayer_Script25    SlbCurCfgEnhGroupTableHealthCheckLayer = 124
	SlbCurCfgEnhGroupTableHealthCheckLayer_Script26    SlbCurCfgEnhGroupTableHealthCheckLayer = 125
	SlbCurCfgEnhGroupTableHealthCheckLayer_Script27    SlbCurCfgEnhGroupTableHealthCheckLayer = 126
	SlbCurCfgEnhGroupTableHealthCheckLayer_Script28    SlbCurCfgEnhGroupTableHealthCheckLayer = 127
	SlbCurCfgEnhGroupTableHealthCheckLayer_Script29    SlbCurCfgEnhGroupTableHealthCheckLayer = 128
	SlbCurCfgEnhGroupTableHealthCheckLayer_Script30    SlbCurCfgEnhGroupTableHealthCheckLayer = 129
	SlbCurCfgEnhGroupTableHealthCheckLayer_Script31    SlbCurCfgEnhGroupTableHealthCheckLayer = 130
	SlbCurCfgEnhGroupTableHealthCheckLayer_Script32    SlbCurCfgEnhGroupTableHealthCheckLayer = 131
	SlbCurCfgEnhGroupTableHealthCheckLayer_Script33    SlbCurCfgEnhGroupTableHealthCheckLayer = 132
	SlbCurCfgEnhGroupTableHealthCheckLayer_Script34    SlbCurCfgEnhGroupTableHealthCheckLayer = 133
	SlbCurCfgEnhGroupTableHealthCheckLayer_Script35    SlbCurCfgEnhGroupTableHealthCheckLayer = 134
	SlbCurCfgEnhGroupTableHealthCheckLayer_Script36    SlbCurCfgEnhGroupTableHealthCheckLayer = 135
	SlbCurCfgEnhGroupTableHealthCheckLayer_Script37    SlbCurCfgEnhGroupTableHealthCheckLayer = 136
	SlbCurCfgEnhGroupTableHealthCheckLayer_Script38    SlbCurCfgEnhGroupTableHealthCheckLayer = 137
	SlbCurCfgEnhGroupTableHealthCheckLayer_Script39    SlbCurCfgEnhGroupTableHealthCheckLayer = 138
	SlbCurCfgEnhGroupTableHealthCheckLayer_Script40    SlbCurCfgEnhGroupTableHealthCheckLayer = 139
	SlbCurCfgEnhGroupTableHealthCheckLayer_Script41    SlbCurCfgEnhGroupTableHealthCheckLayer = 140
	SlbCurCfgEnhGroupTableHealthCheckLayer_Script42    SlbCurCfgEnhGroupTableHealthCheckLayer = 141
	SlbCurCfgEnhGroupTableHealthCheckLayer_Script43    SlbCurCfgEnhGroupTableHealthCheckLayer = 142
	SlbCurCfgEnhGroupTableHealthCheckLayer_Script44    SlbCurCfgEnhGroupTableHealthCheckLayer = 143
	SlbCurCfgEnhGroupTableHealthCheckLayer_Script45    SlbCurCfgEnhGroupTableHealthCheckLayer = 144
	SlbCurCfgEnhGroupTableHealthCheckLayer_Script46    SlbCurCfgEnhGroupTableHealthCheckLayer = 145
	SlbCurCfgEnhGroupTableHealthCheckLayer_Script47    SlbCurCfgEnhGroupTableHealthCheckLayer = 146
	SlbCurCfgEnhGroupTableHealthCheckLayer_Script48    SlbCurCfgEnhGroupTableHealthCheckLayer = 147
	SlbCurCfgEnhGroupTableHealthCheckLayer_Script49    SlbCurCfgEnhGroupTableHealthCheckLayer = 148
	SlbCurCfgEnhGroupTableHealthCheckLayer_Script50    SlbCurCfgEnhGroupTableHealthCheckLayer = 149
	SlbCurCfgEnhGroupTableHealthCheckLayer_Script51    SlbCurCfgEnhGroupTableHealthCheckLayer = 150
	SlbCurCfgEnhGroupTableHealthCheckLayer_Script52    SlbCurCfgEnhGroupTableHealthCheckLayer = 151
	SlbCurCfgEnhGroupTableHealthCheckLayer_Script53    SlbCurCfgEnhGroupTableHealthCheckLayer = 152
	SlbCurCfgEnhGroupTableHealthCheckLayer_Script54    SlbCurCfgEnhGroupTableHealthCheckLayer = 153
	SlbCurCfgEnhGroupTableHealthCheckLayer_Script55    SlbCurCfgEnhGroupTableHealthCheckLayer = 154
	SlbCurCfgEnhGroupTableHealthCheckLayer_Script56    SlbCurCfgEnhGroupTableHealthCheckLayer = 155
	SlbCurCfgEnhGroupTableHealthCheckLayer_Script57    SlbCurCfgEnhGroupTableHealthCheckLayer = 156
	SlbCurCfgEnhGroupTableHealthCheckLayer_Script58    SlbCurCfgEnhGroupTableHealthCheckLayer = 157
	SlbCurCfgEnhGroupTableHealthCheckLayer_Script59    SlbCurCfgEnhGroupTableHealthCheckLayer = 158
	SlbCurCfgEnhGroupTableHealthCheckLayer_Script60    SlbCurCfgEnhGroupTableHealthCheckLayer = 159
	SlbCurCfgEnhGroupTableHealthCheckLayer_Script61    SlbCurCfgEnhGroupTableHealthCheckLayer = 160
	SlbCurCfgEnhGroupTableHealthCheckLayer_Script62    SlbCurCfgEnhGroupTableHealthCheckLayer = 161
	SlbCurCfgEnhGroupTableHealthCheckLayer_Script63    SlbCurCfgEnhGroupTableHealthCheckLayer = 162
	SlbCurCfgEnhGroupTableHealthCheckLayer_Script64    SlbCurCfgEnhGroupTableHealthCheckLayer = 163
	SlbCurCfgEnhGroupTableHealthCheckLayer_None        SlbCurCfgEnhGroupTableHealthCheckLayer = 164
	SlbCurCfgEnhGroupTableHealthCheckLayer_Unknown     SlbCurCfgEnhGroupTableHealthCheckLayer = 165
	SlbCurCfgEnhGroupTableHealthCheckLayer_Unsupported SlbCurCfgEnhGroupTableHealthCheckLayer = 2147483647
)

type SlbCurCfgEnhGroupTableVipHealthCheck int32

const (
	SlbCurCfgEnhGroupTableVipHealthCheck_Enabled     SlbCurCfgEnhGroupTableVipHealthCheck = 1
	SlbCurCfgEnhGroupTableVipHealthCheck_Disabled    SlbCurCfgEnhGroupTableVipHealthCheck = 2
	SlbCurCfgEnhGroupTableVipHealthCheck_Unsupported SlbCurCfgEnhGroupTableVipHealthCheck = 2147483647
)

type SlbCurCfgEnhGroupTableIdsState int32

const (
	SlbCurCfgEnhGroupTableIdsState_Enabled     SlbCurCfgEnhGroupTableIdsState = 1
	SlbCurCfgEnhGroupTableIdsState_Disabled    SlbCurCfgEnhGroupTableIdsState = 2
	SlbCurCfgEnhGroupTableIdsState_Unsupported SlbCurCfgEnhGroupTableIdsState = 2147483647
)

type SlbCurCfgEnhGroupTableIdsFlood int32

const (
	SlbCurCfgEnhGroupTableIdsFlood_Enabled     SlbCurCfgEnhGroupTableIdsFlood = 1
	SlbCurCfgEnhGroupTableIdsFlood_Disabled    SlbCurCfgEnhGroupTableIdsFlood = 2
	SlbCurCfgEnhGroupTableIdsFlood_Unsupported SlbCurCfgEnhGroupTableIdsFlood = 2147483647
)

type SlbCurCfgEnhGroupTableMinmissHash int32

const (
	SlbCurCfgEnhGroupTableMinmissHash_Minmiss24   SlbCurCfgEnhGroupTableMinmissHash = 1
	SlbCurCfgEnhGroupTableMinmissHash_Minmiss32   SlbCurCfgEnhGroupTableMinmissHash = 2
	SlbCurCfgEnhGroupTableMinmissHash_Unsupported SlbCurCfgEnhGroupTableMinmissHash = 2147483647
)

type SlbCurCfgEnhGroupTableRmetric int32

const (
	SlbCurCfgEnhGroupTableRmetric_RoundRobin       SlbCurCfgEnhGroupTableRmetric = 1
	SlbCurCfgEnhGroupTableRmetric_Hash             SlbCurCfgEnhGroupTableRmetric = 2
	SlbCurCfgEnhGroupTableRmetric_LeastConnections SlbCurCfgEnhGroupTableRmetric = 3
	SlbCurCfgEnhGroupTableRmetric_Unsupported      SlbCurCfgEnhGroupTableRmetric = 2147483647
)

type SlbCurCfgEnhGroupTableOperatorAccess int32

const (
	SlbCurCfgEnhGroupTableOperatorAccess_Enabled     SlbCurCfgEnhGroupTableOperatorAccess = 1
	SlbCurCfgEnhGroupTableOperatorAccess_Disabled    SlbCurCfgEnhGroupTableOperatorAccess = 2
	SlbCurCfgEnhGroupTableOperatorAccess_Unsupported SlbCurCfgEnhGroupTableOperatorAccess = 2147483647
)

type SlbCurCfgEnhGroupTableIpVer int32

const (
	SlbCurCfgEnhGroupTableIpVer_Ipv4        SlbCurCfgEnhGroupTableIpVer = 1
	SlbCurCfgEnhGroupTableIpVer_Ipv6        SlbCurCfgEnhGroupTableIpVer = 2
	SlbCurCfgEnhGroupTableIpVer_Mixed       SlbCurCfgEnhGroupTableIpVer = 3
	SlbCurCfgEnhGroupTableIpVer_Unsupported SlbCurCfgEnhGroupTableIpVer = 2147483647
)

type SlbCurCfgEnhGroupTableType int32

const (
	SlbCurCfgEnhGroupTableType_Local       SlbCurCfgEnhGroupTableType = 0
	SlbCurCfgEnhGroupTableType_Wanlink     SlbCurCfgEnhGroupTableType = 2
	SlbCurCfgEnhGroupTableType_Unsupported SlbCurCfgEnhGroupTableType = 2147483647
)

type SlbCurCfgEnhGroupTableIdsChain int32

const (
	SlbCurCfgEnhGroupTableIdsChain_Enabled     SlbCurCfgEnhGroupTableIdsChain = 1
	SlbCurCfgEnhGroupTableIdsChain_Disabled    SlbCurCfgEnhGroupTableIdsChain = 2
	SlbCurCfgEnhGroupTableIdsChain_Unsupported SlbCurCfgEnhGroupTableIdsChain = 2147483647
)

type SlbCurCfgEnhGroupTableSecType int32

const (
	SlbCurCfgEnhGroupTableSecType_None        SlbCurCfgEnhGroupTableSecType = 1
	SlbCurCfgEnhGroupTableSecType_Virtual     SlbCurCfgEnhGroupTableSecType = 2
	SlbCurCfgEnhGroupTableSecType_Layer       SlbCurCfgEnhGroupTableSecType = 3
	SlbCurCfgEnhGroupTableSecType_Passive     SlbCurCfgEnhGroupTableSecType = 4
	SlbCurCfgEnhGroupTableSecType_L3sw        SlbCurCfgEnhGroupTableSecType = 5
	SlbCurCfgEnhGroupTableSecType_Unsupported SlbCurCfgEnhGroupTableSecType = 2147483647
)

type SlbCurCfgEnhGroupTableSecDeviceFlag int32

const (
	SlbCurCfgEnhGroupTableSecDeviceFlag_None        SlbCurCfgEnhGroupTableSecDeviceFlag = 1
	SlbCurCfgEnhGroupTableSecDeviceFlag_Security    SlbCurCfgEnhGroupTableSecDeviceFlag = 2
	SlbCurCfgEnhGroupTableSecDeviceFlag_Unsupported SlbCurCfgEnhGroupTableSecDeviceFlag = 2147483647
)

type SlbCurCfgEnhGroupTableMaxConEx int32

const (
	SlbCurCfgEnhGroupTableMaxConEx_Enabled     SlbCurCfgEnhGroupTableMaxConEx = 1
	SlbCurCfgEnhGroupTableMaxConEx_Disabled    SlbCurCfgEnhGroupTableMaxConEx = 2
	SlbCurCfgEnhGroupTableMaxConEx_Unsupported SlbCurCfgEnhGroupTableMaxConEx = 2147483647
)

type SlbCurCfgEnhGroupTableParams struct {
	// The group alphanumeric index for which the information pertains.
	Index string `json:"Index,omitempty"`
	// The metric used to select next server in group.
	Metric SlbCurCfgEnhGroupTableMetric `json:"Metric,omitempty"`
	// The backup real server for this group.
	BackupServer string `json:"BackupServer,omitempty"`
	// The backup real server group for this group.
	BackupGroup string `json:"BackupGroup,omitempty"`
	// The specific content which is examined during health checks.
	// The content depends on the type of health check.
	HealthCheckUrl string `json:"HealthCheckUrl,omitempty"`
	// The OSI layer at which servers are health checked.
	// From version 29.0.0.0 the following values are not supported:
	// snmp2-snmp5, script1-script64.
	HealthCheckLayer SlbCurCfgEnhGroupTableHealthCheckLayer `json:"HealthCheckLayer,omitempty"`
	// The name of the real server group.
	Name string `json:"Name,omitempty"`
	// The minimum number of real servers available. If at any time, the
	// number reaches this minimum limit, a SYSLOG ALERT message is send to
	// to the configured syslog servers stating that the real server
	// threshold has been reached for the concerned group.
	RealThreshold uint32 `json:"RealThreshold,omitempty"`
	// Enable or disable VIP health checking in DSR mode.
	VipHealthCheck SlbCurCfgEnhGroupTableVipHealthCheck `json:"VipHealthCheck,omitempty"`
	// Enable or disable intrusion detection.
	IdsState SlbCurCfgEnhGroupTableIdsState `json:"IdsState,omitempty"`
	// The intrusion detection port. A value of 1 is invalid.
	IdsPort uint64 `json:"IdsPort,omitempty"`
	// Enable or disable intrusion detection group flood.
	IdsFlood SlbCurCfgEnhGroupTableIdsFlood `json:"IdsFlood,omitempty"`
	// 24|32 number of ip bits used for minmisses hash in the
	// current_configuration block.
	MinmissHash SlbCurCfgEnhGroupTableMinmissHash `json:"MinmissHash,omitempty"`
	// IP address mask used by the persistent hash metric.
	PhashMask string `json:"PhashMask,omitempty"`
	// The metric used to select next rport in server.
	Rmetric SlbCurCfgEnhGroupTableRmetric `json:"Rmetric,omitempty"`
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
	OperatorAccess SlbCurCfgEnhGroupTableOperatorAccess `json:"OperatorAccess,omitempty"`
	// The Workload Manager for this Group.
	Wlm int32 `json:"Wlm,omitempty"`
	// The group RADIUS authentication string. The string is used for
	// generating encrypted authentication string while doing RADIUS
	// health check for this group radius servers.
	RadiusAuthenString string `json:"RadiusAuthenString,omitempty"`
	// The Secondary backup real server group for this group.
	SecBackupGroup string `json:"SecBackupGroup,omitempty"`
	// The slow-start time for this group.
	Slowstart int32 `json:"Slowstart,omitempty"`
	// The minimum threshold value for this group.
	MinThreshold int32 `json:"MinThreshold,omitempty"`
	// The maximum threshold value for this group.
	MaxThreshold int32 `json:"MaxThreshold,omitempty"`
	// The type of real server group IP address.
	IpVer SlbCurCfgEnhGroupTableIpVer `json:"IpVer,omitempty"`
	// The backup real group or real server for this group.
	Backup string `json:"Backup,omitempty"`
	// The Advanced HC ID.
	HealthID string `json:"HealthID,omitempty"`
	// Prefix length used by the persistent hash metric.
	PhashPrefixLength uint32 `json:"PhashPrefixLength,omitempty"`
	// Group type.
	Type SlbCurCfgEnhGroupTableType `json:"Type,omitempty"`
	// Enable or disable IDS group participation in inspection chain.
	IdsChain SlbCurCfgEnhGroupTableIdsChain `json:"IdsChain,omitempty"`
	// The group security device type.
	SecType SlbCurCfgEnhGroupTableSecType `json:"SecType,omitempty"`
	// The Group security device flag.
	SecDeviceFlag SlbCurCfgEnhGroupTableSecDeviceFlag `json:"SecDeviceFlag,omitempty"`
	// Enable or Disable override maximum connections limit.
	MaxConEx SlbCurCfgEnhGroupTableMaxConEx `json:"MaxConEx,omitempty"`
}

func (p SlbCurCfgEnhGroupTableParams) iMABean() {}
