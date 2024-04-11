package beans

import (
	"fmt"
	"reflect"
)

// SlbNewCfgGroupTable The table of groups,
// Note:This mib is not supported on VX instance of Virtualization.
type SlbNewCfgGroupTable struct {
	// The group number for which the information pertains.
	SlbNewCfgGroupIndex int32
	Params              *SlbNewCfgGroupTableParams
}

func NewSlbNewCfgGroupTableList() *SlbNewCfgGroupTable {
	return &SlbNewCfgGroupTable{}
}

func NewSlbNewCfgGroupTable(
	slbNewCfgGroupIndex int32,
	params *SlbNewCfgGroupTableParams,
) *SlbNewCfgGroupTable {
	return &SlbNewCfgGroupTable{
		SlbNewCfgGroupIndex: slbNewCfgGroupIndex,
		Params:              params,
	}
}

func (c *SlbNewCfgGroupTable) Name() string {
	return "SlbNewCfgGroupTable"
}

func (c *SlbNewCfgGroupTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbNewCfgGroupTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbNewCfgGroupTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbNewCfgGroupIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbNewCfgGroupIndex)
}

type SlbNewCfgGroupTableMetric int32

const (
	SlbNewCfgGroupTableMetric_RoundRobin       SlbNewCfgGroupTableMetric = 1
	SlbNewCfgGroupTableMetric_LeastConnections SlbNewCfgGroupTableMetric = 2
	SlbNewCfgGroupTableMetric_MinMisses        SlbNewCfgGroupTableMetric = 3
	SlbNewCfgGroupTableMetric_Hash             SlbNewCfgGroupTableMetric = 4
	SlbNewCfgGroupTableMetric_Response         SlbNewCfgGroupTableMetric = 5
	SlbNewCfgGroupTableMetric_Bandwidth        SlbNewCfgGroupTableMetric = 6
	SlbNewCfgGroupTableMetric_Phash            SlbNewCfgGroupTableMetric = 7
	SlbNewCfgGroupTableMetric_SvcLeast         SlbNewCfgGroupTableMetric = 8
	SlbNewCfgGroupTableMetric_Unsupported      SlbNewCfgGroupTableMetric = 2147483647
)

type SlbNewCfgGroupTableHealthCheckLayer int32

