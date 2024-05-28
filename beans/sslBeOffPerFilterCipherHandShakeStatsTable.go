package beans

import (
	"fmt"
	"reflect"
)

// SslBeOffPerFilterCipherHandShakeStatsTable A table for backend SSL Cipher hand shake current and total statistics per filter.
// Note:This mib is not supported for VX instance of Virtualization.
type SslBeOffPerFilterCipherHandShakeStatsTable struct {
	// Filter ID.
	SslBeOffPerFilterCipherHandShakeFiltId int32
	// Cipher index.
	SslBeOffPerFilterByCipherHandShakeCipherName string
	Params                                       *SslBeOffPerFilterCipherHandShakeStatsTableParams
}

func NewSslBeOffPerFilterCipherHandShakeStatsTableList() *SslBeOffPerFilterCipherHandShakeStatsTable {
	return &SslBeOffPerFilterCipherHandShakeStatsTable{}
}

func NewSslBeOffPerFilterCipherHandShakeStatsTable(
	sslBeOffPerFilterCipherHandShakeFiltId int32,
	sslBeOffPerFilterByCipherHandShakeCipherName string,
	params *SslBeOffPerFilterCipherHandShakeStatsTableParams,
) *SslBeOffPerFilterCipherHandShakeStatsTable {
	return &SslBeOffPerFilterCipherHandShakeStatsTable{
		SslBeOffPerFilterCipherHandShakeFiltId:       sslBeOffPerFilterCipherHandShakeFiltId,
		SslBeOffPerFilterByCipherHandShakeCipherName: sslBeOffPerFilterByCipherHandShakeCipherName,
		Params: params,
	}
}

func (c *SslBeOffPerFilterCipherHandShakeStatsTable) Name() string {
	return "SslBeOffPerFilterCipherHandShakeStatsTable"
}

func (c *SslBeOffPerFilterCipherHandShakeStatsTable) GetParams() BeanType {
	return c.Params
}

func (c *SslBeOffPerFilterCipherHandShakeStatsTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SslBeOffPerFilterCipherHandShakeStatsTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SslBeOffPerFilterCipherHandShakeFiltId).IsZero() &&
		reflect.ValueOf(c.SslBeOffPerFilterByCipherHandShakeCipherName).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SslBeOffPerFilterCipherHandShakeFiltId) + "/" + fmt.Sprint(c.SslBeOffPerFilterByCipherHandShakeCipherName)
}

type SslBeOffPerFilterCipherHandShakeStatsTableParams struct {
	// Filter ID.
	FiltId int32 `json:"FiltId,omitempty"`
	// Cipher index.
	ByCipherHandShakeCipherName string `json:"ByCipherHandShakeCipherName,omitempty"`
	// Number of existing SSL handshakes re-used by AAS to communicate with server per second for filter.
	ByCipherHandShakeCipherHits int32 `json:"ByCipherHandShakeCipherHits,omitempty"`
	// Number of total SSL handshakes re-used by AAS to communicate with server per second for filter.
	ByCipherHandShakeCipherHitsTotal int32 `json:"ByCipherHandShakeCipherHitsTotal,omitempty"`
}

func (p SslBeOffPerFilterCipherHandShakeStatsTableParams) iMABean() {}
