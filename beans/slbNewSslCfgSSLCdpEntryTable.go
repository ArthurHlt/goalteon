package beans

import (
	"fmt"
	"reflect"
)

// SlbNewSslCfgSSLCdpEntryTable The table for configuring CDP Entry.
// Note:This mib is not supported for VX instance of Virtualization.
type SlbNewSslCfgSSLCdpEntryTable struct {
	// The CDP Group (key id) as an index.
	SlbNewSslCfgSSLCdpGrpIdIndex string
	// The Cdp Entry number as an index.
	SlbNewSslCfgSSLCdpEntryIndex int32
	Params                       *SlbNewSslCfgSSLCdpEntryTableParams
}

func NewSlbNewSslCfgSSLCdpEntryTableList() *SlbNewSslCfgSSLCdpEntryTable {
	return &SlbNewSslCfgSSLCdpEntryTable{}
}

func NewSlbNewSslCfgSSLCdpEntryTable(
	slbNewSslCfgSSLCdpGrpIdIndex string,
	slbNewSslCfgSSLCdpEntryIndex int32,
	params *SlbNewSslCfgSSLCdpEntryTableParams,
) *SlbNewSslCfgSSLCdpEntryTable {
	return &SlbNewSslCfgSSLCdpEntryTable{
		SlbNewSslCfgSSLCdpGrpIdIndex: slbNewSslCfgSSLCdpGrpIdIndex,
		SlbNewSslCfgSSLCdpEntryIndex: slbNewSslCfgSSLCdpEntryIndex,
		Params:                       params,
	}
}

func (c *SlbNewSslCfgSSLCdpEntryTable) Name() string {
	return "SlbNewSslCfgSSLCdpEntryTable"
}

func (c *SlbNewSslCfgSSLCdpEntryTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbNewSslCfgSSLCdpEntryTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbNewSslCfgSSLCdpEntryTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbNewSslCfgSSLCdpGrpIdIndex).IsZero() &&
		reflect.ValueOf(c.SlbNewSslCfgSSLCdpEntryIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbNewSslCfgSSLCdpGrpIdIndex) + "/" + fmt.Sprint(c.SlbNewSslCfgSSLCdpEntryIndex)
}

type SlbNewSslCfgSSLCdpEntryTableDelete int32

const (
	SlbNewSslCfgSSLCdpEntryTableDelete_Other       SlbNewSslCfgSSLCdpEntryTableDelete = 1
	SlbNewSslCfgSSLCdpEntryTableDelete_Delete      SlbNewSslCfgSSLCdpEntryTableDelete = 2
	SlbNewSslCfgSSLCdpEntryTableDelete_Unsupported SlbNewSslCfgSSLCdpEntryTableDelete = 2147483647
)

type SlbNewSslCfgSSLCdpEntryTableParams struct {
	// The CDP Group (key id) as an index.
	GrpIdIndex string `json:"GrpIdIndex,omitempty"`
	// The Cdp Entry number as an index.
	Index int32 `json:"Index,omitempty"`
	// Defines URL of specific object (file/folder) to be download by this rule.
	URL string `json:"URL,omitempty"`
	// The Cdp Entry URL User name.
	User string `json:"User,omitempty"`
	// The Cdp Entry URL Password.
	Password string `json:"Password,omitempty"`
	// By setting the value to delete(2), the entire row is deleted.
	Delete SlbNewSslCfgSSLCdpEntryTableDelete `json:"Delete,omitempty"`
}

func (p SlbNewSslCfgSSLCdpEntryTableParams) iMABean() {}
