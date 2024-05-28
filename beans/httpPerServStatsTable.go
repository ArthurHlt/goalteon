package beans

import (
	"fmt"
	"reflect"
)

// HttpPerServStatsTable A table for HTTP statistics per virtual service.
// Note:This mib is not supported for VX instance of Virtualization.
type HttpPerServStatsTable struct {
	// Virtual server number.
	HttpPerServStatsVirtServIndex int32
	// Virtual server service index.
	HttpPerServStatsVirtServiceIndex int32
	Params                           *HttpPerServStatsTableParams
}

func NewHttpPerServStatsTableList() *HttpPerServStatsTable {
	return &HttpPerServStatsTable{}
}

func NewHttpPerServStatsTable(
	httpPerServStatsVirtServIndex int32,
	httpPerServStatsVirtServiceIndex int32,
	params *HttpPerServStatsTableParams,
) *HttpPerServStatsTable {
	return &HttpPerServStatsTable{
		HttpPerServStatsVirtServIndex:    httpPerServStatsVirtServIndex,
		HttpPerServStatsVirtServiceIndex: httpPerServStatsVirtServiceIndex,
		Params:                           params,
	}
}

func (c *HttpPerServStatsTable) Name() string {
	return "HttpPerServStatsTable"
}

func (c *HttpPerServStatsTable) GetParams() BeanType {
	return c.Params
}

func (c *HttpPerServStatsTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *HttpPerServStatsTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.HttpPerServStatsVirtServIndex).IsZero() &&
		reflect.ValueOf(c.HttpPerServStatsVirtServiceIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.HttpPerServStatsVirtServIndex) + "/" + fmt.Sprint(c.HttpPerServStatsVirtServiceIndex)
}

type HttpPerServStatsTableParams struct {
	// Virtual server number.
	VirtServIndex int32 `json:"VirtServIndex,omitempty"`
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
}

func (p HttpPerServStatsTableParams) iMABean() {}
