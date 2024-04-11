package beans

import (
	"fmt"
	"reflect"
)

// SlbNewSslCfgAuthPolTable3 The table for configuring Client Authentication  policy.
// Note:This mib is not supported for VX instance of Virtualization.
type SlbNewSslCfgAuthPolTable3 struct {
	// The Auth policy name(key id) as an index.
	SlbNewSslCfgAuthPolTable3NameIdIndex string
	Params                               *SlbNewSslCfgAuthPolTable3Params
}

func NewSlbNewSslCfgAuthPolTable3List() *SlbNewSslCfgAuthPolTable3 {
	return &SlbNewSslCfgAuthPolTable3{}
}

func NewSlbNewSslCfgAuthPolTable3(
	slbNewSslCfgAuthPolTable3NameIdIndex string,
	params *SlbNewSslCfgAuthPolTable3Params,
) *SlbNewSslCfgAuthPolTable3 {
	return &SlbNewSslCfgAuthPolTable3{
		SlbNewSslCfgAuthPolTable3NameIdIndex: slbNewSslCfgAuthPolTable3NameIdIndex,
		Params:                               params,
	}
}

func (c *SlbNewSslCfgAuthPolTable3) Name() string {
	return "SlbNewSslCfgAuthPolTable3"
}

func (c *SlbNewSslCfgAuthPolTable3) GetParams() BeanType {
	return c.Params
}

func (c *SlbNewSslCfgAuthPolTable3) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbNewSslCfgAuthPolTable3) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbNewSslCfgAuthPolTable3NameIdIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbNewSslCfgAuthPolTable3NameIdIndex)
}

type SlbNewSslCfgAuthPolTable3PassinfoNbeforeFlag int32

const (
	SlbNewSslCfgAuthPolTable3PassinfoNbeforeFlag_Enabled     SlbNewSslCfgAuthPolTable3PassinfoNbeforeFlag = 1
	SlbNewSslCfgAuthPolTable3PassinfoNbeforeFlag_Disabled    SlbNewSslCfgAuthPolTable3PassinfoNbeforeFlag = 2
	SlbNewSslCfgAuthPolTable3PassinfoNbeforeFlag_Unsupported SlbNewSslCfgAuthPolTable3PassinfoNbeforeFlag = 2147483647
)

type SlbNewSslCfgAuthPolTable3PassinfoNafterFlag int32

const (
	SlbNewSslCfgAuthPolTable3PassinfoNafterFlag_Enabled     SlbNewSslCfgAuthPolTable3PassinfoNafterFlag = 1
	SlbNewSslCfgAuthPolTable3PassinfoNafterFlag_Disabled    SlbNewSslCfgAuthPolTable3PassinfoNafterFlag = 2
	SlbNewSslCfgAuthPolTable3PassinfoNafterFlag_Unsupported SlbNewSslCfgAuthPolTable3PassinfoNafterFlag = 2147483647
)

type SlbNewSslCfgAuthPolTable3PassinfoSubjectFlag int32

const (
	SlbNewSslCfgAuthPolTable3PassinfoSubjectFlag_Enabled     SlbNewSslCfgAuthPolTable3PassinfoSubjectFlag = 1
	SlbNewSslCfgAuthPolTable3PassinfoSubjectFlag_Disabled    SlbNewSslCfgAuthPolTable3PassinfoSubjectFlag = 2
	SlbNewSslCfgAuthPolTable3PassinfoSubjectFlag_Unsupported SlbNewSslCfgAuthPolTable3PassinfoSubjectFlag = 2147483647
)

type SlbNewSslCfgAuthPolTable3Params struct {
	// The Auth policy name(key id) as an index.
	NameIdIndex string `json:"NameIdIndex,omitempty"`
	// Certificate 'Not Before Validity Dates' header.
	PassinfoNbeforeName string `json:"PassinfoNbeforeName,omitempty"`
	// Pass certificate 'Not Before' Validity Date to backend server
	PassinfoNbeforeFlag SlbNewSslCfgAuthPolTable3PassinfoNbeforeFlag `json:"PassinfoNbeforeFlag,omitempty"`
	// Certificate 'Not After Validity Dates' header.
	PassinfoNafterName string `json:"PassinfoNafterName,omitempty"`
	// Pass certificate 'Not After' Validity Date to backend server
	PassinfoNafterFlag SlbNewSslCfgAuthPolTable3PassinfoNafterFlag `json:"PassinfoNafterFlag,omitempty"`
	// Certificate subject header name.
	PassinfoSubjectName string `json:"PassinfoSubjectName,omitempty"`
	// Pass certificate subject to backend server
	PassinfoSubjectFlag SlbNewSslCfgAuthPolTable3PassinfoSubjectFlag `json:"PassinfoSubjectFlag,omitempty"`
}

func (p SlbNewSslCfgAuthPolTable3Params) iMABean() {}
