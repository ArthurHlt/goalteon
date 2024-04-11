package beans

import (
	"fmt"
	"reflect"
)

// OverlayCurCfgMapTable The mapping table of slb overlay,
// Note:This mib is not supported for VX instance of Virtualization.
type OverlayCurCfgMapTable struct {
	// Overlay tunnel mapping table index.
	OverlayCurCfgMapIndex string
	Params                *OverlayCurCfgMapTableParams
}

func NewOverlayCurCfgMapTableList() *OverlayCurCfgMapTable {
	return &OverlayCurCfgMapTable{}
}

func NewOverlayCurCfgMapTable(
	overlayCurCfgMapIndex string,
	params *OverlayCurCfgMapTableParams,
) *OverlayCurCfgMapTable {
	return &OverlayCurCfgMapTable{
		OverlayCurCfgMapIndex: overlayCurCfgMapIndex,
		Params:                params,
	}
}

func (c *OverlayCurCfgMapTable) Name() string {
	return "OverlayCurCfgMapTable"
}

func (c *OverlayCurCfgMapTable) GetParams() BeanType {
	return c.Params
}

func (c *OverlayCurCfgMapTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *OverlayCurCfgMapTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.OverlayCurCfgMapIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.OverlayCurCfgMapIndex)
}

type OverlayCurCfgMapTableParams struct {
	// Overlay tunnel mapping table index.
	Index string `json:"Index,omitempty"`
	// Inner packet Dest mac key for mapping table entry:
	// 		get a valid mac number.
	// 			ff:ff:ff:ff:ff:ff as (any) - only for last entry.
	Idmackey string `json:"Idmackey,omitempty"`
	// Outer packet source IP address for tunnel wrapper.
	// 	 	 Address should be 4-byte hexadecimal colon notation.
	// Valid IPv6 address should be in any of the following forms
	// xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx or
	// xxxx::xxxx:xxxx:xxxx:xxxx or ::xxxx
	Map6Osip string `json:"Map6Osip,omitempty"`
	// Outer packet source mac for tunnel wrapper.
	Osmac string `json:"Osmac,omitempty"`
	// Outer packet Dest IP address for tunnel wrapper..
	// 	 	 Address should be 4-byte hexadecimal colon notation.
	// Valid IPv6 address should be in any of the following forms
	// xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx or
	// xxxx::xxxx:xxxx:xxxx:xxxx or ::xxxx
	Map6Odip string `json:"Map6Odip,omitempty"`
	// Outer packet dest mac for tunnel wrapper.
	Odmac string `json:"Odmac,omitempty"`
	// Outer packet Alteon egress port:
	// 	0	- auto
	// 	1-max port	- port number
	Port uint32 `json:"Port,omitempty"`
	// Outer packet VLAN Id:
	// 	0	- auto
	// 	1-4090	- vlan id
	Vlan uint32 `json:"Vlan,omitempty"`
}

func (p OverlayCurCfgMapTableParams) iMABean() {}
