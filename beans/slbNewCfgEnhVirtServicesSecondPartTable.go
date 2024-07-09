package beans

import (
	"fmt"
	"reflect"
)

// SlbNewCfgEnhVirtServicesSecondPartTable The table of virtual Services.
// Note:This mib is not supported for VX instance of Virtualization.
type SlbNewCfgEnhVirtServicesSecondPartTable struct {
	// The number of the virtual server.
	SlbNewCfgEnhVirtServSecondPartIndex string
	// The service index. This has no external meaning
	SlbNewCfgEnhVirtServiceSecondPartIndex int32
	Params                                 *SlbNewCfgEnhVirtServicesSecondPartTableParams
}

func NewSlbNewCfgEnhVirtServicesSecondPartTableList() *SlbNewCfgEnhVirtServicesSecondPartTable {
	return &SlbNewCfgEnhVirtServicesSecondPartTable{}
}

func NewSlbNewCfgEnhVirtServicesSecondPartTable(
	slbNewCfgEnhVirtServSecondPartIndex string,
	slbNewCfgEnhVirtServiceSecondPartIndex int32,
	params *SlbNewCfgEnhVirtServicesSecondPartTableParams,
) *SlbNewCfgEnhVirtServicesSecondPartTable {
	return &SlbNewCfgEnhVirtServicesSecondPartTable{
		SlbNewCfgEnhVirtServSecondPartIndex:    slbNewCfgEnhVirtServSecondPartIndex,
		SlbNewCfgEnhVirtServiceSecondPartIndex: slbNewCfgEnhVirtServiceSecondPartIndex,
		Params:                                 params,
	}
}

func (c *SlbNewCfgEnhVirtServicesSecondPartTable) Name() string {
	return "SlbNewCfgEnhVirtServicesSecondPartTable"
}

func (c *SlbNewCfgEnhVirtServicesSecondPartTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbNewCfgEnhVirtServicesSecondPartTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbNewCfgEnhVirtServicesSecondPartTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbNewCfgEnhVirtServSecondPartIndex).IsZero() &&
		reflect.ValueOf(c.SlbNewCfgEnhVirtServiceSecondPartIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbNewCfgEnhVirtServSecondPartIndex) + "/" + fmt.Sprint(c.SlbNewCfgEnhVirtServiceSecondPartIndex)
}

type SlbNewCfgEnhVirtServicesSecondPartTableConnmgtStatus int32

const (
	SlbNewCfgEnhVirtServicesSecondPartTableConnmgtStatus_Disabled    SlbNewCfgEnhVirtServicesSecondPartTableConnmgtStatus = 0
	SlbNewCfgEnhVirtServicesSecondPartTableConnmgtStatus_Pooling     SlbNewCfgEnhVirtServicesSecondPartTableConnmgtStatus = 1
	SlbNewCfgEnhVirtServicesSecondPartTableConnmgtStatus_Muxenabled  SlbNewCfgEnhVirtServicesSecondPartTableConnmgtStatus = 2
	SlbNewCfgEnhVirtServicesSecondPartTableConnmgtStatus_H2          SlbNewCfgEnhVirtServicesSecondPartTableConnmgtStatus = 3
	SlbNewCfgEnhVirtServicesSecondPartTableConnmgtStatus_Unsupported SlbNewCfgEnhVirtServicesSecondPartTableConnmgtStatus = 2147483647
)

type SlbNewCfgEnhVirtServicesSecondPartTableCloaksrv int32

const (
	SlbNewCfgEnhVirtServicesSecondPartTableCloaksrv_Enabled     SlbNewCfgEnhVirtServicesSecondPartTableCloaksrv = 1
	SlbNewCfgEnhVirtServicesSecondPartTableCloaksrv_Disabled    SlbNewCfgEnhVirtServicesSecondPartTableCloaksrv = 2
	SlbNewCfgEnhVirtServicesSecondPartTableCloaksrv_Unsupported SlbNewCfgEnhVirtServicesSecondPartTableCloaksrv = 2147483647
)

type SlbNewCfgEnhVirtServicesSecondPartTableServErrcodeStatus int32

