package beans

import (
	"fmt"
	"reflect"
)

// SlbStatSapAslrTable The sap aslrr statistics table.  This table shows the statistics
// of sap aslr for a particular vip.
// Note:This mib is not supported on VX instance of Virtualization.
type SlbStatSapAslrTable struct {
	// The Sap Aslr table index.
	SlbStatSapAslrIndex int32
	Params              *SlbStatSapAslrTableParams
}

func NewSlbStatSapAslrTableList() *SlbStatSapAslrTable {
	return &SlbStatSapAslrTable{}
}

func NewSlbStatSapAslrTable(
	slbStatSapAslrIndex int32,
	params *SlbStatSapAslrTableParams,
) *SlbStatSapAslrTable {
	return &SlbStatSapAslrTable{
		SlbStatSapAslrIndex: slbStatSapAslrIndex,
		Params:              params,
	}
}

func (c *SlbStatSapAslrTable) Name() string {
	return "SlbStatSapAslrTable"
}

func (c *SlbStatSapAslrTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbStatSapAslrTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbStatSapAslrTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbStatSapAslrIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbStatSapAslrIndex)
}

type SlbStatSapAslrTableParams struct {
	// The Sap Aslr table index.
	Index int32 `json:"Index,omitempty"`
	// The number of Sap Aslr server lookups.
	Sampling uint32 `json:"Sampling,omitempty"`
	// The number of Sap Aslr server lookup failures.
	Failure uint32 `json:"Failure,omitempty"`
	// The number of Alteon automated configuration updates.
	DeviceUpdates uint32 `json:"DeviceUpdates,omitempty"`
	// The number of Alteon automated configuration failure.
	DeviceFailure uint32 `json:"DeviceFailure,omitempty"`
	// The Sap Aslr last activity time stamp.
	LastAct string `json:"LastAct,omitempty"`
}

func (p SlbStatSapAslrTableParams) iMABean() {}
