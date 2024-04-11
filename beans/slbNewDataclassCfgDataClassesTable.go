package beans

import (
	"fmt"
	"reflect"
)

// SlbNewDataclassCfgDataClassesTable New data classes table.
// Note:This mib is not supported for VX instance of Virtualization.
type SlbNewDataclassCfgDataClassesTable struct {
	// Data class id.
	SlbNewDataclassCfgDataClassesId string
	Params                          *SlbNewDataclassCfgDataClassesTableParams
}

func NewSlbNewDataclassCfgDataClassesTableList() *SlbNewDataclassCfgDataClassesTable {
	return &SlbNewDataclassCfgDataClassesTable{}
}

func NewSlbNewDataclassCfgDataClassesTable(
	slbNewDataclassCfgDataClassesId string,
	params *SlbNewDataclassCfgDataClassesTableParams,
) *SlbNewDataclassCfgDataClassesTable {
	return &SlbNewDataclassCfgDataClassesTable{
		SlbNewDataclassCfgDataClassesId: slbNewDataclassCfgDataClassesId,
		Params:                          params,
	}
}

func (c *SlbNewDataclassCfgDataClassesTable) Name() string {
	return "SlbNewDataclassCfgDataClassesTable"
}

func (c *SlbNewDataclassCfgDataClassesTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbNewDataclassCfgDataClassesTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbNewDataclassCfgDataClassesTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbNewDataclassCfgDataClassesId).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbNewDataclassCfgDataClassesId)
}

type SlbNewDataclassCfgDataClassesTableDataType int32

const (
	SlbNewDataclassCfgDataClassesTableDataType_String      SlbNewDataclassCfgDataClassesTableDataType = 1
	SlbNewDataclassCfgDataClassesTableDataType_Ip          SlbNewDataclassCfgDataClassesTableDataType = 2
	SlbNewDataclassCfgDataClassesTableDataType_Unsupported SlbNewDataclassCfgDataClassesTableDataType = 2147483647
)

type SlbNewDataclassCfgDataClassesTableDel int32

const (
	SlbNewDataclassCfgDataClassesTableDel_Other       SlbNewDataclassCfgDataClassesTableDel = 1
	SlbNewDataclassCfgDataClassesTableDel_Delete      SlbNewDataclassCfgDataClassesTableDel = 2
	SlbNewDataclassCfgDataClassesTableDel_Unsupported SlbNewDataclassCfgDataClassesTableDel = 2147483647
)

type SlbNewDataclassCfgDataClassesTableParams struct {
	// Data class id.
	Id string `json:"Id,omitempty"`
	// Data class data type.
	DataType SlbNewDataclassCfgDataClassesTableDataType `json:"DataType,omitempty"`
	// Data class name.
	Name string `json:"Name,omitempty"`
	// By setting the value to delete(2), the entire row is deleted.
	Del SlbNewDataclassCfgDataClassesTableDel `json:"Del,omitempty"`
	// Duplicating an entire Data class by defining the destination Data class id.
	Copy string `json:"Copy,omitempty"`
}

func (p SlbNewDataclassCfgDataClassesTableParams) iMABean() {}
