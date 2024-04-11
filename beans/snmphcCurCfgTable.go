package beans

import (
	"fmt"
	"reflect"
)

// SnmphcCurCfgTable The SNMP health check table in the current configuration block.
// Note:This mib is not supported for VX instance of Virtualization.
type SnmphcCurCfgTable struct {
	// The index of the SNMP health check.
	SnmphcCurCfgIndex int32
	Params            *SnmphcCurCfgTableParams
}

func NewSnmphcCurCfgTableList() *SnmphcCurCfgTable {
	return &SnmphcCurCfgTable{}
}

func NewSnmphcCurCfgTable(
	snmphcCurCfgIndex int32,
	params *SnmphcCurCfgTableParams,
) *SnmphcCurCfgTable {
	return &SnmphcCurCfgTable{
		SnmphcCurCfgIndex: snmphcCurCfgIndex,
		Params:            params,
	}
}

func (c *SnmphcCurCfgTable) Name() string {
	return "SnmphcCurCfgTable"
}

func (c *SnmphcCurCfgTable) GetParams() BeanType {
	return c.Params
}

func (c *SnmphcCurCfgTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SnmphcCurCfgTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SnmphcCurCfgIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SnmphcCurCfgIndex)
}

type SnmphcCurCfgTableInvert int32

const (
	SnmphcCurCfgTableInvert_Enabled     SnmphcCurCfgTableInvert = 1
	SnmphcCurCfgTableInvert_Disabled    SnmphcCurCfgTableInvert = 2
	SnmphcCurCfgTableInvert_Unsupported SnmphcCurCfgTableInvert = 2147483647
)

type SnmphcCurCfgTableUseWeight int32

const (
	SnmphcCurCfgTableUseWeight_Enabled     SnmphcCurCfgTableUseWeight = 1
	SnmphcCurCfgTableUseWeight_Disabled    SnmphcCurCfgTableUseWeight = 2
	SnmphcCurCfgTableUseWeight_Unsupported SnmphcCurCfgTableUseWeight = 2147483647
)

type SnmphcCurCfgTableParams struct {
	// The index of the SNMP health check.
	Index int32 `json:"Index,omitempty"`
	// The OID to be sent in the SNMP get request packet.
	Oid string `json:"Oid,omitempty"`
	// The community string to be used in the SNMP get request packet.
	CommString string `json:"CommString,omitempty"`
	// The content expected in the SNMP response packet. The content
	// specified can be either a string or an integer value.
	RcvContent string `json:"RcvContent,omitempty"`
	// When the invert option is enabled the health check will fail if
	// the response packet contains the value specified in the
	// receive content field.
	Invert SnmphcCurCfgTableInvert `json:"Invert,omitempty"`
	// When the weight option is enabled the real server weights are
	// adjusted dynamically based on the SNMP health check response.
	UseWeight SnmphcCurCfgTableUseWeight `json:"UseWeight,omitempty"`
}

func (p SnmphcCurCfgTableParams) iMABean() {}
