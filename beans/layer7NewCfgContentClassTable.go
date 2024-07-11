package beans

import (
	"fmt"
	"reflect"
)

// Layer7NewCfgContentClassTable The table for configuring Content Class.
// Note:This mib is not supported for VX instance of virtualization.
type Layer7NewCfgContentClassTable struct {
	// The content Class ID(key id) as an index.
	Layer7NewCfgContentClassID string
	Params                     *Layer7NewCfgContentClassTableParams
}

func NewLayer7NewCfgContentClassTableList() *Layer7NewCfgContentClassTable {
	return &Layer7NewCfgContentClassTable{}
}

func NewLayer7NewCfgContentClassTable(
	layer7NewCfgContentClassID string,
	params *Layer7NewCfgContentClassTableParams,
) *Layer7NewCfgContentClassTable {
	return &Layer7NewCfgContentClassTable{
		Layer7NewCfgContentClassID: layer7NewCfgContentClassID,
		Params:                     params,
	}
}

func (c *Layer7NewCfgContentClassTable) Name() string {
	return "Layer7NewCfgContentClassTable"
}

func (c *Layer7NewCfgContentClassTable) GetParams() BeanType {
	return c.Params
}

func (c *Layer7NewCfgContentClassTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *Layer7NewCfgContentClassTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.Layer7NewCfgContentClassID).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.Layer7NewCfgContentClassID)
}

type Layer7NewCfgContentClassTableHostName int32

const (
	Layer7NewCfgContentClassTableHostName_Yes         Layer7NewCfgContentClassTableHostName = 1
	Layer7NewCfgContentClassTableHostName_No          Layer7NewCfgContentClassTableHostName = 2
	Layer7NewCfgContentClassTableHostName_Unsupported Layer7NewCfgContentClassTableHostName = 2147483647
)

type Layer7NewCfgContentClassTablePath int32

const (
	Layer7NewCfgContentClassTablePath_Yes         Layer7NewCfgContentClassTablePath = 1
	Layer7NewCfgContentClassTablePath_No          Layer7NewCfgContentClassTablePath = 2
	Layer7NewCfgContentClassTablePath_Unsupported Layer7NewCfgContentClassTablePath = 2147483647
)

type Layer7NewCfgContentClassTableFileName int32

const (
	Layer7NewCfgContentClassTableFileName_Yes         Layer7NewCfgContentClassTableFileName = 1
	Layer7NewCfgContentClassTableFileName_No          Layer7NewCfgContentClassTableFileName = 2
	Layer7NewCfgContentClassTableFileName_Unsupported Layer7NewCfgContentClassTableFileName = 2147483647
)

type Layer7NewCfgContentClassTableFileType int32

const (
	Layer7NewCfgContentClassTableFileType_Yes         Layer7NewCfgContentClassTableFileType = 1
	Layer7NewCfgContentClassTableFileType_No          Layer7NewCfgContentClassTableFileType = 2
	Layer7NewCfgContentClassTableFileType_Unsupported Layer7NewCfgContentClassTableFileType = 2147483647
)

type Layer7NewCfgContentClassTableHeader int32

const (
	Layer7NewCfgContentClassTableHeader_Yes         Layer7NewCfgContentClassTableHeader = 1
	Layer7NewCfgContentClassTableHeader_No          Layer7NewCfgContentClassTableHeader = 2
	Layer7NewCfgContentClassTableHeader_Unsupported Layer7NewCfgContentClassTableHeader = 2147483647
)

type Layer7NewCfgContentClassTableCookie int32

const (
	Layer7NewCfgContentClassTableCookie_Yes         Layer7NewCfgContentClassTableCookie = 1
	Layer7NewCfgContentClassTableCookie_No          Layer7NewCfgContentClassTableCookie = 2
	Layer7NewCfgContentClassTableCookie_Unsupported Layer7NewCfgContentClassTableCookie = 2147483647
)

type Layer7NewCfgContentClassTableText int32

const (
	Layer7NewCfgContentClassTableText_Yes         Layer7NewCfgContentClassTableText = 1
	Layer7NewCfgContentClassTableText_No          Layer7NewCfgContentClassTableText = 2
	Layer7NewCfgContentClassTableText_Unsupported Layer7NewCfgContentClassTableText = 2147483647
)

type Layer7NewCfgContentClassTableXMLTag int32

const (
	Layer7NewCfgContentClassTableXMLTag_Yes         Layer7NewCfgContentClassTableXMLTag = 1
	Layer7NewCfgContentClassTableXMLTag_No          Layer7NewCfgContentClassTableXMLTag = 2
	Layer7NewCfgContentClassTableXMLTag_Unsupported Layer7NewCfgContentClassTableXMLTag = 2147483647
)

type Layer7NewCfgContentClassTableDelete int32

const (
	Layer7NewCfgContentClassTableDelete_Other       Layer7NewCfgContentClassTableDelete = 1
	Layer7NewCfgContentClassTableDelete_Delete      Layer7NewCfgContentClassTableDelete = 2
	Layer7NewCfgContentClassTableDelete_Unsupported Layer7NewCfgContentClassTableDelete = 2147483647
)

type Layer7NewCfgContentClassTableType int32

const (
	Layer7NewCfgContentClassTableType_Http        Layer7NewCfgContentClassTableType = 1
	Layer7NewCfgContentClassTableType_Rtsp        Layer7NewCfgContentClassTableType = 3
	Layer7NewCfgContentClassTableType_Ssl         Layer7NewCfgContentClassTableType = 6
	Layer7NewCfgContentClassTableType_Http2       Layer7NewCfgContentClassTableType = 7
	Layer7NewCfgContentClassTableType_Unsupported Layer7NewCfgContentClassTableType = 2147483647
)

type Layer7NewCfgContentClassTableParams struct {
	// The content Class ID(key id) as an index.
	ID string `json:"ID,omitempty"`
	// Content Class name.
	Name string `json:"Name,omitempty"`
	// Enter logical expression needs to be applied between classes..
	LogicalExpression string `json:"LogicalExpression,omitempty"`
	// URL host name table is not empty for current content class.
	HostName Layer7NewCfgContentClassTableHostName `json:"HostName,omitempty"`
	// URL path table is not empty for current content class.
	Path Layer7NewCfgContentClassTablePath `json:"Path,omitempty"`
	// URL file name table is not empty for current content class.
	FileName Layer7NewCfgContentClassTableFileName `json:"FileName,omitempty"`
	// URL file type table is not empty for current content class.
	FileType Layer7NewCfgContentClassTableFileType `json:"FileType,omitempty"`
	// Header table is not empty for current content class.
	Header Layer7NewCfgContentClassTableHeader `json:"Header,omitempty"`
	// Cookie table is not empty for current content class.
	Cookie Layer7NewCfgContentClassTableCookie `json:"Cookie,omitempty"`
	// Text table is not empty for current content class.
	Text Layer7NewCfgContentClassTableText `json:"Text,omitempty"`
	// XML tag table is not empty for current content class.
	XMLTag Layer7NewCfgContentClassTableXMLTag `json:"XMLTag,omitempty"`
	// By setting the value to delete(2), the entire row is deleted.
	Delete Layer7NewCfgContentClassTableDelete `json:"Delete,omitempty"`
	// This is an action object.Enter the content Class ID(key id)
	// to which the current content Class has to be copied.
	// Value 1 is returned always when read this object.
	Copy string `json:"Copy,omitempty"`
	// Content Class Type.
	Type Layer7NewCfgContentClassTableType `json:"Type,omitempty"`
}

func (p Layer7NewCfgContentClassTableParams) iMABean() {}