const (
	SlbNewCfgGroupTableHealthCheckLayer_Icmp        SlbNewCfgGroupTableHealthCheckLayer = 1
	SlbNewCfgGroupTableHealthCheckLayer_Tcp         SlbNewCfgGroupTableHealthCheckLayer = 2
	SlbNewCfgGroupTableHealthCheckLayer_Http        SlbNewCfgGroupTableHealthCheckLayer = 3
	SlbNewCfgGroupTableHealthCheckLayer_Dns         SlbNewCfgGroupTableHealthCheckLayer = 4
	SlbNewCfgGroupTableHealthCheckLayer_Smtp        SlbNewCfgGroupTableHealthCheckLayer = 5
	SlbNewCfgGroupTableHealthCheckLayer_Pop3        SlbNewCfgGroupTableHealthCheckLayer = 6
	SlbNewCfgGroupTableHealthCheckLayer_Nntp        SlbNewCfgGroupTableHealthCheckLayer = 7
	SlbNewCfgGroupTableHealthCheckLayer_Ftp         SlbNewCfgGroupTableHealthCheckLayer = 8
	SlbNewCfgGroupTableHealthCheckLayer_Imap        SlbNewCfgGroupTableHealthCheckLayer = 9
	SlbNewCfgGroupTableHealthCheckLayer_Radius      SlbNewCfgGroupTableHealthCheckLayer = 10
	SlbNewCfgGroupTableHealthCheckLayer_Sslh        SlbNewCfgGroupTableHealthCheckLayer = 11
	SlbNewCfgGroupTableHealthCheckLayer_Script1     SlbNewCfgGroupTableHealthCheckLayer = 12
	SlbNewCfgGroupTableHealthCheckLayer_Script2     SlbNewCfgGroupTableHealthCheckLayer = 13
	SlbNewCfgGroupTableHealthCheckLayer_Script3     SlbNewCfgGroupTableHealthCheckLayer = 14
	SlbNewCfgGroupTableHealthCheckLayer_Script4     SlbNewCfgGroupTableHealthCheckLayer = 15
	SlbNewCfgGroupTableHealthCheckLayer_Script5     SlbNewCfgGroupTableHealthCheckLayer = 16
	SlbNewCfgGroupTableHealthCheckLayer_Script6     SlbNewCfgGroupTableHealthCheckLayer = 17
	SlbNewCfgGroupTableHealthCheckLayer_Script7     SlbNewCfgGroupTableHealthCheckLayer = 18
	SlbNewCfgGroupTableHealthCheckLayer_Script8     SlbNewCfgGroupTableHealthCheckLayer = 19
	SlbNewCfgGroupTableHealthCheckLayer_Script9     SlbNewCfgGroupTableHealthCheckLayer = 20
	SlbNewCfgGroupTableHealthCheckLayer_Script10    SlbNewCfgGroupTableHealthCheckLayer = 21
	SlbNewCfgGroupTableHealthCheckLayer_Script11    SlbNewCfgGroupTableHealthCheckLayer = 22
	SlbNewCfgGroupTableHealthCheckLayer_Script12    SlbNewCfgGroupTableHealthCheckLayer = 23
	SlbNewCfgGroupTableHealthCheckLayer_Script13    SlbNewCfgGroupTableHealthCheckLayer = 24
	SlbNewCfgGroupTableHealthCheckLayer_Script14    SlbNewCfgGroupTableHealthCheckLayer = 25
	SlbNewCfgGroupTableHealthCheckLayer_Script15    SlbNewCfgGroupTableHealthCheckLayer = 26
	SlbNewCfgGroupTableHealthCheckLayer_Script16    SlbNewCfgGroupTableHealthCheckLayer = 27
	SlbNewCfgGroupTableHealthCheckLayer_Link        SlbNewCfgGroupTableHealthCheckLayer = 28
	SlbNewCfgGroupTableHealthCheckLayer_Wsp         SlbNewCfgGroupTableHealthCheckLayer = 29
	SlbNewCfgGroupTableHealthCheckLayer_Wtls        SlbNewCfgGroupTableHealthCheckLayer = 30
	SlbNewCfgGroupTableHealthCheckLayer_Ldap        SlbNewCfgGroupTableHealthCheckLayer = 31
	SlbNewCfgGroupTableHealthCheckLayer_Udpdns      SlbNewCfgGroupTableHealthCheckLayer = 32
	SlbNewCfgGroupTableHealthCheckLayer_Arp         SlbNewCfgGroupTableHealthCheckLayer = 33
	SlbNewCfgGroupTableHealthCheckLayer_Snmp1       SlbNewCfgGroupTableHealthCheckLayer = 34
	SlbNewCfgGroupTableHealthCheckLayer_Snmp2       SlbNewCfgGroupTableHealthCheckLayer = 35
	SlbNewCfgGroupTableHealthCheckLayer_Snmp3       SlbNewCfgGroupTableHealthCheckLayer = 36
	SlbNewCfgGroupTableHealthCheckLayer_Snmp4       SlbNewCfgGroupTableHealthCheckLayer = 37
	SlbNewCfgGroupTableHealthCheckLayer_Snmp5       SlbNewCfgGroupTableHealthCheckLayer = 38
	SlbNewCfgGroupTableHealthCheckLayer_Radiusacs   SlbNewCfgGroupTableHealthCheckLayer = 39
	SlbNewCfgGroupTableHealthCheckLayer_Tftp        SlbNewCfgGroupTableHealthCheckLayer = 40
	SlbNewCfgGroupTableHealthCheckLayer_Wtp         SlbNewCfgGroupTableHealthCheckLayer = 41
	SlbNewCfgGroupTableHealthCheckLayer_Rtsp        SlbNewCfgGroupTableHealthCheckLayer = 42
	SlbNewCfgGroupTableHealthCheckLayer_Sipping     SlbNewCfgGroupTableHealthCheckLayer = 43
	SlbNewCfgGroupTableHealthCheckLayer_Httphead    SlbNewCfgGroupTableHealthCheckLayer = 44
	SlbNewCfgGroupTableHealthCheckLayer_Sipoptions  SlbNewCfgGroupTableHealthCheckLayer = 45
	SlbNewCfgGroupTableHealthCheckLayer_Wts         SlbNewCfgGroupTableHealthCheckLayer = 46
	SlbNewCfgGroupTableHealthCheckLayer_Dhcp        SlbNewCfgGroupTableHealthCheckLayer = 47
	SlbNewCfgGroupTableHealthCheckLayer_Radiusaa    SlbNewCfgGroupTableHealthCheckLayer = 48
	SlbNewCfgGroupTableHealthCheckLayer_Sslv3       SlbNewCfgGroupTableHealthCheckLayer = 49
	SlbNewCfgGroupTableHealthCheckLayer_Script17    SlbNewCfgGroupTableHealthCheckLayer = 116
	SlbNewCfgGroupTableHealthCheckLayer_Script18    SlbNewCfgGroupTableHealthCheckLayer = 117
	SlbNewCfgGroupTableHealthCheckLayer_Script19    SlbNewCfgGroupTableHealthCheckLayer = 118
	SlbNewCfgGroupTableHealthCheckLayer_Script20    SlbNewCfgGroupTableHealthCheckLayer = 119
	SlbNewCfgGroupTableHealthCheckLayer_Script21    SlbNewCfgGroupTableHealthCheckLayer = 120
	SlbNewCfgGroupTableHealthCheckLayer_Script22    SlbNewCfgGroupTableHealthCheckLayer = 121
	SlbNewCfgGroupTableHealthCheckLayer_Script23    SlbNewCfgGroupTableHealthCheckLayer = 122
	SlbNewCfgGroupTableHealthCheckLayer_Script24    SlbNewCfgGroupTableHealthCheckLayer = 123
	SlbNewCfgGroupTableHealthCheckLayer_Script25    SlbNewCfgGroupTableHealthCheckLayer = 124
	SlbNewCfgGroupTableHealthCheckLayer_Script26    SlbNewCfgGroupTableHealthCheckLayer = 125
	SlbNewCfgGroupTableHealthCheckLayer_Script27    SlbNewCfgGroupTableHealthCheckLayer = 126
	SlbNewCfgGroupTableHealthCheckLayer_Script28    SlbNewCfgGroupTableHealthCheckLayer = 127
	SlbNewCfgGroupTableHealthCheckLayer_Script29    SlbNewCfgGroupTableHealthCheckLayer = 128
	SlbNewCfgGroupTableHealthCheckLayer_Script30    SlbNewCfgGroupTableHealthCheckLayer = 129
	SlbNewCfgGroupTableHealthCheckLayer_Script31    SlbNewCfgGroupTableHealthCheckLayer = 130
	SlbNewCfgGroupTableHealthCheckLayer_Script32    SlbNewCfgGroupTableHealthCheckLayer = 131
	SlbNewCfgGroupTableHealthCheckLayer_Script33    SlbNewCfgGroupTableHealthCheckLayer = 132
	SlbNewCfgGroupTableHealthCheckLayer_Script34    SlbNewCfgGroupTableHealthCheckLayer = 133
	SlbNewCfgGroupTableHealthCheckLayer_Script35    SlbNewCfgGroupTableHealthCheckLayer = 134
	SlbNewCfgGroupTableHealthCheckLayer_Script36    SlbNewCfgGroupTableHealthCheckLayer = 135
	SlbNewCfgGroupTableHealthCheckLayer_Script37    SlbNewCfgGroupTableHealthCheckLayer = 136
	SlbNewCfgGroupTableHealthCheckLayer_Script38    SlbNewCfgGroupTableHealthCheckLayer = 137
	SlbNewCfgGroupTableHealthCheckLayer_Script39    SlbNewCfgGroupTableHealthCheckLayer = 138
	SlbNewCfgGroupTableHealthCheckLayer_Script40    SlbNewCfgGroupTableHealthCheckLayer = 139
	SlbNewCfgGroupTableHealthCheckLayer_Script41    SlbNewCfgGroupTableHealthCheckLayer = 140
	SlbNewCfgGroupTableHealthCheckLayer_Script42    SlbNewCfgGroupTableHealthCheckLayer = 141
	SlbNewCfgGroupTableHealthCheckLayer_Script43    SlbNewCfgGroupTableHealthCheckLayer = 142
	SlbNewCfgGroupTableHealthCheckLayer_Script44    SlbNewCfgGroupTableHealthCheckLayer = 143
	SlbNewCfgGroupTableHealthCheckLayer_Script45    SlbNewCfgGroupTableHealthCheckLayer = 144
	SlbNewCfgGroupTableHealthCheckLayer_Script46    SlbNewCfgGroupTableHealthCheckLayer = 145
	SlbNewCfgGroupTableHealthCheckLayer_Script47    SlbNewCfgGroupTableHealthCheckLayer = 146
	SlbNewCfgGroupTableHealthCheckLayer_Script48    SlbNewCfgGroupTableHealthCheckLayer = 147
	SlbNewCfgGroupTableHealthCheckLayer_Script49    SlbNewCfgGroupTableHealthCheckLayer = 148
	SlbNewCfgGroupTableHealthCheckLayer_Script50    SlbNewCfgGroupTableHealthCheckLayer = 149
	SlbNewCfgGroupTableHealthCheckLayer_Script51    SlbNewCfgGroupTableHealthCheckLayer = 150
	SlbNewCfgGroupTableHealthCheckLayer_Script52    SlbNewCfgGroupTableHealthCheckLayer = 151
	SlbNewCfgGroupTableHealthCheckLayer_Script53    SlbNewCfgGroupTableHealthCheckLayer = 152
	SlbNewCfgGroupTableHealthCheckLayer_Script54    SlbNewCfgGroupTableHealthCheckLayer = 153
	SlbNewCfgGroupTableHealthCheckLayer_Script55    SlbNewCfgGroupTableHealthCheckLayer = 154
	SlbNewCfgGroupTableHealthCheckLayer_Script56    SlbNewCfgGroupTableHealthCheckLayer = 155
	SlbNewCfgGroupTableHealthCheckLayer_Script57    SlbNewCfgGroupTableHealthCheckLayer = 156
	SlbNewCfgGroupTableHealthCheckLayer_Script58    SlbNewCfgGroupTableHealthCheckLayer = 157
	SlbNewCfgGroupTableHealthCheckLayer_Script59    SlbNewCfgGroupTableHealthCheckLayer = 158
	SlbNewCfgGroupTableHealthCheckLayer_Script60    SlbNewCfgGroupTableHealthCheckLayer = 159
	SlbNewCfgGroupTableHealthCheckLayer_Script61    SlbNewCfgGroupTableHealthCheckLayer = 160
	SlbNewCfgGroupTableHealthCheckLayer_Script62    SlbNewCfgGroupTableHealthCheckLayer = 161
	SlbNewCfgGroupTableHealthCheckLayer_Script63    SlbNewCfgGroupTableHealthCheckLayer = 162
	SlbNewCfgGroupTableHealthCheckLayer_Script64    SlbNewCfgGroupTableHealthCheckLayer = 163
	SlbNewCfgGroupTableHealthCheckLayer_None        SlbNewCfgGroupTableHealthCheckLayer = 164
	SlbNewCfgGroupTableHealthCheckLayer_Unknown     SlbNewCfgGroupTableHealthCheckLayer = 165
	SlbNewCfgGroupTableHealthCheckLayer_Unsupported SlbNewCfgGroupTableHealthCheckLayer = 2147483647
)

