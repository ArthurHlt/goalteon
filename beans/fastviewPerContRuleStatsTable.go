package beans

import (
	"fmt"
	"reflect"
)

// FastviewPerContRuleStatsTable A table for Fastview statistics per content class.
// Note:This mib is not supported for VX instance of Virtualization.
type FastviewPerContRuleStatsTable struct {
	// Virtual server Index.
	FastviewEnhPerContRuleStatsVirtServIndex string
	// Virtual server service index.
	FastviewEnhPerContRuleStatsVirtServiceIndex int32
	// Content rule index.
	FastviewEnhPerContRuleStatsIndex int32
	Params                           *FastviewPerContRuleStatsTableParams
}

func NewFastviewPerContRuleStatsTableList() *FastviewPerContRuleStatsTable {
	return &FastviewPerContRuleStatsTable{}
}

func NewFastviewPerContRuleStatsTable(
	fastviewEnhPerContRuleStatsVirtServIndex string,
	fastviewEnhPerContRuleStatsVirtServiceIndex int32,
	fastviewEnhPerContRuleStatsIndex int32,
	params *FastviewPerContRuleStatsTableParams,
) *FastviewPerContRuleStatsTable {
	return &FastviewPerContRuleStatsTable{
		FastviewEnhPerContRuleStatsVirtServIndex:    fastviewEnhPerContRuleStatsVirtServIndex,
		FastviewEnhPerContRuleStatsVirtServiceIndex: fastviewEnhPerContRuleStatsVirtServiceIndex,
		FastviewEnhPerContRuleStatsIndex:            fastviewEnhPerContRuleStatsIndex,
		Params:                                      params,
	}
}

func (c *FastviewPerContRuleStatsTable) Name() string {
	return "FastviewPerContRuleStatsTable"
}

func (c *FastviewPerContRuleStatsTable) GetParams() BeanType {
	return c.Params
}

func (c *FastviewPerContRuleStatsTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *FastviewPerContRuleStatsTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.FastviewEnhPerContRuleStatsVirtServIndex).IsZero() &&
		reflect.ValueOf(c.FastviewEnhPerContRuleStatsVirtServiceIndex).IsZero() &&
		reflect.ValueOf(c.FastviewEnhPerContRuleStatsIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.FastviewEnhPerContRuleStatsVirtServIndex) + "/" + fmt.Sprint(c.FastviewEnhPerContRuleStatsVirtServiceIndex) + "/" + fmt.Sprint(c.FastviewEnhPerContRuleStatsIndex)
}

type FastviewPerContRuleStatsTableParams struct {
	// Virtual server Index.
	VirtServIndex string `json:"VirtServIndex,omitempty"`
	// Virtual server service index.
	VirtServiceIndex int32 `json:"VirtServiceIndex,omitempty"`
	// Content rule index.
	Index int32 `json:"Index,omitempty"`
	// Content rule name.
	Name string `json:"Name,omitempty"`
	// Content class associated with the content rule.
	ContClass string `json:"ContClass,omitempty"`
	// Fastview webapp identifier associated with the content rule.
	WebappId string `json:"WebappId,omitempty"`
	// Number of fastview transactions during the measuring period for content rule.
	Transactions int32 `json:"Transactions,omitempty"`
	// Number of pages served by fastview during the measuring period for content rule.
	Pages int32 `json:"Pages,omitempty"`
	// Number of optimized pages served by fastview during the measuring period for content rule.
	OptimizedPages int32 `json:"OptimizedPages,omitempty"`
	// Number of learning pages served by fastview during the measuring period for content rule.
	LearningPages int32 `json:"LearningPages,omitempty"`
	// Number of fastview parsed tokens during the measuring period for content rule.
	TokensParsed int32 `json:"TokensParsed,omitempty"`
	// Number of fastview rewritten tokens during the measuring period for content rule.
	TokensRewritten int32 `json:"TokensRewritten,omitempty"`
	// Number of bytes saved with image reduction by fastview during the measuring period for content rule.
	BytesSavedImgReduction int32 `json:"BytesSavedImgReduction,omitempty"`
	// Number of bytes before image reduction by fastview during the measuring period for content rule.
	BytesBeforeImgReduction int32 `json:"BytesBeforeImgReduction,omitempty"`
	// % of bytes saved with image reduction by fastview during the measuring period for content rule.
	BytesSavedPercent int32 `json:"BytesSavedPercent,omitempty"`
	// Number of responses with expirty modified by fastview during the measuring period for content rule.
	RespWithExpiryModified int32 `json:"RespWithExpiryModified,omitempty"`
	// % of responses with expirty modified by fastview during the measuring period for content rule.
	ExpiryModifiedPercent int32 `json:"ExpiryModifiedPercent,omitempty"`
	// Peak Number of fastview transactions during the measuring period for content rule.
	PeakTransactions int32 `json:"PeakTransactions,omitempty"`
	// Peak Number of pages served by fastview during the measuring period for content rule.
	PeakPages int32 `json:"PeakPages,omitempty"`
	// Peak Number of optimized pages served by fastview during the measuring period for content rule.
	PeakOptimizedPages int32 `json:"PeakOptimizedPages,omitempty"`
	// Peak Number of learning pages served by fastview during the measuring period for content rule.
	PeakLearningPages int32 `json:"PeakLearningPages,omitempty"`
	// Peak Number of fastview parsed tokens during the measuring period for content rule.
	PeakTokensParsed int32 `json:"PeakTokensParsed,omitempty"`
	// Peak Number of fastview rewritten tokens during the measuring period for content rule.
	PeakTokensRewritten int32 `json:"PeakTokensRewritten,omitempty"`
	// Total Number of fastview transactions during the measuring period for content rule.
	TotTransactions int32 `json:"TotTransactions,omitempty"`
	// Total Number of pages served by fastview during the measuring period for content rule.
	TotPages int32 `json:"TotPages,omitempty"`
	// Total Number of optimized pages served by fastview during the measuring period for content rule.
	TotOptimizedPages int32 `json:"TotOptimizedPages,omitempty"`
	// Total Number of learning pages served by fastview during the measuring period for content rule.
	TotLearningPages int32 `json:"TotLearningPages,omitempty"`
	// Total Number of fastview parsed tokens during the measuring period for content rule.
	TotTokensParsed int32 `json:"TotTokensParsed,omitempty"`
	// Total Number of fastview rewritten tokens during the measuring period for content rule.
	TotTokensRewritten int32 `json:"TotTokensRewritten,omitempty"`
	// Total Number of bytes saved with image reduction by fastview during the measuring period for content rule.
	TotBytesSavedImgReduction int32 `json:"TotBytesSavedImgReduction,omitempty"`
	// Total Number of bytes before image reduction by fastview during the measuring period for content rule.
	TotBytesBeforeImgReduction int32 `json:"TotBytesBeforeImgReduction,omitempty"`
	// % of bytes saved with image reduction by fastview during the measuring period for content rule.
	TotBytesSavedPercent int32 `json:"TotBytesSavedPercent,omitempty"`
	// Total Number of responses with expirty modified by fastview during the measuring period for content rule.
	TotRespWithExpiryModified int32 `json:"TotRespWithExpiryModified,omitempty"`
	// % of responses with expirty modified by fastview during the measuring period for content rule.
	TotExpiryModifiedPercent int32 `json:"TotExpiryModifiedPercent,omitempty"`
}

func (p FastviewPerContRuleStatsTableParams) iMABean() {}
