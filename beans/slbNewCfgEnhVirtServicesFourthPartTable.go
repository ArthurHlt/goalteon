package beans

import (
	"fmt"
	"reflect"
)

// SlbNewCfgEnhVirtServicesFourthPartTable The table of virtual Services.
// Note:This mib is not supported for VX instance of Virtualization.
type SlbNewCfgEnhVirtServicesFourthPartTable struct {
	// The number of the virtual server.
	SlbNewCfgEnhVirtServFourthPartIndex string
	// The service index. This has no external meaning
	SlbNewCfgEnhVirtServiceFourthPartIndex int32
	Params                                 *SlbNewCfgEnhVirtServicesFourthPartTableParams
}

func NewSlbNewCfgEnhVirtServicesFourthPartTableList() *SlbNewCfgEnhVirtServicesFourthPartTable {
	return &SlbNewCfgEnhVirtServicesFourthPartTable{}
}

func NewSlbNewCfgEnhVirtServicesFourthPartTable(
	slbNewCfgEnhVirtServFourthPartIndex string,
	slbNewCfgEnhVirtServiceFourthPartIndex int32,
	params *SlbNewCfgEnhVirtServicesFourthPartTableParams,
) *SlbNewCfgEnhVirtServicesFourthPartTable {
	return &SlbNewCfgEnhVirtServicesFourthPartTable{
		SlbNewCfgEnhVirtServFourthPartIndex:    slbNewCfgEnhVirtServFourthPartIndex,
		SlbNewCfgEnhVirtServiceFourthPartIndex: slbNewCfgEnhVirtServiceFourthPartIndex,
		Params:                                 params,
	}
}

func (c *SlbNewCfgEnhVirtServicesFourthPartTable) Name() string {
	return "SlbNewCfgEnhVirtServicesFourthPartTable"
}

func (c *SlbNewCfgEnhVirtServicesFourthPartTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbNewCfgEnhVirtServicesFourthPartTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbNewCfgEnhVirtServicesFourthPartTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbNewCfgEnhVirtServFourthPartIndex).IsZero() &&
		reflect.ValueOf(c.SlbNewCfgEnhVirtServiceFourthPartIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbNewCfgEnhVirtServFourthPartIndex) + "/" + fmt.Sprint(c.SlbNewCfgEnhVirtServiceFourthPartIndex)
}

type SlbNewCfgEnhVirtServicesFourthPartTableServPathHideStatus int32

const (
	SlbNewCfgEnhVirtServicesFourthPartTableServPathHideStatus_Enable      SlbNewCfgEnhVirtServicesFourthPartTableServPathHideStatus = 1
	SlbNewCfgEnhVirtServicesFourthPartTableServPathHideStatus_Disable     SlbNewCfgEnhVirtServicesFourthPartTableServPathHideStatus = 2
	SlbNewCfgEnhVirtServicesFourthPartTableServPathHideStatus_Clear       SlbNewCfgEnhVirtServicesFourthPartTableServPathHideStatus = 3
	SlbNewCfgEnhVirtServicesFourthPartTableServPathHideStatus_Unsupported SlbNewCfgEnhVirtServicesFourthPartTableServPathHideStatus = 2147483647
)

type SlbNewCfgEnhVirtServicesFourthPartTableServPathHideHostType int32

const (
	SlbNewCfgEnhVirtServicesFourthPartTableServPathHideHostType_Sufx        SlbNewCfgEnhVirtServicesFourthPartTableServPathHideHostType = 1
	SlbNewCfgEnhVirtServicesFourthPartTableServPathHideHostType_Prefx       SlbNewCfgEnhVirtServicesFourthPartTableServPathHideHostType = 2
	SlbNewCfgEnhVirtServicesFourthPartTableServPathHideHostType_Eq          SlbNewCfgEnhVirtServicesFourthPartTableServPathHideHostType = 3
	SlbNewCfgEnhVirtServicesFourthPartTableServPathHideHostType_Incl        SlbNewCfgEnhVirtServicesFourthPartTableServPathHideHostType = 4
	SlbNewCfgEnhVirtServicesFourthPartTableServPathHideHostType_Any         SlbNewCfgEnhVirtServicesFourthPartTableServPathHideHostType = 5
	SlbNewCfgEnhVirtServicesFourthPartTableServPathHideHostType_Unsupported SlbNewCfgEnhVirtServicesFourthPartTableServPathHideHostType = 2147483647
)

type SlbNewCfgEnhVirtServicesFourthPartTableServPathHidePathType int32

