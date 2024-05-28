package beans

import (
	"fmt"
	"reflect"
)

// SslBeOffPerServIgnoredCertsReasonsStatsTable A table for backend SSL reasons current and total statistics per service ignored certs.
// Note:This mib is not supported for VX instance of Virtualization.
type SslBeOffPerServIgnoredCertsReasonsStatsTable struct {
	// The virtual backend server number that identifies the service ignored certs.
	SslBeOffPerServIgnoredCertsReasonsVirtServIndex string
	// Virtual backend server service ignored certs index.
	SslBeOffPerServIgnoredCertsReasonsVirtServiceIndex int32
	// Virtual backend server service ignored certs reason index.
	SslBeOffPerServIgnoredCertsReasonsIgnoredCertsName string
	Params                                             *SslBeOffPerServIgnoredCertsReasonsStatsTableParams
}

func NewSslBeOffPerServIgnoredCertsReasonsStatsTableList() *SslBeOffPerServIgnoredCertsReasonsStatsTable {
	return &SslBeOffPerServIgnoredCertsReasonsStatsTable{}
}

func NewSslBeOffPerServIgnoredCertsReasonsStatsTable(
	sslBeOffPerServIgnoredCertsReasonsVirtServIndex string,
	sslBeOffPerServIgnoredCertsReasonsVirtServiceIndex int32,
	sslBeOffPerServIgnoredCertsReasonsIgnoredCertsName string,
	params *SslBeOffPerServIgnoredCertsReasonsStatsTableParams,
) *SslBeOffPerServIgnoredCertsReasonsStatsTable {
	return &SslBeOffPerServIgnoredCertsReasonsStatsTable{
		SslBeOffPerServIgnoredCertsReasonsVirtServIndex:    sslBeOffPerServIgnoredCertsReasonsVirtServIndex,
		SslBeOffPerServIgnoredCertsReasonsVirtServiceIndex: sslBeOffPerServIgnoredCertsReasonsVirtServiceIndex,
		SslBeOffPerServIgnoredCertsReasonsIgnoredCertsName: sslBeOffPerServIgnoredCertsReasonsIgnoredCertsName,
		Params: params,
	}
}

func (c *SslBeOffPerServIgnoredCertsReasonsStatsTable) Name() string {
	return "SslBeOffPerServIgnoredCertsReasonsStatsTable"
}

func (c *SslBeOffPerServIgnoredCertsReasonsStatsTable) GetParams() BeanType {
	return c.Params
}

func (c *SslBeOffPerServIgnoredCertsReasonsStatsTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SslBeOffPerServIgnoredCertsReasonsStatsTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SslBeOffPerServIgnoredCertsReasonsVirtServIndex).IsZero() &&
		reflect.ValueOf(c.SslBeOffPerServIgnoredCertsReasonsVirtServiceIndex).IsZero() &&
		reflect.ValueOf(c.SslBeOffPerServIgnoredCertsReasonsIgnoredCertsName).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SslBeOffPerServIgnoredCertsReasonsVirtServIndex) + "/" + fmt.Sprint(c.SslBeOffPerServIgnoredCertsReasonsVirtServiceIndex) + "/" + fmt.Sprint(c.SslBeOffPerServIgnoredCertsReasonsIgnoredCertsName)
}

type SslBeOffPerServIgnoredCertsReasonsStatsTableParams struct {
	// The virtual backend server number that identifies the service ignored certs.
	VirtServIndex string `json:"VirtServIndex,omitempty"`
	// Virtual backend server service ignored certs index.
	VirtServiceIndex int32 `json:"VirtServiceIndex,omitempty"`
	// Virtual backend server service ignored certs reason index.
	IgnoredCertsName string `json:"IgnoredCertsName,omitempty"`
	// Virtual backend server service ignored certs name.
	Hits int32 `json:"Hits,omitempty"`
	// Number of New backend SSL per second for this virtual service ignored certs .
	HitsTotal int32 `json:"HitsTotal,omitempty"`
}

func (p SslBeOffPerServIgnoredCertsReasonsStatsTableParams) iMABean() {}
