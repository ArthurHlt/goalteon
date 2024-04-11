package beans

import (
	"fmt"
	"reflect"
)

// SlbNewCfgVirtServicesThirdPartTable The table of virtual Services.
// Note:This mib is not supported for VX instance of Virtualization.
type SlbNewCfgVirtServicesThirdPartTable struct {
	// The number of the virtual server.
	SlbNewCfgVirtServThirdPartIndex int32
	// The service index. This has no external meaning
	SlbNewCfgVirtServiceThirdPartIndex int32
	Params                             *SlbNewCfgVirtServicesThirdPartTableParams
}

func NewSlbNewCfgVirtServicesThirdPartTableList() *SlbNewCfgVirtServicesThirdPartTable {
	return &SlbNewCfgVirtServicesThirdPartTable{}
}

func NewSlbNewCfgVirtServicesThirdPartTable(
	slbNewCfgVirtServThirdPartIndex int32,
	slbNewCfgVirtServiceThirdPartIndex int32,
	params *SlbNewCfgVirtServicesThirdPartTableParams,
) *SlbNewCfgVirtServicesThirdPartTable {
	return &SlbNewCfgVirtServicesThirdPartTable{
		SlbNewCfgVirtServThirdPartIndex:    slbNewCfgVirtServThirdPartIndex,
		SlbNewCfgVirtServiceThirdPartIndex: slbNewCfgVirtServiceThirdPartIndex,
		Params:                             params,
	}
}

func (c *SlbNewCfgVirtServicesThirdPartTable) Name() string {
	return "SlbNewCfgVirtServicesThirdPartTable"
}

func (c *SlbNewCfgVirtServicesThirdPartTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbNewCfgVirtServicesThirdPartTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbNewCfgVirtServicesThirdPartTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbNewCfgVirtServThirdPartIndex).IsZero() &&
		reflect.ValueOf(c.SlbNewCfgVirtServiceThirdPartIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbNewCfgVirtServThirdPartIndex) + "/" + fmt.Sprint(c.SlbNewCfgVirtServiceThirdPartIndex)
}

type SlbNewCfgVirtServicesThirdPartTableServUrlchangPathType int32

const (
	SlbNewCfgVirtServicesThirdPartTableServUrlchangPathType_Sufx        SlbNewCfgVirtServicesThirdPartTableServUrlchangPathType = 1
	SlbNewCfgVirtServicesThirdPartTableServUrlchangPathType_Prefx       SlbNewCfgVirtServicesThirdPartTableServUrlchangPathType = 2
	SlbNewCfgVirtServicesThirdPartTableServUrlchangPathType_Eq          SlbNewCfgVirtServicesThirdPartTableServUrlchangPathType = 3
	SlbNewCfgVirtServicesThirdPartTableServUrlchangPathType_Incl        SlbNewCfgVirtServicesThirdPartTableServUrlchangPathType = 4
	SlbNewCfgVirtServicesThirdPartTableServUrlchangPathType_Any         SlbNewCfgVirtServicesThirdPartTableServUrlchangPathType = 5
	SlbNewCfgVirtServicesThirdPartTableServUrlchangPathType_None        SlbNewCfgVirtServicesThirdPartTableServUrlchangPathType = 6
	SlbNewCfgVirtServicesThirdPartTableServUrlchangPathType_Unsupported SlbNewCfgVirtServicesThirdPartTableServUrlchangPathType = 2147483647
)

type SlbNewCfgVirtServicesThirdPartTableServUrlchangActnType int32

const (
	SlbNewCfgVirtServicesThirdPartTableServUrlchangActnType_Insert      SlbNewCfgVirtServicesThirdPartTableServUrlchangActnType = 1
	SlbNewCfgVirtServicesThirdPartTableServUrlchangActnType_Replace     SlbNewCfgVirtServicesThirdPartTableServUrlchangActnType = 2
	SlbNewCfgVirtServicesThirdPartTableServUrlchangActnType_Remove      SlbNewCfgVirtServicesThirdPartTableServUrlchangActnType = 3
	SlbNewCfgVirtServicesThirdPartTableServUrlchangActnType_None        SlbNewCfgVirtServicesThirdPartTableServUrlchangActnType = 4
	SlbNewCfgVirtServicesThirdPartTableServUrlchangActnType_Unsupported SlbNewCfgVirtServicesThirdPartTableServUrlchangActnType = 2147483647
)

type SlbNewCfgVirtServicesThirdPartTableServUrlchangInsrtPostn int32

const (
	SlbNewCfgVirtServicesThirdPartTableServUrlchangInsrtPostn_Before      SlbNewCfgVirtServicesThirdPartTableServUrlchangInsrtPostn = 1
	SlbNewCfgVirtServicesThirdPartTableServUrlchangInsrtPostn_After       SlbNewCfgVirtServicesThirdPartTableServUrlchangInsrtPostn = 2
	SlbNewCfgVirtServicesThirdPartTableServUrlchangInsrtPostn_Unsupported SlbNewCfgVirtServicesThirdPartTableServUrlchangInsrtPostn = 2147483647
)

type SlbNewCfgVirtServicesThirdPartTableParams struct {
	// The number of the virtual server.
	ServThirdPartIndex int32 `json:"ServThirdPartIndex,omitempty"`
	// The service index. This has no external meaning
	ThirdPartIndex int32 `json:"ThirdPartIndex,omitempty"`
	// Enter hostname to match.
	ServUrlchangHostName string `json:"ServUrlchangHostName,omitempty"`
	// Enter path match type [sufx|prefx|eq|incl|any|none].
	ServUrlchangPathType SlbNewCfgVirtServicesThirdPartTableServUrlchangPathType `json:"ServUrlchangPathType,omitempty"`
	// Enter path to match.
	ServUrlchangPathMatch string `json:"ServUrlchangPathMatch,omitempty"`
	// Enter page name to match or none.
	ServUrlchangPageName string `json:"ServUrlchangPageName,omitempty"`
	// Enter page type to match or none.
	ServUrlchangPageType string `json:"ServUrlchangPageType,omitempty"`
	// Enter path action type.
	ServUrlchangActnType SlbNewCfgVirtServicesThirdPartTableServUrlchangActnType `json:"ServUrlchangActnType,omitempty"`
	// Enter path to insert.
	ServUrlchangPathInsrt string `json:"ServUrlchangPathInsrt,omitempty"`
	// Insert the specified path before or after the matched section
	ServUrlchangInsrtPostn SlbNewCfgVirtServicesThirdPartTableServUrlchangInsrtPostn `json:"ServUrlchangInsrtPostn,omitempty"`
}

func (p SlbNewCfgVirtServicesThirdPartTableParams) iMABean() {}
