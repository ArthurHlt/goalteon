package beans

import (
	"fmt"
	"reflect"
)

// SlbNewCfgUrlBwmTable The table of URL based BWM for Virtual Services.
// Note:This mib is not supported for VX instance of Virtualization.
type SlbNewCfgUrlBwmTable struct {
	// The number of the virtual server.
	SlbNewCfgUrlBwmVirtServIndex int32
	// The service index. This has no external meaning
	SlbNewCfgUrlBwmVirtServiceIndex int32
	// The URL Path Identifier.
	SlbNewCfgUrlBwmUrlId int32
	Params               *SlbNewCfgUrlBwmTableParams
}

func NewSlbNewCfgUrlBwmTableList() *SlbNewCfgUrlBwmTable {
	return &SlbNewCfgUrlBwmTable{}
}

func NewSlbNewCfgUrlBwmTable(
	slbNewCfgUrlBwmVirtServIndex int32,
	slbNewCfgUrlBwmVirtServiceIndex int32,
	slbNewCfgUrlBwmUrlId int32,
	params *SlbNewCfgUrlBwmTableParams,
) *SlbNewCfgUrlBwmTable {
	return &SlbNewCfgUrlBwmTable{
		SlbNewCfgUrlBwmVirtServIndex:    slbNewCfgUrlBwmVirtServIndex,
		SlbNewCfgUrlBwmVirtServiceIndex: slbNewCfgUrlBwmVirtServiceIndex,
		SlbNewCfgUrlBwmUrlId:            slbNewCfgUrlBwmUrlId,
		Params:                          params,
	}
}

func (c *SlbNewCfgUrlBwmTable) Name() string {
	return "SlbNewCfgUrlBwmTable"
}

func (c *SlbNewCfgUrlBwmTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbNewCfgUrlBwmTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbNewCfgUrlBwmTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbNewCfgUrlBwmVirtServIndex).IsZero() &&
		reflect.ValueOf(c.SlbNewCfgUrlBwmVirtServiceIndex).IsZero() &&
		reflect.ValueOf(c.SlbNewCfgUrlBwmUrlId).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbNewCfgUrlBwmVirtServIndex) + "/" + fmt.Sprint(c.SlbNewCfgUrlBwmVirtServiceIndex) + "/" + fmt.Sprint(c.SlbNewCfgUrlBwmUrlId)
}

type SlbNewCfgUrlBwmTableDelete int32

const (
	SlbNewCfgUrlBwmTableDelete_Other       SlbNewCfgUrlBwmTableDelete = 1
	SlbNewCfgUrlBwmTableDelete_Delete      SlbNewCfgUrlBwmTableDelete = 2
	SlbNewCfgUrlBwmTableDelete_Unsupported SlbNewCfgUrlBwmTableDelete = 2147483647
)

type SlbNewCfgUrlBwmTableParams struct {
	// The number of the virtual server.
	VirtServIndex int32 `json:"VirtServIndex,omitempty"`
	// The service index. This has no external meaning
	VirtServiceIndex int32 `json:"VirtServiceIndex,omitempty"`
	// The URL Path Identifier.
	UrlId int32 `json:"UrlId,omitempty"`
	// The BW contract number.
	Contract int32 `json:"Contract,omitempty"`
	// When set to the value of 2 (delete), the entire row is deleted.
	// When read, other(1) is returned. Setting the value to anything
	// other than 2(delete) has no effect on the state of the row.
	Delete SlbNewCfgUrlBwmTableDelete `json:"Delete,omitempty"`
}

func (p SlbNewCfgUrlBwmTableParams) iMABean() {}
