package beans

import (
	"fmt"
	"reflect"
)

// SlbOverlaySpStatsMapTable A table for slb overlay mapping statistics .
// Note:This mib is not supported for VX instance of Virtualization.
type SlbOverlaySpStatsMapTable struct {
	// SP index.
	SlbOverlaySpStatsSpIndex int32
	// slb overlay mapping stats table index.
	// This is an alpha numeric index so use it as follows:
	// OID.<str len>.<entry(str)>.<sp val(decimal)> --> OID.2.49.50.10
	SlbOverlaySpStatsMapIndex string
	Params                    *SlbOverlaySpStatsMapTableParams
}

func NewSlbOverlaySpStatsMapTableList() *SlbOverlaySpStatsMapTable {
	return &SlbOverlaySpStatsMapTable{}
}

func NewSlbOverlaySpStatsMapTable(
	slbOverlaySpStatsSpIndex int32,
	slbOverlaySpStatsMapIndex string,
	params *SlbOverlaySpStatsMapTableParams,
) *SlbOverlaySpStatsMapTable {
	return &SlbOverlaySpStatsMapTable{
		SlbOverlaySpStatsSpIndex:  slbOverlaySpStatsSpIndex,
		SlbOverlaySpStatsMapIndex: slbOverlaySpStatsMapIndex,
		Params:                    params,
	}
}

func (c *SlbOverlaySpStatsMapTable) Name() string {
	return "SlbOverlaySpStatsMapTable"
}

func (c *SlbOverlaySpStatsMapTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbOverlaySpStatsMapTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbOverlaySpStatsMapTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbOverlaySpStatsSpIndex).IsZero() &&
		reflect.ValueOf(c.SlbOverlaySpStatsMapIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbOverlaySpStatsSpIndex) + "/" + fmt.Sprint(c.SlbOverlaySpStatsMapIndex)
}

type SlbOverlaySpStatsMapTableParams struct {
	// SP index.
	SpIndex int32 `json:"SpIndex,omitempty"`
	// slb overlay mapping stats table index.
	// This is an alpha numeric index so use it as follows:
	// OID.<str len>.<entry(str)>.<sp val(decimal)> --> OID.2.49.50.10
	Index string `json:"Index,omitempty"`
	// Overlay mapping table hits per SP.
	Hits uint64 `json:"Hits,omitempty"`
	// Overlay mapping table total hits per SP.
	// In oder to retrive the following parameter please do as follow:
	// set entry as 0 with a valid sp number.
	// i.e OID.1.0.<sp num>
	TotalHits uint64 `json:"TotalHits,omitempty"`
	// Overlay mapping table total misses per SP.
	// In oder to retrive the following parameter please do as follow:
	// set entry as 0 with a valid sp number.
	// i.e OID.1.0.<sp num>
	TotalMisses uint64 `json:"TotalMisses,omitempty"`
	// Overlay mapping table total bypass per SP.
	// In oder to retrive the following parameter please do as follow:
	// set entry as 0 with a valid sp number.
	// i.e OID.1.0.<sp num>
	TotalBypass uint64 `json:"TotalBypass,omitempty"`
}

func (p SlbOverlaySpStatsMapTableParams) iMABean() {}
