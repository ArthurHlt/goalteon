package beans

import (
	"fmt"
	"reflect"
)

// SslBeOffPerServRejHsReasonsStatsTable A table for backend SSL Rejected hand shake reasons current and total statistics per service.
// Note:This mib is not supported for VX instance of Virtualization.
type SslBeOffPerServRejHsReasonsStatsTable struct {
	// The virtual server name that identifies the service.
	SslBeOffPerServRejHsReasonsVirtServIndex string
	// Virtual server service index.
	SslBeOffPerServRejHsReasonsVirtServiceIndex int32
	// Virtual server service rejected handshake reason index.
	SslBeOffPerServRejHsReasonsVirtRejHsIndex int32
	Params                                    *SslBeOffPerServRejHsReasonsStatsTableParams
}

func NewSslBeOffPerServRejHsReasonsStatsTableList() *SslBeOffPerServRejHsReasonsStatsTable {
	return &SslBeOffPerServRejHsReasonsStatsTable{}
}

func NewSslBeOffPerServRejHsReasonsStatsTable(
	sslBeOffPerServRejHsReasonsVirtServIndex string,
	sslBeOffPerServRejHsReasonsVirtServiceIndex int32,
	sslBeOffPerServRejHsReasonsVirtRejHsIndex int32,
	params *SslBeOffPerServRejHsReasonsStatsTableParams,
) *SslBeOffPerServRejHsReasonsStatsTable {
	return &SslBeOffPerServRejHsReasonsStatsTable{
		SslBeOffPerServRejHsReasonsVirtServIndex:    sslBeOffPerServRejHsReasonsVirtServIndex,
		SslBeOffPerServRejHsReasonsVirtServiceIndex: sslBeOffPerServRejHsReasonsVirtServiceIndex,
		SslBeOffPerServRejHsReasonsVirtRejHsIndex:   sslBeOffPerServRejHsReasonsVirtRejHsIndex,
		Params: params,
	}
}

func (c *SslBeOffPerServRejHsReasonsStatsTable) Name() string {
	return "SslBeOffPerServRejHsReasonsStatsTable"
}

func (c *SslBeOffPerServRejHsReasonsStatsTable) GetParams() BeanType {
	return c.Params
}

func (c *SslBeOffPerServRejHsReasonsStatsTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SslBeOffPerServRejHsReasonsStatsTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SslBeOffPerServRejHsReasonsVirtServIndex).IsZero() &&
		reflect.ValueOf(c.SslBeOffPerServRejHsReasonsVirtServiceIndex).IsZero() &&
		reflect.ValueOf(c.SslBeOffPerServRejHsReasonsVirtRejHsIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SslBeOffPerServRejHsReasonsVirtServIndex) + "/" + fmt.Sprint(c.SslBeOffPerServRejHsReasonsVirtServiceIndex) + "/" + fmt.Sprint(c.SslBeOffPerServRejHsReasonsVirtRejHsIndex)
}

type SslBeOffPerServRejHsReasonsStatsTableParams struct {
	// The virtual server name that identifies the service.
	VirtServIndex string `json:"VirtServIndex,omitempty"`
	// Virtual server service index.
	VirtServiceIndex int32 `json:"VirtServiceIndex,omitempty"`
	// Virtual server service rejected handshake reason index.
	VirtRejHsIndex int32 `json:"VirtRejHsIndex,omitempty"`
	// Virtual server service rejected handshake name.
	RejHsName string `json:"RejHsName,omitempty"`
	// Number of New SSL rejected handshakes caused due to this reason between AAS and server per second for this virtual service.
	NewHandShake int32 `json:"NewHandShake,omitempty"`
	// Virtual server service port number.
	VirtServPort int32 `json:"VirtServPort,omitempty"`
	// Number of Total SSL rejected handshakes caused due to this reason between AAS and server per second for this virtual service.
	NewHandShakeTotal int32 `json:"NewHandShakeTotal,omitempty"`
}

func (p SslBeOffPerServRejHsReasonsStatsTableParams) iMABean() {}
