package beans

import (
	"fmt"
	"reflect"
)

// SslOffPerCiphServStatsTable A table for SSL statistics per virtual service.
// Note:This mib is not supported on VX instance of Virtualization.
type SslOffPerCiphServStatsTable struct {
	// The virtual server name that identifies the service.
	SslOffPerCiphServVirtServIndex string
	// Virtual server service index.
	SslOffPerCiphServVirtServiceIndex int32
	// Virtual server service cipher index.
	SslOffPerCiphServVirtCiphIndex int32
	Params                         *SslOffPerCiphServStatsTableParams
}

func NewSslOffPerCiphServStatsTableList() *SslOffPerCiphServStatsTable {
	return &SslOffPerCiphServStatsTable{}
}

func NewSslOffPerCiphServStatsTable(
	sslOffPerCiphServVirtServIndex string,
	sslOffPerCiphServVirtServiceIndex int32,
	sslOffPerCiphServVirtCiphIndex int32,
	params *SslOffPerCiphServStatsTableParams,
) *SslOffPerCiphServStatsTable {
	return &SslOffPerCiphServStatsTable{
		SslOffPerCiphServVirtServIndex:    sslOffPerCiphServVirtServIndex,
		SslOffPerCiphServVirtServiceIndex: sslOffPerCiphServVirtServiceIndex,
		SslOffPerCiphServVirtCiphIndex:    sslOffPerCiphServVirtCiphIndex,
		Params:                            params,
	}
}

func (c *SslOffPerCiphServStatsTable) Name() string {
	return "SslOffPerCiphServStatsTable"
}

func (c *SslOffPerCiphServStatsTable) GetParams() BeanType {
	return c.Params
}

func (c *SslOffPerCiphServStatsTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SslOffPerCiphServStatsTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SslOffPerCiphServVirtServIndex).IsZero() &&
		reflect.ValueOf(c.SslOffPerCiphServVirtServiceIndex).IsZero() &&
		reflect.ValueOf(c.SslOffPerCiphServVirtCiphIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SslOffPerCiphServVirtServIndex) + "/" + fmt.Sprint(c.SslOffPerCiphServVirtServiceIndex) + "/" + fmt.Sprint(c.SslOffPerCiphServVirtCiphIndex)
}

type SslOffPerCiphServStatsTableParams struct {
	// The virtual server name that identifies the service.
	VirtServIndex string `json:"VirtServIndex,omitempty"`
	// Virtual server service index.
	VirtServiceIndex int32 `json:"VirtServiceIndex,omitempty"`
	// Virtual server service cipher index.
	VirtCiphIndex int32 `json:"VirtCiphIndex,omitempty"`
	// Virtual server service cipher name.
	CiphName string `json:"CiphName,omitempty"`
	// Number of New SSL handshakes between Clients and AAS per second for this virtual service.
	NewhandShake int32 `json:"NewhandShake,omitempty"`
	// Virtual server service port number.
	VirtServPort int32 `json:"VirtServPort,omitempty"`
	// Number of Total SSL handshakes between Clients and AAS per second for this virtual service.
	NewhandShakeTotal int32 `json:"NewhandShakeTotal,omitempty"`
}

func (p SslOffPerCiphServStatsTableParams) iMABean() {}
