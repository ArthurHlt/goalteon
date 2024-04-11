package beans

import (
	"fmt"
	"reflect"
)

// SlbCurSslCfgSSLInspectFlowRuleTable The table of SSL Inspection Flow Rule.
// Note:This mib is not supported for VX instance of Virtualization.
type SlbCurSslCfgSSLInspectFlowRuleTable struct {
	// The number of the flow.
	SlbCurSslCfgSSLInspectFlowRuleIndex int32
	Params                              *SlbCurSslCfgSSLInspectFlowRuleTableParams
}

func NewSlbCurSslCfgSSLInspectFlowRuleTableList() *SlbCurSslCfgSSLInspectFlowRuleTable {
	return &SlbCurSslCfgSSLInspectFlowRuleTable{}
}

func NewSlbCurSslCfgSSLInspectFlowRuleTable(
	slbCurSslCfgSSLInspectFlowRuleIndex int32,
	params *SlbCurSslCfgSSLInspectFlowRuleTableParams,
) *SlbCurSslCfgSSLInspectFlowRuleTable {
	return &SlbCurSslCfgSSLInspectFlowRuleTable{
		SlbCurSslCfgSSLInspectFlowRuleIndex: slbCurSslCfgSSLInspectFlowRuleIndex,
		Params:                              params,
	}
}

func (c *SlbCurSslCfgSSLInspectFlowRuleTable) Name() string {
	return "SlbCurSslCfgSSLInspectFlowRuleTable"
}

func (c *SlbCurSslCfgSSLInspectFlowRuleTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbCurSslCfgSSLInspectFlowRuleTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbCurSslCfgSSLInspectFlowRuleTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbCurSslCfgSSLInspectFlowRuleIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbCurSslCfgSSLInspectFlowRuleIndex)
}

type SlbCurSslCfgSSLInspectFlowRuleTableProxyType int32

const (
	SlbCurSslCfgSSLInspectFlowRuleTableProxyType_Transparent SlbCurSslCfgSSLInspectFlowRuleTableProxyType = 1
	SlbCurSslCfgSSLInspectFlowRuleTableProxyType_Expicit     SlbCurSslCfgSSLInspectFlowRuleTableProxyType = 2
	SlbCurSslCfgSSLInspectFlowRuleTableProxyType_Unsupported SlbCurSslCfgSSLInspectFlowRuleTableProxyType = 2147483647
)

type SlbCurSslCfgSSLInspectFlowRuleTableInspectHttp int32

const (
	SlbCurSslCfgSSLInspectFlowRuleTableInspectHttp_True        SlbCurSslCfgSSLInspectFlowRuleTableInspectHttp = 1
	SlbCurSslCfgSSLInspectFlowRuleTableInspectHttp_False       SlbCurSslCfgSSLInspectFlowRuleTableInspectHttp = 2
	SlbCurSslCfgSSLInspectFlowRuleTableInspectHttp_Unsupported SlbCurSslCfgSSLInspectFlowRuleTableInspectHttp = 2147483647
)

type SlbCurSslCfgSSLInspectFlowRuleTableOutboundAction int32

const (
	SlbCurSslCfgSSLInspectFlowRuleTableOutboundAction_Forward     SlbCurSslCfgSSLInspectFlowRuleTableOutboundAction = 1
	SlbCurSslCfgSSLInspectFlowRuleTableOutboundAction_Nat         SlbCurSslCfgSSLInspectFlowRuleTableOutboundAction = 2
	SlbCurSslCfgSSLInspectFlowRuleTableOutboundAction_Llb         SlbCurSslCfgSSLInspectFlowRuleTableOutboundAction = 3
	SlbCurSslCfgSSLInspectFlowRuleTableOutboundAction_Unsupported SlbCurSslCfgSSLInspectFlowRuleTableOutboundAction = 2147483647
)

type SlbCurSslCfgSSLInspectFlowRuleTableSrcPortPreserve int32

const (
	SlbCurSslCfgSSLInspectFlowRuleTableSrcPortPreserve_Enabled     SlbCurSslCfgSSLInspectFlowRuleTableSrcPortPreserve = 1
	SlbCurSslCfgSSLInspectFlowRuleTableSrcPortPreserve_Disabled    SlbCurSslCfgSSLInspectFlowRuleTableSrcPortPreserve = 2
	SlbCurSslCfgSSLInspectFlowRuleTableSrcPortPreserve_Unsupported SlbCurSslCfgSSLInspectFlowRuleTableSrcPortPreserve = 2147483647
)

type SlbCurSslCfgSSLInspectFlowRuleTableParams struct {
	// The number of the flow.
	Index int32 `json:"Index,omitempty"`
	// The type of Proxy if 1 so transparent, if 2 so expicit .
	ProxyType SlbCurSslCfgSSLInspectFlowRuleTableProxyType `json:"ProxyType,omitempty"`
	// Flag For inspect HTTP , value 1 is to Inspect value 2 not to Inspect.
	InspectHttp SlbCurSslCfgSSLInspectFlowRuleTableInspectHttp `json:"InspectHttp,omitempty"`
	// HTTPS Port Number
	HttpsPort uint64 `json:"HttpsPort,omitempty"`
	// HTTPS Decrypt Port Number
	HttpsDecryptPort uint64 `json:"HttpsDecryptPort,omitempty"`
	// HTTP Port Number
	HttpPort uint64 `json:"HttpPort,omitempty"`
	// Ingress Ports Number
	IngressPort string `json:"IngressPort,omitempty"`
	// Egress Ports Number
	EgressPort string `json:"EgressPort,omitempty"`
	// The Outbound action if 1 so forward if 2 so nat if 3 llb .
	OutboundAction SlbCurSslCfgSSLInspectFlowRuleTableOutboundAction `json:"OutboundAction,omitempty"`
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
	// NAT MASK address
	NATMASKAddr string `json:"NATMASKAddr,omitempty"`
	// SSL Inspection source port preservation either enable or disable.
	SrcPortPreserve SlbCurSslCfgSSLInspectFlowRuleTableSrcPortPreserve `json:"SrcPortPreserve,omitempty"`
	// Server Authentication Policy
	ServerAuthPol string `json:"ServerAuthPol,omitempty"`
}

func (p SlbCurSslCfgSSLInspectFlowRuleTableParams) iMABean() {}
