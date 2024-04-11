package beans

import (
	"fmt"
	"reflect"
)

// SlbNewSslCfgSSLInspectFlowTable The table of SSL Inspection Flow.
// Note:This mib is not supported for VX instance of Virtualization.
type SlbNewSslCfgSSLInspectFlowTable struct {
	// The number of the flow.
	SlbNewSslCfgSSLInspectFlowIndex int32
	// The The Secarity group index.
	SlbNewSslCfgSSLInspectFlowSeGroupIndex int32
	Params                                 *SlbNewSslCfgSSLInspectFlowTableParams
}

func NewSlbNewSslCfgSSLInspectFlowTableList() *SlbNewSslCfgSSLInspectFlowTable {
	return &SlbNewSslCfgSSLInspectFlowTable{}
}

func NewSlbNewSslCfgSSLInspectFlowTable(
	slbNewSslCfgSSLInspectFlowIndex int32,
	slbNewSslCfgSSLInspectFlowSeGroupIndex int32,
	params *SlbNewSslCfgSSLInspectFlowTableParams,
) *SlbNewSslCfgSSLInspectFlowTable {
	return &SlbNewSslCfgSSLInspectFlowTable{
		SlbNewSslCfgSSLInspectFlowIndex:        slbNewSslCfgSSLInspectFlowIndex,
		SlbNewSslCfgSSLInspectFlowSeGroupIndex: slbNewSslCfgSSLInspectFlowSeGroupIndex,
		Params:                                 params,
	}
}

func (c *SlbNewSslCfgSSLInspectFlowTable) Name() string {
	return "SlbNewSslCfgSSLInspectFlowTable"
}

func (c *SlbNewSslCfgSSLInspectFlowTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbNewSslCfgSSLInspectFlowTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbNewSslCfgSSLInspectFlowTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbNewSslCfgSSLInspectFlowIndex).IsZero() &&
		reflect.ValueOf(c.SlbNewSslCfgSSLInspectFlowSeGroupIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbNewSslCfgSSLInspectFlowIndex) + "/" + fmt.Sprint(c.SlbNewSslCfgSSLInspectFlowSeGroupIndex)
}

type SlbNewSslCfgSSLInspectFlowTableSeGroupState int32

const (
	SlbNewSslCfgSSLInspectFlowTableSeGroupState_Enabled     SlbNewSslCfgSSLInspectFlowTableSeGroupState = 2
	SlbNewSslCfgSSLInspectFlowTableSeGroupState_Disabled    SlbNewSslCfgSSLInspectFlowTableSeGroupState = 3
	SlbNewSslCfgSSLInspectFlowTableSeGroupState_Unsupported SlbNewSslCfgSSLInspectFlowTableSeGroupState = 2147483647
)

type SlbNewSslCfgSSLInspectFlowTableSeGroupFallback int32

const (
	SlbNewSslCfgSSLInspectFlowTableSeGroupFallback_None        SlbNewSslCfgSSLInspectFlowTableSeGroupFallback = 1
	SlbNewSslCfgSSLInspectFlowTableSeGroupFallback_Continue    SlbNewSslCfgSSLInspectFlowTableSeGroupFallback = 2
	SlbNewSslCfgSSLInspectFlowTableSeGroupFallback_Drop        SlbNewSslCfgSSLInspectFlowTableSeGroupFallback = 3
	SlbNewSslCfgSSLInspectFlowTableSeGroupFallback_Unsupported SlbNewSslCfgSSLInspectFlowTableSeGroupFallback = 2147483647
)

type SlbNewSslCfgSSLInspectFlowTableDelete int32

const (
	SlbNewSslCfgSSLInspectFlowTableDelete_Other       SlbNewSslCfgSSLInspectFlowTableDelete = 1
	SlbNewSslCfgSSLInspectFlowTableDelete_Delete      SlbNewSslCfgSSLInspectFlowTableDelete = 2
	SlbNewSslCfgSSLInspectFlowTableDelete_Unsupported SlbNewSslCfgSSLInspectFlowTableDelete = 2147483647
)

type SlbNewSslCfgSSLInspectFlowTableParams struct {
	// The number of the flow.
	Index int32 `json:"Index,omitempty"`
	// The The Secarity group index.
	SeGroupIndex int32 `json:"SeGroupIndex,omitempty"`
	// Enable/disable security group in flow.
	SeGroupState SlbNewSslCfgSSLInspectFlowTableSeGroupState `json:"SeGroupState,omitempty"`
	// This is the the Group location in the Flow .
	SeGroupLocation uint32 `json:"SeGroupLocation,omitempty"`
	// Each security group has it defined fallback action to be
	// 	done in case the group is down. Can be defined as Continue or Drop.
	SeGroupFallback SlbNewSslCfgSSLInspectFlowTableSeGroupFallback `json:"SeGroupFallback,omitempty"`
	// This is the redirct filter Id .
	SeGroupRedFiltId uint32 `json:"SeGroupRedFiltId,omitempty"`
	// Group ID of the Security group
	GroupID string `json:"GroupID,omitempty"`
	// When set to the value of 2 (delete), the entire row is deleted.
	// When read, other(1) is returned. Setting the value to anything
	// other than 2(delete) has no effect on the state of the row.
	Delete SlbNewSslCfgSSLInspectFlowTableDelete `json:"Delete,omitempty"`
}

func (p SlbNewSslCfgSSLInspectFlowTableParams) iMABean() {}
