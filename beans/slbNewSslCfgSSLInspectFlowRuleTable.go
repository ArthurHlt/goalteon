package beans

import (
	"fmt"
	"reflect"
)

// SlbNewSslCfgSSLInspectFlowRuleTable The table of SSL Inspection Flow Rule.
// Note:This mib is not supported for VX instance of Virtualization.
type SlbNewSslCfgSSLInspectFlowRuleTable struct {
	// The number of the flow.
	SlbNewSslCfgSSLInspectFlowRuleIndex int32
	Params                              *SlbNewSslCfgSSLInspectFlowRuleTableParams
}

func NewSlbNewSslCfgSSLInspectFlowRuleTableList() *SlbNewSslCfgSSLInspectFlowRuleTable {
	return &SlbNewSslCfgSSLInspectFlowRuleTable{}
}

func NewSlbNewSslCfgSSLInspectFlowRuleTable(
	slbNewSslCfgSSLInspectFlowRuleIndex int32,
	params *SlbNewSslCfgSSLInspectFlowRuleTableParams,
) *SlbNewSslCfgSSLInspectFlowRuleTable {
	return &SlbNewSslCfgSSLInspectFlowRuleTable{
		SlbNewSslCfgSSLInspectFlowRuleIndex: slbNewSslCfgSSLInspectFlowRuleIndex,
		Params:                              params,
	}
}

func (c *SlbNewSslCfgSSLInspectFlowRuleTable) Name() string {
	return "SlbNewSslCfgSSLInspectFlowRuleTable"
}

func (c *SlbNewSslCfgSSLInspectFlowRuleTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbNewSslCfgSSLInspectFlowRuleTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbNewSslCfgSSLInspectFlowRuleTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbNewSslCfgSSLInspectFlowRuleIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbNewSslCfgSSLInspectFlowRuleIndex)
}

type SlbNewSslCfgSSLInspectFlowRuleTableProxyType int32

const (
	SlbNewSslCfgSSLInspectFlowRuleTableProxyType_Transparent SlbNewSslCfgSSLInspectFlowRuleTableProxyType = 1
	SlbNewSslCfgSSLInspectFlowRuleTableProxyType_Expicit     SlbNewSslCfgSSLInspectFlowRuleTableProxyType = 2
	SlbNewSslCfgSSLInspectFlowRuleTableProxyType_Unsupported SlbNewSslCfgSSLInspectFlowRuleTableProxyType = 2147483647
)

type SlbNewSslCfgSSLInspectFlowRuleTableInspectHttp int32

const (
	SlbNewSslCfgSSLInspectFlowRuleTableInspectHttp_True        SlbNewSslCfgSSLInspectFlowRuleTableInspectHttp = 1
	SlbNewSslCfgSSLInspectFlowRuleTableInspectHttp_False       SlbNewSslCfgSSLInspectFlowRuleTableInspectHttp = 2
	SlbNewSslCfgSSLInspectFlowRuleTableInspectHttp_Unsupported SlbNewSslCfgSSLInspectFlowRuleTableInspectHttp = 2147483647
)

type SlbNewSslCfgSSLInspectFlowRuleTableOutboundAction int32

const (
	SlbNewSslCfgSSLInspectFlowRuleTableOutboundAction_Forward     SlbNewSslCfgSSLInspectFlowRuleTableOutboundAction = 1
	SlbNewSslCfgSSLInspectFlowRuleTableOutboundAction_Nat         SlbNewSslCfgSSLInspectFlowRuleTableOutboundAction = 2
	SlbNewSslCfgSSLInspectFlowRuleTableOutboundAction_Llb         SlbNewSslCfgSSLInspectFlowRuleTableOutboundAction = 3
	SlbNewSslCfgSSLInspectFlowRuleTableOutboundAction_Unsupported SlbNewSslCfgSSLInspectFlowRuleTableOutboundAction = 2147483647
)

type SlbNewSslCfgSSLInspectFlowRuleTableDelete int32

const (
	SlbNewSslCfgSSLInspectFlowRuleTableDelete_Other       SlbNewSslCfgSSLInspectFlowRuleTableDelete = 1
	SlbNewSslCfgSSLInspectFlowRuleTableDelete_Delete      SlbNewSslCfgSSLInspectFlowRuleTableDelete = 2
	SlbNewSslCfgSSLInspectFlowRuleTableDelete_Unsupported SlbNewSslCfgSSLInspectFlowRuleTableDelete = 2147483647
)

