package beans

import (
	"fmt"
	"reflect"
)

// SlbNewCfgVirtServicesSecondPartTable The table of virtual Services.
// Note:This mib is not supported for VX instance of Virtualization.
type SlbNewCfgVirtServicesSecondPartTable struct {
	// The number of the virtual server.
	SlbNewCfgVirtServSecondPartIndex int32
	// The service index. This has no external meaning
	SlbNewCfgVirtServiceSecondPartIndex int32
	Params                              *SlbNewCfgVirtServicesSecondPartTableParams
}

func NewSlbNewCfgVirtServicesSecondPartTableList() *SlbNewCfgVirtServicesSecondPartTable {
	return &SlbNewCfgVirtServicesSecondPartTable{}
}

func NewSlbNewCfgVirtServicesSecondPartTable(
	slbNewCfgVirtServSecondPartIndex int32,
	slbNewCfgVirtServiceSecondPartIndex int32,
	params *SlbNewCfgVirtServicesSecondPartTableParams,
) *SlbNewCfgVirtServicesSecondPartTable {
	return &SlbNewCfgVirtServicesSecondPartTable{
		SlbNewCfgVirtServSecondPartIndex:    slbNewCfgVirtServSecondPartIndex,
		SlbNewCfgVirtServiceSecondPartIndex: slbNewCfgVirtServiceSecondPartIndex,
		Params:                              params,
	}
}

func (c *SlbNewCfgVirtServicesSecondPartTable) Name() string {
	return "SlbNewCfgVirtServicesSecondPartTable"
}

func (c *SlbNewCfgVirtServicesSecondPartTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbNewCfgVirtServicesSecondPartTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbNewCfgVirtServicesSecondPartTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbNewCfgVirtServSecondPartIndex).IsZero() &&
		reflect.ValueOf(c.SlbNewCfgVirtServiceSecondPartIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbNewCfgVirtServSecondPartIndex) + "/" + fmt.Sprint(c.SlbNewCfgVirtServiceSecondPartIndex)
}

type SlbNewCfgVirtServicesSecondPartTableConnmgtStatus int32

const (
	SlbNewCfgVirtServicesSecondPartTableConnmgtStatus_Disabled    SlbNewCfgVirtServicesSecondPartTableConnmgtStatus = 0
	SlbNewCfgVirtServicesSecondPartTableConnmgtStatus_Pooling     SlbNewCfgVirtServicesSecondPartTableConnmgtStatus = 1
	SlbNewCfgVirtServicesSecondPartTableConnmgtStatus_Muxenabled  SlbNewCfgVirtServicesSecondPartTableConnmgtStatus = 2
	SlbNewCfgVirtServicesSecondPartTableConnmgtStatus_Unsupported SlbNewCfgVirtServicesSecondPartTableConnmgtStatus = 2147483647
)

type SlbNewCfgVirtServicesSecondPartTableCloaksrv int32

const (
	SlbNewCfgVirtServicesSecondPartTableCloaksrv_Enabled     SlbNewCfgVirtServicesSecondPartTableCloaksrv = 1
	SlbNewCfgVirtServicesSecondPartTableCloaksrv_Disabled    SlbNewCfgVirtServicesSecondPartTableCloaksrv = 2
	SlbNewCfgVirtServicesSecondPartTableCloaksrv_Unsupported SlbNewCfgVirtServicesSecondPartTableCloaksrv = 2147483647
)

type SlbNewCfgVirtServicesSecondPartTableServErrcodeStatus int32

const (
	SlbNewCfgVirtServicesSecondPartTableServErrcodeStatus_Enabled     SlbNewCfgVirtServicesSecondPartTableServErrcodeStatus = 1
	SlbNewCfgVirtServicesSecondPartTableServErrcodeStatus_Disabled    SlbNewCfgVirtServicesSecondPartTableServErrcodeStatus = 2
	SlbNewCfgVirtServicesSecondPartTableServErrcodeStatus_Clear       SlbNewCfgVirtServicesSecondPartTableServErrcodeStatus = 3
	SlbNewCfgVirtServicesSecondPartTableServErrcodeStatus_Unsupported SlbNewCfgVirtServicesSecondPartTableServErrcodeStatus = 2147483647
)

type SlbNewCfgVirtServicesSecondPartTableServErrcodeHttpRedir int32

