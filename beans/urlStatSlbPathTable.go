package beans

import (
	"fmt"
	"reflect"
)

// UrlStatSlbPathTable The statistics table of instances that the URL path matched.
// Note:This mib is not supported for VX instance of virtualization.
type UrlStatSlbPathTable struct {
	// The URL path table index.
	UrlStatSlbPathIndex int32
	Params              *UrlStatSlbPathTableParams
}

func NewUrlStatSlbPathTableList() *UrlStatSlbPathTable {
	return &UrlStatSlbPathTable{}
}

func NewUrlStatSlbPathTable(
	urlStatSlbPathIndex int32,
	params *UrlStatSlbPathTableParams,
) *UrlStatSlbPathTable {
	return &UrlStatSlbPathTable{
		UrlStatSlbPathIndex: urlStatSlbPathIndex,
		Params:              params,
	}
}

func (c *UrlStatSlbPathTable) Name() string {
	return "UrlStatSlbPathTable"
}

func (c *UrlStatSlbPathTable) GetParams() BeanType {
	return c.Params
}

func (c *UrlStatSlbPathTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *UrlStatSlbPathTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.UrlStatSlbPathIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.UrlStatSlbPathIndex)
}

type UrlStatSlbPathTableParams struct {
	// The URL path table index.
	Index int32 `json:"Index,omitempty"`
	// The number of instances that are load-balanced due to match of
	// the particular URL path.
	Hits uint32 `json:"Hits,omitempty"`
}

func (p UrlStatSlbPathTableParams) iMABean() {}
