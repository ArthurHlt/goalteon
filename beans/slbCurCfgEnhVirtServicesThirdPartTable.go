package beans

import (
	"fmt"
	"reflect"
)

// SlbCurCfgEnhVirtServicesThirdPartTable The table of virtual Services.
// Note:This mib is not supported for VX instance of Virtualization.
type SlbCurCfgEnhVirtServicesThirdPartTable struct {
	// The number of the virtual server.
	SlbCurCfgEnhVirtServThirdPartIndex string
	// The service index. This has no external meaning
	SlbCurCfgEnhVirtServiceThirdPartIndex int32
	Params                                *SlbCurCfgEnhVirtServicesThirdPartTableParams
}

func NewSlbCurCfgEnhVirtServicesThirdPartTableList() *SlbCurCfgEnhVirtServicesThirdPartTable {
	return &SlbCurCfgEnhVirtServicesThirdPartTable{}
}

func NewSlbCurCfgEnhVirtServicesThirdPartTable(
	slbCurCfgEnhVirtServThirdPartIndex string,
	slbCurCfgEnhVirtServiceThirdPartIndex int32,
	params *SlbCurCfgEnhVirtServicesThirdPartTableParams,
) *SlbCurCfgEnhVirtServicesThirdPartTable {
	return &SlbCurCfgEnhVirtServicesThirdPartTable{
		SlbCurCfgEnhVirtServThirdPartIndex:    slbCurCfgEnhVirtServThirdPartIndex,
		SlbCurCfgEnhVirtServiceThirdPartIndex: slbCurCfgEnhVirtServiceThirdPartIndex,
		Params:                                params,
	}
}

func (c *SlbCurCfgEnhVirtServicesThirdPartTable) Name() string {
	return "SlbCurCfgEnhVirtServicesThirdPartTable"
}

func (c *SlbCurCfgEnhVirtServicesThirdPartTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbCurCfgEnhVirtServicesThirdPartTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbCurCfgEnhVirtServicesThirdPartTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbCurCfgEnhVirtServThirdPartIndex).IsZero() &&
		reflect.ValueOf(c.SlbCurCfgEnhVirtServiceThirdPartIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbCurCfgEnhVirtServThirdPartIndex) + "/" + fmt.Sprint(c.SlbCurCfgEnhVirtServiceThirdPartIndex)
}

type SlbCurCfgEnhVirtServicesThirdPartTableServUrlchangPathType int32

const (
	SlbCurCfgEnhVirtServicesThirdPartTableServUrlchangPathType_Sufx        SlbCurCfgEnhVirtServicesThirdPartTableServUrlchangPathType = 1
	SlbCurCfgEnhVirtServicesThirdPartTableServUrlchangPathType_Prefx       SlbCurCfgEnhVirtServicesThirdPartTableServUrlchangPathType = 2
	SlbCurCfgEnhVirtServicesThirdPartTableServUrlchangPathType_Eq          SlbCurCfgEnhVirtServicesThirdPartTableServUrlchangPathType = 3
	SlbCurCfgEnhVirtServicesThirdPartTableServUrlchangPathType_Incl        SlbCurCfgEnhVirtServicesThirdPartTableServUrlchangPathType = 4
	SlbCurCfgEnhVirtServicesThirdPartTableServUrlchangPathType_Any         SlbCurCfgEnhVirtServicesThirdPartTableServUrlchangPathType = 5
	SlbCurCfgEnhVirtServicesThirdPartTableServUrlchangPathType_None        SlbCurCfgEnhVirtServicesThirdPartTableServUrlchangPathType = 6
	SlbCurCfgEnhVirtServicesThirdPartTableServUrlchangPathType_Unsupported SlbCurCfgEnhVirtServicesThirdPartTableServUrlchangPathType = 2147483647
)

type SlbCurCfgEnhVirtServicesThirdPartTableServUrlchangActnType int32

const (
	SlbCurCfgEnhVirtServicesThirdPartTableServUrlchangActnType_Insert      SlbCurCfgEnhVirtServicesThirdPartTableServUrlchangActnType = 1
	SlbCurCfgEnhVirtServicesThirdPartTableServUrlchangActnType_Replace     SlbCurCfgEnhVirtServicesThirdPartTableServUrlchangActnType = 2
	SlbCurCfgEnhVirtServicesThirdPartTableServUrlchangActnType_Remove      SlbCurCfgEnhVirtServicesThirdPartTableServUrlchangActnType = 3
	SlbCurCfgEnhVirtServicesThirdPartTableServUrlchangActnType_None        SlbCurCfgEnhVirtServicesThirdPartTableServUrlchangActnType = 4
	SlbCurCfgEnhVirtServicesThirdPartTableServUrlchangActnType_Unsupported SlbCurCfgEnhVirtServicesThirdPartTableServUrlchangActnType = 2147483647
)

type SlbCurCfgEnhVirtServicesThirdPartTableServUrlchangInsrtPostn int32

const (
	SlbCurCfgEnhVirtServicesThirdPartTableServUrlchangInsrtPostn_Before      SlbCurCfgEnhVirtServicesThirdPartTableServUrlchangInsrtPostn = 1
	SlbCurCfgEnhVirtServicesThirdPartTableServUrlchangInsrtPostn_After       SlbCurCfgEnhVirtServicesThirdPartTableServUrlchangInsrtPostn = 2
	SlbCurCfgEnhVirtServicesThirdPartTableServUrlchangInsrtPostn_Unsupported SlbCurCfgEnhVirtServicesThirdPartTableServUrlchangInsrtPostn = 2147483647
)

type SlbCurCfgEnhVirtServicesThirdPartTableParams struct {
	// The number of the virtual server.
	ServThirdPartIndex string `json:"ServThirdPartIndex,omitempty"`
	// The service index. This has no external meaning
	ThirdPartIndex int32 `json:"ThirdPartIndex,omitempty"`
	// Enter hostname to match.
	ServUrlchangHostName string `json:"ServUrlchangHostName,omitempty"`
	// Enter path match type [sufx|prefx|eq|incl|any].
	ServUrlchangPathType SlbCurCfgEnhVirtServicesThirdPartTableServUrlchangPathType `json:"ServUrlchangPathType,omitempty"`
	// Enter path to match.
	ServUrlchangPathMatch string `json:"ServUrlchangPathMatch,omitempty"`
	// Enter page name to match or none.
	ServUrlchangPageName string `json:"ServUrlchangPageName,omitempty"`
	// Enter page type to match or none.
	ServUrlchangPageType string `json:"ServUrlchangPageType,omitempty"`
	// Enter path action type.
	ServUrlchangActnType SlbCurCfgEnhVirtServicesThirdPartTableServUrlchangActnType `json:"ServUrlchangActnType,omitempty"`
	// Enter path to insert.
	ServUrlchangPathInsrt string `json:"ServUrlchangPathInsrt,omitempty"`
	// Insert the specified path before or after the matched section
	ServUrlchangInsrtPostn SlbCurCfgEnhVirtServicesThirdPartTableServUrlchangInsrtPostn `json:"ServUrlchangInsrtPostn,omitempty"`
}

func (p SlbCurCfgEnhVirtServicesThirdPartTableParams) iMABean() {}
