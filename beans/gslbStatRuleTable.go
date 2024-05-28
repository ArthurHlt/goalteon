package beans

import (
	"fmt"
	"reflect"
)

// GslbStatRuleTable The Global SLB rule statistics table.
// Note:This mib is not supported for VX instance of Virtualization.
type GslbStatRuleTable struct {
	// The rule number that identifies the rule.
	GslbStatRuleIdx int32
	Params          *GslbStatRuleTableParams
}

func NewGslbStatRuleTableList() *GslbStatRuleTable {
	return &GslbStatRuleTable{}
}

func NewGslbStatRuleTable(
	gslbStatRuleIdx int32,
	params *GslbStatRuleTableParams,
) *GslbStatRuleTable {
	return &GslbStatRuleTable{
		GslbStatRuleIdx: gslbStatRuleIdx,
		Params:          params,
	}
}

func (c *GslbStatRuleTable) Name() string {
	return "GslbStatRuleTable"
}

func (c *GslbStatRuleTable) GetParams() BeanType {
	return c.Params
}

func (c *GslbStatRuleTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *GslbStatRuleTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.GslbStatRuleIdx).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.GslbStatRuleIdx)
}

type GslbStatRuleTableParams struct {
	// The rule number that identifies the rule.
	Idx int32 `json:"Idx,omitempty"`
	// The number of times leastconns is used.
	Leastconns uint32 `json:"Leastconns,omitempty"`
	// The number of times roundrobin is used.
	Roundrobin uint32 `json:"Roundrobin,omitempty"`
	// The number of times minmisses is used.
	Minmisses uint32 `json:"Minmisses,omitempty"`
	// The number of times hash is used.
	Hash uint32 `json:"Hash,omitempty"`
	// The number of times response is used.
	Response uint32 `json:"Response,omitempty"`
	// The number of times geographical is used.
	Geographical uint32 `json:"Geographical,omitempty"`
	// The number of times network is used.
	Network uint32 `json:"Network,omitempty"`
	// The number of times random is used.
	Random uint32 `json:"Random,omitempty"`
	// The number of times availability is used.
	Availability uint32 `json:"Availability,omitempty"`
	// The number of times qos is used.
	Qos uint32 `json:"Qos,omitempty"`
	// The number of times persistence is used.
	Persistence uint32 `json:"Persistence,omitempty"`
	// The number of times local is used.
	Local uint32 `json:"Local,omitempty"`
	// The number of times always is used.
	Always uint32 `json:"Always,omitempty"`
	// The number of times remote is used.
	Remote uint32 `json:"Remote,omitempty"`
	// The number of times all the metrics are used.
	Total uint32 `json:"Total,omitempty"`
	// The number of times phash is used.
	Phash uint32 `json:"Phash,omitempty"`
	// The number of times absolute leastconns is used.
	AbsLeastconns uint32 `json:"AbsLeastconns,omitempty"`
}

func (p GslbStatRuleTableParams) iMABean() {}
