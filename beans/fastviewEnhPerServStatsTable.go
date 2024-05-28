package beans

import (
	"fmt"
	"reflect"
)

// FastviewEnhPerServStatsTable A table for Fastview statistics per virtual service.
// Note:This mib is not supported for VX instance of Virtualization.
type FastviewEnhPerServStatsTable struct {
	// Virtual server Index.
	FastviewEnhPerServStatsVirtServIndex string
	// Virtual server service index.
	FastviewEnhPerServStatsVirtServiceIndex int32
	Params                                  *FastviewEnhPerServStatsTableParams
}

func NewFastviewEnhPerServStatsTableList() *FastviewEnhPerServStatsTable {
	return &FastviewEnhPerServStatsTable{}
}

func NewFastviewEnhPerServStatsTable(
	fastviewEnhPerServStatsVirtServIndex string,
	fastviewEnhPerServStatsVirtServiceIndex int32,
	params *FastviewEnhPerServStatsTableParams,
) *FastviewEnhPerServStatsTable {
	return &FastviewEnhPerServStatsTable{
		FastviewEnhPerServStatsVirtServIndex:    fastviewEnhPerServStatsVirtServIndex,
		FastviewEnhPerServStatsVirtServiceIndex: fastviewEnhPerServStatsVirtServiceIndex,
		Params:                                  params,
	}
}

func (c *FastviewEnhPerServStatsTable) Name() string {
	return "FastviewEnhPerServStatsTable"
}

func (c *FastviewEnhPerServStatsTable) GetParams() BeanType {
	return c.Params
}

func (c *FastviewEnhPerServStatsTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *FastviewEnhPerServStatsTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.FastviewEnhPerServStatsVirtServIndex).IsZero() &&
		reflect.ValueOf(c.FastviewEnhPerServStatsVirtServiceIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.FastviewEnhPerServStatsVirtServIndex) + "/" + fmt.Sprint(c.FastviewEnhPerServStatsVirtServiceIndex)
}

type FastviewEnhPerServStatsTableParams struct {
	// Virtual server Index.
	VirtServIndex string `json:"VirtServIndex,omitempty"`
	// Virtual server service index.
	VirtServiceIndex int32 `json:"VirtServiceIndex,omitempty"`
	// Virtual server service port number.
	VirtServPort int32 `json:"VirtServPort,omitempty"`
	// Fastview webapp identifier associated with the virtual service.
	WebappId string `json:"WebappId,omitempty"`
	// Number of fastview transactions during the measuring period for virtual service.
	Transactions int32 `json:"Transactions,omitempty"`
	// Number of pages served by fastview during the measuring period for virtual service.
	Pages int32 `json:"Pages,omitempty"`
	// Number of optimized pages served by fastview during the measuring period for virtual service.
	OptimizedPages int32 `json:"OptimizedPages,omitempty"`
	// Number of learning pages served by fastview during the measuring period for virtual service.
	LearningPages int32 `json:"LearningPages,omitempty"`
	// Number of fastview parsed tokens during the measuring period for virtual service.
	TokensParsed int32 `json:"TokensParsed,omitempty"`
	// Number of fastview rewritten tokens during the measuring period for virtual service.
	TokensRewritten int32 `json:"TokensRewritten,omitempty"`
	// Number of fastview bytes saved with image reduction during the measuring period for virtual service.
	BytesSavedImgReduction int32 `json:"BytesSavedImgReduction,omitempty"`
	// Number of fastview bytes before image reduction during the measuring period for virtual service.
	BytesBeforeImgReduction int32 `json:"BytesBeforeImgReduction,omitempty"`
	// % of fastview bytes saved with image reduction during the measuring period for virtual service.
	BytesSavedPercent int32 `json:"BytesSavedPercent,omitempty"`
	// Number of responses with expiry modified by fastview during the measuring period for virtual service.
	RespWithExpiryModified int32 `json:"RespWithExpiryModified,omitempty"`
	// % of responses with expiry modified by fastview during the measuring period for virtual service.
	ExpiryModifiedPercent int32 `json:"ExpiryModifiedPercent,omitempty"`
	// Peak Number of fastview transactions during the measuring period for virtual service.
	PeakTransactions int32 `json:"PeakTransactions,omitempty"`
	// Peak Number of pages served by fastview during the measuring period for virtual service.
	PeakPages int32 `json:"PeakPages,omitempty"`
	// Peak Number of optimized pages served by fastview during the measuring period for virtual service.
	PeakOptimizedPages int32 `json:"PeakOptimizedPages,omitempty"`
	// Peak Number of learning pages served by fastview during the measuring period for virtual service.
	PeakLearningPages int32 `json:"PeakLearningPages,omitempty"`
	// Peak Number of fastview parsed tokens during the measuring period for virtual service.
	PeakTokensParsed int32 `json:"PeakTokensParsed,omitempty"`
	// Peak Number of fastview rewritten tokens during the measuring period for virtual service.
	PeakTokensRewritten int32 `json:"PeakTokensRewritten,omitempty"`
	// Total Number of fastview transactions during the measuring period for virtual service.
	TotTransactions int32 `json:"TotTransactions,omitempty"`
	// Total Number of pages served by fastview during the measuring period for virtual service.
	TotPages int32 `json:"TotPages,omitempty"`
	// Total Number of optimized pages served by fastview during the measuring period for virtual service.
	TotOptimizedPages int32 `json:"TotOptimizedPages,omitempty"`
	// Total Number of learning pages served by fastview during the measuring period for virtual service.
	TotLearningPages int32 `json:"TotLearningPages,omitempty"`
	// Total Number of fastview parsed tokens during the measuring period for virtual service.
	TotTokensParsed int32 `json:"TotTokensParsed,omitempty"`
	// Total Number of fastview rewritten tokens during the measuring period for virtual service.
	TotTokensRewritten int32 `json:"TotTokensRewritten,omitempty"`
	// Total Number of fastview bytes saved with image reduction during the measuring period for virtual service.
	TotBytesSavedImgReduction int32 `json:"TotBytesSavedImgReduction,omitempty"`
	// Total Number of fastview bytes before image reduction during the measuring period for virtual service.
	TotBytesBeforeImgReduction int32 `json:"TotBytesBeforeImgReduction,omitempty"`
	// % of fastview bytes saved with image reduction during the measuring period for virtual service.
	TotBytesSavedPercent int32 `json:"TotBytesSavedPercent,omitempty"`
	// Total Number of responses with expiry modified by fastview during the measuring period for virtual service.
	TotRespWithExpiryModified int32 `json:"TotRespWithExpiryModified,omitempty"`
	// % of responses with expiry modified by fastview during the measuring period for virtual service.
	TotExpiryModifiedPercent int32 `json:"TotExpiryModifiedPercent,omitempty"`
}

func (p FastviewEnhPerServStatsTableParams) iMABean() {}
