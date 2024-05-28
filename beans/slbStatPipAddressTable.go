package beans

import (
	"fmt"
	"reflect"
)

// SlbStatPipAddressTable The Proxy IP table statistics, configured as service.
// Note:This mib is not supported on VX instance of Virtualization.
type SlbStatPipAddressTable struct {
	// The PIP address for address mode.
	SlbStatPipAddressPip string
	Params               *SlbStatPipAddressTableParams
}

func NewSlbStatPipAddressTableList() *SlbStatPipAddressTable {
	return &SlbStatPipAddressTable{}
}

func NewSlbStatPipAddressTable(
	slbStatPipAddressPip string,
	params *SlbStatPipAddressTableParams,
) *SlbStatPipAddressTable {
	return &SlbStatPipAddressTable{
		SlbStatPipAddressPip: slbStatPipAddressPip,
		Params:               params,
	}
}

func (c *SlbStatPipAddressTable) Name() string {
	return "SlbStatPipAddressTable"
}

func (c *SlbStatPipAddressTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbStatPipAddressTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbStatPipAddressTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbStatPipAddressPip).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbStatPipAddressPip)
}

type SlbStatPipAddressTableParams struct {
	// The PIP address for address mode.
	Pip string `json:"Pip,omitempty"`
	// The PIP address Mask for address mode.
	Subnet uint32 `json:"Subnet,omitempty"`
	// The total number of PIP used counter for address mode.
	Used uint32 `json:"Used,omitempty"`
	// The total number of PIP Failure counter for address mode.
	Failure uint32 `json:"Failure,omitempty"`
}

func (p SlbStatPipAddressTableParams) iMABean() {}