type SlbNewCfgGroupTableVipHealthCheck int32

const (
	SlbNewCfgGroupTableVipHealthCheck_Enabled     SlbNewCfgGroupTableVipHealthCheck = 1
	SlbNewCfgGroupTableVipHealthCheck_Disabled    SlbNewCfgGroupTableVipHealthCheck = 2
	SlbNewCfgGroupTableVipHealthCheck_Unsupported SlbNewCfgGroupTableVipHealthCheck = 2147483647
)

type SlbNewCfgGroupTableIdsState int32

const (
	SlbNewCfgGroupTableIdsState_Enabled     SlbNewCfgGroupTableIdsState = 1
	SlbNewCfgGroupTableIdsState_Disabled    SlbNewCfgGroupTableIdsState = 2
	SlbNewCfgGroupTableIdsState_Unsupported SlbNewCfgGroupTableIdsState = 2147483647
)

type SlbNewCfgGroupTableDelete int32

const (
	SlbNewCfgGroupTableDelete_Other       SlbNewCfgGroupTableDelete = 1
	SlbNewCfgGroupTableDelete_Delete      SlbNewCfgGroupTableDelete = 2
	SlbNewCfgGroupTableDelete_Unsupported SlbNewCfgGroupTableDelete = 2147483647
)

type SlbNewCfgGroupTableIdsFlood int32

