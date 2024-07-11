package beans

import (
	"fmt"
	"reflect"
)

// Layer7CurCfgContentClassTable The table for configuring Content Class.
// Note:This mib is not supported for VX instance of virtualization.
type Layer7CurCfgContentClassTable struct {
	// The content Class ID(key id) as an index.
	Layer7CurCfgContentClassID string
	Params                     *Layer7CurCfgContentClassTableParams
}

func NewLayer7CurCfgContentClassTableList() *Layer7CurCfgContentClassTable {
	return &Layer7CurCfgContentClassTable{}
}

func NewLayer7CurCfgContentClassTable(
	layer7CurCfgContentClassID string,
	params *Layer7CurCfgContentClassTableParams,
) *Layer7CurCfgContentClassTable {
	return &Layer7CurCfgContentClassTable{
		Layer7CurCfgContentClassID: layer7CurCfgContentClassID,
		Params:                     params,
	}
}

func (c *Layer7CurCfgContentClassTable) Name() string {
	return "Layer7CurCfgContentClassTable"
}

func (c *Layer7CurCfgContentClassTable) GetParams() BeanType {
	return c.Params
}

func (c *Layer7CurCfgContentClassTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *Layer7CurCfgContentClassTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.Layer7CurCfgContentClassID).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.Layer7CurCfgContentClassID)
}

type Layer7CurCfgContentClassTableHostName int32

const (
	Layer7CurCfgContentClassTableHostName_Yes         Layer7CurCfgContentClassTableHostName = 1
	Layer7CurCfgContentClassTableHostName_No          Layer7CurCfgContentClassTableHostName = 2
	Layer7CurCfgContentClassTableHostName_Unsupported Layer7CurCfgContentClassTableHostName = 2147483647
)

type Layer7CurCfgContentClassTablePath int32

const (
	Layer7CurCfgContentClassTablePath_Yes         Layer7CurCfgContentClassTablePath = 1
	Layer7CurCfgContentClassTablePath_No          Layer7CurCfgContentClassTablePath = 2
	Layer7CurCfgContentClassTablePath_Unsupported Layer7CurCfgContentClassTablePath = 2147483647
)

type Layer7CurCfgContentClassTableFileName int32

const (
	Layer7CurCfgContentClassTableFileName_Yes         Layer7CurCfgContentClassTableFileName = 1
	Layer7CurCfgContentClassTableFileName_No          Layer7CurCfgContentClassTableFileName = 2
	Layer7CurCfgContentClassTableFileName_Unsupported Layer7CurCfgContentClassTableFileName = 2147483647
)

type Layer7CurCfgContentClassTableFileType int32

const (
	Layer7CurCfgContentClassTableFileType_Yes         Layer7CurCfgContentClassTableFileType = 1
	Layer7CurCfgContentClassTableFileType_No          Layer7CurCfgContentClassTableFileType = 2
	Layer7CurCfgContentClassTableFileType_Unsupported Layer7CurCfgContentClassTableFileType = 2147483647
)

type Layer7CurCfgContentClassTableHeader int32

const (
	Layer7CurCfgContentClassTableHeader_Yes         Layer7CurCfgContentClassTableHeader = 1
	Layer7CurCfgContentClassTableHeader_No          Layer7CurCfgContentClassTableHeader = 2
	Layer7CurCfgContentClassTableHeader_Unsupported Layer7CurCfgContentClassTableHeader = 2147483647
)

type Layer7CurCfgContentClassTableCookie int32

const (
	Layer7CurCfgContentClassTableCookie_Yes         Layer7CurCfgContentClassTableCookie = 1
	Layer7CurCfgContentClassTableCookie_No          Layer7CurCfgContentClassTableCookie = 2
	Layer7CurCfgContentClassTableCookie_Unsupported Layer7CurCfgContentClassTableCookie = 2147483647
)

type Layer7CurCfgContentClassTableText int32

const (
	Layer7CurCfgContentClassTableText_Yes         Layer7CurCfgContentClassTableText = 1
	Layer7CurCfgContentClassTableText_No          Layer7CurCfgContentClassTableText = 2
	Layer7CurCfgContentClassTableText_Unsupported Layer7CurCfgContentClassTableText = 2147483647
)

type Layer7CurCfgContentClassTableXMLTag int32

const (
	Layer7CurCfgContentClassTableXMLTag_Yes         Layer7CurCfgContentClassTableXMLTag = 1
	Layer7CurCfgContentClassTableXMLTag_No          Layer7CurCfgContentClassTableXMLTag = 2
	Layer7CurCfgContentClassTableXMLTag_Unsupported Layer7CurCfgContentClassTableXMLTag = 2147483647
)

type Layer7CurCfgContentClassTableType int32

const (
	Layer7CurCfgContentClassTableType_Http        Layer7CurCfgContentClassTableType = 1
	Layer7CurCfgContentClassTableType_Rtsp        Layer7CurCfgContentClassTableType = 3
	Layer7CurCfgContentClassTableType_Ssl         Layer7CurCfgContentClassTableType = 6
	Layer7CurCfgContentClassTableType_Http2       Layer7CurCfgContentClassTableType = 7
	Layer7CurCfgContentClassTableType_Unsupported Layer7CurCfgContentClassTableType = 2147483647
)

type Layer7CurCfgContentClassTableParams struct {
	// The content Class ID(key id) as an index.
	ID string `json:"ID,omitempty"`
	// Content Class name.
	Name string `json:"Name,omitempty"`
	// Enter logical expression needs to be applied between classes.
	LogicalExpression string `json:"LogicalExpression,omitempty"`
	// URL Hostname table is not empty for current content class.
	HostName Layer7CurCfgContentClassTableHostName `json:"HostName,omitempty"`
	// URL path table is not empty for current content class.
	Path Layer7CurCfgContentClassTablePath `json:"Path,omitempty"`
	// URL file name table is not empty for current content class.
	FileName Layer7CurCfgContentClassTableFileName `json:"FileName,omitempty"`
	// URL file type table is not empty for current content class.
	FileType Layer7CurCfgContentClassTableFileType `json:"FileType,omitempty"`
	// Header table is not empty for current content class.
	Header Layer7CurCfgContentClassTableHeader `json:"Header,omitempty"`
	// Cookie table is not empty for current content class.
	Cookie Layer7CurCfgContentClassTableCookie `json:"Cookie,omitempty"`
	// Text table is not empty for current content class.
	Text Layer7CurCfgContentClassTableText `json:"Text,omitempty"`
	// XML tag table is not empty for current content class.
	XMLTag Layer7CurCfgContentClassTableXMLTag `json:"XMLTag,omitempty"`
	// Content Class type.
	Type Layer7CurCfgContentClassTableType `json:"Type,omitempty"`
}

func (p Layer7CurCfgContentClassTableParams) iMABean() {}
