package beans

import (
	"fmt"
	"reflect"
)

// HttpEnhPerServStatsTable A table for HTTP statistics per virtual service.
// Note:This mib is not supported for VX instance of Virtualization.
type HttpEnhPerServStatsTable struct {
	// Virtual server Index.
	HttpEnhPerServStatsVirtServIndex string
	// Virtual server service index.
	HttpEnhPerServStatsVirtServiceIndex int32
	Params                              *HttpEnhPerServStatsTableParams
}

func NewHttpEnhPerServStatsTableList() *HttpEnhPerServStatsTable {
	return &HttpEnhPerServStatsTable{}
}

func NewHttpEnhPerServStatsTable(
	httpEnhPerServStatsVirtServIndex string,
	httpEnhPerServStatsVirtServiceIndex int32,
	params *HttpEnhPerServStatsTableParams,
) *HttpEnhPerServStatsTable {
	return &HttpEnhPerServStatsTable{
		HttpEnhPerServStatsVirtServIndex:    httpEnhPerServStatsVirtServIndex,
		HttpEnhPerServStatsVirtServiceIndex: httpEnhPerServStatsVirtServiceIndex,
		Params:                              params,
	}
}

func (c *HttpEnhPerServStatsTable) Name() string {
	return "HttpEnhPerServStatsTable"
}

func (c *HttpEnhPerServStatsTable) GetParams() BeanType {
	return c.Params
}

func (c *HttpEnhPerServStatsTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *HttpEnhPerServStatsTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.HttpEnhPerServStatsVirtServIndex).IsZero() &&
		reflect.ValueOf(c.HttpEnhPerServStatsVirtServiceIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.HttpEnhPerServStatsVirtServIndex) + "/" + fmt.Sprint(c.HttpEnhPerServStatsVirtServiceIndex)
}

