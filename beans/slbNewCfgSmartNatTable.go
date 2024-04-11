package beans

import (
	"fmt"
	"reflect"
)

// SlbNewCfgSmartNatTable The table for configuring Smart NAT.
// Note:This mib is not supported on VX instance of Virtualization.
type SlbNewCfgSmartNatTable struct {
	// The Smart NAT id as an index.
	SlbNewCfgSmartNatIndex string
	Params                 *SlbNewCfgSmartNatTableParams
}

func NewSlbNewCfgSmartNatTableList() *SlbNewCfgSmartNatTable {
	return &SlbNewCfgSmartNatTable{}
}

func NewSlbNewCfgSmartNatTable(
	slbNewCfgSmartNatIndex string,
	params *SlbNewCfgSmartNatTableParams,
) *SlbNewCfgSmartNatTable {
	return &SlbNewCfgSmartNatTable{
		SlbNewCfgSmartNatIndex: slbNewCfgSmartNatIndex,
		Params:                 params,
	}
}

func (c *SlbNewCfgSmartNatTable) Name() string {
	return "SlbNewCfgSmartNatTable"
}

func (c *SlbNewCfgSmartNatTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbNewCfgSmartNatTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbNewCfgSmartNatTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbNewCfgSmartNatIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbNewCfgSmartNatIndex)
}

type SlbNewCfgSmartNatTableType int32

const (
	SlbNewCfgSmartNatTableType_Nonat       SlbNewCfgSmartNatTableType = 0
	SlbNewCfgSmartNatTableType_Static      SlbNewCfgSmartNatTableType = 1
	SlbNewCfgSmartNatTableType_Dynamic     SlbNewCfgSmartNatTableType = 2
	SlbNewCfgSmartNatTableType_Unsupported SlbNewCfgSmartNatTableType = 2147483647
)

type SlbNewCfgSmartNatTableIpVer int32

const (
	SlbNewCfgSmartNatTableIpVer_Ipv4        SlbNewCfgSmartNatTableIpVer = 0
	SlbNewCfgSmartNatTableIpVer_Ipv6        SlbNewCfgSmartNatTableIpVer = 1
	SlbNewCfgSmartNatTableIpVer_Unsupported SlbNewCfgSmartNatTableIpVer = 2147483647
)

type SlbNewCfgSmartNatTableMode int32

const (
	SlbNewCfgSmartNatTableMode_Address     SlbNewCfgSmartNatTableMode = 0
	SlbNewCfgSmartNatTableMode_Nwclass     SlbNewCfgSmartNatTableMode = 1
	SlbNewCfgSmartNatTableMode_None        SlbNewCfgSmartNatTableMode = 2
	SlbNewCfgSmartNatTableMode_Unsupported SlbNewCfgSmartNatTableMode = 2147483647
)

type SlbNewCfgSmartNatTableDnatMode int32

const (
	SlbNewCfgSmartNatTableDnatMode_Address     SlbNewCfgSmartNatTableDnatMode = 0
	SlbNewCfgSmartNatTableDnatMode_Nwclass     SlbNewCfgSmartNatTableDnatMode = 1
	SlbNewCfgSmartNatTableDnatMode_None        SlbNewCfgSmartNatTableDnatMode = 2
	SlbNewCfgSmartNatTableDnatMode_Unsupported SlbNewCfgSmartNatTableDnatMode = 2147483647
)

type SlbNewCfgSmartNatTableDel int32

const (
	SlbNewCfgSmartNatTableDel_Other       SlbNewCfgSmartNatTableDel = 1
	SlbNewCfgSmartNatTableDel_Delete      SlbNewCfgSmartNatTableDel = 2
	SlbNewCfgSmartNatTableDel_Unsupported SlbNewCfgSmartNatTableDel = 2147483647
)

type SlbNewCfgSmartNatTableDnatPersist int32

