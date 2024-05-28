package beans

import (
	"fmt"
	"reflect"
)

// SslOffPerFilterStatsTable A table for SSL statistics per filter.
// Note:This mib is not supported for VX instance of Virtualization.
type SslOffPerFilterStatsTable struct {
	// Filter ID.
	SslOffPerFilterId int32
	Params            *SslOffPerFilterStatsTableParams
}

func NewSslOffPerFilterStatsTableList() *SslOffPerFilterStatsTable {
	return &SslOffPerFilterStatsTable{}
}

func NewSslOffPerFilterStatsTable(
	sslOffPerFilterId int32,
	params *SslOffPerFilterStatsTableParams,
) *SslOffPerFilterStatsTable {
	return &SslOffPerFilterStatsTable{
		SslOffPerFilterId: sslOffPerFilterId,
		Params:            params,
	}
}

func (c *SslOffPerFilterStatsTable) Name() string {
	return "SslOffPerFilterStatsTable"
}

func (c *SslOffPerFilterStatsTable) GetParams() BeanType {
	return c.Params
}

func (c *SslOffPerFilterStatsTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SslOffPerFilterStatsTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SslOffPerFilterId).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SslOffPerFilterId)
}

type SslOffPerFilterStatsTableParams struct {
	// Filter ID.
	Id int32 `json:"Id,omitempty"`
	// Number of New SSL handshakes between Clients and AAS per second for this filter.
	NewhandShake int32 `json:"NewhandShake,omitempty"`
	// Number of Reused SSL handshakes between Clients and AAS per second for this filter.
	ReusedhandShake int32 `json:"ReusedhandShake,omitempty"`
	// Number of existing SSL handshakes re-used by clients to communicate with AAS per second for filter.
	RateReusedhandShake int32 `json:"RateReusedhandShake,omitempty"`
	// Percentage of session using SSLv3 out of all session during measuring period for filter.
	PerSessUsingSSLv3 int32 `json:"PerSessUsingSSLv3,omitempty"`
	// Percentage of session using TLS1.0 out of all session during measuring period for filter.
	PerSessUsingTLS10 int32 `json:"PerSessUsingTLS10,omitempty"`
	// Percentage of session using TLS1.1 out of all session during measuring period for filter.
	PerSessUsingTLS11 int32 `json:"PerSessUsingTLS11,omitempty"`
	// Percentage of session using TLS1.2 out of all session during measuring period for filter.
	PerSessUsingTLS12 int32 `json:"PerSessUsingTLS12,omitempty"`
	// Number of rejected SSL handshakes per second for filter.
	RejectedHandShakes int32 `json:"RejectedHandShakes,omitempty"`
	// Number of HTTP redirect location headers updated from HTTP to HTTPS by AAS for filter.
	HttpToHTTPSRedir int32 `json:"HttpToHTTPSRedir,omitempty"`
	// Obsolete Mib! please use SslOffPerFilterCipherHandShakeStatsTable instead
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
	CertificateHostnameMismatch int32 `json:"CertificateHostnameMismatch,omitempty"`
	// Percentage of session using TLS1.3 out of all session during measuring period for filter.
	PerSessUsingTLS13 int32 `json:"PerSessUsingTLS13,omitempty"`
	// Number of Reused SSL handshakes between Clients and AAS per second for this filter (0-RTT).
	Filter0RTTReusedhandShake int32 `json:"Filter0RTTReusedhandShake,omitempty"`
	// Number of existing SSL handshakes re-used by clients to communicate with AAS per second for filter (0-RTT).
	Rate0RTTReusedhandShake int32 `json:"Rate0RTTReusedhandShake,omitempty"`
	// Number of rejected SSL handshakes per second for this filter (0-RTT).
	Filter0RTTRejectedhandShake int32 `json:"Filter0RTTRejectedhandShake,omitempty"`
	// Obsolete Mib! please use SslOffPerFilterRejHsReasonsStatsTable instead
	ByRejectedHandShakeReasons string `json:"ByRejectedHandShakeReasons,omitempty"`
	// Number of New SSL handshakes between Clients and AAS per second for this filter.
	BeNewhandShake int32 `json:"BeNewhandShake,omitempty"`
	// Number of Reused SSL handshakes between servers and AAS per second for this filter.
	BeReusedhandShake int32 `json:"BeReusedhandShake,omitempty"`
	// Number of existing SSL handshakes re-used by servers to communicate with AAS per second for filter.
	BeRateReusedhandShake int32 `json:"BeRateReusedhandShake,omitempty"`
	// Percentage of session using SSLv3 out of all session during measuring period for filter.
	BePerSessUsingSSLv3 int32 `json:"BePerSessUsingSSLv3,omitempty"`
	// Percentage of session using TLS1.0 out of all session during measuring period for filter.
	BePerSessUsingTLS10 int32 `json:"BePerSessUsingTLS10,omitempty"`
	// Percentage of session using TLS1.1 out of all session during measuring period for filter.
	BePerSessUsingTLS11 int32 `json:"BePerSessUsingTLS11,omitempty"`
	// Percentage of session using TLS1.2 out of all session during measuring period for filter.
	BePerSessUsingTLS12 int32 `json:"BePerSessUsingTLS12,omitempty"`
	// Number of rejected SSL handshakes per second for filter.
	BeRejectedHandShakes int32 `json:"BeRejectedHandShakes,omitempty"`
	// Number of HTTP redirect location headers updated from HTTP to HTTPS by AAS for filter.
	BeHttpToHTTPSRedir int32 `json:"BeHttpToHTTPSRedir,omitempty"`
	// Obsolete Mib! please use SslBeOffPerFilterCipherHandShakeStatsTable instead
	BeByCipherHandShake string `json:"BeByCipherHandShake,omitempty"`
	// Obsolete Mib!
	FeRejectedCertificates int32 `json:"FeRejectedCertificates,omitempty"`
	// Obsolete Mib!
	FeIgnoredCertificates int32 `json:"FeIgnoredCertificates,omitempty"`
	// Obsolete Mib!
	FeExpiredCertificates int32 `json:"FeExpiredCertificates,omitempty"`
	// Obsolete Mib!
	FeUntrustedCertificates int32 `json:"FeUntrustedCertificates,omitempty"`
	// Obsolete Mib!
	FeCertificateHostnameMismatch int32 `json:"FeCertificateHostnameMismatch,omitempty"`
	// Percentage of session using TLS1.3 out of all session during measuring period for filter.
	BePerSessUsingTLS13 int32 `json:"BePerSessUsingTLS13,omitempty"`
	// Number of New SSL handshakes between Clients and AAS since last reset for this filter.
	NewhandShakeTotal int32 `json:"NewhandShakeTotal,omitempty"`
	// Obsolete Mib! please use SslBeOffPerFilterRejHsReasonsStatsTable instead
	BeByRejectedHandShakeReasons string `json:"BeByRejectedHandShakeReasons,omitempty"`
	// Number of Reused SSL handshakes between Clients and AAS since last reset for this filter.
	ReusedhandShakeTotal int32 `json:"ReusedhandShakeTotal,omitempty"`
	// Number of existing SSL handshakes re-used by clients to communicate with AAS since last reset for filter.
	RateReusedhandShakeTotal int32 `json:"RateReusedhandShakeTotal,omitempty"`
	// Percentage of session using SSLv3 out of all session since last reset for filter.
	PerSessUsingSSLv3Total int32 `json:"PerSessUsingSSLv3Total,omitempty"`
	// Percentage of session using TLS1.0 out of all session since last reset for filter.
	PerSessUsingTLS10Total int32 `json:"PerSessUsingTLS10Total,omitempty"`
	// Percentage of session using TLS1.1 out of all session since last reset for filter.
	PerSessUsingTLS11Total int32 `json:"PerSessUsingTLS11Total,omitempty"`
	// Percentage of session using TLS1.2 out of all session since last reset for filter.
	PerSessUsingTLS12Total int32 `json:"PerSessUsingTLS12Total,omitempty"`
	// Number of rejected SSL handshakes since last reset for filter.
	RejectedHandShakesTotal int32 `json:"RejectedHandShakesTotal,omitempty"`
	// Number of HTTP redirect location headers updated from HTTP to HTTPS by AAS for filter.
	HttpToHTTPSRedirTotal int32 `json:"HttpToHTTPSRedirTotal,omitempty"`
	// Obsolete Mib! please use SslOffPerFilterCipherHandShakeStatsTable instead
	ByCipherHandShakeTotal string `json:"ByCipherHandShakeTotal,omitempty"`
	// Obsolete Mib!
	RejectedCertificatesTotal int32 `json:"RejectedCertificatesTotal,omitempty"`
	// Obsolete Mib!
	IgnoredCertificatesTotal int32 `json:"IgnoredCertificatesTotal,omitempty"`
	// Obsolete Mib!
	ExpiredCertificatesTotal int32 `json:"ExpiredCertificatesTotal,omitempty"`
	// Obsolete Mib!
	UntrustedCertificatesTotal int32 `json:"UntrustedCertificatesTotal,omitempty"`
	// Obsolete Mib!
	CertificateHostnameMismatchTotal int32 `json:"CertificateHostnameMismatchTotal,omitempty"`
	// Percentage of session using TLS1.3 out of all session since last reset for filter.
	PerSessUsingTLS13Total int32 `json:"PerSessUsingTLS13Total,omitempty"`
	// Number of New SSL handshakes between servers and AAS since last reset for this filter.
	BeNewhandShakeTotal int32 `json:"BeNewhandShakeTotal,omitempty"`
	// Number of Reused SSL handshakes between servers and AAS since last reset for this filter.
	BeReusedhandShakeTotal int32 `json:"BeReusedhandShakeTotal,omitempty"`
	// Obsolete Mib! please use SslOffPerFilterRejHsReasonsStatsTable instead
	ByRejectedHandShakeReasonsTotal string `json:"ByRejectedHandShakeReasonsTotal,omitempty"`
	// Number of existing SSL handshakes re-used by servers to communicate with AAS since last reset for filter.
	BeRateReusedhandShakeTotal int32 `json:"BeRateReusedhandShakeTotal,omitempty"`
	// Percentage of session using SSLv3 out of all session since last reset for filter.
	BePerSessUsingSSLv3Total int32 `json:"BePerSessUsingSSLv3Total,omitempty"`
	// Percentage of session using TLS1.0 out of all session since last reset for filter.
	BePerSessUsingTLS10Total int32 `json:"BePerSessUsingTLS10Total,omitempty"`
	// Percentage of session using TLS1.1 out of all session since last reset for filter.
	BePerSessUsingTLS11Total int32 `json:"BePerSessUsingTLS11Total,omitempty"`
	// Percentage of session using TLS1.2 out of all session since last reset for filter.
	BePerSessUsingTLS12Total int32 `json:"BePerSessUsingTLS12Total,omitempty"`
	// Number of rejected SSL handshakes since last reset for filter.
	BeRejectedHandShakesTotal int32 `json:"BeRejectedHandShakesTotal,omitempty"`
	// Number of HTTP redirect location headers updated from HTTP to HTTPS by AAS for filter.
	BeHttpToHTTPSRedirTotal int32 `json:"BeHttpToHTTPSRedirTotal,omitempty"`
	// Obsolete Mib! please use SslBeOffPerFilterCipherHandShakeStatsTable instead
	BeByCipherHandShakeTotal string `json:"BeByCipherHandShakeTotal,omitempty"`
	// Obsolete Mib!
	FeRejectedCertificatesTotal int32 `json:"FeRejectedCertificatesTotal,omitempty"`
	// Obsolete Mib!
	FeIgnoredCertificatesTotal int32 `json:"FeIgnoredCertificatesTotal,omitempty"`
	// Obsolete Mib!
	FeExpiredCertificatesTotal int32 `json:"FeExpiredCertificatesTotal,omitempty"`
	// Obsolete Mib!
	FeUntrustedCertificatesTotal int32 `json:"FeUntrustedCertificatesTotal,omitempty"`
	// Obsolete Mib!
	FeCertificateHostnameMismatchTotal int32 `json:"FeCertificateHostnameMismatchTotal,omitempty"`
	// Percentage of session using TLS1.3 out of all session since last reset for filter.
	BePerSessUsingTLS13Total int32 `json:"BePerSessUsingTLS13Total,omitempty"`
	// Obsolete Mib! please use SslBeOffPerFilterRejHsReasonsStatsTable instead
	BeByRejectedHandShakeReasonsTotal string `json:"BeByRejectedHandShakeReasonsTotal,omitempty"`
	// Number of existing SSL handshakes re-used by clients to communicate with AAS for filter (0-RTT) since last reset.
	Filter0RTTReusedhandShakeTotal int32 `json:"Filter0RTTReusedhandShakeTotal,omitempty"`
	// Percentage of SSL session re-using keys for filter (0-RTT) since last reset.
	Rate0RTTReusedhandShakeTotal int32 `json:"Rate0RTTReusedhandShakeTotal,omitempty"`
	// Number of rejected SSL handshakes for filter (0-RTT) since last reset.
	Filter0RTTRejectedhandShakeTotal int32 `json:"Filter0RTTRejectedhandShakeTotal,omitempty"`
	// Number of Session Id Reused SSL handshakes between Clients and AAS per second for this filter.
	SessionIdReusedhandShake int32 `json:"SessionIdReusedhandShake,omitempty"`
	// Number of existing SSL handshakes Session Id re-used by clients to communicate with AAS per second for filter.
	RateSessionIdReusedhandShake int32 `json:"RateSessionIdReusedhandShake,omitempty"`
	// Number of Session Id Reused SSL handshakes between servers and AAS per second for this filter.
	BeSessionIdReusedhandShake int32 `json:"BeSessionIdReusedhandShake,omitempty"`
	// Number of existing SSL handshakes Session Id re-used by servers to communicate with AAS per second for filter.
	BeRateSessionIdReusedhandShake int32 `json:"BeRateSessionIdReusedhandShake,omitempty"`
	// Number of Session Id Reused SSL handshakes between Clients and AAS since last reset for this filter.
	SessionIdReusedhandShakeTotal int32 `json:"SessionIdReusedhandShakeTotal,omitempty"`
	// Number of existing SSL handshakes Session Id re-used by clients to communicate with AAS since last reset for filter.
	RateSessionIdReusedhandShakeTotal int32 `json:"RateSessionIdReusedhandShakeTotal,omitempty"`
	// Number of Session Id Reused SSL handshakes between servers and AAS since last reset for this filter.
	BeSessionIdReusedhandShakeTotal int32 `json:"BeSessionIdReusedhandShakeTotal,omitempty"`
	// Number of existing SSL handshakes Session Id re-used by servers to communicate with AAS since last reset for filter.
	BeRateSessionIdReusedhandShakeTotal int32 `json:"BeRateSessionIdReusedhandShakeTotal,omitempty"`
	// Number of Ticket Reused SSL handshakes between Clients and AAS per second for this filter.
	TicketReusedhandShake int32 `json:"TicketReusedhandShake,omitempty"`
	// Number of existing SSL handshakes Ticket re-used by clients to communicate with AAS per second for filter.
	RateTicketReusedhandShake int32 `json:"RateTicketReusedhandShake,omitempty"`
	// Number of Ticket Reused SSL handshakes between servers and AAS per second for this filter.
	BeTicketReusedhandShake int32 `json:"BeTicketReusedhandShake,omitempty"`
	// Number of existing SSL handshakes Ticket re-used by servers to communicate with AAS per second for filter.
	BeRateTicketReusedhandShake int32 `json:"BeRateTicketReusedhandShake,omitempty"`
	// Number of Ticket Reused SSL handshakes between Clients and AAS since last reset for this filter.
	TicketReusedhandShakeTotal int32 `json:"TicketReusedhandShakeTotal,omitempty"`
	// Number of existing SSL handshakes Ticket re-used by clients to communicate with AAS since last reset for filter.
	RateTicketReusedhandShakeTotal int32 `json:"RateTicketReusedhandShakeTotal,omitempty"`
	// Number of Ticket Reused SSL handshakes between servers and AAS since last reset for this filter.
	BeTicketReusedhandShakeTotal int32 `json:"BeTicketReusedhandShakeTotal,omitempty"`
	// Number of existing SSL handshakes Ticket re-used by servers to communicate with AAS since last reset for filter.
	BeRateTicketReusedhandShakeTotal int32 `json:"BeRateTicketReusedhandShakeTotal,omitempty"`
}

func (p SslOffPerFilterStatsTableParams) iMABean() {}