type HttpEnhPerServStatsTableParams struct {
	// Virtual server Index.
	VirtServIndex string `json:"VirtServIndex,omitempty"`
	// Virtual server service index.
	VirtServiceIndex int32 `json:"VirtServiceIndex,omitempty"`
	// Virtual server service port number.
	VirtServPort int32 `json:"VirtServPort,omitempty"`
	// Number of clients sending Connection: Keep-Alive for virtual service.
	CliUseKeepAlive int32 `json:"CliUseKeepAlive,omitempty"`
	// Ratio of requests done using HTTP 1.0 and HTTP 1.1 during measuring period for virtual service.
	Http10VsHttp11Ratio int32 `json:"Http10VsHttp11Ratio,omitempty"`
	// Number of HTTP redirect location headers updated from HTTP to HTTPS by AAS for virtual service.
	HttpToHTTPSRedir int32 `json:"HttpToHTTPSRedir,omitempty"`
	// The average number of requests done over each client connection for virtual service.
	AvgNumReqPerConn int32 `json:"AvgNumReqPerConn,omitempty"`
	// Number of responses for which content size reported smaller than 1KB for virtual service.
	RespSmall1Kb int32 `json:"RespSmall1Kb,omitempty"`
	// Number of responses for which content size reported  between 1KB and 10KB for virtual service.
	Resp1KbTo10Kb int32 `json:"Resp1KbTo10Kb,omitempty"`
	// Number of responses for which content size reported between 11KB and 50KB for virtual service.
	Resp11KbTo50Kb int32 `json:"Resp11KbTo50Kb,omitempty"`
	// Number of responses for which content size reported between 51KB and 100KB for virtual service.
	Resp51KbTo100Kb int32 `json:"Resp51KbTo100Kb,omitempty"`
	// Number of responses for which content size reported larger than 100KB for virtual service.
	RespLarger100Kb int32 `json:"RespLarger100Kb,omitempty"`
	// Number of clients requests from AAS done in the measuring period for virtual service.
	ReqCliToAas int32 `json:"ReqCliToAas,omitempty"`
	// Number of AAS requests from servers done in the measuring period for virtual service.
	ReqAasToSer int32 `json:"ReqAasToSer,omitempty"`
	// Number of servers responses to AAS in the measuring period for virtual service.
	RespSerToAas int32 `json:"RespSerToAas,omitempty"`
	// Number of AAS responses to clients in the measuring period for virtual service.
	RespAasToCli int32 `json:"RespAasToCli,omitempty"`
	// Transactions rate for virtual service.
	TransRate int32 `json:"TransRate,omitempty"`
	// HTTP2 connection count.
	PerServStatsHttp20ConnectionCount int32 `json:"PerServStatsHttp20ConnectionCount,omitempty"`
	// HTTP 1.1 connection count.
	PerServStatsHttp11ConnectionCount int32 `json:"PerServStatsHttp11ConnectionCount,omitempty"`
	// HTTP 1.0 connection count.
	PerServStatsHttp10ConnectionCount int32 `json:"PerServStatsHttp10ConnectionCount,omitempty"`
	// HTTP 2.0 connection count peak.
	PerServStatsHttp20ConnectionPeak int32 `json:"PerServStatsHttp20ConnectionPeak,omitempty"`
	// HTTP 1.1 connection count peak.
	PerServStatsHttp11ConnectionPeak int32 `json:"PerServStatsHttp11ConnectionPeak,omitempty"`
	// HTTP 1.0 connection count peak.
	PerServStatsHttp10ConnectionPeak int32 `json:"PerServStatsHttp10ConnectionPeak,omitempty"`
	// HTTP 2.0 request count.
	PerServStatsHttp20RequestCount int32 `json:"PerServStatsHttp20RequestCount,omitempty"`
	// HTTP 1.1 request count.
	PerServStatsHttp11RequestCount int32 `json:"PerServStatsHttp11RequestCount,omitempty"`
	// HTTP 1.0 request count.
	PerServStatsHttp10RequestCount int32 `json:"PerServStatsHttp10RequestCount,omitempty"`
	// HTTP2 backend proxy connections.
	PerServStatsBackendProxyConnections int32 `json:"PerServStatsBackendProxyConnections,omitempty"`
	// HTTP2 client streams.
	PerServStatsClientStreams int32 `json:"PerServStatsClientStreams,omitempty"`
	// HTTP2 PUSH streams.
	PerServStatsPushStreams int32 `json:"PerServStatsPushStreams,omitempty"`
	// HTTP2 cancelled PUSH streams.
	PerServStatsCanceledPushStreams int32 `json:"PerServStatsCanceledPushStreams,omitempty"`
	// HTTP2 average connection duration (seconds).
	PerServStatsConnectionDurationAvgStr string `json:"PerServStatsConnectionDurationAvgStr,omitempty"`
	// HTTP2 headers request compression ratio.
	PerServStatsHeadersRequestCompRatio int32 `json:"PerServStatsHeadersRequestCompRatio,omitempty"`
	// HTTP2 headers response compression ratio.
	PerServStatsHeadersResponseCompRatio int32 `json:"PerServStatsHeadersResponseCompRatio,omitempty"`
	// HTTP2 big headers count.
	PerServStatsBigHeaders int32 `json:"PerServStatsBigHeaders,omitempty"`
	// HTTP2 average eviction bytes.
	PerServStatsAvgEvictionBytes int32 `json:"PerServStatsAvgEvictionBytes,omitempty"`
	// HTTP2 average HPACK table size.
	PerServStatsAvgHpackTableSize int32 `json:"PerServStatsAvgHpackTableSize,omitempty"`
	// HTTP2 Peak of backend proxy connections.
	PerServStatsPeakBackendProxyConnections int32 `json:"PerServStatsPeakBackendProxyConnections,omitempty"`
	// HTTP2 Peak of client streams.
	PerServStatsPeakClientStreams int32 `json:"PerServStatsPeakClientStreams,omitempty"`
	// HTTP2 Peak of PUSH streams.
	PerServStatsPeakPushStreams int32 `json:"PerServStatsPeakPushStreams,omitempty"`
	// HTTP2 Peak of cancelled PUSH streams.
	PerServStatsPeakCanceledPushStreams int32 `json:"PerServStatsPeakCanceledPushStreams,omitempty"`
	// HTTP2 Peak of connection duration average (seconds).
	PerServStatsPeakConnectionDurationAvgStr string `json:"PerServStatsPeakConnectionDurationAvgStr,omitempty"`
	// HTTP2 Peak of headers request compression ratio.
	PerServStatsPeakHeadersRequestCompRatio int32 `json:"PerServStatsPeakHeadersRequestCompRatio,omitempty"`
	// HTTP2 Peak of headers response compression ratio.
	PerServStatsPeakHeadersResponseCompRatio int32 `json:"PerServStatsPeakHeadersResponseCompRatio,omitempty"`
	// HTTP2 Peak of number of big headers.
	PerServStatsPeakBigHeaders int32 `json:"PerServStatsPeakBigHeaders,omitempty"`
	// HTTP2 Peak of average eviction in bytes.
	PerServStatsPeakAvgEvictionBytes int32 `json:"PerServStatsPeakAvgEvictionBytes,omitempty"`
	// HTTP2 Peak of average HPACK table size in Kb.
	PerServStatsPeakAvgHpackTableSize int32 `json:"PerServStatsPeakAvgHpackTableSize,omitempty"`
}

func (p HttpEnhPerServStatsTableParams) iMABean() {}