const (
	SlbNewCfgEnhVirtServicesSecondPartTableServErrcodeStatus_Enabled     SlbNewCfgEnhVirtServicesSecondPartTableServErrcodeStatus = 1
	SlbNewCfgEnhVirtServicesSecondPartTableServErrcodeStatus_Disabled    SlbNewCfgEnhVirtServicesSecondPartTableServErrcodeStatus = 2
	SlbNewCfgEnhVirtServicesSecondPartTableServErrcodeStatus_Clear       SlbNewCfgEnhVirtServicesSecondPartTableServErrcodeStatus = 3
	SlbNewCfgEnhVirtServicesSecondPartTableServErrcodeStatus_Unsupported SlbNewCfgEnhVirtServicesSecondPartTableServErrcodeStatus = 2147483647
)

type SlbNewCfgEnhVirtServicesSecondPartTableServErrcodeHttpRedir int32

const (
	SlbNewCfgEnhVirtServicesSecondPartTableServErrcodeHttpRedir_Yes         SlbNewCfgEnhVirtServicesSecondPartTableServErrcodeHttpRedir = 1
	SlbNewCfgEnhVirtServicesSecondPartTableServErrcodeHttpRedir_No          SlbNewCfgEnhVirtServicesSecondPartTableServErrcodeHttpRedir = 2
	SlbNewCfgEnhVirtServicesSecondPartTableServErrcodeHttpRedir_Unsupported SlbNewCfgEnhVirtServicesSecondPartTableServErrcodeHttpRedir = 2147483647
)

type SlbNewCfgEnhVirtServicesSecondPartTableServUrlchangStatus int32

const (
	SlbNewCfgEnhVirtServicesSecondPartTableServUrlchangStatus_Enable      SlbNewCfgEnhVirtServicesSecondPartTableServUrlchangStatus = 1
	SlbNewCfgEnhVirtServicesSecondPartTableServUrlchangStatus_Disable     SlbNewCfgEnhVirtServicesSecondPartTableServUrlchangStatus = 2
	SlbNewCfgEnhVirtServicesSecondPartTableServUrlchangStatus_Clear       SlbNewCfgEnhVirtServicesSecondPartTableServUrlchangStatus = 3
	SlbNewCfgEnhVirtServicesSecondPartTableServUrlchangStatus_Unsupported SlbNewCfgEnhVirtServicesSecondPartTableServUrlchangStatus = 2147483647
)

type SlbNewCfgEnhVirtServicesSecondPartTableServUrlchangHostType int32

const (
	SlbNewCfgEnhVirtServicesSecondPartTableServUrlchangHostType_Sufx        SlbNewCfgEnhVirtServicesSecondPartTableServUrlchangHostType = 1
	SlbNewCfgEnhVirtServicesSecondPartTableServUrlchangHostType_Prefx       SlbNewCfgEnhVirtServicesSecondPartTableServUrlchangHostType = 2
	SlbNewCfgEnhVirtServicesSecondPartTableServUrlchangHostType_Eq          SlbNewCfgEnhVirtServicesSecondPartTableServUrlchangHostType = 3
	SlbNewCfgEnhVirtServicesSecondPartTableServUrlchangHostType_Incl        SlbNewCfgEnhVirtServicesSecondPartTableServUrlchangHostType = 4
	SlbNewCfgEnhVirtServicesSecondPartTableServUrlchangHostType_Any         SlbNewCfgEnhVirtServicesSecondPartTableServUrlchangHostType = 5
	SlbNewCfgEnhVirtServicesSecondPartTableServUrlchangHostType_Unsupported SlbNewCfgEnhVirtServicesSecondPartTableServUrlchangHostType = 2147483647
)

type SlbNewCfgEnhVirtServicesSecondPartTableBasicConnmgtStatus int32

const (
	SlbNewCfgEnhVirtServicesSecondPartTableBasicConnmgtStatus_Disabled    SlbNewCfgEnhVirtServicesSecondPartTableBasicConnmgtStatus = 0
	SlbNewCfgEnhVirtServicesSecondPartTableBasicConnmgtStatus_Muxenabled  SlbNewCfgEnhVirtServicesSecondPartTableBasicConnmgtStatus = 2
	SlbNewCfgEnhVirtServicesSecondPartTableBasicConnmgtStatus_Unsupported SlbNewCfgEnhVirtServicesSecondPartTableBasicConnmgtStatus = 2147483647
)

