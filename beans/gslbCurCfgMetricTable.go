package beans

import (
	"fmt"
	"reflect"
)

// GslbCurCfgMetricTable The metric table in the current configuration block.
// Note:This mib is not supported for VX instance of Virtualization.
type GslbCurCfgMetricTable struct {
	// The rule table index in the current configuration block.
	GslbCurCfgRuleMetricIndx int32
	// The metric table index in the current configuration block.
	GslbCurCfgMetricIndx int32
	Params               *GslbCurCfgMetricTableParams
}

func NewGslbCurCfgMetricTableList() *GslbCurCfgMetricTable {
	return &GslbCurCfgMetricTable{}
}

func NewGslbCurCfgMetricTable(
	gslbCurCfgRuleMetricIndx int32,
	gslbCurCfgMetricIndx int32,
	params *GslbCurCfgMetricTableParams,
) *GslbCurCfgMetricTable {
	return &GslbCurCfgMetricTable{
		GslbCurCfgRuleMetricIndx: gslbCurCfgRuleMetricIndx,
		GslbCurCfgMetricIndx:     gslbCurCfgMetricIndx,
		Params:                   params,
	}
}

func (c *GslbCurCfgMetricTable) Name() string {
	return "GslbCurCfgMetricTable"
}

func (c *GslbCurCfgMetricTable) GetParams() BeanType {
	return c.Params
}

func (c *GslbCurCfgMetricTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *GslbCurCfgMetricTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.GslbCurCfgRuleMetricIndx).IsZero() &&
		reflect.ValueOf(c.GslbCurCfgMetricIndx).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.GslbCurCfgRuleMetricIndx) + "/" + fmt.Sprint(c.GslbCurCfgMetricIndx)
}

type GslbCurCfgMetricTableMetric int32

const (
	GslbCurCfgMetricTableMetric_Leastconns    GslbCurCfgMetricTableMetric = 1
	GslbCurCfgMetricTableMetric_Roundrobin    GslbCurCfgMetricTableMetric = 2
	GslbCurCfgMetricTableMetric_Response      GslbCurCfgMetricTableMetric = 3
	GslbCurCfgMetricTableMetric_Geographical  GslbCurCfgMetricTableMetric = 4
	GslbCurCfgMetricTableMetric_Network       GslbCurCfgMetricTableMetric = 5
	GslbCurCfgMetricTableMetric_Random        GslbCurCfgMetricTableMetric = 6
	GslbCurCfgMetricTableMetric_Availability  GslbCurCfgMetricTableMetric = 7
	GslbCurCfgMetricTableMetric_Qos           GslbCurCfgMetricTableMetric = 8
	GslbCurCfgMetricTableMetric_Minmisses     GslbCurCfgMetricTableMetric = 9
	GslbCurCfgMetricTableMetric_Hash          GslbCurCfgMetricTableMetric = 10
	GslbCurCfgMetricTableMetric_Local         GslbCurCfgMetricTableMetric = 11
	GslbCurCfgMetricTableMetric_Always        GslbCurCfgMetricTableMetric = 12
	GslbCurCfgMetricTableMetric_Remote        GslbCurCfgMetricTableMetric = 13
	GslbCurCfgMetricTableMetric_None          GslbCurCfgMetricTableMetric = 14
	GslbCurCfgMetricTableMetric_Persistence   GslbCurCfgMetricTableMetric = 15
	GslbCurCfgMetricTableMetric_Phash         GslbCurCfgMetricTableMetric = 16
	GslbCurCfgMetricTableMetric_Proximity     GslbCurCfgMetricTableMetric = 17
	GslbCurCfgMetricTableMetric_Bandwidth     GslbCurCfgMetricTableMetric = 18
	GslbCurCfgMetricTableMetric_Absleastconns GslbCurCfgMetricTableMetric = 19
	GslbCurCfgMetricTableMetric_Unsupported   GslbCurCfgMetricTableMetric = 2147483647
)

type GslbCurCfgMetricTableParams struct {
	// The rule table index in the current configuration block.
	RuleMetricIndx int32 `json:"RuleMetricIndx,omitempty"`
	// The metric table index in the current configuration block.
	Indx int32 `json:"Indx,omitempty"`
	// The metric to use to select next server in
	// the current configuration block.
	Metric GslbCurCfgMetricTableMetric `json:"Metric,omitempty"`
	// The networks added to the metric.  The networks added are
	// presented in bitmap format.
	// in receiving order:
	// OCTET 1  OCTET 2  .....
	// xxxxxxxx xxxxxxxx .....
	// ||    || |_ network 9
	// ||    ||
	// ||    ||___ network 8
	// ||    |____ network 7
	// ||      .    .   .
	// ||_________ network 2
	// |__________ network 1
	// where x : 1 - The represented network added to the metric
	// 0 - The represented network added to the metric
	NetworkBmap string `json:"NetworkBmap,omitempty"`
}

func (p GslbCurCfgMetricTableParams) iMABean() {}
