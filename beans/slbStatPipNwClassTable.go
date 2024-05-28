package beans

import (
	"fmt"
	"reflect"
)

// SlbStatPipNwClassTable The Proxy IP NwClass table statistics.
// Note:This mib is not supported on VX instance of Virtualization.
type SlbStatPipNwClassTable struct {
	// The PIP Network class id .
	SlbStatPipNwClassId string
	Params              *SlbStatPipNwClassTableParams
}

func NewSlbStatPipNwClassTableList() *SlbStatPipNwClassTable {
	return &SlbStatPipNwClassTable{}
}

func NewSlbStatPipNwClassTable(
	slbStatPipNwClassId string,
	params *SlbStatPipNwClassTableParams,
) *SlbStatPipNwClassTable {
	return &SlbStatPipNwClassTable{
		SlbStatPipNwClassId: slbStatPipNwClassId,
		Params:              params,
	}
}

func (c *SlbStatPipNwClassTable) Name() string {
	return "SlbStatPipNwClassTable"
}

func (c *SlbStatPipNwClassTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbStatPipNwClassTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbStatPipNwClassTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbStatPipNwClassId).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbStatPipNwClassId)
}

type SlbStatPipNwClassTableParams struct {
	// The PIP Network class id .
	Id string `json:"Id,omitempty"`
	// The PIP Network class name .
	Name string `json:"Name,omitempty"`
	// The total number of PIP NwClass used counter.
	Used uint32 `json:"Used,omitempty"`
	// The total number of PIP NwClass Failure counter.
	Failure uint32 `json:"Failure,omitempty"`
}

func (p SlbStatPipNwClassTableParams) iMABean() {}