const (
	SlbNewCfgGroupTableIdsFlood_Enabled     SlbNewCfgGroupTableIdsFlood = 1
	SlbNewCfgGroupTableIdsFlood_Disabled    SlbNewCfgGroupTableIdsFlood = 2
	SlbNewCfgGroupTableIdsFlood_Unsupported SlbNewCfgGroupTableIdsFlood = 2147483647
)

type SlbNewCfgGroupTableMinmissHash int32

const (
	SlbNewCfgGroupTableMinmissHash_Minmiss24   SlbNewCfgGroupTableMinmissHash = 1
	SlbNewCfgGroupTableMinmissHash_Minmiss32   SlbNewCfgGroupTableMinmissHash = 2
	SlbNewCfgGroupTableMinmissHash_Unsupported SlbNewCfgGroupTableMinmissHash = 2147483647
)

type SlbNewCfgGroupTableRmetric int32

const (
	SlbNewCfgGroupTableRmetric_RoundRobin       SlbNewCfgGroupTableRmetric = 1
	SlbNewCfgGroupTableRmetric_Hash             SlbNewCfgGroupTableRmetric = 2
	SlbNewCfgGroupTableRmetric_LeastConnections SlbNewCfgGroupTableRmetric = 3
	SlbNewCfgGroupTableRmetric_Unsupported      SlbNewCfgGroupTableRmetric = 2147483647
)

