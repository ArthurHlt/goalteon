package beans

import (
	"fmt"
	"reflect"
)

// SlbCurCfgVirtServicesFourthPartTable The table of virtual Services.
// Note:This mib is not supported for VX instance of Virtualization.
type SlbCurCfgVirtServicesFourthPartTable struct {
	// The number of the virtual server.
	SlbCurCfgVirtServFourthPartIndex int32
	// The service index. This has no external meaning
	SlbCurCfgVirtServiceFourthPartIndex int32
	Params                              *SlbCurCfgVirtServicesFourthPartTableParams
}

func NewSlbCurCfgVirtServicesFourthPartTableList() *SlbCurCfgVirtServicesFourthPartTable {
	return &SlbCurCfgVirtServicesFourthPartTable{}
}

func NewSlbCurCfgVirtServicesFourthPartTable(
	slbCurCfgVirtServFourthPartIndex int32,
	slbCurCfgVirtServiceFourthPartIndex int32,
	params *SlbCurCfgVirtServicesFourthPartTableParams,
) *SlbCurCfgVirtServicesFourthPartTable {
	return &SlbCurCfgVirtServicesFourthPartTable{
		SlbCurCfgVirtServFourthPartIndex:    slbCurCfgVirtServFourthPartIndex,
		SlbCurCfgVirtServiceFourthPartIndex: slbCurCfgVirtServiceFourthPartIndex,
		Params:                              params,
	}
}

func (c *SlbCurCfgVirtServicesFourthPartTable) Name() string {
	return "SlbCurCfgVirtServicesFourthPartTable"
}

func (c *SlbCurCfgVirtServicesFourthPartTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbCurCfgVirtServicesFourthPartTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbCurCfgVirtServicesFourthPartTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbCurCfgVirtServFourthPartIndex).IsZero() &&
		reflect.ValueOf(c.SlbCurCfgVirtServiceFourthPartIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbCurCfgVirtServFourthPartIndex) + "/" + fmt.Sprint(c.SlbCurCfgVirtServiceFourthPartIndex)
}

type SlbCurCfgVirtServicesFourthPartTableServPathHideStatus int32

const (
	SlbCurCfgVirtServicesFourthPartTableServPathHideStatus_Enable      SlbCurCfgVirtServicesFourthPartTableServPathHideStatus = 1
	SlbCurCfgVirtServicesFourthPartTableServPathHideStatus_Disable     SlbCurCfgVirtServicesFourthPartTableServPathHideStatus = 2
	SlbCurCfgVirtServicesFourthPartTableServPathHideStatus_Clear       SlbCurCfgVirtServicesFourthPartTableServPathHideStatus = 3
	SlbCurCfgVirtServicesFourthPartTableServPathHideStatus_Unsupported SlbCurCfgVirtServicesFourthPartTableServPathHideStatus = 2147483647
)

type SlbCurCfgVirtServicesFourthPartTableServPathHideHostType int32

const (
	SlbCurCfgVirtServicesFourthPartTableServPathHideHostType_Sufx        SlbCurCfgVirtServicesFourthPartTableServPathHideHostType = 1
	SlbCurCfgVirtServicesFourthPartTableServPathHideHostType_Prefx       SlbCurCfgVirtServicesFourthPartTableServPathHideHostType = 2
	SlbCurCfgVirtServicesFourthPartTableServPathHideHostType_Eq          SlbCurCfgVirtServicesFourthPartTableServPathHideHostType = 3
	SlbCurCfgVirtServicesFourthPartTableServPathHideHostType_Incl        SlbCurCfgVirtServicesFourthPartTableServPathHideHostType = 4
	SlbCurCfgVirtServicesFourthPartTableServPathHideHostType_Any         SlbCurCfgVirtServicesFourthPartTableServPathHideHostType = 5
	SlbCurCfgVirtServicesFourthPartTableServPathHideHostType_Unsupported SlbCurCfgVirtServicesFourthPartTableServPathHideHostType = 2147483647
)

type SlbCurCfgVirtServicesFourthPartTableServPathHidePathType int32

