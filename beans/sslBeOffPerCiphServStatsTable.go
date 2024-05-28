package beans

import (
	"fmt"
	"reflect"
)

// SslBeOffPerCiphServStatsTable A table for backend SSL statistics per virtual service.
// Note:This mib is not supported for VX instance of Virtualization.
type SslBeOffPerCiphServStatsTable struct {
	// The virtual server name that identifies the service.
	SslBeOffPerCiphServVirtServIndex string
	// Virtual server service index.
	SslBeOffPerCiphServVirtServiceIndex int32
	// Virtual server service cipher index.
	SslBeOffPerCiphServVirtCiphIndex int32
	Params                           *SslBeOffPerCiphServStatsTableParams
}

func NewSslBeOffPerCiphServStatsTableList() *SslBeOffPerCiphServStatsTable {
	return &SslBeOffPerCiphServStatsTable{}
}

func NewSslBeOffPerCiphServStatsTable(
	sslBeOffPerCiphServVirtServIndex string,
	sslBeOffPerCiphServVirtServiceIndex int32,
	sslBeOffPerCiphServVirtCiphIndex int32,
	params *SslBeOffPerCiphServStatsTableParams,
) *SslBeOffPerCiphServStatsTable {
	return &SslBeOffPerCiphServStatsTable{
		SslBeOffPerCiphServVirtServIndex:    sslBeOffPerCiphServVirtServIndex,
		SslBeOffPerCiphServVirtServiceIndex: sslBeOffPerCiphServVirtServiceIndex,
		SslBeOffPerCiphServVirtCiphIndex:    sslBeOffPerCiphServVirtCiphIndex,
		Params:                              params,
	}
}

func (c *SslBeOffPerCiphServStatsTable) Name() string {
	return "SslBeOffPerCiphServStatsTable"
}

func (c *SslBeOffPerCiphServStatsTable) GetParams() BeanType {
	return c.Params
}

func (c *SslBeOffPerCiphServStatsTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SslBeOffPerCiphServStatsTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SslBeOffPerCiphServVirtServIndex).IsZero() &&
		reflect.ValueOf(c.SslBeOffPerCiphServVirtServiceIndex).IsZero() &&
		reflect.ValueOf(c.SslBeOffPerCiphServVirtCiphIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SslBeOffPerCiphServVirtServIndex) + "/" + fmt.Sprint(c.SslBeOffPerCiphServVirtServiceIndex) + "/" + fmt.Sprint(c.SslBeOffPerCiphServVirtCiphIndex)
}

type SslBeOffPerCiphServStatsTableParams struct {
	// The virtual server name that identifies the service.
	VirtServIndex string `json:"VirtServIndex,omitempty"`
	// Virtual server service index.
	VirtServiceIndex int32 `json:"VirtServiceIndex,omitempty"`
	// Virtual server service cipher index.
	VirtCiphIndex int32 `json:"VirtCiphIndex,omitempty"`
	// Virtual server service cipher name.
	CiphName string `json:"CiphName,omitempty"`
	// Number of New SSL handshakes between AAS and server per second for this virtual service.
	NewhandShake int32 `json:"NewhandShake,omitempty"`
	// Virtual server service port number.
	VirtServPort int32 `json:"VirtServPort,omitempty"`
	// Number of total SSL handshakes between AAS and server per second for this virtual service.
	NewhandShakeTotal int32 `json:"NewhandShakeTotal,omitempty"`
}

func (p SslBeOffPerCiphServStatsTableParams) iMABean() {}
