package beans

import (
	"fmt"
	"reflect"
)

// SlbStatPip6Table The Proxy IP table statistics .
type SlbStatPip6Table struct {
	// The PIP6 address idx .
	SlbStatPip6Index string
	Params           *SlbStatPip6TableParams
}

func NewSlbStatPip6TableList() *SlbStatPip6Table {
	return &SlbStatPip6Table{}
}

func NewSlbStatPip6Table(
	slbStatPip6Index string,
	params *SlbStatPip6TableParams,
) *SlbStatPip6Table {
	return &SlbStatPip6Table{
		SlbStatPip6Index: slbStatPip6Index,
		Params:           params,
	}
}

func (c *SlbStatPip6Table) Name() string {
	return "SlbStatPip6Table"
}

func (c *SlbStatPip6Table) GetParams() BeanType {
	return c.Params
}

func (c *SlbStatPip6Table) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbStatPip6Table) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbStatPip6Index).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbStatPip6Index)
}

type SlbStatPip6TableParams struct {
	// The PIP6 address idx .
	Index string `json:"Index,omitempty"`
	// The PIP6 prefix length .
	PrefixLen int32 `json:"PrefixLen,omitempty"`
	// The total number of PIP6 used counter.
	Used uint32 `json:"Used,omitempty"`
	// The total number of PIP6 Failure counter.
	Failure uint32 `json:"Failure,omitempty"`
	// The total number of PIP6 Free counter.
	Free uint32 `json:"Free,omitempty"`
}

func (p SlbStatPip6TableParams) iMABean() {}
