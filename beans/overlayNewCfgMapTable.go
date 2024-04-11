package beans

import (
	"fmt"
	"reflect"
)

// OverlayNewCfgMapTable The mapping table of slb overlay,
// Note:This mib is not supported for VX instance of Virtualization.
type OverlayNewCfgMapTable struct {
	// Overlay tunnel mapping table index.
	OverlayNewCfgMapIndex string
	Params                *OverlayNewCfgMapTableParams
}

func NewOverlayNewCfgMapTableList() *OverlayNewCfgMapTable {
	return &OverlayNewCfgMapTable{}
}

func NewOverlayNewCfgMapTable(
	overlayNewCfgMapIndex string,
	params *OverlayNewCfgMapTableParams,
) *OverlayNewCfgMapTable {
	return &OverlayNewCfgMapTable{
		OverlayNewCfgMapIndex: overlayNewCfgMapIndex,
		Params:                params,
	}
}

func (c *OverlayNewCfgMapTable) Name() string {
	return "OverlayNewCfgMapTable"
}

func (c *OverlayNewCfgMapTable) GetParams() BeanType {
	return c.Params
}

func (c *OverlayNewCfgMapTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *OverlayNewCfgMapTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.OverlayNewCfgMapIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.OverlayNewCfgMapIndex)
}

type OverlayNewCfgMapTableDelete int32

const (
	OverlayNewCfgMapTableDelete_Other       OverlayNewCfgMapTableDelete = 1
	OverlayNewCfgMapTableDelete_Delete      OverlayNewCfgMapTableDelete = 2
	OverlayNewCfgMapTableDelete_All         OverlayNewCfgMapTableDelete = 3
	OverlayNewCfgMapTableDelete_Unsupported OverlayNewCfgMapTableDelete = 2147483647
)

type OverlayNewCfgMapTableParams struct {
	// Overlay tunnel mapping table index.
	Index string `json:"Index,omitempty"`
	// Inner packet Dest mac key for mapping table entry:
	// 	usage:
	// 		set a valid mac number.
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
	// Outer packet Dest IP address for tunnel wrapper.
	// 	 	 Address should be 4-byte hexadecimal colon notation.
	// Valid IPv6 address should be in any of the following forms
	// xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx or
	// xxxx::xxxx:xxxx:xxxx:xxxx or ::xxxx
	Map6Odip string `json:"Map6Odip,omitempty"`
	// Outer packet dest mac for tunnel wrapper.
	Odmac string `json:"Odmac,omitempty"`
	// Outer packet Alteon egress port:
	// usage:
	// 	0	- auto
	// 	1-max port	- port number
	Port uint32 `json:"Port,omitempty"`
	// Outer packet VLAN Id:
	// usage:
	// 	0	- auto
	// 	1-4090	- vlan id
	Vlan uint32 `json:"Vlan,omitempty"`
	// By setting the value to delete(2), Delete a specific overlay mapping table entry.
	Delete OverlayNewCfgMapTableDelete `json:"Delete,omitempty"`
}

func (p OverlayNewCfgMapTableParams) iMABean() {}
