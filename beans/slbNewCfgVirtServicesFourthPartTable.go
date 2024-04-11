package beans

import (
	"fmt"
	"reflect"
)

// SlbNewCfgVirtServicesFourthPartTable The table of virtual Services.
// Note:This mib is not supported for VX instance of Virtualization.
type SlbNewCfgVirtServicesFourthPartTable struct {
	// The number of the virtual server.
	SlbNewCfgVirtServFourthPartIndex int32
	// The service index. This has no external meaning
	SlbNewCfgVirtServiceFourthPartIndex int32
	Params                              *SlbNewCfgVirtServicesFourthPartTableParams
}

func NewSlbNewCfgVirtServicesFourthPartTableList() *SlbNewCfgVirtServicesFourthPartTable {
	return &SlbNewCfgVirtServicesFourthPartTable{}
}

func NewSlbNewCfgVirtServicesFourthPartTable(
	slbNewCfgVirtServFourthPartIndex int32,
	slbNewCfgVirtServiceFourthPartIndex int32,
	params *SlbNewCfgVirtServicesFourthPartTableParams,
) *SlbNewCfgVirtServicesFourthPartTable {
	return &SlbNewCfgVirtServicesFourthPartTable{
		SlbNewCfgVirtServFourthPartIndex:    slbNewCfgVirtServFourthPartIndex,
		SlbNewCfgVirtServiceFourthPartIndex: slbNewCfgVirtServiceFourthPartIndex,
		Params:                              params,
	}
}

func (c *SlbNewCfgVirtServicesFourthPartTable) Name() string {
	return "SlbNewCfgVirtServicesFourthPartTable"
}

func (c *SlbNewCfgVirtServicesFourthPartTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbNewCfgVirtServicesFourthPartTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbNewCfgVirtServicesFourthPartTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbNewCfgVirtServFourthPartIndex).IsZero() &&
		reflect.ValueOf(c.SlbNewCfgVirtServiceFourthPartIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbNewCfgVirtServFourthPartIndex) + "/" + fmt.Sprint(c.SlbNewCfgVirtServiceFourthPartIndex)
}

type SlbNewCfgVirtServicesFourthPartTableServPathHideStatus int32

const (
	SlbNewCfgVirtServicesFourthPartTableServPathHideStatus_Enable      SlbNewCfgVirtServicesFourthPartTableServPathHideStatus = 1
	SlbNewCfgVirtServicesFourthPartTableServPathHideStatus_Disable     SlbNewCfgVirtServicesFourthPartTableServPathHideStatus = 2
	SlbNewCfgVirtServicesFourthPartTableServPathHideStatus_Clear       SlbNewCfgVirtServicesFourthPartTableServPathHideStatus = 3
	SlbNewCfgVirtServicesFourthPartTableServPathHideStatus_Unsupported SlbNewCfgVirtServicesFourthPartTableServPathHideStatus = 2147483647
)

type SlbNewCfgVirtServicesFourthPartTableServPathHideHostType int32

const (
	SlbNewCfgVirtServicesFourthPartTableServPathHideHostType_Sufx        SlbNewCfgVirtServicesFourthPartTableServPathHideHostType = 1
	SlbNewCfgVirtServicesFourthPartTableServPathHideHostType_Prefx       SlbNewCfgVirtServicesFourthPartTableServPathHideHostType = 2
	SlbNewCfgVirtServicesFourthPartTableServPathHideHostType_Eq          SlbNewCfgVirtServicesFourthPartTableServPathHideHostType = 3
	SlbNewCfgVirtServicesFourthPartTableServPathHideHostType_Incl        SlbNewCfgVirtServicesFourthPartTableServPathHideHostType = 4
	SlbNewCfgVirtServicesFourthPartTableServPathHideHostType_Any         SlbNewCfgVirtServicesFourthPartTableServPathHideHostType = 5
	SlbNewCfgVirtServicesFourthPartTableServPathHideHostType_Unsupported SlbNewCfgVirtServicesFourthPartTableServPathHideHostType = 2147483647
)

type SlbNewCfgVirtServicesFourthPartTableServPathHidePathType int32