type SlbNewCfgGroupTableOperatorAccess int32

const (
	SlbNewCfgGroupTableOperatorAccess_Enabled     SlbNewCfgGroupTableOperatorAccess = 1
	SlbNewCfgGroupTableOperatorAccess_Disabled    SlbNewCfgGroupTableOperatorAccess = 2
	SlbNewCfgGroupTableOperatorAccess_Unsupported SlbNewCfgGroupTableOperatorAccess = 2147483647
)

type SlbNewCfgGroupTableIpVer int32

const (
	SlbNewCfgGroupTableIpVer_Ipv4        SlbNewCfgGroupTableIpVer = 1
	SlbNewCfgGroupTableIpVer_Ipv6        SlbNewCfgGroupTableIpVer = 2
	SlbNewCfgGroupTableIpVer_Mixed       SlbNewCfgGroupTableIpVer = 3
	SlbNewCfgGroupTableIpVer_Unsupported SlbNewCfgGroupTableIpVer = 2147483647
)

type SlbNewCfgGroupTableBackupType int32

const (
	SlbNewCfgGroupTableBackupType_None        SlbNewCfgGroupTableBackupType = 1
	SlbNewCfgGroupTableBackupType_Server      SlbNewCfgGroupTableBackupType = 2
	SlbNewCfgGroupTableBackupType_Group       SlbNewCfgGroupTableBackupType = 3
	SlbNewCfgGroupTableBackupType_Unsupported SlbNewCfgGroupTableBackupType = 2147483647
)

type SlbNewCfgGroupTableIdsChain int32

