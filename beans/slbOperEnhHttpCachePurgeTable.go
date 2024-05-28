package beans

import (
	"fmt"
	"reflect"
)

// SlbOperEnhHttpCachePurgeTable The table of services per server.
// Note:This mib is not supported for VX instance of Virtualization.
type SlbOperEnhHttpCachePurgeTable struct {
	// The Http Cache purge server index.
	SlbOperEnhHttpCachePurgeServerIndex string
	// The Http Cache purge service index.
	SlbOperEnhHttpCachePurgeServiceIndex int32
	Params                               *SlbOperEnhHttpCachePurgeTableParams
}

func NewSlbOperEnhHttpCachePurgeTableList() *SlbOperEnhHttpCachePurgeTable {
	return &SlbOperEnhHttpCachePurgeTable{}
}

func NewSlbOperEnhHttpCachePurgeTable(
	slbOperEnhHttpCachePurgeServerIndex string,
	slbOperEnhHttpCachePurgeServiceIndex int32,
	params *SlbOperEnhHttpCachePurgeTableParams,
) *SlbOperEnhHttpCachePurgeTable {
	return &SlbOperEnhHttpCachePurgeTable{
		SlbOperEnhHttpCachePurgeServerIndex:  slbOperEnhHttpCachePurgeServerIndex,
		SlbOperEnhHttpCachePurgeServiceIndex: slbOperEnhHttpCachePurgeServiceIndex,
		Params:                               params,
	}
}

func (c *SlbOperEnhHttpCachePurgeTable) Name() string {
	return "SlbOperEnhHttpCachePurgeTable"
}

func (c *SlbOperEnhHttpCachePurgeTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbOperEnhHttpCachePurgeTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbOperEnhHttpCachePurgeTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbOperEnhHttpCachePurgeServerIndex).IsZero() &&
		reflect.ValueOf(c.SlbOperEnhHttpCachePurgeServiceIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbOperEnhHttpCachePurgeServerIndex) + "/" + fmt.Sprint(c.SlbOperEnhHttpCachePurgeServiceIndex)
}

type SlbOperEnhHttpCachePurgeTableParams struct {
	// The Http Cache purge server index.
	ServerIndex string `json:"ServerIndex,omitempty"`
	// The Http Cache purge service index.
	ServiceIndex int32 `json:"ServiceIndex,omitempty"`
	// Set object URL or all.
	URL string `json:"URL,omitempty"`
}

func (p SlbOperEnhHttpCachePurgeTableParams) iMABean() {}
