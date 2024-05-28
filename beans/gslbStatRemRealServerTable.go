package beans

import (
	"fmt"
	"reflect"
)

// GslbStatRemRealServerTable The GSLB remote real server statistics table.
// Note:This mib is not supported for VX instance of Virtualization.
type GslbStatRemRealServerTable struct {
	// The server number that identifies the remote real server.
	GslbStatRemRealServerIndex int32
	Params                     *GslbStatRemRealServerTableParams
}

func NewGslbStatRemRealServerTableList() *GslbStatRemRealServerTable {
	return &GslbStatRemRealServerTable{}
}

func NewGslbStatRemRealServerTable(
	gslbStatRemRealServerIndex int32,
	params *GslbStatRemRealServerTableParams,
) *GslbStatRemRealServerTable {
	return &GslbStatRemRealServerTable{
		GslbStatRemRealServerIndex: gslbStatRemRealServerIndex,
		Params:                     params,
	}
}

func (c *GslbStatRemRealServerTable) Name() string {
	return "GslbStatRemRealServerTable"
}

func (c *GslbStatRemRealServerTable) GetParams() BeanType {
	return c.Params
}

func (c *GslbStatRemRealServerTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *GslbStatRemRealServerTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.GslbStatRemRealServerIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.GslbStatRemRealServerIndex)
}

type GslbStatRemRealServerTableParams struct {
	// The server number that identifies the remote real server.
	Index int32 `json:"Index,omitempty"`
	// The number of DNS handoffs by the remote real server.
	DnsHandoffs uint32 `json:"DnsHandoffs,omitempty"`
	// The number of HTTP redirections by the remote real server.
	HttpRedirs uint32 `json:"HttpRedirs,omitempty"`
}

func (p GslbStatRemRealServerTableParams) iMABean() {}
