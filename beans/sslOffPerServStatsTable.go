package beans

import (
	"fmt"
	"reflect"
)

// SslOffPerServStatsTable A table for SSL statistics per virtual service.
// Note:This mib is not supported for VX instance of Virtualization.
type SslOffPerServStatsTable struct {
	// Virtual server number.
	SslOffPerServVirtServIndex int32
	// Virtual server service index.
	SslOffPerServVirtServiceIndex int32
	Params                        *SslOffPerServStatsTableParams
}

func NewSslOffPerServStatsTableList() *SslOffPerServStatsTable {
	return &SslOffPerServStatsTable{}
}

func NewSslOffPerServStatsTable(
	sslOffPerServVirtServIndex int32,
	sslOffPerServVirtServiceIndex int32,
	params *SslOffPerServStatsTableParams,
) *SslOffPerServStatsTable {
	return &SslOffPerServStatsTable{
		SslOffPerServVirtServIndex:    sslOffPerServVirtServIndex,
		SslOffPerServVirtServiceIndex: sslOffPerServVirtServiceIndex,
		Params:                        params,
	}
}

func (c *SslOffPerServStatsTable) Name() string {
	return "SslOffPerServStatsTable"
}

func (c *SslOffPerServStatsTable) GetParams() BeanType {
	return c.Params
}

func (c *SslOffPerServStatsTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SslOffPerServStatsTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SslOffPerServVirtServIndex).IsZero() &&
		reflect.ValueOf(c.SslOffPerServVirtServiceIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SslOffPerServVirtServIndex) + "/" + fmt.Sprint(c.SslOffPerServVirtServiceIndex)
}

type SslOffPerServStatsTableParams struct {
	// Virtual server number.
	VirtServIndex int32 `json:"VirtServIndex,omitempty"`
	// Virtual server service index.
	VirtServiceIndex int32 `json:"VirtServiceIndex,omitempty"`
	// Virtual server service port number.
	VirtServPort int32 `json:"VirtServPort,omitempty"`
	// SSL policy identifier associated with the virtual service.
	SslPolId string `json:"SslPolId,omitempty"`
	// Number of New SSL handshakes between Clients and AAS per second for this virtual service.
	NewhandShake int32 `json:"NewhandShake,omitempty"`
	// Number of existing SSL handshakes re-used by clients to communicate with AAS per second for virtual service.
	ReusedhandShake int32 `json:"ReusedhandShake,omitempty"`
	// Percentage of SSL session re-using keys for virtual service.
	PercReusedhandShake int32 `json:"PercReusedhandShake,omitempty"`
	// Percentage of session using SSLv2 out of all session during measuring period for virtual service.
	SessUsingSSLv2 int32 `json:"SessUsingSSLv2,omitempty"`
	// Percentage of session using SSLv3 out of all session during measuring period for virtual service.
	SessUsingSSLv3 int32 `json:"SessUsingSSLv3,omitempty"`
	// Percentage of session using TLS1.0 out of all session during measuring period for virtual service.
	SessUsingTLS int32 `json:"SessUsingTLS,omitempty"`
	// Percentage of session using TLS1.1 out of all session during measuring period for virtual service.
	SessUsingTLS11 int32 `json:"SessUsingTLS11,omitempty"`
	// Percentage of session using TLS1.2 out of all session during measuring period for virtual service.
	SessUsingTLS12 int32 `json:"SessUsingTLS12,omitempty"`
	// Number of rejected SSL handshakes per second for virtual service.
	RejectedhandShake int32 `json:"RejectedhandShake,omitempty"`
	// Session distribution by cipher for virtual service (in second)
	ByCipherHandShake string `json:"ByCipherHandShake,omitempty"`
	// Obsolete Mib!
	RejectedCertificates int32 `json:"RejectedCertificates,omitempty"`
	// Obsolete Mib!
	IgnoredCertificates int32 `json:"IgnoredCertificates,omitempty"`
	// Obsolete Mib!
	ExpiredCertificates int32 `json:"ExpiredCertificates,omitempty"`
	// Obsolete Mib!
	UntrustedCertificates int32 `json:"UntrustedCertificates,omitempty"`
	// Obsolete Mib!
	CertificateHostnameMismatches int32 `json:"CertificateHostnameMismatches,omitempty"`
	// Number of existing SSL handshakes re-used by clients to communicate with AAS per second for virtual service (0-RTT).
	Serv0RTTReusedhandShake int32 `json:"Serv0RTTReusedhandShake,omitempty"`
	// Percentage of SSL session re-using keys for virtual service (0-RTT).
	Perc0RTTReusedhandShake int32 `json:"Perc0RTTReusedhandShake,omitempty"`
}

func (p SslOffPerServStatsTableParams) iMABean() {}
