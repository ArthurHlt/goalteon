package beans

import (
	"fmt"
	"reflect"
)

// SlbCurSecurePathClassesTable Current secure path table.
// Note:This mib is not supported on VX instance of Virtualization.
type SlbCurSecurePathClassesTable struct {
	// Set secure path class id.
	SlbCurSecurePathId string
	Params             *SlbCurSecurePathClassesTableParams
}

func NewSlbCurSecurePathClassesTableList() *SlbCurSecurePathClassesTable {
	return &SlbCurSecurePathClassesTable{}
}

func NewSlbCurSecurePathClassesTable(
	slbCurSecurePathId string,
	params *SlbCurSecurePathClassesTableParams,
) *SlbCurSecurePathClassesTable {
	return &SlbCurSecurePathClassesTable{
		SlbCurSecurePathId: slbCurSecurePathId,
		Params:             params,
	}
}

func (c *SlbCurSecurePathClassesTable) Name() string {
	return "SlbCurSecurePathClassesTable"
}

func (c *SlbCurSecurePathClassesTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbCurSecurePathClassesTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbCurSecurePathClassesTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbCurSecurePathId).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbCurSecurePathId)
}

type SlbCurSecurePathClassesTableEnaDis int32

const (
	SlbCurSecurePathClassesTableEnaDis_Enabled     SlbCurSecurePathClassesTableEnaDis = 1
	SlbCurSecurePathClassesTableEnaDis_Disabled    SlbCurSecurePathClassesTableEnaDis = 2
	SlbCurSecurePathClassesTableEnaDis_Unsupported SlbCurSecurePathClassesTableEnaDis = 2147483647
)

type SlbCurSecurePathClassesTableBotMng int32

const (
	SlbCurSecurePathClassesTableBotMng_Enabled     SlbCurSecurePathClassesTableBotMng = 1
	SlbCurSecurePathClassesTableBotMng_Disabled    SlbCurSecurePathClassesTableBotMng = 2
	SlbCurSecurePathClassesTableBotMng_Unsupported SlbCurSecurePathClassesTableBotMng = 2147483647
)

type SlbCurSecurePathClassesTableQueryBypass int32

const (
	SlbCurSecurePathClassesTableQueryBypass_Enabled     SlbCurSecurePathClassesTableQueryBypass = 1
	SlbCurSecurePathClassesTableQueryBypass_Disabled    SlbCurSecurePathClassesTableQueryBypass = 2
	SlbCurSecurePathClassesTableQueryBypass_Unsupported SlbCurSecurePathClassesTableQueryBypass = 2147483647
)

type SlbCurSecurePathClassesTableParams struct {
	// Set secure path class id.
	Id string `json:"Id,omitempty"`
	// Set secure path class name.
	Name string `json:"Name,omitempty"`
	// Set secure path policy status either ena or dis.
	EnaDis SlbCurSecurePathClassesTableEnaDis `json:"EnaDis,omitempty"`
	// Set secure path policy secure path capabilities ena or dis.
	BotMng SlbCurSecurePathClassesTableBotMng `json:"BotMng,omitempty"`
	// Set secure path policy token in UUID format.
	Token string `json:"Token,omitempty"`
	// Set secure path policy app id in UUID format.
	AppId string `json:"AppId,omitempty"`
	// Set list of extensions to bypass secure path.
	FileBypass string `json:"FileBypass,omitempty"`
	// Set list of methods to bypass secure path.
	MethodBypass string `json:"MethodBypass,omitempty"`
	// Set secure path policy query bypass ena or dis.
	QueryBypass SlbCurSecurePathClassesTableQueryBypass `json:"QueryBypass,omitempty"`
	// Set the custom web block response code.
	MaxRequestSize uint32 `json:"MaxRequestSize,omitempty"`
}

func (p SlbCurSecurePathClassesTableParams) iMABean() {}
