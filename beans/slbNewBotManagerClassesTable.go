package beans

import (
	"fmt"
	"reflect"
)

// SlbNewBotManagerClassesTable Current bot manager table.
// Note:This mib is not supported on VX instance of Virtualization.
type SlbNewBotManagerClassesTable struct {
	// Set bot manager class id.
	SlbNewBotManagerId string
	Params             *SlbNewBotManagerClassesTableParams
}

func NewSlbNewBotManagerClassesTableList() *SlbNewBotManagerClassesTable {
	return &SlbNewBotManagerClassesTable{}
}

func NewSlbNewBotManagerClassesTable(
	slbNewBotManagerId string,
	params *SlbNewBotManagerClassesTableParams,
) *SlbNewBotManagerClassesTable {
	return &SlbNewBotManagerClassesTable{
		SlbNewBotManagerId: slbNewBotManagerId,
		Params:             params,
	}
}

func (c *SlbNewBotManagerClassesTable) Name() string {
	return "SlbNewBotManagerClassesTable"
}

func (c *SlbNewBotManagerClassesTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbNewBotManagerClassesTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbNewBotManagerClassesTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbNewBotManagerId).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbNewBotManagerId)
}

type SlbNewBotManagerClassesTableEnaDis int32

const (
	SlbNewBotManagerClassesTableEnaDis_Enabled     SlbNewBotManagerClassesTableEnaDis = 1
	SlbNewBotManagerClassesTableEnaDis_Disabled    SlbNewBotManagerClassesTableEnaDis = 2
	SlbNewBotManagerClassesTableEnaDis_Unsupported SlbNewBotManagerClassesTableEnaDis = 2147483647
)

type SlbNewBotManagerClassesTableMode int32

const (
	SlbNewBotManagerClassesTableMode_Active      SlbNewBotManagerClassesTableMode = 1
	SlbNewBotManagerClassesTableMode_ReportOnly  SlbNewBotManagerClassesTableMode = 2
	SlbNewBotManagerClassesTableMode_Unsupported SlbNewBotManagerClassesTableMode = 2147483647
)

type SlbNewBotManagerClassesTableBlock int32

const (
	SlbNewBotManagerClassesTableBlock_Enabled     SlbNewBotManagerClassesTableBlock = 1
	SlbNewBotManagerClassesTableBlock_Disabled    SlbNewBotManagerClassesTableBlock = 2
	SlbNewBotManagerClassesTableBlock_Unsupported SlbNewBotManagerClassesTableBlock = 2147483647
)

type SlbNewBotManagerClassesTableCaptcha int32

const (
	SlbNewBotManagerClassesTableCaptcha_Enabled     SlbNewBotManagerClassesTableCaptcha = 1
	SlbNewBotManagerClassesTableCaptcha_Disabled    SlbNewBotManagerClassesTableCaptcha = 2
	SlbNewBotManagerClassesTableCaptcha_Unsupported SlbNewBotManagerClassesTableCaptcha = 2147483647
)

type SlbNewBotManagerClassesTableSessionIdPicker int32

const (
	SlbNewBotManagerClassesTableSessionIdPicker_Header SlbNewBotManagerClassesTableSessionIdPicker = 1
	SlbNewBotManagerClassesTableSessionIdPicker_Cookie SlbNewBotManagerClassesTableSessionIdPicker = 2
	SlbNewBotManagerClassesTableSessionIdPicker_Query  SlbNewBotManagerClassesTableSessionIdPicker = 3
	SlbNewBotManagerClassesTableSessionIdPicker_None   SlbNewBotManagerClassesTableSessionIdPicker = 4
)

type SlbNewBotManagerClassesTableUserIdPicker int32

const (
	SlbNewBotManagerClassesTableUserIdPicker_Header SlbNewBotManagerClassesTableUserIdPicker = 1
	SlbNewBotManagerClassesTableUserIdPicker_Cookie SlbNewBotManagerClassesTableUserIdPicker = 2
	SlbNewBotManagerClassesTableUserIdPicker_Query  SlbNewBotManagerClassesTableUserIdPicker = 3
	SlbNewBotManagerClassesTableUserIdPicker_None   SlbNewBotManagerClassesTableUserIdPicker = 4
)

type SlbNewBotManagerClassesTableSecureCookie int32

