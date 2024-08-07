package beans

import (
	"fmt"
	"reflect"
)

// AgLicenseInfoTable The table of license Information.
type AgLicenseInfoTable struct {
	// The license index.
	LicenseInfoIdx int32
	Params         *AgLicenseInfoTableParams
}

func NewAgLicenseInfoTableList() *AgLicenseInfoTable {
	return &AgLicenseInfoTable{}
}

func NewAgLicenseInfoTable(
	licenseInfoIdx int32,
	params *AgLicenseInfoTableParams,
) *AgLicenseInfoTable {
	return &AgLicenseInfoTable{
		LicenseInfoIdx: licenseInfoIdx,
		Params:         params,
	}
}

func (c *AgLicenseInfoTable) Name() string {
	return "AgLicenseInfoTable"
}

func (c *AgLicenseInfoTable) GetParams() BeanType {
	return c.Params
}

func (c *AgLicenseInfoTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *AgLicenseInfoTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.LicenseInfoIdx).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.LicenseInfoIdx)
}

type AgLicenseInfoTableLicenseType int32

const (
	AgLicenseInfoTableLicenseType_Permanent    AgLicenseInfoTableLicenseType = 1
	AgLicenseInfoTableLicenseType_Temporary    AgLicenseInfoTableLicenseType = 2
	AgLicenseInfoTableLicenseType_Removed      AgLicenseInfoTableLicenseType = 3
	AgLicenseInfoTableLicenseType_Expired      AgLicenseInfoTableLicenseType = 4
	AgLicenseInfoTableLicenseType_Subscription AgLicenseInfoTableLicenseType = 5
	AgLicenseInfoTableLicenseType_Unsupported  AgLicenseInfoTableLicenseType = 2147483647
)

type AgLicenseInfoTableParams struct {
	// The license index.
	LicenseInfoIdx int32 `json:"licenseInfoIdx,omitempty"`
	// The Sofware Key.
	SoftwareKey string `json:"softwareKey,omitempty"`
	// The License Type information.
	LicenseType AgLicenseInfoTableLicenseType `json:"licenseType,omitempty"`
	// The number of days left for temporary license.
	RemainingDays int32 `json:"remainingDays,omitempty"`
	// The number of days/hours left for temporary license.
	RemainingDaysLeft string `json:"remainingDaysLeft,omitempty"`
	// The available size of license. Value -1 defines unlimited license.
	LicenseSize int32 `json:"licenseSize,omitempty"`
	// The totally allocated license to vADC's. Value -1 defines unlimited license
	LicenseAllocated int32 `json:"licenseAllocated,omitempty"`
	// The time based license status
	TimeBasedLicenseStatus string `json:"timeBasedLicenseStatus,omitempty"`
}

func (p AgLicenseInfoTableParams) iMABean() {}
