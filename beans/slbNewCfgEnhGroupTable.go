package beans

import (
	"fmt"
	"reflect"
)

// SlbNewCfgEnhGroupTable The table of groups,
// Note:This mib is not supported for VX instance of Virtualization.
type SlbNewCfgEnhGroupTable struct {
	// The group alphanumeric index for which the information pertains.
	SlbNewCfgEnhGroupIndex string
	Params                 *SlbNewCfgEnhGroupTableParams
}

func NewSlbNewCfgEnhGroupTableList() *SlbNewCfgEnhGroupTable {
	return &SlbNewCfgEnhGroupTable{}
}

func NewSlbNewCfgEnhGroupTable(
	slbNewCfgEnhGroupIndex string,
	params *SlbNewCfgEnhGroupTableParams,
) *SlbNewCfgEnhGroupTable {
	return &SlbNewCfgEnhGroupTable{
		SlbNewCfgEnhGroupIndex: slbNewCfgEnhGroupIndex,
		Params:                 params,
	}
}

func (c *SlbNewCfgEnhGroupTable) Name() string {
	return "SlbNewCfgEnhGroupTable"
}

func (c *SlbNewCfgEnhGroupTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbNewCfgEnhGroupTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbNewCfgEnhGroupTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbNewCfgEnhGroupIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbNewCfgEnhGroupIndex)
}

type SlbNewCfgEnhGroupTableMetric int32

const (
	SlbNewCfgEnhGroupTableMetric_RoundRobin       SlbNewCfgEnhGroupTableMetric = 1
	SlbNewCfgEnhGroupTableMetric_LeastConnections SlbNewCfgEnhGroupTableMetric = 2
	SlbNewCfgEnhGroupTableMetric_MinMisses        SlbNewCfgEnhGroupTableMetric = 3
	SlbNewCfgEnhGroupTableMetric_Hash             SlbNewCfgEnhGroupTableMetric = 4
	SlbNewCfgEnhGroupTableMetric_Response         SlbNewCfgEnhGroupTableMetric = 5
	SlbNewCfgEnhGroupTableMetric_Bandwidth        SlbNewCfgEnhGroupTableMetric = 6
	SlbNewCfgEnhGroupTableMetric_Phash            SlbNewCfgEnhGroupTableMetric = 7
	SlbNewCfgEnhGroupTableMetric_SvcLeast         SlbNewCfgEnhGroupTableMetric = 8
	SlbNewCfgEnhGroupTableMetric_Hrw              SlbNewCfgEnhGroupTableMetric = 9
	SlbNewCfgEnhGroupTableMetric_Unsupported      SlbNewCfgEnhGroupTableMetric = 2147483647
)

type SlbNewCfgEnhGroupTableHealthCheckLayer int32

