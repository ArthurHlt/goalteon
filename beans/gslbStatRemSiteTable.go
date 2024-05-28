package beans

import (
	"fmt"
	"reflect"
)

// GslbStatRemSiteTable The Global SLB remote site statistics table.
// Note:This mib is not supported for VX instance of Virtualization.
type GslbStatRemSiteTable struct {
	// The remote site number that identifies the remote site.
	GslbStatRemSiteIdx int32
	Params             *GslbStatRemSiteTableParams
}

func NewGslbStatRemSiteTableList() *GslbStatRemSiteTable {
	return &GslbStatRemSiteTable{}
}

func NewGslbStatRemSiteTable(
	gslbStatRemSiteIdx int32,
	params *GslbStatRemSiteTableParams,
) *GslbStatRemSiteTable {
	return &GslbStatRemSiteTable{
		GslbStatRemSiteIdx: gslbStatRemSiteIdx,
		Params:             params,
	}
}

func (c *GslbStatRemSiteTable) Name() string {
	return "GslbStatRemSiteTable"
}

func (c *GslbStatRemSiteTable) GetParams() BeanType {
	return c.Params
}

func (c *GslbStatRemSiteTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *GslbStatRemSiteTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.GslbStatRemSiteIdx).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.GslbStatRemSiteIdx)
}

type GslbStatRemSiteTableParams struct {
	// The remote site number that identifies the remote site.
	Idx int32 `json:"Idx,omitempty"`
	// The number of DSSPv1 remote site updates sent.
	OutUpdates uint32 `json:"OutUpdates,omitempty"`
	// The number of good DSSPv1 remote site updates received.
	InUpdates uint32 `json:"InUpdates,omitempty"`
	// The number of DSSPv2 remote site updates sent.
	OutUpdates2 uint32 `json:"OutUpdates2,omitempty"`
	// The number of good DSSPv2 remote site updates received.
	InUpdates2 uint32 `json:"InUpdates2,omitempty"`
	// The number of bad remote site updates received.
	InBadUpdates uint32 `json:"InBadUpdates,omitempty"`
}

func (p GslbStatRemSiteTableParams) iMABean() {}