const (
	SlbNewCfgGroupTableIdsChain_Enabled     SlbNewCfgGroupTableIdsChain = 1
	SlbNewCfgGroupTableIdsChain_Disabled    SlbNewCfgGroupTableIdsChain = 2
	SlbNewCfgGroupTableIdsChain_Unsupported SlbNewCfgGroupTableIdsChain = 2147483647
)

type SlbNewCfgGroupTableMaxConEx int32

const (
	SlbNewCfgGroupTableMaxConEx_Enabled     SlbNewCfgGroupTableMaxConEx = 1
	SlbNewCfgGroupTableMaxConEx_Disabled    SlbNewCfgGroupTableMaxConEx = 2
	SlbNewCfgGroupTableMaxConEx_Unsupported SlbNewCfgGroupTableMaxConEx = 2147483647
)

type SlbNewCfgGroupTableParams struct {
	// The group number for which the information pertains.
	Index int32 `json:"Index,omitempty"`
	// The Real servers in the group.  The servers are presented
	// in bitmap format.
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
	// The real server to be added to the group. When read, 0 is returned.
	AddServer int32 `json:"AddServer,omitempty"`
	// The real server to be removed from the group. When read, 0 is returned.
	RemoveServer int32 `json:"RemoveServer,omitempty"`
	// The metric used to select next server in group.
	Metric SlbNewCfgGroupTableMetric `json:"Metric,omitempty"`
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
	HealthCheckLayer SlbNewCfgGroupTableHealthCheckLayer `json:"HealthCheckLayer,omitempty"`
	// The name of the real server group.
	Name string `json:"Name,omitempty"`
	// The minimum number of real servers available. If at any time, the
	// number reaches this minimum limit, a SYSLOG ALERT message is send to
	// to the configured syslog servers stating that the real server
	// threshold has been reached for the concerned group.
	RealThreshold uint32 `json:"RealThreshold,omitempty"`
	// Enable or disable VIP health checking in DSR mode.
	VipHealthCheck SlbNewCfgGroupTableVipHealthCheck `json:"VipHealthCheck,omitempty"`
	// Enable or disable intrusion detection.
	IdsState SlbNewCfgGroupTableIdsState `json:"IdsState,omitempty"`
	// The intrusion detection port. A value of 1 is invalid.
	IdsPort uint64 `json:"IdsPort,omitempty"`
	// By setting the value to delete(2), the entire group is deleted.
	Delete SlbNewCfgGroupTableDelete `json:"Delete,omitempty"`
	// Enable or disable intrusion detection group flood.
	IdsFlood SlbNewCfgGroupTableIdsFlood `json:"IdsFlood,omitempty"`
	// 24|32 number of sip bits used for minmisses hash in the
	// new_configuration block.
	MinmissHash SlbNewCfgGroupTableMinmissHash `json:"MinmissHash,omitempty"`
	// IP address mask used by the persistent hash metric.
	PhashMask string `json:"PhashMask,omitempty"`
	// The metric used to select next rport in server.
	Rmetric SlbNewCfgGroupTableRmetric `json:"Rmetric,omitempty"`
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
	OperatorAccess SlbNewCfgGroupTableOperatorAccess `json:"OperatorAccess,omitempty"`
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
	IpVer SlbNewCfgGroupTableIpVer `json:"IpVer,omitempty"`
	// The backup real group or real server for this group.
	Backup string `json:"Backup,omitempty"`
	// Backup type of the real server group.
	BackupType SlbNewCfgGroupTableBackupType `json:"BackupType,omitempty"`
	// The Advanced HC ID.
	HealthID string `json:"HealthID,omitempty"`
	// Prefix length used by the persistent hash metric.
	PhashPrefixLength uint32 `json:"PhashPrefixLength,omitempty"`
	// Enable or disable IDS group participation in inspection chain.
	IdsChain SlbNewCfgGroupTableIdsChain `json:"IdsChain,omitempty"`
	// Enable or Disable override maximum connections limit.
	MaxConEx SlbNewCfgGroupTableMaxConEx `json:"MaxConEx,omitempty"`
}

func (p SlbNewCfgGroupTableParams) iMABean() {}
