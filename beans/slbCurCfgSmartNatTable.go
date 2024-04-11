package beans

import (
	"fmt"
	"reflect"
)

// SlbCurCfgSmartNatTable The table of current Smart NAT configuration.
// Note:This mib is not supported on VX instance of Virtualization.
type SlbCurCfgSmartNatTable struct {
	// The SMART NAT id as an index.
	SlbCurCfgSmartNatIndex string
	Params                 *SlbCurCfgSmartNatTableParams
}

func NewSlbCurCfgSmartNatTableList() *SlbCurCfgSmartNatTable {
	return &SlbCurCfgSmartNatTable{}
}

func NewSlbCurCfgSmartNatTable(
	slbCurCfgSmartNatIndex string,
	params *SlbCurCfgSmartNatTableParams,
) *SlbCurCfgSmartNatTable {
	return &SlbCurCfgSmartNatTable{
		SlbCurCfgSmartNatIndex: slbCurCfgSmartNatIndex,
		Params:                 params,
	}
}

func (c *SlbCurCfgSmartNatTable) Name() string {
	return "SlbCurCfgSmartNatTable"
}

func (c *SlbCurCfgSmartNatTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbCurCfgSmartNatTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbCurCfgSmartNatTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbCurCfgSmartNatIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbCurCfgSmartNatIndex)
}

type SlbCurCfgSmartNatTableType int32

const (
	SlbCurCfgSmartNatTableType_Nonat       SlbCurCfgSmartNatTableType = 0
	SlbCurCfgSmartNatTableType_Static      SlbCurCfgSmartNatTableType = 1
	SlbCurCfgSmartNatTableType_Dynamic     SlbCurCfgSmartNatTableType = 2
	SlbCurCfgSmartNatTableType_Unsupported SlbCurCfgSmartNatTableType = 2147483647
)

type SlbCurCfgSmartNatTableIpVer int32

const (
	SlbCurCfgSmartNatTableIpVer_Ipv4        SlbCurCfgSmartNatTableIpVer = 0
	SlbCurCfgSmartNatTableIpVer_Ipv6        SlbCurCfgSmartNatTableIpVer = 1
	SlbCurCfgSmartNatTableIpVer_Unsupported SlbCurCfgSmartNatTableIpVer = 2147483647
)

type SlbCurCfgSmartNatTableMode int32

const (
	SlbCurCfgSmartNatTableMode_Address     SlbCurCfgSmartNatTableMode = 0
	SlbCurCfgSmartNatTableMode_Nwclass     SlbCurCfgSmartNatTableMode = 1
	SlbCurCfgSmartNatTableMode_None        SlbCurCfgSmartNatTableMode = 2
	SlbCurCfgSmartNatTableMode_Unsupported SlbCurCfgSmartNatTableMode = 2147483647
)

type SlbCurCfgSmartNatTableDnatMode int32

const (
	SlbCurCfgSmartNatTableDnatMode_Address     SlbCurCfgSmartNatTableDnatMode = 0
	SlbCurCfgSmartNatTableDnatMode_Nwclass     SlbCurCfgSmartNatTableDnatMode = 1
	SlbCurCfgSmartNatTableDnatMode_None        SlbCurCfgSmartNatTableDnatMode = 2
	SlbCurCfgSmartNatTableDnatMode_Unsupported SlbCurCfgSmartNatTableDnatMode = 2147483647
)

type SlbCurCfgSmartNatTableDnatPersist int32

const (
	SlbCurCfgSmartNatTableDnatPersist_None        SlbCurCfgSmartNatTableDnatPersist = 1
	SlbCurCfgSmartNatTableDnatPersist_Host        SlbCurCfgSmartNatTableDnatPersist = 2
	SlbCurCfgSmartNatTableDnatPersist_Client      SlbCurCfgSmartNatTableDnatPersist = 3
	SlbCurCfgSmartNatTableDnatPersist_Unsupported SlbCurCfgSmartNatTableDnatPersist = 2147483647
)

type SlbCurCfgSmartNatTableParams struct {
	// The SMART NAT id as an index.
	Index string `json:"Index,omitempty"`
	// Descriptive name for Smart NAT entry.
	Name string `json:"Name,omitempty"`
	// Smart NAT type.
	Type SlbCurCfgSmartNatTableType `json:"Type,omitempty"`
	// IP version for Smart NAT entry.
	IpVer SlbCurCfgSmartNatTableIpVer `json:"IpVer,omitempty"`
	// Smart NAT local mode.
	Mode SlbCurCfgSmartNatTableMode `json:"Mode,omitempty"`
	// Local ipv4 address of the Smart NAT entry. This is valid only when ip version is set to v4.
	LocalIpV4 string `json:"LocalIpV4,omitempty"`
	// Local ipv4 network mask of the Smart NAT entry. This is valid only when ip version is set to v4.
	LocalIpV4Mask string `json:"LocalIpV4Mask,omitempty"`
	// The Local IPV6 address of the Smart NAT entry
	// Address should be 4-byte hexadecimal colon notation. Valid IPv6 address should be in
	// any of the following forms xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx or xxxx::xxxx:xxxx:xxxx:xxxx or ::xxxx
	// This is valid only when ip version is set to ipv6.
	LocalIpV6 string `json:"LocalIpV6,omitempty"`
	// The prefix length associated with the local IPV6 address of the Smart NAT entry.
	// This is valid only when ip version is set to ipv6.
	LocalIpV6Mask uint32 `json:"LocalIpV6Mask,omitempty"`
	// Local network class name.
	LocalNwclss string `json:"LocalNwclss,omitempty"`
	// NAT address mode for this Smart NAT entry
	DnatMode SlbCurCfgSmartNatTableDnatMode `json:"DnatMode,omitempty"`
	// NAT ipv4 address of the Smart NAT entry. This is valid only when ip version is set to ipv4.
	DnatIpV4 string `json:"DnatIpV4,omitempty"`
	// NAT ipv4 network mask of the Smart NAT entry. This is valid only when ip version is set to ipv4.
	DnatIpV4Mask string `json:"DnatIpV4Mask,omitempty"`
	// NAT IPV6 address of the Smart NAT entry
	// Address should be 4-byte hexadecimal colon notation. Valid IPv6 address should be in
	// any of the following forms xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx or xxxx::xxxx:xxxx:xxxx:xxxx or ::xxxx
	// This is valid only when ip version is set to ipv6.
	DnatIpV6 string `json:"DnatIpV6,omitempty"`
	// The prefix length associated with the NAT IPV6 address of the Smart NAT entry.
	// This is valid only when ip version is set to ipv6.
	DnatIpV6Mask uint32 `json:"DnatIpV6Mask,omitempty"`
	// NAT Network class name for the SMART NAT entry.
	DnatNwclss string `json:"DnatNwclss,omitempty"`
	// Wanlink Server name for this Smart NAT entry.
	WanLink string `json:"WanLink,omitempty"`
	// Persistence type for the Smart NAT entry. This is valid only for type Dynamic NAT.
	DnatPersist SlbCurCfgSmartNatTableDnatPersist `json:"DnatPersist,omitempty"`
}

func (p SlbCurCfgSmartNatTableParams) iMABean() {}
