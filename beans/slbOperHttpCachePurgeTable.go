package beans

import (
	"fmt"
	"reflect"
)

// SlbOperHttpCachePurgeTable The table of services per server.
// Note:This mib is not supported for VX instance of Virtualization.
type SlbOperHttpCachePurgeTable struct {
	// The Http Cache purge server index.
	SlbOperHttpCachePurgeServerIndex int32
	// The Http Cache purge service index.
	SlbOperHttpCachePurgeServiceIndex int32
	Params                            *SlbOperHttpCachePurgeTableParams
}

func NewSlbOperHttpCachePurgeTableList() *SlbOperHttpCachePurgeTable {
	return &SlbOperHttpCachePurgeTable{}
}

func NewSlbOperHttpCachePurgeTable(
	slbOperHttpCachePurgeServerIndex int32,
	slbOperHttpCachePurgeServiceIndex int32,
	params *SlbOperHttpCachePurgeTableParams,
) *SlbOperHttpCachePurgeTable {
	return &SlbOperHttpCachePurgeTable{
		SlbOperHttpCachePurgeServerIndex:  slbOperHttpCachePurgeServerIndex,
		SlbOperHttpCachePurgeServiceIndex: slbOperHttpCachePurgeServiceIndex,
		Params:                            params,
	}
}

func (c *SlbOperHttpCachePurgeTable) Name() string {
	return "SlbOperHttpCachePurgeTable"
}

func (c *SlbOperHttpCachePurgeTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbOperHttpCachePurgeTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbOperHttpCachePurgeTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbOperHttpCachePurgeServerIndex).IsZero() &&
		reflect.ValueOf(c.SlbOperHttpCachePurgeServiceIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbOperHttpCachePurgeServerIndex) + "/" + fmt.Sprint(c.SlbOperHttpCachePurgeServiceIndex)
}

type SlbOperHttpCachePurgeTableParams struct {
	// The Http Cache purge server index.
	ServerIndex int32 `json:"ServerIndex,omitempty"`
	// The Http Cache purge service index.
	ServiceIndex int32 `json:"ServiceIndex,omitempty"`
	// Set object URL or all.
	URL string `json:"URL,omitempty"`
}

func (p SlbOperHttpCachePurgeTableParams) iMABean() {}