const (
	SlbNewCfgEnhGroupTableHealthCheckLayer_Icmp        SlbNewCfgEnhGroupTableHealthCheckLayer = 1
	SlbNewCfgEnhGroupTableHealthCheckLayer_Tcp         SlbNewCfgEnhGroupTableHealthCheckLayer = 2
	SlbNewCfgEnhGroupTableHealthCheckLayer_Http        SlbNewCfgEnhGroupTableHealthCheckLayer = 3
	SlbNewCfgEnhGroupTableHealthCheckLayer_Dns         SlbNewCfgEnhGroupTableHealthCheckLayer = 4
	SlbNewCfgEnhGroupTableHealthCheckLayer_Smtp        SlbNewCfgEnhGroupTableHealthCheckLayer = 5
	SlbNewCfgEnhGroupTableHealthCheckLayer_Pop3        SlbNewCfgEnhGroupTableHealthCheckLayer = 6
	SlbNewCfgEnhGroupTableHealthCheckLayer_Nntp        SlbNewCfgEnhGroupTableHealthCheckLayer = 7
	SlbNewCfgEnhGroupTableHealthCheckLayer_Ftp         SlbNewCfgEnhGroupTableHealthCheckLayer = 8
	SlbNewCfgEnhGroupTableHealthCheckLayer_Imap        SlbNewCfgEnhGroupTableHealthCheckLayer = 9
	SlbNewCfgEnhGroupTableHealthCheckLayer_Radius      SlbNewCfgEnhGroupTableHealthCheckLayer = 10
	SlbNewCfgEnhGroupTableHealthCheckLayer_Sslh        SlbNewCfgEnhGroupTableHealthCheckLayer = 11
	SlbNewCfgEnhGroupTableHealthCheckLayer_Script1     SlbNewCfgEnhGroupTableHealthCheckLayer = 12
	SlbNewCfgEnhGroupTableHealthCheckLayer_Script2     SlbNewCfgEnhGroupTableHealthCheckLayer = 13
	SlbNewCfgEnhGroupTableHealthCheckLayer_Script3     SlbNewCfgEnhGroupTableHealthCheckLayer = 14
	SlbNewCfgEnhGroupTableHealthCheckLayer_Script4     SlbNewCfgEnhGroupTableHealthCheckLayer = 15
	SlbNewCfgEnhGroupTableHealthCheckLayer_Script5     SlbNewCfgEnhGroupTableHealthCheckLayer = 16
	SlbNewCfgEnhGroupTableHealthCheckLayer_Script6     SlbNewCfgEnhGroupTableHealthCheckLayer = 17
	SlbNewCfgEnhGroupTableHealthCheckLayer_Script7     SlbNewCfgEnhGroupTableHealthCheckLayer = 18
	SlbNewCfgEnhGroupTableHealthCheckLayer_Script8     SlbNewCfgEnhGroupTableHealthCheckLayer = 19
	SlbNewCfgEnhGroupTableHealthCheckLayer_Script9     SlbNewCfgEnhGroupTableHealthCheckLayer = 20
	SlbNewCfgEnhGroupTableHealthCheckLayer_Script10    SlbNewCfgEnhGroupTableHealthCheckLayer = 21
	SlbNewCfgEnhGroupTableHealthCheckLayer_Script11    SlbNewCfgEnhGroupTableHealthCheckLayer = 22
	SlbNewCfgEnhGroupTableHealthCheckLayer_Script12    SlbNewCfgEnhGroupTableHealthCheckLayer = 23
	SlbNewCfgEnhGroupTableHealthCheckLayer_Script13    SlbNewCfgEnhGroupTableHealthCheckLayer = 24
	SlbNewCfgEnhGroupTableHealthCheckLayer_Script14    SlbNewCfgEnhGroupTableHealthCheckLayer = 25
	SlbNewCfgEnhGroupTableHealthCheckLayer_Script15    SlbNewCfgEnhGroupTableHealthCheckLayer = 26
	SlbNewCfgEnhGroupTableHealthCheckLayer_Script16    SlbNewCfgEnhGroupTableHealthCheckLayer = 27
	SlbNewCfgEnhGroupTableHealthCheckLayer_Link        SlbNewCfgEnhGroupTableHealthCheckLayer = 28
	SlbNewCfgEnhGroupTableHealthCheckLayer_Wsp         SlbNewCfgEnhGroupTableHealthCheckLayer = 29
	SlbNewCfgEnhGroupTableHealthCheckLayer_Wtls        SlbNewCfgEnhGroupTableHealthCheckLayer = 30
	SlbNewCfgEnhGroupTableHealthCheckLayer_Ldap        SlbNewCfgEnhGroupTableHealthCheckLayer = 31
	SlbNewCfgEnhGroupTableHealthCheckLayer_Udpdns      SlbNewCfgEnhGroupTableHealthCheckLayer = 32
	SlbNewCfgEnhGroupTableHealthCheckLayer_Arp         SlbNewCfgEnhGroupTableHealthCheckLayer = 33
	SlbNewCfgEnhGroupTableHealthCheckLayer_Snmp1       SlbNewCfgEnhGroupTableHealthCheckLayer = 34
	SlbNewCfgEnhGroupTableHealthCheckLayer_Snmp2       SlbNewCfgEnhGroupTableHealthCheckLayer = 35
	SlbNewCfgEnhGroupTableHealthCheckLayer_Snmp3       SlbNewCfgEnhGroupTableHealthCheckLayer = 36
	SlbNewCfgEnhGroupTableHealthCheckLayer_Snmp4       SlbNewCfgEnhGroupTableHealthCheckLayer = 37
	SlbNewCfgEnhGroupTableHealthCheckLayer_Snmp5       SlbNewCfgEnhGroupTableHealthCheckLayer = 38
	SlbNewCfgEnhGroupTableHealthCheckLayer_Radiusacs   SlbNewCfgEnhGroupTableHealthCheckLayer = 39
	SlbNewCfgEnhGroupTableHealthCheckLayer_Tftp        SlbNewCfgEnhGroupTableHealthCheckLayer = 40
	SlbNewCfgEnhGroupTableHealthCheckLayer_Wtp         SlbNewCfgEnhGroupTableHealthCheckLayer = 41
	SlbNewCfgEnhGroupTableHealthCheckLayer_Rtsp        SlbNewCfgEnhGroupTableHealthCheckLayer = 42
	SlbNewCfgEnhGroupTableHealthCheckLayer_Sipping     SlbNewCfgEnhGroupTableHealthCheckLayer = 43
	SlbNewCfgEnhGroupTableHealthCheckLayer_Httphead    SlbNewCfgEnhGroupTableHealthCheckLayer = 44
	SlbNewCfgEnhGroupTableHealthCheckLayer_Sipoptions  SlbNewCfgEnhGroupTableHealthCheckLayer = 45
	SlbNewCfgEnhGroupTableHealthCheckLayer_Wts         SlbNewCfgEnhGroupTableHealthCheckLayer = 46
	SlbNewCfgEnhGroupTableHealthCheckLayer_Dhcp        SlbNewCfgEnhGroupTableHealthCheckLayer = 47
	SlbNewCfgEnhGroupTableHealthCheckLayer_Radiusaa    SlbNewCfgEnhGroupTableHealthCheckLayer = 48
	SlbNewCfgEnhGroupTableHealthCheckLayer_Sslv3       SlbNewCfgEnhGroupTableHealthCheckLayer = 49
	SlbNewCfgEnhGroupTableHealthCheckLayer_Script17    SlbNewCfgEnhGroupTableHealthCheckLayer = 116
	SlbNewCfgEnhGroupTableHealthCheckLayer_Script18    SlbNewCfgEnhGroupTableHealthCheckLayer = 117
	SlbNewCfgEnhGroupTableHealthCheckLayer_Script19    SlbNewCfgEnhGroupTableHealthCheckLayer = 118
	SlbNewCfgEnhGroupTableHealthCheckLayer_Script20    SlbNewCfgEnhGroupTableHealthCheckLayer = 119
	SlbNewCfgEnhGroupTableHealthCheckLayer_Script21    SlbNewCfgEnhGroupTableHealthCheckLayer = 120
	SlbNewCfgEnhGroupTableHealthCheckLayer_Script22    SlbNewCfgEnhGroupTableHealthCheckLayer = 121
	SlbNewCfgEnhGroupTableHealthCheckLayer_Script23    SlbNewCfgEnhGroupTableHealthCheckLayer = 122
	SlbNewCfgEnhGroupTableHealthCheckLayer_Script24    SlbNewCfgEnhGroupTableHealthCheckLayer = 123
	SlbNewCfgEnhGroupTableHealthCheckLayer_Script25    SlbNewCfgEnhGroupTableHealthCheckLayer = 124
	SlbNewCfgEnhGroupTableHealthCheckLayer_Script26    SlbNewCfgEnhGroupTableHealthCheckLayer = 125
	SlbNewCfgEnhGroupTableHealthCheckLayer_Script27    SlbNewCfgEnhGroupTableHealthCheckLayer = 126
	SlbNewCfgEnhGroupTableHealthCheckLayer_Script28    SlbNewCfgEnhGroupTableHealthCheckLayer = 127
	SlbNewCfgEnhGroupTableHealthCheckLayer_Script29    SlbNewCfgEnhGroupTableHealthCheckLayer = 128
	SlbNewCfgEnhGroupTableHealthCheckLayer_Script30    SlbNewCfgEnhGroupTableHealthCheckLayer = 129
	SlbNewCfgEnhGroupTableHealthCheckLayer_Script31    SlbNewCfgEnhGroupTableHealthCheckLayer = 130
	SlbNewCfgEnhGroupTableHealthCheckLayer_Script32    SlbNewCfgEnhGroupTableHealthCheckLayer = 131
	SlbNewCfgEnhGroupTableHealthCheckLayer_Script33    SlbNewCfgEnhGroupTableHealthCheckLayer = 132
	SlbNewCfgEnhGroupTableHealthCheckLayer_Script34    SlbNewCfgEnhGroupTableHealthCheckLayer = 133
	SlbNewCfgEnhGroupTableHealthCheckLayer_Script35    SlbNewCfgEnhGroupTableHealthCheckLayer = 134
	SlbNewCfgEnhGroupTableHealthCheckLayer_Script36    SlbNewCfgEnhGroupTableHealthCheckLayer = 135
	SlbNewCfgEnhGroupTableHealthCheckLayer_Script37    SlbNewCfgEnhGroupTableHealthCheckLayer = 136
	SlbNewCfgEnhGroupTableHealthCheckLayer_Script38    SlbNewCfgEnhGroupTableHealthCheckLayer = 137
	SlbNewCfgEnhGroupTableHealthCheckLayer_Script39    SlbNewCfgEnhGroupTableHealthCheckLayer = 138
	SlbNewCfgEnhGroupTableHealthCheckLayer_Script40    SlbNewCfgEnhGroupTableHealthCheckLayer = 139
	SlbNewCfgEnhGroupTableHealthCheckLayer_Script41    SlbNewCfgEnhGroupTableHealthCheckLayer = 140
	SlbNewCfgEnhGroupTableHealthCheckLayer_Script42    SlbNewCfgEnhGroupTableHealthCheckLayer = 141
	SlbNewCfgEnhGroupTableHealthCheckLayer_Script43    SlbNewCfgEnhGroupTableHealthCheckLayer = 142
	SlbNewCfgEnhGroupTableHealthCheckLayer_Script44    SlbNewCfgEnhGroupTableHealthCheckLayer = 143
	SlbNewCfgEnhGroupTableHealthCheckLayer_Script45    SlbNewCfgEnhGroupTableHealthCheckLayer = 144
	SlbNewCfgEnhGroupTableHealthCheckLayer_Script46    SlbNewCfgEnhGroupTableHealthCheckLayer = 145
	SlbNewCfgEnhGroupTableHealthCheckLayer_Script47    SlbNewCfgEnhGroupTableHealthCheckLayer = 146
	SlbNewCfgEnhGroupTableHealthCheckLayer_Script48    SlbNewCfgEnhGroupTableHealthCheckLayer = 147
	SlbNewCfgEnhGroupTableHealthCheckLayer_Script49    SlbNewCfgEnhGroupTableHealthCheckLayer = 148
	SlbNewCfgEnhGroupTableHealthCheckLayer_Script50    SlbNewCfgEnhGroupTableHealthCheckLayer = 149
	SlbNewCfgEnhGroupTableHealthCheckLayer_Script51    SlbNewCfgEnhGroupTableHealthCheckLayer = 150
	SlbNewCfgEnhGroupTableHealthCheckLayer_Script52    SlbNewCfgEnhGroupTableHealthCheckLayer = 151
	SlbNewCfgEnhGroupTableHealthCheckLayer_Script53    SlbNewCfgEnhGroupTableHealthCheckLayer = 152
	SlbNewCfgEnhGroupTableHealthCheckLayer_Script54    SlbNewCfgEnhGroupTableHealthCheckLayer = 153
	SlbNewCfgEnhGroupTableHealthCheckLayer_Script55    SlbNewCfgEnhGroupTableHealthCheckLayer = 154
	SlbNewCfgEnhGroupTableHealthCheckLayer_Script56    SlbNewCfgEnhGroupTableHealthCheckLayer = 155
	SlbNewCfgEnhGroupTableHealthCheckLayer_Script57    SlbNewCfgEnhGroupTableHealthCheckLayer = 156
	SlbNewCfgEnhGroupTableHealthCheckLayer_Script58    SlbNewCfgEnhGroupTableHealthCheckLayer = 157
	SlbNewCfgEnhGroupTableHealthCheckLayer_Script59    SlbNewCfgEnhGroupTableHealthCheckLayer = 158
	SlbNewCfgEnhGroupTableHealthCheckLayer_Script60    SlbNewCfgEnhGroupTableHealthCheckLayer = 159
	SlbNewCfgEnhGroupTableHealthCheckLayer_Script61    SlbNewCfgEnhGroupTableHealthCheckLayer = 160
	SlbNewCfgEnhGroupTableHealthCheckLayer_Script62    SlbNewCfgEnhGroupTableHealthCheckLayer = 161
	SlbNewCfgEnhGroupTableHealthCheckLayer_Script63    SlbNewCfgEnhGroupTableHealthCheckLayer = 162
	SlbNewCfgEnhGroupTableHealthCheckLayer_Script64    SlbNewCfgEnhGroupTableHealthCheckLayer = 163
	SlbNewCfgEnhGroupTableHealthCheckLayer_None        SlbNewCfgEnhGroupTableHealthCheckLayer = 164
	SlbNewCfgEnhGroupTableHealthCheckLayer_Unknown     SlbNewCfgEnhGroupTableHealthCheckLayer = 165
	SlbNewCfgEnhGroupTableHealthCheckLayer_Unsupported SlbNewCfgEnhGroupTableHealthCheckLayer = 2147483647
)

