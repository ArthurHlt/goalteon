package beans

import (
	"fmt"
	"reflect"
)

// SlbNewSslCfgAuthPolTable4 The table for configuring Client Authentication  policy.
// Note:This mib is not supported for VX instance of Virtualization.
type SlbNewSslCfgAuthPolTable4 struct {
	// The Auth policy name(key id) as an index.
	SlbNewSslCfgAuthPolTable4NameIdIndex string
	Params                               *SlbNewSslCfgAuthPolTable4Params
}

func NewSlbNewSslCfgAuthPolTable4List() *SlbNewSslCfgAuthPolTable4 {
	return &SlbNewSslCfgAuthPolTable4{}
}

func NewSlbNewSslCfgAuthPolTable4(
	slbNewSslCfgAuthPolTable4NameIdIndex string,
	params *SlbNewSslCfgAuthPolTable4Params,
) *SlbNewSslCfgAuthPolTable4 {
	return &SlbNewSslCfgAuthPolTable4{
		SlbNewSslCfgAuthPolTable4NameIdIndex: slbNewSslCfgAuthPolTable4NameIdIndex,
		Params:                               params,
	}
}

func (c *SlbNewSslCfgAuthPolTable4) Name() string {
	return "SlbNewSslCfgAuthPolTable4"
}

func (c *SlbNewSslCfgAuthPolTable4) GetParams() BeanType {
	return c.Params
}

func (c *SlbNewSslCfgAuthPolTable4) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbNewSslCfgAuthPolTable4) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbNewSslCfgAuthPolTable4NameIdIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbNewSslCfgAuthPolTable4NameIdIndex)
}

type SlbNewSslCfgAuthPolTable4PassinfoKeytypeFlag int32

const (
	SlbNewSslCfgAuthPolTable4PassinfoKeytypeFlag_Enabled     SlbNewSslCfgAuthPolTable4PassinfoKeytypeFlag = 1
	SlbNewSslCfgAuthPolTable4PassinfoKeytypeFlag_Disabled    SlbNewSslCfgAuthPolTable4PassinfoKeytypeFlag = 2
	SlbNewSslCfgAuthPolTable4PassinfoKeytypeFlag_Unsupported SlbNewSslCfgAuthPolTable4PassinfoKeytypeFlag = 2147483647
)

type SlbNewSslCfgAuthPolTable4PassinfoMd5Flag int32

const (
	SlbNewSslCfgAuthPolTable4PassinfoMd5Flag_Enabled     SlbNewSslCfgAuthPolTable4PassinfoMd5Flag = 1
	SlbNewSslCfgAuthPolTable4PassinfoMd5Flag_Disabled    SlbNewSslCfgAuthPolTable4PassinfoMd5Flag = 2
	SlbNewSslCfgAuthPolTable4PassinfoMd5Flag_Unsupported SlbNewSslCfgAuthPolTable4PassinfoMd5Flag = 2147483647
)

type SlbNewSslCfgAuthPolTable4PassinfoCertFlag int32

const (
	SlbNewSslCfgAuthPolTable4PassinfoCertFlag_Enabled     SlbNewSslCfgAuthPolTable4PassinfoCertFlag = 1
	SlbNewSslCfgAuthPolTable4PassinfoCertFlag_Disabled    SlbNewSslCfgAuthPolTable4PassinfoCertFlag = 2
	SlbNewSslCfgAuthPolTable4PassinfoCertFlag_Unsupported SlbNewSslCfgAuthPolTable4PassinfoCertFlag = 2147483647
)

type SlbNewSslCfgAuthPolTable4Params struct {
	// The Auth policy name(key id) as an index.
	NameIdIndex string `json:"NameIdIndex,omitempty"`
	// Certificate Public Key Type header.
	PassinfoKeytypeName string `json:"PassinfoKeytypeName,omitempty"`
	// Pass certificate Public Key Type information to backend servers
	PassinfoKeytypeFlag SlbNewSslCfgAuthPolTable4PassinfoKeytypeFlag `json:"PassinfoKeytypeFlag,omitempty"`
	// Certificate MD5 Hash header.
	PassinfoMd5Name string `json:"PassinfoMd5Name,omitempty"`
	// Pass certificate MD5 Hash information to backend servers
	PassinfoMd5Flag SlbNewSslCfgAuthPolTable4PassinfoMd5Flag `json:"PassinfoMd5Flag,omitempty"`
	// Certificate header.
	PassinfoCertName string `json:"PassinfoCertName,omitempty"`
	// Certificate Header Lines Format.
	PassinfoCertFormat int32 `json:"PassinfoCertFormat,omitempty"`
	// Pass certificate information to backend servers
	PassinfoCertFlag SlbNewSslCfgAuthPolTable4PassinfoCertFlag `json:"PassinfoCertFlag,omitempty"`
	// Set the character set to be used for information.
	PassinfoCharset int32 `json:"PassinfoCharset,omitempty"`
	// Trusted client's CA certificate name.
	TrustcaChainName string `json:"TrustcaChainName,omitempty"`
	// Trusted client's CA certificate type certificate=cert,Group=group,None=empty string.
	TrustcaChainType string `json:"TrustcaChainType,omitempty"`
	// Client CA for Request certificate or group name.
	ClientcaReqChainName string `json:"ClientcaReqChainName,omitempty"`
	// Client CA for Request type certificate=cert,Group=group,None=none,Default=default.
	ClientcaReqChainType string `json:"ClientcaReqChainType,omitempty"`
}

func (p SlbNewSslCfgAuthPolTable4Params) iMABean() {}
