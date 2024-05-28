package beans

import (
	"fmt"
	"reflect"
)

// SslOffPerFilterRejHsReasonsStatsTable A table for frontend SSL Rejected hand shake reasons current and total statistics per filter.
// Note:This mib is not supported for VX instance of Virtualization.
type SslOffPerFilterRejHsReasonsStatsTable struct {
	// Filter ID.
	SslOffPerFilterRejHsReasonsFiltId int32
	// Cipher index.
	SslOffPerFilterRejHsReasonsName string
	Params                          *SslOffPerFilterRejHsReasonsStatsTableParams
}

func NewSslOffPerFilterRejHsReasonsStatsTableList() *SslOffPerFilterRejHsReasonsStatsTable {
	return &SslOffPerFilterRejHsReasonsStatsTable{}
}

func NewSslOffPerFilterRejHsReasonsStatsTable(
	sslOffPerFilterRejHsReasonsFiltId int32,
	sslOffPerFilterRejHsReasonsName string,
	params *SslOffPerFilterRejHsReasonsStatsTableParams,
) *SslOffPerFilterRejHsReasonsStatsTable {
	return &SslOffPerFilterRejHsReasonsStatsTable{
		SslOffPerFilterRejHsReasonsFiltId: sslOffPerFilterRejHsReasonsFiltId,
		SslOffPerFilterRejHsReasonsName:   sslOffPerFilterRejHsReasonsName,
		Params:                            params,
	}
}

func (c *SslOffPerFilterRejHsReasonsStatsTable) Name() string {
	return "SslOffPerFilterRejHsReasonsStatsTable"
}

func (c *SslOffPerFilterRejHsReasonsStatsTable) GetParams() BeanType {
	return c.Params
}

func (c *SslOffPerFilterRejHsReasonsStatsTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SslOffPerFilterRejHsReasonsStatsTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SslOffPerFilterRejHsReasonsFiltId).IsZero() &&
		reflect.ValueOf(c.SslOffPerFilterRejHsReasonsName).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SslOffPerFilterRejHsReasonsFiltId) + "/" + fmt.Sprint(c.SslOffPerFilterRejHsReasonsName)
}

type SslOffPerFilterRejHsReasonsStatsTableParams struct {
	// Filter ID.
	FiltId int32 `json:"FiltId,omitempty"`
	// Cipher index.
	Name string `json:"Name,omitempty"`
	// Number of existing SSL handshakes re-used by clients to communicate with AAS per second for filter.
	Hits int32 `json:"Hits,omitempty"`
	// Number of existing SSL handshakes re-used by clients to communicate with AAS per second for filter.
	HitsTotal int32 `json:"HitsTotal,omitempty"`
}

func (p SslOffPerFilterRejHsReasonsStatsTableParams) iMABean() {}
