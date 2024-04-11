package beans

import (
	"fmt"
	"reflect"
)

// SlbCurCfgEnhVirtServicesFourthPartTable The table of virtual Services.
// Note:This mib is not supported for VX instance of Virtualization.
type SlbCurCfgEnhVirtServicesFourthPartTable struct {
	// The number of the virtual server.
	SlbCurCfgEnhVirtServFourthPartIndex string
	// The service index. This has no external meaning
	SlbCurCfgEnhVirtServiceFourthPartIndex int32
	Params                                 *SlbCurCfgEnhVirtServicesFourthPartTableParams
}

func NewSlbCurCfgEnhVirtServicesFourthPartTableList() *SlbCurCfgEnhVirtServicesFourthPartTable {
	return &SlbCurCfgEnhVirtServicesFourthPartTable{}
}

func NewSlbCurCfgEnhVirtServicesFourthPartTable(
	slbCurCfgEnhVirtServFourthPartIndex string,
	slbCurCfgEnhVirtServiceFourthPartIndex int32,
	params *SlbCurCfgEnhVirtServicesFourthPartTableParams,
) *SlbCurCfgEnhVirtServicesFourthPartTable {
	return &SlbCurCfgEnhVirtServicesFourthPartTable{
		SlbCurCfgEnhVirtServFourthPartIndex:    slbCurCfgEnhVirtServFourthPartIndex,
		SlbCurCfgEnhVirtServiceFourthPartIndex: slbCurCfgEnhVirtServiceFourthPartIndex,
		Params:                                 params,
	}
}

func (c *SlbCurCfgEnhVirtServicesFourthPartTable) Name() string {
	return "SlbCurCfgEnhVirtServicesFourthPartTable"
}

func (c *SlbCurCfgEnhVirtServicesFourthPartTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbCurCfgEnhVirtServicesFourthPartTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbCurCfgEnhVirtServicesFourthPartTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbCurCfgEnhVirtServFourthPartIndex).IsZero() &&
		reflect.ValueOf(c.SlbCurCfgEnhVirtServiceFourthPartIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbCurCfgEnhVirtServFourthPartIndex) + "/" + fmt.Sprint(c.SlbCurCfgEnhVirtServiceFourthPartIndex)
}

type SlbCurCfgEnhVirtServicesFourthPartTableServPathHideStatus int32

const (
	SlbCurCfgEnhVirtServicesFourthPartTableServPathHideStatus_Enable      SlbCurCfgEnhVirtServicesFourthPartTableServPathHideStatus = 1
	SlbCurCfgEnhVirtServicesFourthPartTableServPathHideStatus_Disable     SlbCurCfgEnhVirtServicesFourthPartTableServPathHideStatus = 2
	SlbCurCfgEnhVirtServicesFourthPartTableServPathHideStatus_Clear       SlbCurCfgEnhVirtServicesFourthPartTableServPathHideStatus = 3
	SlbCurCfgEnhVirtServicesFourthPartTableServPathHideStatus_Unsupported SlbCurCfgEnhVirtServicesFourthPartTableServPathHideStatus = 2147483647
)

type SlbCurCfgEnhVirtServicesFourthPartTableServPathHideHostType int32

const (
	SlbCurCfgEnhVirtServicesFourthPartTableServPathHideHostType_Sufx        SlbCurCfgEnhVirtServicesFourthPartTableServPathHideHostType = 1
	SlbCurCfgEnhVirtServicesFourthPartTableServPathHideHostType_Prefx       SlbCurCfgEnhVirtServicesFourthPartTableServPathHideHostType = 2
	SlbCurCfgEnhVirtServicesFourthPartTableServPathHideHostType_Eq          SlbCurCfgEnhVirtServicesFourthPartTableServPathHideHostType = 3
	SlbCurCfgEnhVirtServicesFourthPartTableServPathHideHostType_Incl        SlbCurCfgEnhVirtServicesFourthPartTableServPathHideHostType = 4
	SlbCurCfgEnhVirtServicesFourthPartTableServPathHideHostType_Any         SlbCurCfgEnhVirtServicesFourthPartTableServPathHideHostType = 5
	SlbCurCfgEnhVirtServicesFourthPartTableServPathHideHostType_Unsupported SlbCurCfgEnhVirtServicesFourthPartTableServPathHideHostType = 2147483647
)

type SlbCurCfgEnhVirtServicesFourthPartTableServPathHidePathType int32

