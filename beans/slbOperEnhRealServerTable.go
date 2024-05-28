package beans

import (
	"fmt"
	"reflect"
)

// SlbOperEnhRealServerTable The table of real servers.
// Note:This mib is not supported for VX instance of Virtualization.
type SlbOperEnhRealServerTable struct {
	// The index for the  real server.
	SlbOperEnhRealServerIndex string
	Params                    *SlbOperEnhRealServerTableParams
}

func NewSlbOperEnhRealServerTableList() *SlbOperEnhRealServerTable {
	return &SlbOperEnhRealServerTable{}
}

func NewSlbOperEnhRealServerTable(
	slbOperEnhRealServerIndex string,
	params *SlbOperEnhRealServerTableParams,
) *SlbOperEnhRealServerTable {
	return &SlbOperEnhRealServerTable{
		SlbOperEnhRealServerIndex: slbOperEnhRealServerIndex,
		Params:                    params,
	}
}

func (c *SlbOperEnhRealServerTable) Name() string {
	return "SlbOperEnhRealServerTable"
}

func (c *SlbOperEnhRealServerTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbOperEnhRealServerTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbOperEnhRealServerTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbOperEnhRealServerIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbOperEnhRealServerIndex)
}

type SlbOperEnhRealServerTableStatus int32

const (
	SlbOperEnhRealServerTableStatus_Enable                  SlbOperEnhRealServerTableStatus = 1
	SlbOperEnhRealServerTableStatus_Disable                 SlbOperEnhRealServerTableStatus = 2
	SlbOperEnhRealServerTableStatus_Cookiepersistent        SlbOperEnhRealServerTableStatus = 3
	SlbOperEnhRealServerTableStatus_Fastage                 SlbOperEnhRealServerTableStatus = 4
	SlbOperEnhRealServerTableStatus_Cookiepersistentfastage SlbOperEnhRealServerTableStatus = 5
	SlbOperEnhRealServerTableStatus_Unsupported             SlbOperEnhRealServerTableStatus = 2147483647
)

type SlbOperEnhRealServerTableParams struct {
	// The index for the  real server.
	Index string `json:"Index,omitempty"`
	// This an action object which is used to temporarily enable/disable a
	// real server. The real server will be returned to its configured
	// operational mode when the switch is reset. Setting the value to
	// 'cookiepersistant' allows cookie persistent HTTP 1.0 sessions
	// when the real server is offline.'fastage' option is for removing the
	// existing sessions. And 'cookiepersistentfastage' allows cookie
	// persistent sessions and remove the existing sessions when real server is offline.
	Status SlbOperEnhRealServerTableStatus `json:"Status,omitempty"`
}

func (p SlbOperEnhRealServerTableParams) iMABean() {}
