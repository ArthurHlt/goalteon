package beans

import (
	"fmt"
	"reflect"
)

// SlbCurCfgVirtServicesSecondPartTable The table of virtual Services.
// Note:This mib is not supported for VX instance of Virtualization.
type SlbCurCfgVirtServicesSecondPartTable struct {
	// The number of the virtual server.
	SlbCurCfgVirtServSecondPartIndex int32
	// The service index. This has no external meaning
	SlbCurCfgVirtServiceSecondPartIndex int32
	Params                              *SlbCurCfgVirtServicesSecondPartTableParams
}

func NewSlbCurCfgVirtServicesSecondPartTableList() *SlbCurCfgVirtServicesSecondPartTable {
	return &SlbCurCfgVirtServicesSecondPartTable{}
}

func NewSlbCurCfgVirtServicesSecondPartTable(
	slbCurCfgVirtServSecondPartIndex int32,
	slbCurCfgVirtServiceSecondPartIndex int32,
	params *SlbCurCfgVirtServicesSecondPartTableParams,
) *SlbCurCfgVirtServicesSecondPartTable {
	return &SlbCurCfgVirtServicesSecondPartTable{
		SlbCurCfgVirtServSecondPartIndex:    slbCurCfgVirtServSecondPartIndex,
		SlbCurCfgVirtServiceSecondPartIndex: slbCurCfgVirtServiceSecondPartIndex,
		Params:                              params,
	}
}

func (c *SlbCurCfgVirtServicesSecondPartTable) Name() string {
	return "SlbCurCfgVirtServicesSecondPartTable"
}

func (c *SlbCurCfgVirtServicesSecondPartTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbCurCfgVirtServicesSecondPartTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbCurCfgVirtServicesSecondPartTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbCurCfgVirtServSecondPartIndex).IsZero() &&
		reflect.ValueOf(c.SlbCurCfgVirtServiceSecondPartIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbCurCfgVirtServSecondPartIndex) + "/" + fmt.Sprint(c.SlbCurCfgVirtServiceSecondPartIndex)
}

type SlbCurCfgVirtServicesSecondPartTableConnmgtStatus int32

const (
	SlbCurCfgVirtServicesSecondPartTableConnmgtStatus_Disabled    SlbCurCfgVirtServicesSecondPartTableConnmgtStatus = 0
	SlbCurCfgVirtServicesSecondPartTableConnmgtStatus_Pooling     SlbCurCfgVirtServicesSecondPartTableConnmgtStatus = 1
	SlbCurCfgVirtServicesSecondPartTableConnmgtStatus_Muxenabled  SlbCurCfgVirtServicesSecondPartTableConnmgtStatus = 2
	SlbCurCfgVirtServicesSecondPartTableConnmgtStatus_Unsupported SlbCurCfgVirtServicesSecondPartTableConnmgtStatus = 2147483647
)

type SlbCurCfgVirtServicesSecondPartTableCloaksrv int32

const (
	SlbCurCfgVirtServicesSecondPartTableCloaksrv_Enabled     SlbCurCfgVirtServicesSecondPartTableCloaksrv = 1
	SlbCurCfgVirtServicesSecondPartTableCloaksrv_Disabled    SlbCurCfgVirtServicesSecondPartTableCloaksrv = 2
	SlbCurCfgVirtServicesSecondPartTableCloaksrv_Unsupported SlbCurCfgVirtServicesSecondPartTableCloaksrv = 2147483647
)

type SlbCurCfgVirtServicesSecondPartTableServErrcodeStatus int32

const (
	SlbCurCfgVirtServicesSecondPartTableServErrcodeStatus_Enabled     SlbCurCfgVirtServicesSecondPartTableServErrcodeStatus = 1
	SlbCurCfgVirtServicesSecondPartTableServErrcodeStatus_Disabled    SlbCurCfgVirtServicesSecondPartTableServErrcodeStatus = 2
	SlbCurCfgVirtServicesSecondPartTableServErrcodeStatus_Clear       SlbCurCfgVirtServicesSecondPartTableServErrcodeStatus = 3
	SlbCurCfgVirtServicesSecondPartTableServErrcodeStatus_Unsupported SlbCurCfgVirtServicesSecondPartTableServErrcodeStatus = 2147483647
)

type SlbCurCfgVirtServicesSecondPartTableServErrcodeHttpRedir int32

const (
	SlbCurCfgVirtServicesSecondPartTableServErrcodeHttpRedir_Yes         SlbCurCfgVirtServicesSecondPartTableServErrcodeHttpRedir = 1
	SlbCurCfgVirtServicesSecondPartTableServErrcodeHttpRedir_No          SlbCurCfgVirtServicesSecondPartTableServErrcodeHttpRedir = 2
	SlbCurCfgVirtServicesSecondPartTableServErrcodeHttpRedir_Unsupported SlbCurCfgVirtServicesSecondPartTableServErrcodeHttpRedir = 2147483647
)

