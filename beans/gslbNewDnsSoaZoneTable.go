package beans

import (
	"fmt"
	"reflect"
)

// GslbNewDnsSoaZoneTable DNS SOA new config zone table.
// Note:This mib is not supported for VX instance of Virtualization.
type GslbNewDnsSoaZoneTable struct {
	// DNS SOA new config zone entry alphanumeric index.
	GslbNewDnsSoaZoneId string
	Params              *GslbNewDnsSoaZoneTableParams
}

func NewGslbNewDnsSoaZoneTableList() *GslbNewDnsSoaZoneTable {
	return &GslbNewDnsSoaZoneTable{}
}

func NewGslbNewDnsSoaZoneTable(
	gslbNewDnsSoaZoneId string,
	params *GslbNewDnsSoaZoneTableParams,
) *GslbNewDnsSoaZoneTable {
	return &GslbNewDnsSoaZoneTable{
		GslbNewDnsSoaZoneId: gslbNewDnsSoaZoneId,
		Params:              params,
	}
}

func (c *GslbNewDnsSoaZoneTable) Name() string {
	return "GslbNewDnsSoaZoneTable"
}

func (c *GslbNewDnsSoaZoneTable) GetParams() BeanType {
	return c.Params
}

func (c *GslbNewDnsSoaZoneTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *GslbNewDnsSoaZoneTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.GslbNewDnsSoaZoneId).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.GslbNewDnsSoaZoneId)
}

type GslbNewDnsSoaZoneTableDelete int32

const (
	GslbNewDnsSoaZoneTableDelete_Other       GslbNewDnsSoaZoneTableDelete = 1
	GslbNewDnsSoaZoneTableDelete_Delete      GslbNewDnsSoaZoneTableDelete = 2
	GslbNewDnsSoaZoneTableDelete_Unsupported GslbNewDnsSoaZoneTableDelete = 2147483647
)

type GslbNewDnsSoaZoneTableParams struct {
	// DNS SOA new config zone entry alphanumeric index.
	Id string `json:"Id,omitempty"`
	// DNS SOA new config zone entry zone name to match.
	Name string `json:"Name,omitempty"`
	// DNS SOA new config zone entry primary nameserver for the zone.
	NameServ string `json:"NameServ,omitempty"`
	// DNS SOA new config zone entry primary responsible mail for the zone.
	RespMail string `json:"RespMail,omitempty"`
	// DNS SOA new config zone entry generated serial, read-only.
	Serial uint32 `json:"Serial,omitempty"`
	// DNS SOA new config zone entry delete operation by setting (2), (1) is ignored
	Delete GslbNewDnsSoaZoneTableDelete `json:"Delete,omitempty"`
}

func (p GslbNewDnsSoaZoneTableParams) iMABean() {}
