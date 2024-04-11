package beans

import (
	"fmt"
	"reflect"
)

// SlbNewSslCfgSSLInspectFlowCriteriaTable The table of SSL Inspection Flow Criteria.
// Note:This mib is not supported for VX instance of Virtualization.
type SlbNewSslCfgSSLInspectFlowCriteriaTable struct {
	// The number of the flow.
	SlbNewSslCfgSSLInspectFlowCriteriaIndex int32
	// The Inspect flow Criteria name id  .
	SlbNewSslCfgSSLInspectFlowCriteriaRuleId string
	Params                                   *SlbNewSslCfgSSLInspectFlowCriteriaTableParams
}

func NewSlbNewSslCfgSSLInspectFlowCriteriaTableList() *SlbNewSslCfgSSLInspectFlowCriteriaTable {
	return &SlbNewSslCfgSSLInspectFlowCriteriaTable{}
}

func NewSlbNewSslCfgSSLInspectFlowCriteriaTable(
	slbNewSslCfgSSLInspectFlowCriteriaIndex int32,
	slbNewSslCfgSSLInspectFlowCriteriaRuleId string,
	params *SlbNewSslCfgSSLInspectFlowCriteriaTableParams,
) *SlbNewSslCfgSSLInspectFlowCriteriaTable {
	return &SlbNewSslCfgSSLInspectFlowCriteriaTable{
		SlbNewSslCfgSSLInspectFlowCriteriaIndex:  slbNewSslCfgSSLInspectFlowCriteriaIndex,
		SlbNewSslCfgSSLInspectFlowCriteriaRuleId: slbNewSslCfgSSLInspectFlowCriteriaRuleId,
		Params:                                   params,
	}
}

func (c *SlbNewSslCfgSSLInspectFlowCriteriaTable) Name() string {
	return "SlbNewSslCfgSSLInspectFlowCriteriaTable"
}

func (c *SlbNewSslCfgSSLInspectFlowCriteriaTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbNewSslCfgSSLInspectFlowCriteriaTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbNewSslCfgSSLInspectFlowCriteriaTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbNewSslCfgSSLInspectFlowCriteriaIndex).IsZero() &&
		reflect.ValueOf(c.SlbNewSslCfgSSLInspectFlowCriteriaRuleId).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbNewSslCfgSSLInspectFlowCriteriaIndex) + "/" + fmt.Sprint(c.SlbNewSslCfgSSLInspectFlowCriteriaRuleId)
}

type SlbNewSslCfgSSLInspectFlowCriteriaTableAction int32

const (
	SlbNewSslCfgSSLInspectFlowCriteriaTableAction_Inspect     SlbNewSslCfgSSLInspectFlowCriteriaTableAction = 1
	SlbNewSslCfgSSLInspectFlowCriteriaTableAction_Bypass      SlbNewSslCfgSSLInspectFlowCriteriaTableAction = 2
	SlbNewSslCfgSSLInspectFlowCriteriaTableAction_Unsupported SlbNewSslCfgSSLInspectFlowCriteriaTableAction = 2147483647
)

type SlbNewSslCfgSSLInspectFlowCriteriaTableContMatch int32

const (
	SlbNewSslCfgSSLInspectFlowCriteriaTableContMatch_None          SlbNewSslCfgSSLInspectFlowCriteriaTableContMatch = 1
	SlbNewSslCfgSSLInspectFlowCriteriaTableContMatch_Urlcategories SlbNewSslCfgSSLInspectFlowCriteriaTableContMatch = 2
	SlbNewSslCfgSSLInspectFlowCriteriaTableContMatch_Hostnamelist  SlbNewSslCfgSSLInspectFlowCriteriaTableContMatch = 3
	SlbNewSslCfgSSLInspectFlowCriteriaTableContMatch_Unsupported   SlbNewSslCfgSSLInspectFlowCriteriaTableContMatch = 2147483647
)

type SlbNewSslCfgSSLInspectFlowCriteriaTableDelete int32

const (
	SlbNewSslCfgSSLInspectFlowCriteriaTableDelete_Other       SlbNewSslCfgSSLInspectFlowCriteriaTableDelete = 1
	SlbNewSslCfgSSLInspectFlowCriteriaTableDelete_Delete      SlbNewSslCfgSSLInspectFlowCriteriaTableDelete = 2
	SlbNewSslCfgSSLInspectFlowCriteriaTableDelete_Unsupported SlbNewSslCfgSSLInspectFlowCriteriaTableDelete = 2147483647
)

type SlbNewSslCfgSSLInspectFlowCriteriaTableReportState int32

const (
	SlbNewSslCfgSSLInspectFlowCriteriaTableReportState_Disabled    SlbNewSslCfgSSLInspectFlowCriteriaTableReportState = 0
	SlbNewSslCfgSSLInspectFlowCriteriaTableReportState_Enabled     SlbNewSslCfgSSLInspectFlowCriteriaTableReportState = 1
	SlbNewSslCfgSSLInspectFlowCriteriaTableReportState_Unsupported SlbNewSslCfgSSLInspectFlowCriteriaTableReportState = 2147483647
)

type SlbNewSslCfgSSLInspectFlowCriteriaTableParams struct {
	// The number of the flow.
	Index int32 `json:"Index,omitempty"`
	// The Inspect flow Criteria name id  .
	RuleId string `json:"RuleId,omitempty"`
	// The Inspect flow Criteria Https filter id.
	HttpsFilter uint32 `json:"HttpsFilter,omitempty"`
	// The Inspect flow Criteria Http filter id.
	HttpFilter uint32 `json:"HttpFilter,omitempty"`
	// The type of action if 1 so inspect, if 2 so bypass .
	Action SlbNewSslCfgSSLInspectFlowCriteriaTableAction `json:"Action,omitempty"`
	// Configures whether URL filtering or a content class is attached to the filter .
	ContMatch SlbNewSslCfgSSLInspectFlowCriteriaTableContMatch `json:"ContMatch,omitempty"`
	// By setting the value to delete(2), the entire row is
	// deleted.
	Delete SlbNewSslCfgSSLInspectFlowCriteriaTableDelete `json:"Delete,omitempty"`
	// report ena or disable.
	ReportState SlbNewSslCfgSSLInspectFlowCriteriaTableReportState `json:"ReportState,omitempty"`
	// Configure data class for the Inspect flow Criteria.
	Dataclss string `json:"Dataclss,omitempty"`
}

func (p SlbNewSslCfgSSLInspectFlowCriteriaTableParams) iMABean() {}
