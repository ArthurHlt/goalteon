package beans

import (
	"fmt"
	"reflect"
)

// SlbCurSslCfgSSLCdpEntryTable The table for configuring CDP Entry.
// Note:This mib is not supported for VX instance of Virtualization.
type SlbCurSslCfgSSLCdpEntryTable struct {
	// The CDP Group (key id) as an index.
	SlbCurSslCfgSSLCdpGrpIdIndex string
	// The Cdp Entry number as an index.
	SlbCurSslCfgSSLCdpEntryIndex int32
	// The Cdp Entry number as an endpoint.
	SlbCurSslCfgSSLCdpEntryEndpoint int32
	Params                          *SlbCurSslCfgSSLCdpEntryTableParams
}

func NewSlbCurSslCfgSSLCdpEntryTableList() *SlbCurSslCfgSSLCdpEntryTable {
	return &SlbCurSslCfgSSLCdpEntryTable{}
}

func NewSlbCurSslCfgSSLCdpEntryTable(
	slbCurSslCfgSSLCdpGrpIdIndex string,
	slbCurSslCfgSSLCdpEntryIndex int32,
	slbCurSslCfgSSLCdpEntryEndpoint int32,
	params *SlbCurSslCfgSSLCdpEntryTableParams,
) *SlbCurSslCfgSSLCdpEntryTable {
	return &SlbCurSslCfgSSLCdpEntryTable{
		SlbCurSslCfgSSLCdpGrpIdIndex:    slbCurSslCfgSSLCdpGrpIdIndex,
		SlbCurSslCfgSSLCdpEntryIndex:    slbCurSslCfgSSLCdpEntryIndex,
		SlbCurSslCfgSSLCdpEntryEndpoint: slbCurSslCfgSSLCdpEntryEndpoint,
		Params:                          params,
	}
}

func (c *SlbCurSslCfgSSLCdpEntryTable) Name() string {
	return "SlbCurSslCfgSSLCdpEntryTable"
}

func (c *SlbCurSslCfgSSLCdpEntryTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbCurSslCfgSSLCdpEntryTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbCurSslCfgSSLCdpEntryTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbCurSslCfgSSLCdpGrpIdIndex).IsZero() &&
		reflect.ValueOf(c.SlbCurSslCfgSSLCdpEntryIndex).IsZero() &&
		reflect.ValueOf(c.SlbCurSslCfgSSLCdpEntryEndpoint).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbCurSslCfgSSLCdpGrpIdIndex) + "/" + fmt.Sprint(c.SlbCurSslCfgSSLCdpEntryIndex) + "/" + fmt.Sprint(c.SlbCurSslCfgSSLCdpEntryEndpoint)
}

type SlbCurSslCfgSSLCdpEntryTableParams struct {
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
	// The Cdp Entry number as an endpoint.
	Endpoint int32 `json:"Endpoint,omitempty"`
}

func (p SlbCurSslCfgSSLCdpEntryTableParams) iMABean() {}
