package beans

import (
	"fmt"
	"reflect"
)

// SlbEnhVirtSpecificServicesInfoTable The table of specific virtual service information.
// Note:This mib is not supported for VX instance of Virtualization.
type SlbEnhVirtSpecificServicesInfoTable struct {
	// The virtual server index.
	SlbEnhVirtSpecificServicesInfoVirtServIndex string
	// The virtual service index.
	SlbEnhVirtSpecificServicesInfoSvcIndex int32
	Params                                 *SlbEnhVirtSpecificServicesInfoTableParams
}

func NewSlbEnhVirtSpecificServicesInfoTableList() *SlbEnhVirtSpecificServicesInfoTable {
	return &SlbEnhVirtSpecificServicesInfoTable{}
}

func NewSlbEnhVirtSpecificServicesInfoTable(
	slbEnhVirtSpecificServicesInfoVirtServIndex string,
	slbEnhVirtSpecificServicesInfoSvcIndex int32,
	params *SlbEnhVirtSpecificServicesInfoTableParams,
) *SlbEnhVirtSpecificServicesInfoTable {
	return &SlbEnhVirtSpecificServicesInfoTable{
		SlbEnhVirtSpecificServicesInfoVirtServIndex: slbEnhVirtSpecificServicesInfoVirtServIndex,
		SlbEnhVirtSpecificServicesInfoSvcIndex:      slbEnhVirtSpecificServicesInfoSvcIndex,
		Params:                                      params,
	}
}

func (c *SlbEnhVirtSpecificServicesInfoTable) Name() string {
	return "SlbEnhVirtSpecificServicesInfoTable"
}

func (c *SlbEnhVirtSpecificServicesInfoTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbEnhVirtSpecificServicesInfoTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbEnhVirtSpecificServicesInfoTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbEnhVirtSpecificServicesInfoVirtServIndex).IsZero() &&
		reflect.ValueOf(c.SlbEnhVirtSpecificServicesInfoSvcIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbEnhVirtSpecificServicesInfoVirtServIndex) + "/" + fmt.Sprint(c.SlbEnhVirtSpecificServicesInfoSvcIndex)
}

type SlbEnhVirtSpecificServicesInfoTableStatus int32

const (
	SlbEnhVirtSpecificServicesInfoTableStatus_Up          SlbEnhVirtSpecificServicesInfoTableStatus = 0
	SlbEnhVirtSpecificServicesInfoTableStatus_Down        SlbEnhVirtSpecificServicesInfoTableStatus = 1
	SlbEnhVirtSpecificServicesInfoTableStatus_Admindown   SlbEnhVirtSpecificServicesInfoTableStatus = 2
	SlbEnhVirtSpecificServicesInfoTableStatus_Warning     SlbEnhVirtSpecificServicesInfoTableStatus = 3
	SlbEnhVirtSpecificServicesInfoTableStatus_Shudown     SlbEnhVirtSpecificServicesInfoTableStatus = 4
	SlbEnhVirtSpecificServicesInfoTableStatus_Unsupported SlbEnhVirtSpecificServicesInfoTableStatus = 2147483647
)

type SlbEnhVirtSpecificServicesInfoTableParams struct {
	// The virtual server index.
	VirtServIndex string `json:"VirtServIndex,omitempty"`
	// The virtual service index.
	SvcIndex int32 `json:"SvcIndex,omitempty"`
	// The status of the specific virtual service.
	Status SlbEnhVirtSpecificServicesInfoTableStatus `json:"Status,omitempty"`
}

func (p SlbEnhVirtSpecificServicesInfoTableParams) iMABean() {}