type SlbNewCfgEnhGroupTableVipHealthCheck int32

const (
	SlbNewCfgEnhGroupTableVipHealthCheck_Enabled     SlbNewCfgEnhGroupTableVipHealthCheck = 1
	SlbNewCfgEnhGroupTableVipHealthCheck_Disabled    SlbNewCfgEnhGroupTableVipHealthCheck = 2
	SlbNewCfgEnhGroupTableVipHealthCheck_Unsupported SlbNewCfgEnhGroupTableVipHealthCheck = 2147483647
)

type SlbNewCfgEnhGroupTableIdsState int32

const (
	SlbNewCfgEnhGroupTableIdsState_Enabled     SlbNewCfgEnhGroupTableIdsState = 1
	SlbNewCfgEnhGroupTableIdsState_Disabled    SlbNewCfgEnhGroupTableIdsState = 2
	SlbNewCfgEnhGroupTableIdsState_Unsupported SlbNewCfgEnhGroupTableIdsState = 2147483647
)

type SlbNewCfgEnhGroupTableDelete int32

const (
	SlbNewCfgEnhGroupTableDelete_Other       SlbNewCfgEnhGroupTableDelete = 1
	SlbNewCfgEnhGroupTableDelete_Delete      SlbNewCfgEnhGroupTableDelete = 2
	SlbNewCfgEnhGroupTableDelete_Unsupported SlbNewCfgEnhGroupTableDelete = 2147483647
)

