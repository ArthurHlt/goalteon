package beans

import (
	"fmt"
	"reflect"
)

// SlbCurCfgVirtServicesThirdPartTable The table of virtual Services.
// Note:This mib is not supported for VX instance of Virtualization.
type SlbCurCfgVirtServicesThirdPartTable struct {
	// The number of the virtual server.
	SlbCurCfgVirtServThirdPartIndex int32
	// The service index. This has no external meaning
	SlbCurCfgVirtServiceThirdPartIndex int32
	Params                             *SlbCurCfgVirtServicesThirdPartTableParams
}

func NewSlbCurCfgVirtServicesThirdPartTableList() *SlbCurCfgVirtServicesThirdPartTable {
	return &SlbCurCfgVirtServicesThirdPartTable{}
}

func NewSlbCurCfgVirtServicesThirdPartTable(
	slbCurCfgVirtServThirdPartIndex int32,
	slbCurCfgVirtServiceThirdPartIndex int32,
	params *SlbCurCfgVirtServicesThirdPartTableParams,
) *SlbCurCfgVirtServicesThirdPartTable {
	return &SlbCurCfgVirtServicesThirdPartTable{
		SlbCurCfgVirtServThirdPartIndex:    slbCurCfgVirtServThirdPartIndex,
		SlbCurCfgVirtServiceThirdPartIndex: slbCurCfgVirtServiceThirdPartIndex,
		Params:                             params,
	}
}

func (c *SlbCurCfgVirtServicesThirdPartTable) Name() string {
	return "SlbCurCfgVirtServicesThirdPartTable"
}

func (c *SlbCurCfgVirtServicesThirdPartTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbCurCfgVirtServicesThirdPartTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbCurCfgVirtServicesThirdPartTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbCurCfgVirtServThirdPartIndex).IsZero() &&
		reflect.ValueOf(c.SlbCurCfgVirtServiceThirdPartIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbCurCfgVirtServThirdPartIndex) + "/" + fmt.Sprint(c.SlbCurCfgVirtServiceThirdPartIndex)
}

type SlbCurCfgVirtServicesThirdPartTableServUrlchangPathType int32

const (
	SlbCurCfgVirtServicesThirdPartTableServUrlchangPathType_Sufx        SlbCurCfgVirtServicesThirdPartTableServUrlchangPathType = 1
	SlbCurCfgVirtServicesThirdPartTableServUrlchangPathType_Prefx       SlbCurCfgVirtServicesThirdPartTableServUrlchangPathType = 2
	SlbCurCfgVirtServicesThirdPartTableServUrlchangPathType_Eq          SlbCurCfgVirtServicesThirdPartTableServUrlchangPathType = 3
	SlbCurCfgVirtServicesThirdPartTableServUrlchangPathType_Incl        SlbCurCfgVirtServicesThirdPartTableServUrlchangPathType = 4
	SlbCurCfgVirtServicesThirdPartTableServUrlchangPathType_Any         SlbCurCfgVirtServicesThirdPartTableServUrlchangPathType = 5
	SlbCurCfgVirtServicesThirdPartTableServUrlchangPathType_None        SlbCurCfgVirtServicesThirdPartTableServUrlchangPathType = 6
	SlbCurCfgVirtServicesThirdPartTableServUrlchangPathType_Unsupported SlbCurCfgVirtServicesThirdPartTableServUrlchangPathType = 2147483647
)

type SlbCurCfgVirtServicesThirdPartTableServUrlchangActnType int32

const (
	SlbCurCfgVirtServicesThirdPartTableServUrlchangActnType_Insert      SlbCurCfgVirtServicesThirdPartTableServUrlchangActnType = 1
	SlbCurCfgVirtServicesThirdPartTableServUrlchangActnType_Replace     SlbCurCfgVirtServicesThirdPartTableServUrlchangActnType = 2
	SlbCurCfgVirtServicesThirdPartTableServUrlchangActnType_Remove      SlbCurCfgVirtServicesThirdPartTableServUrlchangActnType = 3
	SlbCurCfgVirtServicesThirdPartTableServUrlchangActnType_None        SlbCurCfgVirtServicesThirdPartTableServUrlchangActnType = 4
	SlbCurCfgVirtServicesThirdPartTableServUrlchangActnType_Unsupported SlbCurCfgVirtServicesThirdPartTableServUrlchangActnType = 2147483647
)

type SlbCurCfgVirtServicesThirdPartTableServUrlchangInsrtPostn int32

const (
	SlbCurCfgVirtServicesThirdPartTableServUrlchangInsrtPostn_Before      SlbCurCfgVirtServicesThirdPartTableServUrlchangInsrtPostn = 1
	SlbCurCfgVirtServicesThirdPartTableServUrlchangInsrtPostn_After       SlbCurCfgVirtServicesThirdPartTableServUrlchangInsrtPostn = 2
	SlbCurCfgVirtServicesThirdPartTableServUrlchangInsrtPostn_Unsupported SlbCurCfgVirtServicesThirdPartTableServUrlchangInsrtPostn = 2147483647
)

type SlbCurCfgVirtServicesThirdPartTableParams struct {
	// The number of the virtual server.
	ServThirdPartIndex int32 `json:"ServThirdPartIndex,omitempty"`
	// The service index. This has no external meaning
	ThirdPartIndex int32 `json:"ThirdPartIndex,omitempty"`
	// Enter hostname to match.
	ServUrlchangHostName string `json:"ServUrlchangHostName,omitempty"`
	// Enter path match type [sufx|prefx|eq|incl|any].
	ServUrlchangPathType SlbCurCfgVirtServicesThirdPartTableServUrlchangPathType `json:"ServUrlchangPathType,omitempty"`
	// Enter path to match.
	ServUrlchangPathMatch string `json:"ServUrlchangPathMatch,omitempty"`
	// Enter page name to match or none.
	ServUrlchangPageName string `json:"ServUrlchangPageName,omitempty"`
	// Enter page type to match or none.
	ServUrlchangPageType string `json:"ServUrlchangPageType,omitempty"`
	// Enter path action type.
	ServUrlchangActnType SlbCurCfgVirtServicesThirdPartTableServUrlchangActnType `json:"ServUrlchangActnType,omitempty"`
	// Enter path to insert.
	ServUrlchangPathInsrt string `json:"ServUrlchangPathInsrt,omitempty"`
	// Insert the specified path before or after the matched section
	ServUrlchangInsrtPostn SlbCurCfgVirtServicesThirdPartTableServUrlchangInsrtPostn `json:"ServUrlchangInsrtPostn,omitempty"`
}

func (p SlbCurCfgVirtServicesThirdPartTableParams) iMABean() {}
