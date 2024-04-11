package beans

import (
	"fmt"
	"reflect"
)

// SlbNewCfgEnhUrlBwmTable The table of URL based BWM for Virtual Services.
// Note:This mib is not supported for VX instance of Virtualization.
type SlbNewCfgEnhUrlBwmTable struct {
	// The number of the virtual server.
	SlbNewCfgEnhUrlBwmVirtServIndex string
	// The service index. This has no external meaning
	SlbNewCfgEnhUrlBwmVirtServiceIndex int32
	// The URL Path Identifier.
	SlbNewCfgEnhUrlBwmUrlId int32
	Params                  *SlbNewCfgEnhUrlBwmTableParams
}

func NewSlbNewCfgEnhUrlBwmTableList() *SlbNewCfgEnhUrlBwmTable {
	return &SlbNewCfgEnhUrlBwmTable{}
}

func NewSlbNewCfgEnhUrlBwmTable(
	slbNewCfgEnhUrlBwmVirtServIndex string,
	slbNewCfgEnhUrlBwmVirtServiceIndex int32,
	slbNewCfgEnhUrlBwmUrlId int32,
	params *SlbNewCfgEnhUrlBwmTableParams,
) *SlbNewCfgEnhUrlBwmTable {
	return &SlbNewCfgEnhUrlBwmTable{
		SlbNewCfgEnhUrlBwmVirtServIndex:    slbNewCfgEnhUrlBwmVirtServIndex,
		SlbNewCfgEnhUrlBwmVirtServiceIndex: slbNewCfgEnhUrlBwmVirtServiceIndex,
		SlbNewCfgEnhUrlBwmUrlId:            slbNewCfgEnhUrlBwmUrlId,
		Params:                             params,
	}
}

func (c *SlbNewCfgEnhUrlBwmTable) Name() string {
	return "SlbNewCfgEnhUrlBwmTable"
}

func (c *SlbNewCfgEnhUrlBwmTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbNewCfgEnhUrlBwmTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbNewCfgEnhUrlBwmTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbNewCfgEnhUrlBwmVirtServIndex).IsZero() &&
		reflect.ValueOf(c.SlbNewCfgEnhUrlBwmVirtServiceIndex).IsZero() &&
		reflect.ValueOf(c.SlbNewCfgEnhUrlBwmUrlId).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbNewCfgEnhUrlBwmVirtServIndex) + "/" + fmt.Sprint(c.SlbNewCfgEnhUrlBwmVirtServiceIndex) + "/" + fmt.Sprint(c.SlbNewCfgEnhUrlBwmUrlId)
}

type SlbNewCfgEnhUrlBwmTableDelete int32

const (
	SlbNewCfgEnhUrlBwmTableDelete_Other       SlbNewCfgEnhUrlBwmTableDelete = 1
	SlbNewCfgEnhUrlBwmTableDelete_Delete      SlbNewCfgEnhUrlBwmTableDelete = 2
	SlbNewCfgEnhUrlBwmTableDelete_Unsupported SlbNewCfgEnhUrlBwmTableDelete = 2147483647
)

type SlbNewCfgEnhUrlBwmTableParams struct {
	// The number of the virtual server.
	VirtServIndex string `json:"VirtServIndex,omitempty"`
	// The service index. This has no external meaning
	VirtServiceIndex int32 `json:"VirtServiceIndex,omitempty"`
	// The URL Path Identifier.
	UrlId int32 `json:"UrlId,omitempty"`
	// The BW contract number.
	Contract int32 `json:"Contract,omitempty"`
	// When set to the value of 2 (delete), the entire row is deleted.
	// When read, other(1) is returned. Setting the value to anything
	// other than 2(delete) has no effect on the state of the row.
	Delete SlbNewCfgEnhUrlBwmTableDelete `json:"Delete,omitempty"`
}

func (p SlbNewCfgEnhUrlBwmTableParams) iMABean() {}
