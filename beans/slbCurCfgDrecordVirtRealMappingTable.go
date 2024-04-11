package beans

import (
	"fmt"
	"reflect"
)

// SlbCurCfgDrecordVirtRealMappingTable The table of Domain records and virtual server and real
// server mappings.
// Note:This mib is not supported for VX instance of Virtualization.
type SlbCurCfgDrecordVirtRealMappingTable struct {
	// The number of the domain record.
	SlbCurCfgDomainRecordIndex uint64
	// The number of the entry virtual real server mapping.
	SlbCurCfgEntryIndex uint64
	Params              *SlbCurCfgDrecordVirtRealMappingTableParams
}

func NewSlbCurCfgDrecordVirtRealMappingTableList() *SlbCurCfgDrecordVirtRealMappingTable {
	return &SlbCurCfgDrecordVirtRealMappingTable{}
}

func NewSlbCurCfgDrecordVirtRealMappingTable(
	slbCurCfgDomainRecordIndex uint64,
	slbCurCfgEntryIndex uint64,
	params *SlbCurCfgDrecordVirtRealMappingTableParams,
) *SlbCurCfgDrecordVirtRealMappingTable {
	return &SlbCurCfgDrecordVirtRealMappingTable{
		SlbCurCfgDomainRecordIndex: slbCurCfgDomainRecordIndex,
		SlbCurCfgEntryIndex:        slbCurCfgEntryIndex,
		Params:                     params,
	}
}

func (c *SlbCurCfgDrecordVirtRealMappingTable) Name() string {
	return "SlbCurCfgDrecordVirtRealMappingTable"
}

func (c *SlbCurCfgDrecordVirtRealMappingTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbCurCfgDrecordVirtRealMappingTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbCurCfgDrecordVirtRealMappingTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbCurCfgDomainRecordIndex).IsZero() &&
		reflect.ValueOf(c.SlbCurCfgEntryIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbCurCfgDomainRecordIndex) + "/" + fmt.Sprint(c.SlbCurCfgEntryIndex)
}

type SlbCurCfgDrecordVirtRealMappingTableEntryState int32

const (
	SlbCurCfgDrecordVirtRealMappingTableEntryState_Enabled     SlbCurCfgDrecordVirtRealMappingTableEntryState = 1
	SlbCurCfgDrecordVirtRealMappingTableEntryState_Disabled    SlbCurCfgDrecordVirtRealMappingTableEntryState = 2
	SlbCurCfgDrecordVirtRealMappingTableEntryState_Unsupported SlbCurCfgDrecordVirtRealMappingTableEntryState = 2147483647
)

type SlbCurCfgDrecordVirtRealMappingTableParams struct {
	// The number of the domain record.
	DomainRecordIndex uint64 `json:"DomainRecordIndex,omitempty"`
	// The number of the entry virtual real server mapping.
	EntryIndex uint64 `json:"EntryIndex,omitempty"`
	// The virtual server number.
	Server int32 `json:"Server,omitempty"`
	// The Real server number.
	RealServer int32 `json:"RealServer,omitempty"`
	// Enable/disable a mapping in this table.
	EntryState SlbCurCfgDrecordVirtRealMappingTableEntryState `json:"EntryState,omitempty"`
	// The virtual server AlphaNumeric Id.
	EnhVirtServer string `json:"EnhVirtServer,omitempty"`
	// The Real server AlphaNumeric Id.
	EnhRealServer string `json:"EnhRealServer,omitempty"`
}

func (p SlbCurCfgDrecordVirtRealMappingTableParams) iMABean() {}
