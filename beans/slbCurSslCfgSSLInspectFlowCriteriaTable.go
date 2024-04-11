package beans

import (
	"fmt"
	"reflect"
)

// SlbCurSslCfgSSLInspectFlowCriteriaTable The table of SSL Inspection Flow Criteria.
// Note:This mib is not supported for VX instance of Virtualization.
type SlbCurSslCfgSSLInspectFlowCriteriaTable struct {
	// The number of the flow.
	SlbCurSslCfgSSLInspectFlowCriteriaIndex int32
	// The Inspect flow Criteria name id  .
	SlbCurSslCfgSSLInspectFlowCriteriaRuleId string
	Params                                   *SlbCurSslCfgSSLInspectFlowCriteriaTableParams
}

func NewSlbCurSslCfgSSLInspectFlowCriteriaTableList() *SlbCurSslCfgSSLInspectFlowCriteriaTable {
	return &SlbCurSslCfgSSLInspectFlowCriteriaTable{}
}

func NewSlbCurSslCfgSSLInspectFlowCriteriaTable(
	slbCurSslCfgSSLInspectFlowCriteriaIndex int32,
	slbCurSslCfgSSLInspectFlowCriteriaRuleId string,
	params *SlbCurSslCfgSSLInspectFlowCriteriaTableParams,
) *SlbCurSslCfgSSLInspectFlowCriteriaTable {
	return &SlbCurSslCfgSSLInspectFlowCriteriaTable{
		SlbCurSslCfgSSLInspectFlowCriteriaIndex:  slbCurSslCfgSSLInspectFlowCriteriaIndex,
		SlbCurSslCfgSSLInspectFlowCriteriaRuleId: slbCurSslCfgSSLInspectFlowCriteriaRuleId,
		Params:                                   params,
	}
}

func (c *SlbCurSslCfgSSLInspectFlowCriteriaTable) Name() string {
	return "SlbCurSslCfgSSLInspectFlowCriteriaTable"
}

func (c *SlbCurSslCfgSSLInspectFlowCriteriaTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbCurSslCfgSSLInspectFlowCriteriaTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbCurSslCfgSSLInspectFlowCriteriaTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbCurSslCfgSSLInspectFlowCriteriaIndex).IsZero() &&
		reflect.ValueOf(c.SlbCurSslCfgSSLInspectFlowCriteriaRuleId).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbCurSslCfgSSLInspectFlowCriteriaIndex) + "/" + fmt.Sprint(c.SlbCurSslCfgSSLInspectFlowCriteriaRuleId)
}

type SlbCurSslCfgSSLInspectFlowCriteriaTableAction int32

const (
	SlbCurSslCfgSSLInspectFlowCriteriaTableAction_Inspect     SlbCurSslCfgSSLInspectFlowCriteriaTableAction = 1
	SlbCurSslCfgSSLInspectFlowCriteriaTableAction_Bypass      SlbCurSslCfgSSLInspectFlowCriteriaTableAction = 2
	SlbCurSslCfgSSLInspectFlowCriteriaTableAction_Unsupported SlbCurSslCfgSSLInspectFlowCriteriaTableAction = 2147483647
)

type SlbCurSslCfgSSLInspectFlowCriteriaTableContMatch int32

const (
	SlbCurSslCfgSSLInspectFlowCriteriaTableContMatch_None          SlbCurSslCfgSSLInspectFlowCriteriaTableContMatch = 1
	SlbCurSslCfgSSLInspectFlowCriteriaTableContMatch_Urlcategories SlbCurSslCfgSSLInspectFlowCriteriaTableContMatch = 2
	SlbCurSslCfgSSLInspectFlowCriteriaTableContMatch_Hostnamelist  SlbCurSslCfgSSLInspectFlowCriteriaTableContMatch = 3
	SlbCurSslCfgSSLInspectFlowCriteriaTableContMatch_Unsupported   SlbCurSslCfgSSLInspectFlowCriteriaTableContMatch = 2147483647
)

type SlbCurSslCfgSSLInspectFlowCriteriaTableReportState int32

const (
	SlbCurSslCfgSSLInspectFlowCriteriaTableReportState_Disabled    SlbCurSslCfgSSLInspectFlowCriteriaTableReportState = 0
	SlbCurSslCfgSSLInspectFlowCriteriaTableReportState_Enabled     SlbCurSslCfgSSLInspectFlowCriteriaTableReportState = 1
	SlbCurSslCfgSSLInspectFlowCriteriaTableReportState_Unsupported SlbCurSslCfgSSLInspectFlowCriteriaTableReportState = 2147483647
)

type SlbCurSslCfgSSLInspectFlowCriteriaTableParams struct {
	// The number of the flow.
	Index int32 `json:"Index,omitempty"`
	// The Inspect flow Criteria name id  .
	RuleId string `json:"RuleId,omitempty"`
	// The Inspect flow Criteria Https filter id.
	HttpsFilter uint32 `json:"HttpsFilter,omitempty"`
	// The Inspect flow Criteria Http filter id.
	HttpFilter uint32 `json:"HttpFilter,omitempty"`
	// The type of action if 1 so inspect, if 2 so bypass .
	Action SlbCurSslCfgSSLInspectFlowCriteriaTableAction `json:"Action,omitempty"`
	// Configures whether URL filtering or a content class is attached to the filter .
	ContMatch SlbCurSslCfgSSLInspectFlowCriteriaTableContMatch `json:"ContMatch,omitempty"`
	// report state ena or disable .
	ReportState SlbCurSslCfgSSLInspectFlowCriteriaTableReportState `json:"ReportState,omitempty"`
	// The Inspect flow Criteria dataclass.
	Dataclss string `json:"Dataclss,omitempty"`
}

func (p SlbCurSslCfgSSLInspectFlowCriteriaTableParams) iMABean() {}
