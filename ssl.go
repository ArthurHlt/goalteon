package goalteon

import (
	"fmt"
	"github.com/ArthurHlt/goalteon/beans"
	"net/url"
)

type SslType string

const (
	SslTypePair SslType = "pair"
	SslTypeCert SslType = "cert"
	SslTypeKey  SslType = "key"
	SslTypeInCA SslType = "inca"
	SslTypeCLCA SslType = "clca"
)

type ImpExpSslParams struct {
	ID         string
	SslType    SslType
	Passphrase string
	SrcType    SrcType
	RenewId    string
}

func (i *ImpExpSslParams) Values() url.Values {
	v := url.Values{}
	v.Set("id", i.ID)
	if i.Passphrase != "" {
		v.Set("passphrase", i.Passphrase)
	}
	if i.SslType != "" {
		v.Set("type", string(i.SslType))
	}
	if i.SrcType == "" {
		i.SrcType = ConfigSrcTypeTxt
	}
	if i.RenewId != "" {
		v.Set("renew", "1")
		v.Set("id", i.RenewId)
	}
	v.Set("src", string(i.SrcType))
	return v
}

func (c *Client) ExportSsl(params *ImpExpSslParams) ([]byte, error) {
	if params == nil {
		return nil, fmt.Errorf("params are nil")
	}
	if params.ID == "" {
		return nil, fmt.Errorf("id is required")
	}
	res, err := c.Get(&beans.GetCertBean{}, params.Values())
	if err != nil {
		return nil, err
	}
	data := Translate[beans.BytesParams](res)
	return data, nil
}

func (c *Client) ImportSsl(data []byte, params *ImpExpSslParams) (*StatusResponse, error) {
	if params == nil {
		return nil, fmt.Errorf("params are nil")
	}
	if params.ID == "" {
		return nil, fmt.Errorf("id is required")
	}
	return c.Post(&beans.SSLCertImportBean{Content: data}, params.Values())
}
