package beans

import (
	"fmt"
	"reflect"
)

// SlbNewSecurePathClassesTable Current secure path table.
// Note:This mib is not supported on VX instance of Virtualization.
type SlbNewSecurePathClassesTable struct {
	// Set secure path class id.
	SlbNewSecurePathId string
	Params             *SlbNewSecurePathClassesTableParams
}

func NewSlbNewSecurePathClassesTableList() *SlbNewSecurePathClassesTable {
	return &SlbNewSecurePathClassesTable{}
}

func NewSlbNewSecurePathClassesTable(
	slbNewSecurePathId string,
	params *SlbNewSecurePathClassesTableParams,
) *SlbNewSecurePathClassesTable {
	return &SlbNewSecurePathClassesTable{
		SlbNewSecurePathId: slbNewSecurePathId,
		Params:             params,
	}
}

func (c *SlbNewSecurePathClassesTable) Name() string {
	return "SlbNewSecurePathClassesTable"
}

func (c *SlbNewSecurePathClassesTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbNewSecurePathClassesTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbNewSecurePathClassesTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbNewSecurePathId).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbNewSecurePathId)
}

type SlbNewSecurePathClassesTableEnaDis int32

const (
	SlbNewSecurePathClassesTableEnaDis_Enabled     SlbNewSecurePathClassesTableEnaDis = 1
	SlbNewSecurePathClassesTableEnaDis_Disabled    SlbNewSecurePathClassesTableEnaDis = 2
	SlbNewSecurePathClassesTableEnaDis_Unsupported SlbNewSecurePathClassesTableEnaDis = 2147483647
)

type SlbNewSecurePathClassesTableBotMng int32

const (
	SlbNewSecurePathClassesTableBotMng_Enabled     SlbNewSecurePathClassesTableBotMng = 1
	SlbNewSecurePathClassesTableBotMng_Disabled    SlbNewSecurePathClassesTableBotMng = 2
	SlbNewSecurePathClassesTableBotMng_Unsupported SlbNewSecurePathClassesTableBotMng = 2147483647
)

type SlbNewSecurePathClassesTableQueryBypass int32

const (
	SlbNewSecurePathClassesTableQueryBypass_Enabled     SlbNewSecurePathClassesTableQueryBypass = 1
	SlbNewSecurePathClassesTableQueryBypass_Disabled    SlbNewSecurePathClassesTableQueryBypass = 2
	SlbNewSecurePathClassesTableQueryBypass_Unsupported SlbNewSecurePathClassesTableQueryBypass = 2147483647
)

type SlbNewSecurePathClassesTableDel int32

const (
	SlbNewSecurePathClassesTableDel_Other       SlbNewSecurePathClassesTableDel = 1
	SlbNewSecurePathClassesTableDel_Delete      SlbNewSecurePathClassesTableDel = 2
	SlbNewSecurePathClassesTableDel_Unsupported SlbNewSecurePathClassesTableDel = 2147483647
)

type SlbNewSecurePathClassesTableParams struct {
	// Set secure path class id.
	Id string `json:"Id,omitempty"`
	// Set secure path class name.
	Name string `json:"Name,omitempty"`
	// Set secure path policy status either ena or dis.
	EnaDis SlbNewSecurePathClassesTableEnaDis `json:"EnaDis,omitempty"`
	// Set secure path policy secure path capabilities ena or dis.
	BotMng SlbNewSecurePathClassesTableBotMng `json:"BotMng,omitempty"`
	// Set secure path policy token in UUID format.
	Token string `json:"Token,omitempty"`
	// Set secure path policy app id in UUID format.
	AppId string `json:"AppId,omitempty"`
	// Set list of extensions to bypass secure path.
	FileBypass string `json:"FileBypass,omitempty"`
	// Set list of methods to bypass secure path.
	MethodBypass string `json:"MethodBypass,omitempty"`
	// Set secure path policy query bypass ena or dis.
	QueryBypass SlbNewSecurePathClassesTableQueryBypass `json:"QueryBypass,omitempty"`
	// Set the custom web block response code.
	MaxRequestSize uint32 `json:"MaxRequestSize,omitempty"`
	// By setting the value to delete(2), the entire row is deleted.
	Del SlbNewSecurePathClassesTableDel `json:"Del,omitempty"`
}

func (p SlbNewSecurePathClassesTableParams) iMABean() {}
