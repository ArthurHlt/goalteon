package beans

import (
	"fmt"
	"reflect"
)

// SlbCurCfgEnhRealServerThirdPartTable The second table of Real Server configuration in the current_config.
// Note:This mib is not supported on VX instance of Virtualization.
type SlbCurCfgEnhRealServerThirdPartTable struct {
	// The real server number
	SlbCurCfgEnhRealServerThirdPartIndex string
	Params                               *SlbCurCfgEnhRealServerThirdPartTableParams
}

func NewSlbCurCfgEnhRealServerThirdPartTableList() *SlbCurCfgEnhRealServerThirdPartTable {
	return &SlbCurCfgEnhRealServerThirdPartTable{}
}

func NewSlbCurCfgEnhRealServerThirdPartTable(
	slbCurCfgEnhRealServerThirdPartIndex string,
	params *SlbCurCfgEnhRealServerThirdPartTableParams,
) *SlbCurCfgEnhRealServerThirdPartTable {
	return &SlbCurCfgEnhRealServerThirdPartTable{
		SlbCurCfgEnhRealServerThirdPartIndex: slbCurCfgEnhRealServerThirdPartIndex,
		Params:                               params,
	}
}

func (c *SlbCurCfgEnhRealServerThirdPartTable) Name() string {
	return "SlbCurCfgEnhRealServerThirdPartTable"
}

func (c *SlbCurCfgEnhRealServerThirdPartTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbCurCfgEnhRealServerThirdPartTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbCurCfgEnhRealServerThirdPartTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbCurCfgEnhRealServerThirdPartIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbCurCfgEnhRealServerThirdPartIndex)
}

type SlbCurCfgEnhRealServerThirdPartTableParams struct {
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

func (p SlbCurCfgEnhRealServerThirdPartTableParams) iMABean() {}