type SlbCurCfgVirtServicesSecondPartTableServUrlchangStatus int32

const (
	SlbCurCfgVirtServicesSecondPartTableServUrlchangStatus_Enable      SlbCurCfgVirtServicesSecondPartTableServUrlchangStatus = 1
	SlbCurCfgVirtServicesSecondPartTableServUrlchangStatus_Disable     SlbCurCfgVirtServicesSecondPartTableServUrlchangStatus = 2
	SlbCurCfgVirtServicesSecondPartTableServUrlchangStatus_Clear       SlbCurCfgVirtServicesSecondPartTableServUrlchangStatus = 3
	SlbCurCfgVirtServicesSecondPartTableServUrlchangStatus_Unsupported SlbCurCfgVirtServicesSecondPartTableServUrlchangStatus = 2147483647
)

type SlbCurCfgVirtServicesSecondPartTableServUrlchangHostType int32

const (
	SlbCurCfgVirtServicesSecondPartTableServUrlchangHostType_Sufx        SlbCurCfgVirtServicesSecondPartTableServUrlchangHostType = 1
	SlbCurCfgVirtServicesSecondPartTableServUrlchangHostType_Prefx       SlbCurCfgVirtServicesSecondPartTableServUrlchangHostType = 2
	SlbCurCfgVirtServicesSecondPartTableServUrlchangHostType_Eq          SlbCurCfgVirtServicesSecondPartTableServUrlchangHostType = 3
	SlbCurCfgVirtServicesSecondPartTableServUrlchangHostType_Incl        SlbCurCfgVirtServicesSecondPartTableServUrlchangHostType = 4
	SlbCurCfgVirtServicesSecondPartTableServUrlchangHostType_Any         SlbCurCfgVirtServicesSecondPartTableServUrlchangHostType = 5
	SlbCurCfgVirtServicesSecondPartTableServUrlchangHostType_Unsupported SlbCurCfgVirtServicesSecondPartTableServUrlchangHostType = 2147483647
)

type SlbCurCfgVirtServicesSecondPartTableParams struct {
	// The number of the virtual server.
	ServSecondPartIndex int32 `json:"ServSecondPartIndex,omitempty"`
	// The service index. This has no external meaning
	SecondPartIndex int32 `json:"SecondPartIndex,omitempty"`
	// Connection management configuration for HTTP traffic(Enable/disable/pooling) [Default: Disable].
	ConnmgtStatus SlbCurCfgVirtServicesSecondPartTableConnmgtStatus `json:"ConnmgtStatus,omitempty"`
	// Connection management server side connection idle timeout in minutes [0-32768] [Default: 10].
	ConnmgtTimeout uint32 `json:"ConnmgtTimeout,omitempty"`
	// Cache policy name associated with this virtual service.
	Cachepol string `json:"Cachepol,omitempty"`
	// Compression policy name associated with this virtual service.
	Comppol string `json:"Comppol,omitempty"`
	// SSL policy name associated with this virtual service.
	SSLpol string `json:"SSLpol,omitempty"`
	// Server Certificate name associated with this virtual service.
	ServCert string `json:"ServCert,omitempty"`
	// HTTP Content Modifications Rule-list associated with this virtual service.
	HttpmodList string `json:"HttpmodList,omitempty"`
	// Enable/disable server cloaking.
	Cloaksrv SlbCurCfgVirtServicesSecondPartTableCloaksrv `json:"Cloaksrv,omitempty"`
	// Enable/disable/clear error-code configuration.
	ServErrcodeStatus SlbCurCfgVirtServicesSecondPartTableServErrcodeStatus `json:"ServErrcodeStatus,omitempty"`
	// Match error-code(s), e.g 203,204 .
	ServErrcodeMatch string `json:"ServErrcodeMatch,omitempty"`
	// Use http redirection [yes/no] [Default: yes].
	ServErrcodeHttpRedir SlbCurCfgVirtServicesSecondPartTableServErrcodeHttpRedir `json:"ServErrcodeHttpRedir,omitempty"`
	// URL for redirection.
	ServErrcodeUrl string `json:"ServErrcodeUrl,omitempty"`
	// Set error code [Default: 302].
	ServErrcode string `json:"ServErrcode,omitempty"`
	// Enter Cur error code [Default: 302].
	ServErrcodeNew string `json:"ServErrcodeNew,omitempty"`
	// Enter error reason.
	ServErrcodeReason string `json:"ServErrcodeReason,omitempty"`
	// Enter enabled/disabled/clear [Default: clear].
	ServUrlchangStatus SlbCurCfgVirtServicesSecondPartTableServUrlchangStatus `json:"ServUrlchangStatus,omitempty"`
	// Enter hostname match type [sufx|prefx|eq|incl|any] [Default: any]
	ServUrlchangHostType SlbCurCfgVirtServicesSecondPartTableServUrlchangHostType `json:"ServUrlchangHostType,omitempty"`
}

func (p SlbCurCfgVirtServicesSecondPartTableParams) iMABean() {}
