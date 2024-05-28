package beans

import (
	"fmt"
	"reflect"
)

// SlbOperRealServerTable The table of real servers.
// Note:This mib is not supported for VX instance of Virtualization.
type SlbOperRealServerTable struct {
	// The index for the  real server.
	SlbOperRealServerIndex int32
	Params                 *SlbOperRealServerTableParams
}

func NewSlbOperRealServerTableList() *SlbOperRealServerTable {
	return &SlbOperRealServerTable{}
}

func NewSlbOperRealServerTable(
	slbOperRealServerIndex int32,
	params *SlbOperRealServerTableParams,
) *SlbOperRealServerTable {
	return &SlbOperRealServerTable{
		SlbOperRealServerIndex: slbOperRealServerIndex,
		Params:                 params,
	}
}

func (c *SlbOperRealServerTable) Name() string {
	return "SlbOperRealServerTable"
}

func (c *SlbOperRealServerTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbOperRealServerTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbOperRealServerTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbOperRealServerIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbOperRealServerIndex)
}

type SlbOperRealServerTableStatus int32

const (
	SlbOperRealServerTableStatus_Enable                  SlbOperRealServerTableStatus = 1
	SlbOperRealServerTableStatus_Disable                 SlbOperRealServerTableStatus = 2
	SlbOperRealServerTableStatus_Cookiepersistent        SlbOperRealServerTableStatus = 3
	SlbOperRealServerTableStatus_Fastage                 SlbOperRealServerTableStatus = 4
	SlbOperRealServerTableStatus_Cookiepersistentfastage SlbOperRealServerTableStatus = 5
	SlbOperRealServerTableStatus_Unsupported             SlbOperRealServerTableStatus = 2147483647
)

type SlbOperRealServerTableParams struct {
	// The index for the  real server.
	Index int32 `json:"Index,omitempty"`
	// This an action object which is used to temporarily enable/disable a
	// real server. The real server will be returned to its configured
	// operational mode when the switch is reset. Setting the value to
	// 'cookiepersistant' allows cookie persistent HTTP 1.0 sessions
	// when the real server is offline.'fastage' option is for removing the
	// existing sessions. And 'cookiepersistentfastage' allows cookie
	// persistent sessions and remove the existing sessions when real server is offline.
	Status SlbOperRealServerTableStatus `json:"Status,omitempty"`
}

func (p SlbOperRealServerTableParams) iMABean() {}
