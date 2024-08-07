package beans

import (
	"fmt"
	"reflect"
)

// AgLicenseCapacityInfoTable The table of license Capacity Information.
type AgLicenseCapacityInfoTable struct {
	// The capacity license table index.
	LicenseCapacityInfoIdx int32
	Params                 *AgLicenseCapacityInfoTableParams
}

func NewAgLicenseCapacityInfoTableList() *AgLicenseCapacityInfoTable {
	return &AgLicenseCapacityInfoTable{}
}

func NewAgLicenseCapacityInfoTable(
	licenseCapacityInfoIdx int32,
	params *AgLicenseCapacityInfoTableParams,
) *AgLicenseCapacityInfoTable {
	return &AgLicenseCapacityInfoTable{
		LicenseCapacityInfoIdx: licenseCapacityInfoIdx,
		Params:                 params,
	}
}

func (c *AgLicenseCapacityInfoTable) Name() string {
	return "AgLicenseCapacityInfoTable"
}

func (c *AgLicenseCapacityInfoTable) GetParams() BeanType {
	return c.Params
}

func (c *AgLicenseCapacityInfoTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *AgLicenseCapacityInfoTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.LicenseCapacityInfoIdx).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.LicenseCapacityInfoIdx)
}

type AgLicenseCapacityInfoTableParams struct {
	// The capacity license table index.
	LicenseCapacityInfoIdx int32 `json:"licenseCapacityInfoIdx,omitempty"`
	// The available size of license. Value -1 defines unlimited license
	LicenseCapacitySize int32 `json:"licenseCapacitySize,omitempty"`
	// The totally allocated license to vADC's. Value -1 defines unlimited license
	LicenseCapacityAllocated int32 `json:"licenseCapacityAllocated,omitempty"`
	// Capacity License Peak Usage
	LicenseCapacityPeakUsage string `json:"licenseCapacityPeakUsage,omitempty"`
	// Capacity License Current Usage
	LicenseCapacityCurrUsage string `json:"licenseCapacityCurrUsage,omitempty"`
	// Capacity License Peak Usage Collecting Time Stamp (in seconds since the Epoch)
	LicenseCapacityPeakTimeStamp int32 `json:"licenseCapacityPeakTimeStamp,omitempty"`
}

func (p AgLicenseCapacityInfoTableParams) iMABean() {}
