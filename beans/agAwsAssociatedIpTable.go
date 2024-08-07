package beans

import (
	"fmt"
	"reflect"
)

// AgAwsAssociatedIpTable Associated ip table of for AWS ha.
// Note: This MIB is only supported on VA platform.
type AgAwsAssociatedIpTable struct {
	// AWS HA associated ip table index.
	// Note: This MIB is only supported on VA platform.
	AgAwsAssociatedIpIndex int32
	Params                 *AgAwsAssociatedIpTableParams
}

func NewAgAwsAssociatedIpTableList() *AgAwsAssociatedIpTable {
	return &AgAwsAssociatedIpTable{}
}

func NewAgAwsAssociatedIpTable(
	agAwsAssociatedIpIndex int32,
	params *AgAwsAssociatedIpTableParams,
) *AgAwsAssociatedIpTable {
	return &AgAwsAssociatedIpTable{
		AgAwsAssociatedIpIndex: agAwsAssociatedIpIndex,
		Params:                 params,
	}
}

func (c *AgAwsAssociatedIpTable) Name() string {
	return "AgAwsAssociatedIpTable"
}

func (c *AgAwsAssociatedIpTable) GetParams() BeanType {
	return c.Params
}

func (c *AgAwsAssociatedIpTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *AgAwsAssociatedIpTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.AgAwsAssociatedIpIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.AgAwsAssociatedIpIndex)
}

type AgAwsAssociatedIpTableParams struct {
	// AWS HA associated ip table index.
	// Note: This MIB is only supported on VA platform.
	AwsAssociatedIpIndex int32 `json:"AwsAssociatedIpIndex,omitempty"`
	// Get AWS Alteon private IP, as a part of HA associated IP set.
	// Note: This MIB is only supported on VA platform.
	AwsAssociatedIpPrivateIp string `json:"AwsAssociatedIpPrivateIp,omitempty"`
	// Get AWS Alteon backup(s) IP, as a part of HA associated IP set.
	// Note: This MIB is only supported on VA platform.
	AwsAssociatedIpPeerIp string `json:"AwsAssociatedIpPeerIp,omitempty"`
	// Get AWS Alteon floating/elastic IP, as a part of HA associated IP set.
	// Note: This MIB is only supported on VA platform.
	AwsAssociatedIpFloatingIp string `json:"AwsAssociatedIpFloatingIp,omitempty"`
}

func (p AgAwsAssociatedIpTableParams) iMABean() {}
