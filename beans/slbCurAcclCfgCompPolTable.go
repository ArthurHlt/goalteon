package beans

import (
	"fmt"
	"reflect"
)

// SlbCurAcclCfgCompPolTable The table for configuring compression policy.
// Note:This mib is not supported for VX instance of Virtualization.
type SlbCurAcclCfgCompPolTable struct {
	// The compression policy ID(key id) as an index.
	SlbCurAcclCfgCompPolNameIdIndex string
	Params                          *SlbCurAcclCfgCompPolTableParams
}

func NewSlbCurAcclCfgCompPolTableList() *SlbCurAcclCfgCompPolTable {
	return &SlbCurAcclCfgCompPolTable{}
}

func NewSlbCurAcclCfgCompPolTable(
	slbCurAcclCfgCompPolNameIdIndex string,
	params *SlbCurAcclCfgCompPolTableParams,
) *SlbCurAcclCfgCompPolTable {
	return &SlbCurAcclCfgCompPolTable{
		SlbCurAcclCfgCompPolNameIdIndex: slbCurAcclCfgCompPolNameIdIndex,
		Params:                          params,
	}
}

func (c *SlbCurAcclCfgCompPolTable) Name() string {
	return "SlbCurAcclCfgCompPolTable"
}

func (c *SlbCurAcclCfgCompPolTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbCurAcclCfgCompPolTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbCurAcclCfgCompPolTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbCurAcclCfgCompPolNameIdIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbCurAcclCfgCompPolNameIdIndex)
}

type SlbCurAcclCfgCompPolTableAlgrthm int32

const (
	SlbCurAcclCfgCompPolTableAlgrthm_Gzip        SlbCurAcclCfgCompPolTableAlgrthm = 1
	SlbCurAcclCfgCompPolTableAlgrthm_Deflate     SlbCurAcclCfgCompPolTableAlgrthm = 2
	SlbCurAcclCfgCompPolTableAlgrthm_Unsupported SlbCurAcclCfgCompPolTableAlgrthm = 2147483647
)

type SlbCurAcclCfgCompPolTablePreDefBrwsRuleList int32

const (
	SlbCurAcclCfgCompPolTablePreDefBrwsRuleList_Enabled     SlbCurAcclCfgCompPolTablePreDefBrwsRuleList = 1
	SlbCurAcclCfgCompPolTablePreDefBrwsRuleList_Disabled    SlbCurAcclCfgCompPolTablePreDefBrwsRuleList = 2
	SlbCurAcclCfgCompPolTablePreDefBrwsRuleList_Unsupported SlbCurAcclCfgCompPolTablePreDefBrwsRuleList = 2147483647
)

type SlbCurAcclCfgCompPolTableCompsrv int32

const (
	SlbCurAcclCfgCompPolTableCompsrv_Enabled     SlbCurAcclCfgCompPolTableCompsrv = 1
	SlbCurAcclCfgCompPolTableCompsrv_Disabled    SlbCurAcclCfgCompPolTableCompsrv = 2
	SlbCurAcclCfgCompPolTableCompsrv_Unsupported SlbCurAcclCfgCompPolTableCompsrv = 2147483647
)

type SlbCurAcclCfgCompPolTableAdminStatus int32

const (
	SlbCurAcclCfgCompPolTableAdminStatus_Enabled     SlbCurAcclCfgCompPolTableAdminStatus = 1
	SlbCurAcclCfgCompPolTableAdminStatus_Disabled    SlbCurAcclCfgCompPolTableAdminStatus = 2
	SlbCurAcclCfgCompPolTableAdminStatus_Unsupported SlbCurAcclCfgCompPolTableAdminStatus = 2147483647
)

type SlbCurAcclCfgCompPolTableParams struct {
	// The compression policy ID(key id) as an index.
	NameIdIndex string `json:"NameIdIndex,omitempty"`
	// Compression policy name.
	Name string `json:"Name,omitempty"`
	// Preferred compression algorithm.
	Algrthm SlbCurAcclCfgCompPolTableAlgrthm `json:"Algrthm,omitempty"`
	// Compression level.
	Complv1 uint64 `json:"Complv1,omitempty"`
	// Minimum file size to be compressed [Byte].
	MinSize uint64 `json:"MinSize,omitempty"`
	// Maximum file size to be compressed [Byte].
	MaxSize int32 `json:"MaxSize,omitempty"`
	// URL list which is associated with this policy.
	URLList string `json:"URLList,omitempty"`
	// Browser list which is associated with this policy.
	BrwsList string `json:"BrwsList,omitempty"`
	// Enable/Disable predefined browser exceptions rule-list.
	PreDefBrwsRuleList SlbCurAcclCfgCompPolTablePreDefBrwsRuleList `json:"PreDefBrwsRuleList,omitempty"`
	// Enable/Disable compression by real server.
	Compsrv SlbCurAcclCfgCompPolTableCompsrv `json:"Compsrv,omitempty"`
	// Status (enable/disable) of compression policy.
	AdminStatus SlbCurAcclCfgCompPolTableAdminStatus `json:"AdminStatus,omitempty"`
}

func (p SlbCurAcclCfgCompPolTableParams) iMABean() {}