const (
	SlbNewCfgVirtServicesSecondPartTableServErrcodeHttpRedir_Yes         SlbNewCfgVirtServicesSecondPartTableServErrcodeHttpRedir = 1
	SlbNewCfgVirtServicesSecondPartTableServErrcodeHttpRedir_No          SlbNewCfgVirtServicesSecondPartTableServErrcodeHttpRedir = 2
	SlbNewCfgVirtServicesSecondPartTableServErrcodeHttpRedir_Unsupported SlbNewCfgVirtServicesSecondPartTableServErrcodeHttpRedir = 2147483647
)

type SlbNewCfgVirtServicesSecondPartTableServUrlchangStatus int32

const (
	SlbNewCfgVirtServicesSecondPartTableServUrlchangStatus_Enable      SlbNewCfgVirtServicesSecondPartTableServUrlchangStatus = 1
	SlbNewCfgVirtServicesSecondPartTableServUrlchangStatus_Disable     SlbNewCfgVirtServicesSecondPartTableServUrlchangStatus = 2
	SlbNewCfgVirtServicesSecondPartTableServUrlchangStatus_Clear       SlbNewCfgVirtServicesSecondPartTableServUrlchangStatus = 3
	SlbNewCfgVirtServicesSecondPartTableServUrlchangStatus_Unsupported SlbNewCfgVirtServicesSecondPartTableServUrlchangStatus = 2147483647
)

type SlbNewCfgVirtServicesSecondPartTableServUrlchangHostType int32

const (
	SlbNewCfgVirtServicesSecondPartTableServUrlchangHostType_Sufx        SlbNewCfgVirtServicesSecondPartTableServUrlchangHostType = 1
	SlbNewCfgVirtServicesSecondPartTableServUrlchangHostType_Prefx       SlbNewCfgVirtServicesSecondPartTableServUrlchangHostType = 2
	SlbNewCfgVirtServicesSecondPartTableServUrlchangHostType_Eq          SlbNewCfgVirtServicesSecondPartTableServUrlchangHostType = 3
	SlbNewCfgVirtServicesSecondPartTableServUrlchangHostType_Incl        SlbNewCfgVirtServicesSecondPartTableServUrlchangHostType = 4
	SlbNewCfgVirtServicesSecondPartTableServUrlchangHostType_Any         SlbNewCfgVirtServicesSecondPartTableServUrlchangHostType = 5
	SlbNewCfgVirtServicesSecondPartTableServUrlchangHostType_Unsupported SlbNewCfgVirtServicesSecondPartTableServUrlchangHostType = 2147483647
)

type SlbNewCfgVirtServicesSecondPartTableParams struct {
	// The number of the virtual server.
	ServSecondPartIndex int32 `json:"ServSecondPartIndex,omitempty"`
	// The service index. This has no external meaning
	SecondPartIndex int32 `json:"SecondPartIndex,omitempty"`
	// Connection management configuration for HTTP traffic(Enable/disable/pooling) [Default: Disable].
	ConnmgtStatus SlbNewCfgVirtServicesSecondPartTableConnmgtStatus `json:"ConnmgtStatus,omitempty"`
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
	Cloaksrv SlbNewCfgVirtServicesSecondPartTableCloaksrv `json:"Cloaksrv,omitempty"`
	// Enable/disable/clear error-code configuration.
	ServErrcodeStatus SlbNewCfgVirtServicesSecondPartTableServErrcodeStatus `json:"ServErrcodeStatus,omitempty"`
	// Match error-code(s), e.g 203,204 .
	ServErrcodeMatch string `json:"ServErrcodeMatch,omitempty"`
	// Use http redirection [yes/no] [Default: yes].
	ServErrcodeHttpRedir SlbNewCfgVirtServicesSecondPartTableServErrcodeHttpRedir `json:"ServErrcodeHttpRedir,omitempty"`
	// URL for redirection.
	ServErrcodeUrl string `json:"ServErrcodeUrl,omitempty"`
	// set error code [Default: 302].
	ServErrcode string `json:"ServErrcode,omitempty"`
	// Enter new error code [Default: 302].
	ServErrcodeNew string `json:"ServErrcodeNew,omitempty"`
	// Enter error reason.
	ServErrcodeReason string `json:"ServErrcodeReason,omitempty"`
	// Enter enabled/disabled/clear [Default: clear].
	ServUrlchangStatus SlbNewCfgVirtServicesSecondPartTableServUrlchangStatus `json:"ServUrlchangStatus,omitempty"`
	// Enter hostname match type [sufx|prefx|eq|incl|any] [Default: any]
	ServUrlchangHostType SlbNewCfgVirtServicesSecondPartTableServUrlchangHostType `json:"ServUrlchangHostType,omitempty"`
}

func (p SlbNewCfgVirtServicesSecondPartTableParams) iMABean() {}
