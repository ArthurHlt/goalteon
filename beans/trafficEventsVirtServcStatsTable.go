package beans

import (
	"fmt"
	"reflect"
)

// TrafficEventsVirtServcStatsTable Table for traffic event stats per service.
// Note:This mib is not supported for VX instance of Virtualization.
type TrafficEventsVirtServcStatsTable struct {
	// The virtual server index.
	VirtTrafficEventsStatVirtIndex string
	// The virtual service index.
	VirtTrafficEventsStatSvcIndex int32
	Params                        *TrafficEventsVirtServcStatsTableParams
}

func NewTrafficEventsVirtServcStatsTableList() *TrafficEventsVirtServcStatsTable {
	return &TrafficEventsVirtServcStatsTable{}
}

func NewTrafficEventsVirtServcStatsTable(
	virtTrafficEventsStatVirtIndex string,
	virtTrafficEventsStatSvcIndex int32,
	params *TrafficEventsVirtServcStatsTableParams,
) *TrafficEventsVirtServcStatsTable {
	return &TrafficEventsVirtServcStatsTable{
		VirtTrafficEventsStatVirtIndex: virtTrafficEventsStatVirtIndex,
		VirtTrafficEventsStatSvcIndex:  virtTrafficEventsStatSvcIndex,
		Params:                         params,
	}
}

func (c *TrafficEventsVirtServcStatsTable) Name() string {
	return "TrafficEventsVirtServcStatsTable"
}

func (c *TrafficEventsVirtServcStatsTable) GetParams() BeanType {
	return c.Params
}

func (c *TrafficEventsVirtServcStatsTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *TrafficEventsVirtServcStatsTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.VirtTrafficEventsStatVirtIndex).IsZero() &&
		reflect.ValueOf(c.VirtTrafficEventsStatSvcIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.VirtTrafficEventsStatVirtIndex) + "/" + fmt.Sprint(c.VirtTrafficEventsStatSvcIndex)
}

