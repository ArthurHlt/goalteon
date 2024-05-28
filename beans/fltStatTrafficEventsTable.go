package beans

import (
	"fmt"
	"reflect"
)

// FltStatTrafficEventsTable Traffic Events statistics table for filters.
// Note:This mib is not supported for VX instance of Virtualization.
type FltStatTrafficEventsTable struct {
	// Index to identify Filter.
	FltTrafficEventStatFltIndex int32
	Params                      *FltStatTrafficEventsTableParams
}

func NewFltStatTrafficEventsTableList() *FltStatTrafficEventsTable {
	return &FltStatTrafficEventsTable{}
}

func NewFltStatTrafficEventsTable(
	fltTrafficEventStatFltIndex int32,
	params *FltStatTrafficEventsTableParams,
) *FltStatTrafficEventsTable {
	return &FltStatTrafficEventsTable{
		FltTrafficEventStatFltIndex: fltTrafficEventStatFltIndex,
		Params:                      params,
	}
}

func (c *FltStatTrafficEventsTable) Name() string {
	return "FltStatTrafficEventsTable"
}

func (c *FltStatTrafficEventsTable) GetParams() BeanType {
	return c.Params
}

func (c *FltStatTrafficEventsTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *FltStatTrafficEventsTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.FltTrafficEventStatFltIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.FltTrafficEventStatFltIndex)
}

type FltStatTrafficEventsTableParams struct {
	// Index to identify Filter.
	TrafficEventStatFltIndex int32 `json:"TrafficEventStatFltIndex,omitempty"`
	// No of Http Requests sent as traffic events per second.
	HttpReqSuccessPerSec uint32 `json:"HttpReqSuccessPerSec,omitempty"`
	// No of Http Requests sent successfully as traffic events in total.
	HttpReqSuccessTotal uint64 `json:"HttpReqSuccessTotal,omitempty"`
	// No of Http responses sent as traffic events per second.
	HttpRespSuccessPerSec uint32 `json:"HttpRespSuccessPerSec,omitempty"`
	// No of Http responses sent successfully as traffic events in total.
	HttpRespSuccessTotal uint64 `json:"HttpRespSuccessTotal,omitempty"`
	// No of fessl sent successfully as traffice events per second.
	FeSSLInfoSuccessPerSec uint32 `json:"FeSSLInfoSuccessPerSec,omitempty"`
	// No of fessl sent successfully as traffic events in total.
	FeSSLInfoSuccessTotal uint64 `json:"FeSSLInfoSuccessTotal,omitempty"`
	// No of bessl sent successfully as traffice events per second.
	BeSSLInfoSuccessPerSec uint32 `json:"BeSSLInfoSuccessPerSec,omitempty"`
	// No of bessl sent successfully as traffic events in total.
	BeSSLInfoSuccessTotal uint64 `json:"BeSSLInfoSuccessTotal,omitempty"`
	// No of host names bypass sent successfully as traffice events per second.
	HostNameBypassSuccessPerSec uint32 `json:"HostNameBypassSuccessPerSec,omitempty"`
	// No of host names bypass sent successfully as traffic events in total.
	HostNameBypassSuccessTotal uint64 `json:"HostNameBypassSuccessTotal,omitempty"`
	// No of Http requests failed to send as traffic events per second.
	HttpReqFailPerSec uint32 `json:"HttpReqFailPerSec,omitempty"`
	// No of Http requests failed to send as traffic events in total.
	HttpReqFailTotal uint64 `json:"HttpReqFailTotal,omitempty"`
	// No of Http responses failed to send as traffic events per second.
	HttpRespFailPerSec uint32 `json:"HttpRespFailPerSec,omitempty"`
	// No of Http responses failed to send as traffic events in total.
	HttpRespFailTotal uint64 `json:"HttpRespFailTotal,omitempty"`
	// No of fessl failed to send as traffice events per second.
	FeSSLInfoFailPerSec uint32 `json:"FeSSLInfoFailPerSec,omitempty"`
	// No of fessl failed to send as traffice events in total.
	FeSSLInfoFailTotal uint64 `json:"FeSSLInfoFailTotal,omitempty"`
	// No of bessl failed to send as traffice events per second.
	BeSSLInfoFailPerSec uint32 `json:"BeSSLInfoFailPerSec,omitempty"`
	// No of bessl failed to send as traffice events in total.
	BeSSLInfoFailTotal uint64 `json:"BeSSLInfoFailTotal,omitempty"`
	// No of host name bypass failed to send as traffic events per second.
	HostNameBypassFailPerSec uint32 `json:"HostNameBypassFailPerSec,omitempty"`
	// No of host name bypass failed to send as traffic events in total.
	HostNameBypassFailTotal uint64 `json:"HostNameBypassFailTotal,omitempty"`
	// Event field missing failure per second.
	EventsFieldMissingPerSec uint32 `json:"EventsFieldMissingPerSec,omitempty"`
	// Event field missing failure in total.
	EventsFieldMissingTotal uint64 `json:"EventsFieldMissingTotal,omitempty"`
	// Event allocation failure per second.
	EventsAllocFailPerSec uint32 `json:"EventsAllocFailPerSec,omitempty"`
	// Event allocation failure in total.
	EventsAllocFailTotal uint64 `json:"EventsAllocFailTotal,omitempty"`
	// Event queue full failure per second.
	EventsQueueFullPerSec uint32 `json:"EventsQueueFullPerSec,omitempty"`
	// Event queue full failure in total.
	EventsQueueFullTotal uint64 `json:"EventsQueueFullTotal,omitempty"`
	// Layer4 requests success per second.
	Layer4ReqSuccessPerSec uint32 `json:"Layer4ReqSuccessPerSec,omitempty"`
	// Layer4 requests success in total.
	Layer4ReqSuccessTotal uint64 `json:"Layer4ReqSuccessTotal,omitempty"`
	// Layer4 requests failed to send per second.
	Layer4ReqFailPerSec uint32 `json:"Layer4ReqFailPerSec,omitempty"`
	// Layer4 requests failed in total.
	Layer4ReqFailTotal uint64 `json:"Layer4ReqFailTotal,omitempty"`
}

func (p FltStatTrafficEventsTableParams) iMABean() {}
