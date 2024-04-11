package beans

import (
	"fmt"
	"reflect"
)

// SlbCurCfgEnhVirtServicesSecondPartTable The table of virtual Services.
// Note:This mib is not supported for VX instance of Virtualization.
type SlbCurCfgEnhVirtServicesSecondPartTable struct {
	// The number of the virtual server.
	SlbCurCfgEnhVirtServSecondPartIndex string
	// The service index. This has no external meaning
	SlbCurCfgEnhVirtServiceSecondPartIndex int32
	Params                                 *SlbCurCfgEnhVirtServicesSecondPartTableParams
}

func NewSlbCurCfgEnhVirtServicesSecondPartTableList() *SlbCurCfgEnhVirtServicesSecondPartTable {
	return &SlbCurCfgEnhVirtServicesSecondPartTable{}
}

func NewSlbCurCfgEnhVirtServicesSecondPartTable(
	slbCurCfgEnhVirtServSecondPartIndex string,
	slbCurCfgEnhVirtServiceSecondPartIndex int32,
	params *SlbCurCfgEnhVirtServicesSecondPartTableParams,
) *SlbCurCfgEnhVirtServicesSecondPartTable {
	return &SlbCurCfgEnhVirtServicesSecondPartTable{
		SlbCurCfgEnhVirtServSecondPartIndex:    slbCurCfgEnhVirtServSecondPartIndex,
		SlbCurCfgEnhVirtServiceSecondPartIndex: slbCurCfgEnhVirtServiceSecondPartIndex,
		Params:                                 params,
	}
}

func (c *SlbCurCfgEnhVirtServicesSecondPartTable) Name() string {
	return "SlbCurCfgEnhVirtServicesSecondPartTable"
}

func (c *SlbCurCfgEnhVirtServicesSecondPartTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbCurCfgEnhVirtServicesSecondPartTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbCurCfgEnhVirtServicesSecondPartTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbCurCfgEnhVirtServSecondPartIndex).IsZero() &&
		reflect.ValueOf(c.SlbCurCfgEnhVirtServiceSecondPartIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbCurCfgEnhVirtServSecondPartIndex) + "/" + fmt.Sprint(c.SlbCurCfgEnhVirtServiceSecondPartIndex)
}

type SlbCurCfgEnhVirtServicesSecondPartTableConnmgtStatus int32

const (
	SlbCurCfgEnhVirtServicesSecondPartTableConnmgtStatus_Disabled    SlbCurCfgEnhVirtServicesSecondPartTableConnmgtStatus = 0
	SlbCurCfgEnhVirtServicesSecondPartTableConnmgtStatus_Pooling     SlbCurCfgEnhVirtServicesSecondPartTableConnmgtStatus = 1
	SlbCurCfgEnhVirtServicesSecondPartTableConnmgtStatus_Muxenabled  SlbCurCfgEnhVirtServicesSecondPartTableConnmgtStatus = 2
	SlbCurCfgEnhVirtServicesSecondPartTableConnmgtStatus_H2          SlbCurCfgEnhVirtServicesSecondPartTableConnmgtStatus = 3
	SlbCurCfgEnhVirtServicesSecondPartTableConnmgtStatus_Unsupported SlbCurCfgEnhVirtServicesSecondPartTableConnmgtStatus = 2147483647
)

type SlbCurCfgEnhVirtServicesSecondPartTableCloaksrv int32

const (
	SlbCurCfgEnhVirtServicesSecondPartTableCloaksrv_Enabled     SlbCurCfgEnhVirtServicesSecondPartTableCloaksrv = 1
	SlbCurCfgEnhVirtServicesSecondPartTableCloaksrv_Disabled    SlbCurCfgEnhVirtServicesSecondPartTableCloaksrv = 2
	SlbCurCfgEnhVirtServicesSecondPartTableCloaksrv_Unsupported SlbCurCfgEnhVirtServicesSecondPartTableCloaksrv = 2147483647
)

type SlbCurCfgEnhVirtServicesSecondPartTableServErrcodeStatus int32

const (
	SlbCurCfgEnhVirtServicesSecondPartTableServErrcodeStatus_Enabled     SlbCurCfgEnhVirtServicesSecondPartTableServErrcodeStatus = 1
	SlbCurCfgEnhVirtServicesSecondPartTableServErrcodeStatus_Disabled    SlbCurCfgEnhVirtServicesSecondPartTableServErrcodeStatus = 2
	SlbCurCfgEnhVirtServicesSecondPartTableServErrcodeStatus_Clear       SlbCurCfgEnhVirtServicesSecondPartTableServErrcodeStatus = 3
	SlbCurCfgEnhVirtServicesSecondPartTableServErrcodeStatus_Unsupported SlbCurCfgEnhVirtServicesSecondPartTableServErrcodeStatus = 2147483647
)

type SlbCurCfgEnhVirtServicesSecondPartTableServErrcodeHttpRedir int32

