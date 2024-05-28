package beans

import (
	"fmt"
	"reflect"
)

// FastviewPerWebAppStatsTable A table for Fastview statistics per web application.
// Note:This mib is not supported for VX instance of Virtualization.
type FastviewPerWebAppStatsTable struct {
	// Fastview webapp identifier index.
	FastviewEnhPerWebAppStatsWebappId string
	Params                            *FastviewPerWebAppStatsTableParams
}

func NewFastviewPerWebAppStatsTableList() *FastviewPerWebAppStatsTable {
	return &FastviewPerWebAppStatsTable{}
}

func NewFastviewPerWebAppStatsTable(
	fastviewEnhPerWebAppStatsWebappId string,
	params *FastviewPerWebAppStatsTableParams,
) *FastviewPerWebAppStatsTable {
	return &FastviewPerWebAppStatsTable{
		FastviewEnhPerWebAppStatsWebappId: fastviewEnhPerWebAppStatsWebappId,
		Params:                            params,
	}
}

func (c *FastviewPerWebAppStatsTable) Name() string {
	return "FastviewPerWebAppStatsTable"
}

func (c *FastviewPerWebAppStatsTable) GetParams() BeanType {
	return c.Params
}

func (c *FastviewPerWebAppStatsTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *FastviewPerWebAppStatsTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.FastviewEnhPerWebAppStatsWebappId).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.FastviewEnhPerWebAppStatsWebappId)
}

type FastviewPerWebAppStatsTableParams struct {
	// Fastview webapp identifier index.
	WebappId string `json:"WebappId,omitempty"`
	// Number of fastview transactions during the measuring period for web application.
	Transactions int32 `json:"Transactions,omitempty"`
	// Number of pages served by fastview during the measuring period for web application.
	Pages int32 `json:"Pages,omitempty"`
	// Number of optimized pages served by fastview during the measuring period for web application.
	OptimizedPages int32 `json:"OptimizedPages,omitempty"`
	// Number of learning pages served by fastview during the measuring period for web application.
	LearningPages int32 `json:"LearningPages,omitempty"`
	// Number of fastview parsed tokens during the measuring period for web application.
	TokensParsed int32 `json:"TokensParsed,omitempty"`
	// Number of fastview rewritten tokens during the measuring period for web application.
	TokensRewritten int32 `json:"TokensRewritten,omitempty"`
	// Number of bytes saved with image reduction by fastview during the measuring period for web application.
	BytesSavedImgReduction int32 `json:"BytesSavedImgReduction,omitempty"`
	// Number of bytes before image reduction by fastview during the measuring period for web application.
	BytesBeforeImgReduction int32 `json:"BytesBeforeImgReduction,omitempty"`
	// % of bytes saved with image reduction by fastview during the measuring period for web application.
	BytesSavedPercent int32 `json:"BytesSavedPercent,omitempty"`
	// Number of responses with expirty modified by fastview during the measuring period for web application.
	RespWithExpiryModified int32 `json:"RespWithExpiryModified,omitempty"`
	// % of responses with expirty modified by fastview during the measuring period for web application.
	ExpiryModifiedPercent int32 `json:"ExpiryModifiedPercent,omitempty"`
	// Peak Number of fastview transactions during the measuring period for web application.
	PeakTransactions int32 `json:"PeakTransactions,omitempty"`
	// Peak Number of pages served by fastview during the measuring period for web application.
	PeakPages int32 `json:"PeakPages,omitempty"`
	// Peak Number of optimized pages served by fastview during the measuring period for web application.
	PeakOptimizedPages int32 `json:"PeakOptimizedPages,omitempty"`
	// Peak Number of learning pages served by fastview during the measuring period for web application.
	PeakLearningPages int32 `json:"PeakLearningPages,omitempty"`
	// Peak Number of fastview parsed tokens during the measuring period for web application.
	PeakTokensParsed int32 `json:"PeakTokensParsed,omitempty"`
	// Peak Number of fastview rewritten tokens during the measuring period for web application.
	PeakTokensRewritten int32 `json:"PeakTokensRewritten,omitempty"`
	// Total Number of fastview transactions during the measuring period for web application.
	TotTransactions int32 `json:"TotTransactions,omitempty"`
	// Total Number of pages served by fastview during the measuring period for web application.
	TotPages int32 `json:"TotPages,omitempty"`
	// Total Number of optimized pages served by fastview during the measuring period for web application.
	TotOptimizedPages int32 `json:"TotOptimizedPages,omitempty"`
	// Total Number of learning pages served by fastview during the measuring period for web application.
	TotLearningPages int32 `json:"TotLearningPages,omitempty"`
	// Total Number of fastview parsed tokens during the measuring period for web application.
	TotTokensParsed int32 `json:"TotTokensParsed,omitempty"`
	// Total Number of fastview rewritten tokens during the measuring period for web application.
	TotTokensRewritten int32 `json:"TotTokensRewritten,omitempty"`
	// Total Number of bytes saved with image reduction by fastview during the measuring period for web application.
	TotBytesSavedImgReduction int32 `json:"TotBytesSavedImgReduction,omitempty"`
	// Total Number of bytes before image reduction by fastview during the measuring period for web application.
	TotBytesBeforeImgReduction int32 `json:"TotBytesBeforeImgReduction,omitempty"`
	// % of bytes saved with image reduction by fastview during the measuring period for web application.
	TotBytesSavedPercent int32 `json:"TotBytesSavedPercent,omitempty"`
	// Total Number of responses with expirty modified by fastview during the measuring period for web application.
	TotRespWithExpiryModified int32 `json:"TotRespWithExpiryModified,omitempty"`
	// % of responses with expirty modified by fastview during the measuring period for web application.
	TotExpiryModifiedPercent int32 `json:"TotExpiryModifiedPercent,omitempty"`
	// Number of fastview compiled pages during the measuring period for web application.
	CompiledPages int32 `json:"CompiledPages,omitempty"`
	// Peak of fastview compiled pages during the measuring period for web application.
	PeakCompiledPages int32 `json:"PeakCompiledPages,omitempty"`
	// Total of fastview compiled pages during the measuring period for web application.
	TotCompiledPages int32 `json:"TotCompiledPages,omitempty"`
}

func (p FastviewPerWebAppStatsTableParams) iMABean() {}
