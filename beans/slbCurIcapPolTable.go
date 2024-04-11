package beans

import (
	"fmt"
	"reflect"
)

// SlbCurIcapPolTable The table for configuring icap policy.
// Note:This mib is not supported for VX instance of Virtualization.
type SlbCurIcapPolTable struct {
	// The Icap policy ID(key id) as an index.
	SlbCurIcapPolNameIdIndex string
	Params                   *SlbCurIcapPolTableParams
}

func NewSlbCurIcapPolTableList() *SlbCurIcapPolTable {
	return &SlbCurIcapPolTable{}
}

func NewSlbCurIcapPolTable(
	slbCurIcapPolNameIdIndex string,
	params *SlbCurIcapPolTableParams,
) *SlbCurIcapPolTable {
	return &SlbCurIcapPolTable{
		SlbCurIcapPolNameIdIndex: slbCurIcapPolNameIdIndex,
		Params:                   params,
	}
}

func (c *SlbCurIcapPolTable) Name() string {
	return "SlbCurIcapPolTable"
}

func (c *SlbCurIcapPolTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbCurIcapPolTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbCurIcapPolTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbCurIcapPolNameIdIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbCurIcapPolNameIdIndex)
}

type SlbCurIcapPolTableAdminStatus int32

const (
	SlbCurIcapPolTableAdminStatus_Enabled     SlbCurIcapPolTableAdminStatus = 1
	SlbCurIcapPolTableAdminStatus_Disabled    SlbCurIcapPolTableAdminStatus = 2
	SlbCurIcapPolTableAdminStatus_Unsupported SlbCurIcapPolTableAdminStatus = 2147483647
)

type SlbCurIcapPolTableReqmodStatus int32

const (
	SlbCurIcapPolTableReqmodStatus_Enabled     SlbCurIcapPolTableReqmodStatus = 1
	SlbCurIcapPolTableReqmodStatus_Disabled    SlbCurIcapPolTableReqmodStatus = 2
	SlbCurIcapPolTableReqmodStatus_Unsupported SlbCurIcapPolTableReqmodStatus = 2147483647
)

type SlbCurIcapPolTableReqmodFallback int32

const (
	SlbCurIcapPolTableReqmodFallback_Drop        SlbCurIcapPolTableReqmodFallback = 1
	SlbCurIcapPolTableReqmodFallback_ContFlow    SlbCurIcapPolTableReqmodFallback = 2
	SlbCurIcapPolTableReqmodFallback_Unsupported SlbCurIcapPolTableReqmodFallback = 2147483647
)

type SlbCurIcapPolTableReqmodXcip int32

const (
	SlbCurIcapPolTableReqmodXcip_Cip         SlbCurIcapPolTableReqmodXcip = 1
	SlbCurIcapPolTableReqmodXcip_Xff         SlbCurIcapPolTableReqmodXcip = 2
	SlbCurIcapPolTableReqmodXcip_Disabled    SlbCurIcapPolTableReqmodXcip = 3
	SlbCurIcapPolTableReqmodXcip_Unsupported SlbCurIcapPolTableReqmodXcip = 2147483647
)

type SlbCurIcapPolTableRespmodStatus int32

const (
	SlbCurIcapPolTableRespmodStatus_Enabled     SlbCurIcapPolTableRespmodStatus = 1
	SlbCurIcapPolTableRespmodStatus_Disabled    SlbCurIcapPolTableRespmodStatus = 2
	SlbCurIcapPolTableRespmodStatus_Unsupported SlbCurIcapPolTableRespmodStatus = 2147483647
)

type SlbCurIcapPolTableRespmodFallback int32

const (
	SlbCurIcapPolTableRespmodFallback_Drop        SlbCurIcapPolTableRespmodFallback = 1
	SlbCurIcapPolTableRespmodFallback_ContFlow    SlbCurIcapPolTableRespmodFallback = 2
	SlbCurIcapPolTableRespmodFallback_Unsupported SlbCurIcapPolTableRespmodFallback = 2147483647
)

type SlbCurIcapPolTableRespmodXcip int32

const (
	SlbCurIcapPolTableRespmodXcip_Cip         SlbCurIcapPolTableRespmodXcip = 1
	SlbCurIcapPolTableRespmodXcip_Xff         SlbCurIcapPolTableRespmodXcip = 2
	SlbCurIcapPolTableRespmodXcip_Disabled    SlbCurIcapPolTableRespmodXcip = 3
	SlbCurIcapPolTableRespmodXcip_Unsupported SlbCurIcapPolTableRespmodXcip = 2147483647
)

type SlbCurIcapPolTableParams struct {
	// The Icap policy ID(key id) as an index.
	NameIdIndex string `json:"NameIdIndex,omitempty"`
	// Icap policy name.
	Name string `json:"Name,omitempty"`
	// Status (enable/disable) of icap policy.
	AdminStatus SlbCurIcapPolTableAdminStatus `json:"AdminStatus,omitempty"`
	// Status (enable/disable) of icap policy.
	ReqmodStatus SlbCurIcapPolTableReqmodStatus `json:"ReqmodStatus,omitempty"`
	// Icap Uri string.
	ReqmodUri string `json:"ReqmodUri,omitempty"`
	// The Icap Reqmod  server group number.
	ReqmodGroup string `json:"ReqmodGroup,omitempty"`
	// Fallback action (group down).
	ReqmodFallback SlbCurIcapPolTableReqmodFallback `json:"ReqmodFallback,omitempty"`
	// The maximum preview number in bytes.
	ReqmodPreview uint32 `json:"ReqmodPreview,omitempty"`
	// Icap From string.
	ReqmodFrom string `json:"ReqmodFrom,omitempty"`
	// Icap Host string.
	ReqmodHost string `json:"ReqmodHost,omitempty"`
	// Icap Referer string.
	ReqmodReferer string `json:"ReqmodReferer,omitempty"`
	// Icap User-Agent string.
	ReqmodAgent string `json:"ReqmodAgent,omitempty"`
	// Icap X-Client-IP status.
	ReqmodXcip SlbCurIcapPolTableReqmodXcip `json:"ReqmodXcip,omitempty"`
	// Icap xcip header string.
	ReqmodXcipHeader string `json:"ReqmodXcipHeader,omitempty"`
	// Status (enable/disable) of icap policy.
	RespmodStatus SlbCurIcapPolTableRespmodStatus `json:"RespmodStatus,omitempty"`
	// Icap Uri string.
	RespmodUri string `json:"RespmodUri,omitempty"`
	// The Icap Respmod  server group number.
	RespmodGroup string `json:"RespmodGroup,omitempty"`
	// Fallback action (group down).
	RespmodFallback SlbCurIcapPolTableRespmodFallback `json:"RespmodFallback,omitempty"`
	// The maximum preview number in bytes.
	RespmodPreview uint32 `json:"RespmodPreview,omitempty"`
	// Icap From string.
	RespmodFrom string `json:"RespmodFrom,omitempty"`
	// Icap Host string.
	RespmodHost string `json:"RespmodHost,omitempty"`
	// Icap Referer string.
	RespmodReferer string `json:"RespmodReferer,omitempty"`
	// Icap User-Agent string.
	RespmodAgent string `json:"RespmodAgent,omitempty"`
	// Icap X-Client-IP status.
	RespmodXcip SlbCurIcapPolTableRespmodXcip `json:"RespmodXcip,omitempty"`
	// Icap xcip header string.
	RespmodXcipHeader string `json:"RespmodXcipHeader,omitempty"`
}

func (p SlbCurIcapPolTableParams) iMABean() {}
