package beans

import (
	"fmt"
	"reflect"
)

// SlbCurDataclassCfgManualEntriesTable Current data class manual entries table.
// Note:This mib is not supported for VX instance of Virtualization.
type SlbCurDataclassCfgManualEntriesTable struct {
	// Data class ID which the manual entry belongs to.
	SlbCurDataclassCfgManualEntriesEntryDcId string
	// Manual entry Id.
	SlbCurDataclassCfgManualEntriesEntryId int32
	Params                                 *SlbCurDataclassCfgManualEntriesTableParams
}

func NewSlbCurDataclassCfgManualEntriesTableList() *SlbCurDataclassCfgManualEntriesTable {
	return &SlbCurDataclassCfgManualEntriesTable{}
}

func NewSlbCurDataclassCfgManualEntriesTable(
	slbCurDataclassCfgManualEntriesEntryDcId string,
	slbCurDataclassCfgManualEntriesEntryId int32,
	params *SlbCurDataclassCfgManualEntriesTableParams,
) *SlbCurDataclassCfgManualEntriesTable {
	return &SlbCurDataclassCfgManualEntriesTable{
		SlbCurDataclassCfgManualEntriesEntryDcId: slbCurDataclassCfgManualEntriesEntryDcId,
		SlbCurDataclassCfgManualEntriesEntryId:   slbCurDataclassCfgManualEntriesEntryId,
		Params:                                   params,
	}
}

func (c *SlbCurDataclassCfgManualEntriesTable) Name() string {
	return "SlbCurDataclassCfgManualEntriesTable"
}

func (c *SlbCurDataclassCfgManualEntriesTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbCurDataclassCfgManualEntriesTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbCurDataclassCfgManualEntriesTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbCurDataclassCfgManualEntriesEntryDcId).IsZero() &&
		reflect.ValueOf(c.SlbCurDataclassCfgManualEntriesEntryId).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbCurDataclassCfgManualEntriesEntryDcId) + "/" + fmt.Sprint(c.SlbCurDataclassCfgManualEntriesEntryId)
}

type SlbCurDataclassCfgManualEntriesTableParams struct {
	// Data class ID which the manual entry belongs to.
	DcId string `json:"DcId,omitempty"`
	// Manual entry Id.
	Id int32 `json:"Id,omitempty"`
	// Manual entry key. When configuration type is string there are no limitations.
	// When configuration type is ip should be a discrete IPv4/v6 address or subnet.
	Key string `json:"Key,omitempty"`
	// Manual entry value.
	Value string `json:"Value,omitempty"`
}

func (p SlbCurDataclassCfgManualEntriesTableParams) iMABean() {}
