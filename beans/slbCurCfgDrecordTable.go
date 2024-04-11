package beans

import (
	"fmt"
	"reflect"
)

// SlbCurCfgDrecordTable The table of Domain records for Inbound link load balancing
// in the current configuration block.
// Note:This mib is not supported for VX instance of Virtualization.
type SlbCurCfgDrecordTable struct {
	// The index of the Domain record table.
	SlbCurCfgDrecordIndex uint64
	Params                *SlbCurCfgDrecordTableParams
}

func NewSlbCurCfgDrecordTableList() *SlbCurCfgDrecordTable {
	return &SlbCurCfgDrecordTable{}
}

func NewSlbCurCfgDrecordTable(
	slbCurCfgDrecordIndex uint64,
	params *SlbCurCfgDrecordTableParams,
) *SlbCurCfgDrecordTable {
	return &SlbCurCfgDrecordTable{
		SlbCurCfgDrecordIndex: slbCurCfgDrecordIndex,
		Params:                params,
	}
}

func (c *SlbCurCfgDrecordTable) Name() string {
	return "SlbCurCfgDrecordTable"
}

func (c *SlbCurCfgDrecordTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbCurCfgDrecordTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbCurCfgDrecordTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbCurCfgDrecordIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbCurCfgDrecordIndex)
}

type SlbCurCfgDrecordTableDomainRecordState int32

const (
	SlbCurCfgDrecordTableDomainRecordState_Enabled     SlbCurCfgDrecordTableDomainRecordState = 1
	SlbCurCfgDrecordTableDomainRecordState_Disabled    SlbCurCfgDrecordTableDomainRecordState = 2
	SlbCurCfgDrecordTableDomainRecordState_Unsupported SlbCurCfgDrecordTableDomainRecordState = 2147483647
)

type SlbCurCfgDrecordTableParams struct {
	// The index of the Domain record table.
	Index uint64 `json:"Index,omitempty"`
	// Enable/disable a  domain record.
	DomainRecordState SlbCurCfgDrecordTableDomainRecordState `json:"DomainRecordState,omitempty"`
	// The name of the domain record.
	DomainRecordName string `json:"DomainRecordName,omitempty"`
}

func (p SlbCurCfgDrecordTableParams) iMABean() {}
