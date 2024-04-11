package beans

import (
	"fmt"
	"reflect"
)

// SlbCurCfgEnhUrlBwmTable The table of URL based BWM for Virtual Services.
// Note:This mib is not supported for VX instance of Virtualization.
type SlbCurCfgEnhUrlBwmTable struct {
	// The number of the virtual server.
	SlbCurCfgEnhUrlBwmVirtServIndex string
	// The service index. This has no external meaning
	SlbCurCfgEnhUrlBwmVirtServiceIndex int32
	// The URL Path Identifier.
	SlbCurCfgEnhUrlBwmUrlId int32
	Params                  *SlbCurCfgEnhUrlBwmTableParams
}

func NewSlbCurCfgEnhUrlBwmTableList() *SlbCurCfgEnhUrlBwmTable {
	return &SlbCurCfgEnhUrlBwmTable{}
}

func NewSlbCurCfgEnhUrlBwmTable(
	slbCurCfgEnhUrlBwmVirtServIndex string,
	slbCurCfgEnhUrlBwmVirtServiceIndex int32,
	slbCurCfgEnhUrlBwmUrlId int32,
	params *SlbCurCfgEnhUrlBwmTableParams,
) *SlbCurCfgEnhUrlBwmTable {
	return &SlbCurCfgEnhUrlBwmTable{
		SlbCurCfgEnhUrlBwmVirtServIndex:    slbCurCfgEnhUrlBwmVirtServIndex,
		SlbCurCfgEnhUrlBwmVirtServiceIndex: slbCurCfgEnhUrlBwmVirtServiceIndex,
		SlbCurCfgEnhUrlBwmUrlId:            slbCurCfgEnhUrlBwmUrlId,
		Params:                             params,
	}
}

func (c *SlbCurCfgEnhUrlBwmTable) Name() string {
	return "SlbCurCfgEnhUrlBwmTable"
}

func (c *SlbCurCfgEnhUrlBwmTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbCurCfgEnhUrlBwmTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbCurCfgEnhUrlBwmTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbCurCfgEnhUrlBwmVirtServIndex).IsZero() &&
		reflect.ValueOf(c.SlbCurCfgEnhUrlBwmVirtServiceIndex).IsZero() &&
		reflect.ValueOf(c.SlbCurCfgEnhUrlBwmUrlId).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbCurCfgEnhUrlBwmVirtServIndex) + "/" + fmt.Sprint(c.SlbCurCfgEnhUrlBwmVirtServiceIndex) + "/" + fmt.Sprint(c.SlbCurCfgEnhUrlBwmUrlId)
}

type SlbCurCfgEnhUrlBwmTableParams struct {
	// The number of the virtual server.
	VirtServIndex string `json:"VirtServIndex,omitempty"`
	// The service index. This has no external meaning
	VirtServiceIndex int32 `json:"VirtServiceIndex,omitempty"`
	// The URL Path Identifier.
	UrlId int32 `json:"UrlId,omitempty"`
	// The BW contract number.
	Contract int32 `json:"Contract,omitempty"`
}

func (p SlbCurCfgEnhUrlBwmTableParams) iMABean() {}