const (
	SlbCurCfgEnhVirtServicesSecondPartTableServErrcodeHttpRedir_Yes         SlbCurCfgEnhVirtServicesSecondPartTableServErrcodeHttpRedir = 1
	SlbCurCfgEnhVirtServicesSecondPartTableServErrcodeHttpRedir_No          SlbCurCfgEnhVirtServicesSecondPartTableServErrcodeHttpRedir = 2
	SlbCurCfgEnhVirtServicesSecondPartTableServErrcodeHttpRedir_Unsupported SlbCurCfgEnhVirtServicesSecondPartTableServErrcodeHttpRedir = 2147483647
)

type SlbCurCfgEnhVirtServicesSecondPartTableServUrlchangStatus int32

const (
	SlbCurCfgEnhVirtServicesSecondPartTableServUrlchangStatus_Enable      SlbCurCfgEnhVirtServicesSecondPartTableServUrlchangStatus = 1
	SlbCurCfgEnhVirtServicesSecondPartTableServUrlchangStatus_Disable     SlbCurCfgEnhVirtServicesSecondPartTableServUrlchangStatus = 2
	SlbCurCfgEnhVirtServicesSecondPartTableServUrlchangStatus_Clear       SlbCurCfgEnhVirtServicesSecondPartTableServUrlchangStatus = 3
	SlbCurCfgEnhVirtServicesSecondPartTableServUrlchangStatus_Unsupported SlbCurCfgEnhVirtServicesSecondPartTableServUrlchangStatus = 2147483647
)

type SlbCurCfgEnhVirtServicesSecondPartTableServUrlchangHostType int32

const (
	SlbCurCfgEnhVirtServicesSecondPartTableServUrlchangHostType_Sufx        SlbCurCfgEnhVirtServicesSecondPartTableServUrlchangHostType = 1
	SlbCurCfgEnhVirtServicesSecondPartTableServUrlchangHostType_Prefx       SlbCurCfgEnhVirtServicesSecondPartTableServUrlchangHostType = 2
	SlbCurCfgEnhVirtServicesSecondPartTableServUrlchangHostType_Eq          SlbCurCfgEnhVirtServicesSecondPartTableServUrlchangHostType = 3
	SlbCurCfgEnhVirtServicesSecondPartTableServUrlchangHostType_Incl        SlbCurCfgEnhVirtServicesSecondPartTableServUrlchangHostType = 4
	SlbCurCfgEnhVirtServicesSecondPartTableServUrlchangHostType_Any         SlbCurCfgEnhVirtServicesSecondPartTableServUrlchangHostType = 5
	SlbCurCfgEnhVirtServicesSecondPartTableServUrlchangHostType_Unsupported SlbCurCfgEnhVirtServicesSecondPartTableServUrlchangHostType = 2147483647
)

type SlbCurCfgEnhVirtServicesSecondPartTableParams struct {
	// The number of the virtual server.
	ServSecondPartIndex string `json:"ServSecondPartIndex,omitempty"`
	// The service index. This has no external meaning
	SecondPartIndex int32 `json:"SecondPartIndex,omitempty"`
	// Connection management configuration for HTTP traffic(Enable/disable/pooling) [Default: Disable].
	ConnmgtStatus SlbCurCfgEnhVirtServicesSecondPartTableConnmgtStatus `json:"ConnmgtStatus,omitempty"`
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
	Cloaksrv SlbCurCfgEnhVirtServicesSecondPartTableCloaksrv `json:"Cloaksrv,omitempty"`
	// Enable/disable/clear error-code configuration.
	ServErrcodeStatus SlbCurCfgEnhVirtServicesSecondPartTableServErrcodeStatus `json:"ServErrcodeStatus,omitempty"`
	// Match error-code(s), e.g 203,204 .
	ServErrcodeMatch string `json:"ServErrcodeMatch,omitempty"`
	// Use http redirection [yes/no] [Default: yes].
	ServErrcodeHttpRedir SlbCurCfgEnhVirtServicesSecondPartTableServErrcodeHttpRedir `json:"ServErrcodeHttpRedir,omitempty"`
	// URL for redirection.
	ServErrcodeUrl string `json:"ServErrcodeUrl,omitempty"`
	// Set error code [Default: 302].
	ServErrcode string `json:"ServErrcode,omitempty"`
	// Enter Cur error code [Default: 302].
	ServErrcodeNew string `json:"ServErrcodeNew,omitempty"`
	// Enter error reason.
	ServErrcodeReason string `json:"ServErrcodeReason,omitempty"`
	// Enter enabled/disabled/clear [Default: clear].
	ServUrlchangStatus SlbCurCfgEnhVirtServicesSecondPartTableServUrlchangStatus `json:"ServUrlchangStatus,omitempty"`
	// Enter hostname match type [sufx|prefx|eq|incl|any] [Default: any]
	ServUrlchangHostType SlbCurCfgEnhVirtServicesSecondPartTableServUrlchangHostType `json:"ServUrlchangHostType,omitempty"`
	// Frontend TCP optimization policy.
	FeTcpPolId string `json:"FeTcpPolId,omitempty"`
	// Backend TCP optimization policy.
	BeTcpPolId string `json:"BeTcpPolId,omitempty"`
}

func (p SlbCurCfgEnhVirtServicesSecondPartTableParams) iMABean() {}
