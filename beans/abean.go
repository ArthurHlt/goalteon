package beans

import "reflect"

type Bean interface {
	Name() string
	Path() string
	GetParams() BeanType
	GetParamsType() reflect.Type
}

type ScalarBean interface {
	IMAScalarBean()
}

type BeanType interface {
	iMABean()
}

type MapParams map[string]interface{}

func (m MapParams) iMABean() {}

type BytesParams []byte

func (s BytesParams) iMABean() {}

type ConfigBean struct {
	Params MapParams
}

func (c *ConfigBean) Name() string {
	return ""
}

func (c *ConfigBean) Path() string {
	return "/config"
}

func (c *ConfigBean) GetParams() BeanType {
	return c.Params
}

func (c *ConfigBean) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *ConfigBean) IMAScalarBean() {}

type CfgImportBean struct {
	Content BytesParams
}

func (c *CfgImportBean) Name() string {
	return ""
}

func (c *CfgImportBean) Path() string {
	return "/config/configimport"
}

func (c *CfgImportBean) GetParams() BeanType {
	return c.Content
}

func (c *CfgImportBean) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Content)
}

func (c *CfgImportBean) IMAScalarBean() {}

type GetCfgBean struct {
	Content BytesParams
}

func (c *GetCfgBean) Name() string {
	return ""
}

func (c *GetCfgBean) Path() string {
	return "/config/getcfg"
}

func (c *GetCfgBean) GetParams() BeanType {
	return c.Content
}

func (c *GetCfgBean) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Content)
}

func (c *GetCfgBean) IMAScalarBean() {}

type SSLCertImportBean struct {
	Content BytesParams
}

func (c *SSLCertImportBean) Name() string {
	return ""
}

func (c *SSLCertImportBean) Path() string {
	return "/config/sslcertimport"
}

func (c *SSLCertImportBean) GetParams() BeanType {
	return c.Content
}

func (c *SSLCertImportBean) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Content)
}

func (c *SSLCertImportBean) IMAScalarBean() {}

type GetCertBean struct {
	Content BytesParams
}

func (c *GetCertBean) Name() string {
	return ""
}

func (c *GetCertBean) Path() string {
	return "/config/getcert"
}

func (c *GetCertBean) GetParams() BeanType {
	return c.Content
}

func (c *GetCertBean) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Content)
}

func (c *GetCertBean) IMAScalarBean() {}

type AppShapeImportBean struct {
	Content BytesParams
}

func (c *AppShapeImportBean) Name() string {
	return ""
}

func (c *AppShapeImportBean) Path() string {
	return "/config/appshapeimport"
}

func (c *AppShapeImportBean) GetParams() BeanType {
	return c.Content
}

func (c *AppShapeImportBean) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Content)
}

func (c *AppShapeImportBean) IMAScalarBean() {}

type GetAppShapeBean struct {
	Content BytesParams
}

func (c *GetAppShapeBean) Name() string {
	return ""
}

func (c *GetAppShapeBean) Path() string {
	return "/config/getappshape"
}

func (c *GetAppShapeBean) GetParams() BeanType {
	return c.Content
}

func (c *GetAppShapeBean) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Content)
}

func (c *GetAppShapeBean) IMAScalarBean() {}

type HcExtScriptImportBean struct {
	Content BytesParams
}

func (c *HcExtScriptImportBean) Name() string {
	return ""
}

func (c *HcExtScriptImportBean) Path() string {
	return "/config/exthcscriptimport"
}

func (c *HcExtScriptImportBean) GetParams() BeanType {
	return c.Content
}

func (c *HcExtScriptImportBean) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Content)
}

func (c *HcExtScriptImportBean) IMAScalarBean() {}

type GetHcExtScriptBean struct {
	Content BytesParams
}

func (c *GetHcExtScriptBean) Name() string {
	return ""
}

func (c *GetHcExtScriptBean) Path() string {
	return "/config/getexthcscript"
}

func (c *GetHcExtScriptBean) GetParams() BeanType {
	return c.Content
}

func (c *GetHcExtScriptBean) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Content)
}

func (c *GetHcExtScriptBean) IMAScalarBean() {}