const (
	SlbNewCfgSmartNatTableDnatPersist_None        SlbNewCfgSmartNatTableDnatPersist = 1
	SlbNewCfgSmartNatTableDnatPersist_Host        SlbNewCfgSmartNatTableDnatPersist = 2
	SlbNewCfgSmartNatTableDnatPersist_Client      SlbNewCfgSmartNatTableDnatPersist = 3
	SlbNewCfgSmartNatTableDnatPersist_Unsupported SlbNewCfgSmartNatTableDnatPersist = 2147483647
)

type SlbNewCfgSmartNatTableParams struct {
	// The Smart NAT id as an index.
	Index string `json:"Index,omitempty"`
	// Descriptive name for Smart NAT entry.
	Name string `json:"Name,omitempty"`
	// Smart NAT type.
	Type SlbNewCfgSmartNatTableType `json:"Type,omitempty"`
	// IP version for Smart NAT entry.
	IpVer SlbNewCfgSmartNatTableIpVer `json:"IpVer,omitempty"`
	// Smart NAT local address mode.
	Mode SlbNewCfgSmartNatTableMode `json:"Mode,omitempty"`
	// Local ipv4 address of the Smart NAT entry. This is valid only when ip version is set to v4.
	LocalIpV4 string `json:"LocalIpV4,omitempty"`
	// Local ipv4 network mask of the Smart NAT entry. This is valid only ip version is set to v4.
	LocalIpV4Mask string `json:"LocalIpV4Mask,omitempty"`
	// The Local IPV6 address of the Smart NAT entry
	// Address should be 4-byte hexadecimal colon notation. Valid IPv6 address should be in
	// any of the following forms xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx or xxxx::xxxx:xxxx:xxxx:xxxx or ::xxxx
	// This is valid only when ip version is set to ipv6.
	LocalIpV6 string `json:"LocalIpV6,omitempty"`
	// The prefix length associated with the local IPV6 address of the Smart NAT entry.
	// This is valid only when ip version is set to ipv6.
	LocalIpV6Mask uint32 `json:"LocalIpV6Mask,omitempty"`
	// Local Network class name. This is valid only when mode is nwclass.
	LocalNwclss string `json:"LocalNwclss,omitempty"`
	// NAT address mode for this Smart NAT entry.
	DnatMode SlbNewCfgSmartNatTableDnatMode `json:"DnatMode,omitempty"`
	// NAT ipv4 address of the Smart NAT entry.
	// This is valid only ip version is set to ipv4.
	DnatIpV4 string `json:"DnatIpV4,omitempty"`
	// NAT ipv4 network mask of the Smart NAT entry.
	// This is valid only when ip version is set to ipv4.
	DnatIpV4Mask string `json:"DnatIpV4Mask,omitempty"`
	// NAT IPV6 address of the Smart NAT entry
	// Address should be 4-byte hexadecimal colon notation. Valid IPv6 address should be in
	// any of the following forms xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx or xxxx::xxxx:xxxx:xxxx:xxxx or ::xxxx
	// This is valid only when ip version is set to ipv6.
	DnatIpV6 string `json:"DnatIpV6,omitempty"`
	// The prefix length associated with the NAT IPV6 address of the SMART nat entry.
	// This is valid only when ip version is set to ipv6.
	DnatIpV6Mask uint32 `json:"DnatIpV6Mask,omitempty"`
	// NAT Network class name.
	DnatNwclss string `json:"DnatNwclss,omitempty"`
	// Wanlink Server name for this Smart NAT entry.
	WanLink string `json:"WanLink,omitempty"`
	// Delete Smart Nat entry.
	Del SlbNewCfgSmartNatTableDel `json:"Del,omitempty"`
	// Persistence type for this Smart NAT entry. This is valid only for type Dynamic NAT.
	DnatPersist SlbNewCfgSmartNatTableDnatPersist `json:"DnatPersist,omitempty"`
}

func (p SlbNewCfgSmartNatTableParams) iMABean() {}