type SlbNewCfgEnhGroupTableIdsFlood int32

const (
	SlbNewCfgEnhGroupTableIdsFlood_Enabled     SlbNewCfgEnhGroupTableIdsFlood = 1
	SlbNewCfgEnhGroupTableIdsFlood_Disabled    SlbNewCfgEnhGroupTableIdsFlood = 2
	SlbNewCfgEnhGroupTableIdsFlood_Unsupported SlbNewCfgEnhGroupTableIdsFlood = 2147483647
)

type SlbNewCfgEnhGroupTableMinmissHash int32

const (
	SlbNewCfgEnhGroupTableMinmissHash_Minmiss24   SlbNewCfgEnhGroupTableMinmissHash = 1
	SlbNewCfgEnhGroupTableMinmissHash_Minmiss32   SlbNewCfgEnhGroupTableMinmissHash = 2
	SlbNewCfgEnhGroupTableMinmissHash_Unsupported SlbNewCfgEnhGroupTableMinmissHash = 2147483647
)

type SlbNewCfgEnhGroupTableRmetric int32

const (
	SlbNewCfgEnhGroupTableRmetric_RoundRobin       SlbNewCfgEnhGroupTableRmetric = 1
	SlbNewCfgEnhGroupTableRmetric_Hash             SlbNewCfgEnhGroupTableRmetric = 2
	SlbNewCfgEnhGroupTableRmetric_LeastConnections SlbNewCfgEnhGroupTableRmetric = 3
	SlbNewCfgEnhGroupTableRmetric_Unsupported      SlbNewCfgEnhGroupTableRmetric = 2147483647
)

