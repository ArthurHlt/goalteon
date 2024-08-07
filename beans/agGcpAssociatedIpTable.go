package beans

import (
	"fmt"
	"reflect"
)

// AgGcpAssociatedIpTable Associated ip table of for GCP ha.
// Note: This MIB is only supported on VA platform.
type AgGcpAssociatedIpTable struct {
	// GCP HA associated ip table index.
	// Note: This MIB is only supported on VA platform.
	AgGcpAssociatedIpIndex int32
	Params                 *AgGcpAssociatedIpTableParams
}

func NewAgGcpAssociatedIpTableList() *AgGcpAssociatedIpTable {
	return &AgGcpAssociatedIpTable{}
}

func NewAgGcpAssociatedIpTable(
	agGcpAssociatedIpIndex int32,
	params *AgGcpAssociatedIpTableParams,
) *AgGcpAssociatedIpTable {
	return &AgGcpAssociatedIpTable{
		AgGcpAssociatedIpIndex: agGcpAssociatedIpIndex,
		Params:                 params,
	}
}

func (c *AgGcpAssociatedIpTable) Name() string {
	return "AgGcpAssociatedIpTable"
}

func (c *AgGcpAssociatedIpTable) GetParams() BeanType {
	return c.Params
}

func (c *AgGcpAssociatedIpTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *AgGcpAssociatedIpTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.AgGcpAssociatedIpIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.AgGcpAssociatedIpIndex)
}

type AgGcpAssociatedIpTableParams struct {
	// GCP HA associated ip table index.
	// Note: This MIB is only supported on VA platform.
	GcpAssociatedIpIndex int32 `json:"GcpAssociatedIpIndex,omitempty"`
	// Get GCP Alteon private IP, as a part of HA associated IP set.
	// Note: This MIB is only supported on VA platform.
	GcpAssociatedIpPrivateIp string `json:"GcpAssociatedIpPrivateIp,omitempty"`
	// Get GCP Alteon backup(s) IP, as a part of HA associated IP set.
	// Note: This MIB is only supported on VA platform.
	GcpAssociatedIpPeerIp string `json:"GcpAssociatedIpPeerIp,omitempty"`
	// Get GCP Alteon floating/elastic IP, as a part of HA associated IP set.
	// Note: This MIB is only supported on VA platform.
	GcpAssociatedIpFloatingIp string `json:"GcpAssociatedIpFloatingIp,omitempty"`
}

func (p AgGcpAssociatedIpTableParams) iMABean() {}