type TrafficEventsVirtServcStatsTableParams struct {
	// The virtual server index.
	VirtTrafficEventsStatVirtIndex string `json:"virtTrafficEventsStatVirtIndex,omitempty"`
	// The virtual service index.
	VirtTrafficEventsStatSvcIndex int32 `json:"virtTrafficEventsStatSvcIndex,omitempty"`
	// No of Http Requests sent as traffic events per second.
	VirtHttpReqSuccessPerSec uint32 `json:"virtHttpReqSuccessPerSec,omitempty"`
	// No of Http Requests sent successfully as traffic events in total.
	VirtHttpReqSuccessTotal uint64 `json:"virtHttpReqSuccessTotal,omitempty"`
	// No of Http responses sent as traffic events per second.
	VirtHttpRespSuccessPerSec uint32 `json:"virtHttpRespSuccessPerSec,omitempty"`
	// No of Http responses sent successfully as traffic events in total.
	VirtHttpRespSuccessTotal uint64 `json:"virtHttpRespSuccessTotal,omitempty"`
	// No of fessl sent successfully as traffice events per second.
	VirtFeSSLInfoSuccessPerSec uint32 `json:"virtFeSSLInfoSuccessPerSec,omitempty"`
	// No of fessl sent successfully as traffic events in total.
	VirtFeSSLInfoSuccessTotal uint64 `json:"virtFeSSLInfoSuccessTotal,omitempty"`
	// No of bessl sent successfully as traffice events per second.
	VirtBeSSLInfoSuccessPerSec uint32 `json:"virtBeSSLInfoSuccessPerSec,omitempty"`
	// No of bessl sent successfully as traffic events in total.
	VirtBeSSLInfoSuccessTotal uint64 `json:"virtBeSSLInfoSuccessTotal,omitempty"`
	// No of host names bypass sent successfully as traffice events per second.
	VirtHostNameBypassSuccessPerSec uint32 `json:"virtHostNameBypassSuccessPerSec,omitempty"`
	// No of host names bypass sent successfully as traffic events in total.
	VirtHostNameBypassSuccessTotal uint64 `json:"virtHostNameBypassSuccessTotal,omitempty"`
	// No of Http requests failed to send as traffic events per second.
	VirtHttpReqFailPerSec uint32 `json:"virtHttpReqFailPerSec,omitempty"`
	// No of Http requests failed to send as traffic events in total.
	VirtHttpReqFailTotal uint64 `json:"virtHttpReqFailTotal,omitempty"`
	// No of Http responses failed to send as traffic events per second.
	VirtHttpRespFailPerSec uint32 `json:"virtHttpRespFailPerSec,omitempty"`
	// No of Http responses failed to send as traffic events in total.
	VirtHttpRespFailTotal uint64 `json:"virtHttpRespFailTotal,omitempty"`
	// No of fessl failed to send as traffice events per second.
	VirtFeSSLInfoFailPerSec uint32 `json:"virtFeSSLInfoFailPerSec,omitempty"`
	// No of fessl failed to send as traffice events in total.
	VirtFeSSLInfoFailTotal uint64 `json:"virtFeSSLInfoFailTotal,omitempty"`
	// No of bessl failed to send as traffice events per second.
	VirtBeSSLInfoFailPerSec uint32 `json:"virtBeSSLInfoFailPerSec,omitempty"`
	// No of bessl failed to send as traffice events in total.
	VirtBeSSLInfoFailTotal uint64 `json:"virtBeSSLInfoFailTotal,omitempty"`
	// No of host name bypass failed to send as traffic events per second.
	VirtHostNameBypassFailPerSec uint32 `json:"virtHostNameBypassFailPerSec,omitempty"`
	// No of host name bypass failed to send as traffic events in total.
	VirtHostNameBypassFailTotal uint64 `json:"virtHostNameBypassFailTotal,omitempty"`
	// Event field missing failure per second.
	VirtEventsFieldMissingPerSec uint32 `json:"virtEventsFieldMissingPerSec,omitempty"`
	// Event field missing failure in total.
	VirtEventsFieldMissingTotal uint64 `json:"virtEventsFieldMissingTotal,omitempty"`
	// Event allocation failure per second.
	VirtEventsAllocFailPerSec uint32 `json:"virtEventsAllocFailPerSec,omitempty"`
	// Event allocation failure in total.
	VirtEventsAllocFailTotal uint64 `json:"virtEventsAllocFailTotal,omitempty"`
	// Event queue full failure per second.
	VirtEventsQueueFullPerSec uint32 `json:"virtEventsQueueFullPerSec,omitempty"`
	// Event queue full failure in total.
	VirtEventsQueueFullTotal uint64 `json:"virtEventsQueueFullTotal,omitempty"`
	// layer 4 events sent per second.
	VirtLayer4ReqSuccessPerSec uint32 `json:"virtLayer4ReqSuccessPerSec,omitempty"`
	// layer 4 events sent in total.
	VirtLayer4ReqSuccessTotal uint64 `json:"virtLayer4ReqSuccessTotal,omitempty"`
	// layer 4 events failure per second.
	VirtLayer4ReqFailPerSec uint32 `json:"virtLayer4ReqFailPerSec,omitempty"`
	// layer 4 events failure in total.
	VirtLayer4ReqFailTotal uint64 `json:"virtLayer4ReqFailTotal,omitempty"`
	// The layer4 virtual port number of the service.
	VirtLayer4ServicePort uint64 `json:"virtLayer4ServicePort,omitempty"`
	// No. of total unified traffic events that successfully generated.
	VirtUnifiedSucessTotal uint64 `json:"virtUnifiedSucessTotal,omitempty"`
	// No. of unified traffic events per second that successfully generated.
	VirtUnifiedSucessPerSec uint32 `json:"virtUnifiedSucessPerSec,omitempty"`
	// No. of total unified traffic events that failed to be generated.
	VirtUnifiedFailTotal uint64 `json:"virtUnifiedFailTotal,omitempty"`
	// No. of unified traffic events per second that failed to be generated.
	VirtUnifiedFailPerSec uint32 `json:"virtUnifiedFailPerSec,omitempty"`
	// No. of total normal unified traffic events that successfully generated.
	VirtNormalUnifiedSucessTotal uint64 `json:"virtNormalUnifiedSucessTotal,omitempty"`
	// Percentage of total normal unified traffic events that successfully generated based on the defined threshold.
	VirtNormalUnifiedSucessTotalPercentage uint32 `json:"virtNormalUnifiedSucessTotalPercentage,omitempty"`
	// No. of normal unified traffic events per second that successfully generated.
	VirtNormalUnifiedSucessPerSec uint32 `json:"virtNormalUnifiedSucessPerSec,omitempty"`
	// Percentage of normal unified traffic events per second that successfully generated based on the defined threshold.
	VirtNormalUnifiedSucessPerSecPercentage uint32 `json:"virtNormalUnifiedSucessPerSecPercentage,omitempty"`
	// No. of Exception unified traffic events that successfully generated.
	VirtExceptionUnifiedSucessTotal uint64 `json:"virtExceptionUnifiedSucessTotal,omitempty"`
	// Percentage of Exception unified traffic events that successfully generated based on the defined threshold.
	VirtExceptionUnifiedSucessTotalPercentage uint32 `json:"virtExceptionUnifiedSucessTotalPercentage,omitempty"`
	// No. of Exception unified traffic events per second that successfully generated.
	VirtExceptionUnifiedSucessPerSec uint32 `json:"virtExceptionUnifiedSucessPerSec,omitempty"`
	// Percentage of Exception unified traffic events per second that successfully generated based on the defined threshold.
	VirtExceptionUnifiedSucessPerSecPercentage uint32 `json:"virtExceptionUnifiedSucessPerSecPercentage,omitempty"`
	// No. of total security traffic events.
	VirtSecurityTotal uint64 `json:"virtSecurityTotal,omitempty"`
	// No. of security traffic events per second.
	VirtSecurityPerSec uint32 `json:"virtSecurityPerSec,omitempty"`
	// No. of total unified security traffic events.
	VirtUnifiedSecurityTotal uint64 `json:"virtUnifiedSecurityTotal,omitempty"`
	// No. of unified security traffic events per second.
	VirtUnifiedSecurityPerSec uint32 `json:"virtUnifiedSecurityPerSec,omitempty"`
}

func (p TrafficEventsVirtServcStatsTableParams) iMABean() {}
