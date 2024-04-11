package beans

import (
	"fmt"
	"reflect"
)

// GslbNewCfgMetricTable The metric table in the current configuration block.
// Note:This mib is not supported for VX instance of Virtualization.
type GslbNewCfgMetricTable struct {
	// The rule table index in the new configuration block.
	GslbNewCfgRuleMetricIndx int32
	// The metric table index in the new configuration block.
	GslbNewCfgMetricIndx int32
	Params               *GslbNewCfgMetricTableParams
}

func NewGslbNewCfgMetricTableList() *GslbNewCfgMetricTable {
	return &GslbNewCfgMetricTable{}
}

func NewGslbNewCfgMetricTable(
	gslbNewCfgRuleMetricIndx int32,
	gslbNewCfgMetricIndx int32,
	params *GslbNewCfgMetricTableParams,
) *GslbNewCfgMetricTable {
	return &GslbNewCfgMetricTable{
		GslbNewCfgRuleMetricIndx: gslbNewCfgRuleMetricIndx,
		GslbNewCfgMetricIndx:     gslbNewCfgMetricIndx,
		Params:                   params,
	}
}

func (c *GslbNewCfgMetricTable) Name() string {
	return "GslbNewCfgMetricTable"
}

func (c *GslbNewCfgMetricTable) GetParams() BeanType {
	return c.Params
}

func (c *GslbNewCfgMetricTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *GslbNewCfgMetricTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.GslbNewCfgRuleMetricIndx).IsZero() &&
		reflect.ValueOf(c.GslbNewCfgMetricIndx).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.GslbNewCfgRuleMetricIndx) + "/" + fmt.Sprint(c.GslbNewCfgMetricIndx)
}

type GslbNewCfgMetricTableMetric int32

const (
	GslbNewCfgMetricTableMetric_Leastconns    GslbNewCfgMetricTableMetric = 1
	GslbNewCfgMetricTableMetric_Roundrobin    GslbNewCfgMetricTableMetric = 2
	GslbNewCfgMetricTableMetric_Response      GslbNewCfgMetricTableMetric = 3
	GslbNewCfgMetricTableMetric_Geographical  GslbNewCfgMetricTableMetric = 4
	GslbNewCfgMetricTableMetric_Network       GslbNewCfgMetricTableMetric = 5
	GslbNewCfgMetricTableMetric_Random        GslbNewCfgMetricTableMetric = 6
	GslbNewCfgMetricTableMetric_Availability  GslbNewCfgMetricTableMetric = 7
	GslbNewCfgMetricTableMetric_Qos           GslbNewCfgMetricTableMetric = 8
	GslbNewCfgMetricTableMetric_Minmisses     GslbNewCfgMetricTableMetric = 9
	GslbNewCfgMetricTableMetric_Hash          GslbNewCfgMetricTableMetric = 10
	GslbNewCfgMetricTableMetric_Local         GslbNewCfgMetricTableMetric = 11
	GslbNewCfgMetricTableMetric_Always        GslbNewCfgMetricTableMetric = 12
	GslbNewCfgMetricTableMetric_Remote        GslbNewCfgMetricTableMetric = 13
	GslbNewCfgMetricTableMetric_None          GslbNewCfgMetricTableMetric = 14
	GslbNewCfgMetricTableMetric_Persistence   GslbNewCfgMetricTableMetric = 15
	GslbNewCfgMetricTableMetric_Phash         GslbNewCfgMetricTableMetric = 16
	GslbNewCfgMetricTableMetric_Proximity     GslbNewCfgMetricTableMetric = 17
	GslbNewCfgMetricTableMetric_Bandwidth     GslbNewCfgMetricTableMetric = 18
	GslbNewCfgMetricTableMetric_Absleastconns GslbNewCfgMetricTableMetric = 19
	GslbNewCfgMetricTableMetric_Unsupported   GslbNewCfgMetricTableMetric = 2147483647
)

type GslbNewCfgMetricTableParams struct {
	// The rule table index in the new configuration block.
	RuleMetricIndx int32 `json:"RuleMetricIndx,omitempty"`
	// The metric table index in the new configuration block.
	Indx int32 `json:"Indx,omitempty"`
	// The metric to use to select next server in
	// the new configuration block.
	Metric GslbNewCfgMetricTableMetric `json:"Metric,omitempty"`
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
	// This is an action object to add network to a metric. The range of the
	// valid index is between 1 and gslbEnhNetworkTableMaxSize. When read, the
	// value '0' is returned always.
	AddNetwork int32 `json:"AddNetwork,omitempty"`
	// This is an action object to remove network to a metric. The range of the
	// valid index is between 1 and gslbEnhNetworkTableMaxSize. When read, the
	// value '0' is returned always.
	RemNetwork int32 `json:"RemNetwork,omitempty"`
}

func (p GslbNewCfgMetricTableParams) iMABean() {}
