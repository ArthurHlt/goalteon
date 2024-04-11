package beans

import (
	"fmt"
	"reflect"
)

// SlbCurSslCfgSSLInspectFlowTable The table of SSL Inspection Flow.
// Note:This mib is not supported for VX instance of Virtualization.
type SlbCurSslCfgSSLInspectFlowTable struct {
	// The number of the flow.
	SlbCurSslCfgSSLInspectFlowIndex int32
	// The Secarity group index.
	SlbCurSslCfgSSLInspectFlowSeGroupIndex int32
	Params                                 *SlbCurSslCfgSSLInspectFlowTableParams
}

func NewSlbCurSslCfgSSLInspectFlowTableList() *SlbCurSslCfgSSLInspectFlowTable {
	return &SlbCurSslCfgSSLInspectFlowTable{}
}

func NewSlbCurSslCfgSSLInspectFlowTable(
	slbCurSslCfgSSLInspectFlowIndex int32,
	slbCurSslCfgSSLInspectFlowSeGroupIndex int32,
	params *SlbCurSslCfgSSLInspectFlowTableParams,
) *SlbCurSslCfgSSLInspectFlowTable {
	return &SlbCurSslCfgSSLInspectFlowTable{
		SlbCurSslCfgSSLInspectFlowIndex:        slbCurSslCfgSSLInspectFlowIndex,
		SlbCurSslCfgSSLInspectFlowSeGroupIndex: slbCurSslCfgSSLInspectFlowSeGroupIndex,
		Params:                                 params,
	}
}

func (c *SlbCurSslCfgSSLInspectFlowTable) Name() string {
	return "SlbCurSslCfgSSLInspectFlowTable"
}

func (c *SlbCurSslCfgSSLInspectFlowTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbCurSslCfgSSLInspectFlowTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbCurSslCfgSSLInspectFlowTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbCurSslCfgSSLInspectFlowIndex).IsZero() &&
		reflect.ValueOf(c.SlbCurSslCfgSSLInspectFlowSeGroupIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbCurSslCfgSSLInspectFlowIndex) + "/" + fmt.Sprint(c.SlbCurSslCfgSSLInspectFlowSeGroupIndex)
}

type SlbCurSslCfgSSLInspectFlowTableSeGroupState int32

const (
	SlbCurSslCfgSSLInspectFlowTableSeGroupState_Enabled     SlbCurSslCfgSSLInspectFlowTableSeGroupState = 2
	SlbCurSslCfgSSLInspectFlowTableSeGroupState_Disabled    SlbCurSslCfgSSLInspectFlowTableSeGroupState = 3
	SlbCurSslCfgSSLInspectFlowTableSeGroupState_Unsupported SlbCurSslCfgSSLInspectFlowTableSeGroupState = 2147483647
)

type SlbCurSslCfgSSLInspectFlowTableSeGroupFallback int32

const (
	SlbCurSslCfgSSLInspectFlowTableSeGroupFallback_None        SlbCurSslCfgSSLInspectFlowTableSeGroupFallback = 1
	SlbCurSslCfgSSLInspectFlowTableSeGroupFallback_Continue    SlbCurSslCfgSSLInspectFlowTableSeGroupFallback = 2
	SlbCurSslCfgSSLInspectFlowTableSeGroupFallback_Drop        SlbCurSslCfgSSLInspectFlowTableSeGroupFallback = 3
	SlbCurSslCfgSSLInspectFlowTableSeGroupFallback_Unsupported SlbCurSslCfgSSLInspectFlowTableSeGroupFallback = 2147483647
)

type SlbCurSslCfgSSLInspectFlowTableParams struct {
	// The number of the flow.
	Index int32 `json:"Index,omitempty"`
	// The Secarity group index.
	SeGroupIndex int32 `json:"SeGroupIndex,omitempty"`
	// Enable/disable security group in flow.
	SeGroupState SlbCurSslCfgSSLInspectFlowTableSeGroupState `json:"SeGroupState,omitempty"`
	// This is the the Group location in the Flow .
	SeGroupLocation uint32 `json:"SeGroupLocation,omitempty"`
	// Each security group has it defined fallback action to be
	// 	done in case the group is down. Can be defined as Continue or Drop.
	SeGroupFallback SlbCurSslCfgSSLInspectFlowTableSeGroupFallback `json:"SeGroupFallback,omitempty"`
	// This is the redirct filter Id .
	SeGroupRedFiltId uint32 `json:"SeGroupRedFiltId,omitempty"`
	// Group ID of the Security group
	GroupID string `json:"GroupID,omitempty"`
}

func (p SlbCurSslCfgSSLInspectFlowTableParams) iMABean() {}
