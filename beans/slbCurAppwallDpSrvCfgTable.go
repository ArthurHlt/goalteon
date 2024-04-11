package beans

import (
	"fmt"
	"reflect"
)

// SlbCurAppwallDpSrvCfgTable Cur DP signaling servers table.
// Note:This mib is not supported on VX instance of Virtualization.
type SlbCurAppwallDpSrvCfgTable struct {
	// Dp signaling server ID
	SlbCurAppwallDpSrvId string
	Params               *SlbCurAppwallDpSrvCfgTableParams
}

func NewSlbCurAppwallDpSrvCfgTableList() *SlbCurAppwallDpSrvCfgTable {
	return &SlbCurAppwallDpSrvCfgTable{}
}

func NewSlbCurAppwallDpSrvCfgTable(
	slbCurAppwallDpSrvId string,
	params *SlbCurAppwallDpSrvCfgTableParams,
) *SlbCurAppwallDpSrvCfgTable {
	return &SlbCurAppwallDpSrvCfgTable{
		SlbCurAppwallDpSrvId: slbCurAppwallDpSrvId,
		Params:               params,
	}
}

func (c *SlbCurAppwallDpSrvCfgTable) Name() string {
	return "SlbCurAppwallDpSrvCfgTable"
}

func (c *SlbCurAppwallDpSrvCfgTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbCurAppwallDpSrvCfgTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbCurAppwallDpSrvCfgTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbCurAppwallDpSrvId).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbCurAppwallDpSrvId)
}

type SlbCurAppwallDpSrvCfgTableDpSrvDel int32

const (
	SlbCurAppwallDpSrvCfgTableDpSrvDel_Other       SlbCurAppwallDpSrvCfgTableDpSrvDel = 1
	SlbCurAppwallDpSrvCfgTableDpSrvDel_Delete      SlbCurAppwallDpSrvCfgTableDpSrvDel = 2
	SlbCurAppwallDpSrvCfgTableDpSrvDel_Unsupported SlbCurAppwallDpSrvCfgTableDpSrvDel = 2147483647
)

type SlbCurAppwallDpSrvCfgTableParams struct {
	// Dp signaling server ID
	DpSrvId string `json:"DpSrvId,omitempty"`
	// Dp signaling server IP
	DpSrvIpAddress string `json:"DpSrvIpAddress,omitempty"`
	// Dp signaling server Port
	DpSrvPort uint64 `json:"DpSrvPort,omitempty"`
	// By setting the value to delete(2), the entire row is deleted.On GET always other(1)
	DpSrvDel SlbCurAppwallDpSrvCfgTableDpSrvDel `json:"DpSrvDel,omitempty"`
}

func (p SlbCurAppwallDpSrvCfgTableParams) iMABean() {}
