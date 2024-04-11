package beans

import (
	"fmt"
	"reflect"
)

// SlbNewAppwallDpSrvCfgTable New DP signaling servers table.
// Note:This mib is not supported on VX instance of Virtualization.
type SlbNewAppwallDpSrvCfgTable struct {
	// Dp signaling server ID
	SlbNewAppwallDpSrvId string
	Params               *SlbNewAppwallDpSrvCfgTableParams
}

func NewSlbNewAppwallDpSrvCfgTableList() *SlbNewAppwallDpSrvCfgTable {
	return &SlbNewAppwallDpSrvCfgTable{}
}

func NewSlbNewAppwallDpSrvCfgTable(
	slbNewAppwallDpSrvId string,
	params *SlbNewAppwallDpSrvCfgTableParams,
) *SlbNewAppwallDpSrvCfgTable {
	return &SlbNewAppwallDpSrvCfgTable{
		SlbNewAppwallDpSrvId: slbNewAppwallDpSrvId,
		Params:               params,
	}
}

func (c *SlbNewAppwallDpSrvCfgTable) Name() string {
	return "SlbNewAppwallDpSrvCfgTable"
}

func (c *SlbNewAppwallDpSrvCfgTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbNewAppwallDpSrvCfgTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbNewAppwallDpSrvCfgTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbNewAppwallDpSrvId).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbNewAppwallDpSrvId)
}

type SlbNewAppwallDpSrvCfgTableDpSrvDel int32

const (
	SlbNewAppwallDpSrvCfgTableDpSrvDel_Other       SlbNewAppwallDpSrvCfgTableDpSrvDel = 1
	SlbNewAppwallDpSrvCfgTableDpSrvDel_Delete      SlbNewAppwallDpSrvCfgTableDpSrvDel = 2
	SlbNewAppwallDpSrvCfgTableDpSrvDel_Unsupported SlbNewAppwallDpSrvCfgTableDpSrvDel = 2147483647
)

type SlbNewAppwallDpSrvCfgTableParams struct {
	// Dp signaling server ID
	DpSrvId string `json:"DpSrvId,omitempty"`
	// Dp signaling server IP
	DpSrvIpAddress string `json:"DpSrvIpAddress,omitempty"`
	// Dp signaling server Port
	DpSrvPort uint64 `json:"DpSrvPort,omitempty"`
	// By setting the value to delete(2), the entire row is deleted.On GET always other(1)
	DpSrvDel SlbNewAppwallDpSrvCfgTableDpSrvDel `json:"DpSrvDel,omitempty"`
}

func (p SlbNewAppwallDpSrvCfgTableParams) iMABean() {}
