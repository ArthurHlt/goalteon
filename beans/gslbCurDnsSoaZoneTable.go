package beans

import (
	"fmt"
	"reflect"
)

// GslbCurDnsSoaZoneTable DNS SOA current config zone table.
// Note:This mib is not supported for VX instance of Virtualization.
type GslbCurDnsSoaZoneTable struct {
	// DNS SOA current config zone entry alphanumeric index.
	GslbCurDnsSoaZoneId string
	Params              *GslbCurDnsSoaZoneTableParams
}

func NewGslbCurDnsSoaZoneTableList() *GslbCurDnsSoaZoneTable {
	return &GslbCurDnsSoaZoneTable{}
}

func NewGslbCurDnsSoaZoneTable(
	gslbCurDnsSoaZoneId string,
	params *GslbCurDnsSoaZoneTableParams,
) *GslbCurDnsSoaZoneTable {
	return &GslbCurDnsSoaZoneTable{
		GslbCurDnsSoaZoneId: gslbCurDnsSoaZoneId,
		Params:              params,
	}
}

func (c *GslbCurDnsSoaZoneTable) Name() string {
	return "GslbCurDnsSoaZoneTable"
}

func (c *GslbCurDnsSoaZoneTable) GetParams() BeanType {
	return c.Params
}

func (c *GslbCurDnsSoaZoneTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *GslbCurDnsSoaZoneTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.GslbCurDnsSoaZoneId).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.GslbCurDnsSoaZoneId)
}

type GslbCurDnsSoaZoneTableParams struct {
	// DNS SOA current config zone entry alphanumeric index.
	Id string `json:"Id,omitempty"`
	// DNS SOA current config zone entry zone name to match.
	Name string `json:"Name,omitempty"`
	// DNS SOA current config zone entry primary nameserver for the zone.
	NameServ string `json:"NameServ,omitempty"`
	// DNS SOA current config zone entry primary resposible mail for the zone.
	RespMail string `json:"RespMail,omitempty"`
	// DNS SOA current config zone entry generated serial for any change in the zone.
	Serial uint32 `json:"Serial,omitempty"`
}

func (p GslbCurDnsSoaZoneTableParams) iMABean() {}