const (
	SlbNewBotManagerClassesTableSecureCookie_Enabled     SlbNewBotManagerClassesTableSecureCookie = 1
	SlbNewBotManagerClassesTableSecureCookie_Disabled    SlbNewBotManagerClassesTableSecureCookie = 2
	SlbNewBotManagerClassesTableSecureCookie_Unsupported SlbNewBotManagerClassesTableSecureCookie = 2147483647
)

type SlbNewBotManagerClassesTableAllHeaders int32

const (
	SlbNewBotManagerClassesTableAllHeaders_Enabled     SlbNewBotManagerClassesTableAllHeaders = 1
	SlbNewBotManagerClassesTableAllHeaders_Disabled    SlbNewBotManagerClassesTableAllHeaders = 2
	SlbNewBotManagerClassesTableAllHeaders_Unsupported SlbNewBotManagerClassesTableAllHeaders = 2147483647
)

type SlbNewBotManagerClassesTableDel int32

const (
	SlbNewBotManagerClassesTableDel_Other       SlbNewBotManagerClassesTableDel = 1
	SlbNewBotManagerClassesTableDel_Delete      SlbNewBotManagerClassesTableDel = 2
	SlbNewBotManagerClassesTableDel_Unsupported SlbNewBotManagerClassesTableDel = 2147483647
)

type SlbNewBotManagerClassesTableAppType int32

const (
	SlbNewBotManagerClassesTableAppType_Web         SlbNewBotManagerClassesTableAppType = 1
	SlbNewBotManagerClassesTableAppType_Mobile      SlbNewBotManagerClassesTableAppType = 2
	SlbNewBotManagerClassesTableAppType_Unsupported SlbNewBotManagerClassesTableAppType = 2147483647
)

type SlbNewBotManagerClassesTableJsInjectStatus int32

const (
	SlbNewBotManagerClassesTableJsInjectStatus_Enabled     SlbNewBotManagerClassesTableJsInjectStatus = 1
	SlbNewBotManagerClassesTableJsInjectStatus_Disabled    SlbNewBotManagerClassesTableJsInjectStatus = 2
	SlbNewBotManagerClassesTableJsInjectStatus_Unsupported SlbNewBotManagerClassesTableJsInjectStatus = 2147483647
)

type SlbNewBotManagerClassesTableCustomWebBlockStatus int32

const (
	SlbNewBotManagerClassesTableCustomWebBlockStatus_Enabled     SlbNewBotManagerClassesTableCustomWebBlockStatus = 1
	SlbNewBotManagerClassesTableCustomWebBlockStatus_Disabled    SlbNewBotManagerClassesTableCustomWebBlockStatus = 2
	SlbNewBotManagerClassesTableCustomWebBlockStatus_Unsupported SlbNewBotManagerClassesTableCustomWebBlockStatus = 2147483647
)

type SlbNewBotManagerClassesTableCustomMobBlockStatus int32

const (
	SlbNewBotManagerClassesTableCustomMobBlockStatus_Enabled     SlbNewBotManagerClassesTableCustomMobBlockStatus = 1
	SlbNewBotManagerClassesTableCustomMobBlockStatus_Disabled    SlbNewBotManagerClassesTableCustomMobBlockStatus = 2
	SlbNewBotManagerClassesTableCustomMobBlockStatus_Unsupported SlbNewBotManagerClassesTableCustomMobBlockStatus = 2147483647
)

type SlbNewBotManagerClassesTableCustomMobCaptchaStatus int32

const (
	SlbNewBotManagerClassesTableCustomMobCaptchaStatus_Enabled     SlbNewBotManagerClassesTableCustomMobCaptchaStatus = 1
	SlbNewBotManagerClassesTableCustomMobCaptchaStatus_Disabled    SlbNewBotManagerClassesTableCustomMobCaptchaStatus = 2
	SlbNewBotManagerClassesTableCustomMobCaptchaStatus_Unsupported SlbNewBotManagerClassesTableCustomMobCaptchaStatus = 2147483647
)

type SlbNewBotManagerClassesTableTrkevent int32

const (
	SlbNewBotManagerClassesTableTrkevent_Enabled     SlbNewBotManagerClassesTableTrkevent = 1
	SlbNewBotManagerClassesTableTrkevent_Disabled    SlbNewBotManagerClassesTableTrkevent = 2
	SlbNewBotManagerClassesTableTrkevent_Unsupported SlbNewBotManagerClassesTableTrkevent = 2147483647
)

type SlbNewBotManagerClassesTableSameSite int32

