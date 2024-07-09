package beans

import (
	"fmt"
	"reflect"
)

// SlbStatSbSslRealTable Statistic table for Sideband SSL and total real servers.
// Note:This mib is not supported for VX instance of Virtualization.
type SlbStatSbSslRealTable struct {
	// Sideband Index.
	SlbStatSbSslRealSidebandID string
	Params                     *SlbStatSbSslRealTableParams
}

func NewSlbStatSbSslRealTableList() *SlbStatSbSslRealTable {
	return &SlbStatSbSslRealTable{}
}

func NewSlbStatSbSslRealTable(
	slbStatSbSslRealSidebandID string,
	params *SlbStatSbSslRealTableParams,
) *SlbStatSbSslRealTable {
	return &SlbStatSbSslRealTable{
		SlbStatSbSslRealSidebandID: slbStatSbSslRealSidebandID,
		Params:                     params,
	}
}

func (c *SlbStatSbSslRealTable) Name() string {
	return "SlbStatSbSslRealTable"
}

func (c *SlbStatSbSslRealTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbStatSbSslRealTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbStatSbSslRealTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbStatSbSslRealSidebandID).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbStatSbSslRealSidebandID)
}

type SlbStatSbSslRealTableParams struct {
	// Sideband Index.
	SidebandID string `json:"SidebandID,omitempty"`
	// The number of sessions that are currently handled by the
	// sideband.
	CurrSessions uint32 `json:"CurrSessions,omitempty"`
	// The total number of sessions that are handled by the sideband.
	TotalSessions uint32 `json:"TotalSessions,omitempty"`
	// The highest sessions that have been handled by the sideband.
	HighestSessions uint32 `json:"HighestSessions,omitempty"`
	// The total octets received and transmitted by the sideband.
	Octets string `json:"Octets,omitempty"`
	// sideband real server connection per seconds.
	SessionsPerSec uint64 `json:"SessionsPerSec,omitempty"`
	// sideband real server throughput.
	OctetsPerSec uint64 `json:"OctetsPerSec,omitempty"`
	// Number of New SSL handshakes between servers and AAS per second for this sideband.
	NewhandShake uint32 `json:"NewhandShake,omitempty"`
	// Number of existing SSL handshakes re-used by servers to communicate with AAS per second for sideband.
	ReusedhandShake uint32 `json:"ReusedhandShake,omitempty"`
	// Percentage of SSL session re-using keys for sideband.
	PercReusedhandShake uint32 `json:"PercReusedhandShake,omitempty"`
	// Percentage of session using SSLv2 out of all session during measuring period for sideband.
	UsingSSLv2 uint32 `json:"UsingSSLv2,omitempty"`
	// Percentage of session using SSLv3 out of all session during measuring period for sideband.
	UsingSSLv3 uint32 `json:"UsingSSLv3,omitempty"`
	// Percentage of session using TLS1.0 out of all session during measuring period for sideband.
	UsingTLS uint32 `json:"UsingTLS,omitempty"`
	// Percentage of session using TLS1.1 out of all session during measuring period for sideband.
	UsingTLS11 uint32 `json:"UsingTLS11,omitempty"`
	// Percentage of session using TLS1.2 out of all session during measuring period for sideband.
	UsingTLS12 uint32 `json:"UsingTLS12,omitempty"`
	// Percentage of session using TLS1.3 out of all session during measuring period for sideband.
	UsingTLS13 uint32 `json:"UsingTLS13,omitempty"`
	// Number of HTTP redirect location headers updated from HTTP to HTTPS by AAS for filter.
	HttpToHTTPSRedir uint32 `json:"HttpToHTTPSRedir,omitempty"`
	// Number of New SSL handshakes between servers and AAS since last reset for this sideband.
	NewhandShakeTotal uint32 `json:"NewhandShakeTotal,omitempty"`
	// Number of existing SSL handshakes re-used by servers to communicate with AAS since last reset for sideband.
	ReusedhandShakeTotal uint32 `json:"ReusedhandShakeTotal,omitempty"`
	// Percentage of SSL session re-using keys for sideband.
	PercReusedhandShakeTotal uint32 `json:"PercReusedhandShakeTotal,omitempty"`
	// Percentage of session using SSLv2 out of all session since last reset for sideband.
	UsingSSLv2Total uint32 `json:"UsingSSLv2Total,omitempty"`
	// Percentage of session using SSLv3 out of all session since last reset for sideband.
	UsingSSLv3Total uint32 `json:"UsingSSLv3Total,omitempty"`
	// Percentage of session using TLS1.0 out of all session since last reset for sideband.
	UsingTLSTotal uint32 `json:"UsingTLSTotal,omitempty"`
	// Percentage of session using TLS1.1 out of all session since last reset for sideband.
	UsingTLS11Total uint32 `json:"UsingTLS11Total,omitempty"`
	// Percentage of session using TLS1.2 out of all session since last reset for sideband.
	UsingTLS12Total uint32 `json:"UsingTLS12Total,omitempty"`
	// Percentage of session using TLS1.3 out of all session since last reset for sideband.
	UsingTLS13Total uint32 `json:"UsingTLS13Total,omitempty"`
	// Number of HTTP redirect location headers updated from HTTP to HTTPS by AAS for filter.
	HttpToHTTPSRedirTotal uint32 `json:"HttpToHTTPSRedirTotal,omitempty"`
	// Number of existing SSL handshakes Session Id re-used by servers to communicate with AAS per second for sideband.
	SessionIdReusedhandShake uint32 `json:"SessionIdReusedhandShake,omitempty"`
	// Percentage of SSL session Id re-using keys for sideband.
	PercSessionIdReusedhandShake uint32 `json:"PercSessionIdReusedhandShake,omitempty"`
	// Number of existing SSL handshakes Session Id re-used by servers to communicate with AAS since last reset for sideband.
	SessionIdReusedhandShakeTotal uint32 `json:"SessionIdReusedhandShakeTotal,omitempty"`
	// Percentage of SSL session Id re-using keys for sideband.
	PercSessionIdReusedhandShakeTotal uint32 `json:"PercSessionIdReusedhandShakeTotal,omitempty"`
	// Number of existing SSL handshakes Ticket re-used by servers to communicate with AAS per second for sideband.
	TicketReusedhandShake uint32 `json:"TicketReusedhandShake,omitempty"`
	// Percentage of SSL session Ticket re-using keys for sideband.
	PercTicketReusedhandShake uint32 `json:"PercTicketReusedhandShake,omitempty"`
	// Number of existing SSL handshakes Ticket re-used by servers to communicate with AAS since last reset for sideband.
	TicketReusedhandShakeTotal uint32 `json:"TicketReusedhandShakeTotal,omitempty"`
	// Percentage of SSL session Ticket re-using keys for sideband.
	PercTicketReusedhandShakeTotal uint32 `json:"PercTicketReusedhandShakeTotal,omitempty"`
	// The average ServerRTT in ms during the past collection interval.
	ServerRtt string `json:"ServerRtt,omitempty"`
}

func (p SlbStatSbSslRealTableParams) iMABean() {}
