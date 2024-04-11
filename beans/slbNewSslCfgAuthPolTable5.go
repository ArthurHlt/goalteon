package beans

import (
	"fmt"
	"reflect"
)

// SlbNewSslCfgAuthPolTable5 The table for configuring Client Authentication  policy.
// Note:This mib is not supported for VX instance of Virtualization.
type SlbNewSslCfgAuthPolTable5 struct {
	// The Auth policy name(key id) as an index.
	SlbNewSslCfgAuthPolTable5NameIdIndex string
	Params                               *SlbNewSslCfgAuthPolTable5Params
}

func NewSlbNewSslCfgAuthPolTable5List() *SlbNewSslCfgAuthPolTable5 {
	return &SlbNewSslCfgAuthPolTable5{}
}

func NewSlbNewSslCfgAuthPolTable5(
	slbNewSslCfgAuthPolTable5NameIdIndex string,
	params *SlbNewSslCfgAuthPolTable5Params,
) *SlbNewSslCfgAuthPolTable5 {
	return &SlbNewSslCfgAuthPolTable5{
		SlbNewSslCfgAuthPolTable5NameIdIndex: slbNewSslCfgAuthPolTable5NameIdIndex,
		Params:                               params,
	}
}

func (c *SlbNewSslCfgAuthPolTable5) Name() string {
	return "SlbNewSslCfgAuthPolTable5"
}

func (c *SlbNewSslCfgAuthPolTable5) GetParams() BeanType {
	return c.Params
}

func (c *SlbNewSslCfgAuthPolTable5) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbNewSslCfgAuthPolTable5) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbNewSslCfgAuthPolTable5NameIdIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbNewSslCfgAuthPolTable5NameIdIndex)
}

type SlbNewSslCfgAuthPolTable5AdminStatus int32

const (
	SlbNewSslCfgAuthPolTable5AdminStatus_Enabled     SlbNewSslCfgAuthPolTable5AdminStatus = 1
	SlbNewSslCfgAuthPolTable5AdminStatus_Disabled    SlbNewSslCfgAuthPolTable5AdminStatus = 2
	SlbNewSslCfgAuthPolTable5AdminStatus_Unsupported SlbNewSslCfgAuthPolTable5AdminStatus = 2147483647
)

type SlbNewSslCfgAuthPolTable5Delete int32

const (
	SlbNewSslCfgAuthPolTable5Delete_Other       SlbNewSslCfgAuthPolTable5Delete = 1
	SlbNewSslCfgAuthPolTable5Delete_Delete      SlbNewSslCfgAuthPolTable5Delete = 2
	SlbNewSslCfgAuthPolTable5Delete_Unsupported SlbNewSslCfgAuthPolTable5Delete = 2147483647
)

type SlbNewSslCfgAuthPolTable5PassinfoComp2424 int32

const (
	SlbNewSslCfgAuthPolTable5PassinfoComp2424_Enabled     SlbNewSslCfgAuthPolTable5PassinfoComp2424 = 1
	SlbNewSslCfgAuthPolTable5PassinfoComp2424_Disabled    SlbNewSslCfgAuthPolTable5PassinfoComp2424 = 2
	SlbNewSslCfgAuthPolTable5PassinfoComp2424_Unsupported SlbNewSslCfgAuthPolTable5PassinfoComp2424 = 2147483647
)

type SlbNewSslCfgAuthPolTable5Params struct {
	// The Auth policy name(key id) as an index.
	NameIdIndex string `json:"NameIdIndex,omitempty"`
	// Maximum depth to search the trusted CA in the CA certificate.
	Cadepth uint64 `json:"Cadepth,omitempty"`
	// Certificate's CA verification level.
	Caverify int32 `json:"Caverify,omitempty"`
	// URL for redirection when client authentication fails
	Failurl string `json:"Failurl,omitempty"`
	// Enable or disable Auth policy.
	AdminStatus SlbNewSslCfgAuthPolTable5AdminStatus `json:"AdminStatus,omitempty"`
	// Delete Client Authentication Policy.
	Delete SlbNewSslCfgAuthPolTable5Delete `json:"Delete,omitempty"`
	// Enable/Disable the 2424SSL Headers Compliance Mode.
	PassinfoComp2424 SlbNewSslCfgAuthPolTable5PassinfoComp2424 `json:"PassinfoComp2424,omitempty"`
}

func (p SlbNewSslCfgAuthPolTable5Params) iMABean() {}
