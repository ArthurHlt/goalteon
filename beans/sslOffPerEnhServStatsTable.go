package beans

import (
	"fmt"
	"reflect"
)

// SslOffPerEnhServStatsTable A table for SSL statistics per virtual service.
// Note:This mib is not supported for VX instance of Virtualization.
type SslOffPerEnhServStatsTable struct {
	// Virtual server number in alphanumeric.
	SslOffPerEnhServVirtServIndex string
	// Virtual server service index.
	SslOffPerEnhServVirtServiceIndex int32
	Params                           *SslOffPerEnhServStatsTableParams
}

func NewSslOffPerEnhServStatsTableList() *SslOffPerEnhServStatsTable {
	return &SslOffPerEnhServStatsTable{}
}

func NewSslOffPerEnhServStatsTable(
	sslOffPerEnhServVirtServIndex string,
	sslOffPerEnhServVirtServiceIndex int32,
	params *SslOffPerEnhServStatsTableParams,
) *SslOffPerEnhServStatsTable {
	return &SslOffPerEnhServStatsTable{
		SslOffPerEnhServVirtServIndex:    sslOffPerEnhServVirtServIndex,
		SslOffPerEnhServVirtServiceIndex: sslOffPerEnhServVirtServiceIndex,
		Params:                           params,
	}
}

func (c *SslOffPerEnhServStatsTable) Name() string {
	return "SslOffPerEnhServStatsTable"
}

func (c *SslOffPerEnhServStatsTable) GetParams() BeanType {
	return c.Params
}

func (c *SslOffPerEnhServStatsTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SslOffPerEnhServStatsTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SslOffPerEnhServVirtServIndex).IsZero() &&
		reflect.ValueOf(c.SslOffPerEnhServVirtServiceIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SslOffPerEnhServVirtServIndex) + "/" + fmt.Sprint(c.SslOffPerEnhServVirtServiceIndex)
}

