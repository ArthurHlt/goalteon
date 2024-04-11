package beans

import (
	"fmt"
	"reflect"
)

// SlbEnhGroupRealServersTable The table of real server per group.
// Note:This mib is not supported for VX instance of Virtualization.
type SlbEnhGroupRealServersTable struct {
	// The real server group index.
	SlbEnhRealServerGroupIndex string
	// The real server index.
	SlbEnhGroupRealServerIndex string
	Params                     *SlbEnhGroupRealServersTableParams
}

func NewSlbEnhGroupRealServersTableList() *SlbEnhGroupRealServersTable {
	return &SlbEnhGroupRealServersTable{}
}

func NewSlbEnhGroupRealServersTable(
	slbEnhRealServerGroupIndex string,
	slbEnhGroupRealServerIndex string,
	params *SlbEnhGroupRealServersTableParams,
) *SlbEnhGroupRealServersTable {
	return &SlbEnhGroupRealServersTable{
		SlbEnhRealServerGroupIndex: slbEnhRealServerGroupIndex,
		SlbEnhGroupRealServerIndex: slbEnhGroupRealServerIndex,
		Params:                     params,
	}
}

func (c *SlbEnhGroupRealServersTable) Name() string {
	return "SlbEnhGroupRealServersTable"
}

func (c *SlbEnhGroupRealServersTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbEnhGroupRealServersTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbEnhGroupRealServersTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbEnhRealServerGroupIndex).IsZero() &&
		reflect.ValueOf(c.SlbEnhGroupRealServerIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbEnhRealServerGroupIndex) + "/" + fmt.Sprint(c.SlbEnhGroupRealServerIndex)
}

type SlbEnhGroupRealServersTableParams struct {
	// The real server group index.
	RealServerGroupIndex string `json:"RealServerGroupIndex,omitempty"`
	// The real server index.
	Index string `json:"Index,omitempty"`
}

func (p SlbEnhGroupRealServersTableParams) iMABean() {}
