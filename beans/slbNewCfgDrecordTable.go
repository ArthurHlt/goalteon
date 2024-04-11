package beans

import (
	"fmt"
	"reflect"
)

// SlbNewCfgDrecordTable The table of Domain records for Inbound link load balancing
// in the new configuration block.
// Note:This mib is not supported for VX instance of Virtualization.
type SlbNewCfgDrecordTable struct {
	// The index of the Domain record table.
	SlbNewCfgDrecordIndex uint64
	Params                *SlbNewCfgDrecordTableParams
}

func NewSlbNewCfgDrecordTableList() *SlbNewCfgDrecordTable {
	return &SlbNewCfgDrecordTable{}
}

func NewSlbNewCfgDrecordTable(
	slbNewCfgDrecordIndex uint64,
	params *SlbNewCfgDrecordTableParams,
) *SlbNewCfgDrecordTable {
	return &SlbNewCfgDrecordTable{
		SlbNewCfgDrecordIndex: slbNewCfgDrecordIndex,
		Params:                params,
	}
}

func (c *SlbNewCfgDrecordTable) Name() string {
	return "SlbNewCfgDrecordTable"
}

func (c *SlbNewCfgDrecordTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbNewCfgDrecordTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbNewCfgDrecordTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbNewCfgDrecordIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbNewCfgDrecordIndex)
}

type SlbNewCfgDrecordTableDomainRecordState int32

const (
	SlbNewCfgDrecordTableDomainRecordState_Enabled     SlbNewCfgDrecordTableDomainRecordState = 1
	SlbNewCfgDrecordTableDomainRecordState_Disabled    SlbNewCfgDrecordTableDomainRecordState = 2
	SlbNewCfgDrecordTableDomainRecordState_Unsupported SlbNewCfgDrecordTableDomainRecordState = 2147483647
)

type SlbNewCfgDrecordTableDelete int32

const (
	SlbNewCfgDrecordTableDelete_Other       SlbNewCfgDrecordTableDelete = 1
	SlbNewCfgDrecordTableDelete_Delete      SlbNewCfgDrecordTableDelete = 2
	SlbNewCfgDrecordTableDelete_Unsupported SlbNewCfgDrecordTableDelete = 2147483647
)

type SlbNewCfgDrecordTableParams struct {
	// The index of the Domain record table.
	Index uint64 `json:"Index,omitempty"`
	// Enable/disable a  domain record.
	DomainRecordState SlbNewCfgDrecordTableDomainRecordState `json:"DomainRecordState,omitempty"`
	// The name of the domain record.
	DomainRecordName string `json:"DomainRecordName,omitempty"`
	// By setting the value to delete(2), the entire row is deleted.
	Delete SlbNewCfgDrecordTableDelete `json:"Delete,omitempty"`
}

func (p SlbNewCfgDrecordTableParams) iMABean() {}
