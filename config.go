package goalteaon

import (
	"fmt"
	"github.com/ArhurHlt/goalteon/beans"
	"net/url"
	"strings"
)

type SrcType string

const (
	ConfigSrcTypeTxt  SrcType = "txt"
	ConfigSrcTypeFile SrcType = "file"
)

type ConfigType string

const (
	ConfigTypeAll  ConfigType = "all"
	ConfigTypeVADC ConfigType = "vadc"
	ConfigTypePADC ConfigType = "padc"
)

type ConfigRecoveryType string

const (
	ConfigRecoveryTypeAll    ConfigRecoveryType = "all"
	ConfigRecoveryTypeVAdmin ConfigRecoveryType = "vadmin"
)

type ImpExpConfigParams struct {
	Pkey               bool
	Passphrase         string
	ManSync            bool
	ConfigType         ConfigType
	VADCList           []string
	ConfigRecoveryType ConfigRecoveryType
	SrcType            SrcType
}

func (i *ImpExpConfigParams) Values() url.Values {
	v := url.Values{}
	if i.Pkey {
		v.Set("pkey", "yes")
	} else {
		v.Set("pkey", "no")
	}
	if i.Passphrase != "" {
		v.Set("passphrase", i.Passphrase)
	}
	if i.ManSync {
		v.Set("mansync", "yes")
	} else {
		v.Set("mansync", "no")
	}
	if i.ConfigType != "" {
		v.Set("type", string(i.ConfigType))
	}
	if len(i.VADCList) > 0 {
		v.Set("vadc", strings.Join(i.VADCList, ","))
	}
	if i.ConfigRecoveryType != "" {
		v.Set("recovery", string(i.ConfigRecoveryType))
	}
	if i.SrcType == "" {
		i.SrcType = ConfigSrcTypeTxt
	}
	v.Set("src", string(i.SrcType))
	return v
}

func (c *Client) ExportConfig(params *ImpExpConfigParams) ([]byte, error) {
	if params == nil {
		return nil, fmt.Errorf("params are nil")
	}
	res, err := c.Get(&beans.GetCfgBean{}, params.Values())
	if err != nil {
		return nil, err
	}
	data := Translate[beans.BytesParams](res)
	return data, nil
}

func (c *Client) ImportConfig(data []byte, params *ImpExpConfigParams) (*StatusResponse, error) {
	if params == nil {
		return nil, fmt.Errorf("params are nil")
	}
	return c.Post(&beans.CfgImportBean{Content: data}, params.Values())
}
