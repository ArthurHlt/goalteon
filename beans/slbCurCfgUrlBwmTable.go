package beans

import (
	"fmt"
	"reflect"
)

// SlbCurCfgUrlBwmTable The table of URL based BWM for Virtual Services.
// Note:This mib is not supported for VX instance of Virtualization.
type SlbCurCfgUrlBwmTable struct {
	// The number of the virtual server.
	SlbCurCfgUrlBwmVirtServIndex int32
	// The service index. This has no external meaning
	SlbCurCfgUrlBwmVirtServiceIndex int32
	// The URL Path Identifier.
	SlbCurCfgUrlBwmUrlId int32
	Params               *SlbCurCfgUrlBwmTableParams
}

func NewSlbCurCfgUrlBwmTableList() *SlbCurCfgUrlBwmTable {
	return &SlbCurCfgUrlBwmTable{}
}

func NewSlbCurCfgUrlBwmTable(
	slbCurCfgUrlBwmVirtServIndex int32,
	slbCurCfgUrlBwmVirtServiceIndex int32,
	slbCurCfgUrlBwmUrlId int32,
	params *SlbCurCfgUrlBwmTableParams,
) *SlbCurCfgUrlBwmTable {
	return &SlbCurCfgUrlBwmTable{
		SlbCurCfgUrlBwmVirtServIndex:    slbCurCfgUrlBwmVirtServIndex,
		SlbCurCfgUrlBwmVirtServiceIndex: slbCurCfgUrlBwmVirtServiceIndex,
		SlbCurCfgUrlBwmUrlId:            slbCurCfgUrlBwmUrlId,
		Params:                          params,
	}
}

func (c *SlbCurCfgUrlBwmTable) Name() string {
	return "SlbCurCfgUrlBwmTable"
}

func (c *SlbCurCfgUrlBwmTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbCurCfgUrlBwmTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbCurCfgUrlBwmTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbCurCfgUrlBwmVirtServIndex).IsZero() &&
		reflect.ValueOf(c.SlbCurCfgUrlBwmVirtServiceIndex).IsZero() &&
		reflect.ValueOf(c.SlbCurCfgUrlBwmUrlId).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbCurCfgUrlBwmVirtServIndex) + "/" + fmt.Sprint(c.SlbCurCfgUrlBwmVirtServiceIndex) + "/" + fmt.Sprint(c.SlbCurCfgUrlBwmUrlId)
}

type SlbCurCfgUrlBwmTableParams struct {
	// The number of the virtual server.
	VirtServIndex int32 `json:"VirtServIndex,omitempty"`
	// The service index. This has no external meaning
	VirtServiceIndex int32 `json:"VirtServiceIndex,omitempty"`
	// The URL Path Identifier.
	UrlId int32 `json:"UrlId,omitempty"`
	// The BW contract number.
	Contract int32 `json:"Contract,omitempty"`
}

func (p SlbCurCfgUrlBwmTableParams) iMABean() {}