type SlbNewCfgEnhGroupTableOperatorAccess int32

const (
	SlbNewCfgEnhGroupTableOperatorAccess_Enabled     SlbNewCfgEnhGroupTableOperatorAccess = 1
	SlbNewCfgEnhGroupTableOperatorAccess_Disabled    SlbNewCfgEnhGroupTableOperatorAccess = 2
	SlbNewCfgEnhGroupTableOperatorAccess_Unsupported SlbNewCfgEnhGroupTableOperatorAccess = 2147483647
)

type SlbNewCfgEnhGroupTableIpVer int32

const (
	SlbNewCfgEnhGroupTableIpVer_Ipv4        SlbNewCfgEnhGroupTableIpVer = 1
	SlbNewCfgEnhGroupTableIpVer_Ipv6        SlbNewCfgEnhGroupTableIpVer = 2
	SlbNewCfgEnhGroupTableIpVer_Mixed       SlbNewCfgEnhGroupTableIpVer = 3
	SlbNewCfgEnhGroupTableIpVer_Unsupported SlbNewCfgEnhGroupTableIpVer = 2147483647
)

type SlbNewCfgEnhGroupTableBackupType int32

const (
	SlbNewCfgEnhGroupTableBackupType_None        SlbNewCfgEnhGroupTableBackupType = 1
	SlbNewCfgEnhGroupTableBackupType_Server      SlbNewCfgEnhGroupTableBackupType = 2
	SlbNewCfgEnhGroupTableBackupType_Group       SlbNewCfgEnhGroupTableBackupType = 3
	SlbNewCfgEnhGroupTableBackupType_Unsupported SlbNewCfgEnhGroupTableBackupType = 2147483647
)