const (
	SlbNewBotManagerClassesTableSameSite_None        SlbNewBotManagerClassesTableSameSite = 0
	SlbNewBotManagerClassesTableSameSite_Lax         SlbNewBotManagerClassesTableSameSite = 1
	SlbNewBotManagerClassesTableSameSite_Strict      SlbNewBotManagerClassesTableSameSite = 2
	SlbNewBotManagerClassesTableSameSite_Unsupported SlbNewBotManagerClassesTableSameSite = 2147483647
)

type SlbNewBotManagerClassesTableAppClassAppType int32

const (
	SlbNewBotManagerClassesTableAppClassAppType_None   SlbNewBotManagerClassesTableAppClassAppType = 1
	SlbNewBotManagerClassesTableAppClassAppType_Web    SlbNewBotManagerClassesTableAppClassAppType = 2
	SlbNewBotManagerClassesTableAppClassAppType_Mobile SlbNewBotManagerClassesTableAppClassAppType = 3
)

type SlbNewBotManagerClassesTableParams struct {
	// Set bot manager class id.
	Id string `json:"Id,omitempty"`
	// Set bot manager class name.
	Name string `json:"Name,omitempty"`
	// Set bot manager policy status either ena or dis.
	EnaDis SlbNewBotManagerClassesTableEnaDis `json:"EnaDis,omitempty"`
	// Set bot manager policy subscriber id in UUID format.
	SubscriberId string `json:"SubscriberId,omitempty"`
	// Set bot manager sideband policy name.
	SidebandId string `json:"SidebandId,omitempty"`
	// Set bot manager policy protection mode active or report-only.
	Mode SlbNewBotManagerClassesTableMode `json:"Mode,omitempty"`
	// Set bot manager block page url.
	BlockUrl string `json:"BlockUrl,omitempty"`
	// Set bot manager captcha page url.
	CaptchaUrl string `json:"CaptchaUrl,omitempty"`
	// Set bot manager email for block\captcha redirect pages.
	Email string `json:"Email,omitempty"`
	// Set if bot manager should perform block action.
	Block SlbNewBotManagerClassesTableBlock `json:"Block,omitempty"`
	// Set if bot manager should perform captcha action.
	Captcha SlbNewBotManagerClassesTableCaptcha `json:"Captcha,omitempty"`
	// Set Where to search the session id in.
	SessionIdPicker SlbNewBotManagerClassesTableSessionIdPicker `json:"SessionIdPicker,omitempty"`
	// Set bot manager session id.
	SessionId string `json:"SessionId,omitempty"`
	// Set Where to search the user id in.
	UserIdPicker SlbNewBotManagerClassesTableUserIdPicker `json:"UserIdPicker,omitempty"`
	// Set bot manager user id.
	UserId string `json:"UserId,omitempty"`
	// Set user defined ip address header
	UserDefinedIpAddress string `json:"UserDefinedIpAddress,omitempty"`
	// Set bot manager ip address header.
	IpAddress string `json:"IpAddress,omitempty"`
	// Set bot manager domain for cookie grouping.
	DomainGrouping string `json:"DomainGrouping,omitempty"`
	// Set list of extensions to bypass bot manager.
	FileBypass string `json:"FileBypass,omitempty"`
	// Set if bot manger cookies should be sent over secure connection.
	SecureCookie SlbNewBotManagerClassesTableSecureCookie `json:"SecureCookie,omitempty"`
	// Set if all headers should be sent for bot manager for advanced bot detection ena or dis.
	AllHeaders SlbNewBotManagerClassesTableAllHeaders `json:"AllHeaders,omitempty"`
	// By setting the value to delete(2), the entire row is deleted.
	Del SlbNewBotManagerClassesTableDel `json:"Del,omitempty"`
	// Set bot manager application type.
	AppType SlbNewBotManagerClassesTableAppType `json:"AppType,omitempty"`
	// Enable or disable jsinject
	JsInjectStatus SlbNewBotManagerClassesTableJsInjectStatus `json:"JsInjectStatus,omitempty"`
	// Enable or disable custom web block.
	CustomWebBlockStatus SlbNewBotManagerClassesTableCustomWebBlockStatus `json:"CustomWebBlockStatus,omitempty"`
	// Set the custom web block response code.
	CustomWebBlockRespCode uint64 `json:"CustomWebBlockRespCode,omitempty"`
	// Set the custom web block response body.
	CustomWebBlockRespBody string `json:"CustomWebBlockRespBody,omitempty"`
	// Set the custom web block first header name.
	CustomWebBlockFirstHeaderName string `json:"CustomWebBlockFirstHeaderName,omitempty"`
	// Set the custom web block first header value.
	CustomWebBlockFirstHeaderValue string `json:"CustomWebBlockFirstHeaderValue,omitempty"`
	// Set the custom web block second header name.
	CustomWebBlockSecondHeaderName string `json:"CustomWebBlockSecondHeaderName,omitempty"`
	// Set the custom web block second header value.
	CustomWebBlockSecondHeaderValue string `json:"CustomWebBlockSecondHeaderValue,omitempty"`
	// Enable or disable custom mobile block.
	CustomMobBlockStatus SlbNewBotManagerClassesTableCustomMobBlockStatus `json:"CustomMobBlockStatus,omitempty"`
	// Set the custom mobile block response code.
	CustomMobBlockRespCode uint64 `json:"CustomMobBlockRespCode,omitempty"`
	// Set the custom mobile block response body.
	CustomMobBlockRespBody string `json:"CustomMobBlockRespBody,omitempty"`
	// Set the custom mobile block first header name.
	CustomMobBlockFirstHeaderName string `json:"CustomMobBlockFirstHeaderName,omitempty"`
	// Set the custom mobile block first header value.
	CustomMobBlockFirstHeaderValue string `json:"CustomMobBlockFirstHeaderValue,omitempty"`
	// Set the custom mobile block second header name.
	CustomMobBlockSecondHeaderName string `json:"CustomMobBlockSecondHeaderName,omitempty"`
	// Set the custom mobile block second header value.
	CustomMobBlockSecondHeaderValue string `json:"CustomMobBlockSecondHeaderValue,omitempty"`
	// Enable or disable custom mobile Captcha.
	CustomMobCaptchaStatus SlbNewBotManagerClassesTableCustomMobCaptchaStatus `json:"CustomMobCaptchaStatus,omitempty"`
	// Set the custom mobile Captcha response code.
	CustomMobCaptchaRespCode uint64 `json:"CustomMobCaptchaRespCode,omitempty"`
	// Set the custom mobile Captcha response body.
	CustomMobCaptchaRespBody string `json:"CustomMobCaptchaRespBody,omitempty"`
	// Set the custom mobile Captcha first header name.
	CustomMobCaptchaFirstHeaderName string `json:"CustomMobCaptchaFirstHeaderName,omitempty"`
	// Set the custom mobile Captcha first header value.
	CustomMobCaptchaFirstHeaderValue string `json:"CustomMobCaptchaFirstHeaderValue,omitempty"`
	// Set the custom mobile Captcha second header name.
	CustomMobCaptchaSecondHeaderName string `json:"CustomMobCaptchaSecondHeaderName,omitempty"`
	// Set the custom mobile Captcha second header value.
	CustomMobCaptchaSecondHeaderValue string `json:"CustomMobCaptchaSecondHeaderValue,omitempty"`
	// Set Shield Square trkevent header as enabled or disabled
	Trkevent SlbNewBotManagerClassesTableTrkevent `json:"Trkevent,omitempty"`
	// Set the SameSite header header field.
	SameSite SlbNewBotManagerClassesTableSameSite `json:"SameSite,omitempty"`
	// Set which application to use if classification succeeded.
	AppClassAppType SlbNewBotManagerClassesTableAppClassAppType `json:"AppClassAppType,omitempty"`
	// Set content class ID containing only header/s lookup with a key of user-agent.
	AppClassUserAgent string `json:"AppClassUserAgent,omitempty"`
	// Set content class ID containing only hostname/s or path/s (or both) lookups.
	AppClassUrl string `json:"AppClassUrl,omitempty"`
	// Set content class ID containing only header/s lookup.
	AppClassHeader string `json:"AppClassHeader,omitempty"`
	// Set content class ID containing only cookie/s lookup.
	AppClassCookie string `json:"AppClassCookie,omitempty"`
	// Set headers filter when allhdrs is enabled.
	AllHeadersFilter string `json:"AllHeadersFilter,omitempty"`
	// Set a list of header names to include/exclude when allhdrs is enabled.
	AllHeadersList string `json:"AllHeadersList,omitempty"`
}

func (p SlbNewBotManagerClassesTableParams) iMABean() {}
