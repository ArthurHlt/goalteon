package beans

import (
	"fmt"
	"reflect"
)

// SlbCurCfgSdpTable The SDP table in layer7 processing engine.
// Note:This mib is not supported for VX instance of virtualization.
type SlbCurCfgSdpTable struct {
	// The SDP table index.
	SlbCurCfgSdpIndex int32
	Params            *SlbCurCfgSdpTableParams
}

func NewSlbCurCfgSdpTableList() *SlbCurCfgSdpTable {
	return &SlbCurCfgSdpTable{}
}

func NewSlbCurCfgSdpTable(
	slbCurCfgSdpIndex int32,
	params *SlbCurCfgSdpTableParams,
) *SlbCurCfgSdpTable {
	return &SlbCurCfgSdpTable{
		SlbCurCfgSdpIndex: slbCurCfgSdpIndex,
		Params:            params,
	}
}

func (c *SlbCurCfgSdpTable) Name() string {
	return "SlbCurCfgSdpTable"
}

func (c *SlbCurCfgSdpTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbCurCfgSdpTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbCurCfgSdpTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbCurCfgSdpIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbCurCfgSdpIndex)
}

type SlbCurCfgSdpTableParams struct {
	// The SDP table index.
	Index int32 `json:"Index,omitempty"`
	// The private IP address of SDP entry.
	PrivAddr string `json:"PrivAddr,omitempty"`
	// The public IP address of SDP entry.
	PublicAddr string `json:"PublicAddr,omitempty"`
}

func (p SlbCurCfgSdpTableParams) iMABean() {}
