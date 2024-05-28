package beans

import (
	"fmt"
	"reflect"
)

// SslBeOffPerFilterRejHsReasonsStatsTable A table for backend SSL Rejected hand shake reasons current and total statistics per filter.
// Note:This mib is not supported for VX instance of Virtualization.
type SslBeOffPerFilterRejHsReasonsStatsTable struct {
	// Filter ID.
	SslBeOffPerFilterRejHsReasonsFiltId int32
	// Cipher index.
	SslBeOffPerFilterRejHsReasonsName string
	Params                            *SslBeOffPerFilterRejHsReasonsStatsTableParams
}

func NewSslBeOffPerFilterRejHsReasonsStatsTableList() *SslBeOffPerFilterRejHsReasonsStatsTable {
	return &SslBeOffPerFilterRejHsReasonsStatsTable{}
}

func NewSslBeOffPerFilterRejHsReasonsStatsTable(
	sslBeOffPerFilterRejHsReasonsFiltId int32,
	sslBeOffPerFilterRejHsReasonsName string,
	params *SslBeOffPerFilterRejHsReasonsStatsTableParams,
) *SslBeOffPerFilterRejHsReasonsStatsTable {
	return &SslBeOffPerFilterRejHsReasonsStatsTable{
		SslBeOffPerFilterRejHsReasonsFiltId: sslBeOffPerFilterRejHsReasonsFiltId,
		SslBeOffPerFilterRejHsReasonsName:   sslBeOffPerFilterRejHsReasonsName,
		Params:                              params,
	}
}

func (c *SslBeOffPerFilterRejHsReasonsStatsTable) Name() string {
	return "SslBeOffPerFilterRejHsReasonsStatsTable"
}

func (c *SslBeOffPerFilterRejHsReasonsStatsTable) GetParams() BeanType {
	return c.Params
}

func (c *SslBeOffPerFilterRejHsReasonsStatsTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SslBeOffPerFilterRejHsReasonsStatsTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SslBeOffPerFilterRejHsReasonsFiltId).IsZero() &&
		reflect.ValueOf(c.SslBeOffPerFilterRejHsReasonsName).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SslBeOffPerFilterRejHsReasonsFiltId) + "/" + fmt.Sprint(c.SslBeOffPerFilterRejHsReasonsName)
}

type SslBeOffPerFilterRejHsReasonsStatsTableParams struct {
	// Filter ID.
	FiltId int32 `json:"FiltId,omitempty"`
	// Cipher index.
	Name string `json:"Name,omitempty"`
	// Number of existing SSL handshakes re-used by AAS to communicate with server per second for filter.
	Hits int32 `json:"Hits,omitempty"`
	// Number of total SSL handshakes re-used by AAS to communicate with server per second for filter.
	HitsTotal int32 `json:"HitsTotal,omitempty"`
}

func (p SslBeOffPerFilterRejHsReasonsStatsTableParams) iMABean() {}
