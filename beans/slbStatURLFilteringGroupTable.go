package beans

import (
	"fmt"
	"reflect"
)

// SlbStatURLFilteringGroupTable The statistics table of URL Filtering.
// Note:This mib is not supported for VX instance of virtualization.
type SlbStatURLFilteringGroupTable struct {
	// URL Filtering Sub Category count statistics.
	SlbStatURLFilteringSubCategoryIndex int32
	Params                              *SlbStatURLFilteringGroupTableParams
}

func NewSlbStatURLFilteringGroupTableList() *SlbStatURLFilteringGroupTable {
	return &SlbStatURLFilteringGroupTable{}
}

func NewSlbStatURLFilteringGroupTable(
	slbStatURLFilteringSubCategoryIndex int32,
	params *SlbStatURLFilteringGroupTableParams,
) *SlbStatURLFilteringGroupTable {
	return &SlbStatURLFilteringGroupTable{
		SlbStatURLFilteringSubCategoryIndex: slbStatURLFilteringSubCategoryIndex,
		Params:                              params,
	}
}

func (c *SlbStatURLFilteringGroupTable) Name() string {
	return "SlbStatURLFilteringGroupTable"
}

func (c *SlbStatURLFilteringGroupTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbStatURLFilteringGroupTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbStatURLFilteringGroupTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbStatURLFilteringSubCategoryIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbStatURLFilteringSubCategoryIndex)
}

type SlbStatURLFilteringGroupTableParams struct {
	// URL Filtering Sub Category count statistics.
	SubCategoryIndex int32 `json:"SubCategoryIndex,omitempty"`
	// URL Filtering Sub Category statistics.
	SubCategoryName string `json:"SubCategoryName,omitempty"`
	// URL Filtering Category statistics.
	CategoryName string `json:"CategoryName,omitempty"`
	// URL Filtering Sub Category count statistics.
	SubCategoryCount int32 `json:"SubCategoryCount,omitempty"`
}

func (p SlbStatURLFilteringGroupTableParams) iMABean() {}
