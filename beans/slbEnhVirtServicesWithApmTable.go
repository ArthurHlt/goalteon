package beans

import (
	"fmt"
	"reflect"
)

// SlbEnhVirtServicesWithApmTable Table of virtual services with APM enable on the service.
// Note:This mib is not supported for VX instance of Virtualization.
type SlbEnhVirtServicesWithApmTable struct {
	// The virtual server id.
	SlbEnhVirtServerWithApmIndex string
	// The virtual service index.
	SlbEnhVirtServiceWithApmIndex int32
	Params                        *SlbEnhVirtServicesWithApmTableParams
}

func NewSlbEnhVirtServicesWithApmTableList() *SlbEnhVirtServicesWithApmTable {
	return &SlbEnhVirtServicesWithApmTable{}
}

func NewSlbEnhVirtServicesWithApmTable(
	slbEnhVirtServerWithApmIndex string,
	slbEnhVirtServiceWithApmIndex int32,
	params *SlbEnhVirtServicesWithApmTableParams,
) *SlbEnhVirtServicesWithApmTable {
	return &SlbEnhVirtServicesWithApmTable{
		SlbEnhVirtServerWithApmIndex:  slbEnhVirtServerWithApmIndex,
		SlbEnhVirtServiceWithApmIndex: slbEnhVirtServiceWithApmIndex,
		Params:                        params,
	}
}

func (c *SlbEnhVirtServicesWithApmTable) Name() string {
	return "SlbEnhVirtServicesWithApmTable"
}

func (c *SlbEnhVirtServicesWithApmTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbEnhVirtServicesWithApmTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbEnhVirtServicesWithApmTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbEnhVirtServerWithApmIndex).IsZero() &&
		reflect.ValueOf(c.SlbEnhVirtServiceWithApmIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbEnhVirtServerWithApmIndex) + "/" + fmt.Sprint(c.SlbEnhVirtServiceWithApmIndex)
}

type SlbEnhVirtServicesWithApmTableParams struct {
	// The virtual server id.
	ServerWithApmIndex string `json:"ServerWithApmIndex,omitempty"`
	// The virtual service index.
	WithApmIndex int32 `json:"WithApmIndex,omitempty"`
}

func (p SlbEnhVirtServicesWithApmTableParams) iMABean() {}
