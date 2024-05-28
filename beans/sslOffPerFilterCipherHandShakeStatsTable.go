package beans

import (
	"fmt"
	"reflect"
)

// SslOffPerFilterCipherHandShakeStatsTable A table for frontend SSL Cipher hand shake current and total statistics per filter.
// Note:This mib is not supported for VX instance of Virtualization.
type SslOffPerFilterCipherHandShakeStatsTable struct {
	// Filter ID.
	SslOffPerFilterCipherHandShakeFiltId int32
	// Cipher index.
	SslOffPerFilterByCipherHandShakeCipherName string
	Params                                     *SslOffPerFilterCipherHandShakeStatsTableParams
}

func NewSslOffPerFilterCipherHandShakeStatsTableList() *SslOffPerFilterCipherHandShakeStatsTable {
	return &SslOffPerFilterCipherHandShakeStatsTable{}
}

func NewSslOffPerFilterCipherHandShakeStatsTable(
	sslOffPerFilterCipherHandShakeFiltId int32,
	sslOffPerFilterByCipherHandShakeCipherName string,
	params *SslOffPerFilterCipherHandShakeStatsTableParams,
) *SslOffPerFilterCipherHandShakeStatsTable {
	return &SslOffPerFilterCipherHandShakeStatsTable{
		SslOffPerFilterCipherHandShakeFiltId:       sslOffPerFilterCipherHandShakeFiltId,
		SslOffPerFilterByCipherHandShakeCipherName: sslOffPerFilterByCipherHandShakeCipherName,
		Params: params,
	}
}

func (c *SslOffPerFilterCipherHandShakeStatsTable) Name() string {
	return "SslOffPerFilterCipherHandShakeStatsTable"
}

func (c *SslOffPerFilterCipherHandShakeStatsTable) GetParams() BeanType {
	return c.Params
}

func (c *SslOffPerFilterCipherHandShakeStatsTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SslOffPerFilterCipherHandShakeStatsTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SslOffPerFilterCipherHandShakeFiltId).IsZero() &&
		reflect.ValueOf(c.SslOffPerFilterByCipherHandShakeCipherName).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SslOffPerFilterCipherHandShakeFiltId) + "/" + fmt.Sprint(c.SslOffPerFilterByCipherHandShakeCipherName)
}

type SslOffPerFilterCipherHandShakeStatsTableParams struct {
	// Filter ID.
	FiltId int32 `json:"FiltId,omitempty"`
	// Cipher index.
	ByCipherHandShakeCipherName string `json:"ByCipherHandShakeCipherName,omitempty"`
	// Number of existing SSL handshakes re-used by clients to communicate with AAS per second for filter.
	ByCipherHandShakeCipherHits int32 `json:"ByCipherHandShakeCipherHits,omitempty"`
	// Number of existing SSL handshakes re-used by clients to communicate with AAS per second for filter.
	ByCipherHandShakeCipherHitsTotal int32 `json:"ByCipherHandShakeCipherHitsTotal,omitempty"`
}

func (p SslOffPerFilterCipherHandShakeStatsTableParams) iMABean() {}
