package beans

import (
	"fmt"
	"reflect"
)

// BotManagerStatsTable Table for bot manager stats.
type BotManagerStatsTable struct {
	// Bot Manager ID.
	BotManagerStatBotIndex string
	Params                 *BotManagerStatsTableParams
}

func NewBotManagerStatsTableList() *BotManagerStatsTable {
	return &BotManagerStatsTable{}
}

func NewBotManagerStatsTable(
	botManagerStatBotIndex string,
	params *BotManagerStatsTableParams,
) *BotManagerStatsTable {
	return &BotManagerStatsTable{
		BotManagerStatBotIndex: botManagerStatBotIndex,
		Params:                 params,
	}
}

func (c *BotManagerStatsTable) Name() string {
	return "BotManagerStatsTable"
}

func (c *BotManagerStatsTable) GetParams() BeanType {
	return c.Params
}

func (c *BotManagerStatsTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *BotManagerStatsTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.BotManagerStatBotIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.BotManagerStatBotIndex)
}

type BotManagerStatsTableParams struct {
	// Bot Manager ID.
	BotIndex string `json:"BotIndex,omitempty"`
	// Current average Bot Manager latency.
	AvgLatencyPerSec uint32 `json:"AvgLatencyPerSec,omitempty"`
	// Total average Bot Manager latency.
	AvgLatencyTotal uint32 `json:"AvgLatencyTotal,omitempty"`
	// Current number of requests sent to Bot Manager.
	RequestSentPerSec uint32 `json:"RequestSentPerSec,omitempty"`
	// Percentage of requests sent to Bot Manager out of total requests to services with this Bot Manager policy.
	RequestSentPerSecPercentage uint32 `json:"RequestSentPerSecPercentage,omitempty"`
	// Total number of requests sent to Bot Manager.
	RequestSentTotal uint64 `json:"RequestSentTotal,omitempty"`
	// Current number of requests filtered from Bot inspection due to allow lists.
	RequestFilteredPerSec uint32 `json:"RequestFilteredPerSec,omitempty"`
	// Percentage of requests filtered from Bot inspection due to allow lists out of requests reached to Bot module.
	RequestFilteredPerSecPercentage uint32 `json:"RequestFilteredPerSecPercentage,omitempty"`
	// Total number of requests filtered from Bot inspection due to allow list.
	RequestFilteredTotal uint64 `json:"RequestFilteredTotal,omitempty"`
	// Percentage of total requests filtered from Bot inspection due to allow lists out of requests reached to Bot module.
	RequestFilteredTotalPercentage uint32 `json:"RequestFilteredTotalPercentage,omitempty"`
	// Current number of requests which the response from Bot Manager was timed out.
	RequestTimeoutPerSec uint32 `json:"RequestTimeoutPerSec,omitempty"`
	// Percentage of requests which the response from Bot Manager was timed out, out of total request to Bot Manager.
	RequestTimeoutPerSecPercentage uint32 `json:"RequestTimeoutPerSecPercentage,omitempty"`
	// Total number of requests which the response from Bot Manager was timed out.
	RequestTimeoutTotal uint64 `json:"RequestTimeoutTotal,omitempty"`
	// Percentage of tottal requests which the response from Bot Manager was timed out, out of total request to Bot Manager.
	RequestTimeoutTotalPercentage uint32 `json:"RequestTimeoutTotalPercentage,omitempty"`
	// Current number of requests which the response from Bot Manager was allow.
	RequestAllowPerSec uint32 `json:"RequestAllowPerSec,omitempty"`
	// Percentage of requests which the response from Bot Manager was allow out of total requests to Bot Manager.
	RequestAllowPerSecPercentage uint32 `json:"RequestAllowPerSecPercentage,omitempty"`
	// Total number of request which the response from Bot Manager was allow.
	RequestAllowTotal uint64 `json:"RequestAllowTotal,omitempty"`
	// Percentage of total requests which the response from Bot Manager was allow out of total requests to Bot Manager.
	RequestAllowTotalPercentage uint32 `json:"RequestAllowTotalPercentage,omitempty"`
	// Current number of requests which the response from Bot Manager was captcha.
	RequestCaptchaPerSec uint32 `json:"RequestCaptchaPerSec,omitempty"`
	// Percentage of requests which the response from Bot Manager was captcha out of total requests to Bot Manager.
	RequestCaptchaPerSecPercentage uint32 `json:"RequestCaptchaPerSecPercentage,omitempty"`
	// Total number of request which the response from Bot Manager was captcha.
	RequestCaptchaTotal uint64 `json:"RequestCaptchaTotal,omitempty"`
	// Percentage of total requests which the response from Bot Manager was captcha out of total requests to Bot Manager.
	RequestCaptchaTotalPercentage uint32 `json:"RequestCaptchaTotalPercentage,omitempty"`
	// Current number of requests which the response from Bot Manager was block.
	RequestBlockPerSec uint32 `json:"RequestBlockPerSec,omitempty"`
	// Percentage of requests which the response from Bot Manager was block out of total requests to Bot Manager.
	RequestBlockPerSecPercentage uint32 `json:"RequestBlockPerSecPercentage,omitempty"`
	// Total number of request which the response from Bot Manager was block.
	RequestBlockTotal uint64 `json:"RequestBlockTotal,omitempty"`
	// Percentage of total requests which the response from Bot Manager was block out of total requests to Bot Manager.
	RequestBlockTotalPercentage uint32 `json:"RequestBlockTotalPercentage,omitempty"`
	// Current number of requests which the response from Bot Manager was other.
	RequestOtherPerSec uint32 `json:"RequestOtherPerSec,omitempty"`
	// Percentage of requests which the response from Bot Manager was other out of total requests to Bot Manager.
	RequestOtherPerSecPercentage uint32 `json:"RequestOtherPerSecPercentage,omitempty"`
	// Total number of request which the response from Bot Manager was other.
	RequestOtherTotal uint64 `json:"RequestOtherTotal,omitempty"`
	// Percentage of total requests which the response from Bot Manager was other out of total requests to Bot Manager.
	RequestOtherTotalPercentage uint32 `json:"RequestOtherTotalPercentage,omitempty"`
	// Current number of requests which the response from Bot Manager had http error.
	HttpErrPerSec uint32 `json:"HttpErrPerSec,omitempty"`
	// Percentage of requests which the response from Bot Manager had http error out of total requests to Bot Manager.
	HttpErrPerSecPercentage uint32 `json:"HttpErrPerSecPercentage,omitempty"`
	// Total number of request which the response from Bot Manager had http error.
	HttpErrTotal uint64 `json:"HttpErrTotal,omitempty"`
	// Percentage of total requests which the response from Bot Manager had http error out of total requests to Bot Manager.
	HttpErrTotalPercentage uint32 `json:"HttpErrTotalPercentage,omitempty"`
}

func (p BotManagerStatsTableParams) iMABean() {}