const (
	SlbNewCfgEnhVirtServicesFourthPartTableServPathHidePathType_Sufx        SlbNewCfgEnhVirtServicesFourthPartTableServPathHidePathType = 1
	SlbNewCfgEnhVirtServicesFourthPartTableServPathHidePathType_Prefx       SlbNewCfgEnhVirtServicesFourthPartTableServPathHidePathType = 2
	SlbNewCfgEnhVirtServicesFourthPartTableServPathHidePathType_Eq          SlbNewCfgEnhVirtServicesFourthPartTableServPathHidePathType = 3
	SlbNewCfgEnhVirtServicesFourthPartTableServPathHidePathType_None        SlbNewCfgEnhVirtServicesFourthPartTableServPathHidePathType = 4
	SlbNewCfgEnhVirtServicesFourthPartTableServPathHidePathType_Unsupported SlbNewCfgEnhVirtServicesFourthPartTableServPathHidePathType = 2147483647
)

type SlbNewCfgEnhVirtServicesFourthPartTableServTextrepStatus int32

const (
	SlbNewCfgEnhVirtServicesFourthPartTableServTextrepStatus_Enable      SlbNewCfgEnhVirtServicesFourthPartTableServTextrepStatus = 1
	SlbNewCfgEnhVirtServicesFourthPartTableServTextrepStatus_Disable     SlbNewCfgEnhVirtServicesFourthPartTableServTextrepStatus = 2
	SlbNewCfgEnhVirtServicesFourthPartTableServTextrepStatus_Clear       SlbNewCfgEnhVirtServicesFourthPartTableServTextrepStatus = 3
	SlbNewCfgEnhVirtServicesFourthPartTableServTextrepStatus_Unsupported SlbNewCfgEnhVirtServicesFourthPartTableServTextrepStatus = 2147483647
)

type SlbNewCfgEnhVirtServicesFourthPartTableServTextrepAction int32

const (
	SlbNewCfgEnhVirtServicesFourthPartTableServTextrepAction_None        SlbNewCfgEnhVirtServicesFourthPartTableServTextrepAction = 0
	SlbNewCfgEnhVirtServicesFourthPartTableServTextrepAction_Replace     SlbNewCfgEnhVirtServicesFourthPartTableServTextrepAction = 1
	SlbNewCfgEnhVirtServicesFourthPartTableServTextrepAction_Remove      SlbNewCfgEnhVirtServicesFourthPartTableServTextrepAction = 2
	SlbNewCfgEnhVirtServicesFourthPartTableServTextrepAction_Unsupported SlbNewCfgEnhVirtServicesFourthPartTableServTextrepAction = 2147483647
)

type SlbNewCfgEnhVirtServicesFourthPartTableParams struct {
	// The number of the virtual server.
	ServFourthPartIndex string `json:"ServFourthPartIndex,omitempty"`
	// The service index. This has no external meaning
	FourthPartIndex int32 `json:"FourthPartIndex,omitempty"`
	// Enter new page name or none.
	ServUrlchangNewPgName string `json:"ServUrlchangNewPgName,omitempty"`
	// Enter new page type or none.
	ServUrlchangNewPgType string `json:"ServUrlchangNewPgType,omitempty"`
	// Enter enabled/disabled/clear [Default: clear].
	ServPathHideStatus SlbNewCfgEnhVirtServicesFourthPartTableServPathHideStatus `json:"ServPathHideStatus,omitempty"`
	// Enter hostname type [sufx|prefx|eq|incl|any] [Default: any].
	ServPathHideHostType SlbNewCfgEnhVirtServicesFourthPartTableServPathHideHostType `json:"ServPathHideHostType,omitempty"`
	// Enter hostname to match.
	ServPathHideHostName string `json:"ServPathHideHostName,omitempty"`
	// Enter path match type [sufx|prefx|eq|none].
	ServPathHidePathType SlbNewCfgEnhVirtServicesFourthPartTableServPathHidePathType `json:"ServPathHidePathType,omitempty"`
	// Enter path to remove.
	ServPathHidePathName string `json:"ServPathHidePathName,omitempty"`
	// Enter enabled/disabled/clear [Default: clear].
	ServTextrepStatus SlbNewCfgEnhVirtServicesFourthPartTableServTextrepStatus `json:"ServTextrepStatus,omitempty"`
	// Enter action [replace|remove|none].
	ServTextrepAction SlbNewCfgEnhVirtServicesFourthPartTableServTextrepAction `json:"ServTextrepAction,omitempty"`
}

func (p SlbNewCfgEnhVirtServicesFourthPartTableParams) iMABean() {}