type SlbNewCfgEnhGroupTableType int32

const (
	SlbNewCfgEnhGroupTableType_Local       SlbNewCfgEnhGroupTableType = 0
	SlbNewCfgEnhGroupTableType_Wanlink     SlbNewCfgEnhGroupTableType = 2
	SlbNewCfgEnhGroupTableType_Unsupported SlbNewCfgEnhGroupTableType = 2147483647
)

type SlbNewCfgEnhGroupTableIdsChain int32

const (
	SlbNewCfgEnhGroupTableIdsChain_Enabled     SlbNewCfgEnhGroupTableIdsChain = 1
	SlbNewCfgEnhGroupTableIdsChain_Disabled    SlbNewCfgEnhGroupTableIdsChain = 2
	SlbNewCfgEnhGroupTableIdsChain_Unsupported SlbNewCfgEnhGroupTableIdsChain = 2147483647
)

type SlbNewCfgEnhGroupTableSecType int32

const (
	SlbNewCfgEnhGroupTableSecType_None        SlbNewCfgEnhGroupTableSecType = 1
	SlbNewCfgEnhGroupTableSecType_Virtual     SlbNewCfgEnhGroupTableSecType = 2
	SlbNewCfgEnhGroupTableSecType_Layer       SlbNewCfgEnhGroupTableSecType = 3
	SlbNewCfgEnhGroupTableSecType_Passive     SlbNewCfgEnhGroupTableSecType = 4
	SlbNewCfgEnhGroupTableSecType_L3sw        SlbNewCfgEnhGroupTableSecType = 5
	SlbNewCfgEnhGroupTableSecType_Unsupported SlbNewCfgEnhGroupTableSecType = 2147483647
)

type SlbNewCfgEnhGroupTableSecDeviceFlag int32

const (
	SlbNewCfgEnhGroupTableSecDeviceFlag_None        SlbNewCfgEnhGroupTableSecDeviceFlag = 1
	SlbNewCfgEnhGroupTableSecDeviceFlag_Security    SlbNewCfgEnhGroupTableSecDeviceFlag = 2
	SlbNewCfgEnhGroupTableSecDeviceFlag_Unsupported SlbNewCfgEnhGroupTableSecDeviceFlag = 2147483647
)

type SlbNewCfgEnhGroupTableMaxConEx int32

