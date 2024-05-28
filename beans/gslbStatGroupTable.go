package beans

import (
	"fmt"
	"reflect"
)

// GslbStatGroupTable The GSLB group statistics table.
// Note:This mib is not supported for VX instance of Virtualization.
type GslbStatGroupTable struct {
	// The group number that identifies the real server group.
	GslbStatGroupIndex int32
	Params             *GslbStatGroupTableParams
}

func NewGslbStatGroupTableList() *GslbStatGroupTable {
	return &GslbStatGroupTable{}
}

func NewGslbStatGroupTable(
	gslbStatGroupIndex int32,
	params *GslbStatGroupTableParams,
) *GslbStatGroupTable {
	return &GslbStatGroupTable{
		GslbStatGroupIndex: gslbStatGroupIndex,
		Params:             params,
	}
}

func (c *GslbStatGroupTable) Name() string {
	return "GslbStatGroupTable"
}

func (c *GslbStatGroupTable) GetParams() BeanType {
	return c.Params
}

func (c *GslbStatGroupTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *GslbStatGroupTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.GslbStatGroupIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.GslbStatGroupIndex)
}

type GslbStatGroupTableParams struct {
	// The group number that identifies the real server group.
	Index int32 `json:"Index,omitempty"`
	// The total number of DNS hand-offs sent to the remote real servers in
	// the group.
	DnsHandoffs uint32 `json:"DnsHandoffs,omitempty"`
	// The total number of HTTP redirects sent to the remote real servers in
	// the group.
	HttpRedirs uint32 `json:"HttpRedirs,omitempty"`
}

func (p GslbStatGroupTableParams) iMABean() {}
