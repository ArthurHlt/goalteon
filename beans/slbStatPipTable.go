package beans

import (
	"fmt"
	"reflect"
)

// SlbStatPipTable The Proxy IP table statistics.
// Note:This mib is not supported on VX instance of Virtualization.
type SlbStatPipTable struct {
	// The PIP address .
	SlbStatPipIndex string
	Params          *SlbStatPipTableParams
}

func NewSlbStatPipTableList() *SlbStatPipTable {
	return &SlbStatPipTable{}
}

func NewSlbStatPipTable(
	slbStatPipIndex string,
	params *SlbStatPipTableParams,
) *SlbStatPipTable {
	return &SlbStatPipTable{
		SlbStatPipIndex: slbStatPipIndex,
		Params:          params,
	}
}

func (c *SlbStatPipTable) Name() string {
	return "SlbStatPipTable"
}

func (c *SlbStatPipTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbStatPipTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbStatPipTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbStatPipIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbStatPipIndex)
}

type SlbStatPipTableParams struct {
	// The PIP address .
	Index string `json:"Index,omitempty"`
	// The PIP address Mask .
	Mask string `json:"Mask,omitempty"`
	// The total number of PIP used counter.
	Used uint32 `json:"Used,omitempty"`
	// The total number of PIP Failure counter.
	Failure uint32 `json:"Failure,omitempty"`
	// The total number of PIP Free counter.
	Free uint32 `json:"Free,omitempty"`
}

func (p SlbStatPipTableParams) iMABean() {}
