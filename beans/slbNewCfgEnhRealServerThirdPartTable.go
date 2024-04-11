package beans

import (
	"fmt"
	"reflect"
)

// SlbNewCfgEnhRealServerThirdPartTable The Third table of Real Server configuration in the new_config.
// Note:This mib is not supported for VX instance of Virtualization.
type SlbNewCfgEnhRealServerThirdPartTable struct {
	// The real server number
	SlbNewCfgEnhRealServerThirdPartIndex string
	Params                               *SlbNewCfgEnhRealServerThirdPartTableParams
}

func NewSlbNewCfgEnhRealServerThirdPartTableList() *SlbNewCfgEnhRealServerThirdPartTable {
	return &SlbNewCfgEnhRealServerThirdPartTable{}
}

func NewSlbNewCfgEnhRealServerThirdPartTable(
	slbNewCfgEnhRealServerThirdPartIndex string,
	params *SlbNewCfgEnhRealServerThirdPartTableParams,
) *SlbNewCfgEnhRealServerThirdPartTable {
	return &SlbNewCfgEnhRealServerThirdPartTable{
		SlbNewCfgEnhRealServerThirdPartIndex: slbNewCfgEnhRealServerThirdPartIndex,
		Params:                               params,
	}
}

func (c *SlbNewCfgEnhRealServerThirdPartTable) Name() string {
	return "SlbNewCfgEnhRealServerThirdPartTable"
}

func (c *SlbNewCfgEnhRealServerThirdPartTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbNewCfgEnhRealServerThirdPartTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbNewCfgEnhRealServerThirdPartTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbNewCfgEnhRealServerThirdPartIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbNewCfgEnhRealServerThirdPartIndex)
}

type SlbNewCfgEnhRealServerThirdPartTableParams struct {
	// The real server number
	Index string `json:"Index,omitempty"`
	// The OID to be sent in the SNMP get request packet.
	Oid string `json:"Oid,omitempty"`
	// The community string to be used in the SNMP get request packet.
	CommString string `json:"CommString,omitempty"`
	// The backup server number for this server.
	BackUp string `json:"BackUp,omitempty"`
	// The Advanced HC ID.
	HealthID string `json:"HealthID,omitempty"`
	// Critical connection threshold.
	CriticalConnThrsh uint32 `json:"CriticalConnThrsh,omitempty"`
	// High connection threshold.
	HighConnThrsh uint32 `json:"HighConnThrsh,omitempty"`
	// Upload bandwidth limit for WAN Link real server in Mbps.
	UploadBandWidth uint32 `json:"UploadBandWidth,omitempty"`
	// Download bandwidth limit for WAN Link real server in Mbps.
	DownloadBandWidth uint32 `json:"DownloadBandWidth,omitempty"`
}

func (p SlbNewCfgEnhRealServerThirdPartTableParams) iMABean() {}
