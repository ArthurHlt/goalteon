package beans

import (
	"fmt"
	"reflect"
)

// SlbNewCfgDrecordVirtRealMappingTable The table of Domain records and virtual server and real
// server mappings.
// Note:This mib is not supported for VX instance of Virtualization.
type SlbNewCfgDrecordVirtRealMappingTable struct {
	// The number of the domain record.
	SlbNewCfgDomainRecordIndex uint64
	// The number of the virtual real server mapping.
	SlbNewCfgEntryIndex uint64
	Params              *SlbNewCfgDrecordVirtRealMappingTableParams
}

func NewSlbNewCfgDrecordVirtRealMappingTableList() *SlbNewCfgDrecordVirtRealMappingTable {
	return &SlbNewCfgDrecordVirtRealMappingTable{}
}

func NewSlbNewCfgDrecordVirtRealMappingTable(
	slbNewCfgDomainRecordIndex uint64,
	slbNewCfgEntryIndex uint64,
	params *SlbNewCfgDrecordVirtRealMappingTableParams,
) *SlbNewCfgDrecordVirtRealMappingTable {
	return &SlbNewCfgDrecordVirtRealMappingTable{
		SlbNewCfgDomainRecordIndex: slbNewCfgDomainRecordIndex,
		SlbNewCfgEntryIndex:        slbNewCfgEntryIndex,
		Params:                     params,
	}
}

func (c *SlbNewCfgDrecordVirtRealMappingTable) Name() string {
	return "SlbNewCfgDrecordVirtRealMappingTable"
}

func (c *SlbNewCfgDrecordVirtRealMappingTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbNewCfgDrecordVirtRealMappingTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbNewCfgDrecordVirtRealMappingTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbNewCfgDomainRecordIndex).IsZero() &&
		reflect.ValueOf(c.SlbNewCfgEntryIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbNewCfgDomainRecordIndex) + "/" + fmt.Sprint(c.SlbNewCfgEntryIndex)
}

type SlbNewCfgDrecordVirtRealMappingTableEntryState int32

const (
	SlbNewCfgDrecordVirtRealMappingTableEntryState_Enabled     SlbNewCfgDrecordVirtRealMappingTableEntryState = 1
	SlbNewCfgDrecordVirtRealMappingTableEntryState_Disabled    SlbNewCfgDrecordVirtRealMappingTableEntryState = 2
	SlbNewCfgDrecordVirtRealMappingTableEntryState_Unsupported SlbNewCfgDrecordVirtRealMappingTableEntryState = 2147483647
)

type SlbNewCfgDrecordVirtRealMappingTableEntryDelete int32

const (
	SlbNewCfgDrecordVirtRealMappingTableEntryDelete_Other       SlbNewCfgDrecordVirtRealMappingTableEntryDelete = 1
	SlbNewCfgDrecordVirtRealMappingTableEntryDelete_Delete      SlbNewCfgDrecordVirtRealMappingTableEntryDelete = 2
	SlbNewCfgDrecordVirtRealMappingTableEntryDelete_Unsupported SlbNewCfgDrecordVirtRealMappingTableEntryDelete = 2147483647
)

type SlbNewCfgDrecordVirtRealMappingTableParams struct {
	// The number of the domain record.
	DomainRecordIndex uint64 `json:"DomainRecordIndex,omitempty"`
	// The number of the virtual real server mapping.
	EntryIndex uint64 `json:"EntryIndex,omitempty"`
	// The virtual server number.
	Server int32 `json:"Server,omitempty"`
	// The Real server number.
	RealServer int32 `json:"RealServer,omitempty"`
	// Enable/disable a mapping in this table.
	EntryState SlbNewCfgDrecordVirtRealMappingTableEntryState `json:"EntryState,omitempty"`
	// By setting the value to delete(2), the entire row is deleted.
	EntryDelete SlbNewCfgDrecordVirtRealMappingTableEntryDelete `json:"EntryDelete,omitempty"`
	// The virtual server AlphaNumeric Id.
	EnhVirtServer string `json:"EnhVirtServer,omitempty"`
	// The Real server AlphaNumeric Id.
	EnhRealServer string `json:"EnhRealServer,omitempty"`
}

func (p SlbNewCfgDrecordVirtRealMappingTableParams) iMABean() {}
