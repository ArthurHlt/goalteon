package goalteaon

import (
	"fmt"
	"net/url"
	"strings"
)

const (
	UrlParamsKeyProp       = "prop"
	UrlParamsKeyProps      = "props"
	UrlParamsKeyOffset     = "offset"
	UrlParamsKeyCount      = "count"
	UrlParamsKeyAction     = "action"
	UrlParamsKeyFilter     = "filter"
	UrlParamsKeyFilterType = "filtertype"
)

type FilterType string

const (
	FilterTypeExact FilterType = "exact"
	FilterTypeAny   FilterType = "any"
)

type Filter struct {
	Key   string
	Value string
}

func (f Filter) String() string {
	return fmt.Sprintf("%s:%s", f.Key, f.Value)
}

func UrlParamsOffset(prev url.Values, offset int) url.Values {
	if prev == nil {
		prev = url.Values{}
	}
	prev.Set(UrlParamsKeyOffset, fmt.Sprint(offset))
	return prev
}

func UrlParamsCount(prev url.Values, count int) url.Values {
	if prev == nil {
		prev = url.Values{}
	}
	prev.Set(UrlParamsKeyCount, fmt.Sprint(count))
	return prev
}

func UrlParamsProps(prev url.Values, props ...string) url.Values {
	if prev == nil {
		prev = url.Values{}
	}
	preVal := prev.Get(UrlParamsKeyProp)
	propRaw := strings.Join(props, ",")
	if preVal != "" {
		propRaw = preVal + "," + propRaw
	}
	prev.Set(UrlParamsKeyProp, propRaw)
	return prev
}

func UrlParamsAction(prev url.Values, action string) url.Values {
	if prev == nil {
		prev = url.Values{}
	}
	prev.Set(UrlParamsKeyAction, action)
	return prev
}

func UrlParamsFilters(prev url.Values, filterType FilterType, filters ...Filter) url.Values {
	if prev == nil {
		prev = url.Values{}
	}
	filtersRaw := make([]string, 0)
	for _, filter := range filters {
		filtersRaw = append(filtersRaw, filter.String())
	}
	preVal := prev.Get(UrlParamsKeyFilter)
	filterRaw := strings.Join(filtersRaw, ",")
	if preVal != "" {
		filterRaw = preVal + "," + filterRaw
	}
	prev.Set(UrlParamsKeyFilter, filterRaw)
	prev.Set(UrlParamsKeyFilterType, string(filterType))
	return prev
}
