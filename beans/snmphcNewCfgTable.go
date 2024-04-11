package beans

import (
	"fmt"
	"reflect"
)

// SnmphcNewCfgTable The SNMP health check table in the new configuration block.
// Note:This mib is not supported for VX instance of Virtualization.
type SnmphcNewCfgTable struct {
	// The index of the SNMP health check.
	SnmphcNewCfgIndex int32
	Params            *SnmphcNewCfgTableParams
}

func NewSnmphcNewCfgTableList() *SnmphcNewCfgTable {
	return &SnmphcNewCfgTable{}
}

func NewSnmphcNewCfgTable(
	snmphcNewCfgIndex int32,
	params *SnmphcNewCfgTableParams,
) *SnmphcNewCfgTable {
	return &SnmphcNewCfgTable{
		SnmphcNewCfgIndex: snmphcNewCfgIndex,
		Params:            params,
	}
}

func (c *SnmphcNewCfgTable) Name() string {
	return "SnmphcNewCfgTable"
}

func (c *SnmphcNewCfgTable) GetParams() BeanType {
	return c.Params
}

func (c *SnmphcNewCfgTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SnmphcNewCfgTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SnmphcNewCfgIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SnmphcNewCfgIndex)
}

type SnmphcNewCfgTableInvert int32

const (
	SnmphcNewCfgTableInvert_Enabled     SnmphcNewCfgTableInvert = 1
	SnmphcNewCfgTableInvert_Disabled    SnmphcNewCfgTableInvert = 2
	SnmphcNewCfgTableInvert_Unsupported SnmphcNewCfgTableInvert = 2147483647
)

type SnmphcNewCfgTableDeleteHc int32

const (
	SnmphcNewCfgTableDeleteHc_Other       SnmphcNewCfgTableDeleteHc = 1
	SnmphcNewCfgTableDeleteHc_Delete      SnmphcNewCfgTableDeleteHc = 2
	SnmphcNewCfgTableDeleteHc_Unsupported SnmphcNewCfgTableDeleteHc = 2147483647
)

type SnmphcNewCfgTableUseWeight int32

const (
	SnmphcNewCfgTableUseWeight_Enabled     SnmphcNewCfgTableUseWeight = 1
	SnmphcNewCfgTableUseWeight_Disabled    SnmphcNewCfgTableUseWeight = 2
	SnmphcNewCfgTableUseWeight_Unsupported SnmphcNewCfgTableUseWeight = 2147483647
)

type SnmphcNewCfgTableParams struct {
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
	Invert SnmphcNewCfgTableInvert `json:"Invert,omitempty"`
	// When set to the value of 2 (delete), the entire row is deleted.
	// When read, other(1) is returned. Setting the value to anything
	// other than 2(delete) has no effect on the state of the row.
	DeleteHc SnmphcNewCfgTableDeleteHc `json:"DeleteHc,omitempty"`
	// When the weight option is enabled the real server weights are
	// adjusted dynamically based on the SNMP health check response.
	UseWeight SnmphcNewCfgTableUseWeight `json:"UseWeight,omitempty"`
}

func (p SnmphcNewCfgTableParams) iMABean() {}