const (
	SlbCurCfgVirtServicesFourthPartTableServPathHidePathType_Sufx        SlbCurCfgVirtServicesFourthPartTableServPathHidePathType = 1
	SlbCurCfgVirtServicesFourthPartTableServPathHidePathType_Prefx       SlbCurCfgVirtServicesFourthPartTableServPathHidePathType = 2
	SlbCurCfgVirtServicesFourthPartTableServPathHidePathType_Eq          SlbCurCfgVirtServicesFourthPartTableServPathHidePathType = 3
	SlbCurCfgVirtServicesFourthPartTableServPathHidePathType_None        SlbCurCfgVirtServicesFourthPartTableServPathHidePathType = 4
	SlbCurCfgVirtServicesFourthPartTableServPathHidePathType_Unsupported SlbCurCfgVirtServicesFourthPartTableServPathHidePathType = 2147483647
)

type SlbCurCfgVirtServicesFourthPartTableServTextrepStatus int32

const (
	SlbCurCfgVirtServicesFourthPartTableServTextrepStatus_Enable      SlbCurCfgVirtServicesFourthPartTableServTextrepStatus = 1
	SlbCurCfgVirtServicesFourthPartTableServTextrepStatus_Disable     SlbCurCfgVirtServicesFourthPartTableServTextrepStatus = 2
	SlbCurCfgVirtServicesFourthPartTableServTextrepStatus_Clear       SlbCurCfgVirtServicesFourthPartTableServTextrepStatus = 3
	SlbCurCfgVirtServicesFourthPartTableServTextrepStatus_Unsupported SlbCurCfgVirtServicesFourthPartTableServTextrepStatus = 2147483647
)

type SlbCurCfgVirtServicesFourthPartTableServTextrepAction int32

const (
	SlbCurCfgVirtServicesFourthPartTableServTextrepAction_None        SlbCurCfgVirtServicesFourthPartTableServTextrepAction = 0
	SlbCurCfgVirtServicesFourthPartTableServTextrepAction_Replace     SlbCurCfgVirtServicesFourthPartTableServTextrepAction = 1
	SlbCurCfgVirtServicesFourthPartTableServTextrepAction_Remove      SlbCurCfgVirtServicesFourthPartTableServTextrepAction = 2
	SlbCurCfgVirtServicesFourthPartTableServTextrepAction_Unsupported SlbCurCfgVirtServicesFourthPartTableServTextrepAction = 2147483647
)

type SlbCurCfgVirtServicesFourthPartTableParams struct {
	// The number of the virtual server.
	ServFourthPartIndex int32 `json:"ServFourthPartIndex,omitempty"`
	// The service index. This has no external meaning
	FourthPartIndex int32 `json:"FourthPartIndex,omitempty"`
	// Enter Cur page name or none.
	ServUrlchangNewPgName string `json:"ServUrlchangNewPgName,omitempty"`
	// Enter Cur page type or none.
	ServUrlchangNewPgType string `json:"ServUrlchangNewPgType,omitempty"`
	// Enter enabled/disabled/clear [Default: clear].
	ServPathHideStatus SlbCurCfgVirtServicesFourthPartTableServPathHideStatus `json:"ServPathHideStatus,omitempty"`
	// Enter hostname type [sufx|prefx|eq|incl|any] [Default: any].
	ServPathHideHostType SlbCurCfgVirtServicesFourthPartTableServPathHideHostType `json:"ServPathHideHostType,omitempty"`
	// Enter hostname to match.
	ServPathHideHostName string `json:"ServPathHideHostName,omitempty"`
	// Enter path match type [sufx|prefx|eq|none].
	ServPathHidePathType SlbCurCfgVirtServicesFourthPartTableServPathHidePathType `json:"ServPathHidePathType,omitempty"`
	// Enter path to remove.
	ServPathHidePathName string `json:"ServPathHidePathName,omitempty"`
	// Enter enabled/disabled/clear [Default: clear].
	ServTextrepStatus SlbCurCfgVirtServicesFourthPartTableServTextrepStatus `json:"ServTextrepStatus,omitempty"`
	// Enter action [replace|remove|none].
	ServTextrepAction SlbCurCfgVirtServicesFourthPartTableServTextrepAction `json:"ServTextrepAction,omitempty"`
}

func (p SlbCurCfgVirtServicesFourthPartTableParams) iMABean() {}
