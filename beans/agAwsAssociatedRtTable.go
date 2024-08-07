package beans

import (
	"fmt"
	"reflect"
)

// AgAwsAssociatedRtTable Associated route table of for AWS ha.
// Note: This MIB is only supported on VA platform.
type AgAwsAssociatedRtTable struct {
	// AWS HA associated route table index.
	// Note: This MIB is only supported on VA platform.
	AgAwsAssociatedRtIndex int32
	Params                 *AgAwsAssociatedRtTableParams
}

func NewAgAwsAssociatedRtTableList() *AgAwsAssociatedRtTable {
	return &AgAwsAssociatedRtTable{}
}

func NewAgAwsAssociatedRtTable(
	agAwsAssociatedRtIndex int32,
	params *AgAwsAssociatedRtTableParams,
) *AgAwsAssociatedRtTable {
	return &AgAwsAssociatedRtTable{
		AgAwsAssociatedRtIndex: agAwsAssociatedRtIndex,
		Params:                 params,
	}
}

func (c *AgAwsAssociatedRtTable) Name() string {
	return "AgAwsAssociatedRtTable"
}

func (c *AgAwsAssociatedRtTable) GetParams() BeanType {
	return c.Params
}

func (c *AgAwsAssociatedRtTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *AgAwsAssociatedRtTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.AgAwsAssociatedRtIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.AgAwsAssociatedRtIndex)
}

type AgAwsAssociatedRtTableParams struct {
	// AWS HA associated route table index.
	// Note: This MIB is only supported on VA platform.
	AwsAssociatedRtIndex int32 `json:"AwsAssociatedRtIndex,omitempty"`
	// Get AWS Alteon Route ID, as a part of HA associated route set.
	// Note: This MIB is only supported on VA platform.
	AwsAssociatedRtRouteId string `json:"AwsAssociatedRtRouteId,omitempty"`
	// Get AWS Alteon eni ID, as a part of HA associated route set.
	// Note: This MIB is only supported on VA platform.
	AwsAssociatedRtEniId string `json:"AwsAssociatedRtEniId,omitempty"`
	// Get AWS Alteon backup(s) eni ID, as a part of HA associated route set.
	// Note: This MIB is only supported on VA platform.
	AwsAssociatedRtPeerEniId string `json:"AwsAssociatedRtPeerEniId,omitempty"`
}

func (p AgAwsAssociatedRtTableParams) iMABean() {}
