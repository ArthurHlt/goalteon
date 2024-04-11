package beans

import (
	"fmt"
	"reflect"
)

// SlbNewCfgEnhVirtServicesThirdPartTable The table of virtual Services.
// Note:This mib is not supported for VX instance of Virtualization.
type SlbNewCfgEnhVirtServicesThirdPartTable struct {
	// The number of the virtual server.
	SlbNewCfgEnhVirtServThirdPartIndex string
	// The service index. This has no external meaning
	SlbNewCfgEnhVirtServiceThirdPartIndex int32
	Params                                *SlbNewCfgEnhVirtServicesThirdPartTableParams
}

func NewSlbNewCfgEnhVirtServicesThirdPartTableList() *SlbNewCfgEnhVirtServicesThirdPartTable {
	return &SlbNewCfgEnhVirtServicesThirdPartTable{}
}

func NewSlbNewCfgEnhVirtServicesThirdPartTable(
	slbNewCfgEnhVirtServThirdPartIndex string,
	slbNewCfgEnhVirtServiceThirdPartIndex int32,
	params *SlbNewCfgEnhVirtServicesThirdPartTableParams,
) *SlbNewCfgEnhVirtServicesThirdPartTable {
	return &SlbNewCfgEnhVirtServicesThirdPartTable{
		SlbNewCfgEnhVirtServThirdPartIndex:    slbNewCfgEnhVirtServThirdPartIndex,
		SlbNewCfgEnhVirtServiceThirdPartIndex: slbNewCfgEnhVirtServiceThirdPartIndex,
		Params:                                params,
	}
}

func (c *SlbNewCfgEnhVirtServicesThirdPartTable) Name() string {
	return "SlbNewCfgEnhVirtServicesThirdPartTable"
}

func (c *SlbNewCfgEnhVirtServicesThirdPartTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbNewCfgEnhVirtServicesThirdPartTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbNewCfgEnhVirtServicesThirdPartTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbNewCfgEnhVirtServThirdPartIndex).IsZero() &&
		reflect.ValueOf(c.SlbNewCfgEnhVirtServiceThirdPartIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbNewCfgEnhVirtServThirdPartIndex) + "/" + fmt.Sprint(c.SlbNewCfgEnhVirtServiceThirdPartIndex)
}

type SlbNewCfgEnhVirtServicesThirdPartTableServUrlchangPathType int32

const (
	SlbNewCfgEnhVirtServicesThirdPartTableServUrlchangPathType_Sufx        SlbNewCfgEnhVirtServicesThirdPartTableServUrlchangPathType = 1
	SlbNewCfgEnhVirtServicesThirdPartTableServUrlchangPathType_Prefx       SlbNewCfgEnhVirtServicesThirdPartTableServUrlchangPathType = 2
	SlbNewCfgEnhVirtServicesThirdPartTableServUrlchangPathType_Eq          SlbNewCfgEnhVirtServicesThirdPartTableServUrlchangPathType = 3
	SlbNewCfgEnhVirtServicesThirdPartTableServUrlchangPathType_Incl        SlbNewCfgEnhVirtServicesThirdPartTableServUrlchangPathType = 4
	SlbNewCfgEnhVirtServicesThirdPartTableServUrlchangPathType_Any         SlbNewCfgEnhVirtServicesThirdPartTableServUrlchangPathType = 5
	SlbNewCfgEnhVirtServicesThirdPartTableServUrlchangPathType_None        SlbNewCfgEnhVirtServicesThirdPartTableServUrlchangPathType = 6
	SlbNewCfgEnhVirtServicesThirdPartTableServUrlchangPathType_Unsupported SlbNewCfgEnhVirtServicesThirdPartTableServUrlchangPathType = 2147483647
)

type SlbNewCfgEnhVirtServicesThirdPartTableServUrlchangActnType int32

const (
	SlbNewCfgEnhVirtServicesThirdPartTableServUrlchangActnType_Insert      SlbNewCfgEnhVirtServicesThirdPartTableServUrlchangActnType = 1
	SlbNewCfgEnhVirtServicesThirdPartTableServUrlchangActnType_Replace     SlbNewCfgEnhVirtServicesThirdPartTableServUrlchangActnType = 2
	SlbNewCfgEnhVirtServicesThirdPartTableServUrlchangActnType_Remove      SlbNewCfgEnhVirtServicesThirdPartTableServUrlchangActnType = 3
	SlbNewCfgEnhVirtServicesThirdPartTableServUrlchangActnType_None        SlbNewCfgEnhVirtServicesThirdPartTableServUrlchangActnType = 4
	SlbNewCfgEnhVirtServicesThirdPartTableServUrlchangActnType_Unsupported SlbNewCfgEnhVirtServicesThirdPartTableServUrlchangActnType = 2147483647
)

type SlbNewCfgEnhVirtServicesThirdPartTableServUrlchangInsrtPostn int32

const (
	SlbNewCfgEnhVirtServicesThirdPartTableServUrlchangInsrtPostn_Before      SlbNewCfgEnhVirtServicesThirdPartTableServUrlchangInsrtPostn = 1
	SlbNewCfgEnhVirtServicesThirdPartTableServUrlchangInsrtPostn_After       SlbNewCfgEnhVirtServicesThirdPartTableServUrlchangInsrtPostn = 2
	SlbNewCfgEnhVirtServicesThirdPartTableServUrlchangInsrtPostn_Unsupported SlbNewCfgEnhVirtServicesThirdPartTableServUrlchangInsrtPostn = 2147483647
)

type SlbNewCfgEnhVirtServicesThirdPartTableParams struct {
	// The number of the virtual server.
	ServThirdPartIndex string `json:"ServThirdPartIndex,omitempty"`
	// The service index. This has no external meaning
	ThirdPartIndex int32 `json:"ThirdPartIndex,omitempty"`
	// Enter hostname to match.
	ServUrlchangHostName string `json:"ServUrlchangHostName,omitempty"`
	// Enter path match type [sufx|prefx|eq|incl|any|none].
	ServUrlchangPathType SlbNewCfgEnhVirtServicesThirdPartTableServUrlchangPathType `json:"ServUrlchangPathType,omitempty"`
	// Enter path to match.
	ServUrlchangPathMatch string `json:"ServUrlchangPathMatch,omitempty"`
	// Enter page name to match or none.
	ServUrlchangPageName string `json:"ServUrlchangPageName,omitempty"`
	// Enter page type to match or none.
	ServUrlchangPageType string `json:"ServUrlchangPageType,omitempty"`
	// Enter path action type.
	ServUrlchangActnType SlbNewCfgEnhVirtServicesThirdPartTableServUrlchangActnType `json:"ServUrlchangActnType,omitempty"`
	// Enter path to insert.
	ServUrlchangPathInsrt string `json:"ServUrlchangPathInsrt,omitempty"`
	// Insert the specified path before or after the matched section
	ServUrlchangInsrtPostn SlbNewCfgEnhVirtServicesThirdPartTableServUrlchangInsrtPostn `json:"ServUrlchangInsrtPostn,omitempty"`
}

func (p SlbNewCfgEnhVirtServicesThirdPartTableParams) iMABean() {}