type SlbNewSslCfgSSLInspectFlowRuleTableSrcPortPreserve int32

const (
	SlbNewSslCfgSSLInspectFlowRuleTableSrcPortPreserve_Enabled     SlbNewSslCfgSSLInspectFlowRuleTableSrcPortPreserve = 1
	SlbNewSslCfgSSLInspectFlowRuleTableSrcPortPreserve_Disabled    SlbNewSslCfgSSLInspectFlowRuleTableSrcPortPreserve = 2
	SlbNewSslCfgSSLInspectFlowRuleTableSrcPortPreserve_Unsupported SlbNewSslCfgSSLInspectFlowRuleTableSrcPortPreserve = 2147483647
)

type SlbNewSslCfgSSLInspectFlowRuleTableParams struct {
	// The number of the flow.
	Index int32 `json:"Index,omitempty"`
	// The type of Proxy if 1 so transparent, if 2 so expicit .
	ProxyType SlbNewSslCfgSSLInspectFlowRuleTableProxyType `json:"ProxyType,omitempty"`
	// Flag For inspect HTTP , value 1 is to Inspect value 2 not to Inspect.
	InspectHttp SlbNewSslCfgSSLInspectFlowRuleTableInspectHttp `json:"InspectHttp,omitempty"`
	// HTTPS Port Number
	HttpsPort uint64 `json:"HttpsPort,omitempty"`
	// HTTPS Decrypt Port Number
	HttpsDecryptPort uint64 `json:"HttpsDecryptPort,omitempty"`
	// HTTP Port Number
	HttpPort uint64 `json:"HttpPort,omitempty"`
	// Ingress Port Number
	// 	Using the following format: xx, xx
	IngressPort string `json:"IngressPort,omitempty"`
	// Egress Port Number
	// 	Using the following format: xx, xx
	EgressPort string `json:"EgressPort,omitempty"`
	// Ingress Port Number
	AddIngressPort int32 `json:"AddIngressPort,omitempty"`
	// Egress Port Number
	AddEgressPort int32 `json:"AddEgressPort,omitempty"`
	// Ingress Port Number
	RemIngressPort int32 `json:"RemIngressPort,omitempty"`
	// Egress Port Number
	RemEgressPort int32 `json:"RemEgressPort,omitempty"`
	// The Outbound action if 1 so forward if 2 so nat if 3 llb .
	OutboundAction SlbNewSslCfgSSLInspectFlowRuleTableOutboundAction `json:"OutboundAction,omitempty"`
	// The floating IP number in alphanumeric ID.
	FloatId string `json:"FloatId,omitempty"`
	// Outband fillter https Id
	OutboundFilterHttps uint32 `json:"OutboundFilterHttps,omitempty"`
	// Outband fillter http Id
	OutboundFilterHttp uint32 `json:"OutboundFilterHttp,omitempty"`
	// Backend fillter https Id
	BackendFilterHttps uint32 `json:"BackendFilterHttps,omitempty"`
	// Backend fillter http Id
	BackendFilterHttp uint32 `json:"BackendFilterHttp,omitempty"`
	// Network address for interanl interfaces
	NetworkAddr string `json:"NetworkAddr,omitempty"`
	// Gateway address
	GWAddr string `json:"GWAddr,omitempty"`
	// NAT address
	NATAddr string `json:"NATAddr,omitempty"`
	// The WAN Link Group ID.
	WanLinkGrID string `json:"WanLinkGrID,omitempty"`
	// By setting the value to delete(2), the entire row is
	// deleted.
	Delete SlbNewSslCfgSSLInspectFlowRuleTableDelete `json:"Delete,omitempty"`
	// NAT Mask address
	NATMASKAddr string `json:"NATMASKAddr,omitempty"`
	// Enable or disable ssl inspection source port preservation.
	SrcPortPreserve SlbNewSslCfgSSLInspectFlowRuleTableSrcPortPreserve `json:"SrcPortPreserve,omitempty"`
	// Server Authentication Policy
	ServerAuthPol string `json:"ServerAuthPol,omitempty"`
}

func (p SlbNewSslCfgSSLInspectFlowRuleTableParams) iMABean() {}
