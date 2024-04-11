package beans

import (
	"fmt"
	"reflect"
)

// SlbCurDataclassCfgDataClassesTable Current data classes table.
// Note:This mib is not supported for VX instance of Virtualization.
type SlbCurDataclassCfgDataClassesTable struct {
	// Data class id.
	SlbCurDataclassCfgDataClassesId string
	Params                          *SlbCurDataclassCfgDataClassesTableParams
}

func NewSlbCurDataclassCfgDataClassesTableList() *SlbCurDataclassCfgDataClassesTable {
	return &SlbCurDataclassCfgDataClassesTable{}
}

func NewSlbCurDataclassCfgDataClassesTable(
	slbCurDataclassCfgDataClassesId string,
	params *SlbCurDataclassCfgDataClassesTableParams,
) *SlbCurDataclassCfgDataClassesTable {
	return &SlbCurDataclassCfgDataClassesTable{
		SlbCurDataclassCfgDataClassesId: slbCurDataclassCfgDataClassesId,
		Params:                          params,
	}
}

func (c *SlbCurDataclassCfgDataClassesTable) Name() string {
	return "SlbCurDataclassCfgDataClassesTable"
}

func (c *SlbCurDataclassCfgDataClassesTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbCurDataclassCfgDataClassesTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbCurDataclassCfgDataClassesTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbCurDataclassCfgDataClassesId).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbCurDataclassCfgDataClassesId)
}

type SlbCurDataclassCfgDataClassesTableDataType int32

const (
	SlbCurDataclassCfgDataClassesTableDataType_String      SlbCurDataclassCfgDataClassesTableDataType = 1
	SlbCurDataclassCfgDataClassesTableDataType_Ip          SlbCurDataclassCfgDataClassesTableDataType = 2
	SlbCurDataclassCfgDataClassesTableDataType_Unsupported SlbCurDataclassCfgDataClassesTableDataType = 2147483647
)

type SlbCurDataclassCfgDataClassesTableParams struct {
	// Data class id.
	Id string `json:"Id,omitempty"`
	// Data class data type.
	DataType SlbCurDataclassCfgDataClassesTableDataType `json:"DataType,omitempty"`
	// Data class name.
	Name string `json:"Name,omitempty"`
}

func (p SlbCurDataclassCfgDataClassesTableParams) iMABean() {}
