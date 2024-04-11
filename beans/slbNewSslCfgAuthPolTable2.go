package beans

import (
	"fmt"
	"reflect"
)

// SlbNewSslCfgAuthPolTable2 The table for configuring Client Authentication  policy.
// Note:This mib is not supported for VX instance of Virtualization.
type SlbNewSslCfgAuthPolTable2 struct {
	// The Auth policy name(key id) as an index.
	// Note:This mib is not supported for VX instance of Virtualization.
	SlbNewSslCfgAuthPolTable2NameIdIndex string
	Params                               *SlbNewSslCfgAuthPolTable2Params
}

func NewSlbNewSslCfgAuthPolTable2List() *SlbNewSslCfgAuthPolTable2 {
	return &SlbNewSslCfgAuthPolTable2{}
}

func NewSlbNewSslCfgAuthPolTable2(
	slbNewSslCfgAuthPolTable2NameIdIndex string,
	params *SlbNewSslCfgAuthPolTable2Params,
) *SlbNewSslCfgAuthPolTable2 {
	return &SlbNewSslCfgAuthPolTable2{
		SlbNewSslCfgAuthPolTable2NameIdIndex: slbNewSslCfgAuthPolTable2NameIdIndex,
		Params:                               params,
	}
}

func (c *SlbNewSslCfgAuthPolTable2) Name() string {
	return "SlbNewSslCfgAuthPolTable2"
}

func (c *SlbNewSslCfgAuthPolTable2) GetParams() BeanType {
	return c.Params
}

func (c *SlbNewSslCfgAuthPolTable2) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbNewSslCfgAuthPolTable2) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbNewSslCfgAuthPolTable2NameIdIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbNewSslCfgAuthPolTable2NameIdIndex)
}

type SlbNewSslCfgAuthPolTable2PassinfoSerialFlag int32

const (
	SlbNewSslCfgAuthPolTable2PassinfoSerialFlag_Enabled     SlbNewSslCfgAuthPolTable2PassinfoSerialFlag = 1
	SlbNewSslCfgAuthPolTable2PassinfoSerialFlag_Disabled    SlbNewSslCfgAuthPolTable2PassinfoSerialFlag = 2
	SlbNewSslCfgAuthPolTable2PassinfoSerialFlag_Unsupported SlbNewSslCfgAuthPolTable2PassinfoSerialFlag = 2147483647
)

type SlbNewSslCfgAuthPolTable2PassinfoAlgoFlag int32

const (
	SlbNewSslCfgAuthPolTable2PassinfoAlgoFlag_Enabled     SlbNewSslCfgAuthPolTable2PassinfoAlgoFlag = 1
	SlbNewSslCfgAuthPolTable2PassinfoAlgoFlag_Disabled    SlbNewSslCfgAuthPolTable2PassinfoAlgoFlag = 2
	SlbNewSslCfgAuthPolTable2PassinfoAlgoFlag_Unsupported SlbNewSslCfgAuthPolTable2PassinfoAlgoFlag = 2147483647
)

type SlbNewSslCfgAuthPolTable2PassinfoIssuerFlag int32

const (
	SlbNewSslCfgAuthPolTable2PassinfoIssuerFlag_Enabled     SlbNewSslCfgAuthPolTable2PassinfoIssuerFlag = 1
	SlbNewSslCfgAuthPolTable2PassinfoIssuerFlag_Disabled    SlbNewSslCfgAuthPolTable2PassinfoIssuerFlag = 2
	SlbNewSslCfgAuthPolTable2PassinfoIssuerFlag_Unsupported SlbNewSslCfgAuthPolTable2PassinfoIssuerFlag = 2147483647
)

type SlbNewSslCfgAuthPolTable2Params struct {
	// The Auth policy name(key id) as an index.
	// Note:This mib is not supported for VX instance of Virtualization.
	NameIdIndex string `json:"NameIdIndex,omitempty"`
	// Certificate serial-number header.
	PassinfoSerialName string `json:"PassinfoSerialName,omitempty"`
	// Pass certificate serial-number to backend server
	PassinfoSerialFlag SlbNewSslCfgAuthPolTable2PassinfoSerialFlag `json:"PassinfoSerialFlag,omitempty"`
	// Certificate signature algorithm header name.
	PassinfoAlgoName string `json:"PassinfoAlgoName,omitempty"`
	// Pass certificate Signature Algorithm to backend server
	PassinfoAlgoFlag SlbNewSslCfgAuthPolTable2PassinfoAlgoFlag `json:"PassinfoAlgoFlag,omitempty"`
	// Certificate issuer header.
	PassinfoIssuerName string `json:"PassinfoIssuerName,omitempty"`
	// Pass certificate issuer information to backend server
	PassinfoIssuerFlag SlbNewSslCfgAuthPolTable2PassinfoIssuerFlag `json:"PassinfoIssuerFlag,omitempty"`
}

func (p SlbNewSslCfgAuthPolTable2Params) iMABean() {}
