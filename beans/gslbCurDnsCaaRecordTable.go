package beans

import (
	"fmt"
	"reflect"
)

// GslbCurDnsCaaRecordTable DNS CAA config Record table.
// Note:This mib is not supported for VX instance of Virtualization.
type GslbCurDnsCaaRecordTable struct {
	// DNS CAA config Record entry alphanumeric index.
	GslbCurDnsCaaRecordId string
	Params                *GslbCurDnsCaaRecordTableParams
}

func NewGslbCurDnsCaaRecordTableList() *GslbCurDnsCaaRecordTable {
	return &GslbCurDnsCaaRecordTable{}
}

func NewGslbCurDnsCaaRecordTable(
	gslbCurDnsCaaRecordId string,
	params *GslbCurDnsCaaRecordTableParams,
) *GslbCurDnsCaaRecordTable {
	return &GslbCurDnsCaaRecordTable{
		GslbCurDnsCaaRecordId: gslbCurDnsCaaRecordId,
		Params:                params,
	}
}

func (c *GslbCurDnsCaaRecordTable) Name() string {
	return "GslbCurDnsCaaRecordTable"
}

func (c *GslbCurDnsCaaRecordTable) GetParams() BeanType {
	return c.Params
}

func (c *GslbCurDnsCaaRecordTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *GslbCurDnsCaaRecordTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.GslbCurDnsCaaRecordId).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.GslbCurDnsCaaRecordId)
}

type GslbCurDnsCaaRecordTableType int32

const (
	GslbCurDnsCaaRecordTableType_Issue       GslbCurDnsCaaRecordTableType = 0
	GslbCurDnsCaaRecordTableType_Issuewild   GslbCurDnsCaaRecordTableType = 1
	GslbCurDnsCaaRecordTableType_Iodef       GslbCurDnsCaaRecordTableType = 2
	GslbCurDnsCaaRecordTableType_Unsupported GslbCurDnsCaaRecordTableType = 2147483647
)

type GslbCurDnsCaaRecordTableParams struct {
	// DNS CAA config Record entry alphanumeric index.
	Id string `json:"Id,omitempty"`
	// DNS CAA config Record entry domain name.
	DomainName string `json:"DomainName,omitempty"`
	// DNS CAA config Record entry Record Type.
	Type GslbCurDnsCaaRecordTableType `json:"Type,omitempty"`
	// DNS CAA config Record value.
	Value string `json:"Value,omitempty"`
	// DNS CAA config Record TTL.
	Ttl uint64 `json:"Ttl,omitempty"`
}

func (p GslbCurDnsCaaRecordTableParams) iMABean() {}
