package beans

import (
	"fmt"
	"reflect"
)

// SlbCurBotManagerClassesTable Current bot manager table.
// Note:This mib is not supported on VX instance of Virtualization.
type SlbCurBotManagerClassesTable struct {
	// Set bot manager class id.
	SlbCurBotManagerId string
	Params             *SlbCurBotManagerClassesTableParams
}

func NewSlbCurBotManagerClassesTableList() *SlbCurBotManagerClassesTable {
	return &SlbCurBotManagerClassesTable{}
}

func NewSlbCurBotManagerClassesTable(
	slbCurBotManagerId string,
	params *SlbCurBotManagerClassesTableParams,
) *SlbCurBotManagerClassesTable {
	return &SlbCurBotManagerClassesTable{
		SlbCurBotManagerId: slbCurBotManagerId,
		Params:             params,
	}
}

func (c *SlbCurBotManagerClassesTable) Name() string {
	return "SlbCurBotManagerClassesTable"
}

func (c *SlbCurBotManagerClassesTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbCurBotManagerClassesTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbCurBotManagerClassesTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbCurBotManagerId).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbCurBotManagerId)
}

type SlbCurBotManagerClassesTableEnaDis int32

const (
	SlbCurBotManagerClassesTableEnaDis_Enabled     SlbCurBotManagerClassesTableEnaDis = 1
	SlbCurBotManagerClassesTableEnaDis_Disabled    SlbCurBotManagerClassesTableEnaDis = 2
	SlbCurBotManagerClassesTableEnaDis_Unsupported SlbCurBotManagerClassesTableEnaDis = 2147483647
)

type SlbCurBotManagerClassesTableMode int32

const (
	SlbCurBotManagerClassesTableMode_Active      SlbCurBotManagerClassesTableMode = 1
	SlbCurBotManagerClassesTableMode_ReportOnly  SlbCurBotManagerClassesTableMode = 2
	SlbCurBotManagerClassesTableMode_Unsupported SlbCurBotManagerClassesTableMode = 2147483647
)

type SlbCurBotManagerClassesTableBlock int32

const (
	SlbCurBotManagerClassesTableBlock_Enabled     SlbCurBotManagerClassesTableBlock = 1
	SlbCurBotManagerClassesTableBlock_Disabled    SlbCurBotManagerClassesTableBlock = 2
	SlbCurBotManagerClassesTableBlock_Unsupported SlbCurBotManagerClassesTableBlock = 2147483647
)

type SlbCurBotManagerClassesTableCaptcha int32

const (
	SlbCurBotManagerClassesTableCaptcha_Enabled     SlbCurBotManagerClassesTableCaptcha = 1
	SlbCurBotManagerClassesTableCaptcha_Disabled    SlbCurBotManagerClassesTableCaptcha = 2
	SlbCurBotManagerClassesTableCaptcha_Unsupported SlbCurBotManagerClassesTableCaptcha = 2147483647
)

type SlbCurBotManagerClassesTableSessionIdPicker int32

const (
	SlbCurBotManagerClassesTableSessionIdPicker_Header SlbCurBotManagerClassesTableSessionIdPicker = 1
	SlbCurBotManagerClassesTableSessionIdPicker_Cookie SlbCurBotManagerClassesTableSessionIdPicker = 2
	SlbCurBotManagerClassesTableSessionIdPicker_Query  SlbCurBotManagerClassesTableSessionIdPicker = 3
	SlbCurBotManagerClassesTableSessionIdPicker_None   SlbCurBotManagerClassesTableSessionIdPicker = 4
)

type SlbCurBotManagerClassesTableUserIdPicker int32

const (
	SlbCurBotManagerClassesTableUserIdPicker_Header SlbCurBotManagerClassesTableUserIdPicker = 1
	SlbCurBotManagerClassesTableUserIdPicker_Cookie SlbCurBotManagerClassesTableUserIdPicker = 2
	SlbCurBotManagerClassesTableUserIdPicker_Query  SlbCurBotManagerClassesTableUserIdPicker = 3
	SlbCurBotManagerClassesTableUserIdPicker_None   SlbCurBotManagerClassesTableUserIdPicker = 4
)

type SlbCurBotManagerClassesTableSecureCookie int32