const (
	SlbNewCfgEnhGroupTableMaxConEx_Enabled     SlbNewCfgEnhGroupTableMaxConEx = 1
	SlbNewCfgEnhGroupTableMaxConEx_Disabled    SlbNewCfgEnhGroupTableMaxConEx = 2
	SlbNewCfgEnhGroupTableMaxConEx_Unsupported SlbNewCfgEnhGroupTableMaxConEx = 2147483647
)

type SlbNewCfgEnhGroupTableParams struct {
	// The group alphanumeric index for which the information pertains.
	Index string `json:"Index,omitempty"`
	// The real server to be added to the group. When read, 0 is returned.
	AddServer string `json:"AddServer,omitempty"`
	// The real server to be removed from the group. When read, 0 is returned.
	RemoveServer string `json:"RemoveServer,omitempty"`
	// The metric used to select next server in group.
	Metric SlbNewCfgEnhGroupTableMetric `json:"Metric,omitempty"`
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
	HealthCheckLayer SlbNewCfgEnhGroupTableHealthCheckLayer `json:"HealthCheckLayer,omitempty"`
	// The name of the real server group.
	Name string `json:"Name,omitempty"`
	// The minimum number of real servers available. If at any time, the
	// number reaches this minimum limit, a SYSLOG ALERT message is send to
	// to the configured syslog servers stating that the real server
	// threshold has been reached for the concerned group.
	RealThreshold uint32 `json:"RealThreshold,omitempty"`
	// Enable or disable VIP health checking in DSR mode.
	VipHealthCheck SlbNewCfgEnhGroupTableVipHealthCheck `json:"VipHealthCheck,omitempty"`
	// Enable or disable intrusion detection.
	IdsState SlbNewCfgEnhGroupTableIdsState `json:"IdsState,omitempty"`
	// The intrusion detection port. A value of 1 is invalid.
	IdsPort uint64 `json:"IdsPort,omitempty"`
	// By setting the value to delete(2), the entire group is deleted.
	Delete SlbNewCfgEnhGroupTableDelete `json:"Delete,omitempty"`
	// Enable or disable intrusion detection group flood.
	IdsFlood SlbNewCfgEnhGroupTableIdsFlood `json:"IdsFlood,omitempty"`
	// 24|32 number of sip bits used for minmisses hash in the
	// new_configuration block.
	MinmissHash SlbNewCfgEnhGroupTableMinmissHash `json:"MinmissHash,omitempty"`
	// IP address mask used by the persistent hash metric.
	PhashMask string `json:"PhashMask,omitempty"`
	// The metric used to select next rport in server.
	Rmetric SlbNewCfgEnhGroupTableRmetric `json:"Rmetric,omitempty"`
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
	OperatorAccess SlbNewCfgEnhGroupTableOperatorAccess `json:"OperatorAccess,omitempty"`
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
	IpVer SlbNewCfgEnhGroupTableIpVer `json:"IpVer,omitempty"`
	// The backup real group or real server for this group.
	Backup string `json:"Backup,omitempty"`
	// Backup type of the real server group.
	BackupType SlbNewCfgEnhGroupTableBackupType `json:"BackupType,omitempty"`
	// The Advanced HC ID.
	HealthID string `json:"HealthID,omitempty"`
	// Prefix length used by the persistent hash metric.
	PhashPrefixLength uint32 `json:"PhashPrefixLength,omitempty"`
	// Group type.
	Type SlbNewCfgEnhGroupTableType `json:"Type,omitempty"`
	// The alphanumeric index of the new copy to be created.
	Copy string `json:"Copy,omitempty"`
	// Enable or disable IDS group participation in inspection chain.
	IdsChain SlbNewCfgEnhGroupTableIdsChain `json:"IdsChain,omitempty"`
	// The Group security device type.
	SecType SlbNewCfgEnhGroupTableSecType `json:"SecType,omitempty"`
	// The Group security device flag
	SecDeviceFlag SlbNewCfgEnhGroupTableSecDeviceFlag `json:"SecDeviceFlag,omitempty"`
	// Enable or Disable override maximum connections limit.
	MaxConEx SlbNewCfgEnhGroupTableMaxConEx `json:"MaxConEx,omitempty"`
}

func (p SlbNewCfgEnhGroupTableParams) iMABean() {}
