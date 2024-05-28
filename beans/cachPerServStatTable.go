package beans

import (
	"fmt"
	"reflect"
)

// CachPerServStatTable A table for cache statistics per virtual service.
// Note:This mib is not supported for VX instance of Virtualization.
type CachPerServStatTable struct {
	// Virtual server Index.
	CachStatPerServVirtServIndex string
	// Virtual server service index.
	CachStatPerServVirtServiceIndex int32
	Params                          *CachPerServStatTableParams
}

func NewCachPerServStatTableList() *CachPerServStatTable {
	return &CachPerServStatTable{}
}

func NewCachPerServStatTable(
	cachStatPerServVirtServIndex string,
	cachStatPerServVirtServiceIndex int32,
	params *CachPerServStatTableParams,
) *CachPerServStatTable {
	return &CachPerServStatTable{
		CachStatPerServVirtServIndex:    cachStatPerServVirtServIndex,
		CachStatPerServVirtServiceIndex: cachStatPerServVirtServiceIndex,
		Params:                          params,
	}
}

func (c *CachPerServStatTable) Name() string {
	return "CachPerServStatTable"
}

func (c *CachPerServStatTable) GetParams() BeanType {
	return c.Params
}

func (c *CachPerServStatTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *CachPerServStatTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.CachStatPerServVirtServIndex).IsZero() &&
		reflect.ValueOf(c.CachStatPerServVirtServiceIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.CachStatPerServVirtServIndex) + "/" + fmt.Sprint(c.CachStatPerServVirtServiceIndex)
}

type CachPerServStatTableParams struct {
	// Virtual server Index.
	StatPerServVirtServIndex string `json:"StatPerServVirtServIndex,omitempty"`
	// Virtual server service index.
	StatPerServVirtServiceIndex int32 `json:"StatPerServVirtServiceIndex,omitempty"`
	// Virtual server service port number.
	StatPerServVirtServPort int32 `json:"StatPerServVirtServPort,omitempty"`
	// Cache policy identifier associated with the virtual service.
	StatPerServCachePolId string `json:"StatPerServCachePolId,omitempty"`
	// Number of objects served by cache during the measuring period for virtual service.
	StatPerServTotObj int32 `json:"StatPerServTotObj,omitempty"`
	// Percent of HTTP requests served by objects from cache for virtual service.
	StatPerServHitPerc int32 `json:"StatPerServHitPerc,omitempty"`
	// Rate of requests served by cache every second for virtual service.
	StatPerServServRate int32 `json:"StatPerServServRate,omitempty"`
	// Amount of new objects cached during the measuring period for virtual service.
	StatPerServNewCachedObj int32 `json:"StatPerServNewCachedObj,omitempty"`
	// Average number of new objects cached every second for virtual service.
	StatPerServRateNewCachedObj int32 `json:"StatPerServRateNewCachedObj,omitempty"`
	// Amount of new bytes cached during the measuring period for virtual service.
	StatPerServNewCachedBytes int32 `json:"StatPerServNewCachedBytes,omitempty"`
	// Average number of new bytes cached every second for virtual service.
	StatPerServRateNewCachedBytes int32 `json:"StatPerServRateNewCachedBytes,omitempty"`
	// Amount of objects cached during last measuring period of size smaller than 10KB for virtual service.
	StatPerServObjSmaller10K int32 `json:"StatPerServObjSmaller10K,omitempty"`
	// Amount of objects cached during last measuring period of size between 11KB and 50KB for virtual service.
	StatPerServObj11KTO50K int32 `json:"StatPerServObj11KTO50K,omitempty"`
	// Amount of objects cached during last measuring period of size between 51KB and 100KB for virtual service.
	StatPerServObj51KTO100K int32 `json:"StatPerServObj51KTO100K,omitempty"`
	// Amount of objects cached during last measuring period of size between 101KB and 1MB for virtual service.
	StatPerServObj101KTO1M int32 `json:"StatPerServObj101KTO1M,omitempty"`
	// Amount of objects cached during last measuring period of size larger than 1MB for virtual service.
	StatPerServObjLarger1M int32 `json:"StatPerServObjLarger1M,omitempty"`
	// Peak of new objects cached during the measuring period for virtual service.
	StatPerServPeakNewCachedObj int32 `json:"StatPerServPeakNewCachedObj,omitempty"`
}

func (p CachPerServStatTableParams) iMABean() {}
