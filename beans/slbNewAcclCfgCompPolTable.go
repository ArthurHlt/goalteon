package beans

import (
	"fmt"
	"reflect"
)

// SlbNewAcclCfgCompPolTable The table for configuring compression policy.
// Note:This mib is not supported for VX instance of Virtualization.
type SlbNewAcclCfgCompPolTable struct {
	// The compression policy ID(key id) as an index.
	SlbNewAcclCfgCompPolNameIdIndex string
	Params                          *SlbNewAcclCfgCompPolTableParams
}

func NewSlbNewAcclCfgCompPolTableList() *SlbNewAcclCfgCompPolTable {
	return &SlbNewAcclCfgCompPolTable{}
}

func NewSlbNewAcclCfgCompPolTable(
	slbNewAcclCfgCompPolNameIdIndex string,
	params *SlbNewAcclCfgCompPolTableParams,
) *SlbNewAcclCfgCompPolTable {
	return &SlbNewAcclCfgCompPolTable{
		SlbNewAcclCfgCompPolNameIdIndex: slbNewAcclCfgCompPolNameIdIndex,
		Params:                          params,
	}
}

func (c *SlbNewAcclCfgCompPolTable) Name() string {
	return "SlbNewAcclCfgCompPolTable"
}

func (c *SlbNewAcclCfgCompPolTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbNewAcclCfgCompPolTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbNewAcclCfgCompPolTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbNewAcclCfgCompPolNameIdIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbNewAcclCfgCompPolNameIdIndex)
}

type SlbNewAcclCfgCompPolTableAlgrthm int32

const (
	SlbNewAcclCfgCompPolTableAlgrthm_Gzip        SlbNewAcclCfgCompPolTableAlgrthm = 1
	SlbNewAcclCfgCompPolTableAlgrthm_Deflate     SlbNewAcclCfgCompPolTableAlgrthm = 2
	SlbNewAcclCfgCompPolTableAlgrthm_Unsupported SlbNewAcclCfgCompPolTableAlgrthm = 2147483647
)

type SlbNewAcclCfgCompPolTablePreDefBrwsRuleList int32

const (
	SlbNewAcclCfgCompPolTablePreDefBrwsRuleList_Enabled     SlbNewAcclCfgCompPolTablePreDefBrwsRuleList = 1
	SlbNewAcclCfgCompPolTablePreDefBrwsRuleList_Disabled    SlbNewAcclCfgCompPolTablePreDefBrwsRuleList = 2
	SlbNewAcclCfgCompPolTablePreDefBrwsRuleList_Unsupported SlbNewAcclCfgCompPolTablePreDefBrwsRuleList = 2147483647
)

type SlbNewAcclCfgCompPolTableCompsrv int32

const (
	SlbNewAcclCfgCompPolTableCompsrv_Enabled     SlbNewAcclCfgCompPolTableCompsrv = 1
	SlbNewAcclCfgCompPolTableCompsrv_Disabled    SlbNewAcclCfgCompPolTableCompsrv = 2
	SlbNewAcclCfgCompPolTableCompsrv_Unsupported SlbNewAcclCfgCompPolTableCompsrv = 2147483647
)

type SlbNewAcclCfgCompPolTableAdminStatus int32

const (
	SlbNewAcclCfgCompPolTableAdminStatus_Enabled     SlbNewAcclCfgCompPolTableAdminStatus = 1
	SlbNewAcclCfgCompPolTableAdminStatus_Disabled    SlbNewAcclCfgCompPolTableAdminStatus = 2
	SlbNewAcclCfgCompPolTableAdminStatus_Unsupported SlbNewAcclCfgCompPolTableAdminStatus = 2147483647
)

type SlbNewAcclCfgCompPolTableDelete int32

const (
	SlbNewAcclCfgCompPolTableDelete_Other       SlbNewAcclCfgCompPolTableDelete = 1
	SlbNewAcclCfgCompPolTableDelete_Delete      SlbNewAcclCfgCompPolTableDelete = 2
	SlbNewAcclCfgCompPolTableDelete_Unsupported SlbNewAcclCfgCompPolTableDelete = 2147483647
)

type SlbNewAcclCfgCompPolTableParams struct {
	// The compression policy ID(key id) as an index.
	NameIdIndex string `json:"NameIdIndex,omitempty"`
	// Compression policy name.
	Name string `json:"Name,omitempty"`
	// Preferred compression algorithm.
	Algrthm SlbNewAcclCfgCompPolTableAlgrthm `json:"Algrthm,omitempty"`
	// Compression level.
	Complv1 uint64 `json:"Complv1,omitempty"`
	// Minimum file size to be compressed [Byte].
	MinSize uint64 `json:"MinSize,omitempty"`
	// Maximum file size to be compressed [Byte], Values: 1-2000000000, 0-Unlimited
	MaxSize int32 `json:"MaxSize,omitempty"`
	// URL list which is associated with this policy.
	URLList string `json:"URLList,omitempty"`
	// Browser list which is associated with this policy.
	BrwsList string `json:"BrwsList,omitempty"`
	// Enable/Disable predefined browser exceptions rule-list.
	PreDefBrwsRuleList SlbNewAcclCfgCompPolTablePreDefBrwsRuleList `json:"PreDefBrwsRuleList,omitempty"`
	// Enable/Disable compression by real server.
	Compsrv SlbNewAcclCfgCompPolTableCompsrv `json:"Compsrv,omitempty"`
	// Status (enable/disable) of compression policy.
	AdminStatus SlbNewAcclCfgCompPolTableAdminStatus `json:"AdminStatus,omitempty"`
	// Delete compression policy.
	Delete SlbNewAcclCfgCompPolTableDelete `json:"Delete,omitempty"`
}

func (p SlbNewAcclCfgCompPolTableParams) iMABean() {}
