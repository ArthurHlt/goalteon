package beans

import (
	"fmt"
	"reflect"
)

// SlbGslbProxInfoTable Proximity view table.
// Note:This mib is not supported on VX instance of Virtualization.
type SlbGslbProxInfoTable struct {
	// The Subnet index.
	SlbGslbSubnetIndex int32
	Params             *SlbGslbProxInfoTableParams
}

func NewSlbGslbProxInfoTableList() *SlbGslbProxInfoTable {
	return &SlbGslbProxInfoTable{}
}

func NewSlbGslbProxInfoTable(
	slbGslbSubnetIndex int32,
	params *SlbGslbProxInfoTableParams,
) *SlbGslbProxInfoTable {
	return &SlbGslbProxInfoTable{
		SlbGslbSubnetIndex: slbGslbSubnetIndex,
		Params:             params,
	}
}

func (c *SlbGslbProxInfoTable) Name() string {
	return "SlbGslbProxInfoTable"
}

func (c *SlbGslbProxInfoTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbGslbProxInfoTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbGslbProxInfoTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbGslbSubnetIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbGslbSubnetIndex)
}

type SlbGslbProxInfoTableParams struct {
	// The Subnet index.
	SubnetIndex int32 `json:"SubnetIndex,omitempty"`
	// Subnet/Destination ip.
	SubnetIp string `json:"SubnetIp,omitempty"`
	// Rank 1 Nhr ip.
	NhrRank1 string `json:"NhrRank1,omitempty"`
	// Rank 2 Nhr ip.
	NhrRank2 string `json:"NhrRank2,omitempty"`
	// Rank 3 Nhr ip.
	NhrRank3 string `json:"NhrRank3,omitempty"`
	// The roundtrip time for Rank1 Nhr Ip.
	NhrRTT1 string `json:"NhrRTT1,omitempty"`
	// The roundtrip time for Rank2 Nhr Ip.
	NhrRTT2 string `json:"NhrRTT2,omitempty"`
	// The roundtrip time for Rank3 Nhr Ip.
	NhrRTT3 string `json:"NhrRTT3,omitempty"`
	// Number of Hops to reach Destination via Rank1 NHR
	NhrHOP1 int32 `json:"NhrHOP1,omitempty"`
	// Number of Hops to reach Destination via Rank2 NHR
	NhrHOP2 int32 `json:"NhrHOP2,omitempty"`
	// Number of Hops to reach Destination via Rank3 NHR
	NhrHOP3 int32 `json:"NhrHOP3,omitempty"`
	// Aging of each Entry
	NhrAge int32 `json:"NhrAge,omitempty"`
}

func (p SlbGslbProxInfoTableParams) iMABean() {}
