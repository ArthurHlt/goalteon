package beans

import (
	"fmt"
	"reflect"
)

// SslBeOffPerFilterIgnoredCertsReasonsStatsTable A table for backend SSL current and total statistics per filter ignored Certs reasons.
// Note:This mib is not supported for VX instance of Virtualization.
type SslBeOffPerFilterIgnoredCertsReasonsStatsTable struct {
	// Filter backend ignored Certs ID.
	SslBeOffPerFilterIgnoredCertsReasonsFiltId int32
	// Filter backend ignored certs name.
	SslBeOffPerFilterIgnoredCertsReasonsName string
	Params                                   *SslBeOffPerFilterIgnoredCertsReasonsStatsTableParams
}

func NewSslBeOffPerFilterIgnoredCertsReasonsStatsTableList() *SslBeOffPerFilterIgnoredCertsReasonsStatsTable {
	return &SslBeOffPerFilterIgnoredCertsReasonsStatsTable{}
}

func NewSslBeOffPerFilterIgnoredCertsReasonsStatsTable(
	sslBeOffPerFilterIgnoredCertsReasonsFiltId int32,
	sslBeOffPerFilterIgnoredCertsReasonsName string,
	params *SslBeOffPerFilterIgnoredCertsReasonsStatsTableParams,
) *SslBeOffPerFilterIgnoredCertsReasonsStatsTable {
	return &SslBeOffPerFilterIgnoredCertsReasonsStatsTable{
		SslBeOffPerFilterIgnoredCertsReasonsFiltId: sslBeOffPerFilterIgnoredCertsReasonsFiltId,
		SslBeOffPerFilterIgnoredCertsReasonsName:   sslBeOffPerFilterIgnoredCertsReasonsName,
		Params:                                     params,
	}
}

func (c *SslBeOffPerFilterIgnoredCertsReasonsStatsTable) Name() string {
	return "SslBeOffPerFilterIgnoredCertsReasonsStatsTable"
}

func (c *SslBeOffPerFilterIgnoredCertsReasonsStatsTable) GetParams() BeanType {
	return c.Params
}

func (c *SslBeOffPerFilterIgnoredCertsReasonsStatsTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SslBeOffPerFilterIgnoredCertsReasonsStatsTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SslBeOffPerFilterIgnoredCertsReasonsFiltId).IsZero() &&
		reflect.ValueOf(c.SslBeOffPerFilterIgnoredCertsReasonsName).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SslBeOffPerFilterIgnoredCertsReasonsFiltId) + "/" + fmt.Sprint(c.SslBeOffPerFilterIgnoredCertsReasonsName)
}

type SslBeOffPerFilterIgnoredCertsReasonsStatsTableParams struct {
	// Filter backend ignored Certs ID.
	FiltId int32 `json:"FiltId,omitempty"`
	// Filter backend ignored certs name.
	Name string `json:"Name,omitempty"`
	// Number of hits per second for this backend for this filter ignored certs reason.
	Hits int32 `json:"Hits,omitempty"`
	// Number of total hits per second for this backend for this filter ignored certificates reason.
	HitsTotal int32 `json:"HitsTotal,omitempty"`
}

func (p SslBeOffPerFilterIgnoredCertsReasonsStatsTableParams) iMABean() {}