const (
	SlbNewCfgVirtServicesFourthPartTableServPathHidePathType_Sufx        SlbNewCfgVirtServicesFourthPartTableServPathHidePathType = 1
	SlbNewCfgVirtServicesFourthPartTableServPathHidePathType_Prefx       SlbNewCfgVirtServicesFourthPartTableServPathHidePathType = 2
	SlbNewCfgVirtServicesFourthPartTableServPathHidePathType_Eq          SlbNewCfgVirtServicesFourthPartTableServPathHidePathType = 3
	SlbNewCfgVirtServicesFourthPartTableServPathHidePathType_None        SlbNewCfgVirtServicesFourthPartTableServPathHidePathType = 4
	SlbNewCfgVirtServicesFourthPartTableServPathHidePathType_Unsupported SlbNewCfgVirtServicesFourthPartTableServPathHidePathType = 2147483647
)

type SlbNewCfgVirtServicesFourthPartTableServTextrepStatus int32

const (
	SlbNewCfgVirtServicesFourthPartTableServTextrepStatus_Enable      SlbNewCfgVirtServicesFourthPartTableServTextrepStatus = 1
	SlbNewCfgVirtServicesFourthPartTableServTextrepStatus_Disable     SlbNewCfgVirtServicesFourthPartTableServTextrepStatus = 2
	SlbNewCfgVirtServicesFourthPartTableServTextrepStatus_Clear       SlbNewCfgVirtServicesFourthPartTableServTextrepStatus = 3
	SlbNewCfgVirtServicesFourthPartTableServTextrepStatus_Unsupported SlbNewCfgVirtServicesFourthPartTableServTextrepStatus = 2147483647
)

type SlbNewCfgVirtServicesFourthPartTableServTextrepAction int32

const (
	SlbNewCfgVirtServicesFourthPartTableServTextrepAction_None        SlbNewCfgVirtServicesFourthPartTableServTextrepAction = 0
	SlbNewCfgVirtServicesFourthPartTableServTextrepAction_Replace     SlbNewCfgVirtServicesFourthPartTableServTextrepAction = 1
	SlbNewCfgVirtServicesFourthPartTableServTextrepAction_Remove      SlbNewCfgVirtServicesFourthPartTableServTextrepAction = 2
	SlbNewCfgVirtServicesFourthPartTableServTextrepAction_Unsupported SlbNewCfgVirtServicesFourthPartTableServTextrepAction = 2147483647
)

type SlbNewCfgVirtServicesFourthPartTableParams struct {
	// The number of the virtual server.
	ServFourthPartIndex int32 `json:"ServFourthPartIndex,omitempty"`
	// The service index. This has no external meaning
	FourthPartIndex int32 `json:"FourthPartIndex,omitempty"`
	// Enter new page name or none.
	ServUrlchangNewPgName string `json:"ServUrlchangNewPgName,omitempty"`
	// Enter new page type or none.
	ServUrlchangNewPgType string `json:"ServUrlchangNewPgType,omitempty"`
	// Enter enabled/disabled/clear [Default: clear].
	ServPathHideStatus SlbNewCfgVirtServicesFourthPartTableServPathHideStatus `json:"ServPathHideStatus,omitempty"`
	// Enter hostname type [sufx|prefx|eq|incl|any] [Default: any].
	ServPathHideHostType SlbNewCfgVirtServicesFourthPartTableServPathHideHostType `json:"ServPathHideHostType,omitempty"`
	// Enter hostname to match.
	ServPathHideHostName string `json:"ServPathHideHostName,omitempty"`
	// Enter path match type [sufx|prefx|eq|none].
	ServPathHidePathType SlbNewCfgVirtServicesFourthPartTableServPathHidePathType `json:"ServPathHidePathType,omitempty"`
	// Enter path to remove.
	ServPathHidePathName string `json:"ServPathHidePathName,omitempty"`
	// Enter enabled/disabled/clear [Default: clear].
	ServTextrepStatus SlbNewCfgVirtServicesFourthPartTableServTextrepStatus `json:"ServTextrepStatus,omitempty"`
	// Enter action [replace|remove|none].
	ServTextrepAction SlbNewCfgVirtServicesFourthPartTableServTextrepAction `json:"ServTextrepAction,omitempty"`
}

func (p SlbNewCfgVirtServicesFourthPartTableParams) iMABean() {}
