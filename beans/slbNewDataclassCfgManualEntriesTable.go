package beans

import (
	"fmt"
	"reflect"
)

// SlbNewDataclassCfgManualEntriesTable New data class manual entries table.
// Note:This mib is not supported for VX instance of Virtualization.
type SlbNewDataclassCfgManualEntriesTable struct {
	// Data class ID which the manual entry belongs to.
	SlbNewDataclassCfgManualEntriesEntryDcId string
	// Manual entry Id.
	SlbNewDataclassCfgManualEntriesEntryId int32
	Params                                 *SlbNewDataclassCfgManualEntriesTableParams
}

func NewSlbNewDataclassCfgManualEntriesTableList() *SlbNewDataclassCfgManualEntriesTable {
	return &SlbNewDataclassCfgManualEntriesTable{}
}

func NewSlbNewDataclassCfgManualEntriesTable(
	slbNewDataclassCfgManualEntriesEntryDcId string,
	slbNewDataclassCfgManualEntriesEntryId int32,
	params *SlbNewDataclassCfgManualEntriesTableParams,
) *SlbNewDataclassCfgManualEntriesTable {
	return &SlbNewDataclassCfgManualEntriesTable{
		SlbNewDataclassCfgManualEntriesEntryDcId: slbNewDataclassCfgManualEntriesEntryDcId,
		SlbNewDataclassCfgManualEntriesEntryId:   slbNewDataclassCfgManualEntriesEntryId,
		Params:                                   params,
	}
}

func (c *SlbNewDataclassCfgManualEntriesTable) Name() string {
	return "SlbNewDataclassCfgManualEntriesTable"
}

func (c *SlbNewDataclassCfgManualEntriesTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbNewDataclassCfgManualEntriesTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbNewDataclassCfgManualEntriesTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbNewDataclassCfgManualEntriesEntryDcId).IsZero() &&
		reflect.ValueOf(c.SlbNewDataclassCfgManualEntriesEntryId).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbNewDataclassCfgManualEntriesEntryDcId) + "/" + fmt.Sprint(c.SlbNewDataclassCfgManualEntriesEntryId)
}

type SlbNewDataclassCfgManualEntriesTableDel int32

const (
	SlbNewDataclassCfgManualEntriesTableDel_Other       SlbNewDataclassCfgManualEntriesTableDel = 1
	SlbNewDataclassCfgManualEntriesTableDel_Delete      SlbNewDataclassCfgManualEntriesTableDel = 2
	SlbNewDataclassCfgManualEntriesTableDel_Unsupported SlbNewDataclassCfgManualEntriesTableDel = 2147483647
)

type SlbNewDataclassCfgManualEntriesTableParams struct {
	// Data class ID which the manual entry belongs to.
	DcId string `json:"DcId,omitempty"`
	// Manual entry Id.
	Id int32 `json:"Id,omitempty"`
	// Manual entry key. When configuration type is string there are no limitations.
	// When configuration type is ip should be a discrete IPv4/v6 address or subnet.
	Key string `json:"Key,omitempty"`
	// Manual entry value.
	Value string `json:"Value,omitempty"`
	// By setting the value to delete(2), the entire row is deleted.
	Del SlbNewDataclassCfgManualEntriesTableDel `json:"Del,omitempty"`
}

func (p SlbNewDataclassCfgManualEntriesTableParams) iMABean() {}