const (
	SlbCurBotManagerClassesTableSecureCookie_Enabled     SlbCurBotManagerClassesTableSecureCookie = 1
	SlbCurBotManagerClassesTableSecureCookie_Disabled    SlbCurBotManagerClassesTableSecureCookie = 2
	SlbCurBotManagerClassesTableSecureCookie_Unsupported SlbCurBotManagerClassesTableSecureCookie = 2147483647
)

type SlbCurBotManagerClassesTableAllHeaders int32

const (
	SlbCurBotManagerClassesTableAllHeaders_Enabled     SlbCurBotManagerClassesTableAllHeaders = 1
	SlbCurBotManagerClassesTableAllHeaders_Disabled    SlbCurBotManagerClassesTableAllHeaders = 2
	SlbCurBotManagerClassesTableAllHeaders_Unsupported SlbCurBotManagerClassesTableAllHeaders = 2147483647
)

type SlbCurBotManagerClassesTableAppType int32

const (
	SlbCurBotManagerClassesTableAppType_Web         SlbCurBotManagerClassesTableAppType = 1
	SlbCurBotManagerClassesTableAppType_Mobile      SlbCurBotManagerClassesTableAppType = 2
	SlbCurBotManagerClassesTableAppType_Unsupported SlbCurBotManagerClassesTableAppType = 2147483647
)

type SlbCurBotManagerClassesTableJsInjectStatus int32

const (
	SlbCurBotManagerClassesTableJsInjectStatus_Enabled     SlbCurBotManagerClassesTableJsInjectStatus = 1
	SlbCurBotManagerClassesTableJsInjectStatus_Disabled    SlbCurBotManagerClassesTableJsInjectStatus = 2
	SlbCurBotManagerClassesTableJsInjectStatus_Unsupported SlbCurBotManagerClassesTableJsInjectStatus = 2147483647
)

type SlbCurBotManagerClassesTableCustomWebBlockStatus int32

const (
	SlbCurBotManagerClassesTableCustomWebBlockStatus_Enabled     SlbCurBotManagerClassesTableCustomWebBlockStatus = 1
	SlbCurBotManagerClassesTableCustomWebBlockStatus_Disabled    SlbCurBotManagerClassesTableCustomWebBlockStatus = 2
	SlbCurBotManagerClassesTableCustomWebBlockStatus_Unsupported SlbCurBotManagerClassesTableCustomWebBlockStatus = 2147483647
)

type SlbCurBotManagerClassesTableCustomMobBlockStatus int32

const (
	SlbCurBotManagerClassesTableCustomMobBlockStatus_Enabled     SlbCurBotManagerClassesTableCustomMobBlockStatus = 1
	SlbCurBotManagerClassesTableCustomMobBlockStatus_Disabled    SlbCurBotManagerClassesTableCustomMobBlockStatus = 2
	SlbCurBotManagerClassesTableCustomMobBlockStatus_Unsupported SlbCurBotManagerClassesTableCustomMobBlockStatus = 2147483647
)

type SlbCurBotManagerClassesTableCustomMobCaptchaStatus int32

const (
	SlbCurBotManagerClassesTableCustomMobCaptchaStatus_Enabled     SlbCurBotManagerClassesTableCustomMobCaptchaStatus = 1
	SlbCurBotManagerClassesTableCustomMobCaptchaStatus_Disabled    SlbCurBotManagerClassesTableCustomMobCaptchaStatus = 2
	SlbCurBotManagerClassesTableCustomMobCaptchaStatus_Unsupported SlbCurBotManagerClassesTableCustomMobCaptchaStatus = 2147483647
)

type SlbCurBotManagerClassesTableTrkevent int32

const (
	SlbCurBotManagerClassesTableTrkevent_Enabled     SlbCurBotManagerClassesTableTrkevent = 1
	SlbCurBotManagerClassesTableTrkevent_Disabled    SlbCurBotManagerClassesTableTrkevent = 2
	SlbCurBotManagerClassesTableTrkevent_Unsupported SlbCurBotManagerClassesTableTrkevent = 2147483647
)

type SlbCurBotManagerClassesTableSameSite int32

