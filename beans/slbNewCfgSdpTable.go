package beans

import (
	"fmt"
	"reflect"
)

// SlbNewCfgSdpTable The SDP table in layer7 processing engine.
// Note:This mib is not supported for VX instance of virtualization.
type SlbNewCfgSdpTable struct {
	// The SDP table index.Maximum entries you can configure is 16
	SlbNewCfgSdpIndex int32
	Params            *SlbNewCfgSdpTableParams
}

func NewSlbNewCfgSdpTableList() *SlbNewCfgSdpTable {
	return &SlbNewCfgSdpTable{}
}

func NewSlbNewCfgSdpTable(
	slbNewCfgSdpIndex int32,
	params *SlbNewCfgSdpTableParams,
) *SlbNewCfgSdpTable {
	return &SlbNewCfgSdpTable{
		SlbNewCfgSdpIndex: slbNewCfgSdpIndex,
		Params:            params,
	}
}

func (c *SlbNewCfgSdpTable) Name() string {
	return "SlbNewCfgSdpTable"
}

func (c *SlbNewCfgSdpTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbNewCfgSdpTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbNewCfgSdpTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbNewCfgSdpIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbNewCfgSdpIndex)
}

type SlbNewCfgSdpTableDelete int32

const (
	SlbNewCfgSdpTableDelete_Other       SlbNewCfgSdpTableDelete = 1
	SlbNewCfgSdpTableDelete_Delete      SlbNewCfgSdpTableDelete = 2
	SlbNewCfgSdpTableDelete_Unsupported SlbNewCfgSdpTableDelete = 2147483647
)

type SlbNewCfgSdpTableParams struct {
	// The SDP table index.Maximum entries you can configure is 16
	Index int32 `json:"Index,omitempty"`
	// The private IP address of SDP entry.
	PrivAddr string `json:"PrivAddr,omitempty"`
	// The public IP address of SDP entry.
	PublicAddr string `json:"PublicAddr,omitempty"`
	// When set to the value of 2 (delete), the entire row is
	// deleted. When read, other(1) is returned. Setting the value
	// to anything other than 2(delete) has no effect on the state of the row.
	Delete SlbNewCfgSdpTableDelete `json:"Delete,omitempty"`
}

func (p SlbNewCfgSdpTableParams) iMABean() {}