const (
	SlbCurCfgEnhVirtServicesFourthPartTableServPathHidePathType_Sufx        SlbCurCfgEnhVirtServicesFourthPartTableServPathHidePathType = 1
	SlbCurCfgEnhVirtServicesFourthPartTableServPathHidePathType_Prefx       SlbCurCfgEnhVirtServicesFourthPartTableServPathHidePathType = 2
	SlbCurCfgEnhVirtServicesFourthPartTableServPathHidePathType_Eq          SlbCurCfgEnhVirtServicesFourthPartTableServPathHidePathType = 3
	SlbCurCfgEnhVirtServicesFourthPartTableServPathHidePathType_None        SlbCurCfgEnhVirtServicesFourthPartTableServPathHidePathType = 4
	SlbCurCfgEnhVirtServicesFourthPartTableServPathHidePathType_Unsupported SlbCurCfgEnhVirtServicesFourthPartTableServPathHidePathType = 2147483647
)

type SlbCurCfgEnhVirtServicesFourthPartTableServTextrepStatus int32

const (
	SlbCurCfgEnhVirtServicesFourthPartTableServTextrepStatus_Enable      SlbCurCfgEnhVirtServicesFourthPartTableServTextrepStatus = 1
	SlbCurCfgEnhVirtServicesFourthPartTableServTextrepStatus_Disable     SlbCurCfgEnhVirtServicesFourthPartTableServTextrepStatus = 2
	SlbCurCfgEnhVirtServicesFourthPartTableServTextrepStatus_Clear       SlbCurCfgEnhVirtServicesFourthPartTableServTextrepStatus = 3
	SlbCurCfgEnhVirtServicesFourthPartTableServTextrepStatus_Unsupported SlbCurCfgEnhVirtServicesFourthPartTableServTextrepStatus = 2147483647
)

type SlbCurCfgEnhVirtServicesFourthPartTableServTextrepAction int32

const (
	SlbCurCfgEnhVirtServicesFourthPartTableServTextrepAction_None        SlbCurCfgEnhVirtServicesFourthPartTableServTextrepAction = 0
	SlbCurCfgEnhVirtServicesFourthPartTableServTextrepAction_Replace     SlbCurCfgEnhVirtServicesFourthPartTableServTextrepAction = 1
	SlbCurCfgEnhVirtServicesFourthPartTableServTextrepAction_Remove      SlbCurCfgEnhVirtServicesFourthPartTableServTextrepAction = 2
	SlbCurCfgEnhVirtServicesFourthPartTableServTextrepAction_Unsupported SlbCurCfgEnhVirtServicesFourthPartTableServTextrepAction = 2147483647
)

type SlbCurCfgEnhVirtServicesFourthPartTableParams struct {
	// The number of the virtual server.
	ServFourthPartIndex string `json:"ServFourthPartIndex,omitempty"`
	// The service index. This has no external meaning
	FourthPartIndex int32 `json:"FourthPartIndex,omitempty"`
	// Enter Cur page name or none.
	ServUrlchangNewPgName string `json:"ServUrlchangNewPgName,omitempty"`
	// Enter Cur page type or none.
	ServUrlchangNewPgType string `json:"ServUrlchangNewPgType,omitempty"`
	// Enter enabled/disabled/clear [Default: clear].
	ServPathHideStatus SlbCurCfgEnhVirtServicesFourthPartTableServPathHideStatus `json:"ServPathHideStatus,omitempty"`
	// Enter hostname type [sufx|prefx|eq|incl|any] [Default: any].
	ServPathHideHostType SlbCurCfgEnhVirtServicesFourthPartTableServPathHideHostType `json:"ServPathHideHostType,omitempty"`
	// Enter hostname to match.
	ServPathHideHostName string `json:"ServPathHideHostName,omitempty"`
	// Enter path match type [sufx|prefx|eq|none].
	ServPathHidePathType SlbCurCfgEnhVirtServicesFourthPartTableServPathHidePathType `json:"ServPathHidePathType,omitempty"`
	// Enter path to remove.
	ServPathHidePathName string `json:"ServPathHidePathName,omitempty"`
	// Enter enabled/disabled/clear [Default: clear].
	ServTextrepStatus SlbCurCfgEnhVirtServicesFourthPartTableServTextrepStatus `json:"ServTextrepStatus,omitempty"`
	// Enter action [replace|remove|none].
	ServTextrepAction SlbCurCfgEnhVirtServicesFourthPartTableServTextrepAction `json:"ServTextrepAction,omitempty"`
}

func (p SlbCurCfgEnhVirtServicesFourthPartTableParams) iMABean() {}
