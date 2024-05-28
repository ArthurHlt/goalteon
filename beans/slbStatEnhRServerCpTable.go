package beans

import (
	"fmt"
	"reflect"
)

// SlbStatEnhRServerCpTable The real server statistics table.
// Note:This mib is not supported for VX instance of Virtualization.
type SlbStatEnhRServerCpTable struct {
	// The real server index that identifies the server.
	SlbStatEnhRServerIndexCp string
	Params                   *SlbStatEnhRServerCpTableParams
}

func NewSlbStatEnhRServerCpTableList() *SlbStatEnhRServerCpTable {
	return &SlbStatEnhRServerCpTable{}
}

func NewSlbStatEnhRServerCpTable(
	slbStatEnhRServerIndexCp string,
	params *SlbStatEnhRServerCpTableParams,
) *SlbStatEnhRServerCpTable {
	return &SlbStatEnhRServerCpTable{
		SlbStatEnhRServerIndexCp: slbStatEnhRServerIndexCp,
		Params:                   params,
	}
}

func (c *SlbStatEnhRServerCpTable) Name() string {
	return "SlbStatEnhRServerCpTable"
}

func (c *SlbStatEnhRServerCpTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbStatEnhRServerCpTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbStatEnhRServerCpTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbStatEnhRServerIndexCp).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbStatEnhRServerIndexCp)
}

type SlbStatEnhRServerCpTableParams struct {
	// The real server index that identifies the server.
	IndexCp string `json:"IndexCp,omitempty"`
	// The number of sessions that are currently handled by the real server.
	CurrSessionsCp uint32 `json:"CurrSessionsCp,omitempty"`
	// The total number of sessions that are handled by the real server.
	TotalSessionsCp uint32 `json:"TotalSessionsCp,omitempty"`
	// The total number of times that the real server is claimed down.
	FailuresCp uint32 `json:"FailuresCp,omitempty"`
	// The highest sessions that have been handled by the real server.
	HighestSessionsCp uint32 `json:"HighestSessionsCp,omitempty"`
	// The lower 32 bit value of octets received and transmitted out of the
	// real server.
	HCOctetsLow32Cp uint32 `json:"HCOctetsLow32Cp,omitempty"`
	// The higher 32 bit value of octets received and transmitted out
	// of the real server.
	HCOctetsHigh32Cp uint32 `json:"HCOctetsHigh32Cp,omitempty"`
	// The total number of octets received and transmitted by the real
	// server.
	HCOctetsCp uint64 `json:"HCOctetsCp,omitempty"`
}

func (p SlbStatEnhRServerCpTableParams) iMABean() {}