type SslOffPerEnhServStatsTableParams struct {
	// Virtual server number in alphanumeric.
	VirtServIndex string `json:"VirtServIndex,omitempty"`
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
	RejectedHandShake int32 `json:"RejectedHandShake,omitempty"`
	// Obsolete Mib! please use SslOffPerCiphServStatsTable instead
	ByCipherHandShake string `json:"ByCipherHandShake,omitempty"`
	// Number of rejected SSL certificates per second for virtual service.
	RejectedCertificates int32 `json:"RejectedCertificates,omitempty"`
	// Number of ignored SSL certificates per second for virtual service.
	IgnoredCertificates int32 `json:"IgnoredCertificates,omitempty"`
	// Number of expired SSL certificates per second for virtual service.
	ExpiredCertificates int32 `json:"ExpiredCertificates,omitempty"`
	// Number of untrusted SSL certificates per second for virtual service.
	UntrustedCertificates int32 `json:"UntrustedCertificates,omitempty"`
	// Number of SSL certificate hostname mismatches per second for virtual service.
	CertificateHostnameMismatches int32 `json:"CertificateHostnameMismatches,omitempty"`
	// Percentage of session using TLS1.3 out of all session during measuring period for virtual service.
	SessUsingTLS13 int32 `json:"SessUsingTLS13,omitempty"`
	// Number of existing SSL handshakes re-used by clients to communicate with AAS per second for virtual service (0-RTT).
	Serv0RTTReusedhandShake int32 `json:"Serv0RTTReusedhandShake,omitempty"`
	// Percentage of SSL session re-using keys for virtual service (0-RTT).
	Perc0RTTReusedhandShake int32 `json:"Perc0RTTReusedhandShake,omitempty"`
	// Number of rejected SSL handshakes per second for virtual service (0-RTT).
	Serv0RTTRejectedhandShake int32 `json:"Serv0RTTRejectedhandShake,omitempty"`
	// Obsolete Mib! please use SslOffPerServRejHsReasonsStatsTable instead
	ByRejectedHandShakeReasons string `json:"ByRejectedHandShakeReasons,omitempty"`
	// Number of New SSL handshakes between servers and AAS per second for this virtual service.
	BeNewhandShake int32 `json:"BeNewhandShake,omitempty"`
	// Number of existing SSL handshakes re-used by servers to communicate with AAS per second for virtual service.
	BeReusedhandShake int32 `json:"BeReusedhandShake,omitempty"`
	// Percentage of SSL session re-using keys for virtual service.
	BePercReusedhandShake int32 `json:"BePercReusedhandShake,omitempty"`
	// Percentage of session using SSLv2 out of all session during measuring period for virtual service.
	BeSessUsingSSLv2 int32 `json:"BeSessUsingSSLv2,omitempty"`
	// Percentage of session using SSLv3 out of all session during measuring period for virtual service.
	BeSessUsingSSLv3 int32 `json:"BeSessUsingSSLv3,omitempty"`
	// Percentage of session using TLS1.0 out of all session during measuring period for virtual service.
	BeSessUsingTLS int32 `json:"BeSessUsingTLS,omitempty"`
	// Percentage of session using TLS1.1 out of all session during measuring period for virtual service.
	BeSessUsingTLS11 int32 `json:"BeSessUsingTLS11,omitempty"`
	// Percentage of session using TLS1.2 out of all session during measuring period for virtual service.
	BeSessUsingTLS12 int32 `json:"BeSessUsingTLS12,omitempty"`
	// Number of rejected SSL handshakes per second for virtual service.
	BeRejectedHandShake int32 `json:"BeRejectedHandShake,omitempty"`
	// Obsolete Mib! please use SslBeOffPerCiphServStatsTable instead
	BeByCipherHandShake string `json:"BeByCipherHandShake,omitempty"`
	// Number of rejected SSL certificates per second for virtual service.
	FeRejectedCertificates int32 `json:"FeRejectedCertificates,omitempty"`
	// Number of ignored SSL certificates per second for virtual service.
	FeIgnoredCertificates int32 `json:"FeIgnoredCertificates,omitempty"`
	// Number of expired SSL certificates per second for virtual service.
	FeExpiredCertificates int32 `json:"FeExpiredCertificates,omitempty"`
	// Number of untrusted SSL certificates per second for virtual service.
	FeUntrustedCertificates int32 `json:"FeUntrustedCertificates,omitempty"`
	// Number of SSL certificate hostname mismatches per second for virtual service.
	FeCertificateHostnameMismatches int32 `json:"FeCertificateHostnameMismatches,omitempty"`
	// Percentage of session using TLS1.3 out of all session during measuring period for virtual service.
	BeSessUsingTLS13 int32 `json:"BeSessUsingTLS13,omitempty"`
	// Number of New SSL handshakes between Clients and AAS since last reset for this virtual service.
	NewhandShakeTotal int32 `json:"NewhandShakeTotal,omitempty"`
	// Obsolete Mib! please use SslBeOffPerServRejHsReasonsStatsTable instead
	BeByRejectedHandShakeReasons string `json:"BeByRejectedHandShakeReasons,omitempty"`
	// Number of existing SSL handshakes re-used by clients to communicate with AAS since last reset for virtual service.
	ReusedhandShakeTotal int32 `json:"ReusedhandShakeTotal,omitempty"`
	// Percentage of SSL session re-using keys for virtual service.
	PercReusedhandShakeTotal int32 `json:"PercReusedhandShakeTotal,omitempty"`
	// Percentage of session using SSLv2 out of all session since last reset for virtual service.
	SessUsingSSLv2Total int32 `json:"SessUsingSSLv2Total,omitempty"`
	// Percentage of session using SSLv3 out of all session since last reset for virtual service.
	SessUsingSSLv3Total int32 `json:"SessUsingSSLv3Total,omitempty"`
	// Percentage of session using TLS1.0 out of all session since last reset for virtual service.
	SessUsingTLSTotal int32 `json:"SessUsingTLSTotal,omitempty"`
	// Percentage of session using TLS1.1 out of all session since last reset for virtual service.
	SessUsingTLS11Total int32 `json:"SessUsingTLS11Total,omitempty"`
	// Percentage of session using TLS1.2 out of all session since last reset for virtual service.
	SessUsingTLS12Total int32 `json:"SessUsingTLS12Total,omitempty"`
	// Number of rejected SSL handshakes since last reset for virtual service.
	RejectedHandShakeTotal int32 `json:"RejectedHandShakeTotal,omitempty"`
	// Obsolete Mib! please use SslOffPerCiphServStatsTable instead
	ByCipherHandShakeTotal string `json:"ByCipherHandShakeTotal,omitempty"`
	// Number of rejected SSL certificates since last reset for virtual service.
	RejectedCertificatesTotal int32 `json:"RejectedCertificatesTotal,omitempty"`
	// Number of ignored SSL certificates per second for virtual service.
	IgnoredCertificatesTotal int32 `json:"IgnoredCertificatesTotal,omitempty"`
	// Number of expired SSL certificates since last reset for virtual service.
	ExpiredCertificatesTotal int32 `json:"ExpiredCertificatesTotal,omitempty"`
	// Number of untrusted SSL certificates since last reset for virtual service.
	UntrustedCertificatesTotal int32 `json:"UntrustedCertificatesTotal,omitempty"`
	// Number of SSL certificate hostname mismatches since last reset for virtual service.
	CertificateHostnameMismatchesTotal int32 `json:"CertificateHostnameMismatchesTotal,omitempty"`
	// Percentage of session using TLS1.3 out of all session since last reset for virtual service.
	SessUsingTLS13Total int32 `json:"SessUsingTLS13Total,omitempty"`
	// Number of New SSL handshakes between servers and AAS since last reset for this virtual service.
	BeNewhandShakeTotal int32 `json:"BeNewhandShakeTotal,omitempty"`
	// Number of existing SSL handshakes re-used by servers to communicate with AAS since last reset for virtual service.
	BeReusedhandShakeTotal int32 `json:"BeReusedhandShakeTotal,omitempty"`
	// Obsolete Mib! please use SslOffPerServRejHsReasonsStatsTable instead
	ByRejectedHandShakeReasonsTotal string `json:"ByRejectedHandShakeReasonsTotal,omitempty"`
	// Percentage of SSL session re-using keys for virtual service.
	BePercReusedhandShakeTotal int32 `json:"BePercReusedhandShakeTotal,omitempty"`
	// Percentage of session using SSLv2 out of all session since last reset for virtual service.
	BeSessUsingSSLv2Total int32 `json:"BeSessUsingSSLv2Total,omitempty"`
	// Percentage of session using SSLv3 out of all session since last reset for virtual service.
	BeSessUsingSSLv3Total int32 `json:"BeSessUsingSSLv3Total,omitempty"`
	// Percentage of session using TLS1.0 out of all session since last reset for virtual service.
	BeSessUsingTLSTotal int32 `json:"BeSessUsingTLSTotal,omitempty"`
	// Percentage of session using TLS1.1 out of all session since last reset for virtual service.
	BeSessUsingTLS11Total int32 `json:"BeSessUsingTLS11Total,omitempty"`
	// Percentage of session using TLS1.2 out of all session since last reset for virtual service.
	BeSessUsingTLS12Total int32 `json:"BeSessUsingTLS12Total,omitempty"`
	// Number of rejected SSL handshakes since last reset for virtual service.
	BeRejectedHandShakeTotal int32 `json:"BeRejectedHandShakeTotal,omitempty"`
	// Obsolete Mib! please use SslBeOffPerCiphServStatsTable instead
	BeByCipherHandShakeTotal string `json:"BeByCipherHandShakeTotal,omitempty"`
	// Number of rejected SSL certificates since last reset for virtual service.
	FeRejectedCertificatesTotal int32 `json:"FeRejectedCertificatesTotal,omitempty"`
	// Number of ignored SSL certificates since last reset for virtual service.
	FeIgnoredCertificatesTotal int32 `json:"FeIgnoredCertificatesTotal,omitempty"`
	// Number of expired SSL certificates since last reset for virtual service.
	FeExpiredCertificatesTotal int32 `json:"FeExpiredCertificatesTotal,omitempty"`
	// Number of untrusted SSL certificates since last reset for virtual service.
	FeUntrustedCertificatesTotal int32 `json:"FeUntrustedCertificatesTotal,omitempty"`
	// Number of SSL certificate hostname mismatches since last reset for virtual service.
	FeCertificateHostnameMismatchesTotal int32 `json:"FeCertificateHostnameMismatchesTotal,omitempty"`
	// Percentage of session using TLS1.3 out of all session since last reset for virtual service.
	BeSessUsingTLS13Total int32 `json:"BeSessUsingTLS13Total,omitempty"`
	// Obsolete Mib! please use SslBeOffPerServRejHsReasonsStatsTable instead
	BeByRejectedHandShakeReasonsTotal string `json:"BeByRejectedHandShakeReasonsTotal,omitempty"`
	// Number of existing SSL handshakes re-used by clients to communicate with AAS for virtual service (0-RTT) since last reset.
	Serv0RTTReusedhandShakeTotal int32 `json:"Serv0RTTReusedhandShakeTotal,omitempty"`
	// Percentage of SSL session re-using keys for virtual service (0-RTT) since last reset.
	Perc0RTTReusedhandShakeTotal int32 `json:"Perc0RTTReusedhandShakeTotal,omitempty"`
	// Number of rejected SSL handshakes for virtual service (0-RTT) since last reset.
	Serv0RTTRejectedhandShakeTotal int32 `json:"Serv0RTTRejectedhandShakeTotal,omitempty"`
	// Number of existing SSL handshakes Session Id re-used by clients to communicate with AAS per second for virtual service.
	SessionIdReusedhandShake int32 `json:"SessionIdReusedhandShake,omitempty"`
	// Percentage of SSL session Id re-using keys for virtual service.
	PercSessionIdReusedhandShake int32 `json:"PercSessionIdReusedhandShake,omitempty"`
	// Number of existing SSL handshakes Session Id re-used by servers to communicate with AAS per second for virtual service.
	BeSessionIdReusedhandShake int32 `json:"BeSessionIdReusedhandShake,omitempty"`
	// Percentage of SSL session Id re-using keys for virtual service.
	BePercSessionIdReusedhandShake int32 `json:"BePercSessionIdReusedhandShake,omitempty"`
	// Number of existing SSL handshakes Session Id re-used by clients to communicate with AAS since last reset for virtual service.
	SessionIdReusedhandShakeTotal int32 `json:"SessionIdReusedhandShakeTotal,omitempty"`
	// Percentage of SSL session Id re-using keys for virtual service.
	PercSessionIdReusedhandShakeTotal int32 `json:"PercSessionIdReusedhandShakeTotal,omitempty"`
	// Number of existing SSL handshakes Session Id re-used by servers to communicate with AAS since last reset for virtual service.
	BeSessionIdReusedhandShakeTotal int32 `json:"BeSessionIdReusedhandShakeTotal,omitempty"`
	// Percentage of SSL session Id re-using keys for virtual service.
	BePercSessionIdReusedhandShakeTotal int32 `json:"BePercSessionIdReusedhandShakeTotal,omitempty"`
	// Number of existing SSL handshakes Ticket re-used by clients to communicate with AAS per second for virtual service.
	TicketReusedhandShake int32 `json:"TicketReusedhandShake,omitempty"`
	// Percentage of SSL session Ticket re-using keys for virtual service.
	PercTicketReusedhandShake int32 `json:"PercTicketReusedhandShake,omitempty"`
	// Number of existing SSL handshakes Ticket re-used by servers to communicate with AAS per second for virtual service.
	BeTicketReusedhandShake int32 `json:"BeTicketReusedhandShake,omitempty"`
	// Percentage of SSL session Ticket re-using keys for virtual service.
	BePercTicketReusedhandShake int32 `json:"BePercTicketReusedhandShake,omitempty"`
	// Number of existing SSL handshakes Ticket re-used by clients to communicate with AAS since last reset for virtual service.
	TicketReusedhandShakeTotal int32 `json:"TicketReusedhandShakeTotal,omitempty"`
	// Percentage of SSL session Ticket re-using keys for virtual service.
	PercTicketReusedhandShakeTotal int32 `json:"PercTicketReusedhandShakeTotal,omitempty"`
	// Number of existing SSL handshakes Ticket re-used by servers to communicate with AAS since last reset for virtual service.
	BeTicketReusedhandShakeTotal int32 `json:"BeTicketReusedhandShakeTotal,omitempty"`
	// Percentage of SSL session Ticket re-using keys for virtual service.
	BePercTicketReusedhandShakeTotal int32 `json:"BePercTicketReusedhandShakeTotal,omitempty"`
}

func (p SslOffPerEnhServStatsTableParams) iMABean() {}
