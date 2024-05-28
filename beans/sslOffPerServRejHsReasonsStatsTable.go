package beans

import (
	"fmt"
	"reflect"
)

// SslOffPerServRejHsReasonsStatsTable A table for frontend SSL Rejected hand shake reasons current and total statistics per service.
// Note:This mib is not supported for VX instance of Virtualization.
type SslOffPerServRejHsReasonsStatsTable struct {
	// The virtual server name that identifies the service.
	SslOffPerServRejHsReasonsVirtServIndex string
	// Virtual server service index.
	SslOffPerServRejHsReasonsVirtServiceIndex int32
	// Virtual server service rejected handshake reason index.
	SslOffPerServRejHsReasonsVirtRejHsIndex int32
	Params                                  *SslOffPerServRejHsReasonsStatsTableParams
}

func NewSslOffPerServRejHsReasonsStatsTableList() *SslOffPerServRejHsReasonsStatsTable {
	return &SslOffPerServRejHsReasonsStatsTable{}
}

func NewSslOffPerServRejHsReasonsStatsTable(
	sslOffPerServRejHsReasonsVirtServIndex string,
	sslOffPerServRejHsReasonsVirtServiceIndex int32,
	sslOffPerServRejHsReasonsVirtRejHsIndex int32,
	params *SslOffPerServRejHsReasonsStatsTableParams,
) *SslOffPerServRejHsReasonsStatsTable {
	return &SslOffPerServRejHsReasonsStatsTable{
		SslOffPerServRejHsReasonsVirtServIndex:    sslOffPerServRejHsReasonsVirtServIndex,
		SslOffPerServRejHsReasonsVirtServiceIndex: sslOffPerServRejHsReasonsVirtServiceIndex,
		SslOffPerServRejHsReasonsVirtRejHsIndex:   sslOffPerServRejHsReasonsVirtRejHsIndex,
		Params:                                    params,
	}
}

func (c *SslOffPerServRejHsReasonsStatsTable) Name() string {
	return "SslOffPerServRejHsReasonsStatsTable"
}

func (c *SslOffPerServRejHsReasonsStatsTable) GetParams() BeanType {
	return c.Params
}

func (c *SslOffPerServRejHsReasonsStatsTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SslOffPerServRejHsReasonsStatsTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SslOffPerServRejHsReasonsVirtServIndex).IsZero() &&
		reflect.ValueOf(c.SslOffPerServRejHsReasonsVirtServiceIndex).IsZero() &&
		reflect.ValueOf(c.SslOffPerServRejHsReasonsVirtRejHsIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SslOffPerServRejHsReasonsVirtServIndex) + "/" + fmt.Sprint(c.SslOffPerServRejHsReasonsVirtServiceIndex) + "/" + fmt.Sprint(c.SslOffPerServRejHsReasonsVirtRejHsIndex)
}

type SslOffPerServRejHsReasonsStatsTableParams struct {
	// The virtual server name that identifies the service.
	VirtServIndex string `json:"VirtServIndex,omitempty"`
	// Virtual server service index.
	VirtServiceIndex int32 `json:"VirtServiceIndex,omitempty"`
	// Virtual server service rejected handshake reason index.
	VirtRejHsIndex int32 `json:"VirtRejHsIndex,omitempty"`
	// Virtual server service rejected handshake name.
	RejHsName string `json:"RejHsName,omitempty"`
	// Number of New SSL rejected handshakes caused due to this reason between Clients and AAS per second for this virtual service.
	NewHandShake int32 `json:"NewHandShake,omitempty"`
	// Virtual server service port number.
	VirtServPort int32 `json:"VirtServPort,omitempty"`
	// Number of Total SSL rejected handshakes caused due to this reason between Clients and AAS per second for this virtual service.
	NewHandShakeTotal int32 `json:"NewHandShakeTotal,omitempty"`
}

func (p SslOffPerServRejHsReasonsStatsTableParams) iMABean() {}