const (
	SlbCurBotManagerClassesTableSameSite_None        SlbCurBotManagerClassesTableSameSite = 0
	SlbCurBotManagerClassesTableSameSite_Lax         SlbCurBotManagerClassesTableSameSite = 1
	SlbCurBotManagerClassesTableSameSite_Strict      SlbCurBotManagerClassesTableSameSite = 2
	SlbCurBotManagerClassesTableSameSite_Unsupported SlbCurBotManagerClassesTableSameSite = 2147483647
)

type SlbCurBotManagerClassesTableAppClassAppType int32

const (
	SlbCurBotManagerClassesTableAppClassAppType_None   SlbCurBotManagerClassesTableAppClassAppType = 1
	SlbCurBotManagerClassesTableAppClassAppType_Web    SlbCurBotManagerClassesTableAppClassAppType = 2
	SlbCurBotManagerClassesTableAppClassAppType_Mobile SlbCurBotManagerClassesTableAppClassAppType = 3
)

type SlbCurBotManagerClassesTableParams struct {
	// Set bot manager class id.
	Id string `json:"Id,omitempty"`
	// Set bot manager class name.
	Name string `json:"Name,omitempty"`
	// Set bot manager policy status either ena or dis.
	EnaDis SlbCurBotManagerClassesTableEnaDis `json:"EnaDis,omitempty"`
	// Set bot manager policy subscriber id in UUID format.
	SubscriberId string `json:"SubscriberId,omitempty"`
	// Set bot manager sideband policy name.
	SidebandId string `json:"SidebandId,omitempty"`
	// Set bot manager policy protection mode active or report-only.
	Mode SlbCurBotManagerClassesTableMode `json:"Mode,omitempty"`
	// Set bot manager block page url.
	BlockUrl string `json:"BlockUrl,omitempty"`
	// Set bot manager captcha page url.
	CaptchaUrl string `json:"CaptchaUrl,omitempty"`
	// Set bot manager email for block\captcha redirect pages.
	Email string `json:"Email,omitempty"`
	// Set if bot manager should perform block action.
	Block SlbCurBotManagerClassesTableBlock `json:"Block,omitempty"`
	// Set if bot manager should perform captcha action.
	Captcha SlbCurBotManagerClassesTableCaptcha `json:"Captcha,omitempty"`
	// Set Where to search the session id in.
	SessionIdPicker SlbCurBotManagerClassesTableSessionIdPicker `json:"SessionIdPicker,omitempty"`
	// Set bot manager session id.
	SessionId string `json:"SessionId,omitempty"`
	// Set Where to search the user id in.
	UserIdPicker SlbCurBotManagerClassesTableUserIdPicker `json:"UserIdPicker,omitempty"`
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
	SecureCookie SlbCurBotManagerClassesTableSecureCookie `json:"SecureCookie,omitempty"`
	// Set if all headers should be sent for bot manager for advanced bot detection ena or dis.
	AllHeaders SlbCurBotManagerClassesTableAllHeaders `json:"AllHeaders,omitempty"`
	// Set bot manager application type.
	AppType SlbCurBotManagerClassesTableAppType `json:"AppType,omitempty"`
	// Enable or disable jsinject
	JsInjectStatus SlbCurBotManagerClassesTableJsInjectStatus `json:"JsInjectStatus,omitempty"`
	// Enable or disable custom web block.
	CustomWebBlockStatus SlbCurBotManagerClassesTableCustomWebBlockStatus `json:"CustomWebBlockStatus,omitempty"`
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
	CustomMobBlockStatus SlbCurBotManagerClassesTableCustomMobBlockStatus `json:"CustomMobBlockStatus,omitempty"`
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
	CustomMobCaptchaStatus SlbCurBotManagerClassesTableCustomMobCaptchaStatus `json:"CustomMobCaptchaStatus,omitempty"`
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
	// Set the Shield Square trkevent header as enabled or disabled.
	Trkevent SlbCurBotManagerClassesTableTrkevent `json:"Trkevent,omitempty"`
	// Set the SameSite header header field.
	SameSite SlbCurBotManagerClassesTableSameSite `json:"SameSite,omitempty"`
	// Set which application to use if classification succeeded.
	AppClassAppType SlbCurBotManagerClassesTableAppClassAppType `json:"AppClassAppType,omitempty"`
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

func (p SlbCurBotManagerClassesTableParams) iMABean() {}
