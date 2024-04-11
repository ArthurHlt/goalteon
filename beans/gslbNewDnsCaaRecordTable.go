package beans

import (
	"fmt"
	"reflect"
)

// GslbNewDnsCaaRecordTable DNS CAA config record table.
// Note:This mib is not supported for VX instance of Virtualization.
type GslbNewDnsCaaRecordTable struct {
	// DNS CAA config record entry, alphanumeric index.
	GslbNewDnsCaaRecordId string
	Params                *GslbNewDnsCaaRecordTableParams
}

func NewGslbNewDnsCaaRecordTableList() *GslbNewDnsCaaRecordTable {
	return &GslbNewDnsCaaRecordTable{}
}

func NewGslbNewDnsCaaRecordTable(
	gslbNewDnsCaaRecordId string,
	params *GslbNewDnsCaaRecordTableParams,
) *GslbNewDnsCaaRecordTable {
	return &GslbNewDnsCaaRecordTable{
		GslbNewDnsCaaRecordId: gslbNewDnsCaaRecordId,
		Params:                params,
	}
}

func (c *GslbNewDnsCaaRecordTable) Name() string {
	return "GslbNewDnsCaaRecordTable"
}

func (c *GslbNewDnsCaaRecordTable) GetParams() BeanType {
	return c.Params
}

func (c *GslbNewDnsCaaRecordTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *GslbNewDnsCaaRecordTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.GslbNewDnsCaaRecordId).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.GslbNewDnsCaaRecordId)
}

type GslbNewDnsCaaRecordTableType int32

const (
	GslbNewDnsCaaRecordTableType_Issue       GslbNewDnsCaaRecordTableType = 0
	GslbNewDnsCaaRecordTableType_Issuewild   GslbNewDnsCaaRecordTableType = 1
	GslbNewDnsCaaRecordTableType_Iodef       GslbNewDnsCaaRecordTableType = 2
	GslbNewDnsCaaRecordTableType_Unsupported GslbNewDnsCaaRecordTableType = 2147483647
)

type GslbNewDnsCaaRecordTableDelete int32

const (
	GslbNewDnsCaaRecordTableDelete_Other       GslbNewDnsCaaRecordTableDelete = 1
	GslbNewDnsCaaRecordTableDelete_Delete      GslbNewDnsCaaRecordTableDelete = 2
	GslbNewDnsCaaRecordTableDelete_Unsupported GslbNewDnsCaaRecordTableDelete = 2147483647
)

type GslbNewDnsCaaRecordTableParams struct {
	// DNS CAA config record entry, alphanumeric index.
	Id string `json:"Id,omitempty"`
	// DNS CAA config record entry, domain name to match.
	DomainName string `json:"DomainName,omitempty"`
	// DNS CAA config record entry, record type.
	Type GslbNewDnsCaaRecordTableType `json:"Type,omitempty"`
	// DNS CAA config record entry, record value.
	Value string `json:"Value,omitempty"`
	// DNS CAA config record entry, record ttl.
	Ttl uint64 `json:"Ttl,omitempty"`
	// DNS CAA config record entry delete operation by setting (2), (1) is ignored
	Delete GslbNewDnsCaaRecordTableDelete `json:"Delete,omitempty"`
}

func (p GslbNewDnsCaaRecordTableParams) iMABean() {}