type SlbNewCfgEnhVirtServicesSecondPartTableParams struct {
	// The number of the virtual server.
	ServSecondPartIndex string `json:"ServSecondPartIndex,omitempty"`
	// The service index. This has no external meaning
	SecondPartIndex int32 `json:"SecondPartIndex,omitempty"`
	// Connection management configuration for HTTP traffic(Enable/disable/pooling) [Default: Disable].
	ConnmgtStatus SlbNewCfgEnhVirtServicesSecondPartTableConnmgtStatus `json:"ConnmgtStatus,omitempty"`
	// Connection management server side connection idle timeout in minutes [0-32768] [Default: 10].
	ConnmgtTimeout uint32 `json:"ConnmgtTimeout,omitempty"`
	// Cache policy name associated with this virtual service.Set none to delete entry
	Cachepol string `json:"Cachepol,omitempty"`
	// Compression policy name associated with this virtual service.Set none to delete entry
	Comppol string `json:"Comppol,omitempty"`
	// SSL policy name associated with this virtual service.Set none to delete entry
	SSLpol string `json:"SSLpol,omitempty"`
	// Server Certificate name associated with this virtual service.
	ServCert string `json:"ServCert,omitempty"`
	// HTTP Content Modifications Rule-list associated with this virtual service.Set none to delete entry
	HttpmodList string `json:"HttpmodList,omitempty"`
	// Enable/disable server cloaking.
	Cloaksrv SlbNewCfgEnhVirtServicesSecondPartTableCloaksrv `json:"Cloaksrv,omitempty"`
	// Enable/disable/clear error-code configuration.
	ServErrcodeStatus SlbNewCfgEnhVirtServicesSecondPartTableServErrcodeStatus `json:"ServErrcodeStatus,omitempty"`
	// Match error-code(s), e.g 203,204 .
	ServErrcodeMatch string `json:"ServErrcodeMatch,omitempty"`
	// Use http redirection [yes/no] [Default: yes].
	ServErrcodeHttpRedir SlbNewCfgEnhVirtServicesSecondPartTableServErrcodeHttpRedir `json:"ServErrcodeHttpRedir,omitempty"`
	// URL for redirection.
	ServErrcodeUrl string `json:"ServErrcodeUrl,omitempty"`
	// set error code [Default: 302].
	ServErrcode string `json:"ServErrcode,omitempty"`
	// Enter new error code [Default: 302].
	ServErrcodeNew string `json:"ServErrcodeNew,omitempty"`
	// Enter error reason.
	ServErrcodeReason string `json:"ServErrcodeReason,omitempty"`
	// Enter enabled/disabled/clear [Default: clear].
	ServUrlchangStatus SlbNewCfgEnhVirtServicesSecondPartTableServUrlchangStatus `json:"ServUrlchangStatus,omitempty"`
	// Enter hostname match type [sufx|prefx|eq|incl|any] [Default: any]
	ServUrlchangHostType SlbNewCfgEnhVirtServicesSecondPartTableServUrlchangHostType `json:"ServUrlchangHostType,omitempty"`
	// Frontend TCP optimization policy.
	FeTcpPolId string `json:"FeTcpPolId,omitempty"`
	// Backend TCP optimization policy.
	BeTcpPolId string `json:"BeTcpPolId,omitempty"`
	// Connection management configuration for Tcp traffic(Enable/disable) [Default: Disable].
	BasicConnmgtStatus SlbNewCfgEnhVirtServicesSecondPartTableBasicConnmgtStatus `json:"BasicConnmgtStatus,omitempty"`
	// GM SSL Server encryption Certificate name associated with this virtual service.
	ServCertEnc string `json:"ServCertEnc,omitempty"`
	// GM SSL Server sign Certificate name associated with this virtual service.
	ServCertSign string `json:"ServCertSign,omitempty"`
}

func (p SlbNewCfgEnhVirtServicesSecondPartTableParams) iMABean() {}
