package beans

import (
	"fmt"
	"reflect"
)

// SlbNewIcapPolTable The table for configuring icap policy.
// Note:This mib is not supported for VX instance of Virtualization.
type SlbNewIcapPolTable struct {
	// The Icap policy ID(key id) as an index.
	SlbNewIcapPolNameIdIndex string
	Params                   *SlbNewIcapPolTableParams
}

func NewSlbNewIcapPolTableList() *SlbNewIcapPolTable {
	return &SlbNewIcapPolTable{}
}

func NewSlbNewIcapPolTable(
	slbNewIcapPolNameIdIndex string,
	params *SlbNewIcapPolTableParams,
) *SlbNewIcapPolTable {
	return &SlbNewIcapPolTable{
		SlbNewIcapPolNameIdIndex: slbNewIcapPolNameIdIndex,
		Params:                   params,
	}
}

func (c *SlbNewIcapPolTable) Name() string {
	return "SlbNewIcapPolTable"
}

func (c *SlbNewIcapPolTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbNewIcapPolTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbNewIcapPolTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbNewIcapPolNameIdIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbNewIcapPolNameIdIndex)
}

type SlbNewIcapPolTableAdminStatus int32

const (
	SlbNewIcapPolTableAdminStatus_Enabled     SlbNewIcapPolTableAdminStatus = 1
	SlbNewIcapPolTableAdminStatus_Disabled    SlbNewIcapPolTableAdminStatus = 2
	SlbNewIcapPolTableAdminStatus_Unsupported SlbNewIcapPolTableAdminStatus = 2147483647
)

type SlbNewIcapPolTableDel int32

const (
	SlbNewIcapPolTableDel_Other       SlbNewIcapPolTableDel = 1
	SlbNewIcapPolTableDel_Delete      SlbNewIcapPolTableDel = 2
	SlbNewIcapPolTableDel_Unsupported SlbNewIcapPolTableDel = 2147483647
)

type SlbNewIcapPolTableReqmodStatus int32

const (
	SlbNewIcapPolTableReqmodStatus_Enabled     SlbNewIcapPolTableReqmodStatus = 1
	SlbNewIcapPolTableReqmodStatus_Disabled    SlbNewIcapPolTableReqmodStatus = 2
	SlbNewIcapPolTableReqmodStatus_Unsupported SlbNewIcapPolTableReqmodStatus = 2147483647
)

type SlbNewIcapPolTableReqmodFallback int32

const (
	SlbNewIcapPolTableReqmodFallback_Drop        SlbNewIcapPolTableReqmodFallback = 1
	SlbNewIcapPolTableReqmodFallback_ContFlow    SlbNewIcapPolTableReqmodFallback = 2
	SlbNewIcapPolTableReqmodFallback_Unsupported SlbNewIcapPolTableReqmodFallback = 2147483647
)

type SlbNewIcapPolTableReqmodXcip int32

const (
	SlbNewIcapPolTableReqmodXcip_Cip         SlbNewIcapPolTableReqmodXcip = 1
	SlbNewIcapPolTableReqmodXcip_Xff         SlbNewIcapPolTableReqmodXcip = 2
	SlbNewIcapPolTableReqmodXcip_Disabled    SlbNewIcapPolTableReqmodXcip = 3
	SlbNewIcapPolTableReqmodXcip_Unsupported SlbNewIcapPolTableReqmodXcip = 2147483647
)

type SlbNewIcapPolTableRespmodStatus int32

const (
	SlbNewIcapPolTableRespmodStatus_Enabled     SlbNewIcapPolTableRespmodStatus = 1
	SlbNewIcapPolTableRespmodStatus_Disabled    SlbNewIcapPolTableRespmodStatus = 2
	SlbNewIcapPolTableRespmodStatus_Unsupported SlbNewIcapPolTableRespmodStatus = 2147483647
)

type SlbNewIcapPolTableRespmodFallback int32

const (
	SlbNewIcapPolTableRespmodFallback_Drop        SlbNewIcapPolTableRespmodFallback = 1
	SlbNewIcapPolTableRespmodFallback_ContFlow    SlbNewIcapPolTableRespmodFallback = 2
	SlbNewIcapPolTableRespmodFallback_Unsupported SlbNewIcapPolTableRespmodFallback = 2147483647
)

type SlbNewIcapPolTableRespmodXcip int32

const (
	SlbNewIcapPolTableRespmodXcip_Cip         SlbNewIcapPolTableRespmodXcip = 1
	SlbNewIcapPolTableRespmodXcip_Xff         SlbNewIcapPolTableRespmodXcip = 2
	SlbNewIcapPolTableRespmodXcip_Disabled    SlbNewIcapPolTableRespmodXcip = 3
	SlbNewIcapPolTableRespmodXcip_Unsupported SlbNewIcapPolTableRespmodXcip = 2147483647
)

type SlbNewIcapPolTableParams struct {
	// The Icap policy ID(key id) as an index.
	NameIdIndex string `json:"NameIdIndex,omitempty"`
	// Icap policy name.
	Name string `json:"Name,omitempty"`
	// Status (enable/disable) of icap policy.
	AdminStatus SlbNewIcapPolTableAdminStatus `json:"AdminStatus,omitempty"`
	// Delete Icap policy.
	Del SlbNewIcapPolTableDel `json:"Del,omitempty"`
	// Status (enable/disable) of icap policy.
	ReqmodStatus SlbNewIcapPolTableReqmodStatus `json:"ReqmodStatus,omitempty"`
	// Icap Uri string.
	ReqmodUri string `json:"ReqmodUri,omitempty"`
	// The Icap Reqmod  server group number.
	ReqmodGroup string `json:"ReqmodGroup,omitempty"`
	// Fallback action (group down).
	ReqmodFallback SlbNewIcapPolTableReqmodFallback `json:"ReqmodFallback,omitempty"`
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
	ReqmodXcip SlbNewIcapPolTableReqmodXcip `json:"ReqmodXcip,omitempty"`
	// Icap xcip header string.
	ReqmodXcipHeader string `json:"ReqmodXcipHeader,omitempty"`
	// Status (enable/disable) of icap policy.
	RespmodStatus SlbNewIcapPolTableRespmodStatus `json:"RespmodStatus,omitempty"`
	// Icap Uri string.
	RespmodUri string `json:"RespmodUri,omitempty"`
	// The Icap Respmod  server group number.
	RespmodGroup string `json:"RespmodGroup,omitempty"`
	// Fallback action (group down).
	RespmodFallback SlbNewIcapPolTableRespmodFallback `json:"RespmodFallback,omitempty"`
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
	RespmodXcip SlbNewIcapPolTableRespmodXcip `json:"RespmodXcip,omitempty"`
	// Icap xcip header string.
	RespmodXcipHeader string `json:"RespmodXcipHeader,omitempty"`
}

func (p SlbNewIcapPolTableParams) iMABean() {}
